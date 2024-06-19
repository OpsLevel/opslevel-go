package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"

	"github.com/rocktavious/autopilot/v2023"
)

type extractTagsTestCase struct {
	tagsWanted           []ol.Tag
	existingTags         []ol.Tag
	expectedTagsToCreate []ol.Tag
	expectedTagsToDelete []ol.Tag
}

func TestExtractTags(t *testing.T) {
	var noTags []ol.Tag
	fourTags := []ol.Tag{
		{Key: "foo", Value: "bar"},
		{Key: "ping", Value: "pong"},
		{Key: "marco", Value: "pollo"},
		{Key: "env", Value: "prod"},
	}
	// Arrange
	testCases := map[string]extractTagsTestCase{
		"create all delete none": {
			tagsWanted:           fourTags,
			existingTags:         noTags,
			expectedTagsToCreate: fourTags,
			expectedTagsToDelete: noTags,
		},
		"create none delete all": {
			tagsWanted:           noTags,
			existingTags:         fourTags,
			expectedTagsToCreate: noTags,
			expectedTagsToDelete: fourTags,
		},
		"create some delete some": {
			tagsWanted:           fourTags[:3],
			existingTags:         fourTags[1:],
			expectedTagsToCreate: fourTags[:1],
			expectedTagsToDelete: fourTags[3:],
		},
		"no change": {
			tagsWanted:           fourTags,
			existingTags:         fourTags,
			expectedTagsToCreate: noTags,
			expectedTagsToDelete: noTags,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Act
			aliasesToCreate, aliasesToDelete := ol.ExtractTags(tc.existingTags, tc.tagsWanted)

			// Assert
			autopilot.Equals(t, aliasesToCreate, tc.expectedTagsToCreate)
			autopilot.Equals(t, aliasesToDelete, tc.expectedTagsToDelete)
		})
	}
}

type tagHasSameKeyValueTestCase struct {
	tagOne          ol.Tag
	tagTwo          ol.Tag
	hasSameKeyValue bool
}

func TestAssignTagHasSameKeyValue(t *testing.T) {
	// Arrange
	testTag := ol.Tag{Key: "foo", Value: "bar"}
	testCases := map[string]tagHasSameKeyValueTestCase{
		"empty tags match": {
			tagOne:          ol.Tag{},
			tagTwo:          ol.Tag{},
			hasSameKeyValue: true,
		},
		"empty tag does not match non-empty tag": {
			tagOne:          testTag,
			tagTwo:          ol.Tag{},
			hasSameKeyValue: false,
		},
		"tags have different key": {
			tagOne:          testTag,
			tagTwo:          ol.Tag{Key: "env", Value: testTag.Value},
			hasSameKeyValue: false,
		},
		"tags have different value": {
			tagOne:          testTag,
			tagTwo:          ol.Tag{Key: testTag.Key, Value: "prod"},
			hasSameKeyValue: false,
		},
		"tags have same key and value": {
			tagOne:          testTag,
			tagTwo:          ol.Tag{Key: testTag.Key, Value: testTag.Value},
			hasSameKeyValue: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Assert
			autopilot.Equals(t, tc.tagOne.HasSameKeyValue(tc.tagTwo), tc.hasSameKeyValue)
			autopilot.Equals(t, tc.tagTwo.HasSameKeyValue(tc.tagOne), tc.hasSameKeyValue)
		})
	}
}

func TestAssignTagForAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,value},errors{message,path}}}`,
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
		`mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,value},errors{message,path}}}`,
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
		`mutation TagCreate($input:TagCreateInput!){tagCreate(input: $input){tag{id,key,value},errors{message,path}}}`,
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
		`mutation TagUpdate($input:TagUpdateInput!){tagUpdate(input: $input){tag{id,key,value},errors{message,path}}}`,
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
