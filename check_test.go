package opslevel_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

// Temporary solution until repo wide testing is standardized
type TmpCheckTestCase struct {
	fixture autopilot.TestRequest
	body    func(c *ol.Client) (*ol.Check, error)
}

// Helper Data
var (
	templater = template.New("").Funcs(sprig.TxtFuncMap()).Delims("[%", "%]")

	id = ol.ID("Z2lkOi8vb3BzbGV2ZWwvMTIzNDU2")

	predicateInput = &ol.PredicateInput{
		Type:  ol.PredicateTypeEnumEquals,
		Value: ol.RefOf("Requests"),
	}

	predicateUpdateInput = &ol.PredicateUpdateInput{
		Type:  ol.RefOf(ol.PredicateTypeEnumEquals),
		Value: ol.RefOf("Requests"),
	}

	checkNotes = "Hello World Check"

	checkCreateInput = ol.CheckCreateInput{
		Name:     "Hello World",
		Enabled:  ol.RefOf(true),
		Category: id,
		Level:    id,
		Notes:    &checkNotes,
	}

	checkUpdateInput = ol.CheckUpdateInput{
		Id:       id,
		Name:     "Hello World",
		Enabled:  ol.RefOf(true),
		Category: id,
		Level:    id,
		Notes:    &checkNotes,
	}
)

type RequestStyle string

const (
	CreateRequest RequestStyle = "Create"
	UpdateRequest RequestStyle = "Update"
)

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

func BuildCheckMutation(name string, style RequestStyle) string {
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
      ... on ServicePropertyCheck{serviceProperty,propertyDefinition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},propertyValuePredicate{type,value}},
      ... on TagDefinedCheck{tagKey,tagPredicate{type,value}},
      ... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},
      ... on HasDocumentationCheck{documentType,documentSubtype},
      ... on PackageVersionCheck{missingPackageResult,packageConstraint,packageManager,packageName,packageNameIsRegex,versionConstraintPredicate{type,value}}
	},
	errors{message,path}
  }
}`)
	return Template(text, data)
}

func MergeCheckData(extras map[string]any) string {
	data := map[string]any{
		"category": map[string]any{
			"id":   id,
			"name": "Performance",
		},
		"description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.",
		"enabled":     true,
		"filter":      map[string]any{},
		"id":          id,
		"level": map[string]any{
			"alias":       "bronze",
			"description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
			"id":          id,
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

func BuildCheckMutationResponse(name string, style RequestStyle, extras map[string]any) string {
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

func BuildCheckInput(style RequestStyle, extras map[string]any) string {
	base := map[string]any{
		"name":       "Hello World",
		"enabled":    true,
		"categoryId": id,
		"levelId":    id,
		"notes":      "Hello World Check",
	}
	if style == UpdateRequest {
		base["id"] = id
	}
	for k, v := range extras {
		base[k] = v
	}
	b, err := json.Marshal(map[string]any{
		"input": base,
	})
	if err != nil {
		panic(fmt.Errorf("failed to marshal input: %s", err))
	}
	return string(b)
}

func BuildRequest(style RequestStyle, name string, extras map[string]any) autopilot.TestRequest {
	return autopilot.NewTestRequest(
		BuildCheckMutation(name, style),
		BuildCheckInput(style, extras),
		BuildCheckMutationResponse(name, style, extras),
	)
}

func BuildCreateRequest(name string, extras map[string]any) autopilot.TestRequest {
	return autopilot.NewTestRequest(
		BuildCheckMutation(name, CreateRequest),
		BuildCheckInput(CreateRequest, extras),
		BuildCheckMutationResponse(name, CreateRequest, extras),
	)
}

func BuildCreateRequestAlt(name string, input, response map[string]any) autopilot.TestRequest {
	return autopilot.NewTestRequest(
		BuildCheckMutation(name, CreateRequest),
		BuildCheckInput(CreateRequest, input),
		BuildCheckMutationResponse(name, CreateRequest, response),
	)
}

func BuildUpdateRequest(name string, extras map[string]any) autopilot.TestRequest {
	return autopilot.NewTestRequest(
		BuildCheckMutation(name, UpdateRequest),
		BuildCheckInput(UpdateRequest, extras),
		BuildCheckMutationResponse(name, UpdateRequest, extras),
	)
}

func BuildUpdateRequestAlt(name string, input, response map[string]any) autopilot.TestRequest {
	return autopilot.NewTestRequest(
		BuildCheckMutation(name, UpdateRequest),
		BuildCheckInput(UpdateRequest, input),
		BuildCheckMutationResponse(name, UpdateRequest, response),
	)
}

func getCheckTestCases() map[string]TmpCheckTestCase {
	// Test Cases
	testcases := map[string]TmpCheckTestCase{
		"CreateAlertSourceUsage": {
			fixture: BuildCreateRequest("AlertSourceUsage", map[string]any{
				"alertSourceNamePredicate": predicateInput,
				"alertSourceType":          ol.AlertSourceTypeEnumDatadog,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckAlertSourceUsageCreateInput](checkCreateInput)
				input.AlertSourceType = ol.RefOf(ol.AlertSourceTypeEnumDatadog)
				input.AlertSourceNamePredicate = predicateInput
				return c.CreateCheckAlertSourceUsage(*input)
			},
		},
		"UpdateAlertSourceUsage": {
			fixture: BuildUpdateRequest("AlertSourceUsage", map[string]any{
				"alertSourceNamePredicate": predicateUpdateInput,
				"alertSourceType":          ol.AlertSourceTypeEnumDatadog,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckAlertSourceUsageUpdateInput](checkUpdateInput)
				input.AlertSourceType = ol.RefOf(ol.AlertSourceTypeEnumDatadog)
				input.AlertSourceNamePredicate = predicateUpdateInput
				return c.UpdateCheckAlertSourceUsage(*input)
			},
		},

		"CreateGitBranchProtection": {
			fixture: BuildCreateRequest("GitBranchProtection", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckGitBranchProtectionCreateInput](checkCreateInput)
				return c.CreateCheckGitBranchProtection(*input)
			},
		},
		"UpdateGitBranchProtection": {
			fixture: BuildUpdateRequest("GitBranchProtection", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckGitBranchProtectionUpdateInput](checkUpdateInput)
				return c.UpdateCheckGitBranchProtection(*input)
			},
		},

		"CreateHasRecentDeploy": {
			fixture: BuildCreateRequest("HasRecentDeploy", map[string]any{"days": 5}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckHasRecentDeployCreateInput](checkCreateInput)
				input.Days = 5
				return c.CreateCheckHasRecentDeploy(*input)
			},
		},
		"UpdateHasRecentDeploy": {
			fixture: BuildUpdateRequest("HasRecentDeploy", map[string]any{"days": 5}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckHasRecentDeployUpdateInput](checkUpdateInput)
				input.Days = ol.RefOf(5)
				return c.UpdateCheckHasRecentDeploy(*input)
			},
		},

		"CreateHasDocumentation": {
			fixture: BuildCreateRequest("HasDocumentation", map[string]any{
				"documentType":    ol.HasDocumentationTypeEnumAPI,
				"documentSubtype": ol.HasDocumentationSubtypeEnumOpenapi,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckHasDocumentationCreateInput](checkCreateInput)
				input.DocumentType = ol.HasDocumentationTypeEnumAPI
				input.DocumentSubtype = ol.HasDocumentationSubtypeEnumOpenapi
				return c.CreateCheckHasDocumentation(*input)
			},
		},
		"UpdateHasDocumentation": {
			fixture: BuildUpdateRequest("HasDocumentation", map[string]any{
				"documentType":    ol.HasDocumentationTypeEnumAPI,
				"documentSubtype": ol.HasDocumentationSubtypeEnumOpenapi,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckHasDocumentationUpdateInput](checkUpdateInput)
				input.DocumentType = ol.RefOf(ol.HasDocumentationTypeEnumAPI)
				input.DocumentSubtype = ol.RefOf(ol.HasDocumentationSubtypeEnumOpenapi)
				return c.UpdateCheckHasDocumentation(*input)
			},
		},

		"CreateCustomEvent": {
			fixture: BuildCreateRequestAlt("CustomEvent", map[string]any{
				"passPending":      false,
				"serviceSelector":  ".metadata.name",
				"successCondition": ".metadata.name",
				"resultMessage":    "#Hello World",
				"integrationId":    id,
			}, map[string]any{
				"passPending":      false,
				"serviceSelector":  ".metadata.name",
				"successCondition": ".metadata.name",
				"resultMessage":    "#Hello World",
				"integration": ol.IntegrationId{
					Id: id,
				},
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckCustomEventCreateInput](checkCreateInput)
				input.ServiceSelector = ".metadata.name"
				input.SuccessCondition = ".metadata.name"
				input.ResultMessage = ol.RefOf("#Hello World")
				input.IntegrationId = id
				input.PassPending = ol.RefOf(false)
				return c.CreateCheckCustomEvent(*input)
			},
		},
		"UpdateCustomEvent": {
			fixture: BuildUpdateRequestAlt("CustomEvent", map[string]any{
				"passPending":      false,
				"serviceSelector":  ".metadata.name",
				"successCondition": ".metadata.name",
				"resultMessage":    "#Hello World",
				"integrationId":    id,
			}, map[string]any{
				"passPending":      false,
				"serviceSelector":  ".metadata.name",
				"successCondition": ".metadata.name",
				"resultMessage":    "#Hello World",
				"integration": ol.IntegrationId{
					Id: id,
				},
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckCustomEventUpdateInput](checkUpdateInput)
				input.ServiceSelector = ol.RefOf(".metadata.name")
				input.SuccessCondition = ol.RefOf(".metadata.name")
				input.ResultMessage = ol.RefOf("#Hello World")
				input.IntegrationId = &id
				input.PassPending = ol.RefOf(false)
				return c.UpdateCheckCustomEvent(*input)
			},
		},
		"UpdateCustomEventNoMessage": {
			fixture: BuildUpdateRequestAlt("CustomEvent", map[string]any{
				"passPending":      false,
				"serviceSelector":  ".metadata.name",
				"successCondition": ".metadata.name",
				"resultMessage":    "",
				"integrationId":    id,
			}, map[string]any{
				"passPending":      false,
				"serviceSelector":  ".metadata.name",
				"successCondition": ".metadata.name",
				"resultMessage":    "",
				"integration": ol.IntegrationId{
					Id: id,
				},
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckCustomEventUpdateInput](checkUpdateInput)
				input.ServiceSelector = ol.RefOf(".metadata.name")
				input.SuccessCondition = ol.RefOf(".metadata.name")
				input.ResultMessage = ol.RefOf("")
				input.IntegrationId = &id
				input.PassPending = ol.RefOf(false)
				return c.UpdateCheckCustomEvent(*input)
			},
		},
		"CreateManual": {
			fixture: BuildCreateRequest("Manual", map[string]any{
				"updateFrequency":       ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
				"updateRequiresComment": false,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckManualCreateInput](checkCreateInput)
				input.UpdateFrequency = ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1)
				return c.CreateCheckManual(*input)
			},
		},
		"UpdateManual": {
			fixture: BuildUpdateRequest("Manual", map[string]any{
				"updateFrequency": ol.NewManualCheckFrequencyUpdateInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckManualUpdateInput](checkUpdateInput)
				input.UpdateFrequency = ol.NewManualCheckFrequencyUpdateInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1)
				return c.UpdateCheckManual(*input)
			},
		},
		"CreateRepositoryFile": {
			fixture: BuildCreateRequest("RepositoryFile", map[string]any{
				"directorySearch":       true,
				"filePaths":             []string{"/src", "/test"},
				"fileContentsPredicate": predicateInput,
				"useAbsoluteRoot":       true,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryFileCreateInput](checkCreateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = []string{"/src", "/test"}
				input.FileContentsPredicate = predicateInput
				input.UseAbsoluteRoot = ol.RefOf(true)
				return c.CreateCheckRepositoryFile(*input)
			},
		},
		"UpdateRepositoryFile": {
			fixture: BuildUpdateRequest("RepositoryFile", map[string]any{
				"directorySearch":       true,
				"filePaths":             []string{"/src", "/test", "/foo/bar"},
				"fileContentsPredicate": predicateUpdateInput,
				"useAbsoluteRoot":       false,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryFileUpdateInput](checkUpdateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = &[]string{"/src", "/test", "/foo/bar"}
				input.FileContentsPredicate = predicateUpdateInput
				input.UseAbsoluteRoot = ol.RefOf(false)
				return c.UpdateCheckRepositoryFile(*input)
			},
		},
		"CreateRepositoryGrep": {
			fixture: BuildCreateRequest("RepositoryGrep", map[string]any{
				"directorySearch":       true,
				"filePaths":             []string{"**/hello.go"},
				"fileContentsPredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryGrepCreateInput](checkCreateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = []string{"**/hello.go"}
				input.FileContentsPredicate = *predicateInput
				return c.CreateCheckRepositoryGrep(*input)
			},
		},
		"UpdateRepositoryGrep": {
			fixture: BuildUpdateRequest("RepositoryGrep", map[string]any{
				"directorySearch":       true,
				"filePaths":             []string{"go.mod", "**/go.mod"},
				"fileContentsPredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryGrepUpdateInput](checkUpdateInput)
				input.DirectorySearch = ol.RefOf(true)
				input.FilePaths = &[]string{"go.mod", "**/go.mod"}
				input.FileContentsPredicate = predicateUpdateInput
				return c.UpdateCheckRepositoryGrep(*input)
			},
		},
		"UpdateRepositoryGrepDirectorySearchFalse": {
			fixture: BuildUpdateRequest("RepositoryGrep", map[string]any{
				"directorySearch":       false,
				"filePaths":             []string{"**/go.mod"},
				"fileContentsPredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryGrepUpdateInput](checkUpdateInput)
				input.DirectorySearch = ol.RefOf(false)
				input.FilePaths = &[]string{"**/go.mod"}
				input.FileContentsPredicate = predicateUpdateInput
				return c.UpdateCheckRepositoryGrep(*input)
			},
		},
		"CreateRepositoryIntegrated": {
			fixture: BuildCreateRequest("RepositoryIntegrated", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositoryIntegratedCreateInput](checkCreateInput)
				return c.CreateCheckRepositoryIntegrated(*input)
			},
		},
		"UpdateRepositoryIntegrated": {
			fixture: BuildUpdateRequest("RepositoryIntegrated", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositoryIntegratedUpdateInput](checkUpdateInput)
				return c.UpdateCheckRepositoryIntegrated(*input)
			},
		},
		"CreateRepositorySearch": {
			fixture: BuildCreateRequest("RepositorySearch", map[string]any{
				"fileExtensions":        []string{"sbt", "py"},
				"fileContentsPredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckRepositorySearchCreateInput](checkCreateInput)
				input.FileExtensions = &[]string{"sbt", "py"}
				input.FileContentsPredicate = *predicateInput
				return c.CreateCheckRepositorySearch(*input)
			},
		},
		"UpdateRepositorySearch": {
			fixture: BuildUpdateRequest("RepositorySearch", map[string]any{
				"fileExtensions":        []string{"sbt", "py"},
				"fileContentsPredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckRepositorySearchUpdateInput](checkUpdateInput)
				input.FileExtensions = &[]string{"sbt", "py"}
				input.FileContentsPredicate = predicateUpdateInput
				return c.UpdateCheckRepositorySearch(*input)
			},
		},
		"CreateServiceDependency": {
			fixture: BuildCreateRequest("ServiceDependency", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				checkServiceDependencyCreateInput := ol.NewCheckCreateInputTypeOf[ol.CheckServiceDependencyCreateInput](checkCreateInput)
				return c.CreateCheckServiceDependency(*checkServiceDependencyCreateInput)
			},
		},
		"UpdateServiceDependency": {
			fixture: BuildUpdateRequest("ServiceDependency", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceDependencyUpdateInput](checkUpdateInput)
				return c.UpdateCheckServiceDependency(*input)
			},
		},
		"CreateServiceConfiguration": {
			fixture: BuildCreateRequest("ServiceConfiguration", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServiceConfigurationCreateInput](checkCreateInput)
				return c.CreateCheckServiceConfiguration(*input)
			},
		},
		"UpdateServiceConfiguration": {
			fixture: BuildUpdateRequest("ServiceConfiguration", map[string]any{}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceConfigurationUpdateInput](checkUpdateInput)
				return c.UpdateCheckServiceConfiguration(*input)
			},
		},
		"CreateServiceOwnership": {
			fixture: BuildCreateRequest("ServiceOwnership", map[string]any{
				"requireContactMethod": true,
				"contactMethod":        ol.ContactTypeSlack,
				"tagKey":               "updated_at",
				"tagPredicate":         predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServiceOwnershipCreateInput](checkCreateInput)
				input.RequireContactMethod = ol.RefOf(true)
				input.ContactMethod = ol.RefOf(string(ol.ContactTypeSlack))
				input.TagKey = ol.RefOf("updated_at")
				input.TagPredicate = predicateInput
				return c.CreateCheckServiceOwnership(*input)
			},
		},
		"UpdateServiceOwnership": {
			fixture: BuildUpdateRequest("ServiceOwnership", map[string]any{
				"requireContactMethod": true,
				"contactMethod":        ol.ContactTypeEmail,
				"tagKey":               "updated_at",
				"tagPredicate":         predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServiceOwnershipUpdateInput](checkUpdateInput)
				input.RequireContactMethod = ol.RefOf(true)
				input.ContactMethod = ol.RefOf(string(ol.ContactTypeEmail))
				input.TagKey = ol.RefOf("updated_at")
				input.TagPredicate = predicateUpdateInput
				return c.UpdateCheckServiceOwnership(*input)
			},
		},
		"CreateServiceProperty": {
			fixture: BuildCreateRequest("ServiceProperty", map[string]any{
				"serviceProperty":        ol.ServicePropertyTypeEnumFramework,
				"propertyValuePredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServicePropertyCreateInput](checkCreateInput)
				input.ServiceProperty = ol.ServicePropertyTypeEnumFramework
				input.PropertyValuePredicate = predicateInput
				return c.CreateCheckServiceProperty(*input)
			},
		},
		"CreateServicePropertyDefinition": {
			fixture: BuildCreateRequest("ServiceProperty", map[string]any{
				"serviceProperty":        ol.ServicePropertyTypeEnumFramework,
				"propertyDefinition":     ol.NewIdentifier(string(id)),
				"propertyValuePredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckServicePropertyCreateInput](checkCreateInput)
				input.ServiceProperty = ol.ServicePropertyTypeEnumFramework
				input.PropertyDefinition = ol.NewIdentifier(string(id))
				input.PropertyValuePredicate = predicateInput
				return c.CreateCheckServiceProperty(*input)
			},
		},
		"UpdateServiceProperty": {
			fixture: BuildUpdateRequest("ServiceProperty", map[string]any{
				"serviceProperty":        ol.ServicePropertyTypeEnumFramework,
				"propertyValuePredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServicePropertyUpdateInput](checkUpdateInput)
				input.ServiceProperty = ol.RefOf(ol.ServicePropertyTypeEnumFramework)
				input.PropertyValuePredicate = predicateUpdateInput
				return c.UpdateCheckServiceProperty(*input)
			},
		},
		"UpdateServicePropertyDefinition": {
			fixture: BuildUpdateRequest("ServiceProperty", map[string]any{
				"propertyDefinition": ol.NewIdentifier(string(id)),
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckServicePropertyUpdateInput](checkUpdateInput)
				input.PropertyDefinition = ol.NewIdentifier(string(id))
				return c.UpdateCheckServiceProperty(*input)
			},
		},
		"CreateTagDefined": {
			fixture: BuildCreateRequest("TagDefined", map[string]any{
				"tagKey":       "app",
				"tagPredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckTagDefinedCreateInput](checkCreateInput)
				input.TagKey = "app"
				input.TagPredicate = predicateInput
				return c.CreateCheckTagDefined(*input)
			},
		},
		"UpdateTagDefined": {
			fixture: BuildUpdateRequest("TagDefined", map[string]any{
				"tagKey":       "app",
				"tagPredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckTagDefinedUpdateInput](checkUpdateInput)
				input.TagKey = ol.RefOf("app")
				input.TagPredicate = predicateUpdateInput
				return c.UpdateCheckTagDefined(*input)
			},
		},
		"CreateToolUsage": {
			fixture: BuildCreateRequest("ToolUsage", map[string]any{
				"toolCategory":         ol.ToolCategoryMetrics,
				"toolNamePredicate":    predicateInput,
				"toolUrlPredicate":     predicateInput,
				"environmentPredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckToolUsageCreateInput](checkCreateInput)
				input.ToolCategory = ol.ToolCategoryMetrics
				input.ToolNamePredicate = predicateInput
				input.ToolUrlPredicate = predicateInput
				input.EnvironmentPredicate = predicateInput
				return c.CreateCheckToolUsage(*input)
			},
		},
		"UpdateToolUsage": {
			fixture: BuildUpdateRequest("ToolUsage", map[string]any{
				"toolCategory":         ol.ToolCategoryMetrics,
				"toolNamePredicate":    predicateUpdateInput,
				"toolUrlPredicate":     predicateUpdateInput,
				"environmentPredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckToolUsageUpdateInput](checkUpdateInput)
				input.ToolCategory = ol.RefOf(ol.ToolCategoryMetrics)
				input.ToolNamePredicate = predicateUpdateInput
				input.ToolUrlPredicate = predicateUpdateInput
				input.EnvironmentPredicate = predicateUpdateInput
				return c.UpdateCheckToolUsage(*input)
			},
		},
		"UpdateToolUsageNullPredicates": {
			fixture: BuildUpdateRequest("ToolUsage", map[string]any{
				"toolCategory":         ol.ToolCategoryMetrics,
				"toolUrlPredicate":     nil,
				"environmentPredicate": nil,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckToolUsageUpdateInput](checkUpdateInput)
				input.ToolCategory = ol.RefOf(ol.ToolCategoryMetrics)
				input.ToolUrlPredicate = &ol.PredicateUpdateInput{}
				input.EnvironmentPredicate = &ol.PredicateUpdateInput{}
				return c.UpdateCheckToolUsage(*input)
			},
		},
		"CreatePackageVersion": {
			fixture: BuildCreateRequest("PackageVersion", map[string]any{
				"packageManager":             ol.PackageManagerEnumCargo,
				"packageName":                "cult",
				"packageNameIsRegex":         false,
				"packageConstraint":          ol.PackageConstraintEnumDoesNotExist,
				"missingPackageResult":       ol.CheckResultStatusEnumPassed,
				"versionConstraintPredicate": predicateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckCreateInputTypeOf[ol.CheckPackageVersionCreateInput](checkCreateInput)
				input.PackageManager = ol.PackageManagerEnumCargo
				input.PackageName = "cult"
				input.PackageNameIsRegex = ol.RefOf(false)
				input.PackageConstraint = ol.PackageConstraintEnumDoesNotExist
				input.MissingPackageResult = ol.RefOf(ol.CheckResultStatusEnumPassed)
				input.VersionConstraintPredicate = predicateInput
				return c.CreateCheckPackageVersion(*input)
			},
		},
		"UpdatePackageVersion": {
			fixture: BuildUpdateRequest("PackageVersion", map[string]any{
				"packageManager":             ol.PackageManagerEnumCargo,
				"versionConstraintPredicate": predicateUpdateInput,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckPackageVersionUpdateInput](checkUpdateInput)
				input.PackageManager = ol.RefOf(ol.PackageManagerEnumCargo)
				input.VersionConstraintPredicate = predicateUpdateInput
				return c.UpdateCheckPackageVersion(*input)
			},
		},
		"UpdatePackageVersionPredicateNull": {
			fixture: BuildUpdateRequest("PackageVersion", map[string]any{
				"packageManager":             ol.PackageManagerEnumCargo,
				"versionConstraintPredicate": nil,
			}),
			body: func(c *ol.Client) (*ol.Check, error) {
				input := ol.NewCheckUpdateInputTypeOf[ol.CheckPackageVersionUpdateInput](checkUpdateInput)
				input.PackageManager = ol.RefOf(ol.PackageManagerEnumCargo)
				input.VersionConstraintPredicate = &ol.PredicateUpdateInput{}
				return c.UpdateCheckPackageVersion(*input)
			},
		},
		"GetCheck": {
			fixture: autopilot.NewTestRequest(
				`query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyDefinition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype},... on PackageVersionCheck{missingPackageResult,packageConstraint,packageManager,packageName,packageNameIsRegex,versionConstraintPredicate{type,value}}}}}`,
				`{ "id": "Z2lkOi8vb3BzbGV2ZWwvMTIzNDU2" }`,
				`{ "data": { "account": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvMTIzNDU2", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvMTIzNDU2", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Owner Defined", "notes": null } } } }`,
			),
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.GetCheck(id)
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
			autopilot.Equals(t, id, result.Id)
			autopilot.Equals(t, "Performance", result.Category.Name)
			autopilot.Equals(t, "Bronze", result.Level.Name)
		})
	}
}

func TestCanUpdateFilterToNull(t *testing.T) {
	// Arrange
	testRequest := BuildUpdateRequestAlt("CustomEvent", map[string]any{"filterId": nil}, map[string]any{"filter": map[string]any{}})
	client := BestTestClient(t, "check/can_update_filter_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		Id:         id,
		Name:       ol.RefOf("Hello World"),
		CategoryId: ol.RefOf(id),
		Enabled:    ol.RefOf(true),
		LevelId:    ol.RefOf(id),
		FilterId:   ol.NewID(),
		Notes:      ol.RefOf("Hello World Check"),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Hello World", result.Name)
	autopilot.Equals(t, ol.ID(""), result.Filter.Id)
}

func TestCanUpdateNotesToNull(t *testing.T) {
	// Arrange
	testRequest := BuildUpdateRequest("CustomEvent", map[string]any{
		"notes": "",
	})
	client := BestTestClient(t, "check/can_update_notes_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		Id:         id,
		Name:       ol.RefOf("Hello World"),
		CategoryId: ol.RefOf(id),
		Enabled:    ol.RefOf(true),
		LevelId:    ol.RefOf(id),
		Notes:      ol.RefOf(""),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Hello World", result.Name)
	autopilot.Equals(t, "", result.Notes)
}

func TestListChecks(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query CheckList($after:String!$first:Int!){account{rubric{checks(after: $after, first: $first){nodes{category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyDefinition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype},... on PackageVersionCheck{missingPackageResult,packageConstraint,packageManager,packageName,packageNameIsRegex,versionConstraintPredicate{type,value}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "checks": { "nodes": [ { {{ template "common_check_response" }} }, { {{ template "metrics_tool_check" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query CheckList($after:String!$first:Int!){account{rubric{checks(after: $after, first: $first){nodes{category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyDefinition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype},... on PackageVersionCheck{missingPackageResult,packageConstraint,packageManager,packageName,packageNameIsRegex,versionConstraintPredicate{type,value}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
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
		`query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyDefinition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype},... on PackageVersionCheck{missingPackageResult,packageConstraint,packageManager,packageName,packageNameIsRegex,versionConstraintPredicate{type,value}}}}}`,
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
