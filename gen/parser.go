package gen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"github.com/TwiN/go-color"
)

var (
	functions = make(map[string]Function)
	resources = make(map[string]struct{}) // this is a set.
)

func GetFunctions() map[string]Function {
	return functions
}

func GetResources() map[string]struct{} {
	return resources
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

func RunParser() error {
	if err := parse(); err != nil {
		return err
	}
	for _, fn := range functions {
		if fn.Full() || fn.Resource == "" {
			continue
		}
		println(color.InYellow(fn))
	}
	for _, fn := range functions {
		if fn.Full() {
			println(color.InGreen(fn))
			continue
		}
	}
	return nil
}
