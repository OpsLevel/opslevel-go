package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

func TestCreateRubricLevels(t *testing.T) {
	// Arrange
	request := `{
	   "query": "mutation LevelCreate($input:LevelCreateInput!){levelCreate(input: $input){level{alias,description,id,index,name},errors{message,path}}}",
		"variables":{
			"input": {
				"name": "Kyle",
				"description": "Created By Kyle",
				"index": 4
			}
	   }
	}`
	response := `{"data": {
	"levelCreate": {
		"level": {
			"alias": "kyle",
			"description": "Created By Kyle",
			"id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw",
			"index": 4,
			"name": "Kyle"
		},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_create", request, response)
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
	request := `{
	   "query": "query LevelGet($id:ID!){account{level(id: $id){alias,description,id,index,name}}}",
		"variables":{
			"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"
	   }
	}`
	response := `{"data": {
	"account": {
		"level": {
			"alias": "bronze",
			"description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
			"id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
			"index": 1,
			"name": "Bronze"
		}
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_get", request, response)
	// Act
	result, err := client.GetLevel("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Bronze", result.Name)
}

func TestGetMissingRubricLevel(t *testing.T) {
	// Arrange
	request := `{
	   "query": "query LevelGet($id:ID!){account{level(id: $id){alias,description,id,index,name}}}",
		"variables":{
			"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"
	   }
	}`
	response := `{"data": {
	"account": {
		"level": null
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_get_missing", request, response)
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
	request := `{
	   "query": "mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}",
		"variables":{
			"input": {
				"id": "{{ template "id1" }}",
				"name": "{{ template "name1" }}",
				"description": "{{ template "description" }}"
			}
	   }
	}`
	response := `{"data": {
	"levelUpdate": {
		"level": {{ template "level_1" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_update", request, response)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
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
	request := `{
	   "query": "mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}",
		"variables":{
			"input": {
				"id": "{{ template "id1" }}",
				"description": "{{ template "description" }}"
			}
	   }
	}`
	response := `{"data": {
	"levelUpdate": {
		"level": {{ template "level_1" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_update_noname", request, response)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		Description: ol.NewString("An example description"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelEmptyDescription(t *testing.T) {
	// Arrange
	request := `{
	   "query": "mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}",
		"variables":{
			"input": {
				"id": "{{ template "id1" }}",
				"name": "{{ template "name1" }}",
				"description": ""
			}
	   }
	}`
	response := `{"data": {
	"levelUpdate": {
		"level": {{ template "level_1" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_update_emptydescription", request, response)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:          "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		Name:        "Example",
		Description: ol.NewString(""),
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestUpdateRubricLevelNoDescription(t *testing.T) {
	// Arrange
	request := `{
	   "query": "mutation LevelUpdate($input:LevelUpdateInput!){levelUpdate(input: $input){level{alias,description,id,index,name},errors{message,path}}}",
		"variables":{
			"input": {
				"id": "{{ template "id1" }}",
				"name": "{{ template "name1" }}"
			}
	   }
	}`
	response := `{"data": {
	"levelUpdate": {
		"level": {{ template "level_1" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_update_nodescription", request, response)
	// Act
	result, _ := client.UpdateLevel(ol.LevelUpdateInput{
		Id:   "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		Name: "Example",
	})
	// Assert
	autopilot.Equals(t, "example", result.Alias)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "An example description", result.Description)
}

func TestDeleteRubricLevels(t *testing.T) {
	// Arrange
	request := `{
	   "query": "mutation LevelDelete($input:LevelDeleteInput!){levelDelete(input: $input){deletedLevelId,errors{message,path}}}",
		"variables":{
			"input": {
				"id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw"
			}
	   }
	}`
	response := `{"data": {
	"levelDelete": {
		"deletedLevelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw",
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "rubric/level_delete", request, response)
	// Act
	err := client.DeleteLevel("Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDgw")
	// Assert
	autopilot.Equals(t, nil, err)
}
