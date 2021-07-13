package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestListRubricCategories(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/listRubricCategories", autopilot.FixtureResponse("rubric_category_response.json"), FixtureQueryValidation(t, "rubric_category_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, _ := client.ListCategories()
	// Assert
	autopilot.Equals(t, 7, len(result))
	autopilot.Equals(t, "Reliability", result[1].Name)
}
