package parser

import "github.com/lukem570/vulkan-go/generator"

type XMLCommands struct {
    Commands []XMLCommand `xml:"command"`
}

type XMLCommand struct {
    Proto  XMLProto      `xml:"proto"`
    Params []XMLParam    `xml:"param"`
    Alias  string        `xml:"alias,attr"`
}

type XMLProto struct {
    Type string `xml:"type"`
    Name string `xml:"name"`
}

type XMLParam struct {
    Type string `xml:"type"`
    Name string `xml:"name"`
}

func parseCommands(x *XMLRegistry, r *generator.Registry) {
	for _, cmd := range x.Commands.Commands {

		if cmd.Alias != "" {
			continue
		}

		c := &generator.GoCommand{
			Name: stripVk(cmd.Proto.Name),
		}

		for _, p := range cmd.Params {
			param := generator.CommandParam{
				Name: p.Name,
				Type:  parseFieldType(p.Type),
			}
			c.Params = append(c.Params, param)
		}

		r.Commands[c.Name] = c
	}
}