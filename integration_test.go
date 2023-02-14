package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2022"
)

func TestGetIntegration(t *testing.T) {
	// Arrange
	request := `{
	"query": "query IntegrationGet($id:ID!){account{integration(id: $id){id,name,type}}}",
	"variables":{
		"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf"
    }
}`
	response := `{"data": {
	"account": {
		"integration": {
			"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf",
			"name": "Deploy",
			"type": "deploy"
		}
	}
}}`
	client := ABetterTestClient(t, "integration/get", request, response)
	// Act
	result, err := client.GetIntegration("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf", string(result.Id))
	autopilot.Equals(t, "Deploy", result.Name)
}

func TestGetMissingIntegraion(t *testing.T) {
	// Arrange
	request := `{
	"query": "query IntegrationGet($id:ID!){account{integration(id: $id){id,name,type}}}",
	"variables":{
		"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf"
    }
}`
	response := `{"data": {
	"account": {
		"integration": null
	}
}}`
	client := ABetterTestClient(t, "integration/get_missing", request, response)
	// Act
	_, err := client.GetIntegration("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListIntegrations(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query IntegrationList($after:String!$first:Int!){account{integrations(after: $after, first: $first){nodes{id,name,type},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"integrations": {
							"nodes": [
								{
									{{ template "deploy_integration_response" }}
								},
								{
									{{ template "payload_integration_response" }} 
								}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
					}}}}`},
		{`{"query": "query IntegrationList($after:String!$first:Int!){account{integrations(after: $after, first: $first){nodes{id,name,type},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"integrations": {
							"nodes": [
								{
									{{ template "kubernetes_integration_response" }}
								}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
					}}}}`},
	}

	client := APaginatedTestClient(t, "integration/list", requests...)
	// Act
	response, err := client.ListIntegrations(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Payload", result[1].Name)
	autopilot.Equals(t, "Kubernetes", result[2].Name)
	//fmt.Println(Templated(requests[1].Request))
	//panic(true)
}
