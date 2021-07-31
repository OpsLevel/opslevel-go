package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/create")
	// Act
	result, _ := client.CreateCategory(ol.CategoryCreateInput{
		Name:        "Kyle",
		Description: "Created By Kyle",
	})
	// Assert
	autopilot.Equals(t, "Kyle", result.Name)
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
