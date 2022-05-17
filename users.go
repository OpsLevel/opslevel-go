package opslevel

import "github.com/shurcooL/graphql"

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
	Name string
	Role UserRole
}

//#region Helpers

func NewUserIdentifier(value string) *UserIdentifierInput {
	if IsID(value) {
		return &UserIdentifierInput{
			Id: graphql.ID(value),
		}
	}
	return &UserIdentifierInput{
		Email: graphql.String(value),
	}
}

// TODO: func (u *User) Teams(client *Client) ([]Team, error)

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
		"email": email,
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.User, FormatErrors(m.Payload.Errors)
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
