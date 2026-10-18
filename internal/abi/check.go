package abi

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// maxReported caps the mismatches printed per group. A single wrong field
// offset shifts every field after it, so an unbounded report turns one bug into
// thousands of lines.
const maxReported = 40

// Run checks every group against the C ABI.
//
// Each group becomes a subtest. Groups carrying a VK_USE_PLATFORM_* macro are
// skipped when their probe does not compile, since that indicates the
// platform's development headers are absent rather than a binding defect.
func Run(t *testing.T, groups []Group) {
	t.Helper()

	cc := findCC()
	if cc == "" {
		t.Skip("no C compiler (cc, gcc or clang) on PATH: cannot read the Vulkan ABI")
	}

	root, err := repoRoot()
	if err != nil {
		t.Fatalf("locating module root: %v", err)
	}

	includeDir := filepath.Join(root, "mod", "Vulkan-Headers", "include")
	if _, err := os.Stat(filepath.Join(includeDir, "vulkan", "vulkan_core.h")); err != nil {
		t.Skipf("Vulkan-Headers submodule not initialised (%s): run `task setup`", includeDir)
	}

	tmpDir := t.TempDir()

	for _, g := range groups {
		t.Run(g.Name, func(t *testing.T) {
			layout, absent, err := probe(cc, includeDir, tmpDir, g)
			if err != nil {
				if g.Macro != "" {
					t.Skipf("platform group %q did not build; %s development headers are probably missing: %v",
						g.Name, g.Macro, err)
				}
				// The probe is generated from the same specs as the bindings, so a
				// compile failure here means the generator emitted a struct or
				// member that does not exist in the Vulkan headers.
				t.Fatalf("probe for group %q did not build: %v", g.Name, err)
			}

			// Independent subtests: a defect in one dimension leaves the others
			// usable as a gate.
			t.Run("structs", func(t *testing.T) { checkStructs(t, g, layout) })
			t.Run("enums", func(t *testing.T) { checkEnums(t, g, layout) })
			t.Run("enums_declared_but_absent_from_headers", func(t *testing.T) { checkAbsent(t, absent) })
		})
	}
}

// checkAbsent reports enumerants the bindings declare that the Vulkan headers
// do not define at all.
//
// These are not ABI mismatches — there is no C value to disagree with — but
// each one is a Go constant for something that is not in the API the bindings
// target. Known sources are reserved bit placeholders in vk.xml, entries
// belonging to the `vulkansc` API variant, and values for platforms that
// config.yml does not enable.
func checkAbsent(t *testing.T, absent []string) {
	t.Helper()

	if len(absent) == 0 {
		return
	}

	sort.Strings(absent)
	shown := absent
	if len(shown) > maxReported {
		shown = shown[:maxReported]
	}

	t.Errorf("%d enumerants are declared in the bindings but absent from the Vulkan headers:\n\t%s",
		len(absent), strings.Join(shown, "\n\t"))
	if len(absent) > len(shown) {
		t.Errorf("and %d more", len(absent)-len(shown))
	}
}

func checkStructs(t *testing.T, g Group, layout *cLayout) {
	t.Helper()

	reported := 0
	suppressed := 0
	report := func(format string, args ...any) {
		if reported >= maxReported {
			suppressed++
			return
		}
		reported++
		t.Errorf(format, args...)
	}

	for _, spec := range g.Structs {
		got, ok := layout.structs[spec.CName]
		if !ok {
			report("%s: C compiler reported no layout for this struct", spec.CName)
			continue
		}

		if uint64(spec.GoType.Size()) != got.size {
			report("%s: Go size %d, C size %d", spec.CName, spec.GoType.Size(), got.size)
		}

		// Unions are [N]byte in Go, so Go infers alignment 1 by design and the
		// generator pads at each use site instead.
		if !spec.IsUnion && uint64(spec.GoType.Align()) != got.align {
			report("%s: Go alignment %d, C alignment %d (a pointer to this struct may be under-aligned for C)",
				spec.CName, spec.GoType.Align(), got.align)
		}

		if spec.IsUnion {
			continue // all union members sit at offset 0; size is the whole check
		}

		if spec.GoType.NumField() != len(spec.CFields) {
			report("%s: %d Go fields but %d C field names; the generator's spec and struct emission disagree",
				spec.CName, spec.GoType.NumField(), len(spec.CFields))
			continue
		}

		for i, cName := range spec.CFields {
			if cName == "" {
				continue // synthetic: padding or a bitfield pack
			}
			goField := spec.GoType.Field(i)
			cf, ok := got.fields[cName]
			if !ok {
				report("%s.%s: C compiler reported no offset for this member", spec.CName, cName)
				continue
			}
			if uint64(goField.Offset) != cf.offset {
				report("%s.%s: Go offset %d, C offset %d", spec.CName, cName, goField.Offset, cf.offset)
			}
			if uint64(goField.Type.Size()) != cf.size {
				report("%s.%s: Go field size %d (%s), C field size %d",
					spec.CName, cName, goField.Type.Size(), goField.Type, cf.size)
			}
		}
	}

	if suppressed > 0 {
		t.Errorf("%d further struct mismatches suppressed (one wrong offset shifts every later field)", suppressed)
	}
}

func checkEnums(t *testing.T, g Group, layout *cLayout) {
	t.Helper()

	reported := 0
	suppressed := 0

	for _, spec := range g.Enums {
		got, ok := layout.enums[spec.CName]
		if !ok {
			continue // the probe would not have compiled had the name been absent
		}
		if got == spec.GoValue {
			continue
		}
		if reported >= maxReported {
			suppressed++
			continue
		}
		reported++
		// Print signed too: VkResult codes and high bitmask bits read as huge
		// unsigned numbers otherwise.
		t.Errorf("%s: Go value %d (%#x), C value %d (%#x)",
			spec.CName, int64(spec.GoValue), spec.GoValue, int64(got), got)
	}

	if suppressed > 0 {
		t.Errorf("%d further enum mismatches suppressed", suppressed)
	}
}
