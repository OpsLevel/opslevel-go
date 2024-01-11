package gen

import (
	"fmt"
	"slices"
	"strings"
)

type Function struct {
	Name   string   `json:"func"`
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

func (fn *Function) String() string {
	return fmt.Sprintf("%s(%s) (%s)", fn.Name, strings.Join(fn.Input, ", "), strings.Join(fn.Output, ", "))
}

type Resource struct {
	Name     string    `json:"resource"`
	Create   *Function `json:"create,omitempty"`
	Update   *Function `json:"update,omitempty"`
	Get      *Function `json:"get,omitempty"`
	Delete   *Function `json:"delete,omitempty"`
	List     *Function `json:"list,omitempty"`
	Assign   *Function `json:"assign,omitempty"`
	Unassign *Function `json:"unassign,omitempty"`
}

// TODO: using a map would be less lines of code, more dynamic
// TODO: if we use a map we could add more verbs.
func (res *Resource) NumFunctions() int {
	var i int
	for _, fn := range []*Function{res.Create, res.Update, res.Get, res.Delete, res.List, res.Assign, res.Unassign} {
		if fn != nil {
			i++
		}
	}
	return i
}

func (res *Resource) String() string {
	if res.IsPerfect() {
		return fmt.Sprintf("Perfect!\t\t%s", res.Name)
	}
	return fmt.Sprintf("       %d\t\t%s", res.NumFunctions(), res.Name)
}

func (res *Resource) IsPerfect() bool {
	// TODO: find workaround for this hardcode
	if res.NumFunctions() == 5 {
		return true
	} else {
		if res.Name == "Contact" && res.Create != nil && res.Update != nil && res.Delete != nil {
			return true
		} else if res.Name == "Tier" && res.List != nil {
			return true
		} else if res.Name == "Lifecycle" && res.List != nil {
			return true
		} else if res.Name == "ServiceDependency" && res.Create != nil && res.Delete != nil {
			return true
		} else if res.Name == "AlertSourceService" && res.Create != nil && res.Delete != nil {
			return true
		} else if res.Name == "Tool" && res.Create != nil && res.Update != nil && res.Delete != nil {
			return true
		} else if res.Name == "Property" && res.Get != nil && res.Assign != nil && res.Unassign != nil {
			return true
		} else if res.Name == "Tag" && res.Create != nil && res.Update != nil && res.Delete != nil && res.Assign != nil {
			return true
		}
	}
	return false
}

func (res *Resource) PrefCreateInputType() string {
	if strings.HasSuffix(res.Create.Input[0], "Input") {
		return res.Create.Input[0]
	}
	return res.Create.Input[1]
}

func (res *Resource) PrefCreateUpdateType() string {
	if strings.HasSuffix(res.Update.Input[0], "Input") {
		return res.Update.Input[0]
	}
	return res.Update.Input[1]
}

func cleanTypeName(s string) string {
	s = strings.TrimPrefix(s, "*")
	s = strings.TrimPrefix(s, "[]*")
	s = strings.TrimPrefix(s, "[]")
	return s
}

func isResource(s string) bool {
	s = cleanTypeName(s)
	if strings.ToUpper(s)[0] != s[0] || strings.HasSuffix(s, "Input") ||
		strings.HasSuffix(s, "Connection") || strings.HasSuffix(s, "Interface") {
		return false
	}
	if slices.Contains(UNSUPPORTED, s) {
		return false
	}
	return true
}
