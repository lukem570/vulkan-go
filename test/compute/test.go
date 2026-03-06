package main

import (
	"embed"
	"encoding/binary"
	"fmt"
	"log"
	"strings"
	"unsafe"

	vk "github.com/lukem570/vulkan-go/pkg/raw"
)

var _ embed.FS

func VkMakeApiVersion(variant, major, minor, patch uint32) uint32 {
	return (variant << 29) | (major << 22) | (minor << 12) | patch
}

func bytesToUint32(b []byte) []uint32 {
	n := len(b) / 4
	out := make([]uint32, n)

	for i := 0; i < n; i++ {
		out[i] = binary.LittleEndian.Uint32(b[i*4:])
	}

	return out
}

//go:embed test.spv
var computeShaderSPV []byte

const elementCount = 64
const bufferSize = elementCount * 4 // 64 uint32s = 256 bytes

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Initialize Vulkan via volk
	if err := vk.Initialize(); err != nil {
		return fmt.Errorf("initialize: %w", err)
	}

	// Create instance with validation layer and debug utils extension
	instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
		ApplicationInfo: &vk.ApplicationInfo{
			ApplicationName:    "vulkan-go-compute-test",
			ApplicationVersion: 1,
			EngineName:         "vulkan-go",
			EngineVersion:      1,
			ApiVersion:         VkMakeApiVersion(0, 1, 4, 0),
		},
		EnabledLayerNames: []string{
			"VK_LAYER_KHRONOS_validation",
		},
		EnabledExtensionNames: []string{
			"VK_EXT_debug_utils",
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create instance: %w", err)
	}
	defer instance.Destroy(nil)

	vk.LoadInstance(instance)
	fmt.Println("instance created with validation layers")

	// Find a physical device with a compute queue
	physicalDevices, err := instance.EnumeratePhysicalDevices()
	if err != nil {
		return fmt.Errorf("enumerate physical devices: %w", err)
	}

	var physicalDevice *vk.PhysicalDevice
	var computeQueueFamily uint32
	found := false

	for _, phy := range physicalDevices {
		props := phy.GetProperties()
		name := strings.TrimRight(string(props.DeviceName[:]), "\x00")

		queueFamilies := phy.GetQueueFamilyProperties()
		for i, qf := range queueFamilies {
			if qf.QueueFlags&vk.QueueFlags(vk.QueueComputeBit) != 0 {
				fmt.Printf("using: %s (queue family %d)\n", name, i)
				physicalDevice = phy
				computeQueueFamily = uint32(i)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		return fmt.Errorf("no physical device with compute queue found")
	}

	// Create logical device with one compute queue
	device, err := physicalDevice.CreateDevice(&vk.DeviceCreateInfo{
		QueueCreateInfos: []vk.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: computeQueueFamily,
				QueuePriorities:  []float32{1.0},
			},
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create device: %w", err)
	}
	defer device.Destroy(nil)

	vk.LoadDevice(device)
	fmt.Println("device created")

	queue := device.GetQueue(computeQueueFamily, 0)

	// Create storage buffer (host-visible, host-coherent)
	buffer, err := device.CreateBuffer(&vk.BufferCreateInfo{
		Size:        bufferSize,
		Usage:       vk.BufferUsageFlags(vk.BufferUsageStorageBufferBit),
		SharingMode: vk.SharingModeExclusive,
	}, nil)
	if err != nil {
		return fmt.Errorf("create buffer: %w", err)
	}
	defer device.DestroyBuffer(buffer, nil)

	memReqs := device.GetBufferMemoryRequirements(buffer)

	// Find host-visible, host-coherent memory type
	memProps := physicalDevice.GetMemoryProperties()
	memTypeIdx := uint32(0xFFFFFFFF)
	required := vk.MemoryPropertyFlags(vk.MemoryPropertyHostVisibleBit | vk.MemoryPropertyHostCoherentBit)
	for i := uint32(0); i < 32; i++ {
		if memReqs.MemoryTypeBits&(1<<i) == 0 {
			continue
		}
		if memProps.MemoryTypes[i].PropertyFlags&required == required {
			memTypeIdx = i
			break
		}
	}
	if memTypeIdx == 0xFFFFFFFF {
		return fmt.Errorf("no suitable memory type found")
	}

	memory, err := device.AllocateMemory(&vk.MemoryAllocateInfo{
		AllocationSize:  memReqs.Size,
		MemoryTypeIndex: memTypeIdx,
	}, nil)
	if err != nil {
		return fmt.Errorf("allocate memory: %w", err)
	}
	defer device.FreeMemory(memory, nil)

	if err := device.BindBufferMemory(buffer, memory, 0); err != nil {
		return fmt.Errorf("bind buffer memory: %w", err)
	}

	// Map memory and write input data: [1, 2, 3, ..., 64]
	var mappedPtr unsafe.Pointer
	if err := device.MapMemory(memory, 0, bufferSize, 0, &mappedPtr); err != nil {
		return fmt.Errorf("map memory: %w", err)
	}

	data := unsafe.Slice((*uint32)(mappedPtr), elementCount)
	for i := range data {
		data[i] = uint32(i + 1)
	}
	device.UnmapMemory(memory)
	fmt.Println("buffer initialized with [1..64]")

	// Create descriptor set layout
	descSetLayout, err := device.CreateDescriptorSetLayout(&vk.DescriptorSetLayoutCreateInfo{
		Bindings: []vk.DescriptorSetLayoutBinding{
			{
				Binding:         0,
				DescriptorType:  vk.DescriptorTypeStorageBuffer,
				DescriptorCount: 1,
				StageFlags:      vk.ShaderStageFlags(vk.ShaderStageComputeBit),
			},
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create descriptor set layout: %w", err)
	}
	defer device.DestroyDescriptorSetLayout(descSetLayout, nil)

	// Create pipeline layout
	pipelineLayout, err := device.CreatePipelineLayout(&vk.PipelineLayoutCreateInfo{
		SetLayouts: []*vk.DescriptorSetLayout{descSetLayout},
	}, nil)
	if err != nil {
		return fmt.Errorf("create pipeline layout: %w", err)
	}
	defer device.DestroyPipelineLayout(pipelineLayout, nil)

	// Create shader module
	shaderModule, err := device.CreateShaderModule(&vk.ShaderModuleCreateInfo{
		CodeSize: uintptr(len(computeShaderSPV)),
		Code:     bytesToUint32(computeShaderSPV),
	}, nil)
	if err != nil {
		return fmt.Errorf("create shader module: %w", err)
	}
	defer device.DestroyShaderModule(shaderModule, nil)

	// Create compute pipeline
	pipelines, err := device.CreateComputePipelines(nil, 1, []vk.ComputePipelineCreateInfo{
		{
			Stage: vk.PipelineShaderStageCreateInfo{
				Stage:  vk.ShaderStageComputeBit,
				Module: shaderModule,
				Name:   "main",
			},
			Layout: pipelineLayout,
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create compute pipeline: %w", err)
	}
	pipeline := pipelines[0]
	defer device.DestroyPipeline(pipeline, nil)
	fmt.Println("compute pipeline created")

	// Create descriptor pool and allocate descriptor set
	descPool, err := device.CreateDescriptorPool(&vk.DescriptorPoolCreateInfo{
		MaxSets: 1,
		PoolSizes: []vk.DescriptorPoolSize{
			{Type: vk.DescriptorTypeStorageBuffer, DescriptorCount: 1},
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create descriptor pool: %w", err)
	}
	defer device.DestroyDescriptorPool(descPool, nil)

	descSets, err := device.AllocateDescriptorSets(&vk.DescriptorSetAllocateInfo{
		DescriptorPool: descPool,
		SetLayouts:     []*vk.DescriptorSetLayout{descSetLayout},
	})
	if err != nil {
		return fmt.Errorf("allocate descriptor sets: %w", err)
	}
	descSet := descSets[0]

	// Update descriptor set to point to our buffer
	device.UpdateDescriptorSets(1, []vk.WriteDescriptorSet{
		{
			DstSet:          descSet,
			DstBinding:      0,
			DescriptorCount: 1,
			DescriptorType:  vk.DescriptorTypeStorageBuffer,
			BufferInfo: []vk.DescriptorBufferInfo{
				{Buffer: buffer, Offset: 0, Range: bufferSize},
			},
		},
	}, 0, nil)

	// Create command pool and command buffer
	cmdPool, err := device.CreateCommandPool(&vk.CommandPoolCreateInfo{
		QueueFamilyIndex: computeQueueFamily,
	}, nil)
	if err != nil {
		return fmt.Errorf("create command pool: %w", err)
	}
	defer device.DestroyCommandPool(cmdPool, nil)

	cmdBufs, err := device.AllocateCommandBuffers(&vk.CommandBufferAllocateInfo{
		CommandPool:        cmdPool,
		Level:              vk.CommandBufferLevelPrimary,
		CommandBufferCount: 1,
	})
	if err != nil {
		return fmt.Errorf("allocate command buffers: %w", err)
	}
	cmdBuf := cmdBufs[0]

	// Record command buffer
	if err := cmdBuf.Begin(&vk.CommandBufferBeginInfo{
		Flags: vk.CommandBufferUsageFlags(vk.CommandBufferUsageOneTimeSubmitBit),
	}); err != nil {
		return fmt.Errorf("begin command buffer: %w", err)
	}

	cmdBuf.BindPipeline(vk.PipelineBindPointCompute, pipeline)
	cmdBuf.BindDescriptorSets(
		vk.PipelineBindPointCompute,
		pipelineLayout,
		0, 1,
		[]*vk.DescriptorSet{descSet},
		0, nil,
	)
	cmdBuf.Dispatch(1, 1, 1) // 1 workgroup of 64 invocations

	if err := cmdBuf.End(); err != nil {
		return fmt.Errorf("end command buffer: %w", err)
	}

	// Create fence and submit
	fence, err := device.CreateFence(&vk.FenceCreateInfo{}, nil)
	if err != nil {
		return fmt.Errorf("create fence: %w", err)
	}
	defer device.DestroyFence(fence, nil)

	if err := queue.Submit(1, []vk.SubmitInfo{
		{
			CommandBuffers: []*vk.CommandBuffer{cmdBuf},
		},
	}, fence); err != nil {
		return fmt.Errorf("queue submit: %w", err)
	}

	if err := device.WaitForFences(1, []*vk.Fence{fence}, true, ^uint64(0)); err != nil {
		return fmt.Errorf("wait for fences: %w", err)
	}
	fmt.Println("compute shader executed")

	// Read back results
	if err := device.MapMemory(memory, 0, bufferSize, 0, &mappedPtr); err != nil {
		return fmt.Errorf("map memory (readback): %w", err)
	}

	result := unsafe.Slice((*uint32)(mappedPtr), elementCount)
	passed := true
	for i := range result {
		expected := uint32((i + 1) * 2)
		if result[i] != expected {
			fmt.Printf("FAIL: result[%d] = %d, expected %d\n", i, result[i], expected)
			passed = false
		}
	}
	device.UnmapMemory(memory)

	if passed {
		fmt.Println("PASS: all 64 elements correctly doubled")
	} else {
		return fmt.Errorf("compute shader produced incorrect results")
	}

	device.WaitIdle()
	return nil
}
