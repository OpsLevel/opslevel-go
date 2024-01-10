package gen

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

var (
	resources  = make(map[string]*Resource)
	functions  = make(map[string]*Function)
	ignoreGlob = []string{}
)

type Function struct {
	Name   string   `json:"func"`
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

func (fn *Function) String() string {
	b, _ := json.Marshal(fn)
	return string(b)
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

// for typical objects, 5 is perfect
func (res *Resource) NumFunctions() int {
	var i int
	if res.Create != nil {
		i++
	}
	if res.Update != nil {
		i++
	}
	if res.Get != nil {
		i++
	}
	if res.Delete != nil {
		i++
	}
	if res.List != nil {
		i++
	}
	if res.Assign != nil {
		i++
	}
	if res.Unassign != nil {
		i++
	}
	return i
}

func (res *Resource) String() string {
	b, _ := json.Marshal(res)
	return string(b)
}

func cleanTypeName(s string) string {
	s = strings.TrimPrefix(s, "*")
	s = strings.TrimPrefix(s, "[]*")
	s = strings.TrimPrefix(s, "[]")
	return s
}

func isResource(s string) bool {
	var lower string
	s = cleanTypeName(s)
	lower = strings.ToLower(s)
	if strings.Contains(lower, "id") || strings.Contains(lower, "interface") {
		return false
	}
	if strings.HasSuffix(s, "Input") || strings.HasSuffix(s, "Connection") || strings.ToUpper(s)[0] != s[0] {
		return false
	}
	for _, ignore := range ignoreGlob {
		if strings.Contains(s, ignore) {
			return false
		}
	}
	return true
}

func parse() error {
	packages, err := parser.ParseDir(token.NewFileSet(), ".", nil, parser.ParseComments)
	if err != nil {
		return err
	}
	for _, file := range packages["opslevel"].Files {
		ast.FileExports(file)
		ast.Inspect(file, func(n ast.Node) bool {
			var (
				funcName    string
				funcInputs  []string
				funcOutputs []string
				resName     string
			)

			// name must contain verb
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}
			funcName = funcDecl.Name.Name

			// must be a receiver for client
			if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
				return true
			}
			for _, field := range funcDecl.Recv.List {
				if field == nil {
					return true
				}
				for _, name := range field.Names {
					if name == nil || types.ExprString(field.Type) != "*Client" {
						return true
					}
				}
			}

			// parse inputs
			if funcDecl.Type == nil || funcDecl.Type.Params == nil {
				return true
			}
			for _, field := range funcDecl.Type.Params.List {
				if field == nil {
					return true
				}
				funcInputs = append(funcInputs, types.ExprString(field.Type))
			}

			// parse outputs
			if funcDecl.Type == nil || funcDecl.Type.Results == nil {
				return true
			}
			for _, field := range funcDecl.Type.Results.List {
				if field == nil {
					return true
				}
				fieldType := types.ExprString(field.Type)
				funcOutputs = append(funcOutputs, fieldType)
				resName = cleanTypeName(fieldType)
				if isResource(resName) {
					resources[resName] = &Resource{
						Name: resName,
					}
				}
			}

			// save to state
			functions[funcName] = &Function{
				Name:   funcName,
				Input:  funcInputs,
				Output: funcOutputs,
			}
			return true
		})
	}
	return nil
}

func pluralize(s string) string {
	if strings.HasSuffix("s", s) {
		return s + "es"
	}
	if strings.HasSuffix(s, "y") {
		return s[:len(s)-1] + "ies"
	}
	return s + "s"
}

// useful grep command:
// grep -RE '\*Client) [A-Za-z]*Categor[A-Za-z]*' *.go
// TODO: more verbs to support:
// Invite (create equiv for user)
// Connect (service repos)
// Add, Remove (contact)
func mapResourcesToFunctions() {
	for _, res := range resources {
		for _, fn := range functions {
			// Try CreateRes, UpdateRes, GetRes, ...
			if "Create"+res.Name == fn.Name {
				res.Create = fn
			}
			if "Update"+res.Name == fn.Name {
				res.Update = fn
			}
			if "Get"+res.Name == fn.Name {
				res.Get = fn
			}
			if "Delete"+res.Name == fn.Name {
				res.Delete = fn
			}
			// list has multiple cases
			if "List"+res.Name == fn.Name || "List"+pluralize(res.Name) == fn.Name {
				res.List = fn
			}
			// these can be backwards
			if "Assign"+res.Name == fn.Name || res.Name+"Assign" == fn.Name {
				res.Assign = fn
			}
			if "Unassign"+res.Name == fn.Name || res.Name+"Unassign" == fn.Name {
				res.Unassign = fn
			}
		}
	}
}

func validateResources() {
	for _, res := range resources {
		validateResource(res)
	}
}

// TODO: show these as buckets
func showRankings() {
	ranked := maps.Keys(resources)
	sort.Slice(ranked, func(i, j int) bool {
		return resources[ranked[i]].NumFunctions() > resources[ranked[j]].NumFunctions()
	})
	for _, res := range ranked {
		fmt.Printf("parsed %d functions for resource '%s'\n", resources[res].NumFunctions(), resources[res].Name)
	}
	fmt.Println()
}

func writeConfigs() error {
	for _, res := range resources {
		// TODO: move this
		if res.NumFunctions() < 1 {
			continue
		}
		filename := fmt.Sprintf("./gen/config/%s.json", strings.ToLower(res.Name))
		output, err := json.MarshalIndent(res, "", "    ")
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, output, 0o644)
		if err != nil {
			return err
		}
	}
	return nil
}

func wipeConfigs() error {
	contents, err := filepath.Glob("./gen/config/*.json")
	if err != nil {
		return err
	}
	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return err
		}
	}
	fmt.Println("completed writing configs")
	return nil
}

func addManualOverrides() {
	resources["Secret"].List = functions["ListSecretsVaultsSecret"]
	// TODO: is this appropriate?
	resources["User"].Create = functions["InviteUser"]
}

func filterResources() {
	for key, res := range resources {
		if res.NumFunctions() == 5 {
			break
		}
		// TODO: noticed repository has sep. get and alias functions
		// TODO: check this for other resources, make the gen code better.
		// TODO: not complete, automate this by just grepping and seeing how many you got
		switch res.Name {
		case "Tag", "User", "Tool", "Property", "AlertSourceService", "ServiceDependency", "Lifecycle":
			break
		}

		delete(resources, key)
	}
}

func RunParser() error {
	err := parse()
	if err != nil {
		return err
	}
	mapResourcesToFunctions()
	addManualOverrides()
	showRankings()
	validateResources() // shows warnings about only the mapped functions
	// filterResources()   // this is an allow list made by hand
	err = wipeConfigs()
	if err != nil {
		panic(err)
	}
	err = writeConfigs()
	if err != nil {
		panic(err)
	}

	return err
}
