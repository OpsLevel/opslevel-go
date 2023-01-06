package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/hasura/go-graphql-client"
)

type Integration struct {
	Id   graphql.ID `json:"id"`
	Name string     `json:"name"`
	Type string     `json:"type"`
}

type IntegrationConnection struct {
	Nodes      []Integration
	PageInfo   PageInfo
	TotalCount graphql.Int
}

func (self *Integration) Alias() string {
	return fmt.Sprintf("%s-%s", slug.Make(self.Type), slug.Make(self.Name))
}

func (conn *IntegrationConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Integrations IntegrationConnection `graphql:"integrations(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Integrations.PageInfo = conn.PageInfo
	for q.Account.Integrations.PageInfo.HasNextPage {
		v["after"] = q.Account.Integrations.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Integrations.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

//#region Retrieve

func (client *Client) GetIntegration(id graphql.ID) (*Integration, error) {
	var q struct {
		Account struct {
			Integration Integration `graphql:"integration(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if q.Account.Integration.Id == "" {
		return nil, fmt.Errorf("Integration with ID '%s' not found!", id)
	}
	return &q.Account.Integration, nil
}

func (client *Client) ListIntegrations() ([]Integration, error) {
	var q struct {
		Account struct {
			Integrations IntegrationConnection `graphql:"integrations"`
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Integrations.Nodes, err
	}
	if err := q.Account.Integrations.Hydrate(client); err != nil {
		return q.Account.Integrations.Nodes, err
	}
	return q.Account.Integrations.Nodes, nil
}

//#endregion
