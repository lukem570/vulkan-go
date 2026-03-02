package main

import (
	"fmt"
	"strings"
)

type StructElement struct {
	Var CVariable
	Value string
}

type Struct struct {
	Name string
	Elements []StructElement
}

func (s *Struct) Generate() string {
	var structBuilder strings.Builder
	structBuilder.WriteRune('\n')

	structBuilder.WriteString(fmt.Sprintf(
		"type %s struct {\n",
		s.Name,
	))

	seen := make(map[string]struct{}, 0)

	var skipNext bool
	for i, element := range s.Elements {
		if _, ok := seen[element.Var.Name]; ok {
			continue
		}
		seen[element.Var.Name] = struct{}{}
		if skipNext {
			skipNext = false
			continue
		}

		if strings.HasSuffix(element.Var.Name, "Count") {
			if i+1 < len(s.Elements) && strings.HasPrefix(s.Elements[i+1].Var.Name, "p") {
				skipNext = true
				ele := s.Elements[i+1]
				typ := ele.Var.Type
				typ.PtrDepth--
				
				structBuilder.WriteString(fmt.Sprintf(
					"\t%s []%s\n",
					camelToPascal(ele.Var.Name[1:]),
					typ.Generate(),
				))

				continue
			}
		}

		structBuilder.WriteString(fmt.Sprintf(
			"\t%s %s\n",
			camelToPascal(element.Var.Name),
			element.Var.Type.Generate(),
		))
	}

	structBuilder.WriteString("}\n")

	return structBuilder.String()
}

type XmlStructElement struct {
	Value string `xml:"values"`
	Type string `xml:",innerxml"`
}

type XmlStruct struct {
	Name    string  `xml:"name,attr"`
	Members []XmlStructElement `xml:"member"`
}

func (s *XmlStruct) Parse() *Struct {
	out := &Struct{
		Name: s.Name,
		Elements: make([]StructElement, 0),
	}

	for _, member := range s.Members {
		element := StructElement{
			Value: member.Value,
			Var: ParseCVariable(member.Type),
		}

		out.Elements = append(out.Elements, element)
	}

	return out
}