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
	tagFive  = ol.Tag{Id: id4, Key: "foo", Value: "baz"}
)

type reconcileTagsTestCase struct {
	current  []ol.Tag
	desired  []ol.Tag
	toCreate []ol.Tag
	toDelete []ol.Tag
}

func TestReconcileTags(t *testing.T) {
	// Arrange
	testCases := map[string]reconcileTagsTestCase{
		"create all": {
			current:  noTags,
			desired:  []ol.Tag{tagOne, tagTwo},
			toCreate: []ol.Tag{tagOne, tagTwo},
			toDelete: noTags,
		},
		"delete all": {
			current:  []ol.Tag{tagOne, tagTwo},
			desired:  noTags,
			toCreate: noTags,
			toDelete: []ol.Tag{tagOne, tagTwo},
		},
		"create one delete one": {
			current:  []ol.Tag{tagOne, tagThree},
			desired:  []ol.Tag{tagOne, tagTwo},
			toCreate: []ol.Tag{tagTwo},
			toDelete: []ol.Tag{tagThree},
		},
		"create two delete two": {
			current:  []ol.Tag{tagOne, tagTwo, tagThree},
			desired:  []ol.Tag{tagOne, tagFour, tagFive},
			toCreate: []ol.Tag{tagFour, tagFive},
			toDelete: []ol.Tag{tagTwo, tagThree},
		},
		"replace": {
			current:  []ol.Tag{tagOne, tagTwo},
			desired:  []ol.Tag{tagThree, tagFour},
			toCreate: []ol.Tag{tagThree, tagFour},
			toDelete: []ol.Tag{tagOne, tagTwo},
		},
		"no create no delete": {
			current:  []ol.Tag{tagOne, tagTwo},
			desired:  []ol.Tag{tagOne, tagTwo},
			toCreate: noTags,
			toDelete: noTags,
		},
		"null": {
			current:  noTags,
			desired:  noTags,
			toCreate: noTags,
			toDelete: noTags,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Act
			toCreate, toDelete := ol.TestReconcileTags(tc.current, tc.desired)
			// Assert
			autopilot.Equals(t, tc.toCreate, toCreate)
			autopilot.Equals(t, tc.toDelete, toDelete)
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

func TestReconcileTagsAPI(t *testing.T) {
	// Arrange
	domain := &ol.Domain{DomainId: ol.DomainId{Id: id3}}
	requests := []autopilot.TestRequest{
		autopilot.NewTestRequest(
			`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}`,
			`{ {{ template "first_page_variables" }}, "domain": { {{ template "id3" }} } }`,
			`{ "data": { "account": { "domain": { "tags": { "nodes": [{ {{ template "id1" }}, "key": "hello", "value": "world" }], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
		),
		autopilot.NewTestRequest(
			`mutation TagCreate($input:TagCreateInput!){tagCreate(input: $input){tag{id,key,value},errors{message,path}}}`,
			`{"input": { {{ template "id3" }}, "key": "dev", "type": "Domain", "value": "true" }}`,
			`{"data": { "tagCreate": { "tag": { {{ template "id2" }}, "key": "dev", "value": "true" }, "errors": [] }}}`,
		),
		autopilot.NewTestRequest(
			`mutation TagCreate($input:TagCreateInput!){tagCreate(input: $input){tag{id,key,value},errors{message,path}}}`,
			`{"input": { {{ template "id3" }}, "key": "foo", "type": "Domain", "value": "bar" }}`,
			`{"data": { "tagCreate": { "tag": { {{ template "id3" }}, "key": "foo", "value": "bar" }, "errors": [] }}}`,
		),
		autopilot.NewTestRequest(
			`mutation TagDelete($input:TagDeleteInput!){tagDelete(input: $input){errors{message,path}}}`,
			`{"input": { {{ template "id1" }} }}`,
			`{"data": { "tagDelete": { "errors": [] }}}`,
		),
	}
	client := BestTestClient(t, "tags/reconcile_tags", requests...)
	// Act
	err := client.ReconcileTags(domain, []ol.Tag{tagOne, tagTwo})

	// Assert
	autopilot.Ok(t, err)
}
