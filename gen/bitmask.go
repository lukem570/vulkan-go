package main

import "fmt"

type Typedef struct {
	Type  CType
	Alias string
}

func (t *Typedef) Generate() string {
	if t.Type.BaseType == "" || t.Type.BaseType == "void" {
		return ""
	}

	return fmt.Sprintf(
		"\ntype %s %s\n",
		t.Alias,
		t.Type.Generate(),
	)
}

func ParseTypedef(v string) Typedef {
	return Typedef{
		Type: CType{
			BaseType: getTextBetweenTags(v, "type"),
		},
		Alias: getTextBetweenTags(v, "name"),
	}
}
