package parser

import (
	"strings"
	"unicode"

	"github.com/lukem570/vulkan-go/internal/generator"
)

func stripVk(name string) string    { return strings.TrimPrefix(name, "Vk") }
func stripPFNvk(name string) string { return strings.TrimPrefix(name, "PFN_vk") }

func stripP(name string) string {
	i := 0
	for strings.HasPrefix(name[i:], "p") {
		i++
	}

	if strings.ToUpper(name[i:])[0] != name[i:][0] {
		return name
	}

	return name[i:]
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
	"wl_surface": {CTypeName: "struct_wl_surface", GoTypeName: "unsafe.Pointer", PtrInVulkan: true},
	// Win32
	"HINSTANCE":           {CTypeName: "HINSTANCE", GoTypeName: "unsafe.Pointer"},
	"HWND":                {CTypeName: "HWND", GoTypeName: "unsafe.Pointer"},
	"HMONITOR":            {CTypeName: "HMONITOR", GoTypeName: "unsafe.Pointer"},
	"DWORD":               {CTypeName: "DWORD", GoTypeName: "uint32"},
	"LPCWSTR":             {CTypeName: "LPCWSTR", GoTypeName: "unsafe.Pointer"},
	"HANDLE":              {CTypeName: "HANDLE", GoTypeName: "unsafe.Pointer"},
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
		if isArray {
			return &generator.ByteSlice{}
		}
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
			if isDoublePointer {
				// **T (array of pointers to T): represent as unsafe.Pointer in Go.
				return &generator.VoidPtr{}
			}
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
				if isDoublePointer {
					return &generator.PtrSlice{Child: st}
				}
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

var primitiveMap = map[string]string{
	"int": "int32",

	"int8_t":  "int8",
	"int16_t": "int16",
	"int32_t": "int32",
	"int64_t": "int64",

	"uint8_t":  "uint8",
	"uint16_t": "uint16",
	"uint32_t": "uint32",
	"uint64_t": "uint64",

	"size_t": "uintptr",

	"float":  "float32",
	"double": "float64",

	"char": "byte",

	"VkDeviceSize":    "uint64",
	"VkDeviceAddress": "uint64",
	"VkSampleMask":    "uint32",
}

func primitiveType(cName string) generator.FieldType {
	if cName == "VkBool32" {
		return &generator.Bool{}
	}

	goName, ok := primitiveMap[cName]
	if !ok {
		return nil
	}

	return generator.NewPrimitive(cName, goName)
}

func extractDigits(s string) string {
	var digits strings.Builder

	for _, r := range s {
		if unicode.IsDigit(r) || r == '.' {
			digits.WriteRune(r)
		}
	}

	return digits.String()
}
