package main

import (
	"log"

	"github.com/lukem570/vulkan-go/generator"
	"github.com/lukem570/vulkan-go/parser"
)

func main() {
	registry, err := parser.ParseXML("mod/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		log.Fatal(err)
	}

	config, err := generator.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	linker := &generator.Linker{
		Registry: parser.BuildRegistry(registry),
		Config:   config,
	}

	reduced := linker.Link()

	output := reduced.GeneratePackage("vulkan")

	err = generator.WriteFile("pkg/raw/vulkan.go", output)
	if err != nil {
		panic(err)
	}
}