package parser

import (
	"regexp"
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

var fixedArrayRe = regexp.MustCompile(`\[(\d+)\]`)
var namedArrayRe = regexp.MustCompile(`\[<enum>(\w+)</enum>\]`)

// namedArrayConstants maps Vulkan constant names to their integer values.
var namedArrayConstants = map[string]int{
	"VK_MAX_PHYSICAL_DEVICE_NAME_SIZE": 256,
	"VK_UUID_SIZE":                     16,
	"VK_LUID_SIZE":                     8,
	"VK_MAX_EXTENSION_NAME_SIZE":       256,
	"VK_MAX_DESCRIPTION_SIZE":          256,
	"VK_MAX_MEMORY_TYPES":              32,
	"VK_MAX_MEMORY_HEAPS":              16,
	"VK_MAX_DEVICE_GROUP_SIZE":         32,
	"VK_MAX_DRIVER_NAME_SIZE":          256,
	"VK_MAX_DRIVER_INFO_SIZE":          256,
	"VK_MAX_GLOBAL_PRIORITY_SIZE_KHR":  16,
	"VK_MAX_GLOBAL_PRIORITY_SIZE":      16,
	"VK_MAX_SHADER_MODULE_IDENTIFIER_SIZE_EXT": 32,
	"VK_MAX_PIPELINE_BINARY_KEY_SIZE_KHR":      32,
}

type XMLTypes struct {
	Types []XMLType `xml:"type"`
}

type XMLMember struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	Len      string `xml:"len,attr"`
	AltLen   string `xml:"altlen,attr"`
	Optional string `xml:"optional,attr"`
	Values   string `xml:"values,attr"`

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
// Handles both numeric sizes like [4] and named constants like [<enum>VK_MAX_EXTENSION_NAME_SIZE</enum>].
func fixedArraySize(m XMLMember) int {
	idx := strings.Index(m.InnerXML, "</name>")
	if idx < 0 {
		return 0
	}
	after := m.InnerXML[idx+len("</name>"):]

	// Try numeric first
	match := fixedArrayRe.FindStringSubmatch(after)
	if match != nil {
		n := 0
		for _, ch := range match[1] {
			n = n*10 + int(ch-'0')
		}
		return n
	}

	// Try named constant: [<enum>VK_MAX_FOO</enum>]
	enumMatch := namedArrayRe.FindStringSubmatch(after)
	if enumMatch != nil {
		if val, ok := namedArrayConstants[enumMatch[1]]; ok {
			return val
		}
	}

	return 0
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

	// Second pass: register funcpointer names so they can reference each other
	for _, t := range x.Types.Types {
		if t.Category == "funcpointer" && t.Proto.Name != "" {
			fp := &generator.GoFuncPointer{
				CTypeName:  t.Proto.Name,
				GoTypeName: stripPFNvk(t.Proto.Name),
			}
			r.FuncPointers[fp.GoTypeName] = fp
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

	// Fourth pass: resolve funcpointer param/return types now that struct names are registered
	for _, t := range x.Types.Types {
		if t.Category == "funcpointer" && t.Proto.Name != "" {
			resolveFuncPointerTypes(t, r.Handles, r.FuncPointers, r.Structs)
		}
	}

	// Fifth pass: fully parse struct fields (all struct names now registered)
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
func resolveFuncPointerTypes(t XMLType, handles map[string]*generator.GoHandle, funcPointers map[string]*generator.GoFuncPointer, structs map[string]*generator.Structured) {
	fp, ok := funcPointers[stripPFNvk(t.Proto.Name)]
	if !ok {
		return
	}

	// Resolve return type
	protoIsPtr := strings.Contains(t.Proto.InnerXML, "*")
	if t.Proto.Type == "void" && protoIsPtr {
		fp.Return = &generator.VoidPtr{}
	} else if t.Proto.Type != "void" {
		fp.Return = resolveFieldType(t.Proto.Type, protoIsPtr, false, handles, funcPointers, structs)
	}
	// else: void return, fp.Return stays nil

	// Resolve param types
	for _, p := range t.FuncParams {
		isPtr := strings.Contains(p.InnerXML, "*")
		ft := resolveFieldType(p.Type, isPtr, false, handles, funcPointers, structs)
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

	s.HasSType, s.SType = detectSType(t)

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
		fixedSize := fixedArraySize(m)
		// If there's a len attribute that references another member, it's a slice
		// BUT fixed-size arrays take precedence over len-based slices
		isArr := fixedSize == 0 && m.Len != "" && m.Len != "null-terminated"

		isDblPtr := strings.Count(m.InnerXML, "*") >= 2
		ft := resolveFieldType(m.Type, isPtr, isArr, handles, funcPointers, structs, isDblPtr)

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
// A count field is only collapsed when exactly one array references it.
func buildCountMap(members []XMLMember) map[string]string {
	// Count how many arrays reference each count field
	type arrayRef struct {
		name     string
		optional bool
	}
	refs := map[string][]arrayRef{}
	for _, m := range members {
		if m.Len == "" || m.Len == "null-terminated" {
			continue
		}
		parts := strings.Split(m.Len, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" && part != "null-terminated" {
				refs[part] = append(refs[part], arrayRef{
					name:     m.Name,
					optional: m.Optional == "true",
				})
			}
		}
	}
	// Only collapse when exactly one array uses the count.
	// Don't collapse when the array is purely optional (optional="true"),
	// as the count has independent meaning. But optional="false,true" is fine.
	result := map[string]string{}
	for countField, arrayFields := range refs {
		if len(arrayFields) == 1 && !arrayFields[0].optional {
			result[countField] = arrayFields[0].name
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

// detectSType checks if a struct has sType+pNext and extracts the Go sType constant name.
func detectSType(t XMLType) (bool, string) {
	var sTypeValue string
	hasPNext := false
	for _, m := range t.Members {
		switch m.Name {
		case "sType":
			if m.Values != "" {
				sTypeValue = m.Values
			}
		case "pNext":
			hasPNext = true
		}
	}
	if sTypeValue == "" || !hasPNext {
		return false, ""
	}
	// Convert VK_STRUCTURE_TYPE_FOO_BAR to StructureTypeFooBar using the same
	// naming convention as enum elements
	goName := UpperSnakeToPascal(sTypeValue)
	return true, goName
}
