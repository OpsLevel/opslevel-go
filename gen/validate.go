package gen

import (
	"fmt"
	"regexp"
	"strings"
)

// outputs and inputs of client functions should match something like this
// TODO:
// TODO: what is the expectation for assign and unassign?
// res, err := create(input)
// res, err := update(input)
// res, err := update(string|id, input)
// res, err := get(string|id)
// res, err := delete(string|id)
// err := delete(string|id)
var (
	INPUT_LEN_1_STRING_OR_ID     = regexp.MustCompile(`^(string|ID)$`)
	INPUT_LEN_1_INPUT_TYPE       = regexp.MustCompile(`^[A-Za-z]*Input$`)
	INPUT_LEN_1_PAYLOAD_VARS_PTR = regexp.MustCompile(`^\*PayloadVariables$`)
	INPUT_LEN_2                  = regexp.MustCompile(`^(string|ID), [A-Za-z]*Input$`)

	OUTPUT_LEN_1_ERROR          = regexp.MustCompile(`^error$`)
	OUTPUT_LEN_2                = regexp.MustCompile(`^\*[A-Z][A-Za-z]*, error$`)
	OUTPUT_LEN_2_CONNECTION     = regexp.MustCompile(`^[A-Za-z]*Connection, error$`)
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

	// TODO: list should be more consistent
	// list has a special case where input is empty and the output is an array
	LIST_INPUTS  = []*regexp.Regexp{INPUT_LEN_1_PAYLOAD_VARS_PTR}
	LIST_OUTPUTS = []*regexp.Regexp{OUTPUT_LEN_2_CONNECTION, OUTPUT_LEN_2_CONNECTION_PTR, OUTPUT_LEN_2_ARRAY}
)

func complain(res *Resource, fn *Function, msg string) {
	fmt.Printf("%s - %s(%s) (%s)\n", msg, fn.Name, strings.Join(fn.Input, ", "), strings.Join(fn.Output, ", "))
}

func validateFunction(res *Resource, fn *Function, possibleInputs []*regexp.Regexp, possibleOutputs []*regexp.Regexp) {
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
		complain(res, fn, "unexpected input types")
	}
	if !outputIsValid {
		complain(res, fn, "unexpected return types")
	}
}

func validateResource(res *Resource) {
	if res.Create != nil {
		validateFunction(res, res.Create, CREATE_INPUTS, CREATE_OUTPUTS)
	}
	if res.Update != nil {
		validateFunction(res, res.Update, UPDATE_INPUTS, UPDATE_OUTPUTS)
	}
	if res.Get != nil {
		validateFunction(res, res.Get, GET_INPUTS, GET_OUTPUTS)
	}
	if res.Delete != nil {
		validateFunction(res, res.Delete, DELETE_INPUTS, DELETE_OUTPUTS)
	}
	if res.List != nil {
		outputAsBytes := []byte(strings.Join(res.List.Output, ", "))
		// special case for list - ListTiers() ([]Tiers, empty)
		if res.List.Input == nil {
			if !OUTPUT_LEN_2_ARRAY.Match(outputAsBytes) {
				complain(res, res.List, "unexpected list function with no input")
			}
		} else {
			validateFunction(res, res.List, LIST_INPUTS, LIST_OUTPUTS)
		}
	}
}
