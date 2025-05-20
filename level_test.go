package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateRubricLevels(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`mutation LevelCreate($input:LevelCreateInput!){levelCreate(input: $input){level{alias,description,id,index,name},errors{message,path}}}`,
		`{"input": { "name": "Kyle", "description": "Created By Kyle", "index": 4 }}`,
		`{"data": { "levelCreate": { "level": { "alias": "kyle", "description": "Created By Kyle", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw", "index": 4, "name": "Kyle" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "rubric/level_create", testRequest)
	// Act
	index := 4
	result, _ := client.CreateLevel(ol.LevelCreateInput{
		Name:        "Kyle",
		Description: ol.RefOf("Created By Kyle"),
		Index:       &index,
	})
	// Assert
	autopilot.Equals(t, "kyle", result.Alias)
	autopilot.Equals(t, 4, result.Index)
}

func TestGetRubricLevel(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query LevelGet($id:ID!){account{level(id: $id){alias,description,id,index,name}}}`,
		`{"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"}`,
		`{"data": {
        "account": {
          "level": {
            "alias": "bronze",
            "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
            "index": 1,
            "name": "Bronze"
          }}}}`,
	)

	client := BestTestClient(t, "rubric/level_get", testRequest)
	// Act
	result, err := client.GetLevel("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Bronze", result.Name)
}

func TestGetMissingRubricLevel(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query LevelGet($id:ID!){account{level(id: $id){alias,description,id,index,name}}}`,
		`{"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"}`,
		`{"data": { "account": { "level": null }}}`,
	)

	client := BestTestClient(t, "rubric/level_get_missing", testRequest)
	// Act
	_, err := client.GetLevel("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListRubricLevels(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query LevelsList($after:String!$first:Int!){account{rubric{levels(after: $after, first: $first){nodes{alias,description,id,index,name},{{ template "pagination_request" }}}}}}`,
		`{"after":"","first":100}`,
		`{
    "data": {
      "account": {
        "rubric": {
          "levels": {
            "nodes": [
              {
                "alias": "beginner",
                "description": "Services in this level are below the minimum standard to ship to production. You should address your failing checks as soon as possible.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMTAx",
                "index": 0,
                "name": "Beginner"
              },
              {
                "alias": "bronze",
                "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
                "index": 1,
                "name": "Bronze"
              },
              {
                "alias": "silver",
                "description": "Services in this level satisfy important and critical checks. This is considered healthy.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE4",
                "index": 2,
                "name": "Silver"
              },
              {
                "alias": "gold",
                "description": "Services in this level satisfy critical, important and useful checks. This is the requirement for your highest tier services but all services should aspire to be in this level.",
                "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE5",
                "index": 3,
                "name": "Gold"
              }
            ]
          }
        }
      }
    }
  }`,
	)
	client := BestTestClient(t, "rubric/level/list", testRequest)
	// Act
	result, err := client.ListLevels(nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, result.TotalCount)
	autopilot.Equals(t, "Bronze", result.Nodes[1].Name)
}

func TestUpdateRubricLevel(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "name": "{{ template "name1" }}", "description": "{{ template "description" }}" }}`,
		`{"data": { "levelUpdate": { "level": {{ template "level_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "rubric/level_update", testRequest)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          id1,
		Name:        ol.RefOf("Example"),
		Description: ol.RefOf("An example description"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelNoName(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "description": "{{ template "description" }}" } }`,
		`{"data": { "levelUpdate": { "level": {{ template "level_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "rubric/level_update_noname", testRequest)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          id1,
		Description: ol.RefOf("An example description"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelEmptyDescription(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "name": "{{ template "name1" }}", "description": "" }}`,
		`{"data": { "levelUpdate": { "level": {{ template "level_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "rubric/level_update_emptydescription", testRequest)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          id1,
		Name:        ol.RefOf("Example"),
		Description: ol.RefOf(""),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelNoDescription(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "name": "{{ template "name1" }}" }}`,
		`{"data": { "levelUpdate": { "level": {{ template "level_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "rubric/level_update_nodescription", testRequest)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:   id1,
		Name: ol.RefOf("Example"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestDeleteRubricLevels(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation LevelDelete($input:LevelDeleteInput!){levelDelete(input: $input){deletedLevelId,errors{message,path}}}`,
		`{"input": { "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw" }}`,
		`{"data": { "levelDelete": { "deletedLevelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw", "errors": [] }}}`,
	)

	client := BestTestClient(t, "rubric/level_delete", testRequest)
	// Act
	err := client.DeleteLevel("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw")
	// Assert
	autopilot.Equals(t, nil, err)
}
