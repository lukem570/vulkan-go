package generator

import (
	"fmt"
	"strings"
)

type StructField struct {
	GoName    string
	CName     string
	CountFor  string // if non-empty, this field is the count for field CountFor (skip in Go)
	CountCName string // C name of the count field
	Type      FieldType
}

type Structured struct {
	CName    string
	GoName   string
	Fields   []StructField
	Platform string // non-empty for platform-specific structs

	HasSType bool
	SType    string // e.g. StructureTypeInstanceCreateInfo
	IsUnion  bool
}

func (s *Structured) GenerateStructure() string {
	if s == nil {
		return ""
	}
	if s.IsUnion {
		return s.generateUnionType()
	}

	var b strings.Builder
	pfnFieldsCheck := s.pfnFields()
	_, userDataCNameCheck := s.userDataField()
	hasCallbacks := len(pfnFieldsCheck) > 0 && userDataCNameCheck != ""

	b.WriteString(fmt.Sprintf("type %s struct {\n", s.GoName))

	if s.HasSType {
		b.WriteString("\tNext Structure\n")
	}

	for _, field := range s.Fields {
		if field.CountFor != "" {
			continue // skip count fields — collapsed into slice
		}
		goType := field.Type.GoName()
		b.WriteString(fmt.Sprintf("\t%s %s\n", field.GoName, goType))
	}

	if hasCallbacks {
		b.WriteString("\tcallbackCleanupFn func()\n")
	}

	b.WriteString("}\n\n")
	return b.String()
}

func (s *Structured) generateUnionType() string {
	var b strings.Builder

	// Type definition: byte array sized to the C union
	b.WriteString(fmt.Sprintf("type %s [C.sizeof_%s]byte\n\n", s.GoName, s.CName))

	// Constructor per variant
	for _, field := range s.Fields {
		goType := field.Type.GoName()
		b.WriteString(fmt.Sprintf("func New%s%s(val %s) %s {\n", s.GoName, field.GoName, goType, s.GoName))
		b.WriteString(fmt.Sprintf("\tvar u %s\n", s.GoName))
		b.WriteString(fmt.Sprintf("\t*(*%s)(unsafe.Pointer(&u[0])) = val\n", goType))
		b.WriteString("\treturn u\n")
		b.WriteString("}\n\n")
	}

	// Getter per variant
	for _, field := range s.Fields {
		goType := field.Type.GoName()
		b.WriteString(fmt.Sprintf("func (u *%s) %s() %s {\n", s.GoName, field.GoName, goType))
		b.WriteString(fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&u[0]))\n", goType))
		b.WriteString("}\n\n")
	}

	return b.String()
}

func (s *Structured) GenerateGetType() string {
	if s == nil || !s.HasSType {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("func (s *%s) GetType() StructureType {\n", s.GoName))
	b.WriteString(fmt.Sprintf("\treturn %s\n", s.SType))
	b.WriteString("}\n\n")
	return b.String()
}

func (s *Structured) GenerateToC() string {
	if s == nil {
		return ""
	}
	if s.IsUnion {
		return s.generateUnionToC()
	}

	g := &CodeGen{}
	g.Line(fmt.Sprintf("func (s *%s) toC() (unsafe.Pointer, func()) {", s.GoName))
	g.Line("\tcancels := make([]func(), 0)")
	g.Line(fmt.Sprintf("\tp := (*C.%s)(C.malloc(C.size_t(C.sizeof_%s)))", s.CName, s.CName))
	g.Line(fmt.Sprintf("\t*p = C.%s{}", s.CName))

	if s.HasSType {
		g.Line("\tp.sType = C.VkStructureType(s.GetType())")
		g.Line("\tif s.Next != nil {")
		g.Line("\t\tnextPtr, nextCancel := s.Next.toC()")
		g.Line("\t\tcancels = append(cancels, nextCancel)")
		g.Line("\t\tp.pNext = nextPtr")
		g.Line("\t}")
	}

	// Detect PFN callback fields and find the UserData field
	pfnFields := s.pfnFields()
	hasPFN := len(pfnFields) > 0
	userDataGoName, userDataCName := s.userDataField()
	var holderPtrVar string

	if hasPFN && userDataCName != "" {
		holderType := lowerFirstChar(s.GoName) + "Callbacks"
		registryName := lowerFirstChar(s.GoName) + "CallbackRegistry"
		holderVar := g.Var("holder")
		holderPtrVar = g.Var("holderPtr")

		g.Line(fmt.Sprintf("\t%s := &%s{}", holderVar, holderType))
		for _, pf := range pfnFields {
			g.Line(fmt.Sprintf("\tif s.%s != nil { %s.%s = s.%s }", pf.GoName, holderVar, pf.GoName, pf.GoName))
		}
		g.Line(fmt.Sprintf("\t%s.UserData = s.%s", holderVar, userDataGoName))
		g.Line(fmt.Sprintf("\t%s := unsafe.Pointer(%s)", holderPtrVar, holderVar))
		g.Line(fmt.Sprintf("\t%s.Store(%s, %s)", registryName, holderPtrVar, holderVar))
		g.Line(fmt.Sprintf("\ts.callbackCleanupFn = func() { %s.Delete(%s) }", registryName, holderPtrVar))
	}

	for _, field := range s.Fields {
		if field.CountFor != "" {
			continue
		}

		// PFN callback fields: use C trampoline via unsafe pointer assignment
		// (CGo can't directly assign our trampolines to PFN fields due to type differences)
		if _, ok := field.Type.(*GoFuncPointer); ok {
			if hasPFN && userDataCName != "" {
				trampolineName := "trampoline_" + s.GoName + "_" + field.GoName
				cFieldName := sanitizeCField(field.CName)
				g.Line(fmt.Sprintf("\t*(*unsafe.Pointer)(unsafe.Pointer(&p.%s)) = C.%s", cFieldName, trampolineName))
			}
			// else: no pUserData to route through — skip (leave nil)
			continue
		}

		// UserData field: override with holder pointer when callbacks are present
		if hasPFN && userDataCName != "" && field.CName == userDataCName {
			cFieldName := sanitizeCField(field.CName)
			g.Line(fmt.Sprintf("\tp.%s = %s", cFieldName, holderPtrVar))
			continue
		}

		input := "s." + field.GoName
		output := field.Type.GenerateToC(g, input)

		cFieldName := sanitizeCField(field.CName)
		g.Line(fmt.Sprintf("\tp.%s = %s", cFieldName, output))

		// If this is a slice field and the count field was collapsed, auto-set it
		if _, ok := field.Type.(*Slice); ok && field.CountCName != "" {
			// Only auto-set the count if it was collapsed (CountFor is set)
			collapsed := false
			for _, f := range s.Fields {
				if f.CName == field.CountCName && f.CountFor != "" {
					collapsed = true
					break
				}
			}
			if collapsed {
				g.Line(fmt.Sprintf("\tp.%s = C.uint32_t(len(%s))", field.CountCName, input))
			}
		}
	}

	g.Line("\treturn unsafe.Pointer(p), func() {")
	g.Line("\t\tfor _, cancel := range cancels { cancel() }")
	g.Line("\t\tC.free(unsafe.Pointer(p))")
	g.Line("\t}")
	g.Line("}")
	g.Line("")

	return g.String()
}

func (s *Structured) generateUnionToC() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("func (s *%s) toC() (*C.%s, func()) {\n", s.GoName, s.CName))
	b.WriteString(fmt.Sprintf("\tp := (*C.%s)(C.malloc(C.size_t(C.sizeof_%s)))\n", s.CName, s.CName))
	b.WriteString(fmt.Sprintf("\tC.memcpy(unsafe.Pointer(p), unsafe.Pointer(&s[0]), C.sizeof_%s)\n", s.CName))
	b.WriteString("\treturn p, func() { C.free(unsafe.Pointer(p)) }\n")
	b.WriteString("}\n\n")
	return b.String()
}

func (s *Structured) GenerateFromC() string {
	if s == nil {
		return ""
	}
	if s.IsUnion {
		return s.generateUnionFromC()
	}

	g := &CodeGen{}
	g.Line(fmt.Sprintf("func (s *%s) fromC(p *C.%s) {", s.GoName, s.CName))

	for _, field := range s.Fields {
		if field.CountFor != "" {
			continue
		}
		cFieldName := sanitizeCField(field.CName)
		input := "p." + cFieldName
		goField := "s." + field.GoName

		switch ft := field.Type.(type) {
		case *Primitive:
			g.Line(fmt.Sprintf("\t%s = %s(%s)", goField, ft.GoName(), input))
		case *Bool:
			g.Line(fmt.Sprintf("\t%s = %s != 0", goField, input))
		case *NamedType:
			g.Line(fmt.Sprintf("\t%s = %s(%s)", goField, ft.Name, input))
		case *Handle:
			g.Line(fmt.Sprintf("\t%s = &%s{handle: unsafe.Pointer(%s)}", goField, ft.Name, input))
		case *FixedArray:
			switch child := ft.Child.(type) {
			case *Primitive:
				g.Line(fmt.Sprintf("\tfor _i := range %s {", goField))
				g.Line(fmt.Sprintf("\t\t%s[_i] = %s(%s[_i])", goField, child.GoName(), input))
				g.Line("\t}")
			case *NamedType:
				g.Line(fmt.Sprintf("\tfor _i := range %s {", goField))
				g.Line(fmt.Sprintf("\t\t%s[_i] = %s(%s[_i])", goField, child.Name, input))
				g.Line("\t}")
			case *StructType:
				g.Line(fmt.Sprintf("\tfor _i := range %s {", goField))
				g.Line(fmt.Sprintf("\t\t%s[_i].fromC(&%s[_i])", goField, input))
				g.Line("\t}")
			case *Handle:
				g.Line(fmt.Sprintf("\tfor _i := range %s {", goField))
				g.Line(fmt.Sprintf("\t\t%s[_i] = &%s{handle: unsafe.Pointer(%s[_i])}", goField, child.Name, input))
				g.Line("\t}")
			default:
				g.Line(fmt.Sprintf("\t// TODO: fromC for %s (FixedArray of %T)", field.GoName, child))
			}
		case *StructType:
			g.Line(fmt.Sprintf("\t%s.fromC(&%s)", goField, input))
		case *ExternalType:
			if ft.PtrInVulkan || ft.GoTypeName == "unsafe.Pointer" {
				g.Line(fmt.Sprintf("\t%s = unsafe.Pointer(%s)", goField, input))
			} else {
				g.Line(fmt.Sprintf("\t%s = %s(%s)", goField, ft.GoTypeName, input))
			}
		case *VoidPtr:
			g.Line(fmt.Sprintf("\t%s = unsafe.Pointer(%s)", goField, input))
		case *String:
			g.Line(fmt.Sprintf("\tif %s != nil { %s = C.GoString(%s) }", input, goField, input))
		case *Pointer:
			switch child := ft.Child.(type) {
			case *StructType:
				g.Line(fmt.Sprintf("\tif %s != nil { %s.fromC(%s) }", input, goField, input))
			case *Handle:
				g.Line(fmt.Sprintf("\tif %s != nil { %s = &%s{handle: unsafe.Pointer(*%s)} }", input, goField, child.Name, input))
			default:
				g.Line(fmt.Sprintf("\t// TODO: fromC for %s (Pointer of %T)", field.GoName, ft.Child))
			}
		case *Slice:
			if field.CountCName != "" {
				countCField := sanitizeCField(field.CountCName)
				indexVar := g.Var("i")
				elemVar := g.Var("elem")
				switch child := ft.Child.(type) {
				case *StructType:
					g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
					g.Line(fmt.Sprintf("\t\t%s = make([]%s, p.%s)", goField, child.Name, countCField))
					g.Line(fmt.Sprintf("\t\tfor %s := range %s {", indexVar, goField))
					g.Line(fmt.Sprintf("\t\t\t%s := (*[1<<30]C.%s)(unsafe.Pointer(%s))[%s]", elemVar, child.CTypeName, input, indexVar))
					g.Line(fmt.Sprintf("\t\t\t%s[%s].fromC(&%s)", goField, indexVar, elemVar))
					g.Line("\t\t}")
					g.Line("\t}")
				case *Handle:
					g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
					g.Line(fmt.Sprintf("\t\t%s = make([]%s, p.%s)", goField, child.GoName(), countCField))
					g.Line(fmt.Sprintf("\t\tfor %s := range %s {", indexVar, goField))
					g.Line(fmt.Sprintf("\t\t\t%s[%s] = &%s{handle: unsafe.Pointer((*[1<<30]C.%s)(unsafe.Pointer(%s))[%s])}", goField, indexVar, child.Name, child.CTypeName, input, indexVar))
					g.Line("\t\t}")
					g.Line("\t}")
				default:
					g.Line(fmt.Sprintf("\t// TODO: fromC for %s (Slice of %T)", field.GoName, ft.Child))
				}
			}
		default:
			g.Line(fmt.Sprintf("\t// TODO: fromC for %s (%T)", field.GoName, ft))
		}
	}

	g.Line("}")
	g.Line("")
	return g.String()
}

func (s *Structured) generateUnionFromC() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("func (s *%s) fromC(p *C.%s) {\n", s.GoName, s.CName))
	b.WriteString(fmt.Sprintf("\tC.memcpy(unsafe.Pointer(&s[0]), unsafe.Pointer(p), C.sizeof_%s)\n", s.CName))
	b.WriteString("}\n\n")
	return b.String()
}

// GenerateCallbacksSupport generates the Go callback holder struct, registry
// sync.Map, and //export bridge functions for a struct that has PFN fields.
// The output is written into the separate callbacks.go file so //export and
// the main C preamble don't conflict.
func (s *Structured) GenerateCallbacksSupport() string {
	pfnFields := s.pfnFields()
	if len(pfnFields) == 0 {
		return ""
	}
	_, userDataCName := s.userDataField()
	if userDataCName == "" {
		return "" // need pUserData to route callbacks
	}

	holderType := lowerFirstChar(s.GoName) + "Callbacks"
	registryName := lowerFirstChar(s.GoName) + "CallbackRegistry"

	var b strings.Builder

	// Holder struct
	b.WriteString(fmt.Sprintf("type %s struct {\n", holderType))
	for _, f := range pfnFields {
		b.WriteString(fmt.Sprintf("\t%s %s\n", f.GoName, f.Type.GoName()))
	}
	b.WriteString("\tUserData unsafe.Pointer\n")
	b.WriteString("}\n\n")

	// Registry sync.Map
	b.WriteString(fmt.Sprintf("var %s sync.Map\n\n", registryName))

	// Bridge functions
	for _, f := range pfnFields {
		fp := f.Type.(*GoFuncPointer)
		b.WriteString(s.generateBridgeFunction(f, fp, holderType, registryName))
	}

	return b.String()
}

// GenerateCCallbackDecls generates extern + trampoline declarations for the C header.
func (s *Structured) GenerateCCallbackDecls() string {
	pfnFields := s.pfnFields()
	if len(pfnFields) == 0 {
		return ""
	}
	_, userDataCName := s.userDataField()
	if userDataCName == "" {
		return ""
	}

	var b strings.Builder
	for _, f := range pfnFields {
		fp := f.Type.(*GoFuncPointer)
		bridgeName := "go_bridge_" + s.GoName + "_" + f.GoName
		trampolineName := "trampoline_" + s.GoName + "_" + f.GoName

		sig := buildCFuncSig(fp)
		b.WriteString(fmt.Sprintf("extern %s %s(%s);\n", sig.ret, bridgeName, sig.params))
		b.WriteString(fmt.Sprintf("%s %s(%s);\n", sig.ret, trampolineName, sig.params))
	}
	return b.String()
}

// GenerateCCallbackImpls generates trampoline implementations for the C source.
func (s *Structured) GenerateCCallbackImpls() string {
	pfnFields := s.pfnFields()
	if len(pfnFields) == 0 {
		return ""
	}
	_, userDataCName := s.userDataField()
	if userDataCName == "" {
		return ""
	}

	var b strings.Builder
	for _, f := range pfnFields {
		fp := f.Type.(*GoFuncPointer)
		bridgeName := "go_bridge_" + s.GoName + "_" + f.GoName
		trampolineName := "trampoline_" + s.GoName + "_" + f.GoName

		sig := buildCFuncSig(fp)
		b.WriteString(fmt.Sprintf("%s %s(%s) {\n", sig.ret, trampolineName, sig.paramsNamed))
		if sig.ret == "void" {
			b.WriteString(fmt.Sprintf("\t%s(%s);\n", bridgeName, sig.argNames))
		} else {
			b.WriteString(fmt.Sprintf("\treturn %s(%s);\n", bridgeName, sig.argNames))
		}
		b.WriteString("}\n")
	}
	return b.String()
}

type cFuncSig struct {
	ret         string // return type
	params      string // type-only params e.g. "VkBool32, void*"
	paramsNamed string // named params e.g. "VkBool32 p0, void* p1"
	argNames    string // arg names only e.g. "p0, p1"
}

func buildCFuncSig(fp *GoFuncPointer) cFuncSig {
	var retStr string
	if fp.Return == nil {
		retStr = "void"
	} else {
		retStr = ToCTypeName(fp.Return)
	}

	var typeOnly, named, args []string
	for i, p := range fp.Params {
		cType := ToCTypeName(p.Type)
		argName := fmt.Sprintf("p%d", i)
		typeOnly = append(typeOnly, cType)
		named = append(named, cType+" "+argName)
		args = append(args, argName)
	}

	return cFuncSig{
		ret:         retStr,
		params:      strings.Join(typeOnly, ", "),
		paramsNamed: strings.Join(named, ", "),
		argNames:    strings.Join(args, ", "),
	}
}

// generateBridgeFunction generates the //export Go function for a single PFN field.
func (s *Structured) generateBridgeFunction(field StructField, fp *GoFuncPointer, holderType, registryName string) string {
	bridgeName := "go_bridge_" + s.GoName + "_" + field.GoName

	// Build the Go function signature using C types
	var goParams []string
	for i, p := range fp.Params {
		goParams = append(goParams, fmt.Sprintf("p%d %s", i, p.Type.CName()))
	}

	// Return type in C
	var retCType string
	if fp.Return == nil {
		retCType = ""
	} else {
		retCType = fp.Return.CName()
	}

	var b strings.Builder

	// //export + func signature
	b.WriteString(fmt.Sprintf("//export %s\n", bridgeName))
	sig := fmt.Sprintf("func %s(%s)", bridgeName, strings.Join(goParams, ", "))
	if retCType != "" {
		sig += " " + retCType
	}
	b.WriteString(sig + " {\n")

	// Look up holder
	b.WriteString(fmt.Sprintf("\tv, ok := %s.Load(p%d)\n", registryName, pUserDataParamIndex(fp)))
	b.WriteString("\tif !ok {\n")
	b.WriteString(returnZeroValue(fp.Return))
	b.WriteString("\t}\n")
	b.WriteString(fmt.Sprintf("\tholder := v.(*%s)\n", holderType))
	b.WriteString(fmt.Sprintf("\tif holder.%s == nil {\n", field.GoName))
	b.WriteString(returnZeroValue(fp.Return))
	b.WriteString("\t}\n")

	// Convert C params to Go and build call args
	var callArgs []string
	for i, p := range fp.Params {
		paramExpr := fmt.Sprintf("p%d", i)
		// Replace the pUserData parameter with holder.UserData
		if p.Name == "userData" || p.Name == "pUserData" {
			callArgs = append(callArgs, "holder.UserData")
			continue
		}
		callArgs = append(callArgs, convertCParamToGo(paramExpr, p.Type))
	}

	callExpr := fmt.Sprintf("holder.%s(%s)", field.GoName, strings.Join(callArgs, ", "))

	if fp.Return == nil {
		b.WriteString(fmt.Sprintf("\t%s\n", callExpr))
	} else {
		b.WriteString(fmt.Sprintf("\tresult := %s\n", callExpr))
		b.WriteString(convertGoReturnToC(fp.Return))
	}

	b.WriteString("}\n\n")
	return b.String()
}

// pUserDataParamIndex returns the index of the userData/pUserData parameter.
func pUserDataParamIndex(fp *GoFuncPointer) int {
	for i, p := range fp.Params {
		if p.Name == "userData" || p.Name == "pUserData" {
			return i
		}
	}
	// Fallback: last param
	return len(fp.Params) - 1
}

// convertCParamToGo generates a Go expression that converts a C parameter to its Go type.
func convertCParamToGo(expr string, ft FieldType) string {
	switch t := ft.(type) {
	case *Primitive:
		return fmt.Sprintf("%s(%s)", t.GoName(), expr)
	case *Bool:
		return fmt.Sprintf("%s != 0", expr)
	case *NamedType:
		return fmt.Sprintf("%s(%s)", t.Name, expr)
	case *VoidPtr:
		return expr // already unsafe.Pointer
	case *Pointer:
		if st, ok := t.Child.(*StructType); ok {
			// Convert *C.VkFoo to *GoFoo via fromC
			goVar := "_go_" + expr
			return fmt.Sprintf("func() *%s { var %s %s; %s.fromC(%s); return &%s }()",
				st.Name, goVar, st.Name, goVar, expr, goVar)
		}
		return expr
	default:
		return expr
	}
}

// convertGoReturnToC generates a return statement converting a Go result to a C type.
func convertGoReturnToC(ft FieldType) string {
	switch ft.(type) {
	case *Bool:
		return "\tif result { return C.VkBool32(1) }\n\treturn C.VkBool32(0)\n"
	case *VoidPtr:
		return "\treturn result\n"
	default:
		cName := strings.TrimPrefix(ft.CName(), "C.")
		if cName == "unsafe.Pointer" {
			return "\treturn result\n"
		}
		return fmt.Sprintf("\treturn C.%s(result)\n", cName)
	}
}

// returnZeroValue generates an early-return zero value for a return type.
func returnZeroValue(ft FieldType) string {
	if ft == nil {
		return "\t\treturn\n"
	}
	switch ft.(type) {
	case *Bool:
		return "\t\treturn C.VkBool32(0)\n"
	case *VoidPtr:
		return "\t\treturn nil\n"
	default:
		return "\t\treturn 0\n"
	}
}

// pfnFields returns struct fields whose type is *GoFuncPointer.
func (s *Structured) pfnFields() []StructField {
	var fields []StructField
	for _, f := range s.Fields {
		if f.CountFor != "" {
			continue
		}
		if _, ok := f.Type.(*GoFuncPointer); ok {
			fields = append(fields, f)
		}
	}
	return fields
}

// userDataField returns the Go and C names of the pUserData/userData field, if present.
func (s *Structured) userDataField() (goName, cName string) {
	for _, f := range s.Fields {
		if f.CountFor != "" {
			continue
		}
		if f.CName == "pUserData" || f.CName == "userData" {
			return f.GoName, f.CName
		}
	}
	return "", ""
}

func lowerFirstChar(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func sanitizeCField(name string) string {
	switch name {
	case "type":
		return "_type"
	case "range":
		return "_range"
	default:
		return name
	}
}