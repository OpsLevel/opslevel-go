package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation CategoryCreate($input:CategoryCreateInput!){categoryCreate(input: $input){category{description,id,name},errors{message,path}}}`,
		`{ "input": { "name": "Kyle" }}`,
		`{"data": { "categoryCreate": { "category": { {{ template "id1" }}, "name": "Kyle" }, "errors": [] } }}`,
	)
	client := BestTestClient(t, "rubric/category_create", testRequest)
	// Act
	result, _ := client.CreateCategory(ol.CategoryCreateInput{
		Name: "Kyle",
	})
	// Assert
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "Kyle", result.Name)
}

func TestGetRubricCategory(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query CategoryGet($id:ID!){account{category(id: $id){description,id,name}}}`,
		`{ {{ template "id2" }} }`,
		`{"data": { "account": { "category": { {{ template "id3" }}, "name": "Reliability" } }}}`,
	)
	client := BestTestClient(t, "rubric/category_get", testRequest)
	// Act
	result, err := client.GetCategory(id2)
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, id3, result.Id)
	autopilot.Equals(t, "Reliability", result.Name)
}

func TestGetMissingRubricCategory(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query CategoryGet($id:ID!){account{category(id: $id){description,id,name}}}`,
		`{ {{ template "id1" }} }`,
		`{"data": { "account": { "category": null }}}`,
	)
	client := BestTestClient(t, "rubric/category_get_missing", testRequest)
	// Act
	_, err := client.GetCategory(id1)
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListRubricCategories(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query CategoryList($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{description,id,name},{{ template "pagination_request" }},totalCount}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "categories": { "nodes": [ { {{ template "rubric_categories_response1" }} }, { {{ template "rubric_categories_response2" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query CategoryList($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{description,id,name},{{ template "pagination_request" }},totalCount}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "categories": { "nodes": [ { {{ template "rubric_categories_response3" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

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
	testRequest := autopilot.NewTestRequest(
		`mutation CategoryUpdate($input:CategoryUpdateInput!){categoryUpdate(input: $input){category{description,id,name},errors{message,path}}}`,
		`{ "input": { {{ template "id4" }}, "name": "Emily" }}`,
		`{"data": { "categoryUpdate": { "category": { {{ template "id4" }}, "name": "Emily" }, "errors": [] }}}`,
	)
	client := BestTestClient(t, "rubric/category_update", testRequest)
	// Act
	result, err := client.UpdateCategory(ol.CategoryUpdateInput{
		Id:   id4,
		Name: ol.RefOf("Emily"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id4, result.Id)
	autopilot.Equals(t, "Emily", result.Name)
}

func TestDeleteRubricCategory(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation CategoryDelete($input:CategoryDeleteInput!){categoryDelete(input: $input){deletedCategoryId,errors{message,path}}}`,
		`{ "input": { {{ template "id2" }} }}`,
		`{"data": { "categoryDelete": { "deletedCategoryId": "{{ template "id2_string" }}", "errors": [] }}}`,
	)
	client := BestTestClient(t, "rubric/category_delete", testRequest)
	// Act
	err := client.DeleteCategory(id2)
	// Assert
	autopilot.Equals(t, nil, err)
}
