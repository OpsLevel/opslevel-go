package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateRubricLevels(t *testing.T) {
	// Arrange
	client := ANewClient(t, "rubric/level/create")
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
	client := ANewClient(t, "rubric/level/list")
	// Act
	result, _ := client.ListLevels()
	// Assert
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "Bronze", result[1].Name)
}

func TestUpdateRubricLevels(t *testing.T) {
	// Arrange
	client := ANewClient(t, "rubric/level/update")
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
	client := ANewClient(t, "rubric/level/delete")
	// Act
	err := client.DeleteLevel("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw")
	// Assert
	autopilot.Equals(t, nil, err)
}
