package generator

import (
	"fmt"
	"strings"
)

type Enum struct {
	CName     string
	GoName    string
	Base      string
	IsBitmask bool

	Elements []*EnumElement
}

type Bitmask struct {
	CName    string
	GoName   string
	BitsEnum *Enum
}

type EnumElement struct {
	CName  string
	GoName string
	Value  string
	BitPos *int
	Parent *Enum
}

func (e *Enum) Generate() string {
	if e == nil {
		return ""
	}

	// Deduplicate elements by GoName — extensions can re-add elements already
	// present from the core spec, causing "redeclared in this block" errors.
	seen := make(map[string]bool, len(e.Elements))
	deduped := make([]*EnumElement, 0, len(e.Elements))
	for _, el := range e.Elements {
		if el == nil || seen[el.GoName] {
			continue
		}
		seen[el.GoName] = true
		deduped = append(deduped, el)
	}

	// Determine base type: upgrade to uint64 if any bitpos >= 32
	base := e.Base
	if base == "" {
		base = "uint32"
	}
	if base == "uint32" {
		for _, el := range deduped {
			if el.BitPos != nil && *el.BitPos >= 32 {
				base = "uint64"
				break
			}
		}
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("type %s %s\n\n", e.GoName, base))
	b.WriteString("const (\n")

	for _, el := range deduped {
		if el.BitPos != nil {
			b.WriteString(fmt.Sprintf("\t%s %s = 1 << %d\n", el.GoName, e.GoName, *el.BitPos))
		} else {
			val := el.Value
			if val == "" {
				val = "0 // TODO: unknown value"
			}
			b.WriteString(fmt.Sprintf("\t%s %s = %s\n", el.GoName, e.GoName, val))
		}
	}

	b.WriteString(")\n\n")
	return b.String()
}

func (bmask *Bitmask) Generate() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("type %s uint32\n\n", bmask.GoName))
	return b.String()
}