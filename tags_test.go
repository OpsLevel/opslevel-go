package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestAssignTagForAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "tagAssign", "tagAssignWithAlias")
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
	client := ATestClientAlt(t, "tagAssign", "tagAssignWithId")
	// Act
	result, err := client.AssignTagForId("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", "hello", "world")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestAssignTagThrowsValidationError(t *testing.T) {
	// Arrange
	client := ATestClient(t, "tagAssignThrowsValidationError")
	// Act
	_, err := client.AssignTagForId("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", "Hello", "world")
	// Assert
	autopilot.Equals(t, "invalid tag key name 'Hello' - must start with a letter and only lowercase alphanumerics, underscores, hyphens, periods, and slashes are allowed.", err.Error())
}
