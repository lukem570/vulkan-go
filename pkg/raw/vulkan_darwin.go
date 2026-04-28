//go:build darwin

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

type MacOSSurfaceCreateInfoMVK struct {
	Next  Structure
	Flags MacOSSurfaceCreateFlagsMVK
	View  unsafe.Pointer
}

func (s *MacOSSurfaceCreateInfoMVK) GetType() StructureType {
	return StructureTypeMacosSurfaceCreateInfoMvk
}

func (s *MacOSSurfaceCreateInfoMVK) toC() (unsafe.Pointer, func()) {
	cancels := make([]func(), 0)
	p := (*C.VkMacOSSurfaceCreateInfoMVK)(C.malloc(C.size_t(C.sizeof_VkMacOSSurfaceCreateInfoMVK)))
	*p = C.VkMacOSSurfaceCreateInfoMVK{}
	p.sType = C.VkStructureType(s.GetType())
	if s.Next != nil {
		nextPtr, nextCancel := s.Next.toC()
		cancels = append(cancels, nextCancel)
		p.pNext = nextPtr
	}
	val0 := C.VkMacOSSurfaceCreateFlagsMVK(s.Flags)
	p.flags = val0
	p.pView = s.View
	return unsafe.Pointer(p), func() {
		for _, cancel := range cancels {
			cancel()
		}
		C.free(unsafe.Pointer(p))
	}
}

func (s *MacOSSurfaceCreateInfoMVK) fromC(p *C.VkMacOSSurfaceCreateInfoMVK) {
	s.Flags = MacOSSurfaceCreateFlagsMVK(p.flags)
	s.View = unsafe.Pointer(p.pView)
}

func (h Instance) CreateMacOSSurfaceMVK(
	CreateInfo *MacOSSurfaceCreateInfoMVK,
	Allocator *AllocationCallbacks,
) (*SurfaceKHR, error) {
	cancels := make([]func(), 0)
	defer func() {
		for _, c := range cancels {
			c()
		}
	}()

	// param CreateInfo
	var ptr1 *C.VkMacOSSurfaceCreateInfoMVK
	if CreateInfo != nil {
		val2, cancel3 := CreateInfo.toC()
		cancels = append(cancels, cancel3)
		ptr1 = (*C.VkMacOSSurfaceCreateInfoMVK)(val2)
	}
	// param Allocator
	var ptr5 *C.VkAllocationCallbacks
	if Allocator != nil {
		val6, cancel7 := Allocator.toC()
		cancels = append(cancels, cancel7)
		ptr5 = (*C.VkAllocationCallbacks)(val6)
	}
	var surfaceOut C.VkSurfaceKHR
	_result := C.fn_vkCreateMacOSSurfaceMVK(C.VkInstance(unsafe.Pointer(h.handle)), ptr1, ptr5, &surfaceOut)
	if _result != C.VK_SUCCESS {
		return nil, Result(_result)
	}
	h8 := &SurfaceKHR{handle: unsafe.Pointer(surfaceOut)}
	return h8, nil
}
