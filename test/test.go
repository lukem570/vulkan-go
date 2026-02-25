package main

import (
	"fmt"

	vulkan "github.com/lukem570/vulkan-go/pkg/raw"
)

func VkMakeApiVersion(variant, major, minor, patch uint32) uint32 {
	return (variant << 29) | (major << 22) | (minor << 12) | patch
}

func main() {

	appInfo := &vulkan.VkApplicationInfo{
		SType:              vulkan.VkStructureTypeApplicationInfo,
		PNext:              nil,
		PApplicationName:   "vulkan-go-test",
		ApplicationVersion: 1,
		PEngineName:        "vulkan-go-test",
		EngineVersion:      1,
		ApiVersion:         VkMakeApiVersion(0, 1, 4, 0),
	}

	create := &vulkan.VkInstanceCreateInfo{
		SType:            vulkan.VkStructureTypeInstanceCreateInfo,
		PNext:            nil,
		PApplicationInfo: appInfo,
	}

	fmt.Println(*create)
}
