package generator

import "strings"

const header = `
/*
#cgo CFLAGS: -I./../../mod/Vulkan-Headers/include
#cgo LDFLAGS: -lvulkan 
#include <stdlib.h>
#include <vulkan/vulkan.h>
*/
import "C"
import "unsafe"

type Structure interface {
	getType() StructureType
	toC() (*C.void, func())
}
`

type Registry struct {
    Enums      map[string]*Enum
    Bitmasks   map[string]*Bitmask
    Structs    map[string]*Structured
    Commands   map[string]*GoCommand

    Features   map[string]*Feature
    Extensions map[string]*Extension
}

func (r *Registry) AddEnumElement(extends string, el *EnumElement) {
	if e, ok := r.Enums[extends]; ok {
		el.Parent = e
		e.Elements = append(e.Elements, el)
	}
}

func (r *Registry) GeneratePackage(pkg string) string {
	var b strings.Builder

	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(header)

	// Enums
	for _, e := range r.Enums {
		b.WriteString(e.Generate())
	}

	// Bitmasks
	for _, bmask := range r.Bitmasks {
		b.WriteString(bmask.Generate())
	}

	// Structs
	for _, s := range r.Structs {
		b.WriteString(s.GenerateStructure())
		b.WriteString(s.GenerateGetType())
		b.WriteString(s.GenerateToC())
	}

	// Commands
	for _, c := range r.Commands {
		b.WriteString(c.GenerateWrapper())
	}

	return b.String()
}

type Feature struct {
    Name     string // VK_VERSION_1_2
    Requires []RequireBlock
}

type Extension struct {
    Name     string
    Platform string // optional
    Requires []RequireBlock
}

type RequireBlock struct {
    Commands []string
    Types    []string
    Enums    []string
}