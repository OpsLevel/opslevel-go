package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

var input ol.RelationshipDefinitionInput
var resp1 ol.RelationshipDefinitionType
var resp2 ol.RelationshipDefinitionType

func init() {
	AddSetup(func(m *testing.M) {
		input = autopilot.Register[ol.RelationshipDefinitionInput]("relationship_definition_input",
			ol.RelationshipDefinitionInput{
				Alias:         &alias1,
				Name:          &name1,
				Description:   ol.RefOf("Example Description"),
				ComponentType: ol.NewIdentifier("example"),
				Metadata: &ol.RelationshipDefinitionMetadataInput{
					AllowedTypes: []string{"example"},
				},
			})
		resp1 = autopilot.Register[ol.RelationshipDefinitionType]("relationship_definition1_response",
			ol.RelationshipDefinitionType{
				Id:    id1,
				Alias: alias1,
				Name:  name1,
				Metadata: ol.RelationshipDefinitionMetadata{
					AllowedTypes: []string{"example"},
				},
			})
		resp2 = autopilot.Register[ol.RelationshipDefinitionType]("relationship_definition2_response",
			ol.RelationshipDefinitionType{
				Id:    id2,
				Alias: alias2,
				Name:  name2,
				Metadata: ol.RelationshipDefinitionMetadata{
					AllowedTypes: []string{"example2", "example3"},
				},
			})
	})
}

func TestRelationshipDefinitionCreate(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RelationshipDefinitionCreate($input:RelationshipDefinitionInput!){relationshipDefinitionCreate(input: $input){definition{alias,componentType{id,aliases},description,id,metadata{allowedTypes,maxItems,minItems},name},errors{message,path}}}`,
		`{"input": {{ template "relationship_definition_input" }} }`,
		`{"data": {"relationshipDefinitionCreate": {"definition": {{ template "relationship_definition1_response" }} }}}`,
	)

	client := BestTestClient(t, "relationship_definition/create", testRequest)
	// Act
	result, err := client.CreateRelationshipDefinition(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, resp1.Id, result.Id)
	autopilot.Equals(t, resp1.Alias, result.Alias)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Metadata.AllowedTypes)
}

func TestRelationshipDefinitionGet(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query RelationshipDefinitionGet($input:IdentifierInput!){account{relationshipDefinition(input: $input){alias,componentType{id,aliases},description,id,metadata{allowedTypes,maxItems,minItems},name}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": {"account": {"relationshipDefinition": {{ template "relationship_definition1_response" }} }}}`,
	)

	client := BestTestClient(t, "relationship_definition/get", testRequest)
	// Act
	result, err := client.GetRelationshipDefinition(string(id1))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, resp1.Id, result.Id)
	autopilot.Equals(t, resp1.Alias, result.Alias)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Metadata.AllowedTypes)
}

func TestRelationshipDefinitionList(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query RelationshipDefinitionList($after:String!$componentType:IdentifierInput$first:Int!$resource:ID){account{relationshipDefinitions(after: $after, first: $first, componentType: $componentType, resource: $resource){nodes{alias,componentType{id,aliases},description,id,metadata{allowedTypes,maxItems,minItems},name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ {{ template "first_page_variables" }}, "componentType": null, "resource": null}`,
		`{ "data": { "account": { "relationshipDefinitions": { "nodes": [ {{ template "relationship_definition1_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query RelationshipDefinitionList($after:String!$componentType:IdentifierInput$first:Int!$resource:ID){account{relationshipDefinitions(after: $after, first: $first, componentType: $componentType, resource: $resource){nodes{alias,componentType{id,aliases},description,id,metadata{allowedTypes,maxItems,minItems},name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ {{ template "second_page_variables" }}, "componentType": null, "resource": null}`,
		`{ "data": { "account": { "relationshipDefinitions": { "nodes": [ {{ template "relationship_definition2_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "relationship_definition/list", requests...)
	// Act
	result, err := client.ListRelationshipDefinitions(nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, result.TotalCount)
	autopilot.Equals(t, resp1.Id, result.Nodes[0].Id)
	autopilot.Equals(t, resp1.Alias, result.Nodes[0].Alias)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Nodes[0].Metadata.AllowedTypes)
	autopilot.Equals(t, resp2.Id, result.Nodes[1].Id)
	autopilot.Equals(t, resp2.Alias, result.Nodes[1].Alias)
	autopilot.Equals(t, resp2.Metadata.AllowedTypes, result.Nodes[1].Metadata.AllowedTypes)
}

func TestRelationshipDefinitionUpdate(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RelationshipDefinitionUpdate($identifier:IdentifierInput$input:RelationshipDefinitionInput!){relationshipDefinitionUpdate(relationshipDefinition: $identifier, input: $input){definition{alias,componentType{id,aliases},description,id,metadata{allowedTypes,maxItems,minItems},name},errors{message,path}}}`,
		`{"identifier": { {{ template "id1" }} }, "input": {{ template "relationship_definition_input" }} }`,
		`{"data": {"relationshipDefinitionUpdate": {"definition": {{ template "relationship_definition1_response" }} }}}`,
	)

	client := BestTestClient(t, "relationship_definition/update", testRequest)
	// Act
	result, err := client.UpdateRelationshipDefinition(string(id1), input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, resp1.Id, result.Id)
	autopilot.Equals(t, resp1.Alias, result.Alias)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Metadata.AllowedTypes)
}

func TestRelationshipDefinitionDelete(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RelationshipDefinitionDelete($input:IdentifierInput!){relationshipDefinitionDelete(resource: $input){deletedId,errors{message,path}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": {"relationshipDefinitionDelete": {"deletedId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)

	client := BestTestClient(t, "relationship_definition/delete", testRequest)
	// Act
	_, err := client.DeleteRelationshipDefinition(string(id1))
	// Assert
	autopilot.Ok(t, err)
}
