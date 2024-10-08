package opslevel

import (
	"errors"
)

func (client *Client) CreateAliases(ownerId ID, aliases []string) ([]string, error) {
	var output []string
	var allErrors error
	for _, alias := range aliases {
		input := AliasCreateInput{
			Alias:   alias,
			OwnerId: ownerId,
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
			Errors  []OpsLevelErrors
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
			Errors []OpsLevelErrors
		} `graphql:"aliasDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AliasDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
