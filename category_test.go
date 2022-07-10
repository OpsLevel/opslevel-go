package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/create")
	// Act
	result, _ := client.CreateCategory(ol.CategoryCreateInput{
		Name: "Kyle",
	})
	// Assert
	autopilot.Equals(t, "Kyle", result.Name)
}

func TestGetRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/get")
	// Act
	result, err := client.GetCategory("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Reliability", result.Name)
}

func TestGetMissingRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/get_missing")
	// Act
	_, err := client.GetCategory("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListRubricCategories(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/list")
	// Act
	result, _ := client.ListCategories()
	// Assert
	autopilot.Equals(t, 7, len(result))
	autopilot.Equals(t, "Reliability", result[1].Name)
}

func TestUpdateRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/update")
	// Act
	result, _ := client.UpdateCategory(ol.CategoryUpdateInput{
		Id:   ol.NewID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz"),
		Name: "Emily",
	})
	// Assert
	autopilot.Equals(t, "Emily", result.Name)
}

func TestDeleteRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/delete")
	// Act
	err := client.DeleteCategory("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
