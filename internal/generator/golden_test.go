package generator

import (
	"flag"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// The generator's output conventions are load-bearing: every one of them is a
// promise to consumers of package vk, and a change to any of them is a silent
// API break across tens of thousands of generated lines. These tests pin them
// down on hand-written registry fragments small enough to read, so a diff in
// the golden file says exactly which convention moved.
//
// Run with -update to rewrite the golden files after an intended change, then
// read the diff before committing:
//
//	go test ./internal/generator -update

var update = flag.Bool("update", false, "rewrite the .golden files from the current generator output")

// goldenCase is one registry fragment plus the generator entry point to render.
type goldenCase struct {
	// name is both the subtest name and the testdata basename: fixture
	// testdata/<name>.xml is compared against testdata/<name>.golden.
	name string

	// why records the conventions the fixture exists to pin down.
	why string

	config *Config

	// render produces the text to compare. Defaults to renderPackageBody.
	render func(*testing.T, *ReducedRegistry) string
}

func goldenCases() []goldenCase {
	core := func() *Config {
		return &Config{API: "vulkan", Version: "1.0"}
	}

	return []goldenCase{
		{
			name: "handles_and_commands",
			why: "commands whose first parameter is a handle become methods on it; " +
				"output-pointer parameters become return values; VkResult becomes error; " +
				"count+array out-parameter pairs collapse into a returned slice",
			config: core(),
		},
		{
			name: "structs_and_slices",
			why: "pNext becomes Next Structure and sType is set from GetType; " +
				"count+pFoo member pairs collapse into one slice field; " +
				"char* becomes string and char* const* becomes []string",
			config: core(),
		},
		{
			name: "enums_and_bitmasks",
			why: "enum values, negative values and bitpos; extension enumerants " +
				"computed from extnumber/offset/dir; aliases; the Result error method",
			config: &Config{API: "vulkan", Version: "1.0", Extensions: []string{"VK_EXT_golden_test"}},
		},
		{
			name: "arrays_unions_bitfields",
			why: "fixed-size C arrays; unions as [N]byte with per-variant accessors; " +
				"C bitfield members folded into packed uint32 words; explicit padding " +
				"before union members, whose Go alignment is 1",
			config: core(),
		},
		{
			name:   "callbacks",
			why:    "PFN struct members take Go closures, dispatched through a registry keyed by pUserData",
			config: core(),
			render: renderCallbacksFile,
		},
		{
			name:   "platform_split",
			why:    "platform-tagged extensions are emitted into their own build-tagged file",
			config: &Config{API: "vulkan", Version: "1.0", Extensions: []string{"VK_KHR_golden_surface"}, Platforms: []string{"xlib"}},
			render: renderPlatformFiles,
		},
	}
}

func TestGenerateGolden(t *testing.T) {
	for _, tc := range goldenCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("pins: %s", tc.why)

			xmlPath := filepath.Join("testdata", tc.name+".xml")
			parsed, err := ParseXML(xmlPath)
			if err != nil {
				t.Fatalf("parsing %s: %v", xmlPath, err)
			}

			reduced := Filter(BuildRegistry(parsed), tc.config)

			render := tc.render
			if render == nil {
				render = renderPackageBody
			}
			got := render(t, reduced)

			compareGolden(t, filepath.Join("testdata", tc.name+".golden"), got)
		})
	}
}

// renderPackageBody renders the main generated file without the fixed preamble.
//
// The preamble (import block and runtimeHelpers) is the same for every
// registry, so leaving it out keeps the golden files to the part that actually
// depends on the fixture.
func renderPackageBody(t *testing.T, r *ReducedRegistry) string {
	t.Helper()

	full := mustFormat(t, "generated package", r.GeneratePackage("vk"))
	preamble := mustFormat(t, "fixed preamble",
		generatedHeader+"package vk\n\n"+abiHeader+runtimeHelpers)

	body, ok := strings.CutPrefix(full, preamble)
	if !ok {
		t.Fatalf("generated output no longer begins with the fixed preamble; " +
			"update renderPackageBody to match GeneratePackage")
	}
	return body
}

func renderCallbacksFile(t *testing.T, r *ReducedRegistry) string {
	t.Helper()
	return mustFormat(t, "callbacks file", r.GenerateCallbacksFile("vk"))
}

func renderPlatformFiles(t *testing.T, r *ReducedRegistry) string {
	t.Helper()

	files := r.GeneratePlatformFiles("vk")
	if len(files) == 0 {
		t.Fatal("no platform files generated; the fixture's platform extension was not enabled")
	}
	// GeneratePlatformFiles iterates a map, so order is not stable.
	sort.Slice(files, func(i, j int) bool { return files[i].BuildTag < files[j].BuildTag })

	var b strings.Builder
	for _, f := range files {
		b.WriteString(mustFormat(t, "platform file "+f.BuildTag, f.Content))
	}
	return b.String()
}

// mustFormat gofmts generated output, mirroring WriteFile. A formatting failure
// means the generator emitted text that is not valid Go, which is a defect in
// its own right rather than a golden mismatch.
func mustFormat(t *testing.T, what, src string) string {
	t.Helper()

	formatted, err := format.Source([]byte(src))
	if err != nil {
		t.Fatalf("gofmt rejected the %s: %v\n%s", what, err, numberLines(src))
	}
	return string(formatted)
}

func compareGolden(t *testing.T, path, got string) {
	t.Helper()

	if *update {
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatalf("writing %s: %v", path, err)
		}
		t.Logf("updated %s", path)
		return
	}

	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v\n(run `go test ./internal/generator -update` to create it)", path, err)
	}

	if got == string(want) {
		return
	}

	t.Errorf("generated output differs from %s\n%s\n"+
		"if this change is intended, rerun with -update and review the diff",
		path, firstDifference(string(want), got))
}

// firstDifference renders the first differing line with a little context, which
// is far easier to read than a whole-file dump.
func firstDifference(want, got string) string {
	wantLines := strings.Split(want, "\n")
	gotLines := strings.Split(got, "\n")

	for i := 0; i < len(wantLines) || i < len(gotLines); i++ {
		w, g := "", ""
		if i < len(wantLines) {
			w = wantLines[i]
		}
		if i < len(gotLines) {
			g = gotLines[i]
		}
		if w == g {
			continue
		}

		var b strings.Builder
		for j := max(0, i-3); j < i; j++ {
			b.WriteString("  " + wantLines[j] + "\n")
		}
		b.WriteString("- want line " + itoa(i+1) + ": " + w + "\n")
		b.WriteString("+ got  line " + itoa(i+1) + ": " + g + "\n")
		return b.String()
	}

	return "(files differ only in trailing content)"
}

func numberLines(s string) string {
	var b strings.Builder
	for i, line := range strings.Split(s, "\n") {
		b.WriteString(itoa(i+1) + "\t" + line + "\n")
	}
	return b.String()
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
