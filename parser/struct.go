package parser

import (

	"github.com/lukem570/vulkan-go/generator"
)

// parseStruct parses a Vulkan XML type into a generator.Structured
func parseStruct(t XMLType) *generator.Structured {
	name := t.Name

	s := &generator.Structured{
		CName:  name,
		GoName: stripVk(name),
		Fields: []generator.StructField{},
	}

	// Detect if this struct has sType and pNext
	s.HasSType = hasSTypeAndPNext(t)

	// Set sType enum if applicable
	if s.HasSType {
		s.SType = "StructureType" + stripVk(name)
	}

	// Parse fields
	for _, m := range t.Members {
		if m.Name == "sType" || m.Name == "pNext" {
			continue
		}
		
		field := generator.StructField{
			GoName: toPublic(stripVk(m.Name)),
			CName:  m.Name,
			Type:   parseFieldType(m.Type),
			Count:  m.Len, // optional for arrays
		}
		s.Fields = append(s.Fields, field)
	}

	return s
}

// hasSTypeAndPNext returns true if the struct has both sType and pNext members
func hasSTypeAndPNext(t XMLType) bool {
	hasSType := false
	hasPNext := false

	for _, m := range t.Members {
		switch m.Name {
		case "sType":
			hasSType = true
		case "pNext":
			hasPNext = true
		}
		if hasSType && hasPNext {
			return true
		}
	}

	return false
}

// splitCamel splits CamelCase into words
func splitCamel(s string) []string {
	var words []string
	start := 0
	for i := 1; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			words = append(words, s[start:i])
			start = i
		}
	}
	words = append(words, s[start:])
	return words
}