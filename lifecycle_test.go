package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2022"
)

func TestListLifecycles(t *testing.T) {
	// Arrange
	client := ATestClientSkipRequest(t, "lifecycles")
	// Act
	result, err := client.ListLifecycles()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "pre-alpha", result[0].Alias)
}
