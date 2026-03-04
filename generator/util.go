package generator

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
)

func versionLE(featureName string, target string) bool {
	var major, minor int

	// featureName = "VK_VERSION_1_2"
	_, err := fmt.Sscanf(featureName, "VK_VERSION_%d_%d", &major, &minor)
	if err != nil {
		return false
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
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	// Format Go code
	formatted, err := format.Source([]byte(contents))
	if err != nil {
		// If formatting fails, still write raw output
		return os.WriteFile(path, []byte(contents), 0644)
	}

	return os.WriteFile(path, formatted, 0644)
}
