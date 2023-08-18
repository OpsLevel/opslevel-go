package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
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
