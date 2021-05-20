package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestAssignTagForAlias(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/AssignTagForAlias", autopilot.FixtureResponse("tagAssign_response.json"), FixtureQueryValidation(t, "tagAssignWithAlias_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.AssignTagForAlias("coredns", "hello", "world")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestAssignTagForId(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/AssignTagForId", autopilot.FixtureResponse("tagAssign_response.json"), FixtureQueryValidation(t, "tagAssignWithId_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.AssignTagForId("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", "hello", "world")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}
