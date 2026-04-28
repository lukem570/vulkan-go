//go:build windows

package vk

/*
#cgo CFLAGS: -I./../../mod/include
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <string.h>
#include "volk_wrappers.h"
*/
import "C"
import "unsafe"

type Win32SurfaceCreateInfoKHR struct {
	Next      Structure
	Flags     Win32SurfaceCreateFlagsKHR
	Hinstance unsafe.Pointer
	Hwnd      unsafe.Pointer
}

func (s *Win32SurfaceCreateInfoKHR) GetType() StructureType {
	return StructureTypeWin32SurfaceCreateInfoKHR
}

func (s *Win32SurfaceCreateInfoKHR) toC() (unsafe.Pointer, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkWin32SurfaceCreateInfoKHR)(C.malloc(C.size_t(C.sizeof_VkWin32SurfaceCreateInfoKHR)))
	*p = C.VkWin32SurfaceCreateInfoKHR{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = nextPtr
	}
	val0 := C.VkWin32SurfaceCreateFlagsKHR(s.Flags)
	p.flags = val0
	ext1 := (C.HINSTANCE)(s.Hinstance)
	p.hinstance = ext1
	ext2 := (C.HWND)(s.Hwnd)
	p.hwnd = ext2
	return unsafe.Pointer(p), func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func (s *Win32SurfaceCreateInfoKHR) fromC(p *C.VkWin32SurfaceCreateInfoKHR) {
	s.Flags = Win32SurfaceCreateFlagsKHR(p.flags)
	s.Hinstance = unsafe.Pointer(p.hinstance)
	s.Hwnd = unsafe.Pointer(p.hwnd)
}

func (h Instance) CreateWin32SurfaceKHR(
	CreateInfo *Win32SurfaceCreateInfoKHR,
	Allocator *AllocationCallbacks,
) (*SurfaceKHR, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param CreateInfo
	var ptr1 *C.VkWin32SurfaceCreateInfoKHR
	if CreateInfo != nil {
		val2, cancel3 := CreateInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = (*C.VkWin32SurfaceCreateInfoKHR)(val2)
	}
	// param Allocator
	var ptr5 *C.VkAllocationCallbacks
	if Allocator != nil {
		val6, cancel7 := Allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = (*C.VkAllocationCallbacks)(val6)
	}
	var surfaceOut C.VkSurfaceKHR
	_result := C.fn_vkCreateWin32SurfaceKHR(C.VkInstance(unsafe.Pointer(h.handle)), ptr1, ptr5, &surfaceOut)
	if _result != C.VK_SUCCESS {
		return nil, Result(_result)
	}
	h8 := &SurfaceKHR{handle: unsafe.Pointer(surfaceOut)}
	return h8, nil
}

func (h PhysicalDevice) GetWin32PresentationSupportKHR(
	queueFamilyIndex uint32,
) bool {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param queueFamilyIndex
	val1 := C.uint32_t(queueFamilyIndex)
	_result := C.fn_vkGetPhysicalDeviceWin32PresentationSupportKHR(C.VkPhysicalDevice(unsafe.Pointer(h.handle)), val1)
	val2 := _result != 0
	return val2
}
