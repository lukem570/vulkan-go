package generator

// APIConstant is a single entry from the Vulkan API constants block.
type APIConstant struct {
	GoName string
	Value  string // already translated to valid Go syntax
	Type   string // "uint32", "uint64", "float32", "float64", or "" (untyped int)
}

// Registry is the full parsed Vulkan registry. It contains Features and
// Extensions for use by Filter to build a ReducedRegistry.
type Registry struct {
	Enums        map[string]*Enum
	Bitmasks     map[string]*Bitmask
	EnumAliases  map[string]*EnumAlias
	Structs      map[string]*Structured
	Handles      map[string]*GoHandle
	Commands     map[string]*GoCommand
	Features     map[string]*Feature
	Extensions   map[string]*Extension
	FuncPointers map[string]*GoFuncPointer
	STypes       map[string]string
	APIConstants []APIConstant
	Platforms    map[string]string // platform name → protect macro (e.g. "xlib" → "VK_USE_PLATFORM_XLIB_KHR")
}

// ReducedRegistry is the filtered subset of the Registry that is ready for
// code generation. It does not include Features or Extensions.
type ReducedRegistry struct {
	Enums        map[string]*Enum
	Bitmasks     map[string]*Bitmask
	EnumAliases  map[string]*EnumAlias
	Structs      map[string]*Structured
	Handles      map[string]*GoHandle
	Commands     map[string]*GoCommand
	FuncPointers map[string]*GoFuncPointer
	STypes       map[string]string
	APIConstants []APIConstant
	Platforms    map[string]string
}

// PlatformBuildTag maps a Vulkan platform name to a Go build tag.
var PlatformBuildTag = map[string]string{
	"xlib":    "linux",
	"xcb":     "linux",
	"wayland": "linux",
	"win32":   "windows",
	"metal":   "darwin",
	"macos":   "darwin",
	"android": "android",
}

// PlatformIncludes maps a platform name to the C headers it needs.
var PlatformIncludes = map[string]string{
	"xlib":    "#include <X11/Xlib.h>",
	"xcb":     "#include <xcb/xcb.h>",
	"wayland": "#include <wayland-client.h>",
	"win32":   "#include <windows.h>",
}

// PlatformFile represents a generated Go source file for a specific platform.
type PlatformFile struct {
	BuildTag string // e.g. "linux"
	Platform string // e.g. "xlib"
	Content  string
}

// AddEnumElement adds an enum element to the named enum in the Registry.
func (r *Registry) AddEnumElement(extends string, el *EnumElement) {
	if e, ok := r.Enums[extends]; ok {
		el.Parent = e
		e.Elements = append(e.Elements, el)
	}
}

// ---------------------------------------------------------------------------
// C / CGo preamble constants used in generated output files
// ---------------------------------------------------------------------------

const cgoHeader = `/*
#cgo CFLAGS: -I./../../mod/include
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <string.h>
#include "volk_wrappers.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)
`

const cgoPlatformHeader = `/*
#cgo CFLAGS: -I./../../mod/include
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <string.h>
#include "volk_wrappers.h"
*/
import "C"
import "unsafe"
`

const callbacksFileHeader = `/*
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
`

const runtimeHelpers = `
// Structure is the interface implemented by all Vulkan structs that carry
// an sType / pNext chain.
type Structure interface {
	GetType() StructureType
	toC() (unsafe.Pointer, func())
}

// Initialize loads the Vulkan library via Volk.
// Must be called before any other Vulkan function.
func Initialize() error {
	if result := C.vulkan_platform_initialize(); result != C.VK_SUCCESS {
		return Result(result)
	}
	return nil
}

// LoadInstance loads instance-level Vulkan function pointers via Volk.
// Must be called after creating a VkInstance.
func LoadInstance(instance *Instance) {
	C.volkLoadInstance(C.VkInstance(unsafe.Pointer(instance.handle)))
}

// LoadDevice loads device-level Vulkan function pointers via Volk.
// Must be called after creating a VkDevice.
func LoadDevice(device *Device) {
	C.volkLoadDevice(C.VkDevice(unsafe.Pointer(device.handle)))
}

// SurfaceKHRFromHandle wraps a raw VkSurfaceKHR handle value into a
// *SurfaceKHR. Use this when you have the actual handle value (e.g. from
// platform-specific surface creation).
func SurfaceKHRFromHandle(handle uintptr) *SurfaceKHR {
	return &SurfaceKHR{handle: unsafe.Pointer(handle)}
}

// SurfaceKHRFromGLFW wraps the uintptr returned by go-gl/glfw's
// Window.CreateWindowSurface into a *SurfaceKHR. GLFW returns a pointer to
// the VkSurfaceKHR, so this function dereferences it to get the actual handle.
// Must be called immediately after CreateWindowSurface in the same goroutine.
//
// Usage with go-gl/glfw:
//
//	surfPtr, err := window.CreateWindowSurface((*byte)(instance.Handle()), nil)
//	surface := vk.SurfaceKHRFromGLFW(surfPtr)
func SurfaceKHRFromGLFW(glfwSurface uintptr) *SurfaceKHR {
	handle := *(*uintptr)(unsafe.Pointer(glfwSurface))
	return &SurfaceKHR{handle: unsafe.Pointer(handle)}
}
`

// platformInitImpl is injected into volk_wrappers.c.
// On macOS, volk's built-in dlopen doesn't search Homebrew paths, so we do the
// search ourselves and call volkInitializeCustom to bypass volk's loader entirely.
const platformInitImpl = `VkResult vulkan_platform_initialize(void) {
#if defined(__APPLE__)
	const char* paths[] = {
		"libvulkan.dylib",
		"libvulkan.1.dylib",
		"/usr/local/lib/libvulkan.dylib",
		"/opt/homebrew/lib/libvulkan.dylib",
		"libMoltenVK.dylib",
		"/opt/homebrew/lib/libMoltenVK.dylib",
		NULL,
	};
	void* module = NULL;
	for (int i = 0; paths[i] != NULL && module == NULL; i++) {
		module = dlopen(paths[i], RTLD_NOW | RTLD_LOCAL);
	}
	if (!module) return VK_ERROR_INITIALIZATION_FAILED;
	PFN_vkGetInstanceProcAddr proc = (PFN_vkGetInstanceProcAddr)dlsym(module, "vkGetInstanceProcAddr");
	if (!proc) return VK_ERROR_INITIALIZATION_FAILED;
	volkInitializeCustom(proc);
	return VK_SUCCESS;
#else
	return volkInitialize();
#endif
}

`
