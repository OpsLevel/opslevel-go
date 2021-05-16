package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListTeams(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/teams", autopilot.FixtureResponse("teams_response.json"), FixtureQueryValidation(t, "teams_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.ListTeams()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "devops", result[0].Alias)
	autopilot.Equals(t, "developers", result[1].Alias)
}
