package main

import (
	"fmt"
	"strings"
)

var goTypeMap = map[string]string{
	"uint8_t":  "uint8",
	"uint16_t": "uint16",
	"uint32_t": "uint32",
	"uint64_t": "uint64",

	"int8_t":  "int8",
	"int16_t": "int16",
	"int32_t": "int32",
	"int64_t": "int64",

	"size_t": "uint",

	"char": "uint8",

	"float":  "float32",
	"double": "float64",
}

type CType struct {
	BaseType  string   // underlying type, e.g., "int", "char"
	PtrDepth  int      // pointer levels, e.g., 1 for int*, 2 for int**
	IsConst   bool     // const qualifier
	ArrayDims []string // array dimensions, e.g., [10] for int[10], [5, 3] for int[5][3]
}

type CVariable struct {
	Type CType
	Name string
}

func (t *CType) Generate() string {
	var goType strings.Builder

	for _, dim := range t.ArrayDims {
		goType.WriteString(fmt.Sprintf(
			"[%s]",
			dim,
		))
	}

	var typeName string
	ptrDepth := t.PtrDepth
	if ptrDepth > 0 && t.BaseType == "char" && t.IsConst {
		typeName = "string"
		ptrDepth--
	}

	for i := 0; i < ptrDepth; i++ {
		goType.WriteRune('*')
	}

	if typeName != "string" {
		var ok bool
		typeName, ok = goTypeMap[t.BaseType]
		if !ok {
			typeName = t.BaseType
		}
	}

	goType.WriteString(typeName)

	Type := goType.String()
	switch Type {
	case "void":
		Type = ""
	case "*void":
		Type = "unsafe.Pointer"
	}

	return Type
}

// converts
// from:
//
//	const <type>char</type>* <name>pName</name>
//
// to:
//
//		CVariable{
//	 	Name: "pName",
//			CType{
//				BaseType:  "char",
//				PtrDepth:  1,
//				IsConst:   true,
//				ArrayDims: {},
//			},
//		}
func ParseCVariable(v string) CVariable {
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, "<enum>", "")
	v = strings.ReplaceAll(v, "</enum>", "")
	v = removeTags(v, "comment")

	typeText := getTextBetweenTags(v, "type")
	nameText := getTextBetweenTags(v, "name")

	if typeText == "PFN_vkVoidFunction" {
		typeText = "void"
	}

	ptrDepth := strings.Count(v, "*")
	isConst := strings.Contains(v, "const")

	arrayDims := make([]string, 0)
	idx := 0
	for {
		start := strings.Index(v[idx:], "[")
		if start == -1 {
			break
		}
		start += idx

		end := strings.Index(v[start:], "]")
		if end == -1 {
			break
		}
		end += start

		arrayDim := upperToPascal(v[start+1 : end])
		arrayDims = append(arrayDims, arrayDim)

		idx = end + 1
	}

	return CVariable{
		Name: nameText,
		Type: CType{
			BaseType:  typeText,
			PtrDepth:  ptrDepth,
			IsConst:   isConst,
			ArrayDims: arrayDims,
		},
	}
}

func removeTags(input, tag string) string {
	openTag := "<" + tag + ">"
	closeTag := "</" + tag + ">"

	text := getTextBetweenTags(input, tag)

	ta := openTag + text + closeTag

	return strings.ReplaceAll(input, ta, "")
}

func getTextBetweenTags(input, tag string) string {
	openTag := "<" + tag + ">"
	closeTag := "</" + tag + ">"

	start := strings.Index(input, openTag)
	if start == -1 {
		return ""
	}
	start += len(openTag)

	end := strings.Index(input, closeTag)
	if end == -1 {
		return ""
	}

	return input[start:end]
}
