package generator

import "encoding/xml"

// XMLPlatforms is the <platforms> element from vk.xml.
type XMLPlatforms struct {
	Platforms []XMLPlatform `xml:"platform"`
}

// XMLPlatform is a single <platform> element.
type XMLPlatform struct {
	Name    string `xml:"name,attr"`
	Protect string `xml:"protect,attr"`
}

// XMLRegistry is the root <registry> element from vk.xml.
type XMLRegistry struct {
	XMLName    xml.Name      `xml:"registry"`
	Platforms  XMLPlatforms  `xml:"platforms"`
	Types      XMLTypes      `xml:"types"`
	Enums      []XMLEnums    `xml:"enums"`
	Commands   XMLCommands   `xml:"commands"`
	Features   []XMLFeature  `xml:"feature"`
	Extensions XMLExtensions `xml:"extensions"`
}

// XMLFeature is a <feature> element (core Vulkan version).
type XMLFeature struct {
	Name     string       `xml:"name,attr"`
	Depends  string       `xml:"depends,attr"`
	Requires []XMLRequire `xml:"require"`
}

// XMLRequire is a <require> block inside a feature or extension.
type XMLRequire struct {
	Commands []XMLRequireCmd  `xml:"command"`
	Types    []XMLRequireTyp  `xml:"type"`
	Enums    []XMLRequireEnum `xml:"enum"`
}

// XMLRequireCmd is a <command> reference inside a require block.
type XMLRequireCmd struct {
	Name string `xml:"name,attr"`
}

// XMLRequireTyp is a <type> reference inside a require block.
type XMLRequireTyp struct {
	Name string `xml:"name,attr"`
}

// XMLEnums is an <enums> group element.
type XMLEnums struct {
	Name  string    `xml:"name,attr"`
	Type  string    `xml:"type,attr"`
	Enums []XMLEnum `xml:"enum"`
}

// XMLEnum is a single <enum> element within an enums group.
type XMLEnum struct {
	Name      string `xml:"name,attr"`
	Value     string `xml:"value,attr"`
	BitPos    string `xml:"bitpos,attr"`
	Extends   string `xml:"extends,attr"`
	Offset    string `xml:"offset,attr"`
	Dir       string `xml:"dir,attr"`
	Extnumber string `xml:"extnumber,attr"`
}

// XMLRequireEnum is used inside <feature>/<extension> require blocks.
// It has extra fields not present on top-level <enum> elements.
type XMLRequireEnum struct {
	Name      string `xml:"name,attr"`
	Extends   string `xml:"extends,attr"`
	Value     string `xml:"value,attr"`
	BitPos    string `xml:"bitpos,attr"`
	Offset    string `xml:"offset,attr"`
	Dir       string `xml:"dir,attr"`
	Extnumber string `xml:"extnumber,attr"`
	Alias     string `xml:"alias,attr"`
}

// XMLExtensions is the <extensions> element from vk.xml.
type XMLExtensions struct {
	Extensions []XMLExtension `xml:"extension"`
}

// XMLExtension is a single <extension> element.
type XMLExtension struct {
	Name      string       `xml:"name,attr"`
	Number    int          `xml:"number,attr"`
	Platform  string       `xml:"platform,attr"`
	Supported string       `xml:"supported,attr"`
	Requires  []XMLRequire `xml:"require"`
}

// XMLTypes is the <types> element from vk.xml.
type XMLTypes struct {
	Types []XMLType `xml:"type"`
}

// XMLMember is a <member> element inside a struct/union type.
type XMLMember struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	Len      string `xml:"len,attr"`
	AltLen   string `xml:"altlen,attr"`
	Optional string `xml:"optional,attr"`
	Values   string `xml:"values,attr"`

	InnerXML string `xml:",innerxml"`
}

// XMLType is a <type> element from the types section.
type XMLType struct {
	Category string `xml:"category,attr"`
	Name     string `xml:"name,attr"`
	Alias    string `xml:"alias,attr"`
	Requires string `xml:"requires,attr"`
	Parent   string `xml:"parent,attr"`

	Members []XMLMember `xml:"member"`

	// For handle/basetype types the name may be in a child <name> element.
	NameElem string `xml:"name"`

	// For funcpointer types
	Proto      XMLProto   `xml:"proto"`
	FuncParams []XMLParam `xml:"param"`
}

// XMLCommands is the <commands> element from vk.xml.
type XMLCommands struct {
	Commands []XMLCommand `xml:"command"`
}

// XMLCommand is a single <command> element.
type XMLCommand struct {
	Proto  XMLProto   `xml:"proto"`
	Params []XMLParam `xml:"param"`
	Alias  string     `xml:"alias,attr"`
}

// XMLProto is the <proto> element inside a command or funcpointer.
type XMLProto struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	InnerXML string `xml:",innerxml"`
}

// XMLParam is a <param> element inside a command or funcpointer.
type XMLParam struct {
	Type     string `xml:"type"`
	Name     string `xml:"name"`
	Len      string `xml:"len,attr"`
	Optional string `xml:"optional,attr"`
	InnerXML string `xml:",innerxml"`
}
