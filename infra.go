package opslevel

import "fmt"

type InfrastructureResourceSchema struct {
	Type   string `json:"type"`
	Schema JSON   `json:"schema" scalar:"true"`
}

type InfrastructureResourceSchemaConnection struct {
	Nodes      []InfrastructureResourceSchema
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

type InfrastructureResource struct {
	Id          string      `json:"id"`
	Aliases     []string    `json:"aliases"`
	Name        string      `json:"name"`
	Type        string      `json:"type" graphql:"type @include(if: $all)"`
	Owner       EntityOwner `json:"owner" graphql:"owner @include(if: $all)"`
	OwnerLocked bool        `json:"ownerLocked" graphql:"ownerLocked @include(if: $all)"`
	Data        JSON        `json:"data" scalar:"true" graphql:"data @include(if: $all)"`
}

type InfrastructureResourceConnection struct {
	Nodes      []InfrastructureResource
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

type InfrastructureResourceSchemaInput struct {
	Type string `json:"type"`
}

type InfrastructureResourceProviderInput struct {
	AccountName  string `json:"accountName"`
	ExternalURL  string `json:"externalUrl"`
	ProviderName string `json:"providerName"`
}

type InfrastructureResourceInput struct {
	Type         *string                              `json:"providerResourceType,omitempty"`
	Schema       *InfrastructureResourceSchemaInput   `json:"schema,omitempty"`
	ProviderData *InfrastructureResourceProviderInput `json:"providerData,omitempty"`
	Owner        *ID                                  `json:"ownerId,omitempty"`
	Data         JSON                                 `json:"data,omitempty" scalar:"true"`
}

func (client *Client) CreateInfrastructure(input InfrastructureResourceInput) (*InfrastructureResource, error) {
	var m struct {
		Payload struct {
			InfrastructureResource InfrastructureResource
			Warnings               []OpsLevelWarnings
			Errors                 []OpsLevelErrors
		} `graphql:"infrastructureResourceCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
		"all":   true,
	}
	err := client.Mutate(&m, v, WithName("InfrastructureResourceCreate"))
	// TODO: handle m.Payload.Warnings somehow
	return &m.Payload.InfrastructureResource, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetInfrastructure(identifier string) (*InfrastructureResource, error) {
	var q struct {
		Account struct {
			InfrastructureResource InfrastructureResource `graphql:"infrastructureResource(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": NewIdentifier(identifier),
		"all":   true,
	}
	err := client.Query(&q, v, WithName("InfrastructureResourceGet"))
	if q.Account.InfrastructureResource.Id == "" {
		err = fmt.Errorf("InfrastructureResource with identifier '%s' not found", identifier)
	}
	return &q.Account.InfrastructureResource, HandleErrors(err, nil)
}

func (client *Client) ListInfrastructureSchemas(variables *PayloadVariables) (InfrastructureResourceSchemaConnection, error) {
	var q struct {
		Account struct {
			InfrastructureResourceSchemas InfrastructureResourceSchemaConnection `graphql:"infrastructureResourceSchemas(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("IntegrationList")); err != nil {
		return InfrastructureResourceSchemaConnection{}, err
	}
	for q.Account.InfrastructureResourceSchemas.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.InfrastructureResourceSchemas.PageInfo.End
		resp, err := client.ListInfrastructureSchemas(variables)
		if err != nil {
			return InfrastructureResourceSchemaConnection{}, err
		}
		q.Account.InfrastructureResourceSchemas.Nodes = append(q.Account.InfrastructureResourceSchemas.Nodes, resp.Nodes...)
		q.Account.InfrastructureResourceSchemas.PageInfo = resp.PageInfo
		q.Account.InfrastructureResourceSchemas.TotalCount += resp.TotalCount
	}
	return q.Account.InfrastructureResourceSchemas, nil
}

func (client *Client) ListInfrastructure(variables *PayloadVariables) (InfrastructureResourceConnection, error) {
	var q struct {
		Account struct {
			InfrastructureResource InfrastructureResourceConnection `graphql:"infrastructureResources(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
		(*variables)["all"] = true
	}
	if err := client.Query(&q, *variables, WithName("IntegrationList")); err != nil {
		return InfrastructureResourceConnection{}, err
	}
	for q.Account.InfrastructureResource.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.InfrastructureResource.PageInfo.End
		resp, err := client.ListInfrastructure(variables)
		if err != nil {
			return InfrastructureResourceConnection{}, err
		}
		q.Account.InfrastructureResource.Nodes = append(q.Account.InfrastructureResource.Nodes, resp.Nodes...)
		q.Account.InfrastructureResource.PageInfo = resp.PageInfo
		q.Account.InfrastructureResource.TotalCount += resp.TotalCount
	}
	return q.Account.InfrastructureResource, nil
}

func (client *Client) UpdateInfrastructure(identifier string, input InfrastructureResourceInput) (*InfrastructureResource, error) {
	var m struct {
		Payload struct {
			InfrastructureResource InfrastructureResource
			Warnings               []OpsLevelWarnings
			Errors                 []OpsLevelErrors
		} `graphql:"infrastructureResourceUpdate(infrastructureResource: $identifier, input: $input)"`
	}
	v := PayloadVariables{
		"$identifier": *NewIdentifier(identifier),
		"input":       input,
		"all":         true,
	}
	err := client.Mutate(&m, v, WithName("InfrastructureResourceUpdate"))
	// TODO: handle m.Payload.Warnings somehow
	return &m.Payload.InfrastructureResource, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteInfrastructure(identifier string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"infrastructureResourceDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("InfrastructureResourceDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
