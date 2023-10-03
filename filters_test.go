package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"

	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateFilter(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation FilterCreate($input:FilterCreateInput!){filterCreate(input: $input){filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},errors{message,path}}}",
	"variables":{
		"input": {
			"name": "Kubernetes",
			"predicates": [
				{
				"key": "tier_index",
				"type": "equals",
				"value": "1"
				}
			],
			"connective": "and"
		}
    }
}`
	response := `{"data": {
	"filterCreate": {
		"filter": {
			{{ template "filter_tier1service_response" }}
		},
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "filter/create", request, response)
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
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumEquals, result.Predicates[0].Type)
}

func TestCreateFilterNested(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation FilterCreate($input:FilterCreateInput!){filterCreate(input: $input){filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},errors{message,path}}}",
	"variables":{
		"input": {
			{{ template "create_filter_nested_input" }}
		}
    }
}`
	response := `{"data": {
	"filterCreate": {
		"filter": {
			{{ template "create_filter_nested_response" }}
		},
		"errors": []
	}
}}`

	client := ABetterTestClient(t, "filter/create_nested", request, response)
	// Act
	result, err := client.CreateFilter(ol.FilterCreateInput{
		Name:       "Self deployed or Rails",
		Connective: ol.ConnectiveEnumOr,
		Predicates: []ol.FilterPredicate{
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumMatches,
				Value: "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg",
			},
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumMatches,
				Value: "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjQ",
			},
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Self deployed or Rails", result.Name)
	autopilot.Equals(t, ol.ConnectiveEnumOr, result.Connective)
	autopilot.Equals(t, ol.PredicateKeyEnumFilterID, result.Predicates[0].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumMatches, result.Predicates[0].Type)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg", result.Predicates[0].Value)
	autopilot.Equals(t, ol.PredicateKeyEnumFilterID, result.Predicates[1].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumMatches, result.Predicates[1].Type)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjQ", result.Predicates[1].Value)
}

func TestGetFilter(t *testing.T) {
	// Arrange
	request := `{
    "query": "query FilterGet($id:ID!){account{filter(id: $id){connective,htmlUrl,id,name,predicates{key,keyData,type,value}}}}",
	"variables":{
		"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"
    }
}`
	response := `{"data": {
	"account": {
		"filter": {
			{{ template "filter_tier1service_response" }}
		}
	}
}}`
	client := ABetterTestClient(t, "filter/get", request, response)
	// Act
	result, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestGetMissingFilter(t *testing.T) {
	// Arrange
	request := `{
    "query": "query FilterGet($id:ID!){account{filter(id: $id){connective,htmlUrl,id,name,predicates{key,keyData,type,value}}}}",
	"variables":{
		"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMf"
    }
}`
	response := `{"data": {
	"account": {
		"filter": null
	}
}}`
	client := ABetterTestClient(t, "filter/get_missing", request, response)
	// Act
	_, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListFilters(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			Request: `{"query": "query FilterList($after:String!$first:Int!){account{filters(after: $after, first: $first){nodes{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			Response: `{
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
						  }}}}`,
		},
		{
			Request: `{"query": "query FilterList($after:String!$first:Int!){account{filters(after: $after, first: $first){nodes{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			Response: `{
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
						  }}}}`,
		},
	}
	client := APaginatedTestClient(t, "filter/list", requests...)
	// Act
	response, err := client.ListFilters(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Tier 1 Services", result[1].Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result[2].Predicates[0].Key)
	// fmt.Println(Templated(requests[0].Request))
	// panic(true)
}

func TestUpdateFilter(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation FilterUpdate($input:FilterUpdateInput!){filterUpdate(input: $input){filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},errors{message,path}}}",
	"variables":{
		"input": {
			"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg",
			"name": "Test Updated",
			"predicates": [
				{
				"key": "tier_index",
				"type": "equals",
				"value": "1"
				}
			]
		}
    }
}`
	response := `{"data": {
	"filterUpdate": {
		"filter": {
			{{ template "filter_tier1service_response" }}
		},
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "filter/update", request, response)
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
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestUpdateFilterNested(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation FilterUpdate($input:FilterUpdateInput!){filterUpdate(input: $input){filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},errors{message,path}}}",
	"variables":{
		"input": {
			{{ template "update_filter_nested_input" }}
		}
    }
}`
	response := `{"data": {
	"filterUpdate": {
		"filter": {
			{{ template "update_filter_nested_response" }}
		},
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "filter/update_nested", request, response)
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:         "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzIzNDY",
		Name:       "Tier 1-2 not deployed by us",
		Connective: ol.ConnectiveEnumAnd,
		Predicates: []ol.FilterPredicate{
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumDoesNotMatch,
				Value: "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg",
			},
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumMatches,
				Value: "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjY",
			},
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Tier 1-2 not deployed by us", result.Name)
	autopilot.Equals(t, ol.ConnectiveEnumAnd, result.Connective)
	autopilot.Equals(t, ol.PredicateKeyEnumFilterID, result.Predicates[0].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumDoesNotMatch, result.Predicates[0].Type)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg", result.Predicates[0].Value)
	autopilot.Equals(t, ol.PredicateKeyEnumFilterID, result.Predicates[1].Key)
	autopilot.Equals(t, ol.PredicateTypeEnumMatches, result.Predicates[1].Type)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjY", result.Predicates[1].Value)
}

func TestDeleteFilter(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation FilterDelete($input:DeleteInput!){filterDelete(input: $input){deletedId,errors{message,path}}}",
	"variables":{
		"input": {
			"id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz"
		}
    }
}`
	response := `{"data": {
	"filterDelete": {
		"deletedId": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyNQ",
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "filter/delete", request, response)
	// Act
	err := client.DeleteFilter("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
