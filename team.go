package opslevel

import (
	"github.com/shurcooL/graphql"
)

type User struct {
	Name  string
	Email string
}

type Contact struct {
	DisplayName string
	Address     string
}

type ContactInput struct {
	Type        string `json:"type,omitEmpty"`
	DisplayName string `json:"displayName,omitEmpty"`
	Address     string `json:"address,omitEmpty"`
}

type Team struct {
	Alias            string
	Contacts         []Contact
	Id               graphql.ID
	Manager          User
	Name             string
	Responsibilities string
}

type TeamConnection struct {
	Nodes    []Team
	PageInfo PageInfo
}

type TeamCreateInput struct {
	Name             string         `json:"name"`
	ManagerEmail     string         `json:"managerEmail,omitempty"`
	Responsibilities string         `json:"responsibilities,omitempty"`
	Contacts         []ContactInput `json:"contacts,omitempty"`
}

type TeamUpdateInput struct {
	Id               graphql.ID `json:"id,omitempty"`
	Alias            string     `json:"alias,omitempty"`
	Name             string     `json:"name,omitempty"`
	ManagerEmail     string     `json:"managerEmail,omitempty"`
	Responsibilities string     `json:"responsibilities,omitempty"`
}

type TeamDeleteInput struct {
	Id    graphql.ID `json:"id,omitempty"`
	Alias string     `json:"alias,omitempty"`
}

//#region Create

func (client *Client) CreateTeam(input TeamCreateInput) (*Team, error) {
	var m struct {
		Payload struct {
			Team   Team
			Errors []OpsLevelErrors
		} `graphql:"teamCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Team, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetTeamWithAlias(alias string) (*Team, error) {
	var q struct {
		Account struct {
			Team Team `graphql:"team(alias: $alias)"`
		}
	}
	v := PayloadVariables{
		"alias": graphql.String(alias),
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.Team, nil
}

// Deprecated: use GetTeam instead
func (client *Client) GetTeamWithId(id graphql.ID) (*Team, error) {
	return client.GetTeam(id)
}

func (client *Client) GetTeam(id graphql.ID) (*Team, error) {
	var q struct {
		Account struct {
			Team Team `graphql:"team(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.Team, nil
}

func (client *Client) GetTeamCount() (int, error) {
	var q struct {
		Account struct {
			Teams struct {
				TotalCount int
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return 0, err
	}
	return int(q.Account.Teams.TotalCount), nil
}

func (client *Client) ListTeams() ([]Team, error) {
	var output []Team
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
	}
	if err := client.Query(&q, v); err != nil {
		return output, err
	}
	for _, item := range q.Account.Teams.Nodes {
		output = append(output, item)
	}
	for q.Account.Teams.PageInfo.HasNextPage {
		v["after"] = q.Account.Teams.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return output, err
		}
		for _, item := range q.Account.Teams.Nodes {
			output = append(output, item)
		}
	}
	return output, nil
}

func (client *Client) ListTeamsWithManager(email string) ([]Team, error) {
	var output []Team
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(managerEmail: $email, after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
		"email": graphql.String(email),
	}
	if err := client.Query(&q, v); err != nil {
		return output, err
	}
	for _, item := range q.Account.Teams.Nodes {
		output = append(output, item)
	}
	for q.Account.Teams.PageInfo.HasNextPage {
		v["after"] = q.Account.Teams.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return output, err
		}
		for _, item := range q.Account.Teams.Nodes {
			output = append(output, item)
		}
	}
	return output, nil
}

//#endregion

//#region Update

func (client *Client) UpdateTeam(input TeamUpdateInput) (*Team, error) {
	var m struct {
		Payload struct {
			Team   Team
			Errors []OpsLevelErrors
		} `graphql:"teamUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Team, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteTeamWithAlias(alias string) error {
	var m struct {
		Payload struct {
			Id     graphql.ID       `graphql:"deletedTeamId"`
			Alias  graphql.String   `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamDeleteInput{
			Alias: alias,
		},
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

// Deprecated: use DeleteTeam instead
func (client *Client) DeleteTeamWithId(id graphql.ID) error {
	return client.DeleteTeam(id)
}

func (client *Client) DeleteTeam(id graphql.ID) error {
	var m struct {
		Payload struct {
			Id     graphql.ID       `graphql:"deletedTeamId"`
			Alias  graphql.String   `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamDeleteInput{
			Id: id,
		},
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
