package opslevel_test

import (
	"bytes"
	"encoding/json"
	"github.com/Masterminds/sprig/v3"
	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
	"strings"
	"testing"
	"text/template"
)

var templater = template.New("").Funcs(sprig.TxtFuncMap()).Delims("[%", "%]")

var checkCreateInput = ol.CheckCreateInput{
	Name:     "Hello World",
	Enabled:  ol.RefOf(true),
	Category: ol.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1"),
	Level:    ol.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3"),
	Notes:    ol.RefOf("Hello World Check"),
}

var checkUpdateNotes = "Hello World Check"

var checkUpdateInput = ol.CheckUpdateInput{
	Id:       "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
	Name:     "Hello World",
	Enabled:  ol.RefOf(true),
	Category: "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
	Level:    "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
	Notes:    &checkUpdateNotes,
}

// Temporary solution until repo wide testing is standardized
type TmpCheckTestCase struct {
	fixture autopilot.TestRequest
	body    func(c *ol.Client) (*ol.Check, error)
}

func TrimGraphQLString(input string) string {
	processed := strings.ReplaceAll(input, "\n", "")
	processed = strings.ReplaceAll(processed, "\t", "")
	return strings.ReplaceAll(processed, "  ", "")
}

func Template(text string, data map[string]any) string {
	tpl, err := templater.Clone()
	if err != nil {
		panic(err)
	}
	parsed, err := tpl.Parse(text)
	if err != nil {
		panic(err)
	}
	result := bytes.NewBuffer([]byte{})
	if err = parsed.Execute(result, data); err != nil {
		panic(err)
	}
	return result.String()
}

func BuildCheckMutation(name string, style string) string {
	data := map[string]any{
		"Name":  name,
		"Style": style,
	}
	text := TrimGraphQLString(`mutation Check[% .Name %][% .Style %]($input:Check[% .Name %][% .Style %]Input!){
  check[% .Name %][% .Style %](input: $input){
    check{
      category{id,name},
      description,
      enableOn,
      enabled,
      filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},
      id,
      level{alias,description,id,index,name},
      name,
      notes: rawNotes,
      owner{... on Team{alias,id}},
      type,
      ... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},
      ... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},
      ... on HasRecentDeployCheck{days},
      ... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},
      ... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},
      ... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},
      ... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},
      ... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},
      ... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},
      ... on TagDefinedCheck{tagKey,tagPredicate{type,value}},
      ... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},
      ... on HasDocumentationCheck{documentType,documentSubtype}
	},
	errors{message,path}
  }
}`)
	return Template(text, data)
}

func BuildCheckCreateMutation(name string) string {
	return BuildCheckMutation(name, "Create")
}

func BuildCheckUpdateMutation(name string) string {
	return BuildCheckMutation(name, "Update")
}

func MergeCheckData(extras map[string]any) string {
	data := map[string]any{
		"category": map[string]any{
			"id":   "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
			"name": "Performance",
		},
		"description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.",
		"enabled":     true,
		"id":          "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
		"level": map[string]any{
			"alias":       "bronze",
			"description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
			"id":          "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
			"index":       1,
			"name":        "Bronze",
		},
		"name": "Hello World",
	}
	for k, v := range extras {
		data[k] = v
	}
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func BuildCheckMutationResponse(name string, style string, extras map[string]any) string {
	data := map[string]any{
		"Name":  name,
		"Style": style,
		"Body":  MergeCheckData(extras),
	}
	text := TrimGraphQLString(`{
  "data": {
    "check[% .Name %][% .Style %]": {
      "check": [% .Body %],
      "errors": []
    }
  }
}`)
	return Template(text, data)
}

func BuildCheckCreateMutationResponse(name string, extras map[string]any) string {
	return BuildCheckMutationResponse(name, "Create", extras)
}

func BuildCheckUpdateMutationResponse(name string, extras map[string]any) string {
	return BuildCheckMutationResponse(name, "Update", extras)
}

func getCheckTestCases() map[string]TmpCheckTestCase {
	testcases := map[string]TmpCheckTestCase{
		"CreateAlertSourceUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("AlertSourceUsage"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "alertSourceNamePredicate": {"type":"equals", "value":"Requests"}, "alertSourceType":"datadog" } }`,
				BuildCheckCreateMutationResponse("AlertSourceUsage", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckAlertSourceUsageCreateInput](checkCreateInput)
				input.AlertSourceType = ol.RefOf(ol.AlertSourceTypeEnumDatadog)
				input.AlertSourceNamePredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("Requests"),
				}
				return c.CreateCheckAlertSourceUsage(*input)
			},
		},
		"UpdateAlertSourceUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("AlertSourceUsage"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "alertSourceNamePredicate": {"type":"equals", "value":"Requests"}, "alertSourceType":"datadog" } }`,
				BuildCheckUpdateMutationResponse("AlertSourceUsage", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckAlertSourceUsageUpdateInput](checkUpdateInput)
				input.AlertSourceType = ol.RefOf(ol.AlertSourceTypeEnumDatadog)
				input.AlertSourceNamePredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("Requests"),
				}
				return c.UpdateCheckAlertSourceUsage(*input)
			},
		},

		"CreateGitBranchProtection": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("GitBranchProtection"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				BuildCheckCreateMutationResponse("GitBranchProtection", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckGitBranchProtectionCreateInput](checkCreateInput)
				return c.CreateCheckGitBranchProtection(*input)
			},
		},
		"UpdateGitBranchProtection": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("GitBranchProtection"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				BuildCheckUpdateMutationResponse("GitBranchProtection", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckGitBranchProtectionUpdateInput](checkUpdateInput)
				return c.UpdateCheckGitBranchProtection(*input)
			},
		},

		"CreateHasRecentDeploy": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("HasRecentDeploy"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "days": 5 } }`,
				BuildCheckCreateMutationResponse("HasRecentDeploy", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckHasRecentDeployCreateInput](checkCreateInput)
				input.Days = 5
				return c.CreateCheckHasRecentDeploy(*input)
			},
		},
		"UpdateHasRecentDeploy": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("HasRecentDeploy"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "days": 5 } }`,
				BuildCheckUpdateMutationResponse("HasRecentDeploy", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckHasRecentDeployUpdateInput](checkUpdateInput)
				input.Days = ol.RefOf(5)
				return c.UpdateCheckHasRecentDeploy(*input)
			},
		},

		"CreateHasDocumentation": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("HasDocumentation"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "documentType": "api", "documentSubtype": "openapi", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				BuildCheckCreateMutationResponse("HasDocumentation", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckHasDocumentationCreateInput](checkCreateInput)
				input.DocumentType = ol.HasDocumentationTypeEnumAPI
				input.DocumentSubtype = ol.HasDocumentationSubtypeEnumOpenapi
				return c.CreateCheckHasDocumentation(*input)
			},
		},
		"UpdateHasDocumentation": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("HasDocumentation"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "documentType": "api", "documentSubtype": "openapi", {{ template "check_base_vars" }} } }`,
				BuildCheckUpdateMutationResponse("HasDocumentation", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckHasDocumentationUpdateInput](checkUpdateInput)
				input.DocumentType = ol.RefOf(ol.HasDocumentationTypeEnumAPI)
				input.DocumentSubtype = ol.RefOf(ol.HasDocumentationSubtypeEnumOpenapi)
				return c.UpdateCheckHasDocumentation(*input)
			},
		},

		"CreateCustomEvent": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("CustomEvent"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "#Hello World", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				BuildCheckCreateMutationResponse("CustomEvent", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckCustomEventCreateInput](checkCreateInput)
				input.ServiceSelector = ".metadata.name"
				input.SuccessCondition = ".metadata.name"
				input.ResultMessage = ol.RefOf("#Hello World")
				input.IntegrationId = ol.ID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg")
				input.PassPending = ol.RefOf(false)
				return c.CreateCheckCustomEvent(*input)
			},
		},
		"UpdateCustomEvent": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("CustomEvent"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "#Hello World", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				BuildCheckUpdateMutationResponse("CustomEvent", map[string]any{"resultMessage": "#Hello World"}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckCustomEventUpdateInput](checkUpdateInput)
				input.ServiceSelector = ol.RefOf(".metadata.name")
				input.SuccessCondition = ol.RefOf(".metadata.name")
				input.ResultMessage = ol.RefOf("#Hello World")
				input.IntegrationId = ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg")
				input.PassPending = ol.RefOf(false)
				return c.UpdateCheckCustomEvent(*input)
			},
		},
		"UpdateCustomEventNoMessage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("CustomEvent"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				BuildCheckUpdateMutationResponse("CustomEvent", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckCustomEventUpdateInput](checkUpdateInput)
				input.ServiceSelector = ol.RefOf(".metadata.name")
				input.SuccessCondition = ol.RefOf(".metadata.name")
				input.ResultMessage = ol.RefOf("")
				input.IntegrationId = ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg")
				input.PassPending = ol.RefOf(false)
				return c.UpdateCheckCustomEvent(*input)
			},
		},
		"CreateManual": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("Manual"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "updateFrequency": { "startingDate": "2021-07-26T20:22:44.427Z", "frequencyTimeScale": "week", "frequencyValue": 1 }, "updateRequiresComment": false } }`,
				BuildCheckCreateMutationResponse("Manual", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckManualCreateInput](checkCreateInput)
				input.UpdateFrequency = ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1)
				return c.CreateCheckManual(*input)
			},
		},
		"UpdateManual": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("Manual"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "updateFrequency": { "startingDate": "2021-07-26T20:22:44.427Z", "frequencyTimeScale": "week", "frequencyValue": 1 } } }`,
				BuildCheckUpdateMutationResponse("Manual", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckManualUpdateInput](checkUpdateInput)
				input.UpdateFrequency = ol.NewManualCheckFrequencyUpdateInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1)
				return c.UpdateCheckManual(*input)
			},
		},
		"CreateRepositoryFile": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositoryFile"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "directorySearch": true, "filePaths": [ "/src", "/test" ], "fileContentsPredicate": { "type": "equals", "value": "postgres" }, "useAbsoluteRoot": true } }`,
				BuildCheckCreateMutationResponse("RepositoryFile", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryFileCreateInput](checkCreateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = []string{"/src", "/test"}
				input.FileContentsPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("postgres"),
				}
				input.UseAbsoluteRoot = ol.RefOf(true)
				return c.CreateCheckRepositoryFile(*input)
			},
		},
		"UpdateRepositoryFile": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryFile"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "directorySearch": true, "filePaths": [ "/src", "/test" ], "fileContentsPredicate": { "type": "equals", "value": "postgres" }, "useAbsoluteRoot": false } }`,
				BuildCheckUpdateMutationResponse("RepositoryFile", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryFileUpdateInput](checkUpdateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = &[]string{"/src", "/test"}
				input.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("postgres"),
				}
				input.UseAbsoluteRoot = ol.RefOf(false)
				return c.UpdateCheckRepositoryFile(*input)
			},
		},
		"CreateRepositoryGrep": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositoryGrep"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "directorySearch": true, "filePaths": [ "**/hello.go" ], "fileContentsPredicate": { "type": "exists" } } }`,
				BuildCheckCreateMutationResponse("RepositoryGrep", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryGrepCreateInput](checkCreateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = []string{"**/hello.go"}
				input.FileContentsPredicate = ol.PredicateInput{
					Type: ol.PredicateTypeEnumExists,
				}
				return c.CreateCheckRepositoryGrep(*input)
			},
		},
		"UpdateRepositoryGrep": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryGrep"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "directorySearch": true, "filePaths": [ "**/go.mod" ], "fileContentsPredicate": { "type": "exists" } } }`,
				BuildCheckUpdateMutationResponse("RepositoryGrep", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryGrepUpdateInput](checkUpdateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = &[]string{"**/go.mod"}
				input.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type: ol.RefOf(ol.PredicateTypeEnumExists),
				}
				return c.UpdateCheckRepositoryGrep(*input)
			},
		},
		"UpdateRepositoryGrepMissingDirectorySearch": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryGrep"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "filePaths": [ "**/go.mod" ], "directorySearch": false, "fileContentsPredicate": { "type": "exists" } } }`,
				BuildCheckUpdateMutationResponse("RepositoryGrep", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryGrepUpdateInput](checkUpdateInput)
				input.DirectorySearch = ol.RefOf(false)
				input.FilePaths = &[]string{"**/go.mod"}
				input.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type: ol.RefOf(ol.PredicateTypeEnumExists),
				}
				return c.UpdateCheckRepositoryGrep(*input)
			},
		},
		"CreateRepositoryIntegrated": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositoryIntegrated"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				BuildCheckCreateMutationResponse("RepositoryIntegrated", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryIntegratedCreateInput](checkCreateInput)
				return c.CreateCheckRepositoryIntegrated(*input)
			},
		},
		"UpdateRepositoryIntegrated": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryIntegrated"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				BuildCheckUpdateMutationResponse("RepositoryIntegrated", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryIntegratedUpdateInput](checkUpdateInput)
				return c.UpdateCheckRepositoryIntegrated(*input)
			},
		},
		"CreateRepositorySearch": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositorySearch"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "fileExtensions": [ "sbt", "py" ], "fileContentsPredicate": { "type": "contains", "value": "postgres" } } }`,
				BuildCheckCreateMutationResponse("RepositorySearch", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositorySearchCreateInput](checkCreateInput)
				input.FileExtensions = &[]string{"sbt", "py"}
				input.FileContentsPredicate = ol.PredicateInput{
					Type:  ol.PredicateTypeEnumContains,
					Value: ol.RefOf("postgres"),
				}

				return c.CreateCheckRepositorySearch(*input)
			},
		},
		"UpdateRepositorySearch": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositorySearch"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "fileExtensions": [ "sbt", "py" ], "fileContentsPredicate": { "type": "contains", "value": "postgres" } } }`,
				BuildCheckUpdateMutationResponse("RepositorySearch", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositorySearchUpdateInput](checkUpdateInput)
				input.FileExtensions = &[]string{"sbt", "py"}
				input.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumContains),
					Value: ol.RefOf("postgres"),
				}
				return c.UpdateCheckRepositorySearch(*input)
			},
		},
		"CreateServiceDependency": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceDependency"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				BuildCheckCreateMutationResponse("ServiceDependency", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceDependencyCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckServiceDependencyCreateInput](checkCreateInput)
				return c.CreateCheckServiceDependency(*checkServiceDependencyCreateInput)
			},
		},
		"UpdateServiceDependency": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceDependency"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				BuildCheckUpdateMutationResponse("ServiceDependency", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceDependencyUpdateInput](checkUpdateInput)
				return c.UpdateCheckServiceDependency(*input)
			},
		},
		"CreateServiceConfiguration": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceConfiguration"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				BuildCheckCreateMutationResponse("ServiceConfiguration", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServiceConfigurationCreateInput](checkCreateInput)
				return c.CreateCheckServiceConfiguration(*input)
			},
		},
		"UpdateServiceConfiguration": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceConfiguration"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				BuildCheckUpdateMutationResponse("ServiceConfiguration", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceConfigurationUpdateInput](checkUpdateInput)
				return c.UpdateCheckServiceConfiguration(*input)
			},
		},
		"CreateServiceOwnership": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceOwnership"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "requireContactMethod": true, "contactMethod": "slack", "tagKey": "updated_at", "tagPredicate": { "type": "equals", "value": "2-11-2022" } } }`,
				BuildCheckCreateMutationResponse("ServiceOwnership", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServiceOwnershipCreateInput](checkCreateInput)
				input.RequireContactMethod = ol.RefOf(true)
				input.ContactMethod = ol.RefOf(string(ol.ContactTypeSlack))
				input.TagKey = ol.RefOf("updated_at")
				input.TagPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("2-11-2022"),
				}
				return c.CreateCheckServiceOwnership(*input)
			},
		},
		"UpdateServiceOwnership": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceOwnership"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "requireContactMethod": true, "contactMethod": "email", "tagKey": "updated_at", "tagPredicate": { "type": "equals", "value": "2-11-2022" } } }`,
				BuildCheckUpdateMutationResponse("ServiceOwnership", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceOwnershipUpdateInput](checkUpdateInput)
				input.RequireContactMethod = ol.RefOf(true)
				input.ContactMethod = ol.RefOf(string(ol.ContactTypeEmail))
				input.TagKey = ol.RefOf("updated_at")
				input.TagPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("2-11-2022"),
				}
				return c.UpdateCheckServiceOwnership(*input)
			},
		},
		"CreateServiceProperty": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceProperty"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "serviceProperty": "framework", "propertyValuePredicate": { "type": "equals", "value": "postgres" } } }`,
				BuildCheckCreateMutationResponse("ServiceProperty", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServicePropertyCreateInput](checkCreateInput)
				input.ServiceProperty = ol.ServicePropertyTypeEnumFramework
				input.PropertyValuePredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("postgres"),
				}
				return c.CreateCheckServiceProperty(*input)
			},
		},
		"UpdateServiceProperty": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceProperty"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "serviceProperty": "framework", "propertyValuePredicate": { "type": "equals", "value": "postgres" } } }`,
				BuildCheckUpdateMutationResponse("ServiceProperty", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServicePropertyUpdateInput](checkUpdateInput)
				input.ServiceProperty = ol.RefOf(ol.ServicePropertyTypeEnumFramework)
				input.PropertyValuePredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("postgres"),
				}
				return c.UpdateCheckServiceProperty(*input)
			},
		},
		"CreateTagDefined": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("TagDefined"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "tagKey": "app", "tagPredicate": { "type": "equals", "value": "postgres" } } }`,
				BuildCheckCreateMutationResponse("TagDefined", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckTagDefinedCreateInput](checkCreateInput)
				input.TagKey = "app"
				input.TagPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("postgres"),
				}
				return c.CreateCheckTagDefined(*input)
			},
		},
		"UpdateTagDefined": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("TagDefined"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "tagKey": "app", "tagPredicate": { "type": "equals", "value": "postgres" } } }`,
				BuildCheckUpdateMutationResponse("TagDefined", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckTagDefinedUpdateInput](checkUpdateInput)
				input.TagKey = ol.RefOf("app")
				input.TagPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("postgres"),
				}
				return c.UpdateCheckTagDefined(*input)
			},
		},
		"CreateToolUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ToolUsage"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "toolCategory": "metrics", "toolNamePredicate": { "type": "equals", "value": "datadog" }, "toolUrlPredicate": { "type": "contains", "value": "https://" }, "environmentPredicate": { "type": "equals", "value": "production" } } }`,
				BuildCheckCreateMutationResponse("ToolUsage", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckToolUsageCreateInput](checkCreateInput)
				input.ToolCategory = ol.ToolCategoryMetrics
				input.ToolNamePredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("datadog"),
				}
				input.ToolUrlPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumContains,
					Value: ol.RefOf("https://"),
				}
				input.EnvironmentPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("production"),
				}
				return c.CreateCheckToolUsage(*input)
			},
		},
		"UpdateToolUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ToolUsage"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "toolCategory": "metrics", "toolNamePredicate": { "type": "equals", "value": "datadog" }, "toolUrlPredicate": { "type": "contains", "value": "https://" }, "environmentPredicate": { "type": "equals", "value": "production" } } }`,
				BuildCheckUpdateMutationResponse("ToolUsage", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckToolUsageUpdateInput](checkUpdateInput)
				input.ToolCategory = ol.RefOf(ol.ToolCategoryMetrics)
				input.ToolNamePredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("datadog"),
				}
				input.ToolUrlPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumContains),
					Value: ol.RefOf("https://"),
				}
				input.EnvironmentPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("production"),
				}
				return c.UpdateCheckToolUsage(*input)
			},
		},
		"UpdateToolUsageNullPredicates": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ToolUsage"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "toolCategory": "metrics", "toolUrlPredicate": null, "environmentPredicate": null } }`,
				BuildCheckUpdateMutationResponse("ToolUsage", map[string]any{}),
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckToolUsageUpdateInput](checkUpdateInput)
				input.ToolCategory = ol.RefOf(ol.ToolCategoryMetrics)
				input.ToolUrlPredicate = &ol.PredicateUpdateInput{}
				input.EnvironmentPredicate = &ol.PredicateUpdateInput{}
				return c.UpdateCheckToolUsage(*input)
			},
		},
		"GetCheck": {
			fixture: autopilot.NewTestRequest(
				`query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}}}}`,
				`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4" }`,
				`{ "data": { "account": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Owner Defined", "notes": null } } } }`,
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4")
			},
		},
	}
	return testcases
}

func TestChecks(t *testing.T) {
	for name, tc := range getCheckTestCases() {
		t.Run(name, func(t *testing.T) {
			// Arrange
			client := BestTestClient(t, name, tc.fixture)
			// Act
			result, err := tc.body(client)
			// Assert
			autopilot.Equals(t, nil, err)
			autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", string(result.Id))
			autopilot.Equals(t, "Performance", result.Category.Name)
			autopilot.Equals(t, "Bronze", result.Level.Name)
		})
	}
}

func TestCanUpdateFilterToNull(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		BuildCheckUpdateMutation("CustomEvent"),
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "filterId": null }}`,
		BuildCheckUpdateMutationResponse("CustomEvent", map[string]any{}),
	)
	client := BestTestClient(t, "check/can_update_filter_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		Id:       "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
		FilterId: ol.NewID(),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Hello World", result.Name)
	autopilot.Equals(t, ol.ID(""), result.Filter.Id)
}

func TestCanUpdateNotesToNull(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		BuildCheckUpdateMutation("CustomEvent"),
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "notes": "" }}`,
		BuildCheckUpdateMutationResponse("CustomEvent", map[string]any{}),
	)
	client := BestTestClient(t, "check/can_update_notes_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		Id:    "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
		Notes: ol.RefOf(""),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Hello World", result.Name)
	autopilot.Equals(t, "", result.Notes)
}

func TestListChecks(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query CheckList($after:String!$first:Int!){account{rubric{checks(after: $after, first: $first){nodes{category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "checks": { "nodes": [ { {{ template "common_check_response" }} }, { {{ template "metrics_tool_check" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query CheckList($after:String!$first:Int!){account{rubric{checks(after: $after, first: $first){nodes{category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "checks": { "nodes": [ { {{ template "owner_defined_check" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "check/list", requests...)
	// Act
	response, err := client.ListChecks(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Metrics Tool", result[1].Name)
	autopilot.Equals(t, "Tier 1 Services", result[1].Filter.Name)
	autopilot.Equals(t, "Owner Defined", result[2].Name)
	autopilot.Equals(t, "Verifies that the service has an owner defined.", result[2].Description)
}

func TestGetMissingCheck(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}}}}`,
		`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf" }`,
		`{ "data": { "account": { "check": null } } }`,
	)
	client := BestTestClient(t, "check/get_missing", testRequest)
	// Act
	_, err := client.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestDeleteCheck(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation CheckDelete($input:CheckDeleteInput!){checkDelete(input: $input){errors{message,path}}}`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzIxNzI" } }`,
		`{ "data": { "checkDelete": { "errors": [] } } }`,
	)
	client := BestTestClient(t, "check/delete", testRequest)
	// Act
	err := client.DeleteCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzIxNzI")
	// Assert
	autopilot.Equals(t, nil, err)
}

func TestJsonUnmarshalCreateCheckManual(t *testing.T) {
	// Arrange
	data := `{
	"name": "Example",
	"notes": "Example Notes",
	"updateRequiresComment": false
}`
	output := ol.CheckManualCreateInput{
		Name:                  "Example",
		Notes:                 ol.RefOf("Example Notes"),
		UpdateRequiresComment: false,
	}
	// Act
	buf1, err1 := ol.UnmarshalCheckCreateInput(ol.CheckTypeManual, []byte(data))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, &output, buf1.(*ol.CheckManualCreateInput))
}

func TestJsonUnmarshalCreateCheckToolUsage(t *testing.T) {
	// Arrange
	data := `{
	"name": "Example",
	"notes": "Example Notes",
	"environmentPredicate": {
    "type": "exists"
  },
  "toolNamePredicate": {
    "type": "contains",
    "value": "go"
  },
  "toolUrlPredicate": {
    "type": "starts_with",
    "value": "https"
  }
}`
	output := ol.CheckToolUsageCreateInput{
		Name:  "Example",
		Notes: ol.RefOf("Example Notes"),
		EnvironmentPredicate: &ol.PredicateInput{
			Type: ol.PredicateTypeEnum("exists"),
		},
		ToolNamePredicate: &ol.PredicateInput{
			Type:  ol.PredicateTypeEnum("contains"),
			Value: ol.RefOf("go"),
		},
		ToolUrlPredicate: &ol.PredicateInput{
			Type:  ol.PredicateTypeEnum("starts_with"),
			Value: ol.RefOf("https"),
		},
	}
	// Act
	buf1, err1 := ol.UnmarshalCheckCreateInput(ol.CheckTypeToolUsage, []byte(data))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, &output, buf1.(*ol.CheckToolUsageCreateInput))
}

func TestJsonUnmarshalUpdateCheckManual(t *testing.T) {
	// Arrange
	data := `{
	"name": "Example",
	"notes": "Example Notes",
	"updateRequiresComment": true
}`
	output := ol.CheckManualUpdateInput{
		Name:                  ol.RefOf("Example"),
		Notes:                 ol.RefOf("Example Notes"),
		UpdateRequiresComment: ol.RefOf(true),
	}
	// Act
	buf1, err1 := ol.UnmarshalCheckUpdateInput(ol.CheckTypeManual, []byte(data))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, &output, buf1.(*ol.CheckManualUpdateInput))
}

func TestJsonUnmarshalUpdateCheckToolUsage(t *testing.T) {
	// Arrange
	data := `{
	"name": "Example",
	"notes": "Updated Notes",
	"environmentPredicate": {
    "type": "exists"
  },
  "toolNamePredicate": {
    "type": "contains",
    "value": "go"
  },
  "toolUrlPredicate": {
    "type": "starts_with",
    "value": "https"
  }
}`
	output := ol.CheckToolUsageUpdateInput{
		Name:  ol.RefOf("Example"),
		Notes: ol.RefOf("Updated Notes"),
		EnvironmentPredicate: &ol.PredicateUpdateInput{
			Type: ol.RefOf(ol.PredicateTypeEnum("exists")),
		},
		ToolNamePredicate: &ol.PredicateUpdateInput{
			Type:  ol.RefOf(ol.PredicateTypeEnum("contains")),
			Value: ol.RefOf("go"),
		},
		ToolUrlPredicate: &ol.PredicateUpdateInput{
			Type:  ol.RefOf(ol.PredicateTypeEnum("starts_with")),
			Value: ol.RefOf("https"),
		},
	}
	// Act
	buf1, err1 := ol.UnmarshalCheckUpdateInput(ol.CheckTypeToolUsage, []byte(data))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, &output, buf1.(*ol.CheckToolUsageUpdateInput))
}
