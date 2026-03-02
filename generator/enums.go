package generator

import (
	"fmt"
	"strings"
)

type Enum struct {
	cName   string // e.g. VkStructureType
	goName  string // e.g. StructureType
	base    string // usually uint32 / int32
	isBitmask bool

	elements []*EnumElement
}

type Bitmask struct {
	cName    string // VkAccessFlags
	goName   string // AccessFlags
	bitsEnum *Enum  // AccessFlagBits enum
}

type EnumElement struct {
	cName  string // VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	goName string // StructureTypeInstanceCreateInfo

	value  string // literal value or bit position
	bitPos *int   // if it's a bitmask flag

	parent *Enum
}

func (e *Enum) Generate() string {
	var b strings.Builder

	base := e.base
	if base == "" {
		base = "uint32"
	}

	b.WriteString(fmt.Sprintf(
		"type %s %s\n\n",
		e.goName,
		base,
	))

	b.WriteString("const (\n")

	for _, el := range e.elements {
		if el.bitPos != nil {
			b.WriteString(fmt.Sprintf(
				"\t%s %s = 1 << %d\n",
				el.goName,
				e.goName,
				*el.bitPos,
			))
		} else {
			b.WriteString(fmt.Sprintf(
				"\t%s %s = %s\n",
				el.goName,
				e.goName,
				el.value,
			))
		}
	}

	b.WriteString(")\n\n")

	return b.String()
}

func (bmask *Bitmask) Generate() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(
		"type %s uint32\n\n",
		bmask.goName,
	))

	return b.String()
}