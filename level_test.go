package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListRubricLevels(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/listRubricLevels", autopilot.FixtureResponse("rubric_level_response.json"), FixtureQueryValidation(t, "rubric_level_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, _ := client.ListLevels()
	// Assert
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "Bronze", result[1].Name)
}
