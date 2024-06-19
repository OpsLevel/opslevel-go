package opslevel

import (
	"errors"
	"fmt"
	"html"
	"slices"
)

type Contact struct {
	Address     string
	DisplayName string
	DisplayType string
	ExternalId  string
	Id          ID
	IsDefault   bool
	Type        ContactType
}

type TeamId struct {
	Alias string
	Id    ID
}

type Team struct {
	TeamId

	Aliases        []string `graphql:"aliases" json:"aliases" yaml:"aliases"`
	ManagedAliases []string `graphql:"managedAliases" json:"managedAliases" yaml:"managedAliases"`
	Contacts       []Contact

	HTMLUrl          string
	Manager          User
	Memberships      *TeamMembershipConnection
	Name             string
	ParentTeam       TeamId
	Responsibilities string
	Tags             *TagConnection
}

// TeamIdConnection exists to prevent circular references on User because Team has a UserConnection
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

type TeamMembership struct {
	Role string `graphql:"role"`
	Team TeamId `graphql:"team"`
	User UserId `graphql:"user"`
}

type TeamMembershipConnection struct {
	Nodes      []TeamMembership
	PageInfo   PageInfo
	TotalCount int
}

func (team *Team) ReconcileAliases(client *Client, aliasesWanted []string) error {
	var allErrors, err error

	aliasesToCreate, aliasesToDelete := ExtractAliases(team.ManagedAliases, aliasesWanted)
	for _, alias := range aliasesToDelete {
		err := client.DeleteAlias(AliasDeleteInput{
			Alias:     alias,
			OwnerType: AliasOwnerTypeEnumTeam,
		})
		allErrors = errors.Join(allErrors, err)
	}

	if len(aliasesToCreate) > 0 {
		// CreateAliases returns current list of aliases from owned by Team
		team.ManagedAliases, err = client.CreateAliases(team.Id, aliasesToCreate)
		allErrors = errors.Join(allErrors, err)
	} else {
		team.ManagedAliases = slices.DeleteFunc(team.ManagedAliases, func(alias string) bool {
			return slices.Contains(aliasesToDelete, alias)
		})
	}

	return allErrors
}

func (team *Team) ResourceId() ID {
	return team.Id
}

func (team *Team) ResourceType() TaggableResource {
	return TaggableResourceTeam
}

func (team *Team) Hydrate(client *Client) error {
	team.Responsibilities = html.UnescapeString(team.Responsibilities)

	if team.Memberships == nil {
		team.Memberships = &TeamMembershipConnection{}
	}
	if team.Memberships.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = team.Memberships.PageInfo.End
		_, err := team.GetMemberships(client, variables)
		if err != nil {
			return err
		}
	}

	if team.Tags == nil {
		team.Tags = &TagConnection{}
	}
	if team.Tags.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = team.Tags.PageInfo.End
		_, err := team.GetTags(client, variables)
		if err != nil {
			return err
		}
	}
	return nil
}

func (team *Team) GetMemberships(client *Client, variables *PayloadVariables) (*TeamMembershipConnection, error) {
	if team.Id == "" {
		return nil, fmt.Errorf("unable to get Memberships, invalid team id: '%s'", team.Id)
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
	(*variables)["team"] = team.Id
	if err := client.Query(&q, *variables, WithName("TeamMembersList")); err != nil {
		return nil, err
	}
	if team.Memberships == nil {
		memberships := TeamMembershipConnection{}
		team.Memberships = &memberships
	}
	team.Memberships.Nodes = append(team.Memberships.Nodes, q.Account.Team.Memberships.Nodes...)
	team.Memberships.PageInfo = q.Account.Team.Memberships.PageInfo
	team.Memberships.TotalCount += q.Account.Team.Memberships.TotalCount
	for team.Memberships.PageInfo.HasNextPage {
		(*variables)["after"] = team.Memberships.PageInfo.End
		_, err := team.GetMemberships(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return team.Memberships, nil
}

func (team *Team) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			Team struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"team(id: $team)"`
		}
	}
	if team.Id == "" {
		return nil, fmt.Errorf("unable to get Tags, invalid team id: '%s'", team.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["team"] = team.Id
	if err := client.Query(&q, *variables, WithName("TeamTagsList")); err != nil {
		return nil, err
	}
	if team.Tags == nil {
		team.Tags = &TagConnection{}
	}
	// Add unique tags only
	for _, tagNode := range q.Account.Team.Tags.Nodes {
		if !slices.Contains[[]Tag, Tag](team.Tags.Nodes, tagNode) {
			team.Tags.Nodes = append(team.Tags.Nodes, tagNode)
		}
	}
	team.Tags.PageInfo = q.Account.Team.Tags.PageInfo
	team.Tags.TotalCount += q.Account.Team.Tags.TotalCount
	for team.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = team.Tags.PageInfo.End
		_, err := team.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return team.Tags, nil
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

func (team *Team) HasTag(key string, value string) bool {
	for _, tag := range team.Tags.Nodes {
		if tag.Key == key && tag.Value == value {
			return true
		}
	}
	return false
}

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
		contactInput.OwnerId = NewID(team)
	} else {
		contactInput.TeamAlias = &team
	}

	v := PayloadVariables{
		"input": contactInput,
	}
	err := client.Mutate(&m, v, WithName("ContactCreate"))
	return &m.Payload.Contact, HandleErrors(err, m.Payload.Errors)
}

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
	return q.Account.Teams.TotalCount, HandleErrors(err, nil)
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
		return nil, err
	}

	for q.Account.Teams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Teams.PageInfo.End
		resp, err := client.ListTeams(variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
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
		return nil, err
	}

	for q.Account.Teams.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Teams.PageInfo.End
		resp, err := client.ListTeamsWithManager(email, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Teams.Nodes = append(q.Account.Teams.Nodes, node)
		}
		q.Account.Teams.PageInfo = resp.PageInfo
		q.Account.Teams.TotalCount += resp.TotalCount
	}
	return &q.Account.Teams, nil
}

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
		DisplayName: contact.DisplayName,
		Address:     &contact.Address,
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

func (client *Client) DeleteTeam(identifier string) error {
	input := TeamDeleteInput{}
	if IsID(identifier) {
		input.Id = NewID(identifier)
	} else {
		input.Alias = &identifier
	}

	var m struct {
		Payload struct {
			Id     ID               `graphql:"deletedTeamId"`
			Alias  string           `graphql:"deletedTeamAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"teamDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
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
