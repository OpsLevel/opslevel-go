package opslevel

import "fmt"

type PropertyDefinitionInput struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type PropertyDefinition struct {
	Aliases []string `graphql:"aliases"`
	Id      ID       `graphql:"id"`
	Name    string   `graphql:"name"`
	Schema  string   `graphql:"schema"`
}

type PropertyDefinitionConnection struct {
	Nodes      []PropertyDefinition
	PageInfo   PageInfo
	TotalCount int
}

func (client *Client) CreatePropertyDefinition(input PropertyDefinitionInput) (*PropertyDefinition, error) {
	var m struct {
		Payload struct {
			Definition PropertyDefinition `graphql:"definition"`
			Errors     []OpsLevelErrors   `graphql:"errors"`
		} `graphql:"propertyDefinitionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("PropertyDefinitionCreate"))
	return &m.Payload.Definition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetPropertyDefinition(input IdentifierInput) (*PropertyDefinition, error) {
	var q struct {
		Account struct {
			Definition PropertyDefinition `graphql:"propertyDefinition(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Query(&q, v, WithName("PropertyDefinitionGet"))
	if q.Account.Definition.Id == "" {
		err = fmt.Errorf("PropertyDefinition with ID '%s' or Alias '%s' not found", input.Id, input.Alias)
	}
	return &q.Account.Definition, HandleErrors(err, nil)
}

func (client *Client) ListPropertyDefinitions(variables *PayloadVariables) (PropertyDefinitionConnection, error) {
	var q struct {
		Account struct {
			Definitions PropertyDefinitionConnection `graphql:"propertyDefinitions(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("PropertyDefinitionList")); err != nil {
		return PropertyDefinitionConnection{}, err
	}
	for q.Account.Definitions.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Definitions.PageInfo.End
		resp, err := client.ListPropertyDefinitions(variables)
		if err != nil {
			return PropertyDefinitionConnection{}, err
		}
		q.Account.Definitions.Nodes = append(q.Account.Definitions.Nodes, resp.Nodes...)
		q.Account.Definitions.PageInfo = resp.PageInfo
		q.Account.Definitions.TotalCount += resp.TotalCount
	}
	return q.Account.Definitions, nil
}

func (client *Client) DeletePropertyDefinition(input IdentifierInput) error {
	var m struct {
		Payload struct {
			DeletedAlias string           `graphql:"deletedAlias"`
			DeletedID    ID               `graphql:"deletedId"`
			Errors       []OpsLevelErrors `graphql:"errors"`
		} `graphql:"propertyDefinitionDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("PropertyDefinitionDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
