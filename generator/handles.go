package generator

import "fmt"

// GoHandle represents an opaque Vulkan handle like VkInstance, VkDevice.
type GoHandle struct {
	CName    string // VkInstance
	GoName   string // Instance
	Platform string // non-empty for platform-specific handles
}

func (h *GoHandle) Generate() string {
	if h == nil {
		return ""
	}
	return fmt.Sprintf(
		"type %s struct{ handle unsafe.Pointer }\n\n"+
			"func (h *%s) Handle() unsafe.Pointer { return h.handle }\n\n",
		h.GoName, h.GoName,
	)
}