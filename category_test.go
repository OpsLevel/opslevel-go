package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

func TestCreateRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/create")
	// Act
	result, _ := client.CreateCategory(ol.CategoryCreateInput{
		Name: "Kyle",
	})
	// Assert
	autopilot.Equals(t, "Kyle", result.Name)
}

func TestGetRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/get")
	// Act
	result, err := client.GetCategory("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Reliability", result.Name)
}

func TestGetMissingRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/get_missing")
	// Act
	_, err := client.GetCategory("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListRubricCategories(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"rubric": {
							"categories": {
								"nodes": [
									{
										{{ template "rubric_categories_response1" }}
									},
									{
										{{ template "rubric_categories_response2" }} 
									}
								],
								{{ template "pagination_initial_pageInfo_response" }},
								"totalCount": 2
						  }}}}}`},
		{`{"query": "query ($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"rubric": {
							"categories": {
								"nodes": [
									{
										{{ template "rubric_categories_response3" }}
									}
								],
								{{ template "pagination_second_pageInfo_response" }},
								"totalCount": 1
						  }}}}}`},
	}
	client := APaginatedTestClient(t, "rubric/category/list", requests...)
	// Act
	response, err := client.ListCategories(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "🟢 Reliability", result[1].Name)
	autopilot.Equals(t, "🔍 Observability", result[2].Name)

	//fmt.Println(Templated(requests[0].Request))
	//panic(true)
}

func TestUpdateRubricCategory(t *testing.T) {
	// Arrange
	client := ATestClient(t, "rubric/category/update")
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
	client := ATestClient(t, "rubric/category/delete")
	// Act
	err := client.DeleteCategory("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
