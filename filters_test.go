package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"

	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateFilter(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation FilterCreate($input:FilterCreateInput!){filterCreate(input: $input){filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},errors{message,path}}}`,
		`{"input": {"name": "Kubernetes", "predicates": [ { "key": "tier_index", "type": "equals", "value": "1" } ], "connective": "and" }}`,
		`{"data": {"filterCreate": {"filter": { {{ template "filter_tier1service_response" }} }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "filter/create", testRequest)
	// Act
	result, err := client.CreateFilter(ol.FilterCreateInput{
		Name:       "Kubernetes",
		Connective: ol.RefOf(ol.ConnectiveEnumAnd),
		Predicates: &[]ol.FilterPredicateInput{{
			Key:   ol.PredicateKeyEnumTierIndex,
			Type:  ol.PredicateTypeEnumEquals,
			Value: ol.RefOf("1"),
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
	testRequest := autopilot.NewTestRequest(
		`mutation FilterCreate($input:FilterCreateInput!){filterCreate(input: $input){filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},errors{message,path}}}`,
		`{"input": { {{ template "create_filter_nested_input" }} }}`,
		`{"data": {"filterCreate": {"filter": { {{ template "create_filter_nested_response" }} }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "filter/create_nested", testRequest)
	// Act
	result, err := client.CreateFilter(ol.FilterCreateInput{
		Name:       "Self deployed or Rails",
		Connective: ol.RefOf(ol.ConnectiveEnumOr),
		Predicates: &[]ol.FilterPredicateInput{
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumMatches,
				Value: ol.RefOf("Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg"),
			},
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumMatches,
				Value: ol.RefOf("Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjQ"),
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
	testRequest := autopilot.NewTestRequest(
		`query FilterGet($id:ID!){account{filter(id: $id){id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}}}}`,
		`{"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg"}`,
		`{"data": {"account": {"filter": { {{ template "filter_tier1service_response" }} }}}}`,
	)

	client := BestTestClient(t, "filter/get", testRequest)
	// Act
	result, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestGetMissingFilter(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query FilterGet($id:ID!){account{filter(id: $id){id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}}}}`,
		`{"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMf"}`,
		`{"data": {"account": {"filter": null }}}`,
	)

	client := BestTestClient(t, "filter/get_missing", testRequest)
	// Act
	_, err := client.GetFilter("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListFilters(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query FilterList($after:String!$first:Int!){account{filters(after: $after, first: $first){nodes{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},{{ template "pagination_request" }},totalCount}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{"data": { "account": { "filters": { "nodes": [ { {{ template "filter_kubernetes_response" }} }, { {{ template "filter_tier1service_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query FilterList($after:String!$first:Int!){account{filters(after: $after, first: $first){nodes{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},{{ template "pagination_request" }},totalCount}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{"data": { "account": { "filters": { "nodes": [ { {{ template "filter_complex_kubernetes_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "filter/list", requests...)
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
	testRequest := autopilot.NewTestRequest(
		`mutation FilterUpdate($input:FilterUpdateInput!){filterUpdate(input: $input){filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},errors{message,path}}}`,
		`{"input": {"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg", "name": "Test Updated", "predicates": [ { "key": "tier_index", "type": "equals", "value": "1" } ] }}`,
		`{"data": {"filterUpdate": {"filter": { {{ template "filter_tier1service_response" }} }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "filter/update", testRequest)
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:   "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg",
		Name: ol.RefOf("Test Updated"),
		Predicates: &[]ol.FilterPredicateInput{{
			Key:   ol.PredicateKeyEnumTierIndex,
			Type:  ol.PredicateTypeEnumEquals,
			Value: ol.RefOf("1"),
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
}

func TestUpdateFilterNested(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation FilterUpdate($input:FilterUpdateInput!){filterUpdate(input: $input){filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},errors{message,path}}}`,
		`{"input": { {{ template "update_filter_nested_input" }} }}`,
		`{"data": {"filterUpdate": {"filter": { {{ template "update_filter_nested_response" }} }, "errors": [] }}}`,
	)

	client := BestTestClient(t, "filter/update_nested", testRequest)
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:         "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzIzNDY",
		Name:       ol.RefOf("Tier 1-2 not deployed by us"),
		Connective: ol.RefOf(ol.ConnectiveEnumAnd),
		Predicates: &[]ol.FilterPredicateInput{
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumDoesNotMatch,
				Value: ol.RefOf("Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg"),
			},
			{
				Key:   ol.PredicateKeyEnumFilterID,
				Type:  ol.PredicateTypeEnumMatches,
				Value: ol.RefOf("Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjY"),
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

func TestUpdateFilterCaseSensitiveTrue(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation FilterUpdate($input:FilterUpdateInput!){filterUpdate(input: $input){filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},errors{message,path}}}`,
		`{"input": {"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg", "name": "Test Updated", "predicates": [ { "key": "tier_index", "type": "equals", "value": "1", "caseSensitive": true } ] }}`,
		`{"data": {
      "filterUpdate": {
        "filter": {
          "connective": null,
          "htmlUrl": "https://app.opslevel.com/filters/401",
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzQwMQ",
          "name": "Tier 1 Services",
          "predicates": [
          {
            "key": "tier_index",
            "keyData": null,
            "type": "equals",
            "value": "1",
            "caseSensitive": true
          }
          ]
        },
        "errors": []
      }}}`,
	)

	client := BestTestClient(t, "filter/update_case_sensitive_true", testRequest)
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:   "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg",
		Name: ol.RefOf("Test Updated"),
		Predicates: &[]ol.FilterPredicateInput{{
			Key:           ol.PredicateKeyEnumTierIndex,
			Type:          ol.PredicateTypeEnumEquals,
			Value:         ol.RefOf("1"),
			CaseSensitive: ol.RefOf(true),
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
	autopilot.Equals(t, ol.RefOf(true), result.Predicates[0].CaseSensitive)
}

func TestUpdateFilterCaseSensitiveFalse(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation FilterUpdate($input:FilterUpdateInput!){filterUpdate(input: $input){filter{id,name,connective,htmlUrl,predicates{key,keyData,type,value,caseSensitive}},errors{message,path}}}`,
		`{"input": {"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg", "name": "Test Updated", "predicates": [ { "key": "tier_index", "type": "equals", "value": "1", "caseSensitive": false } ] }}`,
		`{"data": {
      "filterUpdate": {
        "filter": {
          "connective": null,
          "htmlUrl": "https://app.opslevel.com/filters/401",
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzQwMQ",
          "name": "Tier 1 Services",
          "predicates": [
          {
            "key": "tier_index",
            "keyData": null,
            "type": "equals",
            "value": "1",
            "caseSensitive": false
          }
          ]
        },
        "errors": []
      }}}`,
	)

	client := BestTestClient(t, "filter/update_case_sensitive_false", testRequest)
	// Act
	result, err := client.UpdateFilter(ol.FilterUpdateInput{
		Id:   "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyMg",
		Name: ol.RefOf("Test Updated"),
		Predicates: &[]ol.FilterPredicateInput{{
			Key:           ol.PredicateKeyEnumTierIndex,
			Type:          ol.PredicateTypeEnumEquals,
			Value:         ol.RefOf("1"),
			CaseSensitive: ol.RefOf(false),
		}},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Tier 1 Services", result.Name)
	autopilot.Equals(t, ol.PredicateKeyEnumTierIndex, result.Predicates[0].Key)
	autopilot.Equals(t, ol.RefOf(false), result.Predicates[0].CaseSensitive)
}

func TestDeleteFilter(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation FilterDelete($input:DeleteInput!){filterDelete(input: $input){deletedId,errors{message,path}}}`,
		`{"input": {"id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz" }}`,
		`{"data": {"filterDelete": {"deletedId": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzYyNQ", "errors": [] }}}`,
	)

	client := BestTestClient(t, "filter/delete", testRequest)
	// Act
	err := client.DeleteFilter("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
