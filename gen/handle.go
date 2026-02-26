package main

import "fmt"

type Handle struct {
	Name string
}

func (h *Handle) Generate() string {
	if h.Name == "" {
		return ""
	}

	return fmt.Sprintf(
		"\ntype %s struct {\n\thandle unsafe.Pointer\n}\n",
		h.Name,
	)
}

func ParseHandle(v string) *Handle {
	return &Handle{
		Name: getTextBetweenTags(v, "name"),
	}
}