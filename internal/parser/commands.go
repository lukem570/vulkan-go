package parser

import (
	"fmt"
	"strings"

	"github.com/lukem570/vulkan-go/internal/generator"
)

var receiverPrefixes = map[string][]string{
	"CommandBuffer": {"Cmd"},
}

type XMLCommands struct {
	Commands []XMLCommand `xml:"command"`
}

type XMLCommand struct {
	Proto  XMLProto   `xml:"proto"`
	Params []XMLParam `xml:"param"`
	Alias  string     `xml:"alias,attr"`
}

type XMLProto struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	InnerXML string `xml:",innerxml"`
}

type XMLParam struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	Len      string `xml:"len,attr"`
	Optional string `xml:"optional,attr"`
	InnerXML string `xml:",innerxml"`
}

func (x *XMLRegistry) parseCommands(r *generator.Registry) {
	for _, cmd := range x.Commands.Commands {
		if cmd.Alias != "" {
			continue
		}
		c := buildCommand(cmd, r.Handles, r.FuncPointers, r.Structs)
		if c != nil {
			r.Commands[c.CName] = c
		}
	}
}

// extractCType reconstructs the full C type string from an XML param's InnerXML.
// e.g. "const <type>VkAllocationCallbacks</type>* <name>pAllocator</name>" -> "const VkAllocationCallbacks*"
// Array params like "const <type>float</type> <name>blendConstants</name>[4]" -> "const float*"
func extractCType(innerXML string) string {
	// Check for array suffix after </name> — treat as pointer
	hasArray := false
	if idx := strings.Index(innerXML, "</name>"); idx >= 0 {
		after := innerXML[idx+len("</name>"):]
		if strings.Contains(after, "[") {
			hasArray = true
		}
	}

	// Remove <name>...</name> and everything after
	if idx := strings.Index(innerXML, "<name>"); idx >= 0 {
		innerXML = innerXML[:idx]
	}
	// Strip XML tags, keeping text content
	var result strings.Builder
	inTag := false
	for _, ch := range innerXML {
		if ch == '<' {
			inTag = true
			continue
		}
		if ch == '>' {
			inTag = false
			continue
		}
		if !inTag {
			result.WriteRune(ch)
		}
	}
	// Clean up whitespace
	s := strings.TrimSpace(result.String())
	// Normalize "const VkFoo *" -> "const VkFoo*"
	s = strings.ReplaceAll(s, " *", "*")
	s = strings.ReplaceAll(s, "  ", " ")

	// Array params decay to pointers in C function signatures
	if hasArray && !strings.HasSuffix(s, "*") {
		s += "*"
	}
	return s
}

func buildCommand(
	cmd XMLCommand,
	handles map[string]*generator.GoHandle,
	funcPointers map[string]*generator.GoFuncPointer,
	structs map[string]*generator.Structured,
) *generator.GoCommand {

	rawName := cmd.Proto.Name
	returnTypeName := cmd.Proto.Type

	c := &generator.GoCommand{
		CName:       rawName,
		Name:        methodName(rawName, ""),
		CReturnType: returnTypeName,
	}

	seen := map[string]bool{}

	// Build CParams from all XML params
	for _, p := range cmd.Params {
		if seen[p.Name] {
			continue
		}
		seen[p.Name] = true
		cType := extractCType(p.InnerXML)
		// For double-pointer non-struct/non-handle array params (e.g. const uint32_t* const*),
		// use const void* in the C wrapper so that Go's unsafe.Pointer is accepted by CGo.
		// Struct double-pointer params (ppGeometries, ppBuildRangeInfos) use PtrSlice instead
		// and keep their correct C types.
		{
			cutoff := strings.Index(p.InnerXML, "<name>")
			if cutoff < 0 {
				cutoff = len(p.InnerXML)
			}
			if strings.Count(p.InnerXML[:cutoff], "*") >= 2 && p.Len != "" {
				goName := stripVk(p.Type)
				_, isStruct := structs[goName]
				_, isHandle := handles[goName]
				if !isStruct && !isHandle && primitiveType(p.Type) != nil {
					cType = "const void*"
				}
			}
		}
		c.CParams = append(c.CParams, generator.CParam{
			Type: cType,
			Name: p.Name,
		})
	}

	if c.CReturnType == "" {
		c.CReturnType = "void"
	}

	// Detect VkResult return → HasError, no Go return type
	switch returnTypeName {
	case "VkResult":
		c.HasError = true
	case "void", "":
		// nothing
	default:
		c.ReturnType = resolveFieldType(returnTypeName, false, false, handles, funcPointers, structs)
	}

	params := cmd.Params
	if len(params) == 0 {
		return c
	}

	// ---- Handle receiver detection ----------------------------------------
	// If the first parameter is a Vulkan handle, make it the method receiver.
	firstParam := params[0]
	firstGoName := stripVk(firstParam.Type)
	if _, ok := handles[firstGoName]; ok {
		c.ReceiverType = firstGoName
		params = params[1:]
		c.Name = methodName(rawName, firstGoName)
	}

	// ---- Byte-data pattern detection ----------------------------------------
	// Detect: second-to-last is size_t* pDataSize, last is void* pData with len="pDataSize".
	byteDataDetected := false
	if len(params) >= 2 && (returnTypeName == "VkResult" || returnTypeName == "void" || returnTypeName == "") {
		secondLast := params[len(params)-2]
		last := params[len(params)-1]

		isSizeParam := secondLast.Type == "size_t" &&
			strings.Contains(secondLast.InnerXML, "*") &&
			!strings.Contains(secondLast.InnerXML, "const")

		isDataParam := last.Type == "void" &&
			last.Len != "" &&
			strings.Contains(last.Len, secondLast.Name) &&
			strings.Contains(last.InnerXML, "*") &&
			!strings.Contains(last.InnerXML, "const")

		if isSizeParam && isDataParam {
			c.ByteDataPattern = &generator.ByteDataInfo{
				SizeCParam: secondLast.Name,
				DataCParam: last.Name,
			}
			params = params[:len(params)-2]
			byteDataDetected = true
		}
	}

	// ---- Enumerate pattern detection ----------------------------------------
	// Detect the two-call pattern: second-to-last is uint32_t* pFooCount,
	// last is T* pFoos with len="pFooCount".
	enumerateDetected := false
	if !byteDataDetected && len(params) >= 2 && (returnTypeName == "VkResult" || returnTypeName == "void" || returnTypeName == "") {
		secondLast := params[len(params)-2]
		last := params[len(params)-1]

		isCountParam := secondLast.Type == "uint32_t" &&
			strings.Contains(secondLast.InnerXML, "*") &&
			!strings.Contains(secondLast.InnerXML, "const")

		isArrayParam := last.Len != "" &&
			strings.Contains(last.Len, secondLast.Name) &&
			strings.Contains(last.InnerXML, "*") &&
			!strings.Contains(last.InnerXML, "const")

		if isCountParam && isArrayParam {
			elemType := resolveFieldType(last.Type, false, false, handles, funcPointers, structs)
			// Only use enumerate pattern for types that have working GenerateFromC
			// (handles, primitives). Structs need a fromC() that doesn't exist yet.
			switch elemType.(type) {
			case *generator.Handle, *generator.Primitive, *generator.NamedType, *generator.Bool, *generator.StructType:
				c.EnumeratePattern = &generator.EnumerateInfo{
					CountCParam: secondLast.Name,
					ArrayCParam: last.Name,
					ElementType: elemType,
				}
				params = params[:len(params)-2]
				enumerateDetected = true
			}
		}
	}

	// ---- Output parameter detection ----------------------------------------
	// The LAST parameter of many Vulkan commands is a pointer to an output value.
	// Heuristics:
	//   - return type is VkResult or void
	//   - last param is a non-const pointer
	//   - param name starts with 'p' followed by uppercase
	//   - optional attribute is NOT set (or is "false")
	if !enumerateDetected && !byteDataDetected && len(params) > 0 {
		last := params[len(params)-1]
		if isOutputParam(last, returnTypeName, handles, structs) {
			ft := outParamType(last, handles, funcPointers, structs)
			op := generator.OutParam{
				Name: last.Name,
				Type: ft,
			}
			// Detect array output params (e.g. pPipelines with len="createInfoCount")
			if last.Len != "" && !strings.Contains(last.Len, "null-terminated") {
				op.IsArray = true
				op.CountGoParam = resolveCountExpr(last.Len, params, structs)
			}
			c.OutParams = append(c.OutParams, op)
			params = params[:len(params)-1]
		}
	}

	seen = map[string]bool{}

	// ---- Regular input params -----------------------------------------------
	for _, p := range params {
		if seen[p.Name] {
			continue
		}
		seen[p.Name] = true

		// Only inspect the part of InnerXML before <name> so that comments
		// (e.g. "Vk*PipelineCreateInfo") don't give false positives.
		dblPtrCutoff := strings.Index(p.InnerXML, "<name>")
		if dblPtrCutoff < 0 {
			dblPtrCutoff = len(p.InnerXML)
		}
		beforeName := p.InnerXML[:dblPtrCutoff]
		isPtr := strings.Contains(beforeName, "*")
		isArr := p.Len != "" && p.Len != "null-terminated"
		isDblPtr := strings.Count(beforeName, "*") >= 2

		ft := resolveFieldType(p.Type, isPtr, isArr, handles, funcPointers, structs, isDblPtr)

		// Check for fixed-size array (e.g. blendConstants[4])
		if fixedSize := fixedArraySizeParam(p.InnerXML); fixedSize > 0 {
			ft = &generator.FixedArray{Child: ft, Size: fixedSize}
		}

		param := &generator.CommandParam{
			Name: goParamName(p.Name),
			Type: ft,
		}

		if _, ok := ft.(*generator.Slice); ok &&
			strings.Contains(strings.ToLower(c.Params[len(c.Params)-1].Name), "count") {

			c.Params[len(c.Params)-1].CountOf = param
		}

		c.Params = append(c.Params, param)
	}

	return c
}

// methodName strips vk prefix and the handle type prefix, returning a clean Go name.
// vkDestroyInstance → Destroy  (when receiver is Instance)
// vkCreateInstance  → CreateInstance  (no receiver)
func methodName(cName string, handle string) string {
	name := strings.TrimPrefix(cName, "vk")
	name = strings.ReplaceAll(name, handle, "")

	if prefixes, ok := receiverPrefixes[handle]; ok {
		for _, prefix := range prefixes {
			name = strings.TrimPrefix(name, prefix)
		}
	}

	return name
}

func isOutputParam(p XMLParam, retType string, handles map[string]*generator.GoHandle, structs map[string]*generator.Structured) bool {
	if p.Name == "" {
		return false
	}
	// must be a pointer
	if !strings.Contains(p.InnerXML, "*") {
		return false
	}
	// const pointers are inputs, not outputs
	if strings.Contains(p.InnerXML, "const") {
		return false
	}
	// must be result-returning or void function
	if retType != "VkResult" && retType != "void" && retType != "" {
		return false
	}
	// void** params are output pointers (e.g. vkMapMemory's ppData).
	// Check this before the p-prefix convention since ppData has double-p.
	// Single void* params are caller-provided buffers, not output values.
	if p.Type == "void" {
		return strings.Contains(p.InnerXML, "**")
	}
	// must have p-prefix convention
	if len(p.Name) < 2 || p.Name[0] != 'p' || p.Name[1] < 'A' || p.Name[1] > 'Z' {
		return false
	}
	goName := stripVk(p.Type)
	if _, isHandle := handles[goName]; isHandle {
		return true
	}
	// Primitives and simple types
	if primitiveType(p.Type) != nil {
		return true
	}
	// Structs as output params (for getter-style commands)
	if structs != nil {
		if _, isStruct := structs[goName]; isStruct {
			return true
		}
	}
	return false
}

func outParamType(p XMLParam, handles map[string]*generator.GoHandle, funcPointers map[string]*generator.GoFuncPointer, structs map[string]*generator.Structured) generator.FieldType {
	// The output param is T* in C — the Go type is T
	// Special case: void* output params are unsafe.Pointer
	if p.Type == "void" {
		return &generator.VoidPtr{}
	}
	return resolveFieldType(p.Type, false, false, handles, funcPointers, structs)
}

// fixedArraySizeParam returns the fixed array size from a command param's InnerXML
// (e.g. "const <type>float</type> <name>blendConstants</name>[4]" → 4), or 0.
func fixedArraySizeParam(innerXML string) int {
	idx := strings.Index(innerXML, "</name>")
	if idx < 0 {
		return 0
	}
	after := innerXML[idx+len("</name>"):]
	// Look for [N] pattern
	start := strings.Index(after, "[")
	end := strings.Index(after, "]")
	if start < 0 || end < 0 || end <= start+1 {
		return 0
	}
	n := 0
	for _, ch := range after[start+1 : end] {
		if ch < '0' || ch > '9' {
			return 0
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

// resolveCountExpr converts a C len expression to a Go expression.
// For simple params like "createInfoCount" it returns the Go param name.
// For struct member access like "pAllocateInfo->descriptorSetCount", it checks
// if the field was collapsed into a slice (CountFor != "") and returns
// len(param.SliceField) if so, otherwise param.Field.
func resolveCountExpr(lenExpr string, params []XMLParam, structs map[string]*generator.Structured) string {
	parts := strings.SplitN(lenExpr, "->", 2)
	if len(parts) != 2 {
		return goParamName(lenExpr)
	}

	paramCName := parts[0]
	fieldCName := parts[1]
	paramGoName := goParamName(paramCName)

	// Find the struct type of the parameter
	var structTypeName string
	for _, p := range params {
		if p.Name == paramCName {
			structTypeName = stripVk(p.Type)
			break
		}
	}

	// Check if the field was collapsed (it's a count for a slice)
	if structTypeName != "" && structs != nil {
		if s, ok := structs[structTypeName]; ok {
			for _, f := range s.Fields {
				if f.CName == fieldCName && f.CountFor != "" {
					// Find the Go name of the slice field it counts
					for _, sf := range s.Fields {
						if sf.GoName == f.CountFor {
							return fmt.Sprintf("uint32(len(%s.%s))", paramGoName, sf.GoName)
						}
					}
				}
			}
		}
	}

	return paramGoName + "." + toPublic(fieldCName)
}

func goParamName(cName string) string {
	name := stripP(cName)
	// avoid Go keywords
	switch name {
	case "type":
		return "typ"
	case "range":
		return "rng"
	case "func":
		return "fn"
	}
	return name
}
