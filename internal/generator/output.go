package generator

import (
	"fmt"
	"sort"
	"strings"
)

// GeneratePackage generates the main Go source file for the package.
func (r *ReducedRegistry) GeneratePackage(pkg string) string {
	var b strings.Builder

	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(cgoHeader)
	b.WriteString(runtimeHelpers)

	// API constants — grouped by type to avoid "mixed untyped constant" errors
	if len(r.APIConstants) > 0 {
		b.WriteString(generateAPIConstants(r.APIConstants))
	}

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
		b.WriteString(s.GenerateStructure())
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
		b.WriteString(fmt.Sprintf("//go:build %s\n\n", tag))
		b.WriteString("package " + pkg + "\n\n")
		b.WriteString(cgoPlatformHeader)

		// Handles
		for _, h := range pc.handles {
			b.WriteString(h.Generate())
		}

		// Structs
		for _, s := range pc.structs {
			b.WriteString(s.GenerateStructure())
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

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
