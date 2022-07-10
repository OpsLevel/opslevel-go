package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
)

func TestCreateRubricLevels(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/create")
	// Act
	result, _ := client.CreateLevel(ol.LevelCreateInput{
		Name:        "Kyle",
		Description: "Created By Kyle",
		Index:       ol.NewInt(4),
	})
	// Assert
	autopilot.Equals(t, "kyle", result.Alias)
	autopilot.Equals(t, 4, result.Index)
}

func TestGetRubricLevel(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/get")
	// Act
	result, err := client.GetLevel("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Bronze", result.Name)
}

func TestGetMissingRubricLevel(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/get_missing")
	// Act
	_, err := client.GetLevel("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListRubricLevels(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/list")
	// Act
	result, _ := client.ListLevels()
	// Assert
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "Bronze", result[1].Name)
}

func TestUpdateRubricLevels(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/update")
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          ol.NewID("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw"),
		Name:        "Kyle",
		Description: "Updated By Kyle",
	})
	// Assert
	autopilot.Equals(t, "kyle", result.Alias)
	autopilot.Equals(t, "Updated By Kyle", result.Description)
}

func TestDeleteRubricLevels(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/delete")
	// Act
	err := client.DeleteLevel("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw")
	// Assert
	autopilot.Equals(t, nil, err)
}
