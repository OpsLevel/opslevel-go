package opslevel

import (
	"fmt"
	"github.com/relvacode/iso8601"

	"github.com/gosimple/slug"
)

type Integration struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	CreatedAt   iso8601.Time `graphql:"createdAt"`
	InstalledAt iso8601.Time `graphql:"installedAt"`

	AWSIntegrationFragment `graphql:"... on AwsIntegration"`
}

type AWSIntegrationFragment struct {
	IAMRole              string   `graphql:"iamRole"`
	ExternalID           string   `graphql:"externalId"`
	OwnershipTagOverride bool     `graphql:"awsTagsOverrideOwnership"`
	OwnershipTagKeys     []string `graphql:"ownershipTagKeys"`
}

type IntegrationConnection struct {
	Nodes      []Integration
	PageInfo   PageInfo
	TotalCount int
}

type AWSIntegrationInput struct {
	IAMRole              string   `json:"iamRole"`
	ExternalID           string   `json:"externalId"`
	OwnershipTagOverride bool     `json:"awsTagsOverrideOwnership"`
	OwnershipTagKeys     []string `json:"ownershipTagKeys"`
}

func (self *Integration) Alias() string {
	return fmt.Sprintf("%s-%s", slug.Make(self.Type), slug.Make(self.Name))
}

//#region Create

func (client *Client) CreateAWSIntegration(input AWSIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"awsIntegrationCreate(identifier: $identifier input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationCreate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

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
		err = fmt.Errorf("Integration with ID '%s' not found!", id)
	}
	return &q.Account.Integration, HandleErrors(err, nil)
}

func (client *Client) ListIntegrations(variables *PayloadVariables) (IntegrationConnection, error) {
	var q struct {
		Account struct {
			Integrations IntegrationConnection `graphql:"integrations(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("IntegrationList")); err != nil {
		return IntegrationConnection{}, err
	}
	for q.Account.Integrations.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Integrations.PageInfo.End
		resp, err := client.ListIntegrations(variables)
		if err != nil {
			return IntegrationConnection{}, err
		}
		q.Account.Integrations.Nodes = append(q.Account.Integrations.Nodes, resp.Nodes...)
		q.Account.Integrations.PageInfo = resp.PageInfo
		q.Account.Integrations.TotalCount += resp.TotalCount
	}
	return q.Account.Integrations, nil
}

//#endregion

//#region Update

func (client *Client) UpdateAWSIntegration(identifier string, input AWSIntegrationInput) (*Integration, error) {
	var m struct {
		Payload struct {
			Integration *Integration
			Errors      []OpsLevelErrors
		} `graphql:"awsIntegrationUpdate(identifier: $identifier input: $input)"`
	}
	v := PayloadVariables{
		"identifier": *NewIdentifier(identifier),
		"input":      input,
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationUpdate"))
	return m.Payload.Integration, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteAWSIntegration(identifier string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"awsIntegrationDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("AWSIntegrationDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
