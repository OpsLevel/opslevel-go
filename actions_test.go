package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateWebhookAction(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation WebhookActionCreate($input:CustomActionsWebhookActionCreateInput!){customActionsWebhookActionCreate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"headers":"{\"Content-Type\":\"application/json\"}","httpMethod":"POST","liquidTemplate":"{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}","name":"Deploy Rollback","webhookUrl":"https://gitlab.com/api/v4/projects/1/trigger/pipeline"}}`,
		`{"data": {"customActionsWebhookActionCreate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)

	client := TmpBetterTestClient(t, "custom_actions/create_action", testRequest)

	// Act
	action, err := client.CreateWebhookAction(ol.CustomActionsWebhookActionCreateInput{
		Name:           "Deploy Rollback",
		LiquidTemplate: "{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}",
		Headers: ol.JSON{
			"Content-Type": "application/json",
		},
		HTTPMethod: ol.CustomActionsHttpMethodEnumPost,
		WebhookURL: "https://gitlab.com/api/v4/projects/1/trigger/pipeline",
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestListCustomActions(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query ExternalActionList($after:String!$first:Int!){account{customActionsExternalActions(after: $after, first: $first){nodes{aliases,id,description,liquidTemplate,name,... on CustomActionsWebhookAction{headers,httpMethod,webhookUrl}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "customActionsExternalActions": { "nodes": [ { {{ template "custom_action1_response" }} }, { {{ template "custom_action2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query ExternalActionList($after:String!$first:Int!){account{customActionsExternalActions(after: $after, first: $first){nodes{aliases,id,description,liquidTemplate,name,... on CustomActionsWebhookAction{headers,httpMethod,webhookUrl}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "customActionsExternalActions": { "nodes": [ { {{ template "custom_action3_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := TmpPaginatedTestClient(t, "custom_actions/list_actions", requests...)
	// Act
	response, err := client.ListCustomActions(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, "Deploy Freeze", result[1].Name)
	autopilot.Equals(t, "Page On-Call", result[2].Name)
	autopilot.Equals(t, "application/json", result[0].Headers["Content-Type"])
}

func TestUpdateWebhookAction(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation WebhookActionUpdate($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"id": "123456789", "httpMethod":"PUT"}}`,
		`{"data": {"customActionsWebhookActionUpdate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/update_action", testRequest)

	// Act
	action, err := client.UpdateWebhookAction(ol.CustomActionsWebhookActionUpdateInput{
		Id:         "123456789",
		HTTPMethod: ol.CustomActionsHttpMethodEnumPut,
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestUpdateWebhookAction2(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation WebhookActionUpdate($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"id": "123456789","description":"","headers":"{\"Accept\":\"application/json\"}"}}`,
		`{"data": {"customActionsWebhookActionUpdate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)

	client := TmpBetterTestClient(t, "custom_actions/update_action2", testRequest)
	headers := ol.JSON{
		"Accept": "application/json",
	}

	// Act
	action, err := client.UpdateWebhookAction(ol.CustomActionsWebhookActionUpdateInput{
		Id:          "123456789",
		Description: ol.NewString(""),
		Headers:     &headers,
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestDeleteWebhookAction(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation WebhookActionDelete($input:IdentifierInput!){customActionsWebhookActionDelete(resource: $input){errors{message,path}}}"`,
		`"variables":{"input":{"id": "123456789"}}`,
		`{"data": {"customActionsWebhookActionDelete": { "errors": [] }}}`,
	)

	client := TmpBetterTestClient(t, "custom_actions/delete_action", testRequest)

	// Act
	err := client.DeleteWebhookAction(ol.IdentifierInput{
		Id: "123456789",
	})

	// Assert
	autopilot.Ok(t, err)
}

func TestCreateTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"SERVICE","filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/create_trigger", testRequest)

	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:        "Deploy Rollback",
		Description: ol.NewString("Disables the Deploy Freeze"),
		Action:      "123456789",
		Owner:       "123456789",
		Filter:      ol.NewID("987654321"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestCreateTriggerDefinitionWithGlobalEntityType(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"GLOBAL","extendedTeamAccess":[{"alias":"example_1"},{"alias":"example_1"}],"filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/create_trigger_with_global_entity", testRequest)

	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:        "Deploy Rollback",
		Description: ol.NewString("Disables the Deploy Freeze"),
		Action:      "123456789",
		Owner:       "123456789",
		Filter:      ol.NewID("987654321"),
		EntityType:  "GLOBAL",
		ExtendedTeamAccess: &[]ol.IdentifierInput{
			*ol.NewIdentifier("example_1"),
			*ol.NewIdentifier("example_1"),
		},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestCreateTriggerDefinitionWithNullExtendedTeams(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"SERVICE","extendedTeamAccess":[],"filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)

	client := TmpBetterTestClient(t, "custom_actions/create_trigger_with_null_extended_teams", testRequest)
	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:               "Deploy Rollback",
		Description:        ol.NewString("Disables the Deploy Freeze"),
		Action:             "123456789",
		Owner:              "123456789",
		Filter:             ol.NewID("987654321"),
		ExtendedTeamAccess: &[]ol.IdentifierInput{},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestGetTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "query TriggerDefinitionGet($input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){{ template "custom_actions_trigger_request" }}}}"`,
		`"variables":{"input":{"id":"123456789"}}`,
		`{"data": {"account": { "customActionsTriggerDefinition": {{ template "custom_action_trigger2" }} }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/get_trigger", testRequest)

	// Act
	trigger, err := client.GetTriggerDefinition(ol.IdentifierInput{Id: "123456789"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
	autopilot.Equals(t, "Uses Ruby", trigger.Filter.Name)
	autopilot.Equals(t, "123456789", string(trigger.Owner.Id))
}

func TestListTriggerDefinitions(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query TriggerDefinitionList($after:String!$first:Int!){account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{action{aliases,id},aliases,description,filter{id,name},id,manualInputsDefinition,name,owner{alias,id},published,timestamps{createdAt,updatedAt},accessControl,responseTemplate,entityType},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "customActionsTriggerDefinitions": { "nodes": [ { {{ template "custom_action_trigger1_response" }} }, { {{ template "custom_action_trigger2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query TriggerDefinitionList($after:String!$first:Int!){account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{action{aliases,id},aliases,description,filter{id,name},id,manualInputsDefinition,name,owner{alias,id},published,timestamps{createdAt,updatedAt},accessControl,responseTemplate,entityType},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "customActionsTriggerDefinitions": { "nodes": [ { {{ template "custom_action_trigger3_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := TmpPaginatedTestClient(t, "custom_actions/list_triggers", requests...)
	// Act
	triggers, err := client.ListTriggerDefinitions(nil)
	result := triggers.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, triggers.TotalCount)
	autopilot.Equals(t, "Release", result[1].Name)
	autopilot.Equals(t, "Uses Ruby", result[1].Filter.Name)
	autopilot.Equals(t, "123456789", string(result[1].Owner.Id))
	autopilot.Equals(t, "Rollback", result[2].Name)
	autopilot.Equals(t, "Uses Go", result[2].Filter.Name)
	autopilot.Equals(t, "123456781", string(result[2].Owner.Id))
}

func TestUpdateTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"id":"123456789", "filterId":null}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/update_trigger", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:     "123456789",
		Filter: ol.NewID(),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestUpdateTriggerDefinition2(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"id":"123456789", "name":"test", "description": "", "extendedTeamAccess": []}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/update_trigger2", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:                 "123456789",
		Name:               ol.NewString("test"),
		Description:        ol.NewString(""),
		ExtendedTeamAccess: &[]ol.IdentifierInput{},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestUpdateTriggerDefinition3(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`"variables":{"input":{"id":"123456789", "name":"test", "description": "", "extendedTeamAccess": [{"alias": "123456789"},{"id":"Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"}]}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := TmpBetterTestClient(t, "custom_actions/update_trigger3", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:          "123456789",
		Name:        ol.NewString("test"),
		Description: ol.NewString(""),
		ExtendedTeamAccess: &[]ol.IdentifierInput{
			*ol.NewIdentifier("123456789"),
			*ol.NewIdentifier("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"),
		},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestDeleteTriggerDefinition(t *testing.T) {
	// Arrange
	request := `{"query":
		"mutation TriggerDefinitionDelete($input:IdentifierInput!){customActionsTriggerDefinitionDelete(resource: $input){errors{message,path}}}",
		"variables":{"input":{"id":"123456789"}}
	}`
	response := `{"data": {"customActionsTriggerDefinitionDelete": {
     "errors": []
 }}}`
	responseErr := `{"data": {"customActionsTriggerDefinitionDelete": {
     "errors": [{{ template "error1" }}]
 }}}`

	client := ABetterTestClient(t, "custom_actions/delete_trigger", request, response)
	clientErr := ABetterTestClient(t, "custom_actions/delete_trigger_err", request, responseErr)
	clientErr2 := ABetterTestClient(t, "custom_actions/delete_trigger_err2", request, "")
	// Act
	err := client.DeleteTriggerDefinition(ol.IdentifierInput{
		Id: "123456789",
	})
	err2 := clientErr.DeleteTriggerDefinition(ol.IdentifierInput{
		Id: "123456789",
	})
	err3 := clientErr2.DeleteTriggerDefinition(ol.IdentifierInput{
		Id: "123456789",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `OpsLevel API Errors:
	- 'one.two.three' Example Error
`, err2.Error())
	autopilot.Assert(t, err3 != nil, "Expected error was not thrown")
}

func TestListExtendedTeamAccess(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query ExtendedTeamAccessList($after:String!$first:Int!$input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){extendedTeamAccess(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},name,responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		`"variables": {{ template "extended_team_access_get_vars_1" }}`,
		`{{ template "extended_team_access_response_1" }}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query ExtendedTeamAccessList($after:String!$first:Int!$input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){extendedTeamAccess(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},name,responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		`"variables": {{ template "extended_team_access_get_vars_2" }}`,
		`{{ template "extended_team_access_response_2" }}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := TmpPaginatedTestClient(t, "custom_actions/list_extended_team_access", requests...)
	id1 := *ol.NewID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	trigger := ol.CustomActionsTriggerDefinition{Id: id1}

	// Act
	resp, err := trigger.ExtendedTeamAccess(client, nil)
	result := resp.Nodes

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, id1, result[0].TeamId.Id)
}
