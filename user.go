package opslevel

import (
	"fmt"
	"slices"
)

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
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
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

	}
	q.Account.User.Tags.TotalCount = len(q.Account.User.Tags.Nodes)
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
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["user"] = user.Id
	if err := client.Query(&q, *variables, WithName("UserTeamsList")); err != nil {
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
	}
	q.Account.User.Teams.TotalCount = len(q.Account.User.Teams.Nodes)
	return &q.Account.User.Teams, nil
}

func (user *User) Hydrate(client *Client) error {
	if user.Tags == nil {
		user.Tags = &TagConnection{}
	}
	if user.Tags.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = user.Tags.PageInfo.End
		resp, err := user.GetTags(client, variables)
		if err != nil {
			return err
		}
		user.Tags.Nodes = append(user.Tags.Nodes, resp.Nodes...)
		user.Tags.PageInfo = resp.PageInfo
	}
	user.Tags.TotalCount = len(user.Tags.Nodes)

	if user.TeamsConnection == nil {
		user.TeamsConnection = &TeamIdConnection{}
	}
	if user.TeamsConnection.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = user.TeamsConnection.PageInfo.End
		resp, err := user.Teams(client, variables)
		if err != nil {
			return err
		}
		user.TeamsConnection.Nodes = append(user.TeamsConnection.Nodes, resp.Nodes...)
		user.TeamsConnection.PageInfo = resp.PageInfo
	}
	user.TeamsConnection.TotalCount = len(user.TeamsConnection.Nodes)

	return nil
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

	// Hydrate inner tags/teams connections for every user on this page.
	// Without this, users on the first outer page would never paginate their
	// nested tag/team connections beyond the server's default page size.
	for i := range q.Account.Users.Nodes {
		if err := q.Account.Users.Nodes[i].Hydrate(client); err != nil {
			return nil, err
		}
	}

	if q.Account.Users.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Users.PageInfo.End
		resp, err := client.ListUsers(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Users.Nodes = append(q.Account.Users.Nodes, resp.Nodes...)
		q.Account.Users.PageInfo = resp.PageInfo
	}
	q.Account.Users.TotalCount = len(q.Account.Users.Nodes)
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
