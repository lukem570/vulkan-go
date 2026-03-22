package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lukem570/vulkan-go/generator"
)

var fixedArrayRe = regexp.MustCompile(`\[(\d+)\]`)
var namedArrayRe = regexp.MustCompile(`\[<enum>(\w+)</enum>\]`)

// namedArrayConstants maps Vulkan constant names to their integer values.
var namedArrayConstants = map[string]int{
	"VK_MAX_PHYSICAL_DEVICE_NAME_SIZE":         256,
	"VK_UUID_SIZE":                             16,
	"VK_LUID_SIZE":                             8,
	"VK_MAX_EXTENSION_NAME_SIZE":               256,
	"VK_MAX_DESCRIPTION_SIZE":                  256,
	"VK_MAX_MEMORY_TYPES":                      32,
	"VK_MAX_MEMORY_HEAPS":                      16,
	"VK_MAX_DEVICE_GROUP_SIZE":                 32,
	"VK_MAX_DRIVER_NAME_SIZE":                  256,
	"VK_MAX_DRIVER_INFO_SIZE":                  256,
	"VK_MAX_GLOBAL_PRIORITY_SIZE_KHR":          16,
	"VK_MAX_GLOBAL_PRIORITY_SIZE":              16,
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
// Only looks before <name> to avoid false positives from <comment> text.
func isPointerMember(m XMLMember) bool {
	cutoff := strings.Index(m.InnerXML, "<name>")
	if cutoff < 0 {
		return strings.Contains(m.InnerXML, "*")
	}
	return strings.Contains(m.InnerXML[:cutoff], "*")
}

// bitfieldWidth returns the bitfield width from InnerXML (e.g. "...<name>foo</name>:24" → 24), or 0 if not a bitfield.
func bitfieldWidth(m XMLMember) int {
	idx := strings.Index(m.InnerXML, "</name>")
	if idx < 0 {
		return 0
	}
	after := strings.TrimSpace(m.InnerXML[idx+len("</name>"):])
	if !strings.HasPrefix(after, ":") {
		return 0
	}
	rest := after[1:]
	var digits string
	for _, ch := range rest {
		if ch >= '0' && ch <= '9' {
			digits += string(ch)
		} else {
			break
		}
	}
	if digits == "" {
		return 0
	}
	n, err := strconv.Atoi(digits)
	if err != nil || n <= 0 {
		return 0
	}
	return n
}

// fixedArraySizes returns all fixed-array dimensions from InnerXML after </name>.
// For example, "matrix[3][4]" → [3, 4], "data[16]" → [16], non-array → nil.
func fixedArraySizes(m XMLMember) []int {
	idx := strings.Index(m.InnerXML, "</name>")
	if idx < 0 {
		return nil
	}
	after := m.InnerXML[idx+len("</name>"):]

	var sizes []int
	for {
		// Try numeric first: [N]
		if match := fixedArrayRe.FindStringIndex(after); match != nil {
			sub := fixedArrayRe.FindStringSubmatch(after)
			n, _ := strconv.Atoi(sub[1])
			sizes = append(sizes, n)
			after = after[match[1]:]
			continue
		}
		// Try named constant: [<enum>VK_MAX_FOO</enum>]
		if match := namedArrayRe.FindStringIndex(after); match != nil {
			sub := namedArrayRe.FindStringSubmatch(after)
			if val, ok := namedArrayConstants[sub[1]]; ok {
				sizes = append(sizes, val)
			}
			after = after[match[1]:]
			continue
		}
		break
	}
	return sizes
}


func (x *XMLRegistry) parseTypes(r *generator.Registry) {
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
		case "enum":
			if t.Alias == "" {
				continue
			}

			name := stripVk(t.Name)

			r.EnumAliases[name] = &generator.EnumAlias{
				CName:  name,
				GoName: name,
				Alias:  stripVk(t.Alias),
			}
		}
	}

	// Fourth pass: resolve funcpointer param/return types now that struct names are registered
	for _, t := range x.Types.Types {
		if t.Category == "funcpointer" && t.Proto.Name != "" {
			resolveFuncPointerTypes(t, r.Handles, r.FuncPointers, r.Structs)
		}
	}

	// Fifth pass: fully parse struct fields (all struct names now registered)
	// Skip alias types — they are resolved in the sixth pass.
	for _, t := range x.Types.Types {
		switch t.Category {
		case "struct", "union":
			if t.Alias != "" {
				continue
			}
			s := parseStruct(t, r.Handles, r.FuncPointers, r.Structs, r.STypes)
			if s != nil {
				r.Structs[s.GoName] = s
			}
		}
	}

	// Sixth pass: resolve struct/union aliases by copying fields from the aliased type.
	for _, t := range x.Types.Types {
		if (t.Category == "struct" || t.Category == "union") && t.Alias != "" {
			aliasedGoName := stripVk(t.Alias)
			aliased, ok := r.Structs[aliasedGoName]
			if !ok {
				continue
			}
			fields := make([]generator.StructField, len(aliased.Fields))
			copy(fields, aliased.Fields)
			r.Structs[stripVk(t.Name)] = &generator.Structured{
				CName:    t.Name,
				GoName:   stripVk(t.Name),
				Fields:   fields,
				IsUnion:  aliased.IsUnion,
				HasSType: aliased.HasSType,
				SType:    aliased.SType,
				Platform: aliased.Platform,
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
func parseStruct(
	t XMLType, 
	handles map[string]*generator.GoHandle, 
	funcPointers map[string]*generator.GoFuncPointer, 
	structs map[string]*generator.Structured,
	sTypes map[string]string,
) *generator.Structured {
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

	if s.HasSType {
		sTypes[s.GoName] = s.SType
	}

	// Build a map of countField → fieldItCounts so we can mark them
	// e.g. descriptorSetCount → pDescriptorSets
	countMap := buildCountMap(t.Members)
	// Reverse map: arrayField → countField (for implicit arrays without len=)
	reverseCountMap := map[string]string{}
	for countField, entry := range countMap {
		reverseCountMap[entry.arrayField] = countField
	}
	seen := map[string]bool{}
	// seenGoNames tracks Go field names already used for visible (non-count) fields.
	// When two C members produce the same Go name, the later one falls back to toPublic(m.Name).
	seenGoNames := map[string]bool{}

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
		if entry, ok := countMap[m.Name]; ok {
			// This is a count field — record it but mark it as CountFor.
			// Preserve the actual C type so auto-set in toC() uses the right cast.
			ft := primitiveType(m.Type)
			if ft == nil {
				ft = generator.NewPrimitive("uint32_t", "uint32")
			}
			field := generator.StructField{
				GoName:     toPublic(stripP(m.Name)),
				CName:      m.Name,
				CountFor:   toPublic(stripP(entry.arrayField)),
				CountCName: m.Name,
				Type:       ft,
			}
			s.Fields = append(s.Fields, field)
			continue
		}

		isPtr := isPointerMember(m)
		fixedSizes := fixedArraySizes(m)
		fixedSize := 0
		if len(fixedSizes) > 0 {
			fixedSize = fixedSizes[0]
		}
		// If there's a len attribute that references another member, it's a slice.
		// Also treat altlen-only fields as arrays (e.g. pCode with latexmath len).
		// Fixed-size arrays take precedence over len-based slices.
		isArr := fixedSize == 0 && (m.Len != "" && m.Len != "null-terminated" || m.AltLen != "")

		// Only count * before <name> to avoid false positives from <comment> text.
		dblPtrCutoff := strings.Index(m.InnerXML, "<name>")
		if dblPtrCutoff < 0 {
			dblPtrCutoff = len(m.InnerXML)
		}
		isDblPtr := strings.Count(m.InnerXML[:dblPtrCutoff], "*") >= 2
		// Double-pointer to char with no len= is an implicit array of strings
		// (e.g. deprecated ppEnabledLayerNames in VkDeviceCreateInfo).
		if isDblPtr && m.Type == "char" && !isArr {
			isArr = true
		}
		ft := resolveFieldType(m.Type, isPtr, isArr, handles, funcPointers, structs, isDblPtr)

		// Wrap in FixedArray for each dimension (e.g. [3][4] → FixedArray{3, FixedArray{4, ft}}).
		// Dimensions are applied innermost-first so matrix[3][4] becomes [3][4]T in Go.
		if len(fixedSizes) > 0 {
			for i := len(fixedSizes) - 1; i >= 0; i-- {
				ft = &generator.FixedArray{Child: ft, Size: fixedSizes[i]}
			}
		}

		// Find the C name of the count field for this slice.
		// For altlen expressions like "codeSize / 4", extract the count field name.
		countCName := ""
		countScale := 0
		if isArr {
			countCName = findCountFieldName(m.Len, t.Members)
			if countCName == "" && m.AltLen != "" {
				if field, divisor := parseAltLenDivExpr(m.AltLen); field != "" {
					countCName = field
					countScale = divisor
				}
			}
			// Fallback: for implicit arrays that have no len= at all,
			// look up the count field from the reverse map.
			if countCName == "" && m.Len == "" && m.AltLen == "" {
				countCName = reverseCountMap[m.Name]
			}
		}

		goFieldName := toPublic(stripP(m.Name))
		if seenGoNames[goFieldName] {
			// Collision: two C members map to the same Go name (e.g. pGeometries and
			// ppGeometries both strip to "Geometries"). Fall back to the raw C name.
			goFieldName = toPublic(m.Name)
		}
		seenGoNames[goFieldName] = true
		field := generator.StructField{
			GoName:     goFieldName,
			CName:      m.Name,
			CountCName: countCName,
			CountScale: countScale,
			Type:       ft,
			BitWidth:   bitfieldWidth(m),
		}
		s.Fields = append(s.Fields, field)
	}

	return s
}

// countEntry describes the relationship between a count field and the array it counts.
type countEntry struct {
	arrayField string
	scale      int // if > 1: C count = len(slice) * scale (e.g. codeSize = len(Code)*4)
}

// buildCountMap returns a map from count-field-name → countEntry describing the array it counts.
// For example: "descriptorSetCount" → {arrayField:"pDescriptorSets", scale:1}
//              "codeSize"           → {arrayField:"pCode", scale:4}  (altlen="codeSize / 4")
// A count field is only collapsed when exactly one array references it.
func buildCountMap(members []XMLMember) map[string]countEntry {
	type arrayRef struct {
		name     string
		optional bool
		scale    int
	}
	refs := map[string][]arrayRef{}
	for _, m := range members {
		lenStr := m.Len
		scale := 1

		if lenStr == "" || lenStr == "null-terminated" {
			// Try altlen alone (e.g. pCode has latexmath in len= but simple altlen)
			if m.AltLen == "" {
				continue
			}
			lenStr = m.AltLen
		}

		// If AltLen provides a simpler human-readable form, prefer it for parsing.
		if m.AltLen != "" {
			if field, divisor := parseAltLenDivExpr(m.AltLen); field != "" {
				lenStr = field
				scale = divisor
			}
		}

		parts := strings.Split(lenStr, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" && part != "null-terminated" {
				refs[part] = append(refs[part], arrayRef{
					name:     m.Name,
					optional: m.Optional == "true",
					scale:    scale,
				})
			}
		}
	}
	// Also pair double-pointer char members that have no len= with their implicit
	// count field (e.g. deprecated ppEnabledLayerNames + enabledLayerCount).
	for i, m := range members {
		if m.Len != "" || m.AltLen != "" {
			continue // already handled via len attribute
		}
		if m.Type != "char" || strings.Count(m.InnerXML, "*") < 2 {
			continue
		}
		// Search backwards for an unclaimed uint32 count member
		for j := i - 1; j >= 0; j-- {
			prev := members[j]
			if _, alreadyRef := refs[prev.Name]; alreadyRef {
				break
			}
			if prev.Type == "uint32_t" && strings.HasSuffix(prev.Name, "Count") {
				refs[prev.Name] = []arrayRef{{name: m.Name, scale: 1}}
				break
			}
		}
	}

	// Only collapse when exactly one array uses the count and it is not optional.
	// Optional arrays (e.g. pImmutableSamplers) mean the count has independent
	// semantic meaning — keep it visible in the Go struct.
	result := map[string]countEntry{}
	for countField, arrayFields := range refs {
		if len(arrayFields) == 1 && !arrayFields[0].optional {
			result[countField] = countEntry{
				arrayField: arrayFields[0].name,
				scale:      arrayFields[0].scale,
			}
		}
	}
	return result
}

// parseAltLenDivExpr parses simple expressions like "codeSize / 4" returning ("codeSize", 4).
// The dividend must be a plain identifier (letters, digits, underscore only).
// Returns ("", 0) for complex expressions like "(rasterizationSamples + 31) / 32".
func parseAltLenDivExpr(s string) (string, int) {
	parts := strings.SplitN(s, "/", 2)
	if len(parts) != 2 {
		return "", 0
	}
	field := strings.TrimSpace(parts[0])
	divisorStr := strings.TrimSpace(parts[1])
	// Reject complex left-hand expressions (spaces, parens, operators)
	for _, ch := range field {
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_') {
			return "", 0
		}
	}
	divisor, err := strconv.Atoi(divisorStr)
	if err != nil || divisor <= 1 {
		return "", 0
	}
	return field, divisor
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
