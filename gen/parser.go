package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type RawXml struct {
	Value string `xml:",innerxml"`
}

type VkTypes struct {
	Structs  []*Struct
	Typedefs []Typedef
}

type VkXML struct {
	Types *VkTypes
	Enums []*Enum
}

func ParseVkXML(file *os.File) (*VkXML, error) {
	out := &VkXML{}

	decoder := xml.NewDecoder(file)

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return out, nil
			}
			return nil, err
		}

		se, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		switch se.Name.Local {
		case "types":
			out.Types, err = ParseTypes(decoder)
			if err != nil {
				return nil, err
			}
		case "enums":
			var enum XmlEnum
			if err := decoder.DecodeElement(&enum, &se); err != nil {
				return nil, err
			}

			out.Enums = append(out.Enums, enum.Parse())
		}
	}
}

func ParseTypes(decoder *xml.Decoder) (*VkTypes, error) {
	out := &VkTypes{
		Structs: make([]*Struct, 0),
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		if ee, ok := token.(xml.EndElement); ok {
			if ee.Name.Local == "types" {
				return out, nil
			}
		}

		se, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		var category string
		for _, attr := range se.Attr {
			if attr.Name.Local == "category" {
				category = attr.Value
				break
			}
		}

		if category == "" {
			continue
		}

		switch category {
		case "struct":
			var str XmlStruct
			if err := decoder.DecodeElement(&str, &se); err != nil {
				return nil, err
			}

			out.Structs = append(out.Structs, str.Parse())
		case "bitmask", "basetype":
			var t RawXml
			if err := decoder.DecodeElement(&t, &se); err != nil {
				return nil, err
			}

			out.Typedefs = append(out.Typedefs, ParseTypedef(t.Value))
		default:
			fmt.Println("Unhandled:", category)
		}
	}
}
