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
	g.Line(fmt.Sprintf("func (s *%s) toC() (*C.%s, func()) {", s.GoName, s.CName))
	g.Line("\tcancels := make([]func(), 0)")
	g.Line(fmt.Sprintf("\tp := (*C.%s)(C.malloc(C.size_t(C.sizeof_%s)))", s.CName, s.CName))
	g.Line(fmt.Sprintf("\t*p = C.%s{}", s.CName))

	if s.HasSType {
		g.Line("\tp.sType = C.VkStructureType(s.GetType())")
		g.Line("\tif s.Next != nil {")
		g.Line("\t\tnextPtr, nextCancel := s.Next.toC()")
		g.Line("\t\tcancels = append(cancels, nextCancel)")
		g.Line("\t\tp.pNext = unsafe.Pointer(nextPtr)")
		g.Line("\t}")
	}

	for _, field := range s.Fields {
		if field.CountFor != "" {
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

	g.Line("\treturn p, func() {")
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