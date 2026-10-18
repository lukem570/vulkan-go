package generator

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------------
// Handle lifetimes
//
// Vulkan destroys handles through explicit vkDestroy*/vkFree* calls. Those
// wrappers are still generated, but unexported: every creation site registers a
// runtime.AddCleanup instead, so destruction happens when the garbage collector
// finds the handle unreachable.
//
// Two invariants keep GC-ordered destruction well formed:
//
//   - A cleanup argument holds its parent handle, so a parent can never be
//     collected (and destroyed) before its children's cleanups have run.
//     Children also store their parent in Handle.parent, which extends the
//     chain upward: Buffer → Device → PhysicalDevice → Instance.
//   - A cleanup argument never holds the handle being destroyed, only that
//     handle's raw unsafe.Pointer, which addresses driver memory. Holding the
//     Go object would keep it reachable, and the cleanup would never run.
//
// What this cannot know is GPU timelines: reachability in Go says nothing about
// whether the device still reads the object. Callers must keep a handle
// reachable for as long as submitted work uses it.
// ---------------------------------------------------------------------------

// destroyer describes how a single handle type is destroyed.
type destroyer struct {
	Handle    string // handle being destroyed, e.g. "Buffer"
	Method    string // unexported method on Parent, e.g. "destroyBuffer"
	Parent    string // receiver type of Method, e.g. "Device"
	Self      bool   // Method's receiver is the handle itself (Instance, Device)
	Allocator bool   // Method takes a trailing *AllocationCallbacks
	HasError  bool   // Method returns error
	PoolType  string // non-empty when Method also takes a pool handle
}

// lifetimes maps a handle's Go name to its destroyer. Populated by
// buildLifetimes before any command is emitted.
var lifetimes = map[string]*destroyer{}

// poolFields maps "<CreateInfoGoName>.<PoolHandleGoName>" to the struct field
// holding that pool, e.g. "CommandBufferAllocateInfo.CommandPool" → "CommandPool".
var poolFields = map[string]string{}

func isDestroyCommand(c *GoCommand) bool {
	return strings.HasPrefix(c.Name, "Destroy") || strings.HasPrefix(c.Name, "Free")
}

func isCreateCommand(c *GoCommand) bool {
	return strings.HasPrefix(c.Name, "Create") || strings.HasPrefix(c.Name, "Allocate")
}

// commandName returns the Go method name to emit for a command. Destroy and
// free commands are unexported: the GC calls them, not the user.
func commandName(c *GoCommand) string {
	if isDestroyCommand(c) {
		return lowerFirstChar(c.Name)
	}
	return c.Name
}

// buildLifetimes pairs every handle with the command that destroys it, and
// records where pool-scoped frees find their pool.
func buildLifetimes(r *ReducedRegistry) {
	lifetimes = map[string]*destroyer{}
	poolFields = map[string]string{}

	for _, k := range sortedKeys(r.Commands) {
		c := r.Commands[k]
		if !isDestroyCommand(c) || c.ReceiverType == "" {
			continue
		}

		d := &destroyer{
			Parent:   c.ReceiverType,
			Method:   commandName(c),
			HasError: c.HasError,
		}

		var handleParams []*Handle
		var sliceElem string
		for _, p := range c.Params {
			if p.CountOf != nil {
				continue
			}
			switch t := p.Type.(type) {
			case *Handle:
				handleParams = append(handleParams, t)
			case *Slice:
				if h, ok := t.Child.(*Handle); ok {
					sliceElem = h.Name
				}
			case *Pointer:
				if st, ok := t.Child.(*StructType); ok && st.Name == "AllocationCallbacks" {
					d.Allocator = true
				}
			}
		}

		switch {
		case sliceElem != "":
			// vkFreeCommandBuffers(device, pool, count, pBuffers) and friends.
			d.Handle = sliceElem
			if len(handleParams) == 1 {
				d.PoolType = handleParams[0].Name
			}
		case len(handleParams) == 1:
			d.Handle = handleParams[0].Name
		case len(handleParams) == 0:
			// vkDestroyDevice(device, pAllocator), vkDestroyInstance(...).
			d.Handle, d.Self = c.ReceiverType, true
		default:
			continue // unrecognised shape: no cleanup for this handle
		}

		if _, exists := lifetimes[d.Handle]; !exists {
			lifetimes[d.Handle] = d
		}
	}

	for _, k := range sortedKeys(r.Structs) {
		s := r.Structs[k]
		for _, f := range s.Fields {
			if h, ok := f.Type.(*Handle); ok {
				poolFields[s.GoName+"."+h.Name] = f.GoName
			}
		}
	}
}

// tableTypeFor returns the dispatch table type for a handle that destroys
// itself, or "" when it carries no table.
func tableTypeFor(handle string) string {
	switch {
	case instanceReceiverSet[handle]:
		return "*instanceTable"
	case deviceReceiverSet[handle]:
		return "*deviceTable"
	default:
		return ""
	}
}

// GenerateCleanups emits the argument struct and cleanup function for every
// handle that has a destroyer.
func (r *ReducedRegistry) GenerateCleanups() string {
	var b strings.Builder

	b.WriteString("// ---------------------------------------------------------------------------\n")
	b.WriteString("// Handle cleanups\n//\n")
	b.WriteString("// Each handle is destroyed by the garbage collector through a runtime.AddCleanup\n")
	b.WriteString("// registered at creation. Cleanup arguments hold the parent handle so parents\n")
	b.WriteString("// outlive their children, and never the handle being destroyed.\n")
	b.WriteString("// ---------------------------------------------------------------------------\n\n")

	for _, name := range sortedKeys(lifetimes) {
		d := lifetimes[name]
		if d.Handle == "" {
			continue
		}

		args := "_cleanupArgs" + d.Handle
		fn := "_cleanup" + d.Handle

		// Argument struct.
		b.WriteString(fmt.Sprintf("type %s struct {\n", args))
		if d.Self {
			if tt := tableTypeFor(d.Handle); tt != "" {
				b.WriteString(fmt.Sprintf("\ttable %s\n", tt))
			}
		} else {
			b.WriteString(fmt.Sprintf("\tparent *%s\n", d.Parent))
		}
		if d.PoolType != "" {
			b.WriteString(fmt.Sprintf("\tpool *%s\n", d.PoolType))
		}
		b.WriteString("\tkeepAlive any\n")
		b.WriteString("\tretain *[]any\n")
		b.WriteString("\thandle unsafe.Pointer\n")
		if d.Allocator {
			b.WriteString("\tallocator *AllocationCallbacks\n")
		}
		b.WriteString("\ttrampoline func()\n")
		b.WriteString("}\n\n")

		// Cleanup function.
		b.WriteString(fmt.Sprintf("func %s(a %s) {\n", fn, args))
		b.WriteString("\tif a.handle == nil {\n\t\treturn\n\t}\n")
		b.WriteString("\t_cleanupMu.Lock()\n\tdefer _cleanupMu.Unlock()\n")

		recv := "a.parent"
		if d.Self {
			fields := []string{"handle: a.handle"}
			if tableTypeFor(d.Handle) != "" {
				fields = append(fields, "table: a.table")
			}
			fields = append(fields, "cleanup: a.trampoline")
			b.WriteString(fmt.Sprintf("\t_self := &%s{%s}\n", d.Handle, strings.Join(fields, ", ")))
			recv = "_self"
		} else {
			b.WriteString(fmt.Sprintf("\tif %s == nil {\n\t\treturn\n\t}\n", recv))
		}

		var callArgs []string
		if d.PoolType != "" {
			b.WriteString("\tif a.pool == nil {\n\t\treturn\n\t}\n")
			callArgs = append(callArgs, "a.pool")
			callArgs = append(callArgs, fmt.Sprintf("[]*%s{{handle: a.handle, cleanup: a.trampoline}}", d.Handle))
		} else if !d.Self {
			callArgs = append(callArgs, fmt.Sprintf("&%s{handle: a.handle, cleanup: a.trampoline}", d.Handle))
		}
		if d.Allocator {
			callArgs = append(callArgs, "a.allocator")
		}

		call := fmt.Sprintf("%s.%s(%s)", recv, d.Method, strings.Join(callArgs, ", "))
		if d.HasError {
			call = "_ = " + call
		}
		b.WriteString("\t" + call + "\n")
		b.WriteString("}\n\n")
	}

	return b.String()
}

// generateCleanupRegistration emits the runtime.AddCleanup call for a handle
// returned by command c. goVar names the *Handle variable, allocParam the
// command's *AllocationCallbacks parameter (empty if it has none), and
// createInfoParam the create-info parameter searched for a pool handle.
//
// It returns "" when the handle has no destroyer (Queue, or a handle obtained
// from a Get call, which the application does not own), or when the creating
// command's receiver is not the destroyer's parent.
func generateCleanupRegistration(c *GoCommand, handle, goVar, allocParam, createInfoParam string) string {
	d := lifetimes[handle]
	if d == nil || !isCreateCommand(c) {
		return ""
	}
	if !d.Self && c.ReceiverType != d.Parent {
		return ""
	}

	var fields []string
	if d.Self {
		if tableTypeFor(d.Handle) != "" {
			fields = append(fields, fmt.Sprintf("table: %s.table", goVar))
		}
	} else {
		fields = append(fields, "parent: h")
	}

	if d.PoolType != "" {
		if createInfoParam == "" {
			return ""
		}
		field := poolFields[paramStructName(c, createInfoParam)+"."+d.PoolType]
		if field == "" {
			return "" // cannot locate the pool; leave this handle without a cleanup
		}
		fields = append(fields, fmt.Sprintf("pool: %s.%s", createInfoParam, field))
	}

	if c.ReceiverType != "" {
		fields = append(fields, "keepAlive: h")
	}
	fields = append(fields, fmt.Sprintf("retain: %s.retain", goVar))
	fields = append(fields, fmt.Sprintf("handle: %s.handle", goVar))
	if d.Allocator {
		if allocParam == "" {
			fields = append(fields, "allocator: nil")
		} else {
			fields = append(fields, fmt.Sprintf("allocator: %s", allocParam))
		}
	}
	fields = append(fields, fmt.Sprintf("trampoline: %s.cleanup", goVar))

	return fmt.Sprintf("\tif %s != nil {\n\t\truntime.AddCleanup(%s, _cleanup%s, _cleanupArgs%s{%s})\n\t}\n",
		goVar, goVar, d.Handle, d.Handle, strings.Join(fields, ", "))
}

// paramStructName returns the Go struct name behind a *Foo parameter.
func paramStructName(c *GoCommand, name string) string {
	for _, p := range c.Params {
		if p.Name != name {
			continue
		}
		if ptr, ok := p.Type.(*Pointer); ok {
			if st, ok := ptr.Child.(*StructType); ok {
				return st.Name
			}
		}
	}
	return ""
}

// commandParamNames finds the parameters a creation site needs for cleanup
// registration: the allocation callbacks, and the create-info struct.
func commandParamNames(c *GoCommand) (allocParam, createInfoParam string) {
	for _, p := range c.Params {
		if p.CountOf != nil {
			continue
		}
		ptr, ok := p.Type.(*Pointer)
		if !ok {
			continue
		}
		st, ok := ptr.Child.(*StructType)
		if !ok {
			continue
		}
		if st.Name == "AllocationCallbacks" {
			allocParam = p.Name
		} else if createInfoParam == "" {
			createInfoParam = p.Name
		}
	}
	return allocParam, createInfoParam
}

// generateRetain emits the retain-list setup for a newly created handle.
//
// Rather than walking create-info structs for handle fields, the create-info
// value itself is retained: it already holds every handle the new object
// depends on, including ones nested inside member structs and slices. That
// keeps a RenderPass alive while a Framebuffer references it, an Image alive
// while its ImageView exists, and so on, which is what makes GC-ordered
// destruction valid. The cost is holding the create-info's non-handle payload
// (notably ShaderModuleCreateInfo.Code) for the lifetime of the object.
func generateRetain(c *GoCommand, goVar string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\t%s.retain = new([]any)\n", goVar))
	for _, p := range c.Params {
		if p.CountOf != nil {
			continue
		}
		switch t := p.Type.(type) {
		case *Pointer:
			st, ok := t.Child.(*StructType)
			if !ok || st.Name == "AllocationCallbacks" {
				continue
			}
		case *Slice:
			if _, ok := t.Child.(*StructType); !ok {
				continue
			}
		case *Handle:
			// A handle passed into a create call (e.g. the pool in an
			// allocate) is a dependency of the result.
		default:
			continue
		}
		b.WriteString(fmt.Sprintf("\tif %s != nil {\n\t\t*%s.retain = append(*%s.retain, %s)\n\t}\n", p.Name, goVar, goVar, p.Name))
	}
	return b.String()
}

// generateMemoryRetain emits retention for the two bind-memory commands, where
// the dependency appears after creation rather than in a create-info.
// The vkBind*Memory2 variants take slices of bind infos and are not covered.
func generateMemoryRetain(c *GoCommand) string {
	if c.CName != "vkBindBufferMemory" && c.CName != "vkBindImageMemory" {
		return ""
	}
	var target, memory string
	for _, p := range c.Params {
		if _, ok := p.Type.(*Handle); !ok {
			continue
		}
		if strings.EqualFold(p.Name, "memory") {
			memory = p.Name
		} else if target == "" {
			target = p.Name
		}
	}
	if target == "" || memory == "" {
		return ""
	}
	return fmt.Sprintf("\tif %s != nil && %s != nil && %s.retain != nil {\n\t\t*%s.retain = append(*%s.retain, %s)\n\t}\n",
		target, memory, target, target, target, memory)
}
