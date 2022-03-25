package opslevel

import (
	"github.com/shurcooL/graphql"
)

type GroupId struct {
	Id graphql.ID `json:"id"`
}

type GroupParent struct {
	GroupId
	Alias string
}

type Group struct {
	GroupId
	Alias       string      `json:"alias,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Parent      GroupParent `json:"parent,omitempty"`
}

func (client *Client) GetGroup(id graphql.ID) (*Group, error) {
	var q struct {
		Account struct {
			Group Group `graphql:"group(id: $group)"`
		}
	}
	v := PayloadVariables{
		"group": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.Group, nil
}
