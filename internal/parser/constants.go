package parser

import (
	"fmt"
	"strings"

	"github.com/lukem570/vulkan-go/internal/generator"
)

func (x *XMLRegistry) parseConstants() []generator.APIConstant {
	constants := make([]generator.APIConstant, 0)

	for _, enums := range x.Enums {
		if enums.Type != "" && enums.Type != "constants" {
			continue
		}
		

		for _, val := range enums.Enums {
			if c := newApiConstant(val); c != nil {
				constants = append(constants, *c)
			}
		}
	}

	return constants
}

func newApiConstant(val XMLEnum) *generator.APIConstant {
	goName := UpperSnakeToPascal(val.Name)
	raw := strings.TrimSpace(val.Value)

	goVal, goType := translateCLiteral(raw)
	if goVal == "" {
		return nil
	}

	return &generator.APIConstant{
		GoName: goName,
		Value:  goVal,
		Type:   goType,
	}
}

// translateCLiteral converts a C constant expression to a Go expression.
// Returns (goValue, goType). goType is empty for plain integer literals.
func translateCLiteral(raw string) (string, string) {
	if raw == "" {
		return "", ""
	}

	digits := extractDigits(raw)

	if strings.ContainsAny(raw, ".fF") {
		return digits, "float32"
	}

	base := ""
	if strings.ContainsAny(raw, "uU") {
		base = "uint"
	} else {
		base = "int"
	}

	bits := ""
	if strings.Contains(strings.ToLower(raw), "ll") {
		bits = "64"
	} else {
		bits = "32"
	}

	not := ""
	if strings.Contains(raw, "~") {
		not = "^"
	}

	typ := fmt.Sprintf("%s%s", base, bits)
	value := fmt.Sprintf("%s%s(%s)", not, typ, digits)

	return value, typ
}