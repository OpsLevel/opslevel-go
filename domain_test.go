package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestDomainCreate(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.DomainInput]("domain_create_input",
		ol.DomainInput{
			Name:        ol.RefOf("platform-test"),
			Description: ol.RefOf("Domain created for testing."),
			OwnerId:     &id1,
			Note:        ol.RefOf("additional note about platform-test domain"),
		})

	testRequest := autopilot.NewTestRequest(
		`mutation DomainCreate($input:DomainInput!){domainCreate(input:$input){domain{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},errors{message,path}}}`,
		`{"input": {{ template "domain_create_input" }} }`,
		`{"data": {"domainCreate": {"domain": {{ template "domain1_response" }} }}}`,
	)

	client := BestTestClient(t, "domain/create", testRequest)
	// Act
	result, err := client.CreateDomain(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "An example description", result.Note)
}

func TestDomainGetSystems(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query DomainChildSystemsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){childSystems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},note},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "domain": { {{ template "id2" }} } }`,
		`{ "data": { "account": { "domain": { "childSystems": { "nodes": [ {{ template "system1_response" }}, {{ template "system2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query DomainChildSystemsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){childSystems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},note},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "domain": { {{ template "id2" }} } }`,
		`{ "data": { "account": { "domain": { "childSystems": { "nodes": [ {{ template "system3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "domain/child_systems", requests...)
	domain := ol.DomainId{
		Id: id2,
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
	testRequestOne := autopilot.NewTestRequest(
		`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "domain": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag1" }}, {{ template "tag2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "domain": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "domain/tags", requests...)
	domain := ol.DomainId{
		Id: id1,
	}
	// Act
	resp, err := domain.GetTags(client, nil)
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
	testRequest := autopilot.NewTestRequest(
		`mutation DomainAssignSystem($childSystems:[IdentifierInput!]!$domain:IdentifierInput!){domainChildAssign(domain:$domain, childSystems:$childSystems){domain{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},errors{message,path}}}`,
		`{"domain":{ {{ template "id1" }} }, "childSystems": [ { {{ template "id3" }} } ] }`,
		`{"data": {"domainChildAssign": {"domain": {{ template "domain1_response" }} }}}`,
	)

	client := BestTestClient(t, "domain/assign_system", testRequest)
	// Act
	domain := ol.Domain{
		DomainId: ol.DomainId{
			Id: id1,
		},
	}
	err := domain.AssignSystem(client, string(id3))
	// Assert
	autopilot.Ok(t, err)
}

func TestDomainGetId(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": {"account": {"domain": {{ template "domain1_response" }} }}}`,
	)

	client := BestTestClient(t, "domain/get_id", testRequest)
	// Act
	result, err := client.GetDomain(string(id1))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestDomainGetAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note}}}`,
		`{"input": {"alias": "my-domain" }}`,
		`{"data": {"account": {"domain": {{ template "domain1_response" }} }}}`,
	)

	client := BestTestClient(t, "domain/get_alias", testRequest)
	// Act
	result, err := client.GetDomain("my-domain")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestDomainList(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "domains": { "nodes": [ {{ template "domain1_response" }}, {{ template "domain2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "domains": { "nodes": [ {{ template "domain3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "domain/list", requests...)
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
	input := autopilot.Register[ol.DomainInput]("domain_update_input",
		ol.DomainInput{
			Name:        ol.RefOf("platform-test-4"),
			Description: ol.RefOf("Domain created for testing."),
			OwnerId:     &id3,
			Note:        ol.RefOf("Please delete me"),
		})

	testRequest := autopilot.NewTestRequest(
		`mutation DomainUpdate($domain:IdentifierInput!$input:DomainInput!){domainUpdate(domain:$domain,input:$input){domain{id,aliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},note},errors{message,path}}}`,
		`{"domain": { {{ template "id1" }} }, "input": {{ template "domain_update_input" }} }`,
		`{"data": {"domainUpdate": {"domain": {{ template "domain1_response" }} }}}`,
	)

	client := BestTestClient(t, "domain/update", testRequest)
	// Act
	result, err := client.UpdateDomain(string(id1), input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "An example description", result.Note)
}

func TestDomainDelete(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation DomainDelete($input:IdentifierInput!){domainDelete(resource: $input){errors{message,path}}}`,
		`{"input":{"alias":"platformdomain3"}}`,
		`{"data": {"domainDelete": {"errors": [] }}}`,
	)

	client := BestTestClient(t, "domain/delete", testRequest)
	// Act
	err := client.DeleteDomain("platformdomain3")
	// Assert
	autopilot.Ok(t, err)
}
