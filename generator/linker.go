package generator

type Linker struct {
	Registry *Registry
	Config   *Config

	// reverse-lookup maps: C name → Go name (commands are keyed by CName directly)
	cToGoType      map[string]string
	cToGoEnum      map[string]string
	cToGoBitmask   map[string]string
	cToGoEnumAlias map[string]string
}

type Enabled struct {
	Commands     map[string]bool // keyed by CName (e.g. "vkCreateInstance")
	Types        map[string]bool // keyed by Go name (structs + handles)
	Enums        map[string]bool // keyed by Go name
	Bitmasks     map[string]bool // keyed by Go name
	FuncPointers map[string]bool // keyed by Go name
	EnumAliases  map[string]bool
}

func (l *Linker) buildIndexes() {
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

	l.cToGoEnumAlias = make(map[string]string, len(l.Registry.EnumAliases))
	for goName, b := range l.Registry.EnumAliases {
		l.cToGoEnumAlias[b.CName] = goName
	}
}

func (l *Linker) Link() *Registry {
	l.buildIndexes()

	enabled := &Enabled{
		Commands:     map[string]bool{},
		Types:        map[string]bool{},
		Enums:        map[string]bool{},
		Bitmasks:     map[string]bool{},
		EnumAliases:  map[string]bool{},
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
			l.enableFeatureRecursive(e, f.Name)
		}
	}
}

// enableFeatureRecursive enables a feature and all its dependencies.
func (l *Linker) enableFeatureRecursive(e *Enabled, name string) {
	f := l.Registry.Features[name]
	if f == nil {
		return
	}
	for _, dep := range f.Depends {
		l.enableFeatureRecursive(e, dep)
	}
	l.enableRequireBlocks(e, f.Requires)
}

func (l *Linker) enableExtensions(e *Enabled) {
	for _, name := range l.Config.Extensions {
		ext := l.Registry.Extensions[name]
		if ext == nil {
			continue
		}
		// If the extension has a platform requirement, check it's in the config
		if ext.Platform != "" && !l.platformEnabled(ext.Platform) {
			continue
		}
		l.enableRequireBlocks(e, ext.Requires)
		// Tag commands and structs with the extension's platform
		if ext.Platform != "" {
			l.tagPlatform(ext)
		}
	}
}

func (l *Linker) platformEnabled(platform string) bool {
	for _, p := range l.Config.Platforms {
		if p == platform {
			return true
		}
	}
	return false
}

func (l *Linker) tagPlatform(ext *Extension) {
	for _, req := range ext.Requires {
		for _, cName := range req.Commands {
			if cmd, ok := l.Registry.Commands[cName]; ok {
				cmd.Platform = ext.Platform
			}
		}
		for _, cName := range req.Types {
			if goName, ok := l.cToGoType[cName]; ok {
				if s, ok := l.Registry.Structs[goName]; ok {
					s.Platform = ext.Platform
				}
				if h, ok := l.Registry.Handles[goName]; ok {
					h.Platform = ext.Platform
				}
			}
		}
	}
}

// enableRequireBlocks translates C names from the XML require blocks into
// enabled sets. Commands are enabled by CName; types/enums/bitmasks by Go name.
func (l *Linker) enableRequireBlocks(e *Enabled, reqs []RequireBlock) {
	for _, r := range reqs {
		for _, cName := range r.Commands {
			if _, ok := l.Registry.Commands[cName]; ok {
				e.Commands[cName] = true
			}
		}
		for _, cName := range r.Types {
			if goName, ok := l.cToGoType[cName]; ok {
				e.Types[goName] = true
			} else if goName, ok := l.cToGoEnum[cName]; ok {
				e.Enums[goName] = true
			} else if goName, ok := l.cToGoBitmask[cName]; ok {
				e.Bitmasks[goName] = true
			} else if goName, ok := l.cToGoEnumAlias[cName]; ok {
				e.EnumAliases[goName] = true
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

		for cName := range e.Commands {
			cmd := l.Registry.Commands[cName]
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
		if !e.EnumAliases[v.Name] {
			if _, isEnumAlias := l.Registry.EnumAliases[v.Name]; isEnumAlias {
				e.EnumAliases[v.Name] = true
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
		EnumAliases:  map[string]*EnumAlias{},
		Handles:      map[string]*GoHandle{},
		FuncPointers: map[string]*GoFuncPointer{},
		STypes:       l.Registry.STypes,
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
	for name := range e.EnumAliases {
		if v, ok := l.Registry.EnumAliases[name]; ok {
			out.EnumAliases[name] = v
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
	for cName := range e.Commands {
		if v, ok := l.Registry.Commands[cName]; ok {
			out.Commands[cName] = v
			// Detect commands that create a handle from a callback-bearing struct.
			// For such commands, the param's callbackCleanupFn must be attached to
			// the returned handle so the holder outlives the create call.
			if v.CallbackStructParamName == "" {
				for _, p := range v.Params {
					ptr, ok := p.Type.(*Pointer)
					if !ok {
						continue
					}
					st, ok := ptr.Child.(*StructType)
					if !ok {
						continue
					}
					s := l.Registry.Structs[st.Name]
					if s == nil {
						continue
					}
					// Only match sType-bearing structs (real CreateInfo types, not AllocationCallbacks).
					if s.HasSType && len(s.pfnFields()) > 0 {
						_, udCName := s.userDataField()
						if udCName != "" {
							v.CallbackStructParamName = p.Name
							break
						}
					}
				}
			}
		}
	}
	for name := range e.FuncPointers {
		if v, ok := l.Registry.FuncPointers[name]; ok {
			out.FuncPointers[name] = v
		}
	}

	out.APIConstants = l.Registry.APIConstants
	out.Platforms = l.Registry.Platforms

	return out
}
