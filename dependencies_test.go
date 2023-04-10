package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestCreateServiceDependency(t *testing.T) {
	// Arrange
	request := `{
   "query": "mutation ServiceDependencyCreate($input:ServiceDependencyCreateInput!){serviceDependencyCreate(inputV2: $input){serviceDependency{id,sourceService{id,aliases},destinationService{id,aliases},notes},errors{message,path}}}",
   "variables": {
		"input": {
			"dependencyKey": {
				"destinationIdentifier": {"alias": "example_3"},
				"sourceIdentifier": {"alias": "example_2"}
			},
			"notes": "An example description"
		}
}}`
	response := `{"data": {
		"serviceDependencyCreate": {
			"serviceDependency": {{ template "serviceDependency" }},
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "serviceDependencyCreate", request, response)
	// Act
	result, err := client.CreateServiceDependency(ol.ServiceDependencyCreateInput{
		Key: ol.ServiceDependencyKey{
			Parent: *ol.NewIdentifier("example_2"),
			Child:  *ol.NewIdentifier("example_3"),
		},
		Notes: "An example description",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), result.Id)
	autopilot.Equals(t, ol.ID("Z2lkOi8vOTg3NjU0MzIxMTIzNDU2Nzg5"), result.Parent.Id)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx"), result.Child.Id)
	autopilot.Equals(t, "An example description", result.Notes)
}

func TestGetServiceDependencies(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ServiceDependenciesList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependencies(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"service": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
			}}`,
			`{"data": {"account": {
				"service": {
					"dependencies": {
						"edges": [
							{{ template "serviceDependencyEdge_1" }},
							{{ template "serviceDependencyEdge_2" }}
						],
						{{ template "pagination_initial_pageInfo_response" }}
					}
				}
			}}}`},
		{`{"query": "query ServiceDependenciesList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependencies(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"service": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
			}}`,
			`{"data": {"account": {
				"service": {
					"dependencies": {
						"edges": [
							{{ template "serviceDependencyEdge_3" }}
						],
						{{ template "pagination_second_pageInfo_response" }}
					}
				}
			}}}`},
	}

	client := APaginatedTestClient(t, "service/get_dependencies", requests...)
	// Act
	resource := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		},
	}
	resp, err := resource.GetDependencies(client, nil)
	result := resp.Edges
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), result[0].Id)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx"), result[2].Id)
}

func TestGetServiceDependents(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ServiceDependentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependents(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"service": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
			}}`,
			`{"data": {"account": {
				"service": {
					"dependents": {
						"edges": [
							{{ template "serviceDependencyEdge_1" }},
							{{ template "serviceDependencyEdge_2" }}
						],
						{{ template "pagination_initial_pageInfo_response" }}
					}
				}
			}}}`},
		{`{"query": "query ServiceDependentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){dependents(after: $after, first: $first){edges{id,locked,node{id,aliases},notes},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"service": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
			}}`,
			`{"data": {"account": {
				"service": {
					"dependents": {
						"edges": [
							{{ template "serviceDependencyEdge_3" }}
						],
						{{ template "pagination_second_pageInfo_response" }}
					}
				}
			}}}`},
	}

	client := APaginatedTestClient(t, "service/get_dependents", requests...)
	// Act
	resource := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		},
	}
	resp, err := resource.GetDependents(client, nil)
	result := resp.Edges
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), result[0].Id)
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx"), result[2].Id)
}

func TestDeleteServiceDependency(t *testing.T) {
	// Arrange
	request := `{
   "query": "mutation ServiceDependencyDelete($input:DeleteInput!){serviceDependencyDelete(input: $input){errors{message,path}}}",
   "variables": {
		"input": {
			"id": "{{ template "id1" }}"
		}
}}`
	response := `{"data": {
		"serviceDependencyDelete": {
			"errors": []
		}
}}`
	client := ABetterTestClient(t, "serviceDependencyDelete", request, response)
	// Act
	err := client.DeleteServiceDependency("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Ok(t, err)
}
