package opslevel

import (
	"html"

	"github.com/hasura/go-graphql-client"
)

type Contact struct {
	Address     string
	DisplayName string
	Id          ID
	Type        ContactType
}

type ContactInput struct {
	Type        ContactType `json:"type"`
	DisplayName string      `json:"displayName,omitEmpty"`
	Address     string      `json:"address"`
}

type ContactCreateInput struct {
	Type        ContactType `json:"type"`
	DisplayName string      `json:"displayName,omitempty"`
	Address     string      `json:"address"`
	TeamId      *ID         `json:"teamId,omitempty"`
	TeamAlias   string      `json:"teamAlias,omitempty"`
}

type ContactUpdateInput struct {
	Id          ID           `json:"id"`
	Type        *ContactType `json:"type,omitempty"`
	DisplayName string       `json:"displayName,omitempty"`
	Address     string       `json:"address,omitempty"`
}

type ContactDeleteInput struct {
	Id ID `json:"id"`
}

type TeamId struct {
	Alias string
	Id    ID
}

type Team struct {
	TeamId

	Aliases          []string
	Contacts         []Contact
	Group            GroupId
	HTMLUrl          string
	Manager          User
	Members          UserConnection
	Name             string
	Responsibilities string
	Tags             TagConnection
}

type TeamConnection struct {
	Nodes    []Team
	PageInfo PageInfo
}

type TeamCreateInput struct {
	Name             string           `json:"name"`
	ManagerEmail     string           `json:"managerEmail,omitempty"`
	Responsibilities string           `json:"responsibilities,omitempty"`
	Group            *IdentifierInput `json:"group"`
	Contacts         []ContactInput   `json:"contacts,omitempty"`
}

type TeamUpdateInput struct {
	Id               ID               `json:"id"`
	Alias            string           `json:"alias,omitempty"`
	Name             string           `json:"name,omitempty"`
	ManagerEmail     string           `json:"managerEmail,omitempty"`
	Group            *IdentifierInput `json:"group"`
	Responsibilities string           `json:"responsibilities,omitempty"`
}

type TeamDeleteInput struct {
	Id    ID     `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type TeamMembershipUserInput struct {
	Email string `json:"email"`
}

type TeamMembershipCreateInput struct {
	TeamId  ID                        `json:"teamId"`
	Members []TeamMembershipUserInput `json:"members"`
}

type TeamMembershipDeleteInput struct {
	TeamId  ID                        `json:"teamId"`
	Members []TeamMembershipUserInput `json:"members"`
}

//#region Helpers

func (conn *UserConnection) Hydrate(id ID, client *Client) error {
	var q struct {
		Account struct {
			Team struct {
				Members UserConnection `graphql:"members(after: $after, first: $first)"`
			} `graphql:"team(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id":    id,
		"first": client.pageSize,
	}
	q.Account.Team.Members.PageInfo = conn.PageInfo
	for q.Account.Team.Members.PageInfo.HasNextPage {
		v["after"] = q.Account.Team.Members.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Team.Members.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (self *Team) Hydrate(client *Client) error {
	self.Responsibilities = html.UnescapeString(self.Responsibilities)
	if err := self.Members.Hydrate(self.Id, client); err != nil {
		return err
	}
	if err := self.Tags.Hydrate(self.Id, client); err != nil {
		return err
	}
	return nil
}

func (conn *TeamConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	for i, item := range conn.Nodes {
		if err := (&item).Hydrate(client); err != nil {
			return err
		}
		conn.Nodes[i] = item
	}
	q.Account.Teams.PageInfo = conn.PageInfo
	for q.Account.Teams.PageInfo.HasNextPage {
		v["after"] = q.Account.Teams.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Teams.Nodes {
			if err := (&item).Hydrate(client); err != nil {
				return err
			}
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (conn *TeamConnection) Query(client *Client, q interface{}, v PayloadVariables) ([]Team, error) {
	if err := client.Query(q, v); err != nil {
		return conn.Nodes, err
	}
	if err := conn.Hydrate(client); err != nil {
		return conn.Nodes, err
	}
	return conn.Nodes, nil
}

func BuildMembershipInput(members []string) (output []TeamMembershipUserInput) {
	for _, email := range members {
		output = append(output, TeamMembershipUserInput{Email: email})
	}
	return
}

func CreateContactSlack(channel string, name string) ContactInput {
	return ContactInput{
		Type:        ContactTypeSlack,
		DisplayName: name,
		Address:     channel,
	}
}

func CreateContactEmail(email string, name string) ContactInput {
	return ContactInput{
		Type:        ContactTypeEmail,
		DisplayName: name,
		Address:     email,
	}
}

func CreateContactWeb(address string, name string) ContactInput {
	return ContactInput{
		Type:        ContactTypeWeb,
		DisplayName: name,
		Address:     address,
	}
}

func (s *Team) HasTag(key string, value string) bool {
	for _, tag := range s.Tags.Nodes {
		if tag.Key == key && tag.Value == value {
			return true
		}
	}
	return false
}

//#endregion

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
	if err := m.Payload.Team.Hydrate(client); err != nil {
		return &m.Payload.Team, err
	}
	return &m.Payload.Team, FormatErrors(m.Payload.Errors)
}

func (client *Client) AddMembers(team *TeamId, emails []string) ([]User, error) {
	var m struct {
		Payload struct {
			Members []User
			Errors  []OpsLevelErrors
		} `graphql:"teamMembershipCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamMembershipCreateInput{
			TeamId:  team.Id,
			Members: BuildMembershipInput(emails),
		},
	}
	err := client.Mutate(&m, v)
	return m.Payload.Members, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) AddMember(team *TeamId, email string) ([]User, error) {
	emails := []string{email}
	return client.AddMembers(team, emails)
}

func (client *Client) AddContact(team string, contact ContactInput) (*Contact, error) {
	var m struct {
		Payload struct {
			Contact Contact
			Errors  []OpsLevelErrors
		} `graphql:"contactCreate(input: $input)"`
	}
	contactInput := ContactCreateInput{
		Type:        contact.Type,
		DisplayName: contact.DisplayName,
		Address:     contact.Address,
	}
	if IsID(team) {
		contactInput.TeamId = NewID(team)
	} else {
		contactInput.TeamAlias = team
	}
	v := PayloadVariables{
		"input": contactInput,
	}
	err := client.Mutate(&m, v)
	return &m.Payload.Contact, HandleErrors(err, m.Payload.Errors)
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
	if err := q.Account.Team.Hydrate(client); err != nil {
		return &q.Account.Team, err
	}
	return &q.Account.Team, nil
}

// Deprecated: use GetTeam instead
func (client *Client) GetTeamWithId(id ID) (*Team, error) {
	return client.GetTeam(id)
}

func (client *Client) GetTeam(id ID) (*Team, error) {
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
	if err := q.Account.Team.Hydrate(client); err != nil {
		return &q.Account.Team, err
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
	err := client.Query(&q, nil)
	return int(q.Account.Teams.TotalCount), HandleErrors(err, nil)
}

func (client *Client) ListTeams() ([]Team, error) {
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	return q.Account.Teams.Query(client, &q, v)
}

func (client *Client) ListTeamsWithManager(email string) ([]Team, error) {
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(managerEmail: $email, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["email"] = graphql.String(email)
	return q.Account.Teams.Query(client, &q, v)
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
	if err := m.Payload.Team.Hydrate(client); err != nil {
		return &m.Payload.Team, err
	}
	return &m.Payload.Team, FormatErrors(m.Payload.Errors)
}

func (client *Client) UpdateContact(id ID, contact ContactInput) (*Contact, error) {
	var m struct {
		Payload struct {
			Contact Contact
			Errors  []OpsLevelErrors
		} `graphql:"contactUpdate(input: $input)"`
	}
	input := ContactUpdateInput{
		Id:          id,
		DisplayName: contact.DisplayName,
		Address:     contact.Address,
	}
	if contact.Type == "" {
		input.Type = nil
	} else {
		input.Type = &contact.Type
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v)
	return &m.Payload.Contact, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteTeamWithAlias(alias string) error {
	var m struct {
		Payload struct {
			Id     ID               `graphql:"deletedTeamId"`
			Alias  graphql.String   `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamDeleteInput{
			Alias: alias,
		},
	}
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

// Deprecated: use DeleteTeam instead
func (client *Client) DeleteTeamWithId(id ID) error {
	return client.DeleteTeam(id)
}

func (client *Client) DeleteTeam(id ID) error {
	var m struct {
		Payload struct {
			Id     ID               `graphql:"deletedTeamId"`
			Alias  graphql.String   `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamDeleteInput{
			Id: id,
		},
	}
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) RemoveMembers(team *TeamId, emails []string) ([]User, error) {
	var m struct {
		Payload struct {
			Members []User `graphql:"deletedMembers"`
			Errors  []OpsLevelErrors
		} `graphql:"teamMembershipDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamMembershipDeleteInput{
			TeamId:  team.Id,
			Members: BuildMembershipInput(emails),
		},
	}
	err := client.Mutate(&m, v)
	return m.Payload.Members, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) RemoveMember(team *TeamId, email string) ([]User, error) {
	emails := []string{email}
	return client.RemoveMembers(team, emails)
}

func (client *Client) RemoveContact(contact ID) error {
	var m struct {
		Payload struct {
			Contact ID `graphql:"deletedContactId"`
			Errors  []OpsLevelErrors
		} `graphql:"contactDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": ContactDeleteInput{
			Id: contact,
		},
	}
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
