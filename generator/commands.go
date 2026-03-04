package generator

import (
	"fmt"
	"strings"
)

type GoCommand struct {
	Name  string

	ReceiverType string // e.g. Instance

	HasError   bool
	ReturnType FieldType

	Params []CommandParam

	// CCommand or something here
}

type CommandParam struct {
	Name string
	Type FieldType
}

func (c *GoCommand) GenerateWrapper() string {
	var b strings.Builder

	if c.ReceiverType != "" {
		b.WriteString(fmt.Sprintf(
			"func (h *%s) %s(\n",
			c.ReceiverType,
			c.Name,
		))
	} else {
		b.WriteString(fmt.Sprintf(
			"func %s(\n",
			c.Name,
		))
	}

	for _, p := range c.Params {
		b.WriteString(fmt.Sprintf(
			"\t%s %s,\n",
			p.Name,
			p.Type.GoName(),
		))
	}

	b.WriteString(")")

	if c.ReturnType != nil {
		if c.HasError {
			b.WriteString(fmt.Sprintf(
				" (%s, error)",
				c.ReturnType.GoName(),
			))
		} else {
			b.WriteString(" " + c.ReturnType.GoName())
		}
	} else if c.HasError {
		b.WriteString(" error")
	}

	b.WriteString(" {\n")

	// TODO:

	b.WriteString("}\n")

	return b.String()
}
