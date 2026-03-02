package parser

type Registry struct {
	Platforms map[string]*Platform
	Tags map[string]*Tag
}

type Platform struct {
	Name string
	Protect string
	Comment string
}

type Tag struct {
	Name string
	Author string
	Contact string
}

type TypeCategory int

const (
	CategoryInclude TypeCategory = iota + 1
	CategoryDefine
	CategoryBitmask
	CategoryHandle
	CategoryEnum
	CategoryFuncptr
	CategoryStruct
	CategoryUnion
)

type Type interface {
	GetCategory()
}