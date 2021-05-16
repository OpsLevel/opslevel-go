package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestGetServiceWithAlias(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/service_alias", autopilot.FixtureResponse("service_response.json"), autopilot.SkipRequestValidation())
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetServiceWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, []byte("alpha"), []byte(result.Lifecycle.Alias))
	autopilot.Equals(t, []byte("developers"), []byte(result.Owner.Alias))
	autopilot.Equals(t, []byte("tier_1"), []byte(result.Tier.Alias))
	autopilot.Equals(t, 3, result.Tags.TotalCount)
	autopilot.Equals(t, 1, result.Tools.TotalCount)
}

func TestGetService(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/service", autopilot.FixtureResponse("service_response.json"), autopilot.SkipRequestValidation())
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, []byte("alpha"), []byte(result.Lifecycle.Alias))
	autopilot.Equals(t, []byte("developers"), []byte(result.Owner.Alias))
	autopilot.Equals(t, []byte("tier_1"), []byte(result.Tier.Alias))
	autopilot.Equals(t, 3, result.Tags.TotalCount)
	autopilot.Equals(t, 1, result.Tools.TotalCount)
}
