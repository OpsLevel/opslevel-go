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
	verbs      = []string{"Create", "Update", "Get", "Delete", "List", "Assign", "Unassign"}
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
func (res *Resource) Score() int {
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
			hasVerb := false
			for _, v := range verbs {
				if strings.Contains(funcName, v) {
					hasVerb = true
					break
				}
			}
			if !hasVerb {
				return true
			}

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

func matchResourcesToFunctions() {
	for _, res := range resources {
		for _, verb := range verbs {
			funcName := verb + res.Name // Try CreateRes, UpdateRes, GetRes, ...
			if verb == "List" {
				funcName += "s"
			}
			if v, ok := functions[funcName]; ok {
				switch verb {
				case "Create":
					res.Create = v
				case "Update":
					res.Update = v
				case "Get":
					res.Get = v
				case "Delete":
					res.Delete = v
				case "List":
					res.List = v
				case "Assign":
					res.Assign = v
				case "Unassign":
					res.Unassign = v
				default:
					panic("hello world")
				}
			}
		}
	}
}

func validatedMatchedFunctions() {
	for _, res := range resources {
		if res.Score() == 0 {
			continue
		}
		if res.Create != nil {
			validateCreate(res.Create)
		}
		if res.Update != nil {
			validateUpdate(res.Update)
		}
		if res.Get != nil {
			validateGet(res.Get)
		}
		if res.Delete != nil {
			validateDelete(res.Delete)
		}
		if res.List != nil {
			validateList(res.List)
		}
	}
}

func showRankings() {
	ranked := maps.Keys(resources)
	sort.Slice(ranked, func(i, j int) bool {
		return resources[ranked[i]].Score() > resources[ranked[j]].Score()
	})
	fmt.Printf("SCORE\tRESOURCE\n")
	for _, res := range ranked {
		fmt.Printf("%d\t%s\n", resources[res].Score(), resources[res].Name)
	}
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

func RunParser() error {
	err := parse()
	if err != nil {
		return err
	}
	matchResourcesToFunctions()
	showRankings()
	validatedMatchedFunctions()
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
