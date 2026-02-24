package main

import (
	"fmt"
	"strconv"
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

	"float": "float32",
	"double": "float64",
}

type CType struct {
	BaseType  string // underlying type, e.g., "int", "char"
	PtrDepth  int    // pointer levels, e.g., 1 for int*, 2 for int**
	IsConst   bool   // const qualifier
	ArrayDims []int  // array dimensions, e.g., [10] for int[10], [5, 3] for int[5][3]
}

type CVariable struct {
	Type CType
	Name string
}

func (t *CType) Generate() string {
	var goType strings.Builder

	for _, dim := range t.ArrayDims {
		goType.WriteString(fmt.Sprintf(
			"[%d]",
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

	return goType.String()
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

	typeText := getTextBetweenTags(v, "type")
	nameText := getTextBetweenTags(v, "name")

	if typeText == "PFN_vkVoidFunction" {
		typeText = "void"
	}

	ptrDepth := strings.Count(v, "*")
	isConst := strings.Contains(v, "const")

	// Handle arrays, e.g., "blendConstants[4][5]"
	arrayDims := []int{}
	// Look for first '[' after variable name
	if idx := strings.Index(v, nameText+"["); idx != -1 {
		arrayPart := v[idx+len(nameText):] // get "[4][5]" part
		for len(arrayPart) > 0 {
			if arrayPart[0] != '[' {
				break
			}
			endIdx := strings.Index(arrayPart, "]")
			if endIdx == -1 {
				break
			}
			dimStr := arrayPart[1:endIdx]
			if dim, err := strconv.Atoi(dimStr); err == nil {
				arrayDims = append(arrayDims, dim)
			}
			arrayPart = arrayPart[endIdx+1:]
		}
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
