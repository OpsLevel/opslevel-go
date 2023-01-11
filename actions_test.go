package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestCreateWebhookAction(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ($input:CustomActionsWebhookActionCreateInput!){customActionsWebhookActionCreate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}",
		"variables":{"input":{"headers":"{\"Content-Type\":\"application/json\"}","httpMethod":"POST","liquidTemplate":"{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}","name":"Deploy Rollback","webhookUrl":"https://gitlab.com/api/v4/projects/1/trigger/pipeline"}}
	}`
	response := `{"data": {"customActionsWebhookActionCreate": {
      "webhookAction": {{ template "custom_action1" }},
      "errors": []
  }}}`

	//fmt.Print(Templated(request))
	//fmt.Print(Templated(response))
	//panic(1)

	client := ABetterTestClient(t, "custom_actions/create_action", request, response)

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
	//Arrange
	request := `{"query":
		"{account{customActionsExternalActions(after: $after, first: $first){nodes{{ template "custom_actions_request" }},{{ template "pagination_request" }}}}}",
		"variables":{}
	}`
	response := `{"data": {"account": {
      "customActionsExternalActions": {
        "nodes": [
          {{ template "custom_action1" }},
          {{ template "custom_action2" }}
        ],
        {{ template "no_pagination_response" }},
        "totalCount": 2
      }
  }}}`

	// An easy way to see the results of templating is by uncommenting this
	//fmt.Print(Templated(request))
	//fmt.Print(Templated(response))
	//panic(1)

	client := ABetterTestClient(t, "custom_actions/list_actions", request, response)
	// Act
	actions, err := client.ListCustomActions(nil)
	headers := actions.Nodes[1].Headers
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, actions.TotalCount)
	autopilot.Assert(t, "MQ" == actions.PageInfo.Start, "Failed to Marshall pagination info")
	autopilot.Equals(t, "Deploy Freeze", actions.Nodes[1].Name)
	autopilot.Equals(t, "application/vnd.github+json", headers["Accept"])
}

func TestUpdateWebhookAction(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}",
		"variables":{"input":{"id": "123456789", "httpMethod":"PUT"}}
	}`
	response := `{"data": {"customActionsWebhookActionUpdate": {
     "webhookAction": {{ template "custom_action1" }},
     "errors": []
 }}}`

	client := ABetterTestClient(t, "custom_actions/update_action", request, response)

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
	//Arrange
	request := `{"query":
		"mutation ($input:CustomActionsWebhookActionUpdateInput!){customActionsWebhookActionUpdate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}",
		"variables":{"input":{"id": "123456789","description":"","headers":"{\"Accept\":\"application/json\"}"}}
	}`
	response := `{"data": {"customActionsWebhookActionUpdate": {
     "webhookAction": {{ template "custom_action1" }},
     "errors": []
 }}}`

	client := ABetterTestClient(t, "custom_actions/update_action2", request, response)
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
	//Arrange
	request := `{"query":
		"mutation ($input:IdentifierInput!){customActionsWebhookActionDelete(resource: $input){deletedAlias,deletedId,errors{message,path}}}",
		"variables":{"input":{"id": "123456789"}}
	}`
	response := `{"data": {"customActionsWebhookActionDelete": {
     "deletedId": "123456789",
     "deletedAlias": "",
     "errors": []
 }}}`

	client := ABetterTestClient(t, "custom_actions/delete_action", request, response)

	// Act
	err := client.DeleteWebhookAction(ol.IdentifierInput{
		Id: "123456789",
	})

	// Assert
	autopilot.Ok(t, err)
}

func TestCreateTriggerDefinition(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}",
		"variables":{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","filterId":"987654321","name":"Deploy Rollback","ownerId":"123456789", "accessControl": "everyone", "responseTemplate": ""}}
	}`
	response := `{"data": {"customActionsTriggerDefinitionCreate": {
     "triggerDefinition": {{ template "custom_action_trigger1" }},
     "errors": []
 }}}`

	client := ABetterTestClient(t, "custom_actions/create_trigger", request, response)
	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:        "Deploy Rollback",
		Description: ol.NewString("Disables the Deploy Freeze"),
		Action:      ol.NewID("123456789"),
		Owner:       "123456789",
		Filter:      ol.NewID("987654321"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestGetTriggerDefinition(t *testing.T) {
	//Arrange
	request := `{"query":
		"query ($input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){{ template "custom_actions_trigger_request" }}}}",
		"variables":{"input":{"id":"123456789"}}
	}`
	response := `{"data": {"account": {
      "customActionsTriggerDefinition": {{ template "custom_action_trigger2" }}
  }}}`

	client := ABetterTestClient(t, "custom_actions/get_trigger", request, response)
	// Act
	trigger, err := client.GetTriggerDefinition(ol.IdentifierInput{Id: "123456789"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
	autopilot.Equals(t, "Uses Ruby", trigger.Filter.Name)
	autopilot.Equals(t, "123456789", string(trigger.Owner.Id))
}

func TestListTriggerDefinitions(t *testing.T) {
	//Arrange
	request := `{"query":
		"{account{customActionsTriggerDefinitions(after: $after, first: $first){nodes{{ template "custom_actions_trigger_request" }},{{ template "pagination_request" }}}}}",
		"variables":{}
	}`
	response := `{"data": {"account": {
      "customActionsTriggerDefinitions": {
        "nodes": [
          {{ template "custom_action_trigger1" }},
          {{ template "custom_action_trigger2" }}
        ],
        {{ template "no_pagination_response" }},
        "totalCount": 2
      }
  }}}`

	client := ABetterTestClient(t, "custom_actions/list_triggers", request, response)
	// Act
	triggers, err := client.ListTriggerDefinitions(nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, triggers.TotalCount)
	autopilot.Equals(t, "Release", triggers.Nodes[1].Name)
	autopilot.Equals(t, "Uses Ruby", triggers.Nodes[1].Filter.Name)
	autopilot.Equals(t, "123456789", string(triggers.Nodes[1].Owner.Id))
}

func TestUpdateTriggerDefinition(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}",
		"variables":{"input":{"id":"123456789", "filterId":null, "accessControl": "everyone"}}
	}`
	response := `{"data": {"customActionsTriggerDefinitionUpdate": {
     "triggerDefinition": {{ template "custom_action_trigger1" }},
     "errors": []
 }}}`

	client := ABetterTestClient(t, "custom_actions/update_trigger", request, response)
	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:     "123456789",
		Filter: ol.NullString(),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestUpdateTriggerDefinition2(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ($input:CustomActionsTriggerDefinitionUpdateInput!){customActionsTriggerDefinitionUpdate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}",
		"variables":{"input":{"id":"123456789", "name":"test", "description": "", "accessControl": "everyone"}}
	}`
	response := `{"data": {"customActionsTriggerDefinitionUpdate": {
     "triggerDefinition": {{ template "custom_action_trigger1" }},
     "errors": []
 }}}`

	client := ABetterTestClient(t, "custom_actions/update_trigger2", request, response)
	// Act
	trigger, err := client.UpdateTriggerDefinition(ol.CustomActionsTriggerDefinitionUpdateInput{
		Id:          "123456789",
		Name:        ol.NewString("test"),
		Description: ol.NewString(""),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestDeleteTriggerDefinition(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ($input:IdentifierInput!){customActionsTriggerDefinitionDelete(resource: $input){deletedAlias,deletedId,errors{message,path}}}",
		"variables":{"input":{"id":"123456789"}}
	}`
	response := `{"data": {"customActionsTriggerDefinitionDelete": {
     "deletedAlias": "",
     "deletedId": "123456789",
     "errors": []
 }}}`
	responseErr := `{"data": {"customActionsTriggerDefinitionDelete": {
     "deletedAlias": null,
     "deletedId": null,
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
	autopilot.Equals(t, "OpsLevel API Errors:\n\t* Example Error", err2.Error())
	autopilot.Assert(t, err3 != nil, "Expected error was not thrown")
}
