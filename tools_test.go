package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
)

func TestCreateTool(t *testing.T) {
	// Arrange
	client := ATestClientSkipRequest(t, "toolCreate")
	// Act
	result, err := client.CreateTool(ol.ToolCreateInput{
		Category:    ol.ToolCategoryOther,
		DisplayName: "example",
		Url:         "https://example.com",
		ServiceId:   "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", result.Service.Id)
	autopilot.Equals(t, ol.ToolCategoryOther, result.Category)
	autopilot.Equals(t, "example", result.DisplayName)
	autopilot.Equals(t, "https://example.com", result.Url)
}

func TestGetToolsForService(t *testing.T) {
	// Arrange
	client := ATestClient(t, "tools")
	// Act
	result, err := client.GetToolsForService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestGetToolsCount(t *testing.T) {
	// Arrange
	client := ATestClient(t, "tools_count")
	// Act
	result, err := client.GetToolCount("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, result)
}
