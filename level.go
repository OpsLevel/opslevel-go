package opslevel

import (
	"fmt"

	"github.com/hasura/go-graphql-client"
)

type LevelConnection struct {
	Nodes      []Level
	PageInfo   PageInfo
	TotalCount int
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
		conn.Nodes = append(conn.Nodes, q.Account.Rubric.Levels.Nodes...)
	}
	return nil
}

func (client *Client) CreateLevel(input LevelCreateInput) (*Level, error) {
	var m struct {
		Payload LevelCreatePayload `graphql:"levelCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("LevelCreate"))
	return &m.Payload.Level, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetLevel(id ID) (*Level, error) {
	var q struct {
		Account struct {
			Level Level `graphql:"level(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("LevelGet"))
	if q.Account.Level.Id == "" {
		err = graphql.Errors{graphql.Error{
			Message: fmt.Sprintf("level with ID '%s' not found", id),
			Path:    []any{"account", "level"},
		}}
	}
	return &q.Account.Level, HandleErrors(err, nil)
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

func (client *Client) UpdateLevel(input LevelUpdateInput) (*Level, error) {
	var m struct {
		Payload LevelUpdatePayload `graphql:"levelUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("LevelUpdate"))
	return &m.Payload.Level, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteLevel(id ID) error {
	var m struct {
		Payload struct { // TODO: fix this
			Id     ID `graphql:"deletedLevelId"`
			Errors []Error
		} `graphql:"levelDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": LevelDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("LevelDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
