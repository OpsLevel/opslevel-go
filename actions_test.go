package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestWebhookActionHeadersMarshalling(t *testing.T) {
	// Arrange
	unstructured := "Accept=\"application/vnd.github+json\"\nAuthorization=\"Bearer XXX\"\nContent-Type=\"application/json\"\nHost=\"app.opslevel.com\"\nOrigin=\"api.opslevel.com\""
	structured := map[string]string{
		"Host":          "app.opslevel.com",
		"Origin":        "api.opslevel.com",
		"Content-Type":  "application/json",
		"Accept":        "application/vnd.github+json",
		"Authorization": "Bearer XXX",
	}
	// Act
	marshalled := ol.CustomActionsToHeaders(unstructured)
	unmarshalled := ol.CustomActionsFromHeaders(structured)
	// Assert
	autopilot.Equals(t, structured, marshalled)
	autopilot.Equals(t, unstructured, unmarshalled)
}

func TestCreateWebhookAction(t *testing.T) {
	//Arrange
	request := `{"query": 
		"mutation($input:CustomActionsWebhookActionCreateInput!){customActionsWebhookActionCreate(input: $input){webhookAction{{ template "custom_actions_request" }},errors{message,path}}}",
		"variables":{"input":{"headers":"Content-Type=\"application/json\"","httpMethod":"POST","liquidTemplate":"{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}","name":"Deploy Rollback","webhookUrl":"https://gitlab.com/api/v4/projects/1/trigger/pipeline"}}
	}`
	response := `{"data": {"customActionsWebhookActionCreate": {
        "webhookAction": {{ template "custom_action1" }},
        "errors": []
    }}}`

	client := ABetterTestClient(t, "custom_actions/create_action", request, response)
	// Act
	trigger, err := client.CreateWebhookAction(ol.CustomActionsWebhookActionCreateInput{
		Name:           "Deploy Rollback",
		LiquidTemplate: "{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		HTTPMethod: ol.CustomActionsHttpMethodEnumPost,
		WebhookURL: "https://gitlab.com/api/v4/projects/1/trigger/pipeline",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Deploy Rollback", trigger.Name)
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
	headers := actions.Nodes[1].Headers()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, actions.TotalCount)
	autopilot.Assert(t, "MQ" == actions.PageInfo.Start, "Failed to Marshall pagination info")
	autopilot.Equals(t, "Deploy Freeze", actions.Nodes[1].Name)
	autopilot.Equals(t, "application/vnd.github+json", headers["Accept"])
}

func TestCreateTriggerDefinition(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation($input:CustomActionsTriggerDefinitionCreateInput!){customActionsTriggerDefinitionCreate(input: $input){triggerDefinition{{ template "custom_actions_trigger_request" }},errors{message,path}}}",
		"variables":{"input":{"actionId":"123456789", "description":"Disables the Deploy Freeze","filterId":"987654321","name":"Deploy Rollback","ownerId":"123456789"}}
	}`
	response := `{"data": {"customActionsTriggerDefinitionCreate": {
       "triggerDefinition": {{ template "custom_action_trigger1" }},
       "errors": []
   }}}`
	description := "Disables the Deploy Freeze"

	client := ABetterTestClient(t, "custom_actions/create_trigger", request, response)
	// Act
	trigger, err := client.CreateTriggerDefinition(ol.CustomActionsTriggerDefinitionCreateInput{
		Name:        "Deploy Rollback",
		Description: &description,
		Action:      ol.NewID("123456789"),
		Owner:       ol.NewID("123456789"),
		Filter:      ol.NewID("987654321"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Release", trigger.Name)
}

func TestGetTriggerDefinition(t *testing.T) {
	//Arrange
	request := `{"query": 
		"query($input:IdentifierInput!){account{customActionsTriggerDefinition(input: $input){{ template "custom_actions_trigger_request" }}}}",
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
	autopilot.Equals(t, "123456789", trigger.Owner.Id)
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
	autopilot.Equals(t, "123456789", triggers.Nodes[1].Owner.Id)
}
