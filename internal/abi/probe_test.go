package abi

import (
	"strings"
	"testing"
)

// The probe decides which enumerants have no C counterpart by reading the
// compiler's diagnostics. If that parsing silently stops matching, every such
// enumerant turns into an unreadable hard failure instead of a precise report,
// so both compiler dialects are pinned here.
func TestUndeclaredIdentifiers(t *testing.T) {
	tests := []struct {
		name        string
		diagnostics string
		want        []string
	}{
		{
			name: "gcc ascii quotes",
			diagnostics: `probe.c: In function 'main':
probe.c:12:5: error: 'VK_GOLDEN_RESERVED_BIT' undeclared (first use in this function)`,
			want: []string{"VK_GOLDEN_RESERVED_BIT"},
		},
		{
			name: "gcc typographic quotes",
			diagnostics: "probe.c:12:5: error: ‘VK_GOLDEN_RESERVED_BIT’ undeclared " +
				"(first use in this function)",
			want: []string{"VK_GOLDEN_RESERVED_BIT"},
		},
		{
			name: "gcc did-you-mean suggestion is not mistaken for the missing name",
			diagnostics: `probe.c:12:5: error: 'VK_ACCESS_2_RESERVED_53_BIT_KHR' undeclared ` +
				`(first use in this function); did you mean 'VK_ACCESS_2_MEMORY_READ_BIT_KHR'?`,
			want: []string{"VK_ACCESS_2_RESERVED_53_BIT_KHR"},
		},
		{
			name:        "clang wording",
			diagnostics: `probe.c:12:5: error: use of undeclared identifier 'VK_GOLDEN_RESERVED_BIT'`,
			want:        []string{"VK_GOLDEN_RESERVED_BIT"},
		},
		{
			name: "deduplicated and in first-seen order",
			diagnostics: `probe.c:12:5: error: 'VK_SECOND' undeclared (first use in this function)
probe.c:13:5: error: 'VK_FIRST' undeclared (first use in this function)
probe.c:14:5: error: 'VK_SECOND' undeclared (first use in this function)`,
			want: []string{"VK_SECOND", "VK_FIRST"},
		},
		{
			name:        "unrelated diagnostics yield nothing",
			diagnostics: `probe.c:12:5: error: unknown type name 'VkGoldenMissing'`,
			want:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := undeclaredIdentifiers(tc.diagnostics)
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("got %v, want %v", got, tc.want)
				}
			}
		})
	}
}

func TestParseReport(t *testing.T) {
	layout, err := parseReport(strings.Join([]string{
		"s|VkGoldenCreateInfo|24|8",
		"f|VkGoldenCreateInfo|sType|0|4",
		"f|VkGoldenCreateInfo|pNext|8|8",
		"e|VK_SUCCESS|0",
		"e|VK_ERROR_OUT_OF_HOST_MEMORY|18446744073709551615",
		"",
	}, "\n"))
	if err != nil {
		t.Fatalf("parseReport: %v", err)
	}

	s, ok := layout.structs["VkGoldenCreateInfo"]
	if !ok {
		t.Fatal("struct record missing")
	}
	if s.size != 24 || s.align != 8 {
		t.Errorf("size/align = %d/%d, want 24/8", s.size, s.align)
	}
	if f := s.fields["pNext"]; f.offset != 8 || f.size != 8 {
		t.Errorf("pNext offset/size = %d/%d, want 8/8", f.offset, f.size)
	}

	// Negative enumerants arrive as the sign-extended bit pattern, matching V.
	if got := layout.enums["VK_ERROR_OUT_OF_HOST_MEMORY"]; got != ^uint64(0) {
		t.Errorf("VK_ERROR_OUT_OF_HOST_MEMORY = %#x, want %#x", got, ^uint64(0))
	}
}

func TestParseReportRejectsMalformedRecords(t *testing.T) {
	tests := map[string]string{
		"short struct record": "s|VkGolden|24",
		"field before struct": "f|VkGolden|sType|0|4",
		"non-numeric size":    "s|VkGolden|big|8",
		"unknown record kind": "x|VkGolden",
	}

	for name, report := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := parseReport(report); err == nil {
				t.Fatalf("parseReport(%q) succeeded, want an error", report)
			}
		})
	}
}

// V must widen by sign extension so that a Go constant and the C value it
// mirrors compare equal for VkResult codes and for bit 63 of a VkFlags64.
func TestVSignExtends(t *testing.T) {
	type result int32
	type flags64 uint64

	if got := V(result(-1)); got != ^uint64(0) {
		t.Errorf("V(result(-1)) = %#x, want %#x", got, ^uint64(0))
	}
	if got := V(flags64(1) << 63); got != 1<<63 {
		t.Errorf("V(1<<63) = %#x, want %#x", got, uint64(1)<<63)
	}
}
