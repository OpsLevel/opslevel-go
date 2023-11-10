package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestSystemCreate(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation SystemCreate($input:SystemInput!){systemCreate(input:$input){system{id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note},errors{message,path}}}"`,
		`{"input": { "name": "PlatformSystem3", "description": "creating this for testing purposes", "ownerId": "{{ template "id4_string" }}", "note": "hello world" } }`,
		`{"data": { "systemCreate": { "system": {{ template "system1_response" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "system/create", testRequest)
	input := ol.SystemInput{
		Name:        ol.NewString("PlatformSystem3"),
		Description: ol.NewString("creating this for testing purposes"),
		Owner:       &id4,
		Note:        ol.NewString("hello world"),
	}
	// Act
	result, err := client.CreateSystem(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "An example description", result.Note)
	autopilot.Equals(t, "An example description", result.Parent.Note)
}

func TestSystemGetServices(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query SystemChildServicesList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){childServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "system": { "id": "Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx" }}`,
		`{ "data": { "account": { "system": { "childServices": { "nodes": [ {{ template "service_1" }}, {{ template "service_2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query SystemChildServicesList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){childServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "system": { "id": "Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx" }}`,
		`{ "data": { "account": { "system": { "childServices": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "system/child_services", requests...)
	system := ol.SystemId{
		Id: id3,
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
	testRequestOne := NewTestRequest(
		`"query SystemTagsList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "system": { {{ template "id3" }} } }`,
		`{ "data": { "account": { "system": { "tags": { "nodes": [ {{ template "tag1" }}, {{ template "tag2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query SystemTagsList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "system": { {{ template "id3" }} }}`,
		`{ "data": { "account": { "system": { "tags": { "nodes": [ {{ template "tag3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "system/tags", requests...)
	system := ol.SystemId{
		Id: id3,
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
	testRequest := NewTestRequest(
		`"mutation SystemAssignService($childServices:[IdentifierInput!]!$system:IdentifierInput!){systemChildAssign(system:$system, childServices:$childServices){system{id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note},errors{message,path}}}"`,
		`{"system": { {{ template "id3" }} }, "childServices": [ { {{ template "id4" }} } ] }`,
		`{"data": { "systemChildAssign": { "system": {{ template "system1_response" }} } }}`,
	)

	client := BestTestClient(t, "system/assign_service", testRequest)
	// Act
	system := ol.System{
		SystemId: ol.SystemId{
			Id: id3,
		},
	}
	err := system.AssignService(client, string(id4))
	// Assert
	autopilot.Ok(t, err)
}

func TestSystemGetId(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note}}}"`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "account": { "system": {{ template "system1_response" }} }}}`,
	)
	client := BestTestClient(t, "system/get_id", testRequest)
	// Act
	result, err := client.GetSystem(string(id1))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestSystemGetAlias(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note}}}"`,
		`{ "input": { "alias": "platformsystem1" } }`,
		`{"data": { "account": { "system": {{ template "system1_response" }} }}}`,
	)
	client := BestTestClient(t, "system/get_alias", testRequest)
	// Act
	result, err := client.GetSystem("platformsystem1")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestListSystems(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note},{{ template "pagination_request" }}}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "systems": { "nodes": [ {{ template "system1_response" }}, {{ template "system2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note},{{ template "pagination_request" }}}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "systems": { "nodes": [ {{ template "system3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "system/list", requests...)
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
	testRequest := NewTestRequest(
		`"mutation SystemUpdate($input:SystemInput!$system:IdentifierInput!){systemUpdate(system:$system,input:$input){system{id,aliases,name,description,htmlUrl,owner{alias,id},parent{id,aliases,name,description,htmlUrl,owner{alias,id},note},note},errors{message,path}}}"`,
		`{"system": { {{ template "id1" }} }, "input":{ "name": "PlatformSystem1", "description":"Yolo!", "ownerId":"{{ template "id4_string" }}", "note": "Please delete me" }}`,
		`{"data": {"systemUpdate": {"system": {{ template "system1_response" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "system/update", testRequest)
	input := ol.SystemInput{
		Name:        ol.NewString("PlatformSystem1"),
		Description: ol.NewString("Yolo!"),
		Owner:       &id4,
		Note:        ol.NewString("Please delete me"),
	}
	// Act
	result, err := client.UpdateSystem(string(id1), input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "An example description", result.Note)
	autopilot.Equals(t, "An example description", result.Parent.Note)
}

func TestSystemDelete(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation SystemDelete($input:IdentifierInput!){systemDelete(resource: $input){errors{message,path}}}"`,
		`{"input":{"alias":"PlatformSystem3"}}`,
		`{"data": { "systemDelete": { "errors": [] } }}`,
	)
	client := BestTestClient(t, "system/delete", testRequest)
	// Act
	err := client.DeleteSystem("PlatformSystem3")
	// Assert
	autopilot.Ok(t, err)
}
