package opslevel

type InfrastructureResourceSchema struct {
	Type   string `json:"type"`
	Schema JSON   `json:"schema"`
}

type InfrastructureResourceSchemaConnection struct {
	Nodes      []InfrastructureResourceSchema
	PageInfo   PageInfo
	TotalCount int
}

type InfrastructureResource struct {
	Id          string      `json:"id"`
	Aliases     []string    `json:"aliases"`
	Name        string      `json:"name"`
	Type        string      `json:"type" graphql:"@include(if: $all)"`
	Owner       EntityOwner `json:"owner" graphql:"@include(if: $all)"`
	OwnerLocked bool        `json:"ownerLocked" graphql:"@include(if: $all)"`
	Data        JSON        `json:"data" graphql:"@include(if: $all)"`
}

type InfrastructureResourceConnection struct {
	Nodes      []InfrastructureResource
	PageInfo   PageInfo
	TotalCount int
}

type InfrastructureResourceSchemaInput struct {
	Type string `json:"type"`
}

type InfrastructureResourceProviderInput struct {
	AccountName  string `json:"accountName"`
	ExternalURL  string `json:"external_url"`
	ProviderName string `json:"providerName"`
}

type InfrastructureResourceInput struct {
	Name         *string                              `json:"name,omitempty"`
	Type         *string                              `json:"providerResourceType,omitempty"`
	Schema       *InfrastructureResourceSchemaInput   `json:"schema,omitempty"`
	ProviderData *InfrastructureResourceProviderInput `json:"providerData,omitempty"`
	OwnerId      *ID                                  `json:"ownerId,omitempty"`
	Data         JSON                                 `json:"data,omitempty"`
}

func (client *Client) CreateInfrastructure(input InfrastructureResourceInput) (*InfrastructureResource, error) {
	var m struct {
		Payload struct {
			InfrastructureResource *InfrastructureResource
			Warnings               []OpsLevelWarnings
			Errors                 []OpsLevelErrors
		} `graphql:"infrastructureResourceCreate(input: $input)"`
	}

	return &InfrastructureResource{}, nil
}

func (client *Client) GetInfrastructure(identifier string) (*InfrastructureResource, error) {
	var q struct {
		Account struct {
			InfrastructureResource InfrastructureResource `graphql:"infrastructureResource(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": NewIdentifier(identifier),
	}
	return &InfrastructureResource{}, nil
}

func (client *Client) ListInfrastructureSchemas(variables *PayloadVariables) (InfrastructureResourceSchemaConnection, error) {
	var q struct {
		Account struct {
			InfrastructureResourceSchemas InfrastructureResourceSchemaConnection `graphql:"infrastructureResourceSchemas(after: $after, first: $first)"`
		}
	}
	return InfrastructureResourceSchemaConnection{}, nil
}

func (client *Client) ListInfrastructure(variables *PayloadVariables) (InfrastructureResourceConnection, error) {
	var q struct {
		Account struct {
			InfrastructureResource InfrastructureResourceConnection `graphql:"infrastructureResource(after: $after, first: $first)"`
		}
	}
	return InfrastructureResourceConnection{}, nil
}

func (client *Client) UpdateInfrastructure(identifier string, input InfrastructureResourceInput) (*InfrastructureResource, error) {
	var m struct {
		Payload struct {
			InfrastructureResource *InfrastructureResource
			Warnings               []OpsLevelWarnings
			Errors                 []OpsLevelErrors
		} `graphql:"infrastructureResourceUpdate(infrastructureResource: $identifier, input: $input)"`
	}

	return &InfrastructureResource{}, nil
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
	return nil
}
