package opslevel

import (
	"fmt"

	"github.com/hasura/go-graphql-client"
)

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

func (client *Client) ListLevels(variables *PayloadVariables) (*LevelConnection, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Levels LevelConnection `graphql:"levels(after: $after, first: $first)"`
			}
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("LevelsList")); err != nil {
		return nil, err
	}
	if q.Account.Rubric.Levels.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Rubric.Levels.PageInfo.End
		resp, err := client.ListLevels(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Rubric.Levels.Nodes = append(q.Account.Rubric.Levels.Nodes, resp.Nodes...)
		q.Account.Rubric.Levels.PageInfo = resp.PageInfo
	}
	q.Account.Rubric.Levels.TotalCount = len(q.Account.Rubric.Levels.Nodes)
	return &q.Account.Rubric.Levels, nil
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
