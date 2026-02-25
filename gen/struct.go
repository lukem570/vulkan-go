package main

import (
	"fmt"
	"strings"
)

type Struct struct {
	Name string
	Elements []CVariable
}

func (s *Struct) Generate() string {
	var structBuilder strings.Builder
	structBuilder.WriteRune('\n')

	structBuilder.WriteString(fmt.Sprintf(
		"type %s struct {\n",
		s.Name,
	))

	for _, element := range s.Elements {
		structBuilder.WriteString(fmt.Sprintf(
			"\t%s %s\n",
			camelToPascal(element.Name),
			element.Type.Generate(),
		))
	}

	structBuilder.WriteString("}\n")

	return structBuilder.String()
}

type XmlStruct struct {
	Name    string  `xml:"name,attr"`
	Members []RawXml `xml:"member"`
}

func (s *XmlStruct) Parse() *Struct {
	out := &Struct{
		Name: s.Name,
		Elements: make([]CVariable, 0),
	}

	for _, member := range s.Members {
		out.Elements = append(out.Elements, ParseCVariable(member.Value))
	}

	return out
}