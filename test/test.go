package main

import (
	"fmt"
	"log"

	vulkan "github.com/lukem570/vulkan-go/pkg/raw"
)

func VkMakeApiVersion(variant, major, minor, patch uint32) uint32 {
	return (variant << 29) | (major << 22) | (minor << 12) | patch
}

func main() {

	err := vulkan.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	appInfo := &vulkan.ApplicationInfo{
		ApplicationName:   "vulkan-go-test",
		ApplicationVersion: 1,
		EngineName:        "vulkan-go-test",
		EngineVersion:      1,
		ApiVersion:         VkMakeApiVersion(0, 1, 4, 0),
	}

	create := &vulkan.InstanceCreateInfo{
		ApplicationInfo: appInfo,
	}

	vulkan.CreateInstance(create)

	fmt.Println(*create)
}
