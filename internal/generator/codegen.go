package generator

import (
	"fmt"
	"strings"
)

type CodeGen struct {
	builder  strings.Builder
	varIndex int
}

func (g *CodeGen) Var(prefix string) string {
	name := fmt.Sprintf("%s%d", prefix, g.varIndex)
	g.varIndex++
	return name
}

func (g *CodeGen) Line(s string) {
	g.builder.WriteString(s)
	g.builder.WriteByte('\n')
}

func (g *CodeGen) String() string {
	return g.builder.String()
}