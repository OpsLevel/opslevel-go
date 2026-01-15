package opslevel_test

import (
	"encoding/json"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2026"
	"github.com/rocktavious/autopilot/v2023"
)

var newID = ol.NewID("123456789")

func TestCreateWebhookAction(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation WebhookActionCreate($input:CustomActionsWebhookActionCreateInput!){customActionsWebhookActionCreate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}`,
		`{"input":{"headers":"{\"Content-Type\":\"application/json\"}","httpMethod":"POST",{{ template "liquid_template_rollback" }},"name":"Deploy Rollback","webhookUrl":"https://gitlab.com/api/v4/projects/1/trigger/pipeline"}}`,
		`{"data": {"customActionsWebhookActionCreate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/create_action", testRequest)

	// Act
	jsonHeaders, err := ol.NewJSON(`{"Content-Type": "application/json"}`)
	autopilot.Ok(t, err)
	action, err := client.CreateWebhookAction(ol.CustomActionsWebhookActionCreateInput{
		Name:           "Deploy Rollback",
		LiquidTemplate: ol.RefOf("{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}"),
		Headers:        jsonHeaders,
		HttpMethod:     ol.CustomActionsHttpMethodEnumPost,
		WebhookUrl:     "https://gitlab.com/api/v4/projects/1/trigger/pipeline",
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestListCustomActions(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ExternalActionList($after:String!$first:Int!){account{customActionsExternalActions(after: $after, first: $first){nodes{{ template "custom_actions_request" }},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "customActionsExternalActions": { "nodes": [ { {{ template "custom_action1_response" }} }, { {{ template "custom_action2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ExternalActionList($after:String!$first:Int!){account{customActionsExternalActions(after: $after, first: $first){nodes{{ template "custom_actions_request" }},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "customActionsExternalActions": { "nodes": [ { {{ template "custom_action3_response" }} } ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "custom_actions/list_actions", requests...)
	// Act
	response, err := client.ListCustomActions(nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, "Deploy Freeze", result[1].Name)
	autopilot.Equals(t, "Page On-Call", result[2].Name)
	autopilot.Equals(t, "application/json", result[0].Headers["Content-Type"])
}

func TestUpdateWebhookAction(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation WebhookActionUpdate($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}`,
		`{"input":{"id": "123456789", "httpMethod":"PUT"}}`,
		`{"data": {"customActionsWebhookActionUpdate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_action", testRequest)

	// Act
	action, err := client.UpdateWebhookAction(ol.CustomActionsWebhookActionUpdateInput{
		Id:         *newID,
		HttpMethod: &ol.CustomActionsHttpMethodEnumPut,
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestUpdateWebhookAction2(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation WebhookActionUpdate($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}`,
		`{"input":{"id": "123456789","description":"","headers":"{\"Accept\":\"application/json\"}"}}`,
		`{"data": {"customActionsWebhookActionUpdate": { "webhookAction": {{ template "custom_action1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/update_action2", testRequest)

	// Act
	jsonHeaders, err := ol.NewJSON(`{"Accept": "application/json"}`)
	autopilot.Ok(t, err)
	action, err := client.UpdateWebhookAction(ol.CustomActionsWebhookActionUpdateInput{
		Id:          *newID,
		Description: ol.RefOf(""),
		Headers:     jsonHeaders,
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestDeleteWebhookAction(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation WebhookActionDelete($input:IdentifierInput!){customActionsWebhookActionDelete(resource: $input){errors{message,path}}}`,
		`{"input":{"alias": "123456789"}}`,
		`{"data": {"customActionsWebhookActionDelete": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/delete_action", testRequest)

	// Act
	err := client.DeleteWebhookAction("123456789")

	// Assert
	autopilot.Ok(t, err)
}

func TestCreateTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}`,
		`{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"SERVICE","filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/create_trigger", testRequest)

	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:                   "Deploy Rollback",
		AccessControl:          &ol.CustomActionsTriggerDefinitionAccessControlEnumEveryone,
		Description:            ol.RefOf("Disables the Deploy Freeze"),
		ResponseTemplate:       ol.RefOf(""),
		ManualInputsDefinition: ol.RefOf(""),
		ActionId:               ol.RefOf(*newID),
		OwnerId:                *newID,
		FilterId:               ol.RefOf(ol.ID("987654321")),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestCreateTriggerDefinitionWithGlobalEntityType(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}`,
		`{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"GLOBAL","extendedTeamAccess":[{"alias":"example_1"},{"alias":"example_1"}],"filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/create_trigger_with_global_entity", testRequest)

	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:                   "Deploy Rollback",
		AccessControl:          &ol.CustomActionsTriggerDefinitionAccessControlEnumEveryone,
		Description:            ol.RefOf("Disables the Deploy Freeze"),
		ActionId:               ol.RefOf(*newID),
		ManualInputsDefinition: ol.RefOf(""),
		ResponseTemplate:       ol.RefOf(""),
		OwnerId:                *newID,
		FilterId:               ol.RefOf(ol.ID("987654321")),
		EntityType:             &ol.CustomActionsEntityTypeEnumGlobal,
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
	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionCreate($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}`,
		`{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","entityType":"SERVICE","extendedTeamAccess":[],"filterId":"987654321","manualInputsDefinition":"", "name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}`,
		`{"data": {"customActionsTriggerDefinitionCreate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "custom_actions/create_trigger_with_null_extended_teams", testRequest)
	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:                   "Deploy Rollback",
		Description:            ol.RefOf("Disables the Deploy Freeze"),
		AccessControl:          &ol.CustomActionsTriggerDefinitionAccessControlEnumEveryone,
		ManualInputsDefinition: ol.RefOf(""),
		ResponseTemplate:       ol.RefOf(""),
		ActionId:               ol.RefOf(*newID),
		OwnerId:                *newID,
		FilterId:               ol.RefOf(ol.ID("987654321")),
		ExtendedTeamAccess:     &[]ol.IdentifierInput{},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestGetTriggerDefinition(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query TriggerDefinitionGet($input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){{ template "custom_actions_trigger_request" }}}}`,
		`{"input":{"alias":"123456789"}}`,
		`{"data": {"account": { "customActionsTriggerDefinition": {{ template "custom_action_trigger2" }} }}}`,
	)
	client := BestTestClient(t, "custom_actions/get_trigger", testRequest)

	// Act
	trigger, err := client.GetTriggerDefinition("123456789")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
	autopilot.Equals(t, "Uses Ruby", trigger.Filter.Name)
	autopilot.Equals(t, *newID, trigger.Owner.Id)
}

func TestListTriggerDefinitions(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query TriggerDefinitionList($after:String!$first:Int!){account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{{ template "custom_actions_trigger_request" }},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "customActionsTriggerDefinitions": { "nodes": [ { {{ template "custom_action_trigger1_response" }} }, { {{ template "custom_action_trigger2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query TriggerDefinitionList($after:String!$first:Int!){account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{{ template "custom_actions_trigger_request" }},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "customActionsTriggerDefinitions": { "nodes": [ { {{ template "custom_action_trigger3_response" }} } ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

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
	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}`,
		`{"input":{"id":"123456789", "filterId":null}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_trigger", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:       *newID,
		FilterId: ol.NewNullOf[ol.ID](),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestUpdateTriggerDefinition2(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}`,
		`{"input":{"id":"123456789", "name":"test", "description": "", "extendedTeamAccess": []}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_trigger2", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:                 *newID,
		Name:               ol.RefOf("test"),
		Description:        ol.RefOf(""),
		ExtendedTeamAccess: &[]ol.IdentifierInput{},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestUpdateTriggerDefinition3(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionUpdate($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}`,
		`{"input":{"id":"123456789", "name":"test", "description": "", "extendedTeamAccess": [{"alias": "123456789"}, { {{ template "id1" }} } ]}}`,
		`{"data": {"customActionsTriggerDefinitionUpdate": { "triggerDefinition": {{ template "custom_action_trigger1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "custom_actions/update_trigger3", testRequest)

	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:          *newID,
		Name:        ol.RefOf("test"),
		Description: ol.RefOf(""),
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
	request := autopilot.NewTestRequest(
		`mutation TriggerDefinitionDelete($input:IdentifierInput!){customActionsTriggerDefinitionDelete(resource: $input){errors{message,path}}}`,
		`{"input":{"alias":"123456789"}}`,
		`{"data": {"customActionsTriggerDefinitionDelete": { "errors": [{{ template "error1" }}] }}}`,
	)

	testRequest := autopilot.NewTestRequest(
		`mutation TriggerDefinitionDelete($input:IdentifierInput!){customActionsTriggerDefinitionDelete(resource: $input){errors{message,path}}}`,
		`{"input":{"alias":"123456789"}}`,
		`{"data": {"customActionsTriggerDefinitionDelete": { "errors": [] }}}`,
	)
	testRequestError := autopilot.NewTestRequest(
		testRequest.Request.Query,
		ol.JSON(testRequest.Request.Variables).ToJSON(),
		`{"data": {"customActionsTriggerDefinitionDelete": { "errors": [{{ template "error1" }}] }}}`,
	)

	client := BestTestClient(t, "custom_actions/delete_trigger", testRequest)
	clientErr := BestTestClient(t, "custom_actions/delete_trigger_err", testRequestError)
	clientErr2 := BestTestClient(t, "custom_actions/delete_trigger_err2", request)
	// Act
	err := client.DeleteTriggerDefinition("123456789")
	err2 := clientErr.DeleteTriggerDefinition("123456789")
	err3 := clientErr2.DeleteTriggerDefinition("123456789")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, `OpsLevel API Errors:
	- 'one.two.three' Example Error`,
		err2.Error())
	autopilot.Assert(t, err3 != nil, "Expected error was not thrown")
}

func TestListExtendedTeamAccess(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ExtendedTeamAccessList($after:String!$first:Int!$input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){extendedTeamAccess(after: $after, first: $first){nodes{alias,id,aliases,managedAliases,contacts{address,displayName,displayType,externalId,id,isDefault,type},htmlUrl,manager{id,email,name,contacts{address,displayName,displayType,externalId,id,isDefault,type},htmlUrl,provisionedBy,role},memberships{nodes{role,team{alias,id},user{id,email,name}},{{ template "pagination_request" }}},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }}}},{{ template "pagination_request" }}}}}}`,
		`{{ template "extended_team_access_get_vars_1" }}`,
		`{{ template "extended_team_access_response_1" }}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ExtendedTeamAccessList($after:String!$first:Int!$input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){extendedTeamAccess(after: $after, first: $first){nodes{alias,id,aliases,managedAliases,contacts{address,displayName,displayType,externalId,id,isDefault,type},htmlUrl,manager{id,email,name,contacts{address,displayName,displayType,externalId,id,isDefault,type},htmlUrl,provisionedBy,role},memberships{nodes{role,team{alias,id},user{id,email,name}},{{ template "pagination_request" }}},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }}}},{{ template "pagination_request" }}}}}}`,
		`{{ template "extended_team_access_get_vars_2" }}`,
		`{{ template "extended_team_access_response_2" }}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "custom_actions/list_extended_team_access", requests...)
	id1 := *ol.NewID(string(id1))
	trigger := ol.CustomActionsTriggerDefinition{Id: id1}

	// Act
	resp, err := trigger.ExtendedTeamAccess(client, nil)
	if resp == nil || resp.Nodes == nil || len(resp.Nodes) == 0 {
		t.Error("Expected team access response to not be nil")
	}
	autopilot.Ok(t, err)
	result := resp.Nodes

	// Assert
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, id1, result[0].Id)
}

func TestApprovalConfigInput(t *testing.T) {
	// Arrange
	v1 := ol.ApprovalConfigInput{}
	v2 := ol.ApprovalConfigInput{Teams: &[]ol.IdentifierInput{}}
	// Act
	j1, err1 := json.Marshal(v1)
	j2, err2 := json.Marshal(v2)
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Ok(t, err2)
	autopilot.Equals(t, `{}`, string(j1))
	autopilot.Equals(t, `{"teams":[]}`, string(j2))
	// Assert
}
