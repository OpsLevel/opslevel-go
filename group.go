package opslevel

import (
	"fmt"
)

type GroupId struct {
	Alias string `json:"alias,omitempty"`
	Id    ID     `json:"id"`
}

type Group struct {
	GroupId
	Description string  `json:"description,omitempty"`
	HtmlURL     string  `json:"htmlUrl,omitempty"`
	Name        string  `json:"name,omitempty"`
	Parent      GroupId `json:"parent,omitempty"`
}

// type SubgroupConnection struct {
// 	nodes      []GroupId
// 	PageInfo   PageInfo
// 	TotalCount graphql.Int
// }

type GroupConnection struct {
	Nodes      []Group
	PageInfo   PageInfo
	TotalCount int
}

type GroupInput struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Parent      *IdentifierInput   `json:"parent"`
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
	err := client.Mutate(&m, v, WithName("GroupCreate"))
	return &m.Payload.Group, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetGroup(id ID) (*Group, error) {
	var q struct {
		Account struct {
			Group Group `graphql:"group(id: $group)"`
		}
	}
	v := PayloadVariables{
		"group": id,
	}
	err := client.Query(&q, v, WithName("GroupGet"))
	return &q.Account.Group, HandleErrors(err, nil)
}

func (client *Client) GetGroupWithAlias(alias string) (*Group, error) {
	var q struct {
		Account struct {
			Group Group `graphql:"group(alias: $group)"`
		}
	}
	v := PayloadVariables{
		"group": alias,
	}
	err := client.Query(&q, v, WithName("GroupGet"))
	return &q.Account.Group, HandleErrors(err, nil)
}

func (client *Client) ListGroups(variables *PayloadVariables) (GroupConnection, error) {
	var q struct {
		Account struct {
			Groups GroupConnection `graphql:"groups(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables); err != nil {
		return GroupConnection{}, err
	}
	for q.Account.Groups.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Groups.PageInfo.End
		resp, err := client.ListGroups(variables)
		if err != nil {
			return GroupConnection{}, err
		}
		q.Account.Groups.Nodes = append(q.Account.Groups.Nodes, resp.Nodes...)
		q.Account.Groups.PageInfo = resp.PageInfo
		q.Account.Groups.TotalCount += resp.TotalCount
	}
	return q.Account.Groups, nil
}

func (g *Group) ChildTeams(client *Client, variables *PayloadVariables) (*TeamConnection, error) {
	var q struct {
		Account struct {
			Group struct {
				ChildTeams TeamConnection `graphql:"childTeams(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Teams, invalid group id: '%s'", g.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["group"] = g.Id

	if err := client.Query(&q, *variables, WithName("GroupChildTeamsList")); err != nil {
		return nil, err
	}
	for q.Account.Group.ChildTeams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Group.ChildTeams.PageInfo.End
		resp, err := g.ChildTeams(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Group.ChildTeams.Nodes = append(q.Account.Group.ChildTeams.Nodes, resp.Nodes...)
		q.Account.Group.ChildTeams.PageInfo = resp.PageInfo
		q.Account.Group.ChildTeams.TotalCount += resp.TotalCount
	}
	return &q.Account.Group.ChildTeams, nil
}

func (g *Group) DescendantTeams(client *Client, variables *PayloadVariables) (*TeamConnection, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantTeams TeamConnection `graphql:"descendantTeams(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Teams, invalid group id: '%s'", g.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["group"] = g.Id
	if err := client.Query(&q, *variables, WithName("GroupDescendantTeamsList")); err != nil {
		return nil, err
	}
	for q.Account.Group.DescendantTeams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Group.DescendantTeams.PageInfo.End
		resp, err := g.DescendantTeams(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Group.DescendantTeams.Nodes = append(q.Account.Group.DescendantTeams.Nodes, resp.Nodes...)
		q.Account.Group.DescendantTeams.PageInfo = resp.PageInfo
		q.Account.Group.DescendantTeams.TotalCount += resp.TotalCount
	}
	return &q.Account.Group.DescendantTeams, nil
}

func (g *Group) DescendantRepositories(client *Client) ([]RepositoryId, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantRepositories struct {
					Nodes    []RepositoryId
					PageInfo PageInfo
				} `graphql:"descendantRepositories(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Repositories, invalid group id: '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": "",
	}
	output := []RepositoryId{}
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

func (g *Group) DescendantServices(client *Client) ([]ServiceId, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantServices struct {
					Nodes    []ServiceId
					PageInfo PageInfo
				} `graphql:"descendantServices(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Services, invalid group id: '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": "",
	}
	output := []ServiceId{}
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

func (g *Group) DescendantSubgroups(client *Client) ([]GroupId, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantSubgroups struct {
					Nodes    []GroupId
					PageInfo PageInfo
				} `graphql:"descendantSubgroups(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Subgroups, invalid group id: '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": "",
	}
	output := []GroupId{}
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

func (g *Group) Members(client *Client) ([]UserId, error) {
	var q struct {
		Account struct {
			Group struct {
				Members struct {
					Nodes    []UserId
					PageInfo PageInfo
				} `graphql:"members(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Members, invalid group id: '%s'", g.Id)
	}
	v := PayloadVariables{
		"group": g.Id,
		"first": client.pageSize,
		"after": "",
	}
	output := []UserId{}
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
	err := client.Mutate(&m, v, WithName("GroupUpdate"))
	return &m.Payload.Group, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

// Deprecated: Please use DeleteGroup instead
func (client *Client) DeleteGroupWithAlias(alias string) error {
	return client.DeleteGroup(alias)
}

func (client *Client) DeleteGroup(identifier string) error {
	var m struct {
		Payload ResourceDeletePayload `graphql:"groupDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("GroupDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
