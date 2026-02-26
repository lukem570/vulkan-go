package main

import (
	"fmt"
	"strings"
)

type Command struct {
	Proto  CVariable
	Params []CVariable
}

func (f *Command) Generate(handles map[string]*Handle) string {
	var cmd strings.Builder
	cmd.WriteRune('\n')

	funcName := f.Proto.Name
	funcName = camelToPascal(funcName)
	funcName, _ = strings.CutPrefix(funcName, "Vk")

	Params := f.Params

	if len(Params) > 0 {
		if _, ok := handles[Params[0].Type.BaseType]; ok {
			typ := Params[0].Type.Generate()
			typN, _ := strings.CutPrefix(typ, "Vk")
			
			funcName = strings.ReplaceAll(funcName, typN, "")

			if typN == "CommandBuffer" {
				funcName = strings.ReplaceAll(funcName, "Cmd", "")
			}
			
			cmd.WriteString(fmt.Sprintf(
				"func (h *%s) %s(\n",
				typ,
				funcName,
			))

			Params = Params[1:]
		} else {
			cmd.WriteString(fmt.Sprintf(
				"func %s(\n",
				funcName,
			))
		}
	} else {
		cmd.WriteString(fmt.Sprintf(
			"func %s(\n",
			funcName,
		))
	}

	for _, param := range Params {
		name := param.Name
		if name == "type" {
			name = "Type"
		}

		cmd.WriteString(fmt.Sprintf(
			"\t%s %s,\n",
			name,
			param.Type.Generate(),
		))
	}

	cmd.WriteString(fmt.Sprintf(
		") %s {\n}\n",
		f.Proto.Type.Generate(),
	))

	return cmd.String()
}

type XmlCommand struct {
	Proto  RawXml   `xml:"proto"`
	Params []RawXml `xml:"param"`
}

func (f *XmlCommand) Parse() *Command {
	out := &Command{
		Proto:  ParseCVariable(f.Proto.Value),
		Params: make([]CVariable, 0),
	}

	for _, param := range f.Params {
		out.Params = append(out.Params, ParseCVariable(param.Value))
	}

	return out
}