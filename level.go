package opslevel

import (
	"github.com/shurcooL/graphql"
)

type Level struct {
	Alias       string
	Description string     `json:"description,omitempty"`
	Id          graphql.ID `json:"id"`
	Index       int
	Name        string
}

type LevelConnection struct {
	Nodes      []Level
	PageInfo   PageInfo
	TotalCount graphql.Int
}

func (conn *LevelConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Rubric struct {
				Levels LevelConnection `graphql:"levels(after: $after, first: $first)"`
			}
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Rubric.Levels.PageInfo = conn.PageInfo
	for q.Account.Rubric.Levels.PageInfo.HasNextPage {
		v["after"] = q.Account.Rubric.Levels.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Rubric.Levels.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

//#region Retrieve

func (client *Client) ListLevels() ([]Level, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Levels LevelConnection
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Rubric.Levels.Nodes, err
	}
	if err := q.Account.Rubric.Levels.Hydrate(client); err != nil {
		return q.Account.Rubric.Levels.Nodes, err
	}
	return q.Account.Rubric.Levels.Nodes, nil
}

//#endregion
