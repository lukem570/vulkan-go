package generator

import (
	"fmt"
	"sort"
	"strings"
)

const cgoHeader = `/*
#cgo CFLAGS: -I./../../mod/Vulkan-Headers/include -I./../../mod/volk
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
#cgo CFLAGS: -I./../../mod/Vulkan-Headers/include -I./../../mod/volk
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <string.h>
#include "volk_wrappers.h"
*/
import "C"
import "unsafe"
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
	if result := C.volkInitialize(); result != C.VK_SUCCESS {
		return vkError(result)
	}
	return nil
}

func vkError(r C.VkResult) error {
	return fmt.Errorf("vulkan error: VkResult(%d)", int(r))
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

// APIConstant is a single entry from the Vulkan API constants block.
type APIConstant struct {
	GoName string
	Value  string // already translated to valid Go syntax
	Type   string // "uint32", "uint64", "float32", "float64", or "" (untyped int)
}

type Registry struct {
	Enums        map[string]*Enum
	Bitmasks     map[string]*Bitmask
	Structs      map[string]*Structured
	Handles      map[string]*GoHandle
	Commands     map[string]*GoCommand
	Features     map[string]*Feature
	Extensions   map[string]*Extension
	FuncPointers map[string]*GoFuncPointer
	APIConstants []APIConstant
	Platforms    map[string]string // platform name → protect macro (e.g. "xlib" → "VK_USE_PLATFORM_XLIB_KHR")
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

func (r *Registry) AddEnumElement(extends string, el *EnumElement) {
	if e, ok := r.Enums[extends]; ok {
		el.Parent = e
		e.Elements = append(e.Elements, el)
	}
}

func (r *Registry) GeneratePackage(pkg string) string {
	var b strings.Builder

	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(cgoHeader)
	b.WriteString(runtimeHelpers)

	// API constants — grouped by type to avoid "mixed untyped constant" errors
	if len(r.APIConstants) > 0 {
		b.WriteString(generateAPIConstants(r.APIConstants))
	}

	// Handles (non-platform only)
	for _, k := range sortedKeys(r.Handles) {
		if r.Handles[k].Platform == "" {
			b.WriteString(r.Handles[k].Generate())
		}
	}

	// FuncPointers
	for _, k := range sortedKeys(r.FuncPointers) {
		b.WriteString(r.FuncPointers[k].Generate())
	}

	// Build reserved names set from structs/handles to detect naming collisions
	reserved := make(map[string]bool)
	for _, k := range sortedKeys(r.Structs) {
		reserved[r.Structs[k].GoName] = true
	}
	for _, k := range sortedKeys(r.Handles) {
		reserved[r.Handles[k].GoName] = true
	}

	// Enums
	for _, k := range sortedKeys(r.Enums) {
		r.Enums[k].SetReserved(reserved)
		b.WriteString(r.Enums[k].Generate())
	}

	// Bitmasks
	for _, k := range sortedKeys(r.Bitmasks) {
		b.WriteString(r.Bitmasks[k].Generate())
	}

	// Structs (non-platform only)
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		b.WriteString(s.GenerateStructure())
		b.WriteString(s.GenerateGetType())
		b.WriteString(s.GenerateToC())
		b.WriteString(s.GenerateFromC())
	}

	// Commands (non-platform only) — deduplicate by (ReceiverType, Name)
	seenMethods := make(map[string]bool)
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform != "" {
			continue
		}
		key := cmd.ReceiverType + "." + cmd.Name
		if seenMethods[key] {
			continue
		}
		seenMethods[key] = true
		b.WriteString(cmd.GenerateWrapper())
	}

	return b.String()
}

// GeneratePlatformFiles generates separate Go source files for each platform.
// Each file has the appropriate build tag and cgo preamble.
func (r *Registry) GeneratePlatformFiles(pkg string) []PlatformFile {
	// Group structs and commands by build tag
	type platformContent struct {
		buildTag  string
		platforms []string // may have multiple platforms per build tag (e.g. xlib+xcb+wayland → linux)
		structs   []*Structured
		commands  []*GoCommand
		handles   []*GoHandle
	}

	byTag := map[string]*platformContent{}
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[s.Platform]
		if tag == "" {
			continue
		}
		if byTag[tag] == nil {
			byTag[tag] = &platformContent{buildTag: tag}
		}
		byTag[tag].structs = append(byTag[tag].structs, s)
	}
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[cmd.Platform]
		if tag == "" {
			continue
		}
		if byTag[tag] == nil {
			byTag[tag] = &platformContent{buildTag: tag}
		}
		byTag[tag].commands = append(byTag[tag].commands, cmd)
	}
	for _, k := range sortedKeys(r.Handles) {
		h := r.Handles[k]
		if h.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[h.Platform]
		if tag == "" {
			continue
		}
		if byTag[tag] == nil {
			byTag[tag] = &platformContent{buildTag: tag}
		}
		byTag[tag].handles = append(byTag[tag].handles, h)
	}

	var files []PlatformFile
	for tag, pc := range byTag {
		if len(pc.structs) == 0 && len(pc.commands) == 0 && len(pc.handles) == 0 {
			continue
		}

		var b strings.Builder
		b.WriteString(fmt.Sprintf("//go:build %s\n\n", tag))
		b.WriteString("package " + pkg + "\n\n")
		b.WriteString(cgoPlatformHeader)

		// Handles
		for _, h := range pc.handles {
			b.WriteString(h.Generate())
		}

		// Structs
		for _, s := range pc.structs {
			b.WriteString(s.GenerateStructure())
			b.WriteString(s.GenerateGetType())
			b.WriteString(s.GenerateToC())
			b.WriteString(s.GenerateFromC())
		}

		// Commands
		seenMethods := make(map[string]bool)
		for _, cmd := range pc.commands {
			key := cmd.ReceiverType + "." + cmd.Name
			if seenMethods[key] {
				continue
			}
			seenMethods[key] = true
			b.WriteString(cmd.GenerateWrapper())
		}

		files = append(files, PlatformFile{
			BuildTag: tag,
			Content:  b.String(),
		})
	}

	return files
}

const callbacksFileHeader = `/*
#cgo CFLAGS: -I./../../mod/Vulkan-Headers/include -I./../../mod/volk
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

// GenerateCallbacksFile generates pkg/raw/callbacks.go containing the
// per-struct callback holder types, sync.Map registries, and //export bridge
// functions needed by the C trampolines.
func (r *Registry) GenerateCallbacksFile(pkg string) string {
	var b strings.Builder
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(callbacksFileHeader)

	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue // only non-platform structs for now
		}
		b.WriteString(s.GenerateCallbacksSupport())
	}
	return b.String()
}

// GenerateCHeader returns the content of volk_wrappers.h.
func (r *Registry) GenerateCHeader() string {
	var b strings.Builder
	b.WriteString("#ifndef VOLK_WRAPPERS_H\n#define VOLK_WRAPPERS_H\n\n")

	// Platform detection and includes
	b.WriteString(r.generatePlatformCIncludes())

	b.WriteString("#include \"volk.h\"\n")
	b.WriteString("#include <stdlib.h>\n\n")
	b.WriteString("#include <vulkan/vulkan.h>\n\n")

	// Non-platform commands
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			b.WriteString(cmd.GenerateCWrapperDecl())
		}
	}

	// Platform-specific commands grouped by protect macro
	r.writePlatformCBlocks(&b, func(cmd *GoCommand) string {
		return cmd.GenerateCWrapperDecl()
	})

	// Callback trampolines
	hasTrampolines := false
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		decls := s.GenerateCCallbackDecls()
		if decls != "" {
			if !hasTrampolines {
				b.WriteString("\n/* Callback trampolines */\n")
				hasTrampolines = true
			}
			b.WriteString(decls)
		}
	}

	b.WriteString("\n#endif\n")
	return b.String()
}

// GenerateCSource returns the content of volk_wrappers.c.
func (r *Registry) GenerateCSource() string {
	var b strings.Builder
	b.WriteString("#define VOLK_IMPLEMENTATION\n")

	// Platform detection and includes
	b.WriteString(r.generatePlatformCIncludes())

	b.WriteString("#include \"volk.h\"\n")
	b.WriteString("#include <vulkan/vulkan.h>\n")
	b.WriteString("#include \"volk_wrappers.h\"\n\n")

	// Non-platform commands
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			b.WriteString(cmd.GenerateCWrapperImpl())
		}
	}

	// Platform-specific commands grouped by protect macro
	r.writePlatformCBlocks(&b, func(cmd *GoCommand) string {
		return cmd.GenerateCWrapperImpl()
	})

	// Callback trampolines
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		b.WriteString(s.GenerateCCallbackImpls())
	}

	return b.String()
}

// generatePlatformCIncludes emits OS-guarded #define and #include lines for
// each platform that has commands in the registry.
func (r *Registry) generatePlatformCIncludes() string {
	// Collect which platforms are actually used
	usedPlatforms := map[string]bool{}
	for _, cmd := range r.Commands {
		if cmd.Platform != "" {
			usedPlatforms[cmd.Platform] = true
		}
	}
	for _, s := range r.Structs {
		if s.Platform != "" {
			usedPlatforms[s.Platform] = true
		}
	}
	if len(usedPlatforms) == 0 {
		return ""
	}

	// Group platforms by build tag (OS)
	type osGroup struct {
		guard     string // e.g. "defined(__linux__) || defined(__unix__)"
		platforms []string
	}

	osGuards := map[string]string{
		"linux":   "defined(__linux__) || defined(__unix__)",
		"windows": "defined(_WIN32)",
		"darwin":  "defined(__APPLE__)",
		"android": "defined(__ANDROID__)",
	}

	groups := map[string]*osGroup{}
	for platform := range usedPlatforms {
		tag := PlatformBuildTag[platform]
		if tag == "" {
			continue
		}
		if groups[tag] == nil {
			groups[tag] = &osGroup{guard: osGuards[tag]}
		}
		groups[tag].platforms = append(groups[tag].platforms, platform)
	}

	var b strings.Builder
	for _, tag := range sortedKeys(groups) {
		g := groups[tag]
		sort.Strings(g.platforms)
		b.WriteString(fmt.Sprintf("#if %s\n", g.guard))
		for _, platform := range g.platforms {
			if inc, ok := PlatformIncludes[platform]; ok {
				b.WriteString(inc + "\n")
			}
			if protect, ok := r.Platforms[platform]; ok {
				b.WriteString(fmt.Sprintf("#define %s\n", protect))
			}
		}
		b.WriteString("#endif\n\n")
	}
	return b.String()
}

// writePlatformCBlocks writes #ifdef-guarded blocks for platform-specific commands.
func (r *Registry) writePlatformCBlocks(b *strings.Builder, generate func(*GoCommand) string) {
	// Group commands by protect macro
	byProtect := map[string][]*GoCommand{}
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			continue
		}
		protect := r.Platforms[cmd.Platform]
		if protect == "" {
			continue
		}
		byProtect[protect] = append(byProtect[protect], cmd)
	}

	for _, protect := range sortedKeys(byProtect) {
		b.WriteString(fmt.Sprintf("\n#ifdef %s\n", protect))
		for _, cmd := range byProtect[protect] {
			b.WriteString(generate(cmd))
		}
		b.WriteString(fmt.Sprintf("#endif // %s\n", protect))
	}
}

// generateAPIConstants emits typed var declarations for each API constant.
// We use var instead of const because some values (^uint32(0), etc.) are not
// representable as Go untyped constants.
func generateAPIConstants(consts []APIConstant) string {
	var b strings.Builder
	b.WriteString("// Vulkan API constants\n")
	b.WriteString("var (\n")
	for _, c := range consts {
		if c.Type != "" {
			b.WriteString(fmt.Sprintf("\t%s %s = %s\n", c.GoName, c.Type, c.Value))
		} else {
			b.WriteString(fmt.Sprintf("\t%s = %s\n", c.GoName, c.Value))
		}
	}
	b.WriteString(")\n\n")
	return b.String()
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type Feature struct {
	Name     string
	Depends  []string // feature names this feature depends on
	Requires []RequireBlock
}

type Extension struct {
	Name      string
	Platform  string
	Supported string // e.g. "vulkan", "vulkan,vulkansc", "disabled"
	Requires  []RequireBlock
}

type RequireBlock struct {
	Commands []string
	Types    []string
	Enums    []string
}