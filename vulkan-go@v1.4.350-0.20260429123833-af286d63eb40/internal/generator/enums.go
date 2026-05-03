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
	Platform string // non-empty for platform-specific bitmasks
	BitsEnum *Enum
	Is64Bit  bool // true for VkFlags64-backed bitmasks
}

type EnumAlias struct {
	CName string
	GoName  string

	Alias string
}

type EnumElement struct {
	CName    string
	GoName   string
	Value    string
	BitPos   *int
	Parent   *Enum
	Platform string // non-empty for platform-specific extension values
}

func (e *EnumAlias) Generate() string {
	return fmt.Sprintf(
		"type %s %s\n\n",
		e.GoName,
		e.Alias,
	)
}

// SetReserved marks element names that conflict with other Go declarations
// (e.g. struct names). Conflicting elements will be skipped during generation.
func (e *Enum) SetReserved(reserved map[string]bool) {
	filtered := make([]*EnumElement, 0, len(e.Elements))
	for _, el := range e.Elements {
		if el != nil && !reserved[el.GoName] {
			filtered = append(filtered, el)
		}
	}
	e.Elements = filtered
}

// Generate emits the type declaration and const block for this enum.
// If platformFilter is non-empty, only elements matching that platform are
// emitted and the type declaration is omitted (it lives in the core file).
func (e *Enum) Generate(platformFilter ...string) string {
	if e == nil {
		return ""
	}

	platform := ""
	if len(platformFilter) > 0 {
		platform = platformFilter[0]
	}

	// Deduplicate elements by GoName — extensions can re-add elements already
	// present from the core spec, causing "redeclared in this block" errors.
	seen := make(map[string]bool, len(e.Elements))
	deduped := make([]*EnumElement, 0, len(e.Elements))
	for _, el := range e.Elements {
		if el == nil || seen[el.GoName] {
			continue
		}
		// Filter by platform
		if platform == "" && el.Platform != "" {
			continue // core generation: skip platform elements
		}
		if platform != "" && el.Platform != platform {
			continue // platform generation: skip non-matching elements
		}
		seen[el.GoName] = true
		deduped = append(deduped, el)
	}

	if len(deduped) == 0 {
		return ""
	}

	// Determine base type (needed for all elements, not just filtered)
	base := e.Base
	if base == "" {
		base = "uint32"
	}
	if base == "uint32" {
		for _, el := range e.Elements {
			if el == nil {
				continue
			}
			if el.Value != "" && strings.HasPrefix(el.Value, "-") {
				base = "int32"
				break
			}
			if el.BitPos != nil && *el.BitPos >= 32 {
				base = "uint64"
				break
			}
		}
	}

	var b strings.Builder
	if platform == "" {
		b.WriteString(fmt.Sprintf("type %s %s\n\n", e.GoName, base))
	}
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

// GenerateError emits an Error() method for the enum using a switch over
// canonical (non-alias) values. This is to implement the error interface
func (e *Enum) GenerateError() string {
	if e == nil {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("func (v %s) Error() string {\n\tswitch v {\n", e.GoName))

	seenValues := make(map[string]bool)
	for _, el := range e.Elements {
		if el == nil || el.Platform != "" {
			continue
		}
		// Skip aliases: Value is another identifier (not a numeric literal)
		val := el.Value
		if val == "" {
			continue
		}
		isNumeric := val[0] == '-' || (val[0] >= '0' && val[0] <= '9')
		if !isNumeric {
			continue
		}
		if seenValues[val] {
			continue
		}
		seenValues[val] = true
		b.WriteString(fmt.Sprintf("\tcase %s:\n\t\treturn %q\n", el.GoName, el.GoName))
	}

	b.WriteString(fmt.Sprintf("\tdefault:\n\t\treturn fmt.Sprintf(\"%s(%%d)\", int(v))\n\t}\n}\n\n", e.GoName))
	return b.String()
}

func (bmask *Bitmask) Generate() string {
	var b strings.Builder
	if bmask.Is64Bit {
		b.WriteString(fmt.Sprintf("type %s uint64\n\n", bmask.GoName))
	} else {
		b.WriteString(fmt.Sprintf("type %s uint32\n\n", bmask.GoName))
	}
	return b.String()
}
