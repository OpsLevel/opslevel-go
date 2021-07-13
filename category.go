package opslevel

import (
	"github.com/shurcooL/graphql"
)

type Category struct {
	Alias       string
	Description string     `json:"description,omitempty"`
	Id          graphql.ID `json:"id"`
	Name        string
}

type CategoryConnection struct {
	Nodes      []Category
	PageInfo   PageInfo
	TotalCount graphql.Int
}

func (conn *CategoryConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Rubric struct {
				Categories CategoryConnection `graphql:"categories(after: $after, first: $first)"`
			}
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Rubric.Categories.PageInfo = conn.PageInfo
	for q.Account.Rubric.Categories.PageInfo.HasNextPage {
		v["after"] = q.Account.Rubric.Categories.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Rubric.Categories.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

//#region Retrieve

func (client *Client) ListCategories() ([]Category, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Categories CategoryConnection
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Rubric.Categories.Nodes, err
	}
	if err := q.Account.Rubric.Categories.Hydrate(client); err != nil {
		return q.Account.Rubric.Categories.Nodes, err
	}
	return q.Account.Rubric.Categories.Nodes, nil
}

//#endregion
