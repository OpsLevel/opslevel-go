package gen

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

var (
	BlockedGlobs   = []string{"check", "execraw", "variables", "ctx", "with", "query"}
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
	fmt.Println(color.With(color.Red, out))
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

func ToResourceName(s string) string {
	for _, verb := range SupportedVerbs {
		s = strings.TrimPrefix(s, verb)
	}
	s = strings.TrimSuffix(s, "CreateInput")
	s = strings.TrimSuffix(s, "UpdateInput")
	s = strings.TrimSuffix(s, "Input")
	s = strings.TrimSuffix(s, "Connection")
	s = strings.TrimSuffix(s, "s")
	switch s {
	case "ID":
		return ""
	}
	return s
}

func setResourceName(fn *Function, resName string) {
	validatedName := ToResourceName(resName)
	if resName != "" && validatedName == "" {
		complain("not a proper resource name: '%s'", resName)
		return
	}
	fn.Resource = validatedName
	resources[validatedName] = struct{}{}
}

func ParseResource(fn *Function) {
	var resourceName string
	if fn.Resource != "" {
		resources[fn.Resource] = struct{}{}
		return
	}
	// TODO: try get from verb?
	if len(fn.Input) == 1 {
		if fn.Input[0][0] != '*' && strings.HasSuffix(fn.Input[0], "Input") {
			setResourceName(fn, resourceName)
			return
		}
	}
	if len(fn.Output) == 2 {
		// get from output
		if fn.Output[0][0] == '*' {
			resourceName = fn.Output[0][1:]
			setResourceName(fn, resourceName)
			return
		} else if fn.Output[0][0] == '[' {
			resourceName = fn.Output[0][2:]
			setResourceName(fn, resourceName)
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
	return fmt.Sprintf("%s(%s) (%s)\n\tresource '%s' case '%s' verb '%s'", fn.Name, strings.Join(fn.Input, ", "),
		strings.Join(fn.Output, ", "), fn.Resource, fn.Case, fn.Verb)
}

func (fn Function) Full() bool {
	return fn.Verb != "" && fn.Resource != "" && fn.Case != ""
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
	ParseResource(function)

	functions = append(functions, *function)
}
