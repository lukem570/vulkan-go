package generator

import "fmt"

// GoHandle represents an opaque Vulkan handle like VkInstance, VkDevice.
type GoHandle struct {
	CName  string // VkInstance
	GoName string // Instance
}

func (h *GoHandle) Generate() string {
	if h == nil {
		return ""
	}
	return fmt.Sprintf(
		"type %s struct{ handle unsafe.Pointer }\n\n",
		h.GoName,
	)
}