package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

var (
	noTags   = []ol.Tag{}
	tagOne   = ol.Tag{Id: id1, Key: "dev", Value: "true"}
	tagTwo   = ol.Tag{Id: id2, Key: "foo", Value: "bar"}
	tagThree = ol.Tag{Id: id3, Key: "prod", Value: "true"}
	tagFour  = ol.Tag{Id: id4, Key: "env", Value: "prod"}
	allTags  = []ol.Tag{tagOne, tagTwo, tagThree, tagFour}
)

type extractTagIdsTestCase struct {
	existingTags           []ol.Tag
	expectedTagIdsToDelete []ol.ID
	tagsWanted             []ol.Tag
}

func TestExtractTagIdsToDelete(t *testing.T) {
	var allTagIds, noTagIds []ol.ID
	for _, tag := range allTags {
		allTagIds = append(allTagIds, tag.Id)
	}

	testCases := map[string]extractTagIdsTestCase{
		"delete none": {
			existingTags:           allTags,
			expectedTagIdsToDelete: noTagIds,
			tagsWanted:             allTags,
		},
		"delete all": {
			existingTags:           allTags,
			expectedTagIdsToDelete: allTagIds,
			tagsWanted:             noTags,
		},
		"delete some": {
			tagsWanted:             []ol.Tag{tagOne, tagTwo},
			existingTags:           []ol.Tag{tagOne, tagTwo, tagThree, tagFour},
			expectedTagIdsToDelete: []ol.ID{tagThree.Id, tagFour.Id},
		},
		"no change": {
			existingTags:           allTags,
			expectedTagIdsToDelete: noTagIds,
			tagsWanted:             allTags,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Act
			tagIdsToDelete := ol.ExtractTagIdsToDelete(tc.existingTags, tc.tagsWanted)

			// Assert
			autopilot.Equals(t, tagIdsToDelete, tc.expectedTagIdsToDelete)
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
	testCases := map[string]tagHasSameKeyValueTestCase{
		"empty tags match": {
			tagOne:          ol.Tag{},
			tagTwo:          ol.Tag{},
			hasSameKeyValue: true,
		},
		"empty tag does not match non-empty tag": {
			tagOne:          tagOne,
			tagTwo:          ol.Tag{},
			hasSameKeyValue: false,
		},
		"tags have different key": {
			tagOne:          tagOne,
			tagTwo:          ol.Tag{Key: "env", Value: tagOne.Value},
			hasSameKeyValue: false,
		},
		"tags have different value": {
			tagOne:          tagOne,
			tagTwo:          ol.Tag{Key: tagOne.Key, Value: "prod"},
			hasSameKeyValue: false,
		},
		"tags have same key and value": {
			tagOne:          tagOne,
			tagTwo:          ol.Tag{Key: tagOne.Key, Value: tagOne.Value},
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

func TestReconcileTagsDeleteAllTags(t *testing.T) {
	// Arrange
	domain := &ol.Domain{DomainId: ol.DomainId{Id: id3}}
	testRequestTagsList := autopilot.NewTestRequest(
		`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "domain": { {{ template "id3" }} } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag1" }}, {{ template "tag2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestDeleteTagOne := autopilot.NewTestRequest(
		`mutation TagDelete($input:TagDeleteInput!){tagDelete(input: $input){errors{message,path}}}`,
		`{"input": { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODg2" }}`,
		`{"data": { "tagDelete": { "errors": [] }}}`,
	)
	testRequestDeleteTagTwo := autopilot.NewTestRequest(
		`mutation TagDelete($input:TagDeleteInput!){tagDelete(input: $input){errors{message,path}}}`,
		`{"input": { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODg3" }}`,
		`{"data": { "tagDelete": { "errors": [] }}}`,
	)

	requests := []autopilot.TestRequest{testRequestTagsList, testRequestDeleteTagOne, testRequestDeleteTagTwo}
	client := BestTestClient(t, "tags/reconcile_tags_delete_all", requests...)
	// Act
	err := client.ReconcileTags(domain, noTags)

	// Assert
	autopilot.Ok(t, err)
}

func TestReconcileTagsAssignAllTags(t *testing.T) {
	// Arrange
	domain := &ol.Domain{DomainId: ol.DomainId{Id: id3}}
	testRequestTagsList := autopilot.NewTestRequest(
		`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "domain": { {{ template "id3" }} } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestAssignTagOne := autopilot.NewTestRequest(
		`mutation TagAssign($input:TagAssignInput!){tagAssign(input: $input){tags{id,key,value},errors{message,path}}}`,
		`{"input": { "id": "{{ template "id3_string" }}", "tags": [
      { "key": "dev", "value": "true" },
      { "key": "foo", "value": "bar" },
      { "key": "prod", "value": "true" },
      { "key": "env", "value": "prod" }
     ] } }`,
		`{"data": {"tagAssign": { "tags": [
      { "key": "dev", "value": "true" },
      { "key": "foo", "value": "bar" },
      { "key": "prod", "value": "true" },
      { "key": "env", "value": "prod" }
     ], "errors": [] }}}`,
	)

	requests := []autopilot.TestRequest{testRequestTagsList, testRequestAssignTagOne}
	client := BestTestClient(t, "tags/reconcile_tags_create_all", requests...)
	// Act
	err := client.ReconcileTags(domain, allTags)

	// Assert
	autopilot.Ok(t, err)
}
