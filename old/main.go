package main

import (
	"log"
	"os"
)

func main() {
	in, err := os.Open("./mod/Vulkan-Headers/registry/vk.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	parser := &VkParser{}

	err = parser.Process(in)
	if err != nil {
		log.Fatal(err)
	}

	err = parser.Save()
	if err != nil {
		log.Fatal(err)
	}
}