package main

import (
	"strings"
	"unicode"
)

func upperToPascal(c string) string {
	if c == "" {
		return ""
	}

	parts := strings.Split(strings.ToLower(c), "_")
	var b strings.Builder

	for _, p := range parts {
		if p == "" {
			continue
		}

		runes := []rune(p)
		runes[0] = unicode.ToUpper(runes[0])
		b.WriteString(string(runes))
	}

	return b.String()
}

func camelToPascal(c string) string {
	member := strings.ToUpper(string(c[0])) + c[1:]

	return member
}