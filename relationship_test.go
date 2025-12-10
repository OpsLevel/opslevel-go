package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

var (
	input ol.RelationshipDefinitionInput
	resp1 ol.RelationshipDefinitionType
	resp2 ol.RelationshipDefinitionType
)

func init() {
	AddSetup(func(m *testing.M) {
		input = autopilot.Register[ol.RelationshipDefinitionInput]("relationship_definition_input",
			ol.RelationshipDefinitionInput{
				Alias:         &alias1,
				Name:          &name1,
				Description:   ol.RefOf("Example Description"),
				ComponentType: ol.NewIdentifier("example"),
				ManagementRules: &[]ol.ManagementRuleInput{
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperator(ol.RelationshipDefinitionManagementRuleOperatorEquals),
						SourceProperty:        "owner",
						SourcePropertyBuiltin: false,
						TargetProperty:        "team",
						TargetPropertyBuiltin: false,
						TargetCategory:        ol.RefOf("infrastructure"),
					},
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperator(ol.RelationshipDefinitionManagementRuleOperatorEquals),
						SourceProperty:        "name",
						SourcePropertyBuiltin: true,
						TargetProperty:        "alias",
						TargetPropertyBuiltin: true,
						TargetType:            ol.RefOf("service"),
					},
				},
				Metadata: &ol.RelationshipDefinitionMetadataInput{
					AllowedCategories: []string{"infrastructure"},
					AllowedTypes:      []string{"service"},
				},
			})
		resp1 = autopilot.Register[ol.RelationshipDefinitionType]("relationship_definition1_response",
			ol.RelationshipDefinitionType{
				Id:    id1,
				Alias: alias1,
				Name:  name1,
				ManagementRules: []ol.RelationshipDefinitionManagementRule{
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperator(ol.RelationshipDefinitionManagementRuleOperatorEquals),
						SourceProperty:        "owner",
						SourcePropertyBuiltin: false,
						TargetProperty:        "team",
						TargetPropertyBuiltin: false,
						TargetCategory:        ol.RefOf("infrastructure"),
					},
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperator(ol.RelationshipDefinitionManagementRuleOperatorEquals),
						SourceProperty:        "name",
						SourcePropertyBuiltin: true,
						TargetProperty:        "alias",
						TargetPropertyBuiltin: true,
						TargetType:            ol.RefOf("service"),
					},
				},
				Metadata: ol.RelationshipDefinitionMetadata{
					AllowedCategories: []string{"infrastructure"},
					AllowedTypes:      []string{"example"},
				},
			})
		resp2 = autopilot.Register[ol.RelationshipDefinitionType]("relationship_definition2_response",
			ol.RelationshipDefinitionType{
				Id:    id2,
				Alias: alias2,
				Name:  name2,
				ManagementRules: []ol.RelationshipDefinitionManagementRule{
					{
						Operator:              ol.RelationshipDefinitionManagementRuleOperator(ol.RelationshipDefinitionManagementRuleOperatorEquals),
						SourceProperty:        "owner",
						SourcePropertyBuiltin: false,
						TargetProperty:        "team",
						TargetPropertyBuiltin: false,
						TargetCategory:        ol.RefOf("infrastructure"),
					},
				},
				Metadata: ol.RelationshipDefinitionMetadata{
					AllowedCategories: []string{"infrastructure"},
					AllowedTypes:      []string{"example2", "example3"},
				},
			})
	})
}

func TestRelationshipDefinitionCreate(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RelationshipDefinitionCreate($input:RelationshipDefinitionInput!){relationshipDefinitionCreate(input: $input){definition{alias,componentType{id,aliases},description,id,managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType},metadata{allowedCategories,allowedTypes,maxItems,minItems},name},errors{message,path}}}`,
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
	autopilot.Equals(t, resp1.Metadata.AllowedCategories, result.Metadata.AllowedCategories)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Metadata.AllowedTypes)
	autopilot.Equals(t, len(resp1.ManagementRules), len(result.ManagementRules))

	autopilot.Equals(t, resp1.ManagementRules[0].Operator, result.ManagementRules[0].Operator)
	autopilot.Equals(t, resp1.ManagementRules[0].SourceProperty, result.ManagementRules[0].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].SourcePropertyBuiltin, result.ManagementRules[0].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetProperty, result.ManagementRules[0].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetPropertyBuiltin, result.ManagementRules[0].TargetPropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetCategory, result.ManagementRules[0].TargetCategory)

	autopilot.Equals(t, resp1.ManagementRules[1].Operator, result.ManagementRules[1].Operator)
	autopilot.Equals(t, resp1.ManagementRules[1].SourceProperty, result.ManagementRules[1].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[1].SourcePropertyBuiltin, result.ManagementRules[1].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetProperty, result.ManagementRules[1].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetPropertyBuiltin, result.ManagementRules[1].TargetPropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetType, result.ManagementRules[1].TargetType)
}

func TestRelationshipDefinitionGet(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query RelationshipDefinitionGet($input:IdentifierInput!){account{relationshipDefinition(input: $input){alias,componentType{id,aliases},description,id,managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType},metadata{allowedCategories,allowedTypes,maxItems,minItems},name}}}`,
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
	autopilot.Equals(t, resp1.Metadata.AllowedCategories, result.Metadata.AllowedCategories)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Metadata.AllowedTypes)
	autopilot.Equals(t, len(resp1.ManagementRules), len(result.ManagementRules))

	autopilot.Equals(t, resp1.ManagementRules[0].Operator, result.ManagementRules[0].Operator)
	autopilot.Equals(t, resp1.ManagementRules[0].SourceProperty, result.ManagementRules[0].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].SourcePropertyBuiltin, result.ManagementRules[0].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetProperty, result.ManagementRules[0].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetPropertyBuiltin, result.ManagementRules[0].TargetPropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetCategory, result.ManagementRules[0].TargetCategory)

	autopilot.Equals(t, resp1.ManagementRules[1].Operator, result.ManagementRules[1].Operator)
	autopilot.Equals(t, resp1.ManagementRules[1].SourceProperty, result.ManagementRules[1].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[1].SourcePropertyBuiltin, result.ManagementRules[1].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetProperty, result.ManagementRules[1].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetPropertyBuiltin, result.ManagementRules[1].TargetPropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetType, result.ManagementRules[1].TargetType)
}

func TestRelationshipDefinitionList(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query RelationshipDefinitionList($after:String!$componentType:IdentifierInput$first:Int!$resource:ID){account{relationshipDefinitions(after: $after, first: $first, componentType: $componentType, resource: $resource){nodes{alias,componentType{id,aliases},description,id,managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType},metadata{allowedCategories,allowedTypes,maxItems,minItems},name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ {{ template "first_page_variables" }}, "componentType": null, "resource": null}`,
		`{ "data": { "account": { "relationshipDefinitions": { "nodes": [ {{ template "relationship_definition1_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query RelationshipDefinitionList($after:String!$componentType:IdentifierInput$first:Int!$resource:ID){account{relationshipDefinitions(after: $after, first: $first, componentType: $componentType, resource: $resource){nodes{alias,componentType{id,aliases},description,id,managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType},metadata{allowedCategories,allowedTypes,maxItems,minItems},name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
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
	autopilot.Equals(t, resp1.Metadata.AllowedCategories, result.Nodes[0].Metadata.AllowedCategories)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Nodes[0].Metadata.AllowedTypes)
	autopilot.Equals(t, resp2.Id, result.Nodes[1].Id)
	autopilot.Equals(t, resp2.Alias, result.Nodes[1].Alias)
	autopilot.Equals(t, resp2.Metadata.AllowedCategories, result.Nodes[1].Metadata.AllowedCategories)
	autopilot.Equals(t, resp2.Metadata.AllowedTypes, result.Nodes[1].Metadata.AllowedTypes)

	autopilot.Equals(t, len(resp1.ManagementRules), len(result.Nodes[0].ManagementRules))

	autopilot.Equals(t, resp1.ManagementRules[0].Operator, result.Nodes[0].ManagementRules[0].Operator)
	autopilot.Equals(t, resp1.ManagementRules[0].SourceProperty, result.Nodes[0].ManagementRules[0].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].SourcePropertyBuiltin, result.Nodes[0].ManagementRules[0].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetProperty, result.Nodes[0].ManagementRules[0].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetPropertyBuiltin, result.Nodes[0].ManagementRules[0].TargetPropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetCategory, result.Nodes[0].ManagementRules[0].TargetCategory)

	autopilot.Equals(t, resp1.ManagementRules[1].Operator, result.Nodes[0].ManagementRules[1].Operator)
	autopilot.Equals(t, resp1.ManagementRules[1].SourceProperty, result.Nodes[0].ManagementRules[1].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[1].SourcePropertyBuiltin, result.Nodes[0].ManagementRules[1].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetProperty, result.Nodes[0].ManagementRules[1].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetPropertyBuiltin, result.Nodes[0].ManagementRules[1].TargetPropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[1].TargetType, result.Nodes[0].ManagementRules[1].TargetType)

	autopilot.Equals(t, len(resp2.ManagementRules), len(result.Nodes[1].ManagementRules))

	autopilot.Equals(t, resp2.ManagementRules[0].Operator, result.Nodes[1].ManagementRules[0].Operator)
	autopilot.Equals(t, resp2.ManagementRules[0].SourceProperty, result.Nodes[1].ManagementRules[0].SourceProperty)
	autopilot.Equals(t, resp2.ManagementRules[0].SourcePropertyBuiltin, result.Nodes[1].ManagementRules[0].SourcePropertyBuiltin)
	autopilot.Equals(t, resp2.ManagementRules[0].TargetProperty, result.Nodes[1].ManagementRules[0].TargetProperty)
	autopilot.Equals(t, resp2.ManagementRules[0].TargetPropertyBuiltin, result.Nodes[1].ManagementRules[0].TargetPropertyBuiltin)
	autopilot.Equals(t, resp2.ManagementRules[0].TargetCategory, result.Nodes[1].ManagementRules[0].TargetCategory)
}

func TestRelationshipDefinitionUpdate(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RelationshipDefinitionUpdate($identifier:IdentifierInput!$input:RelationshipDefinitionInput!){relationshipDefinitionUpdate(relationshipDefinition: $identifier, input: $input){definition{alias,componentType{id,aliases},description,id,managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType},metadata{allowedCategories,allowedTypes,maxItems,minItems},name},errors{message,path}}}`,
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
	autopilot.Equals(t, resp1.Metadata.AllowedCategories, result.Metadata.AllowedCategories)
	autopilot.Equals(t, resp1.Metadata.AllowedTypes, result.Metadata.AllowedTypes)

	autopilot.Equals(t, len(resp1.ManagementRules), len(result.ManagementRules))

	autopilot.Equals(t, resp1.ManagementRules[0].Operator, result.ManagementRules[0].Operator)
	autopilot.Equals(t, resp1.ManagementRules[0].SourceProperty, result.ManagementRules[0].SourceProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].SourcePropertyBuiltin, result.ManagementRules[0].SourcePropertyBuiltin)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetProperty, result.ManagementRules[0].TargetProperty)
	autopilot.Equals(t, resp1.ManagementRules[0].TargetPropertyBuiltin, result.ManagementRules[0].TargetPropertyBuiltin)
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

func TestGetRelationship(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query RelationshipGet($input:ID!){account{relationship(id: $input){id,source{... on Domain{id,aliases},... on InfrastructureResource{id,aliases,name},... on Service{id,aliases},... on System{id,aliases},... on Team{alias,id}},target{... on Domain{id,aliases},... on InfrastructureResource{id,aliases,name},... on Service{id,aliases},... on System{id,aliases},... on Team{alias,id}},type}}}`,
		`{"input": "{{ template "id1_string" }}" }`,
		`{"data": {"account": {"relationship": {
			"id": "{{ template "id1_string" }}",
			"source": {
				"id": "{{ template "id2_string" }}",
				"aliases": ["source-alias"]
			},
			"target": {
				"id": "{{ template "id3_string" }}",
				"aliases": ["target-alias"]
			},
			"type": "belongs_to"
		}}}}`,
	)

	client := BestTestClient(t, "relationship/get", testRequest)
	// Act
	result, err := client.GetRelationship(string(id1))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, id2, result.Source.Domain.Id)
	autopilot.Equals(t, id3, result.Target.Domain.Id)
	autopilot.Equals(t, ol.RelationshipTypeEnumBelongsTo, result.Type)
}
