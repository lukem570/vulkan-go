package main

import (
	"fmt"
	"strings"
)

type Function struct {
	Proto  Param   `xml:"proto"`
	Params []Param `xml:"param"`
}

type Param struct {
	InnerXML string `xml:",innerxml"`
}

type Command struct {
	Function
	Success string `xml:"successcodes,attr"`
	Error   string `xml:"errorcodes,attr"`
}

func (p *VkParser) parseCommand(cmd Command) {
	if cmd.Proto.InnerXML == "" {
		return
	}

	proto := ParseCVariable(cmd.Proto.InnerXML)

	// Parse parameters safely
	parameters := make([]CVariable, 0, len(cmd.Params))
	for _, param := range cmd.Params {
		parameters = append(parameters, ParseCVariable(param.InnerXML))
	}

	isArr, isReturn := checkReturnParameters(parameters)

	var returnStr string
	var ret []CVariable

	n := len(parameters)

	if isReturn && n > 0 {
		// Last parameter is the return value
		ret = append(ret, parameters[n-1])
		parameters = parameters[:n-1]
		n--

		if isArr && n > 0 {
			// Second-to-last is the array length, part of return
			ret = append(ret, parameters[n-1])
			parameters = parameters[:n-1]

			ret[len(ret)-2].Type.PtrDepth--

			// Return type is slice
			returnStr = "[]" + ret[len(ret)-2].Type.ToGoType()
		} else {
			if ret[len(ret)-1].Type.ToGoType() != "string" {
				ret[len(ret)-1].Type.PtrDepth--
			}

			returnStr = ret[len(ret)-1].Type.ToGoType()
		}
	}

	var isResult bool
	if proto.Type.BaseType == "VkResult" {
		isResult = true
		if returnStr != "" {
			returnStr = fmt.Sprintf("(%s, error)", returnStr)
		} else {
			returnStr = "error"
		}
	} else if proto.Type.BaseType != "void" || proto.Type.PtrDepth != 0 {
		returnStr = "bool"
	}

	goLogic := generateLogic(proto.Name, parameters, ret, isResult)

	goParams := generateParameters(parameters)

	p.functions.WriteString(fmt.Sprintf(
		"func %s(\n%s) %s {\n%s}\n\n",
		proto.Name,
		goParams,
		returnStr,
		goLogic,
	))
}

func generateLogic(name string, parameters []CVariable, returnType []CVariable, isResult bool) string {
	if len(returnType) != 1 {
		return "// c-go logic\n"
	}

	var logicBuilder strings.Builder

	retType := returnType[0].Type
	retType.BaseType = "C." + retType.BaseType
	logicBuilder.WriteString(fmt.Sprintf(
		"\tvar ret %s\n\n",
		retType.ToGoType(),
	))

	if isResult {
		logicBuilder.WriteString(fmt.Sprintf(
			"\tresult := C.%s(\n",
			name,
		))
	} else {
		logicBuilder.WriteString(fmt.Sprintf(
			"\tC.%s(\n",
			name,
		))
	}
	for _, param := range parameters {
		param.Type.BaseType = "C." + param.Type.BaseType

		name := param.Name
		if name == "type" {
			name = "Type"
		}

		logicBuilder.WriteString(fmt.Sprintf(
			"\t\t(%s)(%s),\n",
			param.Type.ToGoType(),
			name,
		))
	}
	logicBuilder.WriteString("\t\t&ret,\n\t)\n\n")

	if isResult {
		logicBuilder.WriteString("\tif res != C.VK_SUCCESS {\n")
		logicBuilder.WriteString(fmt.Sprintf(
			"\t\treturn %s(nil), fmt.Errorf(\"%s failed: %%v\", VkResult(result))\n\t}\n\n",
			returnType[0].Type.ToGoType(),
			name,
		))
	}

	if isResult {
		logicBuilder.WriteString(fmt.Sprintf(
			"\treturn %s(ret), nil\n",
			returnType[0].Type.ToGoType(),
		))
	} else {
		logicBuilder.WriteString(fmt.Sprintf(
			"\treturn %s(ret)\n",
			returnType[0].Type.ToGoType(),
		))
	}

	return logicBuilder.String()
}

// Safely check for array and return parameters
func checkReturnParameters(parameters []CVariable) (isArr bool, isReturn bool) {

	n := len(parameters)
	if n == 0 {
		return
	}

	last := parameters[n-1]
	if strings.HasPrefix(last.Name, "p") && last.Name != "pAllocator" {
		isReturn = true
	}

	if n >= 2 {
		secondLast := parameters[n-2]
		if strings.HasSuffix(secondLast.Name, "Count") && strings.HasPrefix(secondLast.Name, "p") {
			isArr = true
		} else if strings.HasSuffix(secondLast.Name, "Count") && !strings.HasPrefix(secondLast.Name, "p") {
			isArr = false
			isReturn = false
		}
	}

	return
}

func generateParameters(parameters []CVariable) string {
	var goParams strings.Builder

	for _, param := range parameters {
		name := param.Name
		if name == "type" {
			name = "Type"
		}

		goParams.WriteString(fmt.Sprintf(
			"\t%s %s,\n",
			name,
			param.Type.ToGoType(),
		))
	}

	return goParams.String()
}
