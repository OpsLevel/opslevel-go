package opslevel_test

import (
	"testing"

	"github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateInfra(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation InfrastructureResourceCreate($all:Boolean!$input:InfrastructureResourceInput!){infrastructureResourceCreate(input: $input){infrastructureResource{aliases,data,href,id,integration,lastSyncedAt,name,owner{... on Team{teamAlias:alias,id}},ownerLocked,providerData{accountName,externalUrl,providerName},providerResourceType,rawData,type},warnings{message},errors{message,path}}}`,
		`{
    "all": true,
    "input": {
      "ownerId": "{{ template "id1_string" }}",
      "data": "{\"endpoint\":\"https://google.com\",\"engine\":\"BigQuery\",\"name\":\"my-big-query\",\"replica\":false,\"storage_size\":{\"unit\":\"GB\",\"value\":1024}}",
      "providerData": {
        "accountName": "Dev - 123456789",
        "externalUrl": "https://google.com",
        "providerName": "Google"
      },
      "providerResourceType": "BigQuery",
      "schema": {
        "type": "Database"
       }
    }}`,
		`{ "data": { "infrastructureResourceCreate": { "infrastructureResource": {{ template "infra_1" }}, "warnings": [], "errors": [] }}}`,
	)
	client := BestTestClient(t, "infra/create", testRequest)
	// Act
	result, err := client.CreateInfrastructure(opslevel.InfraInput{
		Schema: "Database",
		Owner:  &id1,
		Provider: &opslevel.InfraProviderInput{
			Account: "Dev - 123456789",
			Name:    "Google",
			Type:    "BigQuery",
			URL:     "https://google.com",
		},
		Data: &opslevel.JSON{
			"name":     "my-big-query",
			"engine":   "BigQuery",
			"endpoint": "https://google.com",
			"replica":  false,
			"storage_size": map[string]any{
				"value": 1024,
				"unit":  "GB",
			},
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "my-big-query", result.Name)
}

func TestGetInfra(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query InfrastructureResourceGet($all:Boolean!$input:IdentifierInput!){account{infrastructureResource(input: $input){aliases,data,href,id,integration,lastSyncedAt,name,owner{... on Team{teamAlias:alias,id}},ownerLocked,providerData{accountName,externalUrl,providerName},providerResourceType,rawData,type}}}`,
		`{"all": true, "input":{ {{ template "id1" }} }}`,
		`{"data": { "account": { "infrastructureResource": {{ template "infra_1" }} }}}`,
	)
	client := BestTestClient(t, "infra/get", testRequest)
	// Act
	result, err := client.GetInfrastructure(string(id1))
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "my-big-query", result.Name)
}

func TestListInfraSchemas(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query IntegrationList($after:String!$first:Int!){account{infrastructureResourceSchemas(after: $after, first: $first){nodes{schema,type},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "infrastructureResourceSchemas": { "nodes": [ {{ template "infra_schema_1" }}, {{ template "infra_schema_2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query IntegrationList($after:String!$first:Int!){account{infrastructureResourceSchemas(after: $after, first: $first){nodes{schema,type},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "infrastructureResourceSchemas": { "nodes": [ {{ template "infra_schema_3" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "infra/list_schemas", requests...)
	// Act
	response, err := client.ListInfrastructureSchemas(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Database", result[0].Type)
	autopilot.Equals(t, "Compute", result[1].Type)
	autopilot.Equals(t, "Queue", result[2].Type)
}

func TestListInfra(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query IntegrationList($after:String!$all:Boolean!$first:Int!){account{infrastructureResources(after: $after, first: $first){nodes{aliases,data,href,id,integration,lastSyncedAt,name,owner{... on Team{teamAlias:alias,id}},ownerLocked,providerData{accountName,externalUrl,providerName},providerResourceType,rawData,type},pageInfo{endCursor,hasNextPage,hasPreviousPage,startCursor}}}}`,
		`{ "after": "", "all": true, "first": 100 }`,
		`{ "data": { "account": { "infrastructureResources": { "nodes": [ {{ template "infra_1" }}, {{ template "infra_2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query IntegrationList($after:String!$all:Boolean!$first:Int!){account{infrastructureResources(after: $after, first: $first){nodes{aliases,data,href,id,integration,lastSyncedAt,name,owner{... on Team{teamAlias:alias,id}},ownerLocked,providerData{accountName,externalUrl,providerName},providerResourceType,rawData,type},pageInfo{endCursor,hasNextPage,hasPreviousPage,startCursor}}}}`,
		`{ "after": "OA", "all": true, "first": 100 }`,
		`{ "data": { "account": { "infrastructureResources": { "nodes": [ {{ template "infra_3" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "infra/list", requests...)
	// Act
	response, err := client.ListInfrastructure(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "vpc-XXXXXXXXXX", result[1].Name)
	autopilot.Equals(t, "production-demo", result[2].Name)
}

func TestUpdateInfra(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation InfrastructureResourceUpdate($all:Boolean!$identifier:IdentifierInput!$input:InfrastructureResourceInput!){infrastructureResourceUpdate(infrastructureResource: $identifier, input: $input){infrastructureResource{aliases,data,href,id,integration,lastSyncedAt,name,owner{... on Team{teamAlias:alias,id}},ownerLocked,providerData{accountName,externalUrl,providerName},providerResourceType,rawData,type},warnings{message},errors{message,path}}}`,
		`{"all": true, "identifier": { {{ template "id1" }}}, "input": { "ownerId": "{{ template "id1_string" }}", "data": "{\"endpoint\":\"https://google.com\",\"engine\":\"BigQuery\",\"name\":\"my-big-query\",\"replica\":false}" }}`,
		`{"data": { "infrastructureResourceUpdate": { "infrastructureResource": {{ template "infra_1" }}, "warnings": [], "errors": [] }}}`,
	)
	client := BestTestClient(t, "infra/update", testRequest)
	// Act
	result, err := client.UpdateInfrastructure(string(id1), opslevel.InfraInput{
		Owner: &id1,
		Data: &opslevel.JSON{
			"name":     "my-big-query",
			"engine":   "BigQuery",
			"endpoint": "https://google.com",
			"replica":  false,
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "my-big-query", result.Name)
}

func TestDeleteInfra(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation InfrastructureResourceDelete($input:IdentifierInput!){infrastructureResourceDelete(resource: $input){errors{message,path}}}`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "infrastructureResourceDelete": { "errors": [] }}}`,
	)
	client := BestTestClient(t, "infra/delete", testRequest)
	// Act
	err := client.DeleteInfrastructure(string(id1))
	// Assert
	autopilot.Equals(t, nil, err)
}

func TestGetInfrastructureResourceTags(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query InfrastructureResourceTags($after:String!$first:Int!$infrastructureResource:IdentifierInput!){account{infrastructureResource(input: $infrastructureResource){tags(after: $after, first: $first){nodes{id,key,owner{... on Team{alias,id}},value},pageInfo{endCursor,hasNextPage,hasPreviousPage,startCursor},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "infrastructureResource": { {{template "id1" }} } }`,
		`{
                  "data": {
                    "account": {
                      "infrastructureResource": {
                        "tags": {
                          "nodes": [
                            {
                              {{ template "id2" }},
                              "key": "abc",
                              "value": "abc"
                            },
                            {
                              {{ template "id3" }},
                              "key": "db",
                              "value": "mongoqqqq"
                            },
                            {
                              {{ template "id4" }},
                              "key": "db",
                              "value": "prod"
                            }
                          ],
                          {{ template "pagination_initial_pageInfo_response" }},
                          "totalCount": 3
                        }
                      }
                    }
                  }
                }`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query InfrastructureResourceTags($after:String!$first:Int!$infrastructureResource:IdentifierInput!){account{infrastructureResource(input: $infrastructureResource){tags(after: $after, first: $first){nodes{id,key,owner{... on Team{alias,id}},value},pageInfo{endCursor,hasNextPage,hasPreviousPage,startCursor},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "infrastructureResource": { {{template "id1" }} }}`,
		`{
                  "data": {
                    "account": {
                      "infrastructureResource": {
                        "tags": {
                          "nodes": [
                            {
                              {{ template "id3" }},
                              "key": "env",
                              "value": "staging"
                            }
                          ],
                          {{ template "pagination_second_pageInfo_response" }},
                          "totalCount": 1
                        }
                      }
                    }
                  }
                }`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "infrastructureResource/tags", requests...)
	// Act
	infra := opslevel.InfrastructureResource{Id: id1}
	resp, err := infra.GetTags(client, nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, resp.TotalCount)
	autopilot.Equals(t, "abc", result[0].Key)
	autopilot.Equals(t, "abc", result[0].Value)
	autopilot.Equals(t, "db", result[2].Key)
	autopilot.Equals(t, "prod", result[2].Value)
	autopilot.Equals(t, "env", result[3].Key)
	autopilot.Equals(t, "staging", result[3].Value)
}
