package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateServiceDependency(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceDependencyCreate($input:ServiceDependencyCreateInput!){serviceDependencyCreate(inputV2: $input){serviceDependency{id,sourceService{id,aliases},destinationService{id,aliases},notes},errors{message,path}}}`,
		`{ "input": { "dependencyKey": { "destinationIdentifier": {"alias": "example_3"}, "sourceIdentifier": {"alias": "example_2"} }, "notes": "An example description" }}`,
		`{"data": { "serviceDependencyCreate": { "serviceDependency": {{ template "serviceDependency" }}, "errors": [] } }}`,
	)
	client := BestTestClient(t, "serviceDependencyCreate", testRequest)
	// Act
	result, err := client.CreateServiceDependency(ol.ServiceDependencyCreateInput{
		DependencyKey: ol.ServiceDependencyKey{
			SourceIdentifier:      ol.NewIdentifier("example_2"),
			DestinationIdentifier: ol.NewIdentifier("example_3"),
		},
		Notes: ol.RefOf("An example description"),
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
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceDependenciesList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependencies(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data": {"account": { "service": { "dependencies": { "edges": [ {{ template "serviceDependencyEdge_1" }}, {{ template "serviceDependencyEdge_2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceDependenciesList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependencies(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data": {"account": { "service": { "dependencies": { "edges": [ {{ template "serviceDependencyEdge_3" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

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
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceDependentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependents(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data": {"account": { "service": { "dependents": { "edges": [ {{ template "serviceDependencyEdge_1" }}, {{ template "serviceDependencyEdge_2" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceDependentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependents(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data": {"account": { "service": { "dependents": { "edges": [ {{ template "serviceDependencyEdge_3" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

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
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceDependencyDelete($input:DeleteInput!){serviceDependencyDelete(input: $input){errors{message,path}}}`,
		`{ "input": { {{ template "id1" }} } }`,
		`{"data": { "serviceDependencyDelete": { "errors": [] }}}`,
	)
	client := BestTestClient(t, "serviceDependencyDelete", testRequest)
	// Act
	err := client.DeleteServiceDependency(string(id1))
	// Assert
	autopilot.Ok(t, err)
}
