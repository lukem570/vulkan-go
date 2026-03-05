package parser

import "github.com/lukem570/vulkan-go/generator"

type XMLExtensions struct {
	Extensions []XMLExtension `xml:"extension"`
}

type XMLExtension struct {
	Name     string       `xml:"name,attr"`
	Number   int          `xml:"number,attr"`
	Platform string       `xml:"platform,attr"`
	Requires []XMLRequire `xml:"require"`
}

func parseExtensions(x *XMLRegistry, r *generator.Registry) {
	for _, ext := range x.Extensions.Extensions {
		e := &generator.Extension{
			Name:     ext.Name,
			Platform: ext.Platform,
		}

		for _, req := range ext.Requires {
			block := generator.RequireBlock{}
			for _, c := range req.Commands {
				block.Commands = append(block.Commands, c.Name)
			}
			for _, t := range req.Types {
				block.Types = append(block.Types, t.Name)
			}
			for _, en := range req.Enums {
				block.Enums = append(block.Enums, en.Name)
			}
			e.Requires = append(e.Requires, block)
		}

		r.Extensions[e.Name] = e
	}
}