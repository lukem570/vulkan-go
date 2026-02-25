package main

import (
	"fmt"
	"strings"
)

type FuncPointer struct {
	Proto  CVariable
	Params []CVariable
}

func (f *FuncPointer) Generate() string {
	var funcPtr strings.Builder
	funcPtr.WriteRune('\n')

	funcPtr.WriteString(fmt.Sprintf(
		"type %s func(\n",
		f.Proto.Name,
	))

	for _, param := range f.Params {
		funcPtr.WriteString(fmt.Sprintf(
			"\t%s %s,\n",
			param.Name,
			param.Type.Generate(),
		))
	}

	funcPtr.WriteString(fmt.Sprintf(
		")%s\n",
		f.Proto.Type.Generate(),
	))

	return funcPtr.String()
}

type XmlFuncPointer struct {
	Proto  RawXml   `xml:"proto"`
	Params []RawXml `xml:"param"`
}

func (f *XmlFuncPointer) Parse() FuncPointer {
	var out FuncPointer

	out.Proto = ParseCVariable(f.Proto.Value)
	out.Params = make([]CVariable, 0)

	for _, param := range f.Params {
		out.Params = append(out.Params, ParseCVariable(param.Value))
	}

	return out
}
