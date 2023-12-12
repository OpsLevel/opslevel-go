package opslevel

import "fmt"

type PropertyDefinitionInput struct {
	Name   string     `json:"name" default:"John Doe"`
	Schema JSONString `json:"schema,string" default:"{\"$schema\":\"https://json-schema.org/draft/2020-12/schema\",\"title\":\"Packages\",\"description\":\"A list of packages.\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"version\":{\"type\":\"string\"},\"lock_file\":{\"type\":\"string\"},\"manager\":{\"type\":\"string\"},},\"required\":[\"name\",\"version\"],},\"minItems\":0,\"uniqueItems\":true}"`
}

type PropertyDefinition struct {
	Aliases []string `graphql:"aliases" json:"aliases"`
	Id      ID       `graphql:"id" json:"id"`
	Name    string   `graphql:"name" json:"name"`
	Schema  JSON     `json:"schema" scalar:"true"`
}

type PropertyDefinitionConnection struct {
	Nodes      []PropertyDefinition
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
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

func (client *Client) GetPropertyDefinition(input string) (*PropertyDefinition, error) {
	var q struct {
		Account struct {
			Definition PropertyDefinition `graphql:"propertyDefinition(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Query(&q, v, WithName("PropertyDefinitionGet"))
	if q.Account.Definition.Id == "" {
		err = fmt.Errorf("PropertyDefinition with ID or Alias matching '%s' not found", input)
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
		q.Account.Definitions.TotalCount += len(q.Account.Definitions.Nodes)
	}
	return q.Account.Definitions, nil
}

func (client *Client) DeletePropertyDefinition(input string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"propertyDefinitionDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Mutate(&m, v, WithName("PropertyDefinitionDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
