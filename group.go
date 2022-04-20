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
	DescendantRepositoriesConnection Connection `json:"-" graphql:"descendantRepositories"`
	DescendantServicesConnection     Connection `json:"-" graphql:"descendantServices"`
	DescendantSubgroupsConnection    Connection `json:"-" graphql:"descendantSubgroups"`
	DescendantTeamsConnection        Connection `json:"-" graphql:"descendantTeams"`
	GroupId
	Description       string     `json:"description,omitempty"`
	HtmlURL           string     `json:"htmlUrl,omitempty"`
	MembersConnection Connection `json:"-" graphql:"members"`
	Name              string     `json:"name,omitempty"`
	Parent            GroupId    `json:"parent,omitempty"`
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
		"after": g.DescendantTeamsConnection.PageInfo.Start,
	}
	output := []Team{}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Group.DescendantTeams.Nodes...)
	for q.Account.Group.DescendantTeams.PageInfo.HasNextPage {
		v["after"] = q.Account.Group.DescendantTeams.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Group.DescendantTeams.Nodes...)
	}
	return output, nil
}

func (g *Group) DescendantRepositories(client *Client) ([]Repository, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantRepositories struct {
					Nodes    []Repository
					PageInfo PageInfo
				} `graphql:"descendantRepositories(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == nil {
		return nil, fmt.Errorf("failed to request 'DescendantRepositories' for group because there is not a valid id '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": g.DescendantRepositoriesConnection.PageInfo.Start,
	}
	output := []Repository{}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Group.DescendantRepositories.Nodes...)
	for q.Account.Group.DescendantRepositories.PageInfo.HasNextPage {
		v["after"] = q.Account.Group.DescendantRepositories.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Group.DescendantRepositories.Nodes...)
	}
	return output, nil
}

func (g *Group) DescendantServices(client *Client) ([]Service, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantServices struct {
					Nodes    []Service
					PageInfo PageInfo
				} `graphql:"descendantServices(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == nil {
		return nil, fmt.Errorf("failed to request 'DescendantServices' for group because there is not a valid id '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": g.DescendantServicesConnection.PageInfo.Start,
	}
	output := []Service{}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Group.DescendantServices.Nodes...)
	for q.Account.Group.DescendantServices.PageInfo.HasNextPage {
		v["after"] = q.Account.Group.DescendantServices.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Group.DescendantServices.Nodes...)
	}
	return output, nil
}

func (g *Group) DescendantSubgroups(client *Client) ([]Group, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantSubgroups struct {
					Nodes    []Group
					PageInfo PageInfo
				} `graphql:"descendantSubgroups(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == nil {
		return nil, fmt.Errorf("failed to request 'DescendantSubgroups' for group because there is not a valid id '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": g.DescendantSubgroupsConnection.PageInfo.Start,
	}
	output := []Group{}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Group.DescendantSubgroups.Nodes...)
	for q.Account.Group.DescendantSubgroups.PageInfo.HasNextPage {
		v["after"] = q.Account.Group.DescendantSubgroups.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Group.DescendantSubgroups.Nodes...)
	}
	return output, nil
}

func (g *Group) Members(client *Client) ([]User, error) {
	var q struct {
		Account struct {
			Group struct {
				Members struct {
					Nodes    []User
					PageInfo PageInfo
				} `graphql:"members(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == nil {
		return nil, fmt.Errorf("failed to request 'Members' for group because there is not a valid id '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": g.MembersConnection.PageInfo.Start,
	}
	output := []User{}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Group.Members.Nodes...)
	for q.Account.Group.Members.PageInfo.HasNextPage {
		v["after"] = q.Account.Group.Members.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Group.Members.Nodes...)
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
