package generator

import "fmt"

// GoHandle represents an opaque Vulkan handle like VkInstance, VkDevice.
type GoHandle struct {
	CName      string // VkInstance
	GoName     string // Instance
	Platform   string // non-empty for platform-specific handles
	TableLevel string // "instance", "device", or "" (no dispatch table)
}

func (h *GoHandle) Generate() string {
	if h == nil {
		return ""
	}
	var tableField string
	switch h.TableLevel {
	case "instance":
		tableField = "\ttable   *instanceTable\n"
	case "device":
		tableField = "\ttable   *deviceTable\n"
	}
	return fmt.Sprintf(
		"type %s struct {\n\thandle  unsafe.Pointer\n%s\tcleanup func()\n}\n\n"+
			"func (h *%s) Handle() unsafe.Pointer { return h.handle }\n\n",
		h.GoName, tableField, h.GoName,
	)
}