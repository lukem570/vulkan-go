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

// GenerateCIFDecl emits the CIF var and function pointer var (global commands only).
func (c *GoCommand) GenerateCIFDecl() string {
	return fmt.Sprintf("var _cif_%s types.CallInterface\nvar _fn_%s unsafe.Pointer\n", c.CName, c.CName)
}

// GenerateCIFOnly emits just the CIF var (for instance/device commands whose fn ptr lives in a dispatch table).
func (c *GoCommand) GenerateCIFOnly() string {
	return fmt.Sprintf("var _cif_%s types.CallInterface\n", c.CName)
}

// GenerateCIFPrepare emits only the ffi.PrepareCallInterface line for a command.
func (c *GoCommand) GenerateCIFPrepare() string {
	retDesc := goffiReturnDesc(c.CReturnType)
	var argDescs []string
	for _, p := range c.CParams {
		argDescs = append(argDescs, cParamGoffiDesc(p.Type))
	}
	argDescStr := strings.Join(argDescs, ", ")
	return fmt.Sprintf("\tffi.PrepareCallInterface(&_cif_%s, types.DefaultCall, %s, []*types.TypeDescriptor{%s}) //nolint\n",
		c.CName, retDesc, argDescStr)
}

// buildArgExpr returns the ffi argument pointer expression for a C arg value.
// After GenerateToC() emits code and returns result, this maps result to unsafe.Pointer(&slot).
// For StructType (*castN), it emits a temp var and returns unsafe.Pointer(&tempVar).
func buildArgExpr(b *strings.Builder, g *CodeGen, result string, ft FieldType) string {
	switch ft.(type) {
	case *StructType:
		// result = "*castN" — strip * to get the *cVkFoo pointer
		castName := strings.TrimPrefix(result, "*")
		tmpVar := g.Var("_argp")
		b.WriteString(fmt.Sprintf("\t%s := unsafe.Pointer(%s)\n", tmpVar, castName))
		return "unsafe.Pointer(&" + tmpVar + ")"
	case *FixedArray:
		return "unsafe.Pointer(&" + result + "[0])"
	default:
		return "unsafe.Pointer(&" + result + ")"
	}
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

	// Compute the fn pointer expression: global uses package var, others use dispatch table.
	level := commandLevel(c)
	var fnExpr string
	if level == "global" {
		fnExpr = "_fn_" + c.CName
	} else {
		fnExpr = "h.table." + c.CName
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

	var ffiArgExprs []string

	// Receiver
	if c.ReceiverType != "" {
		b.WriteString("\t_recv := h.handle\n")
		ffiArgExprs = append(ffiArgExprs, "unsafe.Pointer(&_recv)")
	}

	// Input params
	for _, p := range c.Params {
		if p.CountOf != nil {
			b.WriteString(fmt.Sprintf("\t%s := uint32(len(%s))\n", p.Name, p.CountOf.Name))
			ffiArgExprs = append(ffiArgExprs, "unsafe.Pointer(&"+p.Name+")")
			continue
		}

		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		argExpr := buildArgExpr(&b, g, result, p.Type)
		ffiArgExprs = append(ffiArgExprs, argExpr)
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

	// Output param slots
	for _, op := range c.OutParams {
		varName := sanitizeIdent(op.Name) + "Out"
		if op.IsArray {
			elemCType := op.Type.CName()
			b.WriteString(fmt.Sprintf("\t%s := make([]%s, %s)\n", varName, elemCType, op.CountGoParam))
			b.WriteString(fmt.Sprintf("\tvar _%sPin runtime.Pinner\n", varName))
			b.WriteString(fmt.Sprintf("\t_%sPtr := unsafe.Pointer(nil)\n", varName))
			b.WriteString(fmt.Sprintf("\tif len(%s) > 0 { _%sPin.Pin(&%s[0]); _%sPtr = unsafe.Pointer(&%s[0]) }\n",
				varName, varName, varName, varName, varName))
			b.WriteString(fmt.Sprintf("\tdefer _%sPin.Unpin()\n", varName))
			ffiArgExprs = append(ffiArgExprs, fmt.Sprintf("unsafe.Pointer(&_%sPtr)", varName))
		} else {
			var cType string
			switch t := op.Type.(type) {
			case *Handle:
				cType = "unsafe.Pointer"
			case *StructType:
				cType = "c" + t.CTypeName
			default:
				cType = op.Type.CName()
			}
			b.WriteString(fmt.Sprintf("\tvar %s %s\n", varName, cType))
			// Pre-set sType for struct outputs that need it
			if st, ok := op.Type.(*StructType); ok {
				if stype, ok := sTypes[st.Name]; ok {
					b.WriteString(fmt.Sprintf("\t%s.sType = int32(%s)\n", varName, stype))
				}
			}
			// Slot holds &varName so goffi passes the pointer-to-storage to C.
			slotVar := "_" + varName + "Slot"
			b.WriteString(fmt.Sprintf("\t%s := unsafe.Pointer(&%s)\n", slotVar, varName))
			ffiArgExprs = append(ffiArgExprs, "unsafe.Pointer(&"+slotVar+")")
		}
	}

	// ffi call
	if c.HasError || c.ReturnType != nil {
		retCType := "int32"
		if c.ReturnType != nil && !c.HasError {
			retCType = c.ReturnType.CName()
		}
		b.WriteString(fmt.Sprintf("\tvar _ret %s\n", retCType))
		b.WriteString(fmt.Sprintf("\tffi.CallFunction(&_cif_%s, %s, unsafe.Pointer(&_ret), []unsafe.Pointer{\n", c.CName, fnExpr)) //nolint
	} else {
		b.WriteString(fmt.Sprintf("\tffi.CallFunction(&_cif_%s, %s, nil, []unsafe.Pointer{\n", c.CName, fnExpr)) //nolint
	}
	for _, ae := range ffiArgExprs {
		b.WriteString("\t\t" + ae + ",\n")
	}
	b.WriteString("\t})\n")

	// Error check
	if c.HasError {
		b.WriteString("\tif _ret != int32(Success) {\n\t\treturn ")
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
		b.WriteString("Result(_ret)\n\t}\n")
	}

	// Convert out params from C → Go
	var retVars []string
	var handleRetVar string
	var handleRetCVar string // C-storage variable name for the returned handle
	for _, op := range c.OutParams {
		varName := sanitizeIdent(op.Name) + "Out"
		if op.IsArray {
			sliceVar := g.Var("out")
			b.WriteString(fmt.Sprintf("\t%s := make([]%s, %s)\n", sliceVar, op.Type.GoName(), op.CountGoParam))
			idxVar := g.Var("i")
			elemExpr := fmt.Sprintf("(*[1<<30]%s)(unsafe.Pointer(_%sPtr))[%s]", op.Type.CName(), varName, idxVar)
			b.WriteString(fmt.Sprintf("\tfor %s := range %s {\n", idxVar, sliceVar))
			outG := &CodeGen{varIndex: g.varIndex}
			goVal := op.Type.GenerateFromC(outG, elemExpr)
			b.WriteString(outG.String())
			g.varIndex = outG.varIndex
			b.WriteString(fmt.Sprintf("\t\t%s[%s] = %s\n", sliceVar, idxVar, goVal))
			// Propagate dispatch table to array elements that are handles.
			if ht, ok := op.Type.(*Handle); ok {
				b.WriteString(fmt.Sprintf("\t\t%s.table = h.table\n", goVal))
				_ = ht
			}
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
				handleRetCVar = varName
			}
		}
	}

	// For Create commands: attach callback cleanup to the returned handle
	if c.CallbackStructParamName != "" && handleRetVar != "" {
		b.WriteString(fmt.Sprintf("\tif %s.callbackCleanupFn != nil { %s.cleanup = %s.callbackCleanupFn }\n",
			c.CallbackStructParamName, handleRetVar, c.CallbackStructParamName))
	}

	// Dispatch table init/propagation for returned handles.
	if handleRetVar != "" {
		switch c.CName {
		case "vkCreateInstance":
			b.WriteString(fmt.Sprintf("\t%s.table = _initInstanceTable(%s)\n", handleRetVar, handleRetCVar))
		case "vkCreateDevice":
			b.WriteString(fmt.Sprintf("\t%s.table = _initDeviceTable(%s)\n", handleRetVar, handleRetCVar))
		default:
			// Propagate the receiver's table to any returned handle that shares the same level.
			if h, ok := c.OutParams[len(c.OutParams)-1].Type.(*Handle); ok {
				if (instanceReceiverSet[c.ReceiverType] && instanceReceiverSet[h.Name]) ||
					(deviceReceiverSet[c.ReceiverType] && deviceReceiverSet[h.Name]) {
					b.WriteString(fmt.Sprintf("\t%s.table = h.table\n", handleRetVar))
				}
			}
		}
	}
	_ = handleRetCVar

	// Convert non-VkResult return value
	if c.ReturnType != nil {
		outG := &CodeGen{varIndex: g.varIndex}
		goVal := c.ReturnType.GenerateFromC(outG, "_ret")
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
// enumerate pattern using goffi.
func (c *GoCommand) generateEnumerateWrapper() string {
	ep := c.EnumeratePattern
	elemGoName := ep.ElementType.GoName()
	elemCName := ep.ElementType.CName()
	sliceType := "[]" + elemGoName

	level := commandLevel(c)
	var fnExpr string
	if level == "global" {
		fnExpr = "_fn_" + c.CName
	} else {
		fnExpr = "h.table." + c.CName
	}

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
	if c.HasError {
		b.WriteString(fmt.Sprintf(" (%s, error)", sliceType))
	} else {
		b.WriteString(" " + sliceType)
	}
	b.WriteString(" {\n")

	b.WriteString("\tcancels := make([]func(), 0)\n")
	b.WriteString("\tdefer func() { for _, c := range cancels { c() } }()\n\n")

	g := &CodeGen{}
	var inputArgExprs []string

	if c.ReceiverType != "" {
		b.WriteString("\t_recv := h.handle\n")
		inputArgExprs = append(inputArgExprs, "unsafe.Pointer(&_recv)")
	}
	for _, p := range c.Params {
		if p.CountOf != nil {
			b.WriteString(fmt.Sprintf("\t%s := uint32(len(%s))\n", p.Name, p.CountOf.Name))
			inputArgExprs = append(inputArgExprs, "unsafe.Pointer(&"+p.Name+")")
			continue
		}
		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		inputArgExprs = append(inputArgExprs, buildArgExpr(&b, g, result, p.Type))
	}

	// Slots for count and array params
	b.WriteString("\tvar _count uint32\n")
	b.WriteString("\t_countSlot := unsafe.Pointer(&_count)\n")
	b.WriteString("\tvar _nilSlot unsafe.Pointer\n")

	var retDecl, retPtr string
	if c.HasError {
		b.WriteString("\tvar _ret int32\n")
		retDecl = ""
		retPtr = "unsafe.Pointer(&_ret)"
	} else {
		retDecl = ""
		retPtr = "nil"
	}
	_ = retDecl

	// First call: count only
	firstArgs := append(append([]string{}, inputArgExprs...), "unsafe.Pointer(&_countSlot)", "unsafe.Pointer(&_nilSlot)")
	b.WriteString(fmt.Sprintf("\tffi.CallFunction(&_cif_%s, %s, %s, []unsafe.Pointer{\n", c.CName, fnExpr, retPtr)) //nolint
	for _, ae := range firstArgs {
		b.WriteString("\t\t" + ae + ",\n")
	}
	b.WriteString("\t})\n")

	if c.HasError {
		b.WriteString("\tif _ret != int32(Success) && _ret != int32(Incomplete) {\n")
		b.WriteString("\t\treturn nil, Result(_ret)\n")
		b.WriteString("\t}\n")
	}
	b.WriteString("\tif _count == 0 {\n")
	if c.HasError {
		b.WriteString("\t\treturn nil, nil\n")
	} else {
		b.WriteString("\t\treturn nil\n")
	}
	b.WriteString("\t}\n\n")

	// Allocate array and pin it
	b.WriteString(fmt.Sprintf("\t_arr := make([]%s, _count)\n", elemCName))
	b.WriteString("\tvar _arrPin runtime.Pinner\n")
	b.WriteString("\t_arrSlot := unsafe.Pointer(&_arr[0])\n")
	b.WriteString("\t_arrPin.Pin(&_arr[0])\n")
	b.WriteString("\tdefer _arrPin.Unpin()\n\n")

	// Second call: fill array
	secondArgs := append(append([]string{}, inputArgExprs...), "unsafe.Pointer(&_countSlot)", "unsafe.Pointer(&_arrSlot)")
	b.WriteString(fmt.Sprintf("\tffi.CallFunction(&_cif_%s, %s, %s, []unsafe.Pointer{\n", c.CName, fnExpr, retPtr)) //nolint
	for _, ae := range secondArgs {
		b.WriteString("\t\t" + ae + ",\n")
	}
	b.WriteString("\t})\n")

	if c.HasError {
		b.WriteString("\tif _ret != int32(Success) {\n")
		b.WriteString("\t\treturn nil, Result(_ret)\n")
		b.WriteString("\t}\n\n")
	}

	// Convert C array to Go slice
	b.WriteString(fmt.Sprintf("\tout := make(%s, int(_count))\n", sliceType))
	b.WriteString("\tfor i, v := range _arr {\n")
	convG := &CodeGen{varIndex: g.varIndex}
	goVal := ep.ElementType.GenerateFromC(convG, "v")
	b.WriteString(convG.String())
	b.WriteString(fmt.Sprintf("\t\tout[i] = %s\n", goVal))
	// Propagate dispatch table to elements that are handles.
	if ht, ok := ep.ElementType.(*Handle); ok {
		if (instanceReceiverSet[c.ReceiverType] && instanceReceiverSet[ht.Name]) ||
			(deviceReceiverSet[c.ReceiverType] && deviceReceiverSet[ht.Name]) {
			b.WriteString(fmt.Sprintf("\t\t%s.table = h.table\n", goVal))
		}
	}
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
// pattern using goffi.
func (c *GoCommand) generateByteDataWrapper() string {
	level := commandLevel(c)
	var fnExpr string
	if level == "global" {
		fnExpr = "_fn_" + c.CName
	} else {
		fnExpr = "h.table." + c.CName
	}

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
	if c.HasError {
		b.WriteString(" ([]byte, error)")
	} else {
		b.WriteString(" []byte")
	}
	b.WriteString(" {\n")

	b.WriteString("\tcancels := make([]func(), 0)\n")
	b.WriteString("\tdefer func() { for _, c := range cancels { c() } }()\n\n")

	g := &CodeGen{}
	var inputArgExprs []string

	if c.ReceiverType != "" {
		b.WriteString("\t_recv := h.handle\n")
		inputArgExprs = append(inputArgExprs, "unsafe.Pointer(&_recv)")
	}
	for _, p := range c.Params {
		if p.CountOf != nil {
			b.WriteString(fmt.Sprintf("\t%s := uint32(len(%s))\n", p.Name, p.CountOf.Name))
			inputArgExprs = append(inputArgExprs, "unsafe.Pointer(&"+p.Name+")")
			continue
		}
		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		inputArgExprs = append(inputArgExprs, buildArgExpr(&b, g, result, p.Type))
	}

	// First call: get data size
	b.WriteString("\tvar _dataSize uintptr\n")
	b.WriteString("\t_dataSizeSlot := unsafe.Pointer(&_dataSize)\n")
	b.WriteString("\tvar _nilSlot unsafe.Pointer\n")

	var retPtr string
	if c.HasError {
		b.WriteString("\tvar _ret int32\n")
		retPtr = "unsafe.Pointer(&_ret)"
	} else {
		retPtr = "nil"
	}

	firstArgs := append(append([]string{}, inputArgExprs...), "unsafe.Pointer(&_dataSizeSlot)", "unsafe.Pointer(&_nilSlot)")
	b.WriteString(fmt.Sprintf("\tffi.CallFunction(&_cif_%s, %s, %s, []unsafe.Pointer{\n", c.CName, fnExpr, retPtr)) //nolint
	for _, ae := range firstArgs {
		b.WriteString("\t\t" + ae + ",\n")
	}
	b.WriteString("\t})\n")

	if c.HasError {
		b.WriteString("\tif _ret != int32(Success) && _ret != int32(Incomplete) {\n")
		b.WriteString("\t\treturn nil, Result(_ret)\n")
		b.WriteString("\t}\n")
	}
	b.WriteString("\tif _dataSize == 0 {\n")
	if c.HasError {
		b.WriteString("\t\treturn nil, nil\n")
	} else {
		b.WriteString("\t\treturn nil\n")
	}
	b.WriteString("\t}\n\n")

	// Allocate buffer
	b.WriteString("\t_buf := make([]byte, _dataSize)\n")
	b.WriteString("\tvar _bufPin runtime.Pinner\n")
	b.WriteString("\t_bufSlot := unsafe.Pointer(&_buf[0])\n")
	b.WriteString("\t_bufPin.Pin(&_buf[0])\n")
	b.WriteString("\tdefer _bufPin.Unpin()\n\n")

	// Second call: fill buffer
	secondArgs := append(append([]string{}, inputArgExprs...), "unsafe.Pointer(&_dataSizeSlot)", "unsafe.Pointer(&_bufSlot)")
	b.WriteString(fmt.Sprintf("\tffi.CallFunction(&_cif_%s, %s, %s, []unsafe.Pointer{\n", c.CName, fnExpr, retPtr)) //nolint
	for _, ae := range secondArgs {
		b.WriteString("\t\t" + ae + ",\n")
	}
	b.WriteString("\t})\n")

	if c.HasError {
		b.WriteString("\tif _ret != int32(Success) {\n")
		b.WriteString("\t\treturn nil, Result(_ret)\n")
		b.WriteString("\t}\n\n")
	}

	if c.HasError {
		b.WriteString("\treturn _buf, nil\n")
	} else {
		b.WriteString("\treturn _buf\n")
	}
	b.WriteString("}\n\n")
	return b.String()
}

