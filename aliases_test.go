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
			aliasesToCreate, aliasesToDelete := ol.TestExtractAliases(tc.existingAliases, tc.aliasesWanted)

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

func TestGetAliasableResource(t *testing.T) {
	service := autopilot.Register[ol.ServiceId]("example_service",
		ol.ServiceId{
			Id:      id1,
			Aliases: []string{alias1, alias2, alias3},
		},
	)
	team := autopilot.Register[ol.Team]("example_team",
		ol.Team{
			TeamId: ol.TeamId{
				Id: id1,
			},
			Aliases: []string{alias1, alias2, alias4},
		},
	)
	system := autopilot.Register[ol.SystemId]("example_system",
		ol.SystemId{
			Id:      id1,
			Aliases: []string{alias3},
		})

	requests := []autopilot.TestRequest{
		autopilot.NewTestRequest(
			`query ServiceGet($service:ID!){account{service(id: $service){{ template "service_get" }}}}}`,
			`{"service": "{{ template "id1_string" }}"}`,
			`{ "data": { "account": {"service": {{ template "example_service" }}}}}`,
		),
		autopilot.NewTestRequest(
			`query ServiceGet($service:String!){account{service(alias: $service){{ template "service_get" }}}}}`,
			`{"service": "{{ template "alias1" }}"}`,
			`{ "data": { "account": {"service": {{ template "example_service" }}}}}`,
		),
		autopilot.NewTestRequest(
			`query TeamGet($id:ID!){account{team(id: $id){{ template "team_get" }}}}`,
			`{"id": "{{ template "id1_string" }}"}`,
			`{ "data": { "account": {"team": {{ template "example_team" }}}}}`,
		),
		autopilot.NewTestRequest(
			`query TeamGet($alias:String!){account{team(alias: $alias){{ template "team_get" }}}}`,
			`{"alias": "{{ template "alias1" }}"}`,
			`{ "data": { "account": {"team": {{ template "example_team" }}}}}`,
		),
		autopilot.NewTestRequest(
			`query SystemGet($input:IdentifierInput!){account{system(input: $input){{ template "system_get" }}}}`,
			`{"input": {"alias": "{{ template "alias1" }}"}}`,
			`{ "data": { "account": {"system": {{ template "example_system"}}}}}`,
		),
		autopilot.NewTestRequest(
			`{{ template "scorecard_get_request" }}`,
			`{{ template "scorecard_get_request_vars" }}`,
			`{{ template "scorecard_get_response" }}`,
		),
		autopilot.NewTestRequest(
			`query InfrastructureResourceGet($all:Boolean!$input:IdentifierInput!){account{infrastructureResource(input: $input){{ template "infra_get" }}}}`,
			`{"all": true, "input":{ {{ template "id1" }} }}`,
			`{"data": { "account": { "infrastructureResource": {{ template "infra_1" }} }}}`,
		),
	}
	client := BestTestClient(t, "tags/get_aliasable_resource", requests...)
	// Act
	service1, err1 := client.GetAliasableResource(ol.AliasOwnerTypeEnumService, string(id1))
	service2, err2 := client.GetAliasableResource(ol.AliasOwnerTypeEnumService, alias1)
	team1, err3 := client.GetAliasableResource(ol.AliasOwnerTypeEnumTeam, string(id1))
	team2, err4 := client.GetAliasableResource(ol.AliasOwnerTypeEnumTeam, alias1)
	system1, err5 := client.GetAliasableResource(ol.AliasOwnerTypeEnumSystem, alias1)
	scorecard1, err6 := client.GetAliasableResource(ol.AliasOwnerTypeEnumScorecard, scorecardId)
	infra1, err7 := client.GetAliasableResource(ol.AliasOwnerTypeEnumInfrastructureResource, string(id1))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, service.Aliases, service1.GetAliases())
	autopilot.Ok(t, err2)
	autopilot.Equals(t, service.Aliases, service2.GetAliases())
	autopilot.Ok(t, err3)
	autopilot.Equals(t, team.Aliases, team1.GetAliases())
	autopilot.Ok(t, err4)
	autopilot.Equals(t, team.Aliases, team2.GetAliases())
	autopilot.Ok(t, err5)
	autopilot.Equals(t, system.Aliases, system1.GetAliases())
	autopilot.Ok(t, err6)
	autopilot.Equals(t, []string{"existing_scorecard"}, scorecard1.GetAliases())
	autopilot.Ok(t, err7)
	autopilot.Equals(t, []string{}, infra1.GetAliases())
}
