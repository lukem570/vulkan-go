package generator

type Config struct {
	API        string   `yaml:"api"`
	Version    string   `yaml:"version"`
	Extensions []string `yaml:"extensions"`
	Platforms  []string `yaml:"platforms,omitempty"`
}

type Linker struct {
	Registry *Registry // full parsed registry
	Config   *Config   // from config.yml
}

type Enabled struct {
	Commands map[string]bool
	Types    map[string]bool
	Enums    map[string]bool
}

func (l *Linker) Link() *Registry {
	enabled := &Enabled{
		Commands: map[string]bool{},
		Types:    map[string]bool{},
		Enums:    map[string]bool{},
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

func (l *Linker) enableRequireBlocks(e *Enabled, reqs []RequireBlock) {
	for _, r := range reqs {
		for _, cmd := range r.Commands {
			e.Commands[cmd] = true
		}
		for _, t := range r.Types {
			e.Types[t] = true
		}
		for _, en := range r.Enums {
			e.Enums[en] = true
		}
	}
}

func (l *Linker) resolveDependencies(e *Enabled) {
	changed := true

	for changed {
		changed = false

		for name := range e.Commands {

			cmd := l.Registry.Commands[name]
			if cmd == nil {
				continue
			}

			// Parameters
			for _, p := range cmd.Params {
				if l.enableTypeRecursive(e, p.Type) {
					changed = true
				}
			}

			// Return type
			if cmd.ReturnType != nil {
				if l.enableTypeRecursive(e, cmd.ReturnType) {
					changed = true
				}
			}
		}

		for typeName := range e.Types {

			s := l.Registry.Structs[typeName]
			if s == nil {
				continue
			}

			for _, f := range s.Fields {
				if l.enableTypeRecursive(e, f.Type) {
					changed = true
				}
			}
		}
	}
}

func (l *Linker) enableTypeRecursive(e *Enabled, t FieldType) bool {
	switch v := t.(type) {
	case *Array:
		return l.enableTypeRecursive(e, v.child)

	case *Pointer:
		return l.enableTypeRecursive(e, v.child)

	case *Primitive:
		return false

	case *TypedStructure:
		// nothing directly

	case *NamedType: // ← YOU NEED THIS TYPE
		if !e.Types[v.Name] {
			e.Types[v.Name] = true
			return true
		}
	}

	return false
}

func (l *Linker) buildReducedRegistry(e *Enabled) *Registry {
	out := &Registry{
		Enums:    map[string]*Enum{},
		Structs:  map[string]*Structured{},
		Commands: map[string]*GoCommand{},
		Bitmasks: map[string]*Bitmask{},
	}

	for name := range e.Enums {
		out.Enums[name] = l.Registry.Enums[name]
	}

	for name := range e.Types {
		out.Structs[name] = l.Registry.Structs[name]
	}

	for name := range e.Commands {
		out.Commands[name] = l.Registry.Commands[name]
	}

	return out
}
