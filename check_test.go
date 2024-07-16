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
	fixture  autopilot.TestRequest
	endpoint string
	body     func(c *ol.Client) (*ol.Check, error)
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

func MergeCheckResponse(extras map[string]any) string {
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

func BuildCheckResponse(name string, style string, extras map[string]any) string {
	data := map[string]any{
		"Name":  name,
		"Style": style,
		"Body":  MergeCheckResponse(extras),
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

func BuildCheckCreateResponse(name string, extras map[string]any) string {
	return BuildCheckResponse(name, "Create", extras)
}

func BuildCheckUpdateResponse(name string, extras map[string]any) string {
	return BuildCheckResponse(name, "Update", extras)
}

func getCheckTestCases() map[string]TmpCheckTestCase {
	testcases := map[string]TmpCheckTestCase{
		"CreateAlertSourceUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("AlertSourceUsage"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "alertSourceNamePredicate": {"type":"equals", "value":"Requests"}, "alertSourceType":"datadog" } }`,
				BuildCheckCreateResponse("AlertSourceUsage", map[string]any{}),
			),
			endpoint: "check/create_alert_source_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkAlertSourceUsageCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckAlertSourceUsageCreateInput](checkCreateInput)
				checkAlertSourceUsageCreateInput.AlertSourceType = ol.RefOf(ol.AlertSourceTypeEnumDatadog)
				checkAlertSourceUsageCreateInput.AlertSourceNamePredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("Requests"),
				}
				return c.CreateCheckAlertSourceUsage(*checkAlertSourceUsageCreateInput)
			},
		},
		"UpdateAlertSourceUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("AlertSourceUsage"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "alertSourceNamePredicate": {"type":"equals", "value":"Requests"}, "alertSourceType":"datadog" } }`,
				BuildCheckUpdateResponse("AlertSourceUsage", map[string]any{}),
			),
			endpoint: "check/update_alert_source_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkAlertSourceUsageUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckAlertSourceUsageUpdateInput](checkUpdateInput)
				checkAlertSourceUsageUpdateInput.AlertSourceType = ol.RefOf(ol.AlertSourceTypeEnumDatadog)
				checkAlertSourceUsageUpdateInput.AlertSourceNamePredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("Requests"),
				}
				return c.UpdateCheckAlertSourceUsage(*checkAlertSourceUsageUpdateInput)
			},
		},

		"CreateGitBranchProtection": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("GitBranchProtection"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkGitBranchProtectionCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_git_branch_protection",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkGitBranchProtectionCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckGitBranchProtectionCreateInput](checkCreateInput)
				return c.CreateCheckGitBranchProtection(*checkGitBranchProtectionCreateInput)
			},
		},
		"UpdateGitBranchProtection": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("GitBranchProtection"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkGitBranchProtectionUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_git_branch_protection",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkGitBranchProtectionUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckGitBranchProtectionUpdateInput](checkUpdateInput)
				return c.UpdateCheckGitBranchProtection(*checkGitBranchProtectionUpdateInput)
			},
		},

		"CreateHasRecentDeploy": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("HasRecentDeploy"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "days": 5 } }`,
				`{ "data": { "checkHasRecentDeployCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_has_recent_deploy",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkHasRecentDeployCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckHasRecentDeployCreateInput](checkCreateInput)
				checkHasRecentDeployCreateInput.Days = 5
				return c.CreateCheckHasRecentDeploy(*checkHasRecentDeployCreateInput)
			},
		},
		"UpdateHasRecentDeploy": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("HasRecentDeploy"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "days": 5 } }`,
				`{ "data": { "checkHasRecentDeployUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_has_recent_deploy",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkHasRecentDeployUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckHasRecentDeployUpdateInput](checkUpdateInput)
				checkHasRecentDeployUpdateInput.Days = ol.RefOf(5)
				return c.UpdateCheckHasRecentDeploy(*checkHasRecentDeployUpdateInput)
			},
		},

		"CreateHasDocumentation": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("HasDocumentation"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "documentType": "api", "documentSubtype": "openapi", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkHasDocumentationCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has valid documentation.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "documentType": "api", "documentSubtype": "openapi" }, "errors": [] } } }`,
			),
			endpoint: "check/create_has_documentation",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkHasDocumentationCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckHasDocumentationCreateInput](checkCreateInput)
				checkHasDocumentationCreateInput.DocumentType = ol.HasDocumentationTypeEnumAPI
				checkHasDocumentationCreateInput.DocumentSubtype = ol.HasDocumentationSubtypeEnumOpenapi
				return c.CreateCheckHasDocumentation(*checkHasDocumentationCreateInput)
			},
		},
		"UpdateHasDocumentation": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("HasDocumentation"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "documentType": "api", "documentSubtype": "openapi", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkHasDocumentationUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World Update", "notes": "Hello World Update", "documentType": "api", "documentSubtype": "openapi" }, "errors": [] } } }`,
			),
			endpoint: "check/update_has_documentation",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkHasDocumentationUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckHasDocumentationUpdateInput](checkUpdateInput)
				checkHasDocumentationUpdateInput.DocumentType = ol.RefOf(ol.HasDocumentationTypeEnumAPI)
				checkHasDocumentationUpdateInput.DocumentSubtype = ol.RefOf(ol.HasDocumentationSubtypeEnumOpenapi)
				return c.UpdateCheckHasDocumentation(*checkHasDocumentationUpdateInput)
			},
		},

		"CreateCustomEvent": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("CustomEvent"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "#Hello World", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				`{ "data": { "checkCustomEventCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_custom_event",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkCustomEventCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckCustomEventCreateInput](checkCreateInput)
				checkCustomEventCreateInput.ServiceSelector = ".metadata.name"
				checkCustomEventCreateInput.SuccessCondition = ".metadata.name"
				checkCustomEventCreateInput.ResultMessage = ol.RefOf("#Hello World")
				checkCustomEventCreateInput.IntegrationId = ol.ID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg")
				checkCustomEventCreateInput.PassPending = ol.RefOf(false)
				return c.CreateCheckCustomEvent(*checkCustomEventCreateInput)
			},
		},
		"UpdateCustomEvent": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("CustomEvent"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "#Hello World", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				BuildCheckUpdateResponse("CustomEvent", map[string]any{"resultMessage": "#Hello World"}),
				//`{ "data": { "checkCustomEventUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_custom_event",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkCustomEventUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckCustomEventUpdateInput](checkUpdateInput)
				checkCustomEventUpdateInput.ServiceSelector = ol.RefOf(".metadata.name")
				checkCustomEventUpdateInput.SuccessCondition = ol.RefOf(".metadata.name")
				checkCustomEventUpdateInput.ResultMessage = ol.RefOf("#Hello World")
				checkCustomEventUpdateInput.IntegrationId = ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg")
				checkCustomEventUpdateInput.PassPending = ol.RefOf(false)
				return c.UpdateCheckCustomEvent(*checkCustomEventUpdateInput)
			},
		},
		"UpdateCustomEventNoMessage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("CustomEvent"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				BuildCheckUpdateResponse("CustomEvent", map[string]any{}),
				//`{ "data": { "checkCustomEventUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_custom_event_no_message",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkCustomEventUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckCustomEventUpdateInput](checkUpdateInput)
				checkCustomEventUpdateInput.ServiceSelector = ol.RefOf(".metadata.name")
				checkCustomEventUpdateInput.SuccessCondition = ol.RefOf(".metadata.name")
				checkCustomEventUpdateInput.ResultMessage = ol.RefOf("")
				checkCustomEventUpdateInput.IntegrationId = ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg")
				checkCustomEventUpdateInput.PassPending = ol.RefOf(false)
				return c.UpdateCheckCustomEvent(*checkCustomEventUpdateInput)
			},
		},
		"CreateManual": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("Manual"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "updateFrequency": { "startingDate": "2021-07-26T20:22:44.427Z", "frequencyTimeScale": "week", "frequencyValue": 1 }, "updateRequiresComment": false } }`,
				`{ "data": { "checkManualCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a service owner to manually complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_manual",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkManualCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckManualCreateInput](checkCreateInput)
				checkManualCreateInput.UpdateFrequency = ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1)
				return c.CreateCheckManual(*checkManualCreateInput)
			},
		},
		"UpdateManual": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("Manual"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "updateFrequency": { "startingDate": "2021-07-26T20:22:44.427Z", "frequencyTimeScale": "week", "frequencyValue": 1 } } }`,
				`{ "data": { "checkManualUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a service owner to manually complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_manual",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkManualUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckManualUpdateInput](checkUpdateInput)
				checkManualUpdateInput.UpdateFrequency = ol.NewManualCheckFrequencyUpdateInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1)
				return c.UpdateCheckManual(*checkManualUpdateInput)
			},
		},
		"CreateRepositoryFile": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositoryFile"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "directorySearch": true, "filePaths": [ "/src", "/test" ], "fileContentsPredicate": { "type": "equals", "value": "postgres" }, "useAbsoluteRoot": true } }`,
				`{ "data": { "checkRepositoryFileCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo directorys ''/src' or '/test'' equals 'postgres'.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_file",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryFileCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryFileCreateInput](checkCreateInput)
				checkRepositoryFileCreateInput.DirectorySearch = ol.RefOf(true)
				checkRepositoryFileCreateInput.FilePaths = []string{"/src", "/test"}
				checkRepositoryFileCreateInput.FileContentsPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("postgres"),
				}
				checkRepositoryFileCreateInput.UseAbsoluteRoot = ol.RefOf(true)
				return c.CreateCheckRepositoryFile(*checkRepositoryFileCreateInput)
			},
		},
		"UpdateRepositoryFile": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryFile"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "directorySearch": true, "filePaths": [ "/src", "/test" ], "fileContentsPredicate": { "type": "equals", "value": "postgres" }, "useAbsoluteRoot": false } }`,
				`{ "data": { "checkRepositoryFileUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo directorys ''/src' or '/test'' equals 'postgres'.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_file",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryFileUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryFileUpdateInput](checkUpdateInput)
				checkRepositoryFileUpdateInput.DirectorySearch = ol.RefOf(true)
				checkRepositoryFileUpdateInput.FilePaths = &[]string{"/src", "/test"}
				checkRepositoryFileUpdateInput.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("postgres"),
				}
				checkRepositoryFileUpdateInput.UseAbsoluteRoot = ol.RefOf(false)
				return c.UpdateCheckRepositoryFile(*checkRepositoryFileUpdateInput)
			},
		},
		"CreateRepositoryGrep": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositoryGrep"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "directorySearch": true, "filePaths": [ "**/hello.go" ], "fileContentsPredicate": { "type": "exists" } } }`,
				`{ "data": { "checkRepositoryGrepCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw", "name": "Performance" }, "description": "Verifies the existence and/or contents of files in a service's attached Git repositories.", "enabled": true, "enableOn": null, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ1", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "owner": null, "type": "repo_grep" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_grep",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryGrepCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryGrepCreateInput](checkCreateInput)
				checkRepositoryGrepCreateInput.DirectorySearch = ol.RefOf(true)
				checkRepositoryGrepCreateInput.FilePaths = []string{"**/hello.go"}
				checkRepositoryGrepCreateInput.FileContentsPredicate = ol.PredicateInput{
					Type: ol.PredicateTypeEnumExists,
				}
				return c.CreateCheckRepositoryGrep(*checkRepositoryGrepCreateInput)
			},
		},
		"UpdateRepositoryGrep": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryGrep"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "directorySearch": true, "filePaths": [ "**/go.mod" ], "fileContentsPredicate": { "type": "exists" } } }`,
				`{ "data": { "checkRepositoryGrepUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw", "name": "Performance" }, "description": "Verifies the existence and/or contents of files in a service's attached Git repositories.", "enabled": true, "enableOn": null, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ1", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "owner": null, "type": "repo_grep" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_grep",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryGrepUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryGrepUpdateInput](checkUpdateInput)
				checkRepositoryGrepUpdateInput.DirectorySearch = ol.RefOf(true)
				checkRepositoryGrepUpdateInput.FilePaths = &[]string{"**/go.mod"}
				checkRepositoryGrepUpdateInput.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type: ol.RefOf(ol.PredicateTypeEnumExists),
				}
				return c.UpdateCheckRepositoryGrep(*checkRepositoryGrepUpdateInput)
			},
		},
		"UpdateRepositoryGrepMissingDirectorySearch": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryGrep"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "filePaths": [ "**/go.mod" ], "directorySearch": false, "fileContentsPredicate": { "type": "exists" } } }`,
				`{ "data": { "checkRepositoryGrepUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw", "name": "Performance" }, "description": "Verifies the existence and/or contents of files in a service's attached Git repositories.", "enabled": true, "enableOn": null, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ1", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "owner": null, "type": "repo_grep" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_grep_missing_directory_search",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryGrepUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryGrepUpdateInput](checkUpdateInput)
				checkRepositoryGrepUpdateInput.DirectorySearch = ol.RefOf(false)
				checkRepositoryGrepUpdateInput.FilePaths = &[]string{"**/go.mod"}
				checkRepositoryGrepUpdateInput.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type: ol.RefOf(ol.PredicateTypeEnumExists),
				}
				return c.UpdateCheckRepositoryGrep(*checkRepositoryGrepUpdateInput)
			},
		},
		"CreateRepositoryIntegrated": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositoryIntegrated"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkRepositoryIntegratedCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has a repository integrated.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_integrated",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryIntegratedCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryIntegratedCreateInput](checkCreateInput)
				return c.CreateCheckRepositoryIntegrated(*checkRepositoryIntegratedCreateInput)
			},
		},
		"UpdateRepositoryIntegrated": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositoryIntegrated"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkRepositoryIntegratedUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has a repository integrated.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_integrated",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositoryIntegratedUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryIntegratedUpdateInput](checkUpdateInput)
				return c.UpdateCheckRepositoryIntegrated(*checkRepositoryIntegratedUpdateInput)
			},
		},
		"CreateRepositorySearch": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("RepositorySearch"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "fileExtensions": [ "sbt", "py" ], "fileContentsPredicate": { "type": "contains", "value": "postgres" } } }`,
				`{ "data": { "checkRepositorySearchCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo contains search term 'postgres' in at least one .sbt or .py file.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_search",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositorySearchCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckRepositorySearchCreateInput](checkCreateInput)
				checkRepositorySearchCreateInput.FileExtensions = &[]string{"sbt", "py"}
				checkRepositorySearchCreateInput.FileContentsPredicate = ol.PredicateInput{
					Type:  ol.PredicateTypeEnumContains,
					Value: ol.RefOf("postgres"),
				}

				return c.CreateCheckRepositorySearch(*checkRepositorySearchCreateInput)
			},
		},
		"UpdateRepositorySearch": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("RepositorySearch"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "fileExtensions": [ "sbt", "py" ], "fileContentsPredicate": { "type": "contains", "value": "postgres" } } }`,
				`{ "data": { "checkRepositorySearchUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo contains search term 'postgres' in at least one .sbt or .py file.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_search",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkRepositorySearchUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositorySearchUpdateInput](checkUpdateInput)
				checkRepositorySearchUpdateInput.FileExtensions = &[]string{"sbt", "py"}
				checkRepositorySearchUpdateInput.FileContentsPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumContains),
					Value: ol.RefOf("postgres"),
				}
				return c.UpdateCheckRepositorySearch(*checkRepositorySearchUpdateInput)
			},
		},
		"CreateServiceDependency": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceDependency"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkServiceDependencyCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has either a dependent or dependency.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_dependency",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceDependencyCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckServiceDependencyCreateInput](checkCreateInput)
				return c.CreateCheckServiceDependency(*checkServiceDependencyCreateInput)
			},
		},
		"UpdateServiceDependency": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceDependency"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkServiceDependencyUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has either a dependent or dependency.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_dependency",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceDependencyUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceDependencyUpdateInput](checkUpdateInput)
				return c.UpdateCheckServiceDependency(*checkServiceDependencyUpdateInput)
			},
		},
		"CreateServiceConfiguration": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceConfiguration"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkServiceConfigurationCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service is maintained though the use of an opslevel.yml service config.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_configuration",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceConfigurationCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckServiceConfigurationCreateInput](checkCreateInput)
				return c.CreateCheckServiceConfiguration(*checkServiceConfigurationCreateInput)
			},
		},
		"UpdateServiceConfiguration": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceConfiguration"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkServiceConfigurationUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service is maintained though the use of an opslevel.yml service config.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_configuration",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceConfigurationUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceConfigurationUpdateInput](checkUpdateInput)
				return c.UpdateCheckServiceConfiguration(*checkServiceConfigurationUpdateInput)
			},
		},
		"CreateServiceOwnership": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceOwnership"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "requireContactMethod": true, "contactMethod": "slack", "tagKey": "updated_at", "tagPredicate": { "type": "equals", "value": "2-11-2022" } } }`,
				`{ "data": { "checkServiceOwnershipCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_ownership",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceOwnershipCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckServiceOwnershipCreateInput](checkCreateInput)
				checkServiceOwnershipCreateInput.RequireContactMethod = ol.RefOf(true)
				checkServiceOwnershipCreateInput.ContactMethod = ol.RefOf(string(ol.ContactTypeSlack))
				checkServiceOwnershipCreateInput.TagKey = ol.RefOf("updated_at")
				checkServiceOwnershipCreateInput.TagPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("2-11-2022"),
				}
				return c.CreateCheckServiceOwnership(*checkServiceOwnershipCreateInput)
			},
		},
		"UpdateServiceOwnership": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceOwnership"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "requireContactMethod": true, "contactMethod": "email", "tagKey": "updated_at", "tagPredicate": { "type": "equals", "value": "2-11-2022" } } }`,
				`{ "data": { "checkServiceOwnershipUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_ownership",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceOwnershipUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceOwnershipUpdateInput](checkUpdateInput)
				checkServiceOwnershipUpdateInput.RequireContactMethod = ol.RefOf(true)
				checkServiceOwnershipUpdateInput.ContactMethod = ol.RefOf(string(ol.ContactTypeEmail))
				checkServiceOwnershipUpdateInput.TagKey = ol.RefOf("updated_at")
				checkServiceOwnershipUpdateInput.TagPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("2-11-2022"),
				}
				return c.UpdateCheckServiceOwnership(*checkServiceOwnershipUpdateInput)
			},
		},
		"CreateServiceProperty": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ServiceProperty"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "serviceProperty": "framework", "propertyValuePredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkServicePropertyCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service has a framework that equals <code>postgres</code>", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_property",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServicePropertyCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckServicePropertyCreateInput](checkCreateInput)
				checkServicePropertyCreateInput.ServiceProperty = ol.ServicePropertyTypeEnumFramework
				checkServicePropertyCreateInput.PropertyValuePredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("postgres"),
				}
				return c.CreateCheckServiceProperty(*checkServicePropertyCreateInput)
			},
		},
		"UpdateServiceProperty": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ServiceProperty"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "serviceProperty": "framework", "propertyValuePredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkServicePropertyUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service has a framework that equals <code>postgres</code>", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_property",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServicePropertyUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckServicePropertyUpdateInput](checkUpdateInput)
				checkServicePropertyUpdateInput.ServiceProperty = ol.RefOf(ol.ServicePropertyTypeEnumFramework)
				checkServicePropertyUpdateInput.PropertyValuePredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("postgres"),
				}
				return c.UpdateCheckServiceProperty(*checkServicePropertyUpdateInput)
			},
		},
		"CreateTagDefined": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("TagDefined"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "tagKey": "app", "tagPredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkTagDefinedCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has the specified tag defined.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_tag_defined",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkTagDefinedCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckTagDefinedCreateInput](checkCreateInput)
				checkTagDefinedCreateInput.TagKey = "app"
				checkTagDefinedCreateInput.TagPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("postgres"),
				}
				return c.CreateCheckTagDefined(*checkTagDefinedCreateInput)
			},
		},
		"UpdateTagDefined": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("TagDefined"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "tagKey": "app", "tagPredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkTagDefinedUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has the specified tag defined.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_tag_defined",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkTagDefinedUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckTagDefinedUpdateInput](checkUpdateInput)
				checkTagDefinedUpdateInput.TagKey = ol.RefOf("app")
				checkTagDefinedUpdateInput.TagPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("postgres"),
				}
				return c.UpdateCheckTagDefined(*checkTagDefinedUpdateInput)
			},
		},
		"CreateToolUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckCreateMutation("ToolUsage"),
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "toolCategory": "metrics", "toolNamePredicate": { "type": "equals", "value": "datadog" }, "toolUrlPredicate": { "type": "contains", "value": "https://" }, "environmentPredicate": { "type": "equals", "value": "production" } } }`,
				`{ "data": { "checkToolUsageCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service is using 'datadog' as a metrics tool in the 'production' environment.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_tool_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkToolUsageCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckToolUsageCreateInput](checkCreateInput)
				checkToolUsageCreateInput.ToolCategory = ol.ToolCategoryMetrics
				checkToolUsageCreateInput.ToolNamePredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("datadog"),
				}
				checkToolUsageCreateInput.ToolUrlPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumContains,
					Value: ol.RefOf("https://"),
				}
				checkToolUsageCreateInput.EnvironmentPredicate = &ol.PredicateInput{
					Type:  ol.PredicateTypeEnumEquals,
					Value: ol.RefOf("production"),
				}
				return c.CreateCheckToolUsage(*checkToolUsageCreateInput)
			},
		},
		"UpdateToolUsage": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ToolUsage"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "toolCategory": "metrics", "toolNamePredicate": { "type": "equals", "value": "datadog" }, "toolUrlPredicate": { "type": "contains", "value": "https://" }, "environmentPredicate": { "type": "equals", "value": "production" } } }`,
				`{ "data": { "checkToolUsageUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service is using 'datadog' as a metrics tool in the 'production' environment.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_tool_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkToolUsageUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckToolUsageUpdateInput](checkUpdateInput)
				checkToolUsageUpdateInput.ToolCategory = ol.RefOf(ol.ToolCategoryMetrics)
				checkToolUsageUpdateInput.ToolNamePredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("datadog"),
				}
				checkToolUsageUpdateInput.ToolUrlPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumContains),
					Value: ol.RefOf("https://"),
				}
				checkToolUsageUpdateInput.EnvironmentPredicate = &ol.PredicateUpdateInput{
					Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
					Value: ol.RefOf("production"),
				}
				return c.UpdateCheckToolUsage(*checkToolUsageUpdateInput)
			},
		},
		"UpdateToolUsageNullPredicates": {
			fixture: autopilot.NewTestRequest(
				BuildCheckUpdateMutation("ToolUsage"),
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "toolCategory": "metrics", "toolUrlPredicate": null, "environmentPredicate": null } }`,
				`{ "data": { "checkToolUsageUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service is using 'datadog' as a metrics tool in the 'production' environment.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_tool_usage_null_predicates",
			body: func(c *ol.Client) (*ol.Check, error) {
				checkToolUsageUpdateInput := ol.NewCheckUpdateInputTypeOf[ol.CheckToolUsageUpdateInput](checkUpdateInput)
				checkToolUsageUpdateInput.ToolCategory = ol.RefOf(ol.ToolCategoryMetrics)
				checkToolUsageUpdateInput.ToolUrlPredicate = &ol.PredicateUpdateInput{}
				checkToolUsageUpdateInput.EnvironmentPredicate = &ol.PredicateUpdateInput{}
				return c.UpdateCheckToolUsage(*checkToolUsageUpdateInput)
			},
		},
		"GetCheck": {
			fixture: autopilot.NewTestRequest(
				`query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}}}}`,
				`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4" }`,
				`{ "data": { "account": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Owner Defined", "notes": null } } } }`,
			),
			endpoint: "check/get",
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
		`mutation CheckCustomEventUpdate($input:CheckCustomEventUpdateInput!){checkCustomEventUpdate(input: $input){check{category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "filterId": null }}`,
		`{ "data": {
      "checkCustomEventUpdate": {
        "check": {
          "category": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
            "name": "Performance"
          },
          "enabled": true,
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
          "level": {
            "alias": "bronze",
            "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
            "index": 1,
            "name": "Bronze"
          },
          "notes": "Hello World Notes",
          "name": "Hello World"
        },
        "errors": []
      }}}`,
	)
	client := BestTestClient(t, "check/can_update_filter_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		Id:       ol.ID("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4"),
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
		`mutation CheckCustomEventUpdate($input:CheckCustomEventUpdateInput!){checkCustomEventUpdate(input: $input){check{category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "notes": "" }}`,
		`{ "data": {
      "checkCustomEventUpdate": {
        "check": {
          "category": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
            "name": "Performance"
          },
          "enabled": true,
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
          "level": {
            "alias": "bronze",
            "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
            "index": 1,
            "name": "Bronze"
          },
          "name": "Hello World"
        },
        "errors": []
      }}}`,
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
