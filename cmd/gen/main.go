package main

import (
	"log"
	"strings"

	"github.com/lukem570/vulkan-go/generator"
	"github.com/lukem570/vulkan-go/parser"
)

// receiverPrefixes lists the prefixes to strip from method names after the
// receiver type prefix itself is removed. For example a CommandBuffer method
// named "CmdBeginRenderPass" after stripping "CommandBuffer" becomes
// "CmdBeginRenderPass" — we additionally strip "Cmd" to get "BeginRenderPass".
var receiverPrefixes = map[string][]string{
	"CommandBuffer": {"Cmd"},
}


func main() {
	registry, err := parser.ParseXML("mod/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		log.Fatal(err)
	}

	config, err := generator.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	base := parser.BuildRegistry(registry)

	linker := &generator.Linker{
		Registry: base,
		Config:   config,
	}

	reduced := linker.Link()

	for _, cmd := range reduced.Commands {
		if cmd.ReceiverType == "" {
			continue
		}

		// Strip the receiver type name prefix (e.g. "Instance" from "InstanceCreate")
		if strings.HasPrefix(cmd.Name, cmd.ReceiverType) {
			if trimmed := strings.TrimPrefix(cmd.Name, cmd.ReceiverType); trimmed != "" {
				cmd.Name = trimmed
			}
		}

		// Strip receiver type after verb prefixes
		// e.g. "GetPhysicalDeviceProperties" on PhysicalDevice → "GetProperties"
		verbs := []string{"Get", "Set", "Create", "Destroy", "Allocate", "Free",
			"Enumerate", "Reset", "Begin", "End", "Bind", "Queue", "Wait"}
		for _, verb := range verbs {
			prefix := verb + cmd.ReceiverType
			if strings.HasPrefix(cmd.Name, prefix) {
				if rest := strings.TrimPrefix(cmd.Name, prefix); rest != "" {
					cmd.Name = verb + rest
				}
				break
			}
		}

		// Strip any additional receiver-specific prefixes (e.g. "Cmd" from CommandBuffer methods)
		for _, extra := range receiverPrefixes[cmd.ReceiverType] {
			if strings.HasPrefix(cmd.Name, extra) {
				if trimmed := strings.TrimPrefix(cmd.Name, extra); trimmed != "" {
					cmd.Name = trimmed
				}
			}
		}
	}

	output := reduced.GeneratePackage("vulkan")

	err = generator.WriteFile("pkg/raw/vulkan.go", output)
	if err != nil {
		log.Fatalf("WriteFile: %v", err)
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