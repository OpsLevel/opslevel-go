package opslevel

import (
	"fmt"
	"slices"
	"strings"
)

type AliasOwnerInterface interface {
	AliasOwnerType() AliasOwnerTypeEnum
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

func (client *Client) DeleteDomainAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumDomain,
	})
}

func (client *Client) DeleteInfraAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumInfrastructureResource,
	})
}

func (client *Client) DeleteScorecardAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumScorecard,
	})
}

func (client *Client) DeleteServiceAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumService,
	})
}

func (client *Client) DeleteSystemAlias(alias string) error {
	return client.DeleteAlias(AliasDeleteInput{
		Alias:     alias,
		OwnerType: AliasOwnerTypeEnumSystem,
	})
}

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
	var deleteAliasFunc func(string) error
	var existingAliases []string
	var resourceId ID

	switch resource := resourceType.(type) {
	case *Domain:
		deleteAliasFunc = client.DeleteDomainAlias
		existingAliases = resource.ManagedAliases
		resourceId = resource.Id
	case *InfrastructureResource:
		deleteAliasFunc = client.DeleteInfraAlias
		existingAliases = resource.Aliases
		resourceId = ID(resource.Id)
	case *Scorecard:
		deleteAliasFunc = client.DeleteScorecardAlias
		existingAliases = resource.Aliases
		resourceId = resource.Id
	case *Service:
		deleteAliasFunc = client.DeleteServiceAlias
		existingAliases = resource.ManagedAliases
		resourceId = resource.Id
	case *System:
		deleteAliasFunc = client.DeleteSystemAlias
		existingAliases = resource.Aliases
		resourceId = resource.Id
	case *Team:
		deleteAliasFunc = client.DeleteTeamAlias
		existingAliases = resource.ManagedAliases
		resourceId = resource.Id
	}

	// delete aliases found in resource but not listed in aliasesWanted
	for _, alias := range existingAliases {
		if !slices.Contains(aliasesWanted, alias) {
			if err := deleteAliasFunc(alias); err != nil {
				return []string{}, err
			}
		}
	}

	newServiceAliases := []string{}
	for _, aliasWanted := range aliasesWanted {
		if !slices.Contains(existingAliases, aliasWanted) {
			newServiceAliases = append(newServiceAliases, aliasWanted)
		}
	}

	return client.CreateAliases(resourceId, newServiceAliases)
}
