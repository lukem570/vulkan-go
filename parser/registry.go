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
		Enums:        make(map[string]*generator.Enum),
		Bitmasks:     make(map[string]*generator.Bitmask),
		Structs:      make(map[string]*generator.Structured),
		Handles:      make(map[string]*generator.GoHandle),
		Commands:     make(map[string]*generator.GoCommand),
		Features:     make(map[string]*generator.Feature),
		Extensions:   make(map[string]*generator.Extension),
		FuncPointers: make(map[string]*generator.GoFuncPointer),
		Platforms:    make(map[string]string),
	}

	// Parse platform name → protect macro mapping
	for _, p := range x.Platforms.Platforms {
		r.Platforms[p.Name] = p.Protect
	}

	parseEnums(x, r)
	parseTypes(x, r)
	parseCommands(x, r)
	parseFeatures(x, r)
	parseExtensions(x, r)
	applyEnumExtensions(x, r)

	return r
}