package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestComponentTypeCreate(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.ComponentTypeInput]("component_type_create_input",
		ol.ComponentTypeInput{
			Alias:       ol.RefOf("example"),
			Name:        ol.RefOf("Example"),
			Description: ol.RefOf("Example Description"),
			Properties:  &[]ol.ComponentTypePropertyDefinitionInput{},
			OwnerRelationship: &ol.OwnerRelationshipInput{
				ManagementRules: &[]ol.ManagementRuleInput{
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperatorEquals,
						SourceProperty:        "tag_key_eq:owner",
						SourcePropertyBuiltin: true,
						TargetProperty:        "name",
						TargetPropertyBuiltin: true,
						TargetType:            ol.NewNullableFrom("team"),
					},
				},
			},
		})

	testRequest := autopilot.NewTestRequest(
		`mutation ComponentTypeCreate($input:ComponentTypeInput!){componentTypeCreate(input:$input){componentType{{ template "component_type_graphql" }},errors{message,path}}}`,
		`{"input": {"alias": "example", "name": "Example", "description": "Example Description", "properties": [], "ownerRelationship": {"managementRules": [{"operator": "EQUALS", "sourceProperty": "tag_key_eq:owner", "sourcePropertyBuiltin": true, "targetProperty": "name", "targetPropertyBuiltin": true, "targetType": "team"}]} }}`,
		`{"data": {"componentTypeCreate": {"componentType": {{ template "component_type_1_response" }} }}}`,
	)

	client := BestTestClient(t, "ComponentType/create", testRequest)
	// Act
	result, err := client.CreateComponentType(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestComponentTypeGet(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ComponentTypeGet($input:IdentifierInput!){account{componentType(input: $input){{ template "component_type_graphql" }}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": {"account": {"componentType": {{ template "component_type_1_response" }} }}}`,
	)

	client := BestTestClient(t, "ComponentType/get", testRequest)
	// Act
	result, err := client.GetComponentType(string(id1))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestComponentTypeList(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ComponentTypeList($after:String!$first:Int!){account{componentTypes(after: $after, first: $first){nodes{{ template "component_type_graphql" }},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "componentTypes": { "nodes": [ {{ template "component_type_1_response" }}, {{ template "component_type_2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ComponentTypeList($after:String!$first:Int!){account{componentTypes(after: $after, first: $first){nodes{{ template "component_type_graphql" }},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "componentTypes": { "nodes": [ {{ template "component_type_3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "ComponentType/list", requests...)
	// Act
	response, err := client.ListComponentTypes(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Example1", result[0].Name)
	autopilot.Equals(t, "Example2", result[1].Name)
	autopilot.Equals(t, "Example3", result[2].Name)
}

func TestComponentTypeUpdate(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.ComponentTypeInput]("component_type_update_input",
		ol.ComponentTypeInput{
			Alias:       ol.RefOf("example"),
			Name:        ol.RefOf("Example"),
			Description: ol.RefOf("Example Description"),
			Properties:  &[]ol.ComponentTypePropertyDefinitionInput{},
			OwnerRelationship: &ol.OwnerRelationshipInput{
				ManagementRules: &[]ol.ManagementRuleInput{
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperatorEquals,
						SourceProperty:        "tag_key_eq:owner",
						SourcePropertyBuiltin: true,
						TargetProperty:        "name",
						TargetPropertyBuiltin: true,
						TargetType:            ol.NewNullableFrom("team"),
					},
				},
			},
		})

	testRequest := autopilot.NewTestRequest(
		`mutation ComponentTypeUpdate($input:ComponentTypeInput!$target:IdentifierInput!){componentTypeUpdate(componentType:$target,input:$input){componentType{{ template "component_type_graphql" }},errors{message,path}}}`,
		`{"input": {"alias": "example", "name": "Example", "description": "Example Description", "properties": [], "ownerRelationship": {"managementRules": [{"operator": "EQUALS", "sourceProperty": "tag_key_eq:owner", "sourcePropertyBuiltin": true, "targetProperty": "name", "targetPropertyBuiltin": true, "targetType": "team"}]}}, "target": { {{ template "id1" }} }}`,
		`{"data": {"componentTypeUpdate": {"componentType": {{ template "component_type_1_response" }} }}}`,
	)

	client := BestTestClient(t, "ComponentType/update", testRequest)
	// Act
	result, err := client.UpdateComponentType(string(id1), input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
}

func TestComponentTypeDelete(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ComponentTypeDelete($target:IdentifierInput!){componentTypeDelete(resource:$target){errors{message,path}}}`,
		`{"target": { {{ template "id1" }} }}`,
		`{"data": {"componentTypeDelete": {"errors": [] }}}`,
	)

	client := BestTestClient(t, "ComponentType/delete", testRequest)
	// Act
	err := client.DeleteComponentType(string(id1))
	// Assert
	autopilot.Ok(t, err)
}
