package main

import (
	"fmt"
	"log"

	"github.com/lukem570/vulkan-go/internal/generator"
)

func main() {
	vkXml, err := generator.ParseXML("mod/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		log.Fatal(err)
	}

	config, err := generator.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	registry := generator.BuildRegistry(vkXml)
	reduced := generator.Filter(registry, config)

	const pkgName string = "vk"

	output := reduced.GeneratePackage(pkgName)
	err = generator.WriteFile("v1/vulkan.go", output)
	if err != nil {
		log.Fatalf("WriteFile (base): %v", err)
	}

	callbacksOutput := reduced.GenerateCallbacksFile(pkgName)
	err = generator.WriteFile("v1/callbacks.go", callbacksOutput)
	if err != nil {
		log.Fatalf("WriteFile (callbacks): %v", err)
	}

	platformFiles := reduced.GeneratePlatformFiles(pkgName)
	for _, pf := range platformFiles {
		path := fmt.Sprintf("v1/vulkan_%s.go", pf.BuildTag)
		err = generator.WriteFile(path, pf.Content)
		if err != nil {
			log.Fatalf("WriteFile (%s): %v", path, err)
		}
		log.Printf("Generated %s", path)
	}

	log.Printf(
		"Generated v1/vulkan.go (%d commands, %d structs, %d enums, %d handles)",
		len(reduced.Commands),
		len(reduced.Structs),
		len(reduced.Enums),
		len(reduced.Handles),
	)
}
