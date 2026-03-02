package generator

import "fmt"

// FieldType generates C conversion code.
// It returns the name of the generated C variable.
type FieldType interface {
	GenerateToC(g *CodeGen, input string) string
	CName() string
	GoName() string
}

type Array struct {
	child FieldType
}

func (t *Array) GenerateToC(g *CodeGen, input string) string {
	lengthVar := g.Var("len")
	ptrVar := g.Var("arr")

	g.Line(fmt.Sprintf("\t%s := len(%s)", lengthVar, input))

	g.Line(fmt.Sprintf(`
	var %s *%s
	if %s > 0 {
		%s = (*%s)(C.malloc(C.size_t(%s) * C.size_t(unsafe.Sizeof(*new(%s)))))
	}
	`,
		ptrVar,
		t.child.CName(),
		lengthVar,
		ptrVar,
		t.child.CName(),
		lengthVar,
		t.child.CName(),
	))

	indexVar := g.Var("i")
	elemVar := g.Var("elem")

	g.Line(fmt.Sprintf("for %s, %s := range %s {",
		indexVar, elemVar, input,
	))

	childOut := t.child.GenerateToC(g, elemVar)

	g.Line(fmt.Sprintf(
		"\t(*[1<<30]%s)(unsafe.Pointer(%s))[%s] = %s",
		t.child.CName(),
		ptrVar,
		indexVar,
		childOut,
	))

	g.Line("}")

	return ptrVar
}

func (t *Array) CName() string {
	return "*" + t.child.CName()
}

func (t *Array) GoName() string {
	return "*" + t.child.CName()
}

type Pointer struct {
	child FieldType
}

func (t *Pointer) GenerateToC(g *CodeGen, input string) string {
	ptrVar := g.Var("ptr")

	g.Line(fmt.Sprintf("var %s %s", ptrVar, t.CName()))
	g.Line(fmt.Sprintf("if %s != nil {", input))

	childOut := t.child.GenerateToC(g, "*"+input)

	g.Line(fmt.Sprintf("\t%s = &%s", ptrVar, childOut))
	g.Line("}")

	return ptrVar
}

func (t *Pointer) CName() string {
	return "*" + t.child.CName()
}

func (t *Pointer) GoName() string {
	return "*" + t.child.CName()
}

type String struct{}

func (t *String) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("cstr")

	g.Line(fmt.Sprintf("%s := C.CString(%s)", out, input))
	g.Line(fmt.Sprintf(
		"cancels = append(cancels, func(){ C.free(unsafe.Pointer(%s)) })",
		out,
	))

	return out
}

func (t *String) CName() string {
	return "*C.char"
}

func (t *String) GoName() string {
	return "string"
}

type TypedStructure struct{}

func (t *TypedStructure) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("typed")
	cancelVar := g.Var("cancel")

	g.Line(fmt.Sprintf("%s, %s := %s.toC()", out, cancelVar, input))
	g.Line(fmt.Sprintf("cancels = append(cancels, %s)", cancelVar))

	return out
}

func (t *TypedStructure) CName() string {
	return "unsafe.Pointer"
}

func (t *TypedStructure) GoName() string {
	return "Structure"
}

type Primitive struct {
	cName  string
	goName string
}

func (t *Primitive) GenerateToC(g *CodeGen, input string) string {
	out := g.Var("val")
	g.Line(fmt.Sprintf("%s := (C.%s)(%s)", out, t.cName, input))
	return out
}

func (t *Primitive) CName() string {
	return "C." + t.cName
}

func (t *Primitive) GoName() string {
	return t.goName
}
