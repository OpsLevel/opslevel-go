package opslevel

import (
	"errors"
	"fmt"
	"slices"
)

type DomainId Identifier

type Domain struct {
	DomainId

	Description    string      `graphql:"description"`
	HTMLUrl        string      `graphql:"htmlUrl"`
	ManagedAliases []string    `graphql:"managedAliases"` // A list of aliases that can be set by users. The unique identifier for the resource is omitted.
	Name           string      `graphql:"name"`
	Note           string      `graphql:"note"`
	Owner          EntityOwner `graphql:"owner"`
}

type DomainConnection struct {
	Nodes      []Domain `json:"nodes"`
	PageInfo   PageInfo `json:"pageInfo"`
	TotalCount int      `json:"totalCount" graphql:"-"`
}

// Returns unique identifiers created by OpsLevel, values in Aliases but not ManagedAliases
func (d *Domain) GetUniqueIdentifiers() []string {
	uniqueIdentifiers := []string{}
	for _, alias := range d.Aliases {
		if !slices.Contains(d.ManagedAliases, alias) {
			uniqueIdentifiers = append(uniqueIdentifiers, alias)
		}
	}

	return uniqueIdentifiers
}

func (d *Domain) ReconcileAliases(client *Client, aliasesWanted []string) error {
	aliasesToCreate, aliasesToDelete := extractAliases(d.Aliases, aliasesWanted)

	// reconcile wanted aliases with actual aliases
	deleteErr := client.DeleteAliases(AliasOwnerTypeEnumDomain, aliasesToDelete)
	_, createErr := client.CreateAliases(d.Id, aliasesToCreate)

	// update domain to reflect API updates
	updatedDomain, getErr := client.GetDomain(string(d.Id))
	if updatedDomain != nil {
		d.Aliases = updatedDomain.Aliases
		d.ManagedAliases = updatedDomain.ManagedAliases
	}

	return errors.Join(deleteErr, createErr, getErr)
}

func (domainId *DomainId) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			Domain struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"domain(input: $domain)"`
		}
	}
	if domainId.Id == "" {
		return nil, fmt.Errorf("unable to get Tags, invalid domain id: '%s'", domainId.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["domain"] = *NewIdentifier(string(domainId.Id))

	if err := client.Query(&q, *variables, WithName("DomainTagsList")); err != nil {
		return nil, err
	}
	for q.Account.Domain.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Domain.Tags.PageInfo.End
		resp, err := domainId.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
		// Add unique tags only
		for _, resp := range resp.Nodes {
			if !slices.Contains[[]Tag, Tag](q.Account.Domain.Tags.Nodes, resp) {
				q.Account.Domain.Tags.Nodes = append(q.Account.Domain.Tags.Nodes, resp)
			}
		}
		q.Account.Domain.Tags.PageInfo = resp.PageInfo
		q.Account.Domain.Tags.TotalCount += resp.TotalCount
	}
	return &q.Account.Domain.Tags, nil
}

func (domainId *DomainId) ResourceId() ID {
	return domainId.Id
}

func (domainId *DomainId) ResourceType() TaggableResource {
	return TaggableResourceDomain
}

func (domainId *DomainId) ChildSystems(client *Client, variables *PayloadVariables) (*SystemConnection, error) {
	var q struct {
		Account struct {
			Domain struct {
				ChildSystems SystemConnection `graphql:"childSystems(after: $after, first: $first)"`
			} `graphql:"domain(input: $domain)"`
		}
	}
	if domainId.Id == "" {
		return nil, fmt.Errorf("unable to get Systems, invalid domain id: '%s'", domainId.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}

	(*variables)["domain"] = *NewIdentifier(string(domainId.Id))

	if err := client.Query(&q, *variables, WithName("DomainChildSystemsList")); err != nil {
		return nil, err
	}
	for q.Account.Domain.ChildSystems.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Domain.ChildSystems.PageInfo.End
		resp, err := domainId.ChildSystems(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Domain.ChildSystems.Nodes = append(q.Account.Domain.ChildSystems.Nodes, resp.Nodes...)
		q.Account.Domain.ChildSystems.PageInfo = resp.PageInfo
	}
	q.Account.Domain.ChildSystems.TotalCount = len(q.Account.Domain.ChildSystems.Nodes)
	return &q.Account.Domain.ChildSystems, nil
}

func (domainId *DomainId) AssignSystem(client *Client, systems ...string) error {
	var m struct {
		Payload struct {
			Domain Domain
			Errors []OpsLevelErrors
		} `graphql:"domainChildAssign(domain:$domain, childSystems:$childSystems)"`
	}
	v := PayloadVariables{
		"domain":       *NewIdentifier(string(domainId.Id)),
		"childSystems": NewIdentifierArray(systems),
	}
	err := client.Mutate(&m, v, WithName("DomainAssignSystem"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateDomain(input DomainInput) (*Domain, error) {
	var m struct {
		Payload struct {
			Domain Domain
			Errors []OpsLevelErrors
		} `graphql:"domainCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("DomainCreate"))
	return &m.Payload.Domain, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetDomain(identifier string) (*Domain, error) {
	var q struct {
		Account struct {
			Domain Domain `graphql:"domain(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Query(&q, v, WithName("DomainGet"))
	return &q.Account.Domain, HandleErrors(err, nil)
}

func (client *Client) ListDomains(variables *PayloadVariables) (*DomainConnection, error) {
	var q struct {
		Account struct {
			Domains DomainConnection `graphql:"domains(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("DomainsList")); err != nil {
		return nil, err
	}
	for q.Account.Domains.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Domains.PageInfo.End
		resp, err := client.ListDomains(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Domains.Nodes = append(q.Account.Domains.Nodes, resp.Nodes...)
		q.Account.Domains.PageInfo = resp.PageInfo
	}
	q.Account.Domains.TotalCount = len(q.Account.Domains.Nodes)
	return &q.Account.Domains, nil
}

func (client *Client) UpdateDomain(identifier string, input DomainInput) (*Domain, error) {
	var m struct {
		Payload struct {
			Domain Domain
			Errors []OpsLevelErrors
		} `graphql:"domainUpdate(domain:$domain,input:$input)"`
	}
	v := PayloadVariables{
		"domain": *NewIdentifier(identifier),
		"input":  input,
	}
	err := client.Mutate(&m, v, WithName("DomainUpdate"))
	return &m.Payload.Domain, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteDomain(identifier string) error {
	var d struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"domainDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&d, v, WithName("DomainDelete"))
	return HandleErrors(err, d.Payload.Errors)
}
