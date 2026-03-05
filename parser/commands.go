package parser

import (
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

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

func parseCommands(x *XMLRegistry, r *generator.Registry) {
	for _, cmd := range x.Commands.Commands {
		if cmd.Alias != "" {
			continue
		}
		c := buildCommand(cmd, r.Handles, r.FuncPointers, r.Structs)
		if c != nil {
			r.Commands[c.Name] = c
		}
	}
}

// extractCType reconstructs the full C type string from an XML param's InnerXML.
// e.g. "const <type>VkAllocationCallbacks</type>* <name>pAllocator</name>" -> "const VkAllocationCallbacks*"
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

func buildCommand(cmd XMLCommand, handles map[string]*generator.GoHandle, funcPointers map[string]*generator.GoFuncPointer, structs map[string]*generator.Structured) *generator.GoCommand {
	rawName := cmd.Proto.Name
	returnTypeName := cmd.Proto.Type

	c := &generator.GoCommand{
		CName:       rawName,
		Name:        methodName(rawName),
		CReturnType: returnTypeName,
	}

	// Build CParams from all XML params
	for _, p := range cmd.Params {
		c.CParams = append(c.CParams, generator.CParam{
			Type: extractCType(p.InnerXML),
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
	}

	// ---- Output parameter detection ----------------------------------------
	// The LAST parameter of many Vulkan commands is a pointer to an output value.
	// Heuristics:
	//   - return type is VkResult or void
	//   - last param is a non-const pointer
	//   - param name starts with 'p' followed by uppercase
	//   - optional attribute is NOT set (or is "false")
	if len(params) > 0 {
		last := params[len(params)-1]
		if isOutputParam(last, returnTypeName) {
			ft := outParamType(last, handles, funcPointers, structs)
			c.OutParams = append(c.OutParams, generator.OutParam{
				Name: last.Name,
				Type: ft,
			})
			params = params[:len(params)-1]
		}
	}

	// ---- Regular input params -----------------------------------------------
	for _, p := range params {
		isPtr := strings.Contains(p.InnerXML, "*")
		isArr := p.Len != "" && p.Len != "null-terminated"

		ft := resolveFieldType(p.Type, isPtr, isArr, handles, funcPointers, structs)

		param := generator.CommandParam{
			Name: goParamName(p.Name),
			Type: ft,
		}
		c.Params = append(c.Params, param)
	}

	return c
}

// methodName strips vk prefix and the handle type prefix, returning a clean Go name.
// vkDestroyInstance → Destroy  (when receiver is Instance)
// vkCreateInstance  → CreateInstance  (no receiver)
func methodName(cName string) string {
	name := strings.TrimPrefix(cName, "vk")
	return name
}

func isOutputParam(p XMLParam, retType string) bool {
	if p.Name == "" {
		return false
	}
	// must be a pointer
	if !strings.Contains(p.InnerXML, "*") {
		return false
	}
	// must have p-prefix convention
	if len(p.Name) < 2 || p.Name[0] != 'p' || p.Name[1] < 'A' || p.Name[1] > 'Z' {
		return false
	}
	// const pointers are inputs, not outputs
	if strings.Contains(p.InnerXML, "const") {
		return false
	}
	// must be result-returning or void function
	return retType == "VkResult" || retType == "void" || retType == ""
}

func outParamType(p XMLParam, handles map[string]*generator.GoHandle, funcPointers map[string]*generator.GoFuncPointer, structs map[string]*generator.Structured) generator.FieldType {
	// The output param is T* in C — the Go type is T
	return resolveFieldType(p.Type, false, false, handles, funcPointers, structs)
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