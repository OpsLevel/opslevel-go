package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot/v2022"
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

func TestUpdateRubricLevel(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/update")
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          ol.NewID("MTIzNDU2Nzg5MTIzNDU2Nzg5"),
		Name:        "Example",
		Description: ol.NewString("An example description"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelNoName(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "rubric/level/update", "rubric/level/update_noname")
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          ol.NewID("MTIzNDU2Nzg5MTIzNDU2Nzg5"),
		Description: ol.NewString("An example description"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelEmptyDescription(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "rubric/level/update", "rubric/level/update_emptydescription")
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          ol.NewID("MTIzNDU2Nzg5MTIzNDU2Nzg5"),
		Name:        "Example",
		Description: ol.NewEmptyString(),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelNoDescription(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "rubric/level/update", "rubric/level/update_nodescription")
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:   ol.NewID("MTIzNDU2Nzg5MTIzNDU2Nzg5"),
		Name: "Example",
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestDeleteRubricLevels(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/level/delete")
	// Act
	err := client.DeleteLevel("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw")
	// Assert
	autopilot.Equals(t, nil, err)
}
