package parser

import (
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

func stripVk(name string) string    { return strings.TrimPrefix(name, "Vk") }
func stripPFNvk(name string) string { return strings.TrimPrefix(name, "PFN_vk") }

func stripP(name string) string {
	if strings.HasPrefix(name, "pp") && len(name) > 2 {
		if name[2] >= 'A' && name[2] <= 'Z' {
			return lowerFirst(name[2:])
		}
	}
	if strings.HasPrefix(name, "p") && len(name) > 1 {
		if name[1] >= 'A' && name[1] <= 'Z' {
			return lowerFirst(name[1:])
		}
	}
	return name
}

func toPublic(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[:1]) + name[1:]
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func UpperSnakeToPascal(vk string) string {
	vk = strings.TrimPrefix(vk, "VK_")
	parts := strings.Split(vk, "_")
	for i, p := range parts {
		if p == "" {
			continue
		}
		upper := strings.ToUpper(p)
		switch upper {
		case "KHR", "EXT", "NV", "NVX", "AMD", "IMG", "QCOM", "ARM", "VALVE",
			"MESA", "INTEL", "GOOGLE", "LUNARG", "FB", "SEC", "HUAWEI":
			parts[i] = upper
		default:
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return strings.Join(parts, "")
}

// externalTypes maps C type names from platform headers to their Go representations.
// These types are not part of the Vulkan type system but appear in platform-specific
// extensions (e.g. VK_KHR_xlib_surface uses Display* and Window from X11).
var externalTypes = map[string]*generator.ExternalType{
	// X11 / Xlib
	"Display":  {CTypeName: "Display", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	"VisualID": {CTypeName: "VisualID", GoTypeName: "uintptr"},
	"Window":   {CTypeName: "Window", GoTypeName: "uintptr"},
	// XCB
	"xcb_connection_t": {CTypeName: "xcb_connection_t", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	"xcb_window_t":     {CTypeName: "xcb_window_t", GoTypeName: "uint32"},
	"xcb_visualid_t":   {CTypeName: "xcb_visualid_t", GoTypeName: "uint32"},
	// Wayland (cgo uses struct_ prefix for C structs)
	"wl_display": {CTypeName: "struct_wl_display", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	"wl_surface":  {CTypeName: "struct_wl_surface", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	// Win32
	"HINSTANCE": {CTypeName: "HINSTANCE", GoTypeName: "unsafe.Pointer"},
	"HWND":      {CTypeName: "HWND", GoTypeName: "unsafe.Pointer"},
	"HMONITOR":  {CTypeName: "HMONITOR", GoTypeName: "unsafe.Pointer"},
	"DWORD":     {CTypeName: "DWORD", GoTypeName: "uint32"},
	"LPCWSTR":   {CTypeName: "LPCWSTR", GoTypeName: "unsafe.Pointer"},
	"HANDLE":    {CTypeName: "HANDLE", GoTypeName: "unsafe.Pointer"},
	"SECURITY_ATTRIBUTES": {CTypeName: "SECURITY_ATTRIBUTES", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	// Metal / macOS
	"CAMetalLayer": {CTypeName: "CAMetalLayer", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	// Android
	"ANativeWindow": {CTypeName: "ANativeWindow", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
}

func resolveFieldType(
	typeName string,
	isPointer bool,
	isArray bool,
	handles map[string]*generator.GoHandle,
	funcPointers map[string]*generator.GoFuncPointer,
	structs map[string]*generator.Structured,
	extraOpts ...bool,
) generator.FieldType {
	isDoublePointer := len(extraOpts) > 0 && extraOpts[0]
	if typeName == "void" && isPointer {
		if isDoublePointer {
			return &generator.Pointer{Child: &generator.VoidPtr{}}
		}
		return &generator.VoidPtr{}
	}
	if typeName == "char" && isPointer {
		if isArray {
			return &generator.Slice{Child: &generator.String{}}
		}
		return &generator.String{}
	}
	if ft := primitiveType(typeName); ft != nil {
		if isArray {
			return &generator.Slice{Child: ft}
		}
		if isPointer {
			return &generator.Pointer{Child: ft}
		}
		return ft
	}
	// Check function pointers (PFN_vk* types)
	if fp, ok := funcPointers[stripPFNvk(typeName)]; ok {
		return fp
	}
	// Check external platform types (Display, HWND, xcb_connection_t, etc.)
	if ext, ok := externalTypes[typeName]; ok {
		if isPointer && ext.PtrInVulkan {
			// The pointer is consumed by the ExternalType (e.g. Display* → unsafe.Pointer)
			return ext
		}
		if isPointer {
			return &generator.Pointer{Child: ext}
		}
		return ext
	}
	goName := stripVk(typeName)
	if _, ok := handles[goName]; ok {
		h := &generator.Handle{Name: goName, CTypeName: typeName}
		if isArray {
			return &generator.Slice{Child: h}
		}
		return h
	}
	// Check if this is a known struct — use StructType which calls toC()
	if structs != nil {
		if _, ok := structs[goName]; ok {
			st := &generator.StructType{Name: goName, CTypeName: typeName}
			if isArray {
				return &generator.Slice{Child: st}
			}
			if isPointer {
				return &generator.Pointer{Child: st}
			}
			return st
		}
	}
	// Otherwise it's an enum/bitmask — simple cast works
	named := &generator.NamedType{Name: goName, CTypeName: typeName}
	if isArray {
		return &generator.Slice{Child: named}
	}
	if isPointer {
		return &generator.Pointer{Child: named}
	}
	return named
}

func primitiveType(name string) generator.FieldType {
	switch name {
	case "uint8_t":
		return generator.NewPrimitive("uint8_t", "uint8")
	case "uint16_t":
		return generator.NewPrimitive("uint16_t", "uint16")
	case "uint32_t":
		return generator.NewPrimitive("uint32_t", "uint32")
	case "uint64_t":
		return generator.NewPrimitive("uint64_t", "uint64")
	case "int8_t":
		return generator.NewPrimitive("int8_t", "int8")
	case "int16_t":
		return generator.NewPrimitive("int16_t", "int16")
	case "int32_t":
		return generator.NewPrimitive("int32_t", "int32")
	case "int64_t":
		return generator.NewPrimitive("int64_t", "int64")
	case "float":
		return generator.NewPrimitive("float", "float32")
	case "double":
		return generator.NewPrimitive("double", "float64")
	case "size_t":
		return generator.NewPrimitive("size_t", "uintptr")
	case "int":
		return generator.NewPrimitive("int", "int32")
	case "char":
		return generator.NewPrimitive("char", "byte")
	case "VkBool32":
		return &generator.Bool{}
	case "VkDeviceSize":
		return generator.NewPrimitive("VkDeviceSize", "uint64")
	case "VkDeviceAddress":
		return generator.NewPrimitive("VkDeviceAddress", "uint64")
	case "VkSampleMask":
		return generator.NewPrimitive("VkSampleMask", "uint32")
	}
	return nil
}