package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

var fakeOwnerId = ol.NewID("Z2lkOi8vMTIzNDU2Nzg5Cg==")

func TestCreateScorecard(t *testing.T) {
	request := `{{ template "scorecard_create_request" }}`
	response := `{{ template "scorecard_create_response" }}`

	client := ABetterTestClient(t, "scorecards/create_scorecard", request, response)
	sc, err := client.CreateScorecard(ol.ScorecardInput{
		Name:    "new scorecard",
		OwnerId: fakeOwnerId,
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, "new scorecard", sc.Name)
	autopilot.Equals(t, *fakeOwnerId, sc.Owner.Id())
}

func TestUpdateScorecard(t *testing.T) {
	request := `{{ template "scorecard_update_request" }}`
	response := `{{ template "scorecard_update_response" }}`

	client := ABetterTestClient(t, "scorecards/update_scorecard", request, response)
	sc, err := client.UpdateScorecard("Z2lkOi8vMTIzNDU2Nzg5MTAK", ol.ScorecardInput{
		Name:    "updated scorecard",
		OwnerId: fakeOwnerId,
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, "updated scorecard", sc.Name)
	autopilot.Equals(t, *fakeOwnerId, sc.Owner.Id())
}

func TestDeleteScorecard(t *testing.T) {
	request := `{{ template "scorecard_delete_request" }}`
	response := `{{ template "scorecard_delete_response" }}`

	client := ABetterTestClient(t, "scorecards/delete_scorecard", request, response)
	deletedScorecardId, err := client.DeleteScorecard("Z2lkOi8vMTIzNDU2Nzg5MTAK")

	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5MTAK"), deletedScorecardId)
}

func TestGetScorecard(t *testing.T) {
	request := `{{ template "scorecard_get_request" }}`
	response := `{{ template "scorecard_get_response" }}`

	client := ABetterTestClient(t, "scorecards/get_scorecard", request, response)
	scorecard, err := client.GetScorecard("Z2lkOi8vMTIzNDU2Nzg5MTAK")

	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5MTAK"), scorecard.Id)
	autopilot.Equals(t, *fakeOwnerId, scorecard.Owner.Id())
}

func TestListScorecards(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			`{"query": "query ScorecardsList($after:String!$first:Int!){account{scorecards(after: $after, first: $first){nodes{aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
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
			`{"query": "query ScorecardsList($after:String!$first:Int!){account{scorecards(after: $after, first: $first){nodes{aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
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
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, "first scorecard", result[0].Name)
	autopilot.Equals(t, "second scorecard", result[1].Name)
	autopilot.Equals(t, "third scorecard", result[2].Name)
}
