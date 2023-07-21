package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"

	"github.com/rocktavious/autopilot/v2022"
)

func TestAssignTagForAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,value},errors{message,path}}}",
    "variables": {
		"input": {
			"alias": "{{ template "alias1" }}",
			"tags": [
				{
				"key": "hello",
				"value": "world"
				}
			]
		}
}}`
	response := `{"data": {
		"tagAssign": {
			"tags": [
				{
				"id": "{{ template "id1" }}",
				"key": "hello",
				"value": "world"
				}
			],
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "tagAssignWithAlias", request, response)
	// Act
	result, err := client.AssignTags("example", map[string]string{"hello": "world"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestAssignTagForId(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,value},errors{message,path}}}",
    "variables": {
		"input": {
			"id": "{{ template "id1" }}",
			"tags": [
				{
				"key": "hello",
				"value": "world"
				}
			]
		}
}}`
	response := `{"data": {
		"tagAssign": {
			"tags": [
				{
				"id": "{{ template "id1" }}",
				"key": "hello",
				"value": "world"
				}
			],
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "tagAssignWithId", request, response)
	// Act
	result, err := client.AssignTags("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", map[string]string{"hello": "world"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestCreateTag(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation TagCreate($input:TagCreateInput!){tagCreate(input: $input){tag{id,key,value},errors{message,path}}}",
    "variables": {
		"input": {
			"id": "{{ template "id1" }}",
			"key": "hello",
			"value": "world"
		}
}}`
	response := `{"data": {
		"tagCreate": {
			"tag": 
			{
			"id": "{{ template "id1" }}",
			"key": "hello",
			"value": "world"
			},
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "tagCreate", request, response)
	// Act
	result, err := client.CreateTags("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", map[string]string{"hello": "world"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestUpdateTag(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation TagUpdate($input:TagUpdateInput!){tagUpdate(input: $input){tag{id,key,value},errors{message,path}}}",
    "variables": {
		"input": {
			"id": "{{ template "id1" }}",
			"key": "hello",
			"value": "world!"
		}
}}`
	response := `{"data": {
		"tagUpdate": {
			"tag": 
			{
			"id": "{{ template "id1" }}",
			"key": "hello",
			"value": "world!"
			},
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "tagUpdate", request, response)
	// Act
	result, err := client.UpdateTag(ol.TagUpdateInput{
		Id:    "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		Key:   "hello",
		Value: "world!",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "hello", result.Key)
	autopilot.Equals(t, "world!", result.Value)
}

func TestDeleteTag(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation TagDelete($input:TagDeleteInput!){tagDelete(input: $input){errors{message,path}}}",
    "variables": {
		"input": {
			"id": "{{ template "id1" }}"
		}
}}`
	response := `{"data": {
		"tagDelete": {
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "tagDelete", request, response)
	// Act
	err := client.DeleteTag("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Ok(t, err)
}
