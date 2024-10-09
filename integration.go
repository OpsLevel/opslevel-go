package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
)

type IntegrationId struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type AWSIntegrationFragment struct {
	IAMRole              string   `graphql:"iamRole"`
	ExternalID           string   `graphql:"externalId"`
	OwnershipTagOverride bool     `graphql:"awsTagsOverrideOwnership"`
	OwnershipTagKeys     []string `graphql:"ownershipTagKeys"`
	RegionOverride       []string `graphql:"regionOverride"`
}

type AzureResourcesIntegrationFragment struct {
	Aliases               []string `graphql:"aliases"`
	OwnershipTagKeys      []string `graphql:"ownershipTagKeys"`
	SubscriptionId        string   `graphql:"subscriptionId"`
	TagsOverrideOwnership bool     `graphql:"tagsOverrideOwnership"`
	TenantId              string   `graphql:"tenantId"`
}

type GoogleCloudProject struct {
	ID   string `graphql:"id"`
	Name string `graphql:"name"`
	URL  string `graphql:"url"`
}

type GoogleCloudIntegrationFragment struct {
	Aliases               []string             `graphql:"aliases"`
	ClientEmail           string               `graphql:"clientEmail"`
	OwnershipTagKeys      []string             `graphql:"ownershipTagKeys"`
	Projects              []GoogleCloudProject `graphql:"projects"`
	TagsOverrideOwnership bool                 `graphql:"tagsOverrideOwnership"`
}

type NewRelicIntegrationFragment struct {
	BaseUrl    string `graphql:"baseUrl"`
	AccountKey string `graphql:"accountKey"`
}

type IntegrationConnection struct {
	Nodes      []Integration
	PageInfo   PageInfo
	TotalCount int
}

type AWSIntegrationInput struct {
	Name                 *string   `json:"name,omitempty"`
	IAMRole              *string   `json:"iamRole,omitempty"`
	ExternalID           *string   `json:"externalId,omitempty"`
	OwnershipTagOverride *bool     `json:"awsTagsOverrideOwnership,omitempty"`
	OwnershipTagKeys     []string  `json:"ownershipTagKeys"`
	RegionOverride       *[]string `json:"regionOverride,omitempty"`
}

func (awsIntegrationInput AWSIntegrationInput) GetGraphQLType() string { return "AwsIntegrationInput" }
func (newRelicIntegrationInput NewRelicIntegrationInput) GetGraphQLType() string {
	return "NewRelicIntegrationInput"
}

func (integrationId *IntegrationId) Alias() string {
	return fmt.Sprintf("%s-%s", slug.Make(integrationId.Type), slug.Make(integrationId.Name))
}

func (client *Client) CreateIntegrationAWS(input AWSIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"awsIntegrationCreate(input: $input)"`
	}
	// This is a default in the UI, so we must maintain it
	if len(input.OwnershipTagKeys) == 0 {
		input.OwnershipTagKeys = append(input.OwnershipTagKeys, "owner")
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationCreate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateIntegrationNewRelic(input NewRelicIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"newRelicIntegrationCreate(input: $input)"`
	}

	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("NewRelicIntegrationCreate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetIntegration(id ID) (*Integration, error) {
	var q struct {
		Account struct {
			Integration Integration `graphql:"integration(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("IntegrationGet"))
	if q.Account.Integration.Id == "" {
		err = fmt.Errorf("integration with ID '%s' not found", id)
	}
	return &q.Account.Integration, HandleErrors(err, nil)
}

func (client *Client) ListIntegrations(variables *PayloadVariables) (*IntegrationConnection, error) {
	var q struct {
		Account struct {
			Integrations IntegrationConnection `graphql:"integrations(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("IntegrationList")); err != nil {
		return nil, err
	}
	for q.Account.Integrations.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Integrations.PageInfo.End
		resp, err := client.ListIntegrations(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Integrations.Nodes = append(q.Account.Integrations.Nodes, resp.Nodes...)
		q.Account.Integrations.PageInfo = resp.PageInfo
		q.Account.Integrations.TotalCount += resp.TotalCount
	}
	return &q.Account.Integrations, nil
}

func (client *Client) UpdateIntegrationAWS(identifier string, input AWSIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"awsIntegrationUpdate(integration: $integration input: $input)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
		"input":       input,
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationUpdate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateIntegrationNewRelic(identifier string, input NewRelicIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"newRelicIntegrationUpdate(input: $input resource: $resource)"`
	}
	v := PayloadVariables{
		"resource": *NewIdentifier(identifier),
		"input":    input,
	}
	err := client.Mutate(&m, v, WithName("NewRelicIntegrationUpdate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteIntegration(identifier string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"integrationDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("IntegrationDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateIntegrationAzureResources(input AzureResourcesIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"azureResourcesIntegrationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AzureResourcesIntegrationCreate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateIntegrationAzureResources(identifier string, input AzureResourcesIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"azureResourcesIntegrationUpdate(integration: $integration input: $input)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
		"input":       input,
	}
	err := client.Mutate(&m, v, WithName("AzureResourcesIntegrationUpdate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateIntegrationGCP(input GoogleCloudIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"googleCloudIntegrationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("GoogleCloudIntegrationCreate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateIntegrationGCP(identifier string, input GoogleCloudIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"googleCloudIntegrationUpdate(integration: $integration input: $input)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
		"input":       input,
	}
	err := client.Mutate(&m, v, WithName("GoogleCloudIntegrationUpdate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) IntegrationReactivate(identifier string) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"integrationReactivate(integration: $integration)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("IntegrationReactivate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}
