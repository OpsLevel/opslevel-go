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

type GroupConnection struct {
	Nodes      []Group
	PageInfo   PageInfo
	TotalCount int
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

func (g *Group) DescendantRepositories(client *Client, variables *PayloadVariables) (*RepositoryConnection, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantRepositories RepositoryConnection `graphql:"descendantRepositories(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Repositories, invalid group id: '%s'", g.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["group"] = g.Id
	if err := client.Query(&q, *variables, WithName("GroupDescendantRepositoriesList")); err != nil {
		return nil, err
	}
	for q.Account.Group.DescendantRepositories.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Group.DescendantRepositories.PageInfo.End
		resp, err := g.DescendantRepositories(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Group.DescendantRepositories.Nodes = append(q.Account.Group.DescendantRepositories.Nodes, resp.Nodes...)
		q.Account.Group.DescendantRepositories.PageInfo = resp.PageInfo
		q.Account.Group.DescendantRepositories.TotalCount += resp.TotalCount
	}
	return &q.Account.Group.DescendantRepositories, nil
}

func (g *Group) DescendantServices(client *Client, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantServices ServiceConnection `graphql:"descendantServices(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Services, invalid group id: '%s'", g.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["group"] = g.Id
	if err := client.Query(&q, *variables, WithName("GroupDescendantServicesList")); err != nil {
		return nil, err
	}
	for q.Account.Group.DescendantServices.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Group.DescendantServices.PageInfo.End
		resp, err := g.DescendantServices(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Group.DescendantServices.Nodes = append(q.Account.Group.DescendantServices.Nodes, resp.Nodes...)
		q.Account.Group.DescendantServices.PageInfo = resp.PageInfo
		q.Account.Group.DescendantServices.TotalCount += resp.TotalCount
	}
	return &q.Account.Group.DescendantServices, nil
}

func (g *Group) DescendantSubgroups(client *Client, variables *PayloadVariables) (*GroupConnection, error) {
	var q struct {
		Account struct {
			Group struct {
				DescendantSubgroups GroupConnection `graphql:"descendantSubgroups(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Subgroups, invalid group id: '%s'", g.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["group"] = g.Id
	if err := client.Query(&q, *variables, WithName("GroupDescendantSubgroupsList")); err != nil {
		return nil, err
	}
	for q.Account.Group.DescendantSubgroups.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Group.DescendantSubgroups.PageInfo.End
		resp, err := g.DescendantSubgroups(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Group.DescendantSubgroups.Nodes = append(q.Account.Group.DescendantSubgroups.Nodes, resp.Nodes...)
		q.Account.Group.DescendantSubgroups.PageInfo = resp.PageInfo
		q.Account.Group.DescendantSubgroups.TotalCount += resp.TotalCount
	}
	return &q.Account.Group.DescendantSubgroups, nil
}

func (g *Group) Members(client *Client, variables *PayloadVariables) (*UserConnection, error) {
	var q struct {
		Account struct {
			Group struct {
				Members UserConnection `graphql:"members(after: $after, first: $first)"`
			} `graphql:"group(id: $group)"`
		}
	}
	if g.Id == "" {
		return nil, fmt.Errorf("Unable to get Members, invalid group id: '%s'", g.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["group"] = g.Id
	if err := client.Query(&q, *variables, WithName("GroupMembersList")); err != nil {
		return nil, err
	}
	for q.Account.Group.Members.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Group.Members.PageInfo.End
		resp, err := g.Members(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Group.Members.Nodes = append(q.Account.Group.Members.Nodes, resp.Nodes...)
		q.Account.Group.Members.PageInfo = resp.PageInfo
		q.Account.Group.Members.TotalCount += resp.TotalCount
	}
	return &q.Account.Group.Members, nil
}

//#endregion

//#region Delete

func (client *Client) DeleteGroup(identifier string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"groupDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("GroupDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
