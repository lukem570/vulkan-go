package main

import (
	"encoding/xml"
	"strings"
)

type FeatureElementType int

const (
	ElementType FeatureElementType = iota
	ElementCommand
	ElementEnum
)

type FeatureBlockType int

const (
	BlockRequire FeatureBlockType = iota
	BlockDeprecate
)

type FeatureElement struct {
	Type FeatureElementType
	Name string
}

type FeatureBlock struct {
	Type     FeatureBlockType
	Comment  string
	Elements []FeatureElement
}

type Feature struct {
	Name    string
	Depends []string
	Comment string
	Blocks  []FeatureBlock
}

type XmlFeature struct {
	XMLName xml.Name `xml:"feature"`
	Name    string   `xml:"name,attr"`
	Depends string   `xml:"depends,attr"`
	Comment string   `xml:"comment,attr"`

	Require   []XmlFeatureBlock `xml:"require"`
	Deprecate []XmlFeatureBlock `xml:"deprecate"`
}

type XmlFeatureBlock struct {
	Comment  string              `xml:"comment,attr"`
	Types    []XmlFeatureElement `xml:"type"`
	Commands []XmlFeatureElement `xml:"command"`
	Enums    []XmlFeatureElement `xml:"enum"`
}

type XmlFeatureElement struct {
	Name string `xml:"name,attr"`
}

func (f *XmlFeature) Parse() *Feature {
	out := &Feature{
		Name:    f.Name,
		Depends: strings.Split(f.Depends, ","),
		Comment: f.Comment,
		Blocks:  make([]FeatureBlock, 0, len(f.Require)+len(f.Deprecate)),
	}

	for _, block := range f.Require {
		blk := FeatureBlock{
			Comment:  block.Comment,
			Elements: make(
				[]FeatureElement, 
				0, 
				len(block.Types) + len(block.Commands) + len(block.Enums),
			),
			Type: BlockRequire,
		}

		for _, element := range block.Commands {
			blk.Elements = append(blk.Elements, element.Parse(ElementCommand))
		}

		for _, element := range block.Enums {
			blk.Elements = append(blk.Elements, element.Parse(ElementEnum))
		}

		for _, element := range block.Types {
			blk.Elements = append(blk.Elements, element.Parse(ElementType))
		}

		out.Blocks = append(out.Blocks, blk)
	}

	return out
}


func (f *XmlFeatureElement) Parse(Type FeatureElementType) FeatureElement {
	return FeatureElement{
		Type: Type,
		Name: f.Name,
	}
}