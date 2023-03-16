package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestDomainCreate(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/create", request, response)
	// Act
	result, err := client.CreateDomain(ol.DomainInput{})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
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
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/get_id", request, response)
	// Act
	result, err := client.GetDomain("MTIzNDU2Nzg5MTIzNDU2Nzg5")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
}

func TestDomainGetAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "domain/get_alias", request, response)
	// Act
	result, err := client.GetDomain("my-domain")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
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
	result, err := client.UpdateDomain("", ol.DomainInput{})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
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
