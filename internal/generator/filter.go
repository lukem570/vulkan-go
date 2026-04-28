package generator

// filter holds state for the filtering pass that reduces a full Registry to
// the subset needed for code generation.
type filter struct {
	registry *Registry
	config   *Config

	// reverse-lookup maps: C name → Go name (commands are keyed by CName directly)
	cToGoType      map[string]string
	cToGoEnum      map[string]string
	cToGoBitmask   map[string]string
	cToGoEnumAlias map[string]string
}

// Enabled tracks which registry items have been enabled for code generation.
type Enabled struct {
	Commands     map[string]bool // keyed by CName (e.g. "vkCreateInstance")
	Types        map[string]bool // keyed by Go name (structs + handles)
	Enums        map[string]bool // keyed by Go name
	Bitmasks     map[string]bool // keyed by Go name
	FuncPointers map[string]bool // keyed by Go name
	EnumAliases  map[string]bool
}

// Filter applies the Config to the Registry and returns a ReducedRegistry
// containing only the types, commands, and enums needed for the enabled
// Vulkan version and extensions.
func Filter(r *Registry, c *Config) *ReducedRegistry {
	f := &filter{registry: r, config: c}
	f.buildIndexes()

	enabled := &Enabled{
		Commands:     map[string]bool{},
		Types:        map[string]bool{},
		Enums:        map[string]bool{},
		Bitmasks:     map[string]bool{},
		EnumAliases:  map[string]bool{},
		FuncPointers: map[string]bool{},
	}

	f.enableCoreVersions(enabled)
	f.enableExtensions(enabled)
	f.resolveDependencies(enabled)

	return f.buildReducedRegistry(enabled)
}

func (f *filter) buildIndexes() {
	f.cToGoType = make(map[string]string)
	for goName, s := range f.registry.Structs {
		f.cToGoType[s.CName] = goName
	}
	for goName, h := range f.registry.Handles {
		f.cToGoType[h.CName] = goName
	}

	f.cToGoEnum = make(map[string]string, len(f.registry.Enums))
	for goName, e := range f.registry.Enums {
		f.cToGoEnum[e.CName] = goName
		// also index individual element C names → parent Go name
		for _, el := range e.Elements {
			f.cToGoEnum[el.CName] = goName
		}
	}

	f.cToGoBitmask = make(map[string]string, len(f.registry.Bitmasks))
	for goName, b := range f.registry.Bitmasks {
		f.cToGoBitmask[b.CName] = goName
	}

	f.cToGoEnumAlias = make(map[string]string, len(f.registry.EnumAliases))
	for goName, b := range f.registry.EnumAliases {
		f.cToGoEnumAlias[b.CName] = goName
	}
}

func (f *filter) enableCoreVersions(e *Enabled) {
	for _, feat := range f.registry.Features {
		if versionLE(feat.Name, f.config.Version) {
			f.enableFeatureRecursive(e, feat.Name)
		}
	}
}

// enableFeatureRecursive enables a feature and all its dependencies.
func (f *filter) enableFeatureRecursive(e *Enabled, name string) {
	feat := f.registry.Features[name]
	if feat == nil {
		return
	}
	for _, dep := range feat.Depends {
		f.enableFeatureRecursive(e, dep)
	}
	f.enableRequireBlocks(e, feat.Requires)
}

func (f *filter) enableExtensions(e *Enabled) {
	for _, name := range f.config.Extensions {
		ext := f.registry.Extensions[name]
		if ext == nil {
			continue
		}
		// If the extension has a platform requirement, check it's in the config
		if ext.Platform != "" && !f.platformEnabled(ext.Platform) {
			continue
		}
		f.enableRequireBlocks(e, ext.Requires)
		// Tag commands and structs with the extension's platform
		if ext.Platform != "" {
			f.tagPlatform(ext)
		}
	}
}

func (f *filter) platformEnabled(platform string) bool {
	for _, p := range f.config.Platforms {
		if p == platform {
			return true
		}
	}
	return false
}

func (f *filter) tagPlatform(ext *Extension) {
	for _, req := range ext.Requires {
		for _, cName := range req.Commands {
			if cmd, ok := f.registry.Commands[cName]; ok {
				cmd.Platform = ext.Platform
			}
		}
		for _, cName := range req.Types {
			if goName, ok := f.cToGoType[cName]; ok {
				if s, ok := f.registry.Structs[goName]; ok {
					s.Platform = ext.Platform
				}
				if h, ok := f.registry.Handles[goName]; ok {
					h.Platform = ext.Platform
				}
			}
		}
	}
}

// enableRequireBlocks translates C names from the XML require blocks into
// enabled sets. Commands are enabled by CName; types/enums/bitmasks by Go name.
func (f *filter) enableRequireBlocks(e *Enabled, reqs []RequireBlock) {
	for _, r := range reqs {
		for _, cName := range r.Commands {
			if _, ok := f.registry.Commands[cName]; ok {
				e.Commands[cName] = true
			}
		}
		for _, cName := range r.Types {
			if goName, ok := f.cToGoType[cName]; ok {
				e.Types[goName] = true
			} else if goName, ok := f.cToGoEnum[cName]; ok {
				e.Enums[goName] = true
			} else if goName, ok := f.cToGoBitmask[cName]; ok {
				e.Bitmasks[goName] = true
			} else if goName, ok := f.cToGoEnumAlias[cName]; ok {
				e.EnumAliases[goName] = true
			}
		}
		for _, cName := range r.Enums {
			// Individual enum element — enable the parent enum
			if goName, ok := f.cToGoEnum[cName]; ok {
				e.Enums[goName] = true
			}
		}
	}
}

func (f *filter) resolveDependencies(e *Enabled) {
	changed := true
	for changed {
		changed = false

		for cName := range e.Commands {
			cmd := f.registry.Commands[cName]
			if cmd == nil {
				continue
			}
			if cmd.ReceiverType != "" && !e.Types[cmd.ReceiverType] {
				if _, ok := f.registry.Handles[cmd.ReceiverType]; ok {
					e.Types[cmd.ReceiverType] = true
					changed = true
				}
			}
			for _, p := range cmd.Params {
				if f.enableTypeRecursive(e, p.Type) {
					changed = true
				}
			}
			for _, op := range cmd.OutParams {
				if f.enableTypeRecursive(e, op.Type) {
					changed = true
				}
			}
			if cmd.ReturnType != nil {
				if f.enableTypeRecursive(e, cmd.ReturnType) {
					changed = true
				}
			}
		}

		for goName := range e.Types {
			s := f.registry.Structs[goName]
			if s == nil {
				continue
			}
			for _, field := range s.Fields {
				if f.enableTypeRecursive(e, field.Type) {
					changed = true
				}
			}
		}

		for goName := range e.FuncPointers {
			fp := f.registry.FuncPointers[goName]
			if fp == nil {
				continue
			}
			for _, p := range fp.Params {
				if f.enableTypeRecursive(e, p.Type) {
					changed = true
				}
			}
			if fp.Return != nil {
				if f.enableTypeRecursive(e, fp.Return) {
					changed = true
				}
			}
		}
	}
}

func (f *filter) enableTypeRecursive(e *Enabled, t FieldType) bool {
	if t == nil {
		return false
	}
	switch v := t.(type) {
	case *Slice:
		return f.enableTypeRecursive(e, v.Child)
	case *FixedArray:
		return f.enableTypeRecursive(e, v.Child)
	case *Pointer:
		return f.enableTypeRecursive(e, v.Child)
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
			if _, isStruct := f.registry.Structs[v.Name]; isStruct {
				e.Types[v.Name] = true
				return true
			}
			if _, isHandle := f.registry.Handles[v.Name]; isHandle {
				e.Types[v.Name] = true
				return true
			}
		}
		if !e.Enums[v.Name] {
			if _, isEnum := f.registry.Enums[v.Name]; isEnum {
				e.Enums[v.Name] = true
				return true
			}
		}
		if !e.Bitmasks[v.Name] {
			if _, isBitmask := f.registry.Bitmasks[v.Name]; isBitmask {
				e.Bitmasks[v.Name] = true
				return true
			}
		}
		if !e.EnumAliases[v.Name] {
			if _, isEnumAlias := f.registry.EnumAliases[v.Name]; isEnumAlias {
				e.EnumAliases[v.Name] = true
				return true
			}
		}
	}
	return false
}

func (f *filter) buildReducedRegistry(e *Enabled) *ReducedRegistry {
	out := &ReducedRegistry{
		Enums:        map[string]*Enum{},
		Structs:      map[string]*Structured{},
		Commands:     map[string]*GoCommand{},
		Bitmasks:     map[string]*Bitmask{},
		EnumAliases:  map[string]*EnumAlias{},
		Handles:      map[string]*GoHandle{},
		FuncPointers: map[string]*GoFuncPointer{},
		STypes:       f.registry.STypes,
	}

	for name := range e.Enums {
		if v, ok := f.registry.Enums[name]; ok {
			out.Enums[name] = v
		}
	}
	for name := range e.Bitmasks {
		if v, ok := f.registry.Bitmasks[name]; ok {
			out.Bitmasks[name] = v
		}
	}
	for name := range e.EnumAliases {
		if v, ok := f.registry.EnumAliases[name]; ok {
			out.EnumAliases[name] = v
		}
	}
	for name := range e.Types {
		if v, ok := f.registry.Structs[name]; ok {
			out.Structs[name] = v
		}
		if v, ok := f.registry.Handles[name]; ok {
			out.Handles[name] = v
		}
	}
	for cName := range e.Commands {
		if v, ok := f.registry.Commands[cName]; ok {
			out.Commands[cName] = v
			// Detect commands that create a handle from a callback-bearing struct.
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
					s := f.registry.Structs[st.Name]
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
		if v, ok := f.registry.FuncPointers[name]; ok {
			out.FuncPointers[name] = v
		}
	}

	out.APIConstants = f.registry.APIConstants
	out.Platforms = f.registry.Platforms

	return out
}
