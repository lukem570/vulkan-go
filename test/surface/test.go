package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/lukem570/vulkan-go/pkg/raw"
	"github.com/lukem570/vulkan-go/pkg/util"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Initialize GLFW first to query required extensions
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("glfw init: %w", err)
	}
	defer glfw.Terminate()

	// Initialize Vulkan via volk
	if err := vk.Initialize(); err != nil {
		return fmt.Errorf("initialize: %w", err)
	}

	// Tell GLFW we want Vulkan
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)

	// Create a temporary window to query required extensions
	tmpWindow, err := glfw.CreateWindow(1, 1, "", nil, nil)
	if err != nil {
		return fmt.Errorf("create tmp window: %w", err)
	}
	glfwExts := tmpWindow.GetRequiredInstanceExtensions()
	tmpWindow.Destroy()

	enabledExts := append(glfwExts, "VK_EXT_debug_utils")

	// Create instance with validation layer and required extensions
	instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
		ApplicationInfo: &vk.ApplicationInfo{
			ApplicationName:    "vulkan-go-surface-test",
			ApplicationVersion: 1,
			EngineName:         "vulkan-go",
			EngineVersion:      1,
			ApiVersion:         vkutil.VkMakeApiVersion(0, 1, 4, 0),
		},
		EnabledLayerNames: []string{
			"VK_LAYER_KHRONOS_validation",
		},
		EnabledExtensionNames: enabledExts,
	}, nil)
	if err != nil {
		return fmt.Errorf("create instance: %w", err)
	}
	defer instance.Destroy(nil)

	fmt.Println("instance created with validation layers")

	// Create the actual window
	window, err := glfw.CreateWindow(800, 600, "Vulkan Triangle", nil, nil)
	if err != nil {
		return fmt.Errorf("create window: %w", err)
	}

	// Create surface using the safe GLFW wrapper.
	// Pass (*byte)(instance.Handle()) so GLFW's reflection-based API
	// receives the correct VkInstance handle value.
	surfPtr, err := window.CreateWindowSurface((*byte)(instance.Handle()), nil)
	if err != nil {
		return fmt.Errorf("create surface: %w", err)
	}
	surface := vk.SurfaceKHRFromGLFW(surfPtr)
	defer instance.DestroySurfaceKHR(surface, nil)
	fmt.Println("surface created")

	// Pick a physical device
	physDevices, err := instance.EnumeratePhysicalDevices()
	if err != nil {
		return fmt.Errorf("enumerate physical devices: %w", err)
	}
	if len(physDevices) == 0 {
		return fmt.Errorf("no Vulkan-capable GPU found")
	}
	physDevice := physDevices[0]

	// Find a queue family that supports graphics and presentation
	queueFamilies := physDevice.GetQueueFamilyProperties()
	graphicsFamily := -1
	for i, qf := range queueFamilies {
		hasGraphics := qf.QueueFlags&vk.QueueFlags(vk.QueueGraphicsBit) != 0

		supported, err := physDevice.GetSurfaceSupportKHR(uint32(i), surface)
		if err != nil {
			continue
		}

		if hasGraphics && supported {
			graphicsFamily = i
			break
		}
	}
	if graphicsFamily < 0 {
		return fmt.Errorf("no suitable queue family found")
	}
	fmt.Printf("using queue family %d\n", graphicsFamily)

	// Query surface capabilities
	caps, err := physDevice.GetSurfaceCapabilitiesKHR(surface)
	if err != nil {
		return fmt.Errorf("get surface capabilities: %w", err)
	}
	fmt.Printf("surface capabilities: min images=%d, current extent=%dx%d\n",
		caps.MinImageCount, caps.CurrentExtent.Width, caps.CurrentExtent.Height)

	// Query surface formats
	formats, err := physDevice.GetSurfaceFormatsKHR(surface)
	if err != nil {
		return fmt.Errorf("get surface formats: %w", err)
	}
	fmt.Printf("surface formats: %d available\n", len(formats))

	// Query present modes
	presentModes, err := physDevice.GetSurfacePresentModesKHR(surface)
	if err != nil {
		return fmt.Errorf("get present modes: %w", err)
	}
	fmt.Printf("present modes: %d available\n", len(presentModes))

	// Create logical device
	device, err := physDevice.CreateDevice(&vk.DeviceCreateInfo{
		QueueCreateInfos: []vk.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: uint32(graphicsFamily),
				QueuePriorities:  []float32{1.0},
			},
		},
		EnabledExtensionNames: []string{
			"VK_KHR_swapchain",
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create device: %w", err)
	}
	defer device.Destroy(nil)
	fmt.Println("device created")

	// Pick a swapchain format
	chosenFormat := formats[0]
	for _, f := range formats {
		if f.Format == vk.FormatB8g8r8a8Srgb && f.ColorSpace == vk.ColorSpaceSrgbNonlinearKHR {
			chosenFormat = f
			break
		}
	}

	// Pick a present mode (prefer mailbox for triple buffering)
	chosenPresentMode := vk.PresentModeFifoKHR
	for _, pm := range presentModes {
		if pm == vk.PresentModeMailboxKHR {
			chosenPresentMode = pm
			break
		}
	}

	imageCount := caps.MinImageCount + 1
	if caps.MaxImageCount > 0 && imageCount > caps.MaxImageCount {
		imageCount = caps.MaxImageCount
	}

	// Create swapchain
	swapchain, err := device.CreateSwapchainKHR(&vk.SwapchainCreateInfoKHR{
		Surface:          surface,
		MinImageCount:    imageCount,
		ImageFormat:      chosenFormat.Format,
		ImageColorSpace:  chosenFormat.ColorSpace,
		ImageExtent:      caps.CurrentExtent,
		ImageArrayLayers: 1,
		ImageUsage:       vk.ImageUsageFlags(vk.ImageUsageColorAttachmentBit),
		ImageSharingMode: vk.SharingModeExclusive,
		PreTransform:     caps.CurrentTransform,
		CompositeAlpha:   vk.CompositeAlphaOpaqueBitKHR,
		PresentMode:      chosenPresentMode,
		Clipped:          true,
	}, nil)
	if err != nil {
		return fmt.Errorf("create swapchain: %w", err)
	}
	defer device.DestroySwapchainKHR(swapchain, nil)
	fmt.Println("swapchain created")

	// Get swapchain images
	images, err := device.GetSwapchainImagesKHR(swapchain)
	if err != nil {
		return fmt.Errorf("get swapchain images: %w", err)
	}
	fmt.Printf("swapchain images: %d\n", len(images))

	fmt.Println("\nsurface test passed! window + swapchain working.")

	// Show the window briefly
	for i := 0; i < 60; i++ {
		if window.ShouldClose() {
			break
		}
		glfw.PollEvents()
	}

	_ = device.WaitIdle()

	return nil
}

// Ensure unsafe is imported for (*byte)(instance.Handle()) cast
var _ unsafe.Pointer
