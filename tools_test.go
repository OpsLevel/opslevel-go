package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateTool(t *testing.T) {
	// Arrange
	toolCreateInput := autopilot.Register[ol.ToolCreateInput]("tool_create_input", ol.NewExampleOf[ol.ToolCreateInput]())
	testRequest := autopilot.NewTestRequest(
		`mutation ToolCreate($input:ToolCreateInput!){toolCreate(input: $input){tool{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},errors{message,path}}}`,
		`{ "input": {{ template "tool_create_input" }}}`,
		`{"data": { "toolCreate": { "tool": {{ template "tool_1" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "toolCreate", testRequest)
	// Act
	result, err := client.CreateTool(toolCreateInput)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Service.Id)
	autopilot.Equals(t, ol.ToolCategoryOther, result.Category)
	autopilot.Equals(t, "Example", result.DisplayName)
	autopilot.Equals(t, "https://example.com", result.Url)
}

func TestUpdateTool(t *testing.T) {
	// Arrange
	toolUpdateInput := autopilot.Register[ol.ToolUpdateInput]("tool_update_input", ol.NewExampleOf[ol.ToolUpdateInput]())
	testRequest := autopilot.NewTestRequest(
		`mutation ToolUpdate($input:ToolUpdateInput!){toolUpdate(input: $input){tool{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},errors{message,path}}}`,
		`{ "input": {{ template "tool_update_input" }} }`,
		`{"data": { "toolUpdate": { "tool": {{ template "tool_1_update" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "toolUpdate", testRequest)
	// Act
	result, err := client.UpdateTool(toolUpdateInput)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ToolCategoryDeployment, result.Category)
	autopilot.Equals(t, "prod", result.Environment)
}

func TestDeleteTool(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ToolDelete($input:ToolDeleteInput!){toolDelete(input: $input){errors{message,path}}}`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "toolDelete": { "errors": [] }}}`,
	)
	client := BestTestClient(t, "toolDelete", testRequest)
	// Act
	err := client.DeleteTool(string(id1))
	// Assert
	autopilot.Ok(t, err)
}
