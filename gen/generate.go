package main

import (
	"os"
	"strings"
)

func Generate(vkXml *VkXML) error {
	var structures strings.Builder
	structures.WriteString("package vulkan\n\n")
	structures.WriteString("import \"unsafe\"\n\n")

	for _, def := range vkXml.Types.Typedefs {
		structures.WriteString(def.Generate())
	}

	for _, str := range vkXml.Types.Structs {
		structures.WriteString(str.Generate())
	}

	err := os.WriteFile("./pkg/raw/structures.go", []byte(structures.String()), 0644)
	if err != nil {
		return err
	}

	var enums strings.Builder
	enums.WriteString("package vulkan\n\n")

	for _, enum := range vkXml.Enums {
		enums.WriteString(enum.Generate())
	}

	err = os.WriteFile("./pkg/raw/enums.go", []byte(enums.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}