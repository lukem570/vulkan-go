package vk

/*
#cgo CFLAGS: -I./../../mod/include
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include "volk_wrappers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

type allocationCallbacksCallbacks struct {
	Allocation         AllocationFunction
	Reallocation       ReallocationFunction
	Free               FreeFunction
	InternalAllocation InternalAllocationNotification
	InternalFree       InternalFreeNotification
	UserData           unsafe.Pointer
}

var allocationCallbacksCallbackRegistry sync.Map

//export go_bridge_AllocationCallbacks_Allocation
func go_bridge_AllocationCallbacks_Allocation(p0 unsafe.Pointer, p1 C.size_t, p2 C.size_t, p3 C.VkSystemAllocationScope) unsafe.Pointer {
	v, ok := allocationCallbacksCallbackRegistry.Load(p3)
	if !ok {
		return nil
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.Allocation == nil {
		return nil
	}
	result := holder.Allocation(p0, uintptr(p1), uintptr(p2), SystemAllocationScope(p3))
	return result
}

//export go_bridge_AllocationCallbacks_Reallocation
func go_bridge_AllocationCallbacks_Reallocation(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.size_t, p3 C.size_t, p4 C.VkSystemAllocationScope) unsafe.Pointer {
	v, ok := allocationCallbacksCallbackRegistry.Load(p4)
	if !ok {
		return nil
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.Reallocation == nil {
		return nil
	}
	result := holder.Reallocation(p0, p1, uintptr(p2), uintptr(p3), SystemAllocationScope(p4))
	return result
}

//export go_bridge_AllocationCallbacks_Free
func go_bridge_AllocationCallbacks_Free(p0 unsafe.Pointer, p1 unsafe.Pointer) {
	v, ok := allocationCallbacksCallbackRegistry.Load(p1)
	if !ok {
		return
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.Free == nil {
		return
	}
	holder.Free(p0, p1)
}

//export go_bridge_AllocationCallbacks_InternalAllocation
func go_bridge_AllocationCallbacks_InternalAllocation(p0 unsafe.Pointer, p1 C.size_t, p2 C.VkInternalAllocationType, p3 C.VkSystemAllocationScope) {
	v, ok := allocationCallbacksCallbackRegistry.Load(p3)
	if !ok {
		return
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.InternalAllocation == nil {
		return
	}
	holder.InternalAllocation(p0, uintptr(p1), InternalAllocationType(p2), SystemAllocationScope(p3))
}

//export go_bridge_AllocationCallbacks_InternalFree
func go_bridge_AllocationCallbacks_InternalFree(p0 unsafe.Pointer, p1 C.size_t, p2 C.VkInternalAllocationType, p3 C.VkSystemAllocationScope) {
	v, ok := allocationCallbacksCallbackRegistry.Load(p3)
	if !ok {
		return
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.InternalFree == nil {
		return
	}
	holder.InternalFree(p0, uintptr(p1), InternalAllocationType(p2), SystemAllocationScope(p3))
}

type debugUtilsMessengerCreateInfoEXTCallbacks struct {
	UserCallback DebugUtilsMessengerCallbackEXT
	UserData     unsafe.Pointer
}

var debugUtilsMessengerCreateInfoEXTCallbackRegistry sync.Map

//export go_bridge_DebugUtilsMessengerCreateInfoEXT_UserCallback
func go_bridge_DebugUtilsMessengerCreateInfoEXT_UserCallback(p0 C.VkDebugUtilsMessageSeverityFlagBitsEXT, p1 C.VkDebugUtilsMessageTypeFlagsEXT, p2 *C.VkDebugUtilsMessengerCallbackDataEXT, p3 unsafe.Pointer) C.VkBool32 {
	v, ok := debugUtilsMessengerCreateInfoEXTCallbackRegistry.Load(p3)
	if !ok {
		return C.VkBool32(0)
	}
	holder := v.(*debugUtilsMessengerCreateInfoEXTCallbacks)
	if holder.UserCallback == nil {
		return C.VkBool32(0)
	}
	result := holder.UserCallback(DebugUtilsMessageSeverityFlagBitsEXT(p0), DebugUtilsMessageTypeFlagsEXT(p1), func() *DebugUtilsMessengerCallbackDataEXT {
		var _go_p2 DebugUtilsMessengerCallbackDataEXT
		_go_p2.fromC(p2)
		return &_go_p2
	}(), p3)
	if result {
		return C.VkBool32(1)
	}
	return C.VkBool32(0)
}
