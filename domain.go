package opslevel

import (
	"errors"
	"fmt"
	"html"
	"slices"
)

// Returns unique identifiers created by OpsLevel, values in Aliases but not ManagedAliases
func (d *Domain) UniqueIdentifiers() []string {
	uniqueIdentifiers := []string{}
	for _, alias := range d.Aliases {
		if !slices.Contains(d.ManagedAliases, alias) {
			uniqueIdentifiers = append(uniqueIdentifiers, alias)
		}
	}

	return uniqueIdentifiers
}

func (d *Domain) Hydrate(client *Client) error {
	d.Description = html.UnescapeString(d.Description)
	d.Note = html.UnescapeString(d.Note)
	return nil
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
	if q.Account.Domain.Tags.PageInfo.HasNextPage {
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
	}
	q.Account.Domain.Tags.TotalCount = len(q.Account.Domain.Tags.Nodes)
	return &q.Account.Domain.Tags, nil
}

func (domainId *DomainId) GetAliases() []string {
	return domainId.Aliases
}

func (domainId *DomainId) ResourceId() ID {
	return domainId.Id
}

func (domainId *DomainId) ResourceType() TaggableResource {
	return TaggableResourceDomain
}

func (DomainId *DomainId) AliasableType() AliasOwnerTypeEnum {
	return AliasOwnerTypeEnumDomain
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
	if q.Account.Domain.ChildSystems.PageInfo.HasNextPage {
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
		Payload DomainPayload `graphql:"domainChildAssign(domain:$domain, childSystems:$childSystems)"`
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
		Payload DomainPayload `graphql:"domainCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v, WithName("DomainCreate")); err != nil {
		return nil, err
	}
	if err := m.Payload.Domain.Hydrate(client); err != nil {
		return &m.Payload.Domain, err
	}
	return &m.Payload.Domain, HandleErrors(nil, m.Payload.Errors)
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
	if err := client.Query(&q, v, WithName("DomainGet")); err != nil {
		return nil, err
	}
	if err := q.Account.Domain.Hydrate(client); err != nil {
		return &q.Account.Domain, err
	}
	return &q.Account.Domain, nil
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
	if q.Account.Domains.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Domains.PageInfo.End
		resp, err := client.ListDomains(variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Domains.Nodes = append(q.Account.Domains.Nodes, node)
		}
		q.Account.Domains.PageInfo = resp.PageInfo
	}
	q.Account.Domains.TotalCount = len(q.Account.Domains.Nodes)
	return &q.Account.Domains, nil
}

func (client *Client) UpdateDomain(identifier string, input DomainInput) (*Domain, error) {
	var m struct {
		Payload DomainPayload `graphql:"domainUpdate(domain:$domain,input:$input)"`
	}
	v := PayloadVariables{
		"domain": *NewIdentifier(identifier),
		"input":  input,
	}
	if err := client.Mutate(&m, v, WithName("DomainUpdate")); err != nil {
		return nil, err
	}
	if err := m.Payload.Domain.Hydrate(client); err != nil {
		return &m.Payload.Domain, err
	}
	return &m.Payload.Domain, HandleErrors(nil, m.Payload.Errors)
}

func (client *Client) DeleteDomain(identifier string) error {
	var d struct {
		Payload BasePayload `graphql:"domainDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Mutate(&d, v, WithName("DomainDelete"))
	return HandleErrors(err, d.Payload.Errors)
}
