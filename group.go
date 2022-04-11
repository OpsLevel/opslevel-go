package opslevel

import (
	"fmt"

	"github.com/shurcooL/graphql"
)

type GroupId struct {
	Alias string     `json:"alias,omitempty"`
	Id    graphql.ID `json:"id"`
}

type Group struct {
	// DescendantRepositories RepositoryConnection `json:"descendantRepositories,omitempty"`
	// DescendantServices ServiceConnection `json:"descendantServices,omitempty"`
	// DescendantSubgroups    SubgroupConnection   `json:"descendantSubgroups,omitempty"`
	DescendantTeamsConnection Connection `json:"descendantTeams,omitempty"`
	GroupId
	Description string `json:"description,omitempty"`
	HtmlURL     string `json:"htmlUrl,omitempty"`
	// Members     UserConnection `json:"members,omitempty"`
	Name   string  `json:"name,omitempty"`
	Parent GroupId `json:"parent,omitempty"`
}

// type SubgroupConnection struct {
// 	nodes      []GroupId
// 	PageInfo   PageInfo
// 	TotalCount graphql.Int
// }

type GroupConnection struct {
	Nodes      []Group
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type GroupInput struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Parent      *IdentifierInput   `json:"parent,omitempty"`
	Members     *[]MemberInput     `json:"members,omitempty"`
	Teams       *[]IdentifierInput `json:"teams,omitempty"`
}

//#region Create

func (client *Client) CreateGroup(input GroupInput) (*Group, error) {
	var m struct {
		Payload struct {
			Group  Group
			Errors []OpsLevelErrors
		} `graphql:"groupCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Group, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

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

func (conn *GroupConnection) Query(client *Client, q interface{}, v PayloadVariables) ([]Group, error) {
	if err := client.Query(q, v); err != nil {
		return conn.Nodes, err
	}
	if err := conn.Hydrate(client); err != nil {
		return conn.Nodes, err
	}
	return conn.Nodes, nil
}

func (conn *GroupConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Groups GroupConnection `graphql:"groups(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Groups.PageInfo = conn.PageInfo
	for q.Account.Groups.PageInfo.HasNextPage {
		v["after"] = q.Account.Groups.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		conn.Nodes = append(conn.Nodes, q.Account.Groups.Nodes...)
	}
	return nil
}

func (client *Client) ListGroups() ([]Group, error) {
	var q struct {
		Account struct {
			Groups GroupConnection `graphql:"groups(after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	return q.Account.Groups.Query(client, &q, v)
}

func (g *Group) DescendantTeams(client *Client) ([]Team, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantTeams struct {
					Nodes    []Team
					PageInfo PageInfo
				} `graphql:"descendantTeams(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == nil {
		return nil, fmt.Errorf("failed to request 'DescendantTeams' for group because there is not a valid id '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
	}
	output := []Team{}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	for q.Account.Group.DescendantTeams.PageInfo.HasNextPage {
		output = append(output, q.Account.Group.DescendantTeams.Nodes...)
		v["after"] = q.Account.Group.DescendantTeams.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return nil, err
		}
	}
	return output, nil
}

//#endregion

//#region Update

func (client *Client) UpdateGroup(identifier string, input GroupInput) (*Group, error) {
	var m struct {
		Payload struct {
			Group  Group
			Errors []OpsLevelErrors
		} `graphql:"groupUpdate(group: $group, input: $input)"`
	}
	v := PayloadVariables{
		"group": *NewIdentifier(identifier),
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Group, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteGroupWithAlias(alias string) error {
	var m struct {
		Payload ResourceDeletePayload `graphql:"groupDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(alias),
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

func (client *Client) DeleteGroup(id graphql.ID) error {
	var m struct {
		Payload ResourceDeletePayload `graphql:"groupDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(id.(string)),
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
