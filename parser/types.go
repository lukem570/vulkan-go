package parser

import (
	"regexp"
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

var fixedArrayRe = regexp.MustCompile(`\[(\d+)\]`)

type XMLTypes struct {
	Types []XMLType `xml:"type"`
}

type XMLMember struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	Len      string `xml:"len,attr"`
	AltLen   string `xml:"altlen,attr"`
	Optional string `xml:"optional,attr"`

	InnerXML string `xml:",innerxml"`
}

type XMLType struct {
	Category string `xml:"category,attr"`
	Name     string `xml:"name,attr"`
	Alias    string `xml:"alias,attr"`
	Requires string `xml:"requires,attr"`
	Parent   string `xml:"parent,attr"`

	Members []XMLMember `xml:"member"`

	// For handle/basetype types the name may be in a child <name> element
	NameElem string `xml:"name"`

	// For funcpointer types
	Proto      XMLProto   `xml:"proto"`
	FuncParams []XMLParam `xml:"param"`
}

// isPointerMember checks whether the raw XML contains a * for the member type.
func isPointerMember(m XMLMember) bool {
	return strings.Contains(m.InnerXML, "*")
}

// fixedArraySize returns the fixed array size from InnerXML (e.g. [2] after </name>), or 0 if not a fixed array.
func fixedArraySize(m XMLMember) int {
	idx := strings.Index(m.InnerXML, "</name>")
	if idx < 0 {
		return 0
	}
	after := m.InnerXML[idx+len("</name>"):]
	match := fixedArrayRe.FindStringSubmatch(after)
	if match == nil {
		return 0
	}
	n := 0
	for _, ch := range match[1] {
		n = n*10 + int(ch-'0')
	}
	return n
}

func parseTypes(x *XMLRegistry, r *generator.Registry) {
	// First pass: collect all handle names so we can reference them in structs/params
	for _, t := range x.Types.Types {
		if t.Category == "handle" {
			name := t.Name
			if name == "" {
				name = t.NameElem
			}
			if name == "" || t.Alias != "" {
				continue
			}
			h := &generator.GoHandle{CName: name, GoName: stripVk(name)}
			r.Handles[h.GoName] = h
		}
	}

	// Second pass (sub-pass 1): register funcpointer names so they can reference each other
	for _, t := range x.Types.Types {
		if t.Category == "funcpointer" && t.Proto.Name != "" {
			fp := &generator.GoFuncPointer{
				CTypeName:  t.Proto.Name,
				GoTypeName: stripPFNvk(t.Proto.Name),
			}
			r.FuncPointers[fp.GoTypeName] = fp
		}
	}

	// Second pass (sub-pass 2): resolve param/return types (all names now registered)
	for _, t := range x.Types.Types {
		if t.Category == "funcpointer" && t.Proto.Name != "" {
			resolveFuncPointerTypes(t, r.Handles, r.FuncPointers)
		}
	}

	// Third pass: register struct names (so fields can reference other structs)
	for _, t := range x.Types.Types {
		switch t.Category {
		case "struct", "union":
			name := t.Name
			if name == "" {
				continue
			}
			r.Structs[stripVk(name)] = &generator.Structured{
				CName:  name,
				GoName: stripVk(name),
			}

		case "bitmask":
			name := t.Name
			if name == "" {
				name = t.NameElem
			}
			if name == "" || t.Alias != "" {
				continue
			}
			b := &generator.Bitmask{CName: name, GoName: stripVk(name)}
			r.Bitmasks[b.GoName] = b
		}
	}

	// Fourth pass: fully parse struct fields (all struct names now registered)
	for _, t := range x.Types.Types {
		switch t.Category {
		case "struct", "union":
			s := parseStruct(t, r.Handles, r.FuncPointers, r.Structs)
			if s != nil {
				r.Structs[s.GoName] = s
			}
		}
	}
}

// resolveFuncPointerTypes populates the Params and Return fields of an already-registered
// GoFuncPointer by parsing the XMLType proto/param elements.
func resolveFuncPointerTypes(t XMLType, handles map[string]*generator.GoHandle, funcPointers map[string]*generator.GoFuncPointer) {
	fp, ok := funcPointers[stripPFNvk(t.Proto.Name)]
	if !ok {
		return
	}

	// Resolve return type
	protoIsPtr := strings.Contains(t.Proto.InnerXML, "*")
	if t.Proto.Type == "void" && protoIsPtr {
		fp.Return = &generator.VoidPtr{}
	} else if t.Proto.Type != "void" {
		fp.Return = resolveFieldType(t.Proto.Type, protoIsPtr, false, handles, funcPointers, nil)
	}
	// else: void return, fp.Return stays nil

	// Resolve param types
	for _, p := range t.FuncParams {
		isPtr := strings.Contains(p.InnerXML, "*")
		ft := resolveFieldType(p.Type, isPtr, false, handles, funcPointers, nil)
		fp.Params = append(fp.Params, generator.FuncPointerParam{
			Name: goParamName(p.Name),
			Type: ft,
		})
	}
}

// parseStruct converts an XMLType of category struct/union into a Structured.
func parseStruct(t XMLType, handles map[string]*generator.GoHandle, funcPointers map[string]*generator.GoFuncPointer, structs map[string]*generator.Structured) *generator.Structured {
	name := t.Name
	if name == "" {
		return nil
	}

	s := &generator.Structured{
		CName:   name,
		GoName:  stripVk(name),
		Fields:  []generator.StructField{},
		IsUnion: t.Category == "union",
	}

	s.HasSType = hasSTypeAndPNext(t)
	if s.HasSType {
		s.SType = "StructureType" + stripVk(name)
	}

	// Build a map of countField → fieldItCounts so we can mark them
	// e.g. descriptorSetCount → pDescriptorSets
	countMap := buildCountMap(t.Members)
	seen := map[string]bool{}

	for _, m := range t.Members {
		if m.Name == "sType" || m.Name == "pNext" {
			continue
		}
		// Skip duplicate members (e.g. same field with different api= attributes)
		if seen[m.Name] {
			continue
		}
		seen[m.Name] = true

		// Determine if this is a count-only field that will be folded into a slice
		if sliceField, ok := countMap[m.Name]; ok {
			// This is a count field — record it but mark it as CountFor
			field := generator.StructField{
				GoName:    toPublic(stripP(m.Name)),
				CName:     m.Name,
				CountFor:  toPublic(stripP(sliceField)),
				CountCName: m.Name,
				Type:      generator.NewPrimitive("uint32_t", "uint32"),
			}
			s.Fields = append(s.Fields, field)
			continue
		}

		isPtr := isPointerMember(m)
		isArr := m.Len != "" && m.Len != "null-terminated"
		fixedSize := fixedArraySize(m)

		// If there's a len attribute that references another member, it's a slice
		ft := resolveFieldType(m.Type, isPtr, isArr, handles, funcPointers, structs)

		// Wrap in FixedArray if this is a fixed-size array (e.g. [2])
		if fixedSize > 0 {
			ft = &generator.FixedArray{Child: ft, Size: fixedSize}
		}

		// Find the C name of the count field for this slice
		countCName := ""
		if isArr {
			countCName = findCountFieldName(m.Len, t.Members)
		}

		goFieldName := toPublic(stripP(m.Name))
		field := generator.StructField{
			GoName:     goFieldName,
			CName:      m.Name,
			CountCName: countCName,
			Type:       ft,
		}
		s.Fields = append(s.Fields, field)
	}

	return s
}

// buildCountMap returns a map from count-field-name → the array field it counts.
// For example: "descriptorSetCount" → "pDescriptorSets"
func buildCountMap(members []XMLMember) map[string]string {
	result := map[string]string{}
	for _, m := range members {
		if m.Len == "" || m.Len == "null-terminated" {
			continue
		}
		// len might be "descriptorSetCount,null-terminated" or just "descriptorSetCount"
		parts := strings.Split(m.Len, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" && part != "null-terminated" {
				result[part] = m.Name
			}
		}
	}
	return result
}

// findCountFieldName returns the C name of the count member referenced by a len attr.
func findCountFieldName(lenAttr string, members []XMLMember) string {
	parts := strings.Split(lenAttr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "null-terminated" {
			continue
		}
		for _, m := range members {
			if m.Name == part {
				return m.Name
			}
		}
	}
	return ""
}

func hasSTypeAndPNext(t XMLType) bool {
	hasSType, hasPNext := false, false
	for _, m := range t.Members {
		switch m.Name {
		case "sType":
			hasSType = true
		case "pNext":
			hasPNext = true
		}
	}
	return hasSType && hasPNext
}