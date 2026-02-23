package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Function struct {
	Proto  FunctionMember   `xml:"proto"`
	Params []FunctionMember `xml:"param"`
}

type FunctionMember struct {
    Type     string `xml:"type"`
    Name     string `xml:"name"`
    InnerXML string `xml:",innerxml"`
}

func (p *VkParser) parseFuncPointer(innerXml []byte) error {
	innerXml = []byte("<root>" + string(innerXml) + "</root>")

	var f Function
	err := xml.Unmarshal(innerXml, &f)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Failed")
			return nil
		}
		return err
	}

	var params strings.Builder
	for _, param := range f.Params {

		cType := extractCType(param)
		goType := convertCToGoType(cType)

		if goType == "" {
			continue
		}

		goParam := fmt.Sprintf("\t%s %s,\n", param.Name, goType)
		params.WriteString(goParam)
	}

	cType := extractCType(f.Proto)
	goType := convertCToGoType(cType)

	goFunction := fmt.Sprintf(
		"func %s (\n%s) %s {\n}\n\n",
		f.Proto.Name,
		params.String(),
		goType,
	)
	p.functions.WriteString(goFunction)

	return nil
}
