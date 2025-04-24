package opslevel

import (
	"errors"
	"fmt"
)

type CustomActionsId struct {
	Aliases []string `graphql:"aliases"`
	Id      ID       `graphql:"id"`
}

func (customActionsTriggerDefinition *CustomActionsTriggerDefinition) ExtendedTeamAccess(client *Client, variables *PayloadVariables) (*TeamConnection, error) {
	var q struct {
		Account struct {
			CustomActionsTriggerDefinition struct {
				ExtendedTeamAccess TeamConnection `graphql:"extendedTeamAccess(after: $after, first: $first)"`
			} `graphql:"customActionsTriggerDefinition(input: $input)"`
		}
	}
	if customActionsTriggerDefinition.Id == "" {
		return nil, fmt.Errorf("unable to get teams with ExtendedTeamAccess, invalid CustomActionsTriggerDefinition id: '%s'", customActionsTriggerDefinition.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["input"] = *NewIdentifier(string(customActionsTriggerDefinition.Id))

	if err := client.Query(&q, *variables, WithName("ExtendedTeamAccessList")); err != nil {
		return nil, err
	}

	if q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess.PageInfo.End
		resp, err := customActionsTriggerDefinition.ExtendedTeamAccess(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess.Nodes = append(q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess.Nodes, resp.Nodes...)
		q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess.PageInfo = resp.PageInfo
		q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess.TotalCount += resp.TotalCount
	}
	return &q.Account.CustomActionsTriggerDefinition.ExtendedTeamAccess, nil
}

type CustomActionsExternalActionsConnection struct {
	Nodes      []CustomActionsExternalAction
	PageInfo   PageInfo
	TotalCount int
}

func (s *CustomActionsExternalActionsConnection) GetNodes() any {
	return s.Nodes
}

type CustomActionsTriggerDefinitionsConnection struct {
	Nodes      []CustomActionsTriggerDefinition
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

func (s *CustomActionsTriggerDefinitionsConnection) GetNodes() any {
	return s.Nodes
}

func (client *Client) CreateWebhookAction(input CustomActionsWebhookActionCreateInput) (*CustomActionsExternalAction, error) {
	var m struct {
		Payload struct { // TODO: fix this
			WebhookAction CustomActionsExternalAction
			Errors        []Error
		} `graphql:"customActionsWebhookActionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("WebhookActionCreate"))
	return &m.Payload.WebhookAction, errors.Join(err, HasAPIErrors(m.Payload.Errors))
}

func (client *Client) GetCustomAction(input string) (*CustomActionsExternalAction, error) {
	var q struct {
		Account struct {
			Action CustomActionsExternalAction `graphql:"customActionsExternalAction(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Query(&q, v, WithName("ExternalActionGet"))
	return &q.Account.Action, HandleErrors(err, IsResourceFound(&q.Account.Action.CustomActionsId))
}

func (client *Client) ListCustomActions(variables *PayloadVariables) (*CustomActionsExternalActionsConnection, error) {
	var q struct {
		Account struct {
			Actions CustomActionsExternalActionsConnection `graphql:"customActionsExternalActions(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("ExternalActionList")); err != nil {
		return nil, err
	}
	if q.Account.Actions.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Actions.PageInfo.End
		resp, err := client.ListCustomActions(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Actions.Nodes = append(q.Account.Actions.Nodes, resp.Nodes...)
		q.Account.Actions.PageInfo = resp.PageInfo
	}
	return &q.Account.Actions, nil
}

func (client *Client) UpdateWebhookAction(input CustomActionsWebhookActionUpdateInput) (*CustomActionsExternalAction, error) {
	var m struct {
		Payload struct { // TODO: fix this
			WebhookAction CustomActionsExternalAction
			Errors        []Error
		} `graphql:"customActionsWebhookActionUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("WebhookActionUpdate"))
	return &m.Payload.WebhookAction, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteWebhookAction(input string) error {
	var m struct {
		Payload BasePayload `graphql:"customActionsWebhookActionDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Mutate(&m, v, WithName("WebhookActionDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateTriggerDefinition(input CustomActionsTriggerDefinitionCreateInput) (*CustomActionsTriggerDefinition, error) {
	var m struct {
		Payload CustomActionsTriggerDefinitionCreatePayload `graphql:"customActionsTriggerDefinitionCreate(input: $input)"`
	}
	if input.AccessControl == nil {
		input.AccessControl = &CustomActionsTriggerDefinitionAccessControlEnumEveryone
	}
	if input.EntityType == nil {
		input.EntityType = &CustomActionsEntityTypeEnumService
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TriggerDefinitionCreate"))
	return &m.Payload.TriggerDefinition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetTriggerDefinition(input string) (*CustomActionsTriggerDefinition, error) {
	var q struct {
		Account struct {
			Definition CustomActionsTriggerDefinition `graphql:"customActionsTriggerDefinition(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Query(&q, v, WithName("TriggerDefinitionGet"))
	return &q.Account.Definition, errors.Join(err, IsResourceFound(&q.Account.Definition))
}

func (client *Client) ListTriggerDefinitions(variables *PayloadVariables) (*CustomActionsTriggerDefinitionsConnection, error) {
	var q struct {
		Account struct {
			Definitions CustomActionsTriggerDefinitionsConnection `graphql:"customActionsTriggerDefinitions(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("TriggerDefinitionList")); err != nil {
		return nil, HandleErrors(err)
	}
	q.Account.Definitions.TotalCount = len(q.Account.Definitions.Nodes)
	if q.Account.Definitions.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Definitions.PageInfo.End
		resp, err := client.ListTriggerDefinitions(variables)
		if err != nil {
			return &q.Account.Definitions, HandleErrors(err)
		}
		q.Account.Definitions.Nodes = append(q.Account.Definitions.Nodes, resp.Nodes...)
		q.Account.Definitions.PageInfo = resp.PageInfo
		q.Account.Definitions.TotalCount += resp.TotalCount
	}
	return &q.Account.Definitions, HandleErrors(IsResourceFound(&q.Account.Definitions))
}

func (client *Client) UpdateTriggerDefinition(input CustomActionsTriggerDefinitionUpdateInput) (*CustomActionsTriggerDefinition, error) {
	var m struct {
		Payload CustomActionsTriggerDefinitionUpdatePayload `graphql:"customActionsTriggerDefinitionUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TriggerDefinitionUpdate"))
	return &m.Payload.TriggerDefinition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteTriggerDefinition(input string) error {
	var m struct {
		Payload BasePayload `graphql:"customActionsTriggerDefinitionDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Mutate(&m, v, WithName("TriggerDefinitionDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
