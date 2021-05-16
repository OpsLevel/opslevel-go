package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListTiers(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/tiers", autopilot.FixtureResponse("tiers_response.json"), autopilot.SkipRequestValidation())
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.ListTiers()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "tier_1", result[0].Alias)
}
