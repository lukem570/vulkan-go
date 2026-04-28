package generator

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
