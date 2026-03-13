package parser

import (
	"encoding/xml"
	"os"

	"github.com/lukem570/vulkan-go/generator"
)

type XMLPlatforms struct {
	Platforms []XMLPlatform `xml:"platform"`
}

type XMLPlatform struct {
	Name    string `xml:"name,attr"`
	Protect string `xml:"protect,attr"`
}

type XMLRegistry struct {
	XMLName    xml.Name      `xml:"registry"`
	Platforms  XMLPlatforms  `xml:"platforms"`
	Types      XMLTypes      `xml:"types"`
	Enums      []XMLEnums    `xml:"enums"`
	Commands   XMLCommands   `xml:"commands"`
	Features   []XMLFeature  `xml:"feature"`
	Extensions XMLExtensions `xml:"extensions"`
}

func ParseXML(path string) (*XMLRegistry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reg := &XMLRegistry{}
	if err := xml.Unmarshal(data, reg); err != nil {
		return nil, err
	}

	return reg, nil
}

func BuildRegistry(x *XMLRegistry) *generator.Registry {
	r := &generator.Registry{
		Enums:        x.parseEnums(),
		APIConstants: x.parseConstants(),
		Bitmasks:     make(map[string]*generator.Bitmask),
		EnumAliases:  make(map[string]*generator.EnumAlias),
		Structs:      make(map[string]*generator.Structured),
		Handles:      make(map[string]*generator.GoHandle),
		Commands:     make(map[string]*generator.GoCommand),
		STypes:       make(map[string]string),
		Features:     x.parseFeatures(),
		Extensions:   x.parseExtensions(),
		FuncPointers: make(map[string]*generator.GoFuncPointer),
		Platforms:    make(map[string]string),
	}

	for _, p := range x.Platforms.Platforms {
		r.Platforms[p.Name] = p.Protect
	}

	x.parseTypes(r)
	x.parseCommands(r)
	x.applyEnumExtensions(r)

	return r
}
