package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListLifecycles(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/query_lifecycles", "query_lifecycles.json")
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.ListLifecycles()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "pre-alpha", result[0].Alias)
}
