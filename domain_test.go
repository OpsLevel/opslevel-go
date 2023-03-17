package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestDomainCreate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation DomainCreate($input:DomainCreateInput!){domainCreate(input:$input){domain{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}},errors{message,path}}}",
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
	input := ol.DomainCreateInput{
		Name:        "platform-test",
		Description: "Domain created for testing.",
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        "additional note about platform-test domain",
	}
	result, err := client.CreateDomain(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMw", string(result.Id))
}

func TestDomainAssignSystem(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{
}
}`
	response := `{"data": {
}}`
	client := ABetterTestClient(t, "domain/assign_system", request, response)
	domain := ol.DomainId{
		Id: "",
	}
	// Act
	err := domain.AssignSystem(client, "", "")
	// Assert
	autopilot.Ok(t, err)
}

func TestDomainGetId(t *testing.T) {
	// Arrange
	request := `{
    "query": "query DomainGet($input:IdentifierInput){account{domain(input: $input){id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}}}}",
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
    "query": "query DomainGet($input:IdentifierInput){account{domain(input: $input){id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}}}}",
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

func TestDomainGetSystems(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/get_systems", request, response)
	domain := ol.DomainId{
		Id: "",
	}
	// Act
	result, err := domain.ChildSystems(client, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, result.TotalCount)
}

func TestDomainGetTags(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/get_tags", request, response)
	domain := ol.DomainId{
		Id: "",
	}
	// Act
	result, err := domain.Tags(client, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, result.TotalCount)
}

func TestDomainList(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"domains": {
							"nodes": [
								{{ template "domain1_response" }},
								{{ template "domain2_response" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
						  }}}}`},
		{`{"query": "query DomainsList($after:String!$first:Int!){account{domains(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"domains": {
							"nodes": [
								{{ template "domain3_response" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }}}}`},
	}

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
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/update", request, response)
	// Act
	result, err := client.UpdateDomain("", ol.DomainUpdateInput{})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
}

func TestDomainDelete(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/delete", request, response)
	// Act
	err := client.DeleteDomain("123456789")
	// Assert
	autopilot.Ok(t, err)
}
