package generator

import (
	"fmt"
	"strings"
)

// FieldType generates C conversion code and describes a type.
type FieldType interface {
	GenerateToC(g *CodeGen, input string) string
	GenerateFromC(g *CodeGen, input string) string
	CName() string
	GoName() string
}

// ---------------------------------------------------------------------------
// Primitive
// ---------------------------------------------------------------------------

type Primitive struct {
	cName  string
	goName string
}

// NewPrimitive creates a Primitive type mapping from a C name to a Go name.
func NewPrimitive(cName, goName string) *Primitive {
	return &Primitive{cName: cName, goName: goName}
}

func (t *Primitive) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := C.%s(%s)", out, t.cName, input))
	return out
}

func (t *Primitive) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.goName, input))
	return out
}

func (t *Primitive) CName() string { return "C." + t.cName }
func (t *Primitive) GoName() string { return t.goName }

// ---------------------------------------------------------------------------
// Bool  (VkBool32 ↔ Go bool)
// ---------------------------------------------------------------------------

type Bool struct{}

func (t *Bool) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := C.VkBool32(0)", out))
	g.Line(fmt.Sprintf("\tif %s { %s = C.VkBool32(1) }", input, out))
	return out
}

func (t *Bool) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s != 0", out, input))
	return out
}

func (t *Bool) CName() string  { return "C.VkBool32" }
func (t *Bool) GoName() string { return "bool" }

// ---------------------------------------------------------------------------
// NamedType  (enums, bitmasks referenced by name — simple cast)
// ---------------------------------------------------------------------------

type NamedType struct {
	Name      string // Go name e.g. "Format"
	CTypeName string // C name  e.g. "VkFormat"
}

func (t *NamedType) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := C.%s(%s)", out, t.CTypeName, input))
	return out
}

func (t *NamedType) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\t%s := %s(%s)", out, t.Name, input))
	return out
}

func (t *NamedType) CName() string  { return "C." + t.CTypeName }
func (t *NamedType) GoName() string { return t.Name }

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
	g.Line(fmt.Sprintf("\t%s, %s := %s.toC()", out, cancelVar, input))
	g.Line(fmt.Sprintf("\tcancels = append(cancels, %s)", cancelVar))
	return "*" + out
}

func (t *StructType) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("\tvar %s %s", out, t.Name))
	g.Line(fmt.Sprintf("\t%s.fromC(&%s)", out, input))
	return out
}

func (t *StructType) CName() string  { return "C." + t.CTypeName }
func (t *StructType) GoName() string { return t.Name }

// ---------------------------------------------------------------------------
// Handle  (opaque Vulkan handles — represented as uintptr in Go)
// ---------------------------------------------------------------------------

type Handle struct {
	Name      string // Go name e.g. "Instance"
	CTypeName string // C name  e.g. "VkInstance"
}

func (t *Handle) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("h")
	g.Line(fmt.Sprintf("\t%s := C.%s(unsafe.Pointer(%s.handle))", out, t.CTypeName, input))
	return out
}

func (t *Handle) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("h")
	g.Line(fmt.Sprintf("\t%s := &%s{handle: unsafe.Pointer(%s)}", out, t.Name, input))
	return out
}

func (t *Handle) CName() string  { return "C." + t.CTypeName }
func (t *Handle) GoName() string { return "*" + t.Name }

// ---------------------------------------------------------------------------
// Pointer
// ---------------------------------------------------------------------------

type Pointer struct {
	Child FieldType
}

func (t *Pointer) GenerateToC(g *CodeGen, input string) string {
	ptrVar := g.Var("ptr")
	g.Line(fmt.Sprintf("\tvar %s %s", ptrVar, t.CName()))
	g.Line(fmt.Sprintf("\tif %s != nil {", input))
	// StructType.toC() takes a pointer receiver and returns a *C.VkFoo directly,
	// so pass input as-is and use the result without wrapping in &.
	if _, ok := t.Child.(*StructType); ok {
		out := g.Var("val")
		cancelVar := g.Var("cancel")
		g.Line(fmt.Sprintf("\t\t%s, %s := %s.toC()", out, cancelVar, input))
		g.Line(fmt.Sprintf("\t\tcancels = append(cancels, %s)", cancelVar))
		g.Line(fmt.Sprintf("\t\t%s = %s", ptrVar, out))
	} else {
		childOut := t.Child.GenerateToC(g, "*"+input)
		g.Line(fmt.Sprintf("\t\t%s = &%s", ptrVar, childOut))
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

func (t *Pointer) CName() string  { return "*" + t.Child.CName() }
func (t *Pointer) GoName() string { return "*" + t.Child.GoName() }

// ---------------------------------------------------------------------------
// Slice  (replaces separate count+pointer pair)
// ---------------------------------------------------------------------------

type Slice struct {
	Child FieldType
}

func (t *Slice) GenerateToC(g *CodeGen, input string) string {
	lengthVar := g.Var("len")
	ptrVar := g.Var("arr")

	g.Line(fmt.Sprintf("\t%s := len(%s)", lengthVar, input))
	g.Line(fmt.Sprintf(`
	var %s *%s
	if %s > 0 {
		%s = (*%s)(C.malloc(C.size_t(%s) * C.size_t(unsafe.Sizeof(*new(%s)))))
		cancels = append(cancels, func() { C.free(unsafe.Pointer(%s)) })
	}`, ptrVar, t.Child.CName(),
		lengthVar,
		ptrVar, t.Child.CName(), lengthVar, t.Child.CName(),
		ptrVar))

	indexVar := g.Var("i")
	elemVar := g.Var("elem")
	g.Line(fmt.Sprintf("\tfor %s, %s := range %s {", indexVar, elemVar, input))
	childOut := t.Child.GenerateToC(g, elemVar)
	g.Line(fmt.Sprintf("\t\t(*[1<<30]%s)(unsafe.Pointer(%s))[%s] = %s",
		t.Child.CName(), ptrVar, indexVar, childOut))
	g.Line("\t}")

	return ptrVar
}

func (t *Slice) GenerateFromC(_ *CodeGen, _ string) string { return "nil" }
func (t *Slice) CName() string                             { return "*" + t.Child.CName() }
func (t *Slice) GoName() string                            { return "[]" + t.Child.GoName() }

// ---------------------------------------------------------------------------
// FixedArray  (fixed-size C arrays like VkOffset3D[2])
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

func (t *FixedArray) CName() string  { return fmt.Sprintf("[%d]%s", t.Size, t.Child.CName()) }
func (t *FixedArray) GoName() string { return fmt.Sprintf("[%d]%s", t.Size, t.Child.GoName()) }

// ---------------------------------------------------------------------------
// String
// ---------------------------------------------------------------------------

type String struct{}

func (t *String) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("cstr")
	g.Line(fmt.Sprintf("\t%s := C.CString(%s)", out, input))
	g.Line(fmt.Sprintf("\tcancels = append(cancels, func() { C.free(unsafe.Pointer(%s)) })", out))
	return out
}

func (t *String) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("str")
	g.Line(fmt.Sprintf("\t%s := C.GoString(%s)", out, input))
	return out
}

func (t *String) CName() string  { return "*C.char" }
func (t *String) GoName() string { return "string" }

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

// ---------------------------------------------------------------------------
// VoidPtr
// ---------------------------------------------------------------------------

type VoidPtr struct{}

func (t *VoidPtr) GenerateToC(_ *CodeGen, input string) string { return input }
func (t *VoidPtr) GenerateFromC(_ *CodeGen, input string) string { return input }
func (t *VoidPtr) CName() string  { return "unsafe.Pointer" }
func (t *VoidPtr) GoName() string { return "unsafe.Pointer" }

// ---------------------------------------------------------------------------
// GoFuncPointer (Vulkan PFN_vk* function pointer types)
// ---------------------------------------------------------------------------

// FuncPointerParam is a single parameter in a funcpointer declaration.
type FuncPointerParam struct {
	Name string
	Type FieldType
}

// GoFuncPointer represents a Vulkan PFN_vk* function pointer type.
// Generate() emits a proper Go func type with the full signature.
type GoFuncPointer struct {
	CTypeName  string // e.g. PFN_vkAllocationFunction
	GoTypeName string // e.g. AllocationFunction
	Return     FieldType          // nil means void return
	Params     []FuncPointerParam // may be empty until resolved
}

// Generate emits: type AllocationFunction func(pUserData unsafe.Pointer, ...) unsafe.Pointer
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

// GenerateToC converts a Go func value to the C function pointer type via unsafe.
// In practice, callers should assign C function pointers via:
//
//	*(*unsafe.Pointer)(unsafe.Pointer(&field)) = unsafe.Pointer(C.myFunc)
func (fp *GoFuncPointer) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("fn")
	g.Line(fmt.Sprintf("\tvar %s C.%s", out, fp.CTypeName))
	g.Line(fmt.Sprintf("\tif %s != nil {", input))
	g.Line(fmt.Sprintf("\t\t%s = *(*C.%s)(unsafe.Pointer(&%s))", out, fp.CTypeName, input))
	g.Line("\t}")
	return out
}

func (fp *GoFuncPointer) GenerateFromC(g *CodeGen, input string) string {
	out := g.Var("fn")
	g.Line(fmt.Sprintf("\tvar %s %s", out, fp.GoTypeName))
	g.Line(fmt.Sprintf("\t*(*unsafe.Pointer)(unsafe.Pointer(&%s)) = *(*unsafe.Pointer)(unsafe.Pointer(&%s))", out, input))
	return out
}

func (fp *GoFuncPointer) CName() string  { return "C." + fp.CTypeName }
func (fp *GoFuncPointer) GoName() string { return fp.GoTypeName }