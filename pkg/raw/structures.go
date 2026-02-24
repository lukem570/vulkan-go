package vulkan

import "unsafe"

type VkSampleMask uint32

type VkBool32 uint32

type VkFlags uint32

type VkFlags64 uint64

type VkDeviceSize uint64

type VkDeviceAddress uint64

type VkFramebufferCreateFlags VkFlags

type VkQueryPoolCreateFlags VkFlags

type VkRenderPassCreateFlags VkFlags

type VkSamplerCreateFlags VkFlags

type VkPipelineLayoutCreateFlags VkFlags

type VkPipelineCacheCreateFlags VkFlags

type VkPipelineDepthStencilStateCreateFlags VkFlags

type VkPipelineDepthStencilStateCreateFlags VkFlags

type VkPipelineDynamicStateCreateFlags VkFlags

type VkPipelineColorBlendStateCreateFlags VkFlags

type VkPipelineColorBlendStateCreateFlags VkFlags

type VkPipelineMultisampleStateCreateFlags VkFlags

type VkPipelineRasterizationStateCreateFlags VkFlags

type VkPipelineViewportStateCreateFlags VkFlags

type VkPipelineTessellationStateCreateFlags VkFlags

type VkPipelineInputAssemblyStateCreateFlags VkFlags

type VkPipelineVertexInputStateCreateFlags VkFlags

type VkPipelineShaderStageCreateFlags VkFlags

type VkDescriptorSetLayoutCreateFlags VkFlags

type VkBufferViewCreateFlags VkFlags

type VkInstanceCreateFlags VkFlags

type VkDeviceCreateFlags VkFlags

type VkDeviceQueueCreateFlags VkFlags

type VkQueueFlags VkFlags

type VkMemoryPropertyFlags VkFlags

type VkMemoryHeapFlags VkFlags

type VkAccessFlags VkFlags

type VkBufferUsageFlags VkFlags

type VkBufferCreateFlags VkFlags

type VkShaderStageFlags VkFlags

type VkImageUsageFlags VkFlags

type VkImageCreateFlags VkFlags

type VkImageViewCreateFlags VkFlags

type VkPipelineCreateFlags VkFlags

type VkColorComponentFlags VkFlags

type VkFenceCreateFlags VkFlags

type VkSemaphoreCreateFlags VkFlags

type VkFormatFeatureFlags VkFlags

type VkQueryControlFlags VkFlags

type VkQueryResultFlags VkFlags

type VkShaderModuleCreateFlags VkFlags

type VkEventCreateFlags VkFlags

type VkCommandPoolCreateFlags VkFlags

type VkCommandPoolResetFlags VkFlags

type VkCommandBufferResetFlags VkFlags

type VkCommandBufferUsageFlags VkFlags

type VkQueryPipelineStatisticFlags VkFlags

type VkMemoryMapFlags VkFlags

type VkMemoryUnmapFlags VkFlags

type VkImageAspectFlags VkFlags

type VkSparseMemoryBindFlags VkFlags

type VkSparseImageFormatFlags VkFlags

type VkSubpassDescriptionFlags VkFlags

type VkPipelineStageFlags VkFlags

type VkSampleCountFlags VkFlags

type VkAttachmentDescriptionFlags VkFlags

type VkStencilFaceFlags VkFlags

type VkCullModeFlags VkFlags

type VkDescriptorPoolCreateFlags VkFlags

type VkDescriptorPoolResetFlags VkFlags

type VkDependencyFlags VkFlags

type VkSubgroupFeatureFlags VkFlags

type VkIndirectCommandsLayoutUsageFlagsNV VkFlags

type VkIndirectStateFlagsNV VkFlags

type VkGeometryFlagsKHR VkFlags

type VkGeometryInstanceFlagsKHR VkFlags

type VkClusterAccelerationStructureGeometryFlagsNV VkFlags

type VkClusterAccelerationStructureClusterFlagsNV VkFlags

type VkClusterAccelerationStructureAddressResolutionFlagsNV VkFlags

type VkBuildAccelerationStructureFlagsKHR VkFlags

type VkPrivateDataSlotCreateFlags VkFlags

type VkAccelerationStructureCreateFlagsKHR VkFlags

type VkDescriptorUpdateTemplateCreateFlags VkFlags

type VkPipelineCreationFeedbackFlags VkFlags

type VkPerformanceCounterDescriptionFlagsKHR VkFlags

type VkAcquireProfilingLockFlagsKHR VkFlags

type VkSemaphoreWaitFlags VkFlags

type VkPipelineCompilerControlFlagsAMD VkFlags

type VkShaderCorePropertiesFlagsAMD VkFlags

type VkDeviceDiagnosticsConfigFlagsNV VkFlags

type VkRefreshObjectFlagsKHR VkFlags

type VkAccessFlags2 VkFlags64

type VkPipelineStageFlags2 VkFlags64

type VkAccelerationStructureMotionInfoFlagsNV VkFlags

type VkAccelerationStructureMotionInstanceFlagsNV VkFlags

type VkFormatFeatureFlags2 VkFlags64

type VkRenderingFlags VkFlags

type VkMemoryDecompressionMethodFlagsEXT VkFlags64

type VkBuildMicromapFlagsEXT VkFlags

type VkMicromapCreateFlagsEXT VkFlags

type VkIndirectCommandsLayoutUsageFlagsEXT VkFlags

type VkIndirectCommandsInputModeFlagsEXT VkFlags

type VkDirectDriverLoadingFlagsLUNARG VkFlags

type VkPipelineCreateFlags2 VkFlags64

type VkBufferUsageFlags2 VkFlags64

type VkAddressCopyFlagsKHR VkFlags

type VkTensorCreateFlagsARM VkFlags64

type VkTensorUsageFlagsARM VkFlags64

type VkTensorViewCreateFlagsARM VkFlags64

type VkDataGraphPipelineSessionCreateFlagsARM VkFlags64

type VkDataGraphPipelineDispatchFlagsARM VkFlags64

type VkVideoEncodeRgbModelConversionFlagsVALVE VkFlags

type VkVideoEncodeRgbRangeCompressionFlagsVALVE VkFlags

type VkVideoEncodeRgbChromaOffsetFlagsVALVE VkFlags

type VkSpirvResourceTypeFlagsEXT VkFlags

type VkCompositeAlphaFlagsKHR VkFlags

type VkDisplayPlaneAlphaFlagsKHR VkFlags

type VkSurfaceTransformFlagsKHR VkFlags

type VkSwapchainCreateFlagsKHR VkFlags

type VkDisplayModeCreateFlagsKHR VkFlags

type VkDisplaySurfaceCreateFlagsKHR VkFlags

type VkAndroidSurfaceCreateFlagsKHR VkFlags

type VkViSurfaceCreateFlagsNN VkFlags

type VkWaylandSurfaceCreateFlagsKHR VkFlags

type VkUbmSurfaceCreateFlagsSEC VkFlags

type VkWin32SurfaceCreateFlagsKHR VkFlags

type VkXlibSurfaceCreateFlagsKHR VkFlags

type VkXcbSurfaceCreateFlagsKHR VkFlags

type VkDirectFBSurfaceCreateFlagsEXT VkFlags

type VkIOSSurfaceCreateFlagsMVK VkFlags

type VkMacOSSurfaceCreateFlagsMVK VkFlags

type VkMetalSurfaceCreateFlagsEXT VkFlags

type VkImagePipeSurfaceCreateFlagsFUCHSIA VkFlags

type VkStreamDescriptorSurfaceCreateFlagsGGP VkFlags

type VkHeadlessSurfaceCreateFlagsEXT VkFlags

type VkScreenSurfaceCreateFlagsQNX VkFlags

type VkPeerMemoryFeatureFlags VkFlags

type VkMemoryAllocateFlags VkFlags

type VkDeviceGroupPresentModeFlagsKHR VkFlags

type VkDebugReportFlagsEXT VkFlags

type VkCommandPoolTrimFlags VkFlags

type VkExternalMemoryHandleTypeFlagsNV VkFlags

type VkClusterAccelerationStructureIndexFormatFlagsNV VkFlags

type VkExternalMemoryFeatureFlagsNV VkFlags

type VkExternalMemoryHandleTypeFlags VkFlags

type VkExternalMemoryFeatureFlags VkFlags

type VkExternalSemaphoreHandleTypeFlags VkFlags

type VkExternalSemaphoreFeatureFlags VkFlags

type VkSemaphoreImportFlags VkFlags

type VkExternalFenceHandleTypeFlags VkFlags

type VkExternalFenceFeatureFlags VkFlags

type VkFenceImportFlags VkFlags

type VkSurfaceCounterFlagsEXT VkFlags

type VkPipelineViewportSwizzleStateCreateFlagsNV VkFlags

type VkPipelineDiscardRectangleStateCreateFlagsEXT VkFlags

type VkPipelineCoverageToColorStateCreateFlagsNV VkFlags

type VkPipelineCoverageModulationStateCreateFlagsNV VkFlags

type VkPipelineCoverageReductionStateCreateFlagsNV VkFlags

type VkValidationCacheCreateFlagsEXT VkFlags

type VkDebugUtilsMessageSeverityFlagsEXT VkFlags

type VkDebugUtilsMessageTypeFlagsEXT VkFlags

type VkDebugUtilsMessengerCreateFlagsEXT VkFlags

type VkDebugUtilsMessengerCallbackDataFlagsEXT VkFlags

type VkDeviceMemoryReportFlagsEXT VkFlags

type VkPipelineRasterizationConservativeStateCreateFlagsEXT VkFlags

type VkDescriptorBindingFlags VkFlags

type VkConditionalRenderingFlagsEXT VkFlags

type VkResolveModeFlags VkFlags

type VkPipelineRasterizationStateStreamCreateFlagsEXT VkFlags

type VkPipelineRasterizationDepthClipStateCreateFlagsEXT VkFlags

type VkSwapchainImageUsageFlagsANDROID VkFlags

type VkToolPurposeFlags VkFlags

type VkSubmitFlags VkFlags

type VkImageFormatConstraintsFlagsFUCHSIA VkFlags

type VkHostImageCopyFlags VkFlags

type VkPartitionedAccelerationStructureInstanceFlagsNV VkFlags

type VkImageConstraintsInfoFlagsFUCHSIA VkFlags

type VkGraphicsPipelineLibraryFlagsEXT VkFlags

type VkImageCompressionFlagsEXT VkFlags

type VkImageCompressionFixedRateFlagsEXT VkFlags

type VkExportMetalObjectTypeFlagsEXT VkFlags

type VkRenderingAttachmentFlagsKHR VkFlags

type VkResolveImageFlagsKHR VkFlags

type VkDeviceAddressBindingFlagsEXT VkFlags

type VkOpticalFlowGridSizeFlagsNV VkFlags

type VkOpticalFlowUsageFlagsNV VkFlags

type VkOpticalFlowSessionCreateFlagsNV VkFlags

type VkOpticalFlowExecuteFlagsNV VkFlags

type VkFrameBoundaryFlagsEXT VkFlags

type VkPresentScalingFlagsKHR VkFlags

type VkPresentGravityFlagsKHR VkFlags

type VkShaderCreateFlagsEXT VkFlags

type VkTileShadingRenderPassFlagsQCOM VkFlags

type VkPhysicalDeviceSchedulingControlsFlagsARM VkFlags64

type VkSurfaceCreateFlagsOHOS VkFlags

type VkPresentStageFlagsEXT VkFlags

type VkPastPresentationTimingFlagsEXT VkFlags

type VkPresentTimingInfoFlagsEXT VkFlags

type VkSwapchainImageUsageFlagsOHOS VkFlags

type VkPerformanceCounterDescriptionFlagsARM VkFlags

type VkVideoCodecOperationFlagsKHR VkFlags

type VkVideoCapabilityFlagsKHR VkFlags

type VkVideoSessionCreateFlagsKHR VkFlags

type VkVideoSessionParametersCreateFlagsKHR VkFlags

type VkVideoBeginCodingFlagsKHR VkFlags

type VkVideoEndCodingFlagsKHR VkFlags

type VkVideoCodingControlFlagsKHR VkFlags

type VkVideoDecodeUsageFlagsKHR VkFlags

type VkVideoDecodeCapabilityFlagsKHR VkFlags

type VkVideoDecodeFlagsKHR VkFlags

type VkVideoDecodeH264PictureLayoutFlagsKHR VkFlags

type VkVideoEncodeFlagsKHR VkFlags

type VkVideoEncodeUsageFlagsKHR VkFlags

type VkVideoEncodeContentFlagsKHR VkFlags

type VkVideoEncodeCapabilityFlagsKHR VkFlags

type VkVideoEncodeFeedbackFlagsKHR VkFlags

type VkVideoEncodeRateControlFlagsKHR VkFlags

type VkVideoEncodeRateControlModeFlagsKHR VkFlags

type VkVideoEncodeIntraRefreshModeFlagsKHR VkFlags

type VkVideoChromaSubsamplingFlagsKHR VkFlags

type VkVideoComponentBitDepthFlagsKHR VkFlags

type VkVideoEncodeH264CapabilityFlagsKHR VkFlags

type VkVideoEncodeH264StdFlagsKHR VkFlags

type VkVideoEncodeH264RateControlFlagsKHR VkFlags

type VkVideoEncodeH265CapabilityFlagsKHR VkFlags

type VkVideoEncodeH265StdFlagsKHR VkFlags

type VkVideoEncodeH265RateControlFlagsKHR VkFlags

type VkVideoEncodeH265CtbSizeFlagsKHR VkFlags

type VkVideoEncodeH265TransformBlockSizeFlagsKHR VkFlags

type VkVideoEncodeAV1CapabilityFlagsKHR VkFlags

type VkVideoEncodeAV1StdFlagsKHR VkFlags

type VkVideoEncodeAV1RateControlFlagsKHR VkFlags

type VkVideoEncodeAV1SuperblockSizeFlagsKHR VkFlags

type VkAccessFlags3KHR VkFlags64

type VkBaseOutStructure struct {
	SType VkStructureType
	PNext *VkBaseOutStructure
}

type VkBaseInStructure struct {
	SType VkStructureType
	PNext *VkBaseInStructure
}

type VkOffset2D struct {
	X int32
	Y int32
}

type VkOffset3D struct {
	X int32
	Y int32
	Z int32
}

type VkExtent2D struct {
	Width  uint32
	Height uint32
}

type VkExtent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

type VkViewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

type VkRect2D struct {
	Offset VkOffset2D
	Extent VkExtent2D
}

type VkClearRect struct {
	Rect           VkRect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

type VkComponentMapping struct {
	R VkComponentSwizzle
	G VkComponentSwizzle
	B VkComponentSwizzle
	A VkComponentSwizzle
}

type VkPhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        VkPhysicalDeviceType
	DeviceName        char
	PipelineCacheUUID uint8
	Limits            VkPhysicalDeviceLimits
	SparseProperties  VkPhysicalDeviceSparseProperties
}

type VkExtensionProperties struct {
	ExtensionName char
	SpecVersion   uint32
}

type VkLayerProperties struct {
	LayerName             char
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           char
}

type VkApplicationInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PApplicationName   string
	ApplicationVersion uint32
	PEngineName        string
	EngineVersion      uint32
	ApiVersion         uint32
}

type VkAllocationCallbacks struct {
	PUserData             unsafe.Pointer
	PfnAllocation         PFN_vkAllocationFunction
	PfnReallocation       PFN_vkReallocationFunction
	PfnFree               PFN_vkFreeFunction
	PfnInternalAllocation PFN_vkInternalAllocationNotification
	PfnInternalFree       PFN_vkInternalFreeNotification
}

type VkDeviceQueueCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueCount       uint32
	PQueuePriorities *float32
}

type VkDeviceCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkDeviceCreateFlags
	QueueCreateInfoCount    uint32
	PQueueCreateInfos       *VkDeviceQueueCreateInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     *string
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames *string
	PEnabledFeatures        *VkPhysicalDeviceFeatures
}

type VkInstanceCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkInstanceCreateFlags
	PApplicationInfo        *VkApplicationInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     *string
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames *string
}

type VkQueueFamilyProperties struct {
	QueueFlags                  VkQueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity VkExtent3D
}

type VkPhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     VkMemoryType
	MemoryHeapCount uint32
	MemoryHeaps     VkMemoryHeap
}

type VkMemoryAllocateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	AllocationSize  VkDeviceSize
	MemoryTypeIndex uint32
}

type VkMemoryRequirements struct {
	Size           VkDeviceSize
	Alignment      VkDeviceSize
	MemoryTypeBits uint32
}

type VkSparseImageFormatProperties struct {
	AspectMask       VkImageAspectFlags
	ImageGranularity VkExtent3D
	Flags            VkSparseImageFormatFlags
}

type VkSparseImageMemoryRequirements struct {
	FormatProperties     VkSparseImageFormatProperties
	ImageMipTailFirstLod uint32
	ImageMipTailSize     VkDeviceSize
	ImageMipTailOffset   VkDeviceSize
	ImageMipTailStride   VkDeviceSize
}

type VkMemoryType struct {
	PropertyFlags VkMemoryPropertyFlags
	HeapIndex     uint32
}

type VkMemoryHeap struct {
	Size  VkDeviceSize
	Flags VkMemoryHeapFlags
}

type VkMappedMemoryRange struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
	Offset VkDeviceSize
	Size   VkDeviceSize
}

type VkFormatProperties struct {
	LinearTilingFeatures  VkFormatFeatureFlags
	OptimalTilingFeatures VkFormatFeatureFlags
	BufferFeatures        VkFormatFeatureFlags
}

type VkImageFormatProperties struct {
	MaxExtent       VkExtent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    VkSampleCountFlags
	MaxResourceSize VkDeviceSize
}

type VkDescriptorBufferInfo struct {
	Buffer VkBuffer
	Offset VkDeviceSize
	Range  VkDeviceSize
}

type VkDescriptorImageInfo struct {
	Sampler     VkSampler
	ImageView   VkImageView
	ImageLayout VkImageLayout
}

type VkWriteDescriptorSet struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DstSet           VkDescriptorSet
	DstBinding       uint32
	DstArrayElement  uint32
	DescriptorCount  uint32
	DescriptorType   VkDescriptorType
	PImageInfo       *VkDescriptorImageInfo
	PBufferInfo      *VkDescriptorBufferInfo
	PTexelBufferView *VkBufferView
}

type VkCopyDescriptorSet struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	SrcSet          VkDescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          VkDescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}

type VkBufferUsageFlags2CreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkBufferUsageFlags2
}

type VkBufferUsageFlags2CreateInfoKHR struct {
}

type VkBufferCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkBufferCreateFlags
	Size                  VkDeviceSize
	Usage                 VkBufferUsageFlags
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}

type VkBufferViewCreateInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkBufferViewCreateFlags
	Buffer VkBuffer
	Format VkFormat
	Offset VkDeviceSize
	Range  VkDeviceSize
}

type VkImageSubresource struct {
	AspectMask VkImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}

type VkImageSubresourceLayers struct {
	AspectMask     VkImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type VkImageSubresourceRange struct {
	AspectMask     VkImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type VkMemoryBarrier struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SrcAccessMask VkAccessFlags
	DstAccessMask VkAccessFlags
}

type VkBufferMemoryBarrier struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	SrcAccessMask       VkAccessFlags
	DstAccessMask       VkAccessFlags
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              VkBuffer
	Offset              VkDeviceSize
	Size                VkDeviceSize
}

type VkImageMemoryBarrier struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	SrcAccessMask       VkAccessFlags
	DstAccessMask       VkAccessFlags
	OldLayout           VkImageLayout
	NewLayout           VkImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               VkImage
	SubresourceRange    VkImageSubresourceRange
}

type VkImageCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkImageCreateFlags
	ImageType             VkImageType
	Format                VkFormat
	Extent                VkExtent3D
	MipLevels             uint32
	ArrayLayers           uint32
	Samples               VkSampleCountFlagBits
	Tiling                VkImageTiling
	Usage                 VkImageUsageFlags
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	InitialLayout         VkImageLayout
}

type VkSubresourceLayout struct {
	Offset     VkDeviceSize
	Size       VkDeviceSize
	RowPitch   VkDeviceSize
	ArrayPitch VkDeviceSize
	DepthPitch VkDeviceSize
}

type VkImageViewCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkImageViewCreateFlags
	Image            VkImage
	ViewType         VkImageViewType
	Format           VkFormat
	Components       VkComponentMapping
	SubresourceRange VkImageSubresourceRange
}

type VkBufferCopy struct {
	SrcOffset VkDeviceSize
	DstOffset VkDeviceSize
	Size      VkDeviceSize
}

type VkSparseMemoryBind struct {
	ResourceOffset VkDeviceSize
	Size           VkDeviceSize
	Memory         VkDeviceMemory
	MemoryOffset   VkDeviceSize
	Flags          VkSparseMemoryBindFlags
}

type VkSparseImageMemoryBind struct {
	Subresource  VkImageSubresource
	Offset       VkOffset3D
	Extent       VkExtent3D
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
	Flags        VkSparseMemoryBindFlags
}

type VkSparseBufferMemoryBindInfo struct {
	Buffer    VkBuffer
	BindCount uint32
	PBinds    *VkSparseMemoryBind
}

type VkSparseImageOpaqueMemoryBindInfo struct {
	Image     VkImage
	BindCount uint32
	PBinds    *VkSparseMemoryBind
}

type VkSparseImageMemoryBindInfo struct {
	Image     VkImage
	BindCount uint32
	PBinds    *VkSparseImageMemoryBind
}

type VkBindSparseInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	WaitSemaphoreCount   uint32
	PWaitSemaphores      *VkSemaphore
	BufferBindCount      uint32
	PBufferBinds         *VkSparseBufferMemoryBindInfo
	ImageOpaqueBindCount uint32
	PImageOpaqueBinds    *VkSparseImageOpaqueMemoryBindInfo
	ImageBindCount       uint32
	PImageBinds          *VkSparseImageMemoryBindInfo
	SignalSemaphoreCount uint32
	PSignalSemaphores    *VkSemaphore
}

type VkImageCopy struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}

type VkImageBlit struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffsets     VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffsets     VkOffset3D
}

type VkBufferImageCopy struct {
	BufferOffset      VkDeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}

type VkStridedDeviceAddressRangeKHR struct {
	Address VkDeviceAddress
	Size    VkDeviceSize
	Stride  VkDeviceSize
}

type VkCopyMemoryIndirectCommandKHR struct {
	SrcAddress VkDeviceAddress
	DstAddress VkDeviceAddress
	Size       VkDeviceSize
}

type VkCopyMemoryIndirectCommandNV struct {
}

type VkCopyMemoryIndirectInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	SrcCopyFlags     VkAddressCopyFlagsKHR
	DstCopyFlags     VkAddressCopyFlagsKHR
	CopyCount        uint32
	CopyAddressRange VkStridedDeviceAddressRangeKHR
}

type VkCopyMemoryToImageIndirectCommandKHR struct {
	SrcAddress        VkDeviceAddress
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}

type VkCopyMemoryToImageIndirectCommandNV struct {
}

type VkCopyMemoryToImageIndirectInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	SrcCopyFlags       VkAddressCopyFlagsKHR
	CopyCount          uint32
	CopyAddressRange   VkStridedDeviceAddressRangeKHR
	DstImage           VkImage
	DstImageLayout     VkImageLayout
	PImageSubresources *VkImageSubresourceLayers
}

type VkImageResolve struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}

type VkShaderModuleCreateInfo struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Flags    VkShaderModuleCreateFlags
	CodeSize uint
	PCode    *uint32
}

type VkDescriptorSetLayoutBinding struct {
	Binding            uint32
	DescriptorType     VkDescriptorType
	DescriptorCount    uint32
	StageFlags         VkShaderStageFlags
	PImmutableSamplers *VkSampler
}

type VkDescriptorSetLayoutCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Flags        VkDescriptorSetLayoutCreateFlags
	BindingCount uint32
	PBindings    *VkDescriptorSetLayoutBinding
}

type VkDescriptorPoolSize struct {
	Type            VkDescriptorType
	DescriptorCount uint32
}

type VkDescriptorPoolCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Flags         VkDescriptorPoolCreateFlags
	MaxSets       uint32
	PoolSizeCount uint32
	PPoolSizes    *VkDescriptorPoolSize
}

type VkDescriptorSetAllocateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	DescriptorPool     VkDescriptorPool
	DescriptorSetCount uint32
	PSetLayouts        *VkDescriptorSetLayout
}

type VkSpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uint
}

type VkSpecializationInfo struct {
	MapEntryCount uint32
	PMapEntries   *VkSpecializationMapEntry
	DataSize      uint
	PData         unsafe.Pointer
}

type VkPipelineShaderStageCreateInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkPipelineShaderStageCreateFlags
	Stage               VkShaderStageFlagBits
	Module              VkShaderModule
	PName               string
	PName               string
	PSpecializationInfo *VkSpecializationInfo
}

type VkComputePipelineCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkPipelineCreateFlags
	Stage              VkPipelineShaderStageCreateInfo
	Layout             VkPipelineLayout
	BasePipelineHandle VkPipeline
	BasePipelineIndex  int32
}

type VkComputePipelineIndirectBufferInfoNV struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	DeviceAddress                      VkDeviceAddress
	Size                               VkDeviceSize
	PipelineDeviceAddressCaptureReplay VkDeviceAddress
}

type VkPipelineCreateFlags2CreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkPipelineCreateFlags2
}

type VkPipelineCreateFlags2CreateInfoKHR struct {
}

type VkVertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VkVertexInputRate
}

type VkVertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   VkFormat
	Offset   uint32
}

type VkPipelineVertexInputStateCreateInfo struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	Flags                           VkPipelineVertexInputStateCreateFlags
	VertexBindingDescriptionCount   uint32
	PVertexBindingDescriptions      *VkVertexInputBindingDescription
	VertexAttributeDescriptionCount uint32
	PVertexAttributeDescriptions    *VkVertexInputAttributeDescription
}

type VkPipelineInputAssemblyStateCreateInfo struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkPipelineInputAssemblyStateCreateFlags
	Topology               VkPrimitiveTopology
	PrimitiveRestartEnable VkBool32
}

type VkPipelineTessellationStateCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkPipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}

type VkPipelineViewportStateCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Flags         VkPipelineViewportStateCreateFlags
	ViewportCount uint32
	PViewports    *VkViewport
	ScissorCount  uint32
	PScissors     *VkRect2D
}

type VkPipelineRasterizationStateCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkPipelineRasterizationStateCreateFlags
	DepthClampEnable        VkBool32
	RasterizerDiscardEnable VkBool32
	PolygonMode             VkPolygonMode
	CullMode                VkCullModeFlags
	FrontFace               VkFrontFace
	DepthBiasEnable         VkBool32
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

type VkPipelineMultisampleStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkPipelineMultisampleStateCreateFlags
	RasterizationSamples  VkSampleCountFlagBits
	SampleShadingEnable   VkBool32
	MinSampleShading      float32
	PSampleMask           *VkSampleMask
	AlphaToCoverageEnable VkBool32
	AlphaToOneEnable      VkBool32
}

type VkPipelineColorBlendAttachmentState struct {
	BlendEnable         VkBool32
	SrcColorBlendFactor VkBlendFactor
	DstColorBlendFactor VkBlendFactor
	ColorBlendOp        VkBlendOp
	SrcAlphaBlendFactor VkBlendFactor
	DstAlphaBlendFactor VkBlendFactor
	AlphaBlendOp        VkBlendOp
	ColorWriteMask      VkColorComponentFlags
}

type VkPipelineColorBlendStateCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkPipelineColorBlendStateCreateFlags
	LogicOpEnable   VkBool32
	LogicOp         VkLogicOp
	AttachmentCount uint32
	PAttachments    *VkPipelineColorBlendAttachmentState
	BlendConstants  float32
}

type VkPipelineDynamicStateCreateInfo struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Flags             VkPipelineDynamicStateCreateFlags
	DynamicStateCount uint32
	PDynamicStates    *VkDynamicState
}

type VkStencilOpState struct {
	FailOp      VkStencilOp
	PassOp      VkStencilOp
	DepthFailOp VkStencilOp
	CompareOp   VkCompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

type VkPipelineDepthStencilStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkPipelineDepthStencilStateCreateFlags
	DepthTestEnable       VkBool32
	DepthWriteEnable      VkBool32
	DepthCompareOp        VkCompareOp
	DepthBoundsTestEnable VkBool32
	StencilTestEnable     VkBool32
	Front                 VkStencilOpState
	Back                  VkStencilOpState
	MinDepthBounds        float32
	MaxDepthBounds        float32
}

type VkGraphicsPipelineCreateInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkPipelineCreateFlags
	StageCount          uint32
	PStages             *VkPipelineShaderStageCreateInfo
	PStages             *VkPipelineShaderStageCreateInfo
	PVertexInputState   *VkPipelineVertexInputStateCreateInfo
	PInputAssemblyState *VkPipelineInputAssemblyStateCreateInfo
	PTessellationState  *VkPipelineTessellationStateCreateInfo
	PViewportState      *VkPipelineViewportStateCreateInfo
	PRasterizationState *VkPipelineRasterizationStateCreateInfo
	PMultisampleState   *VkPipelineMultisampleStateCreateInfo
	PDepthStencilState  *VkPipelineDepthStencilStateCreateInfo
	PColorBlendState    *VkPipelineColorBlendStateCreateInfo
	PDynamicState       *VkPipelineDynamicStateCreateInfo
	Layout              VkPipelineLayout
	RenderPass          VkRenderPass
	Subpass             uint32
	BasePipelineHandle  VkPipeline
	BasePipelineIndex   int32
}

type VkPipelineCacheCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkPipelineCacheCreateFlags
	InitialDataSize uint
	InitialDataSize uint
	PInitialData    unsafe.Pointer
}

type VkPipelineCacheHeaderVersionOne struct {
	HeaderSize        uint32
	HeaderVersion     VkPipelineCacheHeaderVersion
	VendorID          uint32
	DeviceID          uint32
	PipelineCacheUUID uint8
}

type VkPipelineCacheStageValidationIndexEntry struct {
	CodeSize   uint64
	CodeOffset uint64
}

type VkPipelineCacheSafetyCriticalIndexEntry struct {
	PipelineIdentifier uint8
	PipelineMemorySize uint64
	JsonSize           uint64
	JsonOffset         uint64
	StageIndexCount    uint32
	StageIndexStride   uint32
	StageIndexOffset   uint64
}

type VkPipelineCacheHeaderVersionSafetyCriticalOne struct {
	HeaderVersionOne    VkPipelineCacheHeaderVersionOne
	ValidationVersion   VkPipelineCacheValidationVersion
	ImplementationData  uint32
	PipelineIndexCount  uint32
	PipelineIndexStride uint32
	PipelineIndexOffset uint64
}

type VkPipelineCacheHeaderVersionDataGraphQCOM struct {
	HeaderSize       uint32
	HeaderVersion    VkPipelineCacheHeaderVersion
	CacheType        VkDataGraphModelCacheTypeQCOM
	CacheVersion     uint32
	ToolchainVersion uint32
}

type VkPushConstantRange struct {
	StageFlags VkShaderStageFlags
	Offset     uint32
	Size       uint32
}

type VkPipelineBinaryCreateInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	PKeysAndDataInfo    *VkPipelineBinaryKeysAndDataKHR
	Pipeline            VkPipeline
	PPipelineCreateInfo *VkPipelineCreateInfoKHR
}

type VkPipelineBinaryHandlesInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	PipelineBinaryCount uint32
	PPipelineBinaries   *VkPipelineBinaryKHR
}

type VkPipelineBinaryDataKHR struct {
	DataSize uint
	PData    unsafe.Pointer
}

type VkPipelineBinaryKeysAndDataKHR struct {
	BinaryCount         uint32
	PPipelineBinaryKeys *VkPipelineBinaryKeyKHR
	PPipelineBinaryData *VkPipelineBinaryDataKHR
}

type VkPipelineBinaryKeyKHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	KeySize uint32
	Key     uint8
}

type VkPipelineBinaryInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	BinaryCount       uint32
	PPipelineBinaries *VkPipelineBinaryKHR
}

type VkReleaseCapturedPipelineDataInfoKHR struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Pipeline VkPipeline
}

type VkPipelineBinaryDataInfoKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PipelineBinary VkPipelineBinaryKHR
}

type VkPipelineCreateInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkPipelineLayoutCreateInfo struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkPipelineLayoutCreateFlags
	SetLayoutCount         uint32
	PSetLayouts            *VkDescriptorSetLayout
	PushConstantRangeCount uint32
	PPushConstantRanges    *VkPushConstantRange
}

type VkSamplerCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkSamplerCreateFlags
	MagFilter               VkFilter
	MinFilter               VkFilter
	MipmapMode              VkSamplerMipmapMode
	AddressModeU            VkSamplerAddressMode
	AddressModeV            VkSamplerAddressMode
	AddressModeW            VkSamplerAddressMode
	MipLodBias              float32
	AnisotropyEnable        VkBool32
	MaxAnisotropy           float32
	CompareEnable           VkBool32
	CompareOp               VkCompareOp
	MinLod                  float32
	MaxLod                  float32
	BorderColor             VkBorderColor
	UnnormalizedCoordinates VkBool32
}

type VkCommandPoolCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkCommandPoolCreateFlags
	QueueFamilyIndex uint32
}

type VkCommandBufferAllocateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	CommandPool        VkCommandPool
	Level              VkCommandBufferLevel
	CommandBufferCount uint32
}

type VkCommandBufferInheritanceInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	RenderPass           VkRenderPass
	Subpass              uint32
	Framebuffer          VkFramebuffer
	OcclusionQueryEnable VkBool32
	QueryFlags           VkQueryControlFlags
	PipelineStatistics   VkQueryPipelineStatisticFlags
}

type VkCommandBufferBeginInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkCommandBufferUsageFlags
	PInheritanceInfo *VkCommandBufferInheritanceInfo
}

type VkRenderPassBeginInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	RenderPass      VkRenderPass
	Framebuffer     VkFramebuffer
	RenderArea      VkRect2D
	ClearValueCount uint32
	PClearValues    *VkClearValue
}

type VkClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}

type VkClearAttachment struct {
	AspectMask      VkImageAspectFlags
	ColorAttachment uint32
	ClearValue      VkClearValue
}

type VkAttachmentDescription struct {
	Flags          VkAttachmentDescriptionFlags
	Format         VkFormat
	Samples        VkSampleCountFlagBits
	LoadOp         VkAttachmentLoadOp
	StoreOp        VkAttachmentStoreOp
	StencilLoadOp  VkAttachmentLoadOp
	StencilStoreOp VkAttachmentStoreOp
	InitialLayout  VkImageLayout
	FinalLayout    VkImageLayout
}

type VkAttachmentReference struct {
	Attachment uint32
	Layout     VkImageLayout
}

type VkSubpassDescription struct {
	Flags                   VkSubpassDescriptionFlags
	PipelineBindPoint       VkPipelineBindPoint
	InputAttachmentCount    uint32
	PInputAttachments       *VkAttachmentReference
	ColorAttachmentCount    uint32
	PColorAttachments       *VkAttachmentReference
	PResolveAttachments     *VkAttachmentReference
	PDepthStencilAttachment *VkAttachmentReference
	PreserveAttachmentCount uint32
	PPreserveAttachments    *uint32
}

type VkSubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    VkPipelineStageFlags
	DstStageMask    VkPipelineStageFlags
	SrcAccessMask   VkAccessFlags
	DstAccessMask   VkAccessFlags
	DependencyFlags VkDependencyFlags
}

type VkRenderPassCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkRenderPassCreateFlags
	AttachmentCount uint32
	PAttachments    *VkAttachmentDescription
	SubpassCount    uint32
	PSubpasses      *VkSubpassDescription
	DependencyCount uint32
	PDependencies   *VkSubpassDependency
}

type VkEventCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkEventCreateFlags
}

type VkFenceCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkFenceCreateFlags
}

type VkPhysicalDeviceFeatures struct {
	RobustBufferAccess                      VkBool32
	FullDrawIndexUint32                     VkBool32
	ImageCubeArray                          VkBool32
	IndependentBlend                        VkBool32
	GeometryShader                          VkBool32
	TessellationShader                      VkBool32
	SampleRateShading                       VkBool32
	DualSrcBlend                            VkBool32
	LogicOp                                 VkBool32
	MultiDrawIndirect                       VkBool32
	DrawIndirectFirstInstance               VkBool32
	DepthClamp                              VkBool32
	DepthBiasClamp                          VkBool32
	FillModeNonSolid                        VkBool32
	DepthBounds                             VkBool32
	WideLines                               VkBool32
	LargePoints                             VkBool32
	AlphaToOne                              VkBool32
	MultiViewport                           VkBool32
	SamplerAnisotropy                       VkBool32
	TextureCompressionETC2                  VkBool32
	TextureCompressionASTC_LDR              VkBool32
	TextureCompressionBC                    VkBool32
	OcclusionQueryPrecise                   VkBool32
	PipelineStatisticsQuery                 VkBool32
	VertexPipelineStoresAndAtomics          VkBool32
	FragmentStoresAndAtomics                VkBool32
	ShaderTessellationAndGeometryPointSize  VkBool32
	ShaderImageGatherExtended               VkBool32
	ShaderStorageImageExtendedFormats       VkBool32
	ShaderStorageImageMultisample           VkBool32
	ShaderStorageImageReadWithoutFormat     VkBool32
	ShaderStorageImageWriteWithoutFormat    VkBool32
	ShaderUniformBufferArrayDynamicIndexing VkBool32
	ShaderSampledImageArrayDynamicIndexing  VkBool32
	ShaderStorageBufferArrayDynamicIndexing VkBool32
	ShaderStorageImageArrayDynamicIndexing  VkBool32
	ShaderClipDistance                      VkBool32
	ShaderCullDistance                      VkBool32
	ShaderFloat64                           VkBool32
	ShaderInt64                             VkBool32
	ShaderInt16                             VkBool32
	ShaderResourceResidency                 VkBool32
	ShaderResourceMinLod                    VkBool32
	SparseBinding                           VkBool32
	SparseResidencyBuffer                   VkBool32
	SparseResidencyImage2D                  VkBool32
	SparseResidencyImage3D                  VkBool32
	SparseResidency2Samples                 VkBool32
	SparseResidency4Samples                 VkBool32
	SparseResidency8Samples                 VkBool32
	SparseResidency16Samples                VkBool32
	SparseResidencyAliased                  VkBool32
	VariableMultisampleRate                 VkBool32
	InheritedQueries                        VkBool32
}

type VkPhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            VkBool32
	ResidencyStandard2DMultisampleBlockShape VkBool32
	ResidencyStandard3DBlockShape            VkBool32
	ResidencyAlignedMipSize                  VkBool32
	ResidencyNonResidentStrict               VkBool32
}

type VkPhysicalDeviceLimits struct {
	MaxImageDimension1D                             uint32
	MaxImageDimension2D                             uint32
	MaxImageDimension3D                             uint32
	MaxImageDimensionCube                           uint32
	MaxImageArrayLayers                             uint32
	MaxTexelBufferElements                          uint32
	MaxUniformBufferRange                           uint32
	MaxStorageBufferRange                           uint32
	MaxPushConstantsSize                            uint32
	MaxMemoryAllocationCount                        uint32
	MaxSamplerAllocationCount                       uint32
	BufferImageGranularity                          VkDeviceSize
	SparseAddressSpaceSize                          VkDeviceSize
	MaxBoundDescriptorSets                          uint32
	MaxPerStageDescriptorSamplers                   uint32
	MaxPerStageDescriptorUniformBuffers             uint32
	MaxPerStageDescriptorStorageBuffers             uint32
	MaxPerStageDescriptorSampledImages              uint32
	MaxPerStageDescriptorStorageImages              uint32
	MaxPerStageDescriptorInputAttachments           uint32
	MaxPerStageResources                            uint32
	MaxDescriptorSetSamplers                        uint32
	MaxDescriptorSetUniformBuffers                  uint32
	MaxDescriptorSetUniformBuffersDynamic           uint32
	MaxDescriptorSetStorageBuffers                  uint32
	MaxDescriptorSetStorageBuffersDynamic           uint32
	MaxDescriptorSetSampledImages                   uint32
	MaxDescriptorSetStorageImages                   uint32
	MaxDescriptorSetInputAttachments                uint32
	MaxVertexInputAttributes                        uint32
	MaxVertexInputBindings                          uint32
	MaxVertexInputAttributeOffset                   uint32
	MaxVertexInputBindingStride                     uint32
	MaxVertexOutputComponents                       uint32
	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32
	MaxGeometryShaderInvocations                    uint32
	MaxGeometryInputComponents                      uint32
	MaxGeometryOutputComponents                     uint32
	MaxGeometryOutputVertices                       uint32
	MaxGeometryTotalOutputComponents                uint32
	MaxFragmentInputComponents                      uint32
	MaxFragmentOutputAttachments                    uint32
	MaxFragmentDualSrcAttachments                   uint32
	MaxFragmentCombinedOutputResources              uint32
	MaxComputeSharedMemorySize                      uint32
	MaxComputeWorkGroupCount                        uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           uint32
	ViewportBoundsRange                             float32
	ViewportSubPixelBits                            uint32
	MinMemoryMapAlignment                           uint
	MinTexelBufferOffsetAlignment                   VkDeviceSize
	MinUniformBufferOffsetAlignment                 VkDeviceSize
	MinStorageBufferOffsetAlignment                 VkDeviceSize
	MinTexelOffset                                  int32
	MaxTexelOffset                                  uint32
	MinTexelGatherOffset                            int32
	MaxTexelGatherOffset                            uint32
	MinInterpolationOffset                          float32
	MaxInterpolationOffset                          float32
	SubPixelInterpolationOffsetBits                 uint32
	MaxFramebufferWidth                             uint32
	MaxFramebufferHeight                            uint32
	MaxFramebufferLayers                            uint32
	FramebufferColorSampleCounts                    VkSampleCountFlags
	FramebufferDepthSampleCounts                    VkSampleCountFlags
	FramebufferStencilSampleCounts                  VkSampleCountFlags
	FramebufferNoAttachmentsSampleCounts            VkSampleCountFlags
	MaxColorAttachments                             uint32
	SampledImageColorSampleCounts                   VkSampleCountFlags
	SampledImageIntegerSampleCounts                 VkSampleCountFlags
	SampledImageDepthSampleCounts                   VkSampleCountFlags
	SampledImageStencilSampleCounts                 VkSampleCountFlags
	StorageImageSampleCounts                        VkSampleCountFlags
	MaxSampleMaskWords                              uint32
	TimestampComputeAndGraphics                     VkBool32
	TimestampPeriod                                 float32
	MaxClipDistances                                uint32
	MaxCullDistances                                uint32
	MaxCombinedClipAndCullDistances                 uint32
	DiscreteQueuePriorities                         uint32
	PointSizeRange                                  float32
	LineWidthRange                                  float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     VkBool32
	StandardSampleLocations                         VkBool32
	OptimalBufferCopyOffsetAlignment                VkDeviceSize
	OptimalBufferCopyRowPitchAlignment              VkDeviceSize
	NonCoherentAtomSize                             VkDeviceSize
}

type VkSemaphoreCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkSemaphoreCreateFlags
}

type VkQueryPoolCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkQueryPoolCreateFlags
	QueryType          VkQueryType
	QueryCount         uint32
	PipelineStatistics VkQueryPipelineStatisticFlags
}

type VkFramebufferCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkFramebufferCreateFlags
	RenderPass      VkRenderPass
	AttachmentCount uint32
	PAttachments    *VkImageView
	Width           uint32
	Height          uint32
	Layers          uint32
}

type VkDrawIndirectCommand struct {
	VertexCount   uint32
	InstanceCount uint32
	FirstVertex   uint32
	FirstInstance uint32
}

type VkDrawIndexedIndirectCommand struct {
	IndexCount    uint32
	InstanceCount uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}

type VkDispatchIndirectCommand struct {
	X uint32
	Y uint32
	Z uint32
}

type VkMultiDrawInfoEXT struct {
	FirstVertex uint32
	VertexCount uint32
}

type VkMultiDrawIndexedInfoEXT struct {
	FirstIndex   uint32
	IndexCount   uint32
	VertexOffset int32
}

type VkSubmitInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	WaitSemaphoreCount   uint32
	PWaitSemaphores      *VkSemaphore
	PWaitDstStageMask    *VkPipelineStageFlags
	CommandBufferCount   uint32
	PCommandBuffers      *VkCommandBuffer
	SignalSemaphoreCount uint32
	PSignalSemaphores    *VkSemaphore
}

type VkDisplayPropertiesKHR struct {
	Display              VkDisplayKHR
	DisplayName          string
	PhysicalDimensions   VkExtent2D
	PhysicalResolution   VkExtent2D
	SupportedTransforms  VkSurfaceTransformFlagsKHR
	PlaneReorderPossible VkBool32
	PersistentContent    VkBool32
}

type VkDisplayPlanePropertiesKHR struct {
	CurrentDisplay    VkDisplayKHR
	CurrentStackIndex uint32
}

type VkDisplayModeParametersKHR struct {
	VisibleRegion VkExtent2D
	RefreshRate   uint32
}

type VkDisplayModePropertiesKHR struct {
	DisplayMode VkDisplayModeKHR
	Parameters  VkDisplayModeParametersKHR
}

type VkDisplayModeCreateInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkDisplayModeCreateFlagsKHR
	Parameters VkDisplayModeParametersKHR
}

type VkDisplayPlaneCapabilitiesKHR struct {
	SupportedAlpha VkDisplayPlaneAlphaFlagsKHR
	MinSrcPosition VkOffset2D
	MaxSrcPosition VkOffset2D
	MinSrcExtent   VkExtent2D
	MaxSrcExtent   VkExtent2D
	MinDstPosition VkOffset2D
	MaxDstPosition VkOffset2D
	MinDstExtent   VkExtent2D
	MaxDstExtent   VkExtent2D
}

type VkDisplaySurfaceCreateInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkDisplaySurfaceCreateFlagsKHR
	DisplayMode     VkDisplayModeKHR
	PlaneIndex      uint32
	PlaneStackIndex uint32
	Transform       VkSurfaceTransformFlagBitsKHR
	GlobalAlpha     float32
	AlphaMode       VkDisplayPlaneAlphaFlagBitsKHR
	ImageExtent     VkExtent2D
}

type VkDisplaySurfaceStereoCreateInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	StereoType VkDisplaySurfaceStereoTypeNV
}

type VkDisplayPresentInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	SrcRect    VkRect2D
	DstRect    VkRect2D
	Persistent VkBool32
}

type VkSurfaceCapabilitiesKHR struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           VkExtent2D
	MinImageExtent          VkExtent2D
	MaxImageExtent          VkExtent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     VkSurfaceTransformFlagsKHR
	CurrentTransform        VkSurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha VkCompositeAlphaFlagsKHR
	SupportedUsageFlags     VkImageUsageFlags
}

type VkAndroidSurfaceCreateInfoKHR struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkAndroidSurfaceCreateFlagsKHR
	Window *ANativeWindow
}

type VkViSurfaceCreateInfoNN struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkViSurfaceCreateFlagsNN
	Window unsafe.Pointer
}

type VkWaylandSurfaceCreateInfoKHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Flags   VkWaylandSurfaceCreateFlagsKHR
	Display *wl_display
	Surface *wl_surface
}

type VkUbmSurfaceCreateInfoSEC struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Flags   VkUbmSurfaceCreateFlagsSEC
	Device  *ubm_device
	Surface *ubm_surface
}

type VkWin32SurfaceCreateInfoKHR struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Flags     VkWin32SurfaceCreateFlagsKHR
	Hinstance HINSTANCE
	Hwnd      HWND
}

type VkXlibSurfaceCreateInfoKHR struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkXlibSurfaceCreateFlagsKHR
	Dpy    *Display
	Window Window
}

type VkXcbSurfaceCreateInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkXcbSurfaceCreateFlagsKHR
	Connection *xcb_connection_t
	Window     xcb_window_t
}

type VkDirectFBSurfaceCreateInfoEXT struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Flags   VkDirectFBSurfaceCreateFlagsEXT
	Dfb     *IDirectFB
	Surface *IDirectFBSurface
}

type VkImagePipeSurfaceCreateInfoFUCHSIA struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkImagePipeSurfaceCreateFlagsFUCHSIA
	ImagePipeHandle zx_handle_t
}

type VkStreamDescriptorSurfaceCreateInfoGGP struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkStreamDescriptorSurfaceCreateFlagsGGP
	StreamDescriptor GgpStreamDescriptor
}

type VkScreenSurfaceCreateInfoQNX struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Flags   VkScreenSurfaceCreateFlagsQNX
	Context *_screen_context
	Window  *_screen_window
}

type VkSurfaceFormatKHR struct {
	Format     VkFormat
	ColorSpace VkColorSpaceKHR
}

type VkSwapchainCreateInfoKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkSwapchainCreateFlagsKHR
	Surface               VkSurfaceKHR
	MinImageCount         uint32
	ImageFormat           VkFormat
	ImageColorSpace       VkColorSpaceKHR
	ImageExtent           VkExtent2D
	ImageArrayLayers      uint32
	ImageUsage            VkImageUsageFlags
	ImageSharingMode      VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
	PreTransform          VkSurfaceTransformFlagBitsKHR
	CompositeAlpha        VkCompositeAlphaFlagBitsKHR
	PresentMode           VkPresentModeKHR
	Clipped               VkBool32
	OldSwapchain          VkSwapchainKHR
	OldSwapchain          VkSwapchainKHR
}

type VkPresentInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	WaitSemaphoreCount uint32
	PWaitSemaphores    *VkSemaphore
	SwapchainCount     uint32
	PSwapchains        *VkSwapchainKHR
	PImageIndices      *uint32
	PResults           *VkResult
}

type VkDebugReportCallbackCreateInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkDebugReportFlagsEXT
	PfnCallback PFN_vkDebugReportCallbackEXT
	PUserData   unsafe.Pointer
}

type VkValidationFlagsEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	DisabledValidationCheckCount uint32
	PDisabledValidationChecks    *VkValidationCheckEXT
}

type VkValidationFeaturesEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	EnabledValidationFeatureCount  uint32
	PEnabledValidationFeatures     *VkValidationFeatureEnableEXT
	DisabledValidationFeatureCount uint32
	PDisabledValidationFeatures    *VkValidationFeatureDisableEXT
}

type VkLayerSettingsCreateInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	SettingCount uint32
	PSettings    *VkLayerSettingEXT
}

type VkLayerSettingEXT struct {
	PLayerName   string
	PSettingName string
	Type         VkLayerSettingTypeEXT
	ValueCount   uint32
	PValues      unsafe.Pointer
}

type VkApplicationParametersEXT struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	VendorID uint32
	DeviceID uint32
	Key      uint32
	Value    uint64
}

type VkPipelineRasterizationStateRasterizationOrderAMD struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	RasterizationOrder VkRasterizationOrderAMD
}

type VkDebugMarkerObjectNameInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	ObjectType  VkDebugReportObjectTypeEXT
	Object      uint64
	PObjectName string
}

type VkDebugMarkerObjectTagInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	ObjectType VkDebugReportObjectTypeEXT
	Object     uint64
	TagName    uint64
	TagSize    uint
	PTag       unsafe.Pointer
}

type VkDebugMarkerMarkerInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PMarkerName string
	Color       float32
}

type VkDedicatedAllocationImageCreateInfoNV struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	DedicatedAllocation VkBool32
}

type VkDedicatedAllocationBufferCreateInfoNV struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	DedicatedAllocation VkBool32
}

type VkDedicatedAllocationMemoryAllocateInfoNV struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Image  VkImage
	Buffer VkBuffer
}

type VkExternalImageFormatPropertiesNV struct {
	ImageFormatProperties         VkImageFormatProperties
	ExternalMemoryFeatures        VkExternalMemoryFeatureFlagsNV
	ExportFromImportedHandleTypes VkExternalMemoryHandleTypeFlagsNV
	CompatibleHandleTypes         VkExternalMemoryHandleTypeFlagsNV
}

type VkExternalMemoryImageCreateInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlagsNV
}

type VkExportMemoryAllocateInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlagsNV
}

type VkImportMemoryWin32HandleInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagsNV
	Handle     HANDLE
}

type VkExportMemoryWin32HandleInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
}

type VkExportMemorySciBufInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes NvSciBufAttrList
}

type VkImportMemorySciBufInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
	Handle     NvSciBufObj
}

type VkMemoryGetSciBufInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkMemorySciBufPropertiesNV struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

type VkPhysicalDeviceExternalMemorySciBufFeaturesNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	SciBufImport VkBool32
	SciBufExport VkBool32
}

type VkPhysicalDeviceExternalSciBufFeaturesNV struct {
}

type VkWin32KeyedMutexAcquireReleaseInfoNV struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	AcquireCount                uint32
	PAcquireSyncs               *VkDeviceMemory
	PAcquireKeys                *uint64
	PAcquireTimeoutMilliseconds *uint32
	ReleaseCount                uint32
	PReleaseSyncs               *VkDeviceMemory
	PReleaseKeys                *uint64
}

type VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	DeviceGeneratedCommands VkBool32
}

type VkPushConstantBankInfoNV struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Bank  uint32
}

type VkPhysicalDevicePushConstantBankFeaturesNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	PushConstantBank VkBool32
}

type VkPhysicalDevicePushConstantBankPropertiesNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MaxGraphicsPushConstantBanks uint32
	MaxComputePushConstantBanks  uint32
	MaxGraphicsPushDataBanks     uint32
	MaxComputePushDataBanks      uint32
}

type VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	DeviceGeneratedCompute              VkBool32
	DeviceGeneratedComputePipelines     VkBool32
	DeviceGeneratedComputeCaptureReplay VkBool32
}

type VkDevicePrivateDataCreateInfo struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PrivateDataSlotRequestCount uint32
}

type VkDevicePrivateDataCreateInfoEXT struct {
}

type VkPrivateDataSlotCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkPrivateDataSlotCreateFlags
}

type VkPrivateDataSlotCreateInfoEXT struct {
}

type VkPhysicalDevicePrivateDataFeatures struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PrivateData VkBool32
}

type VkPhysicalDevicePrivateDataFeaturesEXT struct {
}

type VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV struct {
	SType                                    VkStructureType
	PNext                                    unsafe.Pointer
	MaxGraphicsShaderGroupCount              uint32
	MaxIndirectSequenceCount                 uint32
	MaxIndirectCommandsTokenCount            uint32
	MaxIndirectCommandsStreamCount           uint32
	MaxIndirectCommandsTokenOffset           uint32
	MaxIndirectCommandsStreamStride          uint32
	MinSequencesCountBufferOffsetAlignment   uint32
	MinSequencesIndexBufferOffsetAlignment   uint32
	MinIndirectCommandsBufferOffsetAlignment uint32
}

type VkPhysicalDeviceClusterAccelerationStructureFeaturesNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	ClusterAccelerationStructure VkBool32
}

type VkPhysicalDeviceClusterAccelerationStructurePropertiesNV struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	MaxVerticesPerCluster              uint32
	MaxTrianglesPerCluster             uint32
	ClusterScratchByteAlignment        uint32
	ClusterByteAlignment               uint32
	ClusterTemplateByteAlignment       uint32
	ClusterBottomLevelByteAlignment    uint32
	ClusterTemplateBoundsByteAlignment uint32
	MaxClusterGeometryIndex            uint32
}

type VkStridedDeviceAddressNV struct {
	StartAddress  VkDeviceAddress
	StrideInBytes VkDeviceSize
}

type VkRayTracingPipelineClusterAccelerationStructureCreateInfoNV struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	AllowClusterAccelerationStructure VkBool32
}

type VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV struct {
	GeometryIndex uint32
	Reserved      uint32
	GeometryFlags uint32
}

type VkClusterAccelerationStructureMoveObjectsInfoNV struct {
	SrcAccelerationStructure VkDeviceAddress
}

type VkClusterAccelerationStructureBuildClustersBottomLevelInfoNV struct {
	ClusterReferencesCount  uint32
	ClusterReferencesStride uint32
	ClusterReferences       VkDeviceAddress
}

type VkClusterAccelerationStructureGetTemplateIndicesInfoNV struct {
	ClusterTemplateAddress VkDeviceAddress
}

type VkClusterAccelerationStructureBuildTriangleClusterInfoNV struct {
	ClusterID                         uint32
	ClusterFlags                      VkClusterAccelerationStructureClusterFlagsNV
	TriangleCount                     uint32
	VertexCount                       uint32
	PositionTruncateBitCount          uint32
	IndexType                         uint32
	OpacityMicromapIndexType          uint32
	BaseGeometryIndexAndGeometryFlags VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV
	IndexBufferStride                 uint16
	VertexBufferStride                uint16
	GeometryIndexAndFlagsBufferStride uint16
	OpacityMicromapIndexBufferStride  uint16
	IndexBuffer                       VkDeviceAddress
	VertexBuffer                      VkDeviceAddress
	GeometryIndexAndFlagsBuffer       VkDeviceAddress
	OpacityMicromapArray              VkDeviceAddress
	OpacityMicromapIndexBuffer        VkDeviceAddress
}

type VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV struct {
	ClusterID                         uint32
	ClusterFlags                      VkClusterAccelerationStructureClusterFlagsNV
	TriangleCount                     uint32
	VertexCount                       uint32
	PositionTruncateBitCount          uint32
	IndexType                         uint32
	OpacityMicromapIndexType          uint32
	BaseGeometryIndexAndGeometryFlags VkClusterAccelerationStructureGeometryIndexAndGeometryFlagsNV
	IndexBufferStride                 uint16
	VertexBufferStride                uint16
	GeometryIndexAndFlagsBufferStride uint16
	OpacityMicromapIndexBufferStride  uint16
	IndexBuffer                       VkDeviceAddress
	VertexBuffer                      VkDeviceAddress
	GeometryIndexAndFlagsBuffer       VkDeviceAddress
	OpacityMicromapArray              VkDeviceAddress
	OpacityMicromapIndexBuffer        VkDeviceAddress
	InstantiationBoundingBoxLimit     VkDeviceAddress
}

type VkClusterAccelerationStructureInstantiateClusterInfoNV struct {
	ClusterIdOffset        uint32
	GeometryIndexOffset    uint32
	Reserved               uint32
	ClusterTemplateAddress VkDeviceAddress
	VertexBuffer           VkStridedDeviceAddressNV
}

type VkClusterAccelerationStructureClustersBottomLevelInputNV struct {
	SType                                   VkStructureType
	PNext                                   unsafe.Pointer
	MaxTotalClusterCount                    uint32
	MaxClusterCountPerAccelerationStructure uint32
}

type VkClusterAccelerationStructureTriangleClusterInputNV struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	VertexFormat                  VkFormat
	MaxGeometryIndexValue         uint32
	MaxClusterUniqueGeometryCount uint32
	MaxClusterTriangleCount       uint32
	MaxClusterVertexCount         uint32
	MaxTotalTriangleCount         uint32
	MaxTotalVertexCount           uint32
	MinPositionTruncateBitCount   uint32
}

type VkClusterAccelerationStructureMoveObjectsInputNV struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Type          VkClusterAccelerationStructureTypeNV
	NoMoveOverlap VkBool32
	MaxMovedBytes VkDeviceSize
}

type VkClusterAccelerationStructureInputInfoNV struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	MaxAccelerationStructureCount uint32
	Flags                         VkBuildAccelerationStructureFlagsKHR
	OpType                        VkClusterAccelerationStructureOpTypeNV
	OpMode                        VkClusterAccelerationStructureOpModeNV
	OpInput                       VkClusterAccelerationStructureOpInputNV
}

type VkClusterAccelerationStructureCommandsInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Input                  VkClusterAccelerationStructureInputInfoNV
	DstImplicitData        VkDeviceAddress
	ScratchData            VkDeviceAddress
	DstAddressesArray      VkStridedDeviceAddressRegionKHR
	DstSizesArray          VkStridedDeviceAddressRegionKHR
	SrcInfosArray          VkStridedDeviceAddressRegionKHR
	SrcInfosCount          VkDeviceAddress
	AddressResolutionFlags VkClusterAccelerationStructureAddressResolutionFlagsNV
}

type VkPhysicalDeviceMultiDrawPropertiesEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	MaxMultiDrawCount uint32
}

type VkGraphicsShaderGroupCreateInfoNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	StageCount         uint32
	PStages            *VkPipelineShaderStageCreateInfo
	PVertexInputState  *VkPipelineVertexInputStateCreateInfo
	PTessellationState *VkPipelineTessellationStateCreateInfo
}

type VkGraphicsPipelineShaderGroupsCreateInfoNV struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	GroupCount    uint32
	PGroups       *VkGraphicsShaderGroupCreateInfoNV
	PipelineCount uint32
	PPipelines    *VkPipeline
}

type VkBindShaderGroupIndirectCommandNV struct {
	GroupIndex uint32
}

type VkBindIndexBufferIndirectCommandNV struct {
	BufferAddress VkDeviceAddress
	Size          uint32
	IndexType     VkIndexType
}

type VkBindVertexBufferIndirectCommandNV struct {
	BufferAddress VkDeviceAddress
	Size          uint32
	Stride        uint32
}

type VkSetStateFlagsIndirectCommandNV struct {
	Data uint32
}

type VkIndirectCommandsStreamNV struct {
	Buffer VkBuffer
	Offset VkDeviceSize
}

type VkIndirectCommandsLayoutTokenNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	TokenType                    VkIndirectCommandsTokenTypeNV
	Stream                       uint32
	Offset                       uint32
	VertexBindingUnit            uint32
	VertexDynamicStride          VkBool32
	PushconstantPipelineLayout   VkPipelineLayout
	PushconstantShaderStageFlags VkShaderStageFlags
	PushconstantOffset           uint32
	PushconstantSize             uint32
	IndirectStateFlags           VkIndirectStateFlagsNV
	IndexTypeCount               uint32
	PIndexTypes                  *VkIndexType
	PIndexTypeValues             *uint32
}

type VkIndirectCommandsLayoutCreateInfoNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Flags             VkIndirectCommandsLayoutUsageFlagsNV
	PipelineBindPoint VkPipelineBindPoint
	TokenCount        uint32
	PTokens           *VkIndirectCommandsLayoutTokenNV
	StreamCount       uint32
	PStreamStrides    *uint32
}

type VkGeneratedCommandsInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	PipelineBindPoint      VkPipelineBindPoint
	Pipeline               VkPipeline
	IndirectCommandsLayout VkIndirectCommandsLayoutNV
	StreamCount            uint32
	PStreams               *VkIndirectCommandsStreamNV
	SequencesCount         uint32
	PreprocessBuffer       VkBuffer
	PreprocessOffset       VkDeviceSize
	PreprocessSize         VkDeviceSize
	SequencesCountBuffer   VkBuffer
	SequencesCountOffset   VkDeviceSize
	SequencesIndexBuffer   VkBuffer
	SequencesIndexOffset   VkDeviceSize
}

type VkGeneratedCommandsMemoryRequirementsInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	PipelineBindPoint      VkPipelineBindPoint
	Pipeline               VkPipeline
	IndirectCommandsLayout VkIndirectCommandsLayoutNV
	MaxSequencesCount      uint32
}

type VkPipelineIndirectDeviceAddressInfoNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PipelineBindPoint VkPipelineBindPoint
	Pipeline          VkPipeline
}

type VkBindPipelineIndirectCommandNV struct {
	PipelineAddress VkDeviceAddress
}

type VkPhysicalDeviceFeatures2 struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Features VkPhysicalDeviceFeatures
}

type VkPhysicalDeviceFeatures2KHR struct {
}

type VkPhysicalDeviceProperties2 struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Properties VkPhysicalDeviceProperties
}

type VkPhysicalDeviceProperties2KHR struct {
}

type VkFormatProperties2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	FormatProperties VkFormatProperties
}

type VkFormatProperties2KHR struct {
}

type VkImageFormatProperties2 struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ImageFormatProperties VkImageFormatProperties
}

type VkImageFormatProperties2KHR struct {
}

type VkPhysicalDeviceImageFormatInfo2 struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Format VkFormat
	Type   VkImageType
	Tiling VkImageTiling
	Usage  VkImageUsageFlags
	Flags  VkImageCreateFlags
}

type VkPhysicalDeviceImageFormatInfo2KHR struct {
}

type VkQueueFamilyProperties2 struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	QueueFamilyProperties VkQueueFamilyProperties
}

type VkQueueFamilyProperties2KHR struct {
}

type VkPhysicalDeviceMemoryProperties2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	MemoryProperties VkPhysicalDeviceMemoryProperties
}

type VkPhysicalDeviceMemoryProperties2KHR struct {
}

type VkSparseImageFormatProperties2 struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Properties VkSparseImageFormatProperties
}

type VkSparseImageFormatProperties2KHR struct {
}

type VkPhysicalDeviceSparseImageFormatInfo2 struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Format  VkFormat
	Type    VkImageType
	Samples VkSampleCountFlagBits
	Usage   VkImageUsageFlags
	Tiling  VkImageTiling
}

type VkPhysicalDeviceSparseImageFormatInfo2KHR struct {
}

type VkPhysicalDevicePushDescriptorProperties struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxPushDescriptors uint32
}

type VkPhysicalDevicePushDescriptorPropertiesKHR struct {
}

type VkConformanceVersion struct {
	Major    uint8
	Minor    uint8
	Subminor uint8
	Patch    uint8
}

type VkConformanceVersionKHR struct {
}

type VkPhysicalDeviceDriverProperties struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	DriverID           VkDriverId
	DriverName         char
	DriverInfo         char
	ConformanceVersion VkConformanceVersion
}

type VkPhysicalDeviceDriverPropertiesKHR struct {
}

type VkPresentRegionsKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PRegions       *VkPresentRegionKHR
}

type VkPresentRegionKHR struct {
	RectangleCount uint32
	PRectangles    *VkRectLayerKHR
}

type VkRectLayerKHR struct {
	Offset VkOffset2D
	Extent VkExtent2D
	Layer  uint32
}

type VkPhysicalDeviceVariablePointersFeatures struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	VariablePointersStorageBuffer VkBool32
	VariablePointers              VkBool32
}

type VkPhysicalDeviceVariablePointersFeaturesKHR struct {
}

type VkPhysicalDeviceVariablePointerFeaturesKHR struct {
}

type VkPhysicalDeviceVariablePointerFeatures struct {
}

type VkExternalMemoryProperties struct {
	ExternalMemoryFeatures        VkExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes VkExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         VkExternalMemoryHandleTypeFlags
}

type VkExternalMemoryPropertiesKHR struct {
}

type VkPhysicalDeviceExternalImageFormatInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkPhysicalDeviceExternalImageFormatInfoKHR struct {
}

type VkExternalImageFormatProperties struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties VkExternalMemoryProperties
}

type VkExternalImageFormatPropertiesKHR struct {
}

type VkPhysicalDeviceExternalBufferInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkBufferCreateFlags
	Usage      VkBufferUsageFlags
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkPhysicalDeviceExternalBufferInfoKHR struct {
}

type VkExternalBufferProperties struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties VkExternalMemoryProperties
}

type VkExternalBufferPropertiesKHR struct {
}

type VkPhysicalDeviceIDProperties struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	DeviceUUID      uint8
	DriverUUID      uint8
	DeviceLUID      uint8
	DeviceNodeMask  uint32
	DeviceLUIDValid VkBool32
}

type VkPhysicalDeviceIDPropertiesKHR struct {
}

type VkExternalMemoryImageCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

type VkExternalMemoryImageCreateInfoKHR struct {
}

type VkExternalMemoryBufferCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

type VkExternalMemoryBufferCreateInfoKHR struct {
}

type VkExportMemoryAllocateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

type VkExportMemoryAllocateInfoKHR struct {
}

type VkImportMemoryWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
	Handle     HANDLE
	Name       LPCWSTR
}

type VkExportMemoryWin32HandleInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}

type VkImportMemoryZirconHandleInfoFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
	Handle     zx_handle_t
}

type VkMemoryZirconHandlePropertiesFUCHSIA struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

type VkMemoryGetZirconHandleInfoFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkMemoryWin32HandlePropertiesKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

type VkMemoryGetWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkImportMemoryFdInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
	Fd         int
}

type VkMemoryFdPropertiesKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

type VkMemoryGetFdInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkWin32KeyedMutexAcquireReleaseInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	AcquireCount     uint32
	PAcquireSyncs    *VkDeviceMemory
	PAcquireKeys     *uint64
	PAcquireTimeouts *uint32
	ReleaseCount     uint32
	PReleaseSyncs    *VkDeviceMemory
	PReleaseKeys     *uint64
}

type VkImportMemoryMetalHandleInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
	Handle     unsafe.Pointer
}

type VkMemoryMetalHandlePropertiesEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

type VkMemoryGetMetalHandleInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkPhysicalDeviceExternalSemaphoreInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}

type VkPhysicalDeviceExternalSemaphoreInfoKHR struct {
}

type VkExternalSemaphoreProperties struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ExportFromImportedHandleTypes VkExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         VkExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     VkExternalSemaphoreFeatureFlags
}

type VkExternalSemaphorePropertiesKHR struct {
}

type VkExportSemaphoreCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalSemaphoreHandleTypeFlags
}

type VkExportSemaphoreCreateInfoKHR struct {
}

type VkImportSemaphoreWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	Flags      VkSemaphoreImportFlags
	HandleType VkExternalSemaphoreHandleTypeFlagBits
	Handle     HANDLE
	Name       LPCWSTR
}

type VkExportSemaphoreWin32HandleInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}

type VkD3D12FenceSubmitInfoKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	WaitSemaphoreValuesCount   uint32
	PWaitSemaphoreValues       *uint64
	SignalSemaphoreValuesCount uint32
	PSignalSemaphoreValues     *uint64
}

type VkSemaphoreGetWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}

type VkImportSemaphoreFdInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	Flags      VkSemaphoreImportFlags
	HandleType VkExternalSemaphoreHandleTypeFlagBits
	Fd         int
}

type VkSemaphoreGetFdInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}

type VkImportSemaphoreZirconHandleInfoFUCHSIA struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Semaphore    VkSemaphore
	Flags        VkSemaphoreImportFlags
	HandleType   VkExternalSemaphoreHandleTypeFlagBits
	ZirconHandle zx_handle_t
}

type VkSemaphoreGetZirconHandleInfoFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}

type VkPhysicalDeviceExternalFenceInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalFenceHandleTypeFlagBits
}

type VkPhysicalDeviceExternalFenceInfoKHR struct {
}

type VkExternalFenceProperties struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ExportFromImportedHandleTypes VkExternalFenceHandleTypeFlags
	CompatibleHandleTypes         VkExternalFenceHandleTypeFlags
	ExternalFenceFeatures         VkExternalFenceFeatureFlags
}

type VkExternalFencePropertiesKHR struct {
}

type VkExportFenceCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalFenceHandleTypeFlags
}

type VkExportFenceCreateInfoKHR struct {
}

type VkImportFenceWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Fence      VkFence
	Flags      VkFenceImportFlags
	HandleType VkExternalFenceHandleTypeFlagBits
	Handle     HANDLE
	Name       LPCWSTR
}

type VkExportFenceWin32HandleInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes *SECURITY_ATTRIBUTES
	DwAccess    DWORD
	Name        LPCWSTR
}

type VkFenceGetWin32HandleInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Fence      VkFence
	HandleType VkExternalFenceHandleTypeFlagBits
}

type VkImportFenceFdInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Fence      VkFence
	Flags      VkFenceImportFlags
	HandleType VkExternalFenceHandleTypeFlagBits
	Fd         int
}

type VkFenceGetFdInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Fence      VkFence
	HandleType VkExternalFenceHandleTypeFlagBits
}

type VkExportFenceSciSyncInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes NvSciSyncAttrList
}

type VkImportFenceSciSyncInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Fence      VkFence
	HandleType VkExternalFenceHandleTypeFlagBits
	Handle     unsafe.Pointer
}

type VkFenceGetSciSyncInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Fence      VkFence
	HandleType VkExternalFenceHandleTypeFlagBits
}

type VkExportSemaphoreSciSyncInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PAttributes NvSciSyncAttrList
}

type VkImportSemaphoreSciSyncInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
	Handle     unsafe.Pointer
}

type VkSemaphoreGetSciSyncInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Semaphore  VkSemaphore
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}

type VkSciSyncAttributesInfoNV struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	ClientType    VkSciSyncClientTypeNV
	PrimitiveType VkSciSyncPrimitiveTypeNV
}

type VkPhysicalDeviceExternalSciSyncFeaturesNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	SciSyncFence     VkBool32
	SciSyncSemaphore VkBool32
	SciSyncImport    VkBool32
	SciSyncExport    VkBool32
}

type VkPhysicalDeviceExternalSciSync2FeaturesNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	SciSyncFence      VkBool32
	SciSyncSemaphore2 VkBool32
	SciSyncImport     VkBool32
	SciSyncExport     VkBool32
}

type VkSemaphoreSciSyncPoolCreateInfoNV struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Handle NvSciSyncObj
}

type VkSemaphoreSciSyncCreateInfoNV struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SemaphorePool VkSemaphoreSciSyncPoolNV
	PFence        *NvSciSyncFence
}

type VkDeviceSemaphoreSciSyncPoolReservationCreateInfoNV struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	SemaphoreSciSyncPoolRequestCount uint32
}

type VkPhysicalDeviceMultiviewFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	Multiview                   VkBool32
	MultiviewGeometryShader     VkBool32
	MultiviewTessellationShader VkBool32
}

type VkPhysicalDeviceMultiviewFeaturesKHR struct {
}

type VkPhysicalDeviceMultiviewProperties struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	MaxMultiviewViewCount     uint32
	MaxMultiviewInstanceIndex uint32
}

type VkPhysicalDeviceMultiviewPropertiesKHR struct {
}

type VkRenderPassMultiviewCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	SubpassCount         uint32
	PViewMasks           *uint32
	DependencyCount      uint32
	PViewOffsets         *int32
	CorrelationMaskCount uint32
	PCorrelationMasks    *uint32
}

type VkRenderPassMultiviewCreateInfoKHR struct {
}

type VkSurfaceCapabilities2EXT struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	MinImageCount            uint32
	MaxImageCount            uint32
	CurrentExtent            VkExtent2D
	MinImageExtent           VkExtent2D
	MaxImageExtent           VkExtent2D
	MaxImageArrayLayers      uint32
	SupportedTransforms      VkSurfaceTransformFlagsKHR
	CurrentTransform         VkSurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha  VkCompositeAlphaFlagsKHR
	SupportedUsageFlags      VkImageUsageFlags
	SupportedSurfaceCounters VkSurfaceCounterFlagsEXT
}

type VkDisplayPowerInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	PowerState VkDisplayPowerStateEXT
}

type VkDeviceEventInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	DeviceEvent VkDeviceEventTypeEXT
}

type VkDisplayEventInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	DisplayEvent VkDisplayEventTypeEXT
}

type VkSwapchainCounterCreateInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	SurfaceCounters VkSurfaceCounterFlagsEXT
}

type VkPhysicalDeviceGroupProperties struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	PhysicalDeviceCount uint32
	PhysicalDevices     VkPhysicalDevice
	SubsetAllocation    VkBool32
}

type VkPhysicalDeviceGroupPropertiesKHR struct {
}

type VkMemoryAllocateFlagsInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkMemoryAllocateFlags
	DeviceMask uint32
}

type VkMemoryAllocateFlagsInfoKHR struct {
}

type VkBindBufferMemoryInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Buffer       VkBuffer
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}

type VkBindBufferMemoryInfoKHR struct {
}

type VkBindBufferMemoryDeviceGroupInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DeviceIndexCount uint32
	PDeviceIndices   *uint32
}

type VkBindBufferMemoryDeviceGroupInfoKHR struct {
}

type VkBindImageMemoryInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Image        VkImage
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}

type VkBindImageMemoryInfoKHR struct {
}

type VkBindImageMemoryDeviceGroupInfo struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	DeviceIndexCount             uint32
	PDeviceIndices               *uint32
	SplitInstanceBindRegionCount uint32
	PSplitInstanceBindRegions    *VkRect2D
}

type VkBindImageMemoryDeviceGroupInfoKHR struct {
}

type VkDeviceGroupRenderPassBeginInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	DeviceMask            uint32
	DeviceRenderAreaCount uint32
	PDeviceRenderAreas    *VkRect2D
}

type VkDeviceGroupRenderPassBeginInfoKHR struct {
}

type VkDeviceGroupCommandBufferBeginInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	DeviceMask uint32
}

type VkDeviceGroupCommandBufferBeginInfoKHR struct {
}

type VkDeviceGroupSubmitInfo struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	WaitSemaphoreCount            uint32
	PWaitSemaphoreDeviceIndices   *uint32
	CommandBufferCount            uint32
	PCommandBufferDeviceMasks     *uint32
	SignalSemaphoreCount          uint32
	PSignalSemaphoreDeviceIndices *uint32
}

type VkDeviceGroupSubmitInfoKHR struct {
}

type VkDeviceGroupBindSparseInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	ResourceDeviceIndex uint32
	MemoryDeviceIndex   uint32
}

type VkDeviceGroupBindSparseInfoKHR struct {
}

type VkDeviceGroupPresentCapabilitiesKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PresentMask uint32
	Modes       VkDeviceGroupPresentModeFlagsKHR
}

type VkImageSwapchainCreateInfoKHR struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Swapchain VkSwapchainKHR
}

type VkBindImageMemorySwapchainInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Swapchain  VkSwapchainKHR
	ImageIndex uint32
}

type VkAcquireNextImageInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Swapchain  VkSwapchainKHR
	Timeout    uint64
	Semaphore  VkSemaphore
	Fence      VkFence
	DeviceMask uint32
}

type VkDeviceGroupPresentInfoKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PDeviceMasks   *uint32
	Mode           VkDeviceGroupPresentModeFlagBitsKHR
}

type VkDeviceGroupDeviceCreateInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	PhysicalDeviceCount uint32
	PPhysicalDevices    *VkPhysicalDevice
}

type VkDeviceGroupDeviceCreateInfoKHR struct {
}

type VkDeviceGroupSwapchainCreateInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Modes VkDeviceGroupPresentModeFlagsKHR
}

type VkDescriptorUpdateTemplateEntry struct {
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  VkDescriptorType
	Offset          uint
	Stride          uint
}

type VkDescriptorUpdateTemplateEntryKHR struct {
}

type VkDescriptorUpdateTemplateCreateInfo struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	Flags                      VkDescriptorUpdateTemplateCreateFlags
	DescriptorUpdateEntryCount uint32
	PDescriptorUpdateEntries   *VkDescriptorUpdateTemplateEntry
	TemplateType               VkDescriptorUpdateTemplateType
	DescriptorSetLayout        VkDescriptorSetLayout
	PipelineBindPoint          VkPipelineBindPoint
	PipelineLayout             VkPipelineLayout
	Set                        uint32
}

type VkDescriptorUpdateTemplateCreateInfoKHR struct {
}

type VkXYColorEXT struct {
	X float32
	Y float32
}

type VkPhysicalDevicePresentIdFeaturesKHR struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	PresentId VkBool32
}

type VkPresentIdKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PPresentIds    *uint64
}

type VkPhysicalDevicePresentId2FeaturesKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	PresentId2 VkBool32
}

type VkPresentId2KHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PPresentIds    *uint64
}

type VkPresentWait2InfoKHR struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	PresentId uint64
	Timeout   uint64
}

type VkPhysicalDevicePresentWaitFeaturesKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PresentWait VkBool32
}

type VkPhysicalDevicePresentWait2FeaturesKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	PresentWait2 VkBool32
}

type VkPhysicalDevicePresentTimingFeaturesEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	PresentTiming         VkBool32
	PresentAtAbsoluteTime VkBool32
	PresentAtRelativeTime VkBool32
}

type VkPresentTimingSurfaceCapabilitiesEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	PresentTimingSupported         VkBool32
	PresentAtAbsoluteTimeSupported VkBool32
	PresentAtRelativeTimeSupported VkBool32
	PresentStageQueries            VkPresentStageFlagsEXT
}

type VkSwapchainTimingPropertiesEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	RefreshDuration uint64
	RefreshInterval uint64
}

type VkSwapchainTimeDomainPropertiesEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	TimeDomainCount uint32
	PTimeDomains    *VkTimeDomainKHR
	PTimeDomainIds  *uint64
}

type VkPresentStageTimeEXT struct {
	Stage VkPresentStageFlagsEXT
	Time  uint64
}

type VkPastPresentationTimingInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Flags     VkPastPresentationTimingFlagsEXT
	Swapchain VkSwapchainKHR
}

type VkPastPresentationTimingPropertiesEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	TimingPropertiesCounter uint64
	TimeDomainsCounter      uint64
	PresentationTimingCount uint32
	PPresentationTimings    *VkPastPresentationTimingEXT
}

type VkPastPresentationTimingEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PresentId         uint64
	TargetTime        uint64
	PresentStageCount uint32
	PPresentStages    *VkPresentStageTimeEXT
	TimeDomain        VkTimeDomainKHR
	TimeDomainId      uint64
	ReportComplete    VkBool32
}

type VkPresentTimingsInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PTimingInfos   *VkPresentTimingInfoEXT
}

type VkPresentTimingInfoEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	Flags                        VkPresentTimingInfoFlagsEXT
	TargetTime                   uint64
	TimeDomainId                 uint64
	PresentStageQueries          VkPresentStageFlagsEXT
	TargetTimeDomainPresentStage VkPresentStageFlagsEXT
}

type VkSwapchainCalibratedTimestampInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Swapchain    VkSwapchainKHR
	PresentStage VkPresentStageFlagsEXT
	TimeDomainId uint64
}

type VkHdrMetadataEXT struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	DisplayPrimaryRed         VkXYColorEXT
	DisplayPrimaryGreen       VkXYColorEXT
	DisplayPrimaryBlue        VkXYColorEXT
	WhitePoint                VkXYColorEXT
	MaxLuminance              float32
	MinLuminance              float32
	MaxContentLightLevel      float32
	MaxFrameAverageLightLevel float32
}

type VkHdrVividDynamicMetadataHUAWEI struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	DynamicMetadataSize uint
	PDynamicMetadata    unsafe.Pointer
}

type VkDisplayNativeHdrSurfaceCapabilitiesAMD struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	LocalDimmingSupport VkBool32
}

type VkSwapchainDisplayNativeHdrCreateInfoAMD struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	LocalDimmingEnable VkBool32
}

type VkRefreshCycleDurationGOOGLE struct {
	RefreshDuration uint64
}

type VkPastPresentationTimingGOOGLE struct {
	PresentID           uint32
	DesiredPresentTime  uint64
	ActualPresentTime   uint64
	EarliestPresentTime uint64
	PresentMargin       uint64
}

type VkPresentTimesInfoGOOGLE struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PTimes         *VkPresentTimeGOOGLE
}

type VkPresentTimeGOOGLE struct {
	PresentID          uint32
	DesiredPresentTime uint64
}

type VkIOSSurfaceCreateInfoMVK struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkIOSSurfaceCreateFlagsMVK
	PView unsafe.Pointer
}

type VkMacOSSurfaceCreateInfoMVK struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkMacOSSurfaceCreateFlagsMVK
	PView unsafe.Pointer
}

type VkMetalSurfaceCreateInfoEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkMetalSurfaceCreateFlagsEXT
	PLayer *CAMetalLayer
}

type VkViewportWScalingNV struct {
	Xcoeff float32
	Ycoeff float32
}

type VkPipelineViewportWScalingStateCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ViewportWScalingEnable VkBool32
	ViewportCount          uint32
	PViewportWScalings     *VkViewportWScalingNV
}

type VkViewportSwizzleNV struct {
	X VkViewportCoordinateSwizzleNV
	Y VkViewportCoordinateSwizzleNV
	Z VkViewportCoordinateSwizzleNV
	W VkViewportCoordinateSwizzleNV
}

type VkPipelineViewportSwizzleStateCreateInfoNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Flags             VkPipelineViewportSwizzleStateCreateFlagsNV
	ViewportCount     uint32
	PViewportSwizzles *VkViewportSwizzleNV
}

type VkPhysicalDeviceDiscardRectanglePropertiesEXT struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	MaxDiscardRectangles uint32
}

type VkPipelineDiscardRectangleStateCreateInfoEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkPipelineDiscardRectangleStateCreateFlagsEXT
	DiscardRectangleMode  VkDiscardRectangleModeEXT
	DiscardRectangleCount uint32
	PDiscardRectangles    *VkRect2D
}

type VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	PerViewPositionAllComponents VkBool32
}

type VkInputAttachmentAspectReference struct {
	Subpass              uint32
	InputAttachmentIndex uint32
	AspectMask           VkImageAspectFlags
}

type VkInputAttachmentAspectReferenceKHR struct {
}

type VkRenderPassInputAttachmentAspectCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	AspectReferenceCount uint32
	PAspectReferences    *VkInputAttachmentAspectReference
}

type VkRenderPassInputAttachmentAspectCreateInfoKHR struct {
}

type VkPhysicalDeviceSurfaceInfo2KHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Surface VkSurfaceKHR
}

type VkSurfaceCapabilities2KHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	SurfaceCapabilities VkSurfaceCapabilitiesKHR
}

type VkSurfaceFormat2KHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SurfaceFormat VkSurfaceFormatKHR
}

type VkDisplayProperties2KHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DisplayProperties VkDisplayPropertiesKHR
}

type VkDisplayPlaneProperties2KHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	DisplayPlaneProperties VkDisplayPlanePropertiesKHR
}

type VkDisplayModeProperties2KHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	DisplayModeProperties VkDisplayModePropertiesKHR
}

type VkDisplayModeStereoPropertiesNV struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Hdmi3DSupported VkBool32
}

type VkDisplayPlaneInfo2KHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Mode       VkDisplayModeKHR
	PlaneIndex uint32
}

type VkDisplayPlaneCapabilities2KHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Capabilities VkDisplayPlaneCapabilitiesKHR
}

type VkSharedPresentSurfaceCapabilitiesKHR struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	SharedPresentSupportedUsageFlags VkImageUsageFlags
}

type VkPhysicalDevice16BitStorageFeatures struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	StorageBuffer16BitAccess           VkBool32
	UniformAndStorageBuffer16BitAccess VkBool32
	StoragePushConstant16              VkBool32
	StorageInputOutput16               VkBool32
}

type VkPhysicalDevice16BitStorageFeaturesKHR struct {
}

type VkPhysicalDeviceSubgroupProperties struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	SubgroupSize              uint32
	SupportedStages           VkShaderStageFlags
	SupportedOperations       VkSubgroupFeatureFlags
	QuadOperationsInAllStages VkBool32
}

type VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	ShaderSubgroupExtendedTypes VkBool32
}

type VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR struct {
}

type VkBufferMemoryRequirementsInfo2 struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
}

type VkBufferMemoryRequirementsInfo2KHR struct {
}

type VkDeviceBufferMemoryRequirements struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PCreateInfo *VkBufferCreateInfo
}

type VkDeviceBufferMemoryRequirementsKHR struct {
}

type VkImageMemoryRequirementsInfo2 struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Image VkImage
}

type VkImageMemoryRequirementsInfo2KHR struct {
}

type VkImageSparseMemoryRequirementsInfo2 struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Image VkImage
}

type VkImageSparseMemoryRequirementsInfo2KHR struct {
}

type VkDeviceImageMemoryRequirements struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PCreateInfo *VkImageCreateInfo
	PlaneAspect VkImageAspectFlagBits
}

type VkDeviceImageMemoryRequirementsKHR struct {
}

type VkMemoryRequirements2 struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MemoryRequirements VkMemoryRequirements
}

type VkMemoryRequirements2KHR struct {
}

type VkSparseImageMemoryRequirements2 struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MemoryRequirements VkSparseImageMemoryRequirements
}

type VkSparseImageMemoryRequirements2KHR struct {
}

type VkPhysicalDevicePointClippingProperties struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	PointClippingBehavior VkPointClippingBehavior
}

type VkPhysicalDevicePointClippingPropertiesKHR struct {
}

type VkMemoryDedicatedRequirements struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PrefersDedicatedAllocation  VkBool32
	RequiresDedicatedAllocation VkBool32
}

type VkMemoryDedicatedRequirementsKHR struct {
}

type VkMemoryDedicatedAllocateInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Image  VkImage
	Buffer VkBuffer
}

type VkMemoryDedicatedAllocateInfoKHR struct {
}

type VkImageViewUsageCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkImageUsageFlags
}

type VkImageViewSlicedCreateInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	SliceOffset uint32
	SliceCount  uint32
}

type VkImageViewUsageCreateInfoKHR struct {
}

type VkPipelineTessellationDomainOriginStateCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	DomainOrigin VkTessellationDomainOrigin
}

type VkPipelineTessellationDomainOriginStateCreateInfoKHR struct {
}

type VkSamplerYcbcrConversionInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Conversion VkSamplerYcbcrConversion
}

type VkSamplerYcbcrConversionInfoKHR struct {
}

type VkSamplerYcbcrConversionCreateInfo struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	Format                      VkFormat
	YcbcrModel                  VkSamplerYcbcrModelConversion
	YcbcrRange                  VkSamplerYcbcrRange
	Components                  VkComponentMapping
	XChromaOffset               VkChromaLocation
	YChromaOffset               VkChromaLocation
	ChromaFilter                VkFilter
	ForceExplicitReconstruction VkBool32
}

type VkSamplerYcbcrConversionCreateInfoKHR struct {
}

type VkBindImagePlaneMemoryInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PlaneAspect VkImageAspectFlagBits
}

type VkBindImagePlaneMemoryInfoKHR struct {
}

type VkImagePlaneMemoryRequirementsInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PlaneAspect VkImageAspectFlagBits
}

type VkImagePlaneMemoryRequirementsInfoKHR struct {
}

type VkPhysicalDeviceSamplerYcbcrConversionFeatures struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	SamplerYcbcrConversion VkBool32
}

type VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR struct {
}

type VkSamplerYcbcrConversionImageFormatProperties struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	CombinedImageSamplerDescriptorCount uint32
}

type VkSamplerYcbcrConversionImageFormatPropertiesKHR struct {
}

type VkTextureLODGatherFormatPropertiesAMD struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	SupportsTextureGatherLODBiasAMD VkBool32
}

type VkConditionalRenderingBeginInfoEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
	Offset VkDeviceSize
	Flags  VkConditionalRenderingFlagsEXT
}

type VkProtectedSubmitInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ProtectedSubmit VkBool32
}

type VkPhysicalDeviceProtectedMemoryFeatures struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ProtectedMemory VkBool32
}

type VkPhysicalDeviceProtectedMemoryProperties struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ProtectedNoFault VkBool32
}

type VkDeviceQueueInfo2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueIndex       uint32
}

type VkPipelineCoverageToColorStateCreateInfoNV struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkPipelineCoverageToColorStateCreateFlagsNV
	CoverageToColorEnable   VkBool32
	CoverageToColorLocation uint32
}

type VkPhysicalDeviceSamplerFilterMinmaxProperties struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	FilterMinmaxSingleComponentFormats VkBool32
	FilterMinmaxImageComponentMapping  VkBool32
}

type VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT struct {
}

type VkSampleLocationEXT struct {
	X float32
	Y float32
}

type VkSampleLocationsInfoEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	SampleLocationsPerPixel VkSampleCountFlagBits
	SampleLocationGridSize  VkExtent2D
	SampleLocationsCount    uint32
	PSampleLocations        *VkSampleLocationEXT
}

type VkAttachmentSampleLocationsEXT struct {
	AttachmentIndex     uint32
	SampleLocationsInfo VkSampleLocationsInfoEXT
}

type VkSubpassSampleLocationsEXT struct {
	SubpassIndex        uint32
	SampleLocationsInfo VkSampleLocationsInfoEXT
}

type VkRenderPassSampleLocationsBeginInfoEXT struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	AttachmentInitialSampleLocationsCount uint32
	PAttachmentInitialSampleLocations     *VkAttachmentSampleLocationsEXT
	PostSubpassSampleLocationsCount       uint32
	PPostSubpassSampleLocations           *VkSubpassSampleLocationsEXT
}

type VkPipelineSampleLocationsStateCreateInfoEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	SampleLocationsEnable VkBool32
	SampleLocationsInfo   VkSampleLocationsInfoEXT
}

type VkPhysicalDeviceSampleLocationsPropertiesEXT struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	SampleLocationSampleCounts    VkSampleCountFlags
	MaxSampleLocationGridSize     VkExtent2D
	SampleLocationCoordinateRange float32
	SampleLocationSubPixelBits    uint32
	VariableSampleLocations       VkBool32
}

type VkMultisamplePropertiesEXT struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	MaxSampleLocationGridSize VkExtent2D
}

type VkSamplerReductionModeCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	ReductionMode VkSamplerReductionMode
}

type VkSamplerReductionModeCreateInfoEXT struct {
}

type VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	AdvancedBlendCoherentOperations VkBool32
}

type VkPhysicalDeviceMultiDrawFeaturesEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	MultiDraw VkBool32
}

type VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	AdvancedBlendMaxColorAttachments      uint32
	AdvancedBlendIndependentBlend         VkBool32
	AdvancedBlendNonPremultipliedSrcColor VkBool32
	AdvancedBlendNonPremultipliedDstColor VkBool32
	AdvancedBlendCorrelatedOverlap        VkBool32
	AdvancedBlendAllOperations            VkBool32
}

type VkPipelineColorBlendAdvancedStateCreateInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	SrcPremultiplied VkBool32
	DstPremultiplied VkBool32
	BlendOverlap     VkBlendOverlapEXT
}

type VkPhysicalDeviceInlineUniformBlockFeatures struct {
	SType                                              VkStructureType
	PNext                                              unsafe.Pointer
	InlineUniformBlock                                 VkBool32
	DescriptorBindingInlineUniformBlockUpdateAfterBind VkBool32
}

type VkPhysicalDeviceInlineUniformBlockFeaturesEXT struct {
}

type VkPhysicalDeviceInlineUniformBlockProperties struct {
	SType                                                   VkStructureType
	PNext                                                   unsafe.Pointer
	MaxInlineUniformBlockSize                               uint32
	MaxPerStageDescriptorInlineUniformBlocks                uint32
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks uint32
	MaxDescriptorSetInlineUniformBlocks                     uint32
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks      uint32
}

type VkPhysicalDeviceInlineUniformBlockPropertiesEXT struct {
}

type VkWriteDescriptorSetInlineUniformBlock struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	DataSize uint32
	PData    unsafe.Pointer
}

type VkWriteDescriptorSetInlineUniformBlockEXT struct {
}

type VkDescriptorPoolInlineUniformBlockCreateInfo struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	MaxInlineUniformBlockBindings uint32
}

type VkDescriptorPoolInlineUniformBlockCreateInfoEXT struct {
}

type VkPipelineCoverageModulationStateCreateInfoNV struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	Flags                         VkPipelineCoverageModulationStateCreateFlagsNV
	CoverageModulationMode        VkCoverageModulationModeNV
	CoverageModulationTableEnable VkBool32
	CoverageModulationTableCount  uint32
	PCoverageModulationTable      *float32
}

type VkImageFormatListCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ViewFormatCount uint32
	PViewFormats    *VkFormat
}

type VkImageFormatListCreateInfoKHR struct {
}

type VkValidationCacheCreateInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkValidationCacheCreateFlagsEXT
	InitialDataSize uint
	PInitialData    unsafe.Pointer
}

type VkShaderModuleValidationCacheCreateInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ValidationCache VkValidationCacheEXT
}

type VkPhysicalDeviceMaintenance3Properties struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	MaxPerSetDescriptors    uint32
	MaxMemoryAllocationSize VkDeviceSize
}

type VkPhysicalDeviceMaintenance3PropertiesKHR struct {
}

type VkPhysicalDeviceMaintenance4Features struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance4 VkBool32
}

type VkPhysicalDeviceMaintenance4FeaturesKHR struct {
}

type VkPhysicalDeviceMaintenance4Properties struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	MaxBufferSize VkDeviceSize
}

type VkPhysicalDeviceMaintenance4PropertiesKHR struct {
}

type VkPhysicalDeviceMaintenance5Features struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance5 VkBool32
}

type VkPhysicalDeviceMaintenance5FeaturesKHR struct {
}

type VkPhysicalDeviceMaintenance5Properties struct {
	SType                                               VkStructureType
	PNext                                               unsafe.Pointer
	EarlyFragmentMultisampleCoverageAfterSampleCounting VkBool32
	EarlyFragmentSampleMaskTestBeforeSampleCounting     VkBool32
	DepthStencilSwizzleOneSupport                       VkBool32
	PolygonModePointSize                                VkBool32
	NonStrictSinglePixelWideLinesUseParallelogram       VkBool32
	NonStrictWideLinesUseParallelogram                  VkBool32
}

type VkPhysicalDeviceMaintenance5PropertiesKHR struct {
}

type VkPhysicalDeviceMaintenance6Features struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance6 VkBool32
}

type VkPhysicalDeviceMaintenance6FeaturesKHR struct {
}

type VkPhysicalDeviceMaintenance6Properties struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	BlockTexelViewCompatibleMultipleLayers VkBool32
	MaxCombinedImageSamplerDescriptorCount uint32
	FragmentShadingRateClampCombinerInputs VkBool32
}

type VkPhysicalDeviceMaintenance6PropertiesKHR struct {
}

type VkPhysicalDeviceMaintenance7FeaturesKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance7 VkBool32
}

type VkPhysicalDeviceMaintenance7PropertiesKHR struct {
	SType                                                     VkStructureType
	PNext                                                     unsafe.Pointer
	RobustFragmentShadingRateAttachmentAccess                 VkBool32
	SeparateDepthStencilAttachmentAccess                      VkBool32
	MaxDescriptorSetTotalUniformBuffersDynamic                uint32
	MaxDescriptorSetTotalStorageBuffersDynamic                uint32
	MaxDescriptorSetTotalBuffersDynamic                       uint32
	MaxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindTotalBuffersDynamic        uint32
}

type VkPhysicalDeviceLayeredApiPropertiesListKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	LayeredApiCount uint32
	PLayeredApis    *VkPhysicalDeviceLayeredApiPropertiesKHR
}

type VkPhysicalDeviceLayeredApiPropertiesKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	VendorID   uint32
	DeviceID   uint32
	LayeredAPI VkPhysicalDeviceLayeredApiKHR
	DeviceName char
}

type VkPhysicalDeviceLayeredApiVulkanPropertiesKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Properties VkPhysicalDeviceProperties2
}

type VkPhysicalDeviceMaintenance8FeaturesKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance8 VkBool32
}

type VkPhysicalDeviceMaintenance9FeaturesKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance9 VkBool32
}

type VkPhysicalDeviceMaintenance9PropertiesKHR struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	Image2DViewOf3DSparse       VkBool32
	DefaultVertexAttributeValue VkDefaultVertexAttributeValueKHR
}

type VkPhysicalDeviceMaintenance10PropertiesKHR struct {
	SType                                            VkStructureType
	PNext                                            unsafe.Pointer
	Rgba4OpaqueBlackSwizzled                         VkBool32
	ResolveSrgbFormatAppliesTransferFunction         VkBool32
	ResolveSrgbFormatSupportsTransferFunctionControl VkBool32
}

type VkPhysicalDeviceMaintenance10FeaturesKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Maintenance10 VkBool32
}

type VkQueueFamilyOwnershipTransferPropertiesKHR struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	OptimalImageTransferToQueueFamilies uint32
}

type VkRenderingAreaInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ViewMask                uint32
	ColorAttachmentCount    uint32
	PColorAttachmentFormats *VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
}

type VkRenderingAreaInfoKHR struct {
}

type VkDescriptorSetLayoutSupport struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Supported VkBool32
}

type VkDescriptorSetLayoutSupportKHR struct {
}

type VkPhysicalDeviceShaderDrawParametersFeatures struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ShaderDrawParameters VkBool32
}

type VkPhysicalDeviceShaderDrawParameterFeatures struct {
}

type VkPhysicalDeviceShaderFloat16Int8Features struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	ShaderFloat16 VkBool32
	ShaderInt8    VkBool32
}

type VkPhysicalDeviceShaderFloat16Int8FeaturesKHR struct {
}

type VkPhysicalDeviceFloat16Int8FeaturesKHR struct {
}

type VkPhysicalDeviceFloatControlsProperties struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	DenormBehaviorIndependence            VkShaderFloatControlsIndependence
	RoundingModeIndependence              VkShaderFloatControlsIndependence
	ShaderSignedZeroInfNanPreserveFloat16 VkBool32
	ShaderSignedZeroInfNanPreserveFloat32 VkBool32
	ShaderSignedZeroInfNanPreserveFloat64 VkBool32
	ShaderDenormPreserveFloat16           VkBool32
	ShaderDenormPreserveFloat32           VkBool32
	ShaderDenormPreserveFloat64           VkBool32
	ShaderDenormFlushToZeroFloat16        VkBool32
	ShaderDenormFlushToZeroFloat32        VkBool32
	ShaderDenormFlushToZeroFloat64        VkBool32
	ShaderRoundingModeRTEFloat16          VkBool32
	ShaderRoundingModeRTEFloat32          VkBool32
	ShaderRoundingModeRTEFloat64          VkBool32
	ShaderRoundingModeRTZFloat16          VkBool32
	ShaderRoundingModeRTZFloat32          VkBool32
	ShaderRoundingModeRTZFloat64          VkBool32
}

type VkPhysicalDeviceFloatControlsPropertiesKHR struct {
}

type VkPhysicalDeviceHostQueryResetFeatures struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	HostQueryReset VkBool32
}

type VkPhysicalDeviceHostQueryResetFeaturesEXT struct {
}

type VkNativeBufferUsage2ANDROID struct {
	Consumer uint64
	Producer uint64
}

type VkNativeBufferANDROID struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Handle unsafe.Pointer
	Stride int
	Format int
	Usage  int
	Usage2 VkNativeBufferUsage2ANDROID
}

type VkSwapchainImageCreateInfoANDROID struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkSwapchainImageUsageFlagsANDROID
}

type VkPhysicalDevicePresentationPropertiesANDROID struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	SharedImage VkBool32
}

type VkShaderResourceUsageAMD struct {
	NumUsedVgprs             uint32
	NumUsedSgprs             uint32
	LdsSizePerLocalWorkGroup uint32
	LdsUsageSizeInBytes      uint
	ScratchMemUsageInBytes   uint
}

type VkShaderStatisticsInfoAMD struct {
	ShaderStageMask      VkShaderStageFlags
	ResourceUsage        VkShaderResourceUsageAMD
	NumPhysicalVgprs     uint32
	NumPhysicalSgprs     uint32
	NumAvailableVgprs    uint32
	NumAvailableSgprs    uint32
	ComputeWorkGroupSize uint32
}

type VkDeviceQueueGlobalPriorityCreateInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	GlobalPriority VkQueueGlobalPriority
}

type VkDeviceQueueGlobalPriorityCreateInfoKHR struct {
}

type VkDeviceQueueGlobalPriorityCreateInfoEXT struct {
}

type VkPhysicalDeviceGlobalPriorityQueryFeatures struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	GlobalPriorityQuery VkBool32
}

type VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR struct {
}

type VkPhysicalDeviceGlobalPriorityQueryFeaturesEXT struct {
}

type VkQueueFamilyGlobalPriorityProperties struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	PriorityCount uint32
	Priorities    VkQueueGlobalPriority
}

type VkQueueFamilyGlobalPriorityPropertiesKHR struct {
}

type VkQueueFamilyGlobalPriorityPropertiesEXT struct {
}

type VkDebugUtilsObjectNameInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	ObjectType   VkObjectType
	ObjectHandle uint64
	PObjectName  string
}

type VkDebugUtilsObjectTagInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	ObjectType   VkObjectType
	ObjectHandle uint64
	TagName      uint64
	TagSize      uint
	PTag         unsafe.Pointer
}

type VkDebugUtilsLabelEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	PLabelName string
	Color      float32
}

type VkDebugUtilsMessengerCreateInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkDebugUtilsMessengerCreateFlagsEXT
	MessageSeverity VkDebugUtilsMessageSeverityFlagsEXT
	MessageType     VkDebugUtilsMessageTypeFlagsEXT
	PfnUserCallback PFN_vkDebugUtilsMessengerCallbackEXT
	PUserData       unsafe.Pointer
}

type VkDebugUtilsMessengerCallbackDataEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkDebugUtilsMessengerCallbackDataFlagsEXT
	PMessageIdName   string
	MessageIdNumber  int32
	PMessage         string
	QueueLabelCount  uint32
	PQueueLabels     *VkDebugUtilsLabelEXT
	CmdBufLabelCount uint32
	PCmdBufLabels    *VkDebugUtilsLabelEXT
	ObjectCount      uint32
	PObjects         *VkDebugUtilsObjectNameInfoEXT
}

type VkPhysicalDeviceDeviceMemoryReportFeaturesEXT struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	DeviceMemoryReport VkBool32
}

type VkDeviceDeviceMemoryReportCreateInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkDeviceMemoryReportFlagsEXT
	PfnUserCallback PFN_vkDeviceMemoryReportCallbackEXT
	PUserData       unsafe.Pointer
}

type VkDeviceMemoryReportCallbackDataEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkDeviceMemoryReportFlagsEXT
	Type           VkDeviceMemoryReportEventTypeEXT
	MemoryObjectId uint64
	Size           VkDeviceSize
	ObjectType     VkObjectType
	ObjectHandle   uint64
	HeapIndex      uint32
}

type VkImportMemoryHostPointerInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	HandleType   VkExternalMemoryHandleTypeFlagBits
	PHostPointer unsafe.Pointer
}

type VkMemoryHostPointerPropertiesEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryTypeBits uint32
}

type VkPhysicalDeviceExternalMemoryHostPropertiesEXT struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	MinImportedHostPointerAlignment VkDeviceSize
}

type VkPhysicalDeviceConservativeRasterizationPropertiesEXT struct {
	SType                                       VkStructureType
	PNext                                       unsafe.Pointer
	PrimitiveOverestimationSize                 float32
	MaxExtraPrimitiveOverestimationSize         float32
	ExtraPrimitiveOverestimationSizeGranularity float32
	PrimitiveUnderestimation                    VkBool32
	ConservativePointAndLineRasterization       VkBool32
	DegenerateTrianglesRasterized               VkBool32
	DegenerateLinesRasterized                   VkBool32
	FullyCoveredFragmentShaderInputVariable     VkBool32
	ConservativeRasterizationPostDepthCoverage  VkBool32
}

type VkCalibratedTimestampInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	TimeDomain VkTimeDomainKHR
}

type VkCalibratedTimestampInfoEXT struct {
}

type VkPhysicalDeviceShaderCorePropertiesAMD struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ShaderEngineCount          uint32
	ShaderArraysPerEngineCount uint32
	ComputeUnitsPerShaderArray uint32
	SimdPerComputeUnit         uint32
	WavefrontsPerSimd          uint32
	WavefrontSize              uint32
	SgprsPerSimd               uint32
	MinSgprAllocation          uint32
	MaxSgprAllocation          uint32
	SgprAllocationGranularity  uint32
	VgprsPerSimd               uint32
	MinVgprAllocation          uint32
	MaxVgprAllocation          uint32
	VgprAllocationGranularity  uint32
}

type VkPhysicalDeviceShaderCoreProperties2AMD struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ShaderCoreFeatures     VkShaderCorePropertiesFlagsAMD
	ActiveComputeUnitCount uint32
}

type VkPipelineRasterizationConservativeStateCreateInfoEXT struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	Flags                            VkPipelineRasterizationConservativeStateCreateFlagsEXT
	ConservativeRasterizationMode    VkConservativeRasterizationModeEXT
	ExtraPrimitiveOverestimationSize float32
}

type VkPhysicalDeviceDescriptorIndexingFeatures struct {
	SType                                              VkStructureType
	PNext                                              unsafe.Pointer
	ShaderInputAttachmentArrayDynamicIndexing          VkBool32
	ShaderUniformTexelBufferArrayDynamicIndexing       VkBool32
	ShaderStorageTexelBufferArrayDynamicIndexing       VkBool32
	ShaderUniformBufferArrayNonUniformIndexing         VkBool32
	ShaderSampledImageArrayNonUniformIndexing          VkBool32
	ShaderStorageBufferArrayNonUniformIndexing         VkBool32
	ShaderStorageImageArrayNonUniformIndexing          VkBool32
	ShaderInputAttachmentArrayNonUniformIndexing       VkBool32
	ShaderUniformTexelBufferArrayNonUniformIndexing    VkBool32
	ShaderStorageTexelBufferArrayNonUniformIndexing    VkBool32
	DescriptorBindingUniformBufferUpdateAfterBind      VkBool32
	DescriptorBindingSampledImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageBufferUpdateAfterBind      VkBool32
	DescriptorBindingUniformTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingStorageTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingUpdateUnusedWhilePending          VkBool32
	DescriptorBindingPartiallyBound                    VkBool32
	DescriptorBindingVariableDescriptorCount           VkBool32
	RuntimeDescriptorArray                             VkBool32
}

type VkPhysicalDeviceDescriptorIndexingFeaturesEXT struct {
}

type VkPhysicalDeviceDescriptorIndexingProperties struct {
	SType                                                VkStructureType
	PNext                                                unsafe.Pointer
	MaxUpdateAfterBindDescriptorsInAllPools              uint32
	ShaderUniformBufferArrayNonUniformIndexingNative     VkBool32
	ShaderSampledImageArrayNonUniformIndexingNative      VkBool32
	ShaderStorageBufferArrayNonUniformIndexingNative     VkBool32
	ShaderStorageImageArrayNonUniformIndexingNative      VkBool32
	ShaderInputAttachmentArrayNonUniformIndexingNative   VkBool32
	RobustBufferAccessUpdateAfterBind                    VkBool32
	QuadDivergentImplicitLod                             VkBool32
	MaxPerStageDescriptorUpdateAfterBindSamplers         uint32
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindSampledImages    uint32
	MaxPerStageDescriptorUpdateAfterBindStorageImages    uint32
	MaxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	MaxPerStageUpdateAfterBindResources                  uint32
	MaxDescriptorSetUpdateAfterBindSamplers              uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffers        uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffers        uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindSampledImages         uint32
	MaxDescriptorSetUpdateAfterBindStorageImages         uint32
	MaxDescriptorSetUpdateAfterBindInputAttachments      uint32
}

type VkPhysicalDeviceDescriptorIndexingPropertiesEXT struct {
}

type VkDescriptorSetLayoutBindingFlagsCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	BindingCount  uint32
	PBindingFlags *VkDescriptorBindingFlags
}

type VkDescriptorSetLayoutBindingFlagsCreateInfoEXT struct {
}

type VkDescriptorSetVariableDescriptorCountAllocateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	DescriptorSetCount uint32
	PDescriptorCounts  *uint32
}

type VkDescriptorSetVariableDescriptorCountAllocateInfoEXT struct {
}

type VkDescriptorSetVariableDescriptorCountLayoutSupport struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	MaxVariableDescriptorCount uint32
}

type VkDescriptorSetVariableDescriptorCountLayoutSupportEXT struct {
}

type VkAttachmentDescription2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkAttachmentDescriptionFlags
	Format         VkFormat
	Samples        VkSampleCountFlagBits
	LoadOp         VkAttachmentLoadOp
	StoreOp        VkAttachmentStoreOp
	StencilLoadOp  VkAttachmentLoadOp
	StencilStoreOp VkAttachmentStoreOp
	InitialLayout  VkImageLayout
	FinalLayout    VkImageLayout
}

type VkAttachmentDescription2KHR struct {
}

type VkAttachmentReference2 struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Attachment uint32
	Layout     VkImageLayout
	AspectMask VkImageAspectFlags
}

type VkAttachmentReference2KHR struct {
}

type VkSubpassDescription2 struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkSubpassDescriptionFlags
	PipelineBindPoint       VkPipelineBindPoint
	ViewMask                uint32
	InputAttachmentCount    uint32
	PInputAttachments       *VkAttachmentReference2
	ColorAttachmentCount    uint32
	PColorAttachments       *VkAttachmentReference2
	PResolveAttachments     *VkAttachmentReference2
	PDepthStencilAttachment *VkAttachmentReference2
	PreserveAttachmentCount uint32
	PPreserveAttachments    *uint32
}

type VkSubpassDescription2KHR struct {
}

type VkSubpassDependency2 struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    VkPipelineStageFlags
	DstStageMask    VkPipelineStageFlags
	SrcAccessMask   VkAccessFlags
	DstAccessMask   VkAccessFlags
	DependencyFlags VkDependencyFlags
	ViewOffset      int32
}

type VkSubpassDependency2KHR struct {
}

type VkRenderPassCreateInfo2 struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkRenderPassCreateFlags
	AttachmentCount         uint32
	PAttachments            *VkAttachmentDescription2
	SubpassCount            uint32
	PSubpasses              *VkSubpassDescription2
	DependencyCount         uint32
	PDependencies           *VkSubpassDependency2
	CorrelatedViewMaskCount uint32
	PCorrelatedViewMasks    *uint32
}

type VkRenderPassCreateInfo2KHR struct {
}

type VkSubpassBeginInfo struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Contents VkSubpassContents
}

type VkSubpassBeginInfoKHR struct {
}

type VkSubpassEndInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkSubpassEndInfoKHR struct {
}

type VkPhysicalDeviceTimelineSemaphoreFeatures struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	TimelineSemaphore VkBool32
}

type VkPhysicalDeviceTimelineSemaphoreFeaturesKHR struct {
}

type VkPhysicalDeviceTimelineSemaphoreProperties struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	MaxTimelineSemaphoreValueDifference uint64
}

type VkPhysicalDeviceTimelineSemaphorePropertiesKHR struct {
}

type VkSemaphoreTypeCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SemaphoreType VkSemaphoreType
	InitialValue  uint64
}

type VkSemaphoreTypeCreateInfoKHR struct {
}

type VkTimelineSemaphoreSubmitInfo struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	WaitSemaphoreValueCount   uint32
	PWaitSemaphoreValues      *uint64
	SignalSemaphoreValueCount uint32
	PSignalSemaphoreValues    *uint64
}

type VkTimelineSemaphoreSubmitInfoKHR struct {
}

type VkSemaphoreWaitInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkSemaphoreWaitFlags
	SemaphoreCount uint32
	PSemaphores    *VkSemaphore
	PValues        *uint64
}

type VkSemaphoreWaitInfoKHR struct {
}

type VkSemaphoreSignalInfo struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Semaphore VkSemaphore
	Value     uint64
}

type VkSemaphoreSignalInfoKHR struct {
}

type VkVertexInputBindingDivisorDescription struct {
	Binding uint32
	Divisor uint32
}

type VkVertexInputBindingDivisorDescriptionKHR struct {
}

type VkVertexInputBindingDivisorDescriptionEXT struct {
}

type VkPipelineVertexInputDivisorStateCreateInfo struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	VertexBindingDivisorCount uint32
	PVertexBindingDivisors    *VkVertexInputBindingDivisorDescription
}

type VkPipelineVertexInputDivisorStateCreateInfoKHR struct {
}

type VkPipelineVertexInputDivisorStateCreateInfoEXT struct {
}

type VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	MaxVertexAttribDivisor uint32
}

type VkPhysicalDeviceVertexAttributeDivisorProperties struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MaxVertexAttribDivisor       uint32
	SupportsNonZeroFirstInstance VkBool32
}

type VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR struct {
}

type VkPhysicalDevicePCIBusInfoPropertiesEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PciDomain   uint32
	PciBus      uint32
	PciDevice   uint32
	PciFunction uint32
}

type VkImportAndroidHardwareBufferInfoANDROID struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer *AHardwareBuffer
}

type VkAndroidHardwareBufferUsageANDROID struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	AndroidHardwareBufferUsage uint64
}

type VkAndroidHardwareBufferPropertiesANDROID struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	AllocationSize VkDeviceSize
	MemoryTypeBits uint32
}

type VkMemoryGetAndroidHardwareBufferInfoANDROID struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
}

type VkAndroidHardwareBufferFormatPropertiesANDROID struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	Format                           VkFormat
	ExternalFormat                   uint64
	FormatFeatures                   VkFormatFeatureFlags
	SamplerYcbcrConversionComponents VkComponentMapping
	SuggestedYcbcrModel              VkSamplerYcbcrModelConversion
	SuggestedYcbcrRange              VkSamplerYcbcrRange
	SuggestedXChromaOffset           VkChromaLocation
	SuggestedYChromaOffset           VkChromaLocation
}

type VkCommandBufferInheritanceConditionalRenderingInfoEXT struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ConditionalRenderingEnable VkBool32
}

type VkExternalFormatANDROID struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	ExternalFormat uint64
}

type VkPhysicalDevice8BitStorageFeatures struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	StorageBuffer8BitAccess           VkBool32
	UniformAndStorageBuffer8BitAccess VkBool32
	StoragePushConstant8              VkBool32
}

type VkPhysicalDevice8BitStorageFeaturesKHR struct {
}

type VkPhysicalDeviceConditionalRenderingFeaturesEXT struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ConditionalRendering          VkBool32
	InheritedConditionalRendering VkBool32
}

type VkPhysicalDeviceVulkanMemoryModelFeatures struct {
	SType                                         VkStructureType
	PNext                                         unsafe.Pointer
	VulkanMemoryModel                             VkBool32
	VulkanMemoryModelDeviceScope                  VkBool32
	VulkanMemoryModelAvailabilityVisibilityChains VkBool32
}

type VkPhysicalDeviceVulkanMemoryModelFeaturesKHR struct {
}

type VkPhysicalDeviceShaderAtomicInt64Features struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ShaderBufferInt64Atomics VkBool32
	ShaderSharedInt64Atomics VkBool32
}

type VkPhysicalDeviceShaderAtomicInt64FeaturesKHR struct {
}

type VkPhysicalDeviceShaderAtomicFloatFeaturesEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	ShaderBufferFloat32Atomics   VkBool32
	ShaderBufferFloat32AtomicAdd VkBool32
	ShaderBufferFloat64Atomics   VkBool32
	ShaderBufferFloat64AtomicAdd VkBool32
	ShaderSharedFloat32Atomics   VkBool32
	ShaderSharedFloat32AtomicAdd VkBool32
	ShaderSharedFloat64Atomics   VkBool32
	ShaderSharedFloat64AtomicAdd VkBool32
	ShaderImageFloat32Atomics    VkBool32
	ShaderImageFloat32AtomicAdd  VkBool32
	SparseImageFloat32Atomics    VkBool32
	SparseImageFloat32AtomicAdd  VkBool32
}

type VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	ShaderBufferFloat16Atomics      VkBool32
	ShaderBufferFloat16AtomicAdd    VkBool32
	ShaderBufferFloat16AtomicMinMax VkBool32
	ShaderBufferFloat32AtomicMinMax VkBool32
	ShaderBufferFloat64AtomicMinMax VkBool32
	ShaderSharedFloat16Atomics      VkBool32
	ShaderSharedFloat16AtomicAdd    VkBool32
	ShaderSharedFloat16AtomicMinMax VkBool32
	ShaderSharedFloat32AtomicMinMax VkBool32
	ShaderSharedFloat64AtomicMinMax VkBool32
	ShaderImageFloat32AtomicMinMax  VkBool32
	SparseImageFloat32AtomicMinMax  VkBool32
}

type VkPhysicalDeviceVertexAttributeDivisorFeatures struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	VertexAttributeInstanceRateDivisor     VkBool32
	VertexAttributeInstanceRateZeroDivisor VkBool32
}

type VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR struct {
}

type VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT struct {
}

type VkQueueFamilyCheckpointPropertiesNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	CheckpointExecutionStageMask VkPipelineStageFlags
}

type VkCheckpointDataNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Stage             VkPipelineStageFlagBits
	PCheckpointMarker unsafe.Pointer
}

type VkPhysicalDeviceDepthStencilResolveProperties struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	SupportedDepthResolveModes   VkResolveModeFlags
	SupportedStencilResolveModes VkResolveModeFlags
	IndependentResolveNone       VkBool32
	IndependentResolve           VkBool32
}

type VkPhysicalDeviceDepthStencilResolvePropertiesKHR struct {
}

type VkSubpassDescriptionDepthStencilResolve struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	DepthResolveMode               VkResolveModeFlagBits
	StencilResolveMode             VkResolveModeFlagBits
	PDepthStencilResolveAttachment *VkAttachmentReference2
}

type VkSubpassDescriptionDepthStencilResolveKHR struct {
}

type VkImageViewASTCDecodeModeEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	DecodeMode VkFormat
}

type VkPhysicalDeviceASTCDecodeFeaturesEXT struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DecodeModeSharedExponent VkBool32
}

type VkPhysicalDeviceTransformFeedbackFeaturesEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	TransformFeedback VkBool32
	GeometryStreams   VkBool32
}

type VkPhysicalDeviceTransformFeedbackPropertiesEXT struct {
	SType                                      VkStructureType
	PNext                                      unsafe.Pointer
	MaxTransformFeedbackStreams                uint32
	MaxTransformFeedbackBuffers                uint32
	MaxTransformFeedbackBufferSize             VkDeviceSize
	MaxTransformFeedbackStreamDataSize         uint32
	MaxTransformFeedbackBufferDataSize         uint32
	MaxTransformFeedbackBufferDataStride       uint32
	TransformFeedbackQueries                   VkBool32
	TransformFeedbackStreamsLinesTriangles     VkBool32
	TransformFeedbackRasterizationStreamSelect VkBool32
	TransformFeedbackDraw                      VkBool32
}

type VkPipelineRasterizationStateStreamCreateInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkPipelineRasterizationStateStreamCreateFlagsEXT
	RasterizationStream uint32
}

type VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	RepresentativeFragmentTest VkBool32
}

type VkPipelineRepresentativeFragmentTestStateCreateInfoNV struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	RepresentativeFragmentTestEnable VkBool32
}

type VkPhysicalDeviceExclusiveScissorFeaturesNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ExclusiveScissor VkBool32
}

type VkPipelineViewportExclusiveScissorStateCreateInfoNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ExclusiveScissorCount uint32
	PExclusiveScissors    *VkRect2D
}

type VkPhysicalDeviceCornerSampledImageFeaturesNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	CornerSampledImage VkBool32
}

type VkPhysicalDeviceComputeShaderDerivativesFeaturesKHR struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	ComputeDerivativeGroupQuads  VkBool32
	ComputeDerivativeGroupLinear VkBool32
}

type VkPhysicalDeviceComputeShaderDerivativesFeaturesNV struct {
}

type VkPhysicalDeviceComputeShaderDerivativesPropertiesKHR struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MeshAndTaskShaderDerivatives VkBool32
}

type VkPhysicalDeviceFragmentShaderBarycentricFeaturesNV struct {
}

type VkPhysicalDeviceShaderImageFootprintFeaturesNV struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	ImageFootprint VkBool32
}

type VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	DedicatedAllocationImageAliasing VkBool32
}

type VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	IndirectMemoryCopy        VkBool32
	IndirectMemoryToImageCopy VkBool32
}

type VkPhysicalDeviceCopyMemoryIndirectFeaturesNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	IndirectCopy VkBool32
}

type VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	SupportedQueues VkQueueFlags
}

type VkPhysicalDeviceCopyMemoryIndirectPropertiesNV struct {
}

type VkPhysicalDeviceMemoryDecompressionFeaturesEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	MemoryDecompression VkBool32
}

type VkPhysicalDeviceMemoryDecompressionFeaturesNV struct {
}

type VkPhysicalDeviceMemoryDecompressionPropertiesEXT struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	DecompressionMethods          VkMemoryDecompressionMethodFlagsEXT
	MaxDecompressionIndirectCount uint64
}

type VkPhysicalDeviceMemoryDecompressionPropertiesNV struct {
}

type VkShadingRatePaletteNV struct {
	ShadingRatePaletteEntryCount uint32
	PShadingRatePaletteEntries   *VkShadingRatePaletteEntryNV
}

type VkPipelineViewportShadingRateImageStateCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ShadingRateImageEnable VkBool32
	ViewportCount          uint32
	PShadingRatePalettes   *VkShadingRatePaletteNV
}

type VkPhysicalDeviceShadingRateImageFeaturesNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	ShadingRateImage             VkBool32
	ShadingRateCoarseSampleOrder VkBool32
}

type VkPhysicalDeviceShadingRateImagePropertiesNV struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	ShadingRateTexelSize        VkExtent2D
	ShadingRatePaletteSize      uint32
	ShadingRateMaxCoarseSamples uint32
}

type VkPhysicalDeviceInvocationMaskFeaturesHUAWEI struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	InvocationMask VkBool32
}

type VkCoarseSampleLocationNV struct {
	PixelX uint32
	PixelY uint32
	Sample uint32
}

type VkCoarseSampleOrderCustomNV struct {
	ShadingRate         VkShadingRatePaletteEntryNV
	SampleCount         uint32
	SampleLocationCount uint32
	PSampleLocations    *VkCoarseSampleLocationNV
}

type VkPipelineViewportCoarseSampleOrderStateCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	SampleOrderType        VkCoarseSampleOrderTypeNV
	CustomSampleOrderCount uint32
	PCustomSampleOrders    *VkCoarseSampleOrderCustomNV
}

type VkPhysicalDeviceMeshShaderFeaturesNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	TaskShader VkBool32
	MeshShader VkBool32
}

type VkPhysicalDeviceMeshShaderPropertiesNV struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	MaxDrawMeshTasksCount             uint32
	MaxTaskWorkGroupInvocations       uint32
	MaxTaskWorkGroupSize              uint32
	MaxTaskTotalMemorySize            uint32
	MaxTaskOutputCount                uint32
	MaxMeshWorkGroupInvocations       uint32
	MaxMeshWorkGroupSize              uint32
	MaxMeshTotalMemorySize            uint32
	MaxMeshOutputVertices             uint32
	MaxMeshOutputPrimitives           uint32
	MaxMeshMultiviewViewCount         uint32
	MeshOutputPerVertexGranularity    uint32
	MeshOutputPerPrimitiveGranularity uint32
}

type VkDrawMeshTasksIndirectCommandNV struct {
	TaskCount uint32
	FirstTask uint32
}

type VkPhysicalDeviceMeshShaderFeaturesEXT struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	TaskShader                             VkBool32
	MeshShader                             VkBool32
	MultiviewMeshShader                    VkBool32
	PrimitiveFragmentShadingRateMeshShader VkBool32
	MeshShaderQueries                      VkBool32
}

type VkPhysicalDeviceMeshShaderPropertiesEXT struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	MaxTaskWorkGroupTotalCount            uint32
	MaxTaskWorkGroupCount                 uint32
	MaxTaskWorkGroupInvocations           uint32
	MaxTaskWorkGroupSize                  uint32
	MaxTaskPayloadSize                    uint32
	MaxTaskSharedMemorySize               uint32
	MaxTaskPayloadAndSharedMemorySize     uint32
	MaxMeshWorkGroupTotalCount            uint32
	MaxMeshWorkGroupCount                 uint32
	MaxMeshWorkGroupInvocations           uint32
	MaxMeshWorkGroupSize                  uint32
	MaxMeshSharedMemorySize               uint32
	MaxMeshPayloadAndSharedMemorySize     uint32
	MaxMeshOutputMemorySize               uint32
	MaxMeshPayloadAndOutputMemorySize     uint32
	MaxMeshOutputComponents               uint32
	MaxMeshOutputVertices                 uint32
	MaxMeshOutputPrimitives               uint32
	MaxMeshOutputLayers                   uint32
	MaxMeshMultiviewViewCount             uint32
	MeshOutputPerVertexGranularity        uint32
	MeshOutputPerPrimitiveGranularity     uint32
	MaxPreferredTaskWorkGroupInvocations  uint32
	MaxPreferredMeshWorkGroupInvocations  uint32
	PrefersLocalInvocationVertexOutput    VkBool32
	PrefersLocalInvocationPrimitiveOutput VkBool32
	PrefersCompactVertexOutput            VkBool32
	PrefersCompactPrimitiveOutput         VkBool32
}

type VkDrawMeshTasksIndirectCommandEXT struct {
	GroupCountX uint32
	GroupCountY uint32
	GroupCountZ uint32
}

type VkRayTracingShaderGroupCreateInfoNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Type               VkRayTracingShaderGroupTypeKHR
	GeneralShader      uint32
	ClosestHitShader   uint32
	AnyHitShader       uint32
	IntersectionShader uint32
}

type VkRayTracingShaderGroupCreateInfoKHR struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	Type                            VkRayTracingShaderGroupTypeKHR
	GeneralShader                   uint32
	ClosestHitShader                uint32
	AnyHitShader                    uint32
	IntersectionShader              uint32
	PShaderGroupCaptureReplayHandle unsafe.Pointer
}

type VkRayTracingPipelineCreateInfoNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkPipelineCreateFlags
	StageCount         uint32
	PStages            *VkPipelineShaderStageCreateInfo
	GroupCount         uint32
	PGroups            *VkRayTracingShaderGroupCreateInfoNV
	MaxRecursionDepth  uint32
	Layout             VkPipelineLayout
	BasePipelineHandle VkPipeline
	BasePipelineIndex  int32
}

type VkRayTracingPipelineCreateInfoKHR struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	Flags                        VkPipelineCreateFlags
	StageCount                   uint32
	PStages                      *VkPipelineShaderStageCreateInfo
	GroupCount                   uint32
	PGroups                      *VkRayTracingShaderGroupCreateInfoKHR
	MaxPipelineRayRecursionDepth uint32
	PLibraryInfo                 *VkPipelineLibraryCreateInfoKHR
	PLibraryInterface            *VkRayTracingPipelineInterfaceCreateInfoKHR
	PDynamicState                *VkPipelineDynamicStateCreateInfo
	Layout                       VkPipelineLayout
	BasePipelineHandle           VkPipeline
	BasePipelineIndex            int32
}

type VkGeometryTrianglesNV struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	VertexData      VkBuffer
	VertexOffset    VkDeviceSize
	VertexCount     uint32
	VertexStride    VkDeviceSize
	VertexFormat    VkFormat
	IndexData       VkBuffer
	IndexOffset     VkDeviceSize
	IndexCount      uint32
	IndexType       VkIndexType
	TransformData   VkBuffer
	TransformOffset VkDeviceSize
}

type VkGeometryAABBNV struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	AabbData VkBuffer
	NumAABBs uint32
	Stride   uint32
	Offset   VkDeviceSize
}

type VkGeometryDataNV struct {
	Triangles VkGeometryTrianglesNV
	Aabbs     VkGeometryAABBNV
}

type VkGeometryNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	GeometryType VkGeometryTypeKHR
	Geometry     VkGeometryDataNV
	Flags        VkGeometryFlagsKHR
}

type VkAccelerationStructureInfoNV struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Type          VkAccelerationStructureTypeNV
	Flags         VkBuildAccelerationStructureFlagsNV
	InstanceCount uint32
	GeometryCount uint32
	PGeometries   *VkGeometryNV
}

type VkAccelerationStructureCreateInfoNV struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	CompactedSize VkDeviceSize
	Info          VkAccelerationStructureInfoNV
}

type VkBindAccelerationStructureMemoryInfoNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	AccelerationStructure VkAccelerationStructureNV
	Memory                VkDeviceMemory
	MemoryOffset          VkDeviceSize
	DeviceIndexCount      uint32
	PDeviceIndices        *uint32
}

type VkWriteDescriptorSetAccelerationStructureKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	AccelerationStructureCount uint32
	PAccelerationStructures    *VkAccelerationStructureKHR
}

type VkWriteDescriptorSetAccelerationStructureNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	AccelerationStructureCount uint32
	PAccelerationStructures    *VkAccelerationStructureNV
}

type VkAccelerationStructureMemoryRequirementsInfoNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Type                  VkAccelerationStructureMemoryRequirementsTypeNV
	AccelerationStructure VkAccelerationStructureNV
}

type VkPhysicalDeviceAccelerationStructureFeaturesKHR struct {
	SType                                                 VkStructureType
	PNext                                                 unsafe.Pointer
	AccelerationStructure                                 VkBool32
	AccelerationStructureCaptureReplay                    VkBool32
	AccelerationStructureIndirectBuild                    VkBool32
	AccelerationStructureHostCommands                     VkBool32
	DescriptorBindingAccelerationStructureUpdateAfterBind VkBool32
}

type VkPhysicalDeviceRayTracingPipelineFeaturesKHR struct {
	SType                                                 VkStructureType
	PNext                                                 unsafe.Pointer
	RayTracingPipeline                                    VkBool32
	RayTracingPipelineShaderGroupHandleCaptureReplay      VkBool32
	RayTracingPipelineShaderGroupHandleCaptureReplayMixed VkBool32
	RayTracingPipelineTraceRaysIndirect                   VkBool32
	RayTraversalPrimitiveCulling                          VkBool32
}

type VkPhysicalDeviceRayQueryFeaturesKHR struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	RayQuery VkBool32
}

type VkPhysicalDeviceAccelerationStructurePropertiesKHR struct {
	SType                                                      VkStructureType
	PNext                                                      unsafe.Pointer
	MaxGeometryCount                                           uint64
	MaxInstanceCount                                           uint64
	MaxPrimitiveCount                                          uint64
	MaxPerStageDescriptorAccelerationStructures                uint32
	MaxPerStageDescriptorUpdateAfterBindAccelerationStructures uint32
	MaxDescriptorSetAccelerationStructures                     uint32
	MaxDescriptorSetUpdateAfterBindAccelerationStructures      uint32
	MinAccelerationStructureScratchOffsetAlignment             uint32
}

type VkPhysicalDeviceRayTracingPipelinePropertiesKHR struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	ShaderGroupHandleSize              uint32
	MaxRayRecursionDepth               uint32
	MaxShaderGroupStride               uint32
	ShaderGroupBaseAlignment           uint32
	ShaderGroupHandleCaptureReplaySize uint32
	MaxRayDispatchInvocationCount      uint32
	ShaderGroupHandleAlignment         uint32
	MaxRayHitAttributeSize             uint32
}

type VkPhysicalDeviceRayTracingPropertiesNV struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	ShaderGroupHandleSize                  uint32
	MaxRecursionDepth                      uint32
	MaxShaderGroupStride                   uint32
	ShaderGroupBaseAlignment               uint32
	MaxGeometryCount                       uint64
	MaxInstanceCount                       uint64
	MaxTriangleCount                       uint64
	MaxDescriptorSetAccelerationStructures uint32
}

type VkStridedDeviceAddressRegionKHR struct {
	DeviceAddress VkDeviceAddress
	Stride        VkDeviceSize
	Size          VkDeviceSize
}

type VkTraceRaysIndirectCommandKHR struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

type VkTraceRaysIndirectCommand2KHR struct {
	RaygenShaderRecordAddress         VkDeviceAddress
	RaygenShaderRecordSize            VkDeviceSize
	MissShaderBindingTableAddress     VkDeviceAddress
	MissShaderBindingTableSize        VkDeviceSize
	MissShaderBindingTableStride      VkDeviceSize
	HitShaderBindingTableAddress      VkDeviceAddress
	HitShaderBindingTableSize         VkDeviceSize
	HitShaderBindingTableStride       VkDeviceSize
	CallableShaderBindingTableAddress VkDeviceAddress
	CallableShaderBindingTableSize    VkDeviceSize
	CallableShaderBindingTableStride  VkDeviceSize
	Width                             uint32
	Height                            uint32
	Depth                             uint32
}

type VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR struct {
	SType                                VkStructureType
	PNext                                unsafe.Pointer
	RayTracingMaintenance1               VkBool32
	RayTracingPipelineTraceRaysIndirect2 VkBool32
}

type VkDrmFormatModifierPropertiesListEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	DrmFormatModifierCount       uint32
	PDrmFormatModifierProperties *VkDrmFormatModifierPropertiesEXT
}

type VkDrmFormatModifierPropertiesEXT struct {
	DrmFormatModifier               uint64
	DrmFormatModifierPlaneCount     uint32
	DrmFormatModifierTilingFeatures VkFormatFeatureFlags
}

type VkPhysicalDeviceImageDrmFormatModifierInfoEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	DrmFormatModifier     uint64
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}

type VkImageDrmFormatModifierListCreateInfoEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	DrmFormatModifierCount uint32
	PDrmFormatModifiers    *uint64
}

type VkImageDrmFormatModifierExplicitCreateInfoEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	DrmFormatModifier           uint64
	DrmFormatModifierPlaneCount uint32
	PPlaneLayouts               *VkSubresourceLayout
}

type VkImageDrmFormatModifierPropertiesEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DrmFormatModifier uint64
}

type VkImageStencilUsageCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	StencilUsage VkImageUsageFlags
}

type VkImageStencilUsageCreateInfoEXT struct {
}

type VkDeviceMemoryOverallocationCreateInfoAMD struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	OverallocationBehavior VkMemoryOverallocationBehaviorAMD
}

type VkPhysicalDeviceFragmentDensityMapFeaturesEXT struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	FragmentDensityMap                    VkBool32
	FragmentDensityMapDynamic             VkBool32
	FragmentDensityMapNonSubsampledImages VkBool32
}

type VkPhysicalDeviceFragmentDensityMap2FeaturesEXT struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	FragmentDensityMapDeferred VkBool32
}

type VkPhysicalDeviceFragmentDensityMapOffsetFeaturesEXT struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	FragmentDensityMapOffset VkBool32
}

type VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM struct {
}

type VkPhysicalDeviceFragmentDensityMapPropertiesEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	MinFragmentDensityTexelSize VkExtent2D
	MaxFragmentDensityTexelSize VkExtent2D
	FragmentDensityInvocations  VkBool32
}

type VkPhysicalDeviceFragmentDensityMap2PropertiesEXT struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	SubsampledLoads                           VkBool32
	SubsampledCoarseReconstructionEarlyAccess VkBool32
	MaxSubsampledArrayLayers                  uint32
	MaxDescriptorSetSubsampledSamplers        uint32
}

type VkPhysicalDeviceFragmentDensityMapOffsetPropertiesEXT struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	FragmentDensityOffsetGranularity VkExtent2D
}

type VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM struct {
}

type VkRenderPassFragmentDensityMapCreateInfoEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	FragmentDensityMapAttachment VkAttachmentReference
}

type VkRenderPassFragmentDensityMapOffsetEndInfoEXT struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	FragmentDensityOffsetCount uint32
	PFragmentDensityOffsets    *VkOffset2D
}

type VkSubpassFragmentDensityMapOffsetEndInfoQCOM struct {
}

type VkPhysicalDeviceScalarBlockLayoutFeatures struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	ScalarBlockLayout VkBool32
}

type VkPhysicalDeviceScalarBlockLayoutFeaturesEXT struct {
}

type VkSurfaceProtectedCapabilitiesKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	SupportsProtected VkBool32
}

type VkPhysicalDeviceUniformBufferStandardLayoutFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	UniformBufferStandardLayout VkBool32
}

type VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR struct {
}

type VkPhysicalDeviceDepthClipEnableFeaturesEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	DepthClipEnable VkBool32
}

type VkPipelineRasterizationDepthClipStateCreateInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkPipelineRasterizationDepthClipStateCreateFlagsEXT
	DepthClipEnable VkBool32
}

type VkPhysicalDeviceMemoryBudgetPropertiesEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HeapBudget VkDeviceSize
	HeapUsage  VkDeviceSize
}

type VkPhysicalDeviceMemoryPriorityFeaturesEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MemoryPriority VkBool32
}

type VkMemoryPriorityAllocateInfoEXT struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Priority float32
}

type VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	PageableDeviceLocalMemory VkBool32
}

type VkPhysicalDeviceBufferDeviceAddressFeatures struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	BufferDeviceAddress              VkBool32
	BufferDeviceAddressCaptureReplay VkBool32
	BufferDeviceAddressMultiDevice   VkBool32
}

type VkPhysicalDeviceBufferDeviceAddressFeaturesKHR struct {
}

type VkPhysicalDeviceBufferDeviceAddressFeaturesEXT struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	BufferDeviceAddress              VkBool32
	BufferDeviceAddressCaptureReplay VkBool32
	BufferDeviceAddressMultiDevice   VkBool32
}

type VkPhysicalDeviceBufferAddressFeaturesEXT struct {
}

type VkBufferDeviceAddressInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
}

type VkBufferDeviceAddressInfoKHR struct {
}

type VkBufferDeviceAddressInfoEXT struct {
}

type VkBufferOpaqueCaptureAddressCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	OpaqueCaptureAddress uint64
}

type VkBufferOpaqueCaptureAddressCreateInfoKHR struct {
}

type VkBufferDeviceAddressCreateInfoEXT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	DeviceAddress VkDeviceAddress
}

type VkPhysicalDeviceImageViewImageFormatInfoEXT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	ImageViewType VkImageViewType
}

type VkFilterCubicImageViewImageFormatPropertiesEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	FilterCubic       VkBool32
	FilterCubicMinmax VkBool32
}

type VkPhysicalDeviceImagelessFramebufferFeatures struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ImagelessFramebuffer VkBool32
}

type VkPhysicalDeviceImagelessFramebufferFeaturesKHR struct {
}

type VkFramebufferAttachmentsCreateInfo struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	AttachmentImageInfoCount uint32
	PAttachmentImageInfos    *VkFramebufferAttachmentImageInfo
}

type VkFramebufferAttachmentsCreateInfoKHR struct {
}

type VkFramebufferAttachmentImageInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkImageCreateFlags
	Usage           VkImageUsageFlags
	Width           uint32
	Height          uint32
	LayerCount      uint32
	ViewFormatCount uint32
	PViewFormats    *VkFormat
}

type VkFramebufferAttachmentImageInfoKHR struct {
}

type VkRenderPassAttachmentBeginInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	AttachmentCount uint32
	PAttachments    *VkImageView
}

type VkRenderPassAttachmentBeginInfoKHR struct {
}

type VkPhysicalDeviceTextureCompressionASTCHDRFeatures struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	TextureCompressionASTC_HDR VkBool32
}

type VkPhysicalDeviceTextureCompressionASTCHDRFeaturesEXT struct {
}

type VkPhysicalDeviceCooperativeMatrixFeaturesNV struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	CooperativeMatrix                   VkBool32
	CooperativeMatrixRobustBufferAccess VkBool32
}

type VkPhysicalDeviceCooperativeMatrixPropertiesNV struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	CooperativeMatrixSupportedStages VkShaderStageFlags
}

type VkCooperativeMatrixPropertiesNV struct {
	SType VkStructureType
	PNext unsafe.Pointer
	MSize uint32
	NSize uint32
	KSize uint32
	AType VkComponentTypeNV
	BType VkComponentTypeNV
	CType VkComponentTypeNV
	DType VkComponentTypeNV
	Scope VkScopeNV
}

type VkPhysicalDeviceYcbcrImageArraysFeaturesEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	YcbcrImageArrays VkBool32
}

type VkImageViewHandleInfoNVX struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	ImageView      VkImageView
	DescriptorType VkDescriptorType
	Sampler        VkSampler
}

type VkImageViewAddressPropertiesNVX struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	DeviceAddress VkDeviceAddress
	Size          VkDeviceSize
}

type VkPresentFrameTokenGGP struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	FrameToken GgpFrameToken
}

type VkPipelineCreationFeedback struct {
	Flags    VkPipelineCreationFeedbackFlags
	Duration uint64
}

type VkPipelineCreationFeedbackEXT struct {
}

type VkPipelineCreationFeedbackCreateInfo struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	PPipelineCreationFeedback          *VkPipelineCreationFeedback
	PipelineStageCreationFeedbackCount uint32
	PPipelineStageCreationFeedbacks    **VkPipelineCreationFeedback
}

type VkPipelineCreationFeedbackCreateInfoEXT struct {
}

type VkSurfaceFullScreenExclusiveInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	FullScreenExclusive VkFullScreenExclusiveEXT
}

type VkSurfaceFullScreenExclusiveWin32InfoEXT struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Hmonitor HMONITOR
}

type VkSurfaceCapabilitiesFullScreenExclusiveEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	FullScreenExclusiveSupported VkBool32
}

type VkPhysicalDevicePresentBarrierFeaturesNV struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PresentBarrier VkBool32
}

type VkSurfaceCapabilitiesPresentBarrierNV struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	PresentBarrierSupported VkBool32
}

type VkSwapchainPresentBarrierCreateInfoNV struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	PresentBarrierEnable VkBool32
}

type VkPhysicalDevicePerformanceQueryFeaturesKHR struct {
	SType                                VkStructureType
	PNext                                unsafe.Pointer
	PerformanceCounterQueryPools         VkBool32
	PerformanceCounterMultipleQueryPools VkBool32
}

type VkPhysicalDevicePerformanceQueryPropertiesKHR struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	AllowCommandBufferQueryCopies VkBool32
}

type VkPerformanceCounterKHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Unit    VkPerformanceCounterUnitKHR
	Scope   VkPerformanceCounterScopeKHR
	Storage VkPerformanceCounterStorageKHR
	Uuid    uint8
}

type VkPerformanceCounterDescriptionKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkPerformanceCounterDescriptionFlagsKHR
	Name        char
	Category    char
	Description char
}

type VkQueryPoolPerformanceCreateInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	QueueFamilyIndex  uint32
	CounterIndexCount uint32
	PCounterIndices   *uint32
}

type VkAcquireProfilingLockInfoKHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Flags   VkAcquireProfilingLockFlagsKHR
	Timeout uint64
}

type VkPerformanceQuerySubmitInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	CounterPassIndex uint32
}

type VkPerformanceQueryReservationInfoKHR struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MaxPerformanceQueriesPerPool uint32
}

type VkHeadlessSurfaceCreateInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkHeadlessSurfaceCreateFlagsEXT
}

type VkPhysicalDeviceCoverageReductionModeFeaturesNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	CoverageReductionMode VkBool32
}

type VkPipelineCoverageReductionStateCreateInfoNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkPipelineCoverageReductionStateCreateFlagsNV
	CoverageReductionMode VkCoverageReductionModeNV
}

type VkFramebufferMixedSamplesCombinationNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	CoverageReductionMode VkCoverageReductionModeNV
	RasterizationSamples  VkSampleCountFlagBits
	DepthStencilSamples   VkSampleCountFlags
	ColorSamples          VkSampleCountFlags
}

type VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ShaderIntegerFunctions2 VkBool32
}

type VkPerformanceValueINTEL struct {
	Type VkPerformanceValueTypeINTEL
	Data VkPerformanceValueDataINTEL
}

type VkInitializePerformanceApiInfoINTEL struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	PUserData unsafe.Pointer
}

type VkQueryPoolPerformanceQueryCreateInfoINTEL struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PerformanceCountersSampling VkQueryPoolSamplingModeINTEL
}

type VkQueryPoolCreateInfoINTEL struct {
}

type VkPerformanceMarkerInfoINTEL struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Marker uint64
}

type VkPerformanceStreamMarkerInfoINTEL struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Marker uint32
}

type VkPerformanceOverrideInfoINTEL struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Type      VkPerformanceOverrideTypeINTEL
	Enable    VkBool32
	Parameter uint64
}

type VkPerformanceConfigurationAcquireInfoINTEL struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Type  VkPerformanceConfigurationTypeINTEL
}

type VkPhysicalDeviceShaderClockFeaturesKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	ShaderSubgroupClock VkBool32
	ShaderDeviceClock   VkBool32
}

type VkPhysicalDeviceIndexTypeUint8Features struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	IndexTypeUint8 VkBool32
}

type VkPhysicalDeviceIndexTypeUint8FeaturesKHR struct {
}

type VkPhysicalDeviceIndexTypeUint8FeaturesEXT struct {
}

type VkPhysicalDeviceShaderSMBuiltinsPropertiesNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ShaderSMCount    uint32
	ShaderWarpsPerSM uint32
}

type VkPhysicalDeviceShaderSMBuiltinsFeaturesNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ShaderSMBuiltins VkBool32
}

type VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	FragmentShaderSampleInterlock      VkBool32
	FragmentShaderPixelInterlock       VkBool32
	FragmentShaderShadingRateInterlock VkBool32
}

type VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	SeparateDepthStencilLayouts VkBool32
}

type VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR struct {
}

type VkAttachmentReferenceStencilLayout struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	StencilLayout VkImageLayout
}

type VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	PrimitiveTopologyListRestart      VkBool32
	PrimitiveTopologyPatchListRestart VkBool32
}

type VkAttachmentReferenceStencilLayoutKHR struct {
}

type VkAttachmentDescriptionStencilLayout struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	StencilInitialLayout VkImageLayout
	StencilFinalLayout   VkImageLayout
}

type VkAttachmentDescriptionStencilLayoutKHR struct {
}

type VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	PipelineExecutableInfo VkBool32
}

type VkPipelineInfoKHR struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Pipeline VkPipeline
}

type VkPipelineInfoEXT struct {
}

type VkPipelineExecutablePropertiesKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Stages       VkShaderStageFlags
	Name         char
	Description  char
	SubgroupSize uint32
}

type VkPipelineExecutableInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Pipeline        VkPipeline
	ExecutableIndex uint32
}

type VkPipelineExecutableStatisticKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Name        char
	Description char
	Format      VkPipelineExecutableStatisticFormatKHR
	Value       VkPipelineExecutableStatisticValueKHR
}

type VkPipelineExecutableInternalRepresentationKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Name        char
	Description char
	IsText      VkBool32
	DataSize    uint
	PData       unsafe.Pointer
}

type VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	ShaderDemoteToHelperInvocation VkBool32
}

type VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT struct {
}

type VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	TexelBufferAlignment VkBool32
}

type VkPhysicalDeviceTexelBufferAlignmentProperties struct {
	SType                                        VkStructureType
	PNext                                        unsafe.Pointer
	StorageTexelBufferOffsetAlignmentBytes       VkDeviceSize
	StorageTexelBufferOffsetSingleTexelAlignment VkBool32
	UniformTexelBufferOffsetAlignmentBytes       VkDeviceSize
	UniformTexelBufferOffsetSingleTexelAlignment VkBool32
}

type VkPhysicalDeviceTexelBufferAlignmentPropertiesEXT struct {
}

type VkPhysicalDeviceSubgroupSizeControlFeatures struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	SubgroupSizeControl  VkBool32
	ComputeFullSubgroups VkBool32
}

type VkPhysicalDeviceSubgroupSizeControlFeaturesEXT struct {
}

type VkPhysicalDeviceSubgroupSizeControlProperties struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MinSubgroupSize              uint32
	MaxSubgroupSize              uint32
	MaxComputeWorkgroupSubgroups uint32
	RequiredSubgroupSizeStages   VkShaderStageFlags
}

type VkPhysicalDeviceSubgroupSizeControlPropertiesEXT struct {
}

type VkPipelineShaderStageRequiredSubgroupSizeCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	RequiredSubgroupSize uint32
}

type VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT struct {
}

type VkShaderRequiredSubgroupSizeCreateInfoEXT struct {
}

type VkSubpassShadingPipelineCreateInfoHUAWEI struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	RenderPass VkRenderPass
	Subpass    uint32
}

type VkPhysicalDeviceSubpassShadingPropertiesHUAWEI struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	MaxSubpassShadingWorkgroupSizeAspectRatio uint32
}

type VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	MaxWorkGroupCount             uint32
	MaxWorkGroupSize              uint32
	MaxOutputClusterCount         uint32
	IndirectBufferOffsetAlignment VkDeviceSize
}

type VkMemoryOpaqueCaptureAddressAllocateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	OpaqueCaptureAddress uint64
}

type VkMemoryOpaqueCaptureAddressAllocateInfoKHR struct {
}

type VkDeviceMemoryOpaqueCaptureAddressInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
}

type VkDeviceMemoryOpaqueCaptureAddressInfoKHR struct {
}

type VkPhysicalDeviceLineRasterizationFeatures struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	RectangularLines         VkBool32
	BresenhamLines           VkBool32
	SmoothLines              VkBool32
	StippledRectangularLines VkBool32
	StippledBresenhamLines   VkBool32
	StippledSmoothLines      VkBool32
}

type VkPhysicalDeviceLineRasterizationFeaturesKHR struct {
}

type VkPhysicalDeviceLineRasterizationFeaturesEXT struct {
}

type VkPhysicalDeviceLineRasterizationProperties struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	LineSubPixelPrecisionBits uint32
}

type VkPhysicalDeviceLineRasterizationPropertiesKHR struct {
}

type VkPhysicalDeviceLineRasterizationPropertiesEXT struct {
}

type VkPipelineRasterizationLineStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	LineRasterizationMode VkLineRasterizationMode
	StippledLineEnable    VkBool32
	LineStippleFactor     uint32
	LineStipplePattern    uint16
}

type VkPipelineRasterizationLineStateCreateInfoKHR struct {
}

type VkPipelineRasterizationLineStateCreateInfoEXT struct {
}

type VkPhysicalDevicePipelineCreationCacheControlFeatures struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	PipelineCreationCacheControl VkBool32
}

type VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT struct {
}

type VkPhysicalDeviceVulkan11Features struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	StorageBuffer16BitAccess           VkBool32
	UniformAndStorageBuffer16BitAccess VkBool32
	StoragePushConstant16              VkBool32
	StorageInputOutput16               VkBool32
	Multiview                          VkBool32
	MultiviewGeometryShader            VkBool32
	MultiviewTessellationShader        VkBool32
	VariablePointersStorageBuffer      VkBool32
	VariablePointers                   VkBool32
	ProtectedMemory                    VkBool32
	SamplerYcbcrConversion             VkBool32
	ShaderDrawParameters               VkBool32
}

type VkPhysicalDeviceVulkan11Properties struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	DeviceUUID                        uint8
	DriverUUID                        uint8
	DeviceLUID                        uint8
	DeviceNodeMask                    uint32
	DeviceLUIDValid                   VkBool32
	SubgroupSize                      uint32
	SubgroupSupportedStages           VkShaderStageFlags
	SubgroupSupportedOperations       VkSubgroupFeatureFlags
	SubgroupQuadOperationsInAllStages VkBool32
	PointClippingBehavior             VkPointClippingBehavior
	MaxMultiviewViewCount             uint32
	MaxMultiviewInstanceIndex         uint32
	ProtectedNoFault                  VkBool32
	MaxPerSetDescriptors              uint32
	MaxMemoryAllocationSize           VkDeviceSize
}

type VkPhysicalDeviceVulkan12Features struct {
	SType                                              VkStructureType
	PNext                                              unsafe.Pointer
	SamplerMirrorClampToEdge                           VkBool32
	DrawIndirectCount                                  VkBool32
	StorageBuffer8BitAccess                            VkBool32
	UniformAndStorageBuffer8BitAccess                  VkBool32
	StoragePushConstant8                               VkBool32
	ShaderBufferInt64Atomics                           VkBool32
	ShaderSharedInt64Atomics                           VkBool32
	ShaderFloat16                                      VkBool32
	ShaderInt8                                         VkBool32
	DescriptorIndexing                                 VkBool32
	ShaderInputAttachmentArrayDynamicIndexing          VkBool32
	ShaderUniformTexelBufferArrayDynamicIndexing       VkBool32
	ShaderStorageTexelBufferArrayDynamicIndexing       VkBool32
	ShaderUniformBufferArrayNonUniformIndexing         VkBool32
	ShaderSampledImageArrayNonUniformIndexing          VkBool32
	ShaderStorageBufferArrayNonUniformIndexing         VkBool32
	ShaderStorageImageArrayNonUniformIndexing          VkBool32
	ShaderInputAttachmentArrayNonUniformIndexing       VkBool32
	ShaderUniformTexelBufferArrayNonUniformIndexing    VkBool32
	ShaderStorageTexelBufferArrayNonUniformIndexing    VkBool32
	DescriptorBindingUniformBufferUpdateAfterBind      VkBool32
	DescriptorBindingSampledImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageImageUpdateAfterBind       VkBool32
	DescriptorBindingStorageBufferUpdateAfterBind      VkBool32
	DescriptorBindingUniformTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingStorageTexelBufferUpdateAfterBind VkBool32
	DescriptorBindingUpdateUnusedWhilePending          VkBool32
	DescriptorBindingPartiallyBound                    VkBool32
	DescriptorBindingVariableDescriptorCount           VkBool32
	RuntimeDescriptorArray                             VkBool32
	SamplerFilterMinmax                                VkBool32
	ScalarBlockLayout                                  VkBool32
	ImagelessFramebuffer                               VkBool32
	UniformBufferStandardLayout                        VkBool32
	ShaderSubgroupExtendedTypes                        VkBool32
	SeparateDepthStencilLayouts                        VkBool32
	HostQueryReset                                     VkBool32
	TimelineSemaphore                                  VkBool32
	BufferDeviceAddress                                VkBool32
	BufferDeviceAddressCaptureReplay                   VkBool32
	BufferDeviceAddressMultiDevice                     VkBool32
	VulkanMemoryModel                                  VkBool32
	VulkanMemoryModelDeviceScope                       VkBool32
	VulkanMemoryModelAvailabilityVisibilityChains      VkBool32
	ShaderOutputViewportIndex                          VkBool32
	ShaderOutputLayer                                  VkBool32
	SubgroupBroadcastDynamicId                         VkBool32
}

type VkPhysicalDeviceVulkan12Properties struct {
	SType                                                VkStructureType
	PNext                                                unsafe.Pointer
	DriverID                                             VkDriverId
	DriverName                                           char
	DriverInfo                                           char
	ConformanceVersion                                   VkConformanceVersion
	DenormBehaviorIndependence                           VkShaderFloatControlsIndependence
	RoundingModeIndependence                             VkShaderFloatControlsIndependence
	ShaderSignedZeroInfNanPreserveFloat16                VkBool32
	ShaderSignedZeroInfNanPreserveFloat32                VkBool32
	ShaderSignedZeroInfNanPreserveFloat64                VkBool32
	ShaderDenormPreserveFloat16                          VkBool32
	ShaderDenormPreserveFloat32                          VkBool32
	ShaderDenormPreserveFloat64                          VkBool32
	ShaderDenormFlushToZeroFloat16                       VkBool32
	ShaderDenormFlushToZeroFloat32                       VkBool32
	ShaderDenormFlushToZeroFloat64                       VkBool32
	ShaderRoundingModeRTEFloat16                         VkBool32
	ShaderRoundingModeRTEFloat32                         VkBool32
	ShaderRoundingModeRTEFloat64                         VkBool32
	ShaderRoundingModeRTZFloat16                         VkBool32
	ShaderRoundingModeRTZFloat32                         VkBool32
	ShaderRoundingModeRTZFloat64                         VkBool32
	MaxUpdateAfterBindDescriptorsInAllPools              uint32
	ShaderUniformBufferArrayNonUniformIndexingNative     VkBool32
	ShaderSampledImageArrayNonUniformIndexingNative      VkBool32
	ShaderStorageBufferArrayNonUniformIndexingNative     VkBool32
	ShaderStorageImageArrayNonUniformIndexingNative      VkBool32
	ShaderInputAttachmentArrayNonUniformIndexingNative   VkBool32
	RobustBufferAccessUpdateAfterBind                    VkBool32
	QuadDivergentImplicitLod                             VkBool32
	MaxPerStageDescriptorUpdateAfterBindSamplers         uint32
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindSampledImages    uint32
	MaxPerStageDescriptorUpdateAfterBindStorageImages    uint32
	MaxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	MaxPerStageUpdateAfterBindResources                  uint32
	MaxDescriptorSetUpdateAfterBindSamplers              uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffers        uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffers        uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindSampledImages         uint32
	MaxDescriptorSetUpdateAfterBindStorageImages         uint32
	MaxDescriptorSetUpdateAfterBindInputAttachments      uint32
	SupportedDepthResolveModes                           VkResolveModeFlags
	SupportedStencilResolveModes                         VkResolveModeFlags
	IndependentResolveNone                               VkBool32
	IndependentResolve                                   VkBool32
	FilterMinmaxSingleComponentFormats                   VkBool32
	FilterMinmaxImageComponentMapping                    VkBool32
	MaxTimelineSemaphoreValueDifference                  uint64
	FramebufferIntegerColorSampleCounts                  VkSampleCountFlags
}

type VkPhysicalDeviceVulkan13Features struct {
	SType                                              VkStructureType
	PNext                                              unsafe.Pointer
	RobustImageAccess                                  VkBool32
	InlineUniformBlock                                 VkBool32
	DescriptorBindingInlineUniformBlockUpdateAfterBind VkBool32
	PipelineCreationCacheControl                       VkBool32
	PrivateData                                        VkBool32
	ShaderDemoteToHelperInvocation                     VkBool32
	ShaderTerminateInvocation                          VkBool32
	SubgroupSizeControl                                VkBool32
	ComputeFullSubgroups                               VkBool32
	Synchronization2                                   VkBool32
	TextureCompressionASTC_HDR                         VkBool32
	ShaderZeroInitializeWorkgroupMemory                VkBool32
	DynamicRendering                                   VkBool32
	ShaderIntegerDotProduct                            VkBool32
	Maintenance4                                       VkBool32
}

type VkPhysicalDeviceVulkan13Properties struct {
	SType                                                                         VkStructureType
	PNext                                                                         unsafe.Pointer
	MinSubgroupSize                                                               uint32
	MaxSubgroupSize                                                               uint32
	MaxComputeWorkgroupSubgroups                                                  uint32
	RequiredSubgroupSizeStages                                                    VkShaderStageFlags
	MaxInlineUniformBlockSize                                                     uint32
	MaxPerStageDescriptorInlineUniformBlocks                                      uint32
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks                       uint32
	MaxDescriptorSetInlineUniformBlocks                                           uint32
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks                            uint32
	MaxInlineUniformTotalSize                                                     uint32
	IntegerDotProduct8BitUnsignedAccelerated                                      VkBool32
	IntegerDotProduct8BitSignedAccelerated                                        VkBool32
	IntegerDotProduct8BitMixedSignednessAccelerated                               VkBool32
	IntegerDotProduct4x8BitPackedUnsignedAccelerated                              VkBool32
	IntegerDotProduct4x8BitPackedSignedAccelerated                                VkBool32
	IntegerDotProduct4x8BitPackedMixedSignednessAccelerated                       VkBool32
	IntegerDotProduct16BitUnsignedAccelerated                                     VkBool32
	IntegerDotProduct16BitSignedAccelerated                                       VkBool32
	IntegerDotProduct16BitMixedSignednessAccelerated                              VkBool32
	IntegerDotProduct32BitUnsignedAccelerated                                     VkBool32
	IntegerDotProduct32BitSignedAccelerated                                       VkBool32
	IntegerDotProduct32BitMixedSignednessAccelerated                              VkBool32
	IntegerDotProduct64BitUnsignedAccelerated                                     VkBool32
	IntegerDotProduct64BitSignedAccelerated                                       VkBool32
	IntegerDotProduct64BitMixedSignednessAccelerated                              VkBool32
	IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated                VkBool32
	IntegerDotProductAccumulatingSaturating8BitSignedAccelerated                  VkBool32
	IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated         VkBool32
	IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated        VkBool32
	IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated          VkBool32
	IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated VkBool32
	IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated               VkBool32
	IntegerDotProductAccumulatingSaturating16BitSignedAccelerated                 VkBool32
	IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated        VkBool32
	IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated               VkBool32
	IntegerDotProductAccumulatingSaturating32BitSignedAccelerated                 VkBool32
	IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated        VkBool32
	IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated               VkBool32
	IntegerDotProductAccumulatingSaturating64BitSignedAccelerated                 VkBool32
	IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated        VkBool32
	StorageTexelBufferOffsetAlignmentBytes                                        VkDeviceSize
	StorageTexelBufferOffsetSingleTexelAlignment                                  VkBool32
	UniformTexelBufferOffsetAlignmentBytes                                        VkDeviceSize
	UniformTexelBufferOffsetSingleTexelAlignment                                  VkBool32
	MaxBufferSize                                                                 VkDeviceSize
}

type VkPhysicalDeviceVulkan14Features struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	GlobalPriorityQuery                    VkBool32
	ShaderSubgroupRotate                   VkBool32
	ShaderSubgroupRotateClustered          VkBool32
	ShaderFloatControls2                   VkBool32
	ShaderExpectAssume                     VkBool32
	RectangularLines                       VkBool32
	BresenhamLines                         VkBool32
	SmoothLines                            VkBool32
	StippledRectangularLines               VkBool32
	StippledBresenhamLines                 VkBool32
	StippledSmoothLines                    VkBool32
	VertexAttributeInstanceRateDivisor     VkBool32
	VertexAttributeInstanceRateZeroDivisor VkBool32
	IndexTypeUint8                         VkBool32
	DynamicRenderingLocalRead              VkBool32
	Maintenance5                           VkBool32
	Maintenance6                           VkBool32
	PipelineProtectedAccess                VkBool32
	PipelineRobustness                     VkBool32
	HostImageCopy                          VkBool32
	PushDescriptor                         VkBool32
}

type VkPhysicalDeviceVulkan14Properties struct {
	SType                                               VkStructureType
	PNext                                               unsafe.Pointer
	LineSubPixelPrecisionBits                           uint32
	MaxVertexAttribDivisor                              uint32
	SupportsNonZeroFirstInstance                        VkBool32
	MaxPushDescriptors                                  uint32
	DynamicRenderingLocalReadDepthStencilAttachments    VkBool32
	DynamicRenderingLocalReadMultisampledAttachments    VkBool32
	EarlyFragmentMultisampleCoverageAfterSampleCounting VkBool32
	EarlyFragmentSampleMaskTestBeforeSampleCounting     VkBool32
	DepthStencilSwizzleOneSupport                       VkBool32
	PolygonModePointSize                                VkBool32
	NonStrictSinglePixelWideLinesUseParallelogram       VkBool32
	NonStrictWideLinesUseParallelogram                  VkBool32
	BlockTexelViewCompatibleMultipleLayers              VkBool32
	MaxCombinedImageSamplerDescriptorCount              uint32
	FragmentShadingRateClampCombinerInputs              VkBool32
	DefaultRobustnessStorageBuffers                     VkPipelineRobustnessBufferBehavior
	DefaultRobustnessUniformBuffers                     VkPipelineRobustnessBufferBehavior
	DefaultRobustnessVertexInputs                       VkPipelineRobustnessBufferBehavior
	DefaultRobustnessImages                             VkPipelineRobustnessImageBehavior
	CopySrcLayoutCount                                  uint32
	PCopySrcLayouts                                     *VkImageLayout
	CopyDstLayoutCount                                  uint32
	PCopyDstLayouts                                     *VkImageLayout
	OptimalTilingLayoutUUID                             uint8
	IdenticalMemoryTypeRequirements                     VkBool32
}

type VkPipelineCompilerControlCreateInfoAMD struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	CompilerControlFlags VkPipelineCompilerControlFlagsAMD
}

type VkPhysicalDeviceCoherentMemoryFeaturesAMD struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	DeviceCoherentMemory VkBool32
}

type VkFaultData struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	FaultLevel VkFaultLevel
	FaultType  VkFaultType
}

type VkFaultCallbackInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	FaultCount       uint32
	PFaults          *VkFaultData
	PfnFaultCallback PFN_vkFaultCallbackFunction
}

type VkPhysicalDeviceToolProperties struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Name        char
	Version     char
	Purposes    VkToolPurposeFlags
	Description char
	Layer       char
}

type VkPhysicalDeviceToolPropertiesEXT struct {
}

type VkSamplerCustomBorderColorCreateInfoEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	CustomBorderColor VkClearColorValue
	Format            VkFormat
}

type VkPhysicalDeviceCustomBorderColorPropertiesEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MaxCustomBorderColorSamplers uint32
}

type VkPhysicalDeviceCustomBorderColorFeaturesEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	CustomBorderColors             VkBool32
	CustomBorderColorWithoutFormat VkBool32
}

type VkSamplerBorderColorComponentMappingCreateInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Components VkComponentMapping
	Srgb       VkBool32
}

type VkPhysicalDeviceBorderColorSwizzleFeaturesEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	BorderColorSwizzle          VkBool32
	BorderColorSwizzleFromImage VkBool32
}

type VkAccelerationStructureGeometryTrianglesDataKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	VertexFormat  VkFormat
	VertexData    VkDeviceOrHostAddressConstKHR
	VertexStride  VkDeviceSize
	MaxVertex     uint32
	IndexType     VkIndexType
	IndexData     VkDeviceOrHostAddressConstKHR
	TransformData VkDeviceOrHostAddressConstKHR
}

type VkAccelerationStructureGeometryAabbsDataKHR struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Data   VkDeviceOrHostAddressConstKHR
	Stride VkDeviceSize
}

type VkAccelerationStructureGeometryInstancesDataKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ArrayOfPointers VkBool32
	Data            VkDeviceOrHostAddressConstKHR
}

type VkAccelerationStructureGeometryLinearSweptSpheresDataNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	VertexFormat VkFormat
	VertexData   VkDeviceOrHostAddressConstKHR
	VertexStride VkDeviceSize
	RadiusFormat VkFormat
	RadiusData   VkDeviceOrHostAddressConstKHR
	RadiusStride VkDeviceSize
	IndexType    VkIndexType
	IndexData    VkDeviceOrHostAddressConstKHR
	IndexStride  VkDeviceSize
	IndexingMode VkRayTracingLssIndexingModeNV
	EndCapsMode  VkRayTracingLssPrimitiveEndCapsModeNV
}

type VkAccelerationStructureGeometrySpheresDataNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	VertexFormat VkFormat
	VertexData   VkDeviceOrHostAddressConstKHR
	VertexStride VkDeviceSize
	RadiusFormat VkFormat
	RadiusData   VkDeviceOrHostAddressConstKHR
	RadiusStride VkDeviceSize
	IndexType    VkIndexType
	IndexData    VkDeviceOrHostAddressConstKHR
	IndexStride  VkDeviceSize
}

type VkAccelerationStructureGeometryKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	GeometryType VkGeometryTypeKHR
	Geometry     VkAccelerationStructureGeometryDataKHR
	Flags        VkGeometryFlagsKHR
}

type VkAccelerationStructureBuildGeometryInfoKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	Type                     VkAccelerationStructureTypeKHR
	Flags                    VkBuildAccelerationStructureFlagsKHR
	Mode                     VkBuildAccelerationStructureModeKHR
	SrcAccelerationStructure VkAccelerationStructureKHR
	DstAccelerationStructure VkAccelerationStructureKHR
	GeometryCount            uint32
	PGeometries              *VkAccelerationStructureGeometryKHR
	PpGeometries             **VkAccelerationStructureGeometryKHR
	ScratchData              VkDeviceOrHostAddressKHR
}

type VkAccelerationStructureBuildRangeInfoKHR struct {
	PrimitiveCount  uint32
	PrimitiveOffset uint32
	FirstVertex     uint32
	TransformOffset uint32
}

type VkAccelerationStructureCreateInfoKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	CreateFlags   VkAccelerationStructureCreateFlagsKHR
	Buffer        VkBuffer
	Offset        VkDeviceSize
	Size          VkDeviceSize
	Type          VkAccelerationStructureTypeKHR
	DeviceAddress VkDeviceAddress
}

type VkAabbPositionsKHR struct {
	MinX float32
	MinY float32
	MinZ float32
	MaxX float32
	MaxY float32
	MaxZ float32
}

type VkAabbPositionsNV struct {
}

type VkTransformMatrixKHR struct {
	Matrix float32
}

type VkTransformMatrixNV struct {
}

type VkAccelerationStructureInstanceKHR struct {
	Transform                              VkTransformMatrixKHR
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  VkGeometryInstanceFlagsKHR
	AccelerationStructureReference         uint64
}

type VkAccelerationStructureInstanceNV struct {
}

type VkAccelerationStructureDeviceAddressInfoKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	AccelerationStructure VkAccelerationStructureKHR
}

type VkAccelerationStructureVersionInfoKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	PVersionData *uint8
}

type VkCopyAccelerationStructureInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Src   VkAccelerationStructureKHR
	Dst   VkAccelerationStructureKHR
	Mode  VkCopyAccelerationStructureModeKHR
}

type VkCopyAccelerationStructureToMemoryInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Src   VkAccelerationStructureKHR
	Dst   VkDeviceOrHostAddressKHR
	Mode  VkCopyAccelerationStructureModeKHR
}

type VkCopyMemoryToAccelerationStructureInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Src   VkDeviceOrHostAddressConstKHR
	Dst   VkAccelerationStructureKHR
	Mode  VkCopyAccelerationStructureModeKHR
}

type VkRayTracingPipelineInterfaceCreateInfoKHR struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	MaxPipelineRayPayloadSize      uint32
	MaxPipelineRayHitAttributeSize uint32
}

type VkPipelineLibraryCreateInfoKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	LibraryCount uint32
	PLibraries   *VkPipeline
}

type VkRefreshObjectKHR struct {
	ObjectType   VkObjectType
	ObjectHandle uint64
	Flags        VkRefreshObjectFlagsKHR
}

type VkRefreshObjectListKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	ObjectCount uint32
	PObjects    *VkRefreshObjectKHR
}

type VkPhysicalDeviceExtendedDynamicStateFeaturesEXT struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ExtendedDynamicState VkBool32
}

type VkPhysicalDeviceExtendedDynamicState2FeaturesEXT struct {
	SType                                   VkStructureType
	PNext                                   unsafe.Pointer
	ExtendedDynamicState2                   VkBool32
	ExtendedDynamicState2LogicOp            VkBool32
	ExtendedDynamicState2PatchControlPoints VkBool32
}

type VkPhysicalDeviceExtendedDynamicState3FeaturesEXT struct {
	SType                                                 VkStructureType
	PNext                                                 unsafe.Pointer
	ExtendedDynamicState3TessellationDomainOrigin         VkBool32
	ExtendedDynamicState3DepthClampEnable                 VkBool32
	ExtendedDynamicState3PolygonMode                      VkBool32
	ExtendedDynamicState3RasterizationSamples             VkBool32
	ExtendedDynamicState3SampleMask                       VkBool32
	ExtendedDynamicState3AlphaToCoverageEnable            VkBool32
	ExtendedDynamicState3AlphaToOneEnable                 VkBool32
	ExtendedDynamicState3LogicOpEnable                    VkBool32
	ExtendedDynamicState3ColorBlendEnable                 VkBool32
	ExtendedDynamicState3ColorBlendEquation               VkBool32
	ExtendedDynamicState3ColorWriteMask                   VkBool32
	ExtendedDynamicState3RasterizationStream              VkBool32
	ExtendedDynamicState3ConservativeRasterizationMode    VkBool32
	ExtendedDynamicState3ExtraPrimitiveOverestimationSize VkBool32
	ExtendedDynamicState3DepthClipEnable                  VkBool32
	ExtendedDynamicState3SampleLocationsEnable            VkBool32
	ExtendedDynamicState3ColorBlendAdvanced               VkBool32
	ExtendedDynamicState3ProvokingVertexMode              VkBool32
	ExtendedDynamicState3LineRasterizationMode            VkBool32
	ExtendedDynamicState3LineStippleEnable                VkBool32
	ExtendedDynamicState3DepthClipNegativeOneToOne        VkBool32
	ExtendedDynamicState3ViewportWScalingEnable           VkBool32
	ExtendedDynamicState3ViewportSwizzle                  VkBool32
	ExtendedDynamicState3CoverageToColorEnable            VkBool32
	ExtendedDynamicState3CoverageToColorLocation          VkBool32
	ExtendedDynamicState3CoverageModulationMode           VkBool32
	ExtendedDynamicState3CoverageModulationTableEnable    VkBool32
	ExtendedDynamicState3CoverageModulationTable          VkBool32
	ExtendedDynamicState3CoverageReductionMode            VkBool32
	ExtendedDynamicState3RepresentativeFragmentTestEnable VkBool32
	ExtendedDynamicState3ShadingRateImageEnable           VkBool32
}

type VkPhysicalDeviceExtendedDynamicState3PropertiesEXT struct {
	SType                                VkStructureType
	PNext                                unsafe.Pointer
	DynamicPrimitiveTopologyUnrestricted VkBool32
}

type VkColorBlendEquationEXT struct {
	SrcColorBlendFactor VkBlendFactor
	DstColorBlendFactor VkBlendFactor
	ColorBlendOp        VkBlendOp
	SrcAlphaBlendFactor VkBlendFactor
	DstAlphaBlendFactor VkBlendFactor
	AlphaBlendOp        VkBlendOp
}

type VkColorBlendAdvancedEXT struct {
	AdvancedBlendOp  VkBlendOp
	SrcPremultiplied VkBool32
	DstPremultiplied VkBool32
	BlendOverlap     VkBlendOverlapEXT
	ClampResults     VkBool32
}

type VkRenderPassTransformBeginInfoQCOM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Transform VkSurfaceTransformFlagBitsKHR
}

type VkCopyCommandTransformInfoQCOM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Transform VkSurfaceTransformFlagBitsKHR
}

type VkCommandBufferInheritanceRenderPassTransformInfoQCOM struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Transform  VkSurfaceTransformFlagBitsKHR
	RenderArea VkRect2D
}

type VkPhysicalDevicePartitionedAccelerationStructureFeaturesNV struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	PartitionedAccelerationStructure VkBool32
}

type VkPhysicalDevicePartitionedAccelerationStructurePropertiesNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	MaxPartitionCount uint32
}

type VkBuildPartitionedAccelerationStructureIndirectCommandNV struct {
	OpType   VkPartitionedAccelerationStructureOpTypeNV
	ArgCount uint32
	ArgData  VkStridedDeviceAddressNV
}

type VkPartitionedAccelerationStructureFlagsNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	EnablePartitionTranslation VkBool32
}

type VkPartitionedAccelerationStructureWriteInstanceDataNV struct {
	Transform                           VkTransformMatrixKHR
	ExplicitAABB                        float32
	InstanceID                          uint32
	InstanceMask                        uint32
	InstanceContributionToHitGroupIndex uint32
	InstanceFlags                       VkPartitionedAccelerationStructureInstanceFlagsNV
	InstanceIndex                       uint32
	PartitionIndex                      uint32
	AccelerationStructure               VkDeviceAddress
}

type VkPartitionedAccelerationStructureUpdateInstanceDataNV struct {
	InstanceIndex                       uint32
	InstanceContributionToHitGroupIndex uint32
	AccelerationStructure               VkDeviceAddress
}

type VkPartitionedAccelerationStructureWritePartitionTranslationDataNV struct {
	PartitionIndex       uint32
	PartitionTranslation float32
}

type VkWriteDescriptorSetPartitionedAccelerationStructureNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	AccelerationStructureCount uint32
	PAccelerationStructures    *VkDeviceAddress
}

type VkPartitionedAccelerationStructureInstancesInputNV struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	Flags                             VkBuildAccelerationStructureFlagsKHR
	InstanceCount                     uint32
	MaxInstancePerPartitionCount      uint32
	PartitionCount                    uint32
	MaxInstanceInGlobalPartitionCount uint32
}

type VkBuildPartitionedAccelerationStructureInfoNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	Input                        VkPartitionedAccelerationStructureInstancesInputNV
	SrcAccelerationStructureData VkDeviceAddress
	DstAccelerationStructureData VkDeviceAddress
	ScratchData                  VkDeviceAddress
	SrcInfos                     VkDeviceAddress
	SrcInfosCount                VkDeviceAddress
}

type VkPhysicalDeviceDiagnosticsConfigFeaturesNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DiagnosticsConfig VkBool32
}

type VkDeviceDiagnosticsConfigCreateInfoNV struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkDeviceDiagnosticsConfigFlagsNV
}

type VkPipelineOfflineCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PipelineIdentifier uint8
	MatchControl       VkPipelineMatchControl
	PoolEntrySize      VkDeviceSize
}

type VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	ShaderZeroInitializeWorkgroupMemory VkBool32
}

type VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR struct {
}

type VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	ShaderSubgroupUniformControlFlow VkBool32
}

type VkPhysicalDeviceRobustness2FeaturesKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	RobustBufferAccess2 VkBool32
	RobustImageAccess2  VkBool32
	NullDescriptor      VkBool32
}

type VkPhysicalDeviceRobustness2FeaturesEXT struct {
}

type VkPhysicalDeviceRobustness2PropertiesKHR struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	RobustStorageBufferAccessSizeAlignment VkDeviceSize
	RobustUniformBufferAccessSizeAlignment VkDeviceSize
}

type VkPhysicalDeviceRobustness2PropertiesEXT struct {
}

type VkPhysicalDeviceImageRobustnessFeatures struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	RobustImageAccess VkBool32
}

type VkPhysicalDeviceImageRobustnessFeaturesEXT struct {
}

type VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR struct {
	SType                                          VkStructureType
	PNext                                          unsafe.Pointer
	WorkgroupMemoryExplicitLayout                  VkBool32
	WorkgroupMemoryExplicitLayoutScalarBlockLayout VkBool32
	WorkgroupMemoryExplicitLayout8BitAccess        VkBool32
	WorkgroupMemoryExplicitLayout16BitAccess       VkBool32
}

type VkPhysicalDevicePortabilitySubsetFeaturesKHR struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	ConstantAlphaColorBlendFactors         VkBool32
	Events                                 VkBool32
	ImageViewFormatReinterpretation        VkBool32
	ImageViewFormatSwizzle                 VkBool32
	ImageView2DOn3DImage                   VkBool32
	MultisampleArrayImage                  VkBool32
	MutableComparisonSamplers              VkBool32
	PointPolygons                          VkBool32
	SamplerMipLodBias                      VkBool32
	SeparateStencilMaskRef                 VkBool32
	ShaderSampleRateInterpolationFunctions VkBool32
	TessellationIsolines                   VkBool32
	TessellationPointMode                  VkBool32
	TriangleFans                           VkBool32
	VertexAttributeAccessBeyondStride      VkBool32
}

type VkPhysicalDevicePortabilitySubsetPropertiesKHR struct {
	SType                                VkStructureType
	PNext                                unsafe.Pointer
	MinVertexInputBindingStrideAlignment uint32
}

type VkPhysicalDevice4444FormatsFeaturesEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	FormatA4R4G4B4 VkBool32
	FormatA4B4G4R4 VkBool32
}

type VkPhysicalDeviceSubpassShadingFeaturesHUAWEI struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SubpassShading VkBool32
}

type VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ClustercullingShader          VkBool32
	MultiviewClusterCullingShader VkBool32
}

type VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ClusterShadingRate VkBool32
}

type VkBufferCopy2 struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	SrcOffset VkDeviceSize
	DstOffset VkDeviceSize
	Size      VkDeviceSize
}

type VkBufferCopy2KHR struct {
}

type VkImageCopy2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}

type VkImageCopy2KHR struct {
}

type VkImageBlit2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcSubresource VkImageSubresourceLayers
	SrcOffsets     VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffsets     VkOffset3D
}

type VkImageBlit2KHR struct {
}

type VkBufferImageCopy2 struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	BufferOffset      VkDeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}

type VkBufferImageCopy2KHR struct {
}

type VkImageResolve2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}

type VkImageResolve2KHR struct {
}

type VkCopyBufferInfo2 struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	SrcBuffer   VkBuffer
	DstBuffer   VkBuffer
	RegionCount uint32
	PRegions    *VkBufferCopy2
}

type VkCopyBufferInfo2KHR struct {
}

type VkCopyImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageCopy2
}

type VkCopyImageInfo2KHR struct {
}

type VkBlitImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageBlit2
	Filter         VkFilter
}

type VkBlitImageInfo2KHR struct {
}

type VkCopyBufferToImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcBuffer      VkBuffer
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkBufferImageCopy2
}

type VkCopyBufferToImageInfo2KHR struct {
}

type VkCopyImageToBufferInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstBuffer      VkBuffer
	RegionCount    uint32
	PRegions       *VkBufferImageCopy2
}

type VkCopyImageToBufferInfo2KHR struct {
}

type VkResolveImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageResolve2
}

type VkResolveImageInfo2KHR struct {
}

type VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ShaderImageInt64Atomics VkBool32
	SparseImageInt64Atomics VkBool32
}

type VkFragmentShadingRateAttachmentInfoKHR struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	PFragmentShadingRateAttachment *VkAttachmentReference2
	ShadingRateAttachmentTexelSize VkExtent2D
}

type VkPipelineFragmentShadingRateStateCreateInfoKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	FragmentSize VkExtent2D
	CombinerOps  VkFragmentShadingRateCombinerOpKHR
}

type VkPhysicalDeviceFragmentShadingRateFeaturesKHR struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	PipelineFragmentShadingRate   VkBool32
	PrimitiveFragmentShadingRate  VkBool32
	AttachmentFragmentShadingRate VkBool32
}

type VkPhysicalDeviceFragmentShadingRatePropertiesKHR struct {
	SType                                                VkStructureType
	PNext                                                unsafe.Pointer
	MinFragmentShadingRateAttachmentTexelSize            VkExtent2D
	MaxFragmentShadingRateAttachmentTexelSize            VkExtent2D
	MaxFragmentShadingRateAttachmentTexelSizeAspectRatio uint32
	PrimitiveFragmentShadingRateWithMultipleViewports    VkBool32
	LayeredShadingRateAttachments                        VkBool32
	FragmentShadingRateNonTrivialCombinerOps             VkBool32
	MaxFragmentSize                                      VkExtent2D
	MaxFragmentSizeAspectRatio                           uint32
	MaxFragmentShadingRateCoverageSamples                uint32
	MaxFragmentShadingRateRasterizationSamples           VkSampleCountFlagBits
	FragmentShadingRateWithShaderDepthStencilWrites      VkBool32
	FragmentShadingRateWithSampleMask                    VkBool32
	FragmentShadingRateWithShaderSampleMask              VkBool32
	FragmentShadingRateWithConservativeRasterization     VkBool32
	FragmentShadingRateWithFragmentShaderInterlock       VkBool32
	FragmentShadingRateWithCustomSampleLocations         VkBool32
	FragmentShadingRateStrictMultiplyCombiner            VkBool32
}

type VkPhysicalDeviceFragmentShadingRateKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	SampleCounts VkSampleCountFlags
	FragmentSize VkExtent2D
}

type VkPhysicalDeviceShaderTerminateInvocationFeatures struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	ShaderTerminateInvocation VkBool32
}

type VkPhysicalDeviceShaderTerminateInvocationFeaturesKHR struct {
}

type VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	FragmentShadingRateEnums         VkBool32
	SupersampleFragmentShadingRates  VkBool32
	NoInvocationFragmentShadingRates VkBool32
}

type VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	MaxFragmentShadingRateInvocationCount VkSampleCountFlagBits
}

type VkPipelineFragmentShadingRateEnumStateCreateInfoNV struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ShadingRateType VkFragmentShadingRateTypeNV
	ShadingRate     VkFragmentShadingRateNV
	CombinerOps     VkFragmentShadingRateCombinerOpKHR
}

type VkAccelerationStructureBuildSizesInfoKHR struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	AccelerationStructureSize VkDeviceSize
	UpdateScratchSize         VkDeviceSize
	BuildScratchSize          VkDeviceSize
}

type VkPhysicalDeviceImage2DViewOf3DFeaturesEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Image2DViewOf3D   VkBool32
	Sampler2DViewOf3D VkBool32
}

type VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	ImageSlicedViewOf3D VkBool32
}

type VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	AttachmentFeedbackLoopDynamicState VkBool32
}

type VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	LegacyVertexAttributes VkBool32
}

type VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	NativeUnalignedPerformance VkBool32
}

type VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	MutableDescriptorType VkBool32
}

type VkPhysicalDeviceMutableDescriptorTypeFeaturesVALVE struct {
}

type VkMutableDescriptorTypeListEXT struct {
	DescriptorTypeCount uint32
	PDescriptorTypes    *VkDescriptorType
}

type VkMutableDescriptorTypeListVALVE struct {
}

type VkMutableDescriptorTypeCreateInfoEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	MutableDescriptorTypeListCount uint32
	PMutableDescriptorTypeLists    *VkMutableDescriptorTypeListEXT
}

type VkMutableDescriptorTypeCreateInfoVALVE struct {
}

type VkPhysicalDeviceDepthClipControlFeaturesEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DepthClipControl VkBool32
}

type VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ZeroInitializeDeviceMemory VkBool32
}

type VkBeginCustomResolveInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkPhysicalDeviceCustomResolveFeaturesEXT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	CustomResolve VkBool32
}

type VkCustomResolveCreateInfoEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	CustomResolve           VkBool32
	ColorAttachmentCount    uint32
	PColorAttachmentFormats *VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
}

type VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	DeviceGeneratedCommands        VkBool32
	DynamicGeneratedPipelineLayout VkBool32
}

type VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT struct {
	SType                                                VkStructureType
	PNext                                                unsafe.Pointer
	MaxIndirectPipelineCount                             uint32
	MaxIndirectShaderObjectCount                         uint32
	MaxIndirectSequenceCount                             uint32
	MaxIndirectCommandsTokenCount                        uint32
	MaxIndirectCommandsTokenOffset                       uint32
	MaxIndirectCommandsIndirectStride                    uint32
	SupportedIndirectCommandsInputModes                  VkIndirectCommandsInputModeFlagsEXT
	SupportedIndirectCommandsShaderStages                VkShaderStageFlags
	SupportedIndirectCommandsShaderStagesPipelineBinding VkShaderStageFlags
	SupportedIndirectCommandsShaderStagesShaderBinding   VkShaderStageFlags
	DeviceGeneratedCommandsTransformFeedback             VkBool32
	DeviceGeneratedCommandsMultiDrawIndirectCount        VkBool32
}

type VkGeneratedCommandsPipelineInfoEXT struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Pipeline VkPipeline
}

type VkGeneratedCommandsShaderInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	ShaderCount uint32
	PShaders    *VkShaderEXT
}

type VkGeneratedCommandsMemoryRequirementsInfoEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	IndirectExecutionSet   VkIndirectExecutionSetEXT
	IndirectCommandsLayout VkIndirectCommandsLayoutEXT
	MaxSequenceCount       uint32
	MaxDrawCount           uint32
}

type VkIndirectExecutionSetPipelineInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	InitialPipeline  VkPipeline
	MaxPipelineCount uint32
}

type VkIndirectExecutionSetShaderLayoutInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SetLayoutCount uint32
	PSetLayouts    *VkDescriptorSetLayout
}

type VkIndirectExecutionSetShaderInfoEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ShaderCount            uint32
	PInitialShaders        *VkShaderEXT
	PSetLayoutInfos        *VkIndirectExecutionSetShaderLayoutInfoEXT
	MaxShaderCount         uint32
	PushConstantRangeCount uint32
	PPushConstantRanges    *VkPushConstantRange
}

type VkIndirectExecutionSetCreateInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Type  VkIndirectExecutionSetInfoTypeEXT
	Info  VkIndirectExecutionSetInfoEXT
}

type VkGeneratedCommandsInfoEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ShaderStages           VkShaderStageFlags
	IndirectExecutionSet   VkIndirectExecutionSetEXT
	IndirectCommandsLayout VkIndirectCommandsLayoutEXT
	IndirectAddress        VkDeviceAddress
	IndirectAddressSize    VkDeviceSize
	PreprocessAddress      VkDeviceAddress
	PreprocessSize         VkDeviceSize
	MaxSequenceCount       uint32
	SequenceCountAddress   VkDeviceAddress
	MaxDrawCount           uint32
}

type VkWriteIndirectExecutionSetPipelineEXT struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Index    uint32
	Pipeline VkPipeline
}

type VkWriteIndirectExecutionSetShaderEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Index  uint32
	Shader VkShaderEXT
}

type VkIndirectCommandsLayoutCreateInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkIndirectCommandsLayoutUsageFlagsEXT
	ShaderStages   VkShaderStageFlags
	IndirectStride uint32
	PipelineLayout VkPipelineLayout
	TokenCount     uint32
	PTokens        *VkIndirectCommandsLayoutTokenEXT
}

type VkIndirectCommandsLayoutTokenEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Type   VkIndirectCommandsTokenTypeEXT
	Data   VkIndirectCommandsTokenDataEXT
	Offset uint32
}

type VkDrawIndirectCountIndirectCommandEXT struct {
	BufferAddress VkDeviceAddress
	Stride        uint32
	CommandCount  uint32
}

type VkIndirectCommandsVertexBufferTokenEXT struct {
	VertexBindingUnit uint32
}

type VkBindVertexBufferIndirectCommandEXT struct {
	BufferAddress VkDeviceAddress
	Size          uint32
	Stride        uint32
}

type VkIndirectCommandsIndexBufferTokenEXT struct {
	Mode VkIndirectCommandsInputModeFlagBitsEXT
}

type VkBindIndexBufferIndirectCommandEXT struct {
	BufferAddress VkDeviceAddress
	Size          uint32
	IndexType     VkIndexType
}

type VkIndirectCommandsPushConstantTokenEXT struct {
	UpdateRange VkPushConstantRange
}

type VkIndirectCommandsExecutionSetTokenEXT struct {
	Type         VkIndirectExecutionSetInfoTypeEXT
	ShaderStages VkShaderStageFlags
}

type VkPipelineViewportDepthClipControlCreateInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	NegativeOneToOne VkBool32
}

type VkPhysicalDeviceDepthClampControlFeaturesEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DepthClampControl VkBool32
}

type VkPipelineViewportDepthClampControlCreateInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DepthClampMode   VkDepthClampModeEXT
	PDepthClampRange *VkDepthClampRangeEXT
}

type VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	VertexInputDynamicState VkBool32
}

type VkPhysicalDeviceExternalMemoryRDMAFeaturesNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ExternalMemoryRDMA VkBool32
}

type VkPhysicalDeviceShaderRelaxedExtendedInstructionFeaturesKHR struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	ShaderRelaxedExtendedInstruction VkBool32
}

type VkVertexInputBindingDescription2EXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Binding   uint32
	Stride    uint32
	InputRate VkVertexInputRate
	Divisor   uint32
}

type VkVertexInputAttributeDescription2EXT struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Location uint32
	Binding  uint32
	Format   VkFormat
	Offset   uint32
}

type VkPhysicalDeviceColorWriteEnableFeaturesEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ColorWriteEnable VkBool32
}

type VkPipelineColorWriteCreateInfoEXT struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	AttachmentCount    uint32
	PColorWriteEnables *VkBool32
}

type VkMemoryBarrier2 struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SrcStageMask  VkPipelineStageFlags2
	SrcAccessMask VkAccessFlags2
	DstStageMask  VkPipelineStageFlags2
	DstAccessMask VkAccessFlags2
}

type VkMemoryBarrier2KHR struct {
}

type VkImageMemoryBarrier2 struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	SrcStageMask        VkPipelineStageFlags2
	SrcAccessMask       VkAccessFlags2
	DstStageMask        VkPipelineStageFlags2
	DstAccessMask       VkAccessFlags2
	OldLayout           VkImageLayout
	NewLayout           VkImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               VkImage
	SubresourceRange    VkImageSubresourceRange
}

type VkImageMemoryBarrier2KHR struct {
}

type VkBufferMemoryBarrier2 struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	SrcStageMask        VkPipelineStageFlags2
	SrcAccessMask       VkAccessFlags2
	DstStageMask        VkPipelineStageFlags2
	DstAccessMask       VkAccessFlags2
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              VkBuffer
	Offset              VkDeviceSize
	Size                VkDeviceSize
}

type VkBufferMemoryBarrier2KHR struct {
}

type VkMemoryBarrierAccessFlags3KHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcAccessMask3 VkAccessFlags3KHR
	DstAccessMask3 VkAccessFlags3KHR
}

type VkDependencyInfo struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DependencyFlags          VkDependencyFlags
	MemoryBarrierCount       uint32
	PMemoryBarriers          *VkMemoryBarrier2
	BufferMemoryBarrierCount uint32
	PBufferMemoryBarriers    *VkBufferMemoryBarrier2
	ImageMemoryBarrierCount  uint32
	PImageMemoryBarriers     *VkImageMemoryBarrier2
}

type VkDependencyInfoKHR struct {
}

type VkSemaphoreSubmitInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Semaphore   VkSemaphore
	Value       uint64
	StageMask   VkPipelineStageFlags2
	DeviceIndex uint32
}

type VkSemaphoreSubmitInfoKHR struct {
}

type VkCommandBufferSubmitInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	CommandBuffer VkCommandBuffer
	DeviceMask    uint32
}

type VkCommandBufferSubmitInfoKHR struct {
}

type VkSubmitInfo2 struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	Flags                    VkSubmitFlags
	WaitSemaphoreInfoCount   uint32
	PWaitSemaphoreInfos      *VkSemaphoreSubmitInfo
	CommandBufferInfoCount   uint32
	PCommandBufferInfos      *VkCommandBufferSubmitInfo
	SignalSemaphoreInfoCount uint32
	PSignalSemaphoreInfos    *VkSemaphoreSubmitInfo
}

type VkSubmitInfo2KHR struct {
}

type VkQueueFamilyCheckpointProperties2NV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	CheckpointExecutionStageMask VkPipelineStageFlags2
}

type VkCheckpointData2NV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Stage             VkPipelineStageFlags2
	PCheckpointMarker unsafe.Pointer
}

type VkPhysicalDeviceSynchronization2Features struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Synchronization2 VkBool32
}

type VkPhysicalDeviceSynchronization2FeaturesKHR struct {
}

type VkPhysicalDeviceUnifiedImageLayoutsFeaturesKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	UnifiedImageLayouts      VkBool32
	UnifiedImageLayoutsVideo VkBool32
}

type VkPhysicalDeviceHostImageCopyFeatures struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	HostImageCopy VkBool32
}

type VkPhysicalDeviceHostImageCopyFeaturesEXT struct {
}

type VkPhysicalDeviceHostImageCopyProperties struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	CopySrcLayoutCount              uint32
	PCopySrcLayouts                 *VkImageLayout
	CopyDstLayoutCount              uint32
	PCopyDstLayouts                 *VkImageLayout
	OptimalTilingLayoutUUID         uint8
	IdenticalMemoryTypeRequirements VkBool32
}

type VkPhysicalDeviceHostImageCopyPropertiesEXT struct {
}

type VkMemoryToImageCopy struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PHostPointer      unsafe.Pointer
	MemoryRowLength   uint32
	MemoryImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}

type VkMemoryToImageCopyEXT struct {
}

type VkImageToMemoryCopy struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PHostPointer      unsafe.Pointer
	MemoryRowLength   uint32
	MemoryImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}

type VkImageToMemoryCopyEXT struct {
}

type VkCopyMemoryToImageInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkHostImageCopyFlags
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkMemoryToImageCopy
}

type VkCopyMemoryToImageInfoEXT struct {
}

type VkCopyImageToMemoryInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkHostImageCopyFlags
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageToMemoryCopy
}

type VkCopyImageToMemoryInfoEXT struct {
}

type VkCopyImageToImageInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkHostImageCopyFlags
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	RegionCount    uint32
	PRegions       *VkImageCopy2
}

type VkCopyImageToImageInfoEXT struct {
}

type VkHostImageLayoutTransitionInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Image            VkImage
	OldLayout        VkImageLayout
	NewLayout        VkImageLayout
	SubresourceRange VkImageSubresourceRange
}

type VkHostImageLayoutTransitionInfoEXT struct {
}

type VkSubresourceHostMemcpySize struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Size  VkDeviceSize
}

type VkSubresourceHostMemcpySizeEXT struct {
}

type VkHostImageCopyDevicePerformanceQuery struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	OptimalDeviceAccess   VkBool32
	IdenticalMemoryLayout VkBool32
}

type VkHostImageCopyDevicePerformanceQueryEXT struct {
}

type VkPhysicalDeviceVulkanSC10Properties struct {
	SType                                            VkStructureType
	PNext                                            unsafe.Pointer
	DeviceNoDynamicHostAllocations                   VkBool32
	DeviceDestroyFreesMemory                         VkBool32
	CommandPoolMultipleCommandBuffersRecording       VkBool32
	CommandPoolResetCommandBuffer                    VkBool32
	CommandBufferSimultaneousUse                     VkBool32
	SecondaryCommandBufferNullOrImagelessFramebuffer VkBool32
	RecycleDescriptorSetMemory                       VkBool32
	RecyclePipelineMemory                            VkBool32
	MaxRenderPassSubpasses                           uint32
	MaxRenderPassDependencies                        uint32
	MaxSubpassInputAttachments                       uint32
	MaxSubpassPreserveAttachments                    uint32
	MaxFramebufferAttachments                        uint32
	MaxDescriptorSetLayoutBindings                   uint32
	MaxQueryFaultCount                               uint32
	MaxCallbackFaultCount                            uint32
	MaxCommandPoolCommandBuffers                     uint32
	MaxCommandBufferSize                             VkDeviceSize
}

type VkPipelinePoolSize struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PoolEntrySize  VkDeviceSize
	PoolEntryCount uint32
}

type VkDeviceObjectReservationCreateInfo struct {
	SType                                      VkStructureType
	PNext                                      unsafe.Pointer
	PipelineCacheCreateInfoCount               uint32
	PPipelineCacheCreateInfos                  *VkPipelineCacheCreateInfo
	PipelinePoolSizeCount                      uint32
	PPipelinePoolSizes                         *VkPipelinePoolSize
	SemaphoreRequestCount                      uint32
	CommandBufferRequestCount                  uint32
	FenceRequestCount                          uint32
	DeviceMemoryRequestCount                   uint32
	BufferRequestCount                         uint32
	ImageRequestCount                          uint32
	EventRequestCount                          uint32
	QueryPoolRequestCount                      uint32
	BufferViewRequestCount                     uint32
	ImageViewRequestCount                      uint32
	LayeredImageViewRequestCount               uint32
	PipelineCacheRequestCount                  uint32
	PipelineLayoutRequestCount                 uint32
	RenderPassRequestCount                     uint32
	GraphicsPipelineRequestCount               uint32
	ComputePipelineRequestCount                uint32
	DescriptorSetLayoutRequestCount            uint32
	SamplerRequestCount                        uint32
	DescriptorPoolRequestCount                 uint32
	DescriptorSetRequestCount                  uint32
	FramebufferRequestCount                    uint32
	CommandPoolRequestCount                    uint32
	SamplerYcbcrConversionRequestCount         uint32
	SurfaceRequestCount                        uint32
	SwapchainRequestCount                      uint32
	DisplayModeRequestCount                    uint32
	SubpassDescriptionRequestCount             uint32
	AttachmentDescriptionRequestCount          uint32
	DescriptorSetLayoutBindingRequestCount     uint32
	DescriptorSetLayoutBindingLimit            uint32
	MaxImageViewMipLevels                      uint32
	MaxImageViewArrayLayers                    uint32
	MaxLayeredImageViewMipLevels               uint32
	MaxOcclusionQueriesPerPool                 uint32
	MaxPipelineStatisticsQueriesPerPool        uint32
	MaxTimestampQueriesPerPool                 uint32
	MaxImmutableSamplersPerDescriptorSetLayout uint32
}

type VkCommandPoolMemoryReservationCreateInfo struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	CommandPoolReservedSize      VkDeviceSize
	CommandPoolMaxCommandBuffers uint32
}

type VkCommandPoolMemoryConsumption struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	CommandPoolAllocated    VkDeviceSize
	CommandPoolReservedSize VkDeviceSize
	CommandBufferAllocated  VkDeviceSize
}

type VkPhysicalDeviceVulkanSC10Features struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ShaderAtomicInstructions VkBool32
}

type VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT struct {
	SType                                         VkStructureType
	PNext                                         unsafe.Pointer
	PrimitivesGeneratedQuery                      VkBool32
	PrimitivesGeneratedQueryWithRasterizerDiscard VkBool32
	PrimitivesGeneratedQueryWithNonZeroStreams    VkBool32
}

type VkPhysicalDeviceLegacyDitheringFeaturesEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	LegacyDithering VkBool32
}

type VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	MultisampledRenderToSingleSampled VkBool32
}

type VkSurfaceCapabilitiesPresentId2KHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	PresentId2Supported VkBool32
}

type VkSurfaceCapabilitiesPresentWait2KHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	PresentWait2Supported VkBool32
}

type VkSubpassResolvePerformanceQueryEXT struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Optimal VkBool32
}

type VkMultisampledRenderToSingleSampledInfoEXT struct {
	SType                                   VkStructureType
	PNext                                   unsafe.Pointer
	MultisampledRenderToSingleSampledEnable VkBool32
	RasterizationSamples                    VkSampleCountFlagBits
}

type VkPhysicalDevicePipelineProtectedAccessFeatures struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	PipelineProtectedAccess VkBool32
}

type VkPhysicalDevicePipelineProtectedAccessFeaturesEXT struct {
}

type VkQueueFamilyVideoPropertiesKHR struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	VideoCodecOperations VkVideoCodecOperationFlagsKHR
}

type VkQueueFamilyQueryResultStatusPropertiesKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	QueryResultStatusSupport VkBool32
}

type VkVideoProfileListInfoKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	ProfileCount uint32
	PProfiles    *VkVideoProfileInfoKHR
}

type VkPhysicalDeviceVideoFormatInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	ImageUsage VkImageUsageFlags
}

type VkVideoFormatPropertiesKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Format           VkFormat
	ComponentMapping VkComponentMapping
	ImageCreateFlags VkImageCreateFlags
	ImageType        VkImageType
	ImageTiling      VkImageTiling
	ImageUsageFlags  VkImageUsageFlags
}

type VkVideoEncodeQuantizationMapCapabilitiesKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	MaxQuantizationMapExtent VkExtent2D
}

type VkVideoEncodeH264QuantizationMapCapabilitiesKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	MinQpDelta int32
	MaxQpDelta int32
}

type VkVideoEncodeH265QuantizationMapCapabilitiesKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	MinQpDelta int32
	MaxQpDelta int32
}

type VkVideoEncodeAV1QuantizationMapCapabilitiesKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MinQIndexDelta int32
	MaxQIndexDelta int32
}

type VkVideoFormatQuantizationMapPropertiesKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	QuantizationMapTexelSize VkExtent2D
}

type VkVideoFormatH265QuantizationMapPropertiesKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	CompatibleCtbSizes VkVideoEncodeH265CtbSizeFlagsKHR
}

type VkVideoFormatAV1QuantizationMapPropertiesKHR struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	CompatibleSuperblockSizes VkVideoEncodeAV1SuperblockSizeFlagsKHR
}

type VkVideoProfileInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	VideoCodecOperation VkVideoCodecOperationFlagBitsKHR
	ChromaSubsampling   VkVideoChromaSubsamplingFlagsKHR
	LumaBitDepth        VkVideoComponentBitDepthFlagsKHR
	ChromaBitDepth      VkVideoComponentBitDepthFlagsKHR
}

type VkVideoCapabilitiesKHR struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	Flags                             VkVideoCapabilityFlagsKHR
	MinBitstreamBufferOffsetAlignment VkDeviceSize
	MinBitstreamBufferSizeAlignment   VkDeviceSize
	PictureAccessGranularity          VkExtent2D
	MinCodedExtent                    VkExtent2D
	MaxCodedExtent                    VkExtent2D
	MaxDpbSlots                       uint32
	MaxActiveReferencePictures        uint32
	StdHeaderVersion                  VkExtensionProperties
}

type VkVideoSessionMemoryRequirementsKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MemoryBindIndex    uint32
	MemoryRequirements VkMemoryRequirements
}

type VkBindVideoSessionMemoryInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	MemoryBindIndex uint32
	Memory          VkDeviceMemory
	MemoryOffset    VkDeviceSize
	MemorySize      VkDeviceSize
}

type VkVideoPictureResourceInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	CodedOffset      VkOffset2D
	CodedExtent      VkExtent2D
	BaseArrayLayer   uint32
	ImageViewBinding VkImageView
}

type VkVideoReferenceSlotInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	SlotIndex        int32
	PPictureResource *VkVideoPictureResourceInfoKHR
}

type VkVideoDecodeCapabilitiesKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkVideoDecodeCapabilityFlagsKHR
}

type VkVideoDecodeUsageInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	VideoUsageHints VkVideoDecodeUsageFlagsKHR
}

type VkVideoDecodeInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkVideoDecodeFlagsKHR
	SrcBuffer           VkBuffer
	SrcBufferOffset     VkDeviceSize
	SrcBufferRange      VkDeviceSize
	DstPictureResource  VkVideoPictureResourceInfoKHR
	PSetupReferenceSlot *VkVideoReferenceSlotInfoKHR
	ReferenceSlotCount  uint32
	PReferenceSlots     *VkVideoReferenceSlotInfoKHR
}

type VkPhysicalDeviceVideoMaintenance1FeaturesKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	VideoMaintenance1 VkBool32
}

type VkPhysicalDeviceVideoMaintenance2FeaturesKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	VideoMaintenance2 VkBool32
}

type VkVideoInlineQueryInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	QueryPool  VkQueryPool
	FirstQuery uint32
	QueryCount uint32
}

type VkVideoDecodeH264ProfileInfoKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	StdProfileIdc StdVideoH264ProfileIdc
	PictureLayout VkVideoDecodeH264PictureLayoutFlagBitsKHR
}

type VkVideoDecodeH264CapabilitiesKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	MaxLevelIdc            StdVideoH264LevelIdc
	FieldOffsetGranularity VkOffset2D
}

type VkVideoDecodeH264SessionParametersAddInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	StdSPSCount uint32
	PStdSPSs    *StdVideoH264SequenceParameterSet
	StdPPSCount uint32
	PStdPPSs    *StdVideoH264PictureParameterSet
}

type VkVideoDecodeH264SessionParametersCreateInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxStdSPSCount     uint32
	MaxStdPPSCount     uint32
	PParametersAddInfo *VkVideoDecodeH264SessionParametersAddInfoKHR
}

type VkVideoDecodeH264InlineSessionParametersInfoKHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	PStdSPS *StdVideoH264SequenceParameterSet
	PStdPPS *StdVideoH264PictureParameterSet
}

type VkVideoDecodeH264PictureInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	PStdPictureInfo *StdVideoDecodeH264PictureInfo
	SliceCount      uint32
	PSliceOffsets   *uint32
}

type VkVideoDecodeH264DpbSlotInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PStdReferenceInfo *StdVideoDecodeH264ReferenceInfo
}

type VkVideoDecodeH265ProfileInfoKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	StdProfileIdc StdVideoH265ProfileIdc
}

type VkVideoDecodeH265CapabilitiesKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	MaxLevelIdc StdVideoH265LevelIdc
}

type VkVideoDecodeH265SessionParametersAddInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	StdVPSCount uint32
	PStdVPSs    *StdVideoH265VideoParameterSet
	StdSPSCount uint32
	PStdSPSs    *StdVideoH265SequenceParameterSet
	StdPPSCount uint32
	PStdPPSs    *StdVideoH265PictureParameterSet
}

type VkVideoDecodeH265SessionParametersCreateInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxStdVPSCount     uint32
	MaxStdSPSCount     uint32
	MaxStdPPSCount     uint32
	PParametersAddInfo *VkVideoDecodeH265SessionParametersAddInfoKHR
}

type VkVideoDecodeH265InlineSessionParametersInfoKHR struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	PStdVPS *StdVideoH265VideoParameterSet
	PStdSPS *StdVideoH265SequenceParameterSet
	PStdPPS *StdVideoH265PictureParameterSet
}

type VkVideoDecodeH265PictureInfoKHR struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	PStdPictureInfo      *StdVideoDecodeH265PictureInfo
	SliceSegmentCount    uint32
	PSliceSegmentOffsets *uint32
}

type VkVideoDecodeH265DpbSlotInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PStdReferenceInfo *StdVideoDecodeH265ReferenceInfo
}

type VkPhysicalDeviceVideoDecodeVP9FeaturesKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	VideoDecodeVP9 VkBool32
}

type VkVideoDecodeVP9ProfileInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	StdProfile StdVideoVP9Profile
}

type VkVideoDecodeVP9CapabilitiesKHR struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	MaxLevel StdVideoVP9Level
}

type VkVideoDecodeVP9PictureInfoKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	PStdPictureInfo          *StdVideoDecodeVP9PictureInfo
	ReferenceNameSlotIndices int32
	UncompressedHeaderOffset uint32
	CompressedHeaderOffset   uint32
	TilesOffset              uint32
}

type VkVideoDecodeAV1ProfileInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	StdProfile       StdVideoAV1Profile
	FilmGrainSupport VkBool32
}

type VkVideoDecodeAV1CapabilitiesKHR struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	MaxLevel StdVideoAV1Level
}

type VkVideoDecodeAV1SessionParametersCreateInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PStdSequenceHeader *StdVideoAV1SequenceHeader
}

type VkVideoDecodeAV1InlineSessionParametersInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PStdSequenceHeader *StdVideoAV1SequenceHeader
}

type VkVideoDecodeAV1PictureInfoKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	PStdPictureInfo          *StdVideoDecodeAV1PictureInfo
	ReferenceNameSlotIndices int32
	FrameHeaderOffset        uint32
	TileCount                uint32
	PTileOffsets             *uint32
	PTileSizes               *uint32
}

type VkVideoDecodeAV1DpbSlotInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PStdReferenceInfo *StdVideoDecodeAV1ReferenceInfo
}

type VkVideoSessionCreateInfoKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	QueueFamilyIndex           uint32
	Flags                      VkVideoSessionCreateFlagsKHR
	PVideoProfile              *VkVideoProfileInfoKHR
	PictureFormat              VkFormat
	MaxCodedExtent             VkExtent2D
	ReferencePictureFormat     VkFormat
	MaxDpbSlots                uint32
	MaxActiveReferencePictures uint32
	PStdHeaderVersion          *VkExtensionProperties
}

type VkVideoSessionParametersCreateInfoKHR struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	Flags                          VkVideoSessionParametersCreateFlagsKHR
	VideoSessionParametersTemplate VkVideoSessionParametersKHR
	VideoSession                   VkVideoSessionKHR
}

type VkVideoSessionParametersUpdateInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	UpdateSequenceCount uint32
}

type VkVideoEncodeSessionParametersGetInfoKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	VideoSessionParameters VkVideoSessionParametersKHR
}

type VkVideoEncodeSessionParametersFeedbackInfoKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	HasOverrides VkBool32
}

type VkVideoBeginCodingInfoKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkVideoBeginCodingFlagsKHR
	VideoSession           VkVideoSessionKHR
	VideoSessionParameters VkVideoSessionParametersKHR
	ReferenceSlotCount     uint32
	PReferenceSlots        *VkVideoReferenceSlotInfoKHR
}

type VkVideoEndCodingInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkVideoEndCodingFlagsKHR
}

type VkVideoCodingControlInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkVideoCodingControlFlagsKHR
}

type VkVideoEncodeUsageInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	VideoUsageHints   VkVideoEncodeUsageFlagsKHR
	VideoContentHints VkVideoEncodeContentFlagsKHR
	TuningMode        VkVideoEncodeTuningModeKHR
}

type VkVideoEncodeInfoKHR struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	Flags                           VkVideoEncodeFlagsKHR
	DstBuffer                       VkBuffer
	DstBufferOffset                 VkDeviceSize
	DstBufferRange                  VkDeviceSize
	SrcPictureResource              VkVideoPictureResourceInfoKHR
	PSetupReferenceSlot             *VkVideoReferenceSlotInfoKHR
	ReferenceSlotCount              uint32
	PReferenceSlots                 *VkVideoReferenceSlotInfoKHR
	PrecedingExternallyEncodedBytes uint32
}

type VkVideoEncodeQuantizationMapInfoKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	QuantizationMap       VkImageView
	QuantizationMapExtent VkExtent2D
}

type VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	QuantizationMapTexelSize VkExtent2D
}

type VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	VideoEncodeQuantizationMap VkBool32
}

type VkQueryPoolVideoEncodeFeedbackCreateInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	EncodeFeedbackFlags VkVideoEncodeFeedbackFlagsKHR
}

type VkVideoEncodeQualityLevelInfoKHR struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	QualityLevel uint32
}

type VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	PVideoProfile *VkVideoProfileInfoKHR
	QualityLevel  uint32
}

type VkVideoEncodeQualityLevelPropertiesKHR struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	PreferredRateControlMode       VkVideoEncodeRateControlModeFlagBitsKHR
	PreferredRateControlLayerCount uint32
}

type VkVideoEncodeRateControlInfoKHR struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	Flags                        VkVideoEncodeRateControlFlagsKHR
	RateControlMode              VkVideoEncodeRateControlModeFlagBitsKHR
	LayerCount                   uint32
	PLayers                      *VkVideoEncodeRateControlLayerInfoKHR
	VirtualBufferSizeInMs        uint32
	InitialVirtualBufferSizeInMs uint32
}

type VkVideoEncodeRateControlLayerInfoKHR struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	AverageBitrate       uint64
	MaxBitrate           uint64
	FrameRateNumerator   uint32
	FrameRateDenominator uint32
}

type VkVideoEncodeCapabilitiesKHR struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	Flags                         VkVideoEncodeCapabilityFlagsKHR
	RateControlModes              VkVideoEncodeRateControlModeFlagsKHR
	MaxRateControlLayers          uint32
	MaxBitrate                    uint64
	MaxQualityLevels              uint32
	EncodeInputPictureGranularity VkExtent2D
	SupportedEncodeFeedbackFlags  VkVideoEncodeFeedbackFlagsKHR
}

type VkVideoEncodeH264CapabilitiesKHR struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	Flags                            VkVideoEncodeH264CapabilityFlagsKHR
	MaxLevelIdc                      StdVideoH264LevelIdc
	MaxSliceCount                    uint32
	MaxPPictureL0ReferenceCount      uint32
	MaxBPictureL0ReferenceCount      uint32
	MaxL1ReferenceCount              uint32
	MaxTemporalLayerCount            uint32
	ExpectDyadicTemporalLayerPattern VkBool32
	MinQp                            int32
	MaxQp                            int32
	PrefersGopRemainingFrames        VkBool32
	RequiresGopRemainingFrames       VkBool32
	StdSyntaxFlags                   VkVideoEncodeH264StdFlagsKHR
}

type VkVideoEncodeH264QualityLevelPropertiesKHR struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	PreferredRateControlFlags         VkVideoEncodeH264RateControlFlagsKHR
	PreferredGopFrameCount            uint32
	PreferredIdrPeriod                uint32
	PreferredConsecutiveBFrameCount   uint32
	PreferredTemporalLayerCount       uint32
	PreferredConstantQp               VkVideoEncodeH264QpKHR
	PreferredMaxL0ReferenceCount      uint32
	PreferredMaxL1ReferenceCount      uint32
	PreferredStdEntropyCodingModeFlag VkBool32
}

type VkVideoEncodeH264SessionCreateInfoKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	UseMaxLevelIdc VkBool32
	MaxLevelIdc    StdVideoH264LevelIdc
}

type VkVideoEncodeH264SessionParametersAddInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	StdSPSCount uint32
	PStdSPSs    *StdVideoH264SequenceParameterSet
	StdPPSCount uint32
	PStdPPSs    *StdVideoH264PictureParameterSet
}

type VkVideoEncodeH264SessionParametersCreateInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxStdSPSCount     uint32
	MaxStdPPSCount     uint32
	PParametersAddInfo *VkVideoEncodeH264SessionParametersAddInfoKHR
}

type VkVideoEncodeH264SessionParametersGetInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	WriteStdSPS VkBool32
	WriteStdPPS VkBool32
	StdSPSId    uint32
	StdPPSId    uint32
}

type VkVideoEncodeH264SessionParametersFeedbackInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	HasStdSPSOverrides VkBool32
	HasStdPPSOverrides VkBool32
}

type VkVideoEncodeH264DpbSlotInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PStdReferenceInfo *StdVideoEncodeH264ReferenceInfo
}

type VkVideoEncodeH264PictureInfoKHR struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	NaluSliceEntryCount uint32
	PNaluSliceEntries   *VkVideoEncodeH264NaluSliceInfoKHR
	PStdPictureInfo     *StdVideoEncodeH264PictureInfo
	GeneratePrefixNalu  VkBool32
}

type VkVideoEncodeH264ProfileInfoKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	StdProfileIdc StdVideoH264ProfileIdc
}

type VkVideoEncodeH264NaluSliceInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ConstantQp      int32
	PStdSliceHeader *StdVideoEncodeH264SliceHeader
}

type VkVideoEncodeH264RateControlInfoKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkVideoEncodeH264RateControlFlagsKHR
	GopFrameCount          uint32
	IdrPeriod              uint32
	ConsecutiveBFrameCount uint32
	TemporalLayerCount     uint32
}

type VkVideoEncodeH264QpKHR struct {
	QpI int32
	QpP int32
	QpB int32
}

type VkVideoEncodeH264FrameSizeKHR struct {
	FrameISize uint32
	FramePSize uint32
	FrameBSize uint32
}

type VkVideoEncodeH264GopRemainingFrameInfoKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	UseGopRemainingFrames VkBool32
	GopRemainingI         uint32
	GopRemainingP         uint32
	GopRemainingB         uint32
}

type VkVideoEncodeH264RateControlLayerInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	UseMinQp        VkBool32
	MinQp           VkVideoEncodeH264QpKHR
	UseMaxQp        VkBool32
	MaxQp           VkVideoEncodeH264QpKHR
	UseMaxFrameSize VkBool32
	MaxFrameSize    VkVideoEncodeH264FrameSizeKHR
}

type VkVideoEncodeH265CapabilitiesKHR struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	Flags                               VkVideoEncodeH265CapabilityFlagsKHR
	MaxLevelIdc                         StdVideoH265LevelIdc
	MaxSliceSegmentCount                uint32
	MaxTiles                            VkExtent2D
	CtbSizes                            VkVideoEncodeH265CtbSizeFlagsKHR
	TransformBlockSizes                 VkVideoEncodeH265TransformBlockSizeFlagsKHR
	MaxPPictureL0ReferenceCount         uint32
	MaxBPictureL0ReferenceCount         uint32
	MaxL1ReferenceCount                 uint32
	MaxSubLayerCount                    uint32
	ExpectDyadicTemporalSubLayerPattern VkBool32
	MinQp                               int32
	MaxQp                               int32
	PrefersGopRemainingFrames           VkBool32
	RequiresGopRemainingFrames          VkBool32
	StdSyntaxFlags                      VkVideoEncodeH265StdFlagsKHR
}

type VkVideoEncodeH265QualityLevelPropertiesKHR struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	PreferredRateControlFlags       VkVideoEncodeH265RateControlFlagsKHR
	PreferredGopFrameCount          uint32
	PreferredIdrPeriod              uint32
	PreferredConsecutiveBFrameCount uint32
	PreferredSubLayerCount          uint32
	PreferredConstantQp             VkVideoEncodeH265QpKHR
	PreferredMaxL0ReferenceCount    uint32
	PreferredMaxL1ReferenceCount    uint32
}

type VkVideoEncodeH265SessionCreateInfoKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	UseMaxLevelIdc VkBool32
	MaxLevelIdc    StdVideoH265LevelIdc
}

type VkVideoEncodeH265SessionParametersAddInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	StdVPSCount uint32
	PStdVPSs    *StdVideoH265VideoParameterSet
	StdSPSCount uint32
	PStdSPSs    *StdVideoH265SequenceParameterSet
	StdPPSCount uint32
	PStdPPSs    *StdVideoH265PictureParameterSet
}

type VkVideoEncodeH265SessionParametersCreateInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxStdVPSCount     uint32
	MaxStdSPSCount     uint32
	MaxStdPPSCount     uint32
	PParametersAddInfo *VkVideoEncodeH265SessionParametersAddInfoKHR
}

type VkVideoEncodeH265SessionParametersGetInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	WriteStdVPS VkBool32
	WriteStdSPS VkBool32
	WriteStdPPS VkBool32
	StdVPSId    uint32
	StdSPSId    uint32
	StdPPSId    uint32
}

type VkVideoEncodeH265SessionParametersFeedbackInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	HasStdVPSOverrides VkBool32
	HasStdSPSOverrides VkBool32
	HasStdPPSOverrides VkBool32
}

type VkVideoEncodeH265PictureInfoKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	NaluSliceSegmentEntryCount uint32
	PNaluSliceSegmentEntries   *VkVideoEncodeH265NaluSliceSegmentInfoKHR
	PStdPictureInfo            *StdVideoEncodeH265PictureInfo
}

type VkVideoEncodeH265NaluSliceSegmentInfoKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ConstantQp             int32
	PStdSliceSegmentHeader *StdVideoEncodeH265SliceSegmentHeader
}

type VkVideoEncodeH265RateControlInfoKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkVideoEncodeH265RateControlFlagsKHR
	GopFrameCount          uint32
	IdrPeriod              uint32
	ConsecutiveBFrameCount uint32
	SubLayerCount          uint32
}

type VkVideoEncodeH265QpKHR struct {
	QpI int32
	QpP int32
	QpB int32
}

type VkVideoEncodeH265FrameSizeKHR struct {
	FrameISize uint32
	FramePSize uint32
	FrameBSize uint32
}

type VkVideoEncodeH265GopRemainingFrameInfoKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	UseGopRemainingFrames VkBool32
	GopRemainingI         uint32
	GopRemainingP         uint32
	GopRemainingB         uint32
}

type VkVideoEncodeH265RateControlLayerInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	UseMinQp        VkBool32
	MinQp           VkVideoEncodeH265QpKHR
	UseMaxQp        VkBool32
	MaxQp           VkVideoEncodeH265QpKHR
	UseMaxFrameSize VkBool32
	MaxFrameSize    VkVideoEncodeH265FrameSizeKHR
}

type VkVideoEncodeH265ProfileInfoKHR struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	StdProfileIdc StdVideoH265ProfileIdc
}

type VkVideoEncodeH265DpbSlotInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PStdReferenceInfo *StdVideoEncodeH265ReferenceInfo
}

type VkVideoEncodeAV1CapabilitiesKHR struct {
	SType                                         VkStructureType
	PNext                                         unsafe.Pointer
	Flags                                         VkVideoEncodeAV1CapabilityFlagsKHR
	MaxLevel                                      StdVideoAV1Level
	CodedPictureAlignment                         VkExtent2D
	MaxTiles                                      VkExtent2D
	MinTileSize                                   VkExtent2D
	MaxTileSize                                   VkExtent2D
	SuperblockSizes                               VkVideoEncodeAV1SuperblockSizeFlagsKHR
	MaxSingleReferenceCount                       uint32
	SingleReferenceNameMask                       uint32
	MaxUnidirectionalCompoundReferenceCount       uint32
	MaxUnidirectionalCompoundGroup1ReferenceCount uint32
	UnidirectionalCompoundReferenceNameMask       uint32
	MaxBidirectionalCompoundReferenceCount        uint32
	MaxBidirectionalCompoundGroup1ReferenceCount  uint32
	MaxBidirectionalCompoundGroup2ReferenceCount  uint32
	BidirectionalCompoundReferenceNameMask        uint32
	MaxTemporalLayerCount                         uint32
	MaxSpatialLayerCount                          uint32
	MaxOperatingPoints                            uint32
	MinQIndex                                     uint32
	MaxQIndex                                     uint32
	PrefersGopRemainingFrames                     VkBool32
	RequiresGopRemainingFrames                    VkBool32
	StdSyntaxFlags                                VkVideoEncodeAV1StdFlagsKHR
}

type VkVideoEncodeAV1QualityLevelPropertiesKHR struct {
	SType                                                  VkStructureType
	PNext                                                  unsafe.Pointer
	PreferredRateControlFlags                              VkVideoEncodeAV1RateControlFlagsKHR
	PreferredGopFrameCount                                 uint32
	PreferredKeyFramePeriod                                uint32
	PreferredConsecutiveBipredictiveFrameCount             uint32
	PreferredTemporalLayerCount                            uint32
	PreferredConstantQIndex                                VkVideoEncodeAV1QIndexKHR
	PreferredMaxSingleReferenceCount                       uint32
	PreferredSingleReferenceNameMask                       uint32
	PreferredMaxUnidirectionalCompoundReferenceCount       uint32
	PreferredMaxUnidirectionalCompoundGroup1ReferenceCount uint32
	PreferredUnidirectionalCompoundReferenceNameMask       uint32
	PreferredMaxBidirectionalCompoundReferenceCount        uint32
	PreferredMaxBidirectionalCompoundGroup1ReferenceCount  uint32
	PreferredMaxBidirectionalCompoundGroup2ReferenceCount  uint32
	PreferredBidirectionalCompoundReferenceNameMask        uint32
}

type VkPhysicalDeviceVideoEncodeAV1FeaturesKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	VideoEncodeAV1 VkBool32
}

type VkVideoEncodeAV1SessionCreateInfoKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	UseMaxLevel VkBool32
	MaxLevel    StdVideoAV1Level
}

type VkVideoEncodeAV1SessionParametersCreateInfoKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	PStdSequenceHeader     *StdVideoAV1SequenceHeader
	PStdDecoderModelInfo   *StdVideoEncodeAV1DecoderModelInfo
	StdOperatingPointCount uint32
	PStdOperatingPoints    *StdVideoEncodeAV1OperatingPointInfo
}

type VkVideoEncodeAV1DpbSlotInfoKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	PStdReferenceInfo *StdVideoEncodeAV1ReferenceInfo
}

type VkVideoEncodeAV1PictureInfoKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	PredictionMode             VkVideoEncodeAV1PredictionModeKHR
	RateControlGroup           VkVideoEncodeAV1RateControlGroupKHR
	ConstantQIndex             uint32
	PStdPictureInfo            *StdVideoEncodeAV1PictureInfo
	ReferenceNameSlotIndices   int32
	PrimaryReferenceCdfOnly    VkBool32
	GenerateObuExtensionHeader VkBool32
}

type VkVideoEncodeAV1ProfileInfoKHR struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	StdProfile StdVideoAV1Profile
}

type VkVideoEncodeAV1RateControlInfoKHR struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	Flags                             VkVideoEncodeAV1RateControlFlagsKHR
	GopFrameCount                     uint32
	KeyFramePeriod                    uint32
	ConsecutiveBipredictiveFrameCount uint32
	TemporalLayerCount                uint32
}

type VkVideoEncodeAV1QIndexKHR struct {
	IntraQIndex        uint32
	PredictiveQIndex   uint32
	BipredictiveQIndex uint32
}

type VkVideoEncodeAV1FrameSizeKHR struct {
	IntraFrameSize        uint32
	PredictiveFrameSize   uint32
	BipredictiveFrameSize uint32
}

type VkVideoEncodeAV1GopRemainingFrameInfoKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	UseGopRemainingFrames    VkBool32
	GopRemainingIntra        uint32
	GopRemainingPredictive   uint32
	GopRemainingBipredictive uint32
}

type VkVideoEncodeAV1RateControlLayerInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	UseMinQIndex    VkBool32
	MinQIndex       VkVideoEncodeAV1QIndexKHR
	UseMaxQIndex    VkBool32
	MaxQIndex       VkVideoEncodeAV1QIndexKHR
	UseMaxFrameSize VkBool32
	MaxFrameSize    VkVideoEncodeAV1FrameSizeKHR
}

type VkPhysicalDeviceInheritedViewportScissorFeaturesNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	InheritedViewportScissor2D VkBool32
}

type VkCommandBufferInheritanceViewportScissorInfoNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ViewportScissor2D  VkBool32
	ViewportDepthCount uint32
	PViewportDepths    *VkViewport
}

type VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Ycbcr2plane444Formats VkBool32
}

type VkPhysicalDeviceProvokingVertexFeaturesEXT struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	ProvokingVertexLast                       VkBool32
	TransformFeedbackPreservesProvokingVertex VkBool32
}

type VkPhysicalDeviceProvokingVertexPropertiesEXT struct {
	SType                                                VkStructureType
	PNext                                                unsafe.Pointer
	ProvokingVertexModePerPipeline                       VkBool32
	TransformFeedbackPreservesTriangleFanProvokingVertex VkBool32
}

type VkPipelineRasterizationProvokingVertexStateCreateInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	ProvokingVertexMode VkProvokingVertexModeEXT
}

type VkVideoEncodeIntraRefreshCapabilitiesKHR struct {
	SType                                   VkStructureType
	PNext                                   unsafe.Pointer
	IntraRefreshModes                       VkVideoEncodeIntraRefreshModeFlagsKHR
	MaxIntraRefreshCycleDuration            uint32
	MaxIntraRefreshActiveReferencePictures  uint32
	PartitionIndependentIntraRefreshRegions VkBool32
	NonRectangularIntraRefreshRegions       VkBool32
}

type VkVideoEncodeSessionIntraRefreshCreateInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	IntraRefreshMode VkVideoEncodeIntraRefreshModeFlagBitsKHR
}

type VkVideoEncodeIntraRefreshInfoKHR struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	IntraRefreshCycleDuration uint32
	IntraRefreshIndex         uint32
}

type VkVideoReferenceIntraRefreshInfoKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DirtyIntraRefreshRegions uint32
}

type VkPhysicalDeviceVideoEncodeIntraRefreshFeaturesKHR struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	VideoEncodeIntraRefresh VkBool32
}

type VkCuModuleCreateInfoNVX struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	DataSize uint
	PData    unsafe.Pointer
}

type VkCuModuleTexturingModeCreateInfoNVX struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Use64bitTexturing VkBool32
}

type VkCuFunctionCreateInfoNVX struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Module VkCuModuleNVX
	PName  string
}

type VkCuLaunchInfoNVX struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Function       VkCuFunctionNVX
	GridDimX       uint32
	GridDimY       uint32
	GridDimZ       uint32
	BlockDimX      uint32
	BlockDimY      uint32
	BlockDimZ      uint32
	SharedMemBytes uint32
	ParamCount     uint
	PParams        **void
	ExtraCount     uint
	PExtras        **void
}

type VkPhysicalDeviceDescriptorBufferFeaturesEXT struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	DescriptorBuffer                   VkBool32
	DescriptorBufferCaptureReplay      VkBool32
	DescriptorBufferImageLayoutIgnored VkBool32
	DescriptorBufferPushDescriptors    VkBool32
}

type VkPhysicalDeviceDescriptorBufferPropertiesEXT struct {
	SType                                                VkStructureType
	PNext                                                unsafe.Pointer
	CombinedImageSamplerDescriptorSingleArray            VkBool32
	BufferlessPushDescriptors                            VkBool32
	AllowSamplerImageViewPostSubmitCreation              VkBool32
	DescriptorBufferOffsetAlignment                      VkDeviceSize
	MaxDescriptorBufferBindings                          uint32
	MaxResourceDescriptorBufferBindings                  uint32
	MaxSamplerDescriptorBufferBindings                   uint32
	MaxEmbeddedImmutableSamplerBindings                  uint32
	MaxEmbeddedImmutableSamplers                         uint32
	BufferCaptureReplayDescriptorDataSize                uint
	ImageCaptureReplayDescriptorDataSize                 uint
	ImageViewCaptureReplayDescriptorDataSize             uint
	SamplerCaptureReplayDescriptorDataSize               uint
	AccelerationStructureCaptureReplayDescriptorDataSize uint
	SamplerDescriptorSize                                uint
	CombinedImageSamplerDescriptorSize                   uint
	SampledImageDescriptorSize                           uint
	StorageImageDescriptorSize                           uint
	UniformTexelBufferDescriptorSize                     uint
	RobustUniformTexelBufferDescriptorSize               uint
	StorageTexelBufferDescriptorSize                     uint
	RobustStorageTexelBufferDescriptorSize               uint
	UniformBufferDescriptorSize                          uint
	RobustUniformBufferDescriptorSize                    uint
	StorageBufferDescriptorSize                          uint
	RobustStorageBufferDescriptorSize                    uint
	InputAttachmentDescriptorSize                        uint
	AccelerationStructureDescriptorSize                  uint
	MaxSamplerDescriptorBufferRange                      VkDeviceSize
	MaxResourceDescriptorBufferRange                     VkDeviceSize
	SamplerDescriptorBufferAddressSpaceSize              VkDeviceSize
	ResourceDescriptorBufferAddressSpaceSize             VkDeviceSize
	DescriptorBufferAddressSpaceSize                     VkDeviceSize
}

type VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT struct {
	SType                                        VkStructureType
	PNext                                        unsafe.Pointer
	CombinedImageSamplerDensityMapDescriptorSize uint
}

type VkDescriptorAddressInfoEXT struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Address VkDeviceAddress
	Range   VkDeviceSize
	Format  VkFormat
}

type VkDescriptorBufferBindingInfoEXT struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Address VkDeviceAddress
	Usage   VkBufferUsageFlags
}

type VkDescriptorBufferBindingPushDescriptorBufferHandleEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
}

type VkDescriptorGetInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Type  VkDescriptorType
	Data  VkDescriptorDataEXT
}

type VkBufferCaptureDescriptorDataInfoEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
}

type VkImageCaptureDescriptorDataInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Image VkImage
}

type VkImageViewCaptureDescriptorDataInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	ImageView VkImageView
}

type VkSamplerCaptureDescriptorDataInfoEXT struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Sampler VkSampler
}

type VkAccelerationStructureCaptureDescriptorDataInfoEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	AccelerationStructure   VkAccelerationStructureKHR
	AccelerationStructureNV VkAccelerationStructureNV
}

type VkOpaqueCaptureDescriptorDataCreateInfoEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	OpaqueCaptureDescriptorData unsafe.Pointer
}

type VkPhysicalDeviceShaderIntegerDotProductFeatures struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ShaderIntegerDotProduct VkBool32
}

type VkPhysicalDeviceShaderIntegerDotProductFeaturesKHR struct {
}

type VkPhysicalDeviceShaderIntegerDotProductProperties struct {
	SType                                                                         VkStructureType
	PNext                                                                         unsafe.Pointer
	IntegerDotProduct8BitUnsignedAccelerated                                      VkBool32
	IntegerDotProduct8BitSignedAccelerated                                        VkBool32
	IntegerDotProduct8BitMixedSignednessAccelerated                               VkBool32
	IntegerDotProduct4x8BitPackedUnsignedAccelerated                              VkBool32
	IntegerDotProduct4x8BitPackedSignedAccelerated                                VkBool32
	IntegerDotProduct4x8BitPackedMixedSignednessAccelerated                       VkBool32
	IntegerDotProduct16BitUnsignedAccelerated                                     VkBool32
	IntegerDotProduct16BitSignedAccelerated                                       VkBool32
	IntegerDotProduct16BitMixedSignednessAccelerated                              VkBool32
	IntegerDotProduct32BitUnsignedAccelerated                                     VkBool32
	IntegerDotProduct32BitSignedAccelerated                                       VkBool32
	IntegerDotProduct32BitMixedSignednessAccelerated                              VkBool32
	IntegerDotProduct64BitUnsignedAccelerated                                     VkBool32
	IntegerDotProduct64BitSignedAccelerated                                       VkBool32
	IntegerDotProduct64BitMixedSignednessAccelerated                              VkBool32
	IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated                VkBool32
	IntegerDotProductAccumulatingSaturating8BitSignedAccelerated                  VkBool32
	IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated         VkBool32
	IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated        VkBool32
	IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated          VkBool32
	IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated VkBool32
	IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated               VkBool32
	IntegerDotProductAccumulatingSaturating16BitSignedAccelerated                 VkBool32
	IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated        VkBool32
	IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated               VkBool32
	IntegerDotProductAccumulatingSaturating32BitSignedAccelerated                 VkBool32
	IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated        VkBool32
	IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated               VkBool32
	IntegerDotProductAccumulatingSaturating64BitSignedAccelerated                 VkBool32
	IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated        VkBool32
}

type VkPhysicalDeviceShaderIntegerDotProductPropertiesKHR struct {
}

type VkPhysicalDeviceDrmPropertiesEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	HasPrimary   VkBool32
	HasRender    VkBool32
	PrimaryMajor int64
	PrimaryMinor int64
	RenderMajor  int64
	RenderMinor  int64
}

type VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	FragmentShaderBarycentric VkBool32
}

type VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR struct {
	SType                                           VkStructureType
	PNext                                           unsafe.Pointer
	TriStripVertexOrderIndependentOfProvokingVertex VkBool32
}

type VkPhysicalDeviceShaderFmaFeaturesKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ShaderFmaFloat16 VkBool32
	ShaderFmaFloat32 VkBool32
	ShaderFmaFloat64 VkBool32
}

type VkPhysicalDeviceRayTracingMotionBlurFeaturesNV struct {
	SType                                         VkStructureType
	PNext                                         unsafe.Pointer
	RayTracingMotionBlur                          VkBool32
	RayTracingMotionBlurPipelineTraceRaysIndirect VkBool32
}

type VkPhysicalDeviceRayTracingValidationFeaturesNV struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	RayTracingValidation VkBool32
}

type VkPhysicalDeviceRayTracingLinearSweptSpheresFeaturesNV struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Spheres            VkBool32
	LinearSweptSpheres VkBool32
}

type VkAccelerationStructureGeometryMotionTrianglesDataNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	VertexData VkDeviceOrHostAddressConstKHR
}

type VkAccelerationStructureMotionInfoNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	MaxInstances uint32
	Flags        VkAccelerationStructureMotionInfoFlagsNV
}

type VkSRTDataNV struct {
	Sx  float32
	A   float32
	B   float32
	Pvx float32
	Sy  float32
	C   float32
	Pvy float32
	Sz  float32
	Pvz float32
	Qx  float32
	Qy  float32
	Qz  float32
	Qw  float32
	Tx  float32
	Ty  float32
	Tz  float32
}

type VkAccelerationStructureSRTMotionInstanceNV struct {
	TransformT0                            VkSRTDataNV
	TransformT1                            VkSRTDataNV
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  VkGeometryInstanceFlagsKHR
	AccelerationStructureReference         uint64
}

type VkAccelerationStructureMatrixMotionInstanceNV struct {
	TransformT0                            VkTransformMatrixKHR
	TransformT1                            VkTransformMatrixKHR
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  VkGeometryInstanceFlagsKHR
	AccelerationStructureReference         uint64
}

type VkAccelerationStructureMotionInstanceNV struct {
	Type  VkAccelerationStructureMotionInstanceTypeNV
	Flags VkAccelerationStructureMotionInstanceFlagsNV
	Data  VkAccelerationStructureMotionInstanceDataNV
}

type VkMemoryGetRemoteAddressInfoNV struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Memory     VkDeviceMemory
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkImportMemoryBufferCollectionFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Collection VkBufferCollectionFUCHSIA
	Index      uint32
}

type VkBufferCollectionImageCreateInfoFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Collection VkBufferCollectionFUCHSIA
	Index      uint32
}

type VkBufferCollectionBufferCreateInfoFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Collection VkBufferCollectionFUCHSIA
	Index      uint32
}

type VkBufferCollectionCreateInfoFUCHSIA struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	CollectionToken zx_handle_t
}

type VkBufferCollectionPropertiesFUCHSIA struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	MemoryTypeBits                   uint32
	BufferCount                      uint32
	CreateInfoIndex                  uint32
	SysmemPixelFormat                uint64
	FormatFeatures                   VkFormatFeatureFlags
	SysmemColorSpaceIndex            VkSysmemColorSpaceFUCHSIA
	SamplerYcbcrConversionComponents VkComponentMapping
	SuggestedYcbcrModel              VkSamplerYcbcrModelConversion
	SuggestedYcbcrRange              VkSamplerYcbcrRange
	SuggestedXChromaOffset           VkChromaLocation
	SuggestedYChromaOffset           VkChromaLocation
}

type VkBufferConstraintsInfoFUCHSIA struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	CreateInfo                  VkBufferCreateInfo
	RequiredFormatFeatures      VkFormatFeatureFlags
	BufferCollectionConstraints VkBufferCollectionConstraintsInfoFUCHSIA
}

type VkSysmemColorSpaceFUCHSIA struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	ColorSpace uint32
}

type VkImageFormatConstraintsInfoFUCHSIA struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ImageCreateInfo        VkImageCreateInfo
	RequiredFormatFeatures VkFormatFeatureFlags
	Flags                  VkImageFormatConstraintsFlagsFUCHSIA
	SysmemPixelFormat      uint64
	ColorSpaceCount        uint32
	PColorSpaces           *VkSysmemColorSpaceFUCHSIA
}

type VkImageConstraintsInfoFUCHSIA struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	FormatConstraintsCount      uint32
	PFormatConstraints          *VkImageFormatConstraintsInfoFUCHSIA
	BufferCollectionConstraints VkBufferCollectionConstraintsInfoFUCHSIA
	Flags                       VkImageConstraintsInfoFlagsFUCHSIA
}

type VkBufferCollectionConstraintsInfoFUCHSIA struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	MinBufferCount                  uint32
	MaxBufferCount                  uint32
	MinBufferCountForCamping        uint32
	MinBufferCountForDedicatedSlack uint32
	MinBufferCountForSharedSlack    uint32
}

type VkCudaModuleCreateInfoNV struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	DataSize uint
	PData    unsafe.Pointer
}

type VkCudaFunctionCreateInfoNV struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Module VkCudaModuleNV
	PName  string
}

type VkCudaLaunchInfoNV struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Function       VkCudaFunctionNV
	GridDimX       uint32
	GridDimY       uint32
	GridDimZ       uint32
	BlockDimX      uint32
	BlockDimY      uint32
	BlockDimZ      uint32
	SharedMemBytes uint32
	ParamCount     uint
	PParams        **void
	ExtraCount     uint
	PExtras        **void
}

type VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	FormatRgba10x6WithoutYCbCrSampler VkBool32
}

type VkFormatProperties3 struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	LinearTilingFeatures  VkFormatFeatureFlags2
	OptimalTilingFeatures VkFormatFeatureFlags2
	BufferFeatures        VkFormatFeatureFlags2
}

type VkFormatProperties3KHR struct {
}

type VkDrmFormatModifierPropertiesList2EXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	DrmFormatModifierCount       uint32
	PDrmFormatModifierProperties *VkDrmFormatModifierProperties2EXT
}

type VkDrmFormatModifierProperties2EXT struct {
	DrmFormatModifier               uint64
	DrmFormatModifierPlaneCount     uint32
	DrmFormatModifierTilingFeatures VkFormatFeatureFlags2
}

type VkAndroidHardwareBufferFormatProperties2ANDROID struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	Format                           VkFormat
	ExternalFormat                   uint64
	FormatFeatures                   VkFormatFeatureFlags2
	SamplerYcbcrConversionComponents VkComponentMapping
	SuggestedYcbcrModel              VkSamplerYcbcrModelConversion
	SuggestedYcbcrRange              VkSamplerYcbcrRange
	SuggestedXChromaOffset           VkChromaLocation
	SuggestedYChromaOffset           VkChromaLocation
}

type VkPipelineRenderingCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ViewMask                uint32
	ColorAttachmentCount    uint32
	PColorAttachmentFormats *VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
}

type VkPipelineRenderingCreateInfoKHR struct {
}

type VkRenderingInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	Flags                VkRenderingFlags
	RenderArea           VkRect2D
	LayerCount           uint32
	ViewMask             uint32
	ColorAttachmentCount uint32
	PColorAttachments    *VkRenderingAttachmentInfo
	PDepthAttachment     *VkRenderingAttachmentInfo
	PStencilAttachment   *VkRenderingAttachmentInfo
}

type VkRenderingInfoKHR struct {
}

type VkRenderingEndInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkRenderingEndInfoEXT struct {
}

type VkRenderingAttachmentInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ImageView          VkImageView
	ImageLayout        VkImageLayout
	ResolveMode        VkResolveModeFlagBits
	ResolveImageView   VkImageView
	ResolveImageLayout VkImageLayout
	LoadOp             VkAttachmentLoadOp
	StoreOp            VkAttachmentStoreOp
	ClearValue         VkClearValue
}

type VkRenderingAttachmentInfoKHR struct {
}

type VkRenderingFragmentShadingRateAttachmentInfoKHR struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	ImageView                      VkImageView
	ImageLayout                    VkImageLayout
	ShadingRateAttachmentTexelSize VkExtent2D
}

type VkRenderingFragmentDensityMapAttachmentInfoEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	ImageView   VkImageView
	ImageLayout VkImageLayout
}

type VkPhysicalDeviceDynamicRenderingFeatures struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DynamicRendering VkBool32
}

type VkPhysicalDeviceDynamicRenderingFeaturesKHR struct {
}

type VkCommandBufferInheritanceRenderingInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkRenderingFlags
	ViewMask                uint32
	ColorAttachmentCount    uint32
	ColorAttachmentCount    uint32
	PColorAttachmentFormats *VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
	RasterizationSamples    VkSampleCountFlagBits
}

type VkCommandBufferInheritanceRenderingInfoKHR struct {
}

type VkAttachmentSampleCountInfoAMD struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ColorAttachmentCount          uint32
	PColorAttachmentSamples       *VkSampleCountFlagBits
	DepthStencilAttachmentSamples VkSampleCountFlagBits
}

type VkAttachmentSampleCountInfoNV struct {
}

type VkMultiviewPerViewAttributesInfoNVX struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	PerViewAttributes              VkBool32
	PerViewAttributesPositionXOnly VkBool32
}

type VkPhysicalDeviceImageViewMinLodFeaturesEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	MinLod VkBool32
}

type VkImageViewMinLodCreateInfoEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	MinLod float32
}

type VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	RasterizationOrderColorAttachmentAccess   VkBool32
	RasterizationOrderDepthAttachmentAccess   VkBool32
	RasterizationOrderStencilAttachmentAccess VkBool32
}

type VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesARM struct {
}

type VkPhysicalDeviceLinearColorAttachmentFeaturesNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	LinearColorAttachment VkBool32
}

type VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	GraphicsPipelineLibrary VkBool32
}

type VkPhysicalDevicePipelineBinaryFeaturesKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	PipelineBinaries VkBool32
}

type VkDevicePipelineBinaryInternalCacheControlKHR struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	DisableInternalCache VkBool32
}

type VkPhysicalDevicePipelineBinaryPropertiesKHR struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	PipelineBinaryInternalCache            VkBool32
	PipelineBinaryInternalCacheControl     VkBool32
	PipelineBinaryPrefersInternalCache     VkBool32
	PipelineBinaryPrecompiledInternalCache VkBool32
	PipelineBinaryCompressedData           VkBool32
}

type VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT struct {
	SType                                                     VkStructureType
	PNext                                                     unsafe.Pointer
	GraphicsPipelineLibraryFastLinking                        VkBool32
	GraphicsPipelineLibraryIndependentInterpolationDecoration VkBool32
}

type VkGraphicsPipelineLibraryCreateInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkGraphicsPipelineLibraryFlagsEXT
}

type VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DescriptorSetHostMapping VkBool32
}

type VkDescriptorSetBindingReferenceVALVE struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	DescriptorSetLayout VkDescriptorSetLayout
	Binding             uint32
}

type VkDescriptorSetLayoutHostMappingInfoVALVE struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DescriptorOffset uint
	DescriptorSize   uint32
}

type VkPhysicalDeviceNestedCommandBufferFeaturesEXT struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	NestedCommandBuffer                VkBool32
	NestedCommandBufferRendering       VkBool32
	NestedCommandBufferSimultaneousUse VkBool32
}

type VkPhysicalDeviceNestedCommandBufferPropertiesEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MaxCommandBufferNestingLevel uint32
}

type VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ShaderModuleIdentifier VkBool32
}

type VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	ShaderModuleIdentifierAlgorithmUUID uint8
}

type VkPipelineShaderStageModuleIdentifierCreateInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	IdentifierSize uint32
	PIdentifier    *uint8
}

type VkShaderModuleIdentifierEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	IdentifierSize uint32
	Identifier     uint8
}

type VkImageCompressionControlEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	Flags                        VkImageCompressionFlagsEXT
	CompressionControlPlaneCount uint32
	PFixedRateFlags              *VkImageCompressionFixedRateFlagsEXT
}

type VkPhysicalDeviceImageCompressionControlFeaturesEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ImageCompressionControl VkBool32
}

type VkImageCompressionPropertiesEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	ImageCompressionFlags          VkImageCompressionFlagsEXT
	ImageCompressionFixedRateFlags VkImageCompressionFixedRateFlagsEXT
}

type VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	ImageCompressionControlSwapchain VkBool32
}

type VkImageSubresource2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ImageSubresource VkImageSubresource
}

type VkImageSubresource2KHR struct {
}

type VkImageSubresource2EXT struct {
}

type VkSubresourceLayout2 struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	SubresourceLayout VkSubresourceLayout
}

type VkSubresourceLayout2KHR struct {
}

type VkSubresourceLayout2EXT struct {
}

type VkRenderPassCreationControlEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	DisallowMerging VkBool32
}

type VkRenderPassCreationFeedbackInfoEXT struct {
	PostMergeSubpassCount uint32
}

type VkRenderPassCreationFeedbackCreateInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	PRenderPassFeedback *VkRenderPassCreationFeedbackInfoEXT
}

type VkRenderPassSubpassFeedbackInfoEXT struct {
	SubpassMergeStatus VkSubpassMergeStatusEXT
	Description        char
	PostMergeIndex     uint32
}

type VkRenderPassSubpassFeedbackCreateInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	PSubpassFeedback *VkRenderPassSubpassFeedbackInfoEXT
}

type VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	SubpassMergeFeedback VkBool32
}

type VkMicromapBuildInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Type                VkMicromapTypeEXT
	Flags               VkBuildMicromapFlagsEXT
	Mode                VkBuildMicromapModeEXT
	DstMicromap         VkMicromapEXT
	UsageCountsCount    uint32
	PUsageCounts        *VkMicromapUsageEXT
	PpUsageCounts       **VkMicromapUsageEXT
	Data                VkDeviceOrHostAddressConstKHR
	ScratchData         VkDeviceOrHostAddressKHR
	TriangleArray       VkDeviceOrHostAddressConstKHR
	TriangleArrayStride VkDeviceSize
}

type VkMicromapCreateInfoEXT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	CreateFlags   VkMicromapCreateFlagsEXT
	Buffer        VkBuffer
	Offset        VkDeviceSize
	Size          VkDeviceSize
	Type          VkMicromapTypeEXT
	DeviceAddress VkDeviceAddress
}

type VkMicromapVersionInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	PVersionData *uint8
}

type VkCopyMicromapInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Src   VkMicromapEXT
	Dst   VkMicromapEXT
	Mode  VkCopyMicromapModeEXT
}

type VkCopyMicromapToMemoryInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Src   VkMicromapEXT
	Dst   VkDeviceOrHostAddressKHR
	Mode  VkCopyMicromapModeEXT
}

type VkCopyMemoryToMicromapInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Src   VkDeviceOrHostAddressConstKHR
	Dst   VkMicromapEXT
	Mode  VkCopyMicromapModeEXT
}

type VkMicromapBuildSizesInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	MicromapSize     VkDeviceSize
	BuildScratchSize VkDeviceSize
	Discardable      VkBool32
}

type VkMicromapUsageEXT struct {
	Count            uint32
	SubdivisionLevel uint32
	Format           uint32
}

type VkMicromapTriangleEXT struct {
	DataOffset       uint32
	SubdivisionLevel uint16
	Format           uint16
}

type VkPhysicalDeviceOpacityMicromapFeaturesEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Micromap              VkBool32
	MicromapCaptureReplay VkBool32
	MicromapHostCommands  VkBool32
}

type VkPhysicalDeviceOpacityMicromapPropertiesEXT struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	MaxOpacity2StateSubdivisionLevel uint32
	MaxOpacity4StateSubdivisionLevel uint32
}

type VkAccelerationStructureTrianglesOpacityMicromapEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	IndexType        VkIndexType
	IndexBuffer      VkDeviceOrHostAddressConstKHR
	IndexStride      VkDeviceSize
	BaseTriangle     uint32
	UsageCountsCount uint32
	PUsageCounts     *VkMicromapUsageEXT
	PpUsageCounts    **VkMicromapUsageEXT
	Micromap         VkMicromapEXT
}

type VkPhysicalDeviceDisplacementMicromapFeaturesNV struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	DisplacementMicromap VkBool32
}

type VkPhysicalDeviceDisplacementMicromapPropertiesNV struct {
	SType                                   VkStructureType
	PNext                                   unsafe.Pointer
	MaxDisplacementMicromapSubdivisionLevel uint32
}

type VkAccelerationStructureTrianglesDisplacementMicromapNV struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	DisplacementBiasAndScaleFormat        VkFormat
	DisplacementVectorFormat              VkFormat
	DisplacementBiasAndScaleBuffer        VkDeviceOrHostAddressConstKHR
	DisplacementBiasAndScaleStride        VkDeviceSize
	DisplacementVectorBuffer              VkDeviceOrHostAddressConstKHR
	DisplacementVectorStride              VkDeviceSize
	DisplacedMicromapPrimitiveFlags       VkDeviceOrHostAddressConstKHR
	DisplacedMicromapPrimitiveFlagsStride VkDeviceSize
	IndexType                             VkIndexType
	IndexBuffer                           VkDeviceOrHostAddressConstKHR
	IndexStride                           VkDeviceSize
	BaseTriangle                          uint32
	UsageCountsCount                      uint32
	PUsageCounts                          *VkMicromapUsageEXT
	PpUsageCounts                         **VkMicromapUsageEXT
	Micromap                              VkMicromapEXT
}

type VkPipelinePropertiesIdentifierEXT struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PipelineIdentifier uint8
}

type VkPhysicalDevicePipelinePropertiesFeaturesEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	PipelinePropertiesIdentifier VkBool32
}

type VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	ShaderEarlyAndLateFragmentTests VkBool32
}

type VkExternalMemoryAcquireUnmodifiedEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	AcquireUnmodifiedMemory VkBool32
}

type VkExportMetalObjectCreateInfoEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ExportObjectType VkExportMetalObjectTypeFlagBitsEXT
}

type VkExportMetalObjectsInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkExportMetalDeviceInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	MtlDevice MTLDevice_id
}

type VkExportMetalCommandQueueInfoEXT struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Queue           VkQueue
	MtlCommandQueue MTLCommandQueue_id
}

type VkExportMetalBufferInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Memory    VkDeviceMemory
	MtlBuffer MTLBuffer_id
}

type VkImportMetalBufferInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	MtlBuffer MTLBuffer_id
}

type VkExportMetalTextureInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Image      VkImage
	ImageView  VkImageView
	BufferView VkBufferView
	Plane      VkImageAspectFlagBits
	MtlTexture MTLTexture_id
}

type VkImportMetalTextureInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Plane      VkImageAspectFlagBits
	MtlTexture MTLTexture_id
}

type VkExportMetalIOSurfaceInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Image     VkImage
	IoSurface IOSurfaceRef
}

type VkImportMetalIOSurfaceInfoEXT struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	IoSurface IOSurfaceRef
}

type VkExportMetalSharedEventInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Semaphore      VkSemaphore
	Event          VkEvent
	MtlSharedEvent MTLSharedEvent_id
}

type VkImportMetalSharedEventInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	MtlSharedEvent MTLSharedEvent_id
}

type VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	NonSeamlessCubeMap VkBool32
}

type VkPhysicalDevicePipelineRobustnessFeatures struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PipelineRobustness VkBool32
}

type VkPhysicalDevicePipelineRobustnessFeaturesEXT struct {
}

type VkPipelineRobustnessCreateInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	StorageBuffers VkPipelineRobustnessBufferBehavior
	UniformBuffers VkPipelineRobustnessBufferBehavior
	VertexInputs   VkPipelineRobustnessBufferBehavior
	Images         VkPipelineRobustnessImageBehavior
}

type VkPipelineRobustnessCreateInfoEXT struct {
}

type VkPhysicalDevicePipelineRobustnessProperties struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	DefaultRobustnessStorageBuffers VkPipelineRobustnessBufferBehavior
	DefaultRobustnessUniformBuffers VkPipelineRobustnessBufferBehavior
	DefaultRobustnessVertexInputs   VkPipelineRobustnessBufferBehavior
	DefaultRobustnessImages         VkPipelineRobustnessImageBehavior
}

type VkPhysicalDevicePipelineRobustnessPropertiesEXT struct {
}

type VkImageViewSampleWeightCreateInfoQCOM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	FilterCenter VkOffset2D
	FilterSize   VkExtent2D
	NumPhases    uint32
}

type VkPhysicalDeviceImageProcessingFeaturesQCOM struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	TextureSampleWeighted VkBool32
	TextureBoxFilter      VkBool32
	TextureBlockMatch     VkBool32
}

type VkPhysicalDeviceImageProcessingPropertiesQCOM struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	MaxWeightFilterPhases    uint32
	MaxWeightFilterDimension VkExtent2D
	MaxBlockMatchRegion      VkExtent2D
	MaxBoxFilterBlockSize    VkExtent2D
}

type VkPhysicalDeviceTilePropertiesFeaturesQCOM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	TileProperties VkBool32
}

type VkTilePropertiesQCOM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	TileSize  VkExtent3D
	ApronSize VkExtent2D
	Origin    VkOffset2D
}

type VkTileMemoryBindInfoQCOM struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
}

type VkPhysicalDeviceAmigoProfilingFeaturesSEC struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	AmigoProfiling VkBool32
}

type VkAmigoProfilingSubmitInfoSEC struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	FirstDrawTimestamp  uint64
	SwapBufferTimestamp uint64
}

type VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	AttachmentFeedbackLoopLayout VkBool32
}

type VkPhysicalDeviceDepthClampZeroOneFeaturesEXT struct {
}

type VkAttachmentFeedbackLoopInfoEXT struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	FeedbackLoopEnable VkBool32
}

type VkPhysicalDeviceAddressBindingReportFeaturesEXT struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ReportAddressBinding VkBool32
}

type VkRenderingAttachmentFlagsInfoKHR struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkRenderingAttachmentFlagsKHR
}

type VkResolveImageModeInfoKHR struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkResolveImageFlagsKHR
	ResolveMode        VkResolveModeFlagBits
	StencilResolveMode VkResolveModeFlagBits
}

type VkDeviceAddressBindingCallbackDataEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkDeviceAddressBindingFlagsEXT
	BaseAddress VkDeviceAddress
	Size        VkDeviceSize
	BindingType VkDeviceAddressBindingTypeEXT
}

type VkPhysicalDeviceOpticalFlowFeaturesNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	OpticalFlow VkBool32
}

type VkPhysicalDeviceOpticalFlowPropertiesNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	SupportedOutputGridSizes   VkOpticalFlowGridSizeFlagsNV
	SupportedHintGridSizes     VkOpticalFlowGridSizeFlagsNV
	HintSupported              VkBool32
	CostSupported              VkBool32
	BidirectionalFlowSupported VkBool32
	GlobalFlowSupported        VkBool32
	MinWidth                   uint32
	MinHeight                  uint32
	MaxWidth                   uint32
	MaxHeight                  uint32
	MaxNumRegionsOfInterest    uint32
}

type VkOpticalFlowImageFormatInfoNV struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkOpticalFlowUsageFlagsNV
}

type VkOpticalFlowImageFormatPropertiesNV struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Format VkFormat
}

type VkOpticalFlowSessionCreateInfoNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Width            uint32
	Height           uint32
	ImageFormat      VkFormat
	FlowVectorFormat VkFormat
	CostFormat       VkFormat
	OutputGridSize   VkOpticalFlowGridSizeFlagsNV
	HintGridSize     VkOpticalFlowGridSizeFlagsNV
	PerformanceLevel VkOpticalFlowPerformanceLevelNV
	Flags            VkOpticalFlowSessionCreateFlagsNV
}

type VkOpticalFlowSessionCreatePrivateDataInfoNV struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Id           uint32
	Size         uint32
	PPrivateData unsafe.Pointer
}

type VkOpticalFlowExecuteInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkOpticalFlowExecuteFlagsNV
	RegionCount uint32
	PRegions    *VkRect2D
}

type VkPhysicalDeviceFaultFeaturesEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	DeviceFault             VkBool32
	DeviceFaultVendorBinary VkBool32
}

type VkDeviceFaultAddressInfoEXT struct {
	AddressType      VkDeviceFaultAddressTypeEXT
	ReportedAddress  VkDeviceAddress
	AddressPrecision VkDeviceSize
}

type VkDeviceFaultVendorInfoEXT struct {
	Description     char
	VendorFaultCode uint64
	VendorFaultData uint64
}

type VkDeviceFaultCountsEXT struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	AddressInfoCount uint32
	VendorInfoCount  uint32
	VendorBinarySize VkDeviceSize
}

type VkDeviceFaultInfoEXT struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Description       char
	PAddressInfos     *VkDeviceFaultAddressInfoEXT
	PVendorInfos      *VkDeviceFaultVendorInfoEXT
	PVendorBinaryData unsafe.Pointer
}

type VkDeviceFaultVendorBinaryHeaderVersionOneEXT struct {
	HeaderSize            uint32
	HeaderVersion         VkDeviceFaultVendorBinaryHeaderVersionEXT
	VendorID              uint32
	DeviceID              uint32
	DriverVersion         uint32
	PipelineCacheUUID     uint8
	ApplicationNameOffset uint32
	ApplicationVersion    uint32
	EngineNameOffset      uint32
	EngineVersion         uint32
	ApiVersion            uint32
}

type VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PipelineLibraryGroupHandles VkBool32
}

type VkDepthBiasInfoEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
}

type VkDepthBiasRepresentationInfoEXT struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	DepthBiasRepresentation VkDepthBiasRepresentationEXT
	DepthBiasExact          VkBool32
}

type VkDecompressMemoryRegionNV struct {
	SrcAddress          VkDeviceAddress
	DstAddress          VkDeviceAddress
	CompressedSize      VkDeviceSize
	DecompressedSize    VkDeviceSize
	DecompressionMethod VkMemoryDecompressionMethodFlagsNV
}

type VkDecompressMemoryRegionEXT struct {
	SrcAddress       VkDeviceAddress
	DstAddress       VkDeviceAddress
	CompressedSize   VkDeviceSize
	DecompressedSize VkDeviceSize
}

type VkDecompressMemoryInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	DecompressionMethod VkMemoryDecompressionMethodFlagsEXT
	RegionCount         uint32
	PRegions            *VkDecompressMemoryRegionEXT
}

type VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ShaderCoreMask     uint64
	ShaderCoreCount    uint32
	ShaderWarpsPerCore uint32
}

type VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ShaderCoreBuiltins VkBool32
}

type VkFrameBoundaryEXT struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkFrameBoundaryFlagsEXT
	FrameID     uint64
	ImageCount  uint32
	PImages     *VkImage
	BufferCount uint32
	PBuffers    *VkBuffer
	TagName     uint64
	TagSize     uint
	PTag        unsafe.Pointer
}

type VkPhysicalDeviceFrameBoundaryFeaturesEXT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	FrameBoundary VkBool32
}

type VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	DynamicRenderingUnusedAttachments VkBool32
}

type VkPhysicalDeviceInternallySynchronizedQueuesFeaturesKHR struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	InternallySynchronizedQueues VkBool32
}

type VkSurfacePresentModeKHR struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PresentMode VkPresentModeKHR
}

type VkSurfacePresentModeEXT struct {
}

type VkSurfacePresentScalingCapabilitiesKHR struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	SupportedPresentScaling  VkPresentScalingFlagsKHR
	SupportedPresentGravityX VkPresentGravityFlagsKHR
	SupportedPresentGravityY VkPresentGravityFlagsKHR
	MinScaledImageExtent     VkExtent2D
	MaxScaledImageExtent     VkExtent2D
}

type VkSurfacePresentScalingCapabilitiesEXT struct {
}

type VkSurfacePresentModeCompatibilityKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	PresentModeCount uint32
	PPresentModes    *VkPresentModeKHR
}

type VkSurfacePresentModeCompatibilityEXT struct {
}

type VkPhysicalDeviceSwapchainMaintenance1FeaturesKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	SwapchainMaintenance1 VkBool32
}

type VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT struct {
}

type VkSwapchainPresentFenceInfoKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PFences        *VkFence
}

type VkSwapchainPresentFenceInfoEXT struct {
}

type VkSwapchainPresentModesCreateInfoKHR struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	PresentModeCount uint32
	PPresentModes    *VkPresentModeKHR
}

type VkSwapchainPresentModesCreateInfoEXT struct {
}

type VkSwapchainPresentModeInfoKHR struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SwapchainCount uint32
	PPresentModes  *VkPresentModeKHR
}

type VkSwapchainPresentModeInfoEXT struct {
}

type VkSwapchainPresentScalingCreateInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ScalingBehavior VkPresentScalingFlagsKHR
	PresentGravityX VkPresentGravityFlagsKHR
	PresentGravityY VkPresentGravityFlagsKHR
}

type VkSwapchainPresentScalingCreateInfoEXT struct {
}

type VkReleaseSwapchainImagesInfoKHR struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Swapchain       VkSwapchainKHR
	ImageIndexCount uint32
	PImageIndices   *uint32
}

type VkReleaseSwapchainImagesInfoEXT struct {
}

type VkPhysicalDeviceDepthBiasControlFeaturesEXT struct {
	SType                                           VkStructureType
	PNext                                           unsafe.Pointer
	DepthBiasControl                                VkBool32
	LeastRepresentableValueForceUnormRepresentation VkBool32
	FloatRepresentation                             VkBool32
	DepthBiasExact                                  VkBool32
}

type VkPhysicalDeviceRayTracingInvocationReorderFeaturesEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	RayTracingInvocationReorder VkBool32
}

type VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	RayTracingInvocationReorder VkBool32
}

type VkPhysicalDeviceRayTracingInvocationReorderPropertiesEXT struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	RayTracingInvocationReorderReorderingHint VkRayTracingInvocationReorderModeEXT
	MaxShaderBindingTableRecordIndex          uint32
}

type VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	RayTracingInvocationReorderReorderingHint VkRayTracingInvocationReorderModeEXT
}

type VkPhysicalDeviceExtendedSparseAddressSpaceFeaturesNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ExtendedSparseAddressSpace VkBool32
}

type VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	ExtendedSparseAddressSpaceSize VkDeviceSize
	ExtendedSparseImageUsageFlags  VkImageUsageFlags
	ExtendedSparseBufferUsageFlags VkBufferUsageFlags
}

type VkDirectDriverLoadingInfoLUNARG struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkDirectDriverLoadingFlagsLUNARG
	PfnGetInstanceProcAddr PFN_vkGetInstanceProcAddrLUNARG
}

type VkDirectDriverLoadingListLUNARG struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Mode        VkDirectDriverLoadingModeLUNARG
	DriverCount uint32
	PDrivers    *VkDirectDriverLoadingInfoLUNARG
}

type VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	MultiviewPerViewViewports VkBool32
}

type VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	RayTracingPositionFetch VkBool32
}

type VkDeviceImageSubresourceInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	PCreateInfo  *VkImageCreateInfo
	PSubresource *VkImageSubresource2
}

type VkDeviceImageSubresourceInfoKHR struct {
}

type VkPhysicalDeviceShaderCorePropertiesARM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	PixelRate uint32
	TexelRate uint32
	FmaRate   uint32
}

type VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	MultiviewPerViewRenderAreas VkBool32
}

type VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	PerViewRenderAreaCount uint32
	PPerViewRenderAreas    *VkRect2D
}

type VkQueryLowLatencySupportNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	PQueriedLowLatencyData unsafe.Pointer
}

type VkMemoryMapInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkMemoryMapFlags
	Memory VkDeviceMemory
	Offset VkDeviceSize
	Size   VkDeviceSize
}

type VkMemoryMapInfoKHR struct {
}

type VkMemoryUnmapInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkMemoryUnmapFlags
	Memory VkDeviceMemory
}

type VkMemoryUnmapInfoKHR struct {
}

type VkPhysicalDeviceShaderObjectFeaturesEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	ShaderObject VkBool32
}

type VkPhysicalDeviceShaderObjectPropertiesEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	ShaderBinaryUUID    uint8
	ShaderBinaryVersion uint32
}

type VkShaderCreateInfoEXT struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkShaderCreateFlagsEXT
	Stage                  VkShaderStageFlagBits
	NextStage              VkShaderStageFlags
	CodeType               VkShaderCodeTypeEXT
	CodeSize               uint
	PCode                  unsafe.Pointer
	PName                  string
	SetLayoutCount         uint32
	PSetLayouts            *VkDescriptorSetLayout
	PushConstantRangeCount uint32
	PPushConstantRanges    *VkPushConstantRange
	PSpecializationInfo    *VkSpecializationInfo
}

type VkPhysicalDeviceShaderTileImageFeaturesEXT struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	ShaderTileImageColorReadAccess   VkBool32
	ShaderTileImageDepthReadAccess   VkBool32
	ShaderTileImageStencilReadAccess VkBool32
}

type VkPhysicalDeviceShaderTileImagePropertiesEXT struct {
	SType                                            VkStructureType
	PNext                                            unsafe.Pointer
	ShaderTileImageCoherentReadAccelerated           VkBool32
	ShaderTileImageReadSampleFromPixelRateInvocation VkBool32
	ShaderTileImageReadFromHelperInvocation          VkBool32
}

type VkImportScreenBufferInfoQNX struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer *_screen_buffer
}

type VkScreenBufferPropertiesQNX struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	AllocationSize VkDeviceSize
	MemoryTypeBits uint32
}

type VkScreenBufferFormatPropertiesQNX struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	Format                           VkFormat
	ExternalFormat                   uint64
	ScreenUsage                      uint64
	FormatFeatures                   VkFormatFeatureFlags
	SamplerYcbcrConversionComponents VkComponentMapping
	SuggestedYcbcrModel              VkSamplerYcbcrModelConversion
	SuggestedYcbcrRange              VkSamplerYcbcrRange
	SuggestedXChromaOffset           VkChromaLocation
	SuggestedYChromaOffset           VkChromaLocation
}

type VkExternalFormatQNX struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	ExternalFormat uint64
}

type VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ScreenBufferImport VkBool32
}

type VkPhysicalDeviceCooperativeMatrixFeaturesKHR struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	CooperativeMatrix                   VkBool32
	CooperativeMatrixRobustBufferAccess VkBool32
}

type VkCooperativeMatrixPropertiesKHR struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	MSize                  uint32
	NSize                  uint32
	KSize                  uint32
	AType                  VkComponentTypeKHR
	BType                  VkComponentTypeKHR
	CType                  VkComponentTypeKHR
	ResultType             VkComponentTypeKHR
	SaturatingAccumulation VkBool32
	Scope                  VkScopeKHR
}

type VkPhysicalDeviceCooperativeMatrixPropertiesKHR struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	CooperativeMatrixSupportedStages VkShaderStageFlags
}

type VkPhysicalDeviceCooperativeMatrixConversionFeaturesQCOM struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	CooperativeMatrixConversion VkBool32
}

type VkPhysicalDeviceShaderEnqueuePropertiesAMDX struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	MaxExecutionGraphDepth                 uint32
	MaxExecutionGraphShaderOutputNodes     uint32
	MaxExecutionGraphShaderPayloadSize     uint32
	MaxExecutionGraphShaderPayloadCount    uint32
	ExecutionGraphDispatchAddressAlignment uint32
	MaxExecutionGraphWorkgroupCount        uint32
	MaxExecutionGraphWorkgroups            uint32
}

type VkPhysicalDeviceShaderEnqueueFeaturesAMDX struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	ShaderEnqueue     VkBool32
	ShaderMeshEnqueue VkBool32
}

type VkExecutionGraphPipelineCreateInfoAMDX struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkPipelineCreateFlags
	StageCount         uint32
	PStages            *VkPipelineShaderStageCreateInfo
	PLibraryInfo       *VkPipelineLibraryCreateInfoKHR
	Layout             VkPipelineLayout
	BasePipelineHandle VkPipeline
	BasePipelineIndex  int32
}

type VkPipelineShaderStageNodeCreateInfoAMDX struct {
	SType VkStructureType
	PNext unsafe.Pointer
	PName string
	Index uint32
}

type VkExecutionGraphPipelineScratchSizeAMDX struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	MinSize         VkDeviceSize
	MaxSize         VkDeviceSize
	SizeGranularity VkDeviceSize
}

type VkDispatchGraphInfoAMDX struct {
	NodeIndex     uint32
	PayloadCount  uint32
	Payloads      VkDeviceOrHostAddressConstAMDX
	PayloadStride uint64
}

type VkDispatchGraphCountInfoAMDX struct {
	Count  uint32
	Infos  VkDeviceOrHostAddressConstAMDX
	Stride uint64
}

type VkPhysicalDeviceAntiLagFeaturesAMD struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	AntiLag VkBool32
}

type VkAntiLagDataAMD struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Mode              VkAntiLagModeAMD
	MaxFPS            uint32
	PPresentationInfo *VkAntiLagPresentationInfoAMD
}

type VkAntiLagPresentationInfoAMD struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Stage      VkAntiLagStageAMD
	FrameIndex uint64
}

type VkBindMemoryStatus struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	PResult *VkResult
}

type VkPhysicalDeviceTileMemoryHeapFeaturesQCOM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	TileMemoryHeap VkBool32
}

type VkPhysicalDeviceTileMemoryHeapPropertiesQCOM struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	QueueSubmitBoundary VkBool32
	TileBufferTransfers VkBool32
}

type VkTileMemorySizeInfoQCOM struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Size  VkDeviceSize
}

type VkTileMemoryRequirementsQCOM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Size      VkDeviceSize
	Alignment VkDeviceSize
}

type VkBindMemoryStatusKHR struct {
}

type VkBindDescriptorSetsInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	StageFlags         VkShaderStageFlags
	Layout             VkPipelineLayout
	FirstSet           uint32
	DescriptorSetCount uint32
	PDescriptorSets    *VkDescriptorSet
	DynamicOffsetCount uint32
	PDynamicOffsets    *uint32
}

type VkBindDescriptorSetsInfoKHR struct {
}

type VkPushConstantsInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Layout     VkPipelineLayout
	StageFlags VkShaderStageFlags
	Offset     uint32
	Size       uint32
	PValues    unsafe.Pointer
}

type VkPushConstantsInfoKHR struct {
}

type VkPushDescriptorSetInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	StageFlags           VkShaderStageFlags
	Layout               VkPipelineLayout
	Set                  uint32
	DescriptorWriteCount uint32
	PDescriptorWrites    *VkWriteDescriptorSet
}

type VkPushDescriptorSetInfoKHR struct {
}

type VkPushDescriptorSetWithTemplateInfo struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DescriptorUpdateTemplate VkDescriptorUpdateTemplate
	Layout                   VkPipelineLayout
	Set                      uint32
	PData                    unsafe.Pointer
}

type VkPushDescriptorSetWithTemplateInfoKHR struct {
}

type VkSetDescriptorBufferOffsetsInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	StageFlags     VkShaderStageFlags
	Layout         VkPipelineLayout
	FirstSet       uint32
	SetCount       uint32
	PBufferIndices *uint32
	POffsets       *VkDeviceSize
}

type VkBindDescriptorBufferEmbeddedSamplersInfoEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	StageFlags VkShaderStageFlags
	Layout     VkPipelineLayout
	Set        uint32
}

type VkPhysicalDeviceCubicClampFeaturesQCOM struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	CubicRangeClamp VkBool32
}

type VkPhysicalDeviceYcbcrDegammaFeaturesQCOM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	YcbcrDegamma VkBool32
}

type VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	EnableYDegamma    VkBool32
	EnableCbCrDegamma VkBool32
}

type VkPhysicalDeviceCubicWeightsFeaturesQCOM struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	SelectableCubicWeights VkBool32
}

type VkSamplerCubicWeightsCreateInfoQCOM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	CubicWeights VkCubicFilterWeightsQCOM
}

type VkBlitImageCubicWeightsInfoQCOM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	CubicWeights VkCubicFilterWeightsQCOM
}

type VkPhysicalDeviceImageProcessing2FeaturesQCOM struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	TextureBlockMatch2 VkBool32
}

type VkPhysicalDeviceImageProcessing2PropertiesQCOM struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	MaxBlockMatchWindow VkExtent2D
}

type VkSamplerBlockMatchWindowCreateInfoQCOM struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	WindowExtent      VkExtent2D
	WindowCompareMode VkBlockMatchWindowCompareModeQCOM
}

type VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	DescriptorPoolOverallocation VkBool32
}

type VkPhysicalDeviceLayeredDriverPropertiesMSFT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	UnderlyingAPI VkLayeredDriverUnderlyingApiMSFT
}

type VkPhysicalDevicePerStageDescriptorSetFeaturesNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	PerStageDescriptorSet VkBool32
	DynamicPipelineLayout VkBool32
}

type VkPhysicalDeviceExternalFormatResolveFeaturesANDROID struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ExternalFormatResolve VkBool32
}

type VkPhysicalDeviceExternalFormatResolvePropertiesANDROID struct {
	SType                                        VkStructureType
	PNext                                        unsafe.Pointer
	NullColorAttachmentWithExternalFormatResolve VkBool32
	ExternalFormatResolveChromaOffsetX           VkChromaLocation
	ExternalFormatResolveChromaOffsetY           VkChromaLocation
}

type VkAndroidHardwareBufferFormatResolvePropertiesANDROID struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ColorAttachmentFormat VkFormat
}

type VkLatencySleepModeInfoNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	LowLatencyMode    VkBool32
	LowLatencyBoost   VkBool32
	MinimumIntervalUs uint32
}

type VkLatencySleepInfoNV struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	SignalSemaphore VkSemaphore
	Value           uint64
}

type VkSetLatencyMarkerInfoNV struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	PresentID uint64
	Marker    VkLatencyMarkerNV
}

type VkGetLatencyMarkerInfoNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	TimingCount uint32
	PTimings    *VkLatencyTimingsFrameReportNV
}

type VkLatencyTimingsFrameReportNV struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	PresentID                uint64
	InputSampleTimeUs        uint64
	SimStartTimeUs           uint64
	SimEndTimeUs             uint64
	RenderSubmitStartTimeUs  uint64
	RenderSubmitEndTimeUs    uint64
	PresentStartTimeUs       uint64
	PresentEndTimeUs         uint64
	DriverStartTimeUs        uint64
	DriverEndTimeUs          uint64
	OsRenderQueueStartTimeUs uint64
	OsRenderQueueEndTimeUs   uint64
	GpuRenderStartTimeUs     uint64
	GpuRenderEndTimeUs       uint64
}

type VkOutOfBandQueueTypeInfoNV struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	QueueType VkOutOfBandQueueTypeNV
}

type VkLatencySubmissionPresentIdNV struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	PresentID uint64
}

type VkSwapchainLatencyCreateInfoNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	LatencyModeEnable VkBool32
}

type VkLatencySurfaceCapabilitiesNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	PresentModeCount uint32
	PPresentModes    *VkPresentModeKHR
}

type VkPhysicalDeviceCudaKernelLaunchFeaturesNV struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	CudaKernelLaunchFeatures VkBool32
}

type VkPhysicalDeviceCudaKernelLaunchPropertiesNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ComputeCapabilityMinor uint32
	ComputeCapabilityMajor uint32
}

type VkDeviceQueueShaderCoreControlCreateInfoARM struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ShaderCoreCount uint32
}

type VkPhysicalDeviceSchedulingControlsFeaturesARM struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	SchedulingControls VkBool32
}

type VkPhysicalDeviceSchedulingControlsPropertiesARM struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	SchedulingControlsFlags VkPhysicalDeviceSchedulingControlsFlagsARM
}

type VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	RelaxedLineRasterization VkBool32
}

type VkPhysicalDeviceRenderPassStripedFeaturesARM struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	RenderPassStriped VkBool32
}

type VkPhysicalDeviceRenderPassStripedPropertiesARM struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	RenderPassStripeGranularity VkExtent2D
	MaxRenderPassStripes        uint32
}

type VkRenderPassStripeInfoARM struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	StripeArea VkRect2D
}

type VkRenderPassStripeBeginInfoARM struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	StripeInfoCount uint32
	PStripeInfos    *VkRenderPassStripeInfoARM
}

type VkRenderPassStripeSubmitInfoARM struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	StripeSemaphoreInfoCount uint32
	PStripeSemaphoreInfos    *VkSemaphoreSubmitInfo
}

type VkPhysicalDevicePipelineOpacityMicromapFeaturesARM struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	PipelineOpacityMicromap VkBool32
}

type VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ShaderMaximalReconvergence VkBool32
}

type VkPhysicalDeviceShaderSubgroupRotateFeatures struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ShaderSubgroupRotate          VkBool32
	ShaderSubgroupRotateClustered VkBool32
}

type VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR struct {
}

type VkPhysicalDeviceShaderExpectAssumeFeatures struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ShaderExpectAssume VkBool32
}

type VkPhysicalDeviceShaderExpectAssumeFeaturesKHR struct {
}

type VkPhysicalDeviceShaderFloatControls2Features struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ShaderFloatControls2 VkBool32
}

type VkPhysicalDeviceShaderFloatControls2FeaturesKHR struct {
}

type VkPhysicalDeviceDynamicRenderingLocalReadFeatures struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	DynamicRenderingLocalRead VkBool32
}

type VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR struct {
}

type VkRenderingAttachmentLocationInfo struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	ColorAttachmentCount      uint32
	PColorAttachmentLocations *uint32
}

type VkRenderingAttachmentLocationInfoKHR struct {
}

type VkRenderingInputAttachmentIndexInfo struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	ColorAttachmentCount         uint32
	PColorAttachmentInputIndices *uint32
	PDepthInputAttachmentIndex   *uint32
	PStencilInputAttachmentIndex *uint32
}

type VkRenderingInputAttachmentIndexInfoKHR struct {
}

type VkPhysicalDeviceShaderQuadControlFeaturesKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	ShaderQuadControl VkBool32
}

type VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ShaderFloat16VectorAtomics VkBool32
}

type VkPhysicalDeviceMapMemoryPlacedFeaturesEXT struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	MemoryMapPlaced      VkBool32
	MemoryMapRangePlaced VkBool32
	MemoryUnmapReserve   VkBool32
}

type VkPhysicalDeviceMapMemoryPlacedPropertiesEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	MinPlacedMemoryMapAlignment VkDeviceSize
}

type VkMemoryMapPlacedInfoEXT struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PPlacedAddress unsafe.Pointer
}

type VkPhysicalDeviceShaderBfloat16FeaturesKHR struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	ShaderBFloat16Type              VkBool32
	ShaderBFloat16DotProduct        VkBool32
	ShaderBFloat16CooperativeMatrix VkBool32
}

type VkPhysicalDeviceRawAccessChainsFeaturesNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ShaderRawAccessChains VkBool32
}

type VkPhysicalDeviceCommandBufferInheritanceFeaturesNV struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	CommandBufferInheritance VkBool32
}

type VkPhysicalDeviceImageAlignmentControlFeaturesMESA struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ImageAlignmentControl VkBool32
}

type VkPhysicalDeviceImageAlignmentControlPropertiesMESA struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	SupportedImageAlignmentMask uint32
}

type VkImageAlignmentControlCreateInfoMESA struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	MaximumRequestedAlignment uint32
}

type VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	ShaderReplicatedComposites VkBool32
}

type VkPhysicalDevicePresentModeFifoLatestReadyFeaturesEXT struct {
}

type VkPhysicalDevicePresentModeFifoLatestReadyFeaturesKHR struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	PresentModeFifoLatestReady VkBool32
}

type VkDepthClampRangeEXT struct {
	MinDepthClamp float32
	MaxDepthClamp float32
}

type VkPhysicalDeviceCooperativeMatrix2FeaturesNV struct {
	SType                                 VkStructureType
	PNext                                 unsafe.Pointer
	CooperativeMatrixWorkgroupScope       VkBool32
	CooperativeMatrixFlexibleDimensions   VkBool32
	CooperativeMatrixReductions           VkBool32
	CooperativeMatrixConversions          VkBool32
	CooperativeMatrixPerElementOperations VkBool32
	CooperativeMatrixTensorAddressing     VkBool32
	CooperativeMatrixBlockLoads           VkBool32
}

type VkPhysicalDeviceCooperativeMatrix2PropertiesNV struct {
	SType                                               VkStructureType
	PNext                                               unsafe.Pointer
	CooperativeMatrixWorkgroupScopeMaxWorkgroupSize     uint32
	CooperativeMatrixFlexibleDimensionsMaxDimension     uint32
	CooperativeMatrixWorkgroupScopeReservedSharedMemory uint32
}

type VkCooperativeMatrixFlexibleDimensionsPropertiesNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	MGranularity           uint32
	NGranularity           uint32
	KGranularity           uint32
	AType                  VkComponentTypeKHR
	BType                  VkComponentTypeKHR
	CType                  VkComponentTypeKHR
	ResultType             VkComponentTypeKHR
	SaturatingAccumulation VkBool32
	Scope                  VkScopeKHR
	WorkgroupInvocations   uint32
}

type VkPhysicalDeviceHdrVividFeaturesHUAWEI struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	HdrVivid VkBool32
}

type VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	VertexAttributeRobustness VkBool32
}

type VkPhysicalDeviceDenseGeometryFormatFeaturesAMDX struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	DenseGeometryFormat VkBool32
}

type VkAccelerationStructureDenseGeometryFormatTrianglesDataAMDX struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	CompressedData    VkDeviceOrHostAddressConstKHR
	DataSize          VkDeviceSize
	NumTriangles      uint32
	NumVertices       uint32
	MaxPrimitiveIndex uint32
	MaxGeometryIndex  uint32
	Format            VkCompressedTriangleFormatAMDX
}

type VkPhysicalDeviceDepthClampZeroOneFeaturesKHR struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DepthClampZeroOne VkBool32
}

type VkPhysicalDeviceCooperativeVectorFeaturesNV struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	CooperativeVector         VkBool32
	CooperativeVectorTraining VkBool32
}

type VkCooperativeVectorPropertiesNV struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	InputType            VkComponentTypeKHR
	InputInterpretation  VkComponentTypeKHR
	MatrixInterpretation VkComponentTypeKHR
	BiasInterpretation   VkComponentTypeKHR
	ResultType           VkComponentTypeKHR
	Transpose            VkBool32
}

type VkPhysicalDeviceCooperativeVectorPropertiesNV struct {
	SType                                        VkStructureType
	PNext                                        unsafe.Pointer
	CooperativeVectorSupportedStages             VkShaderStageFlags
	CooperativeVectorTrainingFloat16Accumulation VkBool32
	CooperativeVectorTrainingFloat32Accumulation VkBool32
	MaxCooperativeVectorComponents               uint32
}

type VkConvertCooperativeVectorMatrixInfoNV struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	SrcSize          uint
	SrcData          VkDeviceOrHostAddressConstKHR
	PDstSize         *uint
	DstData          VkDeviceOrHostAddressKHR
	SrcComponentType VkComponentTypeKHR
	DstComponentType VkComponentTypeKHR
	NumRows          uint32
	NumColumns       uint32
	SrcLayout        VkCooperativeVectorMatrixLayoutNV
	SrcStride        uint
	DstLayout        VkCooperativeVectorMatrixLayoutNV
	DstStride        uint
}

type VkPhysicalDeviceTileShadingFeaturesQCOM struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	TileShading                   VkBool32
	TileShadingFragmentStage      VkBool32
	TileShadingColorAttachments   VkBool32
	TileShadingDepthAttachments   VkBool32
	TileShadingStencilAttachments VkBool32
	TileShadingInputAttachments   VkBool32
	TileShadingSampledAttachments VkBool32
	TileShadingPerTileDraw        VkBool32
	TileShadingPerTileDispatch    VkBool32
	TileShadingDispatchTile       VkBool32
	TileShadingApron              VkBool32
	TileShadingAnisotropicApron   VkBool32
	TileShadingAtomicOps          VkBool32
	TileShadingImageProcessing    VkBool32
}

type VkPhysicalDeviceTileShadingPropertiesQCOM struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxApronSize       uint32
	PreferNonCoherent  VkBool32
	TileGranularity    VkExtent2D
	MaxTileShadingRate VkExtent2D
}

type VkRenderPassTileShadingCreateInfoQCOM struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Flags         VkTileShadingRenderPassFlagsQCOM
	TileApronSize VkExtent2D
}

type VkPerTileBeginInfoQCOM struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkPerTileEndInfoQCOM struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkDispatchTileInfoQCOM struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkPhysicalDeviceFragmentDensityMapLayeredPropertiesVALVE struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	MaxFragmentDensityMapLayers uint32
}

type VkPhysicalDeviceFragmentDensityMapLayeredFeaturesVALVE struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	FragmentDensityMapLayered VkBool32
}

type VkPipelineFragmentDensityMapLayeredCreateInfoVALVE struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	MaxFragmentDensityMapLayers uint32
}

type VkSetPresentConfigNV struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	NumFramesPerBatch     uint32
	PresentConfigFeedback uint32
}

type VkPhysicalDevicePresentMeteringFeaturesNV struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	PresentMetering VkBool32
}

type VkExternalComputeQueueDeviceCreateInfoNV struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	ReservedExternalQueues uint32
}

type VkExternalComputeQueueCreateInfoNV struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PreferredQueue VkQueue
}

type VkExternalComputeQueueDataParamsNV struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	DeviceIndex uint32
}

type VkPhysicalDeviceExternalComputeQueuePropertiesNV struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	ExternalDataSize  uint32
	MaxExternalQueues uint32
}

type VkPhysicalDeviceShaderUniformBufferUnsizedArrayFeaturesEXT struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	ShaderUniformBufferUnsizedArray VkBool32
}

type VkPhysicalDeviceShaderMixedFloatDotProductFeaturesVALVE struct {
	SType                                       VkStructureType
	PNext                                       unsafe.Pointer
	ShaderMixedFloatDotProductFloat16AccFloat32 VkBool32
	ShaderMixedFloatDotProductFloat16AccFloat16 VkBool32
	ShaderMixedFloatDotProductBFloat16Acc       VkBool32
	ShaderMixedFloatDotProductFloat8AccFloat32  VkBool32
}

type VkPhysicalDeviceFormatPackFeaturesARM struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	FormatPack VkBool32
}

type VkTensorDescriptionARM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Tiling         VkTensorTilingARM
	Format         VkFormat
	DimensionCount uint32
	PDimensions    *int64
	PStrides       *int64
	Usage          VkTensorUsageFlagsARM
}

type VkTensorCreateInfoARM struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	Flags                 VkTensorCreateFlagsARM
	PDescription          *VkTensorDescriptionARM
	SharingMode           VkSharingMode
	QueueFamilyIndexCount uint32
	PQueueFamilyIndices   *uint32
}

type VkTensorViewCreateInfoARM struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkTensorViewCreateFlagsARM
	Tensor VkTensorARM
	Format VkFormat
}

type VkTensorMemoryRequirementsInfoARM struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Tensor VkTensorARM
}

type VkBindTensorMemoryInfoARM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Tensor       VkTensorARM
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}

type VkWriteDescriptorSetTensorARM struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	TensorViewCount uint32
	PTensorViews    *VkTensorViewARM
}

type VkTensorFormatPropertiesARM struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	OptimalTilingTensorFeatures VkFormatFeatureFlags2
	LinearTilingTensorFeatures  VkFormatFeatureFlags2
}

type VkPhysicalDeviceTensorPropertiesARM struct {
	SType                                              VkStructureType
	PNext                                              unsafe.Pointer
	MaxTensorDimensionCount                            uint32
	MaxTensorElements                                  uint64
	MaxPerDimensionTensorElements                      uint64
	MaxTensorStride                                    int64
	MaxTensorSize                                      uint64
	MaxTensorShaderAccessArrayLength                   uint32
	MaxTensorShaderAccessSize                          uint32
	MaxDescriptorSetStorageTensors                     uint32
	MaxPerStageDescriptorSetStorageTensors             uint32
	MaxDescriptorSetUpdateAfterBindStorageTensors      uint32
	MaxPerStageDescriptorUpdateAfterBindStorageTensors uint32
	ShaderStorageTensorArrayNonUniformIndexingNative   VkBool32
	ShaderTensorSupportedStages                        VkShaderStageFlags
}

type VkTensorMemoryBarrierARM struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	SrcStageMask        VkPipelineStageFlags2
	SrcAccessMask       VkAccessFlags2
	DstStageMask        VkPipelineStageFlags2
	DstAccessMask       VkAccessFlags2
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Tensor              VkTensorARM
}

type VkTensorDependencyInfoARM struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	TensorMemoryBarrierCount uint32
	PTensorMemoryBarriers    *VkTensorMemoryBarrierARM
}

type VkPhysicalDeviceTensorFeaturesARM struct {
	SType                                         VkStructureType
	PNext                                         unsafe.Pointer
	TensorNonPacked                               VkBool32
	ShaderTensorAccess                            VkBool32
	ShaderStorageTensorArrayDynamicIndexing       VkBool32
	ShaderStorageTensorArrayNonUniformIndexing    VkBool32
	DescriptorBindingStorageTensorUpdateAfterBind VkBool32
	Tensors                                       VkBool32
}

type VkDeviceTensorMemoryRequirementsARM struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PCreateInfo *VkTensorCreateInfoARM
}

type VkCopyTensorInfoARM struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	SrcTensor   VkTensorARM
	DstTensor   VkTensorARM
	RegionCount uint32
	PRegions    *VkTensorCopyARM
}

type VkTensorCopyARM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	DimensionCount uint32
	PSrcOffset     *uint64
	PDstOffset     *uint64
	PExtent        *uint64
}

type VkMemoryDedicatedAllocateInfoTensorARM struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Tensor VkTensorARM
}

type VkPhysicalDeviceDescriptorBufferTensorPropertiesARM struct {
	SType                                     VkStructureType
	PNext                                     unsafe.Pointer
	TensorCaptureReplayDescriptorDataSize     uint
	TensorViewCaptureReplayDescriptorDataSize uint
	TensorDescriptorSize                      uint
}

type VkPhysicalDeviceDescriptorBufferTensorFeaturesARM struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	DescriptorBufferTensorDescriptors VkBool32
}

type VkTensorCaptureDescriptorDataInfoARM struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Tensor VkTensorARM
}

type VkTensorViewCaptureDescriptorDataInfoARM struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	TensorView VkTensorViewARM
}

type VkDescriptorGetTensorInfoARM struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	TensorView VkTensorViewARM
}

type VkFrameBoundaryTensorsARM struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	TensorCount uint32
	PTensors    *VkTensorARM
}

type VkPhysicalDeviceExternalTensorInfoARM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Flags        VkTensorCreateFlagsARM
	PDescription *VkTensorDescriptionARM
	HandleType   VkExternalMemoryHandleTypeFlagBits
}

type VkExternalTensorPropertiesARM struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties VkExternalMemoryProperties
}

type VkExternalMemoryTensorCreateInfoARM struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

type VkPhysicalDeviceShaderFloat8FeaturesEXT struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ShaderFloat8                  VkBool32
	ShaderFloat8CooperativeMatrix VkBool32
}

type VkSurfaceCreateInfoOHOS struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkSurfaceCreateFlagsOHOS
	Window *OHNativeWindow
}

type VkPhysicalDeviceDataGraphFeaturesARM struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	DataGraph                        VkBool32
	DataGraphUpdateAfterBind         VkBool32
	DataGraphSpecializationConstants VkBool32
	DataGraphDescriptorBuffer        VkBool32
	DataGraphShaderModule            VkBool32
}

type VkDataGraphPipelineConstantTensorSemiStructuredSparsityInfoARM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Dimension uint32
	ZeroCount uint32
	GroupSize uint32
}

type VkDataGraphPipelineConstantARM struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Id            uint32
	PConstantData unsafe.Pointer
}

type VkDataGraphPipelineResourceInfoARM struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	DescriptorSet uint32
	Binding       uint32
	ArrayElement  uint32
}

type VkDataGraphPipelineCompilerControlCreateInfoARM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PVendorOptions string
}

type VkDataGraphPipelineCreateInfoARM struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Flags             VkPipelineCreateFlags2KHR
	Layout            VkPipelineLayout
	ResourceInfoCount uint32
	PResourceInfos    *VkDataGraphPipelineResourceInfoARM
}

type VkDataGraphPipelineShaderModuleCreateInfoARM struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Module              VkShaderModule
	PName               string
	PSpecializationInfo *VkSpecializationInfo
	ConstantCount       uint32
	PConstants          *VkDataGraphPipelineConstantARM
}

type VkDataGraphPipelineSessionCreateInfoARM struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Flags             VkDataGraphPipelineSessionCreateFlagsARM
	DataGraphPipeline VkPipeline
}

type VkDataGraphPipelineSessionBindPointRequirementsInfoARM struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	Session VkDataGraphPipelineSessionARM
}

type VkDataGraphPipelineSessionBindPointRequirementARM struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	BindPoint     VkDataGraphPipelineSessionBindPointARM
	BindPointType VkDataGraphPipelineSessionBindPointTypeARM
	NumObjects    uint32
}

type VkDataGraphPipelineSessionMemoryRequirementsInfoARM struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Session     VkDataGraphPipelineSessionARM
	BindPoint   VkDataGraphPipelineSessionBindPointARM
	ObjectIndex uint32
}

type VkBindDataGraphPipelineSessionMemoryInfoARM struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Session      VkDataGraphPipelineSessionARM
	BindPoint    VkDataGraphPipelineSessionBindPointARM
	ObjectIndex  uint32
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}

type VkDataGraphPipelineInfoARM struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DataGraphPipeline VkPipeline
}

type VkDataGraphPipelinePropertyQueryResultARM struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Property VkDataGraphPipelinePropertyARM
	IsText   VkBool32
	DataSize uint
	PData    unsafe.Pointer
}

type VkDataGraphPipelineIdentifierCreateInfoARM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	IdentifierSize uint32
	PIdentifier    *uint8
}

type VkDataGraphPipelineDispatchInfoARM struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkDataGraphPipelineDispatchFlagsARM
}

type VkPhysicalDeviceDataGraphProcessingEngineARM struct {
	Type      VkPhysicalDeviceDataGraphProcessingEngineTypeARM
	IsForeign VkBool32
}

type VkPhysicalDeviceDataGraphOperationSupportARM struct {
	OperationType VkPhysicalDeviceDataGraphOperationTypeARM
	Name          char
	Version       uint32
}

type VkQueueFamilyDataGraphPropertiesARM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Engine    VkPhysicalDeviceDataGraphProcessingEngineARM
	Operation VkPhysicalDeviceDataGraphOperationSupportARM
}

type VkPhysicalDeviceQueueFamilyDataGraphProcessingEngineInfoARM struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	QueueFamilyIndex uint32
	EngineType       VkPhysicalDeviceDataGraphProcessingEngineTypeARM
}

type VkQueueFamilyDataGraphProcessingEnginePropertiesARM struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	ForeignSemaphoreHandleTypes VkExternalSemaphoreHandleTypeFlags
	ForeignMemoryHandleTypes    VkExternalMemoryHandleTypeFlags
}

type VkDataGraphProcessingEngineCreateInfoARM struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ProcessingEngineCount uint32
	PProcessingEngines    *VkPhysicalDeviceDataGraphProcessingEngineARM
}

type VkPhysicalDevicePipelineCacheIncrementalModeFeaturesSEC struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	PipelineCacheIncrementalMode VkBool32
}

type VkDataGraphPipelineBuiltinModelCreateInfoQCOM struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	POperation *VkPhysicalDeviceDataGraphOperationSupportARM
}

type VkPhysicalDeviceDataGraphModelFeaturesQCOM struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	DataGraphModel VkBool32
}

type VkPhysicalDeviceShaderUntypedPointersFeaturesKHR struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ShaderUntypedPointers VkBool32
}

type VkNativeBufferOHOS struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Handle *OHBufferHandle
}

type VkSwapchainImageCreateInfoOHOS struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkSwapchainImageUsageFlagsOHOS
}

type VkPhysicalDevicePresentationPropertiesOHOS struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	SharedImage VkBool32
}

type VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	VideoEncodeRgbConversion VkBool32
}

type VkVideoEncodeRgbConversionCapabilitiesVALVE struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	RgbModels      VkVideoEncodeRgbModelConversionFlagsVALVE
	RgbRanges      VkVideoEncodeRgbRangeCompressionFlagsVALVE
	XChromaOffsets VkVideoEncodeRgbChromaOffsetFlagsVALVE
	YChromaOffsets VkVideoEncodeRgbChromaOffsetFlagsVALVE
}

type VkVideoEncodeProfileRgbConversionInfoVALVE struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	PerformEncodeRgbConversion VkBool32
}

type VkVideoEncodeSessionRgbConversionCreateInfoVALVE struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	RgbModel      VkVideoEncodeRgbModelConversionFlagBitsVALVE
	RgbRange      VkVideoEncodeRgbRangeCompressionFlagBitsVALVE
	XChromaOffset VkVideoEncodeRgbChromaOffsetFlagBitsVALVE
	YChromaOffset VkVideoEncodeRgbChromaOffsetFlagBitsVALVE
}

type VkPhysicalDeviceShader64BitIndexingFeaturesEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Shader64BitIndexing VkBool32
}

type VkNativeBufferUsageOHOS struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	OHOSNativeBufferUsage uint64
}

type VkNativeBufferPropertiesOHOS struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	AllocationSize VkDeviceSize
	MemoryTypeBits uint32
}

type VkNativeBufferFormatPropertiesOHOS struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	Format                           VkFormat
	ExternalFormat                   uint64
	FormatFeatures                   VkFormatFeatureFlags
	SamplerYcbcrConversionComponents VkComponentMapping
	SuggestedYcbcrModel              VkSamplerYcbcrModelConversion
	SuggestedYcbcrRange              VkSamplerYcbcrRange
	SuggestedXChromaOffset           VkChromaLocation
	SuggestedYChromaOffset           VkChromaLocation
}

type VkImportNativeBufferInfoOHOS struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer *OH_NativeBuffer
}

type VkMemoryGetNativeBufferInfoOHOS struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
}

type VkExternalFormatOHOS struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	ExternalFormat uint64
}

type VkPhysicalDevicePerformanceCountersByRegionFeaturesARM struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PerformanceCountersByRegion VkBool32
}

type VkPhysicalDevicePerformanceCountersByRegionPropertiesARM struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	MaxPerRegionPerformanceCounters uint32
	PerformanceCounterRegionSize    VkExtent2D
	RowStrideAlignment              uint32
	RegionAlignment                 uint32
	IdentityTransformOrder          VkBool32
}

type VkPerformanceCounterARM struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	CounterID uint32
}

type VkPerformanceCounterDescriptionARM struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkPerformanceCounterDescriptionFlagsARM
	Name  char
}

type VkRenderPassPerformanceCountersByRegionBeginInfoARM struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	CounterAddressCount uint32
	PCounterAddresses   *VkDeviceAddress
	SerializeRegions    VkBool32
	CounterIndexCount   uint32
	PCounterIndices     *uint32
}

type VkComputeOccupancyPriorityParametersNV struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	OccupancyPriority   float32
	OccupancyThrottling float32
}

type VkPhysicalDeviceComputeOccupancyPriorityFeaturesNV struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ComputeOccupancyPriority VkBool32
}

type VkPhysicalDeviceShaderLongVectorFeaturesEXT struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	LongVector VkBool32
}

type VkPhysicalDeviceShaderLongVectorPropertiesEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	MaxVectorComponents uint32
}

type VkPhysicalDeviceTextureCompressionASTC3DFeaturesEXT struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	TextureCompressionASTC_3D VkBool32
}

type VkPhysicalDeviceShaderSubgroupPartitionedFeaturesEXT struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	ShaderSubgroupPartitioned VkBool32
}

type VkHostAddressRangeEXT struct {
	Address unsafe.Pointer
	Size    uint
}

type VkHostAddressRangeConstEXT struct {
	Address unsafe.Pointer
	Size    uint
}

type VkDeviceAddressRangeEXT struct {
	Address VkDeviceAddress
	Size    VkDeviceSize
}

type VkTexelBufferDescriptorInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Format       VkFormat
	AddressRange VkDeviceAddressRangeEXT
}

type VkImageDescriptorInfoEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	PView  *VkImageViewCreateInfo
	Layout VkImageLayout
}

type VkResourceDescriptorInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Type  VkDescriptorType
	Data  VkResourceDescriptorDataEXT
}

type VkBindHeapInfoEXT struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	HeapRange           VkDeviceAddressRangeEXT
	ReservedRangeOffset VkDeviceSize
	ReservedRangeSize   VkDeviceSize
}

type VkPushDataInfoEXT struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Offset uint32
	Data   VkHostAddressRangeConstEXT
}

type VkDescriptorMappingSourceConstantOffsetEXT struct {
	HeapOffset             uint32
	HeapArrayStride        uint32
	PEmbeddedSampler       *VkSamplerCreateInfo
	SamplerHeapOffset      uint32
	SamplerHeapArrayStride uint32
}

type VkDescriptorMappingSourcePushIndexEXT struct {
	HeapOffset                   uint32
	PushOffset                   uint32
	HeapIndexStride              uint32
	HeapArrayStride              uint32
	PEmbeddedSampler             *VkSamplerCreateInfo
	UseCombinedImageSamplerIndex VkBool32
	SamplerHeapOffset            uint32
	SamplerPushOffset            uint32
	SamplerHeapIndexStride       uint32
	SamplerHeapArrayStride       uint32
}

type VkDescriptorMappingSourceIndirectIndexEXT struct {
	HeapOffset                   uint32
	PushOffset                   uint32
	AddressOffset                uint32
	HeapIndexStride              uint32
	HeapArrayStride              uint32
	PEmbeddedSampler             *VkSamplerCreateInfo
	UseCombinedImageSamplerIndex VkBool32
	SamplerHeapOffset            uint32
	SamplerPushOffset            uint32
	SamplerAddressOffset         uint32
	SamplerHeapIndexStride       uint32
	SamplerHeapArrayStride       uint32
}

type VkDescriptorMappingSourceIndirectIndexArrayEXT struct {
	HeapOffset                   uint32
	PushOffset                   uint32
	AddressOffset                uint32
	HeapIndexStride              uint32
	PEmbeddedSampler             *VkSamplerCreateInfo
	UseCombinedImageSamplerIndex VkBool32
	SamplerHeapOffset            uint32
	SamplerPushOffset            uint32
	SamplerAddressOffset         uint32
	SamplerHeapIndexStride       uint32
}

type VkDescriptorMappingSourceHeapDataEXT struct {
	HeapOffset uint32
	PushOffset uint32
}

type VkDescriptorMappingSourceShaderRecordIndexEXT struct {
	HeapOffset                   uint32
	ShaderRecordOffset           uint32
	HeapIndexStride              uint32
	HeapArrayStride              uint32
	PEmbeddedSampler             *VkSamplerCreateInfo
	UseCombinedImageSamplerIndex VkBool32
	SamplerHeapOffset            uint32
	SamplerShaderRecordOffset    uint32
	SamplerHeapIndexStride       uint32
	SamplerHeapArrayStride       uint32
}

type VkDescriptorMappingSourceIndirectAddressEXT struct {
	PushOffset    uint32
	AddressOffset uint32
}

type VkDescriptorSetAndBindingMappingEXT struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	DescriptorSet uint32
	FirstBinding  uint32
	BindingCount  uint32
	ResourceMask  VkSpirvResourceTypeFlagsEXT
	Source        VkDescriptorMappingSourceEXT
	SourceData    VkDescriptorMappingSourceDataEXT
}

type VkShaderDescriptorSetAndBindingMappingInfoEXT struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	MappingCount uint32
	PMappings    *VkDescriptorSetAndBindingMappingEXT
}

type VkSamplerCustomBorderColorIndexCreateInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Index uint32
}

type VkOpaqueCaptureDataCreateInfoEXT struct {
	SType VkStructureType
	PNext unsafe.Pointer
	PData *VkHostAddressRangeConstEXT
}

type VkIndirectCommandsLayoutPushDataTokenNV struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	PushDataOffset uint32
	PushDataSize   uint32
}

type VkSubsampledImageFormatPropertiesEXT struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	SubsampledImageDescriptorCount uint32
}

type VkPhysicalDeviceDescriptorHeapFeaturesEXT struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	DescriptorHeap              VkBool32
	DescriptorHeapCaptureReplay VkBool32
}

type VkPhysicalDeviceDescriptorHeapPropertiesEXT struct {
	SType                                   VkStructureType
	PNext                                   unsafe.Pointer
	SamplerHeapAlignment                    VkDeviceSize
	ResourceHeapAlignment                   VkDeviceSize
	MaxSamplerHeapSize                      VkDeviceSize
	MaxResourceHeapSize                     VkDeviceSize
	MinSamplerHeapReservedRange             VkDeviceSize
	MinSamplerHeapReservedRangeWithEmbedded VkDeviceSize
	MinResourceHeapReservedRange            VkDeviceSize
	SamplerDescriptorSize                   VkDeviceSize
	ImageDescriptorSize                     VkDeviceSize
	BufferDescriptorSize                    VkDeviceSize
	SamplerDescriptorAlignment              VkDeviceSize
	ImageDescriptorAlignment                VkDeviceSize
	BufferDescriptorAlignment               VkDeviceSize
	MaxPushDataSize                         VkDeviceSize
	ImageCaptureReplayOpaqueDataSize        uint
	MaxDescriptorHeapEmbeddedSamplers       uint32
	SamplerYcbcrConversionCount             uint32
	SparseDescriptorHeaps                   VkBool32
	ProtectedDescriptorHeaps                VkBool32
}

type VkCommandBufferInheritanceDescriptorHeapInfoEXT struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	PSamplerHeapBindInfo  *VkBindHeapInfoEXT
	PResourceHeapBindInfo *VkBindHeapInfoEXT
}

type VkPhysicalDeviceDescriptorHeapTensorPropertiesARM struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	TensorDescriptorSize              VkDeviceSize
	TensorDescriptorAlignment         VkDeviceSize
	TensorCaptureReplayOpaqueDataSize uint
}
