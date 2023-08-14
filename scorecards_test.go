package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

func TestCreateScorecard(t *testing.T) {
	request := `{{ template "scorecard_create_request" }}`
	response := `{{ template "scorecard_create_response" }}`

	client := ABetterTestClient(t, "scorecards/create_scorecard", request, response)
	sc, err := client.CreateScorecard(ol.ScorecardInput{
		Name:  "new scorecard",
		Owner: ol.IdentifierInput{Alias: "platform"},
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, "new scorecard", sc.Name)
	autopilot.Equals(t, "platform", sc.Owner.OnTeam.Alias)
}

func TestUpdateScorecard(t *testing.T) {
	request := `{{ template "scorecard_update_request" }}`
	response := `{{ template "scorecard_update_response" }}`

	client := ABetterTestClient(t, "scorecards/update_scorecard", request, response)
	sc, err := client.UpdateScorecard(ol.IdentifierInput{Id: "scorecard-id"}, ol.ScorecardInput{
		Name:  "updated scorecard",
		Owner: ol.IdentifierInput{Id: "team-id"},
	})

	autopilot.Ok(t, err)
	autopilot.Equals(t, "updated scorecard", sc.Name)
	autopilot.Equals(t, ol.ID("team-id"), sc.Owner.OnTeam.Id)
}

func TestDeleteScorecard(t *testing.T) {
	request := `{{ template "scorecard_delete_request" }}`
	response := `{{ template "scorecard_delete_response" }}`

	client := ABetterTestClient(t, "scorecards/delete_scorecard", request, response)
	deletedScorecardId, err := client.DeleteScorecard(ol.IdentifierInput{Alias: "scorecard-alias"})

	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("scorecard-id"), deletedScorecardId)
}

func TestGetScorecard(t *testing.T) {
	request := `{{ template "scorecard_get_request" }}`
	response := `{{ template "scorecard_get_response" }}`

	client := ABetterTestClient(t, "scorecards/get_scorecard", request, response)
	scorecard, err := client.GetScorecard(ol.IdentifierInput{Id: "scorecard-id"})

	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("scorecard-id"), scorecard.Id)
	// this test does not care about group vs team alias,
	// so both teamAlias and groupAlias can be used in the template successfully
	// this needs to be fixed on the API level
}
