package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"

	"github.com/rocktavious/autopilot/v2023"
)

func TestAssignTagForAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,owner{... on Team{alias,id}},value},errors{message,path}}}`,
		`{"input": { "alias": "{{ template "alias1" }}", "tags": [ { "key": "hello", "value": "world" } ] } }`,
		`{"data": {"tagAssign": { "tags": [ { {{ template "id1" }}, "key": "hello", "value": "world" } ], "errors": [] }}}`,
	)

	client := BestTestClient(t, "tagAssignWithAlias", testRequest)
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
	testRequest := autopilot.NewTestRequest(
		`mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,owner{... on Team{alias,id}},value},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "tags": [ { "key": "hello", "value": "world" } ] }}`,
		`{"data": { "tagAssign": { "tags": [ { {{ template "id1" }}, "key": "hello", "value": "world" } ], "errors": [] }}}`,
	)

	client := BestTestClient(t, "tagAssignWithId", testRequest)
	// Act
	result, err := client.AssignTags(string(id1), map[string]string{"hello": "world"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestCreateTag(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TagCreate($input:TagCreateInput!){tagCreate(input: $input){tag{id,key,owner{... on Team{alias,id}},value},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "key": "hello", "value": "world" }}`,
		`{"data": { "tagCreate": { "tag": { {{ template "id1" }}, "key": "hello", "value": "world" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "tagCreate", testRequest)
	// Act
	result, err := client.CreateTags(string(id1), map[string]string{"hello": "world"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
	autopilot.Equals(t, "hello", result[0].Key)
	autopilot.Equals(t, "world", result[0].Value)
}

func TestUpdateTag(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TagUpdate($input:TagUpdateInput!){tagUpdate(input: $input){tag{id,key,owner{... on Team{alias,id}},value},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "key": "hello", "value": "world!" }}`,
		`{"data": { "tagUpdate": { "tag": { {{ template "id1" }}, "key": "hello", "value": "world!" }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "tagUpdate", testRequest)
	// Act
	result, err := client.UpdateTag(ol.TagUpdateInput{
		Id:    id1,
		Key:   ol.RefOf("hello"),
		Value: ol.RefOf("world!"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "hello", result.Key)
	autopilot.Equals(t, "world!", result.Value)
}

func TestDeleteTag(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TagDelete($input:TagDeleteInput!){tagDelete(input: $input){errors{message,path}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": { "tagDelete": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "tagDelete", testRequest)
	// Act
	err := client.DeleteTag(id1)
	// Assert
	autopilot.Ok(t, err)
}
