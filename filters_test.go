package opslevel_test

import (
	"fmt"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"

	"github.com/rocktavious/autopilot/v2022"
)

func TestCreateFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/create")
	// Act
	result, err := client.CreateFilter(ol.FilterCreateInput{
		Name:       "Kubernetes",
		Connective: ol.ConnectiveEnumAnd,
		Predicates: []ol.FilterPredicate{{
			Key:   ol.PredicateKeyEnumTierIndex,
			Type:  ol.PredicateTypeEnumEquals,
			Value: "1",
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Kubernetes", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumEquals, result.Predicates[0].Type)
}

func TestGetFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/get")
	// Act
	result, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Test", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestGetMissingFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/get_missing")
	// Act
	_, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListFilters(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ($after:String!$first:Int!){account{filters{nodes{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"filters": {
							"nodes": [
								{
									{{ template "filter_kubernetes_response" }}
								},
								{
									{{ template "filter_tier1service_response" }} 
								}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
						  }}}}`},
		{`{"query": "query ($after:String!$first:Int!){account{filters{nodes{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"filters": {
							"nodes": [
								{
									{{ template "filter_complex_kubernetes_response" }}
								}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }}}}`},
	}
	//client := APaginatedTestClient(t, "filter/list", requests...)
	//// Act
	//response, err := client.ListFilters(nil)
	//result := response.Nodes
	//// Assert
	//autopilot.Ok(t, err)
	//autopilot.Equals(t, 3, len(result))
	//autopilot.Equals(t, "Tier 1 Services", result[1].Name)
	//autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result[2].Predicates[0].Key)
	fmt.Println(Templated(requests[0].Request))
	panic(true)
}

func TestUpdateFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/update")
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:   "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg",
		Name: "Test Updated",
		Predicates: []ol.FilterPredicate{{
			Key:   ol.PredicateKeyEnumTierIndex,
			Type:  ol.PredicateTypeEnumEquals,
			Value: "1",
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Test Updated", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestDeleteFilter(t *testing.T) {
	// Arrange
	client := ATestClient(t, "filter/delete")
	// Act
	err := client.DeleteFilter("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
