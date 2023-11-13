package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

var newID *ol.ID = ol.NewID("123456789")

func TestCreateWebhookAction(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation WebhookActionCreate($input:CustomActionsWebhookActionCreateInput!){customActionsWebhookActionCreate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}"`,
		`{"input":{"headers":"{\"Content-Type\":\"application/json\"}","httpMethod":"POST",{{ template "liquid_template_rollback" }},"name":"Deploy Rollback","webhookUrl":"https://gitlab.com/api/v4/projects/1/trigger/pipeline"}}`,
		`{"data": {"customActionsWebhookActionCreate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/create_action", testRequest)

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
		`"query ExternalActionList($after:String!$first:Int!){account{customActionsExternalActions(after: $after, first: $first){nodes{aliases,id,description,liquidTemplate,name,... on CustomActionsWebhookAction{headers,httpMethod,webhookUrl}},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "customActionsExternalActions": { "nodes": [ { {{ template "custom_action1_response" }} }, { {{ template "custom_action2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query ExternalActionList($after:String!$first:Int!){account{customActionsExternalActions(after: $after, first: $first){nodes{aliases,id,description,liquidTemplate,name,... on CustomActionsWebhookAction{headers,httpMethod,webhookUrl}},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "customActionsExternalActions": { "nodes": [ { {{ template "custom_action3_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "custom_actions/list_actions", requests...)
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
		`"mutation WebhookActionUpdate($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}"`,
		`{"input":{"id": "123456789", "httpMethod":"PUT"}}`,
		`{"data": {"customActionsWebhookActionUpdate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_action", testRequest)

	// Act
	action, err := client.UpdateWebhookAction(ol.CustomActionsWebhookActionUpdateInput{
		Id:         *newID,
		HTTPMethod: ol.CustomActionsHttpMethodEnumPut,
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestUpdateWebhookAction2(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation WebhookActionUpdate($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}"`,
		`{"input":{"id": "123456789","description":"","headers":"{\"Accept\":\"application/json\"}"}}`,
		`{"data": {"customActionsWebhookActionUpdate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/update_action2", testRequest)
	headers := ol.JSON{
		"Accept": "application/json",
	}

	// Act
	action, err := client.UpdateWebhookAction(ol.CustomActionsWebhookActionUpdateInput{
		Id:          *newID,
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
		`"mutation WebhookActionDelete($input:IdentifierInput!){customActionsWebhookActionDelete(resource: $input){errors{message,path}}}"`,
		`{"input":{"id": "123456789"}}`,
		`{"data": {"customActionsWebhookActionDelete": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/delete_action", testRequest)

	// Act
	err := client.DeleteWebhookAction(ol.IdentifierInput{
		Id: newID,
	})

	// Assert
	autopilot.Ok(t, err)
}

func TestCreateTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"SERVICE","filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/create_trigger", testRequest)

	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:        "Deploy Rollback",
		Description: ol.NewString("Disables the Deploy Freeze"),
		Action:      *newID,
		Owner:       *newID,
		Filter:      ol.NewID("987654321"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestCreateTriggerDefinitionWithGlobalEntityType(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"GLOBAL","extendedTeamAccess":[{"alias":"example_1"},{"alias":"example_1"}],"filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/create_trigger_with_global_entity", testRequest)

	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:        "Deploy Rollback",
		Description: ol.NewString("Disables the Deploy Freeze"),
		Action:      *newID,
		Owner:       *newID,
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
		`"mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"SERVICE","extendedTeamAccess":[],"filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/create_trigger_with_null_extended_teams", testRequest)
	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:               "Deploy Rollback",
		Description:        ol.NewString("Disables the Deploy Freeze"),
		Action:             *newID,
		Owner:              *newID,
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
		`"query TriggerDefinitionGet($input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){{ template "custom_actions_trigger_request" }}}}"`,
		`{"input":{"id":"123456789"}}`,
		`{"data": {"account": { "customActionsTriggerDefinition": {{ template "custom_action_trigger2" }} }}}`,
	)
	client := BestTestClient(t, "custom_actions/get_trigger", testRequest)

	// Act
	trigger, err := client.GetTriggerDefinition(ol.IdentifierInput{Id: newID})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
	autopilot.Equals(t, "Uses Ruby", trigger.Filter.Name)
	autopilot.Equals(t, *newID, trigger.Owner.Id)
}

func TestListTriggerDefinitions(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query TriggerDefinitionList($after:String!$first:Int!){account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{action{aliases,id},aliases,description,filter{id,name},id,manualInputsDefinition,name,owner{alias,id},published,timestamps{createdAt,updatedAt},accessControl,responseTemplate,entityType},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "customActionsTriggerDefinitions": { "nodes": [ { {{ template "custom_action_trigger1_response" }} }, { {{ template "custom_action_trigger2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query TriggerDefinitionList($after:String!$first:Int!){account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{action{aliases,id},aliases,description,filter{id,name},id,manualInputsDefinition,name,owner{alias,id},published,timestamps{createdAt,updatedAt},accessControl,responseTemplate,entityType},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "customActionsTriggerDefinitions": { "nodes": [ { {{ template "custom_action_trigger3_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "custom_actions/list_triggers", requests...)
	// Act
	triggers, err := client.ListTriggerDefinitions(nil)
	result := triggers.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, triggers.TotalCount)
	autopilot.Equals(t, "Release", result[1].Name)
	autopilot.Equals(t, "Uses Ruby", result[1].Filter.Name)
	autopilot.Equals(t, *newID, result[1].Owner.Id)
	autopilot.Equals(t, "Rollback", result[2].Name)
	autopilot.Equals(t, "Uses Go", result[2].Filter.Name)
	autopilot.Equals(t, "123456781", string(result[2].Owner.Id))
}

func TestUpdateTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`{"input":{"id":"123456789", "filterId":null}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_trigger", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:     *newID,
		Filter: ol.NewID(),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestUpdateTriggerDefinition2(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`{"input":{"id":"123456789", "name":"test", "description": "", "extendedTeamAccess": []}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_trigger2", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:                 *newID,
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
		`"mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}"`,
		`{"input":{"id":"123456789", "name":"test", "description": "", "extendedTeamAccess": [{"alias": "123456789"}, { {{ template "id1" }} } ]}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_trigger3", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:          *newID,
		Name:        ol.NewString("test"),
		Description: ol.NewString(""),
		ExtendedTeamAccess: &[]ol.IdentifierInput{
			*ol.NewIdentifier(string(*newID)),
			*ol.NewIdentifier(string(id1)),
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
		{"input":{"id":"123456789"}}
	}`

	testRequest := NewTestRequest(
		`"mutation TriggerDefinitionDelete($input:IdentifierInput!){customActionsTriggerDefinitionDelete(resource: $input){errors{message,path}}}"`,
		`{"input":{"id":"123456789"}}`,
		`{"data": {"customActionsTriggerDefinitionDelete": { "errors": [] }}}`,
	)
	testRequestError := NewTestRequest(
		testRequest.Request,
		testRequest.Variables,
		`{"data": {"customActionsTriggerDefinitionDelete": { "errors": [{{ template "error1" }}] }}}`,
	)

	client := BestTestClient(t, "custom_actions/delete_trigger", testRequest)
	clientErr := BestTestClient(t, "custom_actions/delete_trigger_err", testRequestError)
	clientErr2 := ABetterTestClient(t, "custom_actions/delete_trigger_err2", request, "")
	// Act
	err := client.DeleteTriggerDefinition(ol.IdentifierInput{Id: newID})
	err2 := clientErr.DeleteTriggerDefinition(ol.IdentifierInput{Id: newID})
	err3 := clientErr2.DeleteTriggerDefinition(ol.IdentifierInput{Id: newID})
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
		`"query ExtendedTeamAccessList($after:String!$first:Int!$input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){extendedTeamAccess(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{{ template "extended_team_access_get_vars_1" }}`,
		`{{ template "extended_team_access_response_1" }}`,
	)
	testRequestTwo := NewTestRequest(
		`"query ExtendedTeamAccessList($after:String!$first:Int!$input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){extendedTeamAccess(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{{ template "extended_team_access_get_vars_2" }}`,
		`{{ template "extended_team_access_response_2" }}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "custom_actions/list_extended_team_access", requests...)
	id1 := *ol.NewID(string(id1))
	trigger := ol.CustomActionsTriggerDefinition{Id: id1}

	// Act
	resp, err := trigger.ExtendedTeamAccess(client, nil)
	result := resp.Nodes

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, id1, result[0].TeamId.Id)
}
