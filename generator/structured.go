package generator

import (
	"fmt"
	"strings"
)

type CodeGen struct {
	builder  strings.Builder
	varIndex int
}

func (g *CodeGen) Var(prefix string) string {
	name := fmt.Sprintf("%s_%d", prefix, g.varIndex)
	g.varIndex++
	return name
}

func (g *CodeGen) Line(s string) {
	g.builder.WriteString(s)
	g.builder.WriteByte('\n')
}

func (g *CodeGen) String() string {
	return g.builder.String()
}

type StructField struct {
	goName string
	cName  string
	count  string // optional count field for arrays
	typ    FieldType
}

type Structured struct {
	sType  string
	goName string
	cName  string
	fields []StructField
}

func (s *Structured) GenerateToC() string {
	g := &CodeGen{}

	g.Line(fmt.Sprintf(
		"func (s *%s) toC() (*C.%s, func()) {",
		s.goName,
		s.cName,
	))

	g.Line("\tcancels := make([]func(), 0)")

	// Allocate struct
	g.Line(fmt.Sprintf(
		"\tp := (*C.%s)(C.malloc(C.size_t(C.sizeof_%s)))",
		s.cName,
		s.cName,
	))

	// sType
	g.Line("\tp.sType = (C.VkStructureType)(s.GetType())")

	// pNext
	g.Line(`
	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}
	`)

	// Fields
	for _, field := range s.fields {
		input := "s." + field.goName
		output := field.typ.GenerateToC(g, input)

		g.Line(fmt.Sprintf("\tp.%s = %s", field.cName, output))

		if field.count != "" {
			g.Line(fmt.Sprintf("\tp.%s = C.uint32_t(len(%s))", field.count, input))
		}
	}

	// Cleanup
	g.Line(`
	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}
`)

	return g.String()
}

func (s *Structured) GenerateGetType() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(
		"func (s *%s) GetType() StructureType {\n",
		s.goName,
	))

	b.WriteString(fmt.Sprintf(
		"\treturn %s\n",
		s.sType,
	))

	b.WriteString("}\n\n")

	return b.String()
}

func (s *Structured) GenerateStructure() string {
	var b strings.Builder

	// Struct header
	b.WriteString(fmt.Sprintf("type %s struct {\n", s.goName))

	// Safe pNext replacement
	b.WriteString("\tNext Structure\n")

	// Fields
	for _, field := range s.fields {
		goType := field.typ.CName()

		// For Go struct definition we want the Go type,
		// not the C type name. So we strip C. prefix if present.
		goType = strings.TrimPrefix(goType, "C.")

		b.WriteString(fmt.Sprintf(
			"\t%s %s\n",
			field.goName,
			goType,
		))
	}

	b.WriteString("}\n\n")

	return b.String()
}
