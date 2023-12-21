package opslevel

import "fmt"

type PropertyDefinitionInput struct {
	Name                  string                    `json:"name,omitempty" yaml:"name,omitempty" default:"Example Package Schema"`
	Description           string                    `json:"description,omitempty" yaml:"description,omitempty" default:"Place description here"`
	Schema                JSON                      `json:"schema,omitempty" yaml:"schema,omitempty" default:"{\"$schema\":\"https://json-schema.org/draft/2020-12/schema\",\"title\":\"Packages\",\"description\":\"A list of packages.\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"version\":{\"type\":\"string\"},\"lock_file\":{\"type\":\"string\"},\"manager\":{\"type\":\"string\"}},\"required\":[\"name\",\"version\"]},\"minItems\":0,\"uniqueItems\":true}"`
	PropertyDisplayStatus PropertyDisplayStatusEnum `json:"propertyDisplayStatus,omitempty" yaml:"propertyDisplayStatus,omitempty" default:"visible"`
}

type PropertyDefinition struct {
	Aliases               []string                          `graphql:"aliases" json:"aliases"`
	Id                    ID                                `graphql:"id" json:"id"`
	Name                  string                            `graphql:"name" json:"name"`
	Description           string                            `graphql:"description" json:"description"`
	DisplaySubtype        PropertyDefinitionDisplayTypeEnum `graphql:"displaySubtype" json:"displaySubtype"`
	DisplayType           PropertyDefinitionDisplayTypeEnum `graphql:"displayType" json:"displayType"`
	PropertyDisplayStatus PropertyDisplayStatusEnum         `graphql:"propertyDisplayStatus" json:"propertyDisplayStatus"`
	Schema                JSON                              `json:"schema" scalar:"true"`
}

type PropertyDefinitionConnection struct {
	Nodes      []PropertyDefinition
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

type PropertyDefinitionId struct {
	Id      ID       `json:"id"`
	Aliases []string `json:"aliases,omitempty"`
}

type PropertyInput struct {
	Owner         IdentifierInput `json:"owner"`
	Definition    IdentifierInput `json:"definition"`
	Value         JSONString      `json:"value"`
	RunValidation *bool           `json:"runValidation,omitempty"`
}

type Property struct {
	Definition       PropertyDefinitionId `graphql:"definition"`
	Owner            EntityOwnerService   `graphql:"owner"`
	ValidationErrors []OpsLevelErrors     `graphql:"validationErrors"`
	Value            *JSONString          `graphql:"value"`
}

type ServicePropertiesConnection struct {
	Nodes      []Property
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

func (client *Client) UpdatePropertyDefinition(identifier string, input PropertyDefinitionInput) (*PropertyDefinition, error) {
	var m struct {
		Payload struct {
			Definition PropertyDefinition `graphql:"definition"`
			Errors     []OpsLevelErrors   `graphql:"errors"`
		} `graphql:"propertyDefinitionUpdate(propertyDefinition: $propertyDefinition, input: $input)"`
	}
	v := PayloadVariables{
		"propertyDefinition": *NewIdentifier(identifier),
		"input":              input,
	}
	err := client.Mutate(&m, v, WithName("PropertyDefinitionUpdate"))
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

func (client *Client) GetProperty(owner string, definition string) (*Property, error) {
	var q struct {
		Account struct {
			Property Property `graphql:"property(owner: $owner, definition: $definition)"`
		}
	}
	v := PayloadVariables{
		"owner":      *NewIdentifier(owner),
		"definition": *NewIdentifier(definition),
	}
	err := client.Query(&q, v, WithName("PropertyGet"))
	if q.Account.Property.Definition.Id == "" {
		err = fmt.Errorf("Property with ID or alias matching '%s' on Service with ID or alias matching '%s' not found", definition, owner)
	}
	return &q.Account.Property, HandleErrors(err, nil)
}

func (client *Client) PropertyAssign(input PropertyInput) (*Property, error) {
	var m struct {
		Payload struct {
			Property Property         `graphql:"property"`
			Errors   []OpsLevelErrors `graphql:"errors"`
		} `graphql:"propertyAssign(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("PropertyAssign"))
	return &m.Payload.Property, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) PropertyUnassign(owner string, definition string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"propertyUnassign(owner: $owner, definition: $definition)"`
	}
	v := PayloadVariables{
		"owner":      *NewIdentifier(owner),
		"definition": *NewIdentifier(definition),
	}
	err := client.Mutate(&m, v, WithName("PropertyUnassign"))
	return HandleErrors(err, m.Payload.Errors)
}

func (service *Service) GetProperties(client *Client, variables *PayloadVariables) (*ServicePropertiesConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Properties ServicePropertiesConnection `graphql:"properties(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}

	if service.Id == "" {
		return nil, fmt.Errorf("Unable to get properties, invalid Service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServicePropertiesList")); err != nil {
		return nil, err
	}
	if service.Properties == nil {
		service.Properties = &ServicePropertiesConnection{}
	}
	service.Properties.Nodes = append(service.Properties.Nodes, q.Account.Service.Properties.Nodes...)
	service.Properties.PageInfo = q.Account.Service.Properties.PageInfo
	for service.Properties.PageInfo.HasNextPage {
		(*variables)["after"] = service.Properties.PageInfo.End
		_, err := service.GetProperties(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return service.Properties, nil
}
