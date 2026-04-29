package generator

import (
	"fmt"
	"sort"
	"strings"
)

// GeneratePackage generates the main Go source file for the package.
func (r *ReducedRegistry) GeneratePackage(pkg string) string {
	var b strings.Builder

	b.WriteString(generatedHeader)
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(goffiHeader)
	b.WriteString(runtimeHelpers)

	// API constants — grouped by type to avoid "mixed untyped constant" errors
	if len(r.APIConstants) > 0 {
		b.WriteString(generateAPIConstants(r.APIConstants))
	}

	// Assign dispatch table levels to handles before generating them.
	for _, k := range sortedKeys(r.Handles) {
		h := r.Handles[k]
		if instanceReceiverSet[h.GoName] {
			h.TableLevel = "instance"
		} else if deviceReceiverSet[h.GoName] {
			h.TableLevel = "device"
		}
	}

	// Dispatch table struct types.
	b.WriteString(r.generateDispatchTableStructs())

	// Handles (non-platform only)
	for _, k := range sortedKeys(r.Handles) {
		if r.Handles[k].Platform == "" {
			b.WriteString(r.Handles[k].Generate())
		}
	}

	for _, k := range r.EnumAliases {
		b.WriteString(k.Generate())
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
		if r.Enums[k].GoName == "Result" {
			b.WriteString(r.Enums[k].GenerateError())
		}
	}

	// Bitmasks
	for _, k := range sortedKeys(r.Bitmasks) {
		b.WriteString(r.Bitmasks[k].Generate())
	}

	// Structs (non-platform only)
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		b.WriteString(s.GenerateCLayoutStruct(r.Structs))
		b.WriteString(s.GenerateStructure(r.Structs))
		b.WriteString(s.GenerateGetType())
		b.WriteString(s.GenerateToC())
		b.WriteString(s.GenerateFromC())
	}

	// Commands (non-platform only) — deduplicate by (ReceiverType, Name)
	seenMethods := make(map[string]bool)
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform != "" {
			continue
		}
		key := cmd.ReceiverType + "." + cmd.Name
		if seenMethods[key] {
			continue
		}
		seenMethods[key] = true
		b.WriteString(cmd.GenerateWrapper(r.STypes))
	}

	// CIF var declarations (all commands) + fn vars (global commands only).
	seenDecls := make(map[string]bool)
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if seenDecls[cmd.CName] {
			continue
		}
		seenDecls[cmd.CName] = true
		if commandLevel(cmd) == "global" {
			b.WriteString(cmd.GenerateCIFDecl())
		} else {
			b.WriteString(cmd.GenerateCIFOnly())
		}
	}
	b.WriteString("\n")

	b.WriteString(r.generateInitTableFuncs())
	b.WriteString(r.generateInitFunctions())

	return b.String()
}

// GeneratePlatformFiles generates separate Go source files for each platform.
func (r *ReducedRegistry) GeneratePlatformFiles(pkg string) []PlatformFile {
	type platformContent struct {
		buildTag  string
		platforms []string
		structs   []*Structured
		commands  []*GoCommand
		handles   []*GoHandle
	}

	byTag := map[string]*platformContent{}
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[s.Platform]
		if tag == "" {
			continue
		}
		if byTag[tag] == nil {
			byTag[tag] = &platformContent{buildTag: tag}
		}
		byTag[tag].structs = append(byTag[tag].structs, s)
	}
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[cmd.Platform]
		if tag == "" {
			continue
		}
		if byTag[tag] == nil {
			byTag[tag] = &platformContent{buildTag: tag}
		}
		byTag[tag].commands = append(byTag[tag].commands, cmd)
	}
	for _, k := range sortedKeys(r.Handles) {
		h := r.Handles[k]
		if h.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[h.Platform]
		if tag == "" {
			continue
		}
		if byTag[tag] == nil {
			byTag[tag] = &platformContent{buildTag: tag}
		}
		byTag[tag].handles = append(byTag[tag].handles, h)
	}

	var files []PlatformFile
	for tag, pc := range byTag {
		if len(pc.structs) == 0 && len(pc.commands) == 0 && len(pc.handles) == 0 {
			continue
		}

		var b strings.Builder
		b.WriteString(generatedHeader)
		b.WriteString(fmt.Sprintf("//go:build %s\n\n", tag))
		b.WriteString("package " + pkg + "\n\n")
		b.WriteString(goffiPlatformHeader)

		// Handles
		for _, h := range pc.handles {
			b.WriteString(h.Generate())
		}

		// Structs
		for _, s := range pc.structs {
			b.WriteString(s.GenerateCLayoutStruct(r.Structs))
			b.WriteString(s.GenerateStructure(r.Structs))
			b.WriteString(s.GenerateGetType())
			b.WriteString(s.GenerateToC())
			b.WriteString(s.GenerateFromC())
		}

		// Commands
		seenMethods := make(map[string]bool)
		for _, cmd := range pc.commands {
			key := cmd.ReceiverType + "." + cmd.Name
			if seenMethods[key] {
				continue
			}
			seenMethods[key] = true
			b.WriteString(cmd.GenerateWrapper(r.STypes))
		}

		files = append(files, PlatformFile{
			BuildTag: tag,
			Content:  b.String(),
		})
	}

	return files
}

// generateAPIConstants emits typed var declarations for each API constant.
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

// instanceReceiverSet is the set of Vulkan handle types whose commands are
// instance-level (loaded via vkGetInstanceProcAddr with a real instance).
var instanceReceiverSet = map[string]bool{
	"Instance": true, "PhysicalDevice": true,
}

// deviceReceiverSet is the set of Vulkan handle types whose commands are
// device-level (loaded via vkGetDeviceProcAddr).
var deviceReceiverSet = map[string]bool{
	"Device": true, "Queue": true, "CommandBuffer": true,
	"Fence": true, "Semaphore": true, "Event": true, "QueryPool": true,
	"Buffer": true, "BufferView": true, "Image": true, "ImageView": true,
	"ShaderModule": true, "PipelineCache": true, "PipelineLayout": true,
	"RenderPass": true, "Pipeline": true, "DescriptorSetLayout": true,
	"Sampler": true, "DescriptorPool": true, "DescriptorSet": true,
	"Framebuffer": true, "CommandPool": true,
	// Extensions
	"SwapchainKHR": true, "DescriptorUpdateTemplate": true,
	"PrivateDataSlot": true, "AccelerationStructureKHR": true,
	"AccelerationStructureNV": true, "MicromapEXT": true,
	"DeferredOperationKHR": true, "VideoSessionKHR": true,
	"VideoSessionParametersKHR": true, "CuModuleNVX": true,
	"CuFunctionNVX": true, "CudaModuleNV": true, "CudaFunctionNV": true,
	"IndirectCommandsLayoutNV": true, "IndirectCommandsLayoutEXT": true,
	"IndirectExecutionSetEXT": true, "ShaderEXT": true,
	"OpticalFlowSessionNV": true, "PerformanceConfigurationINTEL": true,
	"BufferCollectionFUCHSIA": true, "ValidationCacheEXT": true,
}

// commandLevel returns "global", "instance", or "device".
func commandLevel(cmd *GoCommand) string {
	if deviceReceiverSet[cmd.ReceiverType] {
		return "device"
	}
	if cmd.ReceiverType != "" {
		return "instance"
	}
	return "global"
}

// generateDispatchTableStructs emits the instanceTable and deviceTable struct types.
func (r *ReducedRegistry) generateDispatchTableStructs() string {
	var b strings.Builder

	seen := make(map[string]bool)
	b.WriteString("type instanceTable struct {\n")
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if seen[cmd.CName] || commandLevel(cmd) != "instance" {
			continue
		}
		seen[cmd.CName] = true
		b.WriteString(fmt.Sprintf("\t%s unsafe.Pointer\n", cmd.CName))
	}
	b.WriteString("}\n\n")

	seen = make(map[string]bool)
	b.WriteString("type deviceTable struct {\n")
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if seen[cmd.CName] || commandLevel(cmd) != "device" {
			continue
		}
		seen[cmd.CName] = true
		b.WriteString(fmt.Sprintf("\t%s unsafe.Pointer\n", cmd.CName))
	}
	b.WriteString("}\n\n")

	return b.String()
}

// generateInitTableFuncs emits _initInstanceTable and _initDeviceTable helper functions.
func (r *ReducedRegistry) generateInitTableFuncs() string {
	var b strings.Builder

	seen := make(map[string]bool)
	var instanceCmds, deviceCmds []*GoCommand
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if seen[cmd.CName] {
			continue
		}
		seen[cmd.CName] = true
		switch commandLevel(cmd) {
		case "instance":
			instanceCmds = append(instanceCmds, cmd)
		case "device":
			deviceCmds = append(deviceCmds, cmd)
		}
	}

	// _initInstanceTable populates all instance-level fn pointers and also
	// caches vkGetDeviceProcAddr in the package-level _vkGetDeviceProcAddr
	// so that _initDeviceTable can use it without needing the instance handle.
	b.WriteString("func _initInstanceTable(handle unsafe.Pointer) *instanceTable {\n")
	b.WriteString("\t_vkGetDeviceProcAddr = _loadProc(handle, \"vkGetDeviceProcAddr\")\n")
	b.WriteString("\tt := &instanceTable{}\n")
	for _, cmd := range instanceCmds {
		b.WriteString(fmt.Sprintf("\tt.%s = _loadProc(handle, %q)\n", cmd.CName, cmd.CName))
	}
	b.WriteString("\treturn t\n}\n\n")

	b.WriteString("func _initDeviceTable(handle unsafe.Pointer) *deviceTable {\n")
	b.WriteString("\tt := &deviceTable{}\n")
	for _, cmd := range deviceCmds {
		b.WriteString(fmt.Sprintf("\tt.%s = _loadDeviceProc(handle, %q)\n", cmd.CName, cmd.CName))
	}
	b.WriteString("\treturn t\n}\n\n")

	return b.String()
}

// generateInitFunctions emits Initialize() which bootstraps the Vulkan library
// and global-level function pointers. Instance/device function pointers are now
// loaded automatically by CreateInstance and CreateDevice via dispatch tables.
func (r *ReducedRegistry) generateInitFunctions() string {
	var b strings.Builder

	seen := make(map[string]bool)
	var globalCmds, allCmds []*GoCommand
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if seen[cmd.CName] {
			continue
		}
		seen[cmd.CName] = true
		allCmds = append(allCmds, cmd)
		if commandLevel(cmd) == "global" {
			globalCmds = append(globalCmds, cmd)
		}
	}

	b.WriteString("// Initialize loads the Vulkan library and global function pointers.\n")
	b.WriteString("// Must be called before any other Vulkan function.\n")
	b.WriteString("// Instance and device function pointers are loaded automatically\n")
	b.WriteString("// by CreateInstance and CreateDevice respectively.\n")
	b.WriteString("func Initialize() error {\n")
	b.WriteString("\tlib, err := ffi.LoadLibrary(_vulkanLibName())\n")
	b.WriteString("\tif err != nil {\n\t\treturn fmt.Errorf(\"vulkan: %w\", err)\n\t}\n")
	b.WriteString("\t_vkLib = lib\n\n")
	b.WriteString("\tgipaFn, err := ffi.GetSymbol(lib, \"vkGetInstanceProcAddr\")\n")
	b.WriteString("\tif err != nil {\n\t\treturn fmt.Errorf(\"vulkan: %w\", err)\n\t}\n")
	b.WriteString("\t_vkGetInstanceProcAddr = gipaFn\n")
	b.WriteString("\tffi.PrepareCallInterface(&_cif_vkGetInstanceProcAddr, types.DefaultCall, types.PointerTypeDescriptor, []*types.TypeDescriptor{types.PointerTypeDescriptor, types.PointerTypeDescriptor}) //nolint\n")
	b.WriteString("\tffi.PrepareCallInterface(&_cif_vkGetDeviceProcAddr, types.DefaultCall, types.PointerTypeDescriptor, []*types.TypeDescriptor{types.PointerTypeDescriptor, types.PointerTypeDescriptor}) //nolint\n\n")

	b.WriteString("\t// Load global commands\n")
	for _, cmd := range globalCmds {
		b.WriteString(fmt.Sprintf("\t_fn_%s = _loadProc(nil, %q)\n", cmd.CName, cmd.CName))
	}

	b.WriteString("\n\t// Prepare all call interfaces\n")
	for _, cmd := range allCmds {
		b.WriteString(cmd.GenerateCIFPrepare())
	}

	b.WriteString("\n\treturn nil\n}\n\n")
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
