package opslevel

import (
	"fmt"

	"github.com/hasura/go-graphql-client"
)

type Level struct {
	Alias       string
	Description string `json:"description,omitempty"`
	Id          ID     `json:"id"`
	Index       int
	Name        string
}

type LevelConnection struct {
	Nodes      []Level
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type LevelCreateInput struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Index       *int   `json:"index,omitempty"`
}

type LevelUpdateInput struct {
	Id          ID              `json:"id"`
	Name        graphql.String  `json:"name,omitempty"`
	Description *graphql.String `json:"description,omitempty"`
}

type LevelDeleteInput struct {
	Id ID `json:"id"`
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

//#region Create

func (client *Client) CreateLevel(input LevelCreateInput) (*Level, error) {
	var m struct {
		Payload struct {
			Level  Level
			Errors []OpsLevelErrors
		} `graphql:"levelCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Level, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetLevel(id ID) (*Level, error) {
	var q struct {
		Account struct {
			Level Level `graphql:"level(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if q.Account.Level.Id == "" {
		return nil, fmt.Errorf("Level with ID '%s' not found!", id)
	}
	return &q.Account.Level, nil
}

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

//#region Update

func (client *Client) UpdateLevel(input LevelUpdateInput) (*Level, error) {
	var m struct {
		Payload struct {
			Level  Level
			Errors []OpsLevelErrors
		} `graphql:"levelUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Level, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteLevel(id ID) error {
	var m struct {
		Payload struct {
			Id     graphql.ID `graphql:"deletedLevelId"`
			Errors []OpsLevelErrors
		} `graphql:"levelDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": LevelDeleteInput{Id: id},
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
