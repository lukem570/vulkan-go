package generator

import (
	"fmt"
	"strings"
)

type Enum struct {
	CName   string // e.g. VkStructureType
	GoName  string // e.g. StructureType
	Base    string // usually uint32 / int32
	IsBitmask bool

	Elements []*EnumElement
}

type Bitmask struct {
	CName    string // VkAccessFlags
	GoName   string // AccessFlags
	BitsEnum *Enum  // AccessFlagBits enum
}

type EnumElement struct {
	CName  string // VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	GoName string // StructureTypeInstanceCreateInfo

	Value  string // literal value or bit position
	BitPos *int   // if it's a bitmask flag

	Parent *Enum
}

func (e *Enum) Generate() string {
	if e == nil {
		return ""
	}

	var b strings.Builder

	base := e.Base
	if base == "" {
		base = "uint32"
	}

	b.WriteString(fmt.Sprintf("type %s %s\n\n", e.GoName, base))
	b.WriteString("const (\n")

	for _, el := range e.Elements {
		if el == nil {
			continue
		}

		if el.BitPos != nil {
			b.WriteString(fmt.Sprintf("\t%s %s = 1 << %d\n", el.GoName, e.GoName, *el.BitPos))
		} else {
			b.WriteString(fmt.Sprintf("\t%s %s = %s\n", el.GoName, e.GoName, el.Value))
		}
	}

	b.WriteString(")\n\n")

	return b.String()
}

func (bmask *Bitmask) Generate() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(
		"type %s uint32\n\n",
		bmask.GoName,
	))

	return b.String()
}