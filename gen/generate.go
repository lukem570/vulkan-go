package main

import (
	"fmt"
	"os"
	"strings"
)

const header = `package vulkan

import (
	"unsafe"
	"fmt"
)

/*
#cgo CFLAGS: -I./../../mod/volk -I./../../mod/Vulkan-Headers/include/vulkan
#define VOLK_IMPLEMENTATION
#include <stdlib.h>
#include <volk.h>
#include <vulkan.h>
*/
import "C" 

// Target: GO_VULKAN

func Initialize() error {
	res := C.volkInitialize()

	if res != C.VK_SUCCESS {
		return fmt.Errorf("Failed to find Vulkan library.") 
	}

	return nil
}
`

type Generatable interface {
	Generate() string
}

type FeatureLookup struct {
	Features map[string]*Feature

	Types    map[string]Generatable
	Commands map[string]*Command

	Handles map[string]*Handle
}

func NewFeatureLookup(vkXml *VkXML) *FeatureLookup {
	out := &FeatureLookup{
		Features: make(map[string]*Feature),
		Types:    make(map[string]Generatable),
		Commands: make(map[string]*Command),

		Handles: make(map[string]*Handle),
	}

	for _, feat := range vkXml.Features {
		out.Features[feat.Name] = feat
	}

	for _, cmd := range vkXml.Commands {
		out.Commands[cmd.Proto.Name] = cmd
	}

	for _, str := range vkXml.Types.Structs {
		out.Types[str.Name] = str
	}

	for _, def := range vkXml.Types.Typedefs {
		out.Types[def.Alias] = def
	}

	for _, hwd := range vkXml.Types.Handles {
		out.Handles[hwd.Name] = hwd
		out.Types[hwd.Name] = hwd
	}

	for _, fn := range vkXml.Types.Funcs {
		out.Types[fn.Proto.Name] = fn
	}

	return out
}

func Generate(vkXml *VkXML, target string) error {

	fmt.Println("Generating", target)

	var out strings.Builder
	out.WriteString(header)

	lookup := NewFeatureLookup(vkXml)
	err := lookup.Generate(target, &out, make(map[string]struct{}))
	if err != nil {
		return err
	}

	for _, enum := range vkXml.Enums {
		out.WriteString(enum.Generate())
	}

	err = os.WriteFile("./pkg/raw/vulkan.go", []byte(out.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *FeatureLookup) Generate(
	target string,
	out *strings.Builder,
	seen map[string]struct{},
) error {
	if target == "" {
		return nil
	}

	if _, ok := seen[target]; ok {
		return nil
	}
	seen[target] = struct{}{}

	feat, ok := f.Features[target]
	if !ok {
		return fmt.Errorf("Unknown target '%s'", target)
	}

	for _, depend := range feat.Depends {
		err := f.Generate(depend, out, seen)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(out,
		"// Target: %s\n\n",
		target)

	for _, block := range feat.Blocks {
		if block.Type != BlockRequire {
			continue
		}

		if block.Comment != "" {
			fmt.Fprintf(out,
				"// %s\n",
				block.Comment)
		}

		for _, element := range block.Elements {
			switch element.Type {
			case ElementType:
				typ, ok := f.Types[element.Name]
				if !ok {
					fmt.Println("Missing type:", element.Name)
					continue
				}

				out.WriteString(typ.Generate())
			case ElementCommand:
				cmd, ok := f.Commands[element.Name]
				if !ok {
					fmt.Println("Missing command:", element.Name)
					continue
				}

				out.WriteString(cmd.Generate(f.Handles))
			default:
			}
		}
	}

	return nil
}
