package opslevel

import "fmt"

// PropertyDefinition represents the definition of a property.
type PropertyDefinition struct {
	Aliases               []string                          `graphql:"aliases" json:"aliases"`
	AllowedInConfigFiles  bool                              `graphql:"allowedInConfigFiles"` // Whether or not the property is allowed to be set in opslevel.yml config files.
	Id                    ID                                `graphql:"id" json:"id"`
	Name                  string                            `graphql:"name" json:"name"`
	Description           string                            `graphql:"description" json:"description"`
	DisplaySubtype        PropertyDefinitionDisplayTypeEnum `graphql:"displaySubtype" json:"displaySubtype"`
	DisplayType           PropertyDefinitionDisplayTypeEnum `graphql:"displayType" json:"displayType"`
	PropertyDisplayStatus PropertyDisplayStatusEnum         `graphql:"propertyDisplayStatus" json:"propertyDisplayStatus"`
	LockedStatus          PropertyLockedStatusEnum          `graphql:"lockedStatus" json:"lockedStatus"`
	Schema                JSONSchema                        `json:"schema" scalar:"true"`
}

type PropertyDefinitionId struct {
	Id      ID       `json:"id"`
	Aliases []string `json:"aliases,omitempty"`
}

// Property represents a custom property value assigned to an entity.
type Property struct {
	Definition       PropertyDefinitionId `graphql:"definition"`
	Locked           bool                 `graphql:"locked"`
	Owner            EntityOwnerService   `graphql:"owner"`
	ValidationErrors []Error              `graphql:"validationErrors"`
	Value            *JsonString          `graphql:"value"`
}

type ServicePropertiesConnection struct {
	Nodes      []Property
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

func (client *Client) CreatePropertyDefinition(input PropertyDefinitionInput) (*PropertyDefinition, error) {
	var m struct {
		Payload PropertyDefinitionPayload `graphql:"propertyDefinitionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("PropertyDefinitionCreate"))
	return &m.Payload.Definition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdatePropertyDefinition(identifier string, input PropertyDefinitionInput) (*PropertyDefinition, error) {
	var m struct {
		Payload PropertyDefinitionPayload `graphql:"propertyDefinitionUpdate(propertyDefinition: $propertyDefinition, input: $input)"`
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

func (client *Client) ListPropertyDefinitions(variables *PayloadVariables) (*PropertyDefinitionConnection, error) {
	var q struct {
		Account struct {
			Definitions PropertyDefinitionConnection `graphql:"propertyDefinitions(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("PropertyDefinitionList")); err != nil {
		return nil, err
	}
	if q.Account.Definitions.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Definitions.PageInfo.End
		resp, err := client.ListPropertyDefinitions(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Definitions.Nodes = append(q.Account.Definitions.Nodes, resp.Nodes...)
		q.Account.Definitions.PageInfo = resp.PageInfo
		q.Account.Definitions.TotalCount += len(q.Account.Definitions.Nodes)
	}
	return &q.Account.Definitions, nil
}

func (client *Client) DeletePropertyDefinition(input string) error {
	var m struct {
		Payload BasePayload `graphql:"propertyDefinitionDelete(resource: $input)"`
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
	return &q.Account.Property, HandleErrors(err, nil)
}

func (client *Client) PropertyAssign(input PropertyInput) (*Property, error) {
	var m struct {
		Payload PropertyPayload `graphql:"propertyAssign(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("PropertyAssign"))
	return &m.Payload.Property, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) PropertyUnassign(owner string, definition string) error {
	var m struct {
		Payload BasePayload `graphql:"propertyUnassign(owner: $owner, definition: $definition)"`
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
		return nil, fmt.Errorf("unable to get properties, invalid Service id: '%s'", service.Id)
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
	if service.Properties.PageInfo.HasNextPage {
		(*variables)["after"] = service.Properties.PageInfo.End
		resp, err := service.GetProperties(client, variables)
		if err != nil {
			return nil, err
		}
		service.Properties.TotalCount += resp.TotalCount
	}
	return service.Properties, nil
}
