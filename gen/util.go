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
func caseMember(c string) string {
	member := strings.ToUpper(string(c[0])) + c[1:]

	return member
}

func extractCType(m FunctionMember) string {
	s := m.InnerXML

	// Remove name
	s = strings.ReplaceAll(s, "<name>"+m.Name+"</name>", "")

	// Replace <type>...</type> with actual type text
	s = strings.ReplaceAll(s, "<type>"+m.Type+"</type>", m.Type)

	return strings.TrimSpace(s)
}

func convertCToGoType(cType string) string {
	cType = strings.TrimSpace(cType)

	

	// Count pointer depth
	ptrDepth := strings.Count(cType, "*")

	// Remove all *
	base := strings.ReplaceAll(cType, "*", "")
	base = strings.TrimSpace(base)

	// Map base C type → Go type
	goBase, ok := goTypeMap[base]
	if !ok {
		goBase = base
	}

	// Special case: void*
	if base == "void" && ptrDepth > 0 {
		return "unsafe.Pointer"
	}

	if base == "void" && ptrDepth == 0 {
		return ""
	}

	if base == "const char" && ptrDepth > 1 {
		goBase = "string"
		ptrDepth--
	}

	// Rebuild pointer chain
	for i := 0; i < ptrDepth; i++ {
		goBase = "*" + goBase
	}

	return goBase
}