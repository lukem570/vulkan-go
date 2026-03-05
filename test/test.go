package main

import (
	"fmt"
	"log"
	"strings"

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
		EnabledLayerNames: []string{
			"VK_LAYER_KHRONOS_validation",
		},
	}

	instance, err := vulkan.CreateInstance(create, nil)
	if err != nil {
		log.Fatal(err)
	}

	vulkan.LoadInstance(instance)

	physicalDevices, err := instance.EnumeratePhysicalDevices()
	if err != nil {
		log.Fatal(err)
	}

	var physicalDevice *vulkan.PhysicalDevice
	for _, phy := range physicalDevices {
		properties := phy.GetProperties()
		name := string(properties.DeviceName[:])
		// Trim null bytes
		if i := strings.IndexByte(name, 0); i >= 0 {
			name = name[:i]
		}
		
		switch properties.DeviceType {
		case vulkan.PhysicalDeviceTypeDiscreteGpu:
			fmt.Println("using:", name)
			physicalDevice = phy
		}
	}

	device, err := physicalDevice.CreateDevice(&vulkan.DeviceCreateInfo{
		
	}, nil)

	vulkan.LoadDevice(device)

	fmt.Println(device)
}
