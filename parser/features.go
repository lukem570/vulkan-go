package parser

import (
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

type XMLFeature struct {
	Name     string       `xml:"name,attr"`
	Depends  string       `xml:"depends,attr"`
	Requires []XMLRequire `xml:"require"`
}

type XMLRequire struct {
	Commands []XMLRequireCmd `xml:"command"`
	Types    []XMLRequireTyp `xml:"type"`
	Enums    []XMLRequireEnum `xml:"enum"`
}

type XMLRequireCmd struct {
	Name string `xml:"name,attr"`
}

type XMLRequireTyp struct {
	Name string `xml:"name,attr"`
}

func (x *XMLRegistry) parseFeatures() map[string]*generator.Feature {
	features := make(map[string]*generator.Feature)

	for _, f := range x.Features {
		feat := &generator.Feature{Name: f.Name}
		if f.Depends != "" {
			for _, d := range strings.Split(f.Depends, ",") {
				d = strings.TrimSpace(d)
				if d != "" {
					feat.Depends = append(feat.Depends, d)
				}
			}
		}

		for _, req := range f.Requires {
			block := generator.RequireBlock{}
			for _, c := range req.Commands {
				block.Commands = append(block.Commands, c.Name)
			}
			for _, t := range req.Types {
				block.Types = append(block.Types, t.Name)
			}
			for _, e := range req.Enums {
				block.Enums = append(block.Enums, e.Name)
			}
			feat.Requires = append(feat.Requires, block)
		}

		features[feat.Name] = feat
	}

	return features
}