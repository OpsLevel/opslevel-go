package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestGetServiceIdWithAlias(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/service_id_with_alias", autopilot.FixtureResponse("service_id_response.json"), FixtureQueryValidation(t, "service_id_with_alias_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetServiceIdWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", result.Id)
}

func TestGetServiceWithAlias(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/service_alias", autopilot.FixtureResponse("service_response.json"), FixtureQueryValidation(t, "service_with_alias_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetServiceWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
	autopilot.Equals(t, 1, result.Tools.TotalCount)
}

func TestGetService(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/service", autopilot.FixtureResponse("service_response.json"), FixtureQueryValidation(t, "service_with_id_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
	autopilot.Equals(t, 1, result.Tools.TotalCount)
}
