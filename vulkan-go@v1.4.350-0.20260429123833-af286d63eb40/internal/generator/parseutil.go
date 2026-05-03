package generator

import (
	"strings"
	"unicode"
)

func stripVk(name string) string    { return strings.TrimPrefix(name, "Vk") }
func stripPFNvk(name string) string { return strings.TrimPrefix(name, "PFN_vk") }

func stripP(name string) string {
	// Strip pfn (pointer-to-function) prefix, lowercasing the first real char
	// so that toPublic/goParamName produce the right case for each context.
	// e.g. pfnUserCallback → userCallback → toPublic → UserCallback
	if strings.HasPrefix(name, "pfn") && len(name) > 3 && name[3] >= 'A' && name[3] <= 'Z' {
		return strings.ToLower(name[3:4]) + name[4:]
	}

	i := 0
	for strings.HasPrefix(name[i:], "p") {
		i++
	}

	if i == 0 || strings.ToUpper(name[i:])[0] != name[i:][0] {
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

// UpperSnakeToPascal converts a VK_UPPER_SNAKE_CASE name to PascalCase.
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
var externalTypes = map[string]*ExternalType{
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
	handles map[string]*GoHandle,
	funcPointers map[string]*GoFuncPointer,
	structs map[string]*Structured,
	bitmasks map[string]*Bitmask,
	extraOpts ...bool,
) FieldType {
	isDoublePointer := len(extraOpts) > 0 && extraOpts[0]
	if typeName == "void" && isPointer {
		if isArray {
			return &ByteSlice{}
		}
		if isDoublePointer {
			return &Pointer{Child: &VoidPtr{}}
		}
		return &VoidPtr{}
	}
	if typeName == "char" && isPointer {
		if isArray {
			return &Slice{Child: &String{}}
		}
		return &String{}
	}
	if ft := primitiveType(typeName); ft != nil {
		if isArray {
			if isDoublePointer {
				// **T (array of pointers to T): represent as unsafe.Pointer in Go.
				return &VoidPtr{}
			}
			return &Slice{Child: ft}
		}
		if isPointer {
			return &Pointer{Child: ft}
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
			return &Pointer{Child: ext}
		}
		return ext
	}
	goName := stripVk(typeName)
	if _, ok := handles[goName]; ok {
		h := &Handle{Name: goName, CTypeName: typeName}
		if isArray {
			return &Slice{Child: h}
		}
		return h
	}
	// Check if this is a known struct — use StructType which calls toC()
	if structs != nil {
		if _, ok := structs[goName]; ok {
			st := &StructType{Name: goName, CTypeName: typeName}
			if isArray {
				if isDoublePointer {
					return &PtrSlice{Child: st}
				}
				return &Slice{Child: st}
			}
			if isPointer {
				return &Pointer{Child: st}
			}
			return st
		}
	}
	// Check if it's a known 64-bit bitmask
	if bitmasks != nil {
		if bm, ok := bitmasks[goName]; ok && bm.Is64Bit {
			named := &NamedType{Name: goName, CTypeName: typeName, Is64Bit: true}
			if isArray {
				return &Slice{Child: named}
			}
			if isPointer {
				return &Pointer{Child: named}
			}
			return named
		}
	}

	// Otherwise it's an enum/bitmask — simple cast works
	named := &NamedType{Name: goName, CTypeName: typeName}
	if isArray {
		return &Slice{Child: named}
	}
	if isPointer {
		return &Pointer{Child: named}
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

func primitiveType(cName string) FieldType {
	if cName == "VkBool32" {
		return &Bool{}
	}

	goName, ok := primitiveMap[cName]
	if !ok {
		return nil
	}

	return NewPrimitive(cName, goName)
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
