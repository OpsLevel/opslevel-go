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
	request := `{{ template "scorecard_create_request" }}`
	response := `{{ template "scorecard_create_response" }}`
	name := "new scorecard"
	description := "a new scorecard with an attached filter id"

	client := ABetterTestClient(t, "scorecards/create_scorecard", request, response)
	sc, err := client.CreateScorecard(ol.ScorecardInput{
		Name:        name,
		Description: &description,
		OwnerId:     *fakeOwnerId,
		FilterId:    fakeFilterId,
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, name, sc.Name)
	autopilot.Equals(t, description, sc.Description)
	autopilot.Equals(t, *fakeOwnerId, sc.Owner.Id())
	autopilot.Equals(t, *fakeFilterId, sc.Filter.Id)
}

func TestUpdateScorecard(t *testing.T) {
	request := `{{ template "scorecard_update_request" }}`
	response := `{{ template "scorecard_update_response" }}`
	name := "updated scorecard"
	description := "this scorecard was updated"

	client := ABetterTestClient(t, "scorecards/update_scorecard", request, response)
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
}

func TestDeleteScorecard(t *testing.T) {
	request := `{{ template "scorecard_delete_request" }}`
	response := `{{ template "scorecard_delete_response" }}`

	client := ABetterTestClient(t, "scorecards/delete_scorecard", request, response)
	deletedScorecardId, err := client.DeleteScorecard(scorecardId)

	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID(scorecardId), deletedScorecardId)
}

func TestGetScorecard(t *testing.T) {
	request := `{{ template "scorecard_get_request" }}`
	response := `{{ template "scorecard_get_response" }}`
	name := "fetched scorecard"
	description := "hello there!"

	client := ABetterTestClient(t, "scorecards/get_scorecard", request, response)
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
	requests := []TestRequest{
		{
			`{"query": "{{ template "scorecard_list_query" }}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"scorecards": {
							"nodes": [
								{
									{{ template "scorecard_1_response" }}
								},
								{
									{{ template "scorecard_2_response" }}
								}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
						  }}}}`,
		},
		{
			`{"query": "{{ template "scorecard_list_query" }}",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"scorecards": {
							"nodes": [
								{
									{{ template "scorecard_3_response" }}
								}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }}}}`,
		},
	}
	// An easy way to see the results of templating is by uncommenting this
	// fmt.Print(Templated(request))
	// fmt.Print(Templated(response))
	// panic(1)

	client := APaginatedTestClient(t, "scorecards/list_scorecards", requests...)
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
