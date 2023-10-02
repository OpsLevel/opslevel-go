package opslevel_test

import (
	"testing"

	"github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "mutation InfrastructureResourceCreate($all:Boolean!$input:InfrastructureResourceInput!){infrastructureResourceCreate(input: $input){infrastructureResource{id,aliases,name,type @include(if: $all),providerResourceType @include(if: $all),providerData @include(if: $all){accountName,externalUrl,providerName},owner @include(if: $all){... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},ownerLocked @include(if: $all),data @include(if: $all),rawData @include(if: $all)},warnings{message},errors{message,path}}}",
  "variables":{
    "all": true,
    "input": {
      "ownerId":"Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
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
    }
  }}`
	response := `{
  "data": {
    "infrastructureResourceCreate": {
      "infrastructureResource": {{ template "infra_1" }},
      "warnings": [],
      "errors": []
    }
  }
}`
	client := ABetterTestClient(t, "infra/create", request, response)
	// Act
	result, err := client.CreateInfrastructure(opslevel.InfraInput{
		Schema: "Database",
		Owner:  opslevel.NewID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"),
		Provider: &opslevel.InfraProviderInput{
			Account: "Dev - 123456789",
			Name:    "Google",
			Type:    "BigQuery",
			URL:     "https://google.com",
		},
		Data: opslevel.JSON{
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
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "my-big-query", result.Name)
}

func TestGetInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "query InfrastructureResourceGet($all:Boolean!$input:IdentifierInput!){account{infrastructureResource(input: $input){id,aliases,name,type @include(if: $all),providerResourceType @include(if: $all),providerData @include(if: $all){accountName,externalUrl,providerName},owner @include(if: $all){... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},ownerLocked @include(if: $all),data @include(if: $all),rawData @include(if: $all)}}}",
	"variables":{
		"all": true,
		"input":{
			"id":"Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
		}
	}
}`
	response := `{"data": {
	"account": {
		"infrastructureResource": {{ template "infra_1" }}
	}
}}`
	client := ABetterTestClient(t, "infra/get", request, response)
	// Act
	result, err := client.GetInfrastructure("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "my-big-query", result.Name)
}

func TestListInfraSchemas(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			`{"query": "query IntegrationList($after:String!$first:Int!){account{infrastructureResourceSchemas(after: $after, first: $first){nodes{type,schema},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
				{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"infrastructureResourceSchemas": {
							"nodes": [
								{{ template "infra_schema_1" }},
								{{ template "infra_schema_2" }}
							],
							{{ template "pagination_initial_pageInfo_response" }}
					}}}}`,
		},
		{
			`{"query": "query IntegrationList($after:String!$first:Int!){account{infrastructureResourceSchemas(after: $after, first: $first){nodes{type,schema},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
				{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"infrastructureResourceSchemas": {
							"nodes": [
								{{ template "infra_schema_3" }}
							],
							{{ template "pagination_second_pageInfo_response" }}
					}}}}`,
		},
	}

	client := APaginatedTestClient(t, "infra/list_schemas", requests...)
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
	requests := []TestRequest{
		{
			`{"query": "query IntegrationList($after:String!$all:Boolean!$first:Int!){account{infrastructureResources(after: $after, first: $first){nodes{id,aliases,name,type @include(if: $all),providerResourceType @include(if: $all),providerData @include(if: $all){accountName,externalUrl,providerName},owner @include(if: $all){... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},ownerLocked @include(if: $all),data @include(if: $all),rawData @include(if: $all)},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
				"variables": {
					"after": "",
					"all": true,
					"first": 100
				}
			}`,
			`{
				"data": {
					"account": {
						"infrastructureResources": {
							"nodes": [
								{{ template "infra_1" }},
								{{ template "infra_2" }}
							],
							{{ template "pagination_initial_pageInfo_response" }}
					}}}}`,
		},
		{
			`{"query": "query IntegrationList($after:String!$all:Boolean!$first:Int!){account{infrastructureResources(after: $after, first: $first){nodes{id,aliases,name,type @include(if: $all),providerResourceType @include(if: $all),providerData @include(if: $all){accountName,externalUrl,providerName},owner @include(if: $all){... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},ownerLocked @include(if: $all),data @include(if: $all),rawData @include(if: $all)},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
				"variables": {
					"after": "OA",
					"all": true,
					"first": 100
				}
			}`,
			`{
				"data": {
					"account": {
						"infrastructureResources": {
							"nodes": [
								{{ template "infra_3" }}
							],
							{{ template "pagination_second_pageInfo_response" }}
					}}}}`,
		},
	}

	client := APaginatedTestClient(t, "infra/list", requests...)
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
	request := `{
	"query": "mutation InfrastructureResourceUpdate($all:Boolean!$identifier:IdentifierInput!$input:InfrastructureResourceInput!){infrastructureResourceUpdate(infrastructureResource: $identifier, input: $input){infrastructureResource{id,aliases,name,type @include(if: $all),providerResourceType @include(if: $all),providerData @include(if: $all){accountName,externalUrl,providerName},owner @include(if: $all){... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},ownerLocked @include(if: $all),data @include(if: $all),rawData @include(if: $all)},warnings{message},errors{message,path}}}",
  "variables":{
    "all": true,
    "identifier": {"id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"},
    "input": {
      "ownerId":"Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
      "data": "{\"endpoint\":\"https://google.com\",\"engine\":\"BigQuery\",\"name\":\"my-big-query\",\"replica\":false}"
    }
}
}`
	response := `{"data": {
    "infrastructureResourceUpdate": {
      "infrastructureResource": {{ template "infra_1" }},
      "warnings": [],
      "errors": []
    }
}}`
	client := ABetterTestClient(t, "infra/update", request, response)
	// Act
	result, err := client.UpdateInfrastructure("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", opslevel.InfraInput{
		Owner: opslevel.NewID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"),
		Data: opslevel.JSON{
			"name":     "my-big-query",
			"engine":   "BigQuery",
			"endpoint": "https://google.com",
			"replica":  false,
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "my-big-query", result.Name)
}

func TestDeleteInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "mutation InfrastructureResourceDelete($input:IdentifierInput!){infrastructureResourceDelete(resource: $input){errors{message,path}}}",
	"variables":{
    "input": {
      "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
    }
  }}`
	response := `{"data": {
		"infrastructureResourceDelete": {
			"errors": []
		}
	}}`
	client := ABetterTestClient(t, "infra/delete", request, response)
	// Act
	err := client.DeleteInfrastructure("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Equals(t, nil, err)
}

func TestGetInfrastructureResourceTags(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			`{"query": "query InfrastructureResourceTagsList($after:String!$first:Int!$infrastructureResource:IdentifierInput!){account{infrastructureResources(input: $infrastructureResource){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
            "variables": {
                {{ template "first_page_variables" }},
                "infrastructureResource": {"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"}
            }
            }`,
			`{
                  "data": {
                    "account": {
                      "infrastructureResource": {
                        "tags": {
                          "nodes": [
                            {
                              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIw",
                              "key": "abc",
                              "value": "abc"
                            },
                            {
                              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIx",
                              "key": "db",
                              "value": "mongoqqqq"
                            },
                            {
                              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIy",
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
		},
		{
			`{"query": "query InfrastructureResourceTagsList($after:String!$first:Int!$infrastructureResource:IdentifierInput!){account{infrastructureResources(input: $infrastructureResource){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
            "variables": {
                {{ template "second_page_variables" }},
                "infrastructureResource": {"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"}
            }
            }`,
			`{
                  "data": {
                    "account": {
                      "infrastructureResource": {
                        "tags": {
                          "nodes": [
                            {
                              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIz",
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
		},
	}
	client := APaginatedTestClient(t, "infrastructureResource/tags", requests...)
	// Act
	infra := opslevel.InfrastructureResource{Id: "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"}
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
