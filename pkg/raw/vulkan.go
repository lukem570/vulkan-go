package vulkan

/*
#cgo CFLAGS: -I./../../mod/Vulkan-Headers/include
#cgo LDFLAGS: -lvulkan
#include <stdlib.h>
#include <vulkan/vulkan.h>
*/
import "C"
import "unsafe"

type Structure interface {
	getType() VkStructureType
	toC() (*C.void, func())
}
type DebugUtilsMessengerCreateInfoEXT struct {
	Next            Structure
	Flags           DebugUtilsMessengerCreateFlagsEXT
	MessageSeverity DebugUtilsMessageSeverityFlagsEXT
	MessageType     DebugUtilsMessageTypeFlagsEXT
	PfnUserCallback PFN_vkDebugUtilsMessengerCallbackEXT
	PUserData       void
}

func (s *DebugUtilsMessengerCreateInfoEXT) GetType() StructureType {
	return StructureTypeDebugUtilsMessengerCreateInfoEXT
}

func (s *DebugUtilsMessengerCreateInfoEXT) toC() (*C.VkDebugUtilsMessengerCreateInfoEXT, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDebugUtilsMessengerCreateInfoEXT)(C.malloc(C.size_t(C.sizeof_VkDebugUtilsMessengerCreateInfoEXT)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.VkDebugUtilsMessengerCreateFlagsEXT)(s.Flags)
	p.flags = val_0
	val_1 := (C.VkDebugUtilsMessageSeverityFlagsEXT)(s.MessageSeverity)
	p.messageSeverity = val_1
	val_2 := (C.VkDebugUtilsMessageTypeFlagsEXT)(s.MessageType)
	p.messageType = val_2
	val_3 := (C.PFN_vkDebugUtilsMessengerCallbackEXT)(s.PfnUserCallback)
	p.pfnUserCallback = val_3
	val_4 := (C.void)(s.PUserData)
	p.pUserData = val_4

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SwapchainCreateInfoKHR struct {
	Next                  Structure
	Flags                 SwapchainCreateFlagsKHR
	Surface               SurfaceKHR
	MinImageCount         uint32_t
	ImageFormat           Format
	ImageColorSpace       ColorSpaceKHR
	ImageExtent           Extent2D
	ImageArrayLayers      uint32_t
	ImageUsage            ImageUsageFlags
	ImageSharingMode      SharingMode
	QueueFamilyIndexCount uint32_t
	PQueueFamilyIndices   uint32_t
	PreTransform          SurfaceTransformFlagBitsKHR
	CompositeAlpha        CompositeAlphaFlagBitsKHR
	PresentMode           PresentModeKHR
	Clipped               Bool32
	OldSwapchain          SwapchainKHR
	OldSwapchain          SwapchainKHR
}

func (s *SwapchainCreateInfoKHR) GetType() StructureType {
	return StructureTypeSwapchainCreateInfoKHR
}

func (s *SwapchainCreateInfoKHR) toC() (*C.VkSwapchainCreateInfoKHR, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSwapchainCreateInfoKHR)(C.malloc(C.size_t(C.sizeof_VkSwapchainCreateInfoKHR)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.VkSwapchainCreateFlagsKHR)(s.Flags)
	p.flags = val_0
	val_1 := (C.VkSurfaceKHR)(s.Surface)
	p.surface = val_1
	val_2 := (C.uint32_t)(s.MinImageCount)
	p.minImageCount = val_2
	val_3 := (C.VkFormat)(s.ImageFormat)
	p.imageFormat = val_3
	val_4 := (C.VkColorSpaceKHR)(s.ImageColorSpace)
	p.imageColorSpace = val_4
	val_5 := (C.VkExtent2D)(s.ImageExtent)
	p.imageExtent = val_5
	val_6 := (C.uint32_t)(s.ImageArrayLayers)
	p.imageArrayLayers = val_6
	val_7 := (C.VkImageUsageFlags)(s.ImageUsage)
	p.imageUsage = val_7
	val_8 := (C.VkSharingMode)(s.ImageSharingMode)
	p.imageSharingMode = val_8
	val_9 := (C.uint32_t)(s.QueueFamilyIndexCount)
	p.queueFamilyIndexCount = val_9
	val_10 := (C.uint32_t)(s.PQueueFamilyIndices)
	p.pQueueFamilyIndices = val_10
	p.queueFamilyIndexCount = C.uint32_t(len(s.PQueueFamilyIndices))
	val_11 := (C.VkSurfaceTransformFlagBitsKHR)(s.PreTransform)
	p.preTransform = val_11
	val_12 := (C.VkCompositeAlphaFlagBitsKHR)(s.CompositeAlpha)
	p.compositeAlpha = val_12
	val_13 := (C.VkPresentModeKHR)(s.PresentMode)
	p.presentMode = val_13
	val_14 := (C.VkBool32)(s.Clipped)
	p.clipped = val_14
	val_15 := (C.VkSwapchainKHR)(s.OldSwapchain)
	p.oldSwapchain = val_15
	val_16 := (C.VkSwapchainKHR)(s.OldSwapchain)
	p.oldSwapchain = val_16

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type PresentInfoKHR struct {
	Next               Structure
	WaitSemaphoreCount uint32_t
	PWaitSemaphores    Semaphore
	SwapchainCount     uint32_t
	PSwapchains        SwapchainKHR
	PImageIndices      uint32_t
	PResults           Result
}

func (s *PresentInfoKHR) GetType() StructureType {
	return StructureTypePresentInfoKHR
}

func (s *PresentInfoKHR) toC() (*C.VkPresentInfoKHR, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkPresentInfoKHR)(C.malloc(C.size_t(C.sizeof_VkPresentInfoKHR)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.uint32_t)(s.WaitSemaphoreCount)
	p.waitSemaphoreCount = val_0
	val_1 := (C.VkSemaphore)(s.PWaitSemaphores)
	p.pWaitSemaphores = val_1
	p.waitSemaphoreCount = C.uint32_t(len(s.PWaitSemaphores))
	val_2 := (C.uint32_t)(s.SwapchainCount)
	p.swapchainCount = val_2
	val_3 := (C.VkSwapchainKHR)(s.PSwapchains)
	p.pSwapchains = val_3
	p.swapchainCount = C.uint32_t(len(s.PSwapchains))
	val_4 := (C.uint32_t)(s.PImageIndices)
	p.pImageIndices = val_4
	p.swapchainCount = C.uint32_t(len(s.PImageIndices))
	val_5 := (C.VkResult)(s.PResults)
	p.pResults = val_5
	p.swapchainCount = C.uint32_t(len(s.PResults))

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SurfaceFormatKHR struct {
	Format     Format
	ColorSpace ColorSpaceKHR
}

func (s *SurfaceFormatKHR) toC() (*C.VkSurfaceFormatKHR, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSurfaceFormatKHR)(C.malloc(C.size_t(C.sizeof_VkSurfaceFormatKHR)))
	val_0 := (C.VkFormat)(s.Format)
	p.format = val_0
	val_1 := (C.VkColorSpaceKHR)(s.ColorSpace)
	p.colorSpace = val_1

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type SurfaceCapabilitiesKHR struct {
	MinImageCount           uint32_t
	MaxImageCount           uint32_t
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32_t
	SupportedTransforms     SurfaceTransformFlagsKHR
	CurrentTransform        SurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha CompositeAlphaFlagsKHR
	SupportedUsageFlags     ImageUsageFlags
}

func (s *SurfaceCapabilitiesKHR) toC() (*C.VkSurfaceCapabilitiesKHR, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkSurfaceCapabilitiesKHR)(C.malloc(C.size_t(C.sizeof_VkSurfaceCapabilitiesKHR)))
	val_0 := (C.uint32_t)(s.MinImageCount)
	p.minImageCount = val_0
	val_1 := (C.uint32_t)(s.MaxImageCount)
	p.maxImageCount = val_1
	val_2 := (C.VkExtent2D)(s.CurrentExtent)
	p.currentExtent = val_2
	val_3 := (C.VkExtent2D)(s.MinImageExtent)
	p.minImageExtent = val_3
	val_4 := (C.VkExtent2D)(s.MaxImageExtent)
	p.maxImageExtent = val_4
	val_5 := (C.uint32_t)(s.MaxImageArrayLayers)
	p.maxImageArrayLayers = val_5
	val_6 := (C.VkSurfaceTransformFlagsKHR)(s.SupportedTransforms)
	p.supportedTransforms = val_6
	val_7 := (C.VkSurfaceTransformFlagBitsKHR)(s.CurrentTransform)
	p.currentTransform = val_7
	val_8 := (C.VkCompositeAlphaFlagsKHR)(s.SupportedCompositeAlpha)
	p.supportedCompositeAlpha = val_8
	val_9 := (C.VkImageUsageFlags)(s.SupportedUsageFlags)
	p.supportedUsageFlags = val_9

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}

func (s *Rect2D) toC() (*C.VkRect2D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkRect2D)(C.malloc(C.size_t(C.sizeof_VkRect2D)))
	val_0 := (C.VkOffset2D)(s.Offset)
	p.offset = val_0
	val_1 := (C.VkExtent2D)(s.Extent)
	p.extent = val_1

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DebugUtilsLabelEXT struct {
	Next       Structure
	PLabelName char
	Color      float
}

func (s *DebugUtilsLabelEXT) GetType() StructureType {
	return StructureTypeDebugUtilsLabelEXT
}

func (s *DebugUtilsLabelEXT) toC() (*C.VkDebugUtilsLabelEXT, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDebugUtilsLabelEXT)(C.malloc(C.size_t(C.sizeof_VkDebugUtilsLabelEXT)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.char)(s.PLabelName)
	p.pLabelName = val_0
	p.null - terminated = C.uint32_t(len(s.PLabelName))
	val_1 := (C.float)(s.Color)
	p.color = val_1

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AcquireNextImageInfoKHR struct {
	Next       Structure
	Swapchain  SwapchainKHR
	Timeout    uint64_t
	Semaphore  Semaphore
	Fence      Fence
	DeviceMask uint32_t
}

func (s *AcquireNextImageInfoKHR) GetType() StructureType {
	return StructureTypeAcquireNextImageInfoKHR
}

func (s *AcquireNextImageInfoKHR) toC() (*C.VkAcquireNextImageInfoKHR, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAcquireNextImageInfoKHR)(C.malloc(C.size_t(C.sizeof_VkAcquireNextImageInfoKHR)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.VkSwapchainKHR)(s.Swapchain)
	p.swapchain = val_0
	val_1 := (C.uint64_t)(s.Timeout)
	p.timeout = val_1
	val_2 := (C.VkSemaphore)(s.Semaphore)
	p.semaphore = val_2
	val_3 := (C.VkFence)(s.Fence)
	p.fence = val_3
	val_4 := (C.uint32_t)(s.DeviceMask)
	p.deviceMask = val_4

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DebugUtilsObjectTagInfoEXT struct {
	Next         Structure
	ObjectType   ObjectType
	ObjectHandle uint64_t
	TagName      uint64_t
	TagSize      size_t
	PTag         void
}

func (s *DebugUtilsObjectTagInfoEXT) GetType() StructureType {
	return StructureTypeDebugUtilsObjectTagInfoEXT
}

func (s *DebugUtilsObjectTagInfoEXT) toC() (*C.VkDebugUtilsObjectTagInfoEXT, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDebugUtilsObjectTagInfoEXT)(C.malloc(C.size_t(C.sizeof_VkDebugUtilsObjectTagInfoEXT)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.VkObjectType)(s.ObjectType)
	p.objectType = val_0
	val_1 := (C.uint64_t)(s.ObjectHandle)
	p.objectHandle = val_1
	val_2 := (C.uint64_t)(s.TagName)
	p.tagName = val_2
	val_3 := (C.size_t)(s.TagSize)
	p.tagSize = val_3
	val_4 := (C.void)(s.PTag)
	p.pTag = val_4
	p.tagSize = C.uint32_t(len(s.PTag))

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Offset2D struct {
	X int32_t
	Y int32_t
}

func (s *Offset2D) toC() (*C.VkOffset2D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkOffset2D)(C.malloc(C.size_t(C.sizeof_VkOffset2D)))
	val_0 := (C.int32_t)(s.X)
	p.x = val_0
	val_1 := (C.int32_t)(s.Y)
	p.y = val_1

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type Extent2D struct {
	Width  uint32_t
	Height uint32_t
}

func (s *Extent2D) toC() (*C.VkExtent2D, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkExtent2D)(C.malloc(C.size_t(C.sizeof_VkExtent2D)))
	val_0 := (C.uint32_t)(s.Width)
	p.width = val_0
	val_1 := (C.uint32_t)(s.Height)
	p.height = val_1

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type AllocationCallbacks struct {
	PUserData             void
	PfnAllocation         PFN_vkAllocationFunction
	PfnReallocation       PFN_vkReallocationFunction
	PfnFree               PFN_vkFreeFunction
	PfnInternalAllocation PFN_vkInternalAllocationNotification
	PfnInternalFree       PFN_vkInternalFreeNotification
}

func (s *AllocationCallbacks) toC() (*C.VkAllocationCallbacks, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkAllocationCallbacks)(C.malloc(C.size_t(C.sizeof_VkAllocationCallbacks)))
	val_0 := (C.void)(s.PUserData)
	p.pUserData = val_0
	val_1 := (C.PFN_vkAllocationFunction)(s.PfnAllocation)
	p.pfnAllocation = val_1
	val_2 := (C.PFN_vkReallocationFunction)(s.PfnReallocation)
	p.pfnReallocation = val_2
	val_3 := (C.PFN_vkFreeFunction)(s.PfnFree)
	p.pfnFree = val_3
	val_4 := (C.PFN_vkInternalAllocationNotification)(s.PfnInternalAllocation)
	p.pfnInternalAllocation = val_4
	val_5 := (C.PFN_vkInternalFreeNotification)(s.PfnInternalFree)
	p.pfnInternalFree = val_5

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DeviceGroupPresentCapabilitiesKHR struct {
	Next        Structure
	PresentMask uint32_t
	Modes       DeviceGroupPresentModeFlagsKHR
}

func (s *DeviceGroupPresentCapabilitiesKHR) GetType() StructureType {
	return StructureTypeDeviceGroupPresentCapabilitiesKHR
}

func (s *DeviceGroupPresentCapabilitiesKHR) toC() (*C.VkDeviceGroupPresentCapabilitiesKHR, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDeviceGroupPresentCapabilitiesKHR)(C.malloc(C.size_t(C.sizeof_VkDeviceGroupPresentCapabilitiesKHR)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.uint32_t)(s.PresentMask)
	p.presentMask = val_0
	val_1 := (C.VkDeviceGroupPresentModeFlagsKHR)(s.Modes)
	p.modes = val_1

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DebugUtilsMessengerCallbackDataEXT struct {
	Next             Structure
	Flags            DebugUtilsMessengerCallbackDataFlagsEXT
	PMessageIdName   char
	MessageIdNumber  int32_t
	PMessage         char
	QueueLabelCount  uint32_t
	PQueueLabels     DebugUtilsLabelEXT
	CmdBufLabelCount uint32_t
	PCmdBufLabels    DebugUtilsLabelEXT
	ObjectCount      uint32_t
	PObjects         DebugUtilsObjectNameInfoEXT
}

func (s *DebugUtilsMessengerCallbackDataEXT) GetType() StructureType {
	return StructureTypeDebugUtilsMessengerCallbackDataEXT
}

func (s *DebugUtilsMessengerCallbackDataEXT) toC() (*C.VkDebugUtilsMessengerCallbackDataEXT, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDebugUtilsMessengerCallbackDataEXT)(C.malloc(C.size_t(C.sizeof_VkDebugUtilsMessengerCallbackDataEXT)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.VkDebugUtilsMessengerCallbackDataFlagsEXT)(s.Flags)
	p.flags = val_0
	val_1 := (C.char)(s.PMessageIdName)
	p.pMessageIdName = val_1
	p.null - terminated = C.uint32_t(len(s.PMessageIdName))
	val_2 := (C.int32_t)(s.MessageIdNumber)
	p.messageIdNumber = val_2
	val_3 := (C.char)(s.PMessage)
	p.pMessage = val_3
	p.null - terminated = C.uint32_t(len(s.PMessage))
	val_4 := (C.uint32_t)(s.QueueLabelCount)
	p.queueLabelCount = val_4
	val_5 := (C.VkDebugUtilsLabelEXT)(s.PQueueLabels)
	p.pQueueLabels = val_5
	p.queueLabelCount = C.uint32_t(len(s.PQueueLabels))
	val_6 := (C.uint32_t)(s.CmdBufLabelCount)
	p.cmdBufLabelCount = val_6
	val_7 := (C.VkDebugUtilsLabelEXT)(s.PCmdBufLabels)
	p.pCmdBufLabels = val_7
	p.cmdBufLabelCount = C.uint32_t(len(s.PCmdBufLabels))
	val_8 := (C.uint32_t)(s.ObjectCount)
	p.objectCount = val_8
	val_9 := (C.VkDebugUtilsObjectNameInfoEXT)(s.PObjects)
	p.pObjects = val_9
	p.objectCount = C.uint32_t(len(s.PObjects))

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

type DebugUtilsObjectNameInfoEXT struct {
	Next         Structure
	ObjectType   ObjectType
	ObjectHandle uint64_t
	PObjectName  char
}

func (s *DebugUtilsObjectNameInfoEXT) GetType() StructureType {
	return StructureTypeDebugUtilsObjectNameInfoEXT
}

func (s *DebugUtilsObjectNameInfoEXT) toC() (*C.VkDebugUtilsObjectNameInfoEXT, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkDebugUtilsObjectNameInfoEXT)(C.malloc(C.size_t(C.sizeof_VkDebugUtilsObjectNameInfoEXT)))
	p.sType = (C.VkStructureType)(s.GetType())

	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = unsafe.Pointer(next)
	}

	val_0 := (C.VkObjectType)(s.ObjectType)
	p.objectType = val_0
	val_1 := (C.uint64_t)(s.ObjectHandle)
	p.objectHandle = val_1
	val_2 := (C.char)(s.PObjectName)
	p.pObjectName = val_2
	p.null - terminated = C.uint32_t(len(s.PObjectName))

	return p, func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func vkSetDebugUtilsObjectNameEXT(
	device Device,
	pNameInfo DebugUtilsObjectNameInfoEXT,
) {
}
func vkQueueInsertDebugUtilsLabelEXT(
	queue Queue,
	pLabelInfo DebugUtilsLabelEXT,
) {
}
func vkSubmitDebugUtilsMessageEXT(
	instance Instance,
	messageSeverity DebugUtilsMessageSeverityFlagBitsEXT,
	messageTypes DebugUtilsMessageTypeFlagsEXT,
	pCallbackData DebugUtilsMessengerCallbackDataEXT,
) {
}
func vkDestroySurfaceKHR(
	instance Instance,
	surface SurfaceKHR,
	pAllocator AllocationCallbacks,
) {
}
func vkAcquireNextImage2KHR(
	device Device,
	pAcquireInfo AcquireNextImageInfoKHR,
	pImageIndex uint32_t,
) {
}
func vkGetPhysicalDeviceSurfacePresentModesKHR(
	physicalDevice PhysicalDevice,
	surface SurfaceKHR,
	pPresentModeCount uint32_t,
	pPresentModes PresentModeKHR,
) {
}
func vkCreateSwapchainKHR(
	device Device,
	pCreateInfo SwapchainCreateInfoKHR,
	pCreateInfo SwapchainCreateInfoKHR,
	pAllocator AllocationCallbacks,
	pSwapchain SwapchainKHR,
) {
}
func vkDestroySwapchainKHR(
	device Device,
	swapchain SwapchainKHR,
	pAllocator AllocationCallbacks,
) {
}
func vkDestroyDebugUtilsMessengerEXT(
	instance Instance,
	messenger DebugUtilsMessengerEXT,
	pAllocator AllocationCallbacks,
) {
}
func vkAcquireNextImageKHR(
	device Device,
	swapchain SwapchainKHR,
	timeout uint64_t,
	semaphore Semaphore,
	fence Fence,
	pImageIndex uint32_t,
) {
}
func vkQueuePresentKHR(
	queue Queue,
	pPresentInfo PresentInfoKHR,
) {
}
func vkQueueEndDebugUtilsLabelEXT(
	queue Queue,
) {
}
func vkGetPhysicalDeviceSurfaceFormatsKHR(
	physicalDevice PhysicalDevice,
	surface SurfaceKHR,
	pSurfaceFormatCount uint32_t,
	pSurfaceFormats SurfaceFormatKHR,
) {
}
func vkSetDebugUtilsObjectTagEXT(
	device Device,
	pTagInfo DebugUtilsObjectTagInfoEXT,
) {
}
func vkGetDeviceGroupSurfacePresentModesKHR(
	device Device,
	surface SurfaceKHR,
	pModes DeviceGroupPresentModeFlagsKHR,
) {
}
func vkCmdBeginDebugUtilsLabelEXT(
	commandBuffer CommandBuffer,
	pLabelInfo DebugUtilsLabelEXT,
) {
}
func vkGetPhysicalDevicePresentRectanglesKHR(
	physicalDevice PhysicalDevice,
	surface SurfaceKHR,
	pRectCount uint32_t,
	pRects Rect2D,
) {
}
func vkGetDeviceGroupPresentCapabilitiesKHR(
	device Device,
	pDeviceGroupPresentCapabilities DeviceGroupPresentCapabilitiesKHR,
) {
}
func vkGetPhysicalDeviceSurfaceSupportKHR(
	physicalDevice PhysicalDevice,
	queueFamilyIndex uint32_t,
	surface SurfaceKHR,
	pSupported Bool32,
) {
}
func vkGetSwapchainImagesKHR(
	device Device,
	swapchain SwapchainKHR,
	pSwapchainImageCount uint32_t,
	pSwapchainImages Image,
) {
}
func vkQueueBeginDebugUtilsLabelEXT(
	queue Queue,
	pLabelInfo DebugUtilsLabelEXT,
) {
}
func vkCmdEndDebugUtilsLabelEXT(
	commandBuffer CommandBuffer,
) {
}
func vkCreateDebugUtilsMessengerEXT(
	instance Instance,
	pCreateInfo DebugUtilsMessengerCreateInfoEXT,
	pAllocator AllocationCallbacks,
	pMessenger DebugUtilsMessengerEXT,
) {
}
func vkCmdInsertDebugUtilsLabelEXT(
	commandBuffer CommandBuffer,
	pLabelInfo DebugUtilsLabelEXT,
) {
}
func vkGetPhysicalDeviceSurfaceCapabilitiesKHR(
	physicalDevice PhysicalDevice,
	surface SurfaceKHR,
	pSurfaceCapabilities SurfaceCapabilitiesKHR,
) {
}
