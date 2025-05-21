package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

var (
	scorecardId  = "Z2lkOi8vMTIzNDU2Nzg5MTAK"           // 12345678910
	fakeOwnerId  = ol.NewID("Z2lkOi8vMTIzNDU2Nzg5Cg==") // 123456789
	newOwnerId   = ol.NewID("Z2lkOi8vMTIzNDU2Nzc3Cg==") // 123456777
	fakeFilterId = ol.NewID("Z2lkOi8vMTIzNDU2MTIzCg==") // 123456123
	newFilterId  = ol.NewID("Z2lkOi8vMTIzNDU2NDU2Cg==") // 123456456
)

func TestCreateScorecard(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`{{ template "scorecard_create_request" }}`,
		`{{ template "scorecard_create_request_vars" }}`,
		`{{ template "scorecard_create_response" }}`,
	)
	name := "new scorecard"
	description := "a new scorecard with an attached filter id"

	client := BestTestClient(t, "scorecards/create_scorecard", testRequest)
	sc, err := client.CreateScorecard(ol.ScorecardInput{
		Name:                        name,
		Description:                 ol.RefOf(description),
		OwnerId:                     *fakeOwnerId,
		FilterId:                    ol.RefOf(*fakeFilterId),
		AffectsOverallServiceLevels: ol.RefOf(true),
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, name, sc.Name)
	autopilot.Equals(t, description, sc.Description)
	autopilot.Equals(t, *fakeOwnerId, sc.Owner.Id())
	autopilot.Equals(t, *fakeFilterId, sc.Filter.Id)
	autopilot.Equals(t, true, sc.AffectsOverallServiceLevels)
}

func TestCreateScorecardDoesNotAffectServiceLevels(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`{{ template "scorecard_create_request" }}`,
		`{{ template "scorecard_create_request_vars_affects_service_levels_false" }}`,
		`{{ template "scorecard_create_response_affects_service_levels_false" }}`,
	)
	name := "new scorecard"
	description := "a new scorecard with an attached filter id"

	client := BestTestClient(t, "scorecards/create_scorecard_not_affects_service_levels", testRequest)
	sc, err := client.CreateScorecard(ol.ScorecardInput{
		Name:                        name,
		Description:                 ol.RefOf(description),
		OwnerId:                     *fakeOwnerId,
		FilterId:                    ol.RefOf(*fakeFilterId),
		AffectsOverallServiceLevels: ol.RefOf(false),
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, name, sc.Name)
	autopilot.Equals(t, description, sc.Description)
	autopilot.Equals(t, *fakeOwnerId, sc.Owner.Id())
	autopilot.Equals(t, *fakeFilterId, sc.Filter.Id)
	autopilot.Equals(t, false, sc.AffectsOverallServiceLevels)
}

func TestUpdateScorecard(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`{{ template "scorecard_update_request" }}`,
		`{{ template "scorecard_update_request_vars" }}`,
		`{{ template "scorecard_update_response" }}`,
	)

	name := "updated scorecard"
	description := "this scorecard was updated"

	client := BestTestClient(t, "scorecards/update_scorecard", testRequest)
	sc, err := client.UpdateScorecard(scorecardId, ol.ScorecardInput{
		Description: ol.RefOf(description),
		Name:        name,
		OwnerId:     *newOwnerId,
		FilterId:    ol.RefOf(*newFilterId),
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, *ol.NewID(scorecardId), sc.Id)
	autopilot.Equals(t, name, sc.Name)
	autopilot.Equals(t, description, sc.Description)
	autopilot.Equals(t, *newOwnerId, sc.Owner.Id())
	autopilot.Equals(t, *newFilterId, sc.Filter.Id)
	autopilot.Equals(t, false, sc.AffectsOverallServiceLevels)
}

func TestDeleteScorecard(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`{{ template "scorecard_delete_request" }}`,
		`{{ template "scorecard_delete_request_vars" }}`,
		`{{ template "scorecard_delete_response" }}`,
	)
	client := BestTestClient(t, "scorecards/delete_scorecard", testRequest)
	deletedScorecardId, err := client.DeleteScorecard(scorecardId)

	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.NewID(scorecardId), deletedScorecardId)
}

func TestGetScorecard(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`{{ template "scorecard_get_request" }}`,
		`{{ template "scorecard_get_request_vars" }}`,
		`{{ template "scorecard_get_response" }}`,
	)
	name := "fetched scorecard"
	description := "hello there!"

	client := BestTestClient(t, "scorecards/get_scorecard", testRequest)
	sc, err := client.GetScorecard(scorecardId)

	autopilot.Ok(t, err)
	autopilot.Equals(t, *ol.NewID(scorecardId), sc.Id)
	autopilot.Equals(t, name, sc.Name)
	autopilot.Equals(t, description, sc.Description)
	autopilot.Equals(t, *fakeOwnerId, sc.Owner.Id())
	autopilot.Equals(t, *fakeFilterId, sc.Filter.Id)
	autopilot.Equals(t, 10, sc.PassingChecks)
	autopilot.Equals(t, 20, sc.ServiceCount)
	autopilot.Equals(t, 30, sc.TotalChecks)
}

func TestListScorecards(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`{{ template "scorecard_list_query" }}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "scorecards": { "nodes": [ { {{ template "scorecard_1_response" }} }, { {{ template "scorecard_2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`{{ template "scorecard_list_query" }}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "scorecards": { "nodes": [ { {{ template "scorecard_3_response" }} } ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "scorecards/list_scorecards", requests...)
	// Act
	response, err := client.ListScorecards(nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "first scorecard", result[0].Name)
	autopilot.Equals(t, ol.ConnectiveEnumAnd, result[0].Filter.Connective)
	autopilot.Equals(t, *fakeOwnerId, result[0].Owner.Id())
	autopilot.Equals(t, 150, result[0].ServiceCount)
	autopilot.Equals(t, "second scorecard", result[1].Name)
	autopilot.Equals(t, ol.ConnectiveEnumOr, result[1].Filter.Connective)
	autopilot.Equals(t, *fakeOwnerId, result[0].Owner.Id())
	autopilot.Equals(t, 22, result[1].ServiceCount)
	autopilot.Equals(t, "third scorecard", result[2].Name)
	autopilot.Equals(t, ol.Filter{}, result[2].Filter)
	autopilot.Equals(t, *newOwnerId, result[2].Owner.Id())
	autopilot.Equals(t, 33, result[2].ServiceCount)
}

func TestListScorecardCategories(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ScorecardCategoryList($after:String!$first:Int!$scorecard:IdentifierInput!){account{scorecard(input: $scorecard){categories(after: $after, first: $first){nodes{description,id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}`,
		`{ {{ template "first_page_variables" }}, "scorecard": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "scorecard": { "categories": { "nodes": [ { {{ template "id2" }}, "name": "quality" } ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ScorecardCategoryList($after:String!$first:Int!$scorecard:IdentifierInput!){account{scorecard(input: $scorecard){categories(after: $after, first: $first){nodes{description,id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}`,
		`{ {{ template "second_page_variables" }}, "scorecard": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "scorecard": { "categories": { "nodes": [ { {{ template "id3" }}, "name": "ownership" } ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "scorecard/categories", requests...)
	// Act
	scorecard := ol.Scorecard{
		ScorecardId: ol.ScorecardId{
			Id: id1,
		},
	}
	resp, err := scorecard.ListCategories(client, nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert

	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, id2, result[0].Id)
	autopilot.Equals(t, "quality", result[0].Name)
	autopilot.Equals(t, id3, result[1].Id)
	autopilot.Equals(t, "ownership", result[1].Name)
}

func TestScorecardReconcileAliasesDeleteAll(t *testing.T) {
	// Arrange
	aliasesWanted := []string{}
	scorecard := ol.ScorecardId{
		Id:      id1,
		Aliases: []string{"one", "two"},
	}

	// delete "one" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "one", "ownerType": "scorecard" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "two" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerType": "scorecard" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// get service
	testRequestThree := autopilot.NewTestRequest(
		`{{ template "scorecard_get_request" }}`,
		`{"input": { {{ template "id1" }} }}`,
		`{ "data": { "account": { "scorecard": { {{ template "id1" }}, "aliases": [] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree}
	client := BestTestClient(t, "scorecard/reconcile_aliases_delete_all", requests...)

	// Act
	err := scorecard.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, len(scorecard.Aliases), 0)
}

func TestScorecardReconcileAliases(t *testing.T) {
	// Arrange
	aliasesWanted := []string{"one", "two", "three"}
	scorecard := ol.ScorecardId{
		Id:      id1,
		Aliases: []string{"one", "alpha", "beta"},
	}

	// delete "alpha" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "alpha", "ownerType": "scorecard" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "beta" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "beta", "ownerType": "scorecard" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// create "two" alias
	testRequestThree := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "one", "two" ], "ownerId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)
	// create "three" alias
	testRequestFour := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input":{ "alias": "three", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "one", "two", "three" ], "ownerId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)
	// get service
	testRequestFive := autopilot.NewTestRequest(
		`{{ template "scorecard_get_request" }}`,
		`{"input": { {{ template "id1" }} }}`,
		`{ "data": { "account": { "scorecard": { {{ template "id1" }}, "aliases": ["one", "two", "three"] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree, testRequestFour, testRequestFive}
	client := BestTestClient(t, "scorecard/reconcile_aliases", requests...)

	// Act
	err := scorecard.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, scorecard.Aliases, aliasesWanted)
}
