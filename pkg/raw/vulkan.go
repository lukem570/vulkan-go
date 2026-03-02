package vulkan

import (
	"fmt"
	"unsafe"
)

/*
#cgo CFLAGS: -I./../../mod/volk -I./../../mod/Vulkan-Headers/include/vulkan
#define VOLK_IMPLEMENTATION
#include <stdlib.h>
#include <volk.h>
#include <vulkan.h>
*/
import "C"

// Target: GO_VULKAN

func Initialize() error {
	res := C.volkInitialize()

	if res != C.VK_SUCCESS {
		return fmt.Errorf("Failed to find Vulkan library.")
	}

	return nil
}

// Target: VK_BASE_VERSION_1_0

// Header boilerplate
// Fundamental types used by many commands and structures

type VkBool32 uint32

type VkDeviceAddress uint64

type VkDeviceSize uint64

type VkExtent2D struct {
	Width  uint32
	Height uint32
}

type VkExtent3D struct {
	Width  uint32
	Height uint32
	Depth  uint32
}

type VkFlags uint32

type VkOffset2D struct {
	X int32
	Y int32
}

type VkOffset3D struct {
	X int32
	Y int32
	Z int32
}

type VkRect2D struct {
	Offset VkOffset2D
	Extent VkExtent2D
}

// API constants
// These types are part of the API, though not directly used in API commands or data structures

type VkBaseInStructure struct {
	sType struct{} structure
 	PNext *VkBaseInStructure
}

type VkBaseOutStructure struct {
	SType VkStructureType
	PNext *VkBaseOutStructure
}

// API version macros
// Device initialization

func CreateInstance(
	pCreateInfo *VkInstanceCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pInstance *VkInstance,
) VkResult {
}

func (h *VkInstance) Destroy(
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkInstance) EnumeratePhysicalDevices(
	pPhysicalDeviceCount *uint32,
	pPhysicalDevices *VkPhysicalDevice,
) VkResult {
}

func (h *VkPhysicalDevice) GetFeatures(
	pFeatures *VkPhysicalDeviceFeatures,
) {
}

func (h *VkPhysicalDevice) GetFormatProperties(
	format VkFormat,
	pFormatProperties *VkFormatProperties,
) {
}

func (h *VkPhysicalDevice) GetImageFormatProperties(
	format VkFormat,
	Type VkImageType,
	tiling VkImageTiling,
	usage VkImageUsageFlags,
	flags VkImageCreateFlags,
	pImageFormatProperties *VkImageFormatProperties,
) VkResult {
}

func (h *VkPhysicalDevice) GetProperties(
	pProperties *VkPhysicalDeviceProperties,
) {
}

func (h *VkPhysicalDevice) GetQueueFamilyProperties(
	pQueueFamilyPropertyCount *uint32,
	pQueueFamilyProperties *VkQueueFamilyProperties,
) {
}

func (h *VkPhysicalDevice) GetMemoryProperties(
	pMemoryProperties *VkPhysicalDeviceMemoryProperties,
) {
}

func (h *VkInstance) GetProcAddr(
	pName string,
) {
}

func (h *VkDevice) GetProcAddr(
	pName string,
) {
}

type PFN_vkAllocationFunction func(
	pUserData unsafe.Pointer,
	size uint,
	alignment uint,
	allocationScope VkSystemAllocationScope,
) unsafe.Pointer

type PFN_vkFreeFunction func(
	pUserData unsafe.Pointer,
	pMemory unsafe.Pointer,
)

type PFN_vkInternalAllocationNotification func(
	pUserData unsafe.Pointer,
	size uint,
	allocationType VkInternalAllocationType,
	allocationScope VkSystemAllocationScope,
)

type PFN_vkInternalFreeNotification func(
	pUserData unsafe.Pointer,
	size uint,
	allocationType VkInternalAllocationType,
	allocationScope VkSystemAllocationScope,
)

type PFN_vkReallocationFunction func(
	pUserData unsafe.Pointer,
	pOriginal unsafe.Pointer,
	size uint,
	alignment uint,
	allocationScope VkSystemAllocationScope,
) unsafe.Pointer

type PFN_vkVoidFunction func()

type VkAllocationCallbacks struct {
	PUserData             unsafe.Pointer
	PfnAllocation         PFN_vkAllocationFunction
	PfnReallocation       PFN_vkReallocationFunction
	PfnFree               PFN_vkFreeFunction
	PfnInternalAllocation PFN_vkInternalAllocationNotification
	PfnInternalFree       PFN_vkInternalFreeNotification
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

type VkFormatFeatureFlags VkFlags

type VkFormatProperties struct {
	LinearTilingFeatures  VkFormatFeatureFlags
	OptimalTilingFeatures VkFormatFeatureFlags
	BufferFeatures        VkFormatFeatureFlags
}

type VkImageCreateFlags VkFlags

type VkImageFormatProperties struct {
	MaxExtent       VkExtent3D
	MaxMipLevels    uint32
	MaxArrayLayers  uint32
	SampleCounts    VkSampleCountFlags
	MaxResourceSize VkDeviceSize
}

type VkImageUsageFlags VkFlags

type VkInstance struct {
	handle unsafe.Pointer
}

type VkInstanceCreateFlags VkFlags

type Structure interface {
	getType() VkStructureType
	toC() *C.void
}

type VkInstanceCreateInfo struct {
	Next                   Structure
	Flags                  VkInstanceCreateFlags
	PApplicationInfo       *VkApplicationInfo
	PEnabledLayerNames     []string
	PEnabledExtensionNames []string
}

func (s *VkInstanceCreateInfo) getType() VkStructureType {
	return VkStructureTypeInstanceCreateInfo
}

func (s *VkInstanceCreateInfo) toC() (*C.void, func()) {
	p := (*C.VkInstanceCreateInfo)(C.malloc(
		C.size_t(
			unsafe.Sizeof(C.VkInstanceCreateInfo{}),
		),
	))

	cancels = make(func(), 0)

	p.sType = (*C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, c := s.Next.toC()
		cancels = append(cancels, c)
		p.pNext = next;
	}

	p.Flags = (*C.VkInstanceCreateFlags)(s.Flags)

	if s.PApplicationInfo != nil {
		p.PApplicationInfo = (*C.VkApplicationInfo)()
	}

	p.enabledLayerCount = len(s.PEnabledLayerNames)	
	p.ppEnabledLayerNames
	p.enabledExtensionCount
	p.ppEnabledExtensionNames

	return p, func() {
		C.free(unsafe.Pointer(p))

		for _, c := range cancels {
			C.free(unsafe.Pointer(c))
		}
	}
}

type VkMemoryHeap struct {
	Size  VkDeviceSize
	Flags VkMemoryHeapFlags
}

type VkMemoryHeapFlags VkFlags

type VkMemoryPropertyFlags VkFlags

type VkMemoryType struct {
	PropertyFlags VkMemoryPropertyFlags
	HeapIndex     uint32
}

type VkPhysicalDevice struct {
	handle unsafe.Pointer
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
	MaxComputeWorkGroupCount                        [3]uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         [3]uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           [2]uint32
	ViewportBoundsRange                             [2]float32
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
	PointSizeRange                                  [2]float32
	LineWidthRange                                  [2]float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     VkBool32
	StandardSampleLocations                         VkBool32
	OptimalBufferCopyOffsetAlignment                VkDeviceSize
	OptimalBufferCopyRowPitchAlignment              VkDeviceSize
	NonCoherentAtomSize                             VkDeviceSize
}

type VkPhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [VkMaxMemoryTypes]VkMemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [VkMaxMemoryHeaps]VkMemoryHeap
}

type VkPhysicalDeviceSparseProperties struct {
	ResidencyStandard2DBlockShape            VkBool32
	ResidencyStandard2DMultisampleBlockShape VkBool32
	ResidencyStandard3DBlockShape            VkBool32
	ResidencyAlignedMipSize                  VkBool32
	ResidencyNonResidentStrict               VkBool32
}

type VkPhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        VkPhysicalDeviceType
	DeviceName        [VkMaxPhysicalDeviceNameSize]uint8
	PipelineCacheUUID [VkUuidSize]uint8
	Limits            VkPhysicalDeviceLimits
	SparseProperties  VkPhysicalDeviceSparseProperties
}

type VkQueueFamilyProperties struct {
	QueueFlags                  VkQueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity VkExtent3D
}

type VkQueueFlags VkFlags

type VkSampleCountFlags VkFlags

type VkShaderStageFlags VkFlags

// Device commands

func (h *VkPhysicalDevice) CreateDevice(
	pCreateInfo *VkDeviceCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pDevice *VkDevice,
) VkResult {
}

func (h *VkDevice) Destroy(
	pAllocator *VkAllocationCallbacks,
) {
}

type VkDevice struct {
	handle unsafe.Pointer
}

type VkDeviceCreateFlags VkFlags

type VkDeviceQueueCreateFlags VkFlags

type VkDeviceQueueCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueuePriorities  []float32
}

type VkDeviceCreateInfo struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkDeviceCreateFlags
	QueueCreateInfos       []VkDeviceQueueCreateInfo
	PEnabledLayerNames     []string
	PEnabledExtensionNames []string
	PEnabledFeatures       *VkPhysicalDeviceFeatures
}

// Extension discovery commands

func EnumerateInstanceExtensionProperties(
	pLayerName string,
	pPropertyCount *uint32,
	pProperties *VkExtensionProperties,
) VkResult {
}

func (h *VkPhysicalDevice) EnumerateDeviceExtensionProperties(
	pLayerName string,
	pPropertyCount *uint32,
	pProperties *VkExtensionProperties,
) VkResult {
}

type VkExtensionProperties struct {
	ExtensionName [VkMaxExtensionNameSize]uint8
	SpecVersion   uint32
}

// Layer discovery commands

func EnumerateInstanceLayerProperties(
	pPropertyCount *uint32,
	pProperties *VkLayerProperties,
) VkResult {
}

func (h *VkPhysicalDevice) EnumerateDeviceLayerProperties(
	pPropertyCount *uint32,
	pProperties *VkLayerProperties,
) VkResult {
}

type VkLayerProperties struct {
	LayerName             [VkMaxExtensionNameSize]uint8
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [VkMaxDescriptionSize]uint8
}

// Queue commands

func (h *VkDevice) GetQueue(
	queueFamilyIndex uint32,
	queueIndex uint32,
	pQueue *VkQueue,
) {
}

func (h *VkQueue) Submit(
	submitCount uint32,
	pSubmits *VkSubmitInfo,
	fence VkFence,
) VkResult {
}

func (h *VkQueue) WaitIdle() VkResult {
}

func (h *VkDevice) WaitIdle() VkResult {
}

type VkPipelineStageFlags VkFlags

type VkQueue struct {
	handle unsafe.Pointer
}

type VkSubmitInfo struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	WaitSemaphores    []VkSemaphore
	PWaitDstStageMask *VkPipelineStageFlags
	CommandBuffers    []VkCommandBuffer
	SignalSemaphores  []VkSemaphore
}

// Memory commands

func (h *VkDevice) AllocateMemory(
	pAllocateInfo *VkMemoryAllocateInfo,
	pAllocator *VkAllocationCallbacks,
	pMemory *VkDeviceMemory,
) VkResult {
}

func (h *VkDevice) FreeMemory(
	memory VkDeviceMemory,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) MapMemory(
	memory VkDeviceMemory,
	offset VkDeviceSize,
	size VkDeviceSize,
	flags VkMemoryMapFlags,
	ppData **void,
) VkResult {
}

func (h *VkDevice) UnmapMemory(
	memory VkDeviceMemory,
) {
}

func (h *VkDevice) FlushMappedMemoryRanges(
	memoryRangeCount uint32,
	pMemoryRanges *VkMappedMemoryRange,
) VkResult {
}

func (h *VkDevice) InvalidateMappedMemoryRanges(
	memoryRangeCount uint32,
	pMemoryRanges *VkMappedMemoryRange,
) VkResult {
}

func (h *VkDevice) GetMemoryCommitment(
	memory VkDeviceMemory,
	pCommittedMemoryInBytes *VkDeviceSize,
) {
}

type VkMappedMemoryRange struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
	Offset VkDeviceSize
	Size   VkDeviceSize
}

type VkMemoryAllocateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	AllocationSize  VkDeviceSize
	MemoryTypeIndex uint32
}

type VkMemoryMapFlags VkFlags

// Memory management API commands

func (h *VkDevice) BindBufferMemory(
	buffer VkBuffer,
	memory VkDeviceMemory,
	memoryOffset VkDeviceSize,
) VkResult {
}

func (h *VkDevice) BindImageMemory(
	image VkImage,
	memory VkDeviceMemory,
	memoryOffset VkDeviceSize,
) VkResult {
}

func (h *VkDevice) GetBufferMemoryRequirements(
	buffer VkBuffer,
	pMemoryRequirements *VkMemoryRequirements,
) {
}

func (h *VkDevice) GetImageMemoryRequirements(
	image VkImage,
	pMemoryRequirements *VkMemoryRequirements,
) {
}

type VkDeviceMemory struct {
	handle unsafe.Pointer
}

type VkMemoryRequirements struct {
	Size           VkDeviceSize
	Alignment      VkDeviceSize
	MemoryTypeBits uint32
}

// Sparse resource memory management API commands (optional)

func (h *VkDevice) GetImageSparseMemoryRequirements(
	image VkImage,
	pSparseMemoryRequirementCount *uint32,
	pSparseMemoryRequirements *VkSparseImageMemoryRequirements,
) {
}

func (h *VkPhysicalDevice) GetSparseImageFormatProperties(
	format VkFormat,
	Type VkImageType,
	samples VkSampleCountFlagBits,
	usage VkImageUsageFlags,
	tiling VkImageTiling,
	pPropertyCount *uint32,
	pProperties *VkSparseImageFormatProperties,
) {
}

func (h *VkQueue) BindSparse(
	bindInfoCount uint32,
	pBindInfo *VkBindSparseInfo,
	fence VkFence,
) VkResult {
}

type VkImageAspectFlags VkFlags

type VkImageSubresource struct {
	AspectMask VkImageAspectFlags
	MipLevel   uint32
	ArrayLayer uint32
}

type VkSparseImageFormatFlags VkFlags

type VkSparseImageFormatProperties struct {
	AspectMask       VkImageAspectFlags
	ImageGranularity VkExtent3D
	Flags            VkSparseImageFormatFlags
}

type VkSparseImageMemoryBind struct {
	Subresource  VkImageSubresource
	Offset       VkOffset3D
	Extent       VkExtent3D
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
	Flags        VkSparseMemoryBindFlags
}

type VkSparseImageMemoryBindInfo struct {
	Image VkImage
	Binds []VkSparseImageMemoryBind
}

type VkSparseImageMemoryRequirements struct {
	FormatProperties     VkSparseImageFormatProperties
	ImageMipTailFirstLod uint32
	ImageMipTailSize     VkDeviceSize
	ImageMipTailOffset   VkDeviceSize
	ImageMipTailStride   VkDeviceSize
}

type VkSparseMemoryBind struct {
	ResourceOffset VkDeviceSize
	Size           VkDeviceSize
	Memory         VkDeviceMemory
	MemoryOffset   VkDeviceSize
	Flags          VkSparseMemoryBindFlags
}

type VkSparseMemoryBindFlags VkFlags

type VkSparseBufferMemoryBindInfo struct {
	Buffer VkBuffer
	Binds  []VkSparseMemoryBind
}

type VkSparseImageOpaqueMemoryBindInfo struct {
	Image VkImage
	Binds []VkSparseMemoryBind
}

type VkBindSparseInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	WaitSemaphores   []VkSemaphore
	BufferBinds      []VkSparseBufferMemoryBindInfo
	ImageOpaqueBinds []VkSparseImageOpaqueMemoryBindInfo
	ImageBinds       []VkSparseImageMemoryBindInfo
	SignalSemaphores []VkSemaphore
}

// Fence commands

func (h *VkDevice) CreateFence(
	pCreateInfo *VkFenceCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pFence *VkFence,
) VkResult {
}

func (h *VkDevice) DestroyFence(
	fence VkFence,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) ResetFences(
	fenceCount uint32,
	pFences *VkFence,
) VkResult {
}

func (h *VkDevice) GetFenceStatus(
	fence VkFence,
) VkResult {
}

func (h *VkDevice) WaitForFences(
	fenceCount uint32,
	pFences *VkFence,
	waitAll VkBool32,
	timeout uint64,
) VkResult {
}

type VkFence struct {
	handle unsafe.Pointer
}

type VkFenceCreateFlags VkFlags

type VkFenceCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkFenceCreateFlags
}

// Queue semaphore commands

func (h *VkDevice) CreateSemaphore(
	pCreateInfo *VkSemaphoreCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pSemaphore *VkSemaphore,
) VkResult {
}

func (h *VkDevice) DestroySemaphore(
	semaphore VkSemaphore,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkSemaphore struct {
	handle unsafe.Pointer
}

type VkSemaphoreCreateFlags VkFlags

type VkSemaphoreCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkSemaphoreCreateFlags
}

// Query commands

func (h *VkDevice) CreateQueryPool(
	pCreateInfo *VkQueryPoolCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pQueryPool *VkQueryPool,
) VkResult {
}

func (h *VkDevice) DestroyQueryPool(
	queryPool VkQueryPool,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) GetQueryPoolResults(
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
	dataSize uint,
	pData unsafe.Pointer,
	stride VkDeviceSize,
	flags VkQueryResultFlags,
) VkResult {
}

type VkQueryPool struct {
	handle unsafe.Pointer
}

type VkQueryPoolCreateFlags VkFlags

type VkQueryPipelineStatisticFlags VkFlags

type VkQueryPoolCreateInfo struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	Flags             VkQueryPoolCreateFlags
	QueryType         VkQueryType
	IpelineStatistics []VkQueryPipelineStatisticFlags
}

type VkQueryResultFlags VkFlags

// Buffer commands

func (h *VkDevice) CreateBuffer(
	pCreateInfo *VkBufferCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pBuffer *VkBuffer,
) VkResult {
}

func (h *VkDevice) DestroyBuffer(
	buffer VkBuffer,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkBuffer struct {
	handle unsafe.Pointer
}

type VkBufferCreateFlags VkFlags

type VkBufferCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkBufferCreateFlags
	Size               VkDeviceSize
	Usage              VkBufferUsageFlags
	SharingMode        VkSharingMode
	QueueFamilyIndices []uint32
}

type VkBufferUsageFlags VkFlags

// Image commands

func (h *VkDevice) CreateImage(
	pCreateInfo *VkImageCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pImage *VkImage,
) VkResult {
}

func (h *VkDevice) DestroyImage(
	image VkImage,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) GetImageSubresourceLayout(
	image VkImage,
	pSubresource *VkImageSubresource,
	pLayout *VkSubresourceLayout,
) {
}

type VkImage struct {
	handle unsafe.Pointer
}

type VkImageCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkImageCreateFlags
	ImageType          VkImageType
	Format             VkFormat
	Extent             VkExtent3D
	MipLevels          uint32
	ArrayLayers        uint32
	Samples            VkSampleCountFlagBits
	Tiling             VkImageTiling
	Usage              VkImageUsageFlags
	SharingMode        VkSharingMode
	QueueFamilyIndices []uint32
	InitialLayout      VkImageLayout
}

type VkSubresourceLayout struct {
	Offset     VkDeviceSize
	Size       VkDeviceSize
	RowPitch   VkDeviceSize
	ArrayPitch VkDeviceSize
	DepthPitch VkDeviceSize
}

// Image view commands

func (h *VkDevice) CreateImageView(
	pCreateInfo *VkImageViewCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pView *VkImageView,
) VkResult {
}

func (h *VkDevice) DestroyImageView(
	imageView VkImageView,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkComponentMapping struct {
	R VkComponentSwizzle
	G VkComponentSwizzle
	B VkComponentSwizzle
	A VkComponentSwizzle
}

type VkImageSubresourceRange struct {
	AspectMask     VkImageAspectFlags
	BaseMipLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type VkImageView struct {
	handle unsafe.Pointer
}

type VkImageViewCreateFlags VkFlags

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

// Pass commands

type VkAccessFlags VkFlags

type VkDependencyFlags VkFlags

// Command pool commands

func (h *VkDevice) CreateCommandPool(
	pCreateInfo *VkCommandPoolCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pCommandPool *VkCommandPool,
) VkResult {
}

func (h *VkDevice) DestroyCommandPool(
	commandPool VkCommandPool,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) ResetCommandPool(
	commandPool VkCommandPool,
	flags VkCommandPoolResetFlags,
) VkResult {
}

type VkCommandPool struct {
	handle unsafe.Pointer
}

type VkCommandPoolCreateFlags VkFlags

type VkCommandPoolCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Flags            VkCommandPoolCreateFlags
	QueueFamilyIndex uint32
}

type VkCommandPoolResetFlags VkFlags

// Command buffer commands

func (h *VkDevice) AllocateCommandBuffers(
	pAllocateInfo *VkCommandBufferAllocateInfo,
	pCommandBuffers *VkCommandBuffer,
) VkResult {
}

func (h *VkDevice) FreeCommandBuffers(
	commandPool VkCommandPool,
	commandBufferCount uint32,
	pCommandBuffers *VkCommandBuffer,
) {
}

func (h *VkCommandBuffer) Begin(
	pBeginInfo *VkCommandBufferBeginInfo,
) VkResult {
}

func (h *VkCommandBuffer) End() VkResult {
}

func (h *VkCommandBuffer) Reset(
	flags VkCommandBufferResetFlags,
) VkResult {
}

type VkCommandBuffer struct {
	handle unsafe.Pointer
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

type VkCommandBufferResetFlags VkFlags

type VkCommandBufferUsageFlags VkFlags

type VkQueryControlFlags VkFlags

// Command buffer building commands

func (h *VkCommandBuffer) CopyBuffer(
	srcBuffer VkBuffer,
	dstBuffer VkBuffer,
	regionCount uint32,
	pRegions *VkBufferCopy,
) {
}

func (h *VkCommandBuffer) CopyImage(
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkImageCopy,
) {
}

func (h *VkCommandBuffer) CopyBufferToImage(
	srcBuffer VkBuffer,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkBufferImageCopy,
) {
}

func (h *VkCommandBuffer) CopyImageToBuffer(
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstBuffer VkBuffer,
	regionCount uint32,
	pRegions *VkBufferImageCopy,
) {
}

func (h *VkCommandBuffer) UpdateBuffer(
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	dataSize VkDeviceSize,
	pData unsafe.Pointer,
) {
}

func (h *VkCommandBuffer) FillBuffer(
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	size VkDeviceSize,
	data uint32,
) {
}

func (h *VkCommandBuffer) PipelineBarrier(
	srcStageMask VkPipelineStageFlags,
	dstStageMask VkPipelineStageFlags,
	dependencyFlags VkDependencyFlags,
	memoryBarrierCount uint32,
	pMemoryBarriers *VkMemoryBarrier,
	bufferMemoryBarrierCount uint32,
	pBufferMemoryBarriers *VkBufferMemoryBarrier,
	imageMemoryBarrierCount uint32,
	pImageMemoryBarriers *VkImageMemoryBarrier,
) {
}

func (h *VkCommandBuffer) BeginQuery(
	queryPool VkQueryPool,
	query uint32,
	flags VkQueryControlFlags,
) {
}

func (h *VkCommandBuffer) EndQuery(
	queryPool VkQueryPool,
	query uint32,
) {
}

func (h *VkCommandBuffer) ResetQueryPool(
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
) {
}

func (h *VkCommandBuffer) WriteTimestamp(
	pipelineStage VkPipelineStageFlagBits,
	queryPool VkQueryPool,
	query uint32,
) {
}

func (h *VkCommandBuffer) CopyQueryPoolResults(
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
	dstBuffer VkBuffer,
	dstOffset VkDeviceSize,
	stride VkDeviceSize,
	flags VkQueryResultFlags,
) {
}

func (h *VkCommandBuffer) ExecuteCommands(
	commandBufferCount uint32,
	pCommandBuffers *VkCommandBuffer,
) {
}

type VkBufferCopy struct {
	SrcOffset VkDeviceSize
	DstOffset VkDeviceSize
	Size      VkDeviceSize
}

type VkImageSubresourceLayers struct {
	AspectMask     VkImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type VkBufferImageCopy struct {
	BufferOffset      VkDeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  VkImageSubresourceLayers
	ImageOffset       VkOffset3D
	ImageExtent       VkExtent3D
}

type VkImageCopy struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
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

type VkMemoryBarrier struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SrcAccessMask VkAccessFlags
	DstAccessMask VkAccessFlags
}

// Retroactively promoted from VK_EXT_debug_report for compatibility with VulkanSC
// Target: VK_COMPUTE_VERSION_1_0

// These types are part of the API, though not directly used in API commands or data structures

type VkDispatchIndirectCommand struct {
	X uint32
	Y uint32
	Z uint32
}

type VkPipelineCacheHeaderVersionOne struct {
	HeaderSize        uint32
	HeaderVersion     VkPipelineCacheHeaderVersion
	VendorID          uint32
	DeviceID          uint32
	PipelineCacheUUID [VkUuidSize]uint8
}

// Event commands

func (h *VkDevice) CreateEvent(
	pCreateInfo *VkEventCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pEvent *VkEvent,
) VkResult {
}

func (h *VkDevice) DestroyEvent(
	event VkEvent,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) GetEventStatus(
	event VkEvent,
) VkResult {
}

func (h *VkDevice) SetEvent(
	event VkEvent,
) VkResult {
}

func (h *VkDevice) ResetEvent(
	event VkEvent,
) VkResult {
}

type VkEvent struct {
	handle unsafe.Pointer
}

type VkEventCreateFlags VkFlags

type VkEventCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkEventCreateFlags
}

// Buffer view commands

func (h *VkDevice) CreateBufferView(
	pCreateInfo *VkBufferViewCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pView *VkBufferView,
) VkResult {
}

func (h *VkDevice) DestroyBufferView(
	bufferView VkBufferView,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkBufferView struct {
	handle unsafe.Pointer
}

type VkBufferViewCreateFlags VkFlags

type VkBufferViewCreateInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkBufferViewCreateFlags
	Buffer VkBuffer
	Format VkFormat
	Offset VkDeviceSize
	Range  VkDeviceSize
}

// Shader commands

func (h *VkDevice) CreateShaderModule(
	pCreateInfo *VkShaderModuleCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pShaderModule *VkShaderModule,
) VkResult {
}

func (h *VkDevice) DestroyShaderModule(
	shaderModule VkShaderModule,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkShaderModule struct {
	handle unsafe.Pointer
}

type VkShaderModuleCreateFlags VkFlags

type VkShaderModuleCreateInfo struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Flags    VkShaderModuleCreateFlags
	CodeSize uint
	PCode    *uint32
}

// Pipeline Cache commands

func (h *VkDevice) CreatePipelineCache(
	pCreateInfo *VkPipelineCacheCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pPipelineCache *VkPipelineCache,
) VkResult {
}

func (h *VkDevice) DestroyPipelineCache(
	pipelineCache VkPipelineCache,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) GetPipelineCacheData(
	pipelineCache VkPipelineCache,
	pDataSize *uint,
	pData unsafe.Pointer,
) VkResult {
}

func (h *VkDevice) MergePipelineCaches(
	dstCache VkPipelineCache,
	srcCacheCount uint32,
	pSrcCaches *VkPipelineCache,
) VkResult {
}

type VkPipelineCache struct {
	handle unsafe.Pointer
}

type VkPipelineCacheCreateFlags VkFlags

type VkPipelineCacheCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	Flags           VkPipelineCacheCreateFlags
	InitialDataSize uint
	PInitialData    unsafe.Pointer
}

// Compute Pipeline commands

func (h *VkDevice) CreateComputePipelines(
	pipelineCache VkPipelineCache,
	createInfoCount uint32,
	pCreateInfos *VkComputePipelineCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pPipelines *VkPipeline,
) VkResult {
}

func (h *VkDevice) DestroyPipeline(
	pipeline VkPipeline,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkPipeline struct {
	handle unsafe.Pointer
}

type VkPipelineCreateFlags VkFlags

type VkPipelineLayoutCreateFlags VkFlags

type VkPipelineShaderStageCreateFlags VkFlags

type VkSpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       uint
}

type VkSpecializationInfo struct {
	MapEntries []VkSpecializationMapEntry
	DataSize   uint
	PData      unsafe.Pointer
}

type VkPipelineShaderStageCreateInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkPipelineShaderStageCreateFlags
	Stage               VkShaderStageFlagBits
	Module              VkShaderModule
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

// Pipeline layout commands

func (h *VkDevice) CreatePipelineLayout(
	pCreateInfo *VkPipelineLayoutCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pPipelineLayout *VkPipelineLayout,
) VkResult {
}

func (h *VkDevice) DestroyPipelineLayout(
	pipelineLayout VkPipelineLayout,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkPipelineLayout struct {
	handle unsafe.Pointer
}

type VkPushConstantRange struct {
	StageFlags VkShaderStageFlags
	Offset     uint32
	Size       uint32
}

type VkPipelineLayoutCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkPipelineLayoutCreateFlags
	SetLayouts         []VkDescriptorSetLayout
	PushConstantRanges []VkPushConstantRange
}

// Sampler commands

func (h *VkDevice) CreateSampler(
	pCreateInfo *VkSamplerCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pSampler *VkSampler,
) VkResult {
}

func (h *VkDevice) DestroySampler(
	sampler VkSampler,
	pAllocator *VkAllocationCallbacks,
) {
}

type VkSampler struct {
	handle unsafe.Pointer
}

type VkSamplerCreateFlags VkFlags

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

// Descriptor set commands

func (h *VkDevice) CreateDescriptorSetLayout(
	pCreateInfo *VkDescriptorSetLayoutCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pSetLayout *VkDescriptorSetLayout,
) VkResult {
}

func (h *VkDevice) DestroyDescriptorSetLayout(
	descriptorSetLayout VkDescriptorSetLayout,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) CreateDescriptorPool(
	pCreateInfo *VkDescriptorPoolCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pDescriptorPool *VkDescriptorPool,
) VkResult {
}

func (h *VkDevice) DestroyDescriptorPool(
	descriptorPool VkDescriptorPool,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) ResetDescriptorPool(
	descriptorPool VkDescriptorPool,
	flags VkDescriptorPoolResetFlags,
) VkResult {
}

func (h *VkDevice) AllocateDescriptorSets(
	pAllocateInfo *VkDescriptorSetAllocateInfo,
	pDescriptorSets *VkDescriptorSet,
) VkResult {
}

func (h *VkDevice) FreeDescriptorSets(
	descriptorPool VkDescriptorPool,
	descriptorSetCount uint32,
	pDescriptorSets *VkDescriptorSet,
) VkResult {
}

func (h *VkDevice) UpdateDescriptorSets(
	descriptorWriteCount uint32,
	pDescriptorWrites *VkWriteDescriptorSet,
	descriptorCopyCount uint32,
	pDescriptorCopies *VkCopyDescriptorSet,
) {
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

type VkDescriptorPool struct {
	handle unsafe.Pointer
}

type VkDescriptorPoolSize struct {
	Type            VkDescriptorType
	DescriptorCount uint32
}

type VkDescriptorPoolCreateFlags VkFlags

type VkDescriptorPoolCreateInfo struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Flags     VkDescriptorPoolCreateFlags
	MaxSets   uint32
	PoolSizes []VkDescriptorPoolSize
}

type VkDescriptorPoolResetFlags VkFlags

type VkDescriptorSet struct {
	handle unsafe.Pointer
}

type VkDescriptorSetAllocateInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	DescriptorPool VkDescriptorPool
	SetLayouts     []VkDescriptorSetLayout
}

type VkDescriptorSetLayout struct {
	handle unsafe.Pointer
}

type VkDescriptorSetLayoutBinding struct {
	Binding            uint32
	DescriptorType     VkDescriptorType
	DescriptorCount    uint32
	StageFlags         VkShaderStageFlags
	PImmutableSamplers *VkSampler
}

type VkDescriptorSetLayoutCreateFlags VkFlags

type VkDescriptorSetLayoutCreateInfo struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Flags    VkDescriptorSetLayoutCreateFlags
	Bindings []VkDescriptorSetLayoutBinding
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

// Pass commands
// Command buffer building commands

func (h *VkCommandBuffer) BindPipeline(
	pipelineBindPoint VkPipelineBindPoint,
	pipeline VkPipeline,
) {
}

func (h *VkCommandBuffer) BindDescriptorSets(
	pipelineBindPoint VkPipelineBindPoint,
	layout VkPipelineLayout,
	firstSet uint32,
	descriptorSetCount uint32,
	pDescriptorSets *VkDescriptorSet,
	dynamicOffsetCount uint32,
	pDynamicOffsets *uint32,
) {
}

func (h *VkCommandBuffer) ClearColorImage(
	image VkImage,
	imageLayout VkImageLayout,
	pColor *VkClearColorValue,
	rangeCount uint32,
	pRanges *VkImageSubresourceRange,
) {
}

func (h *VkCommandBuffer) Dispatch(
	groupCountX uint32,
	groupCountY uint32,
	groupCountZ uint32,
) {
}

func (h *VkCommandBuffer) DispatchIndirect(
	buffer VkBuffer,
	offset VkDeviceSize,
) {
}

func (h *VkCommandBuffer) SetEvent(
	event VkEvent,
	stageMask VkPipelineStageFlags,
) {
}

func (h *VkCommandBuffer) ResetEvent(
	event VkEvent,
	stageMask VkPipelineStageFlags,
) {
}

func (h *VkCommandBuffer) WaitEvents(
	eventCount uint32,
	pEvents *VkEvent,
	srcStageMask VkPipelineStageFlags,
	dstStageMask VkPipelineStageFlags,
	memoryBarrierCount uint32,
	pMemoryBarriers *VkMemoryBarrier,
	bufferMemoryBarrierCount uint32,
	pBufferMemoryBarriers *VkBufferMemoryBarrier,
	imageMemoryBarrierCount uint32,
	pImageMemoryBarriers *VkImageMemoryBarrier,
) {
}

func (h *VkCommandBuffer) PushConstants(
	layout VkPipelineLayout,
	stageFlags VkShaderStageFlags,
	offset uint32,
	size uint32,
	pValues unsafe.Pointer,
) {
}

type VkClearColorValue struct {
	Float32 [4]float32
	Int32   [4]int32
	Uint32  [4]uint32
}

// Target: VK_GRAPHICS_VERSION_1_0

// API constants
// These types are part of the API, though not directly used in API commands or data structures

type VkDrawIndexedIndirectCommand struct {
	IndexCount    uint32
	InstanceCount uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}

type VkDrawIndirectCommand struct {
	VertexCount   uint32
	InstanceCount uint32
	FirstVertex   uint32
	FirstInstance uint32
}

// Graphics Pipeline commands

func (h *VkDevice) CreateGraphicsPipelines(
	pipelineCache VkPipelineCache,
	createInfoCount uint32,
	pCreateInfos *VkGraphicsPipelineCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pPipelines *VkPipeline,
) VkResult {
}

type VkColorComponentFlags VkFlags

type VkCullModeFlags VkFlags

type VkStencilOpState struct {
	FailOp      VkStencilOp
	PassOp      VkStencilOp
	DepthFailOp VkStencilOp
	CompareOp   VkCompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

type VkVertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   VkFormat
	Offset   uint32
}

type VkVertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VkVertexInputRate
}

type VkViewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
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

type VkPipelineColorBlendStateCreateFlags VkFlags

type VkPipelineColorBlendStateCreateInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkPipelineColorBlendStateCreateFlags
	LogicOpEnable  VkBool32
	LogicOp        VkLogicOp
	Attachments    []VkPipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

type VkPipelineDepthStencilStateCreateFlags VkFlags

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

type VkPipelineDynamicStateCreateFlags VkFlags

type VkPipelineDynamicStateCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	Flags         VkPipelineDynamicStateCreateFlags
	DynamicStates []VkDynamicState
}

type VkPipelineInputAssemblyStateCreateFlags VkFlags

type VkPipelineInputAssemblyStateCreateInfo struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	Flags                  VkPipelineInputAssemblyStateCreateFlags
	Topology               VkPrimitiveTopology
	PrimitiveRestartEnable VkBool32
}

type VkPipelineMultisampleStateCreateFlags VkFlags

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

type VkPipelineRasterizationStateCreateFlags VkFlags

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

type VkPipelineTessellationStateCreateFlags VkFlags

type VkPipelineTessellationStateCreateInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkPipelineTessellationStateCreateFlags
	PatchControlPoints uint32
}

type VkPipelineVertexInputStateCreateFlags VkFlags

type VkPipelineVertexInputStateCreateInfo struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	Flags                       VkPipelineVertexInputStateCreateFlags
	VertexBindingDescriptions   []VkVertexInputBindingDescription
	VertexAttributeDescriptions []VkVertexInputAttributeDescription
}

type VkPipelineViewportStateCreateFlags VkFlags

type VkPipelineViewportStateCreateInfo struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Flags     VkPipelineViewportStateCreateFlags
	Viewports []VkViewport
	Scissors  []VkRect2D
}

type VkSampleMask uint32

type VkGraphicsPipelineCreateInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkPipelineCreateFlags
	Stages              []VkPipelineShaderStageCreateInfo
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

// Pass commands

func (h *VkDevice) CreateFramebuffer(
	pCreateInfo *VkFramebufferCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pFramebuffer *VkFramebuffer,
) VkResult {
}

func (h *VkDevice) DestroyFramebuffer(
	framebuffer VkFramebuffer,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) CreateRenderPass(
	pCreateInfo *VkRenderPassCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pRenderPass *VkRenderPass,
) VkResult {
}

func (h *VkDevice) DestroyRenderPass(
	renderPass VkRenderPass,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) GetRenderAreaGranularity(
	renderPass VkRenderPass,
	pGranularity *VkExtent2D,
) {
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

type VkAttachmentDescriptionFlags VkFlags

type VkAttachmentReference struct {
	Attachment uint32
	Layout     VkImageLayout
}

type VkFramebuffer struct {
	handle unsafe.Pointer
}

type VkFramebufferCreateFlags VkFlags

type VkFramebufferCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkFramebufferCreateFlags
	RenderPass  VkRenderPass
	Attachments []VkImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

type VkRenderPass struct {
	handle unsafe.Pointer
}

type VkRenderPassCreateFlags VkFlags

type VkSubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    VkPipelineStageFlags
	DstStageMask    VkPipelineStageFlags
	SrcAccessMask   VkAccessFlags
	DstAccessMask   VkAccessFlags
	DependencyFlags VkDependencyFlags
}

type VkSubpassDescription struct {
	Flags                   VkSubpassDescriptionFlags
	PipelineBindPoint       VkPipelineBindPoint
	InputAttachments        []VkAttachmentReference
	ColorAttachments        []VkAttachmentReference
	PResolveAttachments     *VkAttachmentReference
	PDepthStencilAttachment *VkAttachmentReference
	PreserveAttachments     []uint32
}

type VkSubpassDescriptionFlags VkFlags

type VkRenderPassCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Flags        VkRenderPassCreateFlags
	Attachments  []VkAttachmentDescription
	Subpasses    []VkSubpassDescription
	Dependencies []VkSubpassDependency
}

// Command buffer building commands

func (h *VkCommandBuffer) SetViewport(
	firstViewport uint32,
	viewportCount uint32,
	pViewports *VkViewport,
) {
}

func (h *VkCommandBuffer) SetScissor(
	firstScissor uint32,
	scissorCount uint32,
	pScissors *VkRect2D,
) {
}

func (h *VkCommandBuffer) SetLineWidth(
	lineWidth float32,
) {
}

func (h *VkCommandBuffer) SetDepthBias(
	depthBiasConstantFactor float32,
	depthBiasClamp float32,
	depthBiasSlopeFactor float32,
) {
}

func (h *VkCommandBuffer) SetBlendConstants(
	blendConstants [4]float32,
) {
}

func (h *VkCommandBuffer) SetDepthBounds(
	minDepthBounds float32,
	maxDepthBounds float32,
) {
}

func (h *VkCommandBuffer) SetStencilCompareMask(
	faceMask VkStencilFaceFlags,
	compareMask uint32,
) {
}

func (h *VkCommandBuffer) SetStencilWriteMask(
	faceMask VkStencilFaceFlags,
	writeMask uint32,
) {
}

func (h *VkCommandBuffer) SetStencilReference(
	faceMask VkStencilFaceFlags,
	reference uint32,
) {
}

func (h *VkCommandBuffer) BindIndexBuffer(
	buffer VkBuffer,
	offset VkDeviceSize,
	indexType VkIndexType,
) {
}

func (h *VkCommandBuffer) BindVertexBuffers(
	firstBinding uint32,
	bindingCount uint32,
	pBuffers *VkBuffer,
	pOffsets *VkDeviceSize,
) {
}

func (h *VkCommandBuffer) Draw(
	vertexCount uint32,
	instanceCount uint32,
	firstVertex uint32,
	firstInstance uint32,
) {
}

func (h *VkCommandBuffer) DrawIndexed(
	indexCount uint32,
	instanceCount uint32,
	firstIndex uint32,
	vertexOffset int32,
	firstInstance uint32,
) {
}

func (h *VkCommandBuffer) DrawIndirect(
	buffer VkBuffer,
	offset VkDeviceSize,
	drawCount uint32,
	stride uint32,
) {
}

func (h *VkCommandBuffer) DrawIndexedIndirect(
	buffer VkBuffer,
	offset VkDeviceSize,
	drawCount uint32,
	stride uint32,
) {
}

func (h *VkCommandBuffer) BlitImage(
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkImageBlit,
	filter VkFilter,
) {
}

func (h *VkCommandBuffer) ClearDepthStencilImage(
	image VkImage,
	imageLayout VkImageLayout,
	pDepthStencil *VkClearDepthStencilValue,
	rangeCount uint32,
	pRanges *VkImageSubresourceRange,
) {
}

func (h *VkCommandBuffer) ClearAttachments(
	attachmentCount uint32,
	pAttachments *VkClearAttachment,
	rectCount uint32,
	pRects *VkClearRect,
) {
}

func (h *VkCommandBuffer) ResolveImage(
	srcImage VkImage,
	srcImageLayout VkImageLayout,
	dstImage VkImage,
	dstImageLayout VkImageLayout,
	regionCount uint32,
	pRegions *VkImageResolve,
) {
}

func (h *VkCommandBuffer) BeginRenderPass(
	pRenderPassBegin *VkRenderPassBeginInfo,
	contents VkSubpassContents,
) {
}

func (h *VkCommandBuffer) NextSubpass(
	contents VkSubpassContents,
) {
}

func (h *VkCommandBuffer) EndRenderPass() {
}

type VkClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}

type VkClearRect struct {
	Rect           VkRect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

type VkClearValue struct {
	Color        VkClearColorValue
	DepthStencil VkClearDepthStencilValue
}

type VkClearAttachment struct {
	AspectMask      VkImageAspectFlags
	ColorAttachment uint32
	ClearValue      VkClearValue
}

type VkImageBlit struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffsets     [2]VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffsets     [2]VkOffset3D
}

type VkImageResolve struct {
	SrcSubresource VkImageSubresourceLayers
	SrcOffset      VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffset      VkOffset3D
	Extent         VkExtent3D
}

type VkRenderPassBeginInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	RenderPass  VkRenderPass
	Framebuffer VkFramebuffer
	RenderArea  VkRect2D
	ClearValues []VkClearValue
}

type VkStencilFaceFlags VkFlags

// Target: VK_VERSION_1_0

// Feature requirements
// Target: VK_BASE_VERSION_1_1

// API version macros
// Device Initialization

func EnumerateInstanceVersion(
	pApiVersion *uint32,
) VkResult {
}

type VkSubgroupFeatureFlags VkFlags

// Promoted from VK_KHR_bind_memory2

func (h *VkDevice) BindBufferMemory2(
	bindInfoCount uint32,
	pBindInfos *VkBindBufferMemoryInfo,
) VkResult {
}

func (h *VkDevice) BindImageMemory2(
	bindInfoCount uint32,
	pBindInfos *VkBindImageMemoryInfo,
) VkResult {
}

type VkBindBufferMemoryInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Buffer       VkBuffer
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}

type VkBindImageMemoryInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Image        VkImage
	Memory       VkDeviceMemory
	MemoryOffset VkDeviceSize
}

// Promoted from VK_KHR_dedicated_allocation

type VkMemoryDedicatedRequirements struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PrefersDedicatedAllocation  VkBool32
	RequiresDedicatedAllocation VkBool32
}

type VkMemoryDedicatedAllocateInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Image  VkImage
	Buffer VkBuffer
}

// Promoted from VK_KHR_device_group

func (h *VkDevice) GetGroupPeerMemoryFeatures(
	heapIndex uint32,
	localDeviceIndex uint32,
	remoteDeviceIndex uint32,
	pPeerMemoryFeatures *VkPeerMemoryFeatureFlags,
) {
}

func (h *VkCommandBuffer) SetDeviceMask(
	deviceMask uint32,
) {
}

type VkPeerMemoryFeatureFlags VkFlags

type VkMemoryAllocateFlags VkFlags

type VkMemoryAllocateFlagsInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkMemoryAllocateFlags
	DeviceMask uint32
}

type VkDeviceGroupCommandBufferBeginInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	DeviceMask uint32
}

type VkDeviceGroupSubmitInfo struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	WaitSemaphoreDeviceIndices   []uint32
	CommandBufferDeviceMasks     []uint32
	SignalSemaphoreDeviceIndices []uint32
}

type VkDeviceGroupBindSparseInfo struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	ResourceDeviceIndex uint32
	MemoryDeviceIndex   uint32
}

// Promoted from VK_KHR_device_group + VK_KHR_bind_memory2

type VkBindBufferMemoryDeviceGroupInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	DeviceIndices []uint32
}

type VkBindImageMemoryDeviceGroupInfo struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DeviceIndices            []uint32
	SplitInstanceBindRegions []VkRect2D
}

// Promoted from VK_KHR_device_group_creation

func (h *VkInstance) EnumeratePhysicalDeviceGroups(
	pPhysicalDeviceGroupCount *uint32,
	pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupProperties,
) VkResult {
}

type VkPhysicalDeviceGroupProperties struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	HysicalDevices   [][VkMaxDeviceGroupSize]VkPhysicalDevice
	SubsetAllocation VkBool32
}

type VkDeviceGroupDeviceCreateInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	PhysicalDevices []VkPhysicalDevice
}

// Promoted from VK_KHR_get_memory_requirements2

func (h *VkDevice) GetImageMemoryRequirements2(
	pInfo *VkImageMemoryRequirementsInfo2,
	pMemoryRequirements *VkMemoryRequirements2,
) {
}

func (h *VkDevice) GetBufferMemoryRequirements2(
	pInfo *VkBufferMemoryRequirementsInfo2,
	pMemoryRequirements *VkMemoryRequirements2,
) {
}

func (h *VkDevice) GetImageSparseMemoryRequirements2(
	pInfo *VkImageSparseMemoryRequirementsInfo2,
	pSparseMemoryRequirementCount *uint32,
	pSparseMemoryRequirements *VkSparseImageMemoryRequirements2,
) {
}

type VkBufferMemoryRequirementsInfo2 struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
}

type VkImageMemoryRequirementsInfo2 struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Image VkImage
}

type VkImageSparseMemoryRequirementsInfo2 struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Image VkImage
}

type VkMemoryRequirements2 struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MemoryRequirements VkMemoryRequirements
}

type VkSparseImageMemoryRequirements2 struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MemoryRequirements VkSparseImageMemoryRequirements
}

// Promoted from VK_KHR_get_physical_device_properties2

func (h *VkPhysicalDevice) GetFeatures2(
	pFeatures *VkPhysicalDeviceFeatures2,
) {
}

func (h *VkPhysicalDevice) GetProperties2(
	pProperties *VkPhysicalDeviceProperties2,
) {
}

func (h *VkPhysicalDevice) GetFormatProperties2(
	format VkFormat,
	pFormatProperties *VkFormatProperties2,
) {
}

func (h *VkPhysicalDevice) GetImageFormatProperties2(
	pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2,
	pImageFormatProperties *VkImageFormatProperties2,
) VkResult {
}

func (h *VkPhysicalDevice) GetQueueFamilyProperties2(
	pQueueFamilyPropertyCount *uint32,
	pQueueFamilyProperties *VkQueueFamilyProperties2,
) {
}

func (h *VkPhysicalDevice) GetMemoryProperties2(
	pMemoryProperties *VkPhysicalDeviceMemoryProperties2,
) {
}

func (h *VkPhysicalDevice) GetSparseImageFormatProperties2(
	pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2,
	pPropertyCount *uint32,
	pProperties *VkSparseImageFormatProperties2,
) {
}

type VkPhysicalDeviceFeatures2 struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Features VkPhysicalDeviceFeatures
}

type VkPhysicalDeviceProperties2 struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Properties VkPhysicalDeviceProperties
}

type VkFormatProperties2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	FormatProperties VkFormatProperties
}

type VkImageFormatProperties2 struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	ImageFormatProperties VkImageFormatProperties
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

type VkQueueFamilyProperties2 struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	QueueFamilyProperties VkQueueFamilyProperties
}

type VkPhysicalDeviceMemoryProperties2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	MemoryProperties VkPhysicalDeviceMemoryProperties
}

type VkSparseImageFormatProperties2 struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Properties VkSparseImageFormatProperties
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

// Promoted from VK_KHR_maintenance1

func (h *VkDevice) TrimCommandPool(
	commandPool VkCommandPool,
	flags VkCommandPoolTrimFlags,
) {
}

type VkCommandPoolTrimFlags VkFlags

// Promoted from VK_KHR_maintenance2

type VkImageViewUsageCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkImageUsageFlags
}

// Originally based on VK_KHR_protected_memory (extension 146), which was never published; thus the mystifying large value= numbers below. These are not aliased since they were not actually promoted from an extension.

func (h *VkDevice) GetQueue2(
	pQueueInfo *VkDeviceQueueInfo2,
	pQueue *VkQueue,
) {
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

type VkProtectedSubmitInfo struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	ProtectedSubmit VkBool32
}

// Promoted from VK_KHR_sampler_ycbcr_conversion

type VkBindImagePlaneMemoryInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PlaneAspect VkImageAspectFlagBits
}

type VkImagePlaneMemoryRequirementsInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PlaneAspect VkImageAspectFlagBits
}

// Promoted from VK_KHR_external_memory_capabilities

func (h *VkPhysicalDevice) GetExternalBufferProperties(
	pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfo,
	pExternalBufferProperties *VkExternalBufferProperties,
) {
}

type VkExternalMemoryHandleTypeFlags VkFlags

type VkExternalMemoryFeatureFlags VkFlags

type VkExternalMemoryProperties struct {
	ExternalMemoryFeatures        VkExternalMemoryFeatureFlags
	ExportFromImportedHandleTypes VkExternalMemoryHandleTypeFlags
	CompatibleHandleTypes         VkExternalMemoryHandleTypeFlags
}

type VkPhysicalDeviceExternalImageFormatInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkExternalImageFormatProperties struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties VkExternalMemoryProperties
}

type VkPhysicalDeviceExternalBufferInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkBufferCreateFlags
	Usage      VkBufferUsageFlags
	HandleType VkExternalMemoryHandleTypeFlagBits
}

type VkExternalBufferProperties struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ExternalMemoryProperties VkExternalMemoryProperties
}

type VkPhysicalDeviceIDProperties struct {
	SType           VkStructureType
	PNext           unsafe.Pointer
	DeviceUUID      [VkUuidSize]uint8
	DriverUUID      [VkUuidSize]uint8
	DeviceLUID      [VkLuidSize]uint8
	DeviceNodeMask  uint32
	DeviceLUIDValid VkBool32
}

// Promoted from VK_KHR_external_memory

type VkExternalMemoryImageCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

type VkExternalMemoryBufferCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

type VkExportMemoryAllocateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalMemoryHandleTypeFlags
}

// Promoted from VK_KHR_external_fence_capabilities

func (h *VkPhysicalDevice) GetExternalFenceProperties(
	pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfo,
	pExternalFenceProperties *VkExternalFenceProperties,
) {
}

type VkExternalFenceHandleTypeFlags VkFlags

type VkExternalFenceFeatureFlags VkFlags

type VkPhysicalDeviceExternalFenceInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalFenceHandleTypeFlagBits
}

type VkExternalFenceProperties struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ExportFromImportedHandleTypes VkExternalFenceHandleTypeFlags
	CompatibleHandleTypes         VkExternalFenceHandleTypeFlags
	ExternalFenceFeatures         VkExternalFenceFeatureFlags
}

// Promoted from VK_KHR_external_fence

type VkFenceImportFlags VkFlags

type VkExportFenceCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalFenceHandleTypeFlags
}

// Promoted from VK_KHR_external_semaphore

type VkSemaphoreImportFlags VkFlags

type VkExportSemaphoreCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	HandleTypes VkExternalSemaphoreHandleTypeFlags
}

// Promoted from VK_KHR_external_semaphore_capabilities

func (h *VkPhysicalDevice) GetExternalSemaphoreProperties(
	pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfo,
	pExternalSemaphoreProperties *VkExternalSemaphoreProperties,
) {
}

type VkExternalSemaphoreHandleTypeFlags VkFlags

type VkExternalSemaphoreFeatureFlags VkFlags

type VkPhysicalDeviceExternalSemaphoreInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	HandleType VkExternalSemaphoreHandleTypeFlagBits
}

type VkExternalSemaphoreProperties struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ExportFromImportedHandleTypes VkExternalSemaphoreHandleTypeFlags
	CompatibleHandleTypes         VkExternalSemaphoreHandleTypeFlags
	ExternalSemaphoreFeatures     VkExternalSemaphoreFeatureFlags
}

// Target: VK_COMPUTE_VERSION_1_1

// Promoted from VK_KHR_relaxed_block_layout, which has no API
// Promoted from VK_KHR_storage_buffer_storage_class, which has no API
// Originally based on VK_KHR_subgroup (extension 94), but the actual enum block used was, incorrectly, that of extension 95

type VkPhysicalDeviceSubgroupProperties struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	SubgroupSize              uint32
	SupportedStages           VkShaderStageFlags
	SupportedOperations       VkSubgroupFeatureFlags
	QuadOperationsInAllStages VkBool32
}

// Promoted from VK_KHR_16bit_storage

type VkPhysicalDevice16BitStorageFeatures struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	StorageBuffer16BitAccess           VkBool32
	UniformAndStorageBuffer16BitAccess VkBool32
	StoragePushConstant16              VkBool32
	StorageInputOutput16               VkBool32
}

// Promoted from VK_KHR_device_group

func (h *VkCommandBuffer) DispatchBase(
	baseGroupX uint32,
	baseGroupY uint32,
	baseGroupZ uint32,
	groupCountX uint32,
	groupCountY uint32,
	groupCountZ uint32,
) {
}

// Promoted from VK_KHR_variable_pointers

type VkPhysicalDeviceVariablePointersFeatures struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	VariablePointersStorageBuffer VkBool32
	VariablePointers              VkBool32
}

type VkPhysicalDeviceVariablePointerFeatures struct {
}

// Promoted from VK_KHR_descriptor_update_template

func (h *VkDevice) CreateDescriptorUpdateTemplate(
	pCreateInfo *VkDescriptorUpdateTemplateCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pDescriptorUpdateTemplate *VkDescriptorUpdateTemplate,
) VkResult {
}

func (h *VkDevice) DestroyDescriptorUpdateTemplate(
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) UpdateDescriptorSetWithTemplate(
	descriptorSet VkDescriptorSet,
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	pData unsafe.Pointer,
) {
}

type VkDescriptorUpdateTemplate struct {
	handle unsafe.Pointer
}

type VkDescriptorUpdateTemplateCreateFlags VkFlags

type VkDescriptorUpdateTemplateEntry struct {
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  VkDescriptorType
	Offset          uint
	Stride          uint
}

type VkDescriptorUpdateTemplateCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkDescriptorUpdateTemplateCreateFlags
	DescriptorUpdateEntries []VkDescriptorUpdateTemplateEntry
	TemplateType            VkDescriptorUpdateTemplateType
	DescriptorSetLayout     VkDescriptorSetLayout
	PipelineBindPoint       VkPipelineBindPoint
	PipelineLayout          VkPipelineLayout
	Set                     uint32
}

// Promoted from VK_KHR_maintenance3

func (h *VkDevice) GetDescriptorSetLayoutSupport(
	pCreateInfo *VkDescriptorSetLayoutCreateInfo,
	pSupport *VkDescriptorSetLayoutSupport,
) {
}

type VkPhysicalDeviceMaintenance3Properties struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	MaxPerSetDescriptors    uint32
	MaxMemoryAllocationSize VkDeviceSize
}

type VkDescriptorSetLayoutSupport struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Supported VkBool32
}

// Promoted from VK_KHR_sampler_ycbcr_conversion

func (h *VkDevice) CreateSamplerYcbcrConversion(
	pCreateInfo *VkSamplerYcbcrConversionCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pYcbcrConversion *VkSamplerYcbcrConversion,
) VkResult {
}

func (h *VkDevice) DestroySamplerYcbcrConversion(
	ycbcrConversion VkSamplerYcbcrConversion,
	pAllocator *VkAllocationCallbacks,
) {
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

type VkSamplerYcbcrConversionInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Conversion VkSamplerYcbcrConversion
}

type VkPhysicalDeviceSamplerYcbcrConversionFeatures struct {
	SType                  VkStructureType
	PNext                  unsafe.Pointer
	SamplerYcbcrConversion VkBool32
}

type VkSamplerYcbcrConversionImageFormatProperties struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	CombinedImageSamplerDescriptorCount uint32
}

type VkSamplerYcbcrConversion struct {
	handle unsafe.Pointer
}

// Target: VK_GRAPHICS_VERSION_1_1

// Promoted from VK_KHR_device_group

type VkDeviceGroupRenderPassBeginInfo struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	DeviceMask        uint32
	DeviceRenderAreas []VkRect2D
}

// Promoted from VK_KHR_maintenance2

type VkPhysicalDevicePointClippingProperties struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	PointClippingBehavior VkPointClippingBehavior
}

type VkInputAttachmentAspectReference struct {
	Subpass              uint32
	InputAttachmentIndex uint32
	AspectMask           VkImageAspectFlags
}

type VkRenderPassInputAttachmentAspectCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	AspectReferences []VkInputAttachmentAspectReference
}

type VkPipelineTessellationDomainOriginStateCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	DomainOrigin VkTessellationDomainOrigin
}

// Promoted from VK_KHR_multiview

type VkRenderPassMultiviewCreateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ViewMasks        []uint32
	ViewOffsets      []int32
	CorrelationMasks []uint32
}

type VkPhysicalDeviceMultiviewFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	Multiview                   VkBool32
	MultiviewGeometryShader     VkBool32
	MultiviewTessellationShader VkBool32
}

type VkPhysicalDeviceMultiviewProperties struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	MaxMultiviewViewCount     uint32
	MaxMultiviewInstanceIndex uint32
}

// Promoted from VK_KHR_shader_draw_parameters, with a feature support query added

type VkPhysicalDeviceShaderDrawParametersFeatures struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ShaderDrawParameters VkBool32
}

type VkPhysicalDeviceShaderDrawParameterFeatures struct {
}

// Target: VK_VERSION_1_1

// Feature requirements
// Target: VK_BASE_VERSION_1_2

// API version macros
// Promoted from VK_KHR_driver_properties (extension 197)

type VkConformanceVersion struct {
	Major    uint8
	Minor    uint8
	Subminor uint8
	Patch    uint8
}

type VkPhysicalDeviceDriverProperties struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	DriverID           VkDriverId
	DriverName         [VkMaxDriverNameSize]uint8
	DriverInfo         [VkMaxDriverInfoSize]uint8
	ConformanceVersion VkConformanceVersion
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
	DeviceUUID                        [VkUuidSize]uint8
	DriverUUID                        [VkUuidSize]uint8
	DeviceLUID                        [VkLuidSize]uint8
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

type VkResolveModeFlags VkFlags

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
	DriverName                                           [VkMaxDriverNameSize]uint8
	DriverInfo                                           [VkMaxDriverInfoSize]uint8
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

// Promoted from VK_KHR_image_format_list (extension 148)

type VkImageFormatListCreateInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	ViewFormats []VkFormat
}

// Promoted from VK_KHR_sampler_mirror_clamp_to_edge (extension 15)
// Promoted from VK_KHR_vulkan_memory_model (extension 212)

type VkPhysicalDeviceVulkanMemoryModelFeatures struct {
	SType                                         VkStructureType
	PNext                                         unsafe.Pointer
	VulkanMemoryModel                             VkBool32
	VulkanMemoryModelDeviceScope                  VkBool32
	VulkanMemoryModelAvailabilityVisibilityChains VkBool32
}

// Promoted from VK_EXT_host_query_reset (extension 262)

func (h *VkDevice) ResetQueryPool(
	queryPool VkQueryPool,
	firstQuery uint32,
	queryCount uint32,
) {
}

type VkPhysicalDeviceHostQueryResetFeatures struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	HostQueryReset VkBool32
}

// Promoted from VK_KHR_timeline_semaphore (extension 208)

func (h *VkDevice) GetSemaphoreCounterValue(
	semaphore VkSemaphore,
	pValue *uint64,
) VkResult {
}

func (h *VkDevice) WaitSemaphores(
	pWaitInfo *VkSemaphoreWaitInfo,
	timeout uint64,
) VkResult {
}

func (h *VkDevice) SignalSemaphore(
	pSignalInfo *VkSemaphoreSignalInfo,
) VkResult {
}

type VkPhysicalDeviceTimelineSemaphoreFeatures struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	TimelineSemaphore VkBool32
}

type VkPhysicalDeviceTimelineSemaphoreProperties struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	MaxTimelineSemaphoreValueDifference uint64
}

type VkSemaphoreTypeCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SemaphoreType VkSemaphoreType
	InitialValue  uint64
}

type VkTimelineSemaphoreSubmitInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	WaitSemaphoreValues   []uint64
	SignalSemaphoreValues []uint64
}

type VkSemaphoreWaitFlags VkFlags

type VkSemaphoreWaitInfo struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Flags      VkSemaphoreWaitFlags
	Semaphores []VkSemaphore
	PValues    *uint64
}

type VkSemaphoreSignalInfo struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Semaphore VkSemaphore
	Value     uint64
}

// Promoted from VK_KHR_buffer_device_address (extension 258)

func (h *VkDevice) GetBufferAddress(
	pInfo *VkBufferDeviceAddressInfo,
) VkDeviceAddress {
}

func (h *VkDevice) GetBufferOpaqueCaptureAddress(
	pInfo *VkBufferDeviceAddressInfo,
) uint64 {
}

func (h *VkDevice) GetMemoryOpaqueCaptureAddress(
	pInfo *VkDeviceMemoryOpaqueCaptureAddressInfo,
) uint64 {
}

type VkPhysicalDeviceBufferDeviceAddressFeatures struct {
	SType                            VkStructureType
	PNext                            unsafe.Pointer
	BufferDeviceAddress              VkBool32
	BufferDeviceAddressCaptureReplay VkBool32
	BufferDeviceAddressMultiDevice   VkBool32
}

type VkBufferDeviceAddressInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Buffer VkBuffer
}

type VkBufferOpaqueCaptureAddressCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	OpaqueCaptureAddress uint64
}

type VkMemoryOpaqueCaptureAddressAllocateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	OpaqueCaptureAddress uint64
}

type VkDeviceMemoryOpaqueCaptureAddressInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Memory VkDeviceMemory
}

// Target: VK_COMPUTE_VERSION_1_2

// Promoted from VK_KHR_8bit_storage (extension 178)

type VkPhysicalDevice8BitStorageFeatures struct {
	SType                             VkStructureType
	PNext                             unsafe.Pointer
	StorageBuffer8BitAccess           VkBool32
	UniformAndStorageBuffer8BitAccess VkBool32
	StoragePushConstant8              VkBool32
}

// Promoted from VK_KHR_shader_atomic_int64 (extension 181)

type VkPhysicalDeviceShaderAtomicInt64Features struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ShaderBufferInt64Atomics VkBool32
	ShaderSharedInt64Atomics VkBool32
}

// Promoted from VK_KHR_shader_float16_int8 (extension 83)

type VkPhysicalDeviceShaderFloat16Int8Features struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	ShaderFloat16 VkBool32
	ShaderInt8    VkBool32
}

// Promoted from VK_KHR_shader_float_controls (extension 198)

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

// Promoted from VK_EXT_descriptor_indexing (extension 162)

type VkDescriptorSetLayoutBindingFlagsCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	BindingFlags []VkDescriptorBindingFlags
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

type VkDescriptorSetVariableDescriptorCountAllocateInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DescriptorCounts []uint32
}

type VkDescriptorSetVariableDescriptorCountLayoutSupport struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	MaxVariableDescriptorCount uint32
}

type VkDescriptorBindingFlags VkFlags

// Promoted from VK_EXT_scalar_block_layout (extension 222))

type VkPhysicalDeviceScalarBlockLayoutFeatures struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	ScalarBlockLayout VkBool32
}

// Promoted from VK_EXT_sampler_filter_minmax (extension 131)

type VkSamplerReductionModeCreateInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	ReductionMode VkSamplerReductionMode
}

type VkPhysicalDeviceSamplerFilterMinmaxProperties struct {
	SType                              VkStructureType
	PNext                              unsafe.Pointer
	FilterMinmaxSingleComponentFormats VkBool32
	FilterMinmaxImageComponentMapping  VkBool32
}

// Promoted from VK_KHR_uniform_buffer_standard_layout (extension 254)

type VkPhysicalDeviceUniformBufferStandardLayoutFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	UniformBufferStandardLayout VkBool32
}

// Promoted from VK_KHR_shader_subgroup_extended_types (extension 176)

type VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	ShaderSubgroupExtendedTypes VkBool32
}

// Promoted from VK_KHR_spirv_1_4 (extension 237)
// Target: VK_GRAPHICS_VERSION_1_2

// Promoted from VK_KHR_draw_indirect_count (extension 170)

func (h *VkCommandBuffer) DrawIndirectCount(
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
}

func (h *VkCommandBuffer) DrawIndexedIndirectCount(
	buffer VkBuffer,
	offset VkDeviceSize,
	countBuffer VkBuffer,
	countBufferOffset VkDeviceSize,
	maxDrawCount uint32,
	stride uint32,
) {
}

// Promoted from VK_KHR_create_renderpass2 (extension 110)

func (h *VkDevice) CreateRenderPass2(
	pCreateInfo *VkRenderPassCreateInfo2,
	pAllocator *VkAllocationCallbacks,
	pRenderPass *VkRenderPass,
) VkResult {
}

func (h *VkCommandBuffer) BeginRenderPass2(
	pRenderPassBegin *VkRenderPassBeginInfo,
	pSubpassBeginInfo *VkSubpassBeginInfo,
) {
}

func (h *VkCommandBuffer) NextSubpass2(
	pSubpassBeginInfo *VkSubpassBeginInfo,
	pSubpassEndInfo *VkSubpassEndInfo,
) {
}

func (h *VkCommandBuffer) EndRenderPass2(
	pSubpassEndInfo *VkSubpassEndInfo,
) {
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

type VkAttachmentReference2 struct {
	SType      VkStructureType
	PNext      unsafe.Pointer
	Attachment uint32
	Layout     VkImageLayout
	AspectMask VkImageAspectFlags
}

type VkSubpassDescription2 struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkSubpassDescriptionFlags
	PipelineBindPoint       VkPipelineBindPoint
	ViewMask                uint32
	InputAttachments        []VkAttachmentReference2
	ColorAttachments        []VkAttachmentReference2
	PResolveAttachments     *VkAttachmentReference2
	PDepthStencilAttachment *VkAttachmentReference2
	PreserveAttachments     []uint32
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

type VkSubpassBeginInfo struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	Contents VkSubpassContents
}

type VkSubpassEndInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
}

type VkRenderPassCreateInfo2 struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	Flags               VkRenderPassCreateFlags
	Attachments         []VkAttachmentDescription2
	Subpasses           []VkSubpassDescription2
	Dependencies        []VkSubpassDependency2
	CorrelatedViewMasks []uint32
}

// Promoted from VK_KHR_depth_stencil_resolve (extension 200)

type VkSubpassDescriptionDepthStencilResolve struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	DepthResolveMode               VkResolveModeFlagBits
	StencilResolveMode             VkResolveModeFlagBits
	PDepthStencilResolveAttachment *VkAttachmentReference2
}

type VkPhysicalDeviceDepthStencilResolveProperties struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	SupportedDepthResolveModes   VkResolveModeFlags
	SupportedStencilResolveModes VkResolveModeFlags
	IndependentResolveNone       VkBool32
	IndependentResolve           VkBool32
}

// Promoted from VK_EXT_separate_stencil_usage (extension 247)

type VkImageStencilUsageCreateInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	StencilUsage VkImageUsageFlags
}

// Promoted from VK_KHR_imageless_framebuffer (extension 109)

type VkPhysicalDeviceImagelessFramebufferFeatures struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ImagelessFramebuffer VkBool32
}

type VkFramebufferAttachmentImageInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Flags       VkImageCreateFlags
	Usage       VkImageUsageFlags
	Width       uint32
	Height      uint32
	LayerCount  uint32
	ViewFormats []VkFormat
}

type VkRenderPassAttachmentBeginInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Attachments []VkImageView
}

type VkFramebufferAttachmentsCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	AttachmentImageInfos []VkFramebufferAttachmentImageInfo
}

// Promoted from VK_KHR_separate_depth_stencil_layouts (extension 242)

type VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	SeparateDepthStencilLayouts VkBool32
}

type VkAttachmentReferenceStencilLayout struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	StencilLayout VkImageLayout
}

type VkAttachmentDescriptionStencilLayout struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	StencilInitialLayout VkImageLayout
	StencilFinalLayout   VkImageLayout
}

// Target: VK_VERSION_1_2

// Feature requirements
// Target: VK_BASE_VERSION_1_3

// API version macros

type VkFlags64 uint64

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

// Promoted from VK_EXT_tooling_info (extension 246)

func (h *VkPhysicalDevice) GetToolProperties(
	pToolCount *uint32,
	pToolProperties *VkPhysicalDeviceToolProperties,
) VkResult {
}

type VkToolPurposeFlags VkFlags

type VkPhysicalDeviceToolProperties struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Name        [VkMaxExtensionNameSize]uint8
	Version     [VkMaxExtensionNameSize]uint8
	Purposes    VkToolPurposeFlags
	Description [VkMaxDescriptionSize]uint8
	Layer       [VkMaxExtensionNameSize]uint8
}

// Promoted from VK_EXT_private_data (extension 296)

func (h *VkDevice) CreatePrivateDataSlot(
	pCreateInfo *VkPrivateDataSlotCreateInfo,
	pAllocator *VkAllocationCallbacks,
	pPrivateDataSlot *VkPrivateDataSlot,
) VkResult {
}

func (h *VkDevice) DestroyPrivateDataSlot(
	privateDataSlot VkPrivateDataSlot,
	pAllocator *VkAllocationCallbacks,
) {
}

func (h *VkDevice) SetPrivateData(
	objectType VkObjectType,
	objectHandle uint64,
	privateDataSlot VkPrivateDataSlot,
	data uint64,
) VkResult {
}

func (h *VkDevice) GetPrivateData(
	objectType VkObjectType,
	objectHandle uint64,
	privateDataSlot VkPrivateDataSlot,
	pData *uint64,
) {
}

type VkPhysicalDevicePrivateDataFeatures struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PrivateData VkBool32
}

type VkDevicePrivateDataCreateInfo struct {
	SType                       VkStructureType
	PNext                       unsafe.Pointer
	PrivateDataSlotRequestCount uint32
}

type VkPrivateDataSlotCreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkPrivateDataSlotCreateFlags
}

type VkPrivateDataSlot struct {
	handle unsafe.Pointer
}

type VkPrivateDataSlotCreateFlags VkFlags

// Promoted from VK_KHR_synchronization2 (extension 315)

func (h *VkCommandBuffer) PipelineBarrier2(
	pDependencyInfo *VkDependencyInfo,
) {
}

func (h *VkCommandBuffer) WriteTimestamp2(
	stage VkPipelineStageFlags2,
	queryPool VkQueryPool,
	query uint32,
) {
}

func (h *VkQueue) Submit2(
	submitCount uint32,
	pSubmits *VkSubmitInfo2,
	fence VkFence,
) VkResult {
}

type VkPipelineStageFlags2 VkFlags64

type VkAccessFlags2 VkFlags64

type VkMemoryBarrier2 struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	SrcStageMask  VkPipelineStageFlags2
	SrcAccessMask VkAccessFlags2
	DstStageMask  VkPipelineStageFlags2
	DstAccessMask VkAccessFlags2
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

type VkDependencyInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	DependencyFlags      VkDependencyFlags
	MemoryBarriers       []VkMemoryBarrier2
	BufferMemoryBarriers []VkBufferMemoryBarrier2
	ImageMemoryBarriers  []VkImageMemoryBarrier2
}

type VkSemaphoreSubmitInfo struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	Semaphore   VkSemaphore
	Value       uint64
	StageMask   VkPipelineStageFlags2
	DeviceIndex uint32
}

type VkCommandBufferSubmitInfo struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	CommandBuffer VkCommandBuffer
	DeviceMask    uint32
}

type VkSubmitFlags VkFlags

type VkSubmitInfo2 struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	Flags                VkSubmitFlags
	WaitSemaphoreInfos   []VkSemaphoreSubmitInfo
	CommandBufferInfos   []VkCommandBufferSubmitInfo
	SignalSemaphoreInfos []VkSemaphoreSubmitInfo
}

type VkPhysicalDeviceSynchronization2Features struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Synchronization2 VkBool32
}

// Promoted from VK_KHR_copy_commands2 (extension 338)

func (h *VkCommandBuffer) CopyBuffer2(
	pCopyBufferInfo *VkCopyBufferInfo2,
) {
}

func (h *VkCommandBuffer) CopyImage2(
	pCopyImageInfo *VkCopyImageInfo2,
) {
}

func (h *VkCommandBuffer) CopyBufferToImage2(
	pCopyBufferToImageInfo *VkCopyBufferToImageInfo2,
) {
}

func (h *VkCommandBuffer) CopyImageToBuffer2(
	pCopyImageToBufferInfo *VkCopyImageToBufferInfo2,
) {
}

type VkBufferCopy2 struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	SrcOffset VkDeviceSize
	DstOffset VkDeviceSize
	Size      VkDeviceSize
}

type VkCopyBufferInfo2 struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	SrcBuffer VkBuffer
	DstBuffer VkBuffer
	Regions   []VkBufferCopy2
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

type VkCopyImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	Regions        []VkImageCopy2
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

type VkCopyBufferToImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcBuffer      VkBuffer
	DstImage       VkImage
	DstImageLayout VkImageLayout
	Regions        []VkBufferImageCopy2
}

type VkCopyImageToBufferInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstBuffer      VkBuffer
	Regions        []VkBufferImageCopy2
}

// Promoted from VK_EXT_ycbcr_2plane_444_formats (does not promote the Feature struct, just the formats) (extension 331)
// Promoted from VK_EXT_4444_formats (does not promote the Feature struct, just the formats) (extension 341)
// Promoted from VK_EXT_texture_compression_astc_hdr (Feature struct is promoted, but becomes optional) (extension 67)

type VkPhysicalDeviceTextureCompressionASTCHDRFeatures struct {
	SType                      VkStructureType
	PNext                      unsafe.Pointer
	TextureCompressionASTC_HDR VkBool32
}

// Promoted from VK_KHR_format_feature_flags2 (extension 361)

type VkFormatFeatureFlags2 VkFlags64

type VkFormatProperties3 struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	LinearTilingFeatures  VkFormatFeatureFlags2
	OptimalTilingFeatures VkFormatFeatureFlags2
	BufferFeatures        VkFormatFeatureFlags2
}

// Promoted from VK_KHR_maintenance4 (extension 414)

func (h *VkDevice) GetBufferMemoryRequirements(
	pInfo *VkDeviceBufferMemoryRequirements,
	pMemoryRequirements *VkMemoryRequirements2,
) {
}

func (h *VkDevice) GetImageMemoryRequirements(
	pInfo *VkDeviceImageMemoryRequirements,
	pMemoryRequirements *VkMemoryRequirements2,
) {
}

func (h *VkDevice) GetImageSparseMemoryRequirements(
	pInfo *VkDeviceImageMemoryRequirements,
	pSparseMemoryRequirementCount *uint32,
	pSparseMemoryRequirements *VkSparseImageMemoryRequirements2,
) {
}

type VkPhysicalDeviceMaintenance4Features struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance4 VkBool32
}

type VkPhysicalDeviceMaintenance4Properties struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	MaxBufferSize VkDeviceSize
}

type VkDeviceBufferMemoryRequirements struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PCreateInfo *VkBufferCreateInfo
}

type VkDeviceImageMemoryRequirements struct {
	SType       VkStructureType
	PNext       unsafe.Pointer
	PCreateInfo *VkImageCreateInfo
	PlaneAspect VkImageAspectFlagBits
}

// Target: VK_COMPUTE_VERSION_1_3

// Promoted from VK_EXT_pipeline_creation_feedback (extension 193)

type VkPipelineCreationFeedbackFlags VkFlags

type VkPipelineCreationFeedback struct {
	Flags    VkPipelineCreationFeedbackFlags
	Duration uint64
}

type VkPipelineCreationFeedbackCreateInfo struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	PPipelineCreationFeedback      *VkPipelineCreationFeedback
	PipelineStageCreationFeedbacks []VkPipelineCreationFeedback
}

// Promoted from VK_KHR_shader_terminate_invocation (extension 216)

type VkPhysicalDeviceShaderTerminateInvocationFeatures struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	ShaderTerminateInvocation VkBool32
}

// Promoted from VK_EXT_shader_demote_to_helper_invocation (extension 277)

type VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures struct {
	SType                          VkStructureType
	PNext                          unsafe.Pointer
	ShaderDemoteToHelperInvocation VkBool32
}

// Promoted from VK_KHR_shader_non_semantic_info (extension 294)
// Promoted from VK_EXT_pipeline_creation_cache_control (extension 298)

type VkPhysicalDevicePipelineCreationCacheControlFeatures struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	PipelineCreationCacheControl VkBool32
}

// Promoted from VK_KHR_synchronization2 (extension 315)

func (h *VkCommandBuffer) SetEvent2(
	event VkEvent,
	pDependencyInfo *VkDependencyInfo,
) {
}

func (h *VkCommandBuffer) ResetEvent2(
	event VkEvent,
	stageMask VkPipelineStageFlags2,
) {
}

func (h *VkCommandBuffer) WaitEvents2(
	eventCount uint32,
	pEvents *VkEvent,
	pDependencyInfos *VkDependencyInfo,
) {
}

// Promoted from VK_KHR_zero_initialize_workgroup_memory (extension 326)

type VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures struct {
	SType                               VkStructureType
	PNext                               unsafe.Pointer
	ShaderZeroInitializeWorkgroupMemory VkBool32
}

// Promoted from VK_EXT_image_robustness (extension 336)

type VkPhysicalDeviceImageRobustnessFeatures struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	RobustImageAccess VkBool32
}

// Promoted from VK_EXT_subgroup_size_control (STDPROMOTE/PROPLIMCHANGE) (extension 226)

type VkPhysicalDeviceSubgroupSizeControlFeatures struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	SubgroupSizeControl  VkBool32
	ComputeFullSubgroups VkBool32
}

type VkPhysicalDeviceSubgroupSizeControlProperties struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MinSubgroupSize              uint32
	MaxSubgroupSize              uint32
	MaxComputeWorkgroupSubgroups uint32
	RequiredSubgroupSizeStages   VkShaderStageFlags
}

type VkPipelineShaderStageRequiredSubgroupSizeCreateInfo struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	RequiredSubgroupSize uint32
}

// Promoted from VK_EXT_inline_uniform_block (STDPROMOTE/PROPLIMCHANGE) (extension 139)

type VkPhysicalDeviceInlineUniformBlockFeatures struct {
	SType                                              VkStructureType
	PNext                                              unsafe.Pointer
	InlineUniformBlock                                 VkBool32
	DescriptorBindingInlineUniformBlockUpdateAfterBind VkBool32
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

type VkWriteDescriptorSetInlineUniformBlock struct {
	SType    VkStructureType
	PNext    unsafe.Pointer
	DataSize uint32
	PData    unsafe.Pointer
}

type VkDescriptorPoolInlineUniformBlockCreateInfo struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	MaxInlineUniformBlockBindings uint32
}

// Promoted from VK_KHR_shader_integer_dot_product (extension 281)

type VkPhysicalDeviceShaderIntegerDotProductFeatures struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ShaderIntegerDotProduct VkBool32
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

// Promoted from VK_EXT_texel_buffer_alignment (extension 282)

type VkPhysicalDeviceTexelBufferAlignmentProperties struct {
	SType                                        VkStructureType
	PNext                                        unsafe.Pointer
	StorageTexelBufferOffsetAlignmentBytes       VkDeviceSize
	StorageTexelBufferOffsetSingleTexelAlignment VkBool32
	UniformTexelBufferOffsetAlignmentBytes       VkDeviceSize
	UniformTexelBufferOffsetSingleTexelAlignment VkBool32
}

// Target: VK_GRAPHICS_VERSION_1_3

// Promoted from VK_KHR_copy_commands2 (extension 338)

func (h *VkCommandBuffer) BlitImage2(
	pBlitImageInfo *VkBlitImageInfo2,
) {
}

func (h *VkCommandBuffer) ResolveImage2(
	pResolveImageInfo *VkResolveImageInfo2,
) {
}

type VkImageBlit2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcSubresource VkImageSubresourceLayers
	SrcOffsets     [2]VkOffset3D
	DstSubresource VkImageSubresourceLayers
	DstOffsets     [2]VkOffset3D
}

type VkBlitImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	Regions        []VkImageBlit2
	Filter         VkFilter
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

type VkResolveImageInfo2 struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	Regions        []VkImageResolve2
}

// Promoted from VK_KHR_dynamic_rendering (extension 45)

func (h *VkCommandBuffer) BeginRendering(
	pRenderingInfo *VkRenderingInfo,
) {
}

func (h *VkCommandBuffer) EndRendering() {
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

type VkRenderingInfo struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	Flags              VkRenderingFlags
	RenderArea         VkRect2D
	LayerCount         uint32
	ViewMask           uint32
	ColorAttachments   []VkRenderingAttachmentInfo
	PDepthAttachment   *VkRenderingAttachmentInfo
	PStencilAttachment *VkRenderingAttachmentInfo
}

type VkPipelineRenderingCreateInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ViewMask                uint32
	ColorAttachmentFormats  []VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
}

type VkPhysicalDeviceDynamicRenderingFeatures struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	DynamicRendering VkBool32
}

type VkCommandBufferInheritanceRenderingInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	Flags                   VkRenderingFlags
	ViewMask                uint32
	ColorAttachmentCount    uint32
	PColorAttachmentFormats *VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
	RasterizationSamples    VkSampleCountFlagBits
}

type VkRenderingFlags VkFlags

// Promoted from VK_EXT_extended_dynamic_state (Feature struct is not promoted) (extension 268)

func (h *VkCommandBuffer) SetCullMode(
	cullMode VkCullModeFlags,
) {
}

func (h *VkCommandBuffer) SetFrontFace(
	frontFace VkFrontFace,
) {
}

func (h *VkCommandBuffer) SetPrimitiveTopology(
	primitiveTopology VkPrimitiveTopology,
) {
}

func (h *VkCommandBuffer) SetViewportWithCount(
	viewportCount uint32,
	pViewports *VkViewport,
) {
}

func (h *VkCommandBuffer) SetScissorWithCount(
	scissorCount uint32,
	pScissors *VkRect2D,
) {
}

func (h *VkCommandBuffer) BindVertexBuffers2(
	firstBinding uint32,
	bindingCount uint32,
	pBuffers *VkBuffer,
	pOffsets *VkDeviceSize,
	pSizes *VkDeviceSize,
	pStrides *VkDeviceSize,
) {
}

func (h *VkCommandBuffer) SetDepthTestEnable(
	depthTestEnable VkBool32,
) {
}

func (h *VkCommandBuffer) SetDepthWriteEnable(
	depthWriteEnable VkBool32,
) {
}

func (h *VkCommandBuffer) SetDepthCompareOp(
	depthCompareOp VkCompareOp,
) {
}

func (h *VkCommandBuffer) SetDepthBoundsTestEnable(
	depthBoundsTestEnable VkBool32,
) {
}

func (h *VkCommandBuffer) SetStencilTestEnable(
	stencilTestEnable VkBool32,
) {
}

func (h *VkCommandBuffer) SetStencilOp(
	faceMask VkStencilFaceFlags,
	failOp VkStencilOp,
	passOp VkStencilOp,
	depthFailOp VkStencilOp,
	compareOp VkCompareOp,
) {
}

// Promoted from VK_EXT_extended_dynamic_state2 (Feature struct and optional state are not promoted) (extension 378)

func (h *VkCommandBuffer) SetRasterizerDiscardEnable(
	rasterizerDiscardEnable VkBool32,
) {
}

func (h *VkCommandBuffer) SetDepthBiasEnable(
	depthBiasEnable VkBool32,
) {
}

func (h *VkCommandBuffer) SetPrimitiveRestartEnable(
	primitiveRestartEnable VkBool32,
) {
}

// Target: VK_VERSION_1_3

// Feature requirements
// Target: VK_BASE_VERSION_1_4

// API version macros

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
	CopySrcLayouts                                      []VkImageLayout
	CopyDstLayouts                                      []VkImageLayout
	OptimalTilingLayoutUUID                             [VkUuidSize]uint8
	IdenticalMemoryTypeRequirements                     VkBool32
}

// Promoted from VK_KHR_global_priority (extension 189)

type VkDeviceQueueGlobalPriorityCreateInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	GlobalPriority VkQueueGlobalPriority
}

type VkPhysicalDeviceGlobalPriorityQueryFeatures struct {
	SType               VkStructureType
	PNext               unsafe.Pointer
	GlobalPriorityQuery VkBool32
}

type VkQueueFamilyGlobalPriorityProperties struct {
	SType     VkStructureType
	PNext     unsafe.Pointer
	Riorities [][VkMaxGlobalPrioritySize]VkQueueGlobalPriority
}

// Promoted from VK_KHR_index_type_uint8 (extension 534) 'Roadmap 2024'

type VkPhysicalDeviceIndexTypeUint8Features struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	IndexTypeUint8 VkBool32
}

// Promoted from VK_KHR_map_memory2 (extension 272) 'Roadmap 2024'

func (h *VkDevice) MapMemory2(
	pMemoryMapInfo *VkMemoryMapInfo,
	ppData **void,
) VkResult {
}

func (h *VkDevice) UnmapMemory2(
	pMemoryUnmapInfo *VkMemoryUnmapInfo,
) VkResult {
}

type VkMemoryMapInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkMemoryMapFlags
	Memory VkDeviceMemory
	Offset VkDeviceSize
	Size   VkDeviceSize
}

type VkMemoryUnmapInfo struct {
	SType  VkStructureType
	PNext  unsafe.Pointer
	Flags  VkMemoryUnmapFlags
	Memory VkDeviceMemory
}

type VkMemoryUnmapFlags VkFlags

// Promoted from VK_KHR_maintenance5 (extension 471) 'Roadmap 2024'

func (h *VkDevice) GetImageSubresourceLayout(
	pInfo *VkDeviceImageSubresourceInfo,
	pLayout *VkSubresourceLayout2,
) {
}

func (h *VkDevice) GetImageSubresourceLayout2(
	image VkImage,
	pSubresource *VkImageSubresource2,
	pLayout *VkSubresourceLayout2,
) {
}

type VkPhysicalDeviceMaintenance5Features struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance5 VkBool32
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

type VkSubresourceLayout2 struct {
	SType             VkStructureType
	PNext             unsafe.Pointer
	SubresourceLayout VkSubresourceLayout
}

type VkImageSubresource2 struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	ImageSubresource VkImageSubresource
}

type VkDeviceImageSubresourceInfo struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	PCreateInfo  *VkImageCreateInfo
	PSubresource *VkImageSubresource2
}

type VkBufferUsageFlags2 VkFlags64

type VkBufferUsageFlags2CreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Usage VkBufferUsageFlags2
}

// Promoted from VK_KHR_maintenance6 (extension 546) 'additional functionality'

type VkPhysicalDeviceMaintenance6Features struct {
	SType        VkStructureType
	PNext        unsafe.Pointer
	Maintenance6 VkBool32
}

type VkPhysicalDeviceMaintenance6Properties struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	BlockTexelViewCompatibleMultipleLayers VkBool32
	MaxCombinedImageSamplerDescriptorCount uint32
	FragmentShadingRateClampCombinerInputs VkBool32
}

type VkBindMemoryStatus struct {
	SType   VkStructureType
	PNext   unsafe.Pointer
	PResult *VkResult
}

// Promoted (as optional feature) from VK_EXT_host_image_copy (extension 271) 'streaming transfers'

func (h *VkDevice) CopyMemoryToImage(
	pCopyMemoryToImageInfo *VkCopyMemoryToImageInfo,
) VkResult {
}

func (h *VkDevice) CopyImageToMemory(
	pCopyImageToMemoryInfo *VkCopyImageToMemoryInfo,
) VkResult {
}

func (h *VkDevice) CopyImageToImage(
	pCopyImageToImageInfo *VkCopyImageToImageInfo,
) VkResult {
}

func (h *VkDevice) TransitionImageLayout(
	transitionCount uint32,
	pTransitions *VkHostImageLayoutTransitionInfo,
) VkResult {
}

type VkPhysicalDeviceHostImageCopyFeatures struct {
	SType         VkStructureType
	PNext         unsafe.Pointer
	HostImageCopy VkBool32
}

type VkPhysicalDeviceHostImageCopyProperties struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	CopySrcLayouts                  []VkImageLayout
	CopyDstLayouts                  []VkImageLayout
	OptimalTilingLayoutUUID         [VkUuidSize]uint8
	IdenticalMemoryTypeRequirements VkBool32
}

type VkHostImageCopyFlags VkFlags

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

type VkCopyMemoryToImageInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkHostImageCopyFlags
	DstImage       VkImage
	DstImageLayout VkImageLayout
	Regions        []VkMemoryToImageCopy
}

type VkCopyImageToMemoryInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkHostImageCopyFlags
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	Regions        []VkImageToMemoryCopy
}

type VkCopyImageToImageInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	Flags          VkHostImageCopyFlags
	SrcImage       VkImage
	SrcImageLayout VkImageLayout
	DstImage       VkImage
	DstImageLayout VkImageLayout
	Regions        []VkImageCopy2
}

type VkHostImageLayoutTransitionInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	Image            VkImage
	OldLayout        VkImageLayout
	NewLayout        VkImageLayout
	SubresourceRange VkImageSubresourceRange
}

type VkSubresourceHostMemcpySize struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Size  VkDeviceSize
}

type VkHostImageCopyDevicePerformanceQuery struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	OptimalDeviceAccess   VkBool32
	IdenticalMemoryLayout VkBool32
}

// Target: VK_COMPUTE_VERSION_1_4

// Promoted from VK_KHR_shader_subgroup_rotate (extension 417) 'Roadmap 2024'

type VkPhysicalDeviceShaderSubgroupRotateFeatures struct {
	SType                         VkStructureType
	PNext                         unsafe.Pointer
	ShaderSubgroupRotate          VkBool32
	ShaderSubgroupRotateClustered VkBool32
}

// Promoted from VK_KHR_shader_float_controls2 (extension 529) 'Roadmap 2024'

type VkPhysicalDeviceShaderFloatControls2Features struct {
	SType                VkStructureType
	PNext                unsafe.Pointer
	ShaderFloatControls2 VkBool32
}

// Promoted from VK_KHR_shader_expect_assume (extension 545) 'Roadmap 2024'

type VkPhysicalDeviceShaderExpectAssumeFeatures struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	ShaderExpectAssume VkBool32
}

// Promoted from VK_KHR_maintenance5 (extension 471) 'Roadmap 2024'

type VkPipelineCreateFlags2 VkFlags64

type VkPipelineCreateFlags2CreateInfo struct {
	SType VkStructureType
	PNext unsafe.Pointer
	Flags VkPipelineCreateFlags2
}

// Promoted as an interaction between VK_KHR_maintenance5 (extension 471) 'Roadmap 2024' and VK_EXT_pipeline_protected_access (extension 467) 'additional functionality'
// Promoted from VK_KHR_push_descriptor (extension 81) 'Roadmap 2024'

func (h *VkCommandBuffer) PushDescriptorSet(
	pipelineBindPoint VkPipelineBindPoint,
	layout VkPipelineLayout,
	set uint32,
	descriptorWriteCount uint32,
	pDescriptorWrites *VkWriteDescriptorSet,
) {
}

func (h *VkCommandBuffer) PushDescriptorSetWithTemplate(
	descriptorUpdateTemplate VkDescriptorUpdateTemplate,
	layout VkPipelineLayout,
	set uint32,
	pData unsafe.Pointer,
) {
}

type VkPhysicalDevicePushDescriptorProperties struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	MaxPushDescriptors uint32
}

// Promoted from VK_KHR_maintenance6 (extension 546) 'additional functionality'

func (h *VkCommandBuffer) BindDescriptorSets2(
	pBindDescriptorSetsInfo *VkBindDescriptorSetsInfo,
) {
}

func (h *VkCommandBuffer) PushConstants2(
	pPushConstantsInfo *VkPushConstantsInfo,
) {
}

func (h *VkCommandBuffer) PushDescriptorSet2(
	pPushDescriptorSetInfo *VkPushDescriptorSetInfo,
) {
}

func (h *VkCommandBuffer) PushDescriptorSetWithTemplate2(
	pPushDescriptorSetWithTemplateInfo *VkPushDescriptorSetWithTemplateInfo,
) {
}

type VkBindDescriptorSetsInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	StageFlags     VkShaderStageFlags
	Layout         VkPipelineLayout
	FirstSet       uint32
	DescriptorSets []VkDescriptorSet
	DynamicOffsets []uint32
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

type VkPushDescriptorSetInfo struct {
	SType            VkStructureType
	PNext            unsafe.Pointer
	StageFlags       VkShaderStageFlags
	Layout           VkPipelineLayout
	Set              uint32
	DescriptorWrites []VkWriteDescriptorSet
}

type VkPushDescriptorSetWithTemplateInfo struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	DescriptorUpdateTemplate VkDescriptorUpdateTemplate
	Layout                   VkPipelineLayout
	Set                      uint32
	PData                    unsafe.Pointer
}

// Promoted from VK_EXT_pipeline_protected_access (extension 467) 'additional functionality'

type VkPhysicalDevicePipelineProtectedAccessFeatures struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	PipelineProtectedAccess VkBool32
}

// Promoted from VK_EXT_pipeline_robustness (extension 69) 'additional functionality'

type VkPhysicalDevicePipelineRobustnessFeatures struct {
	SType              VkStructureType
	PNext              unsafe.Pointer
	PipelineRobustness VkBool32
}

type VkPhysicalDevicePipelineRobustnessProperties struct {
	SType                           VkStructureType
	PNext                           unsafe.Pointer
	DefaultRobustnessStorageBuffers VkPipelineRobustnessBufferBehavior
	DefaultRobustnessUniformBuffers VkPipelineRobustnessBufferBehavior
	DefaultRobustnessVertexInputs   VkPipelineRobustnessBufferBehavior
	DefaultRobustnessImages         VkPipelineRobustnessImageBehavior
}

type VkPipelineRobustnessCreateInfo struct {
	SType          VkStructureType
	PNext          unsafe.Pointer
	StorageBuffers VkPipelineRobustnessBufferBehavior
	UniformBuffers VkPipelineRobustnessBufferBehavior
	VertexInputs   VkPipelineRobustnessBufferBehavior
	Images         VkPipelineRobustnessImageBehavior
}

// Target: VK_GRAPHICS_VERSION_1_4

// Promoted from VK_KHR_load_store_op_none (extension 527) 'Roadmap 2024' (VK_ATTACHMENT_STORE_OP_NONE is defined in Vulkan 1.3)
// Promoted from VK_KHR_line_rasterization (extension 535) 'Roadmap 2024'

func (h *VkCommandBuffer) SetLineStipple(
	lineStippleFactor uint32,
	lineStipplePattern uint16,
) {
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

type VkPhysicalDeviceLineRasterizationProperties struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	LineSubPixelPrecisionBits uint32
}

type VkPipelineRasterizationLineStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	LineRasterizationMode VkLineRasterizationMode
	StippledLineEnable    VkBool32
	LineStippleFactor     uint32
	LineStipplePattern    uint16
}

// Promoted from VK_KHR_vertex_attribute_divisor (extension 526) 'Roadmap 2024'

type VkPhysicalDeviceVertexAttributeDivisorProperties struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	MaxVertexAttribDivisor       uint32
	SupportsNonZeroFirstInstance VkBool32
}

type VkVertexInputBindingDivisorDescription struct {
	Binding uint32
	Divisor uint32
}

type VkPipelineVertexInputDivisorStateCreateInfo struct {
	SType                 VkStructureType
	PNext                 unsafe.Pointer
	VertexBindingDivisors []VkVertexInputBindingDivisorDescription
}

type VkPhysicalDeviceVertexAttributeDivisorFeatures struct {
	SType                                  VkStructureType
	PNext                                  unsafe.Pointer
	VertexAttributeInstanceRateDivisor     VkBool32
	VertexAttributeInstanceRateZeroDivisor VkBool32
}

// Promoted from VK_KHR_maintenance5 (extension 471) 'Roadmap 2024'

func (h *VkCommandBuffer) BindIndexBuffer2(
	buffer VkBuffer,
	offset VkDeviceSize,
	size VkDeviceSize,
	indexType VkIndexType,
) {
}

func (h *VkDevice) GetRenderingAreaGranularity(
	pRenderingAreaInfo *VkRenderingAreaInfo,
	pGranularity *VkExtent2D,
) {
}

type VkRenderingAreaInfo struct {
	SType                   VkStructureType
	PNext                   unsafe.Pointer
	ViewMask                uint32
	ColorAttachmentFormats  []VkFormat
	DepthAttachmentFormat   VkFormat
	StencilAttachmentFormat VkFormat
}

// Promoted from VK_KHR_dynamic_rendering_local_read (extension 233) 'Roadmap 2024'

func (h *VkCommandBuffer) SetRenderingAttachmentLocations(
	pLocationInfo *VkRenderingAttachmentLocationInfo,
) {
}

func (h *VkCommandBuffer) SetRenderingInputAttachmentIndices(
	pInputAttachmentIndexInfo *VkRenderingInputAttachmentIndexInfo,
) {
}

type VkPhysicalDeviceDynamicRenderingLocalReadFeatures struct {
	SType                     VkStructureType
	PNext                     unsafe.Pointer
	DynamicRenderingLocalRead VkBool32
}

type VkRenderingAttachmentLocationInfo struct {
	SType                    VkStructureType
	PNext                    unsafe.Pointer
	ColorAttachmentLocations []uint32
}

type VkRenderingInputAttachmentIndexInfo struct {
	SType                        VkStructureType
	PNext                        unsafe.Pointer
	ColorAttachmentInputIndices  []uint32
	PDepthInputAttachmentIndex   *uint32
	PStencilInputAttachmentIndex *uint32
}

// Target: VK_VERSION_1_4

// Feature requirements

const (
	VkMaxPhysicalDeviceNameSize                              = 256
	VkUuidSize                                               = 16
	VkLuidSize                                               = 8
	VkMaxExtensionNameSize                                   = 256
	VkMaxDescriptionSize                                     = 256
	VkMaxMemoryTypes                                         = 32
	VkMaxMemoryHeaps                                         = 16
	VkLodClampNone                                           = 1000.0
	VkRemainingMipLevels                                     = ^0
	VkRemainingArrayLayers                                   = ^0
	VkRemaining3dSlicesExt                                   = ^0
	VkWholeSize                                              = ^0
	VkAttachmentUnused                                       = ^0
	VkTrue                                                   = 1
	VkFalse                                                  = 0
	VkQueueFamilyIgnored                                     = ^0
	VkQueueFamilyExternal                                    = ^1
	VkQueueFamilyForeignExt                                  = ^2
	VkSubpassExternal                                        = ^0
	VkMaxDeviceGroupSize                                     = 32
	VkMaxDriverNameSize                                      = 256
	VkMaxDriverInfoSize                                      = 256
	VkShaderUnusedKhr                                        = ^0
	VkMaxGlobalPrioritySize                                  = 16
	VkMaxShaderModuleIdentifierSizeExt                       = 32
	VkMaxPipelineBinaryKeySizeKhr                            = 32
	VkMaxVideoAv1ReferencesPerFrameKhr                       = 7
	VkMaxVideoVp9ReferencesPerFrameKhr                       = 3
	VkShaderIndexUnusedAmdx                                  = ^0
	VkPartitionedAccelerationStructurePartitionIndexGlobalNv = ^0
	VkCompressedTriangleFormatDgf1ByteAlignmentAmdx          = 128
	VkCompressedTriangleFormatDgf1ByteStrideAmdx             = 128
	VkMaxPhysicalDeviceDataGraphOperationSetNameSizeArm      = 128
	VkDataGraphModelToolchainVersionLengthQcom               = 3
	VkComputeOccupancyPriorityLowNv                          = 0.25
	VkComputeOccupancyPriorityNormalNv                       = 0.50
	VkComputeOccupancyPriorityHighNv                         = 0.75
)

type VkImageLayout int64

var (
	VkImageLayoutUndefined                     VkImageLayout = 0
	VkImageLayoutGeneral                       VkImageLayout = 1
	VkImageLayoutColorAttachmentOptimal        VkImageLayout = 2
	VkImageLayoutDepthStencilAttachmentOptimal VkImageLayout = 3
	VkImageLayoutDepthStencilReadOnlyOptimal   VkImageLayout = 4
	VkImageLayoutShaderReadOnlyOptimal         VkImageLayout = 5
	VkImageLayoutTransferSrcOptimal            VkImageLayout = 6
	VkImageLayoutTransferDstOptimal            VkImageLayout = 7
	VkImageLayoutPreinitialized                VkImageLayout = 8
)

type VkAttachmentLoadOp int64

var (
	VkAttachmentLoadOpLoad     VkAttachmentLoadOp = 0
	VkAttachmentLoadOpClear    VkAttachmentLoadOp = 1
	VkAttachmentLoadOpDontCare VkAttachmentLoadOp = 2
)

type VkAttachmentStoreOp int64

var (
	VkAttachmentStoreOpStore    VkAttachmentStoreOp = 0
	VkAttachmentStoreOpDontCare VkAttachmentStoreOp = 1
)

type VkImageType int64

var (
	VkImageType1d VkImageType = 0
	VkImageType2d VkImageType = 1
	VkImageType3d VkImageType = 2
)

type VkImageTiling int64

var (
	VkImageTilingOptimal VkImageTiling = 0
	VkImageTilingLinear  VkImageTiling = 1
)

type VkImageViewType int64

var (
	VkImageViewType1d        VkImageViewType = 0
	VkImageViewType2d        VkImageViewType = 1
	VkImageViewType3d        VkImageViewType = 2
	VkImageViewTypeCube      VkImageViewType = 3
	VkImageViewType1dArray   VkImageViewType = 4
	VkImageViewType2dArray   VkImageViewType = 5
	VkImageViewTypeCubeArray VkImageViewType = 6
)

type VkCommandBufferLevel int64

var (
	VkCommandBufferLevelPrimary   VkCommandBufferLevel = 0
	VkCommandBufferLevelSecondary VkCommandBufferLevel = 1
)

type VkComponentSwizzle int64

var (
	VkComponentSwizzleIdentity VkComponentSwizzle = 0
	VkComponentSwizzleZero     VkComponentSwizzle = 1
	VkComponentSwizzleOne      VkComponentSwizzle = 2
	VkComponentSwizzleR        VkComponentSwizzle = 3
	VkComponentSwizzleG        VkComponentSwizzle = 4
	VkComponentSwizzleB        VkComponentSwizzle = 5
	VkComponentSwizzleA        VkComponentSwizzle = 6
)

type VkDescriptorType int64

var (
	VkDescriptorTypeSampler              VkDescriptorType = 0
	VkDescriptorTypeCombinedImageSampler VkDescriptorType = 1
	VkDescriptorTypeSampledImage         VkDescriptorType = 2
	VkDescriptorTypeStorageImage         VkDescriptorType = 3
	VkDescriptorTypeUniformTexelBuffer   VkDescriptorType = 4
	VkDescriptorTypeStorageTexelBuffer   VkDescriptorType = 5
	VkDescriptorTypeUniformBuffer        VkDescriptorType = 6
	VkDescriptorTypeStorageBuffer        VkDescriptorType = 7
	VkDescriptorTypeUniformBufferDynamic VkDescriptorType = 8
	VkDescriptorTypeStorageBufferDynamic VkDescriptorType = 9
	VkDescriptorTypeInputAttachment      VkDescriptorType = 10
)

type VkQueryType int64

var (
	VkQueryTypeOcclusion          VkQueryType = 0
	VkQueryTypePipelineStatistics VkQueryType = 1
	VkQueryTypeTimestamp          VkQueryType = 2
)

type VkBorderColor int64

var (
	VkBorderColorFloatTransparentBlack VkBorderColor = 0
	VkBorderColorIntTransparentBlack   VkBorderColor = 1
	VkBorderColorFloatOpaqueBlack      VkBorderColor = 2
	VkBorderColorIntOpaqueBlack        VkBorderColor = 3
	VkBorderColorFloatOpaqueWhite      VkBorderColor = 4
	VkBorderColorIntOpaqueWhite        VkBorderColor = 5
)

type VkPipelineBindPoint int64

var (
	VkPipelineBindPointGraphics VkPipelineBindPoint = 0
	VkPipelineBindPointCompute  VkPipelineBindPoint = 1
)

type VkPipelineCacheHeaderVersion int64

var (
	Vkone VkPipelineCacheHeaderVersion = 1
)

type VkPipelineCacheCreateFlagBits int64

var ()

type VkPrimitiveTopology int64

var (
	VkPrimitiveTopologyPointList                  VkPrimitiveTopology = 0
	VkPrimitiveTopologyLineList                   VkPrimitiveTopology = 1
	VkPrimitiveTopologyLineStrip                  VkPrimitiveTopology = 2
	VkPrimitiveTopologyTriangleList               VkPrimitiveTopology = 3
	VkPrimitiveTopologyTriangleStrip              VkPrimitiveTopology = 4
	VkPrimitiveTopologyTriangleFan                VkPrimitiveTopology = 5
	VkPrimitiveTopologyLineListWithAdjacency      VkPrimitiveTopology = 6
	VkPrimitiveTopologyLineStripWithAdjacency     VkPrimitiveTopology = 7
	VkPrimitiveTopologyTriangleListWithAdjacency  VkPrimitiveTopology = 8
	VkPrimitiveTopologyTriangleStripWithAdjacency VkPrimitiveTopology = 9
	VkPrimitiveTopologyPatchList                  VkPrimitiveTopology = 10
)

type VkSharingMode int64

var (
	VkSharingModeExclusive  VkSharingMode = 0
	VkSharingModeConcurrent VkSharingMode = 1
)

type VkIndexType int64

var (
	VkIndexTypeUint16 VkIndexType = 0
	VkIndexTypeUint32 VkIndexType = 1
)

type VkFilter int64

var (
	VkFilterNearest VkFilter = 0
	VkFilterLinear  VkFilter = 1
)

type VkSamplerMipmapMode int64

var (
	VkSamplerMipmapModeNearest VkSamplerMipmapMode = 0
	VkSamplerMipmapModeLinear  VkSamplerMipmapMode = 1
)

type VkSamplerAddressMode int64

var (
	VkSamplerAddressModeRepeat         VkSamplerAddressMode = 0
	VkSamplerAddressModeMirroredRepeat VkSamplerAddressMode = 1
	VkSamplerAddressModeClampToEdge    VkSamplerAddressMode = 2
	VkSamplerAddressModeClampToBorder  VkSamplerAddressMode = 3
)

type VkCompareOp int64

var (
	VkCompareOpNever          VkCompareOp = 0
	VkCompareOpLess           VkCompareOp = 1
	VkCompareOpEqual          VkCompareOp = 2
	VkCompareOpLessOrEqual    VkCompareOp = 3
	VkCompareOpGreater        VkCompareOp = 4
	VkCompareOpNotEqual       VkCompareOp = 5
	VkCompareOpGreaterOrEqual VkCompareOp = 6
	VkCompareOpAlways         VkCompareOp = 7
)

type VkPolygonMode int64

var (
	VkPolygonModeFill  VkPolygonMode = 0
	VkPolygonModeLine  VkPolygonMode = 1
	VkPolygonModePoint VkPolygonMode = 2
)

type VkFrontFace int64

var (
	VkFrontFaceCounterClockwise VkFrontFace = 0
	VkFrontFaceClockwise        VkFrontFace = 1
)

type VkBlendFactor int64

var (
	VkBlendFactorZero                  VkBlendFactor = 0
	VkBlendFactorOne                   VkBlendFactor = 1
	VkBlendFactorSrcColor              VkBlendFactor = 2
	VkBlendFactorOneMinusSrcColor      VkBlendFactor = 3
	VkBlendFactorDstColor              VkBlendFactor = 4
	VkBlendFactorOneMinusDstColor      VkBlendFactor = 5
	VkBlendFactorSrcAlpha              VkBlendFactor = 6
	VkBlendFactorOneMinusSrcAlpha      VkBlendFactor = 7
	VkBlendFactorDstAlpha              VkBlendFactor = 8
	VkBlendFactorOneMinusDstAlpha      VkBlendFactor = 9
	VkBlendFactorConstantColor         VkBlendFactor = 10
	VkBlendFactorOneMinusConstantColor VkBlendFactor = 11
	VkBlendFactorConstantAlpha         VkBlendFactor = 12
	VkBlendFactorOneMinusConstantAlpha VkBlendFactor = 13
	VkBlendFactorSrcAlphaSaturate      VkBlendFactor = 14
	VkBlendFactorSrc1Color             VkBlendFactor = 15
	VkBlendFactorOneMinusSrc1Color     VkBlendFactor = 16
	VkBlendFactorSrc1Alpha             VkBlendFactor = 17
	VkBlendFactorOneMinusSrc1Alpha     VkBlendFactor = 18
)

type VkBlendOp int64

var (
	VkBlendOpAdd             VkBlendOp = 0
	VkBlendOpSubtract        VkBlendOp = 1
	VkBlendOpReverseSubtract VkBlendOp = 2
	VkBlendOpMin             VkBlendOp = 3
	VkBlendOpMax             VkBlendOp = 4
)

type VkStencilOp int64

var (
	VkStencilOpKeep              VkStencilOp = 0
	VkStencilOpZero              VkStencilOp = 1
	VkStencilOpReplace           VkStencilOp = 2
	VkStencilOpIncrementAndClamp VkStencilOp = 3
	VkStencilOpDecrementAndClamp VkStencilOp = 4
	VkStencilOpInvert            VkStencilOp = 5
	VkStencilOpIncrementAndWrap  VkStencilOp = 6
	VkStencilOpDecrementAndWrap  VkStencilOp = 7
)

type VkLogicOp int64

var (
	VkLogicOpClear        VkLogicOp = 0
	VkLogicOpAnd          VkLogicOp = 1
	VkLogicOpAndReverse   VkLogicOp = 2
	VkLogicOpCopy         VkLogicOp = 3
	VkLogicOpAndInverted  VkLogicOp = 4
	VkLogicOpNoOp         VkLogicOp = 5
	VkLogicOpXor          VkLogicOp = 6
	VkLogicOpOr           VkLogicOp = 7
	VkLogicOpNor          VkLogicOp = 8
	VkLogicOpEquivalent   VkLogicOp = 9
	VkLogicOpInvert       VkLogicOp = 10
	VkLogicOpOrReverse    VkLogicOp = 11
	VkLogicOpCopyInverted VkLogicOp = 12
	VkLogicOpOrInverted   VkLogicOp = 13
	VkLogicOpNand         VkLogicOp = 14
	VkLogicOpSet          VkLogicOp = 15
)

type VkInternalAllocationType int64

var (
	VkInternalAllocationTypeExecutable VkInternalAllocationType = 0
)

type VkSystemAllocationScope int64

var (
	VkSystemAllocationScopeCommand  VkSystemAllocationScope = 0
	VkSystemAllocationScopeObject   VkSystemAllocationScope = 1
	VkSystemAllocationScopeCache    VkSystemAllocationScope = 2
	VkSystemAllocationScopeDevice   VkSystemAllocationScope = 3
	VkSystemAllocationScopeInstance VkSystemAllocationScope = 4
)

type VkPhysicalDeviceType int64

var (
	VkPhysicalDeviceTypeOther         VkPhysicalDeviceType = 0
	VkPhysicalDeviceTypeIntegratedGpu VkPhysicalDeviceType = 1
	VkPhysicalDeviceTypeDiscreteGpu   VkPhysicalDeviceType = 2
	VkPhysicalDeviceTypeVirtualGpu    VkPhysicalDeviceType = 3
	VkPhysicalDeviceTypeCpu           VkPhysicalDeviceType = 4
)

type VkVertexInputRate int64

var (
	VkVertexInputRateVertex   VkVertexInputRate = 0
	VkVertexInputRateInstance VkVertexInputRate = 1
)

type VkFormat int64

var (
	VkFormatUndefined                VkFormat = 0
	VkFormatR4g4UnormPack8           VkFormat = 1
	VkFormatR4g4b4a4UnormPack16      VkFormat = 2
	VkFormatB4g4r4a4UnormPack16      VkFormat = 3
	VkFormatR5g6b5UnormPack16        VkFormat = 4
	VkFormatB5g6r5UnormPack16        VkFormat = 5
	VkFormatR5g5b5a1UnormPack16      VkFormat = 6
	VkFormatB5g5r5a1UnormPack16      VkFormat = 7
	VkFormatA1r5g5b5UnormPack16      VkFormat = 8
	VkFormatR8Unorm                  VkFormat = 9
	VkFormatR8Snorm                  VkFormat = 10
	VkFormatR8Uscaled                VkFormat = 11
	VkFormatR8Sscaled                VkFormat = 12
	VkFormatR8Uint                   VkFormat = 13
	VkFormatR8Sint                   VkFormat = 14
	VkFormatR8Srgb                   VkFormat = 15
	VkFormatR8g8Unorm                VkFormat = 16
	VkFormatR8g8Snorm                VkFormat = 17
	VkFormatR8g8Uscaled              VkFormat = 18
	VkFormatR8g8Sscaled              VkFormat = 19
	VkFormatR8g8Uint                 VkFormat = 20
	VkFormatR8g8Sint                 VkFormat = 21
	VkFormatR8g8Srgb                 VkFormat = 22
	VkFormatR8g8b8Unorm              VkFormat = 23
	VkFormatR8g8b8Snorm              VkFormat = 24
	VkFormatR8g8b8Uscaled            VkFormat = 25
	VkFormatR8g8b8Sscaled            VkFormat = 26
	VkFormatR8g8b8Uint               VkFormat = 27
	VkFormatR8g8b8Sint               VkFormat = 28
	VkFormatR8g8b8Srgb               VkFormat = 29
	VkFormatB8g8r8Unorm              VkFormat = 30
	VkFormatB8g8r8Snorm              VkFormat = 31
	VkFormatB8g8r8Uscaled            VkFormat = 32
	VkFormatB8g8r8Sscaled            VkFormat = 33
	VkFormatB8g8r8Uint               VkFormat = 34
	VkFormatB8g8r8Sint               VkFormat = 35
	VkFormatB8g8r8Srgb               VkFormat = 36
	VkFormatR8g8b8a8Unorm            VkFormat = 37
	VkFormatR8g8b8a8Snorm            VkFormat = 38
	VkFormatR8g8b8a8Uscaled          VkFormat = 39
	VkFormatR8g8b8a8Sscaled          VkFormat = 40
	VkFormatR8g8b8a8Uint             VkFormat = 41
	VkFormatR8g8b8a8Sint             VkFormat = 42
	VkFormatR8g8b8a8Srgb             VkFormat = 43
	VkFormatB8g8r8a8Unorm            VkFormat = 44
	VkFormatB8g8r8a8Snorm            VkFormat = 45
	VkFormatB8g8r8a8Uscaled          VkFormat = 46
	VkFormatB8g8r8a8Sscaled          VkFormat = 47
	VkFormatB8g8r8a8Uint             VkFormat = 48
	VkFormatB8g8r8a8Sint             VkFormat = 49
	VkFormatB8g8r8a8Srgb             VkFormat = 50
	VkFormatA8b8g8r8UnormPack32      VkFormat = 51
	VkFormatA8b8g8r8SnormPack32      VkFormat = 52
	VkFormatA8b8g8r8UscaledPack32    VkFormat = 53
	VkFormatA8b8g8r8SscaledPack32    VkFormat = 54
	VkFormatA8b8g8r8UintPack32       VkFormat = 55
	VkFormatA8b8g8r8SintPack32       VkFormat = 56
	VkFormatA8b8g8r8SrgbPack32       VkFormat = 57
	VkFormatA2r10g10b10UnormPack32   VkFormat = 58
	VkFormatA2r10g10b10SnormPack32   VkFormat = 59
	VkFormatA2r10g10b10UscaledPack32 VkFormat = 60
	VkFormatA2r10g10b10SscaledPack32 VkFormat = 61
	VkFormatA2r10g10b10UintPack32    VkFormat = 62
	VkFormatA2r10g10b10SintPack32    VkFormat = 63
	VkFormatA2b10g10r10UnormPack32   VkFormat = 64
	VkFormatA2b10g10r10SnormPack32   VkFormat = 65
	VkFormatA2b10g10r10UscaledPack32 VkFormat = 66
	VkFormatA2b10g10r10SscaledPack32 VkFormat = 67
	VkFormatA2b10g10r10UintPack32    VkFormat = 68
	VkFormatA2b10g10r10SintPack32    VkFormat = 69
	VkFormatR16Unorm                 VkFormat = 70
	VkFormatR16Snorm                 VkFormat = 71
	VkFormatR16Uscaled               VkFormat = 72
	VkFormatR16Sscaled               VkFormat = 73
	VkFormatR16Uint                  VkFormat = 74
	VkFormatR16Sint                  VkFormat = 75
	VkFormatR16Sfloat                VkFormat = 76
	VkFormatR16g16Unorm              VkFormat = 77
	VkFormatR16g16Snorm              VkFormat = 78
	VkFormatR16g16Uscaled            VkFormat = 79
	VkFormatR16g16Sscaled            VkFormat = 80
	VkFormatR16g16Uint               VkFormat = 81
	VkFormatR16g16Sint               VkFormat = 82
	VkFormatR16g16Sfloat             VkFormat = 83
	VkFormatR16g16b16Unorm           VkFormat = 84
	VkFormatR16g16b16Snorm           VkFormat = 85
	VkFormatR16g16b16Uscaled         VkFormat = 86
	VkFormatR16g16b16Sscaled         VkFormat = 87
	VkFormatR16g16b16Uint            VkFormat = 88
	VkFormatR16g16b16Sint            VkFormat = 89
	VkFormatR16g16b16Sfloat          VkFormat = 90
	VkFormatR16g16b16a16Unorm        VkFormat = 91
	VkFormatR16g16b16a16Snorm        VkFormat = 92
	VkFormatR16g16b16a16Uscaled      VkFormat = 93
	VkFormatR16g16b16a16Sscaled      VkFormat = 94
	VkFormatR16g16b16a16Uint         VkFormat = 95
	VkFormatR16g16b16a16Sint         VkFormat = 96
	VkFormatR16g16b16a16Sfloat       VkFormat = 97
	VkFormatR32Uint                  VkFormat = 98
	VkFormatR32Sint                  VkFormat = 99
	VkFormatR32Sfloat                VkFormat = 100
	VkFormatR32g32Uint               VkFormat = 101
	VkFormatR32g32Sint               VkFormat = 102
	VkFormatR32g32Sfloat             VkFormat = 103
	VkFormatR32g32b32Uint            VkFormat = 104
	VkFormatR32g32b32Sint            VkFormat = 105
	VkFormatR32g32b32Sfloat          VkFormat = 106
	VkFormatR32g32b32a32Uint         VkFormat = 107
	VkFormatR32g32b32a32Sint         VkFormat = 108
	VkFormatR32g32b32a32Sfloat       VkFormat = 109
	VkFormatR64Uint                  VkFormat = 110
	VkFormatR64Sint                  VkFormat = 111
	VkFormatR64Sfloat                VkFormat = 112
	VkFormatR64g64Uint               VkFormat = 113
	VkFormatR64g64Sint               VkFormat = 114
	VkFormatR64g64Sfloat             VkFormat = 115
	VkFormatR64g64b64Uint            VkFormat = 116
	VkFormatR64g64b64Sint            VkFormat = 117
	VkFormatR64g64b64Sfloat          VkFormat = 118
	VkFormatR64g64b64a64Uint         VkFormat = 119
	VkFormatR64g64b64a64Sint         VkFormat = 120
	VkFormatR64g64b64a64Sfloat       VkFormat = 121
	VkFormatB10g11r11UfloatPack32    VkFormat = 122
	VkFormatE5b9g9r9UfloatPack32     VkFormat = 123
	VkFormatD16Unorm                 VkFormat = 124
	VkFormatX8D24UnormPack32         VkFormat = 125
	VkFormatD32Sfloat                VkFormat = 126
	VkFormatS8Uint                   VkFormat = 127
	VkFormatD16UnormS8Uint           VkFormat = 128
	VkFormatD24UnormS8Uint           VkFormat = 129
	VkFormatD32SfloatS8Uint          VkFormat = 130
	VkFormatBc1RgbUnormBlock         VkFormat = 131
	VkFormatBc1RgbSrgbBlock          VkFormat = 132
	VkFormatBc1RgbaUnormBlock        VkFormat = 133
	VkFormatBc1RgbaSrgbBlock         VkFormat = 134
	VkFormatBc2UnormBlock            VkFormat = 135
	VkFormatBc2SrgbBlock             VkFormat = 136
	VkFormatBc3UnormBlock            VkFormat = 137
	VkFormatBc3SrgbBlock             VkFormat = 138
	VkFormatBc4UnormBlock            VkFormat = 139
	VkFormatBc4SnormBlock            VkFormat = 140
	VkFormatBc5UnormBlock            VkFormat = 141
	VkFormatBc5SnormBlock            VkFormat = 142
	VkFormatBc6hUfloatBlock          VkFormat = 143
	VkFormatBc6hSfloatBlock          VkFormat = 144
	VkFormatBc7UnormBlock            VkFormat = 145
	VkFormatBc7SrgbBlock             VkFormat = 146
	VkFormatEtc2R8g8b8UnormBlock     VkFormat = 147
	VkFormatEtc2R8g8b8SrgbBlock      VkFormat = 148
	VkFormatEtc2R8g8b8a1UnormBlock   VkFormat = 149
	VkFormatEtc2R8g8b8a1SrgbBlock    VkFormat = 150
	VkFormatEtc2R8g8b8a8UnormBlock   VkFormat = 151
	VkFormatEtc2R8g8b8a8SrgbBlock    VkFormat = 152
	VkFormatEacR11UnormBlock         VkFormat = 153
	VkFormatEacR11SnormBlock         VkFormat = 154
	VkFormatEacR11g11UnormBlock      VkFormat = 155
	VkFormatEacR11g11SnormBlock      VkFormat = 156
	VkFormatAstc4x4UnormBlock        VkFormat = 157
	VkFormatAstc4x4SrgbBlock         VkFormat = 158
	VkFormatAstc5x4UnormBlock        VkFormat = 159
	VkFormatAstc5x4SrgbBlock         VkFormat = 160
	VkFormatAstc5x5UnormBlock        VkFormat = 161
	VkFormatAstc5x5SrgbBlock         VkFormat = 162
	VkFormatAstc6x5UnormBlock        VkFormat = 163
	VkFormatAstc6x5SrgbBlock         VkFormat = 164
	VkFormatAstc6x6UnormBlock        VkFormat = 165
	VkFormatAstc6x6SrgbBlock         VkFormat = 166
	VkFormatAstc8x5UnormBlock        VkFormat = 167
	VkFormatAstc8x5SrgbBlock         VkFormat = 168
	VkFormatAstc8x6UnormBlock        VkFormat = 169
	VkFormatAstc8x6SrgbBlock         VkFormat = 170
	VkFormatAstc8x8UnormBlock        VkFormat = 171
	VkFormatAstc8x8SrgbBlock         VkFormat = 172
	VkFormatAstc10x5UnormBlock       VkFormat = 173
	VkFormatAstc10x5SrgbBlock        VkFormat = 174
	VkFormatAstc10x6UnormBlock       VkFormat = 175
	VkFormatAstc10x6SrgbBlock        VkFormat = 176
	VkFormatAstc10x8UnormBlock       VkFormat = 177
	VkFormatAstc10x8SrgbBlock        VkFormat = 178
	VkFormatAstc10x10UnormBlock      VkFormat = 179
	VkFormatAstc10x10SrgbBlock       VkFormat = 180
	VkFormatAstc12x10UnormBlock      VkFormat = 181
	VkFormatAstc12x10SrgbBlock       VkFormat = 182
	VkFormatAstc12x12UnormBlock      VkFormat = 183
	VkFormatAstc12x12SrgbBlock       VkFormat = 184
)

type VkStructureType int64

var (
	VkStructureTypeApplicationInfo                      VkStructureType = 0
	VkStructureTypeInstanceCreateInfo                   VkStructureType = 1
	VkStructureTypeDeviceQueueCreateInfo                VkStructureType = 2
	VkStructureTypeDeviceCreateInfo                     VkStructureType = 3
	VkStructureTypeSubmitInfo                           VkStructureType = 4
	VkStructureTypeMemoryAllocateInfo                   VkStructureType = 5
	VkStructureTypeMappedMemoryRange                    VkStructureType = 6
	VkStructureTypeBindSparseInfo                       VkStructureType = 7
	VkStructureTypeFenceCreateInfo                      VkStructureType = 8
	VkStructureTypeSemaphoreCreateInfo                  VkStructureType = 9
	VkStructureTypeEventCreateInfo                      VkStructureType = 10
	VkStructureTypeQueryPoolCreateInfo                  VkStructureType = 11
	VkStructureTypeBufferCreateInfo                     VkStructureType = 12
	VkStructureTypeBufferViewCreateInfo                 VkStructureType = 13
	VkStructureTypeImageCreateInfo                      VkStructureType = 14
	VkStructureTypeImageViewCreateInfo                  VkStructureType = 15
	VkStructureTypeShaderModuleCreateInfo               VkStructureType = 16
	VkStructureTypePipelineCacheCreateInfo              VkStructureType = 17
	VkStructureTypePipelineShaderStageCreateInfo        VkStructureType = 18
	VkStructureTypePipelineVertexInputStateCreateInfo   VkStructureType = 19
	VkStructureTypePipelineInputAssemblyStateCreateInfo VkStructureType = 20
	VkStructureTypePipelineTessellationStateCreateInfo  VkStructureType = 21
	VkStructureTypePipelineViewportStateCreateInfo      VkStructureType = 22
	VkStructureTypePipelineRasterizationStateCreateInfo VkStructureType = 23
	VkStructureTypePipelineMultisampleStateCreateInfo   VkStructureType = 24
	VkStructureTypePipelineDepthStencilStateCreateInfo  VkStructureType = 25
	VkStructureTypePipelineColorBlendStateCreateInfo    VkStructureType = 26
	VkStructureTypePipelineDynamicStateCreateInfo       VkStructureType = 27
	VkStructureTypeGraphicsPipelineCreateInfo           VkStructureType = 28
	VkStructureTypeComputePipelineCreateInfo            VkStructureType = 29
	VkStructureTypePipelineLayoutCreateInfo             VkStructureType = 30
	VkStructureTypeSamplerCreateInfo                    VkStructureType = 31
	VkStructureTypeDescriptorSetLayoutCreateInfo        VkStructureType = 32
	VkStructureTypeDescriptorPoolCreateInfo             VkStructureType = 33
	VkStructureTypeDescriptorSetAllocateInfo            VkStructureType = 34
	VkStructureTypeWriteDescriptorSet                   VkStructureType = 35
	VkStructureTypeCopyDescriptorSet                    VkStructureType = 36
	VkStructureTypeFramebufferCreateInfo                VkStructureType = 37
	VkStructureTypeRenderPassCreateInfo                 VkStructureType = 38
	VkStructureTypeCommandPoolCreateInfo                VkStructureType = 39
	VkStructureTypeCommandBufferAllocateInfo            VkStructureType = 40
	VkStructureTypeCommandBufferInheritanceInfo         VkStructureType = 41
	VkStructureTypeCommandBufferBeginInfo               VkStructureType = 42
	VkStructureTypeRenderPassBeginInfo                  VkStructureType = 43
	VkStructureTypeBufferMemoryBarrier                  VkStructureType = 44
	VkStructureTypeImageMemoryBarrier                   VkStructureType = 45
	VkStructureTypeMemoryBarrier                        VkStructureType = 46
	VkStructureTypeLoaderInstanceCreateInfo             VkStructureType = 47
	VkStructureTypeLoaderDeviceCreateInfo               VkStructureType = 48
)

type VkSubpassContents int64

var (
	VkSubpassContentsInline                  VkSubpassContents = 0
	VkSubpassContentsSecondaryCommandBuffers VkSubpassContents = 1
)

type VkResult int64

var (
	VkSuccess                   VkResult = 0
	VkNotReady                  VkResult = 1
	VkTimeout                   VkResult = 2
	VkEventSet                  VkResult = 3
	VkEventReset                VkResult = 4
	VkIncomplete                VkResult = 5
	VkErrorOutOfHostMemory      VkResult = -1
	VkErrorOutOfDeviceMemory    VkResult = -2
	VkErrorInitializationFailed VkResult = -3
	VkErrorDeviceLost           VkResult = -4
	VkErrorMemoryMapFailed      VkResult = -5
	VkErrorLayerNotPresent      VkResult = -6
	VkErrorExtensionNotPresent  VkResult = -7
	VkErrorFeatureNotPresent    VkResult = -8
	VkErrorIncompatibleDriver   VkResult = -9
	VkErrorTooManyObjects       VkResult = -10
	VkErrorFormatNotSupported   VkResult = -11
	VkErrorFragmentedPool       VkResult = -12
	VkErrorUnknown              VkResult = -13
)

type VkDynamicState int64

var (
	VkDynamicStateViewport           VkDynamicState = 0
	VkDynamicStateScissor            VkDynamicState = 1
	VkDynamicStateLineWidth          VkDynamicState = 2
	VkDynamicStateDepthBias          VkDynamicState = 3
	VkDynamicStateBlendConstants     VkDynamicState = 4
	VkDynamicStateDepthBounds        VkDynamicState = 5
	VkDynamicStateStencilCompareMask VkDynamicState = 6
	VkDynamicStateStencilWriteMask   VkDynamicState = 7
	VkDynamicStateStencilReference   VkDynamicState = 8
)

type VkDescriptorUpdateTemplateType int64

var (
	VkDescriptorUpdateTemplateTypeDescriptorSet VkDescriptorUpdateTemplateType = 0
)

type VkObjectType int64

var (
	VkObjectTypeUnknown             VkObjectType = 0
	VkObjectTypeInstance            VkObjectType = 1
	VkObjectTypePhysicalDevice      VkObjectType = 2
	VkObjectTypeDevice              VkObjectType = 3
	VkObjectTypeQueue               VkObjectType = 4
	VkObjectTypeSemaphore           VkObjectType = 5
	VkObjectTypeCommandBuffer       VkObjectType = 6
	VkObjectTypeFence               VkObjectType = 7
	VkObjectTypeDeviceMemory        VkObjectType = 8
	VkObjectTypeBuffer              VkObjectType = 9
	VkObjectTypeImage               VkObjectType = 10
	VkObjectTypeEvent               VkObjectType = 11
	VkObjectTypeQueryPool           VkObjectType = 12
	VkObjectTypeBufferView          VkObjectType = 13
	VkObjectTypeImageView           VkObjectType = 14
	VkObjectTypeShaderModule        VkObjectType = 15
	VkObjectTypePipelineCache       VkObjectType = 16
	VkObjectTypePipelineLayout      VkObjectType = 17
	VkObjectTypeRenderPass          VkObjectType = 18
	VkObjectTypePipeline            VkObjectType = 19
	VkObjectTypeDescriptorSetLayout VkObjectType = 20
	VkObjectTypeSampler             VkObjectType = 21
	VkObjectTypeDescriptorPool      VkObjectType = 22
	VkObjectTypeDescriptorSet       VkObjectType = 23
	VkObjectTypeFramebuffer         VkObjectType = 24
	VkObjectTypeCommandPool         VkObjectType = 25
)

type VkRayTracingInvocationReorderModeEXT int64

var (
	VkRayTracingInvocationReorderModeNoneExt    VkRayTracingInvocationReorderModeEXT = 0
	VkRayTracingInvocationReorderModeReorderExt VkRayTracingInvocationReorderModeEXT = 1
)

type VkRayTracingLssIndexingModeNV int64

var (
	VkRayTracingLssIndexingModeListNv       VkRayTracingLssIndexingModeNV = 0
	VkRayTracingLssIndexingModeSuccessiveNv VkRayTracingLssIndexingModeNV = 1
)

type VkRayTracingLssPrimitiveEndCapsModeNV int64

var (
	VkRayTracingLssPrimitiveEndCapsModeNoneNv    VkRayTracingLssPrimitiveEndCapsModeNV = 0
	VkRayTracingLssPrimitiveEndCapsModeChainedNv VkRayTracingLssPrimitiveEndCapsModeNV = 1
)

type VkDirectDriverLoadingModeLUNARG int64

var (
	VkDirectDriverLoadingModeExclusiveLunarg VkDirectDriverLoadingModeLUNARG = 0
	VkDirectDriverLoadingModeInclusiveLunarg VkDirectDriverLoadingModeLUNARG = 1
)

type VkAntiLagModeAMD int64

var (
	VkAntiLagModeDriverControlAmd VkAntiLagModeAMD = 0
	VkAntiLagModeOnAmd            VkAntiLagModeAMD = 1
	VkAntiLagModeOffAmd           VkAntiLagModeAMD = 2
)

type VkAntiLagStageAMD int64

var (
	VkAntiLagStageInputAmd   VkAntiLagStageAMD = 0
	VkAntiLagStagePresentAmd VkAntiLagStageAMD = 1
)

type VkQueueFlagBits int64

var (
	VkQueueGraphicsBit      VkQueueFlagBits = (1 << 0)
	VkQueueComputeBit       VkQueueFlagBits = (1 << 1)
	VkQueueTransferBit      VkQueueFlagBits = (1 << 2)
	VkQueueSparseBindingBit VkQueueFlagBits = (1 << 3)
)

type VkCullModeFlagBits int64

var (
	VkCullModeNone         VkCullModeFlagBits = 0
	VkCullModeFrontBit     VkCullModeFlagBits = (1 << 0)
	VkCullModeBackBit      VkCullModeFlagBits = (1 << 1)
	VkCullModeFrontAndBack VkCullModeFlagBits = 0
)

type VkRenderPassCreateFlagBits int64

var ()

type VkDeviceQueueCreateFlagBits int64

var ()

type VkMemoryPropertyFlagBits int64

var (
	VkMemoryPropertyDeviceLocalBit     VkMemoryPropertyFlagBits = (1 << 0)
	VkMemoryPropertyHostVisibleBit     VkMemoryPropertyFlagBits = (1 << 1)
	VkMemoryPropertyHostCoherentBit    VkMemoryPropertyFlagBits = (1 << 2)
	VkMemoryPropertyHostCachedBit      VkMemoryPropertyFlagBits = (1 << 3)
	VkMemoryPropertyLazilyAllocatedBit VkMemoryPropertyFlagBits = (1 << 4)
)

type VkMemoryHeapFlagBits int64

var (
	VkMemoryHeapDeviceLocalBit VkMemoryHeapFlagBits = (1 << 0)
)

type VkAccessFlagBits int64

var (
	VkAccessIndirectCommandReadBit         VkAccessFlagBits = (1 << 0)
	VkAccessIndexReadBit                   VkAccessFlagBits = (1 << 1)
	VkAccessVertexAttributeReadBit         VkAccessFlagBits = (1 << 2)
	VkAccessUniformReadBit                 VkAccessFlagBits = (1 << 3)
	VkAccessInputAttachmentReadBit         VkAccessFlagBits = (1 << 4)
	VkAccessShaderReadBit                  VkAccessFlagBits = (1 << 5)
	VkAccessShaderWriteBit                 VkAccessFlagBits = (1 << 6)
	VkAccessColorAttachmentReadBit         VkAccessFlagBits = (1 << 7)
	VkAccessColorAttachmentWriteBit        VkAccessFlagBits = (1 << 8)
	VkAccessDepthStencilAttachmentReadBit  VkAccessFlagBits = (1 << 9)
	VkAccessDepthStencilAttachmentWriteBit VkAccessFlagBits = (1 << 10)
	VkAccessTransferReadBit                VkAccessFlagBits = (1 << 11)
	VkAccessTransferWriteBit               VkAccessFlagBits = (1 << 12)
	VkAccessHostReadBit                    VkAccessFlagBits = (1 << 13)
	VkAccessHostWriteBit                   VkAccessFlagBits = (1 << 14)
	VkAccessMemoryReadBit                  VkAccessFlagBits = (1 << 15)
	VkAccessMemoryWriteBit                 VkAccessFlagBits = (1 << 16)
)

type VkBufferUsageFlagBits int64

var (
	VkBufferUsageTransferSrcBit        VkBufferUsageFlagBits = (1 << 0)
	VkBufferUsageTransferDstBit        VkBufferUsageFlagBits = (1 << 1)
	VkBufferUsageUniformTexelBufferBit VkBufferUsageFlagBits = (1 << 2)
	VkBufferUsageStorageTexelBufferBit VkBufferUsageFlagBits = (1 << 3)
	VkBufferUsageUniformBufferBit      VkBufferUsageFlagBits = (1 << 4)
	VkBufferUsageStorageBufferBit      VkBufferUsageFlagBits = (1 << 5)
	VkBufferUsageIndexBufferBit        VkBufferUsageFlagBits = (1 << 6)
	VkBufferUsageVertexBufferBit       VkBufferUsageFlagBits = (1 << 7)
	VkBufferUsageIndirectBufferBit     VkBufferUsageFlagBits = (1 << 8)
)

type VkBufferUsageFlagBits2 int64

var (
	VkBufferUsage2TransferSrcBit        VkBufferUsageFlagBits2 = (1 << 0)
	VkBufferUsage2TransferDstBit        VkBufferUsageFlagBits2 = (1 << 1)
	VkBufferUsage2UniformTexelBufferBit VkBufferUsageFlagBits2 = (1 << 2)
	VkBufferUsage2StorageTexelBufferBit VkBufferUsageFlagBits2 = (1 << 3)
	VkBufferUsage2UniformBufferBit      VkBufferUsageFlagBits2 = (1 << 4)
	VkBufferUsage2StorageBufferBit      VkBufferUsageFlagBits2 = (1 << 5)
	VkBufferUsage2IndexBufferBit        VkBufferUsageFlagBits2 = (1 << 6)
	VkBufferUsage2VertexBufferBit       VkBufferUsageFlagBits2 = (1 << 7)
	VkBufferUsage2IndirectBufferBit     VkBufferUsageFlagBits2 = (1 << 8)
)

type VkBufferCreateFlagBits int64

var (
	VkBufferCreateSparseBindingBit   VkBufferCreateFlagBits = (1 << 0)
	VkBufferCreateSparseResidencyBit VkBufferCreateFlagBits = (1 << 1)
	VkBufferCreateSparseAliasedBit   VkBufferCreateFlagBits = (1 << 2)
)

type VkShaderStageFlagBits int64

var (
	VkShaderStageVertexBit                 VkShaderStageFlagBits = (1 << 0)
	VkShaderStageTessellationControlBit    VkShaderStageFlagBits = (1 << 1)
	VkShaderStageTessellationEvaluationBit VkShaderStageFlagBits = (1 << 2)
	VkShaderStageGeometryBit               VkShaderStageFlagBits = (1 << 3)
	VkShaderStageFragmentBit               VkShaderStageFlagBits = (1 << 4)
	VkShaderStageComputeBit                VkShaderStageFlagBits = (1 << 5)
	VkShaderStageAllGraphics               VkShaderStageFlagBits = 0
	VkShaderStageAll                       VkShaderStageFlagBits = 0
)

type VkImageUsageFlagBits int64

var (
	VkImageUsageTransferSrcBit            VkImageUsageFlagBits = (1 << 0)
	VkImageUsageTransferDstBit            VkImageUsageFlagBits = (1 << 1)
	VkImageUsageSampledBit                VkImageUsageFlagBits = (1 << 2)
	VkImageUsageStorageBit                VkImageUsageFlagBits = (1 << 3)
	VkImageUsageColorAttachmentBit        VkImageUsageFlagBits = (1 << 4)
	VkImageUsageDepthStencilAttachmentBit VkImageUsageFlagBits = (1 << 5)
	VkImageUsageTransientAttachmentBit    VkImageUsageFlagBits = (1 << 6)
	VkImageUsageInputAttachmentBit        VkImageUsageFlagBits = (1 << 7)
)

type VkImageCreateFlagBits int64

var (
	VkImageCreateSparseBindingBit   VkImageCreateFlagBits = (1 << 0)
	VkImageCreateSparseResidencyBit VkImageCreateFlagBits = (1 << 1)
	VkImageCreateSparseAliasedBit   VkImageCreateFlagBits = (1 << 2)
	VkImageCreateMutableFormatBit   VkImageCreateFlagBits = (1 << 3)
	VkImageCreateCubeCompatibleBit  VkImageCreateFlagBits = (1 << 4)
)

type VkImageViewCreateFlagBits int64

var ()

type VkSamplerCreateFlagBits int64

var ()

type VkPipelineCreateFlagBits int64

var (
	VkPipelineCreateDisableOptimizationBit VkPipelineCreateFlagBits = (1 << 0)
	VkPipelineCreateAllowDerivativesBit    VkPipelineCreateFlagBits = (1 << 1)
	VkPipelineCreateDerivativeBit          VkPipelineCreateFlagBits = (1 << 2)
)

type VkPipelineCreateFlagBits2 int64

var (
	VkPipelineCreate2DisableOptimizationBit           VkPipelineCreateFlagBits2 = (1 << 0)
	VkPipelineCreate2AllowDerivativesBit              VkPipelineCreateFlagBits2 = (1 << 1)
	VkPipelineCreate2DerivativeBit                    VkPipelineCreateFlagBits2 = (1 << 2)
	VkPipelineCreate2ViewIndexFromDeviceIndexBit      VkPipelineCreateFlagBits2 = (1 << 3)
	VkPipelineCreate2DispatchBaseBit                  VkPipelineCreateFlagBits2 = (1 << 4)
	VkPipelineCreate2FailOnPipelineCompileRequiredBit VkPipelineCreateFlagBits2 = (1 << 8)
	VkPipelineCreate2EarlyReturnOnFailureBit          VkPipelineCreateFlagBits2 = (1 << 9)
	VkPipelineCreate2NoProtectedAccessBit             VkPipelineCreateFlagBits2 = (1 << 27)
	VkPipelineCreate2ProtectedAccessOnlyBit           VkPipelineCreateFlagBits2 = (1 << 30)
)

type VkPipelineShaderStageCreateFlagBits int64

var ()

type VkColorComponentFlagBits int64

var (
	VkColorComponentRBit VkColorComponentFlagBits = (1 << 0)
	VkColorComponentGBit VkColorComponentFlagBits = (1 << 1)
	VkColorComponentBBit VkColorComponentFlagBits = (1 << 2)
	VkColorComponentABit VkColorComponentFlagBits = (1 << 3)
)

type VkFenceCreateFlagBits int64

var (
	VkFenceCreateSignaledBit VkFenceCreateFlagBits = (1 << 0)
)

type VkSemaphoreCreateFlagBits int64

var ()

type VkFormatFeatureFlagBits int64

var (
	VkFormatFeatureSampledImageBit             VkFormatFeatureFlagBits = (1 << 0)
	VkFormatFeatureStorageImageBit             VkFormatFeatureFlagBits = (1 << 1)
	VkFormatFeatureStorageImageAtomicBit       VkFormatFeatureFlagBits = (1 << 2)
	VkFormatFeatureUniformTexelBufferBit       VkFormatFeatureFlagBits = (1 << 3)
	VkFormatFeatureStorageTexelBufferBit       VkFormatFeatureFlagBits = (1 << 4)
	VkFormatFeatureStorageTexelBufferAtomicBit VkFormatFeatureFlagBits = (1 << 5)
	VkFormatFeatureVertexBufferBit             VkFormatFeatureFlagBits = (1 << 6)
	VkFormatFeatureColorAttachmentBit          VkFormatFeatureFlagBits = (1 << 7)
	VkFormatFeatureColorAttachmentBlendBit     VkFormatFeatureFlagBits = (1 << 8)
	VkFormatFeatureDepthStencilAttachmentBit   VkFormatFeatureFlagBits = (1 << 9)
	VkFormatFeatureBlitSrcBit                  VkFormatFeatureFlagBits = (1 << 10)
	VkFormatFeatureBlitDstBit                  VkFormatFeatureFlagBits = (1 << 11)
	VkFormatFeatureSampledImageFilterLinearBit VkFormatFeatureFlagBits = (1 << 12)
)

type VkQueryControlFlagBits int64

var (
	VkQueryControlPreciseBit VkQueryControlFlagBits = (1 << 0)
)

type VkQueryResultFlagBits int64

var (
	VkQueryResult64Bit               VkQueryResultFlagBits = (1 << 0)
	VkQueryResultWaitBit             VkQueryResultFlagBits = (1 << 1)
	VkQueryResultWithAvailabilityBit VkQueryResultFlagBits = (1 << 2)
	VkQueryResultPartialBit          VkQueryResultFlagBits = (1 << 3)
)

type VkCommandBufferUsageFlagBits int64

var (
	VkCommandBufferUsageOneTimeSubmitBit      VkCommandBufferUsageFlagBits = (1 << 0)
	VkCommandBufferUsageRenderPassContinueBit VkCommandBufferUsageFlagBits = (1 << 1)
	VkCommandBufferUsageSimultaneousUseBit    VkCommandBufferUsageFlagBits = (1 << 2)
)

type VkQueryPipelineStatisticFlagBits int64

var (
	VkQueryPipelineStatisticInputAssemblyVerticesBit                   VkQueryPipelineStatisticFlagBits = (1 << 0)
	VkQueryPipelineStatisticInputAssemblyPrimitivesBit                 VkQueryPipelineStatisticFlagBits = (1 << 1)
	VkQueryPipelineStatisticVertexShaderInvocationsBit                 VkQueryPipelineStatisticFlagBits = (1 << 2)
	VkQueryPipelineStatisticGeometryShaderInvocationsBit               VkQueryPipelineStatisticFlagBits = (1 << 3)
	VkQueryPipelineStatisticGeometryShaderPrimitivesBit                VkQueryPipelineStatisticFlagBits = (1 << 4)
	VkQueryPipelineStatisticClippingInvocationsBit                     VkQueryPipelineStatisticFlagBits = (1 << 5)
	VkQueryPipelineStatisticClippingPrimitivesBit                      VkQueryPipelineStatisticFlagBits = (1 << 6)
	VkQueryPipelineStatisticFragmentShaderInvocationsBit               VkQueryPipelineStatisticFlagBits = (1 << 7)
	VkQueryPipelineStatisticTessellationControlShaderPatchesBit        VkQueryPipelineStatisticFlagBits = (1 << 8)
	VkQueryPipelineStatisticTessellationEvaluationShaderInvocationsBit VkQueryPipelineStatisticFlagBits = (1 << 9)
	VkQueryPipelineStatisticComputeShaderInvocationsBit                VkQueryPipelineStatisticFlagBits = (1 << 10)
)

type VkMemoryMapFlagBits int64

var ()

type VkImageAspectFlagBits int64

var (
	VkImageAspectColorBit    VkImageAspectFlagBits = (1 << 0)
	VkImageAspectDepthBit    VkImageAspectFlagBits = (1 << 1)
	VkImageAspectStencilBit  VkImageAspectFlagBits = (1 << 2)
	VkImageAspectMetadataBit VkImageAspectFlagBits = (1 << 3)
)

type VkSparseImageFormatFlagBits int64

var (
	VkSparseImageFormatSingleMiptailBit        VkSparseImageFormatFlagBits = (1 << 0)
	VkSparseImageFormatAlignedMipSizeBit       VkSparseImageFormatFlagBits = (1 << 1)
	VkSparseImageFormatNonstandardBlockSizeBit VkSparseImageFormatFlagBits = (1 << 2)
)

type VkSparseMemoryBindFlagBits int64

var (
	VkSparseMemoryBindMetadataBit VkSparseMemoryBindFlagBits = (1 << 0)
)

type VkPipelineStageFlagBits int64

var (
	VkPipelineStageTopOfPipeBit                    VkPipelineStageFlagBits = (1 << 0)
	VkPipelineStageDrawIndirectBit                 VkPipelineStageFlagBits = (1 << 1)
	VkPipelineStageVertexInputBit                  VkPipelineStageFlagBits = (1 << 2)
	VkPipelineStageVertexShaderBit                 VkPipelineStageFlagBits = (1 << 3)
	VkPipelineStageTessellationControlShaderBit    VkPipelineStageFlagBits = (1 << 4)
	VkPipelineStageTessellationEvaluationShaderBit VkPipelineStageFlagBits = (1 << 5)
	VkPipelineStageGeometryShaderBit               VkPipelineStageFlagBits = (1 << 6)
	VkPipelineStageFragmentShaderBit               VkPipelineStageFlagBits = (1 << 7)
	VkPipelineStageEarlyFragmentTestsBit           VkPipelineStageFlagBits = (1 << 8)
	VkPipelineStageLateFragmentTestsBit            VkPipelineStageFlagBits = (1 << 9)
	VkPipelineStageColorAttachmentOutputBit        VkPipelineStageFlagBits = (1 << 10)
	VkPipelineStageComputeShaderBit                VkPipelineStageFlagBits = (1 << 11)
	VkPipelineStageTransferBit                     VkPipelineStageFlagBits = (1 << 12)
	VkPipelineStageBottomOfPipeBit                 VkPipelineStageFlagBits = (1 << 13)
	VkPipelineStageHostBit                         VkPipelineStageFlagBits = (1 << 14)
	VkPipelineStageAllGraphicsBit                  VkPipelineStageFlagBits = (1 << 15)
	VkPipelineStageAllCommandsBit                  VkPipelineStageFlagBits = (1 << 16)
)

type VkCommandPoolCreateFlagBits int64

var (
	VkCommandPoolCreateTransientBit          VkCommandPoolCreateFlagBits = (1 << 0)
	VkCommandPoolCreateResetCommandBufferBit VkCommandPoolCreateFlagBits = (1 << 1)
)

type VkCommandPoolResetFlagBits int64

var (
	VkCommandPoolResetReleaseResourcesBit VkCommandPoolResetFlagBits = (1 << 0)
)

type VkCommandBufferResetFlagBits int64

var (
	VkCommandBufferResetReleaseResourcesBit VkCommandBufferResetFlagBits = (1 << 0)
)

type VkSampleCountFlagBits int64

var (
	VkSampleCount1Bit  VkSampleCountFlagBits = (1 << 0)
	VkSampleCount2Bit  VkSampleCountFlagBits = (1 << 1)
	VkSampleCount4Bit  VkSampleCountFlagBits = (1 << 2)
	VkSampleCount8Bit  VkSampleCountFlagBits = (1 << 3)
	VkSampleCount16Bit VkSampleCountFlagBits = (1 << 4)
	VkSampleCount32Bit VkSampleCountFlagBits = (1 << 5)
	VkSampleCount64Bit VkSampleCountFlagBits = (1 << 6)
)

type VkAttachmentDescriptionFlagBits int64

var (
	VkAttachmentDescriptionMayAliasBit VkAttachmentDescriptionFlagBits = (1 << 0)
)

type VkStencilFaceFlagBits int64

var (
	VkStencilFaceFrontBit     VkStencilFaceFlagBits = (1 << 0)
	VkStencilFaceBackBit      VkStencilFaceFlagBits = (1 << 1)
	VkStencilFaceFrontAndBack VkStencilFaceFlagBits = 0
	VkStencilFrontAndBack     VkStencilFaceFlagBits = 0
)

type VkDescriptorPoolCreateFlagBits int64

var (
	VkDescriptorPoolCreateFreeDescriptorSetBit VkDescriptorPoolCreateFlagBits = (1 << 0)
)

type VkDependencyFlagBits int64

var (
	VkDependencyByRegionBit VkDependencyFlagBits = (1 << 0)
)

type VkSemaphoreType int64

var (
	VkSemaphoreTypeBinary   VkSemaphoreType = 0
	VkSemaphoreTypeTimeline VkSemaphoreType = 1
)

type VkSemaphoreWaitFlagBits int64

var (
	VkSemaphoreWaitAnyBit VkSemaphoreWaitFlagBits = (1 << 0)
)

type VkPresentModeKHR int64

var (
	VkPresentModeImmediateKhr   VkPresentModeKHR = 0
	VkPresentModeMailboxKhr     VkPresentModeKHR = 1
	VkPresentModeFifoKhr        VkPresentModeKHR = 2
	VkPresentModeFifoRelaxedKhr VkPresentModeKHR = 3
)

type VkColorSpaceKHR int64

var (
	VkColorSpaceSrgbNonlinearKhr VkColorSpaceKHR = 0
	VkColorspaceSrgbNonlinearKhr VkColorSpaceKHR = 0
)

type VkDisplayPlaneAlphaFlagBitsKHR int64

var (
	VkDisplayPlaneAlphaOpaqueBitKhr                VkDisplayPlaneAlphaFlagBitsKHR = (1 << 0)
	VkDisplayPlaneAlphaGlobalBitKhr                VkDisplayPlaneAlphaFlagBitsKHR = (1 << 1)
	VkDisplayPlaneAlphaPerPixelBitKhr              VkDisplayPlaneAlphaFlagBitsKHR = (1 << 2)
	VkDisplayPlaneAlphaPerPixelPremultipliedBitKhr VkDisplayPlaneAlphaFlagBitsKHR = (1 << 3)
)

type VkCompositeAlphaFlagBitsKHR int64

var (
	VkCompositeAlphaOpaqueBitKhr         VkCompositeAlphaFlagBitsKHR = (1 << 0)
	VkCompositeAlphaPreMultipliedBitKhr  VkCompositeAlphaFlagBitsKHR = (1 << 1)
	VkCompositeAlphaPostMultipliedBitKhr VkCompositeAlphaFlagBitsKHR = (1 << 2)
	VkCompositeAlphaInheritBitKhr        VkCompositeAlphaFlagBitsKHR = (1 << 3)
)

type VkSurfaceTransformFlagBitsKHR int64

var (
	VkSurfaceTransformIdentityBitKhr                  VkSurfaceTransformFlagBitsKHR = (1 << 0)
	VkSurfaceTransformRotate90BitKhr                  VkSurfaceTransformFlagBitsKHR = (1 << 1)
	VkSurfaceTransformRotate180BitKhr                 VkSurfaceTransformFlagBitsKHR = (1 << 2)
	VkSurfaceTransformRotate270BitKhr                 VkSurfaceTransformFlagBitsKHR = (1 << 3)
	VkSurfaceTransformHorizontalMirrorBitKhr          VkSurfaceTransformFlagBitsKHR = (1 << 4)
	VkSurfaceTransformHorizontalMirrorRotate90BitKhr  VkSurfaceTransformFlagBitsKHR = (1 << 5)
	VkSurfaceTransformHorizontalMirrorRotate180BitKhr VkSurfaceTransformFlagBitsKHR = (1 << 6)
	VkSurfaceTransformHorizontalMirrorRotate270BitKhr VkSurfaceTransformFlagBitsKHR = (1 << 7)
	VkSurfaceTransformInheritBitKhr                   VkSurfaceTransformFlagBitsKHR = (1 << 8)
)

type VkDisplaySurfaceStereoTypeNV int64

var (
	VkDisplaySurfaceStereoTypeNoneNv              VkDisplaySurfaceStereoTypeNV = 0
	VkDisplaySurfaceStereoTypeOnboardDinNv        VkDisplaySurfaceStereoTypeNV = 1
	VkDisplaySurfaceStereoTypeHdmi3dNv            VkDisplaySurfaceStereoTypeNV = 2
	VkDisplaySurfaceStereoTypeInbandDisplayportNv VkDisplaySurfaceStereoTypeNV = 3
)

type VkSwapchainImageUsageFlagBitsANDROID int64

var (
	VkSwapchainImageUsageSharedBitAndroid VkSwapchainImageUsageFlagBitsANDROID = (1 << 0)
)

type VkTimeDomainKHR int64

var (
	VkTimeDomainDeviceKhr                  VkTimeDomainKHR = 0
	VkTimeDomainClockMonotonicKhr          VkTimeDomainKHR = 1
	VkTimeDomainClockMonotonicRawKhr       VkTimeDomainKHR = 2
	VkTimeDomainQueryPerformanceCounterKhr VkTimeDomainKHR = 3
)

type VkDebugReportFlagBitsEXT int64

var (
	VkDebugReportInformationBitExt        VkDebugReportFlagBitsEXT = (1 << 0)
	VkDebugReportWarningBitExt            VkDebugReportFlagBitsEXT = (1 << 1)
	VkDebugReportPerformanceWarningBitExt VkDebugReportFlagBitsEXT = (1 << 2)
	VkDebugReportErrorBitExt              VkDebugReportFlagBitsEXT = (1 << 3)
	VkDebugReportDebugBitExt              VkDebugReportFlagBitsEXT = (1 << 4)
)

type VkDebugReportObjectTypeEXT int64

var (
	VkDebugReportObjectTypeUnknownExt                VkDebugReportObjectTypeEXT = 0
	VkDebugReportObjectTypeInstanceExt               VkDebugReportObjectTypeEXT = 1
	VkDebugReportObjectTypePhysicalDeviceExt         VkDebugReportObjectTypeEXT = 2
	VkDebugReportObjectTypeDeviceExt                 VkDebugReportObjectTypeEXT = 3
	VkDebugReportObjectTypeQueueExt                  VkDebugReportObjectTypeEXT = 4
	VkDebugReportObjectTypeSemaphoreExt              VkDebugReportObjectTypeEXT = 5
	VkDebugReportObjectTypeCommandBufferExt          VkDebugReportObjectTypeEXT = 6
	VkDebugReportObjectTypeFenceExt                  VkDebugReportObjectTypeEXT = 7
	VkDebugReportObjectTypeDeviceMemoryExt           VkDebugReportObjectTypeEXT = 8
	VkDebugReportObjectTypeBufferExt                 VkDebugReportObjectTypeEXT = 9
	VkDebugReportObjectTypeImageExt                  VkDebugReportObjectTypeEXT = 10
	VkDebugReportObjectTypeEventExt                  VkDebugReportObjectTypeEXT = 11
	VkDebugReportObjectTypeQueryPoolExt              VkDebugReportObjectTypeEXT = 12
	VkDebugReportObjectTypeBufferViewExt             VkDebugReportObjectTypeEXT = 13
	VkDebugReportObjectTypeImageViewExt              VkDebugReportObjectTypeEXT = 14
	VkDebugReportObjectTypeShaderModuleExt           VkDebugReportObjectTypeEXT = 15
	VkDebugReportObjectTypePipelineCacheExt          VkDebugReportObjectTypeEXT = 16
	VkDebugReportObjectTypePipelineLayoutExt         VkDebugReportObjectTypeEXT = 17
	VkDebugReportObjectTypeRenderPassExt             VkDebugReportObjectTypeEXT = 18
	VkDebugReportObjectTypePipelineExt               VkDebugReportObjectTypeEXT = 19
	VkDebugReportObjectTypeDescriptorSetLayoutExt    VkDebugReportObjectTypeEXT = 20
	VkDebugReportObjectTypeSamplerExt                VkDebugReportObjectTypeEXT = 21
	VkDebugReportObjectTypeDescriptorPoolExt         VkDebugReportObjectTypeEXT = 22
	VkDebugReportObjectTypeDescriptorSetExt          VkDebugReportObjectTypeEXT = 23
	VkDebugReportObjectTypeFramebufferExt            VkDebugReportObjectTypeEXT = 24
	VkDebugReportObjectTypeCommandPoolExt            VkDebugReportObjectTypeEXT = 25
	VkDebugReportObjectTypeSurfaceKhrExt             VkDebugReportObjectTypeEXT = 26
	VkDebugReportObjectTypeSwapchainKhrExt           VkDebugReportObjectTypeEXT = 27
	VkDebugReportObjectTypeDebugReportCallbackExtExt VkDebugReportObjectTypeEXT = 28
	VkDebugReportObjectTypeDebugReportExt            VkDebugReportObjectTypeEXT = 0
	VkDebugReportObjectTypeDisplayKhrExt             VkDebugReportObjectTypeEXT = 29
	VkDebugReportObjectTypeDisplayModeKhrExt         VkDebugReportObjectTypeEXT = 30
	VkDebugReportObjectTypeValidationCacheExtExt     VkDebugReportObjectTypeEXT = 33
	VkDebugReportObjectTypeValidationCacheExt        VkDebugReportObjectTypeEXT = 0
)

type VkDeviceMemoryReportEventTypeEXT int64

var (
	VkDeviceMemoryReportEventTypeAllocateExt         VkDeviceMemoryReportEventTypeEXT = 0
	VkDeviceMemoryReportEventTypeFreeExt             VkDeviceMemoryReportEventTypeEXT = 1
	VkDeviceMemoryReportEventTypeImportExt           VkDeviceMemoryReportEventTypeEXT = 2
	VkDeviceMemoryReportEventTypeUnimportExt         VkDeviceMemoryReportEventTypeEXT = 3
	VkDeviceMemoryReportEventTypeAllocationFailedExt VkDeviceMemoryReportEventTypeEXT = 4
)

type VkRasterizationOrderAMD int64

var (
	VkRasterizationOrderStrictAmd  VkRasterizationOrderAMD = 0
	VkRasterizationOrderRelaxedAmd VkRasterizationOrderAMD = 1
)

type VkExternalMemoryHandleTypeFlagBitsNV int64

var (
	VkExternalMemoryHandleTypeOpaqueWin32BitNv    VkExternalMemoryHandleTypeFlagBitsNV = (1 << 0)
	VkExternalMemoryHandleTypeOpaqueWin32KmtBitNv VkExternalMemoryHandleTypeFlagBitsNV = (1 << 1)
	VkExternalMemoryHandleTypeD3d11ImageBitNv     VkExternalMemoryHandleTypeFlagBitsNV = (1 << 2)
	VkExternalMemoryHandleTypeD3d11ImageKmtBitNv  VkExternalMemoryHandleTypeFlagBitsNV = (1 << 3)
)

type VkExternalMemoryFeatureFlagBitsNV int64

var (
	VkExternalMemoryFeatureDedicatedOnlyBitNv VkExternalMemoryFeatureFlagBitsNV = (1 << 0)
	VkExternalMemoryFeatureExportableBitNv    VkExternalMemoryFeatureFlagBitsNV = (1 << 1)
	VkExternalMemoryFeatureImportableBitNv    VkExternalMemoryFeatureFlagBitsNV = (1 << 2)
)

type VkClusterAccelerationStructureIndexFormatFlagBitsNV int64

var (
	VkClusterAccelerationStructureIndexFormat8bitNv  VkClusterAccelerationStructureIndexFormatFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureIndexFormat16bitNv VkClusterAccelerationStructureIndexFormatFlagBitsNV = (1 << 1)
	VkClusterAccelerationStructureIndexFormat32bitNv VkClusterAccelerationStructureIndexFormatFlagBitsNV = (1 << 2)
)

type VkClusterAccelerationStructureTypeNV int64

var (
	VkClusterAccelerationStructureTypeClustersBottomLevelNv     VkClusterAccelerationStructureTypeNV = 0
	VkClusterAccelerationStructureTypeTriangleClusterNv         VkClusterAccelerationStructureTypeNV = 1
	VkClusterAccelerationStructureTypeTriangleClusterTemplateNv VkClusterAccelerationStructureTypeNV = 2
)

type VkClusterAccelerationStructureOpTypeNV int64

var (
	VkClusterAccelerationStructureOpTypeMoveObjectsNv                  VkClusterAccelerationStructureOpTypeNV = 0
	VkClusterAccelerationStructureOpTypeBuildClustersBottomLevelNv     VkClusterAccelerationStructureOpTypeNV = 1
	VkClusterAccelerationStructureOpTypeBuildTriangleClusterNv         VkClusterAccelerationStructureOpTypeNV = 2
	VkClusterAccelerationStructureOpTypeBuildTriangleClusterTemplateNv VkClusterAccelerationStructureOpTypeNV = 3
	VkClusterAccelerationStructureOpTypeInstantiateTriangleClusterNv   VkClusterAccelerationStructureOpTypeNV = 4
	VkClusterAccelerationStructureOpTypeGetClusterTemplateIndicesNv    VkClusterAccelerationStructureOpTypeNV = 5
)

type VkClusterAccelerationStructureOpModeNV int64

var (
	VkClusterAccelerationStructureOpModeImplicitDestinationsNv VkClusterAccelerationStructureOpModeNV = 0
	VkClusterAccelerationStructureOpModeExplicitDestinationsNv VkClusterAccelerationStructureOpModeNV = 1
	VkClusterAccelerationStructureOpModeComputeSizesNv         VkClusterAccelerationStructureOpModeNV = 2
)

type VkClusterAccelerationStructureClusterFlagBitsNV int64

var (
	VkClusterAccelerationStructureClusterAllowDisableOpacityMicromapsNv VkClusterAccelerationStructureClusterFlagBitsNV = (1 << 0)
)

type VkClusterAccelerationStructureGeometryFlagBitsNV int64

var (
	VkClusterAccelerationStructureGeometryCullDisableBitNv                 VkClusterAccelerationStructureGeometryFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureGeometryNoDuplicateAnyhitInvocationBitNv VkClusterAccelerationStructureGeometryFlagBitsNV = (1 << 1)
	VkClusterAccelerationStructureGeometryOpaqueBitNv                      VkClusterAccelerationStructureGeometryFlagBitsNV = (1 << 2)
)

type VkClusterAccelerationStructureAddressResolutionFlagBitsNV int64

var (
	VkClusterAccelerationStructureAddressResolutionNoneNv                         VkClusterAccelerationStructureAddressResolutionFlagBitsNV = 0
	VkClusterAccelerationStructureAddressResolutionIndirectedDstImplicitDataBitNv VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 0)
	VkClusterAccelerationStructureAddressResolutionIndirectedScratchDataBitNv     VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 1)
	VkClusterAccelerationStructureAddressResolutionIndirectedDstAddressArrayBitNv VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 2)
	VkClusterAccelerationStructureAddressResolutionIndirectedDstSizesArrayBitNv   VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 3)
	VkClusterAccelerationStructureAddressResolutionIndirectedSrcInfosArrayBitNv   VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 4)
	VkClusterAccelerationStructureAddressResolutionIndirectedSrcInfosCountBitNv   VkClusterAccelerationStructureAddressResolutionFlagBitsNV = (1 << 5)
)

type VkValidationCheckEXT int64

var (
	VkValidationCheckAllExt     VkValidationCheckEXT = 0
	VkValidationCheckShadersExt VkValidationCheckEXT = 1
)

type VkValidationFeatureEnableEXT int64

var (
	VkValidationFeatureEnableGpuAssistedExt                   VkValidationFeatureEnableEXT = 0
	VkValidationFeatureEnableGpuAssistedReserveBindingSlotExt VkValidationFeatureEnableEXT = 1
	VkValidationFeatureEnableBestPracticesExt                 VkValidationFeatureEnableEXT = 2
	VkValidationFeatureEnableDebugPrintfExt                   VkValidationFeatureEnableEXT = 3
	VkValidationFeatureEnableSynchronizationValidationExt     VkValidationFeatureEnableEXT = 4
)

type VkValidationFeatureDisableEXT int64

var (
	VkValidationFeatureDisableAllExt                   VkValidationFeatureDisableEXT = 0
	VkValidationFeatureDisableShadersExt               VkValidationFeatureDisableEXT = 1
	VkValidationFeatureDisableThreadSafetyExt          VkValidationFeatureDisableEXT = 2
	VkValidationFeatureDisableApiParametersExt         VkValidationFeatureDisableEXT = 3
	VkValidationFeatureDisableObjectLifetimesExt       VkValidationFeatureDisableEXT = 4
	VkValidationFeatureDisableCoreChecksExt            VkValidationFeatureDisableEXT = 5
	VkValidationFeatureDisableUniqueHandlesExt         VkValidationFeatureDisableEXT = 6
	VkValidationFeatureDisableShaderValidationCacheExt VkValidationFeatureDisableEXT = 7
)

type VkLayerSettingTypeEXT int64

var (
	VkLayerSettingTypeBool32Ext  VkLayerSettingTypeEXT = 0
	VkLayerSettingTypeInt32Ext   VkLayerSettingTypeEXT = 1
	VkLayerSettingTypeInt64Ext   VkLayerSettingTypeEXT = 2
	VkLayerSettingTypeUint32Ext  VkLayerSettingTypeEXT = 3
	VkLayerSettingTypeUint64Ext  VkLayerSettingTypeEXT = 4
	VkLayerSettingTypeFloat32Ext VkLayerSettingTypeEXT = 5
	VkLayerSettingTypeFloat64Ext VkLayerSettingTypeEXT = 6
	VkLayerSettingTypeStringExt  VkLayerSettingTypeEXT = 7
)

type VkSubgroupFeatureFlagBits int64

var (
	VkSubgroupFeatureBasicBit           VkSubgroupFeatureFlagBits = (1 << 0)
	VkSubgroupFeatureVoteBit            VkSubgroupFeatureFlagBits = (1 << 1)
	VkSubgroupFeatureArithmeticBit      VkSubgroupFeatureFlagBits = (1 << 2)
	VkSubgroupFeatureBallotBit          VkSubgroupFeatureFlagBits = (1 << 3)
	VkSubgroupFeatureShuffleBit         VkSubgroupFeatureFlagBits = (1 << 4)
	VkSubgroupFeatureShuffleRelativeBit VkSubgroupFeatureFlagBits = (1 << 5)
	VkSubgroupFeatureClusteredBit       VkSubgroupFeatureFlagBits = (1 << 6)
	VkSubgroupFeatureQuadBit            VkSubgroupFeatureFlagBits = (1 << 7)
)

type VkIndirectCommandsLayoutUsageFlagBitsNV int64

var (
	VkIndirectCommandsLayoutUsageExplicitPreprocessBitNv VkIndirectCommandsLayoutUsageFlagBitsNV = (1 << 0)
	VkIndirectCommandsLayoutUsageIndexedSequencesBitNv   VkIndirectCommandsLayoutUsageFlagBitsNV = (1 << 1)
	VkIndirectCommandsLayoutUsageUnorderedSequencesBitNv VkIndirectCommandsLayoutUsageFlagBitsNV = (1 << 2)
)

type VkIndirectStateFlagBitsNV int64

var (
	VkIndirectStateFlagFrontfaceBitNv VkIndirectStateFlagBitsNV = (1 << 0)
)

type VkIndirectCommandsTokenTypeNV int64

var (
	VkIndirectCommandsTokenTypeShaderGroupNv  VkIndirectCommandsTokenTypeNV = 0
	VkIndirectCommandsTokenTypeStateFlagsNv   VkIndirectCommandsTokenTypeNV = 1
	VkIndirectCommandsTokenTypeIndexBufferNv  VkIndirectCommandsTokenTypeNV = 2
	VkIndirectCommandsTokenTypeVertexBufferNv VkIndirectCommandsTokenTypeNV = 3
	VkIndirectCommandsTokenTypePushConstantNv VkIndirectCommandsTokenTypeNV = 4
	VkIndirectCommandsTokenTypeDrawIndexedNv  VkIndirectCommandsTokenTypeNV = 5
	VkIndirectCommandsTokenTypeDrawNv         VkIndirectCommandsTokenTypeNV = 6
	VkIndirectCommandsTokenTypeDrawTasksNv    VkIndirectCommandsTokenTypeNV = 7
)

type VkPrivateDataSlotCreateFlagBits int64

var ()

type VkDescriptorSetLayoutCreateFlagBits int64

var ()

type VkExternalMemoryHandleTypeFlagBits int64

var (
	VkExternalMemoryHandleTypeOpaqueFdBit        VkExternalMemoryHandleTypeFlagBits = (1 << 0)
	VkExternalMemoryHandleTypeOpaqueWin32Bit     VkExternalMemoryHandleTypeFlagBits = (1 << 1)
	VkExternalMemoryHandleTypeOpaqueWin32KmtBit  VkExternalMemoryHandleTypeFlagBits = (1 << 2)
	VkExternalMemoryHandleTypeD3d11TextureBit    VkExternalMemoryHandleTypeFlagBits = (1 << 3)
	VkExternalMemoryHandleTypeD3d11TextureKmtBit VkExternalMemoryHandleTypeFlagBits = (1 << 4)
	VkExternalMemoryHandleTypeD3d12HeapBit       VkExternalMemoryHandleTypeFlagBits = (1 << 5)
	VkExternalMemoryHandleTypeD3d12ResourceBit   VkExternalMemoryHandleTypeFlagBits = (1 << 6)
)

type VkExternalMemoryFeatureFlagBits int64

var (
	VkExternalMemoryFeatureDedicatedOnlyBit VkExternalMemoryFeatureFlagBits = (1 << 0)
	VkExternalMemoryFeatureExportableBit    VkExternalMemoryFeatureFlagBits = (1 << 1)
	VkExternalMemoryFeatureImportableBit    VkExternalMemoryFeatureFlagBits = (1 << 2)
)

type VkExternalSemaphoreHandleTypeFlagBits int64

var (
	VkExternalSemaphoreHandleTypeOpaqueFdBit       VkExternalSemaphoreHandleTypeFlagBits = (1 << 0)
	VkExternalSemaphoreHandleTypeOpaqueWin32Bit    VkExternalSemaphoreHandleTypeFlagBits = (1 << 1)
	VkExternalSemaphoreHandleTypeOpaqueWin32KmtBit VkExternalSemaphoreHandleTypeFlagBits = (1 << 2)
	VkExternalSemaphoreHandleTypeD3d12FenceBit     VkExternalSemaphoreHandleTypeFlagBits = (1 << 3)
	VkExternalSemaphoreHandleTypeD3d11FenceBit     VkExternalSemaphoreHandleTypeFlagBits = 0
	VkExternalSemaphoreHandleTypeSyncFdBit         VkExternalSemaphoreHandleTypeFlagBits = (1 << 4)
)

type VkExternalSemaphoreFeatureFlagBits int64

var (
	VkExternalSemaphoreFeatureExportableBit VkExternalSemaphoreFeatureFlagBits = (1 << 0)
	VkExternalSemaphoreFeatureImportableBit VkExternalSemaphoreFeatureFlagBits = (1 << 1)
)

type VkSemaphoreImportFlagBits int64

var (
	VkSemaphoreImportTemporaryBit VkSemaphoreImportFlagBits = (1 << 0)
)

type VkExternalFenceHandleTypeFlagBits int64

var (
	VkExternalFenceHandleTypeOpaqueFdBit       VkExternalFenceHandleTypeFlagBits = (1 << 0)
	VkExternalFenceHandleTypeOpaqueWin32Bit    VkExternalFenceHandleTypeFlagBits = (1 << 1)
	VkExternalFenceHandleTypeOpaqueWin32KmtBit VkExternalFenceHandleTypeFlagBits = (1 << 2)
	VkExternalFenceHandleTypeSyncFdBit         VkExternalFenceHandleTypeFlagBits = (1 << 3)
)

type VkExternalFenceFeatureFlagBits int64

var (
	VkExternalFenceFeatureExportableBit VkExternalFenceFeatureFlagBits = (1 << 0)
	VkExternalFenceFeatureImportableBit VkExternalFenceFeatureFlagBits = (1 << 1)
)

type VkFenceImportFlagBits int64

var (
	VkFenceImportTemporaryBit VkFenceImportFlagBits = (1 << 0)
)

type VkSurfaceCounterFlagBitsEXT int64

var (
	VkSurfaceCounterVblankBitExt VkSurfaceCounterFlagBitsEXT = (1 << 0)
	VkSurfaceCounterVblankExt    VkSurfaceCounterFlagBitsEXT = 0
)

type VkDisplayPowerStateEXT int64

var (
	VkDisplayPowerStateOffExt     VkDisplayPowerStateEXT = 0
	VkDisplayPowerStateSuspendExt VkDisplayPowerStateEXT = 1
	VkDisplayPowerStateOnExt      VkDisplayPowerStateEXT = 2
)

type VkDeviceEventTypeEXT int64

var (
	VkDeviceEventTypeDisplayHotplugExt VkDeviceEventTypeEXT = 0
)

type VkDisplayEventTypeEXT int64

var (
	VkDisplayEventTypeFirstPixelOutExt VkDisplayEventTypeEXT = 0
)

type VkPeerMemoryFeatureFlagBits int64

var (
	VkPeerMemoryFeatureCopySrcBit    VkPeerMemoryFeatureFlagBits = (1 << 0)
	VkPeerMemoryFeatureCopyDstBit    VkPeerMemoryFeatureFlagBits = (1 << 1)
	VkPeerMemoryFeatureGenericSrcBit VkPeerMemoryFeatureFlagBits = (1 << 2)
	VkPeerMemoryFeatureGenericDstBit VkPeerMemoryFeatureFlagBits = (1 << 3)
)

type VkMemoryAllocateFlagBits int64

var (
	VkMemoryAllocateDeviceMaskBit VkMemoryAllocateFlagBits = (1 << 0)
)

type VkDeviceGroupPresentModeFlagBitsKHR int64

var (
	VkDeviceGroupPresentModeLocalBitKhr            VkDeviceGroupPresentModeFlagBitsKHR = (1 << 0)
	VkDeviceGroupPresentModeRemoteBitKhr           VkDeviceGroupPresentModeFlagBitsKHR = (1 << 1)
	VkDeviceGroupPresentModeSumBitKhr              VkDeviceGroupPresentModeFlagBitsKHR = (1 << 2)
	VkDeviceGroupPresentModeLocalMultiDeviceBitKhr VkDeviceGroupPresentModeFlagBitsKHR = (1 << 3)
)

type VkSwapchainCreateFlagBitsKHR int64

var ()

type VkViewportCoordinateSwizzleNV int64

var (
	VkViewportCoordinateSwizzlePositiveXNv VkViewportCoordinateSwizzleNV = 0
	VkViewportCoordinateSwizzleNegativeXNv VkViewportCoordinateSwizzleNV = 1
	VkViewportCoordinateSwizzlePositiveYNv VkViewportCoordinateSwizzleNV = 2
	VkViewportCoordinateSwizzleNegativeYNv VkViewportCoordinateSwizzleNV = 3
	VkViewportCoordinateSwizzlePositiveZNv VkViewportCoordinateSwizzleNV = 4
	VkViewportCoordinateSwizzleNegativeZNv VkViewportCoordinateSwizzleNV = 5
	VkViewportCoordinateSwizzlePositiveWNv VkViewportCoordinateSwizzleNV = 6
	VkViewportCoordinateSwizzleNegativeWNv VkViewportCoordinateSwizzleNV = 7
)

type VkDiscardRectangleModeEXT int64

var (
	VkDiscardRectangleModeInclusiveExt VkDiscardRectangleModeEXT = 0
	VkDiscardRectangleModeExclusiveExt VkDiscardRectangleModeEXT = 1
)

type VkSubpassDescriptionFlagBits int64

var ()

type VkPointClippingBehavior int64

var (
	VkPointClippingBehaviorAllClipPlanes      VkPointClippingBehavior = 0
	VkPointClippingBehaviorUserClipPlanesOnly VkPointClippingBehavior = 1
)

type VkSamplerReductionMode int64

var (
	VkSamplerReductionModeWeightedAverage VkSamplerReductionMode = 0
	VkSamplerReductionModeMin             VkSamplerReductionMode = 1
	VkSamplerReductionModeMax             VkSamplerReductionMode = 2
)

type VkTessellationDomainOrigin int64

var (
	VkTessellationDomainOriginUpperLeft VkTessellationDomainOrigin = 0
	VkTessellationDomainOriginLowerLeft VkTessellationDomainOrigin = 1
)

type VkSamplerYcbcrModelConversion int64

var (
	VkSamplerYcbcrModelConversionRgbIdentity   VkSamplerYcbcrModelConversion = 0
	VkSamplerYcbcrModelConversionYcbcrIdentity VkSamplerYcbcrModelConversion = 1
	VkSamplerYcbcrModelConversionYcbcr709      VkSamplerYcbcrModelConversion = 2
	VkSamplerYcbcrModelConversionYcbcr601      VkSamplerYcbcrModelConversion = 3
	VkSamplerYcbcrModelConversionYcbcr2020     VkSamplerYcbcrModelConversion = 4
)

type VkSamplerYcbcrRange int64

var (
	VkSamplerYcbcrRangeItuFull   VkSamplerYcbcrRange = 0
	VkSamplerYcbcrRangeItuNarrow VkSamplerYcbcrRange = 1
)

type VkChromaLocation int64

var (
	VkChromaLocationCositedEven VkChromaLocation = 0
	VkChromaLocationMidpoint    VkChromaLocation = 1
)

type VkBlendOverlapEXT int64

var (
	VkBlendOverlapUncorrelatedExt VkBlendOverlapEXT = 0
	VkBlendOverlapDisjointExt     VkBlendOverlapEXT = 1
	VkBlendOverlapConjointExt     VkBlendOverlapEXT = 2
)

type VkCoverageModulationModeNV int64

var (
	VkCoverageModulationModeNoneNv  VkCoverageModulationModeNV = 0
	VkCoverageModulationModeRgbNv   VkCoverageModulationModeNV = 1
	VkCoverageModulationModeAlphaNv VkCoverageModulationModeNV = 2
	VkCoverageModulationModeRgbaNv  VkCoverageModulationModeNV = 3
)

type VkCoverageReductionModeNV int64

var (
	VkCoverageReductionModeMergeNv    VkCoverageReductionModeNV = 0
	VkCoverageReductionModeTruncateNv VkCoverageReductionModeNV = 1
)

type VkValidationCacheHeaderVersionEXT int64

var (
	VkValidationCacheHeaderVersionOneExt VkValidationCacheHeaderVersionEXT = 1
)

type VkShaderInfoTypeAMD int64

var (
	VkShaderInfoTypeStatisticsAmd  VkShaderInfoTypeAMD = 0
	VkShaderInfoTypeBinaryAmd      VkShaderInfoTypeAMD = 1
	VkShaderInfoTypeDisassemblyAmd VkShaderInfoTypeAMD = 2
)

type VkQueueGlobalPriority int64

var (
	VkQueueGlobalPriorityLow      VkQueueGlobalPriority = 128
	VkQueueGlobalPriorityMedium   VkQueueGlobalPriority = 256
	VkQueueGlobalPriorityHigh     VkQueueGlobalPriority = 512
	VkQueueGlobalPriorityRealtime VkQueueGlobalPriority = 1024
)

type VkDebugUtilsMessageSeverityFlagBitsEXT int64

var (
	VkDebugUtilsMessageSeverityVerboseBitExt VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 0)
	VkDebugUtilsMessageSeverityInfoBitExt    VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 4)
	VkDebugUtilsMessageSeverityWarningBitExt VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 8)
	VkDebugUtilsMessageSeverityErrorBitExt   VkDebugUtilsMessageSeverityFlagBitsEXT = (1 << 12)
)

type VkDebugUtilsMessageTypeFlagBitsEXT int64

var (
	VkDebugUtilsMessageTypeGeneralBitExt     VkDebugUtilsMessageTypeFlagBitsEXT = (1 << 0)
	VkDebugUtilsMessageTypeValidationBitExt  VkDebugUtilsMessageTypeFlagBitsEXT = (1 << 1)
	VkDebugUtilsMessageTypePerformanceBitExt VkDebugUtilsMessageTypeFlagBitsEXT = (1 << 2)
)

type VkConservativeRasterizationModeEXT int64

var (
	VkConservativeRasterizationModeDisabledExt      VkConservativeRasterizationModeEXT = 0
	VkConservativeRasterizationModeOverestimateExt  VkConservativeRasterizationModeEXT = 1
	VkConservativeRasterizationModeUnderestimateExt VkConservativeRasterizationModeEXT = 2
)

type VkDescriptorBindingFlagBits int64

var (
	VkDescriptorBindingUpdateAfterBindBit          VkDescriptorBindingFlagBits = (1 << 0)
	VkDescriptorBindingUpdateUnusedWhilePendingBit VkDescriptorBindingFlagBits = (1 << 1)
	VkDescriptorBindingPartiallyBoundBit           VkDescriptorBindingFlagBits = (1 << 2)
	VkDescriptorBindingVariableDescriptorCountBit  VkDescriptorBindingFlagBits = (1 << 3)
)

type VkVendorId int64

var (
	VkVendorIdKhronos  VkVendorId = 0x10000
	VkVendorIdViv      VkVendorId = 0x10001
	VkVendorIdVsi      VkVendorId = 0x10002
	VkVendorIdKazan    VkVendorId = 0x10003
	VkVendorIdCodeplay VkVendorId = 0x10004
	VkVendorIdMesa     VkVendorId = 0x10005
	VkVendorIdPocl     VkVendorId = 0x10006
	VkVendorIdMobileye VkVendorId = 0x10007
)

type VkDriverId int64

var (
	VkDriverIdAmdProprietary            VkDriverId = 1
	VkDriverIdAmdOpenSource             VkDriverId = 2
	VkDriverIdMesaRadv                  VkDriverId = 3
	VkDriverIdNvidiaProprietary         VkDriverId = 4
	VkDriverIdIntelProprietaryWindows   VkDriverId = 5
	VkDriverIdIntelOpenSourceMesa       VkDriverId = 6
	VkDriverIdImaginationProprietary    VkDriverId = 7
	VkDriverIdQualcommProprietary       VkDriverId = 8
	VkDriverIdArmProprietary            VkDriverId = 9
	VkDriverIdGoogleSwiftshader         VkDriverId = 10
	VkDriverIdGgpProprietary            VkDriverId = 11
	VkDriverIdBroadcomProprietary       VkDriverId = 12
	VkDriverIdMesaLlvmpipe              VkDriverId = 13
	VkDriverIdMoltenvk                  VkDriverId = 14
	VkDriverIdCoreaviProprietary        VkDriverId = 15
	VkDriverIdJuiceProprietary          VkDriverId = 16
	VkDriverIdVerisiliconProprietary    VkDriverId = 17
	VkDriverIdMesaTurnip                VkDriverId = 18
	VkDriverIdMesaV3dv                  VkDriverId = 19
	VkDriverIdMesaPanvk                 VkDriverId = 20
	VkDriverIdSamsungProprietary        VkDriverId = 21
	VkDriverIdMesaVenus                 VkDriverId = 22
	VkDriverIdMesaDozen                 VkDriverId = 23
	VkDriverIdMesaNvk                   VkDriverId = 24
	VkDriverIdImaginationOpenSourceMesa VkDriverId = 25
	VkDriverIdMesaHoneykrisp            VkDriverId = 26
	VkDriverIdVulkanScEmulationOnVulkan VkDriverId = 27
	VkDriverIdMesaKosmickrisp           VkDriverId = 28
)

type VkConditionalRenderingFlagBitsEXT int64

var (
	VkConditionalRenderingInvertedBitExt VkConditionalRenderingFlagBitsEXT = (1 << 0)
)

type VkResolveModeFlagBits int64

var (
	VkResolveModeNone          VkResolveModeFlagBits = 0
	VkResolveModeSampleZeroBit VkResolveModeFlagBits = (1 << 0)
	VkResolveModeAverageBit    VkResolveModeFlagBits = (1 << 1)
	VkResolveModeMinBit        VkResolveModeFlagBits = (1 << 2)
	VkResolveModeMaxBit        VkResolveModeFlagBits = (1 << 3)
)

type VkShadingRatePaletteEntryNV int64

var (
	VkShadingRatePaletteEntryNoInvocationsNv           VkShadingRatePaletteEntryNV = 0
	VkShadingRatePaletteEntry16InvocationsPerPixelNv   VkShadingRatePaletteEntryNV = 1
	VkShadingRatePaletteEntry8InvocationsPerPixelNv    VkShadingRatePaletteEntryNV = 2
	VkShadingRatePaletteEntry4InvocationsPerPixelNv    VkShadingRatePaletteEntryNV = 3
	VkShadingRatePaletteEntry2InvocationsPerPixelNv    VkShadingRatePaletteEntryNV = 4
	VkShadingRatePaletteEntry1InvocationPerPixelNv     VkShadingRatePaletteEntryNV = 5
	VkShadingRatePaletteEntry1InvocationPer2x1PixelsNv VkShadingRatePaletteEntryNV = 6
	VkShadingRatePaletteEntry1InvocationPer1x2PixelsNv VkShadingRatePaletteEntryNV = 7
	VkShadingRatePaletteEntry1InvocationPer2x2PixelsNv VkShadingRatePaletteEntryNV = 8
	VkShadingRatePaletteEntry1InvocationPer4x2PixelsNv VkShadingRatePaletteEntryNV = 9
	VkShadingRatePaletteEntry1InvocationPer2x4PixelsNv VkShadingRatePaletteEntryNV = 10
	VkShadingRatePaletteEntry1InvocationPer4x4PixelsNv VkShadingRatePaletteEntryNV = 11
)

type VkCoarseSampleOrderTypeNV int64

var (
	VkCoarseSampleOrderTypeDefaultNv     VkCoarseSampleOrderTypeNV = 0
	VkCoarseSampleOrderTypeCustomNv      VkCoarseSampleOrderTypeNV = 1
	VkCoarseSampleOrderTypePixelMajorNv  VkCoarseSampleOrderTypeNV = 2
	VkCoarseSampleOrderTypeSampleMajorNv VkCoarseSampleOrderTypeNV = 3
)

type VkGeometryInstanceFlagBitsKHR int64

var (
	VkGeometryInstanceTriangleFacingCullDisableBitKhr     VkGeometryInstanceFlagBitsKHR = (1 << 0)
	VkGeometryInstanceTriangleFlipFacingBitKhr            VkGeometryInstanceFlagBitsKHR = (1 << 1)
	VkGeometryInstanceForceOpaqueBitKhr                   VkGeometryInstanceFlagBitsKHR = (1 << 2)
	VkGeometryInstanceForceNoOpaqueBitKhr                 VkGeometryInstanceFlagBitsKHR = (1 << 3)
	VkGeometryInstanceTriangleFrontCounterclockwiseBitKhr VkGeometryInstanceFlagBitsKHR = 0
)

type VkGeometryFlagBitsKHR int64

var (
	VkGeometryOpaqueBitKhr                      VkGeometryFlagBitsKHR = (1 << 0)
	VkGeometryNoDuplicateAnyHitInvocationBitKhr VkGeometryFlagBitsKHR = (1 << 1)
)

type VkBuildAccelerationStructureFlagBitsKHR int64

var (
	VkBuildAccelerationStructureAllowUpdateBitKhr     VkBuildAccelerationStructureFlagBitsKHR = (1 << 0)
	VkBuildAccelerationStructureAllowCompactionBitKhr VkBuildAccelerationStructureFlagBitsKHR = (1 << 1)
	VkBuildAccelerationStructurePreferFastTraceBitKhr VkBuildAccelerationStructureFlagBitsKHR = (1 << 2)
	VkBuildAccelerationStructurePreferFastBuildBitKhr VkBuildAccelerationStructureFlagBitsKHR = (1 << 3)
	VkBuildAccelerationStructureLowMemoryBitKhr       VkBuildAccelerationStructureFlagBitsKHR = (1 << 4)
)

type VkAccelerationStructureCreateFlagBitsKHR int64

var (
	VkAccelerationStructureCreateDeviceAddressCaptureReplayBitKhr VkAccelerationStructureCreateFlagBitsKHR = (1 << 0)
)

type VkCopyAccelerationStructureModeKHR int64

var (
	VkCopyAccelerationStructureModeCloneKhr   VkCopyAccelerationStructureModeKHR = 0
	VkCopyAccelerationStructureModeCompactKhr VkCopyAccelerationStructureModeKHR = 1
)

type VkBuildAccelerationStructureModeKHR int64

var (
	VkBuildAccelerationStructureModeBuildKhr  VkBuildAccelerationStructureModeKHR = 0
	VkBuildAccelerationStructureModeUpdateKhr VkBuildAccelerationStructureModeKHR = 1
)

type VkAccelerationStructureTypeKHR int64

var (
	VkAccelerationStructureTypeTopLevelKhr    VkAccelerationStructureTypeKHR = 0
	VkAccelerationStructureTypeBottomLevelKhr VkAccelerationStructureTypeKHR = 1
	VkAccelerationStructureTypeGenericKhr     VkAccelerationStructureTypeKHR = 2
)

type VkGeometryTypeKHR int64

var (
	VkGeometryTypeTrianglesKhr VkGeometryTypeKHR = 0
	VkGeometryTypeAabbsKhr     VkGeometryTypeKHR = 1
	VkGeometryTypeInstancesKhr VkGeometryTypeKHR = 2
)

type VkAccelerationStructureMemoryRequirementsTypeNV int64

var (
	VkAccelerationStructureMemoryRequirementsTypeObjectNv        VkAccelerationStructureMemoryRequirementsTypeNV = 0
	VkAccelerationStructureMemoryRequirementsTypeBuildScratchNv  VkAccelerationStructureMemoryRequirementsTypeNV = 1
	VkAccelerationStructureMemoryRequirementsTypeUpdateScratchNv VkAccelerationStructureMemoryRequirementsTypeNV = 2
)

type VkAccelerationStructureBuildTypeKHR int64

var (
	VkAccelerationStructureBuildTypeHostKhr         VkAccelerationStructureBuildTypeKHR = 0
	VkAccelerationStructureBuildTypeDeviceKhr       VkAccelerationStructureBuildTypeKHR = 1
	VkAccelerationStructureBuildTypeHostOrDeviceKhr VkAccelerationStructureBuildTypeKHR = 2
)

type VkRayTracingShaderGroupTypeKHR int64

var (
	VkRayTracingShaderGroupTypeGeneralKhr            VkRayTracingShaderGroupTypeKHR = 0
	VkRayTracingShaderGroupTypeTrianglesHitGroupKhr  VkRayTracingShaderGroupTypeKHR = 1
	VkRayTracingShaderGroupTypeProceduralHitGroupKhr VkRayTracingShaderGroupTypeKHR = 2
)

type VkAccelerationStructureCompatibilityKHR int64

var (
	VkAccelerationStructureCompatibilityCompatibleKhr   VkAccelerationStructureCompatibilityKHR = 0
	VkAccelerationStructureCompatibilityIncompatibleKhr VkAccelerationStructureCompatibilityKHR = 1
)

type VkShaderGroupShaderKHR int64

var (
	VkShaderGroupShaderGeneralKhr      VkShaderGroupShaderKHR = 0
	VkShaderGroupShaderClosestHitKhr   VkShaderGroupShaderKHR = 1
	VkShaderGroupShaderAnyHitKhr       VkShaderGroupShaderKHR = 2
	VkShaderGroupShaderIntersectionKhr VkShaderGroupShaderKHR = 3
)

type VkMemoryOverallocationBehaviorAMD int64

var (
	VkMemoryOverallocationBehaviorDefaultAmd    VkMemoryOverallocationBehaviorAMD = 0
	VkMemoryOverallocationBehaviorAllowedAmd    VkMemoryOverallocationBehaviorAMD = 1
	VkMemoryOverallocationBehaviorDisallowedAmd VkMemoryOverallocationBehaviorAMD = 2
)

type VkFramebufferCreateFlagBits int64

var ()

type VkQueryPoolCreateFlagBits int64

var ()

type VkDeviceDiagnosticsConfigFlagBitsNV int64

var (
	VkDeviceDiagnosticsConfigEnableShaderDebugInfoBitNv      VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 0)
	VkDeviceDiagnosticsConfigEnableResourceTrackingBitNv     VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 1)
	VkDeviceDiagnosticsConfigEnableAutomaticCheckpointsBitNv VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 2)
	VkDeviceDiagnosticsConfigEnableShaderErrorReportingBitNv VkDeviceDiagnosticsConfigFlagBitsNV = (1 << 3)
)

type VkPipelineCreationFeedbackFlagBits int64

var (
	VkPipelineCreationFeedbackValidBit                       VkPipelineCreationFeedbackFlagBits = (1 << 0)
	VkPipelineCreationFeedbackApplicationPipelineCacheHitBit VkPipelineCreationFeedbackFlagBits = (1 << 1)
	VkPipelineCreationFeedbackBasePipelineAccelerationBit    VkPipelineCreationFeedbackFlagBits = (1 << 2)
)

type VkFullScreenExclusiveEXT int64

var (
	VkFullScreenExclusiveDefaultExt               VkFullScreenExclusiveEXT = 0
	VkFullScreenExclusiveAllowedExt               VkFullScreenExclusiveEXT = 1
	VkFullScreenExclusiveDisallowedExt            VkFullScreenExclusiveEXT = 2
	VkFullScreenExclusiveApplicationControlledExt VkFullScreenExclusiveEXT = 3
)

type VkPerformanceCounterScopeKHR int64

var (
	VkPerformanceCounterScopeCommandBufferKhr VkPerformanceCounterScopeKHR = 0
	VkPerformanceCounterScopeRenderPassKhr    VkPerformanceCounterScopeKHR = 1
	VkPerformanceCounterScopeCommandKhr       VkPerformanceCounterScopeKHR = 2
	VkQueryScopeCommandBufferKhr              VkPerformanceCounterScopeKHR = 0
	VkQueryScopeRenderPassKhr                 VkPerformanceCounterScopeKHR = 0
	VkQueryScopeCommandKhr                    VkPerformanceCounterScopeKHR = 0
)

type VkMemoryDecompressionMethodFlagBitsEXT int64

var (
	VkMemoryDecompressionMethodGdeflate10BitExt VkMemoryDecompressionMethodFlagBitsEXT = (1 << 0)
	VkMemoryDecompressionMethodGdeflate10BitNv  VkMemoryDecompressionMethodFlagBitsEXT = 0
)

type VkPerformanceCounterUnitKHR int64

var (
	VkPerformanceCounterUnitGenericKhr        VkPerformanceCounterUnitKHR = 0
	VkPerformanceCounterUnitPercentageKhr     VkPerformanceCounterUnitKHR = 1
	VkPerformanceCounterUnitNanosecondsKhr    VkPerformanceCounterUnitKHR = 2
	VkPerformanceCounterUnitBytesKhr          VkPerformanceCounterUnitKHR = 3
	VkPerformanceCounterUnitBytesPerSecondKhr VkPerformanceCounterUnitKHR = 4
	VkPerformanceCounterUnitKelvinKhr         VkPerformanceCounterUnitKHR = 5
	VkPerformanceCounterUnitWattsKhr          VkPerformanceCounterUnitKHR = 6
	VkPerformanceCounterUnitVoltsKhr          VkPerformanceCounterUnitKHR = 7
	VkPerformanceCounterUnitAmpsKhr           VkPerformanceCounterUnitKHR = 8
	VkPerformanceCounterUnitHertzKhr          VkPerformanceCounterUnitKHR = 9
	VkPerformanceCounterUnitCyclesKhr         VkPerformanceCounterUnitKHR = 10
)

type VkPerformanceCounterStorageKHR int64

var (
	VkPerformanceCounterStorageInt32Khr   VkPerformanceCounterStorageKHR = 0
	VkPerformanceCounterStorageInt64Khr   VkPerformanceCounterStorageKHR = 1
	VkPerformanceCounterStorageUint32Khr  VkPerformanceCounterStorageKHR = 2
	VkPerformanceCounterStorageUint64Khr  VkPerformanceCounterStorageKHR = 3
	VkPerformanceCounterStorageFloat32Khr VkPerformanceCounterStorageKHR = 4
	VkPerformanceCounterStorageFloat64Khr VkPerformanceCounterStorageKHR = 5
)

type VkPerformanceCounterDescriptionFlagBitsKHR int64

var (
	VkPerformanceCounterDescriptionPerformanceImpactingBitKhr VkPerformanceCounterDescriptionFlagBitsKHR = (1 << 0)
	VkPerformanceCounterDescriptionPerformanceImpactingKhr    VkPerformanceCounterDescriptionFlagBitsKHR = 0
	VkPerformanceCounterDescriptionConcurrentlyImpactedBitKhr VkPerformanceCounterDescriptionFlagBitsKHR = (1 << 1)
	VkPerformanceCounterDescriptionConcurrentlyImpactedKhr    VkPerformanceCounterDescriptionFlagBitsKHR = 0
)

type VkAcquireProfilingLockFlagBitsKHR int64

var ()

type VkShaderCorePropertiesFlagBitsAMD int64

var ()

type VkRefreshObjectFlagBitsKHR int64

var ()

type VkPerformanceConfigurationTypeINTEL int64

var (
	VkPerformanceConfigurationTypeCommandQueueMetricsDiscoveryActivatedIntel VkPerformanceConfigurationTypeINTEL = 0
)

type VkQueryPoolSamplingModeINTEL int64

var (
	VkQueryPoolSamplingModeManualIntel VkQueryPoolSamplingModeINTEL = 0
)

type VkPerformanceOverrideTypeINTEL int64

var (
	VkPerformanceOverrideTypeNullHardwareIntel   VkPerformanceOverrideTypeINTEL = 0
	VkPerformanceOverrideTypeFlushGpuCachesIntel VkPerformanceOverrideTypeINTEL = 1
)

type VkPerformanceParameterTypeINTEL int64

var (
	VkPerformanceParameterTypeHwCountersSupportedIntel   VkPerformanceParameterTypeINTEL = 0
	VkPerformanceParameterTypeStreamMarkerValidBitsIntel VkPerformanceParameterTypeINTEL = 1
)

type VkPerformanceValueTypeINTEL int64

var (
	VkPerformanceValueTypeUint32Intel VkPerformanceValueTypeINTEL = 0
	VkPerformanceValueTypeUint64Intel VkPerformanceValueTypeINTEL = 1
	VkPerformanceValueTypeFloatIntel  VkPerformanceValueTypeINTEL = 2
	VkPerformanceValueTypeBoolIntel   VkPerformanceValueTypeINTEL = 3
	VkPerformanceValueTypeStringIntel VkPerformanceValueTypeINTEL = 4
)

type VkShaderFloatControlsIndependence int64

var (
	VkShaderFloatControlsIndependence32BitOnly VkShaderFloatControlsIndependence = 0
	VkShaderFloatControlsIndependenceAll       VkShaderFloatControlsIndependence = 1
	VkShaderFloatControlsIndependenceNone      VkShaderFloatControlsIndependence = 2
)

type VkPipelineExecutableStatisticFormatKHR int64

var (
	VkPipelineExecutableStatisticFormatBool32Khr  VkPipelineExecutableStatisticFormatKHR = 0
	VkPipelineExecutableStatisticFormatInt64Khr   VkPipelineExecutableStatisticFormatKHR = 1
	VkPipelineExecutableStatisticFormatUint64Khr  VkPipelineExecutableStatisticFormatKHR = 2
	VkPipelineExecutableStatisticFormatFloat64Khr VkPipelineExecutableStatisticFormatKHR = 3
)

type VkLineRasterizationMode int64

var (
	VkLineRasterizationModeDefault           VkLineRasterizationMode = 0
	VkLineRasterizationModeRectangular       VkLineRasterizationMode = 1
	VkLineRasterizationModeBresenham         VkLineRasterizationMode = 2
	VkLineRasterizationModeRectangularSmooth VkLineRasterizationMode = 3
)

type VkShaderModuleCreateFlagBits int64

var ()

type VkPipelineCompilerControlFlagBitsAMD int64

var ()

type VkFaultLevel int64

var (
	VkFaultLevelUnassigned  VkFaultLevel = 0
	VkFaultLevelCritical    VkFaultLevel = 1
	VkFaultLevelRecoverable VkFaultLevel = 2
	VkFaultLevelWarning     VkFaultLevel = 3
)

type VkFaultType int64

var (
	VkFaultTypeInvalid           VkFaultType = 0
	VkFaultTypeUnassigned        VkFaultType = 1
	VkFaultTypeImplementation    VkFaultType = 2
	VkFaultTypeSystem            VkFaultType = 3
	VkFaultTypePhysicalDevice    VkFaultType = 4
	VkFaultTypeCommandBufferFull VkFaultType = 5
	VkFaultTypeInvalidApiUsage   VkFaultType = 6
)

type VkFaultQueryBehavior int64

var (
	VkFaultQueryBehaviorGetAndClearAllFaults VkFaultQueryBehavior = 0
)

type VkToolPurposeFlagBits int64

var (
	VkToolPurposeValidationBit         VkToolPurposeFlagBits = (1 << 0)
	VkToolPurposeProfilingBit          VkToolPurposeFlagBits = (1 << 1)
	VkToolPurposeTracingBit            VkToolPurposeFlagBits = (1 << 2)
	VkToolPurposeAdditionalFeaturesBit VkToolPurposeFlagBits = (1 << 3)
	VkToolPurposeModifyingFeaturesBit  VkToolPurposeFlagBits = (1 << 4)
)

type VkPipelineMatchControl int64

var (
	VkPipelineMatchControlApplicationUuidExactMatch VkPipelineMatchControl = 0
)

type VkFragmentShadingRateCombinerOpKHR int64

var (
	VkFragmentShadingRateCombinerOpKeepKhr    VkFragmentShadingRateCombinerOpKHR = 0
	VkFragmentShadingRateCombinerOpReplaceKhr VkFragmentShadingRateCombinerOpKHR = 1
	VkFragmentShadingRateCombinerOpMinKhr     VkFragmentShadingRateCombinerOpKHR = 2
	VkFragmentShadingRateCombinerOpMaxKhr     VkFragmentShadingRateCombinerOpKHR = 3
	VkFragmentShadingRateCombinerOpMulKhr     VkFragmentShadingRateCombinerOpKHR = 4
)

type VkFragmentShadingRateNV int64

var (
	VkFragmentShadingRate1InvocationPerPixelNv     VkFragmentShadingRateNV = 0
	VkFragmentShadingRate1InvocationPer1x2PixelsNv VkFragmentShadingRateNV = 1
	VkFragmentShadingRate1InvocationPer2x1PixelsNv VkFragmentShadingRateNV = 4
	VkFragmentShadingRate1InvocationPer2x2PixelsNv VkFragmentShadingRateNV = 5
	VkFragmentShadingRate1InvocationPer2x4PixelsNv VkFragmentShadingRateNV = 6
	VkFragmentShadingRate1InvocationPer4x2PixelsNv VkFragmentShadingRateNV = 9
	VkFragmentShadingRate1InvocationPer4x4PixelsNv VkFragmentShadingRateNV = 10
	VkFragmentShadingRate2InvocationsPerPixelNv    VkFragmentShadingRateNV = 11
	VkFragmentShadingRate4InvocationsPerPixelNv    VkFragmentShadingRateNV = 12
	VkFragmentShadingRate8InvocationsPerPixelNv    VkFragmentShadingRateNV = 13
	VkFragmentShadingRate16InvocationsPerPixelNv   VkFragmentShadingRateNV = 14
	VkFragmentShadingRateNoInvocationsNv           VkFragmentShadingRateNV = 15
)

type VkFragmentShadingRateTypeNV int64

var (
	VkFragmentShadingRateTypeFragmentSizeNv VkFragmentShadingRateTypeNV = 0
	VkFragmentShadingRateTypeEnumsNv        VkFragmentShadingRateTypeNV = 1
)

type VkSubpassMergeStatusEXT int64

var (
	VkSubpassMergeStatusMergedExt                               VkSubpassMergeStatusEXT = 0
	VkSubpassMergeStatusDisallowedExt                           VkSubpassMergeStatusEXT = 1
	VkSubpassMergeStatusNotMergedSideEffectsExt                 VkSubpassMergeStatusEXT = 2
	VkSubpassMergeStatusNotMergedSamplesMismatchExt             VkSubpassMergeStatusEXT = 3
	VkSubpassMergeStatusNotMergedViewsMismatchExt               VkSubpassMergeStatusEXT = 4
	VkSubpassMergeStatusNotMergedAliasingExt                    VkSubpassMergeStatusEXT = 5
	VkSubpassMergeStatusNotMergedDependenciesExt                VkSubpassMergeStatusEXT = 6
	VkSubpassMergeStatusNotMergedIncompatibleInputAttachmentExt VkSubpassMergeStatusEXT = 7
	VkSubpassMergeStatusNotMergedTooManyAttachmentsExt          VkSubpassMergeStatusEXT = 8
	VkSubpassMergeStatusNotMergedInsufficientStorageExt         VkSubpassMergeStatusEXT = 9
	VkSubpassMergeStatusNotMergedDepthStencilCountExt           VkSubpassMergeStatusEXT = 10
	VkSubpassMergeStatusNotMergedResolveAttachmentReuseExt      VkSubpassMergeStatusEXT = 11
	VkSubpassMergeStatusNotMergedSingleSubpassExt               VkSubpassMergeStatusEXT = 12
	VkSubpassMergeStatusNotMergedUnspecifiedExt                 VkSubpassMergeStatusEXT = 13
)

type VkAccessFlagBits2 int64

var (
	VkAccess2None                           VkAccessFlagBits2 = 0
	VkAccess2IndirectCommandReadBit         VkAccessFlagBits2 = (1 << 0)
	VkAccess2IndexReadBit                   VkAccessFlagBits2 = (1 << 1)
	VkAccess2VertexAttributeReadBit         VkAccessFlagBits2 = (1 << 2)
	VkAccess2UniformReadBit                 VkAccessFlagBits2 = (1 << 3)
	VkAccess2InputAttachmentReadBit         VkAccessFlagBits2 = (1 << 4)
	VkAccess2ShaderReadBit                  VkAccessFlagBits2 = (1 << 5)
	VkAccess2ShaderWriteBit                 VkAccessFlagBits2 = (1 << 6)
	VkAccess2ColorAttachmentReadBit         VkAccessFlagBits2 = (1 << 7)
	VkAccess2ColorAttachmentWriteBit        VkAccessFlagBits2 = (1 << 8)
	VkAccess2DepthStencilAttachmentReadBit  VkAccessFlagBits2 = (1 << 9)
	VkAccess2DepthStencilAttachmentWriteBit VkAccessFlagBits2 = (1 << 10)
	VkAccess2TransferReadBit                VkAccessFlagBits2 = (1 << 11)
	VkAccess2TransferWriteBit               VkAccessFlagBits2 = (1 << 12)
	VkAccess2HostReadBit                    VkAccessFlagBits2 = (1 << 13)
	VkAccess2HostWriteBit                   VkAccessFlagBits2 = (1 << 14)
	VkAccess2MemoryReadBit                  VkAccessFlagBits2 = (1 << 15)
	VkAccess2MemoryWriteBit                 VkAccessFlagBits2 = (1 << 16)
	VkAccess2ShaderSampledReadBit           VkAccessFlagBits2 = (1 << 32)
	VkAccess2ShaderStorageReadBit           VkAccessFlagBits2 = (1 << 33)
	VkAccess2ShaderStorageWriteBit          VkAccessFlagBits2 = (1 << 34)
)

type VkPipelineStageFlagBits2 int64

var (
	VkPipelineStage2None                            VkPipelineStageFlagBits2 = 0
	VkPipelineStage2TopOfPipeBit                    VkPipelineStageFlagBits2 = (1 << 0)
	VkPipelineStage2DrawIndirectBit                 VkPipelineStageFlagBits2 = (1 << 1)
	VkPipelineStage2VertexInputBit                  VkPipelineStageFlagBits2 = (1 << 2)
	VkPipelineStage2VertexShaderBit                 VkPipelineStageFlagBits2 = (1 << 3)
	VkPipelineStage2TessellationControlShaderBit    VkPipelineStageFlagBits2 = (1 << 4)
	VkPipelineStage2TessellationEvaluationShaderBit VkPipelineStageFlagBits2 = (1 << 5)
	VkPipelineStage2GeometryShaderBit               VkPipelineStageFlagBits2 = (1 << 6)
	VkPipelineStage2FragmentShaderBit               VkPipelineStageFlagBits2 = (1 << 7)
	VkPipelineStage2EarlyFragmentTestsBit           VkPipelineStageFlagBits2 = (1 << 8)
	VkPipelineStage2LateFragmentTestsBit            VkPipelineStageFlagBits2 = (1 << 9)
	VkPipelineStage2ColorAttachmentOutputBit        VkPipelineStageFlagBits2 = (1 << 10)
	VkPipelineStage2ComputeShaderBit                VkPipelineStageFlagBits2 = (1 << 11)
	VkPipelineStage2AllTransferBit                  VkPipelineStageFlagBits2 = (1 << 12)
	VkPipelineStage2TransferBit                     VkPipelineStageFlagBits2 = 0
	VkPipelineStage2BottomOfPipeBit                 VkPipelineStageFlagBits2 = (1 << 13)
	VkPipelineStage2HostBit                         VkPipelineStageFlagBits2 = (1 << 14)
	VkPipelineStage2AllGraphicsBit                  VkPipelineStageFlagBits2 = (1 << 15)
	VkPipelineStage2AllCommandsBit                  VkPipelineStageFlagBits2 = (1 << 16)
	VkPipelineStage2CopyBit                         VkPipelineStageFlagBits2 = (1 << 32)
	VkPipelineStage2ResolveBit                      VkPipelineStageFlagBits2 = (1 << 33)
	VkPipelineStage2BlitBit                         VkPipelineStageFlagBits2 = (1 << 34)
	VkPipelineStage2ClearBit                        VkPipelineStageFlagBits2 = (1 << 35)
	VkPipelineStage2IndexInputBit                   VkPipelineStageFlagBits2 = (1 << 36)
	VkPipelineStage2VertexAttributeInputBit         VkPipelineStageFlagBits2 = (1 << 37)
	VkPipelineStage2PreRasterizationShadersBit      VkPipelineStageFlagBits2 = (1 << 38)
)

type VkSubmitFlagBits int64

var (
	VkSubmitProtectedBit VkSubmitFlagBits = (1 << 0)
)

type VkEventCreateFlagBits int64

var ()

type VkPipelineLayoutCreateFlagBits int64

var ()

type VkSciSyncClientTypeNV int64

var (
	VkSciSyncClientTypeSignalerNv       VkSciSyncClientTypeNV = 0
	VkSciSyncClientTypeWaiterNv         VkSciSyncClientTypeNV = 1
	VkSciSyncClientTypeSignalerWaiterNv VkSciSyncClientTypeNV = 2
)

type VkSciSyncPrimitiveTypeNV int64

var (
	VkSciSyncPrimitiveTypeFenceNv     VkSciSyncPrimitiveTypeNV = 0
	VkSciSyncPrimitiveTypeSemaphoreNv VkSciSyncPrimitiveTypeNV = 1
)

type VkProvokingVertexModeEXT int64

var (
	VkProvokingVertexModeFirstVertexExt VkProvokingVertexModeEXT = 0
	VkProvokingVertexModeLastVertexExt  VkProvokingVertexModeEXT = 1
)

type VkPipelineCacheValidationVersion int64

var (
	VkPipelineCacheValidationVersionSafetyCriticalOne VkPipelineCacheValidationVersion = 1
)

type VkAccelerationStructureMotionInstanceTypeNV int64

var (
	VkAccelerationStructureMotionInstanceTypeStaticNv       VkAccelerationStructureMotionInstanceTypeNV = 0
	VkAccelerationStructureMotionInstanceTypeMatrixMotionNv VkAccelerationStructureMotionInstanceTypeNV = 1
	VkAccelerationStructureMotionInstanceTypeSrtMotionNv    VkAccelerationStructureMotionInstanceTypeNV = 2
)

type VkPipelineColorBlendStateCreateFlagBits int64

var ()

type VkPipelineDepthStencilStateCreateFlagBits int64

var ()

type VkGraphicsPipelineLibraryFlagBitsEXT int64

var (
	VkGraphicsPipelineLibraryVertexInputInterfaceBitExt    VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 0)
	VkGraphicsPipelineLibraryPreRasterizationShadersBitExt VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 1)
	VkGraphicsPipelineLibraryFragmentShaderBitExt          VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 2)
	VkGraphicsPipelineLibraryFragmentOutputInterfaceBitExt VkGraphicsPipelineLibraryFlagBitsEXT = (1 << 3)
)

type VkRenderingAttachmentFlagBitsKHR int64

var ()

type VkResolveImageFlagBitsKHR int64

var ()

type VkDeviceAddressBindingFlagBitsEXT int64

var (
	VkDeviceAddressBindingInternalObjectBitExt VkDeviceAddressBindingFlagBitsEXT = (1 << 0)
)

type VkDeviceAddressBindingTypeEXT int64

var (
	VkDeviceAddressBindingTypeBindExt   VkDeviceAddressBindingTypeEXT = 0
	VkDeviceAddressBindingTypeUnbindExt VkDeviceAddressBindingTypeEXT = 1
)

type VkFrameBoundaryFlagBitsEXT int64

var (
	VkFrameBoundaryFrameEndBitExt VkFrameBoundaryFlagBitsEXT = (1 << 0)
)

type VkPresentScalingFlagBitsKHR int64

var (
	VkPresentScalingOneToOneBitKhr           VkPresentScalingFlagBitsKHR = (1 << 0)
	VkPresentScalingOneToOneBitExt           VkPresentScalingFlagBitsKHR = 0
	VkPresentScalingAspectRatioStretchBitKhr VkPresentScalingFlagBitsKHR = (1 << 1)
	VkPresentScalingAspectRatioStretchBitExt VkPresentScalingFlagBitsKHR = 0
	VkPresentScalingStretchBitKhr            VkPresentScalingFlagBitsKHR = (1 << 2)
	VkPresentScalingStretchBitExt            VkPresentScalingFlagBitsKHR = 0
)

type VkPresentGravityFlagBitsKHR int64

var (
	VkPresentGravityMinBitKhr      VkPresentGravityFlagBitsKHR = (1 << 0)
	VkPresentGravityMinBitExt      VkPresentGravityFlagBitsKHR = 0
	VkPresentGravityMaxBitKhr      VkPresentGravityFlagBitsKHR = (1 << 1)
	VkPresentGravityMaxBitExt      VkPresentGravityFlagBitsKHR = 0
	VkPresentGravityCenteredBitKhr VkPresentGravityFlagBitsKHR = (1 << 2)
	VkPresentGravityCenteredBitExt VkPresentGravityFlagBitsKHR = 0
)

type VkPhysicalDeviceSchedulingControlsFlagBitsARM int64

var (
	VkPhysicalDeviceSchedulingControlsShaderCoreCountArm VkPhysicalDeviceSchedulingControlsFlagBitsARM = (1 << 0)
)

type VkPresentStageFlagBitsEXT int64

var (
	VkPresentStageQueueOperationsEndBitExt     VkPresentStageFlagBitsEXT = (1 << 0)
	VkPresentStageRequestDequeuedBitExt        VkPresentStageFlagBitsEXT = (1 << 1)
	VkPresentStageImageFirstPixelOutBitExt     VkPresentStageFlagBitsEXT = (1 << 2)
	VkPresentStageImageFirstPixelVisibleBitExt VkPresentStageFlagBitsEXT = (1 << 3)
)

type VkPastPresentationTimingFlagBitsEXT int64

var (
	VkPastPresentationTimingAllowPartialResultsBitExt    VkPastPresentationTimingFlagBitsEXT = (1 << 0)
	VkPastPresentationTimingAllowOutOfOrderResultsBitExt VkPastPresentationTimingFlagBitsEXT = (1 << 1)
)

type VkPresentTimingInfoFlagBitsEXT int64

var (
	VkPresentTimingInfoPresentAtRelativeTimeBitExt        VkPresentTimingInfoFlagBitsEXT = (1 << 0)
	VkPresentTimingInfoPresentAtNearestRefreshCycleBitExt VkPresentTimingInfoFlagBitsEXT = (1 << 1)
)

type VkVideoCodecOperationFlagBitsKHR int64

var (
	VkVideoCodecOperationNoneKhr VkVideoCodecOperationFlagBitsKHR = 0
)

type VkVideoChromaSubsamplingFlagBitsKHR int64

var (
	VkVideoChromaSubsamplingInvalidKhr       VkVideoChromaSubsamplingFlagBitsKHR = 0
	VkVideoChromaSubsamplingMonochromeBitKhr VkVideoChromaSubsamplingFlagBitsKHR = (1 << 0)
	VkVideoChromaSubsampling420BitKhr        VkVideoChromaSubsamplingFlagBitsKHR = (1 << 1)
	VkVideoChromaSubsampling422BitKhr        VkVideoChromaSubsamplingFlagBitsKHR = (1 << 2)
	VkVideoChromaSubsampling444BitKhr        VkVideoChromaSubsamplingFlagBitsKHR = (1 << 3)
)

type VkVideoComponentBitDepthFlagBitsKHR int64

var (
	VkVideoComponentBitDepthInvalidKhr VkVideoComponentBitDepthFlagBitsKHR = 0
	VkVideoComponentBitDepth8BitKhr    VkVideoComponentBitDepthFlagBitsKHR = (1 << 0)
	VkVideoComponentBitDepth10BitKhr   VkVideoComponentBitDepthFlagBitsKHR = (1 << 2)
	VkVideoComponentBitDepth12BitKhr   VkVideoComponentBitDepthFlagBitsKHR = (1 << 4)
)

type VkVideoCapabilityFlagBitsKHR int64

var (
	VkVideoCapabilityProtectedContentBitKhr        VkVideoCapabilityFlagBitsKHR = (1 << 0)
	VkVideoCapabilitySeparateReferenceImagesBitKhr VkVideoCapabilityFlagBitsKHR = (1 << 1)
)

type VkVideoSessionCreateFlagBitsKHR int64

var (
	VkVideoSessionCreateProtectedContentBitKhr VkVideoSessionCreateFlagBitsKHR = (1 << 0)
)

type VkVideoSessionParametersCreateFlagBitsKHR int64

var ()

type VkVideoDecodeH264PictureLayoutFlagBitsKHR int64

var (
	VkVideoDecodeH264PictureLayoutProgressiveKhr                   VkVideoDecodeH264PictureLayoutFlagBitsKHR = 0
	VkVideoDecodeH264PictureLayoutInterlacedInterleavedLinesBitKhr VkVideoDecodeH264PictureLayoutFlagBitsKHR = (1 << 0)
	VkVideoDecodeH264PictureLayoutInterlacedSeparatePlanesBitKhr   VkVideoDecodeH264PictureLayoutFlagBitsKHR = (1 << 1)
)

type VkVideoCodingControlFlagBitsKHR int64

var (
	VkVideoCodingControlResetBitKhr VkVideoCodingControlFlagBitsKHR = (1 << 0)
)

type VkQueryResultStatusKHR int64

var (
	VkQueryResultStatusErrorKhr    VkQueryResultStatusKHR = -1
	VkQueryResultStatusNotReadyKhr VkQueryResultStatusKHR = 0
	VkQueryResultStatusCompleteKhr VkQueryResultStatusKHR = 1
)

type VkVideoDecodeUsageFlagBitsKHR int64

var (
	VkVideoDecodeUsageDefaultKhr        VkVideoDecodeUsageFlagBitsKHR = 0
	VkVideoDecodeUsageTranscodingBitKhr VkVideoDecodeUsageFlagBitsKHR = (1 << 0)
	VkVideoDecodeUsageOfflineBitKhr     VkVideoDecodeUsageFlagBitsKHR = (1 << 1)
	VkVideoDecodeUsageStreamingBitKhr   VkVideoDecodeUsageFlagBitsKHR = (1 << 2)
)

type VkVideoDecodeCapabilityFlagBitsKHR int64

var (
	VkVideoDecodeCapabilityDpbAndOutputCoincideBitKhr VkVideoDecodeCapabilityFlagBitsKHR = (1 << 0)
	VkVideoDecodeCapabilityDpbAndOutputDistinctBitKhr VkVideoDecodeCapabilityFlagBitsKHR = (1 << 1)
)

type VkVideoEncodeFlagBitsKHR int64

var ()

type VkVideoEncodeUsageFlagBitsKHR int64

var (
	VkVideoEncodeUsageDefaultKhr         VkVideoEncodeUsageFlagBitsKHR = 0
	VkVideoEncodeUsageTranscodingBitKhr  VkVideoEncodeUsageFlagBitsKHR = (1 << 0)
	VkVideoEncodeUsageStreamingBitKhr    VkVideoEncodeUsageFlagBitsKHR = (1 << 1)
	VkVideoEncodeUsageRecordingBitKhr    VkVideoEncodeUsageFlagBitsKHR = (1 << 2)
	VkVideoEncodeUsageConferencingBitKhr VkVideoEncodeUsageFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeContentFlagBitsKHR int64

var (
	VkVideoEncodeContentDefaultKhr     VkVideoEncodeContentFlagBitsKHR = 0
	VkVideoEncodeContentCameraBitKhr   VkVideoEncodeContentFlagBitsKHR = (1 << 0)
	VkVideoEncodeContentDesktopBitKhr  VkVideoEncodeContentFlagBitsKHR = (1 << 1)
	VkVideoEncodeContentRenderedBitKhr VkVideoEncodeContentFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeTuningModeKHR int64

var (
	VkVideoEncodeTuningModeDefaultKhr         VkVideoEncodeTuningModeKHR = 0
	VkVideoEncodeTuningModeHighQualityKhr     VkVideoEncodeTuningModeKHR = 1
	VkVideoEncodeTuningModeLowLatencyKhr      VkVideoEncodeTuningModeKHR = 2
	VkVideoEncodeTuningModeUltraLowLatencyKhr VkVideoEncodeTuningModeKHR = 3
	VkVideoEncodeTuningModeLosslessKhr        VkVideoEncodeTuningModeKHR = 4
)

type VkVideoEncodeCapabilityFlagBitsKHR int64

var (
	VkVideoEncodeCapabilityPrecedingExternallyEncodedBytesBitKhr           VkVideoEncodeCapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeCapabilityInsufficientBitstreamBufferRangeDetectionBitKhr VkVideoEncodeCapabilityFlagBitsKHR = (1 << 1)
)

type VkVideoEncodeFeedbackFlagBitsKHR int64

var (
	VkVideoEncodeFeedbackBitstreamBufferOffsetBitKhr VkVideoEncodeFeedbackFlagBitsKHR = (1 << 0)
	VkVideoEncodeFeedbackBitstreamBytesWrittenBitKhr VkVideoEncodeFeedbackFlagBitsKHR = (1 << 1)
	VkVideoEncodeFeedbackBitstreamHasOverridesBitKhr VkVideoEncodeFeedbackFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeRateControlModeFlagBitsKHR int64

var (
	VkVideoEncodeRateControlModeDefaultKhr     VkVideoEncodeRateControlModeFlagBitsKHR = 0
	VkVideoEncodeRateControlModeDisabledBitKhr VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 0)
	VkVideoEncodeRateControlModeCbrBitKhr      VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 1)
	VkVideoEncodeRateControlModeVbrBitKhr      VkVideoEncodeRateControlModeFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeIntraRefreshModeFlagBitsKHR int64

var (
	VkVideoEncodeIntraRefreshModeNoneKhr                   VkVideoEncodeIntraRefreshModeFlagBitsKHR = 0
	VkVideoEncodeIntraRefreshModePerPicturePartitionBitKhr VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 0)
	VkVideoEncodeIntraRefreshModeBlockBasedBitKhr          VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 1)
	VkVideoEncodeIntraRefreshModeBlockRowBasedBitKhr       VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 2)
	VkVideoEncodeIntraRefreshModeBlockColumnBasedBitKhr    VkVideoEncodeIntraRefreshModeFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeH264CapabilityFlagBitsKHR int64

var (
	VkVideoEncodeH264CapabilityHrdComplianceBitKhr                  VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeH264CapabilityPredictionWeightTableGeneratedBitKhr VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 1)
	VkVideoEncodeH264CapabilityRowUnalignedSliceBitKhr              VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 2)
	VkVideoEncodeH264CapabilityDifferentSliceTypeBitKhr             VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 3)
	VkVideoEncodeH264CapabilityBFrameInL0ListBitKhr                 VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 4)
	VkVideoEncodeH264CapabilityBFrameInL1ListBitKhr                 VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 5)
	VkVideoEncodeH264CapabilityPerPictureTypeMinMaxQpBitKhr         VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 6)
	VkVideoEncodeH264CapabilityPerSliceConstantQpBitKhr             VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 7)
	VkVideoEncodeH264CapabilityGeneratePrefixNaluBitKhr             VkVideoEncodeH264CapabilityFlagBitsKHR = (1 << 8)
)

type VkVideoEncodeH264StdFlagBitsKHR int64

var (
	VkVideoEncodeH264StdSeparateColorPlaneFlagSetBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 0)
	VkVideoEncodeH264StdQpprimeYZeroTransformBypassFlagSetBitKhr VkVideoEncodeH264StdFlagBitsKHR = (1 << 1)
	VkVideoEncodeH264StdScalingMatrixPresentFlagSetBitKhr        VkVideoEncodeH264StdFlagBitsKHR = (1 << 2)
	VkVideoEncodeH264StdChromaQpIndexOffsetBitKhr                VkVideoEncodeH264StdFlagBitsKHR = (1 << 3)
	VkVideoEncodeH264StdSecondChromaQpIndexOffsetBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 4)
	VkVideoEncodeH264StdPicInitQpMinus26BitKhr                   VkVideoEncodeH264StdFlagBitsKHR = (1 << 5)
	VkVideoEncodeH264StdWeightedPredFlagSetBitKhr                VkVideoEncodeH264StdFlagBitsKHR = (1 << 6)
	VkVideoEncodeH264StdWeightedBipredIdcExplicitBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 7)
	VkVideoEncodeH264StdWeightedBipredIdcImplicitBitKhr          VkVideoEncodeH264StdFlagBitsKHR = (1 << 8)
	VkVideoEncodeH264StdTransform8x8ModeFlagSetBitKhr            VkVideoEncodeH264StdFlagBitsKHR = (1 << 9)
	VkVideoEncodeH264StdDirectSpatialMvPredFlagUnsetBitKhr       VkVideoEncodeH264StdFlagBitsKHR = (1 << 10)
	VkVideoEncodeH264StdEntropyCodingModeFlagUnsetBitKhr         VkVideoEncodeH264StdFlagBitsKHR = (1 << 11)
	VkVideoEncodeH264StdEntropyCodingModeFlagSetBitKhr           VkVideoEncodeH264StdFlagBitsKHR = (1 << 12)
	VkVideoEncodeH264StdDirect8x8InferenceFlagUnsetBitKhr        VkVideoEncodeH264StdFlagBitsKHR = (1 << 13)
	VkVideoEncodeH264StdConstrainedIntraPredFlagSetBitKhr        VkVideoEncodeH264StdFlagBitsKHR = (1 << 14)
	VkVideoEncodeH264StdDeblockingFilterDisabledBitKhr           VkVideoEncodeH264StdFlagBitsKHR = (1 << 15)
	VkVideoEncodeH264StdDeblockingFilterEnabledBitKhr            VkVideoEncodeH264StdFlagBitsKHR = (1 << 16)
	VkVideoEncodeH264StdDeblockingFilterPartialBitKhr            VkVideoEncodeH264StdFlagBitsKHR = (1 << 17)
	VkVideoEncodeH264StdSliceQpDeltaBitKhr                       VkVideoEncodeH264StdFlagBitsKHR = (1 << 19)
	VkVideoEncodeH264StdDifferentSliceQpDeltaBitKhr              VkVideoEncodeH264StdFlagBitsKHR = (1 << 20)
)

type VkVideoEncodeH264RateControlFlagBitsKHR int64

var (
	VkVideoEncodeH264RateControlAttemptHrdComplianceBitKhr       VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 0)
	VkVideoEncodeH264RateControlRegularGopBitKhr                 VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 1)
	VkVideoEncodeH264RateControlReferencePatternFlatBitKhr       VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 2)
	VkVideoEncodeH264RateControlReferencePatternDyadicBitKhr     VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 3)
	VkVideoEncodeH264RateControlTemporalLayerPatternDyadicBitKhr VkVideoEncodeH264RateControlFlagBitsKHR = (1 << 4)
)

type VkHostImageCopyFlagBits int64

var (
	VkHostImageCopyMemcpyBit VkHostImageCopyFlagBits = (1 << 0)
	VkHostImageCopyMemcpy    VkHostImageCopyFlagBits = 0
)

type VkPartitionedAccelerationStructureOpTypeNV int64

var (
	VkPartitionedAccelerationStructureOpTypeWriteInstanceNv             VkPartitionedAccelerationStructureOpTypeNV = 0
	VkPartitionedAccelerationStructureOpTypeUpdateInstanceNv            VkPartitionedAccelerationStructureOpTypeNV = 1
	VkPartitionedAccelerationStructureOpTypeWritePartitionTranslationNv VkPartitionedAccelerationStructureOpTypeNV = 2
)

type VkPartitionedAccelerationStructureInstanceFlagBitsNV int64

var (
	VkPartitionedAccelerationStructureInstanceFlagTriangleFacingCullDisableBitNv VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 0)
	VkPartitionedAccelerationStructureInstanceFlagTriangleFlipFacingBitNv        VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 1)
	VkPartitionedAccelerationStructureInstanceFlagForceOpaqueBitNv               VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 2)
	VkPartitionedAccelerationStructureInstanceFlagForceNoOpaqueBitNv             VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 3)
	VkPartitionedAccelerationStructureInstanceFlagEnableExplicitBoundingBoxNv    VkPartitionedAccelerationStructureInstanceFlagBitsNV = (1 << 4)
)

type VkImageFormatConstraintsFlagBitsFUCHSIA int64

var ()

type VkImageConstraintsInfoFlagBitsFUCHSIA int64

var (
	VkImageConstraintsInfoCpuReadRarelyFuchsia     VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 0)
	VkImageConstraintsInfoCpuReadOftenFuchsia      VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 1)
	VkImageConstraintsInfoCpuWriteRarelyFuchsia    VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 2)
	VkImageConstraintsInfoCpuWriteOftenFuchsia     VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 3)
	VkImageConstraintsInfoProtectedOptionalFuchsia VkImageConstraintsInfoFlagBitsFUCHSIA = (1 << 4)
)

type VkFormatFeatureFlagBits2 int64

var (
	VkFormatFeature2SampledImageBit                                                     VkFormatFeatureFlagBits2 = (1 << 0)
	VkFormatFeature2StorageImageBit                                                     VkFormatFeatureFlagBits2 = (1 << 1)
	VkFormatFeature2StorageImageAtomicBit                                               VkFormatFeatureFlagBits2 = (1 << 2)
	VkFormatFeature2UniformTexelBufferBit                                               VkFormatFeatureFlagBits2 = (1 << 3)
	VkFormatFeature2StorageTexelBufferBit                                               VkFormatFeatureFlagBits2 = (1 << 4)
	VkFormatFeature2StorageTexelBufferAtomicBit                                         VkFormatFeatureFlagBits2 = (1 << 5)
	VkFormatFeature2VertexBufferBit                                                     VkFormatFeatureFlagBits2 = (1 << 6)
	VkFormatFeature2ColorAttachmentBit                                                  VkFormatFeatureFlagBits2 = (1 << 7)
	VkFormatFeature2ColorAttachmentBlendBit                                             VkFormatFeatureFlagBits2 = (1 << 8)
	VkFormatFeature2DepthStencilAttachmentBit                                           VkFormatFeatureFlagBits2 = (1 << 9)
	VkFormatFeature2BlitSrcBit                                                          VkFormatFeatureFlagBits2 = (1 << 10)
	VkFormatFeature2BlitDstBit                                                          VkFormatFeatureFlagBits2 = (1 << 11)
	VkFormatFeature2SampledImageFilterLinearBit                                         VkFormatFeatureFlagBits2 = (1 << 12)
	VkFormatFeature2TransferSrcBit                                                      VkFormatFeatureFlagBits2 = (1 << 14)
	VkFormatFeature2TransferDstBit                                                      VkFormatFeatureFlagBits2 = (1 << 15)
	VkFormatFeature2SampledImageFilterMinmaxBit                                         VkFormatFeatureFlagBits2 = (1 << 16)
	VkFormatFeature2MidpointChromaSamplesBit                                            VkFormatFeatureFlagBits2 = (1 << 17)
	VkFormatFeature2SampledImageYcbcrConversionLinearFilterBit                          VkFormatFeatureFlagBits2 = (1 << 18)
	VkFormatFeature2SampledImageYcbcrConversionSeparateReconstructionFilterBit          VkFormatFeatureFlagBits2 = (1 << 19)
	VkFormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitBit          VkFormatFeatureFlagBits2 = (1 << 20)
	VkFormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitForceableBit VkFormatFeatureFlagBits2 = (1 << 21)
	VkFormatFeature2DisjointBit                                                         VkFormatFeatureFlagBits2 = (1 << 22)
	VkFormatFeature2CositedChromaSamplesBit                                             VkFormatFeatureFlagBits2 = (1 << 23)
	VkFormatFeature2StorageReadWithoutFormatBit                                         VkFormatFeatureFlagBits2 = (1 << 31)
	VkFormatFeature2StorageWriteWithoutFormatBit                                        VkFormatFeatureFlagBits2 = (1 << 32)
	VkFormatFeature2SampledImageDepthComparisonBit                                      VkFormatFeatureFlagBits2 = (1 << 33)
)

type VkRenderingFlagBits int64

var (
	VkRenderingContentsSecondaryCommandBuffersBit VkRenderingFlagBits = (1 << 0)
	VkRenderingSuspendingBit                      VkRenderingFlagBits = (1 << 1)
	VkRenderingResumingBit                        VkRenderingFlagBits = (1 << 2)
)

type VkVideoEncodeH265CapabilityFlagBitsKHR int64

var (
	VkVideoEncodeH265CapabilityHrdComplianceBitKhr                  VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265CapabilityPredictionWeightTableGeneratedBitKhr VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265CapabilityRowUnalignedSliceSegmentBitKhr       VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265CapabilityDifferentSliceSegmentTypeBitKhr      VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 3)
	VkVideoEncodeH265CapabilityBFrameInL0ListBitKhr                 VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 4)
	VkVideoEncodeH265CapabilityBFrameInL1ListBitKhr                 VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 5)
	VkVideoEncodeH265CapabilityPerPictureTypeMinMaxQpBitKhr         VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 6)
	VkVideoEncodeH265CapabilityPerSliceSegmentConstantQpBitKhr      VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 7)
	VkVideoEncodeH265CapabilityMultipleTilesPerSliceSegmentBitKhr   VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 8)
	VkVideoEncodeH265CapabilityMultipleSliceSegmentsPerTileBitKhr   VkVideoEncodeH265CapabilityFlagBitsKHR = (1 << 9)
)

type VkVideoEncodeH265StdFlagBitsKHR int64

var (
	VkVideoEncodeH265StdSeparateColorPlaneFlagSetBitKhr              VkVideoEncodeH265StdFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265StdSampleAdaptiveOffsetEnabledFlagSetBitKhr     VkVideoEncodeH265StdFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265StdScalingListDataPresentFlagSetBitKhr          VkVideoEncodeH265StdFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265StdPcmEnabledFlagSetBitKhr                      VkVideoEncodeH265StdFlagBitsKHR = (1 << 3)
	VkVideoEncodeH265StdSpsTemporalMvpEnabledFlagSetBitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 4)
	VkVideoEncodeH265StdInitQpMinus26BitKhr                          VkVideoEncodeH265StdFlagBitsKHR = (1 << 5)
	VkVideoEncodeH265StdWeightedPredFlagSetBitKhr                    VkVideoEncodeH265StdFlagBitsKHR = (1 << 6)
	VkVideoEncodeH265StdWeightedBipredFlagSetBitKhr                  VkVideoEncodeH265StdFlagBitsKHR = (1 << 7)
	VkVideoEncodeH265StdLog2ParallelMergeLevelMinus2BitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 8)
	VkVideoEncodeH265StdSignDataHidingEnabledFlagSetBitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 9)
	VkVideoEncodeH265StdTransformSkipEnabledFlagSetBitKhr            VkVideoEncodeH265StdFlagBitsKHR = (1 << 10)
	VkVideoEncodeH265StdTransformSkipEnabledFlagUnsetBitKhr          VkVideoEncodeH265StdFlagBitsKHR = (1 << 11)
	VkVideoEncodeH265StdPpsSliceChromaQpOffsetsPresentFlagSetBitKhr  VkVideoEncodeH265StdFlagBitsKHR = (1 << 12)
	VkVideoEncodeH265StdTransquantBypassEnabledFlagSetBitKhr         VkVideoEncodeH265StdFlagBitsKHR = (1 << 13)
	VkVideoEncodeH265StdConstrainedIntraPredFlagSetBitKhr            VkVideoEncodeH265StdFlagBitsKHR = (1 << 14)
	VkVideoEncodeH265StdEntropyCodingSyncEnabledFlagSetBitKhr        VkVideoEncodeH265StdFlagBitsKHR = (1 << 15)
	VkVideoEncodeH265StdDeblockingFilterOverrideEnabledFlagSetBitKhr VkVideoEncodeH265StdFlagBitsKHR = (1 << 16)
	VkVideoEncodeH265StdDependentSliceSegmentsEnabledFlagSetBitKhr   VkVideoEncodeH265StdFlagBitsKHR = (1 << 17)
	VkVideoEncodeH265StdDependentSliceSegmentFlagSetBitKhr           VkVideoEncodeH265StdFlagBitsKHR = (1 << 18)
	VkVideoEncodeH265StdSliceQpDeltaBitKhr                           VkVideoEncodeH265StdFlagBitsKHR = (1 << 19)
	VkVideoEncodeH265StdDifferentSliceQpDeltaBitKhr                  VkVideoEncodeH265StdFlagBitsKHR = (1 << 20)
)

type VkVideoEncodeH265RateControlFlagBitsKHR int64

var (
	VkVideoEncodeH265RateControlAttemptHrdComplianceBitKhr          VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265RateControlRegularGopBitKhr                    VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265RateControlReferencePatternFlatBitKhr          VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265RateControlReferencePatternDyadicBitKhr        VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 3)
	VkVideoEncodeH265RateControlTemporalSubLayerPatternDyadicBitKhr VkVideoEncodeH265RateControlFlagBitsKHR = (1 << 4)
)

type VkVideoEncodeH265CtbSizeFlagBitsKHR int64

var (
	VkVideoEncodeH265CtbSize16BitKhr VkVideoEncodeH265CtbSizeFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265CtbSize32BitKhr VkVideoEncodeH265CtbSizeFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265CtbSize64BitKhr VkVideoEncodeH265CtbSizeFlagBitsKHR = (1 << 2)
)

type VkVideoEncodeH265TransformBlockSizeFlagBitsKHR int64

var (
	VkVideoEncodeH265TransformBlockSize4BitKhr  VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 0)
	VkVideoEncodeH265TransformBlockSize8BitKhr  VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 1)
	VkVideoEncodeH265TransformBlockSize16BitKhr VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 2)
	VkVideoEncodeH265TransformBlockSize32BitKhr VkVideoEncodeH265TransformBlockSizeFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeAV1CapabilityFlagBitsKHR int64

var (
	VkVideoEncodeAv1CapabilityPerRateControlGroupMinMaxQIndexBitKhr VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1CapabilityGenerateObuExtensionHeaderBitKhr      VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 1)
	VkVideoEncodeAv1CapabilityPrimaryReferenceCdfOnlyBitKhr         VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 2)
	VkVideoEncodeAv1CapabilityFrameSizeOverrideBitKhr               VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 3)
	VkVideoEncodeAv1CapabilityMotionVectorScalingBitKhr             VkVideoEncodeAV1CapabilityFlagBitsKHR = (1 << 4)
)

type VkVideoEncodeAV1StdFlagBitsKHR int64

var (
	VkVideoEncodeAv1StdUniformTileSpacingFlagSetBitKhr VkVideoEncodeAV1StdFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1StdSkipModePresentUnsetBitKhr      VkVideoEncodeAV1StdFlagBitsKHR = (1 << 1)
	VkVideoEncodeAv1StdPrimaryRefFrameBitKhr           VkVideoEncodeAV1StdFlagBitsKHR = (1 << 2)
	VkVideoEncodeAv1StdDeltaQBitKhr                    VkVideoEncodeAV1StdFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeAV1RateControlFlagBitsKHR int64

var (
	VkVideoEncodeAv1RateControlRegularGopBitKhr                 VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1RateControlTemporalLayerPatternDyadicBitKhr VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 1)
	VkVideoEncodeAv1RateControlReferencePatternFlatBitKhr       VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 2)
	VkVideoEncodeAv1RateControlReferencePatternDyadicBitKhr     VkVideoEncodeAV1RateControlFlagBitsKHR = (1 << 3)
)

type VkVideoEncodeAV1SuperblockSizeFlagBitsKHR int64

var (
	VkVideoEncodeAv1SuperblockSize64BitKhr  VkVideoEncodeAV1SuperblockSizeFlagBitsKHR = (1 << 0)
	VkVideoEncodeAv1SuperblockSize128BitKhr VkVideoEncodeAV1SuperblockSizeFlagBitsKHR = (1 << 1)
)

type VkVideoEncodeAV1PredictionModeKHR int64

var (
	VkVideoEncodeAv1PredictionModeIntraOnlyKhr              VkVideoEncodeAV1PredictionModeKHR = 0
	VkVideoEncodeAv1PredictionModeSingleReferenceKhr        VkVideoEncodeAV1PredictionModeKHR = 1
	VkVideoEncodeAv1PredictionModeUnidirectionalCompoundKhr VkVideoEncodeAV1PredictionModeKHR = 2
	VkVideoEncodeAv1PredictionModeBidirectionalCompoundKhr  VkVideoEncodeAV1PredictionModeKHR = 3
)

type VkVideoEncodeAV1RateControlGroupKHR int64

var (
	VkVideoEncodeAv1RateControlGroupIntraKhr        VkVideoEncodeAV1RateControlGroupKHR = 0
	VkVideoEncodeAv1RateControlGroupPredictiveKhr   VkVideoEncodeAV1RateControlGroupKHR = 1
	VkVideoEncodeAv1RateControlGroupBipredictiveKhr VkVideoEncodeAV1RateControlGroupKHR = 2
)

type VkExportMetalObjectTypeFlagBitsEXT int64

var (
	VkExportMetalObjectTypeMetalDeviceBitExt       VkExportMetalObjectTypeFlagBitsEXT = (1 << 0)
	VkExportMetalObjectTypeMetalCommandQueueBitExt VkExportMetalObjectTypeFlagBitsEXT = (1 << 1)
	VkExportMetalObjectTypeMetalBufferBitExt       VkExportMetalObjectTypeFlagBitsEXT = (1 << 2)
	VkExportMetalObjectTypeMetalTextureBitExt      VkExportMetalObjectTypeFlagBitsEXT = (1 << 3)
	VkExportMetalObjectTypeMetalIosurfaceBitExt    VkExportMetalObjectTypeFlagBitsEXT = (1 << 4)
	VkExportMetalObjectTypeMetalSharedEventBitExt  VkExportMetalObjectTypeFlagBitsEXT = (1 << 5)
)

type VkInstanceCreateFlagBits int64

var ()

type VkImageCompressionFlagBitsEXT int64

var (
	VkImageCompressionDefaultExt           VkImageCompressionFlagBitsEXT = 0
	VkImageCompressionFixedRateDefaultExt  VkImageCompressionFlagBitsEXT = (1 << 0)
	VkImageCompressionFixedRateExplicitExt VkImageCompressionFlagBitsEXT = (1 << 1)
	VkImageCompressionDisabledExt          VkImageCompressionFlagBitsEXT = (1 << 2)
)

type VkImageCompressionFixedRateFlagBitsEXT int64

var (
	VkImageCompressionFixedRateNoneExt     VkImageCompressionFixedRateFlagBitsEXT = 0
	VkImageCompressionFixedRate1bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 0)
	VkImageCompressionFixedRate2bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 1)
	VkImageCompressionFixedRate3bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 2)
	VkImageCompressionFixedRate4bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 3)
	VkImageCompressionFixedRate5bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 4)
	VkImageCompressionFixedRate6bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 5)
	VkImageCompressionFixedRate7bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 6)
	VkImageCompressionFixedRate8bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 7)
	VkImageCompressionFixedRate9bpcBitExt  VkImageCompressionFixedRateFlagBitsEXT = (1 << 8)
	VkImageCompressionFixedRate10bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 9)
	VkImageCompressionFixedRate11bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 10)
	VkImageCompressionFixedRate12bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 11)
	VkImageCompressionFixedRate13bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 12)
	VkImageCompressionFixedRate14bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 13)
	VkImageCompressionFixedRate15bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 14)
	VkImageCompressionFixedRate16bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 15)
	VkImageCompressionFixedRate17bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 16)
	VkImageCompressionFixedRate18bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 17)
	VkImageCompressionFixedRate19bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 18)
	VkImageCompressionFixedRate20bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 19)
	VkImageCompressionFixedRate21bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 20)
	VkImageCompressionFixedRate22bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 21)
	VkImageCompressionFixedRate23bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 22)
	VkImageCompressionFixedRate24bpcBitExt VkImageCompressionFixedRateFlagBitsEXT = (1 << 23)
)

type VkPipelineRobustnessBufferBehavior int64

var (
	VkPipelineRobustnessBufferBehaviorDeviceDefault       VkPipelineRobustnessBufferBehavior = 0
	VkPipelineRobustnessBufferBehaviorDisabled            VkPipelineRobustnessBufferBehavior = 1
	VkPipelineRobustnessBufferBehaviorRobustBufferAccess  VkPipelineRobustnessBufferBehavior = 2
	VkPipelineRobustnessBufferBehaviorRobustBufferAccess2 VkPipelineRobustnessBufferBehavior = 3
)

type VkPipelineRobustnessImageBehavior int64

var (
	VkPipelineRobustnessImageBehaviorDeviceDefault      VkPipelineRobustnessImageBehavior = 0
	VkPipelineRobustnessImageBehaviorDisabled           VkPipelineRobustnessImageBehavior = 1
	VkPipelineRobustnessImageBehaviorRobustImageAccess  VkPipelineRobustnessImageBehavior = 2
	VkPipelineRobustnessImageBehaviorRobustImageAccess2 VkPipelineRobustnessImageBehavior = 3
)

type VkOpticalFlowGridSizeFlagBitsNV int64

var (
	VkOpticalFlowGridSizeUnknownNv VkOpticalFlowGridSizeFlagBitsNV = 0
	VkOpticalFlowGridSize1x1BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 0)
	VkOpticalFlowGridSize2x2BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 1)
	VkOpticalFlowGridSize4x4BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 2)
	VkOpticalFlowGridSize8x8BitNv  VkOpticalFlowGridSizeFlagBitsNV = (1 << 3)
)

type VkOpticalFlowUsageFlagBitsNV int64

var (
	VkOpticalFlowUsageUnknownNv       VkOpticalFlowUsageFlagBitsNV = 0
	VkOpticalFlowUsageInputBitNv      VkOpticalFlowUsageFlagBitsNV = (1 << 0)
	VkOpticalFlowUsageOutputBitNv     VkOpticalFlowUsageFlagBitsNV = (1 << 1)
	VkOpticalFlowUsageHintBitNv       VkOpticalFlowUsageFlagBitsNV = (1 << 2)
	VkOpticalFlowUsageCostBitNv       VkOpticalFlowUsageFlagBitsNV = (1 << 3)
	VkOpticalFlowUsageGlobalFlowBitNv VkOpticalFlowUsageFlagBitsNV = (1 << 4)
)

type VkOpticalFlowPerformanceLevelNV int64

var (
	VkOpticalFlowPerformanceLevelUnknownNv VkOpticalFlowPerformanceLevelNV = 0
	VkOpticalFlowPerformanceLevelSlowNv    VkOpticalFlowPerformanceLevelNV = 1
	VkOpticalFlowPerformanceLevelMediumNv  VkOpticalFlowPerformanceLevelNV = 2
	VkOpticalFlowPerformanceLevelFastNv    VkOpticalFlowPerformanceLevelNV = 3
)

type VkOpticalFlowSessionBindingPointNV int64

var (
	VkOpticalFlowSessionBindingPointUnknownNv            VkOpticalFlowSessionBindingPointNV = 0
	VkOpticalFlowSessionBindingPointInputNv              VkOpticalFlowSessionBindingPointNV = 1
	VkOpticalFlowSessionBindingPointReferenceNv          VkOpticalFlowSessionBindingPointNV = 2
	VkOpticalFlowSessionBindingPointHintNv               VkOpticalFlowSessionBindingPointNV = 3
	VkOpticalFlowSessionBindingPointFlowVectorNv         VkOpticalFlowSessionBindingPointNV = 4
	VkOpticalFlowSessionBindingPointBackwardFlowVectorNv VkOpticalFlowSessionBindingPointNV = 5
	VkOpticalFlowSessionBindingPointCostNv               VkOpticalFlowSessionBindingPointNV = 6
	VkOpticalFlowSessionBindingPointBackwardCostNv       VkOpticalFlowSessionBindingPointNV = 7
	VkOpticalFlowSessionBindingPointGlobalFlowNv         VkOpticalFlowSessionBindingPointNV = 8
)

type VkOpticalFlowSessionCreateFlagBitsNV int64

var (
	VkOpticalFlowSessionCreateEnableHintBitNv       VkOpticalFlowSessionCreateFlagBitsNV = (1 << 0)
	VkOpticalFlowSessionCreateEnableCostBitNv       VkOpticalFlowSessionCreateFlagBitsNV = (1 << 1)
	VkOpticalFlowSessionCreateEnableGlobalFlowBitNv VkOpticalFlowSessionCreateFlagBitsNV = (1 << 2)
	VkOpticalFlowSessionCreateAllowRegionsBitNv     VkOpticalFlowSessionCreateFlagBitsNV = (1 << 3)
	VkOpticalFlowSessionCreateBothDirectionsBitNv   VkOpticalFlowSessionCreateFlagBitsNV = (1 << 4)
)

type VkOpticalFlowExecuteFlagBitsNV int64

var (
	VkOpticalFlowExecuteDisableTemporalHintsBitNv VkOpticalFlowExecuteFlagBitsNV = (1 << 0)
)

type VkMicromapTypeEXT int64

var (
	VkMicromapTypeOpacityMicromapExt VkMicromapTypeEXT = 0
)

type VkBuildMicromapFlagBitsEXT int64

var (
	VkBuildMicromapPreferFastTraceBitExt VkBuildMicromapFlagBitsEXT = (1 << 0)
	VkBuildMicromapPreferFastBuildBitExt VkBuildMicromapFlagBitsEXT = (1 << 1)
	VkBuildMicromapAllowCompactionBitExt VkBuildMicromapFlagBitsEXT = (1 << 2)
)

type VkMicromapCreateFlagBitsEXT int64

var (
	VkMicromapCreateDeviceAddressCaptureReplayBitExt VkMicromapCreateFlagBitsEXT = (1 << 0)
)

type VkCopyMicromapModeEXT int64

var (
	VkCopyMicromapModeCloneExt       VkCopyMicromapModeEXT = 0
	VkCopyMicromapModeSerializeExt   VkCopyMicromapModeEXT = 1
	VkCopyMicromapModeDeserializeExt VkCopyMicromapModeEXT = 2
	VkCopyMicromapModeCompactExt     VkCopyMicromapModeEXT = 3
)

type VkBuildMicromapModeEXT int64

var (
	VkBuildMicromapModeBuildExt VkBuildMicromapModeEXT = 0
)

type VkOpacityMicromapFormatEXT int64

var (
	VkOpacityMicromapFormat2StateExt VkOpacityMicromapFormatEXT = 1
	VkOpacityMicromapFormat4StateExt VkOpacityMicromapFormatEXT = 2
)

type VkOpacityMicromapSpecialIndexEXT int64

var (
	VkOpacityMicromapSpecialIndexFullyTransparentExt        VkOpacityMicromapSpecialIndexEXT = -1
	VkOpacityMicromapSpecialIndexFullyOpaqueExt             VkOpacityMicromapSpecialIndexEXT = -2
	VkOpacityMicromapSpecialIndexFullyUnknownTransparentExt VkOpacityMicromapSpecialIndexEXT = -3
	VkOpacityMicromapSpecialIndexFullyUnknownOpaqueExt      VkOpacityMicromapSpecialIndexEXT = -4
)

type VkDepthBiasRepresentationEXT int64

var (
	VkDepthBiasRepresentationLeastRepresentableValueFormatExt     VkDepthBiasRepresentationEXT = 0
	VkDepthBiasRepresentationLeastRepresentableValueForceUnormExt VkDepthBiasRepresentationEXT = 1
	VkDepthBiasRepresentationFloatExt                             VkDepthBiasRepresentationEXT = 2
)

type VkDeviceFaultAddressTypeEXT int64

var (
	VkDeviceFaultAddressTypeNoneExt                      VkDeviceFaultAddressTypeEXT = 0
	VkDeviceFaultAddressTypeReadInvalidExt               VkDeviceFaultAddressTypeEXT = 1
	VkDeviceFaultAddressTypeWriteInvalidExt              VkDeviceFaultAddressTypeEXT = 2
	VkDeviceFaultAddressTypeExecuteInvalidExt            VkDeviceFaultAddressTypeEXT = 3
	VkDeviceFaultAddressTypeInstructionPointerUnknownExt VkDeviceFaultAddressTypeEXT = 4
	VkDeviceFaultAddressTypeInstructionPointerInvalidExt VkDeviceFaultAddressTypeEXT = 5
	VkDeviceFaultAddressTypeInstructionPointerFaultExt   VkDeviceFaultAddressTypeEXT = 6
)

type VkDeviceFaultVendorBinaryHeaderVersionEXT int64

var (
	VkDeviceFaultVendorBinaryHeaderVersionOneExt VkDeviceFaultVendorBinaryHeaderVersionEXT = 1
)

type VkIndirectCommandsLayoutUsageFlagBitsEXT int64

var (
	VkIndirectCommandsLayoutUsageExplicitPreprocessBitExt VkIndirectCommandsLayoutUsageFlagBitsEXT = (1 << 0)
	VkIndirectCommandsLayoutUsageUnorderedSequencesBitExt VkIndirectCommandsLayoutUsageFlagBitsEXT = (1 << 1)
)

type VkIndirectExecutionSetInfoTypeEXT int64

var (
	VkIndirectExecutionSetInfoTypePipelinesExt     VkIndirectExecutionSetInfoTypeEXT = 0
	VkIndirectExecutionSetInfoTypeShaderObjectsExt VkIndirectExecutionSetInfoTypeEXT = 1
)

type VkIndirectCommandsInputModeFlagBitsEXT int64

var (
	VkIndirectCommandsInputModeVulkanIndexBufferExt VkIndirectCommandsInputModeFlagBitsEXT = (1 << 0)
	VkIndirectCommandsInputModeDxgiIndexBufferExt   VkIndirectCommandsInputModeFlagBitsEXT = (1 << 1)
)

type VkIndirectCommandsTokenTypeEXT int64

var (
	VkIndirectCommandsTokenTypeExecutionSetExt     VkIndirectCommandsTokenTypeEXT = 0
	VkIndirectCommandsTokenTypePushConstantExt     VkIndirectCommandsTokenTypeEXT = 1
	VkIndirectCommandsTokenTypeSequenceIndexExt    VkIndirectCommandsTokenTypeEXT = 2
	VkIndirectCommandsTokenTypeIndexBufferExt      VkIndirectCommandsTokenTypeEXT = 3
	VkIndirectCommandsTokenTypeVertexBufferExt     VkIndirectCommandsTokenTypeEXT = 4
	VkIndirectCommandsTokenTypeDrawIndexedExt      VkIndirectCommandsTokenTypeEXT = 5
	VkIndirectCommandsTokenTypeDrawExt             VkIndirectCommandsTokenTypeEXT = 6
	VkIndirectCommandsTokenTypeDrawIndexedCountExt VkIndirectCommandsTokenTypeEXT = 7
	VkIndirectCommandsTokenTypeDrawCountExt        VkIndirectCommandsTokenTypeEXT = 8
	VkIndirectCommandsTokenTypeDispatchExt         VkIndirectCommandsTokenTypeEXT = 9
)

type VkDisplacementMicromapFormatNV int64

var (
	VkDisplacementMicromapFormat64Triangles64BytesNv    VkDisplacementMicromapFormatNV = 1
	VkDisplacementMicromapFormat256Triangles128BytesNv  VkDisplacementMicromapFormatNV = 2
	VkDisplacementMicromapFormat1024Triangles128BytesNv VkDisplacementMicromapFormatNV = 3
)

type VkShaderCreateFlagBitsEXT int64

var (
	VkShaderCreateLinkStageBitExt VkShaderCreateFlagBitsEXT = (1 << 0)
)

type VkShaderCodeTypeEXT int64

var (
	VkShaderCodeTypeBinaryExt VkShaderCodeTypeEXT = 0
	VkShaderCodeTypeSpirvExt  VkShaderCodeTypeEXT = 1
)

type VkScopeKHR int64

var (
	VkScopeDeviceKhr      VkScopeKHR = 1
	VkScopeWorkgroupKhr   VkScopeKHR = 2
	VkScopeSubgroupKhr    VkScopeKHR = 3
	VkScopeQueueFamilyKhr VkScopeKHR = 5
)

type VkComponentTypeKHR int64

var (
	VkComponentTypeFloat16Khr VkComponentTypeKHR = 0
	VkComponentTypeFloat32Khr VkComponentTypeKHR = 1
	VkComponentTypeFloat64Khr VkComponentTypeKHR = 2
	VkComponentTypeSint8Khr   VkComponentTypeKHR = 3
	VkComponentTypeSint16Khr  VkComponentTypeKHR = 4
	VkComponentTypeSint32Khr  VkComponentTypeKHR = 5
	VkComponentTypeSint64Khr  VkComponentTypeKHR = 6
	VkComponentTypeUint8Khr   VkComponentTypeKHR = 7
	VkComponentTypeUint16Khr  VkComponentTypeKHR = 8
	VkComponentTypeUint32Khr  VkComponentTypeKHR = 9
	VkComponentTypeUint64Khr  VkComponentTypeKHR = 10
)

type VkCubicFilterWeightsQCOM int64

var (
	VkCubicFilterWeightsCatmullRomQcom          VkCubicFilterWeightsQCOM = 0
	VkCubicFilterWeightsZeroTangentCardinalQcom VkCubicFilterWeightsQCOM = 1
	VkCubicFilterWeightsBSplineQcom             VkCubicFilterWeightsQCOM = 2
	VkCubicFilterWeightsMitchellNetravaliQcom   VkCubicFilterWeightsQCOM = 3
)

type VkBlockMatchWindowCompareModeQCOM int64

var (
	VkBlockMatchWindowCompareModeMinQcom VkBlockMatchWindowCompareModeQCOM = 0
	VkBlockMatchWindowCompareModeMaxQcom VkBlockMatchWindowCompareModeQCOM = 1
)

type VkPhysicalDeviceLayeredApiKHR int64

var (
	VkPhysicalDeviceLayeredApiVulkanKhr   VkPhysicalDeviceLayeredApiKHR = 0
	VkPhysicalDeviceLayeredApiD3d12Khr    VkPhysicalDeviceLayeredApiKHR = 1
	VkPhysicalDeviceLayeredApiMetalKhr    VkPhysicalDeviceLayeredApiKHR = 2
	VkPhysicalDeviceLayeredApiOpenglKhr   VkPhysicalDeviceLayeredApiKHR = 3
	VkPhysicalDeviceLayeredApiOpenglesKhr VkPhysicalDeviceLayeredApiKHR = 4
)

type VkLayeredDriverUnderlyingApiMSFT int64

var (
	VkLayeredDriverUnderlyingApiNoneMsft  VkLayeredDriverUnderlyingApiMSFT = 0
	VkLayeredDriverUnderlyingApiD3d12Msft VkLayeredDriverUnderlyingApiMSFT = 1
)

type VkLatencyMarkerNV int64

var (
	VkLatencyMarkerSimulationStartNv            VkLatencyMarkerNV = 0
	VkLatencyMarkerSimulationEndNv              VkLatencyMarkerNV = 1
	VkLatencyMarkerRendersubmitStartNv          VkLatencyMarkerNV = 2
	VkLatencyMarkerRendersubmitEndNv            VkLatencyMarkerNV = 3
	VkLatencyMarkerPresentStartNv               VkLatencyMarkerNV = 4
	VkLatencyMarkerPresentEndNv                 VkLatencyMarkerNV = 5
	VkLatencyMarkerInputSampleNv                VkLatencyMarkerNV = 6
	VkLatencyMarkerTriggerFlashNv               VkLatencyMarkerNV = 7
	VkLatencyMarkerOutOfBandRendersubmitStartNv VkLatencyMarkerNV = 8
	VkLatencyMarkerOutOfBandRendersubmitEndNv   VkLatencyMarkerNV = 9
	VkLatencyMarkerOutOfBandPresentStartNv      VkLatencyMarkerNV = 10
	VkLatencyMarkerOutOfBandPresentEndNv        VkLatencyMarkerNV = 11
)

type VkOutOfBandQueueTypeNV int64

var (
	VkOutOfBandQueueTypeRenderNv  VkOutOfBandQueueTypeNV = 0
	VkOutOfBandQueueTypePresentNv VkOutOfBandQueueTypeNV = 1
)

type VkMemoryUnmapFlagBits int64

var ()

type VkCompressedTriangleFormatAMDX int64

var (
	VkCompressedTriangleFormatDgf1Amdx VkCompressedTriangleFormatAMDX = 0
)

type VkWaylandSurfaceCreateFlagBitsKHR int64

var ()

type VkDepthClampModeEXT int64

var (
	VkDepthClampModeViewportRangeExt    VkDepthClampModeEXT = 0
	VkDepthClampModeUserDefinedRangeExt VkDepthClampModeEXT = 1
)

type VkAccessFlagBits3KHR int64

var (
	VkAccess3NoneKhr VkAccessFlagBits3KHR = 0
)

type VkTileShadingRenderPassFlagBitsQCOM int64

var (
	VkTileShadingRenderPassEnableBitQcom           VkTileShadingRenderPassFlagBitsQCOM = (1 << 0)
	VkTileShadingRenderPassPerTileExecutionBitQcom VkTileShadingRenderPassFlagBitsQCOM = (1 << 1)
)

type VkCooperativeVectorMatrixLayoutNV int64

var (
	VkCooperativeVectorMatrixLayoutRowMajorNv           VkCooperativeVectorMatrixLayoutNV = 0
	VkCooperativeVectorMatrixLayoutColumnMajorNv        VkCooperativeVectorMatrixLayoutNV = 1
	VkCooperativeVectorMatrixLayoutInferencingOptimalNv VkCooperativeVectorMatrixLayoutNV = 2
	VkCooperativeVectorMatrixLayoutTrainingOptimalNv    VkCooperativeVectorMatrixLayoutNV = 3
)

type VkAddressCopyFlagBitsKHR int64

var (
	VkAddressCopyDeviceLocalBitKhr VkAddressCopyFlagBitsKHR = (1 << 0)
	VkAddressCopySparseBitKhr      VkAddressCopyFlagBitsKHR = (1 << 1)
	VkAddressCopyProtectedBitKhr   VkAddressCopyFlagBitsKHR = (1 << 2)
)

type VkTensorCreateFlagBitsARM int64

var (
	VkTensorCreateMutableFormatBitArm VkTensorCreateFlagBitsARM = (1 << 0)
	VkTensorCreateProtectedBitArm     VkTensorCreateFlagBitsARM = (1 << 1)
)

type VkTensorUsageFlagBitsARM int64

var (
	VkTensorUsageShaderBitArm        VkTensorUsageFlagBitsARM = (1 << 1)
	VkTensorUsageTransferSrcBitArm   VkTensorUsageFlagBitsARM = (1 << 2)
	VkTensorUsageTransferDstBitArm   VkTensorUsageFlagBitsARM = (1 << 3)
	VkTensorUsageImageAliasingBitArm VkTensorUsageFlagBitsARM = (1 << 4)
)

type VkTensorTilingARM int64

var (
	VkTensorTilingOptimalArm VkTensorTilingARM = 0
	VkTensorTilingLinearArm  VkTensorTilingARM = 1
)

type VkTensorViewCreateFlagBitsARM int64

var ()

type VkDefaultVertexAttributeValueKHR int64

var (
	VkDefaultVertexAttributeValueZeroZeroZeroZeroKhr VkDefaultVertexAttributeValueKHR = 0
	VkDefaultVertexAttributeValueZeroZeroZeroOneKhr  VkDefaultVertexAttributeValueKHR = 1
)

type VkDataGraphPipelineSessionCreateFlagBitsARM int64

var (
	VkDataGraphPipelineSessionCreateProtectedBitArm VkDataGraphPipelineSessionCreateFlagBitsARM = (1 << 0)
)

type VkDataGraphPipelineSessionBindPointARM int64

var (
	VkDataGraphPipelineSessionBindPointTransientArm VkDataGraphPipelineSessionBindPointARM = 0
)

type VkDataGraphPipelineSessionBindPointTypeARM int64

var (
	VkDataGraphPipelineSessionBindPointTypeMemoryArm VkDataGraphPipelineSessionBindPointTypeARM = 0
)

type VkDataGraphPipelinePropertyARM int64

var (
	VkDataGraphPipelinePropertyCreationLogArm VkDataGraphPipelinePropertyARM = 0
	VkDataGraphPipelinePropertyIdentifierArm  VkDataGraphPipelinePropertyARM = 1
)

type VkDataGraphPipelineDispatchFlagBitsARM int64

var ()

type VkPhysicalDeviceDataGraphProcessingEngineTypeARM int64

var (
	VkPhysicalDeviceDataGraphProcessingEngineTypeDefaultArm VkPhysicalDeviceDataGraphProcessingEngineTypeARM = 0
)

type VkPhysicalDeviceDataGraphOperationTypeARM int64

var (
	VkPhysicalDeviceDataGraphOperationTypeSpirvExtendedInstructionSetArm VkPhysicalDeviceDataGraphOperationTypeARM = 0
)

type VkDataGraphModelCacheTypeQCOM int64

var (
	VkDataGraphModelCacheTypeGenericBinaryQcom VkDataGraphModelCacheTypeQCOM = 0
)

type VkVideoEncodeRgbModelConversionFlagBitsVALVE int64

var (
	VkVideoEncodeRgbModelConversionRgbIdentityBitValve   VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 0)
	VkVideoEncodeRgbModelConversionYcbcrIdentityBitValve VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 1)
	VkVideoEncodeRgbModelConversionYcbcr709BitValve      VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 2)
	VkVideoEncodeRgbModelConversionYcbcr601BitValve      VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 3)
	VkVideoEncodeRgbModelConversionYcbcr2020BitValve     VkVideoEncodeRgbModelConversionFlagBitsVALVE = (1 << 4)
)

type VkVideoEncodeRgbRangeCompressionFlagBitsVALVE int64

var (
	VkVideoEncodeRgbRangeCompressionFullRangeBitValve   VkVideoEncodeRgbRangeCompressionFlagBitsVALVE = (1 << 0)
	VkVideoEncodeRgbRangeCompressionNarrowRangeBitValve VkVideoEncodeRgbRangeCompressionFlagBitsVALVE = (1 << 1)
)

type VkVideoEncodeRgbChromaOffsetFlagBitsVALVE int64

var (
	VkVideoEncodeRgbChromaOffsetCositedEvenBitValve VkVideoEncodeRgbChromaOffsetFlagBitsVALVE = (1 << 0)
	VkVideoEncodeRgbChromaOffsetMidpointBitValve    VkVideoEncodeRgbChromaOffsetFlagBitsVALVE = (1 << 1)
)

type VkSwapchainImageUsageFlagBitsOHOS int64

var (
	VkSwapchainImageUsageSharedBitOhos VkSwapchainImageUsageFlagBitsOHOS = (1 << 0)
)

type VkDescriptorMappingSourceEXT int64

var (
	VkDescriptorMappingSourceHeapWithConstantOffsetExt     VkDescriptorMappingSourceEXT = 0
	VkDescriptorMappingSourceHeapWithPushIndexExt          VkDescriptorMappingSourceEXT = 1
	VkDescriptorMappingSourceHeapWithIndirectIndexExt      VkDescriptorMappingSourceEXT = 2
	VkDescriptorMappingSourceHeapWithIndirectIndexArrayExt VkDescriptorMappingSourceEXT = 3
	VkDescriptorMappingSourceResourceHeapDataExt           VkDescriptorMappingSourceEXT = 4
	VkDescriptorMappingSourcePushDataExt                   VkDescriptorMappingSourceEXT = 5
	VkDescriptorMappingSourcePushAddressExt                VkDescriptorMappingSourceEXT = 6
	VkDescriptorMappingSourceIndirectAddressExt            VkDescriptorMappingSourceEXT = 7
)

type VkSpirvResourceTypeFlagBitsEXT int64

var (
	VkSpirvResourceTypeAllExt                       VkSpirvResourceTypeFlagBitsEXT = 0
	VkSpirvResourceTypeSamplerBitExt                VkSpirvResourceTypeFlagBitsEXT = (1 << 0)
	VkSpirvResourceTypeSampledImageBitExt           VkSpirvResourceTypeFlagBitsEXT = (1 << 1)
	VkSpirvResourceTypeReadOnlyImageBitExt          VkSpirvResourceTypeFlagBitsEXT = (1 << 2)
	VkSpirvResourceTypeReadWriteImageBitExt         VkSpirvResourceTypeFlagBitsEXT = (1 << 3)
	VkSpirvResourceTypeCombinedSampledImageBitExt   VkSpirvResourceTypeFlagBitsEXT = (1 << 4)
	VkSpirvResourceTypeUniformBufferBitExt          VkSpirvResourceTypeFlagBitsEXT = (1 << 5)
	VkSpirvResourceTypeReadOnlyStorageBufferBitExt  VkSpirvResourceTypeFlagBitsEXT = (1 << 6)
	VkSpirvResourceTypeReadWriteStorageBufferBitExt VkSpirvResourceTypeFlagBitsEXT = (1 << 7)
)
