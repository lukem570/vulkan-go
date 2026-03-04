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
	GoName string
	CName  string
	Count  string // optional count field for arrays
	Type   FieldType
}

type Structured struct {
	CName  string
	GoName string
	Fields []StructField

	// Only for structs that have sType / pNext
	HasSType bool
	SType    string // e.g. StructureTypeInstanceCreateInfo
}

func (s *Structured) GenerateToC() string {
	if s == nil {
		return ""
	}

	g := &CodeGen{}
	g.Line(fmt.Sprintf("func (s *%s) toC() (*C.%s, func()) {", s.GoName, s.CName))
	g.Line("\tcancels := make([]func(), 0)")
	g.Line(fmt.Sprintf("\tp := (*C.%s)(C.malloc(C.size_t(C.sizeof_%s)))", s.CName, s.CName))

	if s.HasSType {
		g.Line("\tp.sType = (C.VkStructureType)(s.GetType())")
		g.Line(`
	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}
	`)
	}

	for _, field := range s.Fields {
		input := "s." + field.GoName
		output := field.Type.GenerateToC(g, input)
		g.Line(fmt.Sprintf("\tp.%s = %s", field.CName, output))
		if field.Count != "" {
			g.Line(fmt.Sprintf("\tp.%s = C.uint32_t(len(%s))", field.Count, input))
		}
	}

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
	if s == nil {
		return ""
	}

	if !s.HasSType {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("func (s *%s) GetType() StructureType {\n", s.GoName))
	b.WriteString(fmt.Sprintf("\treturn %s\n", s.SType))
	b.WriteString("}\n\n")
	return b.String()
}

func (s *Structured) GenerateStructure() string {
	if s == nil {
		return ""
	}

	var b strings.Builder

	b.WriteString(fmt.Sprintf("type %s struct {\n", s.GoName))

	if s.HasSType {
		b.WriteString("\tNext Structure\n")
	}

	for _, field := range s.Fields {
		goType := field.Type.GoName()
		goType = strings.TrimPrefix(goType, "C.") // convert C type to Go type
		b.WriteString(fmt.Sprintf("\t%s %s\n", field.GoName, goType))
	}

	b.WriteString("}\n\n")
	return b.String()
}
