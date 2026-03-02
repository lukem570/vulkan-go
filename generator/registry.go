package generator

type Registry struct {
	Enums    map[string]*Enum
	Bitmasks map[string]*Bitmask
}

func (r *Registry) AddEnumElement(extends string, el *EnumElement) {
	if e, ok := r.Enums[extends]; ok {
		el.parent = e
		e.elements = append(e.elements, el)
	}
}