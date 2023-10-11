package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestDomainCreate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation DomainCreate($input:DomainInput!){domainCreate(input:$input){domain{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},errors{message,path}}}",
	"variables":{
		"input": {
			"name": "platform-test",
			"description": "Domain created for testing.",
			"ownerId": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU",
			"note": "additional note about platform-test domain"
		}
    }
}`
	response := `{"data": {
		"domainCreate": {
			"domain": {{ template "domain1_response" }}
		}
}}`
	client := ABetterTestClient(t, "domain/create", request, response)
	// Act
	input := ol.DomainInput{
		Name:        ol.NewString("platform-test"),
		Description: ol.NewString("Domain created for testing."),
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        ol.NewString("additional note about platform-test domain"),
	}
	result, err := client.CreateDomain(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw", string(result.Id))
	autopilot.Equals(t, "An example description", result.Note)
}

func TestDomainGetSystems(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query DomainChildSystemsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){childSystems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}"`,
		`"variables": { {{ template "first_page_variables" }}, "domain": { "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx" } }`,
		`{ "data": { "account": { "domain": { "childSystems": { "nodes": [ {{ template "system1_response" }}, {{ template "system2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query DomainChildSystemsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){childSystems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},note},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}"`,
		`"variables": { {{ template "second_page_variables" }}, "domain": { "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx" } }`,
		`{ "data": { "account": { "domain": { "childSystems": { "nodes": [ {{ template "system3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := APaginatedTestClient(t, "domain/child_systems", requests...)
	domain := ol.DomainId{
		Id: "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx",
	}
	// Act
	resp, err := domain.ChildSystems(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "PlatformSystem1", result[0].Name)
	autopilot.Equals(t, "PlatformSystem2", result[1].Name)
	autopilot.Equals(t, "PlatformSystem3", result[2].Name)
}

func TestDomainGetTags(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		`"variables": { {{ template "first_page_variables" }}, "domain": { "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx" } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag1" }}, {{ template "tag2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		`"variables": { {{ template "second_page_variables" }}, "domain": { "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx" } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := APaginatedTestClient(t, "domain/tags", requests...)
	domain := ol.DomainId{
		Id: "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx",
	}
	// Act
	resp, err := domain.Tags(client, nil)
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

func TestDomainAssignSystem(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation DomainAssignSystem($childSystems:[IdentifierInput!]!$domain:IdentifierInput!){domainChildAssign(domain:$domain, childSystems:$childSystems){domain{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},errors{message,path}}}",
	"variables":{
		"domain":{
			"id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx"
	  	},
	  	"childSystems": [
			{"id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy"}
	  	]
	}
}`
	response := `{"data": {
		"domainChildAssign": {
			"domain": {{ template "domain1_response" }}
		}
}}`
	client := ABetterTestClient(t, "domain/assign_system", request, response)
	// Act
	domain := ol.Domain{
		DomainId: ol.DomainId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUx",
		},
	}
	childSystems := "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy"
	err := domain.AssignSystem(client, childSystems)
	// Assert
	autopilot.Ok(t, err)
}

func TestDomainGetId(t *testing.T) {
	// Arrange
	request := `{
    "query": "query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note}}}",
	"variables":{
		"input": {
			"id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw"
    	}
	}
}`
	response := `{"data": {
		"account": {
			"domain": {{ template "domain1_response" }}
		}
}}`
	client := ABetterTestClient(t, "domain/get_id", request, response)
	// Act
	result, err := client.GetDomain("Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw", string(result.Id))
}

func TestDomainGetAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note}}}",
	"variables":{
		"input": {
			"alias": "my-domain"
    }
}}`
	response := `{"data": {
		"account": {
			"domain": {{ template "domain1_response" }}
		}
}}`
	client := ABetterTestClient(t, "domain/get_alias", request, response)
	// Act
	result, err := client.GetDomain("my-domain")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw", string(result.Id))
}

func TestDomainList(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "domains": { "nodes": [ {{ template "domain1_response" }}, {{ template "domain2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "domains": { "nodes": [ {{ template "domain3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := APaginatedTestClient(t, "domain/list", requests...)
	// Act
	response, err := client.ListDomains(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "PlatformDomain1", result[0].Name)
	autopilot.Equals(t, "PlatformDomain2", result[1].Name)
	autopilot.Equals(t, "PlatformDomain3", result[2].Name)
}

func TestDomainUpdate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation DomainUpdate($domain:IdentifierInput!$input:DomainInput!){domainUpdate(domain:$domain,input:$input){domain{id,aliases,name,description,htmlUrl,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},note},errors{message,path}}}",
	"variables":{
	   "domain":{
		    "id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw"
	    },
		"input":{
			"name": "platform-test-4",
			"description":"Domain created for testing.",
			"ownerId":"Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU",
			"note": "Please delete me"
		}
    }
}`
	response := `{"data": {
		"domainUpdate": {
			"domain": {{ template "domain1_response" }}
		}
}}`
	client := ABetterTestClient(t, "domain/update", request, response)
	input := ol.DomainInput{
		Name:        ol.NewString("platform-test-4"),
		Description: ol.NewString("Domain created for testing."),
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        ol.NewString("Please delete me"),
	}
	// Act
	result, err := client.UpdateDomain("Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw", input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw", string(result.Id))
	autopilot.Equals(t, "An example description", result.Note)
}

func TestDomainDelete(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation DomainDelete($input:IdentifierInput!){domainDelete(resource: $input){errors{message,path}}}",
	"variables":{"input":{"alias":"platformdomain3"}}
	}`
	response := `{"data": {
		"domainDelete": {
      "errors": []
    }
}}`
	client := ABetterTestClient(t, "domain/delete", request, response)
	// Act
	err := client.DeleteDomain("platformdomain3")
	// Assert
	autopilot.Ok(t, err)
}
