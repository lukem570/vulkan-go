package generator

type Config struct {
	API        string   `yaml:"api"`
	Version    string   `yaml:"version"`
	Extensions []string `yaml:"extensions"`
	Platforms  []string `yaml:"platforms,omitempty"`
}

type Linker struct {
	Registry *Registry
	Config   *Config

	// reverse-lookup maps: C name → Go name
	cToGoCmd    map[string]string
	cToGoType   map[string]string
	cToGoEnum   map[string]string
	cToGoBitmask map[string]string
}

type Enabled struct {
	Commands     map[string]bool // keyed by Go name
	Types        map[string]bool // keyed by Go name (structs + handles)
	Enums        map[string]bool // keyed by Go name
	Bitmasks     map[string]bool // keyed by Go name
	FuncPointers map[string]bool // keyed by Go name
}

func (l *Linker) buildIndexes() {
	l.cToGoCmd = make(map[string]string, len(l.Registry.Commands))
	for goName, cmd := range l.Registry.Commands {
		l.cToGoCmd[cmd.CName] = goName
	}

	l.cToGoType = make(map[string]string)
	for goName, s := range l.Registry.Structs {
		l.cToGoType[s.CName] = goName
	}
	for goName, h := range l.Registry.Handles {
		l.cToGoType[h.CName] = goName
	}

	l.cToGoEnum = make(map[string]string, len(l.Registry.Enums))
	for goName, e := range l.Registry.Enums {
		l.cToGoEnum[e.CName] = goName
		// also index individual element C names → parent Go name
		for _, el := range e.Elements {
			l.cToGoEnum[el.CName] = goName
		}
	}

	l.cToGoBitmask = make(map[string]string, len(l.Registry.Bitmasks))
	for goName, b := range l.Registry.Bitmasks {
		l.cToGoBitmask[b.CName] = goName
	}
}

func (l *Linker) Link() *Registry {
	l.buildIndexes()

	enabled := &Enabled{
		Commands:     map[string]bool{},
		Types:        map[string]bool{},
		Enums:        map[string]bool{},
		Bitmasks:     map[string]bool{},
		FuncPointers: map[string]bool{},
	}

	l.enableCoreVersions(enabled)
	l.enableExtensions(enabled)
	l.resolveDependencies(enabled)

	return l.buildReducedRegistry(enabled)
}

func (l *Linker) enableCoreVersions(e *Enabled) {
	for _, f := range l.Registry.Features {
		if versionLE(f.Name, l.Config.Version) {
			l.enableRequireBlocks(e, f.Requires)
		}
	}
}

func (l *Linker) enableExtensions(e *Enabled) {
	for _, name := range l.Config.Extensions {
		ext := l.Registry.Extensions[name]
		if ext == nil {
			continue
		}
		l.enableRequireBlocks(e, ext.Requires)
	}
}

// enableRequireBlocks translates C names from the XML require blocks into Go
// names and marks them enabled.
func (l *Linker) enableRequireBlocks(e *Enabled, reqs []RequireBlock) {
	for _, r := range reqs {
		for _, cName := range r.Commands {
			if goName, ok := l.cToGoCmd[cName]; ok {
				e.Commands[goName] = true
			}
		}
		for _, cName := range r.Types {
			if goName, ok := l.cToGoType[cName]; ok {
				e.Types[goName] = true
			} else if goName, ok := l.cToGoEnum[cName]; ok {
				e.Enums[goName] = true
			} else if goName, ok := l.cToGoBitmask[cName]; ok {
				e.Bitmasks[goName] = true
			}
		}
		for _, cName := range r.Enums {
			// Individual enum element — enable the parent enum
			if goName, ok := l.cToGoEnum[cName]; ok {
				e.Enums[goName] = true
			}
		}
	}
}

func (l *Linker) resolveDependencies(e *Enabled) {
	changed := true
	for changed {
		changed = false

		for goName := range e.Commands {
			cmd := l.Registry.Commands[goName]
			if cmd == nil {
				continue
			}
			if cmd.ReceiverType != "" && !e.Types[cmd.ReceiverType] {
				if _, ok := l.Registry.Handles[cmd.ReceiverType]; ok {
					e.Types[cmd.ReceiverType] = true
					changed = true
				}
			}
			for _, p := range cmd.Params {
				if l.enableTypeRecursive(e, p.Type) {
					changed = true
				}
			}
			for _, op := range cmd.OutParams {
				if l.enableTypeRecursive(e, op.Type) {
					changed = true
				}
			}
			if cmd.ReturnType != nil {
				if l.enableTypeRecursive(e, cmd.ReturnType) {
					changed = true
				}
			}
		}

		for goName := range e.Types {
			s := l.Registry.Structs[goName]
			if s == nil {
				continue
			}
			for _, f := range s.Fields {
				if l.enableTypeRecursive(e, f.Type) {
					changed = true
				}
			}
		}

		for goName := range e.FuncPointers {
			fp := l.Registry.FuncPointers[goName]
			if fp == nil {
				continue
			}
			for _, p := range fp.Params {
				if l.enableTypeRecursive(e, p.Type) {
					changed = true
				}
			}
			if fp.Return != nil {
				if l.enableTypeRecursive(e, fp.Return) {
					changed = true
				}
			}
		}
	}
}

func (l *Linker) enableTypeRecursive(e *Enabled, t FieldType) bool {
	if t == nil {
		return false
	}
	switch v := t.(type) {
	case *Slice:
		return l.enableTypeRecursive(e, v.Child)
	case *FixedArray:
		return l.enableTypeRecursive(e, v.Child)
	case *Pointer:
		return l.enableTypeRecursive(e, v.Child)
	case *Primitive, *Bool, *VoidPtr, *String, *TypedStructure:
		return false
	case *GoFuncPointer:
		if !e.FuncPointers[v.GoTypeName] {
			e.FuncPointers[v.GoTypeName] = true
			return true
		}
	case *Handle:
		if !e.Types[v.Name] {
			e.Types[v.Name] = true
			return true
		}
	case *StructType:
		if !e.Types[v.Name] {
			e.Types[v.Name] = true
			return true
		}
	case *NamedType:
		// Could be an enum or bitmask — check all maps
		if !e.Types[v.Name] {
			if _, isStruct := l.Registry.Structs[v.Name]; isStruct {
				e.Types[v.Name] = true
				return true
			}
			if _, isHandle := l.Registry.Handles[v.Name]; isHandle {
				e.Types[v.Name] = true
				return true
			}
		}
		if !e.Enums[v.Name] {
			if _, isEnum := l.Registry.Enums[v.Name]; isEnum {
				e.Enums[v.Name] = true
				return true
			}
		}
		if !e.Bitmasks[v.Name] {
			if _, isBitmask := l.Registry.Bitmasks[v.Name]; isBitmask {
				e.Bitmasks[v.Name] = true
				return true
			}
		}
	}
	return false
}

func (l *Linker) buildReducedRegistry(e *Enabled) *Registry {
	out := &Registry{
		Enums:        map[string]*Enum{},
		Structs:      map[string]*Structured{},
		Commands:     map[string]*GoCommand{},
		Bitmasks:     map[string]*Bitmask{},
		Handles:      map[string]*GoHandle{},
		FuncPointers: map[string]*GoFuncPointer{},
	}

	for name := range e.Enums {
		if v, ok := l.Registry.Enums[name]; ok {
			out.Enums[name] = v
		}
	}
	for name := range e.Bitmasks {
		if v, ok := l.Registry.Bitmasks[name]; ok {
			out.Bitmasks[name] = v
		}
	}
	for name := range e.Types {
		if v, ok := l.Registry.Structs[name]; ok {
			out.Structs[name] = v
		}
		if v, ok := l.Registry.Handles[name]; ok {
			out.Handles[name] = v
		}
	}
	for name := range e.Commands {
		if v, ok := l.Registry.Commands[name]; ok {
			out.Commands[name] = v
		}
	}
	for name := range e.FuncPointers {
		if v, ok := l.Registry.FuncPointers[name]; ok {
			out.FuncPointers[name] = v
		}
	}

	out.APIConstants = l.Registry.APIConstants

	return out
}