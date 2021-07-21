package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
	"github.com/shurcooL/graphql"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/createRubricCategory", autopilot.FixtureResponse("rubric_category_create_response.json"), FixtureQueryValidation(t, "rubric_category_create_request.json"))
	client := NewClient("X", SetURL(url))
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
	url := autopilot.RegisterEndpoint("/listRubricCategories", autopilot.FixtureResponse("rubric_category_response.json"), FixtureQueryValidation(t, "rubric_category_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, _ := client.ListCategories()
	// Assert
	autopilot.Equals(t, 7, len(result))
	autopilot.Equals(t, "Reliability", result[1].Name)
}

func TestUpdateRubricCategory(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/updateRubricCategory", autopilot.FixtureResponse("rubric_category_update_response.json"), FixtureQueryValidation(t, "rubric_category_update_request.json"))
	client := NewClient("X", SetURL(url))
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
	url := autopilot.RegisterEndpoint("/deleteRubricCategory", autopilot.FixtureResponse("rubric_category_delete_response.json"), FixtureQueryValidation(t, "rubric_category_delete_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	err := client.DeleteCategory("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
