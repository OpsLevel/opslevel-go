package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateTool(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/mutate_toolCreate", "mutate_toolCreate.json")
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
	autopilot.Equals(t, graphql.ID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2"), result.Service.Id)
	autopilot.Equals(t, "https://example.com", result.Url)
}

func TestGetToolsForService(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/query_tools", "query_tools.json")
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetToolsForService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestGetToolsCount(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/query_tools_count", "query_tools.json")
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetToolCount("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, result)
}
