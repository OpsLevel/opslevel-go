package gen

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	INPUT_LEN_1_STRING_OR_ID     = regexp.MustCompile(`^(string|ID)$`)
	INPUT_LEN_1_INPUT_TYPE       = regexp.MustCompile(`^[A-Za-z]*Input$`)
	INPUT_LEN_1_PAYLOAD_VARS_PTR = regexp.MustCompile(`^\*PayloadVariables$`)
	INPUT_LEN_2                  = regexp.MustCompile(`^(string|ID), [A-Za-z]*Input$`)
	INPUT_LEN_2_STRINGS          = regexp.MustCompile(`^(string|ID), (string|ID)$`)

	OUTPUT_LEN_1_ERROR          = regexp.MustCompile(`^error$`)
	OUTPUT_LEN_2                = regexp.MustCompile(`^\*[A-Z][A-Za-z]*, error$`)
	OUTPUT_LEN_2_CONNECTION_PTR = regexp.MustCompile(`^\*[A-Za-z]*Connection, error$`)
	OUTPUT_LEN_2_ARRAY          = regexp.MustCompile(`^\[\][A-Za-z]*, error$`)

	CREATE_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_INPUT_TYPE}
	CREATE_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_2}

	UPDATE_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_INPUT_TYPE, INPUT_LEN_2}
	UPDATE_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_2}

	GET_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_STRING_OR_ID}
	GET_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_2}

	DELETE_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_STRING_OR_ID}
	DELETE_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_1_ERROR, OUTPUT_LEN_2}

	// note: list has a special case where input is empty and the output is an array
	LIST_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_PAYLOAD_VARS_PTR}
	LIST_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_2_CONNECTION_PTR}

	ASSIGN_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_INPUT_TYPE}
	ASSIGN_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_2_ARRAY, OUTPUT_LEN_2}

	UNASSIGN_INPUTS  = []*regexp.Regexp{INPUT_LEN_2_STRINGS}
	UNASSIGN_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_1_ERROR}
)

func complain(fn *Function, msg string) {
	fmt.Printf("%s - %s\n", msg, fn)
}

func validateFunction(fn *Function, possibleInputs []*regexp.Regexp, possibleOutputs []*regexp.Regexp) {
	var (
		inputAsBytes  = []byte(strings.Join(fn.Input, ", "))
		outputAsBytes = []byte(strings.Join(fn.Output, ", "))
		inputIsValid  = false
		outputIsValid = false
	)

	for _, r := range possibleInputs {
		if r.Match(inputAsBytes) {
			inputIsValid = true
			break
		}
	}
	for _, r := range possibleOutputs {
		if r.Match(outputAsBytes) {
			outputIsValid = true
			break
		}
	}

	if !inputIsValid {
		complain(fn, "unexpected input types")
	}
	if !outputIsValid {
		complain(fn, "unexpected return types")
	}
}

func validateFunctions() {
	if len(functions) == 0 {
		return
	}
	fmt.Printf("---> validating total of %d mapped functions\n", len(functions))
	for _, fn := range functions {
		if strings.Contains(fn.Name, "Create") || strings.Contains(fn.Name, "Add") || strings.Contains(fn.Name, "Invite") {
			validateFunction(fn, CREATE_INPUTS, CREATE_OUTPUTS)
		} else if strings.Contains(fn.Name, "Update") {
			validateFunction(fn, UPDATE_INPUTS, UPDATE_OUTPUTS)
		} else if strings.Contains(fn.Name, "Get") {
			validateFunction(fn, GET_INPUTS, GET_OUTPUTS)
		} else if strings.Contains(fn.Name, "Delete") || strings.Contains(fn.Name, "Remove") {
			validateFunction(fn, DELETE_INPUTS, DELETE_OUTPUTS)
		} else if strings.Contains(fn.Name, "List") {
			if fn.Input == nil {
				if OUTPUT_LEN_2_ARRAY.Match([]byte(strings.Join(fn.Output, ", "))) {
				} else {
					complain(fn, "unexpected list function with no input")
				}
			} else {
				validateFunction(fn, LIST_INPUTS, LIST_OUTPUTS)
			}
		} else if strings.Contains(fn.Name, "Assign") {
			validateFunction(fn, ASSIGN_INPUTS, ASSIGN_OUTPUTS)
		} else if strings.Contains(fn.Name, "Unassign") {
			validateFunction(fn, UNASSIGN_INPUTS, UNASSIGN_OUTPUTS)
		} else {
			complain(fn, "has no validation rule")
		}
	}
	fmt.Println()
}
