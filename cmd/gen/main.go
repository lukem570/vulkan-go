package main

import (
	"fmt"
	"log"

	"github.com/lukem570/vulkan-go/generator"
	"github.com/lukem570/vulkan-go/parser"
)

func main() {
	vkXml, err := parser.ParseXML("mod/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		log.Fatal(err)
	}

	config, err := generator.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	registry := parser.BuildRegistry(vkXml)

	linker := &generator.Linker{
		Registry: registry,
		Config:   config,
	}

	reduced := linker.Link()

	output := reduced.GeneratePackage("vulkan")

	err = generator.WriteFile("pkg/raw/vulkan.go", output)
	if err != nil {
		log.Fatalf("WriteFile: %v", err)
	}

	callbacksOutput := reduced.GenerateCallbacksFile("vulkan")
	err = generator.WriteFile("pkg/raw/callbacks.go", callbacksOutput)
	if err != nil {
		log.Fatalf("WriteFile (callbacks): %v", err)
	}

	// Generate platform-specific Go files
	platformFiles := reduced.GeneratePlatformFiles("vulkan")
	for _, pf := range platformFiles {
		path := fmt.Sprintf("pkg/raw/vulkan_%s.go", pf.BuildTag)
		err = generator.WriteFile(path, pf.Content)
		if err != nil {
			log.Fatalf("WriteFile (%s): %v", path, err)
		}
		log.Printf("Generated %s", path)
	}

	err = generator.WriteCFile("pkg/raw/volk_wrappers.h", reduced.GenerateCHeader())
	if err != nil {
		log.Fatalf("WriteCFile (header): %v", err)
	}

	err = generator.WriteCFile("pkg/raw/volk_wrappers.c", reduced.GenerateCSource())
	if err != nil {
		log.Fatalf("WriteCFile (source): %v", err)
	}

	log.Printf("Generated pkg/raw/vulkan.go, volk_wrappers.{c,h} (%d commands, %d structs, %d enums, %d handles)",
		len(reduced.Commands), len(reduced.Structs), len(reduced.Enums), len(reduced.Handles))
}