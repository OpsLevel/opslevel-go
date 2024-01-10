package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
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
		Description:                 &description,
		OwnerId:                     *fakeOwnerId,
		FilterId:                    fakeFilterId,
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
		Description:                 &description,
		OwnerId:                     *fakeOwnerId,
		FilterId:                    fakeFilterId,
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
		Description: &description,
		Name:        name,
		OwnerId:     *newOwnerId,
		FilterId:    newFilterId,
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
	autopilot.Equals(t, 30, sc.ChecksCount)
}

func TestListScorecards(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`{{ template "scorecard_list_query" }}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "scorecards": { "nodes": [ { {{ template "scorecard_1_response" }} }, { {{ template "scorecard_2_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`{{ template "scorecard_list_query" }}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "scorecards": { "nodes": [ { {{ template "scorecard_3_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "scorecards/list_scorecards", requests...)
	// Act
	response, err := client.ListScorecards(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
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
