package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestCreateAliases(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/aliasCreate", autopilot.FixtureResponse("aliasCreate_response.json"), FixtureQueryValidation(t, "aliasCreate_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result := client.CreateAliases("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", []string{"MyAwesomeService"})
	// Assert
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "MyAwesomeService", result[1])
}
