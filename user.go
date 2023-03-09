package opslevel

import (
	"errors"
	"fmt"
)

type MemberInput struct {
	Email string `json:"email"`
}

type UserId struct {
	Id    ID
	Email string
}

type User struct {
	UserId
	HTMLUrl string
	Name    string
	Role    UserRole
}

type UserConnection struct {
	Nodes      []User
	PageInfo   PageInfo
	TotalCount int
}

type UserIdentifierInput struct {
	Id    ID             `graphql:"id" json:"id,omitempty"`
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
			Id: ID(value),
		}
	}
	return UserIdentifierInput{
		Email: graphql.String(value),
	}
}

func (u *User) Teams(client *Client, variables *PayloadVariables) ([]Team, error) { // not sure what to refactor []Team to?  a pointer to something new?
	var q struct {
		Account struct {
			User struct {
				// can I collapse this to something flat like Members in Edgar's example?  where is that created/derived from?
				Teams struct {
					Nodes    []Team
					PageInfo PageInfo
				} `graphql:"teams(after: $after, first: $first)"`
			} `graphql:"user(id: $user)"`
		}
	}
	if u.Id == "" {
		return nil, fmt.Errorf("unable to get teams, nil user id")  // do I need to have a u.ID here for the error to return?
	}

	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)[""] = u.Id // what goes in the "" here?
	if err := client.Query(&q, *variables, WithName("")); err != nil { // what goes in "" here and how is it derived?
		return nil, err
	}
	if t.ID == nil {
		team :=  // whatever we turn the []Team into from line 53 goes here
		t.ID = &team
	}
	// pausing here to hand-off to Kyle
	//v := PayloadVariables{
	//	"user":  u.Id,
	//	"first": client.pageSize,
	//	"after": "",
	//}

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
	err := client.Mutate(&m, v, WithName("UserInvite"))
	return &m.Payload.User, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetUser(value string) (*User, error) {
	var q struct {
		Account struct {
			User User `graphql:"user(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": NewUserIdentifier(value),
	}
	err := client.Query(&q, v, WithName("UserGet"))
	return &q.Account.User, HandleErrors(err, nil)
}

func (client *Client) ListUsers(variables *PayloadVariables) (UserConnection, error) {
	var q struct {
		Account struct {
			Users UserConnection `graphql:"users(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}

	if err := client.Query(&q, *variables, WithName("UserList")); err != nil {
		return UserConnection{}, err
	}
	//output = append(output, q.Account.Users.Nodes...)
	for q.Account.Users.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Users.PageInfo.End
		resp, err := client.ListUsers(variables)
		if err != nil {
			return UserConnection{}, err
		}
		q.Account.Users.Nodes = append(q.Account.Users.Nodes, resp.Nodes...)
		q.Account.Users.PageInfo = resp.PageInfo
		q.Account.Users.TotalCount += resp.TotalCount
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
	err := client.Mutate(&m, v, WithName("UserUpdate"))
	return &m.Payload.User, HandleErrors(err, m.Payload.Errors)
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
	err := client.Mutate(&m, v, WithName("UserDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
