package generator

import (
	"fmt"
	"strings"
)

type GoCommand struct {
	name  string

	receiverType string // e.g. Instance

	hasError   bool
	returnType FieldType

	params []CommandParam

	// CCommand or something here
}

type CommandParam struct {
	name string
	typ FieldType
}

func (c *GoCommand) GenerateWrapper() string {
	var b strings.Builder

	if c.receiverType != "" {
		b.WriteString(fmt.Sprintf(
			"func (h *%s) %s(\n",
			c.receiverType,
			c.name,
		))
	} else {
		b.WriteString(fmt.Sprintf(
			"func %s(\n",
			c.name,
		))
	}

	for _, p := range c.params {
		b.WriteString(fmt.Sprintf(
			"\t%s %s,\n",
			p.name,
			p.typ.GoName(),
		))
	}

	b.WriteString(")")

	if c.returnType != nil {
		if c.hasError {
			b.WriteString(fmt.Sprintf(
				" (%s, error)",
				c.returnType.GoName(),
			))
		} else {
			b.WriteString(" " + c.returnType.GoName())
		}
	} else if c.hasError {
		b.WriteString(" error")
	}

	b.WriteString(" {\n")

	// TODO:

	b.WriteString("}\n")

	return b.String()
}
