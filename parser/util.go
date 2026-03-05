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

func resolveFieldType(
	typeName string,
	isPointer bool,
	isArray bool,
	handles map[string]*generator.GoHandle,
	funcPointers map[string]*generator.GoFuncPointer,
	structs map[string]*generator.Structured,
) generator.FieldType {
	if typeName == "void" && isPointer {
		return &generator.VoidPtr{}
	}
	if typeName == "char" && isPointer {
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