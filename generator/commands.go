package generator

import (
	"fmt"
	"strings"
)

type GoCommand struct {
	Name         string
	CName        string // original C name e.g. vkCreateInstance
	ReceiverType string // e.g. "Instance" if first param was a handle
	HasError     bool
	ReturnType   FieldType  // non-VkResult return, if any
	OutParams    []OutParam // params that are actually return values
	Params       []*CommandParam
	CReturnType  string   // C return type e.g. "VkResult", "void"
	CParams      []CParam // full C prototype params for wrapper generation
	Platform     string   // non-empty for platform-specific commands (e.g. "xlib")

	// EnumeratePattern, when non-nil, indicates the two-call Vulkan
	// enumerate pattern (call once for count, then again with array).
	EnumeratePattern *EnumerateInfo

	// ByteDataPattern, when non-nil, indicates the two-call Vulkan
	// byte-data pattern (call once for size, then again with void* buffer).
	ByteDataPattern *ByteDataInfo

	// CallbackStructParamName, when non-empty, names the input param whose
	// callbackCleanupFn should be attached to the returned handle after a
	// successful create call.
	CallbackStructParamName string
}

// EnumerateInfo describes a Vulkan two-call enumerate pattern.
type EnumerateInfo struct {
	CountCParam string    // C param name for the count (e.g. "pPhysicalDeviceCount")
	ArrayCParam string    // C param name for the array (e.g. "pPhysicalDevices")
	ElementType FieldType // Go element type (e.g. Handle for PhysicalDevice)
}

// ByteDataInfo describes a Vulkan two-call byte-data pattern:
// first call to get pDataSize, second to fill void* pData.
type ByteDataInfo struct {
	SizeCParam string // C param name for the size (e.g. "pDataSize")
	DataCParam string // C param name for the data (e.g. "pData")
}

// CParam holds the full C type and name for a command parameter.
type CParam struct {
	Type string // e.g. "const VkInstanceCreateInfo*"
	Name string // e.g. "pCreateInfo"
}

// OutParam describes a parameter that is an output pointer in C but a return
// value in Go.
type OutParam struct {
	Name         string
	Type         FieldType
	IsArray      bool   // true when the output is an array (e.g. pPipelines with len=createInfoCount)
	CountGoParam string // Go param name that holds the count (e.g. "createInfoCount")
}

type CommandParam struct {
	Name    string
	Type    FieldType
	CountOf *CommandParam
}

func (c *GoCommand) GenerateWrapper(sTypes map[string]string) string {
	if c == nil {
		return ""
	}
	if c.EnumeratePattern != nil {
		return c.generateEnumerateWrapper()
	}
	if c.ByteDataPattern != nil {
		return c.generateByteDataWrapper()
	}

	var b strings.Builder
	g := &CodeGen{}

	// ---- Function signature ------------------------------------------------
	if c.ReceiverType != "" {
		b.WriteString(fmt.Sprintf("func (h %s) %s(\n", c.ReceiverType, c.Name))
	} else {
		b.WriteString(fmt.Sprintf("func %s(\n", c.Name))
	}

	for _, p := range c.Params {
		if p.CountOf != nil {
			continue
		}

		b.WriteString(fmt.Sprintf("\t%s %s,\n", p.Name, p.Type.GoName()))
	}
	b.WriteString(")")

	// ---- Return types -------------------------------------------------------
	var returns []string
	for _, op := range c.OutParams {
		if op.IsArray {
			returns = append(returns, "[]"+op.Type.GoName())
		} else {
			returns = append(returns, op.Type.GoName())
		}
	}
	if c.ReturnType != nil {
		returns = append(returns, c.ReturnType.GoName())
	}
	if c.HasError {
		returns = append(returns, "error")
	}
	if len(returns) == 1 {
		b.WriteString(" " + returns[0])
	} else if len(returns) > 1 {
		b.WriteString(" (" + strings.Join(returns, ", ") + ")")
	}
	b.WriteString(" {\n")

	// ---- Body ---------------------------------------------------------------
	b.WriteString("\tcancels := make([]func(), 0)\n")
	b.WriteString("\tdefer func() { for _, c := range cancels { c() } }()\n\n")

	// Convert Go → C for each input param
	cArgs := make([]string, 0)

	if c.ReceiverType != "" {
		cArgs = append(cArgs, fmt.Sprintf("C.%s(unsafe.Pointer(h.handle))", "Vk"+c.ReceiverType))
	}

	for _, p := range c.Params {
		if p.CountOf != nil {
			g.builder.WriteString(p.Name+" := len("+p.CountOf.Name+")\n")
			cArgs = append(cArgs, "C.uint32_t("+p.Name+")")
			continue
		}
		
		cVar := g.Var("c_" + sanitizeIdent(p.Name))
		b.WriteString(fmt.Sprintf("\t// param %s\n", p.Name))
		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		_ = result
		_ = cVar
		// Fixed arrays need to be passed as pointer to first element for C
		if _, ok := p.Type.(*FixedArray); ok {
			cArgs = append(cArgs, "&"+result+"[0]")
		} else {
			cArgs = append(cArgs, result)
		}
	}

	// Allocate output param slots
	for _, op := range c.OutParams {
		varName := sanitizeIdent(op.Name) + "Out"
		if op.IsArray {
			// Allocate a C array sized by the count param
			b.WriteString(fmt.Sprintf("\t%s := (*%s)(C.malloc(C.size_t(%s) * C.size_t(unsafe.Sizeof(*new(%s)))))\n",
				varName, op.Type.CName(), op.CountGoParam, op.Type.CName()))
			b.WriteString(fmt.Sprintf("\tdefer C.free(unsafe.Pointer(%s))\n", varName))
			cArgs = append(cArgs, varName)
		} else {
			b.WriteString(fmt.Sprintf("\tvar %s %s\n", varName, op.Type.CName()))

			if s, ok := op.Type.(*StructType); ok {
				if stype, ok := sTypes[s.Name]; ok {
					b.WriteString(fmt.Sprintf(
						"\t%s.sType = (C.VkStructureType)(%s)\n", 
						varName, stype,
					))
				}
			}
			cArgs = append(cArgs, "&"+varName)
		}
	}

	// For Destroy/Free commands: run handle cleanup before the C call
	isDestroy := strings.HasPrefix(c.Name, "Destroy") || strings.HasPrefix(c.Name, "Free")
	if isDestroy {
		for _, p := range c.Params {
			if _, ok := p.Type.(*Handle); ok {
				b.WriteString(fmt.Sprintf("\tif %s != nil && %s.cleanup != nil { %s.cleanup() }\n", p.Name, p.Name, p.Name))
			}
		}
	}

	// Call C function
	var callLine string
	if c.HasError || c.ReturnType != nil {
		callLine = fmt.Sprintf("\t_result := C.fn_%s(%s)\n", c.CName, strings.Join(cArgs, ", "))
	} else {
		callLine = fmt.Sprintf("\tC.fn_%s(%s)\n", c.CName, strings.Join(cArgs, ", "))
	}
	b.WriteString(callLine)

	// Error check
	if c.HasError {
		b.WriteString("\tif _result != C.VK_SUCCESS {\n")
		b.WriteString("\t\treturn ")
		for _, op := range c.OutParams {
			if op.IsArray {
				b.WriteString("nil, ")
			} else {
				b.WriteString(zeroValue(op.Type) + ", ")
			}
		}
		if c.ReturnType != nil {
			b.WriteString(zeroValue(c.ReturnType) + ", ")
		}
		b.WriteString("vkError(_result)\n\t}\n")
	}

	// Convert out params from C → Go
	var retVars []string
	var handleRetVar string // tracks the var name of a Handle out param, for callback cleanup attachment
	for _, op := range c.OutParams {
		varName := sanitizeIdent(op.Name) + "Out"

		if op.IsArray {
			// Convert C array to Go slice
			sliceVar := g.Var("out")
			b.WriteString(fmt.Sprintf("\t%s := make([]%s, %s)\n", sliceVar, op.Type.GoName(), op.CountGoParam))
			idxVar := g.Var("i")
			b.WriteString(fmt.Sprintf("\tfor %s := range %s {\n", idxVar, sliceVar))
			elemExpr := fmt.Sprintf("(*[1<<30]%s)(unsafe.Pointer(%s))[%s]", op.Type.CName(), varName, idxVar)
			outG := &CodeGen{varIndex: g.varIndex}
			goVal := op.Type.GenerateFromC(outG, elemExpr)
			b.WriteString(outG.String())
			g.varIndex = outG.varIndex
			b.WriteString(fmt.Sprintf("\t\t%s[%s] = %s\n", sliceVar, idxVar, goVal))
			b.WriteString("\t}\n")
			retVars = append(retVars, sliceVar)
		} else {
			outG := &CodeGen{varIndex: g.varIndex}
			goVal := op.Type.GenerateFromC(outG, varName)
			b.WriteString(outG.String())
			g.varIndex = outG.varIndex
			retVars = append(retVars, goVal)
			if _, ok := op.Type.(*Handle); ok {
				handleRetVar = goVal
			}
		}
	}

	// For Create commands: attach callback cleanup to the returned handle
	if c.CallbackStructParamName != "" && handleRetVar != "" {
		b.WriteString(fmt.Sprintf("\tif %s.callbackCleanupFn != nil { %s.cleanup = %s.callbackCleanupFn }\n",
			c.CallbackStructParamName, handleRetVar, c.CallbackStructParamName))
	}

	// Convert non-VkResult return value
	if c.ReturnType != nil {
		outG := &CodeGen{varIndex: g.varIndex}
		goVal := c.ReturnType.GenerateFromC(outG, "_result")
		b.WriteString(outG.String())
		g.varIndex = outG.varIndex
		retVars = append(retVars, goVal)
	}

	if c.HasError {
		retVars = append(retVars, "nil")
	}
	if len(retVars) > 0 {
		b.WriteString("\treturn " + strings.Join(retVars, ", ") + "\n")
	}

	b.WriteString("}\n\n")
	return b.String()
}

// generateEnumerateWrapper generates a Go wrapper for the Vulkan two-call
// enumerate pattern. It calls the C function twice: once to get the count,
// then again to fill an allocated array, returning a Go slice.
func (c *GoCommand) generateEnumerateWrapper() string {
	ep := c.EnumeratePattern
	elemGoName := ep.ElementType.GoName()
	elemCName := ep.ElementType.CName()
	sliceType := "[]" + elemGoName

	var b strings.Builder

	// Signature
	if c.ReceiverType != "" {
		b.WriteString(fmt.Sprintf("func (h %s) %s(\n", c.ReceiverType, c.Name))
	} else {
		b.WriteString(fmt.Sprintf("func %s(\n", c.Name))
	}
	for _, p := range c.Params {
		if p.CountOf != nil {
			continue
		}

		b.WriteString(fmt.Sprintf("\t%s %s,\n", p.Name, p.Type.GoName()))
	}
	b.WriteString(")")

	// Return type
	if c.HasError {
		b.WriteString(fmt.Sprintf(" (%s, error)", sliceType))
	} else {
		b.WriteString(" " + sliceType)
	}
	b.WriteString(" {\n")

	// Body
	b.WriteString("\tcancels := make([]func(), 0)\n")
	b.WriteString("\tdefer func() { for _, c := range cancels { c() } }()\n\n")

	// Build the C args for the receiver + input params
	g := &CodeGen{}
	var cArgs []string
	if c.ReceiverType != "" {
		cArgs = append(cArgs, fmt.Sprintf("C.%s(unsafe.Pointer(h.handle))", "Vk"+c.ReceiverType))
	}
	for _, p := range c.Params {
		if p.CountOf != nil {
			g.builder.WriteString(p.Name+" := len("+p.CountOf.Name+")\n")
			cArgs = append(cArgs, "C.uint32_t("+p.Name+")")
			continue
		}

		b.WriteString(fmt.Sprintf("\t// param %s\n", p.Name))
		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		if _, ok := p.Type.(*FixedArray); ok {
			cArgs = append(cArgs, "&"+result+"[0]")
		} else {
			cArgs = append(cArgs, result)
		}
	}

	// First call: get count
	b.WriteString("\tvar count C.uint32_t\n")
	firstCallArgs := append(append([]string{}, cArgs...), "&count", "nil")
	if c.HasError {
		b.WriteString(fmt.Sprintf("\t_result := C.fn_%s(%s)\n", c.CName, strings.Join(firstCallArgs, ", ")))
		b.WriteString("\tif _result != C.VK_SUCCESS && _result != C.VK_INCOMPLETE {\n")
		b.WriteString("\t\treturn nil, vkError(_result)\n")
		b.WriteString("\t}\n")
	} else {
		b.WriteString(fmt.Sprintf("\tC.fn_%s(%s)\n", c.CName, strings.Join(firstCallArgs, ", ")))
	}
	b.WriteString("\tif count == 0 {\n")
	if c.HasError {
		b.WriteString("\t\treturn nil, nil\n")
	} else {
		b.WriteString("\t\treturn nil\n")
	}
	b.WriteString("\t}\n\n")

	// Allocate C array
	b.WriteString(fmt.Sprintf("\tcArr := (*%s)(C.malloc(C.size_t(count) * C.size_t(unsafe.Sizeof(*new(%s)))))\n", elemCName, elemCName))
	b.WriteString("\tcancels = append(cancels, func() { C.free(unsafe.Pointer(cArr)) })\n\n")

	// Second call: fill array
	secondCallArgs := append(append([]string{}, cArgs...), "&count", "cArr")
	if c.HasError {
		b.WriteString(fmt.Sprintf("\t_result = C.fn_%s(%s)\n", c.CName, strings.Join(secondCallArgs, ", ")))
		b.WriteString("\tif _result != C.VK_SUCCESS {\n")
		b.WriteString("\t\treturn nil, vkError(_result)\n")
		b.WriteString("\t}\n\n")
	} else {
		b.WriteString(fmt.Sprintf("\tC.fn_%s(%s)\n", c.CName, strings.Join(secondCallArgs, ", ")))
	}

	// Convert C array to Go slice
	b.WriteString(fmt.Sprintf("\tout := make(%s, int(count))\n", sliceType))
	b.WriteString(fmt.Sprintf("\tcSlice := (*[1 << 30]%s)(unsafe.Pointer(cArr))[:count:count]\n", elemCName))
	b.WriteString("\tfor i, v := range cSlice {\n")

	convG := &CodeGen{varIndex: g.varIndex}
	goVal := ep.ElementType.GenerateFromC(convG, "v")
	b.WriteString(convG.String())
	b.WriteString(fmt.Sprintf("\t\tout[i] = %s\n", goVal))
	b.WriteString("\t}\n")

	if c.HasError {
		b.WriteString("\treturn out, nil\n")
	} else {
		b.WriteString("\treturn out\n")
	}
	b.WriteString("}\n\n")
	return b.String()
}

// generateByteDataWrapper generates a Go wrapper for the Vulkan two-call byte-data
// pattern: first call with nil data to get pDataSize, then allocate and fill.
func (c *GoCommand) generateByteDataWrapper() string {
	var b strings.Builder

	// Signature
	if c.ReceiverType != "" {
		b.WriteString(fmt.Sprintf("func (h %s) %s(\n", c.ReceiverType, c.Name))
	} else {
		b.WriteString(fmt.Sprintf("func %s(\n", c.Name))
	}
	for _, p := range c.Params {
		if p.CountOf != nil {
			continue
		}
		b.WriteString(fmt.Sprintf("\t%s %s,\n", p.Name, p.Type.GoName()))
	}
	b.WriteString(")")

	// Return type
	if c.HasError {
		b.WriteString(" ([]byte, error)")
	} else {
		b.WriteString(" []byte")
	}
	b.WriteString(" {\n")

	// Body
	b.WriteString("\tcancels := make([]func(), 0)\n")
	b.WriteString("\tdefer func() { for _, c := range cancels { c() } }()\n\n")

	// Build C args for receiver + input params
	g := &CodeGen{}
	var cArgs []string
	if c.ReceiverType != "" {
		cArgs = append(cArgs, fmt.Sprintf("C.%s(unsafe.Pointer(h.handle))", "Vk"+c.ReceiverType))
	}
	for _, p := range c.Params {
		if p.CountOf != nil {
			g.builder.WriteString(p.Name + " := len(" + p.CountOf.Name + ")\n")
			cArgs = append(cArgs, "C.uint32_t("+p.Name+")")
			continue
		}
		b.WriteString(fmt.Sprintf("\t// param %s\n", p.Name))
		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		if _, ok := p.Type.(*FixedArray); ok {
			cArgs = append(cArgs, "&"+result+"[0]")
		} else {
			cArgs = append(cArgs, result)
		}
	}

	// First call: get data size
	b.WriteString("\tvar dataSize C.size_t\n")
	firstCallArgs := append(append([]string{}, cArgs...), "&dataSize", "nil")
	if c.HasError {
		b.WriteString(fmt.Sprintf("\t_result := C.fn_%s(%s)\n", c.CName, strings.Join(firstCallArgs, ", ")))
		b.WriteString("\tif _result != C.VK_SUCCESS && _result != C.VK_INCOMPLETE {\n")
		b.WriteString("\t\treturn nil, vkError(_result)\n")
		b.WriteString("\t}\n")
	} else {
		b.WriteString(fmt.Sprintf("\tC.fn_%s(%s)\n", c.CName, strings.Join(firstCallArgs, ", ")))
	}
	b.WriteString("\tif dataSize == 0 {\n")
	if c.HasError {
		b.WriteString("\t\treturn nil, nil\n")
	} else {
		b.WriteString("\t\treturn nil\n")
	}
	b.WriteString("\t}\n\n")

	// Allocate buffer
	b.WriteString("\tbuf := C.malloc(dataSize)\n")
	b.WriteString("\tcancels = append(cancels, func() { C.free(buf) })\n\n")

	// Second call: fill buffer
	secondCallArgs := append(append([]string{}, cArgs...), "&dataSize", "buf")
	if c.HasError {
		b.WriteString(fmt.Sprintf("\t_result = C.fn_%s(%s)\n", c.CName, strings.Join(secondCallArgs, ", ")))
		b.WriteString("\tif _result != C.VK_SUCCESS {\n")
		b.WriteString("\t\treturn nil, vkError(_result)\n")
		b.WriteString("\t}\n\n")
	} else {
		b.WriteString(fmt.Sprintf("\tC.fn_%s(%s)\n", c.CName, strings.Join(secondCallArgs, ", ")))
	}

	if c.HasError {
		b.WriteString("\treturn C.GoBytes(buf, C.int(dataSize)), nil\n")
	} else {
		b.WriteString("\treturn C.GoBytes(buf, C.int(dataSize))\n")
	}
	b.WriteString("}\n\n")
	return b.String()
}

// GenerateCWrapperDecl returns the C function declaration for the .h file.
func (c *GoCommand) GenerateCWrapperDecl() string {
	params := make([]string, len(c.CParams))
	for i, p := range c.CParams {
		params[i] = p.Type + " " + p.Name
	}
	paramStr := strings.Join(params, ", ")
	if paramStr == "" {
		paramStr = "void"
	}
	return fmt.Sprintf("%s fn_%s(%s);\n", c.CReturnType, c.CName, paramStr)
}

// GenerateCWrapperImpl returns the C function implementation for the .c file.
func (c *GoCommand) GenerateCWrapperImpl() string {
	params := make([]string, len(c.CParams))
	args := make([]string, len(c.CParams))
	for i, p := range c.CParams {
		params[i] = p.Type + " " + p.Name
		args[i] = p.Name
	}
	paramStr := strings.Join(params, ", ")
	if paramStr == "" {
		paramStr = "void"
	}
	argStr := strings.Join(args, ", ")

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s fn_%s(%s) {\n", c.CReturnType, c.CName, paramStr))
	if c.CReturnType == "void" {
		b.WriteString(fmt.Sprintf("\t%s(%s);\n", c.CName, argStr))
	} else {
		b.WriteString(fmt.Sprintf("\treturn %s(%s);\n", c.CName, argStr))
	}
	b.WriteString("}\n\n")
	return b.String()
}
