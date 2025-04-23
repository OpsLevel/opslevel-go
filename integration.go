package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/hasura/go-graphql-client"
)

type IntegrationId struct {
	Id   ID     `json:"id"`   // The unique identifier of the integration.
	Name string `json:"name"` // The name of the integration.
	Type string `json:"type"` // The type of the integration.
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
	ExternalID           *Nullable[string] `json:"externalId,omitempty"`
	IAMRole              *Nullable[string] `json:"iamRole,omitempty"`
	Name                 *Nullable[string] `json:"name,omitempty"`
	OwnershipTagKeys     []string          `json:"ownershipTagKeys"`
	OwnershipTagOverride *Nullable[bool]   `json:"awsTagsOverrideOwnership,omitempty"`
	RegionOverride       *[]string         `json:"regionOverride,omitempty"`
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
		Payload IntegrationCreatePayload `graphql:"awsIntegrationCreate(input: $input)"`
	}
	// This is a default in the UI, so we must maintain it
	if len(input.OwnershipTagKeys) == 0 {
		input.OwnershipTagKeys = append(input.OwnershipTagKeys, "owner")
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationCreate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateEventIntegration(input EventIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationCreatePayload `graphql:"eventIntegrationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("EventIntegrationCreate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateIntegrationNewRelic(input NewRelicIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationCreatePayload `graphql:"newRelicIntegrationCreate(input: $input)"`
	}

	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("NewRelicIntegrationCreate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
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
		err = graphql.Errors{graphql.Error{
			Message: fmt.Sprintf("integration with ID '%s' not found", id),
			Path:    []any{"account", "integration"},
		}}
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
	if q.Account.Integrations.PageInfo.HasNextPage {
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
		Payload IntegrationUpdatePayload `graphql:"awsIntegrationUpdate(integration: $integration input: $input)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
		"input":       input,
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationUpdate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateEventIntegration(input EventIntegrationUpdateInput) (*Integration, error) {
	var m struct {
		Payload IntegrationUpdatePayload `graphql:"eventIntegrationUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("EventIntegrationUpdate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateIntegrationNewRelic(identifier string, input NewRelicIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationUpdatePayload `graphql:"newRelicIntegrationUpdate(input: $input resource: $resource)"`
	}
	v := PayloadVariables{
		"resource": *NewIdentifier(identifier),
		"input":    input,
	}
	err := client.Mutate(&m, v, WithName("NewRelicIntegrationUpdate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteIntegration(identifier string) error {
	var m struct {
		Payload BasePayload `graphql:"integrationDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("IntegrationDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateIntegrationAzureResources(input AzureResourcesIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationCreatePayload `graphql:"azureResourcesIntegrationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AzureResourcesIntegrationCreate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateIntegrationAzureResources(identifier string, input AzureResourcesIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationUpdatePayload `graphql:"azureResourcesIntegrationUpdate(integration: $integration input: $input)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
		"input":       input,
	}
	err := client.Mutate(&m, v, WithName("AzureResourcesIntegrationUpdate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateIntegrationGCP(input GoogleCloudIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationCreatePayload `graphql:"googleCloudIntegrationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("GoogleCloudIntegrationCreate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateIntegrationGCP(identifier string, input GoogleCloudIntegrationInput) (*Integration, error) {
	var m struct {
		Payload IntegrationUpdatePayload `graphql:"googleCloudIntegrationUpdate(integration: $integration input: $input)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
		"input":       input,
	}
	err := client.Mutate(&m, v, WithName("GoogleCloudIntegrationUpdate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) IntegrationReactivate(identifier string) (*Integration, error) {
	var m struct {
		Payload IntegrationReactivatePayload `graphql:"integrationReactivate(integration: $integration)"`
	}
	v := PayloadVariables{
		"integration": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("IntegrationReactivate"))
	return &m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}
