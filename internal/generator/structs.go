package generator

import (
	"fmt"
	"strings"
)

type StructField struct {
	GoName     string
	CName      string
	CountFor   string // if non-empty, this field is the count for field CountFor (skip in Go)
	CountCName string // C name of the count field
	CountScale int    // if > 1, C count = len(slice) * CountScale (e.g. codeSize = len(Code)*4)
	Type       FieldType
	BitWidth   int // non-zero for C bitfield members (e.g. instanceCustomIndex:24)
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

// bitpackEntry records where a bitfield member lives within a packed uint32 field.
type bitpackEntry struct {
	packName string // e.g. "_bitpack0"
	offset   int    // bit offset within the pack
}

// bitfieldMap builds a map from CName → bitpackEntry for all bitfield members.
func (s *Structured) bitfieldMap() map[string]bitpackEntry {
	result := make(map[string]bitpackEntry)
	packIdx := 0
	bitOffset := 0
	inPack := false

	for _, f := range s.Fields {
		if f.BitWidth <= 0 {
			if inPack {
				packIdx++
				bitOffset = 0
				inPack = false
			}
			continue
		}
		if !inPack {
			bitOffset = 0
			inPack = true
		} else if bitOffset+f.BitWidth > 32 {
			packIdx++
			bitOffset = 0
		}
		result[f.CName] = bitpackEntry{
			packName: fmt.Sprintf("_bitpack%d", packIdx),
			offset:   bitOffset,
		}
		bitOffset += f.BitWidth
	}
	return result
}

// GenerateCLayoutStruct emits the unexported cVkFoo struct (or [N]byte for unions)
// that mirrors the C memory layout used for FFI calls.
func (s *Structured) GenerateCLayoutStruct(structs map[string]*Structured) string {
	if s == nil {
		return ""
	}
	if s.IsUnion {
		sz := computeUnionLayoutSize(s, structs)
		return fmt.Sprintf("type c%s [%d]byte\n\n", s.CName, sz)
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("type c%s struct {\n", s.CName))

	if s.HasSType {
		b.WriteString("\tsType int32\n")
		b.WriteString("\tpNext unsafe.Pointer\n")
	}

	packIdx := 0
	bitOffset := 0
	inPack := false
	emittedPacks := make(map[string]bool)

	for _, f := range s.Fields {
		if f.BitWidth <= 0 {
			if inPack {
				packIdx++
				bitOffset = 0
				inPack = false
			}
			fieldName := sanitizeCField(f.CName)
			b.WriteString(fmt.Sprintf("\t%s %s\n", fieldName, f.Type.CName()))
			continue
		}
		// Bitfield member — fold into a packed uint32 field.
		if !inPack {
			bitOffset = 0
			inPack = true
		} else if bitOffset+f.BitWidth > 32 {
			packIdx++
			bitOffset = 0
		}
		packName := fmt.Sprintf("_bitpack%d", packIdx)
		if !emittedPacks[packName] {
			emittedPacks[packName] = true
			b.WriteString(fmt.Sprintf("\t%s uint32\n", packName))
		}
		bitOffset += f.BitWidth
	}

	b.WriteString("}\n\n")
	return b.String()
}

func (s *Structured) GenerateStructure(structs map[string]*Structured) string {
	if s == nil {
		return ""
	}
	if s.IsUnion {
		return s.generateUnionType(structs)
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

func (s *Structured) generateUnionType(structs map[string]*Structured) string {
	var b strings.Builder

	byteSize := computeUnionLayoutSize(s, structs)

	// Type definition: byte array sized to the C union
	b.WriteString(fmt.Sprintf("type %s [%d]byte\n\n", s.GoName, byteSize))

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
	g.Line(fmt.Sprintf("\tp := new(c%s)", s.CName))

	if s.HasSType {
		g.Line("\tp.sType = int32(s.GetType())")
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

	bfMap := s.bitfieldMap()

	for _, field := range s.Fields {
		if field.CountFor != "" {
			continue
		}

		// Bitfield member: pack into _bitpackN using bit manipulation.
		if field.BitWidth > 0 {
			info := bfMap[field.CName]
			mask := uint32((1 << field.BitWidth) - 1)
			input := "s." + field.GoName
			g.Line(fmt.Sprintf("\tp.%s |= (uint32(%s) & 0x%x) << %d", info.packName, input, mask, info.offset))
			continue
		}

		// PFN callback fields: assign goffi trampoline var.
		if _, ok := field.Type.(*GoFuncPointer); ok {
			if hasPFN && userDataCName != "" {
				trampolineName := "_trampoline_" + s.GoName + "_" + field.GoName
				cFieldName := sanitizeCField(field.CName)
				g.Line(fmt.Sprintf("\tp.%s = %s", cFieldName, trampolineName))
			}
			continue
		}

		// UserData field: override with holder pointer when callbacks are present.
		if hasPFN && userDataCName != "" && field.CName == userDataCName {
			cFieldName := sanitizeCField(field.CName)
			g.Line(fmt.Sprintf("\tp.%s = %s", cFieldName, holderPtrVar))
			continue
		}

		input := "s." + field.GoName
		output := field.Type.GenerateToC(g, input)
		cFieldName := sanitizeCField(field.CName)
		g.Line(fmt.Sprintf("\tp.%s = %s", cFieldName, output))

		// Auto-set collapsed count field.
		_, isSliceField := field.Type.(*Slice)
		_, isByteSliceField := field.Type.(*ByteSlice)
		if (isSliceField || isByteSliceField) && field.CountCName != "" {
			var countFieldType FieldType
			for _, f := range s.Fields {
				if f.CName == field.CountCName && f.CountFor != "" {
					countFieldType = f.Type
					break
				}
			}
			if countFieldType != nil {
				cTypeName := countFieldType.CName()
				if field.CountScale > 1 {
					g.Line(fmt.Sprintf("\tp.%s = %s(len(%s) * %d)", field.CountCName, cTypeName, input, field.CountScale))
				} else {
					g.Line(fmt.Sprintf("\tp.%s = %s(len(%s))", field.CountCName, cTypeName, input))
				}
			}
		}
	}

	g.Line("\treturn unsafe.Pointer(p), func() {")
	g.Line("\t\tfor _, cancel := range cancels { cancel() }")
	g.Line("\t}")
	g.Line("}")
	g.Line("")

	return g.String()
}

func (s *Structured) generateUnionToC() string {
	var b strings.Builder
	cTypeName := "c" + s.CName
	b.WriteString(fmt.Sprintf("func (s *%s) toC() (unsafe.Pointer, func()) {\n", s.GoName))
	b.WriteString(fmt.Sprintf("\tp := new(%s)\n", cTypeName))
	b.WriteString("\tcopy((*p)[:], s[:])\n")
	b.WriteString("\treturn unsafe.Pointer(p), func() {}\n")
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
	g.Line(fmt.Sprintf("func (s *%s) fromC(p *c%s) {", s.GoName, s.CName))

	bfMap := s.bitfieldMap()

	for _, field := range s.Fields {
		if field.CountFor != "" {
			continue
		}

		// Bitfield member: extract via bit mask and shift.
		if field.BitWidth > 0 {
			info := bfMap[field.CName]
			mask := uint32((1 << field.BitWidth) - 1)
			goType := field.Type.GoName()
			g.Line(fmt.Sprintf("\ts.%s = %s((p.%s >> %d) & 0x%x)", field.GoName, goType, info.packName, info.offset, mask))
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
			g.Line(fmt.Sprintf("\t%s = &%s{handle: %s}", goField, ft.Name, input))
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
				g.Line(fmt.Sprintf("\t\t%s[_i] = &%s{handle: %s[_i]}", goField, child.Name, input))
				g.Line("\t}")
			case *FixedArray:
				switch inner := child.Child.(type) {
				case *Primitive:
					g.Line(fmt.Sprintf("\tfor _i := range %s {", goField))
					g.Line(fmt.Sprintf("\t\tfor _j := range %s[_i] {", goField))
					g.Line(fmt.Sprintf("\t\t\t%s[_i][_j] = %s(%s[_i][_j])", goField, inner.GoName(), input))
					g.Line("\t\t}")
					g.Line("\t}")
				default:
					g.Line(fmt.Sprintf("\t// TODO: fromC for %s (FixedArray[FixedArray[%T]])", field.GoName, inner))
				}
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
			g.Line(fmt.Sprintf("\t%s = %s", goField, input))
		case *ByteSlice:
			if field.CountCName != "" {
				countCField := sanitizeCField(field.CountCName)
				g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
				g.Line(fmt.Sprintf("\t\t%s = make([]byte, p.%s)", goField, countCField))
				g.Line(fmt.Sprintf("\t\tcopy(%s, unsafe.Slice((*byte)(%s), p.%s))", goField, input, countCField))
				g.Line("\t}")
			}
		case *String:
			g.Line(fmt.Sprintf("\tif %s != nil { %s = _cGoString(%s) }", input, goField, input))
		case *Pointer:
			switch child := ft.Child.(type) {
			case *StructType:
				g.Line(fmt.Sprintf("\tif %s != nil { %s.fromC((*c%s)(%s)) }", input, goField, child.CTypeName, input))
			case *Handle:
				g.Line(fmt.Sprintf("\tif %s != nil { %s = &%s{handle: *(*unsafe.Pointer)(%s)} }", input, goField, child.Name, input))
			default:
				g.Line(fmt.Sprintf("\t// TODO: fromC for %s (Pointer of %T)", field.GoName, ft.Child))
			}
		case *Slice:
			if field.CountCName != "" {
				countCField := sanitizeCField(field.CountCName)
				var countExpr string
				if field.CountScale > 1 {
					countExpr = fmt.Sprintf("int(p.%s) / %d", countCField, field.CountScale)
				} else {
					countExpr = fmt.Sprintf("p.%s", countCField)
				}
				indexVar := g.Var("i")
				elemVar := g.Var("elem")
				switch child := ft.Child.(type) {
				case *Primitive:
					g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
					g.Line(fmt.Sprintf("\t\t%s = make([]%s, %s)", goField, child.GoName(), countExpr))
					g.Line(fmt.Sprintf("\t\tfor %s := range %s {", indexVar, goField))
					g.Line(fmt.Sprintf("\t\t\t%s[%s] = %s((*[1<<30]%s)(unsafe.Pointer(%s))[%s])", goField, indexVar, child.GoName(), child.CName(), input, indexVar))
					g.Line("\t\t}")
					g.Line("\t}")
				case *StructType:
					g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
					g.Line(fmt.Sprintf("\t\t%s = make([]%s, %s)", goField, child.Name, countExpr))
					g.Line(fmt.Sprintf("\t\tfor %s := range %s {", indexVar, goField))
					g.Line(fmt.Sprintf("\t\t\t%s := (*[1<<30]c%s)(unsafe.Pointer(%s))[%s]", elemVar, child.CTypeName, input, indexVar))
					g.Line(fmt.Sprintf("\t\t\t%s[%s].fromC(&%s)", goField, indexVar, elemVar))
					g.Line("\t\t}")
					g.Line("\t}")
				case *Handle:
					g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
					g.Line(fmt.Sprintf("\t\t%s = make([]%s, %s)", goField, child.GoName(), countExpr))
					g.Line(fmt.Sprintf("\t\tfor %s := range %s {", indexVar, goField))
					g.Line(fmt.Sprintf("\t\t\t%s[%s] = &%s{handle: (*[1<<30]unsafe.Pointer)(unsafe.Pointer(%s))[%s]}", goField, indexVar, child.Name, input, indexVar))
					g.Line("\t\t}")
					g.Line("\t}")
				default:
					g.Line(fmt.Sprintf("\t// TODO: fromC for %s (Slice of %T)", field.GoName, ft.Child))
				}
			}
		case *PtrSlice:
			if field.CountCName != "" {
				countCField := sanitizeCField(field.CountCName)
				if st, ok := ft.Child.(*StructType); ok {
					indexVar := g.Var("i")
					elemVar := g.Var("elem")
					goElemVar := g.Var("goElem")
					g.Line(fmt.Sprintf("\tif p.%s > 0 && %s != nil {", countCField, input))
					g.Line(fmt.Sprintf("\t\t%s = make([]*%s, p.%s)", goField, st.Name, countCField))
					g.Line(fmt.Sprintf("\t\tfor %s := range %s {", indexVar, goField))
					g.Line(fmt.Sprintf("\t\t\t%s := (*[1<<30]unsafe.Pointer)(unsafe.Pointer(%s))[%s]", elemVar, input, indexVar))
					g.Line(fmt.Sprintf("\t\t\tif %s != nil {", elemVar))
					g.Line(fmt.Sprintf("\t\t\t\tvar %s %s", goElemVar, st.Name))
					g.Line(fmt.Sprintf("\t\t\t\t%s.fromC((*c%s)(%s))", goElemVar, st.CTypeName, elemVar))
					g.Line(fmt.Sprintf("\t\t\t\t%s[%s] = &%s", goField, indexVar, goElemVar))
					g.Line("\t\t\t}")
					g.Line("\t\t}")
					g.Line("\t}")
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
	cTypeName := "c" + s.CName
	b.WriteString(fmt.Sprintf("func (s *%s) fromC(p *%s) {\n", s.GoName, cTypeName))
	b.WriteString("\tcopy(s[:], (*p)[:])\n")
	b.WriteString("}\n\n")
	return b.String()
}

// GenerateCallbacksSupport generates the Go callback holder struct, registry
// sync.Map, and ffi.NewCallback trampolines for a struct that has PFN fields.
func (s *Structured) GenerateCallbacksSupport() string {
	pfnFields := s.pfnFields()
	if len(pfnFields) == 0 {
		return ""
	}
	_, userDataCName := s.userDataField()
	if userDataCName == "" {
		return ""
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

	// ffi.NewCallback trampolines (package-level vars)
	for _, f := range pfnFields {
		fp := f.Type.(*GoFuncPointer)
		b.WriteString(s.generateGoffiTrampoline(f, fp, holderType, registryName))
	}

	return b.String()
}

// generateGoffiTrampoline emits a package-level var holding a ffi.NewCallback
// trampoline for a single PFN field.
func (s *Structured) generateGoffiTrampoline(field StructField, fp *GoFuncPointer, holderType, registryName string) string {
	trampolineName := "_trampoline_" + s.GoName + "_" + field.GoName

	// Build Go func params using CName() (Go primitive types matching C ABI).
	var goParams []string
	for i, p := range fp.Params {
		goParams = append(goParams, fmt.Sprintf("p%d %s", i, p.Type.CName()))
	}

	// Return type in Go primitive form.
	var retCType string
	if fp.Return != nil {
		retCType = fp.Return.CName()
	}

	var b strings.Builder

	funcSig := fmt.Sprintf("func(%s)", strings.Join(goParams, ", "))
	if retCType != "" {
		funcSig += " " + retCType
	}
	b.WriteString(fmt.Sprintf("var %s uintptr = ffi.NewCallback(%s {\n", trampolineName, funcSig))

	// Look up holder via pUserData param.
	udIdx := pUserDataParamIndex(fp)
	b.WriteString(fmt.Sprintf("\tv, ok := %s.Load(p%d)\n", registryName, udIdx))
	b.WriteString("\tif !ok {\n")
	b.WriteString(returnZeroValueGo(fp.Return))
	b.WriteString("\t}\n")
	b.WriteString(fmt.Sprintf("\tholder := v.(*%s)\n", holderType))
	b.WriteString(fmt.Sprintf("\tif holder.%s == nil {\n", field.GoName))
	b.WriteString(returnZeroValueGo(fp.Return))
	b.WriteString("\t}\n")

	// Build call args converting C params to Go.
	var callArgs []string
	for i, p := range fp.Params {
		if p.Name == "userData" || p.Name == "pUserData" {
			callArgs = append(callArgs, "holder.UserData")
			continue
		}
		callArgs = append(callArgs, convertCParamToGoGoffi(fmt.Sprintf("p%d", i), p.Type))
	}

	callExpr := fmt.Sprintf("holder.%s(%s)", field.GoName, strings.Join(callArgs, ", "))

	if fp.Return == nil {
		b.WriteString(fmt.Sprintf("\t%s\n", callExpr))
	} else {
		b.WriteString(fmt.Sprintf("\tresult := %s\n", callExpr))
		b.WriteString(convertGoReturnToCGoffi(fp.Return))
	}

	b.WriteString("})\n\n")
	return b.String()
}

// pUserDataParamIndex returns the index of the userData/pUserData parameter.
func pUserDataParamIndex(fp *GoFuncPointer) int {
	for i, p := range fp.Params {
		if p.Name == "userData" || p.Name == "pUserData" {
			return i
		}
	}
	return len(fp.Params) - 1
}

// convertCParamToGoGoffi converts a raw C param (Go primitive type) to the Go type
// expected by the user-facing callback signature.
func convertCParamToGoGoffi(expr string, ft FieldType) string {
	switch t := ft.(type) {
	case *Primitive:
		return fmt.Sprintf("%s(%s)", t.GoName(), expr)
	case *Bool:
		return fmt.Sprintf("%s != 0", expr)
	case *NamedType:
		return fmt.Sprintf("%s(%s)", t.Name, expr)
	case *VoidPtr:
		return expr
	case *Pointer:
		if st, ok := t.Child.(*StructType); ok {
			goVar := "_go_" + expr
			return fmt.Sprintf("func() *%s { var %s %s; %s.fromC((*c%s)(%s)); return &%s }()",
				st.Name, goVar, st.Name, goVar, st.CTypeName, expr, goVar)
		}
		return expr
	default:
		return expr
	}
}

// convertGoReturnToCGoffi emits a return statement converting a Go result to the
// C primitive type expected by the ffi.NewCallback trampoline.
func convertGoReturnToCGoffi(ft FieldType) string {
	switch ft.(type) {
	case *Bool:
		return "\tif result { return uint32(1) }\n\treturn uint32(0)\n"
	case *VoidPtr:
		return "\treturn result\n"
	default:
		cName := ft.CName()
		if cName == "unsafe.Pointer" {
			return "\treturn result\n"
		}
		return fmt.Sprintf("\treturn %s(result)\n", cName)
	}
}

// returnZeroValueGo emits an early-return zero value for a goffi trampoline return type.
func returnZeroValueGo(ft FieldType) string {
	if ft == nil {
		return "\t\treturn\n"
	}
	switch ft.(type) {
	case *Bool:
		return "\t\treturn uint32(0)\n"
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
