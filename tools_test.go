package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestCreateTool(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/toolCreate", autopilot.FixtureResponse("toolCreate_response.json"), autopilot.SkipRequestValidation())
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.CreateTool(ToolCreateInput{
		Category:    ToolCategoryOther,
		DisplayName: "example",
		Url:         "https://example.com",
		ServiceId:   "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", result.Service.Id)
	autopilot.Equals(t, ToolCategoryOther, result.Category)
	autopilot.Equals(t, "example", result.DisplayName)
	autopilot.Equals(t, "https://example.com", result.Url)
}

func TestGetToolsForService(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/tools", autopilot.FixtureResponse("tools_response.json"), FixtureQueryValidation(t, "tools_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetToolsForService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestGetToolsCount(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/tools_count", autopilot.FixtureResponse("tools_count_response.json"), FixtureQueryValidation(t, "tools_count_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetToolCount("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, result)
}
