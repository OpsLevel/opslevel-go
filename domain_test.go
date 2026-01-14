package opslevel_test

import (
	"strings"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestDomainCreate(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.DomainInput]("domain_create_input",
		ol.DomainInput{
			Name:        ol.RefOf("platform-test"),
			Description: ol.RefOf("Domain created for testing."),
			OwnerId:     ol.RefOf(id1),
			Note:        ol.RefOf("additional note about platform-test domain"),
		})

	testRequest := autopilot.NewTestRequest(
		`mutation DomainCreate($input:DomainInput!){domainCreate(input:$input){domain{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}},errors{message,path}}}`,
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
		`query DomainChildSystemsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){childSystems(after: $after, first: $first){nodes{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}`,
		`{ {{ template "first_page_variables" }}, "domain": { {{ template "id2" }} } }`,
		`{ "data": { "account": { "domain": { "childSystems": { "nodes": [ {{ template "system1_response" }}, {{ template "system2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query DomainChildSystemsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){childSystems(after: $after, first: $first){nodes{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}`,
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

func TestDomainReconcileAliasesDeleteAll(t *testing.T) {
	// Arrange
	aliasesWanted := []string{}
	domain := ol.Domain{
		DomainId: ol.DomainId{
			Id:      id1,
			Aliases: []string{"one", "two"},
		},
		ManagedAliases: []string{"one", "two"},
	}

	// delete "one" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "one", "ownerType": "domain" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "two" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerType": "domain" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// get domain
	testRequestThree := autopilot.NewTestRequest(
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": {"account": {"domain": { {{ template "id1" }}, "aliases": [], "managedAliases": [] } }}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree}
	client := BestTestClient(t, "domain/reconcile_aliases_delete_all", requests...)

	// Act
	err := domain.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, len(domain.Aliases), 0)
	autopilot.Equals(t, len(domain.ManagedAliases), 0)
}

func TestDomainReconcileAliases(t *testing.T) {
	// Arrange
	aliasesWanted := []string{"one", "two", "three"}
	domain := ol.Domain{
		DomainId: ol.DomainId{
			Id:      id1,
			Aliases: []string{"one", "alpha", "beta"},
		},
		ManagedAliases: []string{"one", "alpha", "beta"},
	}

	// delete "alpha" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "alpha", "ownerType": "domain" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "beta" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "beta", "ownerType": "domain" }}`,
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
	// get domain
	testRequestFive := autopilot.NewTestRequest(
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": {"account": {"domain": { {{ template "id1" }}, "aliases": ["one", "two", "three"], "managedAliases": ["one", "two", "three"] } }}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree, testRequestFour, testRequestFive}
	client := BestTestClient(t, "domain/reconcile_aliases", requests...)

	// Act
	err := domain.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, domain.Aliases, aliasesWanted)
	autopilot.Equals(t, domain.ManagedAliases, aliasesWanted)
}

func TestDomainGetTags(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "domain": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag1" }}, {{ template "tag2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query DomainTagsList($after:String!$domain:IdentifierInput!$first:Int!){account{domain(input: $domain){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "domain": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "domain": { "tags": { "nodes": [ {{ template "tag3" }} ], {{ template "pagination_second_pageInfo_response" }} } }}}}`,
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
		`mutation DomainAssignSystem($childSystems:[IdentifierInput!]!$domain:IdentifierInput!){domainChildAssign(domain:$domain, childSystems:$childSystems){domain{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}},errors{message,path}}}`,
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
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}`,
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
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}`,
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
		`query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "domains": { "nodes": [ {{ template "domain1_response" }}, {{ template "domain2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
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
			OwnerId:     ol.RefOf(id3),
			Note:        ol.RefOf("Please delete me"),
		})

	testRequest := autopilot.NewTestRequest(
		`mutation DomainUpdate($domain:IdentifierInput!$input:DomainInput!){domainUpdate(domain:$domain,input:$input){domain{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}},errors{message,path}}}`,
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

func TestDomainGetWithHTMLEntities(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query DomainGet($input:IdentifierInput!){account{domain(input: $input){id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}`,
		`{"input": { {{ template "id1" }} } }`,
		`{"data": {"account": {"domain": {
			{{ template "id1" }},
			"aliases": ["test-domain"],
			"name": "TestDomain",
			"description": "A domain with &lt;html&gt; &amp; special characters &quot;quoted&quot;",
			"htmlUrl": "https://app.opslevel.com/catalog/domains/test-domain",
			"note": "Note with &lt;b&gt;bold&lt;/b&gt; and &amp; ampersand",
			"managedAliases": []
		}}}}`,
	)

	client := BestTestClient(t, "domain/html_entities", testRequest)
	// Act
	result, err := client.GetDomain(string(id1))
	// Assert
	autopilot.Ok(t, err)

	// Verify HTML entities are unescaped
	if strings.Contains(result.Description, "&lt;") || strings.Contains(result.Description, "&gt;") ||
		strings.Contains(result.Description, "&amp;") || strings.Contains(result.Description, "&quot;") {
		t.Errorf("Domain Description still contains HTML entities: %s", result.Description)
	}
	if strings.Contains(result.Note, "&lt;") || strings.Contains(result.Note, "&gt;") ||
		strings.Contains(result.Note, "&amp;") {
		t.Errorf("Domain Note still contains HTML entities: %s", result.Note)
	}

	// Verify expected unescaped values
	expectedDescription := `A domain with <html> & special characters "quoted"`
	expectedNote := "Note with <b>bold</b> and & ampersand"

	autopilot.Equals(t, expectedDescription, result.Description)
	autopilot.Equals(t, expectedNote, result.Note)
}
