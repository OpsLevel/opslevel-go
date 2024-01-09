package localparser

import (
	"fmt"
	"strings"
)

const (
	INPUT_NOT_LEN_ONE         = "input is not len 1"
	INPUT_NOT_LEN_TWO         = "input is not len 2"
	INPUT_NOT_LEN_ONE_OR_TWO  = "input is not len 1 or len 2"
	INPUT_NOT_STRING          = "input is not string"
	INPUT_NOT_INPUT_TYPE      = "input is not an Input"
	OUTPUT_NOT_POINTER        = "output is not a pointer"
	OUTPUT_NOT_ERROR          = "output is not an error"
	OUTPUT_NOT_LEN_ONE        = "output is not len 1"
	OUTPUT_NOT_LEN_TWO        = "output is not len 2"
	OUTPUT_NOT_LEN_ONE_OR_TWO = "output is not len 1 or len 2"
)

// res, err := create(input)
func validateCreate(fn *Function) {
	inputLenOneInputType(fn)
	outputLenTwo(fn)
}

// res, err := update(string, input)
// res, err := update(input)
func validateUpdate(fn *Function) {
	if len(fn.Input) != 1 && len(fn.Input) != 2 {
		complain(fn, INPUT_NOT_LEN_ONE_OR_TWO)
	} else if len(fn.Input) == 1 {
		inputLenOneInputType(fn)
	} else {
		inputLenTwo(fn)
	}
	outputLenTwo(fn)
}

// res, err := get(string)
func validateGet(fn *Function) {
	inputLenOneString(fn)
	outputLenTwo(fn)
}

// res, err := delete(string)
// err := delete(string)
func validateDelete(fn *Function) {
	inputLenOneString(fn)
	if len(fn.Output) != 1 && len(fn.Output) != 2 {
		complain(fn, OUTPUT_NOT_LEN_ONE_OR_TWO)
	} else if len(fn.Output) == 1 {
		outputLenOne(fn)
	} else {
		outputLenTwo(fn)
	}
}

// res, err := list(nil)
func validateList(fn *Function) {
	if len(fn.Input) != 1 {
		complain(fn, INPUT_NOT_LEN_ONE)
	}
	outputLenTwo(fn)
}

func inputLenOneString(fn *Function) {
	if len(fn.Input) != 1 {
		complain(fn, INPUT_NOT_LEN_ONE)
	} else {
		if fn.Input[0] != "string" {
			complain(fn, INPUT_NOT_STRING)
		}
	}
}

func inputLenOneInputType(fn *Function) {
	if len(fn.Input) != 1 {
		complain(fn, INPUT_NOT_LEN_ONE)
	} else {
		if !isInputType(fn.Input[0]) {
			complain(fn, INPUT_NOT_INPUT_TYPE)
		}
	}
}

func inputLenTwo(fn *Function) {
	if len(fn.Input) != 2 {
		complain(fn, INPUT_NOT_LEN_TWO)
	} else {
		if fn.Input[0] != "string" {
			complain(fn, INPUT_NOT_STRING)
		}
		if !isInputType(fn.Input[1]) {
			complain(fn, INPUT_NOT_INPUT_TYPE)
		}
	}
}

func outputLenOne(fn *Function) {
	if len(fn.Output) != 1 {
		complain(fn, OUTPUT_NOT_LEN_ONE)
	} else {
		if fn.Output[0] != "error" {
			complain(fn, OUTPUT_NOT_ERROR)
		}
	}
}

func outputLenTwo(fn *Function) {
	if len(fn.Output) != 2 {
		complain(fn, OUTPUT_NOT_LEN_TWO)
	} else {
		if !isPointer(fn.Output[0]) {
			complain(fn, OUTPUT_NOT_POINTER)
		}
		if fn.Output[1] != "error" {
			complain(fn, OUTPUT_NOT_ERROR)
		}
	}
}

func complain(fn *Function, msg string) {
	fmt.Printf("%s\n\t%s\n", fn, msg)
}

func isPointer(s string) bool {
	return s[0] == '*'
}

func isInputType(s string) bool {
	return !isPointer(s) && strings.HasSuffix(s, "Input")
}
