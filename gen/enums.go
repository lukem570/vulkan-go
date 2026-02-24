package main

import (
	"fmt"
	"strings"
)

type EnumType int

const (
	EnumDefault EnumType = iota
	EnumBitmask
	EnumConst
)

type EnumElement struct {
	Name  string 
	Value string 
}

type Enum struct {
	Type     EnumType
	Name     string
	Elements []EnumElement
}

func (e *Enum) Generate() string {
	var enumBuilder strings.Builder
	enumBuilder.WriteRune('\n')

	if e.Type != EnumConst {
		enumBuilder.WriteString(fmt.Sprintf(
			"type %s int64\n\n",
			e.Name,
		))
	}

	enumBuilder.WriteString("var (\n")

	for _, element := range e.Elements {
		var value string
		name := e.Name + " "
		switch e.Type {
		case EnumConst:
			name = ""

			value = strings.ReplaceAll(element.Value, "U", "")
			value = strings.ReplaceAll(value, "F", "")
			value = strings.ReplaceAll(value, "f", "")
			value = strings.ReplaceAll(value, "LL", "")
			value = strings.ReplaceAll(value, "(", "")
			value = strings.ReplaceAll(value, ")", "")
			value = strings.ReplaceAll(value, "~", "^")

		case EnumDefault:
			value = element.Value
		case EnumBitmask:
			value = fmt.Sprintf("(1 << %s)", element.Value)
		}

		enumBuilder.WriteString(fmt.Sprintf(
			"\t%s %s= %s\n",
			upperToPascal(element.Name),
			name,
			value,
		))
	}

	enumBuilder.WriteString(")\n")

	return enumBuilder.String()
}

type XmlEnumElement struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	Bitpos string `xml:"bitpos,attr"`
}

type XmlEnum struct {
	Name     string        `xml:"name,attr"`
	Type     string        `xml:"type,attr"`
	Elements []XmlEnumElement `xml:"enum"`
}

func (e *XmlEnum) Parse() *Enum {
	out := &Enum{
		Name:     e.Name,
		Elements: make([]EnumElement, 0),
	}

	switch e.Type {
	case "constants":
		out.Type = EnumConst
	case "bitmask":
		out.Type = EnumBitmask
	case "enum":
		out.Type = EnumDefault
	}

	for _, element := range e.Elements {
		var ee EnumElement
		ee.Name = element.Name

		if out.Type == EnumBitmask {
			ee.Value = element.Bitpos
		} else {
			ee.Value = element.Value
		}

		if ee.Value == "" {
			ee.Value = "0"
		}

		out.Elements = append(out.Elements, ee)
	}

	return out
}
