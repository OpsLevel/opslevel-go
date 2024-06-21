package opslevel

import (
	"fmt"
	"slices"
)

type SystemId Identifier

type System struct {
	SystemId
	Name        string      `graphql:"name"`
	Description string      `graphql:"description"`
	HTMLUrl     string      `graphql:"htmlUrl"`
	Owner       EntityOwner `graphql:"owner"`
	Parent      Domain      `graphql:"parent"`
	Note        string      `graphql:"note"`
}

type SystemConnection struct {
	Nodes      []System `json:"nodes"`
	PageInfo   PageInfo `json:"pageInfo"`
	TotalCount int      `json:"totalCount" graphql:"-"`
}

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
	for q.Account.System.Tags.PageInfo.HasNextPage {
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
		q.Account.System.Tags.TotalCount += resp.TotalCount
	}
	return &q.Account.System.Tags, nil
}

func (systemId *SystemId) ResourceId() ID {
	return systemId.Id
}

func (systemId *SystemId) ResourceType() TaggableResource {
	return TaggableResourceSystem
}

func (system *SystemId) ReconcileAliases(client *Client, aliasesWanted []string) error {
	var err error
	aliasesToCreate := getSliceWithStringsRemoved(aliasesWanted, system.Aliases)
	aliasesToDelete := getSliceWithStringsRemoved(system.Aliases, aliasesWanted)

	if err := client.DeleteAliases(AliasOwnerTypeEnumSystem, aliasesToDelete); err != nil {
		return err
	}

	// Update this system's aliases
	if len(aliasesToCreate) > 0 {
		// aliases retrieved from API
		system.Aliases, err = client.CreateAliases(system.Id, aliasesToCreate)
	} else {
		// aliases updated by hand - no more API calls here
		system.Aliases = getSliceWithStringsRemoved(system.Aliases, aliasesToDelete)
	}

	return err
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
	for q.Account.System.ChildServices.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.System.ChildServices.PageInfo.End
		resp, err := systemId.ChildServices(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.System.ChildServices.Nodes = append(q.Account.System.ChildServices.Nodes, resp.Nodes...)
		q.Account.System.ChildServices.PageInfo = resp.PageInfo
		q.Account.System.ChildServices.TotalCount += resp.TotalCount
	}
	return &q.Account.System.ChildServices, nil
}

func (systemId *SystemId) AssignService(client *Client, services ...string) error {
	var m struct {
		Payload struct {
			System System
			Errors []OpsLevelErrors
		} `graphql:"systemChildAssign(system:$system, childServices:$childServices)"`
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
		Payload struct {
			System System
			Errors []OpsLevelErrors
		} `graphql:"systemCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("SystemCreate"))
	return &m.Payload.System, HandleErrors(err, m.Payload.Errors)
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
	err := client.Query(&q, v, WithName("SystemGet"))
	return &q.Account.System, HandleErrors(err, nil)
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
	for q.Account.Systems.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Systems.PageInfo.End
		resp, err := client.ListSystems(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Systems.Nodes = append(q.Account.Systems.Nodes, resp.Nodes...)
		q.Account.Systems.PageInfo = resp.PageInfo
	}
	q.Account.Systems.TotalCount = len(q.Account.Systems.Nodes)
	return &q.Account.Systems, nil
}

func (client *Client) UpdateSystem(identifier string, input SystemInput) (*System, error) {
	var s struct {
		Payload struct {
			System System
			Errors []OpsLevelErrors
		} `graphql:"systemUpdate(system:$system,input:$input)"`
	}
	v := PayloadVariables{
		"system": *NewIdentifier(identifier),
		"input":  input,
	}
	err := client.Mutate(&s, v, WithName("SystemUpdate"))
	return &s.Payload.System, HandleErrors(err, s.Payload.Errors)
}

func (client *Client) DeleteSystem(identifier string) error {
	var s struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"systemDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&s, v, WithName("SystemDelete"))
	return HandleErrors(err, s.Payload.Errors)
}
