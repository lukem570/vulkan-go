package main

import (
	"fmt"
	"strings"
)

type Enumeration struct {
	Name     string        `xml:"name,attr"`
	Type     string        `xml:"type,attr"`
	Elements []EnumElement `xml:"enum"`
}

type EnumElement struct {
	Value  string `xml:"value,attr"`
	Bitpos int    `xml:"bitpos,attr"`
	Name   string `xml:"name,attr"`
}

func (p *VkParser) parseEnum(enum Enumeration) error {

	if enum.Type != "constants" {
		p.enums.WriteString(fmt.Sprintf(
			"type %s int64\n\n",
			enum.Name,
		))
	}

	p.enums.WriteString("var (\n")

	switch enum.Type {
	case "enum":
		p.parseBasicEnum(enum)
	case "bitmask":
		p.parseBitmask(enum)
	case "constants":
		p.parseConstants(enum)
	default:
		fmt.Println("Unhandled enum:", enum.Type)
	}

	p.enums.WriteString(")\n\n")

	return nil
}

func (p *VkParser) parseBasicEnum(enum Enumeration) {
	for _, e := range enum.Elements {

		if e.Value == "" {
			continue
		}

		p.enums.WriteString(fmt.Sprintf(
			"\t%s %s = %s\n",
			caseUpper(e.Name),
			enum.Name,
			e.Value,
		))
	}
}

func (p *VkParser) parseBitmask(enum Enumeration) {
	for _, e := range enum.Elements {

		p.enums.WriteString(fmt.Sprintf(
			"\t%s %s = %d\n",
			caseUpper(e.Name),
			enum.Name,
			1<<e.Bitpos,
		))
	}
}

func (p *VkParser) parseConstants(enum Enumeration) {
	for _, e := range enum.Elements {

		val := strings.ReplaceAll(e.Value, "U", "")
		val = strings.ReplaceAll(val, "F", "")
		val = strings.ReplaceAll(val, "f", "")
		val = strings.ReplaceAll(val, "LL", "")
		val = strings.ReplaceAll(val, "(", "")
		val = strings.ReplaceAll(val, ")", "")
		val = strings.ReplaceAll(val, "~", "^")

		p.enums.WriteString(fmt.Sprintf(
			"\t%s = %s\n",
			caseUpper(e.Name),
			val,
		))
	}
}
