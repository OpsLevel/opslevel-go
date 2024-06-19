package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

type extractAliasesTestCase struct {
	aliasesWanted           []string
	existingAliases         []string
	expectedAliasesToCreate []string
	expectedAliasesToDelete []string
}

func TestExtractAliases(t *testing.T) {
	var noAliases []string
	aliases := []string{"foo", "alpha", "beta", "gamma"}
	// Arrange
	testCases := map[string]extractAliasesTestCase{
		"create all delete none": {
			aliasesWanted:           aliases,
			existingAliases:         noAliases,
			expectedAliasesToCreate: aliases,
			expectedAliasesToDelete: noAliases,
		},
		"create none delete all": {
			aliasesWanted:           noAliases,
			existingAliases:         aliases,
			expectedAliasesToCreate: noAliases,
			expectedAliasesToDelete: aliases,
		},
		"create some delete some": {
			aliasesWanted:           []string{"foo", "alpha", "beta"},
			existingAliases:         []string{"alpha", "beta", "gamma"},
			expectedAliasesToCreate: []string{"foo"},
			expectedAliasesToDelete: []string{"gamma"},
		},
		"no change": {
			aliasesWanted:           aliases,
			existingAliases:         aliases,
			expectedAliasesToCreate: noAliases,
			expectedAliasesToDelete: noAliases,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Act
			aliasesToCreate, aliasesToDelete := ol.ExtractAliases(tc.existingAliases, tc.aliasesWanted)

			// Assert
			autopilot.Equals(t, aliasesToCreate, tc.expectedAliasesToCreate)
			autopilot.Equals(t, aliasesToDelete, tc.expectedAliasesToDelete)
		})
	}
}

func TestCreateAliases(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input": { "alias": "MyAwesomeAlias", "ownerId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "MyCoolService", "MyAwesomeAlias" ], "ownerId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", "errors": [] }}}`,
	)

	client := BestTestClient(t, "aliases/create", testRequest)
	// Act
	result, err := client.CreateAliases("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", []string{"MyAwesomeAlias"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "MyAwesomeAlias", result[1])
}

func TestCreateAliasesOwnerNotFound(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input": { "alias": "MyAwesomeAlias", "ownerId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2" }}`,
		`{ "data": { "aliasCreate": { "aliases": null, "ownerId": null, "errors": [ { "message": "'Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2' does not identify any record on this account", "path": [ "ownerId" ] } ] } } }`,
	)

	client := BestTestClient(t, "aliases/create_owner_not_found", testRequest)
	// Act
	result, err := client.CreateAliases("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", []string{"MyAwesomeAlias"})
	// Assert
	if err == nil {
		t.Error("expected number of errors to be > 1")
	}
	autopilot.Equals(t, 0, len(result))
}

func TestDeleteServiceAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input": { "alias": "MyAwesomeAlias", "ownerType": "service" }}`,
		`{"data": { "aliasDelete": { "deletedAlias": "MyAwesomeAlias", "errors": [] }}}`,
	)

	client := BestTestClient(t, "aliases/delete_service", testRequest)
	// Act
	err := client.DeleteServiceAlias("MyAwesomeAlias")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteTeamAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input": { "alias": "MyAwesomeAlias", "ownerType": "team" }}`,
		`{"data": { "aliasDelete": { "deletedAlias": "MyAwesomeAlias", "errors": [] }}}`,
	)

	client := BestTestClient(t, "aliases/delete_team", testRequest)
	// Act
	err := client.DeleteTeamAlias("MyAwesomeAlias")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteAliasNotFound(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input": { "alias": "MyAwesomeAlias", "ownerType": "team" }}`,
		`{ "data": { "aliasDelete": { "deletedAlias": null, "errors": [ { "message": "'MyAwesomeAlias' does not identify a team on this account", "path": [ "alias" ] } ] } } }`,
	)

	client := BestTestClient(t, "aliases/delete_alias_not_found", testRequest)
	// Act
	err := client.DeleteTeamAlias("MyAwesomeAlias")
	// Assert
	if err == nil {
		t.Error("expected number of errors to be > 1")
	}
}
