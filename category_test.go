package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation CategoryCreate($input:CategoryCreateInput!){categoryCreate(input: $input){category{id,name},errors{message,path}}}"`,
		`{ "input": { "name": "Kyle" }}`,
		`{"data": { "categoryCreate": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz", "name": "Kyle" }, "errors": [] } }}`,
	)
	client := BestTestClient(t, "rubric/category_create", testRequest)
	// Act
	result, _ := client.CreateCategory(ol.CategoryCreateInput{
		Name: "Kyle",
	})
	// Assert
	autopilot.Equals(t, "Kyle", result.Name)
}

func TestGetRubricCategory(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query CategoryGet($id:ID!){account{category(id: $id){id,name}}}"`,
		`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg" }`,
		`{"data": { "account": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA0", "name": "Reliability" } }}}`,
	)
	client := BestTestClient(t, "rubric/category_get", testRequest)
	// Act
	result, err := client.GetCategory("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Reliability", result.Name)
}

func TestGetMissingRubricCategory(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query CategoryGet($id:ID!){account{category(id: $id){id,name}}}"`,
		`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg" }`,
		`{"data": { "account": { "category": null }}}`,
	)
	client := BestTestClient(t, "rubric/category_get_missing", testRequest)
	// Act
	_, err := client.GetCategory("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListRubricCategories(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query CategoryList($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{id,name},{{ template "pagination_request" }},totalCount}}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "categories": { "nodes": [ { {{ template "rubric_categories_response1" }} }, { {{ template "rubric_categories_response2" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query CategoryList($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{id,name},{{ template "pagination_request" }},totalCount}}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "categories": { "nodes": [ { {{ template "rubric_categories_response3" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "rubric/category_list", requests...)
	// Act
	response, err := client.ListCategories(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "🟢 Reliability", result[1].Name)
	autopilot.Equals(t, "🔍 Observability", result[2].Name)

	// fmt.Println(Templated(requests[0].Request))
	// panic(true)
}

func TestUpdateRubricCategory(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation CategoryUpdate($input:CategoryUpdateInput!){categoryUpdate(input: $input){category{id,name},errors{message,path}}}"`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz", "name": "Emily" }}`,
		`{"data": { "categoryUpdate": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz", "name": "Emily" }, "errors": [] }}}`,
	)
	client := BestTestClient(t, "rubric/category_update", testRequest)
	// Act
	result, _ := client.UpdateCategory(ol.CategoryUpdateInput{
		Id:   "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz",
		Name: "Emily",
	})
	// Assert
	autopilot.Equals(t, "Emily", result.Name)
}

func TestDeleteRubricCategory(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation CategoryDelete($input:CategoryDeleteInput!){categoryDelete(input: $input){deletedCategoryId,errors{message,path}}}"`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz" }}`,
		`{"data": { "categoryDelete": { "deletedCategoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz", "errors": [] }}}`,
	)
	client := BestTestClient(t, "rubric/category_delete", testRequest)
	// Act
	err := client.DeleteCategory("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
