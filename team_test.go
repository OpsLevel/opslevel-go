package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListTeams(t *testing.T) {
	// Arrange
	client := ATestClient(t, "teams")
	// Act
	result, err := client.ListTeams()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "devops", result[0].Alias)
	autopilot.Equals(t, "developers", result[1].Alias)
}
