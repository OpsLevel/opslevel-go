package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	client := ANewClient(t, "rubric/category/create")
	// Act
	result, _ := client.CreateCategory(CategoryCreateInput{
		Name:        "Kyle",
		Description: "Created By Kyle",
	})
	// Assert
	autopilot.Equals(t, "Kyle", result.Name)
}

func TestListRubricCategories(t *testing.T) {
	// Arrange
	client := ANewClient(t, "rubric/category/list")
	// Act
	result, _ := client.ListCategories()
	// Assert
	autopilot.Equals(t, 7, len(result))
	autopilot.Equals(t, "Reliability", result[1].Name)
}

func TestUpdateRubricCategory(t *testing.T) {
	// Arrange
	client := ANewClient(t, "rubric/category/update")
	// Act
	result, _ := client.UpdateCategory(CategoryUpdateInput{
		Id:   graphql.ID("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz"),
		Name: "Emily",
	})
	// Assert
	autopilot.Equals(t, "Emily", result.Name)
}

func TestDeleteRubricCategory(t *testing.T) {
	// Arrange
	client := ANewClient(t, "rubric/category/delete")
	// Act
	err := client.DeleteCategory("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
