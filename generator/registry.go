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

	// Handles
	for _, k := range sortedKeys(r.Handles) {
		b.WriteString(r.Handles[k].Generate())
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

	// Structs
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		b.WriteString(s.GenerateStructure())
		b.WriteString(s.GenerateGetType())
		b.WriteString(s.GenerateToC())
		b.WriteString(s.GenerateFromC())
	}

	// Commands — deduplicate by (ReceiverType, Name) to avoid collisions
	// from different C names mapping to the same Go method name.
	seenMethods := make(map[string]bool)
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		key := cmd.ReceiverType + "." + cmd.Name
		if seenMethods[key] {
			continue
		}
		seenMethods[key] = true
		b.WriteString(cmd.GenerateWrapper())
	}

	return b.String()
}

// GenerateCHeader returns the content of volk_wrappers.h.
func (r *Registry) GenerateCHeader() string {
	var b strings.Builder
	b.WriteString("#ifndef VOLK_WRAPPERS_H\n#define VOLK_WRAPPERS_H\n\n")
	b.WriteString("#include \"volk.h\"\n")
	b.WriteString("#include <stdlib.h>\n\n")
	b.WriteString("#include <vulkan/vulkan.h>\n\n")
	for _, k := range sortedKeys(r.Commands) {
		b.WriteString(r.Commands[k].GenerateCWrapperDecl())
	}
	b.WriteString("\n#endif\n")
	return b.String()
}

// GenerateCSource returns the content of volk_wrappers.c.
func (r *Registry) GenerateCSource() string {
	var b strings.Builder
	b.WriteString("#define VOLK_IMPLEMENTATION\n")
	b.WriteString("#include \"volk.h\"\n")
	b.WriteString("#include <vulkan/vulkan.h>\n")
	b.WriteString("#include \"volk_wrappers.h\"\n\n")
	for _, k := range sortedKeys(r.Commands) {
		b.WriteString(r.Commands[k].GenerateCWrapperImpl())
	}
	return b.String()
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
	Name     string
	Platform string
	Requires []RequireBlock
}

type RequireBlock struct {
	Commands []string
	Types    []string
	Enums    []string
}