package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateServiceDependency(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation ServiceDependencyCreate($input:ServiceDependencyCreateInput!){serviceDependencyCreate(inputV2: $input){serviceDependency{id,sourceService{id,aliases},destinationService{id,aliases},notes},errors{message,path}}}"`,
		`{ "input": { "dependencyKey": { "destinationIdentifier": {"alias": "example_3"}, "sourceIdentifier": {"alias": "example_2"} }, "notes": "An example description" }}`,
		`{"data": { "serviceDependencyCreate": { "serviceDependency": {{ template "serviceDependency" }}, "errors": [] } }}`,
	)
	client := BestTestClient(t, "serviceDependencyCreate", testRequest)
	// Act
	result, err := client.CreateServiceDependency(ol.ServiceDependencyCreateInput{
		Key: ol.ServiceDependencyKey{
			Service:   *ol.NewIdentifier("example_2"),
			DependsOn: *ol.NewIdentifier("example_3"),
		},
		Notes: "An example description",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, id2, result.Service.Id)
	autopilot.Equals(t, id3, result.DependsOn.Id)
	autopilot.Equals(t, "An example description", result.Notes)
}

func TestGetServiceDependencies(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query ServiceDependenciesList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependencies(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}"`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1" }}" }`,
		`{"data": {"account": { "service": { "dependencies": { "edges": [ {{ template "serviceDependencyEdge_1" }}, {{ template "serviceDependencyEdge_2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query ServiceDependenciesList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependencies(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}"`,
		`{ {{ template "second_page_variables" }}, "service": "{{ template "id1" }}" }`,
		`{"data": {"account": { "service": { "dependencies": { "edges": [ {{ template "serviceDependencyEdge_3" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/get_dependencies", requests...)
	// Act
	resource := ol.Service{
		ServiceId: ol.ServiceId{
			Id: id1,
		},
	}
	resp, err := resource.GetDependencies(client, nil)
	result := resp.Edges
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result[0].Id)
	autopilot.Equals(t, id3, result[2].Id)
}

func TestGetServiceDependents(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query ServiceDependentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependents(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}"`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1" }}" }`,
		`{"data": {"account": { "service": { "dependents": { "edges": [ {{ template "serviceDependencyEdge_1" }}, {{ template "serviceDependencyEdge_2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query ServiceDependentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependents(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}"`,
		`{ {{ template "second_page_variables" }}, "service": "{{ template "id1" }}" }`,
		`{"data": {"account": { "service": { "dependents": { "edges": [ {{ template "serviceDependencyEdge_3" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/get_dependents", requests...)
	// Act
	resource := ol.Service{
		ServiceId: ol.ServiceId{
			Id: id1,
		},
	}
	resp, err := resource.GetDependents(client, nil)
	result := resp.Edges
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result[0].Id)
	autopilot.Equals(t, id3, result[2].Id)
}

func TestDeleteServiceDependency(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation ServiceDependencyDelete($input:DeleteInput!){serviceDependencyDelete(input: $input){errors{message,path}}}"`,
		`{ "input": { "id": "{{ template "id1" }}" } }`,
		`{"data": { "serviceDependencyDelete": { "errors": [] }}}`,
	)
	client := BestTestClient(t, "serviceDependencyDelete", testRequest)
	// Act
	err := client.DeleteServiceDependency(id1)
	// Assert
	autopilot.Ok(t, err)
}
