//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"github.com/hasura/go-graphql-client/ident"
	"github.com/opslevel/opslevel-go/v2024"
)

const (
	// connectionFile  string = "pkg/gen/connection.go"
	enumFile        string = "enum.go"
	inputObjectFile string = "input.go"
	interfacesFile  string = "interfaces.go"
	objectFile      string = "object.go"
	// mutationFile    string = "pkg/gen/mutation.go"
	// payloadFile     string = "pkg/gen/payload.go"
	// queryFile       string = "pkg/gen/query.go"
	// scalarFile      string = "pkg/gen/scalar.go"
	// unionFile       string = "pkg/gen/union.go"
)

var knownTypeIsName = []string{
	"category",
	"filter",
	"level",
}

var knownBoolsByName = []string{
	"enabled",
	"published",
}

var knownIsoTimeByName = []string{
	"createdat",
	"installedat",
	"enableon",
	"updatedat",
}

var stringTypeSuffixes = []string{
	"actionmessage",
	"address",
	"alias",
	"aliases",
	"apidocsdefaultpath",
	"createdat",
	"cursor",
	"description",
	"email",
	"externaluuid",
	"htmlurl",
	"id",
	"kind",
	"liquidtemplate",
	"message",
	"name",
	"processedat",
	"role",
	"queryparams",
	"updatedat",
	"userdeletepayload",
	"url",
	"yaml",
}

var knownTypeMappings = map[string]string{
	"data":                           "JSON",
	"deletedmembers":                 "User",
	"edges":                          "any",
	"filteredcount":                  "int",
	"memberships":                    "TeamMembership",
	"node":                           "any",
	"nodes":                          "[]any",
	"notupdatedrepositories":         "RepositoryOperationErrorPayload",
	"promotedchecks":                 "Check",
	"relationship":                   "RelationshipType",
	"teamsbeingnotified":             "CampaignSendReminderOutcomeTeams",
	"teamsbeingnotifiedcount":        "int",
	"teamsmissingcontactmethod":      "int",
	"teamsmissingcontactmethodcount": "int",
	"type":                           "any",
	"totalcount":                     "int",
	"triggerdefinition":              "CustomActionsTriggerDefinition",
	"updatedrepositories":            "Repository",
	"webhookaction":                  "CustomActionsWebhookAction",
}

const header = `// Code generated by gen.go; DO NOT EDIT.

package opslevel`

type GraphQLSchema struct {
	Types []GraphQLTypes `graphql:"types" json:"types"`
}

type IntrospectiveType struct {
	Name   string `graphql:"name" json:"name"`
	Kind   string `graphql:"kind" json:"kind"`
	OfType struct {
		OfTypeName string `graphql:"name" json:"name"`
	} `graphql:"ofType" json:"ofType"`
}

type GraphQLInputValue struct {
	Name         string            `graphql:"name" json:"name"`
	DefaultValue string            `graphql:"defaultValue" json:"defaultValue"`
	Description  string            `graphql:"description" json:"description"`
	Type         IntrospectiveType `graphql:"type" json:"type"`
}

type GraphQLField struct {
	Args         []GraphQLInputValue `graphql:"args" json:"args"`
	Description  string              `graphql:"description" json:"description"`
	IsDeprecated bool                `graphql:"isDeprecated" json:"isDeprecated"`
	Name         string              `graphql:"name" json:"name"`
}

type GraphQLTypes struct {
	Name          string                `graphql:"name" json:"name"`
	Kind          string                `graphql:"kind" json:"kind"`
	Description   string                `graphql:"description" json:"description"`
	PossibleTypes []GraphQLPossibleType `graphql:"possibleTypes"`
	EnumValues    []GraphQLEnumValues   `graphql:"enumValues" json:"enumValues"`
	Fields        []GraphQLField        `graphql:"fields" json:"fields"`
	InputFields   []GraphQLInputValue   `graphql:"inputFields" json:"inputFields"`
}

type GraphQLEnumValues struct {
	Name        string `graphql:"name" json:"name"`
	Description string `graphql:"description" json:"description"`
}

type GraphQLPossibleType struct {
	Name   string
	Kind   string
	OfType GraphQLOfType
}

type GraphQLOfType struct {
	Name string
	Kind string
}

func GetSchema(client *opslevel.Client) (*GraphQLSchema, error) {
	var q struct {
		Schema GraphQLSchema `graphql:"__schema"`
	}
	if err := client.Query(&q, nil); err != nil {
		return nil, err
	}
	return &q.Schema, nil
}

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func getRootSchema() (*GraphQLSchema, error) {
	token, ok := os.LookupEnv("OPSLEVEL_API_TOKEN")
	if !ok {
		return nil, fmt.Errorf("OPSLEVEL_API_TOKEN environment variable not set")
	}
	client := opslevel.NewGQLClient(opslevel.SetAPIToken(token), opslevel.SetAPIVisibility("public"))
	schema, err := GetSchema(client)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func run() error {
	schema, err := getRootSchema()
	if err != nil {
		return err
	}

	enumSchema := GraphQLSchema{}
	inputObjectSchema := GraphQLSchema{}
	interfaceSchema := GraphQLSchema{}
	objectSchema := GraphQLSchema{}
	scalarSchema := GraphQLSchema{}
	unionSchema := GraphQLSchema{}
	for _, t := range schema.Types {
		switch t.Kind {
		case "ENUM":
			enumSchema.Types = append(enumSchema.Types, t)
		case "INPUT_OBJECT":
			inputObjectSchema.Types = append(inputObjectSchema.Types, t)
		case "INTERFACE":
			interfaceSchema.Types = append(interfaceSchema.Types, t)
		case "OBJECT":
			objectSchema.Types = append(objectSchema.Types, t)
		case "SCALAR":
			scalarSchema.Types = append(scalarSchema.Types, t)
		case "UNION":
			unionSchema.Types = append(unionSchema.Types, t)
		default:
			panic("Unknown GraphQL type: " + t.Kind)
		}
	}

	var buf bytes.Buffer
	var subSchema GraphQLSchema
	for filename, t := range templates {
		switch filename {
		// case connectionFile:
		// 	subSchema = objectSchema
		case enumFile:
			subSchema = enumSchema
		case inputObjectFile:
			subSchema = inputObjectSchema
		case interfacesFile:
			subSchema = interfaceSchema
		case objectFile:
			subSchema = objectSchema
		// case mutationFile:
		// 	subSchema = objectSchema
		// case payloadFile:
		// 	subSchema = objectSchema
		// case queryFile:
		// 	subSchema = objectSchema
		// case scalarFile:
		// 	subSchema = scalarSchema
		// case unionFile:
		// 	subSchema = unionSchema
		default:
			panic("Unknown file: " + filename)
		}
		err := t.Execute(&buf, subSchema)
		if err != nil {
			return err
		}
		out, err := format.Source(buf.Bytes())
		if err != nil {
			log.Println(err)
			out = []byte("// gofmt error: " + err.Error() + "\n\n" + buf.String())
		}
		buf.Reset()
		fmt.Println("writing", filename)
		err = os.WriteFile(filename, out, 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}

const (
	convertedTypeTmpl = `
{{- define "converted_type" -}}
  {{ .Name | title | convertPayloadType }}
{{- end }}`
	descriptionTmpl = `
{{- define "description" -}}
 {{.Description | clean | endSentence}}
{{- end }}`
	graphqlStructTagTmpl = `
{{- define "graphql_struct_tag" -}}` + "`" + `graphql:"
  {{- .Name | lowerFirst }}"` + "`" + `
{{- end }}`
	graphqlStructTagWithArgsTmpl = `
{{- define "graphql_struct_tag_with_args" -}}` + "`" + `graphql:"
  {{- .Name}}( {{- range $index, $element := .Args }}
    {{- if gt $index 0 }}, {{ end -}}
    {{- .Name}}: ${{.Name}}
  {{- end}})"` + "`" + `
{{- end }}`
	fragmentsTmpl = `
{{- define "fragments" -}}
  {{ if eq .Name "Check" }}{{ check_fragments }}
  {{ else if eq .Name "CustomActionsExternalAction" }}{{ custom_actions_ext_action_fragments }}
  {{ else if eq .Name "Integration" }}{{ integration_fragments }}
  {{- end }}
{{- end }}`
	nameToSingularTmpl = `
{{- define "name_to_singular" -}}
  {{- .Name | title | makeSingular }}
{{- end }}`
	typeCommentDescriptionTmpl = `
{{- define "type_comment_description" -}}
  // {{.Name | title}} {{ template "description" . }}
{{- end }}`
	fieldCommentDescriptionTmpl = `
{{- define "field_comment_description" -}}
  // {{ .Description | clean | fullSentence }}
{{- end }}`
)

// Filename -> Template.
var templates = map[string]*template.Template{
	// 	connectionFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{ if hasSuffix "Connection" .Name }}
	//       {{- template "object" . }}
	//     {{end}}
	//   {{- end}}
	// {{- end}}

	// {{- define "object" -}}
	// {{ if hasSuffix "Connection" .Name }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	//   Nodes []{{- if eq .Name "AncestorGroupsConnection"}}Group
	//           {{- else}}{{.Name | trimSuffix "Connection" | trimSuffix "V2" }} ` + "`" + `graphql:"nodes"` + "`" + `
	//           {{- end }}
	//   Edges []{{.Name | trimSuffix "Connection" }}Edge ` + "`" + `graphql:"edges"` + "`" + `
	// {{ range .Fields }} {{ if and (ne "edges" .Name) (ne "nodes" .Name) }}
	//     {{- .Name | title}} {{ template "converted_type" . }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	//   {{- end }}
	// {{- end }}
	// }
	// {{- end }}{{- end -}}
	//   `),
	enumFile: t(header + `
{{range .Types | sortByName}}{{if and (eq .Kind "ENUM") (not (internal .Name))}}
{{template "enum" .}}
{{end}}{{end}}


{{- define "enum" -}}
{{ template "type_comment_description" . }}
type {{.Name}} string

const (
{{- range .EnumValues }}
	{{$.Name}}{{.Name | enumIdentifier}} {{$.Name}} = {{.Name | quote}} {{ template "field_comment_description" . }}
{{- end }}
)
// All {{$.Name}} as []string
var All{{$.Name}} = []string {
	{{range .EnumValues}}string({{$.Name}}{{.Name | enumIdentifier}}),
	{{end}}
}
{{- end -}}
`),
	inputObjectFile: t(header + `
import "github.com/relvacode/iso8601"

{{range .Types | sortByName}}{{if and (eq .Kind "INPUT_OBJECT") (not (internal .Name))}}
{{ if and (not (hasPrefix "Campaign" .Name)) (not (hasPrefix "Group" .Name)) -}}
{{template "input_object" .}}
{{end}}{{end}}{{end}}

{{- define "input_object" -}}
{{ template "type_comment_description" . }}
type {{.Name}} struct { {{range .InputFields }}
  {{.Name | title}} {{if ne .Type.Kind "NON_NULL"}}*{{end -}}
    {{- if isListType .Name }}[]{{ end -}}
    {{- if and (hasSuffix "Id" .Name) (not (eq .Name "externalId")) }}ID
    {{- else if hasSuffix "Access" .Name }}IdentifierInput
    {{- else if eq .Name "predicates" }}FilterPredicateInput
    {{- else if eq .Name "tags" }}TagInput
    {{- else if eq .Name "members" }}TeamMembershipUserInput
    {{- else if eq .Name "contacts" }}ContactInput
    {{- else if eq .Type.Name "UserRole" }}UserRole
    {{- else if .Type.Name }}{{ template "converted_type" .Type }}
    {{- else }}{{ .Type.OfType.OfTypeName | convertPayloadType  }}{{ end -}} ` + "`" +
		`json:"{{.Name | lowerFirst }}{{if ne .Type.Kind "NON_NULL"}},omitempty{{end}}"` +
		` yaml:"{{.Name | lowerFirst }}{{if ne .Type.Kind "NON_NULL"}},omitempty{{end}}"` + `

  {{-  if and (not (hasSuffix "Input" .Type.Name)) (not (hasSuffix "Input" .Type.OfType.OfTypeName)) }} example:"
   {{- if isListType .Name }}[{{ end -}}
     {{ . | example_tag_value }}
   {{- if isListType .Name }}]{{ end -}}"{{- end}}` +
		"`" + `{{ template "field_comment_description" . }} {{if eq .Type.Kind "NON_NULL"}}(Required.){{else}}(Optional.){{end}}
  {{- end}}
}
{{- end -}}
`),
	interfacesFile: t(header + `
  import "github.com/relvacode/iso8601"

	{{range .Types | sortByName}}{{if and (eq .Kind "INTERFACE") (not (internal .Name))}}
	{{template "interface_object" .}}
	{{end}}{{end}}

	{{- define "interface_object" -}}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct { {{ add_special_interfaces_fields .Name }}
    {{ range .Fields }}{{ if not (skip_interface_field $.Name .Name) }}
	  {{.Name | title}} {{ . | get_field_type }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	{{- end }}{{ end }}

    {{ template "fragments" . }}
	}
	{{- end -}}
		`),
	// 	payloadFile: t(header + `
	// {{range .Types | sortByName}}{{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	// {{template "payload_object" .}}
	// {{- end}}{{- end}}

	// {{- define "payload_object" -}}
	// {{ if hasSuffix "Payload" .Name }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	// {{ range .Fields }}
	//   {{.Name | title}} {{ if isListType .Name }}[]{{- end -}}
	//     {{ template "converted_type" . }} {{ template "field_comment_description" . }}
	// {{- end }}
	// }
	// {{- end }}{{ end -}}
	// `),
	// NOTE: "account" == objectSchema.Types[0]
	// NOTE: "mutation" == objectSchema.Types[134]
	// NOTE: may have to use interfaceSchema to derive details for objects
	// 	queryFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{- if eq .Name "Account" }}
	//       {{ template "account_queries" . }}
	//     {{- end}}
	//   {{- end}}
	// {{- end}}

	// {{ define "account_queries" -}}
	//     {{- range .Fields }}
	// {{ template "type_comment_description" . }}
	// func (client *Client) {{ if isListType .Name }}List{{ .Name | title }}(input any) ({{ template "name_to_singular" . }}Connection, error) {
	//     {{- else }}Get{{ .Name | title }}(input any) ({{.Name | title}}, error) {
	//     {{end -}}
	//     var q struct {
	//       Account struct {
	//         {{ .Name | title }} {{ if isListType .Name }}{{ template "name_to_singular" . }}Connection
	//                             {{- else }}Get{{ template "name_to_singular" . }}{{end -}}` + "`" + `graphql:"{{.Name}}(input: $input)"` + "`" + `
	//       }
	//     }
	//     v := PayloadVariables{ {{ range .Args }}
	//       "{{.Name}}": input, {{ end}}
	//     }
	//     err := client.Query(&q, v, WithName("{{ template "name_to_singular" . }}{{ if isListType .Name }}List{{else}}Get{{end}}"))
	//     return &q.Account.{{ .Name | title }}, HandleErrors(err, nil)
	// }
	// {{- end}}{{- end}}

	// {{- define "object" -}}
	// {{ if not (hasSuffix "Payload" .Name) }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	//     {{ range .Fields }}
	//   {{.Name | title}} string  {{ template "field_comment_description" . }}
	//     {{- end }}
	// }
	// {{- end }}{{- end -}}
	// `),
	// 	mutationFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{- if eq .Name "Mutation" }}
	//       {{- template "mutation" .}}
	//     {{end}}
	//   {{- end}}
	// {{- end}}

	// {{ define "mutation" -}}
	// {{- range .Fields }}
	// // {{ .Name | title | renameMutation }} {{ template "description" . }}
	// func (client *Client) {{ .Name | title | renameMutation }}(
	//   {{- range $index, $element := .Args }}{{- if gt $index 0 }}, {{ end -}}
	//     {{- if eq "IdentifierInput" .Type.OfType.OfTypeName }}identifier string
	//     {{- else }}{{- .Name }} {{ with .Type.OfType.OfTypeName }}{{.}}{{else}}any{{end}}
	//     {{- end }}
	//   {{- end -}} ) (*{{.Name | title | renameMutationReturnType}}, error) {
	//     var m struct {
	//       Payload struct {
	//         {{ .Name | title | renameMutationReturnType}} {{ .Name | title | renameMutationReturnType}}
	//         Errors []OpsLevelErrors
	//       }{{ template "graphql_struct_tag_with_args" . }}
	//     }
	//     v := PayloadVariables{ {{ range .Args }}
	//       "{{.Name}}": {{- if eq "IdentifierInput" .Type.OfType.OfTypeName }}*NewIdentifier(identifier),
	//                    {{- else}}{{.Name}},{{ end }}
	//                            {{- end}}
	//     }
	//     err := client.Mutate(&m, v, WithName("{{ .Name | title }}"))
	//     return &m.Account.{{ .Name | title | renameMutationReturnType}}, HandleErrors(err, m.Payload.Errors)
	// }
	// {{- end}}
	// {{- end}}
	// `),
	objectFile: t(header + `
  import "github.com/relvacode/iso8601"

	{{range .Types | sortByName}}
	  {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	    {{- if eq .Name "Account" }}
	      {{- template "account_struct" . }}
	    {{- else}}{{template "object" .}}{{end}}
	  {{- end}}
	{{- end}}

	{{ define "account_struct" -}}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct { {{range .Fields }}
	  {{.Name | title}} *{{ if isListType .Name }}[]{{ end }}{{ template "converted_type" . }}  {{ template "field_comment_description" . }}
	 {{- end }}
	}
	{{- end }}

	{{- define "object" -}}
	{{ if and (and (not (hasSuffix "Payload" .Name)) (not (hasSuffix "Connection" .Name))) (not (hasSuffix "Edge" .Name)) }}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct {
	  {{ range .Fields -}}
	    {{ if not (len .Args) }}{{.Name | title}} {{ template "converted_type" . }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	    {{- end}}
	  {{ end -}}
	}
	{{- end }}{{- end -}}
		`),
	// 	scalarFile: t(header + `
	// import (
	// 	"encoding/base64"
	// 	"strconv"
	// 	"strings"
	// )

	// {{range .Types | sortByName}}{{if and (eq .Kind "SCALAR") (not (internal .Name))}}
	// {{template "scalar" .}}
	// {{end}}{{end}}

	// {{- define "scalar" -}}
	// {{ template "type_comment_description" . }}
	// type {{.Name}}
	// {{- if eq .Name "Boolean" }} bool
	// {{- else if eq .Name "Float" }} float64
	// {{- else if eq .Name "ID" }} string
	// {{- else if eq .Name "ISO8601DateTime" }} string
	// {{- else if eq .Name "Int" }} int
	// {{- else if eq .Name "JSON" }} map[string]any
	// {{- else if eq .Name "JSONSchema" }} map[string]any
	// {{- else if eq .Name "String" }} string
	// {{- end -}}{{end}}

	// func NewID(id ...string) *ID {
	// 	var output ID
	// 	if len(id) == 1 {
	// 		output = ID(id[0])
	// 	}
	// 	return &output
	// }

	// func (s ID) GetGraphQLType() string { return "ID" }

	// func (s *ID) MarshalJSON() ([]byte, error) {
	// 	if *s == "" {
	// 		return []byte("null"), nil
	// 	}
	// 	return []byte(strconv.Quote(string(*s))), nil
	// }

	// type Identifier struct {
	// 	Id      ID       ` + "`" + `graphql:"id"` + "`" + `
	// 	Aliases []string ` + "`" + `graphql:"aliases"` + "`" + `
	// }

	// func NewIdentifier(value string) *IdentifierInput {
	// 	if IsID(value) {
	// 		return &IdentifierInput{
	// 			Id: NewID(value),
	// 		}
	// 	}
	// 	return &IdentifierInput{
	// 		Alias: NewString(value),
	// 	}
	// }

	// func NewIdentifierArray(values []string) []IdentifierInput {
	// 	output := []IdentifierInput{}
	// 	for _, value := range values {
	// 		output = append(output, *NewIdentifier(value))
	// 	}
	// 	return output
	// }

	// func IsID(value string) bool {
	// 	decoded, err := base64.RawURLEncoding.DecodeString(value)
	// 	if err != nil {
	// 		return false
	// 	}
	// 	return strings.HasPrefix(string(decoded), "gid://")
	// }

	// func NewString(value string) *string {
	// 	return &value
	// }`),
	// 	unionFile: t(header + `
	// {{range .Types | sortByName}}{{if and (eq .Kind "UNION") (not (internal .Name))}}
	// {{template "union_object" .}}
	// {{end}}{{end}}

	// {{- define "union_object" -}}
	// // Union{{.Name}} {{ template "description" . }}
	// type Union{{.Name}} interface { {{range .PossibleTypes }}
	//
	//	    {{.Name}}Fragment() {{.Name}}Fragment{{end}}
	//	}
	//
	// {{- end -}}
	//
	//	`),
}

func t(text string) *template.Template {
	// typeString returns a string representation of GraphQL type t.
	var typeString func(t map[string]interface{}) string
	typeString = func(t map[string]interface{}) string {
		switch t["kind"] {
		case "NON_NULL":
			s := typeString(t["ofType"].(map[string]interface{}))
			if !strings.HasPrefix(s, "*") {
				panic(fmt.Errorf("nullable type %q doesn't begin with '*'", s))
			}
			return s[1:] // Strip star from nullable type to make it non-null.
		case "LIST":
			return "*[]" + typeString(t["ofType"].(map[string]interface{}))
		default:
			return "*" + t["name"].(string)
		}
	}

	genTemplate := template.New("")
	genTemplate.Funcs(templFuncMap)
	genTemplate.Funcs(sprig.TxtFuncMap())
	genTemplate.Funcs(template.FuncMap{"type": typeString})
	genTemplate.Parse(convertedTypeTmpl)
	genTemplate.Parse(descriptionTmpl)
	genTemplate.Parse(fieldCommentDescriptionTmpl)
	genTemplate.Parse(fragmentsTmpl)
	genTemplate.Parse(graphqlStructTagTmpl)
	genTemplate.Parse(graphqlStructTagWithArgsTmpl)
	genTemplate.Parse(nameToSingularTmpl)
	genTemplate.Parse(typeCommentDescriptionTmpl)
	return template.Must(genTemplate.Parse(text))
}

func makeSingular(s string) string {
	value := strings.ToLower(s)
	if strings.HasSuffix(value, "ies") {
		return strings.ReplaceAll(s, "ies", "y")
	}
	if isPlural(s) {
		return strings.TrimSuffix(s, "s")
	}
	return s
}

func convertPayloadType(s string) string {
	switch s {
	case "Boolean":
		return "bool"
	case "Int":
		return "int"
	case "String":
		return "string"
	case "ISO8601DateTime":
		return "iso8601.Time"
	case "":
		return "string"
	}

	value := strings.ToLower(s)
	if strings.HasSuffix(value, "id") {
		return "ID"
	} else if slices.Contains(knownBoolsByName, value) {
		return "bool"
	} else if slices.Contains(knownIsoTimeByName, value) {
		return "iso8601.Time"
	}
	for k, v := range knownTypeMappings {
		if value == k {
			return v
		}
	}
	for _, knownStringTypeSuffix := range stringTypeSuffixes {
		if strings.HasSuffix(value, knownStringTypeSuffix) {
			return "string"
		}
	}
	return makeSingular(s)
}

// TODO fix up later
func renameMutationReturnType(s string) string {
	create := "Create"
	delete := "Delete"
	update := "Update"
	if strings.HasSuffix(s, create) {
		s = strings.TrimSuffix(s, create)
	} else if strings.HasSuffix(s, delete) {
		s = strings.TrimSuffix(s, delete)
	} else if strings.HasSuffix(s, update) {
		s = strings.TrimSuffix(s, update)
	}
	return s
}

// TODO fix up later
func renameMutation(s string) string {
	create := "Create"
	delete := "Delete"
	update := "Update"
	if strings.HasSuffix(s, create) {
		s = strings.TrimSuffix(s, create)
		s = fmt.Sprintf("%s%s", create, s)
	} else if strings.HasSuffix(s, delete) {
		s = strings.TrimSuffix(s, delete)
		s = fmt.Sprintf("%s%s", delete, s)
	} else if strings.HasSuffix(s, update) {
		s = strings.TrimSuffix(s, update)
		s = fmt.Sprintf("%s%s", update, s)
	}
	return s
}

func isPlural(s string) bool {
	value := strings.ToLower(s)
	// Examples: "alias", "address", "status", "levels", "responsibilities"
	if value == "notes" || value == "days" || value == "headers" ||
		strings.HasSuffix(value, "ies") ||
		strings.HasSuffix(value, "ias") ||
		strings.HasSuffix(value, "ls") ||
		(!strings.HasSuffix(value, "access") && strings.HasSuffix(value, "ss")) ||
		strings.HasSuffix(value, "us") {
		return false
	}
	if strings.HasSuffix(value, "s") {
		return true
	}
	return false
}

func getFragmentWithStructTag(fragmentNames ...string) string {
	output := make([]string, len(fragmentNames))
	for _, fragment := range fragmentNames {
		fragmentWithTag := fragment + "`" + `graphql:"... on ` + strings.TrimSuffix(fragment, "Fragment") + "\"`"
		output = append(output, fragmentWithTag)
	}
	return strings.Join(output, "\n")
}

// Check
func fragmentsForCheck() string {
	checkFragments := []string{
		"AlertSourceUsageCheckFragment",
		"CustomEventCheckFragment",
		"HasRecentDeployCheckFragment",
		"ManualCheckFragment",
		"RepositoryFileCheckFragment",
		"RepositoryGrepCheckFragment",
		"RepositorySearchCheckFragment",
		"ServiceOwnershipCheckFragment",
		"ServicePropertyCheckFragment",
		"TagDefinedCheckFragment",
		"ToolUsageCheckFragment",
		"HasDocumentationCheckFragment",
	}
	return getFragmentWithStructTag(checkFragments...)
}

func fragmentsForIntegration() string {
	integrationFragments := []string{
		"AwsIntegrationFragment",
		"NewRelicIntegrationFragment",
	}

	stuff := getFragmentWithStructTag(integrationFragments...)
	return strings.Replace(stuff, "AwsIntegration", "AWSIntegration", 1)
}

func fragmentsForCustomActionsExtAction() string {
	return getFragmentWithStructTag("CustomActionsWebhookAction")
}

var templFuncMap = template.FuncMap{
	"internal":                            func(s string) bool { return strings.HasPrefix(s, "__") },
	"quote":                               strconv.Quote,
	"join":                                strings.Join,
	"check_fragments":                     fragmentsForCheck,
	"custom_actions_ext_action_fragments": fragmentsForCustomActionsExtAction,
	"integration_fragments":               fragmentsForIntegration,
	"get_field_type":                      getFieldType,
	"add_special_interfaces_fields":       addSpecialInterfacesFields,
	"skip_interface_field":                skipInterfaceField,
	"example_tag_value":                   getExampleValue,
	"isListType":                          isPlural,
	"renameMutation":                      renameMutation,
	"renameMutationReturnType":            renameMutationReturnType,
	"convertPayloadType":                  convertPayloadType,
	"makeSingular":                        makeSingular,
	"lowerFirst": func(value string) string {
		for i, v := range value {
			return string(unicode.ToLower(v)) + value[i+1:]
		}
		return value
	},
	"sortByName": func(types []GraphQLTypes) []GraphQLTypes {
		sort.Slice(types, func(i, j int) bool {
			ni := types[i].Name
			nj := types[j].Name
			return ni < nj
		})
		return types
	},
	"inputObjects": func(types []interface{}) []string {
		var names []string
		for _, t := range types {
			t := t.(map[string]interface{})
			if t["kind"].(string) != "INPUT_OBJECT" {
				continue
			}
			names = append(names, t["name"].(string))
		}
		sort.Strings(names)
		return names
	},
	"identifier":     func(name string) string { return ident.ParseLowerCamelCase(name).ToMixedCaps() },
	"enumIdentifier": func(name string) string { return ident.ParseScreamingSnakeCase(name).ToMixedCaps() },
	"clean":          func(s string) string { return strings.Join(strings.Fields(s), " ") },
	"endSentence": func(s string) string {
		if len(s) == 0 {
			// Do nothing.
			return ""
		}

		s = strings.ToLower(s[0:1]) + s[1:]
		switch {
		default:
			s = "represents " + s
		case strings.HasPrefix(s, "autogenerated "):
			s = "is an " + s
		case strings.HasPrefix(s, "specifies "):
			// Do nothing.
		}
		if !strings.HasSuffix(s, ".") {
			s += "."
		}
		return s
	},
	"fullSentence": func(s string) string {
		if !strings.HasSuffix(s, ".") {
			s += "."
		}
		return s
	},
}

func addSpecialInterfacesFields(interfaceName string) string {
	switch interfaceName {
	case "CustomActionsExternalAction":
		return "CustomActionsId"
	case "Integration":
		return "IntegrationId"
	}
	return ""
}

func skipInterfaceField(interfaceName, fieldName string) bool {
	if interfaceName == "Check" {
		switch fieldName {
		case "campaign", "rawNotes", "url":
			return true
		}
	} else if interfaceName == "CustomActionsExternalAction" {
		switch fieldName {
		case "id", "aliases":
			return true
		}
	} else if interfaceName == "Integration" {
		switch fieldName {
		case "id", "name", "type":
			return true
		}
	}
	return false
}

func getFieldType(inputField GraphQLField) string {
	lowercaseFieldName := strings.ToLower(inputField.Name)
	switch {
	case "id" == lowercaseFieldName:
		return "ID"
	case "aliases" == lowercaseFieldName:
		return "[]string"
	case "owner" == lowercaseFieldName:
		return "CheckOwner"
	case "type" == lowercaseFieldName:
		return "CheckType"
	case slices.Contains(knownTypeIsName, lowercaseFieldName):
		return strings.ToUpper(inputField.Name[0:1]) + inputField.Name[1:]
	case slices.Contains(knownBoolsByName, lowercaseFieldName):
		return "bool"
	case slices.Contains(knownIsoTimeByName, lowercaseFieldName):
		return "iso8601.Time"
	}
	return "string"
}

func getExampleValueByFieldName(inputField GraphQLInputValue) string {
	mapFieldTypeToExampleValue := map[string]string{
		"DocumentSubtype":      "openapi",
		"Address":              "support@company.com",
		"Id":                   "Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk",
		"Definition":           "example_definition",
		"Template":             `{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}`,
		"Name":                 "example_name",
		"Language":             "example_language",
		"Alias":                "example_alias",
		"Description":          "example_description",
		"Email":                "first.last@domain.com",
		"Data":                 "example_data",
		"Note":                 "example_note",
		"IamRole":              "example_role",
		"DisplayType":          "example_type",
		"HttpMethod":           "GET",
		"Notes":                "example_notes",
		"Value":                "example_value",
		"Product":              "example_product",
		"Framework":            "example_framework",
		"Url":                  "john.doe@example.com",
		"BaseDirectory":        "/home/opslevel.yaml",
		"ExternalUrl":          "https://google.com",
		"Responsibilities":     "example description of responsibilities",
		"Environment":          "environment that tool belongs to",
		"Arg":                  "example_arg",
		"Extensions":           "'go', 'py', 'rb'",
		"Paths":                "'/usr/local/bin', '/home/opslevel'",
		"Ids":                  "'Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk', 'Z2lkOi8vc2VydmljZS85ODc2NTQzMjE'",
		"TagKeys":              "'tag_key1', 'tag_key2'",
		"Selector":             "example_selector",
		"Condition":            "example_condition",
		"Message":              "example_message",
		"RequireContactMethod": "false",
		"Identifier":           "example_identifier",
		"DocumentType":         "api",
	}
	for k, v := range mapFieldTypeToExampleValue {
		if k == inputField.Name ||
			strings.ToLower(k[:1])+k[1:] == inputField.Name ||
			strings.HasSuffix(inputField.Name, k) {
			return v
		}
	}
	return ""
}

func getExampleValueByFieldType(inputField GraphQLInputValue) string {
	mapFieldTypeToExampleValue := map[string]string{
		"Time":                        "2024-01-05T01:00:00.000Z",
		"FrequencyTimeScale":          "week",
		"ContactType":                 "slack",
		"AlertSourceTypeEnum":         "pagerduty",
		"AliasOwnerTypeEnum":          "scorecard",
		"BasicTypeEnum":               "does_not_equal",
		"ConnectiveEnum":              "or",
		"CustomActionsEntityTypeEnum": "GLOBAL",
		"CustomActionsHttpMethodEnum": "GET",
		"CustomActionsTriggerDefinitionAccessControlEnum": "service_owners",
		"RelationshipTypeEnum":                            "depends_on",
		"PredicateKeyEnum":                                "filter_id",
		"PredicateTypeEnum":                               "satisfies_jq_expression",
		"ServicePropertyTypeEnum":                         "language",
		"UsersFilterEnum":                                 "last_sign_in_at",
		"UserRole":                                        "admin",
		"ToolCategory":                                    "api_documentation",
	}
	for k, v := range mapFieldTypeToExampleValue {
		if inputFieldMatchesType(inputField, k) {
			return v
		}
	}
	return ""
}

func inputFieldMatchesType(inputField GraphQLInputValue, fieldType string) bool {
	if fieldType == inputField.Type.Name ||
		fieldType == inputField.Type.OfType.OfTypeName ||
		strings.ToLower(fieldType) == inputField.Type.Name ||
		strings.HasSuffix(inputField.Type.Name, fieldType) ||
		strings.HasSuffix(inputField.Type.OfType.OfTypeName, fieldType) {
		return true
	}
	return false
}

func inputFieldNameMatchesName(inputField GraphQLInputValue, fieldName string) bool {
	if fieldName == inputField.Name ||
		strings.ToLower(fieldName[:1])+fieldName[1:] == inputField.Name ||
		strings.HasSuffix(inputField.Name, fieldName) {
		return true
	}
	return false
}

func getExampleValue(inputField GraphQLInputValue) string {
	switch {
	case inputFieldMatchesType(inputField, "Boolean"):
		return "false"
	case inputFieldMatchesType(inputField, "Int"):
		return "3"
	case inputFieldMatchesType(inputField, "JSON"):
		return `{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}`
	}

	if valueByName := getExampleValueByFieldName(inputField); valueByName != "" {
		return valueByName
	}
	if valueByType := getExampleValueByFieldType(inputField); valueByType != "" {
		return valueByType
	}

	switch {
	case inputFieldNameMatchesName(inputField, "Role"):
		return "example_role"
	case inputFieldNameMatchesName(inputField, "Key"):
		return "XXX_example_key_XXX"
	case inputFieldNameMatchesName(inputField, "Type"):
		return "example_type"
	case inputFieldNameMatchesName(inputField, "Method"):
		return "example_method"
	case inputFieldMatchesType(inputField, "Enum"):
		return "NEW_ENUM_SET_DEFAULT"
	}
	return ""
}
