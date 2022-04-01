package opslevel

import (
	"github.com/shurcooL/graphql"
)

type GroupId struct {
	Alias string     `json:"alias,omitempty"`
	Id    graphql.ID `json:"id"`
}

type GroupParent struct {
	GroupId
	Alias string
}

type Group struct {
	// DescendantRepositories RepositoryConnection `json:"descendantRepositories,omitempty"`
	// DescendantServices ServiceConnection `json:"descendantServices,omitempty"`
	// DescendantSubgroups    SubgroupConnection   `json:"descendantSubgroups,omitempty"`
	// DescendantTeams        TeamConnection       `json:"descendantTeams,omitempty"`
	GroupId
	Description string `json:"description,omitempty"`
	HtmlURL     string `json:"htmlUrl,omitempty"`
	// Members     UserConnection `json:"members,omitempty"`
	Name   string      `json:"name,omitempty"`
	Parent GroupParent `json:"parent,omitempty"`
}

// type SubgroupConnection struct {
// 	nodes      []GroupId
// 	PageInfo   PageInfo
// 	TotalCount graphql.Int
// }

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

func (client *Client) GetGroupWithAlias(alias string) (*Group, error) {
	var q struct {
		Account struct {
			Group Group `graphql:"group(alias: $group)"`
		}
	}
	v := PayloadVariables{
		"group": graphql.String(alias),
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.Group, nil
}
