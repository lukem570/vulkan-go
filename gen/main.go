package main

import (
	"log"
	"os"
)

// vk.xml structure:
// platforms
// types
// enums
// commands
// features
// extensions

func main() {
	in, err := os.Open("./mod/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	vkXml, err := ParseVkXML(in)
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(vkXml, "VK_VERSION_1_4")
	if err != nil {
		log.Fatal(err)
	}
}