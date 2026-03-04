package parser

import "github.com/lukem570/vulkan-go/generator"

type XMLEnums struct {
	Name  string    `xml:"name,attr"`
	Type  string    `xml:"type,attr"`
	Enums []XMLEnum `xml:"enum"`
}

type XMLEnum struct {
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
	BitPos  *int   `xml:"bitpos,attr"`
	Extends string `xml:"extends,attr"`
}

func parseEnums(x *XMLRegistry, r *generator.Registry) {
	for _, enums := range x.Enums {

		// Skip API constants block
		if enums.Type == "" {
			continue
		}

		e := &generator.Enum{
			CName:     enums.Name,
			GoName:    stripVk(enums.Name),
			Base:      "uint32",
			IsBitmask: enums.Type == "bitmask",
		}

		for _, val := range enums.Enums {
			el := &generator.EnumElement{
				CName:  val.Name,
				GoName: stripVk(val.Name),
				Value:  val.Value,
				BitPos: val.BitPos,
				Parent: e,
			}
			e.Elements = append(e.Elements, el)
		}

		r.Enums[e.GoName] = e
	}
}
