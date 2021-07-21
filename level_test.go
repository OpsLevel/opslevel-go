package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateRubricLevels(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/createRubricLevels", autopilot.FixtureResponse("rubric_level_create_response.json"), FixtureQueryValidation(t, "rubric_level_create_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, _ := client.CreateLevel(LevelCreateInput{
		Name:        "Kyle",
		Description: "Created By Kyle",
	})
	// Assert
	autopilot.Equals(t, "kyle", result.Alias)
	autopilot.Equals(t, 4, result.Index)
}

func TestListRubricLevels(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/listRubricLevels", autopilot.FixtureResponse("rubric_level_response.json"), FixtureQueryValidation(t, "rubric_level_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, _ := client.ListLevels()
	// Assert
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "Bronze", result[1].Name)
}

func TestUpdateRubricLevels(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/updateRubricLevels", autopilot.FixtureResponse("rubric_level_update_response.json"), FixtureQueryValidation(t, "rubric_level_update_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, _ := client.UpdateLevel(LevelUpdateInput{
		Id:          graphql.ID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw"),
		Name:        "Kyle",
		Description: "Updated By Kyle",
	})
	// Assert
	autopilot.Equals(t, "kyle", result.Alias)
	autopilot.Equals(t, "Updated By Kyle", result.Description)
}

func TestDeleteRubricLevels(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/deleteRubricLevels", autopilot.FixtureResponse("rubric_level_delete_response.json"), FixtureQueryValidation(t, "rubric_level_delete_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	err := client.DeleteLevel("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw")
	// Assert
	autopilot.Equals(t, nil, err)
}
