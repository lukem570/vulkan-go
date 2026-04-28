package parser

import (
	"strconv"

	"github.com/lukem570/vulkan-go/internal/generator"
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

func (x *XMLRegistry) parseEnums() map[string]*generator.Enum {
	parsed := make(map[string]*generator.Enum)

	for _, enums := range x.Enums {
		if enums.Type == "" || enums.Type == "constants" {
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

		parsed[e.GoName] = e
	}

	return parsed
}



// applyEnumExtensions processes enum values that appear inside <feature> and
// <extension> require blocks and adds them to already-parsed enums.
func (x *XMLRegistry) applyEnumExtensions(r *generator.Registry) {
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
