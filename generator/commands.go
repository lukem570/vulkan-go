package generator

import (
	"fmt"
	"strings"
)

type GoCommand struct {
	Name         string
	CName        string // original C name e.g. vkCreateInstance
	ReceiverType string // e.g. "Instance" if first param was a handle
	HasError     bool
	ReturnType   FieldType    // non-VkResult return, if any
	OutParams    []OutParam   // params that are actually return values
	Params       []CommandParam
	CReturnType  string   // C return type e.g. "VkResult", "void"
	CParams      []CParam // full C prototype params for wrapper generation
}

// CParam holds the full C type and name for a command parameter.
type CParam struct {
	Type string // e.g. "const VkInstanceCreateInfo*"
	Name string // e.g. "pCreateInfo"
}

// OutParam describes a parameter that is an output pointer in C but a return
// value in Go.
type OutParam struct {
	Name string
	Type FieldType
}

type CommandParam struct {
	Name string
	Type FieldType
}

func (c *GoCommand) GenerateWrapper() string {
	if c == nil {
		return ""
	}

	var b strings.Builder
	g := &CodeGen{}

	// ---- Function signature ------------------------------------------------
	if c.ReceiverType != "" {
		b.WriteString(fmt.Sprintf("func (h %s) %s(\n", c.ReceiverType, c.Name))
	} else {
		b.WriteString(fmt.Sprintf("func %s(\n", c.Name))
	}

	for _, p := range c.Params {
		b.WriteString(fmt.Sprintf("\t%s %s,\n", p.Name, p.Type.GoName()))
	}
	b.WriteString(")")

	// ---- Return types -------------------------------------------------------
	var returns []string
	for _, op := range c.OutParams {
		returns = append(returns, op.Type.GoName())
	}
	if c.ReturnType != nil {
		returns = append(returns, c.ReturnType.GoName())
	}
	if c.HasError {
		returns = append(returns, "error")
	}
	if len(returns) == 1 {
		b.WriteString(" " + returns[0])
	} else if len(returns) > 1 {
		b.WriteString(" (" + strings.Join(returns, ", ") + ")")
	}
	b.WriteString(" {\n")

	// ---- Body ---------------------------------------------------------------
	b.WriteString("\tcancels := make([]func(), 0)\n")
	b.WriteString("\tdefer func() { for _, c := range cancels { c() } }()\n\n")

	// Convert Go → C for each input param
	cArgs := make([]string, 0)

	if c.ReceiverType != "" {
		cArgs = append(cArgs, fmt.Sprintf("C.%s(unsafe.Pointer(h.handle))", "Vk"+c.ReceiverType))
	}

	for _, p := range c.Params {
		cVar := g.Var("c_" + sanitizeIdent(p.Name))
		b.WriteString(fmt.Sprintf("\t// param %s\n", p.Name))
		result := p.Type.GenerateToC(g, p.Name)
		b.WriteString(g.String())
		g = &CodeGen{varIndex: g.varIndex}
		_ = result
		_ = cVar
		cArgs = append(cArgs, result)
	}

	// Allocate output param slots
	for _, op := range c.OutParams {
		varName := sanitizeIdent(op.Name) + "Out"
		b.WriteString(fmt.Sprintf("\tvar %s %s\n", varName, op.Type.CName()))
		cArgs = append(cArgs, "&"+varName)
	}

	// Call C function
	callLine := fmt.Sprintf("\t_result := C.fn_%s(%s)\n", c.CName, strings.Join(cArgs, ", "))
	if !c.HasError && c.ReturnType == nil && len(c.OutParams) == 0 {
		callLine = fmt.Sprintf("\tC.fn_%s(%s)\n", c.CName, strings.Join(cArgs, ", "))
	}
	b.WriteString(callLine)

	// Error check
	if c.HasError {
		b.WriteString("\tif _result != C.VK_SUCCESS {\n")
		b.WriteString("\t\treturn ")
		for range c.OutParams {
			b.WriteString("nil, ")
		}
		if c.ReturnType != nil {
			b.WriteString("nil, ")
		}
		b.WriteString("vkError(_result)\n\t}\n")
	}

	// Convert out params from C → Go
	var retVars []string
	for _, op := range c.OutParams {
		varName := sanitizeIdent(op.Name) + "Out"
		outG := &CodeGen{varIndex: g.varIndex}
		goVal := op.Type.GenerateFromC(outG, varName)
		b.WriteString(outG.String())
		g.varIndex = outG.varIndex
		retVars = append(retVars, goVal)
	}

	if c.HasError {
		retVars = append(retVars, "nil")
	}
	if len(retVars) > 0 {
		b.WriteString("\treturn " + strings.Join(retVars, ", ") + "\n")
	}

	b.WriteString("}\n\n")
	return b.String()
}

// GenerateCWrapperDecl returns the C function declaration for the .h file.
func (c *GoCommand) GenerateCWrapperDecl() string {
	params := make([]string, len(c.CParams))
	for i, p := range c.CParams {
		params[i] = p.Type + " " + p.Name
	}
	paramStr := strings.Join(params, ", ")
	if paramStr == "" {
		paramStr = "void"
	}
	return fmt.Sprintf("%s fn_%s(%s);\n", c.CReturnType, c.CName, paramStr)
}

// GenerateCWrapperImpl returns the C function implementation for the .c file.
func (c *GoCommand) GenerateCWrapperImpl() string {
	params := make([]string, len(c.CParams))
	args := make([]string, len(c.CParams))
	for i, p := range c.CParams {
		params[i] = p.Type + " " + p.Name
		args[i] = p.Name
	}
	paramStr := strings.Join(params, ", ")
	if paramStr == "" {
		paramStr = "void"
	}
	argStr := strings.Join(args, ", ")

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s fn_%s(%s) {\n", c.CReturnType, c.CName, paramStr))
	if c.CReturnType == "void" {
		b.WriteString(fmt.Sprintf("\t%s(%s);\n", c.CName, argStr))
	} else {
		b.WriteString(fmt.Sprintf("\treturn %s(%s);\n", c.CName, argStr))
	}
	b.WriteString("}\n\n")
	return b.String()
}

func sanitizeIdent(s string) string {
	s = strings.TrimPrefix(s, "p")
	if s == "" {
		return "val"
	}
	return strings.ToLower(s[:1]) + s[1:]
}