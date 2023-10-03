package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestSystemCreate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemCreate($input:SystemInput!){systemCreate(input:$input){system{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},errors{message,path}}}",
	"variables":{
		"input": {
			"name": "PlatformSystem3",
			"description": "creating this for testing purposes",
			"ownerId": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU",
			"note": "hello world"
		  }
    }}`
	response := `{"data": {
		"systemCreate": {
			"system": {{ template "system1_response" }},
			"errors": []
	}}}`
	client := ABetterTestClient(t, "system/create", request, response)
	input := ol.SystemInput{
		Name:        ol.NewString("PlatformSystem3"),
		Description: ol.NewString("creating this for testing purposes"),
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        ol.NewString("hello world"),
	}
	// Act
	result, err := client.CreateSystem(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy", string(result.Id))
	autopilot.Equals(t, "An example description", result.Note)
	autopilot.Equals(t, "An example description", result.Parent.Note)
}

func TestSystemGetServices(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			Request: `{"query": "query SystemChildServicesList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){childServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"system": {
					"id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx"
				}}}`,
			Response: `{
				"data": {
					"account": {
						"system": {
							"childServices": {
								"nodes": [
									{{ template "service_1" }},
									{{ template "service_2" }}
								],
								{{ template "pagination_initial_pageInfo_response" }},
								"totalCount": 2
							}}}}}`,
		},
		{
			Request: `{"query": "query SystemChildServicesList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){childServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"system": {
					"id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx"
				}}}`,
			Response: `{
				"data": {
					"account": {
						"system": {
							"childServices": {
								"nodes": [
									{{ template "service_2" }}
								],
								{{ template "pagination_second_pageInfo_response" }},
								"totalCount": 1
							}}}}}`,
		},
	}

	client := APaginatedTestClient(t, "system/child_services", requests...)
	system := ol.SystemId{
		Id: "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx",
	}
	// Act
	resp, err := system.ChildServices(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, "Bar", result[1].Name)
	autopilot.Equals(t, "Bar", result[2].Name)
}

func TestSystemGetTags(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			Request: `{"query": "query SystemTagsList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"system": {
					"id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx"
				}
			}
			}`,
			Response: `{
				"data": {
					"account": {
						"system": {
							"tags": {
								"nodes": [
									{{ template "tag1" }},
									{{ template "tag2" }}
								],
								{{ template "pagination_initial_pageInfo_response" }},
								"totalCount": 2
							}
						  }}}}`,
		},
		{
			Request: `{"query": "query SystemTagsList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"system": {
					"id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx"
				}}}`,
			Response: `{
				"data": {
					"account": {
						"system": {
							"tags": {
								"nodes": [
									{{ template "tag3" }}
								],
								{{ template "pagination_second_pageInfo_response" }},
								"totalCount": 1
							}}}}}`,
		},
	}

	client := APaginatedTestClient(t, "system/tags", requests...)
	system := ol.SystemId{
		Id: "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx",
	}
	// Act
	resp, err := system.Tags(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "dev", result[0].Key)
	autopilot.Equals(t, "true", result[0].Value)
	autopilot.Equals(t, "foo", result[1].Key)
	autopilot.Equals(t, "bar", result[1].Value)
	autopilot.Equals(t, "prod", result[2].Key)
	autopilot.Equals(t, "true", result[2].Value)
}

func TestSystemAssignService(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemAssignService($childServices:[IdentifierInput!]!$system:IdentifierInput!){systemChildAssign(system:$system, childServices:$childServices){system{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},errors{message,path}}}",
	"variables":{
		"system":{
			"id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx"
	  	},
	  	"childServices": [
			{"id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy"}
	  	]
	}
}`
	response := `{"data": {
		"systemChildAssign": {
			"system": {{ template "system1_response" }}
		}
}}`
	client := ABetterTestClient(t, "system/assign_service", request, response)
	// Act
	system := ol.System{
		SystemId: ol.SystemId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx",
		},
	}
	childServices := "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy"
	err := system.AssignService(client, childServices)
	// Assert
	autopilot.Ok(t, err)
}

func TestSystemGetId(t *testing.T) {
	// Arrange
	request := `{
    "query": "query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note}}}",
	"variables":{
		"input": {
			"id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy"
    	}
    }}`
	response := `{"data": {
		"account": {
			"system":
				{{ template "system1_response" }}
		}
	}}`
	client := ABetterTestClient(t, "system/get_id", request, response)
	// Act
	result, err := client.GetSystem("Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy", string(result.Id))
}

func TestSystemGetAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note}}}",
	"variables":{
		"input": {
			"alias": "platformsystem1"
		}
    }}`
	response := `{"data": {
		"account": {
			"system": {{ template "system1_response" }}
		}
	}}`
	client := ABetterTestClient(t, "system/get_alias", request, response)
	// Act
	result, err := client.GetSystem("platformsystem1")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy", string(result.Id))
}

func TestListSystems(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			Request: `{"query": "query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			Response: `{
				"data": {
					"account": {
						"systems": {
							"nodes": [
								{{ template "system1_response" }},
								{{ template "system2_response" }}
							],
							{{ template "pagination_initial_pageInfo_response" }}
						  }}}}`,
		},
		{
			Request: `{"query": "query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			Response: `{
				"data": {
					"account": {
						"systems": {
							"nodes": [
								{{ template "system3_response" }}
							],
							{{ template "pagination_second_pageInfo_response" }}
						  }}}}`,
		},
	}

	client := APaginatedTestClient(t, "system/list", requests...)
	// Act
	response, err := client.ListSystems(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "PlatformSystem1", result[0].Name)
	autopilot.Equals(t, "PlatformSystem2", result[1].Name)
	autopilot.Equals(t, "", string(result[1].Parent.Id))
	autopilot.Equals(t, "PlatformSystem3", result[2].Name)
}

func TestSystemUpdate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemUpdate($input:SystemInput!$system:IdentifierInput!){systemUpdate(system:$system,input:$input){system{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},errors{message,path}}}",
	"variables":{
		"system":{"id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy"},
		"input":{
			"name": "PlatformSystem1",
			"description":"Yolo!",
			"ownerId":"Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU",
			"note": "Please delete me"
		}
	}}`
	response := `{"data": {
		"systemUpdate": {
			"system": {{ template "system1_response" }},
			"errors": []
	}}}`
	client := ABetterTestClient(t, "system/update", request, response)
	input := ol.SystemInput{
		Name:        ol.NewString("PlatformSystem1"),
		Description: ol.NewString("Yolo!"),
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        ol.NewString("Please delete me"),
	}
	// Act
	result, err := client.UpdateSystem("Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy", input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy", string(result.Id))
	autopilot.Equals(t, "An example description", result.Note)
	autopilot.Equals(t, "An example description", result.Parent.Note)
}

func TestSystemDelete(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemDelete($input:IdentifierInput!){systemDelete(resource: $input){errors{message,path}}}",
	"variables":{"input":{"alias":"PlatformSystem3"}}
}`
	response := `{"data": {
	"systemDelete": {
      "errors": []
    }
}}`
	client := ABetterTestClient(t, "system/delete", request, response)
	// Act
	err := client.DeleteSystem("PlatformSystem3")
	// Assert
	autopilot.Ok(t, err)
}
