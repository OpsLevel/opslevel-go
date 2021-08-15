package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestCreateAliases(t *testing.T) {
	// Arrange
	client := ATestClient(t, "aliasCreate")
	// Act
	result, err := client.CreateAliases("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", []string{"MyAwesomeService"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "MyAwesomeService", result[1])
}
