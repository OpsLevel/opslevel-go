package gen

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

var (
	BlockedGlobs   = []string{"execraw", "variables", "ctx", "with", "query"}
	SupportedVerbs = []string{"Create", "Update", "Get", "Delete", "List", "Assign", "Unassign", "Invite"}
	// TODO: add, remove, connect
	// checks
	// tags
	// aliases
	// property
	// runner
	// job
	// membership
	// contact

	// TODO: CLI
	// be able to load this all
	// example, get, create, list
)

func blocked(funcName string) bool {
	for _, glob := range BlockedGlobs {
		if strings.Contains(strings.ToLower(funcName), strings.ToLower(glob)) {
			return true
		}
	}
	return false
}

func complain(template string, args ...any) {
	out := fmt.Sprintf(template, args...)
	println(color.With(color.Red, out))
}

func ParseVerb(fn *Function) {
	if fn.Verb != "" {
		return
	}
	for _, verb := range SupportedVerbs {
		if strings.Contains(fn.Name, verb) {
			fn.Verb = verb
			return
		}
	}
}

func ParseResource(fn *Function) {
	if fn.Resource != "" {
		return
	}
	// TODO: resource association
	// attempt 1: get from output
	// attempt 2: get from input
	// attempt 2: get from verb
	if len(fn.Output) == 2 {
		if fn.Output[0][0] == '*' {
			fn.Resource = fn.Output[0][1:]
			return
		}
	}
}

type Function struct {
	Name     string   `json:"name,omitempty"`
	Verb     string   `json:"verb,omitempty"`
	Case     string   `json:"case,omitempty"`
	Resource string   `json:"resource,omitempty"`
	Input    []string `json:"input,omitempty"`
	Output   []string `json:"output,omitempty"`
}

func (fn Function) String() string {
	return fmt.Sprintf("func %s(%s) (%s) // resource '%s' case '%s' verb '%s'", fn.Name, strings.Join(fn.Input, ", "),
		strings.Join(fn.Output, ", "), fn.Resource, fn.Case, fn.Verb)
}

func (fn Function) Full() bool {
	return fn.Verb != "" && fn.Resource != ""
}

func addFunction(name string, input []string, output []string) {
	var (
		function = &Function{Name: name, Input: input, Output: output}
		blocked  = blocked(name)
	)
	if blocked {
		return
	}

	// TODO: override here.

	ParseVerb(function)
	if function.Verb == "" {
		complain("ParseVerb: couldn't parse '%s'", function.Name)
		return
	}
	GetCase(function)
	if function.Case == "" {
		complain("GetCase: no combo works for function '%s' with verb '%s'", function.Name,
			function.Verb)
		return
	}
	ParseResource(function) // tries again after all functions are added

	functions[function.Name] = *function
}
