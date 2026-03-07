package parser

import (
	"strconv"
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

type XMLEnums struct {
	Name  string    `xml:"name,attr"`
	Type  string    `xml:"type,attr"`
	Enums []XMLEnum `xml:"enum"`
}

type XMLEnum struct {
	Name      string `xml:"name,attr"`
	Value     string `xml:"value,attr"`
	BitPos    string `xml:"bitpos,attr"`
	Extends   string `xml:"extends,attr"`
	Offset    string `xml:"offset,attr"`
	Dir       string `xml:"dir,attr"`
	Extnumber string `xml:"extnumber,attr"`
}

// XMLRequireEnum is used inside <feature>/<extension> require blocks.
// It has extra fields not present on top-level <enum> elements.
type XMLRequireEnum struct {
	Name      string `xml:"name,attr"`
	Extends   string `xml:"extends,attr"`
	Value     string `xml:"value,attr"`
	BitPos    string `xml:"bitpos,attr"`
	Offset    string `xml:"offset,attr"`
	Dir       string `xml:"dir,attr"`
	Extnumber string `xml:"extnumber,attr"`
	Alias     string `xml:"alias,attr"`
}

func parseEnums(x *XMLRegistry, r *generator.Registry) {
	for _, enums := range x.Enums {
		// API constants block has no type attribute — parse into APIConstants
		if enums.Type == "" || enums.Type == "constants" {
			for _, val := range enums.Enums {
				if c := translateAPIConstant(val); c != nil {
					r.APIConstants = append(r.APIConstants, *c)
				}
			}
			continue
		}

		e := &generator.Enum{
			CName:     enums.Name,
			GoName:    stripVk(enums.Name),
			Base:      "uint32",
			IsBitmask: enums.Type == "bitmask",
		}

		for _, val := range enums.Enums {
			el := makeEnumElement(val, e)
			if el != nil {
				e.Elements = append(e.Elements, el)
			}
		}

		r.Enums[e.GoName] = e
	}
}

// translateAPIConstant converts a raw C literal value from the XML into a
// valid Go expression, along with a type annotation where needed.
func translateAPIConstant(val XMLEnum) *generator.APIConstant {
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

	// (~0U)   → ^uint32(0)
	// (~0ULL) → ^uint64(0)
	// (~1U)   → ^uint32(0) - 1  ... actually (^uint32(0) - 1)
	// (~2U)   → ^uint32(0) - 2
	if strings.HasPrefix(raw, "(~") && strings.HasSuffix(raw, ")") {
		inner := raw[2 : len(raw)-1] // e.g. "0U", "0ULL", "1U", "2U"
		inner = strings.TrimRight(inner, "UuLl")
		n, err := strconv.ParseInt(inner, 0, 64)
		if err != nil {
			return "", ""
		}
		if strings.Contains(raw, "ULL") || strings.Contains(raw, "ull") {
			if n == 0 {
				return "^uint64(0)", "uint64"
			}
			return "^uint64(0) - " + strconv.FormatInt(n-1, 10), "uint64"
		}
		if n == 0 {
			return "^uint32(0)", "uint32"
		}
		return "^uint32(0) - " + strconv.FormatInt(n-1, 10), "uint32"
	}

	// Float literals: "1000.0F", "0.25f", "0.50f", "0.75f"
	lower := strings.ToLower(raw)
	if strings.HasSuffix(lower, "f") && strings.ContainsAny(raw, ".") {
		trimmed := raw[:len(raw)-1] // remove trailing F/f
		_, err := strconv.ParseFloat(trimmed, 32)
		if err != nil {
			return "", ""
		}
		return trimmed, "float32"
	}

	// Plain integer or hex with U/ULL suffix
	stripped := strings.TrimRight(raw, "UuLl")
	_, err := strconv.ParseInt(stripped, 0, 64)
	if err == nil {
		return stripped, ""
	}
	_, err = strconv.ParseUint(stripped, 0, 64)
	if err == nil {
		return stripped, ""
	}

	// Unrecognised — emit as a comment so we don't break compilation
	return "0 // TODO: untranslated: " + raw, ""
}

// applyEnumExtensions processes enum values that appear inside <feature> and
// <extension> require blocks and adds them to already-parsed enums.
func applyEnumExtensions(x *XMLRegistry, r *generator.Registry) {
	apply := func(enums []XMLRequireEnum, extNumber int) {
		for _, e := range enums {
			if e.Extends == "" {
				continue
			}
			parentGoName := stripVk(e.Extends)
			parent, ok := r.Enums[parentGoName]
			if !ok {
				continue
			}
			el := makeExtendedEnumElement(e, parent, extNumber)
			if el != nil {
				parent.Elements = append(parent.Elements, el)
			}
		}
	}

	for _, f := range x.Features {
		for _, req := range f.Requires {
			apply(req.Enums, 0)
		}
	}

	for _, ext := range x.Extensions.Extensions {
		for _, req := range ext.Requires {
			apply(req.Enums, ext.Number)
		}
	}
}

func makeEnumElement(val XMLEnum, parent *generator.Enum) *generator.EnumElement {
	el := &generator.EnumElement{
		CName:  val.Name,
		GoName: UpperSnakeToPascal(val.Name),
		Parent: parent,
	}
	if val.BitPos != "" {
		bp, err := strconv.Atoi(val.BitPos)
		if err == nil {
			el.BitPos = &bp
		}
	} else {
		el.Value = val.Value
	}
	return el
}

func makeExtendedEnumElement(val XMLRequireEnum, parent *generator.Enum, extNumber int) *generator.EnumElement {
	el := &generator.EnumElement{
		CName:  val.Name,
		GoName: UpperSnakeToPascal(val.Name),
		Parent: parent,
	}

	switch {
	case val.BitPos != "":
		bp, err := strconv.Atoi(val.BitPos)
		if err == nil {
			el.BitPos = &bp
		}
	case val.Offset != "":
		offset, err := strconv.Atoi(val.Offset)
		if err != nil {
			return nil
		}
		num := extNumber
		if val.Extnumber != "" {
			n, err := strconv.Atoi(val.Extnumber)
			if err == nil {
				num = n
			}
		}
		value := 1000000000 + (num-1)*1000 + offset
		if val.Dir == "-" {
			el.Value = strconv.Itoa(-value)
		} else {
			el.Value = strconv.Itoa(value)
		}
	case val.Value != "":
		el.Value = val.Value
	default:
		if val.Alias == "" {
			return nil
		}
		el.Value = UpperSnakeToPascal(val.Alias)
	}

	return el
}