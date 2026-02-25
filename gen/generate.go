package main

import (
	"fmt"
	"os"
	"strings"
)

type Generatable interface {
	Generate() string
}

type FeatureLookup struct {
	Features map[string]*Feature

	Types    map[string]Generatable
	Commands map[string]Generatable
}

func NewFeatureLookup(vkXml *VkXML) *FeatureLookup {
	out := &FeatureLookup{
		Features: make(map[string]*Feature),
		Types:    make(map[string]Generatable),
		Commands: make(map[string]Generatable),
	}

	for _, feat := range vkXml.Features {
		out.Features[feat.Name] = feat
	}

	for _, str := range vkXml.Types.Structs {
		out.Types[str.Name] = str
	}

	for _, def := range vkXml.Types.Typedefs {
		out.Types[def.Alias] = def
	}

	for _, hwd := range vkXml.Types.Handles {
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
	out.WriteString("package vulkan\n\n")
	out.WriteString("import \"unsafe\"\n\n")

	lookup := NewFeatureLookup(vkXml)
	err := lookup.Generate(target, &out)
	if err != nil {
		return err
	}

	for _, enum := range vkXml.Enums {
		out.WriteString(enum.Generate())
	}

	err = os.WriteFile("./pkg/raw/structs.go", []byte(out.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *FeatureLookup) Generate(target string, out *strings.Builder) error {
	if target == "" {
		return nil
	}

	feat, ok := f.Features[target]
	if !ok {
		return fmt.Errorf("Unknown target '%s'", target)
	}

	for _, depend := range feat.Depends {
		err := f.Generate(depend, out)
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
			default:
			}
		}
	}

	return nil
}
