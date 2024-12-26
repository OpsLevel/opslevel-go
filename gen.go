//go:build ignore
// +build ignore

package main

import (
	"bytes"
	_ "embed"
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
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/types"
	"github.com/hasura/go-graphql-client/ident"
	"github.com/opslevel/opslevel-go/v2024"
)

const (
	connectionFile string = "connection.go"
	interfacesFile string = "interfaces.go"
	objectFile     string = "object.go"
	queryFile      string = "query.go"
	mutationFile   string = "mutation.go"
	payloadFile    string = "payload.go"
	// scalarFile      string = "scalar.go" // NOTE: probably not useful
	// unionFile       string = "union.go" // NOTE: probably not useful
)

var knownTypeIsName = []string{
	"category",
	"filter",
	"frequencytimescale",
	"lifecycle",
	"level",
	"tier",
	"timestamps",
}

var knownBoolsByName = []string{
	"affectsoverallservicelevels",
	"casesensitive",
	"enabled",
	"forked",
	"hasnextpage",
	"haspreviouspage",
	"locked",
	"lockedfromgraphqlmodification",
	"ownerlocked",
	"passingchecks",
	"published",
	"servicecount",
	"totalchecks",
	"visible",
}

var knownIntsByName = []string{
	"hiddencount",
	"index",
	"frequencyvalue",
}

var knownIsoTimeByName = []string{
	"archivedat",
	"createdat",
	"createdon",
	"enableon",
	"installedat",
	"lastownerchangedat",
	"lastsyncedat",
	"startingdate",
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
	"notes",
	"processedat",
	"queryparams",
	"role",
	"status",
	"updatedat",
	"userdeletepayload",
	"url",
	"yaml",
}

var enumExamples = map[string]string{
	"AlertSourceStatusTypeEnum":                       opslevel.AllAlertSourceStatusTypeEnum[0],
	"ApiDocumentSourceEnum":                           opslevel.AllApiDocumentSourceEnum[0],
	"AlertSourceTypeEnum":                             opslevel.AllAlertSourceTypeEnum[0],
	"AliasOwnerTypeEnum":                              opslevel.AllAliasOwnerTypeEnum[0],
	"BasicTypeEnum":                                   opslevel.AllBasicTypeEnum[0],
	"CampaignFilterEnum":                              opslevel.AllCampaignFilterEnum[0],
	"CampaignReminderChannelEnum":                     opslevel.AllCampaignReminderChannelEnum[0],
	"CampaignReminderFrequencyUnitEnum":               opslevel.AllCampaignReminderFrequencyUnitEnum[0],
	"CampaignReminderTypeEnum":                        opslevel.AllCampaignReminderTypeEnum[0],
	"CampaignServiceStatusEnum":                       opslevel.AllCampaignServiceStatusEnum[0],
	"CampaignSortEnum":                                opslevel.AllCampaignSortEnum[0],
	"CampaignStatusEnum":                              opslevel.AllCampaignStatusEnum[0],
	"CheckCodeIssueConstraintEnum":                    opslevel.AllCheckCodeIssueConstraintEnum[0],
	"CheckResultStatusEnum":                           opslevel.AllCheckResultStatusEnum[0],
	"CheckStatus":                                     opslevel.AllCheckStatus[0],
	"CheckType":                                       opslevel.AllCheckType[0],
	"CodeIssueResolutionTimeUnitEnum":                 opslevel.AllCodeIssueResolutionTimeUnitEnum[0],
	"ConnectiveEnum":                                  opslevel.AllConnectiveEnum[0],
	"ContactType":                                     opslevel.AllContactType[0],
	"CustomActionsEntityTypeEnum":                     opslevel.AllCustomActionsEntityTypeEnum[0],
	"CustomActionsHttpMethodEnum":                     opslevel.AllCustomActionsHttpMethodEnum[0],
	"CustomActionsTriggerDefinitionAccessControlEnum": opslevel.AllCustomActionsTriggerDefinitionAccessControlEnum[0],
	"CustomActionsTriggerEventStatusEnum":             opslevel.AllCustomActionsTriggerEventStatusEnum[0],
	"DayOfWeekEnum":                                   opslevel.AllDayOfWeekEnum[0],
	"EventIntegrationEnum":                            opslevel.AllEventIntegrationEnum[0],
	"FrequencyTimeScale":                              opslevel.AllFrequencyTimeScale[0],
	"HasDocumentationSubtypeEnum":                     opslevel.AllHasDocumentationSubtypeEnum[0],
	"HasDocumentationTypeEnum":                        opslevel.AllHasDocumentationTypeEnum[0],
	"PackageConstraintEnum":                           opslevel.AllPackageConstraintEnum[0],
	"PackageManagerEnum":                              opslevel.AllPackageManagerEnum[0],
	"PayloadFilterEnum":                               opslevel.AllPayloadFilterEnum[0],
	"PayloadSortEnum":                                 opslevel.AllPayloadSortEnum[0],
	"PredicateKeyEnum":                                opslevel.AllPredicateKeyEnum[0],
	"PredicateTypeEnum":                               opslevel.AllPredicateTypeEnum[0],
	"PropertyDefinitionDisplayTypeEnum":               opslevel.AllPropertyDefinitionDisplayTypeEnum[0],
	"PropertyDisplayStatusEnum":                       opslevel.AllPropertyDisplayStatusEnum[0],
	"RelatedResourceRelationshipTypeEnum":             opslevel.AllRelatedResourceRelationshipTypeEnum[0],
	"RelationshipTypeEnum":                            opslevel.AllRelationshipTypeEnum[0],
	"RepositoryVisibilityEnum":                        opslevel.AllRepositoryVisibilityEnum[0],
	"ResourceDocumentStatusTypeEnum":                  opslevel.AllResourceDocumentStatusTypeEnum[0],
	"ScorecardSortEnum":                               opslevel.AllScorecardSortEnum[0],
	"ServicePropertyTypeEnum":                         opslevel.AllServicePropertyTypeEnum[0],
	"ServiceSortEnum":                                 opslevel.AllServiceSortEnum[0],
	"SnykIntegrationRegionEnum":                       opslevel.AllSnykIntegrationRegionEnum[0],
	"TaggableResource":                                opslevel.AllTaggableResource[0],
	"ToolCategory":                                    opslevel.AllToolCategory[0],
	"UserRole":                                        opslevel.AllUserRole[0],
	"UsersFilterEnum":                                 opslevel.AllUsersFilterEnum[0],
	"UsersInviteScopeEnum":                            opslevel.AllUsersInviteScopeEnum[0],
	"VaultSecretsSortEnum":                            opslevel.AllVaultSecretsSortEnum[0],
}

var listExamples = map[string]string{
	"channels":           "[]",
	"checkIds":           "['Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk', 'Z2lkOi8vc2VydmljZS85ODc2NTQzMjE']",
	"checkIdsToCopy":     "[]",
	"checksToPromote":    "[]",
	"contacts":           "[]",
	"daysOfWeek":         "[]",
	"dependsOn":          "[]",
	"dependencyOf":       "[]",
	"extendedTeamAccess": "[]",
	"fileExtensions":     "['go', 'py', 'rb']",
	"filePaths":          "['/usr/local/bin', '/home/opslevel']",
	"issueType":          "['bug', 'error']",
	"members":            "[]",
	"ownershipTagKeys":   "['tag_key1', 'tag_key2']",
	"predicates":         "[]",
	"properties":         "[]",
	"regionOverride":     "['us-east-1', 'eu-west-1']",
	"reminderTypes":      "[]",
	"severity":           "['sev1', 'sev2']",
	"tags":               "[]",
	"teamIds":            "[]",
	"teams":              "[]",
	"users":              "[]",
}

var scalarExamples = map[string]string{
	"Boolean":         "false",
	"ID":              "Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk",
	"Int":             `"3"`,
	"ISO8601DateTime": "2025-01-05T01:00:00.000Z",
	"Float":           "4.2069",
	"String":          "example_value",
	"JSON":            `{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}`,
	"JSONSchema":      `SCHEMA_TBD`,
	"JsonString":      "JSON_TBD",
}

var knownTypeMappings = map[string]string{
	"data":                           "JSON",
	"deletedmembers":                 "User",
	"filteredcount":                  "int",
	"headers":                        "JSON",
	"highlights":                     "JSON",
	"memberships":                    "TeamMembership",
	"notupdatedrepositories":         "RepositoryOperationErrorPayload",
	"promotedchecks":                 "Check",
	"relationship":                   "RelationshipType",
	"teamsbeingnotified":             "CampaignSendReminderOutcomeTeams",
	"teamsbeingnotifiedcount":        "int",
	"teamsmissingcontactmethod":      "int",
	"teamsmissingcontactmethodcount": "int",
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

//go:embed schema.graphql
var graphqlSchema string

func main() {
	flag.Parse()

	opts := []graphql.SchemaOpt{
		graphql.UseStringDescriptions(),
	}
	schema := graphql.MustParseSchema(graphqlSchema, nil, opts...)
	schemaAst := schema.ASTSchema()

	inputObjects := map[string]*types.InputObject{}
	objects := map[string]*types.ObjectTypeDefinition{}
	enums := map[string]*types.EnumTypeDefinition{}
	interfaces := map[string]*types.InterfaceTypeDefinition{}
	unions := map[string]*types.Union{}
	scalars := map[string]*types.ScalarTypeDefinition{}
	for name, graphqlType := range schemaAst.Types {
		switch v := graphqlType.(type) {
		case *types.EnumTypeDefinition:
			enums[name] = v
		case *types.InputObject:
			inputObjects[name] = v
		case *types.InterfaceTypeDefinition:
			interfaces[name] = v
		case *types.ObjectTypeDefinition:
			objects[name] = v
		case *types.ScalarTypeDefinition:
			scalars[name] = v
		case *types.Union:
			unions[name] = v
		default:
			panic(fmt.Errorf("Unknown GraphQL type: %v", v))
		}
	}
	// genEnums2(enums)
	genEnums(schemaAst.Enums)
	genInputObjects(inputObjects)

	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func sortedMapKeys[T any](schemaMap map[string]T) []string {
	sortedNames := make([]string, 0, len(schemaMap))
	for k := range schemaMap {
		sortedNames = append(sortedNames, k)
	}
	slices.Sort(sortedNames)
	return sortedNames
}

func genEnums2(schemaEnums map[string]*types.EnumTypeDefinition) {
	var buf bytes.Buffer
	buf.WriteString(header + "\n\n")

	enumTmpl := template.New("enum")
	enumTmpl.Funcs(sprig.TxtFuncMap())
	enumTmpl.Funcs(templFuncMap)
	template.Must(enumTmpl.ParseFiles("./templates/enum.tpl"))

	for _, enumName := range sortedMapKeys(schemaEnums) {
		if err := enumTmpl.ExecuteTemplate(&buf, "enum", schemaEnums[enumName]); err != nil {
			panic(err)
		}
	}
	out, err := format.Source(buf.Bytes())
	err = os.WriteFile("enum.go", out, 0o644)
	if err != nil {
		log.Fatalln(err)
	}
}

func genEnums(schemaEnums []*types.EnumTypeDefinition) {
	var buf bytes.Buffer

	buf.WriteString(header + "\n\n")

	tmpl := template.New("enum")
	tmpl.Funcs(sprig.TxtFuncMap())
	tmpl.Funcs(templFuncMap)
	template.Must(tmpl.ParseFiles("./templates/enum.tpl"))

	for _, enum := range schemaEnums {
		if err := tmpl.ExecuteTemplate(&buf, "enum", enum); err != nil {
			panic(err)
		}
	}
	out, err := format.Source(buf.Bytes())
	err = os.WriteFile("enum.go", out, 0o644)
	if err != nil {
		panic(err)
	}
}

func genInputObjects(inputObjects map[string]*types.InputObject) {
	var buf bytes.Buffer
	buf.WriteString(header + "\n\nimport \"github.com/relvacode/iso8601\"\n")

	tmpl := template.New("inputs")
	tmpl.Funcs(sprig.TxtFuncMap())
	tmpl.Funcs(templFuncMap)
	template.Must(tmpl.ParseFiles("./templates/inputObjects.tpl"))

	for _, enumName := range sortedMapKeys(inputObjects) {
		if err := tmpl.ExecuteTemplate(&buf, "inputs", inputObjects[enumName]); err != nil {
			panic(err)
		}
	}
	err := os.WriteFile("input.go", buf.Bytes(), 0o644)
	if err != nil {
		panic(err)
	}
}

func getRootSchema() (*GraphQLSchema, error) {
	visibility, ok := os.LookupEnv("GRAPHQL_VISIBILITY")
	if !ok {
		visibility = "public"
	}
	token, ok := os.LookupEnv("OPSLEVEL_API_TOKEN")
	if !ok {
		return nil, fmt.Errorf("OPSLEVEL_API_TOKEN environment variable not set")
	}
	client := opslevel.NewGQLClient(opslevel.SetAPIToken(token), opslevel.SetAPIVisibility(visibility))
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

	inputObjectSchema := GraphQLSchema{}
	interfaceSchema := GraphQLSchema{}
	objectSchema := GraphQLSchema{}
	scalarSchema := GraphQLSchema{}
	unionSchema := GraphQLSchema{}
	for _, t := range schema.Types {
		switch t.Kind {
		case "ENUM":
			continue
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
		case connectionFile:
			subSchema = objectSchema
		case interfacesFile:
			subSchema = interfaceSchema
		case objectFile:
			subSchema = objectSchema
		case mutationFile:
			subSchema = objectSchema
		case payloadFile:
			subSchema = objectSchema
		case queryFile:
			subSchema = objectSchema
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
  {{- if eq .Name "Check" }}{{ check_fragments }}
  {{- else if eq .Name "CustomActionsExternalAction" }}{{ custom_actions_ext_action_fragments }}
  {{- else if eq .Name "Integration" }}{{ integration_fragments }}
  {{ end }}
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
	connectionFile: t(header + `
	{{range .Types | sortByName}}
	  {{if and (not (skip_if_campaign_or_group .Name)) (not (internal .Name)) }}
	    {{ if hasSuffix "Connection" .Name }}
	      {{- template "connection_object" . }}
	    {{ else if hasSuffix "Edge" .Name }}
	      {{- template "edge_object" . }}
	    {{end}}
	  {{- end}}
	{{- end}}

	{{- define "connection_object" -}}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct {
    Nodes []{{.Name | trimSuffix "Connection" | trimSuffix "V2" | makeSingular }} ` +
		"`" + `graphql:"nodes"` + "`" + ` // A list of nodes.
    Edges []{{.Name | trimSuffix "Connection" }}Edge ` + "`" + `graphql:"edges"` +
		"`" + ` // A list of edges.
	{{ range .Fields }}
      {{- if and (ne "edges" .Name) (ne "nodes" .Name) }}
	    {{ .Name | title}} {{ template "converted_type" . }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	  {{- end }}
	{{- end }}
	}
	{{- end }}

  {{- define "edge_object" -}}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct {
    {{ range .Fields }}
	    {{ .Name | title }} {{ if eq .Name "node"}}{{ $.Name | trimSuffix "Edge" | trimSuffix "V2" | makeSingular }}
      {{- else }}{{ template "converted_type" . }}
      {{- end }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
    {{- end }}
	}
	{{- end }}
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
	  {{.Name | title}} {{ get_input_field_type . }} {{ if eq .Name "notes" }}` + "`" + `graphql:"notes: rawNotes"` + "`" + `{{ else }}{{ template "graphql_struct_tag" . }}
    {{- end}} {{ template "field_comment_description" . }}
	{{- end }}{{ end }}
    {{ template "fragments" . -}}
	}
	{{- end -}}
		`),
	payloadFile: t(header + `
  import "github.com/relvacode/iso8601"

	{{range .Types | sortByName}}{{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	{{template "payload_object" .}}
	{{- end}}{{- end}}

	{{- define "payload_object" -}}
	{{ if and (hasSuffix "Payload" .Name) (not (skip_query .Name)) }}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct {
	{{ range .Fields }}
	  {{.Name | title}} {{ if isListType .Name }}[]{{- end -}}
      {{- if eq .Name "targetCategory" -}}Category
      {{- else }}{{ template "converted_type" . }}
      {{- end }} {{ template "field_comment_description" . }}
	{{- end }}
	}
	{{- end }}{{ end -}}
	`),
	// NOTE: "account" == objectSchema.Types[0]
	// NOTE: "mutation" == objectSchema.Types[134]
	queryFile: t(header + `
  import "fmt"

	{{range .Types | sortByName}}
	  {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	    {{- if eq .Name "Account" }}
	      {{- template "account_queries" . }}
	    {{- else }}
	      {{- template "non_account_queries" . }}
	    {{- end}}
	  {{- end}}
	{{- end}}

	{{ define "account_queries" -}}
	    {{- range .Fields }} {{- if and (len .Args) (not (skip_query .Name)) }}
  // {{ if gt (len .Args) 3 -}} List {{- else -}} Get {{- end -}}
     {{- .Name | title}} {{ .Description | clean | endSentence }}
	func (client *Client) {{ if gt (len .Args) 3 }}List{{ .Name | title | makePlural }}(variables *PayloadVariables) (*
                                      {{- if eq .Name "customActionsExternalActions" }}{{ .Name | title  }}Connection, error) {
                                      {{- else}}{{ .Name | title | makeSingular }}Connection, error) {
                                      {{- end -}}
                        {{- else -}}              Get{{ .Name | title }}(value string) (*{{.Name | title | trimSuffix "sVaultsSecret" }}, error) {
                        {{- end -}}
	    var q struct {
	      Account struct {
          {{ if gt (len .Args) 3 -}}
            {{- .Name | title | makePlural }} {{ if eq .Name "customActionsExternalActions" -}}{{ .Name | title  }}{{ else }}{{ .Name | title | makeSingular }}{{ end }}Connection {{ template "graphql_struct_tag_with_args" . }}
          {{- else -}}
            {{- .Name | title  }} {{ .Name | title | makeSingular | trimSuffix "sVaultsSecret" }} {{ template "graphql_struct_tag_with_args" . }}
          {{- end -}}
	      }
	    }
      {{- if gt (len .Args) 3 }}
    	if variables == nil {
        variables = client.InitialPageVariablesPointer()
      }
      {{ else }}
	    v := PayloadVariables{ {{ range .Args }}
	      "{{.Name}}": value, {{ end}}
	    }
      {{- end }}

      {{ if gt (len .Args) 3 -}}
      if err := client.Query(&q, *variables, WithName("{{ template "name_to_singular" . }}List")); err != nil {
        return nil, err
      }

      for q.Account.{{ .Name | title }}.PageInfo.HasNextPage {
        (*variables)["after"] = q.Account.{{ .Name | title }}.PageInfo.End
        resp, err := client.List{{ .Name | title }}(variables)
        if err != nil {
          return nil, err
        }
        q.Account.{{ .Name | title }}.Nodes = append(q.Account.{{ .Name | title }}.Nodes, resp.Nodes...)
        q.Account.{{ .Name | title }}.PageInfo = resp.PageInfo
        q.Account.{{ .Name | title }}.TotalCount += resp.TotalCount
      }
	    return &q.Account.{{ .Name | title }}, nil
      {{ else }}
	    err := client.Query(&q, v, WithName("{{ template "name_to_singular" . }}{{ if isListType .Name }}List{{else}}Get{{end}}"))
	    return &q.Account.{{ .Name | title }}, HandleErrors(err, nil)
      {{- end }}
	}
  {{end}}{{- end}}{{- end}}

	{{ define "non_account_queries" -}}
    {{- range .Fields }} {{- if and (len .Args) (not (skip_query $.Name)) }}
	// {{ if gt (len .Args) 3 }}List{{- else }}Get{{ end }}{{.Name | title}} {{ .Description | clean | endSentence }}
	func ( {{- $.Name | first_char_lowered }} *{{ $.Name | title | makeSingular }})

    {{- if gt (len .Args) 3 }}List{{ .Name | title }}(client *Client, variables *PayloadVariables) (*
    {{- if or (hasPrefix "ancestor" .Name) (hasPrefix "child" .Name) }} {{- $.Name }}Connection, error
    {{- else if hasPrefix "descendant" .Name }}{{ .Name | title | makeSingular | trimPrefix "Descendant" }}Connection, error
    {{- else }}{{ if eq .Name "memberships" }}Team{{end}}{{ .Name | title | makeSingular | trimPrefix "Child" }}Connection, error
    {{- end -}} ) {
      if {{ $.Name | first_char_lowered }}.Id == "" {
        return nil, fmt.Errorf("Unable to get {{ .Name | title }}, invalid {{ $.Name | lower }} id: '%s'", {{ $.Name | first_char_lowered }}.Id)
      }
      var q struct {
        Account struct {
          {{ $.Name | title | makeSingular }} struct {
            {{- if or (hasPrefix "ancestor" .Name) (hasPrefix "child" .Name) -}}
              {{ .Name | title }} {{ $.Name | title | makeSingular }}Connection
            {{- else if hasPrefix "descendant" .Name }}
              {{ .Name | title }} {{ .Name | title | makeSingular | trimPrefix "Descendant" }}Connection
            {{- else -}}
              {{ .Name | title }} {{ if eq .Name "memberships" }}Team{{end}}{{ .Name | title | makeSingular }}Connection
            {{- end }} ` + "`" + `graphql:"{{.Name}}(after: $after, first: $first)"` + "`" + `
          } ` + "`" + `graphql:"{{$.Name | word_first_char_lowered }}(id: $id)"` + "`" + `
        }
      }
      if variables == nil {
        variables = client.InitialPageVariablesPointer()
      }
      (*variables)["id"] = {{ $.Name | first_char_lowered }}.Id
      if err := client.Query(&q, *variables, WithName("{{ template "name_to_singular" . }}List")); err != nil {
        return nil, err
      }

      for q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.PageInfo.HasNextPage {
        (*variables)["after"] = q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.PageInfo.End
        connection, err := {{ $.Name | first_char_lowered }}.List{{ .Name | title }}(client, variables)
        if err != nil {
          return nil, err
        }
        q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.Nodes = append(q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.Nodes, connection.Nodes...)
        q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.PageInfo = q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.PageInfo
        q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.TotalCount += q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}.TotalCount
      }

      return &q.Account.{{ $.Name | title | makeSingular }}.{{ .Name | title }}, nil

    {{- else }}Get{{ .Name | title }}({{ query_args . }}) (*{{.Name | title | makeSingular }}, error) {
      var q struct {
        Account struct {
          {{ .Name | title }} {{ template "name_to_singular" . }} ` + "`" + `graphql:"{{.Name}}(input: $input)"` + "`" + `
        }
      }
      v := PayloadVariables{"input": *NewIdentifier(identifier)}
      if err := client.Query(&q, v, WithName("{{ template "name_to_singular" . }}Get")); err != nil {
        return nil, err
      }
      return &q.Account.{{ .Name | title }}, nil
    {{- end -}}
  }
  {{- end -}}{{- end -}}{{- end -}}
	`),
	mutationFile: t(header + `
	{{range .Types | sortByName}}
	  {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	    {{- if eq .Name "Mutation" }}
	      {{- template "mutation" .}}
	    {{end}}
	  {{- end}}
	{{- end}}

	{{ define "mutation" -}}
	{{- range .Fields }} {{- if not (skip_if_campaign_or_group .Name) }}
	// {{ .Name | title | renameMutation }} {{ .Description | clean | fullSentence }}
	func (client *Client) {{ .Name | title | renameMutation }}(
    {{- if hasSuffix "Delete" .Name }}id ID
    {{- else }}
      {{- range $index, $element := .Args }} {{- if gt $index 0 }}, {{ end -}}
	    {{- if eq "IdentifierInput" .Type.OfType.OfTypeName }}identifier string
	    {{- else if hasSuffix "Delete" $.Name }}id ID
	    {{- else if eq "String" .Type.OfType.OfTypeName }}{{ .Name }} string
	    {{- else }}{{- .Name }} {{ with .Type.OfType.OfTypeName }}{{.}}{{else}}any{{end}}
	    {{- end }}
    {{- end }}
	  {{- end -}} ) {{ if hasSuffix "Delete" .Name -}}
                  error {
	                {{- else -}}
                  (*{{.Name | title | renameMutationReturnType}}, error) {
	                {{- end -}}
      {{- if hasSuffix "Delete" .Name }}
      input := {{.Name | title}}Input{Id: id}
      {{ end }}
	    var m struct {
	      {{ .Name | title }}Payload {{ template "graphql_struct_tag_with_args" . }}
	    }
	    v := PayloadVariables{ {{ range .Args }}
	      "{{.Name}}": {{- if eq "IdentifierInput" .Type.OfType.OfTypeName }}*NewIdentifier(identifier),
	                   {{- else}}input,{{ end }}
	                           {{- end}}
	    }
	    err := client.Mutate(&m, v, WithName("{{ .Name | title }}"))
      {{- if hasSuffix "Delete" .Name }}
      return HandleErrors(err, m.{{ .Name | title }}Payload.Errors)
      {{- else }}
	    return &m.{{ .Name | title }}Payload.{{ .Name | title | renameMutationReturnType}}, HandleErrors(err, m.{{ .Name | title }}Payload.Errors)
      {{- end }}
	}
	{{- end}}{{ end }}{{- end}}
	`),
	objectFile: t(header + `
  import "github.com/relvacode/iso8601"

	{{range .Types | sortByName}}
	  {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	  {{- if ne .Name "Account" }}
	    {{template "object" .}}{{end}}
	  {{- end}}
	{{- end}}

	{{- define "object" -}}
	{{ if not (skip_object .Name) }}
	{{ template "type_comment_description" . }}
	type {{.Name}} struct { {{ add_special_fields .Name }}
    {{ range .Fields }}
    {{- if and (not (skip_object_field $.Name .Name)) (not (len .Args)) }}
      {{ .Name | title}} {{ get_field_type $.Name . }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	  {{- end -}}{{ end }}
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

func makePlural(s string) string {
	value := strings.ToLower(s)
	if isPlural(s) {
		return s
	} else if strings.HasSuffix(value, "y") {
		return strings.ReplaceAll(s, "y", "ies")
	} else if strings.HasSuffix(value, "s") {
		return s + "es"
	}
	return s
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
	} else if slices.Contains(knownIntsByName, value) {
		return "int"
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

func renameMutationReturnType(s string) string {
	create, delete, update := "Create", "Delete", "Update"
	if strings.HasSuffix(s, create) {
		s = strings.TrimSuffix(s, create)
	} else if strings.HasSuffix(s, delete) {
		s = strings.TrimSuffix(s, delete)
	} else if strings.HasSuffix(s, update) {
		s = strings.TrimSuffix(s, update)
	}
	return s
}

func renameMutation(s string) string {
	create, delete, update := "Create", "Delete", "Update"
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

func fieldCommentDescription(fieldType *types.InputValueDefinition) string {
	oneLineDescription := strings.ReplaceAll(fieldType.Desc, "\n", " ")
	if _, ok := fieldType.Type.(*types.NonNull); ok {
		return fmt.Sprintf(" // %s (Required.)", oneLineDescription)
	}
	return fmt.Sprintf(" // %s (Optional.)", oneLineDescription)
}

func exampleStructTag(field *types.InputValueDefinition) string {
	var exampleValue string
	var unwrappedType types.Type

	if nonNullType, ok := field.Type.(*types.NonNull); ok {
		unwrappedType = nonNullType.OfType
	} else {
		unwrappedType = field.Type
	}

	switch fieldType := unwrappedType.(type) {
	case *types.EnumTypeDefinition:
		if enumValue, ok := enumExamples[strings.TrimSuffix(fieldType.String(), "!")]; ok {
			exampleValue = enumValue
		} else {
			exampleValue = "ENUM_TODO"
		}
	case *types.InputObject:
		return "" // omit 'example' struct tag, implicit nested tag is used
	case *types.ScalarTypeDefinition:
		if scalarValue, ok := scalarExamples[strings.TrimSuffix(fieldType.String(), "!")]; ok {
			exampleValue = scalarValue
		} else {
			exampleValue = "SCALAR_TODO"
		}
	case *types.List:
		if listValue, ok := listExamples[field.Name.Name]; ok {
			exampleValue = listValue
		} else {
			exampleValue = "LIST_TODO"
		}
	case *types.TypeName:
		exampleValue = "TYPENAME_TODO"
	default:
		exampleValue = "UNKNOWN_TODO"
	}

	return fmt.Sprintf(` example:"%s"`, exampleValue)
}

func yamlStructTag(fieldType *types.InputValueDefinition) string {
	jsonStructTag := jsonStructTag(fieldType)
	return strings.Replace(jsonStructTag, "json", "yaml", 1)
}

func jsonStructTag(fieldType *types.InputValueDefinition) string {
	fieldName := fieldType.Name.Name
	if isNullable(fieldType.Type) {
		return fmt.Sprintf(`json:"%s,omitempty"`, fieldName)
	}
	return fmt.Sprintf(`json:"%s"`, fieldName)
}

func graphqlTypeToGolang(graphqlType string) string {
	var convertedType string
	nullable := "*Nullable["
	if strings.HasSuffix(graphqlType, "!") {
		graphqlType = strings.TrimSuffix(graphqlType, "!")
	} else {
		// GraphQL nullable --> Go pointer
		convertedType += nullable
	}

	// GraphQL list --> Go slice
	if strings.HasPrefix(graphqlType, "[") {
		graphqlType = strings.TrimPrefix(graphqlType, "[")
		graphqlType = strings.TrimSuffix(graphqlType, "]")
		// NOTE: pretty sure we don't support slices containing pointers
		graphqlType = strings.TrimSuffix(graphqlType, "!")
		convertedType += "[]"
	}

	// GraphQL scalars/types --> Go type
	switch graphqlType {
	case "Boolean":
		convertedType += "bool"
	case "Int":
		convertedType += "int"
	case "ISO8601DateTime":
		convertedType += "iso8601.Time"
	case "Float", "String":
		convertedType += "string"
	default:
		convertedType += graphqlType
	}

	if strings.HasPrefix(convertedType, nullable) {
		convertedType += "]"
	}

	return convertedType
}

func getFieldTypeNew(fieldType types.Type) string {
	return graphqlTypeToGolang(fieldType.String())
}

func isNullable(fieldType types.Type) bool {
	return fieldType.Kind() != "NON_NULL"
}

func isPlural(s string) bool {
	value := strings.ToLower(s)
	// Examples: "alias", "address", "status", "levels"
	if value == "notes" || value == "days" || value == "headers" || value == "responsibilities" ||
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

func wordFirstCharLowered(s string) string {
	return firstCharLowered(s) + s[1:]
}

func firstCharLowered(s string) string {
	return strings.ToLower(s[:1])
}

var templFuncMap = template.FuncMap{
	"internal":                            func(s string) bool { return strings.HasPrefix(s, "__") },
	"quote":                               strconv.Quote,
	"join":                                strings.Join,
	"check_fragments":                     fragmentsForCheck,
	"custom_actions_ext_action_fragments": fragmentsForCustomActionsExtAction,
	"integration_fragments":               fragmentsForIntegration,
	"get_field_type":                      getFieldTypeOld,
	"get_input_field_type":                getInputFieldType,
	"add_special_fields":                  addSpecialFields,
	"add_special_interfaces_fields":       addSpecialInterfacesFields,
	"query_args":                          queryArgs,
	"skip_object":                         skipObject,
	"skip_if_campaign_or_group":           skipCampaignAndGroup,
	"skip_object_field":                   skipObjectField,
	"skip_query":                          skipQuery,
	"skip_interface_field":                skipInterfaceField,
	"isListType":                          isPlural,
	"getFieldType":                        getFieldTypeNew,
	"exampleStructTag":                    exampleStructTag,
	"jsonStructTag":                       jsonStructTag,
	"yamlStructTag":                       yamlStructTag,
	"fieldCommentDescription":             fieldCommentDescription,
	"isTypeNullable":                      isNullable,
	"renameMutation":                      renameMutation,
	"renameMutationReturnType":            renameMutationReturnType,
	"convertPayloadType":                  convertPayloadType,
	"first_char_lowered":                  firstCharLowered,
	"word_first_char_lowered":             wordFirstCharLowered,
	"makePlural":                          makePlural,
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

func addSpecialFields(objectName string) string {
	switch objectName {
	case "Domain":
		return "DomainId"
	case "Filter":
		return "FilterId"
	case "Repository":
		return "Services *RepositoryServiceConnection\nTags *TagConnection"
	case "Scorecard":
		return "ScorecardId"
	case "Service":
		return "ServiceId"
	case "System":
		return "SystemId"
	case "Team":
		return "TeamId"
	case "User":
		return "UserId"
	}
	return ""
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

func skipCampaignAndGroup(objectName string) bool {
	nameLowerCased := strings.ToLower(objectName)
	if strings.Contains(nameLowerCased, "group") ||
		strings.Contains(nameLowerCased, "campaign") {
		return true
	}
	return false
}

func queryArgs(fieldObject GraphQLField) string {
	// var output string
	// for _, b := range fieldObject.Args {
	// 	output += fmt.Sprintf("%s", b.Name)
	// 	fmt.Println(b)
	// }
	// return output
	return "input any"
}

func skipObject(objectName string) bool {
	nameLowerCased := strings.ToLower(objectName)
	switch nameLowerCased {
	case "group", "mutation":
		return true
	}
	if strings.HasPrefix(nameLowerCased, "campaign") ||
		strings.HasSuffix(nameLowerCased, "payload") ||
		strings.HasSuffix(nameLowerCased, "edge") ||
		strings.HasSuffix(nameLowerCased, "connection") {
		return true
	}
	return false
}

func skipObjectField(objectName, fieldName string) bool {
	nameLowerCased := strings.ToLower(fieldName)
	switch nameLowerCased {
	case "metadata", "rawnotes":
		return true
	}
	switch objectName {
	case "Filter":
		switch nameLowerCased {
		case "id", "name":
			return true
		}
	case "Domain", "Service", "Scorecard", "System":
		switch nameLowerCased {
		case "id", "aliases":
			return true
		}
	case "Team":
		switch nameLowerCased {
		case "id", "alias":
			return true
		}
	case "User":
		switch nameLowerCased {
		case "id", "email":
			return true
		}
	}
	return false
}

func skipQuery(objectName string) bool {
	nameLowerCased := strings.ToLower(objectName)
	if strings.Contains(nameLowerCased, "campaign") ||
		strings.Contains(nameLowerCased, "group") ||
		strings.Contains(nameLowerCased, "mutation") {
		return true
	}
	return false
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

func getInputFieldType(inputField GraphQLField) string {
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
	case slices.Contains(knownIntsByName, lowercaseFieldName):
		return "int"
	case slices.Contains(knownIsoTimeByName, lowercaseFieldName):
		return "iso8601.Time"
	}
	return "string"
}

func getFieldTypeOld(objectName string, inputField GraphQLField) string {
	lowercaseFieldName := strings.ToLower(inputField.Name)
	switch {
	case "type" == lowercaseFieldName:
		switch objectName {
		case "AlertSource":
			return "AlertSourceTypeEnum"
		case "AlertSourceUsageCheck", "CustomCheck", "CustomEventCheck",
			"GitBranchProtectionCheck", "HasDocumentationCheck", "HasRecentDeployCheck",
			"ManualCheck", "PayloadCheck", "RepositoryFileCheck", "RepositoryGrepCheck",
			"RepositoryIntegratedCheck", "RepositorySearchCheck", "ServiceConfigurationCheck",
			"ServiceDependencyCheck", "ServiceOwnershipCheck", "ServicePropertyCheck",
			"TagDefinedCheck", "ToolUsageCheck":
			return "CheckType"
		case "ApiDocIntegration", "AwsIntegration", "AzureDevopsIntegration",
			"AzureDevopsPermissionError", "BitbucketIntegration", "CheckIntegration",
			"DatadogIntegration", "DeployIntegration", "FluxIntegration", "GenericIntegration",
			"GithubActionsIntegration", "GithubIntegration", "GitlabIntegration",
			"InfrastructureResource", "InfrastructureResourceSchema", "Repository",
			"IssueTrackingIntegration", "JenkinsIntegration", "KubernetesIntegration",
			"NewRelicIntegration", "OctopusDeployIntegration", "OnPremGitlabIntegration",
			"OpsgenieIntegration", "PagerdutyIntegration", "PayloadIntegration",
			"RelationshipType", "ScimIntegration", "SlackIntegration", "TerraformIntegration":
			return "string"
		case "Contact":
			return "ContactType"
		case "FilterPredicate":
			return "PredicateTypeEnum"
		case "InfrastructureResourceProviderData":
			return "DoubleCheckThis"
		case "Predicate":
			return "PredicateTypeEnum"
		default:
			return "any"
		}
	case objectName == "AlertSource" && lowercaseFieldName == "integration":
		return "IntegrationId"
	case objectName == "AlertSourceService":
		switch lowercaseFieldName {
		case "alertsource":
			return "AlertSource"
		case "service":
			return "ServiceId"
		case "status":
			return "AlertSourceStatusTypeEnum"
		}
	case objectName == "CustomActionsTriggerDefinition":
		switch lowercaseFieldName {
		case "accesscontrol":
			return "CustomActionsTriggerDefinitionAccessControlEnum"
		case "action":
			return "CustomActionsId"
		case "entitytype":
			return "CustomActionsEntityTypeEnum"
		case "filter":
			return "FilterId"
		case "owner":
			return "TeamId"
		}
	case objectName == "CustomActionsWebhookAction":
		switch lowercaseFieldName {
		case "headers":
			return "JSON"
		case "httpmethod":
			return "CustomActionsHttpMethodEnum"
		}
	case objectName == "Domain":
		switch lowercaseFieldName {
		case "managedaliases":
			return "[]string"
		case "owner":
			return "EntityOwner"
		}
	case objectName == "Filter":
		switch lowercaseFieldName {
		case "connective":
			return "ConnectiveEnum"
		case "predicates":
			return "[]FilterPredicate"
		}
	case objectName == "FilterPredicate":
		switch lowercaseFieldName {
		case "casesensitive":
			return "*bool"
		case "key":
			return "PredicateKeyEnum"
		}
	case objectName == "InfrastructureResource":
		switch lowercaseFieldName {
		case "data", "rawdata":
			return "JSON"
		case "id":
			return "ID"
		case "owner":
			return "EntityOwner"
		case "providerdata":
			return "InfrastructureResourceProviderData"
		}
	case objectName == "InfrastructureResourceSchema" && lowercaseFieldName == "schema":
		return "JSON"
	case objectName == "Language" && lowercaseFieldName == "usage":
		return "float64"
	case objectName == "Repository":
		switch lowercaseFieldName {
		case "languages":
			return "[]Language"
		case "owner":
			return "TeamId"
		}
	case objectName == "Scorecard":
		switch lowercaseFieldName {
		case "owner":
			return "EntityOwner"
		case "passingchecks", "servicecount", "totalchecks":
			return "int"
		}
	case objectName == "Secret" && lowercaseFieldName == "owner":
		return "TeamId"
	case objectName == "Repository" && lowercaseFieldName == "owner":
		return "TeamId"
	case objectName == "Service":
		switch lowercaseFieldName {
		case "managedaliases":
			return "[]string"
		case "owner":
			return "TeamId"
		case "preferredapidocument":
			return "*ServiceDocument"
		case "preferredapidocumentsource":
			return "*ApiDocumentSourceEnum"
		}
	case objectName == "ServiceDependency":
		switch lowercaseFieldName {
		case "destinationservice", "sourceservice":
			return "ServiceId"
		}
	case objectName == "ServiceDocument" && lowercaseFieldName == "source":
		return "ServiceDocumentSource"
	case objectName == "ServiceRepository":
		switch lowercaseFieldName {
		case "repository":
			return "RepositoryId"
		case "service":
			return "ServiceId"
		}
	case objectName == "System":
		switch lowercaseFieldName {
		case "managedaliases":
			return "[]string"
		case "owner":
			return "EntityOwner"
		case "parent":
			return "Domain"
		}
	case objectName == "Team":
		switch lowercaseFieldName {
		case "contacts":
			return "Contact"
		case "managedaliases":
			return "[]string"
		case "parentteam":
			return "TeamId"
		}
	case objectName == "TeamMembership":
		switch lowercaseFieldName {
		case "team":
			return "TeamId"
		case "user":
			return "UserId"
		}
	case objectName == "Tool":
		switch lowercaseFieldName {
		case "category":
			return "ToolCategory"
		case "service":
			return "ServiceId"
		}
	case objectName == "User" && lowercaseFieldName == "role":
		return "UserRole"
	}

	return getInputFieldType(inputField)
}
