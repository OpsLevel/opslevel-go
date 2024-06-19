package opslevel

import (
	"fmt"
	"slices"
	"strings"
)

type AliasOwnerInterface interface {
	ReconcileAliases(*Client, []string) ([]string, error)
}

func (client *Client) CreateAliases(ownerId ID, aliases []string) ([]string, error) {
	var output []string
	var errors []string
	for _, alias := range aliases {
		input := AliasCreateInput{
			Alias:   alias,
			OwnerId: ownerId,
		}
		result, err := client.CreateAlias(input)
		if err != nil {
			errors = append(errors, err.Error())
		}
		output = append(output, result...)
	}
	output = removeDuplicates(output)
	if len(errors) > 0 {
		return output, fmt.Errorf(strings.Join(errors, "\n"))
	} else {
		return output, nil
	}
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

// ReconcileAliases manages aliases API operations for AliasOwnerInterface implementations
//
// Aliases not in 'aliasesWanted' will be deleted, new tags from 'aliasesWanted' will be created. Reconciled aliases are returned.
func (client *Client) ReconcileAliases(resourceType AliasOwnerInterface, aliasesWanted []string) ([]string, error) {
	return resourceType.ReconcileAliases(client, aliasesWanted)
}

// Given actual aliases and wanted aliases, returns aliasesToCreate and aliasesToDelete lists
func extractAliases(existingAliases, aliasesWanted []string) ([]string, []string) {
	var aliasesToCreate, aliasesToDelete []string

	for _, alias := range existingAliases {
		if !slices.Contains(aliasesWanted, alias) {
			aliasesToDelete = append(aliasesToDelete, alias)
		}
	}

	for _, aliasWanted := range aliasesWanted {
		if !slices.Contains(existingAliases, aliasWanted) {
			aliasesToCreate = append(aliasesToCreate, aliasWanted)
		}
	}
	return aliasesToCreate, aliasesToDelete
}
