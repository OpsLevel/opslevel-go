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
	"strings"
)

// TODO:
// write hardcodes to custom.json
// write generated to generated.json
// only wipe generated.json
// make CLI alert if unknown things are added or something changes.

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

func (res *Resource) isPerfect() bool {
	// TODO: special cases - not all resources have the standard create/update/get/delete/list.
	// TODO: update PR template to ask users to verify they checked the output of .json files
	switch res.Name {
	case "Check":
		return false // TODO: checks are very complicated, exclude them for now.
	case "Tag":
		return res.NumFunctions() == 4
	case "Repository", "Property", "Contact":
		return res.NumFunctions() == 3
	case "AlertSourceService", "Dependency", "Alias":
		return res.NumFunctions() == 2
	case "Tiers", "Lifecycle":
		return res.NumFunctions() == 1
	}

	// general case
	return res.NumFunctions() == 5
}

func (res *Resource) String() string {
	b, _ := json.Marshal(res)
	return string(b)
}

func (res *Resource) PrefCreateInputType() string {
	fmt.Println("I am a " + res.Name)
	if res.Create == nil {
		return ""
	}
	switch res.Name {
	case "User", "Secret":
		return res.Create.Input[1]
	}
	return res.Create.Input[0]
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
			// TODO: there are still some separate get with alias functions, like on Repository.
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
	fmt.Println()
}

func showRankings() {
	buckets := make(map[int][]string)
	for _, res := range resources {
		if buckets[res.NumFunctions()] == nil {
			buckets[res.NumFunctions()] = []string{}
		}
		buckets[res.NumFunctions()] = append(buckets[res.NumFunctions()], res.Name)
	}
	for i := len(buckets); i > 0; i-- {
		if _, ok := buckets[i]; !ok {
			continue
		}
		fmt.Printf("has %d functions: [%s]\n", i, strings.Join(buckets[i], ", "))
	}
	fmt.Println()
}

func writeConfigs() error {
	for _, res := range resources {
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
	// TODO: should we add more verbs, like Connect and Invite?
	// TODO: the following are not linked fully
	// ServiceRepository - ConnectServiceRepository
	// Integrations - CreateIntegrationAWS, CreateIntegrationNewRelic, UpdateIntegrationAWS, UpdateIntegrationNewRelic
	// AlertSource - GetAlertSourceWithExternalIdentifier

	// Could not link part of resource
	if resources["Secret"].List == nil {
		resources["Secret"].List = functions["ListSecretsVaultsSecret"]
	}
	if resources["User"].List == nil {
		resources["User"].Create = functions["InviteUser"]
	}

	// Could not link resource
	if resources["Action"] == nil {
		resources["Action"] = &Resource{
			Name:   "Action",
			Create: functions["CreateWebhookAction"],
		}
	}
	if resources["Dependency"] == nil {
		resources["Dependency"] = &Resource{
			Name:   "Dependency", // TODO: this has a different name
			Create: functions["CreateServiceDependency"],
			Delete: functions["DeleteServiceDependency"],
		}
	}
	if resources["Alias"] == nil {
		resources["Alias"] = &Resource{
			Name:   "Alias",
			Create: functions["CreateAlias"],
			Delete: functions["DeleteAlias"],
		}
	}
	if resources["Infra"] == nil {
		resources["Infra"] = &Resource{
			Name:   "Infra", // TODO: this has a different name
			Create: functions["CreateInfrastructure"],
			Update: functions["UpdateInfrastructure"],
			Get:    functions["GetInfrastructure"],
			Delete: functions["DeleteInfrastructure"],
			List:   functions["ListInfrastructure"],
		}
	}
	if resources["Member"] == nil { // TODO: uses verbs add, remove
		resources["Member"] = &Resource{
			Name: "Member", // TODO: this has a different name
		}
		resources["Member"].Create = &Function{
			Name:  "TODO: I am not real",
			Input: []string{"TeamMembershipUserInput"},
		}
	}
}

func filterResources() {
	for key, res := range resources {
		if !res.isPerfect() {
			// delete(resources, key)
			fmt.Println("not perfect - " + key)
		}
	}
	fmt.Println()
}

func RunParser() error {
	err := parse()
	if err != nil {
		return err
	}
	mapResourcesToFunctions()
	addManualOverrides()
	showRankings()
	validateResources() // can only show warnings of functions that were mapped.
	filterResources()
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
