package opslevel

import (
	"errors"
	"fmt"
)

type AliasableResourceInterface interface {
	GetAliases() []string
	ResourceId() ID
	AliasableType() AliasOwnerTypeEnum
}

func (client *Client) GetAliasableResource(resourceType AliasOwnerTypeEnum, identifier string) (AliasableResourceInterface, error) {
	var err error
	var aliasableResource AliasableResourceInterface

	switch resourceType {
	case AliasOwnerTypeEnumService:
		if IsID(identifier) {
			aliasableResource, err = client.GetService(ID(identifier))
		} else {
			aliasableResource, err = client.GetServiceWithAlias(identifier)
		}
	case AliasOwnerTypeEnumTeam:
		if IsID(identifier) {
			aliasableResource, err = client.GetTeam(ID(identifier))
		} else {
			aliasableResource, err = client.GetTeamWithAlias(identifier)
		}
	case AliasOwnerTypeEnumSystem:
		aliasableResource, err = client.GetSystem(identifier)
	case AliasOwnerTypeEnumDomain:
		aliasableResource, err = client.GetDomain(identifier)
	case AliasOwnerTypeEnumInfrastructureResource:
		aliasableResource, err = client.GetInfrastructure(identifier)
	case AliasOwnerTypeEnumScorecard:
		aliasableResource, err = client.GetScorecard(identifier)
	default:
		err = fmt.Errorf("not an aliasable resource type '%s'", resourceType)
	}

	return aliasableResource, err
}

func (client *Client) CreateAliases(ownerId ID, aliases []string) ([]string, error) {
	var output []string
	var allErrors error
	for _, alias := range aliases {
		input := AliasCreateInput{
			Alias:   alias,
			OwnerId: ID(ownerId),
		}
		result, err := client.CreateAlias(input)
		allErrors = errors.Join(allErrors, err)
		output = append(output, result...)
	}
	output = removeDuplicates(output)
	return output, allErrors
}

func (client *Client) CreateAlias(input AliasCreateInput) ([]string, error) {
	var m struct {
		Payload struct {
			Aliases []string
			OwnerId string
			Errors  []Error
		} `graphql:"aliasCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AliasCreate"))
	output := make([]string, len(m.Payload.Aliases))
	copy(output, m.Payload.Aliases)
	return output, HandleErrors(err, m.Payload.Errors)
}

// Deprecated: use client.DeleteAlias instead
func (client *Client) DeleteInfraAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumInfrastructureResource,
	})
}

// Deprecated: use client.DeleteAlias instead
func (client *Client) DeleteServiceAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumService,
	})
}

// Deprecated: use client.DeleteAlias instead
func (client *Client) DeleteTeamAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumTeam,
	})
}

func (client *Client) DeleteAliases(aliasOwnerType AliasOwnerTypeEnum, aliases []string) error {
	var allErrors error

	for _, alias := range aliases {
		input := AliasDeleteInput{
			Alias:     alias,
			OwnerType: aliasOwnerType,
		}
		allErrors = errors.Join(allErrors, client.DeleteAlias(input))
	}

	return allErrors
}

func (client *Client) DeleteAlias(input AliasDeleteInput) error {
	var m struct {
		Payload struct {
			Alias  string `graphql:"deletedAlias"`
			Errors []Error
		} `graphql:"aliasDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AliasDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
