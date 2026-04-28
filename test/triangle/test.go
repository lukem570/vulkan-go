package main

import (
	"embed"
	"encoding/binary"
	"fmt"
	"log"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lukem570/vulkan-go/pkg/raw"
	"github.com/lukem570/vulkan-go/pkg/util"
)

var _ embed.FS

func bytesToUint32(b []byte) []uint32 {
	n := len(b) / 4
	out := make([]uint32, n)
	for i := range n {
		out[i] = binary.LittleEndian.Uint32(b[i*4:])
	}
	return out
}

//go:embed vert.spv
var vertShaderSPV []byte

//go:embed frag.spv
var fragShaderSPV []byte

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("glfw init: %w", err)
	}
	defer glfw.Terminate()

	if err := vk.Initialize(); err != nil {
		return fmt.Errorf("initialize: %w", err)
	}

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)

	window, err := glfw.CreateWindow(800, 600, "Vulkan go triangle", nil, nil)
	if err != nil {
		return fmt.Errorf("create window: %w", err)
	}
	glfwExts := window.GetRequiredInstanceExtensions()
	defer window.Destroy()

	enabledExts := append(glfwExts, "VK_EXT_debug_utils")

	messengerCreateInfo := &vk.DebugUtilsMessengerCreateInfoEXT{
		MessageSeverity: vk.DebugUtilsMessageSeverityFlagsEXT(
			vk.DebugUtilsMessageSeverityErrorBitEXT |
				vk.DebugUtilsMessageSeverityWarningBitEXT,
		),
		MessageType: vk.DebugUtilsMessageTypeFlagsEXT(
			vk.DebugUtilsMessageTypeValidationBitEXT |
				vk.DebugUtilsMessageTypePerformanceBitEXT,
		),
		UserCallback: func(messageSeverity vk.DebugUtilsMessageSeverityFlagBitsEXT, messageTypes vk.DebugUtilsMessageTypeFlagsEXT, callbackData *vk.DebugUtilsMessengerCallbackDataEXT, userData unsafe.Pointer) bool {
			fmt.Println("vulkan:", callbackData.Message)
			return false
		},
	}

	instNext := &vk.ValidationFeaturesEXT{
		Next: messengerCreateInfo,
		EnabledValidationFeatures: []vk.ValidationFeatureEnableEXT{
			vk.ValidationFeatureEnableBestPracticesEXT,
			vk.ValidationFeatureEnableSynchronizationValidationEXT,
			vk.ValidationFeatureEnableGpuAssistedEXT,
		},
	}

	instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
		Next: instNext,
		ApplicationInfo: &vk.ApplicationInfo{
			ApplicationName:    "vulkan-go-triangle-test",
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

	vk.LoadInstance(instance)

	messenger, err := instance.CreateDebugUtilsMessengerEXT(messengerCreateInfo, nil)
	if err != nil {
		return fmt.Errorf("create debug messenger: %w", err)
	}
	defer instance.DestroyDebugUtilsMessengerEXT(messenger, nil)

	fmt.Println("instance created")

	surfPtr, err := window.CreateWindowSurface((*byte)(instance.Handle()), nil)
	if err != nil {
		return fmt.Errorf("create surface: %w", err)
	}
	surface := vk.SurfaceKHRFromGLFW(surfPtr)
	defer instance.DestroySurfaceKHR(surface, nil)

	// Pick physical device
	physDevices, err := instance.EnumeratePhysicalDevices()
	if err != nil {
		return fmt.Errorf("enumerate physical devices: %w", err)
	}
	if len(physDevices) == 0 {
		return fmt.Errorf("no Vulkan-capable GPU found")
	}
	physDevice := physDevices[0]

	// Find graphics+present queue family
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
	vk.LoadDevice(device)

	queue := device.GetQueue(uint32(graphicsFamily), 0)

	// Query surface info
	caps, err := physDevice.GetSurfaceCapabilitiesKHR(surface)
	if err != nil {
		return fmt.Errorf("get surface capabilities: %w", err)
	}
	formats, err := physDevice.GetSurfaceFormatsKHR(surface)
	if err != nil {
		return fmt.Errorf("get surface formats: %w", err)
	}
	presentModes, err := physDevice.GetSurfacePresentModesKHR(surface)
	if err != nil {
		return fmt.Errorf("get present modes: %w", err)
	}

	// Choose format
	chosenFormat := formats[0]
	for _, f := range formats {
		if f.Format == vk.FormatB8g8r8a8Srgb && f.ColorSpace == vk.ColorSpaceSrgbNonlinearKHR {
			chosenFormat = f
			break
		}
	}

	// Choose present mode
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

	extent := caps.CurrentExtent

	// Create swapchain
	swapchain, err := device.CreateSwapchainKHR(&vk.SwapchainCreateInfoKHR{
		Surface:          surface,
		MinImageCount:    imageCount,
		ImageFormat:      chosenFormat.Format,
		ImageColorSpace:  chosenFormat.ColorSpace,
		ImageExtent:      extent,
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

	images, err := device.GetSwapchainImagesKHR(swapchain)
	if err != nil {
		return fmt.Errorf("get swapchain images: %w", err)
	}

	// Create image views
	imageViews := make([]*vk.ImageView, len(images))
	for i, img := range images {
		iv, err := device.CreateImageView(&vk.ImageViewCreateInfo{
			Image:    img,
			ViewType: vk.ImageViewType2d,
			Format:   chosenFormat.Format,
			Components: vk.ComponentMapping{
				R: vk.ComponentSwizzleIdentity,
				G: vk.ComponentSwizzleIdentity,
				B: vk.ComponentSwizzleIdentity,
				A: vk.ComponentSwizzleIdentity,
			},
			SubresourceRange: vk.ImageSubresourceRange{
				AspectMask:     vk.ImageAspectFlags(vk.ImageAspectColorBit),
				BaseMipLevel:   0,
				LevelCount:     1,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
		}, nil)
		if err != nil {
			return fmt.Errorf("create image view %d: %w", i, err)
		}
		imageViews[i] = iv
		defer device.DestroyImageView(iv, nil)
	}

	// Create render pass
	renderPass, err := device.CreateRenderPass(&vk.RenderPassCreateInfo{
		Attachments: []vk.AttachmentDescription{
			{
				Format:         chosenFormat.Format,
				Samples:        vk.SampleCount1Bit,
				LoadOp:         vk.AttachmentLoadOpClear,
				StoreOp:        vk.AttachmentStoreOpStore,
				StencilLoadOp:  vk.AttachmentLoadOpDontCare,
				StencilStoreOp: vk.AttachmentStoreOpDontCare,
				InitialLayout:  vk.ImageLayoutUndefined,
				FinalLayout:    vk.ImageLayoutPresentSrcKHR,
			},
		},
		Subpasses: []vk.SubpassDescription{
			{
				PipelineBindPoint:    vk.PipelineBindPointGraphics,
				ColorAttachmentCount: 1,
				ColorAttachments: []vk.AttachmentReference{
					{
						Attachment: 0,
						Layout:     vk.ImageLayoutColorAttachmentOptimal,
					},
				},
			},
		},
		Dependencies: []vk.SubpassDependency{
			{
				SrcSubpass:    vk.SubpassExternal,
				DstSubpass:    0,
				SrcStageMask:  vk.PipelineStageFlags(vk.PipelineStageColorAttachmentOutputBit),
				DstStageMask:  vk.PipelineStageFlags(vk.PipelineStageColorAttachmentOutputBit),
				SrcAccessMask: 0,
				DstAccessMask: vk.AccessFlags(vk.AccessColorAttachmentWriteBit),
			},
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create render pass: %w", err)
	}
	defer device.DestroyRenderPass(renderPass, nil)

	// Create shader modules
	vertModule, err := device.CreateShaderModule(&vk.ShaderModuleCreateInfo{
		Code:     bytesToUint32(vertShaderSPV),
	}, nil)
	if err != nil {
		return fmt.Errorf("create vertex shader module: %w", err)
	}
	defer device.DestroyShaderModule(vertModule, nil)

	fragModule, err := device.CreateShaderModule(&vk.ShaderModuleCreateInfo{
		Code:     bytesToUint32(fragShaderSPV),
	}, nil)
	if err != nil {
		return fmt.Errorf("create fragment shader module: %w", err)
	}
	defer device.DestroyShaderModule(fragModule, nil)

	// Create pipeline layout (no descriptors needed)
	pipelineLayout, err := device.CreatePipelineLayout(&vk.PipelineLayoutCreateInfo{}, nil)
	if err != nil {
		return fmt.Errorf("create pipeline layout: %w", err)
	}
	defer device.DestroyPipelineLayout(pipelineLayout, nil)

	// Create graphics pipeline
	pipelines, err := device.CreateGraphicsPipelines(nil, []vk.GraphicsPipelineCreateInfo{
		{
			StageCount: 2,
			Stages: []vk.PipelineShaderStageCreateInfo{
				{
					Stage:  vk.ShaderStageVertexBit,
					Module: vertModule,
					Name:   "main",
				},
				{
					Stage:  vk.ShaderStageFragmentBit,
					Module: fragModule,
					Name:   "main",
				},
			},
			VertexInputState: &vk.PipelineVertexInputStateCreateInfo{},
			InputAssemblyState: &vk.PipelineInputAssemblyStateCreateInfo{
				Topology: vk.PrimitiveTopologyTriangleList,
			},
			ViewportState: &vk.PipelineViewportStateCreateInfo{
				ViewportCount: 1,
				ScissorCount:  1,
			},
			RasterizationState: &vk.PipelineRasterizationStateCreateInfo{
				PolygonMode: vk.PolygonModeFill,
				CullMode:    vk.CullModeFlags(vk.CullModeBackBit),
				FrontFace:   vk.FrontFaceClockwise,
				LineWidth:   1.0,
			},
			MultisampleState: &vk.PipelineMultisampleStateCreateInfo{
				RasterizationSamples: vk.SampleCount1Bit,
			},
			ColorBlendState: &vk.PipelineColorBlendStateCreateInfo{
				AttachmentCount: 1,
				Attachments: []vk.PipelineColorBlendAttachmentState{
					{
						ColorWriteMask: vk.ColorComponentFlags(
							vk.ColorComponentRBit | vk.ColorComponentGBit |
								vk.ColorComponentBBit | vk.ColorComponentABit,
						),
					},
				},
			},
			DynamicState: &vk.PipelineDynamicStateCreateInfo{
				DynamicStates: []vk.DynamicState{
					vk.DynamicStateViewport,
					vk.DynamicStateScissor,
				},
			},
			Layout:     pipelineLayout,
			RenderPass: renderPass,
			Subpass:    0,
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("create graphics pipeline: %w", err)
	}
	pipeline := pipelines[0]
	defer device.DestroyPipeline(pipeline, nil)
	fmt.Println("graphics pipeline created")

	// Create framebuffers
	framebuffers := make([]*vk.Framebuffer, len(imageViews))
	for i, iv := range imageViews {
		fb, err := device.CreateFramebuffer(&vk.FramebufferCreateInfo{
			RenderPass:  renderPass,
			Attachments: []*vk.ImageView{iv},
			Width:       extent.Width,
			Height:      extent.Height,
			Layers:      1,
		}, nil)
		if err != nil {
			return fmt.Errorf("create framebuffer %d: %w", i, err)
		}
		framebuffers[i] = fb
		defer device.DestroyFramebuffer(fb, nil)
	}

	// Create command pool and command buffers
	cmdPool, err := device.CreateCommandPool(&vk.CommandPoolCreateInfo{
		Flags:            vk.CommandPoolCreateFlags(vk.CommandPoolCreateResetCommandBufferBit),
		QueueFamilyIndex: uint32(graphicsFamily),
	}, nil)
	if err != nil {
		return fmt.Errorf("create command pool: %w", err)
	}
	defer device.DestroyCommandPool(cmdPool, nil)

	cmdBufs, err := device.AllocateCommandBuffers(&vk.CommandBufferAllocateInfo{
		CommandPool:        cmdPool,
		Level:              vk.CommandBufferLevelPrimary,
		CommandBufferCount: uint32(len(images)),
	})
	if err != nil {
		return fmt.Errorf("allocate command buffers: %w", err)
	}

	// Create sync objects — one set per swapchain image to avoid semaphore reuse issues
	swapchainImageCount := len(images)
	imageAvailableSemaphores := make([]*vk.Semaphore, swapchainImageCount)
	renderFinishedSemaphores := make([]*vk.Semaphore, swapchainImageCount)
	inFlightFences := make([]*vk.Fence, swapchainImageCount)

	for i := range swapchainImageCount {
		sem1, err := device.CreateSemaphore(&vk.SemaphoreCreateInfo{}, nil)
		if err != nil {
			return fmt.Errorf("create semaphore: %w", err)
		}
		defer device.DestroySemaphore(sem1, nil)
		imageAvailableSemaphores[i] = sem1

		sem2, err := device.CreateSemaphore(&vk.SemaphoreCreateInfo{}, nil)
		if err != nil {
			return fmt.Errorf("create semaphore: %w", err)
		}
		defer device.DestroySemaphore(sem2, nil)
		renderFinishedSemaphores[i] = sem2

		fence, err := device.CreateFence(&vk.FenceCreateInfo{
			Flags: vk.FenceCreateFlags(vk.FenceCreateSignaledBit),
		}, nil)
		if err != nil {
			return fmt.Errorf("create fence: %w", err)
		}
		defer device.DestroyFence(fence, nil)
		inFlightFences[i] = fence
	}

	fmt.Println("rendering triangle...")

	currentFrame := 0
	for {
		if window.ShouldClose() {
			break
		}
		glfw.PollEvents()

		// Wait for previous frame
		if err := device.WaitForFences([]*vk.Fence{inFlightFences[currentFrame]}, true, ^uint64(0)); err != nil {
			return fmt.Errorf("wait for fence: %w", err)
		}
		if err := device.ResetFences([]*vk.Fence{inFlightFences[currentFrame]}); err != nil {
			return fmt.Errorf("reset fence: %w", err)
		}

		// Acquire next image
		imageIndex, err := device.AcquireNextImageKHR(swapchain, ^uint64(0), imageAvailableSemaphores[currentFrame], nil)
		if err != nil {
			return fmt.Errorf("acquire next image: %w", err)
		}

		// Record command buffer
		cmdBuf := cmdBufs[currentFrame]
		cmdBuf.Reset(0)

		if err := cmdBuf.Begin(&vk.CommandBufferBeginInfo{}); err != nil {
			return fmt.Errorf("begin command buffer: %w", err)
		}

		clearColor := vk.NewClearValueColor(vk.NewClearColorValueFloat32([4]float32{0.0, 0.0, 0.0, 1.0}))

		cmdBuf.BeginRenderPass(&vk.RenderPassBeginInfo{
			RenderPass:  renderPass,
			Framebuffer: framebuffers[imageIndex],
			RenderArea: vk.Rect2D{
				Offset: vk.Offset2D{X: 0, Y: 0},
				Extent: extent,
			},
			ClearValues: []vk.ClearValue{clearColor},
		}, vk.SubpassContentsInline)

		cmdBuf.BindPipeline(vk.PipelineBindPointGraphics, pipeline)

		cmdBuf.SetViewport(0, []vk.Viewport{
			{
				X:        0,
				Y:        0,
				Width:    float32(extent.Width),
				Height:   float32(extent.Height),
				MinDepth: 0.0,
				MaxDepth: 1.0,
			},
		})

		cmdBuf.SetScissor(0, []vk.Rect2D{
			{
				Offset: vk.Offset2D{X: 0, Y: 0},
				Extent: extent,
			},
		})

		cmdBuf.Draw(3, 1, 0, 0)

		cmdBuf.EndRenderPass()

		if err := cmdBuf.End(); err != nil {
			return fmt.Errorf("end command buffer: %w", err)
		}

		// Submit
		if err := queue.Submit([]vk.SubmitInfo{
			{
				WaitSemaphoreCount: 1,
				WaitSemaphores:     []*vk.Semaphore{imageAvailableSemaphores[currentFrame]},
				WaitDstStageMask:   []vk.PipelineStageFlags{vk.PipelineStageFlags(vk.PipelineStageColorAttachmentOutputBit)},
				CommandBuffers:     []*vk.CommandBuffer{cmdBuf},
				SignalSemaphores:   []*vk.Semaphore{renderFinishedSemaphores[imageIndex]},
			},
		}, inFlightFences[currentFrame]); err != nil {
			return fmt.Errorf("queue submit: %w", err)
		}

		// Present
		if err := queue.PresentKHR(&vk.PresentInfoKHR{
			WaitSemaphores: []*vk.Semaphore{renderFinishedSemaphores[imageIndex]},
			SwapchainCount: 1,
			Swapchains:     []*vk.SwapchainKHR{swapchain},
			ImageIndices:   []uint32{imageIndex},
		}); err != nil {
			return fmt.Errorf("queue present: %w", err)
		}

		currentFrame = (currentFrame + 1) % swapchainImageCount
	}

	_ = device.WaitIdle()
	fmt.Println("triangle test passed!")

	return nil
}

var _ unsafe.Pointer
