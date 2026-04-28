package main

import (
	"fmt"
	"log"

	"github.com/lukem570/vulkan-go/internal/generator"
	"github.com/lukem570/vulkan-go/internal/parser"
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
	
	const pkgName string = "vk"

	output := reduced.GeneratePackage(pkgName)

	err = generator.WriteFile("pkg/raw/vulkan.go", output)
	if err != nil {
		log.Fatalf("WriteFile (base): %v", err)
	}

	callbacksOutput := reduced.GenerateCallbacksFile(pkgName)
	err = generator.WriteFile("pkg/raw/callbacks.go", callbacksOutput)
	if err != nil {
		log.Fatalf("WriteFile (callbacks): %v", err)
	}

	// Generate platform-specific Go files
	platformFiles := reduced.GeneratePlatformFiles(pkgName)
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

	log.Printf(
		"Generated pkg/raw/vulkan.go, volk_wrappers.{c,h} (%d commands, %d structs, %d enums, %d handles)",
		len(reduced.Commands), 
		len(reduced.Structs), 
		len(reduced.Enums), 
		len(reduced.Handles),
	)
}