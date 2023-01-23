package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/hasura/go-graphql-client"
)

type Integration struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type IntegrationConnection struct {
	Nodes      []Integration
	PageInfo   PageInfo
	TotalCount graphql.Int
}

func (self *Integration) Alias() string {
	return fmt.Sprintf("%s-%s", slug.Make(self.Type), slug.Make(self.Name))
}

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
	err := client.Query(&q, v)
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
	if err := client.Query(&q, *variables); err != nil {
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
	}
	return q.Account.Integrations, nil
}

//#endregion
