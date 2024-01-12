package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"sort"

	"github.com/TwiN/go-color"
)

var (
	functions = []Function{}
	resources = make(map[string]struct{}) // this is a set.
)

func sortFunctions() {
	sort.Slice(functions, func(i, j int) bool {
		return functions[i].Name < functions[j].Name
	})
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// GetService = Function
// verb[resource] = function
func GetData() map[string]map[string]Function {
	var (
		localResources = Keys(resources)
		result         = map[string]map[string]Function{}
	)
	sortFunctions()
	sort.Strings(localResources)
	for _, fn := range functions {
		// add verb
		if _, ok := result[fn.Verb]; !ok {
			result[fn.Verb] = map[string]Function{}
		}
		// add resource
		if _, ok := result[fn.Verb][fn.Resource]; !ok {
			result[fn.Verb][fn.Resource] = fn
		}
	}
	return result
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
			)

			// parse function decl
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}
			funcName = funcDecl.Name.Name
			if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
				return true
			}
			// parse receiver
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
				funcOutputs = append(funcOutputs, types.ExprString(field.Type))
			}
			// save to state
			addFunction(funcName, funcInputs, funcOutputs)
			return true
		})
	}
	return nil
}

func RunParser(verbose bool) error {
	var countFnPartial, countFnFull, countRes int

	if err := parse(); err != nil {
		return err
	}

	sortFunctions()
	if verbose {
		for _, fn := range functions {
			if !fn.Full() {
				fmt.Println(color.InYellow(fn))
				countFnPartial++
			}
		}
		for _, fn := range functions {
			if fn.Full() {
				fmt.Println(color.InGreen(fn))
				countFnFull++
			}
		}
		for res := range resources {
			fmt.Println(res)
			countRes++
		}
	}
	fmt.Printf("partial functions     %d\n", countFnPartial)
	fmt.Printf("full functions        %d\n", countFnFull)
	fmt.Printf("resources             %d\n", countRes)

	return nil
}
