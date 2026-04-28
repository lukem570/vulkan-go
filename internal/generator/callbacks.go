package generator

import "strings"

// GenerateCallbacksFile generates pkg/raw/callbacks.go containing the
// per-struct callback holder types, sync.Map registries, and //export bridge
// functions needed by the C trampolines.
func (r *ReducedRegistry) GenerateCallbacksFile(pkg string) string {
	var b strings.Builder
	b.WriteString("package " + pkg + "\n\n")
	b.WriteString(callbacksFileHeader)

	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		if s.Platform != "" {
			continue // only non-platform structs for now
		}
		b.WriteString(s.GenerateCallbacksSupport())
	}
	return b.String()
}
