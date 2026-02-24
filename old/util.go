package main

import (
	"strings"
)

// Converts VAR_NAME -> VarName
func caseUpper(c string) string {
	parts := strings.Split(strings.ToLower(c), "_")

	for i := range parts {
		parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
	}

	return strings.Join(parts, "")
}

// Converts varName -> VarName
func camelToPascal(c string) string {
	member := strings.ToUpper(string(c[0])) + c[1:]

	return member
}