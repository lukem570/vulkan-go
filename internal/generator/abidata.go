package generator

import (
	"fmt"
	"strings"
)

// ABIDataFile is a generated Go test file carrying ABI conformance data.
type ABIDataFile struct {
	// BuildTag is the Go build tag the file is constrained to, or "" for the
	// core file.
	BuildTag string
	Content  string
}

// abiImports is the import block for the core ABI data file.
const abiImports = `import (
	"reflect"
	"testing"

	"github.com/lukem570/vulkan-go/internal/abi"
)
`

// abiPlatformImports is the import block for platform ABI data files, which
// register groups but do not declare the test function.
const abiPlatformImports = `import (
	"reflect"

	"github.com/lukem570/vulkan-go/internal/abi"
)
`

// abiRunner declares the group registry and the test entry point. Groups are
// appended by init functions, including those in platform files, so the set
// under test follows the build constraints of the bindings themselves.
const abiRunner = `
// abiGroups collects the ABI probe groups registered by the generated init
// functions below, one per C translation unit the probe needs.
var abiGroups []abi.Group

// TestABIConformance checks every generated C-layout struct and enumerant
// against the Vulkan headers using a C compiler. It is skipped when no
// compiler is available or the Vulkan-Headers submodule is missing.
func TestABIConformance(t *testing.T) {
	abi.Run(t, abiGroups)
}
`

// GenerateABIDataFiles emits the ABI conformance data: the layout of every
// generated C-layout struct and the value of every generated enumerant, paired
// with the C names to compare them against.
//
// It must run after GeneratePackage, which applies SetReserved to the enums and
// so determines which enumerants are actually emitted.
func (r *ReducedRegistry) GenerateABIDataFiles(pkg string) []ABIDataFile {
	files := []ABIDataFile{{
		BuildTag: "",
		Content:  r.generateCoreABIData(pkg),
	}}

	for _, tag := range sortedKeys(r.platformsByBuildTag()) {
		if content := r.generatePlatformABIData(pkg, tag); content != "" {
			files = append(files, ABIDataFile{BuildTag: tag, Content: content})
		}
	}

	return files
}

// platformsByBuildTag groups the platforms that own at least one struct by the
// Go build tag their generated code lives under.
func (r *ReducedRegistry) platformsByBuildTag() map[string][]string {
	byTag := map[string][]string{}
	seen := map[string]bool{}

	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform == "" {
			continue
		}
		tag := PlatformBuildTag[s.Platform]
		if tag == "" || seen[s.Platform] {
			continue
		}
		seen[s.Platform] = true
		byTag[tag] = append(byTag[tag], s.Platform)
	}

	return byTag
}

func (r *ReducedRegistry) generateCoreABIData(pkg string) string {
	var b strings.Builder

	b.WriteString(generatedHeader)
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(abiImports)
	b.WriteString(abiRunner)

	b.WriteString("\nfunc init() {\n")
	b.WriteString("\tabiGroups = append(abiGroups, abi.Group{\n")
	b.WriteString("\t\tName: \"core\",\n")
	b.WriteString("\t\tStructs: []abi.StructSpec{\n")
	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue
		}
		b.WriteString(r.structSpec(s))
	}
	b.WriteString("\t\t},\n")
	b.WriteString("\t\tEnums: []abi.EnumSpec{\n")
	b.WriteString(r.enumSpecs(""))
	b.WriteString("\t\t},\n")
	b.WriteString("\t})\n")
	b.WriteString("}\n")

	return b.String()
}

func (r *ReducedRegistry) generatePlatformABIData(pkg, tag string) string {
	platforms := r.platformsByBuildTag()[tag]
	if len(platforms) == 0 {
		return ""
	}

	var b strings.Builder

	b.WriteString(generatedHeader)
	b.WriteString(fmt.Sprintf("//go:build %s\n\n", tag))
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(abiPlatformImports)
	b.WriteString("\nfunc init() {\n")

	for _, platform := range platforms {
		b.WriteString("\tabiGroups = append(abiGroups, abi.Group{\n")
		fmt.Fprintf(&b, "\t\tName: %q,\n", platform)
		// The platform's structs only become visible to the C preprocessor once
		// this macro is defined, which also pulls in its system headers.
		fmt.Fprintf(&b, "\t\tMacro: %q,\n", r.Platforms[platform])
		b.WriteString("\t\tStructs: []abi.StructSpec{\n")
		for _, k := range sortedKeys(r.Structs) {
			s := r.Structs[k]
			if s.Platform != platform {
				continue
			}
			b.WriteString(r.structSpec(s))
		}
		b.WriteString("\t\t},\n")
		b.WriteString("\t})\n")
	}

	b.WriteString("}\n")
	return b.String()
}

// structSpec emits one abi.StructSpec literal.
//
// The C field names are taken from cLayoutFields, the same helper that emits
// the struct, so the list is parallel to the Go fields by construction.
func (r *ReducedRegistry) structSpec(s *Structured) string {
	if s.IsUnion {
		// Unions are [N]byte in Go: there are no fields to line up, and the
		// size is the whole of the check.
		return fmt.Sprintf("\t\t\t{CName: %q, GoType: reflect.TypeOf(c%s{}), IsUnion: true},\n",
			s.CName, s.CName)
	}

	cNames := make([]string, 0, len(s.Fields)+2)
	for _, f := range s.cLayoutFields(r.Structs) {
		cNames = append(cNames, fmt.Sprintf("%q", f.CName))
	}

	return fmt.Sprintf("\t\t\t{CName: %q, GoType: reflect.TypeOf(c%s{}), CFields: []string{%s}},\n",
		s.CName, s.CName, strings.Join(cNames, ", "))
}

// enumSpecs emits an abi.EnumSpec literal for every enumerant the package
// declares a constant for.
func (r *ReducedRegistry) enumSpecs(platform string) string {
	var b strings.Builder

	for _, k := range sortedKeys(r.Enums) {
		for _, el := range r.Enums[k].generatedElements(platform) {
			fmt.Fprintf(&b, "\t\t\t{CName: %q, GoValue: abi.V(%s)},\n", el.CName, el.GoName)
		}
	}

	return b.String()
}
