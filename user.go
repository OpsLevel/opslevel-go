package opslevel

import (
	"fmt"
	"slices"
)

type UserConnection struct {
	Nodes      []User
	PageInfo   PageInfo
	TotalCount int
}

func (s *UserConnection) GetNodes() any {
	return s.Nodes
}

func (user *User) ResourceId() ID {
	return user.Id
}

func (user *User) ResourceType() TaggableResource {
	return TaggableResourceUser
}

func NewUserIdentifier(value string) *UserIdentifierInput {
	if IsID(value) {
		return &UserIdentifierInput{
			Id: RefOf(ID(value)),
		}
	}
	return &UserIdentifierInput{
		Email: RefOf(value),
	}
}

func (userId *UserId) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			User struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"user(id: $user)"`
		}
	}
	if userId.Id == "" {
		return nil, fmt.Errorf("unable to get Tags, invalid User id: '%s'", userId.Id)
	}

	variables = client.PopulatePaginationParams(variables)
	(*variables)["user"] = userId.Id
	if err := client.Query(&q, *variables, WithName("UserTagsList")); err != nil {
		return nil, err
	}
	if q.Account.User.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.User.Tags.PageInfo.End
		resp, err := userId.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
		// Add unique tags only
		for _, resp := range resp.Nodes {
			if !slices.Contains[[]Tag, Tag](q.Account.User.Tags.Nodes, resp) {
				q.Account.User.Tags.Nodes = append(q.Account.User.Tags.Nodes, resp)
			}
		}
		q.Account.User.Tags.PageInfo = resp.PageInfo
		q.Account.User.Tags.TotalCount += resp.TotalCount
	}

	return &q.Account.User.Tags, nil
}

func (user *User) Teams(client *Client, variables *PayloadVariables) (*TeamIdConnection, error) {
	var q struct {
		Account struct {
			User struct {
				Teams TeamIdConnection `graphql:"teams(after: $after, first: $first)"`
			} `graphql:"user(id: $user)"`
		}
	}
	if user.Id == "" {
		return nil, fmt.Errorf("unable to get teams, nil user id")
	}

	variables = client.PopulatePaginationParams(variables)
	(*variables)["user"] = user.Id
	if err := client.Query(&q, *variables, WithName("UserTeamsList")); err != nil { // what goes in "" here and how is it derived?
		return nil, err
	}
	if q.Account.User.Teams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.User.Teams.PageInfo.End
		conn, err := user.Teams(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.User.Teams.Nodes = append(q.Account.User.Teams.Nodes, conn.Nodes...)
		q.Account.User.Teams.PageInfo = conn.PageInfo
		q.Account.User.Teams.TotalCount += conn.TotalCount
	}
	return &q.Account.User.Teams, nil
}

func (client *Client) InviteUser(email string, input UserInput, sendInvite bool) (*User, error) {
	var m struct {
		Payload struct { // TODO: need to fix this
			User   User
			Errors []Error
		} `graphql:"userInvite(email: $email input: $input, forceSendInvite: $forceSendInvite)"`
	}
	v := PayloadVariables{
		"email":           email,
		"input":           input,
		"forceSendInvite": &sendInvite,
	}
	err := client.Mutate(&m, v, WithName("UserInvite"))
	return &m.Payload.User, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetUser(value string) (*User, error) {
	var q struct {
		Account struct {
			User User `graphql:"user(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewUserIdentifier(value),
	}
	err := client.Query(&q, v, WithName("UserGet"))
	return &q.Account.User, HandleErrors(err, nil)
}

func (client *Client) ListUsers(variables *PayloadVariables) (*UserConnection, error) {
	var q struct {
		Account struct {
			Users UserConnection `graphql:"users(after: $after, first: $first, filter: $filter)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
		(*variables)["filter"] = new([]UsersFilterInput)
	}

	if err := client.Query(&q, *variables, WithName("UserList")); err != nil {
		return nil, err
	}

	if q.Account.Users.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Users.PageInfo.End
		resp, err := client.ListUsers(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Users.Nodes = append(q.Account.Users.Nodes, resp.Nodes...)
		q.Account.Users.PageInfo = resp.PageInfo
		q.Account.Users.TotalCount += resp.TotalCount
	}
	return &q.Account.Users, nil
}

func (client *Client) UpdateUser(user string, input UserInput) (*User, error) {
	var m struct {
		Payload UserPayload `graphql:"userUpdate(user: $user input: $input)"`
	}
	v := PayloadVariables{
		"user":  *NewUserIdentifier(user),
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("UserUpdate"))
	return &m.Payload.User, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteUser(user string) error {
	var m struct {
		Payload BasePayload `graphql:"userDelete(user: $user)"`
	}
	v := PayloadVariables{
		"user": *NewUserIdentifier(user),
	}
	err := client.Mutate(&m, v, WithName("UserDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
