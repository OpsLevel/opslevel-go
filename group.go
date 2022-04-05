package opslevel

import (
	"github.com/shurcooL/graphql"
)

type GroupCreateInput struct {
	Description string             `json:"description,omitempty"`
	Members     []GroupMemberInput `json:"members,omitempty"`
	Name        string             `json:"name,omitempty"`
	Parent      graphql.ID         `json:"parent,omitempty"`
	Teams       graphql.ID         `json:"teams,omitempty"`
}

type GroupId struct {
	Alias string     `json:"alias,omitempty"`
	Id    graphql.ID `json:"id"`
}

type GroupMemberInput struct {
	Email string `json:"email"`
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
	Description string         `json:"description,omitempty"`
	HtmlURL     string         `json:"htmlUrl,omitempty"`
	Members     UserConnection `json:"members,omitempty"`
	Name        string         `json:"name,omitempty"`
	Parent      GroupParent    `json:"parent,omitempty"`
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

func (client *Client) CreateGroup(input GroupCreateInput) (*Group, error) {
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
	if err := m.Payload.Group.Hydrate(client); err != nil {
		return &m.Payload.Group, err
	}
	return &m.Payload.Group, FormatErrors(m.Payload.Errors)
}

func (s *Group) Hydrate(client *Client) error {
	// TODO: Will need to hydrate descendants and members here
	return nil
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
	if err := q.Account.Group.Hydrate(client); err != nil {
		return &q.Account.Group, err
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
	hydratedGroups := GroupConnection{}
	for _, item := range conn.Nodes {
		if err := (&item).Hydrate(client); err != nil {
			return err
		}
		hydratedGroups.Nodes = append(hydratedGroups.Nodes, item)
	}
	q.Account.Groups.PageInfo = conn.PageInfo
	for q.Account.Groups.PageInfo.HasNextPage {
		v["after"] = q.Account.Groups.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Groups.Nodes {
			if err := (&item).Hydrate(client); err != nil {
				return err
			}
			hydratedGroups.Nodes = append(hydratedGroups.Nodes, item)
		}
	}
	conn.Nodes = hydratedGroups.Nodes
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
