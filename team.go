package opslevel

import (
	"fmt"
	"html"
)

type Contact struct {
	Address     string
	DisplayName string
	Id          ID
	Type        ContactType
}

type ContactInput struct {
	Type        ContactType `json:"type"`
	DisplayName *string     `json:"displayName,omitempty"`
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
	Members          *UserConnection
	Name             string
	Responsibilities string
}

// Had to create this to prevent circular references on User because Team has UserConnection
type TeamIdConnection struct {
	Nodes      []TeamId
	PageInfo   PageInfo
	TotalCount int
}

type TeamConnection struct {
	Nodes      []Team
	PageInfo   PageInfo
	TotalCount int
}

type TeamCreateInput struct {
	Name             string           `json:"name"`
	ManagerEmail     string           `json:"managerEmail,omitempty"`
	Responsibilities string           `json:"responsibilities,omitempty"`
	Group            *IdentifierInput `json:"group"`
	Contacts         *[]ContactInput  `json:"contacts,omitempty"`
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

func (t *Team) ResourceId() ID {
	return t.Id
}

func (t *Team) ResourceType() TaggableResource {
	return TaggableResourceTeam
}

//#region Helpers

func (t *Team) Hydrate(client *Client) error {
	if t == nil || t.Id == "" {
		return nil
	}

	t.Responsibilities = html.UnescapeString(t.Responsibilities)

	if t.Members == nil {
		t.Members = &UserConnection{}
	}
	if t.Members.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = t.Members.PageInfo.End
		_, err := t.GetMembers(client, variables)
		if err != nil {
			return err
		}
	}

	if _, err := t.Tags(client, nil); err != nil {
		return err
	}
	return nil
}

func (t *Team) GetMembers(client *Client, variables *PayloadVariables) (*UserConnection, error) {
	var q struct {
		Account struct {
			Team struct {
				Members UserConnection `graphql:"members(after: $after, first: $first)"`
			} `graphql:"team(id: $team)"`
		}
	}
	if t.Id == "" {
		return nil, fmt.Errorf("Unable to get Members, invalid team id: '%s'", t.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["team"] = t.Id
	if err := client.Query(&q, *variables, WithName("TeamMembersList")); err != nil {
		return nil, err
	}
	if t.Members == nil {
		members := UserConnection{}
		t.Members = &members
	}
	t.Members.Nodes = append(t.Members.Nodes, q.Account.Team.Members.Nodes...)
	t.Members.PageInfo = q.Account.Team.Members.PageInfo
	t.Members.TotalCount += q.Account.Team.Members.TotalCount
	for t.Members.PageInfo.HasNextPage {
		(*variables)["after"] = t.Members.PageInfo.End
		_, err := t.GetMembers(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return t.Members, nil
}

func (t *TeamId) Tags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	if t.Id == "" {
		return nil, fmt.Errorf("Unable to get Tags, invalid team id: '%s'", t.Id)
	}

	var q struct {
		Account struct {
			Team struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"team(id: $team)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["team"] = t.Id
	if err := client.Query(&q, *variables, WithName("TeamTagsList")); err != nil {
		return nil, err
	}

	for q.Account.Team.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Team.Tags.PageInfo.End
		resp, err := t.Tags(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Team.Tags.Nodes = append(q.Account.Team.Tags.Nodes, resp.Nodes...)
		q.Account.Team.Tags.PageInfo = resp.PageInfo
		q.Account.Team.Tags.TotalCount += resp.TotalCount
	}
	return &q.Account.Team.Tags, nil
}

func BuildMembershipInput(members []string) (output []TeamMembershipUserInput) {
	for _, email := range members {
		output = append(output, TeamMembershipUserInput{Email: email})
	}
	return
}

func CreateContactSlack(channel string, name *string) ContactInput {
	return ContactInput{
		Type:        ContactTypeSlack,
		DisplayName: name,
		Address:     channel,
	}
}

func CreateContactSlackHandle(channel string, name *string) ContactInput {
	return ContactInput{
		Type:        ContactTypeSlackHandle,
		DisplayName: name,
		Address:     channel,
	}
}

func CreateContactEmail(email string, name *string) ContactInput {
	return ContactInput{
		Type:        ContactTypeEmail,
		DisplayName: name,
		Address:     email,
	}
}

func CreateContactWeb(address string, name *string) ContactInput {
	return ContactInput{
		Type:        ContactTypeWeb,
		DisplayName: name,
		Address:     address,
	}
}

func (t *Team) HasTag(key string, value string) bool {
	tags, err := t.Tags(NewGQLClient(), nil)
	if err != nil {
		return false
	}
	for _, tag := range tags.Nodes {
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
	if err := client.Mutate(&m, v, WithName("TeamCreate")); err != nil {
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
	err := client.Mutate(&m, v, WithName("TeamMembershipCreate"))
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
		DisplayName: *contact.DisplayName,
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
	err := client.Mutate(&m, v, WithName("ContactCreate"))
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
		"alias": alias,
	}
	if err := client.Query(&q, v, WithName("TeamGet")); err != nil {
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
	if err := client.Query(&q, v, WithName("TeamGet")); err != nil {
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
	err := client.Query(&q, nil, WithName("TeamCount"))
	return int(q.Account.Teams.TotalCount), HandleErrors(err, nil)
}

func (client *Client) ListTeams(variables *PayloadVariables) (*TeamConnection, error) {
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}

	if err := client.Query(&q, *variables, WithName("TeamList")); err != nil {
		return &TeamConnection{}, err
	}

	for q.Account.Teams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Teams.PageInfo.End
		resp, err := client.ListTeams(variables)
		if err != nil {
			return &TeamConnection{}, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return &TeamConnection{}, err
			}
			q.Account.Teams.Nodes = append(q.Account.Teams.Nodes, node)
		}
		q.Account.Teams.PageInfo = resp.PageInfo
		q.Account.Teams.TotalCount += resp.TotalCount
	}
	return &q.Account.Teams, nil
}

func (client *Client) ListTeamsWithManager(email string, variables *PayloadVariables) (*TeamConnection, error) {
	var q struct {
		Account struct {
			Teams TeamConnection `graphql:"teams(managerEmail: $email, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["email"] = email

	if err := client.Query(&q, *variables, WithName("TeamList")); err != nil {
		return &TeamConnection{}, err
	}

	for q.Account.Teams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Teams.PageInfo.End
		resp, err := client.ListTeamsWithManager(email, variables)
		if err != nil {
			return &TeamConnection{}, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return &TeamConnection{}, err
			}
			q.Account.Teams.Nodes = append(q.Account.Teams.Nodes, node)
		}
		q.Account.Teams.PageInfo = resp.PageInfo
		q.Account.Teams.TotalCount += resp.TotalCount
	}
	return &q.Account.Teams, nil
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
	if err := client.Mutate(&m, v, WithName("TeamUpdate")); err != nil {
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
		DisplayName: *contact.DisplayName,
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
	err := client.Mutate(&m, v, WithName("ContactUpdate"))
	return &m.Payload.Contact, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteTeamWithAlias(alias string) error {
	var m struct {
		Payload struct {
			Id     ID               `graphql:"deletedTeamId"`
			Alias  string           `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamDeleteInput{
			Alias: alias,
		},
	}
	err := client.Mutate(&m, v, WithName("TeamDelete"))
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
			Alias  string           `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamDeleteInput{
			Id: id,
		},
	}
	err := client.Mutate(&m, v, WithName("TeamDelete"))
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
	err := client.Mutate(&m, v, WithName("TeamMembershipDelete"))
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
	err := client.Mutate(&m, v, WithName("ContactDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
