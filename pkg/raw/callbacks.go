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
	PfnAllocation         AllocationFunction
	PfnReallocation       ReallocationFunction
	PfnFree               FreeFunction
	PfnInternalAllocation InternalAllocationNotification
	PfnInternalFree       InternalFreeNotification
	UserData              unsafe.Pointer
}

var allocationCallbacksCallbackRegistry sync.Map

//export go_bridge_AllocationCallbacks_PfnAllocation
func go_bridge_AllocationCallbacks_PfnAllocation(p0 unsafe.Pointer, p1 C.size_t, p2 C.size_t, p3 C.VkSystemAllocationScope) unsafe.Pointer {
	v, ok := allocationCallbacksCallbackRegistry.Load(p3)
	if !ok {
		return nil
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.PfnAllocation == nil {
		return nil
	}
	result := holder.PfnAllocation(p0, uintptr(p1), uintptr(p2), SystemAllocationScope(p3))
	return result
}

//export go_bridge_AllocationCallbacks_PfnReallocation
func go_bridge_AllocationCallbacks_PfnReallocation(p0 unsafe.Pointer, p1 unsafe.Pointer, p2 C.size_t, p3 C.size_t, p4 C.VkSystemAllocationScope) unsafe.Pointer {
	v, ok := allocationCallbacksCallbackRegistry.Load(p4)
	if !ok {
		return nil
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.PfnReallocation == nil {
		return nil
	}
	result := holder.PfnReallocation(p0, p1, uintptr(p2), uintptr(p3), SystemAllocationScope(p4))
	return result
}

//export go_bridge_AllocationCallbacks_PfnFree
func go_bridge_AllocationCallbacks_PfnFree(p0 unsafe.Pointer, p1 unsafe.Pointer) {
	v, ok := allocationCallbacksCallbackRegistry.Load(p1)
	if !ok {
		return
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.PfnFree == nil {
		return
	}
	holder.PfnFree(p0, p1)
}

//export go_bridge_AllocationCallbacks_PfnInternalAllocation
func go_bridge_AllocationCallbacks_PfnInternalAllocation(p0 unsafe.Pointer, p1 C.size_t, p2 C.VkInternalAllocationType, p3 C.VkSystemAllocationScope) {
	v, ok := allocationCallbacksCallbackRegistry.Load(p3)
	if !ok {
		return
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.PfnInternalAllocation == nil {
		return
	}
	holder.PfnInternalAllocation(p0, uintptr(p1), InternalAllocationType(p2), SystemAllocationScope(p3))
}

//export go_bridge_AllocationCallbacks_PfnInternalFree
func go_bridge_AllocationCallbacks_PfnInternalFree(p0 unsafe.Pointer, p1 C.size_t, p2 C.VkInternalAllocationType, p3 C.VkSystemAllocationScope) {
	v, ok := allocationCallbacksCallbackRegistry.Load(p3)
	if !ok {
		return
	}
	holder := v.(*allocationCallbacksCallbacks)
	if holder.PfnInternalFree == nil {
		return
	}
	holder.PfnInternalFree(p0, uintptr(p1), InternalAllocationType(p2), SystemAllocationScope(p3))
}

type debugUtilsMessengerCreateInfoEXTCallbacks struct {
	PfnUserCallback DebugUtilsMessengerCallbackEXT
	UserData        unsafe.Pointer
}

var debugUtilsMessengerCreateInfoEXTCallbackRegistry sync.Map

//export go_bridge_DebugUtilsMessengerCreateInfoEXT_PfnUserCallback
func go_bridge_DebugUtilsMessengerCreateInfoEXT_PfnUserCallback(p0 C.VkDebugUtilsMessageSeverityFlagBitsEXT, p1 C.VkDebugUtilsMessageTypeFlagsEXT, p2 *C.VkDebugUtilsMessengerCallbackDataEXT, p3 unsafe.Pointer) C.VkBool32 {
	v, ok := debugUtilsMessengerCreateInfoEXTCallbackRegistry.Load(p3)
	if !ok {
		return C.VkBool32(0)
	}
	holder := v.(*debugUtilsMessengerCreateInfoEXTCallbacks)
	if holder.PfnUserCallback == nil {
		return C.VkBool32(0)
	}
	result := holder.PfnUserCallback(DebugUtilsMessageSeverityFlagBitsEXT(p0), DebugUtilsMessageTypeFlagsEXT(p1), func() *DebugUtilsMessengerCallbackDataEXT {
		var _go_p2 DebugUtilsMessengerCallbackDataEXT
		_go_p2.fromC(p2)
		return &_go_p2
	}(), p3)
	if result {
		return C.VkBool32(1)
	}
	return C.VkBool32(0)
}
