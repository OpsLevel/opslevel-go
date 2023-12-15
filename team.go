package opslevel

import (
	"fmt"
	"html"
	"slices"
)

type Contact struct {
	Address     string
	DisplayName string
	Id          ID
	Type        ContactType
}

type ContactInput struct {
	Type        ContactType `json:"type" yaml:"type" default:"email"`
	DisplayName *string     `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Address     string      `json:"address" yaml:"address" default:"team@example.com"`
}

type ContactCreateInput struct {
	Type        ContactType `json:"type" yaml:"type" default:"email"`
	DisplayName string      `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Address     string      `json:"address" yaml:"address" default:"team@example.com"`
	TeamId      *ID         `json:"teamId,omitempty" yaml:"teamId,omitempty"`
	TeamAlias   string      `json:"teamAlias,omitempty" yaml:"teamAlias,omitempty"`
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

// Has no json struct tags as this is nested in returned data structs
type TeamId struct {
	Alias string
	Id    ID
}

type Team struct {
	TeamId

	Aliases  []string
	Contacts []Contact

	Group            GroupId // Deprecated: Group field will be removed in a future release
	HTMLUrl          string
	Manager          User
	Members          *UserConnection // Deprecated: Members field will be removed in a future release
	Memberships      *TeamMembershipConnection
	Name             string
	ParentTeam       TeamId
	Responsibilities string
	Tags             *TagConnection
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
	Name             string           `json:"name" default:"Platform"`
	ManagerEmail     string           `json:"managerEmail,omitempty" yaml:"managerEmail,omitempty" default:"manager@example.com"`
	Responsibilities string           `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty" default:"Makes Tools"`
	Contacts         *[]ContactInput  `json:"contacts,omitempty" yaml:"contacts,omitempty" default:"[{\"Type\":\"slack\",\"DisplayName\":\"Team Eng\",\"Address\":\"#team-engineering\"}]"`
	ParentTeam       *IdentifierInput `json:"parentTeam" yaml:"parentTeam" default:"{\"alias\":\"Engineering\"}"`
}

type TeamUpdateInput struct {
	Id               ID               `json:"id"`
	Alias            string           `json:"alias,omitempty"`
	Name             string           `json:"name,omitempty"`
	ManagerEmail     string           `json:"managerEmail,omitempty"`
	Responsibilities string           `json:"responsibilities,omitempty"`
	ParentTeam       *IdentifierInput `json:"parentTeam"`
}

type TeamDeleteInput struct {
	Id    ID     `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type TeamMembership struct {
	Team TeamId `graphql:"team"`
	Role string `graphql:"role"`
	User UserId `graphql:"user"`
}

type TeamMembershipConnection struct {
	Nodes      []TeamMembership
	PageInfo   PageInfo
	TotalCount int
}

type TeamMembershipUserInput struct {
	User UserIdentifierInput `json:"user"`
	Role string              `json:"role" default:"admin"`
}

type TeamMembershipCreateInput struct {
	TeamId  ID                        `json:"teamId" yaml:"teamId" default:"XXX_team_id_XXX"`
	Members []TeamMembershipUserInput `json:"members"`
}

type TeamMembershipDeleteInput struct {
	TeamId  ID                        `json:"teamId" yaml:"teamId" default:"XXX_team_id_XXX"`
	Members []TeamMembershipUserInput `json:"members"`
}

func (t *Team) ResourceId() ID {
	return t.Id
}

func (t *Team) ResourceType() TaggableResource {
	return TaggableResourceTeam
}

//#region Helpers

func (self *Team) Hydrate(client *Client) error {
	self.Responsibilities = html.UnescapeString(self.Responsibilities)

	if self.Memberships == nil {
		self.Memberships = &TeamMembershipConnection{}
	}
	if self.Memberships.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = self.Memberships.PageInfo.End
		_, err := self.GetMemberships(client, variables)
		if err != nil {
			return err
		}
	}

	if self.Tags == nil {
		self.Tags = &TagConnection{}
	}
	if self.Tags.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = self.Tags.PageInfo.End
		_, err := self.GetTags(client, variables)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Team) GetMemberships(client *Client, variables *PayloadVariables) (*TeamMembershipConnection, error) {
	if t.Id == "" {
		return nil, fmt.Errorf("Unable to get Memberships, invalid team id: '%s'", t.Id)
	}
	var q struct {
		Account struct {
			Team struct {
				Memberships TeamMembershipConnection `graphql:"memberships(after: $after, first: $first)"`
			} `graphql:"team(id: $team)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["team"] = t.Id
	if err := client.Query(&q, *variables, WithName("TeamMembersList")); err != nil {
		return nil, err
	}
	if t.Memberships == nil {
		memberships := TeamMembershipConnection{}
		t.Memberships = &memberships
	}
	t.Memberships.Nodes = append(t.Memberships.Nodes, q.Account.Team.Memberships.Nodes...)
	t.Memberships.PageInfo = q.Account.Team.Memberships.PageInfo
	t.Memberships.TotalCount += q.Account.Team.Memberships.TotalCount
	for t.Memberships.PageInfo.HasNextPage {
		(*variables)["after"] = t.Memberships.PageInfo.End
		_, err := t.GetMemberships(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return t.Memberships, nil
}

// Deprecated: use GetMemberships instead
func (t *Team) GetMembers(client *Client, variables *PayloadVariables) (*UserConnection, error) {
	return nil, fmt.Errorf("Deprecated, use GetMemberships instead")
}

func (t *Team) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			Team struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"team(id: $team)"`
		}
	}
	if t.Id == "" {
		return nil, fmt.Errorf("Unable to get Tags, invalid team id: '%s'", t.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["team"] = t.Id
	if err := client.Query(&q, *variables, WithName("TeamTagsList")); err != nil {
		return nil, err
	}
	if t.Tags == nil {
		t.Tags = &TagConnection{}
	}
	// Add unique tags only
	for _, tagNode := range q.Account.Team.Tags.Nodes {
		if !slices.Contains[[]Tag, Tag](t.Tags.Nodes, tagNode) {
			t.Tags.Nodes = append(t.Tags.Nodes, tagNode)
		}
	}
	t.Tags.PageInfo = q.Account.Team.Tags.PageInfo
	t.Tags.TotalCount += q.Account.Team.Tags.TotalCount
	for t.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = t.Tags.PageInfo.End
		_, err := t.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return t.Tags, nil
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
	if err := client.Mutate(&m, v, WithName("TeamCreate")); err != nil {
		return nil, err
	}
	if err := m.Payload.Team.Hydrate(client); err != nil {
		return &m.Payload.Team, err
	}
	return &m.Payload.Team, FormatErrors(m.Payload.Errors)
}

func (client *Client) AddMemberships(team *TeamId, memberships ...TeamMembershipUserInput) ([]TeamMembership, error) {
	var m struct {
		Payload struct {
			Memberships []TeamMembership `graphql:"memberships"`
			Errors      []OpsLevelErrors
		} `graphql:"teamMembershipCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamMembershipCreateInput{
			TeamId:  team.Id,
			Members: memberships,
		},
	}
	err := client.Mutate(&m, v, WithName("TeamMembershipCreate"))
	return m.Payload.Memberships, HandleErrors(err, m.Payload.Errors)
}

// Deprecated: use AddMemberships instead
func (client *Client) AddMembers(team *TeamId, emails []string) ([]TeamMembership, error) {
	return nil, fmt.Errorf("Deprecated, use AddMemberships instead")
}

// Deprecated: use AddMemberships instead
func (client *Client) AddMember(team *TeamId, email string) ([]TeamMembership, error) {
	return nil, fmt.Errorf("Deprecated, use AddMemberships instead")
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

func (client *Client) RemoveMemberships(team *TeamId, memberships ...TeamMembershipUserInput) ([]User, error) {
	var m struct {
		Payload struct {
			Members []User `graphql:"deletedMembers"`
			Errors  []OpsLevelErrors
		} `graphql:"teamMembershipDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TeamMembershipDeleteInput{
			TeamId:  team.Id,
			Members: memberships,
		},
	}
	err := client.Mutate(&m, v, WithName("TeamMembershipDelete"))
	return m.Payload.Members, HandleErrors(err, m.Payload.Errors)
}

// Deprecated: use RemoveMembers instead
func (client *Client) RemoveMembers(team *TeamId, emails []string) ([]User, error) {
	return nil, fmt.Errorf("Deprecated, use RemoveMemberships instead")
}

// Deprecated: use RemoveMembers instead
func (client *Client) RemoveMember(team *TeamId, membership TeamMembershipUserInput) ([]User, error) {
	return nil, fmt.Errorf("Deprecated, use RemoveMemberships instead")
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
