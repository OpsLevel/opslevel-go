package opslevel

import (
	"errors"
	"fmt"
	"html"
	"slices"
)

func (systemId *SystemId) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			System struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"system(input: $system)"`
		}
	}
	if systemId.Id == "" {
		return nil, fmt.Errorf("unable to get Tags, invalid system id: '%s'", systemId.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["system"] = *NewIdentifier(string(systemId.Id))

	if err := client.Query(&q, *variables, WithName("SystemTagsList")); err != nil {
		return nil, err
	}
	if q.Account.System.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.System.Tags.PageInfo.End
		resp, err := systemId.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
		// Add unique tags only
		for _, resp := range resp.Nodes {
			if !slices.Contains[[]Tag, Tag](q.Account.System.Tags.Nodes, resp) {
				q.Account.System.Tags.Nodes = append(q.Account.System.Tags.Nodes, resp)
			}
		}
		q.Account.System.Tags.PageInfo = resp.PageInfo
	}
	q.Account.System.Tags.TotalCount = len(q.Account.System.Tags.Nodes)
	return &q.Account.System.Tags, nil
}

func (systemId *SystemId) GetAliases() []string {
	return systemId.Aliases
}

func (systemId *SystemId) ResourceId() ID {
	return systemId.Id
}

func (systemId *SystemId) ResourceType() TaggableResource {
	return TaggableResourceSystem
}

func (systemId *SystemId) AliasableType() AliasOwnerTypeEnum {
	return AliasOwnerTypeEnumSystem
}

// Returns unique identifiers created by OpsLevel, values in Aliases but not ManagedAliases
func (system *System) UniqueIdentifiers() []string {
	uniqueIdentifiers := []string{}
	for _, alias := range system.Aliases {
		if !slices.Contains(system.ManagedAliases, alias) {
			uniqueIdentifiers = append(uniqueIdentifiers, alias)
		}
	}

	return uniqueIdentifiers
}

func (system *System) Hydrate(client *Client) error {
	system.Description = html.UnescapeString(system.Description)
	system.Note = html.UnescapeString(system.Note)
	return nil
}

func (system *SystemId) ReconcileAliases(client *Client, aliasesWanted []string) error {
	aliasesToCreate, aliasesToDelete := extractAliases(system.Aliases, aliasesWanted)

	// reconcile wanted aliases with actual aliases
	deleteErr := client.DeleteAliases(AliasOwnerTypeEnumSystem, aliasesToDelete)
	_, createErr := client.CreateAliases(system.Id, aliasesToCreate)

	// update system to reflect API updates
	updatedSystem, getErr := client.GetSystem(string(system.Id))
	if updatedSystem != nil {
		system.Aliases = updatedSystem.Aliases
	}

	return errors.Join(deleteErr, createErr, getErr)
}

func (systemId *SystemId) ChildServices(client *Client, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			System struct {
				ChildServices ServiceConnection `graphql:"childServices(after: $after, first: $first)"`
			} `graphql:"system(input: $system)"`
		}
	}
	if systemId.Id == "" {
		return nil, fmt.Errorf("unable to get Services, invalid system id: '%s'", systemId.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}

	(*variables)["system"] = *NewIdentifier(string(systemId.Id))

	if err := client.Query(&q, *variables, WithName("SystemChildServicesList")); err != nil {
		return nil, err
	}
	if q.Account.System.ChildServices.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.System.ChildServices.PageInfo.End
		resp, err := systemId.ChildServices(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.System.ChildServices.Nodes = append(q.Account.System.ChildServices.Nodes, resp.Nodes...)
		q.Account.System.ChildServices.PageInfo = resp.PageInfo
	}
	q.Account.System.ChildServices.TotalCount = len(q.Account.System.ChildServices.Nodes)
	return &q.Account.System.ChildServices, nil
}

func (systemId *SystemId) AssignService(client *Client, services ...string) error {
	var m struct {
		Payload SystemPayload `graphql:"systemChildAssign(system:$system, childServices:$childServices)"`
	}
	v := PayloadVariables{
		"system":        *NewIdentifier(string(systemId.Id)),
		"childServices": NewIdentifierArray(services),
	}
	err := client.Mutate(&m, v, WithName("SystemAssignService"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateSystem(input SystemInput) (*System, error) {
	var m struct {
		Payload SystemPayload `graphql:"systemCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v, WithName("SystemCreate")); err != nil {
		return nil, err
	}
	if err := m.Payload.System.Hydrate(client); err != nil {
		return &m.Payload.System, err
	}
	return &m.Payload.System, HandleErrors(nil, m.Payload.Errors)
}

func (client *Client) GetSystem(identifier string) (*System, error) {
	var q struct {
		Account struct {
			System System `graphql:"system(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	if err := client.Query(&q, v, WithName("SystemGet")); err != nil {
		return nil, err
	}
	if err := q.Account.System.Hydrate(client); err != nil {
		return &q.Account.System, err
	}
	return &q.Account.System, nil
}

func (client *Client) ListSystems(variables *PayloadVariables) (*SystemConnection, error) {
	var q struct {
		Account struct {
			Systems SystemConnection `graphql:"systems(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("SystemsList")); err != nil {
		return nil, err
	}
	if q.Account.Systems.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Systems.PageInfo.End
		resp, err := client.ListSystems(variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Systems.Nodes = append(q.Account.Systems.Nodes, node)
		}
		q.Account.Systems.PageInfo = resp.PageInfo
	}
	q.Account.Systems.TotalCount = len(q.Account.Systems.Nodes)
	return &q.Account.Systems, nil
}

func (client *Client) UpdateSystem(identifier string, input SystemInput) (*System, error) {
	var s struct {
		Payload SystemPayload `graphql:"systemUpdate(system:$system,input:$input)"`
	}
	v := PayloadVariables{
		"system": *NewIdentifier(identifier),
		"input":  input,
	}
	if err := client.Mutate(&s, v, WithName("SystemUpdate")); err != nil {
		return nil, err
	}
	if err := s.Payload.System.Hydrate(client); err != nil {
		return &s.Payload.System, err
	}
	return &s.Payload.System, HandleErrors(nil, s.Payload.Errors)
}

func (client *Client) DeleteSystem(identifier string) error {
	var s struct {
		Payload BasePayload `graphql:"systemDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&s, v, WithName("SystemDelete"))
	return HandleErrors(err, s.Payload.Errors)
}
