package generator

import (
	"fmt"
	"strings"
)

// FieldType generates conversion code and describes a type's layout.
// CName() now returns the C-layout Go type name (no "C." prefix).
// ABIDesc() returns the goffi types.TypeDescriptor expression for the type
// when used as a command parameter.
type FieldType interface {
	GenerateToC(g *CodeGen, input string) string
	GenerateFromC(g *CodeGen, input string) string
	CName() string      // C-layout Go type (e.g. "uint32", "unsafe.Pointer", "cVkExtent2D")
	GoName() string     // Go user-facing type name
	ABIDesc() string  // goffi TypeDescriptor expression for use in CIF
}

// ---------------------------------------------------------------------------
// Primitive
// ---------------------------------------------------------------------------

type Primitive struct {
	cName  string
	goName string
}

func NewPrimitive(cName, goName string) *Primitive {
	return &Primitive{cName: cName, goName: goName}
}

func (t *Primitive) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.goName, input))
	return out
}

func (t *Primitive) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.goName, input))
	return out
}

func (t *Primitive) CName() string  { return t.goName }
func (t *Primitive) GoName() string { return t.goName }
func (t *Primitive) ABIDesc() string {
	switch t.goName {
	case "uint8":
		return "ffi.UInt8"
	case "int8":
		return "ffi.SInt8"
	case "uint16":
		return "ffi.UInt16"
	case "int16":
		return "ffi.SInt16"
	case "uint32":
		return "ffi.UInt32"
	case "int32":
		return "ffi.SInt32"
	case "uint64":
		return "ffi.UInt64"
	case "int64":
		return "ffi.SInt64"
	case "float32":
		return "ffi.Float"
	case "float64":
		return "ffi.Double"
	default:
		return "ffi.Pointer" // size_t / uintptr
	}
}

// ---------------------------------------------------------------------------
// Bool  (VkBool32 ↔ Go bool)
// ---------------------------------------------------------------------------

type Bool struct{}

func (t *Bool) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := uint32(0)", out))
	g.Line(fmt.Sprintf("\tif %s { %s = 1 }", input, out))
	return out
}

func (t *Bool) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s != 0", out, input))
	return out
}

func (t *Bool) CName() string     { return "uint32" }
func (t *Bool) GoName() string    { return "bool" }
func (t *Bool) ABIDesc() string { return "ffi.UInt32" }

// ---------------------------------------------------------------------------
// NamedType  (enums, bitmasks referenced by name — simple cast)
// ---------------------------------------------------------------------------

type NamedType struct {
	Name      string // Go name e.g. "Format"
	CTypeName string // C name  e.g. "VkFormat"
	Is64Bit   bool   // true for VkFlags64-backed bitmask types
}

func (t *NamedType) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	if t.Is64Bit {
		g.Line(fmt.Sprintf("\t%s := uint64(%s)", out, input))
	} else {
		g.Line(fmt.Sprintf("\t%s := uint32(%s)", out, input))
	}
	return out
}

func (t *NamedType) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.Name, input))
	return out
}

func (t *NamedType) CName() string {
	if t.Is64Bit {
		return "uint64"
	}
	return "uint32"
}
func (t *NamedType) GoName() string { return t.Name }
func (t *NamedType) ABIDesc() string {
	if t.Is64Bit {
		return "ffi.UInt64"
	}
	return "ffi.UInt32"
}

// ---------------------------------------------------------------------------
// StructType  (structs referenced by name — uses toC()/fromC)
// ---------------------------------------------------------------------------

type StructType struct {
	Name      string // Go name e.g. "Offset3D"
	CTypeName string // C name  e.g. "VkOffset3D"
}

func (t *StructType) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	cancelVar := g.Var("cancel")
	castVar := g.Var("cast")
	g.Line(fmt.Sprintf("\t%s, %s := %s.toC()", out, cancelVar, input))
	g.Line(fmt.Sprintf("\tcancels = append(cancels, %s)", cancelVar))
	g.Line(fmt.Sprintf("\t%s := (*c%s)(%s)", castVar, t.CTypeName, out))
	return "*" + castVar
}

func (t *StructType) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\tvar %s %s", out, t.Name))
	g.Line(fmt.Sprintf("\t%s.fromC(&%s)", out, input))
	return out
}

func (t *StructType) CName() string     { return "c" + t.CTypeName }
func (t *StructType) GoName() string    { return t.Name }
func (t *StructType) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// Handle  (opaque Vulkan handles — represented as unsafe.Pointer in Go)
// ---------------------------------------------------------------------------

type Handle struct {
	Name      string // Go name e.g. "Instance"
	CTypeName string // C name  e.g. "VkInstance"
}

func (t *Handle) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("h")
	g.Line(fmt.Sprintf("\tvar %s unsafe.Pointer", out))
	g.Line(fmt.Sprintf("\tif %s != nil { %s = %s.handle }", input, out, input))
	return out
}

func (t *Handle) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("h")
	g.Line(fmt.Sprintf("\t%s := &%s{handle: unsafe.Pointer(%s)}", out, t.Name, input))
	return out
}

func (t *Handle) CName() string     { return "unsafe.Pointer" }
func (t *Handle) GoName() string    { return "*" + t.Name }
func (t *Handle) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// Pointer
// ---------------------------------------------------------------------------

type Pointer struct {
	Child FieldType
}

func (t *Pointer) GenerateToC(g *CodeGen, input string) string {
	ptrVar := g.Var("ptr")
	g.Line(fmt.Sprintf("\tvar %s unsafe.Pointer", ptrVar))
	g.Line(fmt.Sprintf("\tif %s != nil {", input))
	if st, ok := t.Child.(*StructType); ok {
		out := g.Var("val")
		cancelVar := g.Var("cancel")
		g.Line(fmt.Sprintf("\t\t%s, %s := %s.toC()", out, cancelVar, input))
		g.Line(fmt.Sprintf("\t\tcancels = append(cancels, %s)", cancelVar))
		g.Line(fmt.Sprintf("\t\t%s = %s", ptrVar, out))
		_ = st
	} else {
		childOut := t.Child.GenerateToC(g, "*"+input)
		innerVar := g.Var("inner")
		g.Line(fmt.Sprintf("\t\t%s := %s", innerVar, childOut))
		g.Line(fmt.Sprintf("\t\t%s = unsafe.Pointer(&%s)", ptrVar, innerVar))
	}
	g.Line("\t}")
	return ptrVar
}

func (t *Pointer) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("ptr")
	g.Line(fmt.Sprintf("\tvar %s %s", out, t.GoName()))
	g.Line(fmt.Sprintf("\tif %s != nil {", input))
	childOut := t.Child.GenerateFromC(g, "*"+input)
	g.Line(fmt.Sprintf("\t\t%s = &%s", out, childOut))
	g.Line("\t}")
	return out
}

func (t *Pointer) CName() string     { return "unsafe.Pointer" }
func (t *Pointer) GoName() string    { return "*" + t.Child.GoName() }
func (t *Pointer) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// Slice  (replaces separate count+pointer pair)
// ---------------------------------------------------------------------------

type Slice struct {
	Child FieldType
}

func (t *Slice) GenerateToC(g *CodeGen, input string) string {
	lengthVar := g.Var("len")
	arrVar := g.Var("arr")
	bufVar := g.Var("buf")
	pinVar := g.Var("pin")

	g.Line(fmt.Sprintf("\t%s := len(%s)", lengthVar, input))
	g.Line(fmt.Sprintf("\tvar %s unsafe.Pointer", arrVar))
	g.Line(fmt.Sprintf("\tif %s > 0 {", lengthVar))
	g.Line(fmt.Sprintf("\t\t%s := make([]%s, %s)", bufVar, t.Child.CName(), lengthVar))
	g.Line(fmt.Sprintf("\t\tvar %s runtime.Pinner", pinVar))
	g.Line(fmt.Sprintf("\t\t%s.Pin(&%s[0])", pinVar, bufVar))
	g.Line(fmt.Sprintf("\t\tcancels = append(cancels, func() { %s.Unpin(); runtime.KeepAlive(%s) })", pinVar, bufVar))
	g.Line(fmt.Sprintf("\t\t%s = unsafe.Pointer(&%s[0])", arrVar, bufVar))

	indexVar := g.Var("i")
	elemVar := g.Var("elem")
	g.Line(fmt.Sprintf("\t\tfor %s, %s := range %s {", indexVar, elemVar, input))
	childOut := t.Child.GenerateToC(g, elemVar)
	// For struct children, childOut is "*castVar" (dereferenced) — assign directly.
	// For all others it's a scalar value.
	g.Line(fmt.Sprintf("\t\t\t%s[%s] = %s", bufVar, indexVar, childOut))
	g.Line("\t\t}")
	g.Line("\t}")

	return arrVar
}

func (t *Slice) GenerateFromC(_ *CodeGen, _ string) string { return "nil" }
func (t *Slice) CName() string                             { return "unsafe.Pointer" }
func (t *Slice) GoName() string                            { return "[]" + t.Child.GoName() }
func (t *Slice) ABIDesc() string                         { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// FixedArray  (fixed-size C arrays like float[2])
// ---------------------------------------------------------------------------

type FixedArray struct {
	Child FieldType
	Size  int
}

func (t *FixedArray) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("arr")
	g.Line(fmt.Sprintf("\tvar %s [%d]%s", out, t.Size, t.Child.CName()))
	indexVar := g.Var("i")
	elemVar := g.Var("elem")
	g.Line(fmt.Sprintf("\tfor %s, %s := range %s {", indexVar, elemVar, input))
	childOut := t.Child.GenerateToC(g, elemVar)
	g.Line(fmt.Sprintf("\t\t%s[%s] = %s", out, indexVar, childOut))
	g.Line("\t}")
	return out
}

func (t *FixedArray) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("arr")
	g.Line(fmt.Sprintf("\tvar %s [%d]%s", out, t.Size, t.Child.GoName()))
	indexVar := g.Var("i")
	elemVar := g.Var("elem")
	g.Line(fmt.Sprintf("\tfor %s, %s := range %s {", indexVar, elemVar, input))
	childOut := t.Child.GenerateFromC(g, elemVar)
	g.Line(fmt.Sprintf("\t\t%s[%s] = %s", out, indexVar, childOut))
	g.Line("\t}")
	return out
}

func (t *FixedArray) CName() string {
	return fmt.Sprintf("[%d]%s", t.Size, t.Child.CName())
}
func (t *FixedArray) GoName() string {
	return fmt.Sprintf("[%d]%s", t.Size, t.Child.GoName())
}
func (t *FixedArray) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// PtrSlice  (double-pointer array: **T in C, []*T in Go)
// ---------------------------------------------------------------------------

type PtrSlice struct {
	Child FieldType // always *StructType in practice
}

func (t *PtrSlice) GenerateToC(g *CodeGen, input string) string {
	lengthVar := g.Var("len")
	arrVar := g.Var("arr")
	bufVar := g.Var("buf")
	pinVar := g.Var("pin")

	g.Line(fmt.Sprintf("\t%s := len(%s)", lengthVar, input))
	g.Line(fmt.Sprintf("\tvar %s unsafe.Pointer", arrVar))
	g.Line(fmt.Sprintf("\tif %s > 0 {", lengthVar))
	g.Line(fmt.Sprintf("\t\t%s := make([]unsafe.Pointer, %s)", bufVar, lengthVar))
	g.Line(fmt.Sprintf("\t\tvar %s runtime.Pinner", pinVar))
	g.Line(fmt.Sprintf("\t\t%s.Pin(&%s[0])", pinVar, bufVar))
	g.Line(fmt.Sprintf("\t\tcancels = append(cancels, func() { %s.Unpin(); runtime.KeepAlive(%s) })", pinVar, bufVar))
	g.Line(fmt.Sprintf("\t\t%s = unsafe.Pointer(&%s[0])", arrVar, bufVar))

	indexVar := g.Var("i")
	elemVar := g.Var("elem")
	g.Line(fmt.Sprintf("\t\tfor %s, %s := range %s {", indexVar, elemVar, input))
	if st, ok := t.Child.(*StructType); ok {
		outVar := g.Var("val")
		cancelVar := g.Var("cancel")
		g.Line(fmt.Sprintf("\t\t\t%s, %s := %s.toC()", outVar, cancelVar, elemVar))
		g.Line(fmt.Sprintf("\t\t\tcancels = append(cancels, %s)", cancelVar))
		g.Line(fmt.Sprintf("\t\t\t%s[%s] = %s", bufVar, indexVar, outVar))
		_ = st
	}
	g.Line("\t\t}")
	g.Line("\t}")

	return arrVar
}

func (t *PtrSlice) GenerateFromC(_ *CodeGen, _ string) string { return "nil" }
func (t *PtrSlice) CName() string                             { return "unsafe.Pointer" }
func (t *PtrSlice) GoName() string                            { return "[]*" + t.Child.GoName() }
func (t *PtrSlice) ABIDesc() string                         { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// String  (null-terminated const char* ↔ Go string)
// ---------------------------------------------------------------------------

type String struct{}

func (t *String) GenerateToC(g *CodeGen, input string) string {
	bufVar := g.Var("cstr")
	pinVar := g.Var("pin")
	ptrVar := g.Var("ptr")
	g.Line(fmt.Sprintf("\t%s := append([]byte(%s), 0)", bufVar, input))
	g.Line(fmt.Sprintf("\tvar %s runtime.Pinner", pinVar))
	g.Line(fmt.Sprintf("\t%s.Pin(&%s[0])", pinVar, bufVar))
	g.Line(fmt.Sprintf("\tcancels = append(cancels, func() { %s.Unpin(); runtime.KeepAlive(%s) })", pinVar, bufVar))
	g.Line(fmt.Sprintf("\t%s := unsafe.Pointer(&%s[0])", ptrVar, bufVar))
	return ptrVar
}

func (t *String) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("str")
	g.Line(fmt.Sprintf("\t%s := _cGoString(%s)", out, input))
	return out
}

func (t *String) CName() string     { return "unsafe.Pointer" }
func (t *String) GoName() string    { return "string" }
func (t *String) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// TypedStructure  (pNext / Next field — implements Structure interface)
// ---------------------------------------------------------------------------

type TypedStructure struct{}

func (t *TypedStructure) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("typed")
	cancelVar := g.Var("cancel")
	g.Line(fmt.Sprintf("\t%s, %s := %s.toC()", out, cancelVar, input))
	g.Line(fmt.Sprintf("\tcancels = append(cancels, %s)", cancelVar))
	return out
}

func (t *TypedStructure) GenerateFromC(_ *CodeGen, _ string) string { return "nil" }
func (t *TypedStructure) CName() string                             { return "unsafe.Pointer" }
func (t *TypedStructure) GoName() string                            { return "Structure" }
func (t *TypedStructure) ABIDesc() string                         { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// VoidPtr
// ---------------------------------------------------------------------------

type VoidPtr struct{}

func (t *VoidPtr) GenerateToC(_ *CodeGen, input string) string { return input }
func (t *VoidPtr) GenerateFromC(_ *CodeGen, input string) string {
	return "unsafe.Pointer(" + input + ")"
}
func (t *VoidPtr) CName() string     { return "unsafe.Pointer" }
func (t *VoidPtr) GoName() string    { return "unsafe.Pointer" }
func (t *VoidPtr) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// ByteSlice  (void* + paired size — represented as []byte)
// ---------------------------------------------------------------------------

type ByteSlice struct{}

func (t *ByteSlice) GenerateToC(g *CodeGen, input string) string {
	ptrVar := g.Var("ptr")
	g.Line(fmt.Sprintf("\tvar %s unsafe.Pointer", ptrVar))
	g.Line(fmt.Sprintf("\tif len(%s) > 0 { %s = unsafe.Pointer(&%s[0]) }", input, ptrVar, input))
	return ptrVar
}

func (t *ByteSlice) GenerateFromC(_ *CodeGen, _ string) string { return "nil" }
func (t *ByteSlice) CName() string                             { return "unsafe.Pointer" }
func (t *ByteSlice) GoName() string                            { return "[]byte" }
func (t *ByteSlice) ABIDesc() string                         { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// ExternalType (platform-specific C types not part of the Vulkan type system)
// ---------------------------------------------------------------------------

type ExternalType struct {
	CTypeName   string // e.g., "Display", "HWND", "DWORD"
	GoTypeName  string // Go representation: "unsafe.Pointer", "uintptr", "uint32"
	PtrInVulkan bool   // whether Vulkan uses it with * (e.g., Display* dpy)
}

func (t *ExternalType) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("ext")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.GoTypeName, input))
	return out
}

func (t *ExternalType) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("ext")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.GoTypeName, input))
	return out
}

func (t *ExternalType) CName() string {
	return t.GoTypeName
}

func (t *ExternalType) GoName() string    { return t.GoTypeName }
func (t *ExternalType) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// GoFuncPointer (Vulkan PFN_vk* function pointer types)
// ---------------------------------------------------------------------------

type FuncPointerParam struct {
	Name string
	Type FieldType
}

type GoFuncPointer struct {
	CTypeName  string
	GoTypeName string
	Return     FieldType
	Params     []FuncPointerParam
}

func (fp *GoFuncPointer) Generate() string {
	var b strings.Builder
	b.WriteString("type " + fp.GoTypeName + " func(")
	for i, p := range fp.Params {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(p.Name + " " + p.Type.GoName())
	}
	b.WriteString(")")
	if fp.Return != nil {
		b.WriteString(" " + fp.Return.GoName())
	}
	b.WriteString("\n\n")
	return b.String()
}

func (fp *GoFuncPointer) GenerateToC(g *CodeGen, input string) string {
	// Function pointer fields with callbacks are handled specially by the struct generator.
	// This fallback just returns a nil pointer.
	out := g.Var("fn")
	g.Line(fmt.Sprintf("\t%s := uintptr(0)", out))
	g.Line(fmt.Sprintf("\t_ = %s", input))
	return out
}

func (fp *GoFuncPointer) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("fn")
	g.Line(fmt.Sprintf("\tvar %s %s", out, fp.GoTypeName))
	g.Line(fmt.Sprintf("\t*(*uintptr)(unsafe.Pointer(&%s)) = %s", out, input))
	return out
}

func (fp *GoFuncPointer) CName() string     { return "uintptr" }
func (fp *GoFuncPointer) GoName() string    { return fp.GoTypeName }
func (fp *GoFuncPointer) ABIDesc() string { return "ffi.Pointer" }

// ---------------------------------------------------------------------------
// Helpers used by output generation (no longer depend on C. type names)
// ---------------------------------------------------------------------------

func alignUpInt(n, align int) int {
	if align <= 1 {
		return n
	}
	return (n + align - 1) &^ (align - 1)
}

// cTypeAlignment returns the C alignment requirement for the given type.
func cTypeAlignment(ft FieldType, structs map[string]*Structured) int {
	switch t := ft.(type) {
	case *Primitive:
		switch t.goName {
		case "uint8", "int8", "byte":
			return 1
		case "uint16", "int16":
			return 2
		case "uint32", "int32", "float32":
			return 4
		default:
			return 8
		}
	case *Bool:
		return 4
	case *NamedType:
		if t.Is64Bit {
			return 8
		}
		return 4
	case *FixedArray:
		return cTypeAlignment(t.Child, structs)
	case *StructType:
		if s2, ok := structs[t.Name]; ok {
			return structCAlignment(s2, structs)
		}
		return 8
	default:
		return 8
	}
}

// structCAlignment returns the max field alignment of a struct or union for C layout.
func structCAlignment(s *Structured, structs map[string]*Structured) int {
	maxAlign := 1
	if s.HasSType {
		maxAlign = 8 // pNext unsafe.Pointer
	}
	for _, f := range s.Fields {
		al := cTypeAlignment(f.Type, structs)
		if al > maxAlign {
			maxAlign = al
		}
	}
	return maxAlign
}

// goTypeAlignment returns the alignment Go will infer for this type when used as a struct field.
// For union StructTypes (represented as [N]byte), Go infers alignment 1, not the C alignment.
func goTypeAlignment(ft FieldType, structs map[string]*Structured) int {
	if st, ok := ft.(*StructType); ok {
		if s2, ok2 := structs[st.Name]; ok2 && s2.IsUnion {
			return 1
		}
	}
	return cTypeAlignment(ft, structs)
}

// fieldTypeLayoutSize returns the byte size of ft in a C-layout struct.
// Used to compute union sizes.
func fieldTypeLayoutSize(ft FieldType, structs map[string]*Structured) int {
	switch t := ft.(type) {
	case *Primitive:
		switch t.goName {
		case "uint8", "int8", "byte":
			return 1
		case "uint16", "int16":
			return 2
		case "uint32", "int32", "float32":
			return 4
		case "uint64", "int64", "float64":
			return 8
		}
		return 8 // uintptr / size_t
	case *Bool:
		return 4
	case *NamedType:
		if t.Is64Bit {
			return 8
		}
		return 4
	case *FixedArray:
		return t.Size * fieldTypeLayoutSize(t.Child, structs)
	case *StructType:
		if s2, ok := structs[t.Name]; ok {
			return computeStructLayoutSize(s2, structs)
		}
		return 8
	default:
		return 8 // pointers, handles, slices, strings, etc.
	}
}

func computeStructLayoutSize(s *Structured, structs map[string]*Structured) int {
	if s.IsUnion {
		return computeUnionLayoutSize(s, structs)
	}
	total := 0
	maxAlign := 1
	for _, f := range s.Fields {
		sz := fieldTypeLayoutSize(f.Type, structs)
		al := sz
		if al > 8 {
			al = 8
		}
		if al > maxAlign {
			maxAlign = al
		}
		total = (total + al - 1) &^ (al - 1)
		total += sz
	}
	// sType (int32) + pNext (unsafe.Pointer) for sType-bearing structs
	if s.HasSType {
		total += 4 + 8 // conservative; actual layout computed per-struct
	}
	return (total + maxAlign - 1) &^ (maxAlign - 1)
}

func computeUnionLayoutSize(s *Structured, structs map[string]*Structured) int {
	max := 1
	for _, f := range s.Fields {
		sz := fieldTypeLayoutSize(f.Type, structs)
		if sz > max {
			max = sz
		}
	}
	return max
}

// abiReturnDesc returns the goffi TypeDescriptor for a command's C return type.
func abiReturnDesc(cReturnType string) string {
	switch cReturnType {
	case "void", "":
		return "ffi.Void"
	case "VkResult":
		return "ffi.SInt32"
	default:
		return "ffi.Pointer"
	}
}

// cParamABIDesc maps a C param type string to a goffi TypeDescriptor expression.
func cParamABIDesc(cType string) string {
	s := strings.TrimSpace(cType)
	// Any pointer type → pointer descriptor
	if strings.Contains(s, "*") {
		return "ffi.Pointer"
	}
	s = strings.TrimPrefix(s, "const ")
	s = strings.TrimSpace(s)
	switch s {
	case "uint8_t":
		return "ffi.UInt8"
	case "uint16_t":
		return "ffi.UInt16"
	case "uint32_t", "VkFlags", "VkBool32", "VkSampleMask":
		return "ffi.UInt32"
	case "uint64_t", "VkDeviceSize", "VkDeviceAddress", "VkFlags64":
		return "ffi.UInt64"
	case "int8_t":
		return "ffi.SInt8"
	case "int16_t":
		return "ffi.SInt16"
	case "int32_t", "int", "VkResult":
		return "ffi.SInt32"
	case "int64_t":
		return "ffi.SInt64"
	case "float":
		return "ffi.Float"
	case "double":
		return "ffi.Double"
	case "size_t":
		return "ffi.Pointer"
	default:
		// Handles (VkInstance, VkDevice, ...) and any remaining types → pointer-sized
		return "ffi.Pointer"
	}
}
