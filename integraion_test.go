package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestGetIntegraion(t *testing.T) {
	// Arrange
	client := ATestClient(t, "integration/get")
	// Act
	result, err := client.GetIntegration("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf", result.Id)
	autopilot.Equals(t, "Deploy", result.Name)
}

func TestGetMissingIntegraion(t *testing.T) {
	// Arrange
	client := ATestClient(t, "integration/get_missing")
	// Act
	_, err := client.GetIntegration("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListIntegrations(t *testing.T) {
	// Arrange
	client := ATestClient(t, "integration/list")
	// Act
	result, _ := client.ListIntegrations()
	// Assert
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "Deploy", result[1].Name)
}
