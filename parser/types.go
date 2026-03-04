package parser

import "github.com/lukem570/vulkan-go/generator"

type XMLTypes struct {
	Types []XMLType `xml:"type"`
}

type XMLMember struct {
	Type string `xml:"type"`
	Name string `xml:"name"`

	Len   string `xml:"len,attr"`
	AltLen string `xml:"altlen,attr"`

	InnerXML string `xml:",innerxml"`
}

type XMLType struct {
	Category string `xml:"category,attr"`
	Name     string `xml:"name,attr"`
	Alias    string `xml:"alias,attr"`
	Requires string `xml:"requires,attr"`
	Parent   string `xml:"parent,attr"`

	Members  []XMLMember `xml:"member"`
}

func parseTypes(x *XMLRegistry, r *generator.Registry) {
	for _, t := range x.Types.Types {

		switch t.Category {

		case "struct":
			s := parseStruct(t)
			r.Structs[s.GoName] = s

		case "bitmask":
			b := &generator.Bitmask{
				CName:  t.Name,
				GoName: stripVk(t.Name),
			}
			r.Bitmasks[b.GoName] = b

		case "handle":
			// Optional: implement later
			// For now treat as named type

		}
	}
}