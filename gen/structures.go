package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

var goTypeMap = map[string]string{
	"uint8_t":  "uint8",
	"uint16_t": "uint16",
	"uint32_t": "uint32",
	"uint64_t": "uint64",

	"int8_t":  "int8",
	"int16_t": "int16",
	"int32_t": "int32",
	"int64_t": "int64",

	"size_t": "uint",

	"float": "float32",
	"double": "float64",
}

type Struct struct {
	Members []StructMember `xml:"member"`
}

type StructMember struct {
	Type     string `xml:"type"`                // maps to <type>
	Name     string `xml:"name"`                // maps to <name>
	Values   string `xml:"values,attr"`         // optional attribute
	Optional string `xml:"optional,attr"`       // optional attribute
	NoAuto   string `xml:"noautovalidity,attr"` // optional attribute
	Len      string `xml:"len,attr"`            // optional attribute
}

func (p *VkParser) parseStruct(name string, innerXml []byte) error {
	innerXml = []byte("<root>" + string(innerXml) + "</root>")

	var s Struct
	err := xml.Unmarshal(innerXml, &s)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	var members strings.Builder
	for _, member := range s.Members {
		conv, ok := goTypeMap[member.Type]
		if !ok {
			conv = member.Type
		}

		goMember := fmt.Sprintf("\t%s %s\n", caseMember(member.Name), conv)
		members.WriteString(goMember)
	}

	goStruct := fmt.Sprintf("type %s struct {\n%s}\n\n", name, members.String())
	p.structures.WriteString(goStruct)

	return nil
}
