package opslevel

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var StructTagAnalyzer = &analysis.Analyzer{
	Name: "lintStructs",
	Doc:  "validates struct tags",
	Run:  run,
}

// TODO: differentiate between empty tag and unset tag
type TagContents struct {
	graphql string
	json    string
	yaml    string
}

func getTagContents(value string) TagContents {
	// TODO: possible to use single regex to grab all?
	graphqlTagRegex := regexp.MustCompile(`graphql:"([^"]*)"`)
	jsonTagRegex := regexp.MustCompile(`json:"([^"]*)"`)
	yamlTagRegex := regexp.MustCompile(`yaml:"([^"]*)"`)

	var result TagContents
	var match []string
	match = graphqlTagRegex.FindStringSubmatch(value)
	if len(match) > 1 {
		found := match[len(match)-1]
		result.graphql = found
	}
	match = jsonTagRegex.FindStringSubmatch(value)
	if len(match) > 1 {
		found := match[len(match)-1]
		result.json = found
	}
	match = yamlTagRegex.FindStringSubmatch(value)
	if len(match) > 1 {
		found := match[len(match)-1]
		result.yaml = found
	}

	return result
}

func isPointer(field *ast.Field) bool {
	_, ok := field.Type.(*ast.StarExpr)
	return ok
}

func evaluateTags(pass *analysis.Pass, typeSpec *ast.TypeSpec, field *ast.Field) {
	render := fmt.Sprintf("%s.%s %s", typeSpec.Name, field.Names[0].Name, field.Tag.Value)
	tagContents := getTagContents(field.Tag.Value)

	if tagContents.json != tagContents.yaml {
		pass.Reportf(field.Tag.Pos(), "json and yaml tags should be the exact same: %s", render)
	}

	if strings.HasSuffix(typeSpec.Name.Name, "Input") {
		if tagContents.graphql != "" {
			pass.Reportf(field.Tag.Pos(), "remove graphql tag on input field: %s", render)
		}
		if tagContents.json == "" {
			pass.Reportf(field.Tag.Pos(), "missing json tag on input field: %s", render)
		}
		if tagContents.yaml == "" {
			pass.Reportf(field.Tag.Pos(), "missing yaml tag on input field: %s", render)
		}

		if isPointer(field) {
			if !strings.Contains(tagContents.json, "omitempty") || !strings.Contains(tagContents.yaml, "omitempty") {
				pass.Reportf(field.Tag.Pos(), "missing omitempty on pointer input field: %s",
					render)
			}
		} else {
			if strings.Contains(tagContents.json, "omitempty") || strings.Contains(tagContents.yaml, "omitempty") {
				pass.Reportf(field.Tag.Pos(), "remove omitempty on non-pointer input field: %s",
					render)
			}
		}
	} else {
		if tagContents.graphql == "" {
			pass.Reportf(field.Tag.Pos(), "missing graphql tag on non-input field: %s", render)
		}
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		// TODO: is there a way to restrict packages? maybe with CLI?
		if file.Name.Name != "opslevel" {
			continue
		}
		ast.Inspect(file, func(n ast.Node) bool {
			typeSpec, ok := n.(*ast.TypeSpec)
			if !ok {
				// TODO: use real return values (not just here)
				return true
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				return true
			}

			for _, field := range structType.Fields.List {
				// TODO: deal with embedded struct case in a more readable way
				if field.Tag != nil && len(field.Names) > 0 {
					evaluateTags(pass, typeSpec, field)
				}
			}

			return true
		})
	}

	return nil, nil
}
