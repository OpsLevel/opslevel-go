package opslevel

import (
	"errors"
	"github.com/hasura/go-graphql-client"
)

type MemberInput struct {
	Email string `json:"email"`
}

type User struct {
	Email   string
	HTMLUrl string
	Id      graphql.ID
	Name    string
	Role    UserRole
}

type UserId struct {
	Email string
	Id    graphql.ID
}

type UserConnection struct {
	Nodes    []User
	PageInfo PageInfo
}

type UserIdentifierInput struct {
	Id    graphql.ID     `graphql:"id" json:"id,omitempty"`
	Email graphql.String `graphql:"email" json:"email,omitempty"`
}

type UserInput struct {
	Name string   `json:"name,omitempty"`
	Role UserRole `json:"role,omitempty"`
}

//#region Helpers

func NewUserIdentifier(value string) UserIdentifierInput {
	if IsID(value) {
		return UserIdentifierInput{
			Id: graphql.ID(value),
		}
	}
	return UserIdentifierInput{
		Email: graphql.String(value),
	}
}

func (u *User) Teams(client *Client) ([]Team, error) {
	var q struct {
		Account struct {
			User struct {
				Teams struct {
					Nodes    []Team
					PageInfo PageInfo
				} `graphql:"teams(after: $after, first: $first)"`
			} `graphql:"user(id: $user)"`
		}
	}
	if u.Id == "" {
		return nil, errors.New("unable to get teams, nil user id")
	}
	v := PayloadVariables{
		"user":  u.Id,
		"first": client.pageSize,
		"after": graphql.String(""),
	}
	var output []Team
	if err := client.Query(&q, v); err != nil {
		return output, err
	}
	output = append(output, q.Account.User.Teams.Nodes...)
	for q.Account.User.Teams.PageInfo.HasNextPage {
		v["after"] = q.Account.User.Teams.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return output, err
		}
		output = append(output, q.Account.User.Teams.Nodes...)
	}
	return output, nil
}

//#endregion

//#region Create

func (client *Client) InviteUser(email string, input UserInput) (*User, error) {
	var m struct {
		Payload struct {
			User   User
			Errors []OpsLevelErrors
		} `graphql:"userInvite(email: $email input: $input)"`
	}
	v := PayloadVariables{
		"email": graphql.String(email),
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.User, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetUser(id graphql.ID) (*User, error) {
	var q struct {
		Account struct {
			User User `graphql:"user(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.User, nil
}

func (client *Client) ListUsers(variables *PayloadVariables) (UserConnection, error) {
	var q struct {
		Account struct {
			Users UserConnection `graphql:"users(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = &PayloadVariables{
			"after": "",
			"first": client.pageSize,
		}
	}
	if err := client.Query(&q, *variables); err != nil {
		return UserConnection{}, err
	}
	for q.Account.Users.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Users.PageInfo.End
		resp, err := client.ListUsers(variables)
		if err != nil {
			return UserConnection{}, err
		}
		q.Account.Users.Nodes = append(q.Account.Users.Nodes, resp.Nodes...)
		q.Account.Users.PageInfo = resp.PageInfo
	}
	return q.Account.Users, nil
}

//#endregion

//#region Update

func (client *Client) UpdateUser(user string, input UserInput) (*User, error) {
	var m struct {
		Payload struct {
			User   User
			Errors []OpsLevelErrors
		} `graphql:"userUpdate(user: $user input: $input)"`
	}
	v := PayloadVariables{
		"user":  NewUserIdentifier(user),
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.User, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteUser(user string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"userDelete(user: $user)"`
	}
	v := PayloadVariables{
		"user": NewUserIdentifier(user),
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
