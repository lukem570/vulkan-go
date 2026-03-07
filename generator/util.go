package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func versionLE(featureName string, target string) bool {
	var major, minor int

	// Handle both VK_VERSION_1_2 and VK_GRAPHICS_VERSION_1_2
	_, err := fmt.Sscanf(featureName, "VK_VERSION_%d_%d", &major, &minor)
	if err != nil {
		_, err = fmt.Sscanf(featureName, "VK_GRAPHICS_VERSION_%d_%d", &major, &minor)
		if err != nil {
			return false
		}
	}

	var targetMajor, targetMinor int
	_, err = fmt.Sscanf(target, "%d.%d", &targetMajor, &targetMinor)
	if err != nil {
		return false
	}

	if major < targetMajor {
		return true
	}
	if major == targetMajor && minor <= targetMinor {
		return true
	}
	return false
}

func WriteFile(path string, contents string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	formatted, err := format.Source([]byte(contents))
	if err != nil {
		_ = os.WriteFile(path+".unformatted", []byte(contents), 0644)
		return fmt.Errorf("gofmt failed: %w\nraw output written to %s.unformatted", err, path)
	}

	return os.WriteFile(path, formatted, 0644)
}

func WriteCFile(path string, contents string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(contents), 0644)
}

// lowerFirstChar lowercases the first character of s.
func lowerFirstChar(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// sanitizeCField escapes C field names that are Go keywords.
func sanitizeCField(name string) string {
	switch name {
	case "type":
		return "_type"
	case "range":
		return "_range"
	default:
		return name
	}
}

// sanitizeIdent strips a leading 'p' pointer prefix and lowercases the result,
// turning e.g. "pCreateInfo" into "createInfo".
func sanitizeIdent(s string) string {
	s = strings.TrimPrefix(s, "p")
	if s == "" {
		return "val"
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// zeroValue returns the Go zero-value expression for a FieldType.
func zeroValue(t FieldType) string {
	switch t.(type) {
	case *Pointer, *Handle, *VoidPtr:
		return "nil"
	case *StructType:
		return t.GoName() + "{}"
	case *Bool:
		return "false"
	case *String:
		return `""`
	default:
		goName := t.GoName()
		if len(goName) > 0 && goName[0] >= 'A' && goName[0] <= 'Z' {
			return goName + "{}"
		}
		return "0"
	}
}