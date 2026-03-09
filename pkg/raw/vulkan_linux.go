//go:build linux

package vulkan

/*
#cgo CFLAGS: -I./../../mod/cp-Vulkan-Headers/include -I./../../mod/cp-volk
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <string.h>
#include "volk_wrappers.h"
*/
import "C"
import "unsafe"

type WaylandSurfaceCreateInfoKHR struct {
	Next    Structure
	Flags   WaylandSurfaceCreateFlagsKHR
	Display unsafe.Pointer
	Surface unsafe.Pointer
}

func (s *WaylandSurfaceCreateInfoKHR) GetType() StructureType {
	return StructureTypeWaylandSurfaceCreateInfoKHR
}

func (s *WaylandSurfaceCreateInfoKHR) toC() (unsafe.Pointer, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkWaylandSurfaceCreateInfoKHR)(C.malloc(C.size_t(C.sizeof_VkWaylandSurfaceCreateInfoKHR)))
	*p = C.VkWaylandSurfaceCreateInfoKHR{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = nextPtr
	}
	val0 := C.VkWaylandSurfaceCreateFlagsKHR(s.Flags)
	p.flags = val0
	ext1 := (*C.struct_wl_display)(s.Display)
	p.display = ext1
	ext2 := (*C.struct_wl_surface)(s.Surface)
	p.surface = ext2
	return unsafe.Pointer(p), func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func (s *WaylandSurfaceCreateInfoKHR) fromC(p *C.VkWaylandSurfaceCreateInfoKHR) {
	s.Flags = WaylandSurfaceCreateFlagsKHR(p.flags)
	s.Display = unsafe.Pointer(p.display)
	s.Surface = unsafe.Pointer(p.surface)
}

type XcbSurfaceCreateInfoKHR struct {
	Next       Structure
	Flags      XcbSurfaceCreateFlagsKHR
	Connection unsafe.Pointer
	Window     uint32
}

func (s *XcbSurfaceCreateInfoKHR) GetType() StructureType {
	return StructureTypeXcbSurfaceCreateInfoKHR
}

func (s *XcbSurfaceCreateInfoKHR) toC() (unsafe.Pointer, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkXcbSurfaceCreateInfoKHR)(C.malloc(C.size_t(C.sizeof_VkXcbSurfaceCreateInfoKHR)))
	*p = C.VkXcbSurfaceCreateInfoKHR{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = nextPtr
	}
	val0 := C.VkXcbSurfaceCreateFlagsKHR(s.Flags)
	p.flags = val0
	ext1 := (*C.xcb_connection_t)(s.Connection)
	p.connection = ext1
	ext2 := C.xcb_window_t(s.Window)
	p.window = ext2
	return unsafe.Pointer(p), func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func (s *XcbSurfaceCreateInfoKHR) fromC(p *C.VkXcbSurfaceCreateInfoKHR) {
	s.Flags = XcbSurfaceCreateFlagsKHR(p.flags)
	s.Connection = unsafe.Pointer(p.connection)
	s.Window = uint32(p.window)
}

type XlibSurfaceCreateInfoKHR struct {
	Next   Structure
	Flags  XlibSurfaceCreateFlagsKHR
	Dpy    unsafe.Pointer
	Window uintptr
}

func (s *XlibSurfaceCreateInfoKHR) GetType() StructureType {
	return StructureTypeXlibSurfaceCreateInfoKHR
}

func (s *XlibSurfaceCreateInfoKHR) toC() (unsafe.Pointer, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkXlibSurfaceCreateInfoKHR)(C.malloc(C.size_t(C.sizeof_VkXlibSurfaceCreateInfoKHR)))
	*p = C.VkXlibSurfaceCreateInfoKHR{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = nextPtr
	}
	val0 := C.VkXlibSurfaceCreateFlagsKHR(s.Flags)
	p.flags = val0
	ext1 := (*C.Display)(s.Dpy)
	p.dpy = ext1
	ext2 := C.Window(s.Window)
	p.window = ext2
	return unsafe.Pointer(p), func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func (s *XlibSurfaceCreateInfoKHR) fromC(p *C.VkXlibSurfaceCreateInfoKHR) {
	s.Flags = XlibSurfaceCreateFlagsKHR(p.flags)
	s.Dpy = unsafe.Pointer(p.dpy)
	s.Window = uintptr(p.window)
}

func (h Instance) CreateWaylandSurfaceKHR(
	CreateInfo *WaylandSurfaceCreateInfoKHR,
	Allocator *AllocationCallbacks,
) (*SurfaceKHR, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param CreateInfo
	var ptr1 *C.VkWaylandSurfaceCreateInfoKHR
	if CreateInfo != nil {
		val2, cancel3 := CreateInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = (*C.VkWaylandSurfaceCreateInfoKHR)(val2)
	}
	// param Allocator
	var ptr5 *C.VkAllocationCallbacks
	if Allocator != nil {
		val6, cancel7 := Allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = (*C.VkAllocationCallbacks)(val6)
	}
	var surfaceOut C.VkSurfaceKHR
	_result := C.fn_vkCreateWaylandSurfaceKHR(C.VkInstance(unsafe.Pointer(h.handle)), ptr1, ptr5, &surfaceOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h8 := &SurfaceKHR{handle: unsafe.Pointer(surfaceOut)}
	return h8, nil
}

func (h Instance) CreateXcbSurfaceKHR(
	CreateInfo *XcbSurfaceCreateInfoKHR,
	Allocator *AllocationCallbacks,
) (*SurfaceKHR, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param CreateInfo
	var ptr1 *C.VkXcbSurfaceCreateInfoKHR
	if CreateInfo != nil {
		val2, cancel3 := CreateInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = (*C.VkXcbSurfaceCreateInfoKHR)(val2)
	}
	// param Allocator
	var ptr5 *C.VkAllocationCallbacks
	if Allocator != nil {
		val6, cancel7 := Allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = (*C.VkAllocationCallbacks)(val6)
	}
	var surfaceOut C.VkSurfaceKHR
	_result := C.fn_vkCreateXcbSurfaceKHR(C.VkInstance(unsafe.Pointer(h.handle)), ptr1, ptr5, &surfaceOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h8 := &SurfaceKHR{handle: unsafe.Pointer(surfaceOut)}
	return h8, nil
}

func (h Instance) CreateXlibSurfaceKHR(
	CreateInfo *XlibSurfaceCreateInfoKHR,
	Allocator *AllocationCallbacks,
) (*SurfaceKHR, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param CreateInfo
	var ptr1 *C.VkXlibSurfaceCreateInfoKHR
	if CreateInfo != nil {
		val2, cancel3 := CreateInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = (*C.VkXlibSurfaceCreateInfoKHR)(val2)
	}
	// param Allocator
	var ptr5 *C.VkAllocationCallbacks
	if Allocator != nil {
		val6, cancel7 := Allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = (*C.VkAllocationCallbacks)(val6)
	}
	var surfaceOut C.VkSurfaceKHR
	_result := C.fn_vkCreateXlibSurfaceKHR(C.VkInstance(unsafe.Pointer(h.handle)), ptr1, ptr5, &surfaceOut)
	if _result != C.VK_SUCCESS {
		return nil, vkError(_result)
	}
	h8 := &SurfaceKHR{handle: unsafe.Pointer(surfaceOut)}
	return h8, nil
}

func (h PhysicalDevice) GetWaylandPresentationSupportKHR(
	queueFamilyIndex uint32,
	display unsafe.Pointer,
) bool {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param queueFamilyIndex
	val1 := C.uint32_t(queueFamilyIndex)
	// param display
	ext3 := (*C.struct_wl_display)(display)
	_result := C.fn_vkGetPhysicalDeviceWaylandPresentationSupportKHR(C.VkPhysicalDevice(unsafe.Pointer(h.handle)), val1, ext3)
	val4 := _result != 0
	return val4
}

func (h PhysicalDevice) GetXcbPresentationSupportKHR(
	queueFamilyIndex uint32,
	connection unsafe.Pointer,
	visual_id uint32,
) bool {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param queueFamilyIndex
	val1 := C.uint32_t(queueFamilyIndex)
	// param connection
	ext3 := (*C.xcb_connection_t)(connection)
	// param visual_id
	ext5 := C.xcb_visualid_t(visual_id)
	_result := C.fn_vkGetPhysicalDeviceXcbPresentationSupportKHR(C.VkPhysicalDevice(unsafe.Pointer(h.handle)), val1, ext3, ext5)
	val6 := _result != 0
	return val6
}

func (h PhysicalDevice) GetXlibPresentationSupportKHR(
	queueFamilyIndex uint32,
	dpy unsafe.Pointer,
	visualID uintptr,
) bool {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param queueFamilyIndex
	val1 := C.uint32_t(queueFamilyIndex)
	// param dpy
	ext3 := (*C.Display)(dpy)
	// param visualID
	ext5 := C.VisualID(visualID)
	_result := C.fn_vkGetPhysicalDeviceXlibPresentationSupportKHR(C.VkPhysicalDevice(unsafe.Pointer(h.handle)), val1, ext3, ext5)
	val6 := _result != 0
	return val6
}
