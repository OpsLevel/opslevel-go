package gen

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

// hardcoded cases
var (
	NAME_TRANSLATION = map[string]string{
		"InfrastructureResource":         "Infrastructure",
		"CustomActionsTriggerDefinition": "TriggerDefinition",
	}
	UNSUPPORTED = []string{
		"ServiceId", "ID", "PayloadVariables", "ServiceMaturity", "Runner", "RunnerJob",
		"RunnerScale", "AlertSource", "TeamMembership", "CustomActionsExternalAction", "Integration",
		"Repository", "ServiceRepository", "Check",
	} // also: Alias (does not get parsed.)
)

var (
	resources         = make(map[string]*Resource)
	functions         = make(map[string]*Function)
	unmappedFunctions = make(map[string]*Function)
)

func MustImportGeneratedResources(relativePath string) map[string]*Resource {
	var (
		b        []byte
		err      error
		imported = make(map[string]*Resource)
	)
	b, err = os.ReadFile(relativePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &imported)
	if err != nil {
		panic(err)
	}
	return imported
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
					if _, ok := NAME_TRANSLATION[resName]; ok {
						resName = NAME_TRANSLATION[resName]
					}
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

func mapping() {
fnLoop:
	for fnKey, fn := range functions {
		for _, res := range resources {
			// TODO: would be nice to extract this to top of file
			// hardcoded mappings - make sure to stay inside this loop.
			if res.Name == "Secret" && fn.Name == "ListSecretsVaultsSecret" {
				res.List = fn
				continue fnLoop
			} else if res.Name == "User" && fn.Name == "InviteUser" {
				res.Create = fn
				continue fnLoop
			} else if res.Name == "Contact" {
				if fn.Name == "AddContact" {
					res.Create = fn
					continue fnLoop
				} else if fn.Name == "UpdateContact" {
					res.Update = fn
					continue fnLoop
				} else if fn.Name == "RemoveContact" {
					res.Delete = fn
					continue fnLoop
				}
			}

			name := res.Name
			if "Create"+name == fn.Name {
				res.Create = fn
				continue fnLoop
			} else if "Update"+name == fn.Name {
				res.Update = fn
				continue fnLoop
			} else if "Get"+name == fn.Name {
				res.Get = fn
				continue fnLoop
			} else if "Delete"+name == fn.Name {
				res.Delete = fn
				continue fnLoop
			} else if "List"+name == fn.Name || "List"+pluralize(name) == fn.Name {
				res.List = fn
				continue fnLoop
			} else if "Assign"+name == fn.Name || name+"Assign" == fn.Name {
				res.Assign = fn
				continue fnLoop
			} else if "Unassign"+name == fn.Name || name+"Unassign" == fn.Name {
				res.Unassign = fn
				continue fnLoop
			}
		}
		unmappedFunctions[fnKey] = fn
		delete(functions, fnKey)
	}
}

func showUnmapped() {
	fmt.Printf("---> total of %d unmapped functions\n", len(unmappedFunctions))
	if len(unmappedFunctions) == 0 {
		return
	}
	for _, fn := range unmappedFunctions {
		fmt.Println(fn)
	}
	fmt.Println()
}

func showRankings() {
	fmt.Printf("---> total of %d resources\n", len(resources))
	if len(resources) == 0 {
		return
	}
	for _, resource := range resources {
		fmt.Println(resource)
	}
	fmt.Println()
}

func writeFile() error {
	var (
		err              error
		name             = "./gen/generated.json"
		orderedKeys      = maps.Keys(resources)
		orderedResources = make([]*Resource, len(orderedKeys))
		output           []byte
	)
	sort.Slice(orderedKeys, func(i, j int) bool {
		return orderedKeys[i] < orderedKeys[j]
	})
	for i, resName := range orderedKeys {
		orderedResources[i] = resources[resName]
	}
	output, err = json.MarshalIndent(resources, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(name, output, 0o644)
	if err != nil {
		return err
	}
	return err
}

func RunParser() error {
	var err error
	err = parse()
	if err != nil {
		return err
	}
	mapping()
	showUnmapped()
	showRankings()
	validateFunctions() // only validates mapped functions
	err = writeFile()
	if err != nil {
		panic(err)
	}
	return err
}
