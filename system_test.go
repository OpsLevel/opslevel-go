package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestSystemCreate(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation SystemCreate($input:SystemInput!){systemCreate(input:$input){system{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},errors{message,path}}}`,
		`{"input": { "name": "PlatformSystem3", "description": "creating this for testing purposes", "ownerId": "{{ template "id4_string" }}", "note": "hello world" } }`,
		`{"data": { "systemCreate": { "system": {{ template "system1_response" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "system/create", testRequest)
	input := ol.SystemInput{
		Name:        ol.RefOf("PlatformSystem3"),
		Description: ol.RefOf("creating this for testing purposes"),
		OwnerId:     ol.RefOf(id4),
		Note:        ol.RefOf("hello world"),
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
	testRequestOne := autopilot.NewTestRequest(
		`query SystemChildServicesList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){childServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "system": { "id": "Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx" }}`,
		`{ "data": { "account": { "system": { "childServices": { "nodes": [ {{ template "service_1" }}, {{ template "service_2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query SystemChildServicesList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){childServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "system": { "id": "Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx" }}`,
		`{ "data": { "account": { "system": { "childServices": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

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
	testRequestOne := autopilot.NewTestRequest(
		`query SystemTagsList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "system": { {{ template "id3" }} } }`,
		`{ "data": { "account": { "system": { "tags": { "nodes": [ {{ template "tag1" }}, {{ template "tag2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query SystemTagsList($after:String!$first:Int!$system:IdentifierInput!){account{system(input: $system){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "system": { {{ template "id3" }} }}`,
		`{ "data": { "account": { "system": { "tags": { "nodes": [ {{ template "tag3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "system/tags", requests...)
	system := ol.SystemId{
		Id: id3,
	}
	// Act
	resp, err := system.GetTags(client, nil)
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
	testRequest := autopilot.NewTestRequest(
		`mutation SystemAssignService($childServices:[IdentifierInput!]!$system:IdentifierInput!){systemChildAssign(system:$system, childServices:$childServices){system{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},errors{message,path}}}`,
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
	testRequest := autopilot.NewTestRequest(
		`query SystemGet($input:IdentifierInput!){account{system(input: $input){{ template "system_get" }}}}`,
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
	testRequest := autopilot.NewTestRequest(
		`query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}}`,
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
	testRequestOne := autopilot.NewTestRequest(
		`query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "systems": { "nodes": [ {{ template "system1_response" }}, {{ template "system2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "systems": { "nodes": [ {{ template "system3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

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
	testRequest := autopilot.NewTestRequest(
		`mutation SystemUpdate($input:SystemInput!$system:IdentifierInput!){systemUpdate(system:$system,input:$input){system{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},errors{message,path}}}`,
		`{"system": { {{ template "id1" }} }, "input":{ "name": "PlatformSystem1", "description":"Yolo!", "ownerId":"{{ template "id4_string" }}", "note": "Please delete me" }}`,
		`{"data": {"systemUpdate": {"system": {{ template "system1_response" }}, "errors": [] }}}`,
	)
	client := BestTestClient(t, "system/update", testRequest)
	input := ol.SystemInput{
		Name:        ol.RefOf("PlatformSystem1"),
		Description: ol.RefOf("Yolo!"),
		OwnerId:     ol.RefOf(id4),
		Note:        ol.RefOf("Please delete me"),
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
	testRequest := autopilot.NewTestRequest(
		`mutation SystemDelete($input:IdentifierInput!){systemDelete(resource: $input){errors{message,path}}}`,
		`{"input":{"alias":"PlatformSystem3"}}`,
		`{"data": { "systemDelete": { "errors": [] } }}`,
	)
	client := BestTestClient(t, "system/delete", testRequest)
	// Act
	err := client.DeleteSystem("PlatformSystem3")
	// Assert
	autopilot.Ok(t, err)
}

func TestSystemReconcileAliasesDeleteAll(t *testing.T) {
	// Arrange
	aliasesWanted := []string{}
	system := ol.System{
		SystemId: ol.SystemId{
			Id:      id1,
			Aliases: []string{"one", "two"},
		},
	}

	// delete "one" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "one", "ownerType": "system" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "two" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerType": "system" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	testRequestThree := autopilot.NewTestRequest(
		`query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}}`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "account": { "system": { {{ template "id1" }}, "aliases": [] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree}
	client := BestTestClient(t, "system/reconcile_aliases_delete_all", requests...)

	// Act
	err := system.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, len(system.Aliases), 0)
}

func TestSystemReconcileAliasesDeleteSome(t *testing.T) {
	// Arrange
	aliasesWanted := []string{"two"}
	system := ol.System{
		SystemId: ol.SystemId{
			Id:      id1,
			Aliases: []string{"one", "two"},
		},
	}

	// delete "one" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "one", "ownerType": "system" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}}`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "account": { "system": { {{ template "id1" }}, "aliases": ["two"] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "system/reconcile_aliases_delete_some", requests...)

	// Act
	err := system.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, len(system.Aliases), 1)
}

func TestSystemReconcileAliases(t *testing.T) {
	// Arrange
	aliasesWanted := []string{"one", "two", "three"}
	system := ol.SystemId{
		Id:      id1,
		Aliases: []string{"one", "alpha", "beta"},
	}

	// delete "alpha" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "alpha", "ownerType": "system" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "beta" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "beta", "ownerType": "system" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// create "two" alias
	testRequestThree := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "one", "two" ], "ownerId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)
	// create "three" alias
	testRequestFour := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input":{ "alias": "three", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "one", "two", "three" ], "ownerId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)
	// get service
	testRequestFive := autopilot.NewTestRequest(
		`query SystemGet($input:IdentifierInput!){account{system(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}}`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "account": { "system": { {{ template "id1" }}, "aliases": ["one", "two", "three"] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree, testRequestFour, testRequestFive}
	client := BestTestClient(t, "system/reconcile_aliases", requests...)

	// Act
	err := system.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, system.Aliases, aliasesWanted)
}
