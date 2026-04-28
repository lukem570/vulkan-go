package generator

import (
	"fmt"
	"sort"
	"strings"
)

// GenerateCHeader returns the content of volk_wrappers.h.
func (r *ReducedRegistry) GenerateCHeader() string {
	var b strings.Builder
	b.WriteString("#ifndef VOLK_WRAPPERS_H\n#define VOLK_WRAPPERS_H\n\n")

	// Platform detection and includes
	b.WriteString(r.generatePlatformCIncludes())

	b.WriteString("#include \"volk.h\"\n")
	b.WriteString("#include <stdlib.h>\n\n")
	b.WriteString("#include <vulkan/vulkan.h>\n\n")

	// Bitfield helpers: static inline getters/setters for structs with C bitfields.
	for _, k := range sortedKeys(r.Structs) {
		b.WriteString(r.Structs[k].GenerateCBitfieldHelpers())
	}

	b.WriteString("VkResult vulkan_platform_initialize(void);\n\n")

	// Non-platform commands
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			b.WriteString(cmd.GenerateCWrapperDecl())
		}
	}

	// Platform-specific commands grouped by protect macro
	r.writePlatformCBlocks(&b, func(cmd *GoCommand) string {
		return cmd.GenerateCWrapperDecl()
	})

	// Callback trampolines
	hasTrampolines := false
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		decls := s.GenerateCCallbackDecls()
		if decls != "" {
			if !hasTrampolines {
				b.WriteString("\n/* Callback trampolines */\n")
				hasTrampolines = true
			}
			b.WriteString(decls)
		}
	}

	b.WriteString("\n#endif\n")
	return b.String()
}

// GenerateCSource returns the content of volk_wrappers.c.
func (r *ReducedRegistry) GenerateCSource() string {
	var b strings.Builder
	b.WriteString("#define VOLK_IMPLEMENTATION\n")

	// Platform detection and includes
	b.WriteString(r.generatePlatformCIncludes())

	b.WriteString("#include \"volk.h\"\n")
	b.WriteString("#include <vulkan/vulkan.h>\n")
	b.WriteString("#include \"volk_wrappers.h\"\n\n")

	b.WriteString(platformInitImpl)

	// Non-platform commands
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			b.WriteString(cmd.GenerateCWrapperImpl())
		}
	}

	// Platform-specific commands grouped by protect macro
	r.writePlatformCBlocks(&b, func(cmd *GoCommand) string {
		return cmd.GenerateCWrapperImpl()
	})

	// Callback trampolines
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		b.WriteString(s.GenerateCCallbackImpls())
	}

	return b.String()
}

// generatePlatformCIncludes emits OS-guarded #define and #include lines for
// each platform that has commands in the registry.
func (r *ReducedRegistry) generatePlatformCIncludes() string {
	// Collect which platforms are actually used
	usedPlatforms := map[string]bool{}
	for _, cmd := range r.Commands {
		if cmd.Platform != "" {
			usedPlatforms[cmd.Platform] = true
		}
	}
	for _, s := range r.Structs {
		if s.Platform != "" {
			usedPlatforms[s.Platform] = true
		}
	}
	if len(usedPlatforms) == 0 {
		return ""
	}

	// Group platforms by build tag (OS)
	type osGroup struct {
		guard     string
		platforms []string
	}

	osGuards := map[string]string{
		"linux":   "defined(__linux__)",
		"windows": "defined(_WIN32)",
		"darwin":  "defined(__APPLE__)",
		"android": "defined(__ANDROID__)",
	}

	groups := map[string]*osGroup{}
	for platform := range usedPlatforms {
		tag := PlatformBuildTag[platform]
		if tag == "" {
			continue
		}
		if groups[tag] == nil {
			groups[tag] = &osGroup{guard: osGuards[tag]}
		}
		groups[tag].platforms = append(groups[tag].platforms, platform)
	}

	var b strings.Builder
	for _, tag := range sortedKeys(groups) {
		g := groups[tag]
		sort.Strings(g.platforms)
		b.WriteString(fmt.Sprintf("#if %s\n", g.guard))
		for _, platform := range g.platforms {
			if inc, ok := PlatformIncludes[platform]; ok {
				b.WriteString(inc + "\n")
			}
			if protect, ok := r.Platforms[platform]; ok {
				b.WriteString(fmt.Sprintf("#define %s\n", protect))
			}
		}
		b.WriteString("#endif\n\n")
	}
	return b.String()
}

// writePlatformCBlocks writes #ifdef-guarded blocks for platform-specific commands.
func (r *ReducedRegistry) writePlatformCBlocks(b *strings.Builder, generate func(*GoCommand) string) {
	// Group commands by protect macro
	byProtect := map[string][]*GoCommand{}
	for _, k := range sortedKeys(r.Commands) {
		cmd := r.Commands[k]
		if cmd.Platform == "" {
			continue
		}
		protect := r.Platforms[cmd.Platform]
		if protect == "" {
			continue
		}
		byProtect[protect] = append(byProtect[protect], cmd)
	}

	for _, protect := range sortedKeys(byProtect) {
		b.WriteString(fmt.Sprintf("\n#ifdef %s\n", protect))
		for _, cmd := range byProtect[protect] {
			b.WriteString(generate(cmd))
		}
		b.WriteString(fmt.Sprintf("#endif // %s\n", protect))
	}
}
