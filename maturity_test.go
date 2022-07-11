package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListServicesMaturity(t *testing.T) {
	// Arrange
	client := ATestClient(t, "maturity/services")
	// Act
	result, err := client.ListServicesMaturity()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "Example", result[0].Name)
	autopilot.Equals(t, "Gold", result[0].MaturityReport.Get("Quality").Name)
}
