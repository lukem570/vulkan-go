package generator

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ParseXML reads and unmarshals vk.xml at the given path.
func ParseXML(path string) (*XMLRegistry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reg := &XMLRegistry{}
	if err := xml.Unmarshal(data, reg); err != nil {
		return nil, err
	}

	return reg, nil
}

// BuildRegistry builds the full Registry from a parsed XMLRegistry.
func BuildRegistry(x *XMLRegistry) *Registry {
	r := &Registry{
		Enums:        x.parseEnums(),
		APIConstants: x.parseConstants(),
		Bitmasks:     make(map[string]*Bitmask),
		EnumAliases:  make(map[string]*EnumAlias),
		Structs:      make(map[string]*Structured),
		Handles:      make(map[string]*GoHandle),
		Commands:     make(map[string]*GoCommand),
		STypes:       make(map[string]string),
		Features:     x.parseFeatures(),
		Extensions:   x.parseExtensions(),
		FuncPointers: make(map[string]*GoFuncPointer),
		Platforms:    make(map[string]string),
	}

	for _, p := range x.Platforms.Platforms {
		r.Platforms[p.Name] = p.Protect
	}

	x.parseTypes(r)
	x.parseCommands(r)
	x.applyEnumExtensions(r)

	return r
}

// ---------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------

func (x *XMLRegistry) parseConstants() []APIConstant {
	constants := make([]APIConstant, 0)

	for _, enums := range x.Enums {
		if enums.Type != "" && enums.Type != "constants" {
			continue
		}

		for _, val := range enums.Enums {
			if c := newApiConstant(val); c != nil {
				constants = append(constants, *c)
			}
		}
	}

	return constants
}

func newApiConstant(val XMLEnum) *APIConstant {
	goName := UpperSnakeToPascal(val.Name)
	raw := strings.TrimSpace(val.Value)

	goVal, goType := translateCLiteral(raw)
	if goVal == "" {
		return nil
	}

	return &APIConstant{
		GoName: goName,
		Value:  goVal,
		Type:   goType,
	}
}

// translateCLiteral converts a C constant expression to a Go expression.
// Returns (goValue, goType). goType is empty for plain integer literals.
func translateCLiteral(raw string) (string, string) {
	if raw == "" {
		return "", ""
	}

	digits := extractDigits(raw)

	if strings.ContainsAny(raw, ".fF") {
		return digits, "float32"
	}

	base := ""
	if strings.ContainsAny(raw, "uU") {
		base = "uint"
	} else {
		base = "int"
	}

	bits := ""
	if strings.Contains(strings.ToLower(raw), "ll") {
		bits = "64"
	} else {
		bits = "32"
	}

	not := ""
	if strings.Contains(raw, "~") {
		not = "^"
	}

	typ := fmt.Sprintf("%s%s", base, bits)
	value := fmt.Sprintf("%s%s(%s)", not, typ, digits)

	return value, typ
}

// ---------------------------------------------------------------------------
// Features
// ---------------------------------------------------------------------------

func (x *XMLRegistry) parseFeatures() map[string]*Feature {
	features := make(map[string]*Feature)

	for _, f := range x.Features {
		feat := &Feature{Name: f.Name}
		if f.Depends != "" {
			for _, d := range strings.Split(f.Depends, ",") {
				d = strings.TrimSpace(d)
				if d != "" {
					feat.Depends = append(feat.Depends, d)
				}
			}
		}

		for _, req := range f.Requires {
			block := RequireBlock{}
			for _, c := range req.Commands {
				block.Commands = append(block.Commands, c.Name)
			}
			for _, t := range req.Types {
				block.Types = append(block.Types, t.Name)
			}
			for _, e := range req.Enums {
				block.Enums = append(block.Enums, e.Name)
			}
			feat.Requires = append(feat.Requires, block)
		}

		features[feat.Name] = feat
	}

	return features
}

// ---------------------------------------------------------------------------
// Extensions
// ---------------------------------------------------------------------------

func (x *XMLRegistry) parseExtensions() map[string]*Extension {
	extensions := make(map[string]*Extension)

	for _, ext := range x.Extensions.Extensions {
		e := &Extension{
			Name:      ext.Name,
			Platform:  ext.Platform,
			Supported: ext.Supported,
		}

		for _, req := range ext.Requires {
			block := RequireBlock{}
			for _, c := range req.Commands {
				block.Commands = append(block.Commands, c.Name)
			}
			for _, t := range req.Types {
				block.Types = append(block.Types, t.Name)
			}
			for _, en := range req.Enums {
				block.Enums = append(block.Enums, en.Name)
			}
			e.Requires = append(e.Requires, block)
		}

		extensions[e.Name] = e
	}

	return extensions
}

// ---------------------------------------------------------------------------
// Enums
// ---------------------------------------------------------------------------

func (x *XMLRegistry) parseEnums() map[string]*Enum {
	parsed := make(map[string]*Enum)

	for _, enums := range x.Enums {
		if enums.Type == "" || enums.Type == "constants" {
			continue
		}

		e := &Enum{
			CName:     enums.Name,
			GoName:    stripVk(enums.Name),
			Base:      "uint32",
			IsBitmask: enums.Type == "bitmask",
		}

		for _, val := range enums.Enums {
			el := makeEnumElement(val, e)
			if el != nil {
				e.Elements = append(e.Elements, el)
			}
		}

		parsed[e.GoName] = e
	}

	return parsed
}

// applyEnumExtensions processes enum values that appear inside <feature> and
// <extension> require blocks and adds them to already-parsed enums.
func (x *XMLRegistry) applyEnumExtensions(r *Registry) {
	apply := func(enums []XMLRequireEnum, extNumber int) {
		for _, e := range enums {
			if e.Extends == "" {
				continue
			}
			parentGoName := stripVk(e.Extends)
			parent, ok := r.Enums[parentGoName]
			if !ok {
				continue
			}
			el := makeExtendedEnumElement(e, parent, extNumber)
			if el != nil {
				parent.Elements = append(parent.Elements, el)
			}
		}
	}

	for _, f := range x.Features {
		for _, req := range f.Requires {
			apply(req.Enums, 0)
		}
	}

	for _, ext := range x.Extensions.Extensions {
		for _, req := range ext.Requires {
			apply(req.Enums, ext.Number)
		}
	}
}

func makeEnumElement(val XMLEnum, parent *Enum) *EnumElement {
	el := &EnumElement{
		CName:  val.Name,
		GoName: UpperSnakeToPascal(val.Name),
		Parent: parent,
	}
	if val.BitPos != "" {
		bp, err := strconv.Atoi(val.BitPos)
		if err == nil {
			el.BitPos = &bp
		}
	} else {
		el.Value = val.Value
	}
	return el
}

func makeExtendedEnumElement(val XMLRequireEnum, parent *Enum, extNumber int) *EnumElement {
	el := &EnumElement{
		CName:  val.Name,
		GoName: UpperSnakeToPascal(val.Name),
		Parent: parent,
	}

	switch {
	case val.BitPos != "":
		bp, err := strconv.Atoi(val.BitPos)
		if err == nil {
			el.BitPos = &bp
		}
	case val.Offset != "":
		offset, err := strconv.Atoi(val.Offset)
		if err != nil {
			return nil
		}
		num := extNumber
		if val.Extnumber != "" {
			n, err := strconv.Atoi(val.Extnumber)
			if err == nil {
				num = n
			}
		}
		value := 1000000000 + (num-1)*1000 + offset
		if val.Dir == "-" {
			el.Value = strconv.Itoa(-value)
		} else {
			el.Value = strconv.Itoa(value)
		}
	case val.Value != "":
		el.Value = val.Value
	default:
		if val.Alias == "" {
			return nil
		}
		el.Value = UpperSnakeToPascal(val.Alias)
	}

	return el
}

// ---------------------------------------------------------------------------
// Types
// ---------------------------------------------------------------------------

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

func (x *XMLRegistry) parseTypes(r *Registry) {
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
			h := &GoHandle{CName: name, GoName: stripVk(name)}
			r.Handles[h.GoName] = h
		}
	}

	// Second pass: register funcpointer names so they can reference each other
	for _, t := range x.Types.Types {
		if t.Category == "funcpointer" && t.Proto.Name != "" {
			fp := &GoFuncPointer{
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
			r.Structs[stripVk(name)] = &Structured{
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
			b := &Bitmask{CName: name, GoName: stripVk(name)}
			r.Bitmasks[b.GoName] = b
		case "enum":
			if t.Alias == "" {
				continue
			}

			name := stripVk(t.Name)

			r.EnumAliases[name] = &EnumAlias{
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
			fields := make([]StructField, len(aliased.Fields))
			copy(fields, aliased.Fields)
			r.Structs[stripVk(t.Name)] = &Structured{
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
func resolveFuncPointerTypes(t XMLType, handles map[string]*GoHandle, funcPointers map[string]*GoFuncPointer, structs map[string]*Structured) {
	fp, ok := funcPointers[stripPFNvk(t.Proto.Name)]
	if !ok {
		return
	}

	// Resolve return type
	protoIsPtr := strings.Contains(t.Proto.InnerXML, "*")
	if t.Proto.Type == "void" && protoIsPtr {
		fp.Return = &VoidPtr{}
	} else if t.Proto.Type != "void" {
		fp.Return = resolveFieldType(t.Proto.Type, protoIsPtr, false, handles, funcPointers, structs)
	}
	// else: void return, fp.Return stays nil

	// Resolve param types
	for _, p := range t.FuncParams {
		isPtr := strings.Contains(p.InnerXML, "*")
		ft := resolveFieldType(p.Type, isPtr, false, handles, funcPointers, structs)
		fp.Params = append(fp.Params, FuncPointerParam{
			Name: goParamName(p.Name),
			Type: ft,
		})
	}
}

// parseStruct converts an XMLType of category struct/union into a Structured.
func parseStruct(
	t XMLType,
	handles map[string]*GoHandle,
	funcPointers map[string]*GoFuncPointer,
	structs map[string]*Structured,
	sTypes map[string]string,
) *Structured {
	name := t.Name
	if name == "" {
		return nil
	}

	s := &Structured{
		CName:   name,
		GoName:  stripVk(name),
		Fields:  []StructField{},
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
				ft = NewPrimitive("uint32_t", "uint32")
			}
			field := StructField{
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
				ft = &FixedArray{Child: ft, Size: fixedSizes[i]}
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
		field := StructField{
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

// ---------------------------------------------------------------------------
// Commands
// ---------------------------------------------------------------------------

var receiverPrefixes = map[string][]string{
	"CommandBuffer": {"Cmd"},
}

func (x *XMLRegistry) parseCommands(r *Registry) {
	for _, cmd := range x.Commands.Commands {
		if cmd.Alias != "" {
			continue
		}
		c := buildCommand(cmd, r.Handles, r.FuncPointers, r.Structs)
		if c != nil {
			r.Commands[c.CName] = c
		}
	}
}

// extractCType reconstructs the full C type string from an XML param's InnerXML.
func extractCType(innerXML string) string {
	// Check for array suffix after </name> — treat as pointer
	hasArray := false
	if idx := strings.Index(innerXML, "</name>"); idx >= 0 {
		after := innerXML[idx+len("</name>"):]
		if strings.Contains(after, "[") {
			hasArray = true
		}
	}

	// Remove <name>...</name> and everything after
	if idx := strings.Index(innerXML, "<name>"); idx >= 0 {
		innerXML = innerXML[:idx]
	}
	// Strip XML tags, keeping text content
	var result strings.Builder
	inTag := false
	for _, ch := range innerXML {
		if ch == '<' {
			inTag = true
			continue
		}
		if ch == '>' {
			inTag = false
			continue
		}
		if !inTag {
			result.WriteRune(ch)
		}
	}
	// Clean up whitespace
	s := strings.TrimSpace(result.String())
	// Normalize "const VkFoo *" -> "const VkFoo*"
	s = strings.ReplaceAll(s, " *", "*")
	s = strings.ReplaceAll(s, "  ", " ")

	// Array params decay to pointers in C function signatures
	if hasArray && !strings.HasSuffix(s, "*") {
		s += "*"
	}
	return s
}

func buildCommand(
	cmd XMLCommand,
	handles map[string]*GoHandle,
	funcPointers map[string]*GoFuncPointer,
	structs map[string]*Structured,
) *GoCommand {

	rawName := cmd.Proto.Name
	returnTypeName := cmd.Proto.Type

	c := &GoCommand{
		CName:       rawName,
		Name:        methodName(rawName, ""),
		CReturnType: returnTypeName,
	}

	seen := map[string]bool{}

	// Build CParams from all XML params
	for _, p := range cmd.Params {
		if seen[p.Name] {
			continue
		}
		seen[p.Name] = true
		cType := extractCType(p.InnerXML)
		// For double-pointer non-struct/non-handle array params (e.g. const uint32_t* const*),
		// use const void* in the C wrapper so that Go's unsafe.Pointer is accepted by CGo.
		// Struct double-pointer params (ppGeometries, ppBuildRangeInfos) use PtrSlice instead
		// and keep their correct C types.
		{
			cutoff := strings.Index(p.InnerXML, "<name>")
			if cutoff < 0 {
				cutoff = len(p.InnerXML)
			}
			if strings.Count(p.InnerXML[:cutoff], "*") >= 2 && p.Len != "" {
				goName := stripVk(p.Type)
				_, isStruct := structs[goName]
				_, isHandle := handles[goName]
				if !isStruct && !isHandle && primitiveType(p.Type) != nil {
					cType = "const void*"
				}
			}
		}
		c.CParams = append(c.CParams, CParam{
			Type: cType,
			Name: p.Name,
		})
	}

	if c.CReturnType == "" {
		c.CReturnType = "void"
	}

	// Detect VkResult return → HasError, no Go return type
	switch returnTypeName {
	case "VkResult":
		c.HasError = true
	case "void", "":
		// nothing
	default:
		c.ReturnType = resolveFieldType(returnTypeName, false, false, handles, funcPointers, structs)
	}

	params := cmd.Params
	if len(params) == 0 {
		return c
	}

	// ---- Handle receiver detection ----------------------------------------
	// If the first parameter is a Vulkan handle, make it the method receiver.
	firstParam := params[0]
	firstGoName := stripVk(firstParam.Type)
	if _, ok := handles[firstGoName]; ok {
		c.ReceiverType = firstGoName
		params = params[1:]
		c.Name = methodName(rawName, firstGoName)
	}

	// ---- Byte-data pattern detection ----------------------------------------
	byteDataDetected := false
	if len(params) >= 2 && (returnTypeName == "VkResult" || returnTypeName == "void" || returnTypeName == "") {
		secondLast := params[len(params)-2]
		last := params[len(params)-1]

		isSizeParam := secondLast.Type == "size_t" &&
			strings.Contains(secondLast.InnerXML, "*") &&
			!strings.Contains(secondLast.InnerXML, "const")

		isDataParam := last.Type == "void" &&
			last.Len != "" &&
			strings.Contains(last.Len, secondLast.Name) &&
			strings.Contains(last.InnerXML, "*") &&
			!strings.Contains(last.InnerXML, "const")

		if isSizeParam && isDataParam {
			c.ByteDataPattern = &ByteDataInfo{
				SizeCParam: secondLast.Name,
				DataCParam: last.Name,
			}
			params = params[:len(params)-2]
			byteDataDetected = true
		}
	}

	// ---- Enumerate pattern detection ----------------------------------------
	enumerateDetected := false
	if !byteDataDetected && len(params) >= 2 && (returnTypeName == "VkResult" || returnTypeName == "void" || returnTypeName == "") {
		secondLast := params[len(params)-2]
		last := params[len(params)-1]

		isCountParam := secondLast.Type == "uint32_t" &&
			strings.Contains(secondLast.InnerXML, "*") &&
			!strings.Contains(secondLast.InnerXML, "const")

		isArrayParam := last.Len != "" &&
			strings.Contains(last.Len, secondLast.Name) &&
			strings.Contains(last.InnerXML, "*") &&
			!strings.Contains(last.InnerXML, "const")

		if isCountParam && isArrayParam {
			elemType := resolveFieldType(last.Type, false, false, handles, funcPointers, structs)
			// Only use enumerate pattern for types that have working GenerateFromC
			switch elemType.(type) {
			case *Handle, *Primitive, *NamedType, *Bool, *StructType:
				c.EnumeratePattern = &EnumerateInfo{
					CountCParam: secondLast.Name,
					ArrayCParam: last.Name,
					ElementType: elemType,
				}
				params = params[:len(params)-2]
				enumerateDetected = true
			}
		}
	}

	// ---- Output parameter detection ----------------------------------------
	if !enumerateDetected && !byteDataDetected && len(params) > 0 {
		last := params[len(params)-1]
		if isOutputParam(last, returnTypeName, handles, structs) {
			ft := outParamType(last, handles, funcPointers, structs)
			op := OutParam{
				Name: last.Name,
				Type: ft,
			}
			// Detect array output params (e.g. pPipelines with len="createInfoCount")
			if last.Len != "" && !strings.Contains(last.Len, "null-terminated") {
				op.IsArray = true
				op.CountGoParam = resolveCountExpr(last.Len, params, structs)
			}
			c.OutParams = append(c.OutParams, op)
			params = params[:len(params)-1]
		}
	}

	seen = map[string]bool{}

	// ---- Regular input params -----------------------------------------------
	for _, p := range params {
		if seen[p.Name] {
			continue
		}
		seen[p.Name] = true

		// Only inspect the part of InnerXML before <name> so that comments
		// don't give false positives.
		dblPtrCutoff := strings.Index(p.InnerXML, "<name>")
		if dblPtrCutoff < 0 {
			dblPtrCutoff = len(p.InnerXML)
		}
		beforeName := p.InnerXML[:dblPtrCutoff]
		isPtr := strings.Contains(beforeName, "*")
		isArr := p.Len != "" && p.Len != "null-terminated"
		isDblPtr := strings.Count(beforeName, "*") >= 2

		ft := resolveFieldType(p.Type, isPtr, isArr, handles, funcPointers, structs, isDblPtr)

		// Check for fixed-size array (e.g. blendConstants[4])
		if fixedSize := fixedArraySizeParam(p.InnerXML); fixedSize > 0 {
			ft = &FixedArray{Child: ft, Size: fixedSize}
		}

		param := &CommandParam{
			Name: goParamName(p.Name),
			Type: ft,
		}

		if _, ok := ft.(*Slice); ok &&
			strings.Contains(strings.ToLower(c.Params[len(c.Params)-1].Name), "count") {

			c.Params[len(c.Params)-1].CountOf = param
		}

		c.Params = append(c.Params, param)
	}

	return c
}

// methodName strips vk prefix and the handle type prefix, returning a clean Go name.
func methodName(cName string, handle string) string {
	name := strings.TrimPrefix(cName, "vk")
	name = strings.ReplaceAll(name, handle, "")

	if prefixes, ok := receiverPrefixes[handle]; ok {
		for _, prefix := range prefixes {
			name = strings.TrimPrefix(name, prefix)
		}
	}

	return name
}

func isOutputParam(p XMLParam, retType string, handles map[string]*GoHandle, structs map[string]*Structured) bool {
	if p.Name == "" {
		return false
	}
	// must be a pointer
	if !strings.Contains(p.InnerXML, "*") {
		return false
	}
	// const pointers are inputs, not outputs
	if strings.Contains(p.InnerXML, "const") {
		return false
	}
	// must be result-returning or void function
	if retType != "VkResult" && retType != "void" && retType != "" {
		return false
	}
	// void** params are output pointers (e.g. vkMapMemory's ppData).
	if p.Type == "void" {
		return strings.Contains(p.InnerXML, "**")
	}
	// must have p-prefix convention
	if len(p.Name) < 2 || p.Name[0] != 'p' || p.Name[1] < 'A' || p.Name[1] > 'Z' {
		return false
	}
	goName := stripVk(p.Type)
	if _, isHandle := handles[goName]; isHandle {
		return true
	}
	// Primitives and simple types
	if primitiveType(p.Type) != nil {
		return true
	}
	// Structs as output params (for getter-style commands)
	if structs != nil {
		if _, isStruct := structs[goName]; isStruct {
			return true
		}
	}
	return false
}

func outParamType(p XMLParam, handles map[string]*GoHandle, funcPointers map[string]*GoFuncPointer, structs map[string]*Structured) FieldType {
	// The output param is T* in C — the Go type is T
	// Special case: void* output params are unsafe.Pointer
	if p.Type == "void" {
		return &VoidPtr{}
	}
	return resolveFieldType(p.Type, false, false, handles, funcPointers, structs)
}

// fixedArraySizeParam returns the fixed array size from a command param's InnerXML or 0.
func fixedArraySizeParam(innerXML string) int {
	idx := strings.Index(innerXML, "</name>")
	if idx < 0 {
		return 0
	}
	after := innerXML[idx+len("</name>"):]
	// Look for [N] pattern
	start := strings.Index(after, "[")
	end := strings.Index(after, "]")
	if start < 0 || end < 0 || end <= start+1 {
		return 0
	}
	n := 0
	for _, ch := range after[start+1 : end] {
		if ch < '0' || ch > '9' {
			return 0
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

// resolveCountExpr converts a C len expression to a Go expression.
func resolveCountExpr(lenExpr string, params []XMLParam, structs map[string]*Structured) string {
	parts := strings.SplitN(lenExpr, "->", 2)
	if len(parts) != 2 {
		return goParamName(lenExpr)
	}

	paramCName := parts[0]
	fieldCName := parts[1]
	paramGoName := goParamName(paramCName)

	// Find the struct type of the parameter
	var structTypeName string
	for _, p := range params {
		if p.Name == paramCName {
			structTypeName = stripVk(p.Type)
			break
		}
	}

	// Check if the field was collapsed (it's a count for a slice)
	if structTypeName != "" && structs != nil {
		if s, ok := structs[structTypeName]; ok {
			for _, f := range s.Fields {
				if f.CName == fieldCName && f.CountFor != "" {
					// Find the Go name of the slice field it counts
					for _, sf := range s.Fields {
						if sf.GoName == f.CountFor {
							return fmt.Sprintf("uint32(len(%s.%s))", paramGoName, sf.GoName)
						}
					}
				}
			}
		}
	}

	return paramGoName + "." + toPublic(fieldCName)
}

func goParamName(cName string) string {
	name := stripP(cName)
	// avoid Go keywords
	switch name {
	case "type":
		return "typ"
	case "range":
		return "rng"
	case "func":
		return "fn"
	}
	return name
}
