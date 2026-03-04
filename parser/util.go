package parser

import (
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

func stripVk(name string) string {
	return strings.TrimPrefix(name, "Vk")
}

func toPublic(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[:1]) + name[1:]
}

func parseFieldType(name string) generator.FieldType {
	return &generator.NamedType{
		Name:      stripVk(name),
		CTypeName: name,
	}
}
