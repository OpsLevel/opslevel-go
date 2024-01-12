package gen

import (
	"regexp"
	"strings"
)

var (
	INPUT_LEN_1_STRING_OR_ID     = `^(string|ID)$`
	INPUT_LEN_1_INPUT_TYPE       = `^[A-Za-z]*Input$`
	INPUT_LEN_1_PAYLOAD_VARS_PTR = `^\*PayloadVariables$`
	INPUT_LEN_2                  = `^(string|ID), [A-Za-z]*Input$`
	INPUT_LEN_2_STRINGS          = `^(string|ID), (string|ID)$`
	OUTPUT_LEN_1_ERROR           = `^error$`
	OUTPUT_LEN_2                 = `^\*[A-Z][A-Za-z]*, error$`
	OUTPUT_LEN_2_ARRAY           = `^\[\][A-Za-z]*, error$`

	AssignArray = Combo{Name: "AssignArray", Input: INPUT_LEN_1_INPUT_TYPE, Output: OUTPUT_LEN_2_ARRAY}
	AssignStd   = Combo{Name: "AssignStd", Input: INPUT_LEN_1_INPUT_TYPE, Output: OUTPUT_LEN_2}
	CreateStd   = Combo{Name: "CreateStd", Input: INPUT_LEN_1_INPUT_TYPE, Output: OUTPUT_LEN_2}
	DeleteOut1  = Combo{Name: "DeleteOut1", Input: INPUT_LEN_1_STRING_OR_ID, Output: OUTPUT_LEN_1_ERROR}
	DeleteOut2  = Combo{Name: "DeleteOut2", Input: INPUT_LEN_1_STRING_OR_ID, Output: OUTPUT_LEN_2}
	GetStd      = Combo{Name: "GetStd", Input: INPUT_LEN_1_STRING_OR_ID, Output: OUTPUT_LEN_2}
	InviteStd   = Combo{Name: "CreateStd", Input: INPUT_LEN_2, Output: OUTPUT_LEN_2}
	ListArray   = Combo{Name: "ListArray", Input: "", Output: OUTPUT_LEN_2_ARRAY}
	ListStd     = Combo{Name: "ListStd", Input: INPUT_LEN_1_PAYLOAD_VARS_PTR, Output: OUTPUT_LEN_2}
	UnassignStd = Combo{Name: "UnassignStd", Input: INPUT_LEN_1_INPUT_TYPE, Output: OUTPUT_LEN_2}
	UpdateLen1  = Combo{Name: "UpdateLen1", Input: INPUT_LEN_1_INPUT_TYPE, Output: OUTPUT_LEN_2}
	UpdateLen2  = Combo{Name: "UpdateLen2", Input: INPUT_LEN_2, Output: OUTPUT_LEN_2}

	Combos = map[string][]Combo{
		"Assign":   {AssignStd, AssignArray},
		"Create":   {CreateStd},
		"Delete":   {DeleteOut1, DeleteOut2},
		"Get":      {GetStd},
		"Invite":   {InviteStd},
		"List":     {ListStd, ListArray},
		"Unassign": {UnassignStd},
		"Update":   {UpdateLen1, UpdateLen2},
	}
)

type Combo struct {
	Name   string
	Input  string
	Output string
}

func GetCase(fn *Function) {
	var (
		input  = strings.Join(fn.Input, ", ")
		output = strings.Join(fn.Output, ", ")
		combos []Combo
	)

	if _, ok := Combos[fn.Verb]; !ok {
		complain("GetCase: unmapped verb on function '%s' with verb '%s'", fn.Name, fn.Verb)
		return
	} else {
		combos = Combos[fn.Verb]
	}
	for _, combo := range combos {
		var (
			inOk, outOk = false, false
			err         error
		)

		inOk, err = regexp.MatchString(combo.Input, input)
		if err != nil {
			complain("GetCase: error ocurred on '%s'", fn.Name)
			panic(err)
		}
		outOk, err = regexp.MatchString(combo.Output, output)
		if err != nil {
			complain("GetCase: error ocurred on '%s'", fn.Name)
			panic(err)
		}
		if inOk && outOk {
			fn.Case = combo.Name
			return
		}
	}
}
