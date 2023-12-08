package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

const schemaString = `{"$ref":"#/$defs/MyProp","$defs":{"MyProp":{"properties":{"name":{"type":"string","title":"the name","description":"The name of a friend","default":"alex","examples":["joe","lucy"]}},"additionalProperties":false,"type":"object","required":["name"]}}}`

func TestCreatePropertyDefinition(t *testing.T) {
	// Arrange
	expectedPropertyDefinition := autopilot.Register[ol.PropertyDefinition]("expected_property_definition", ol.PropertyDefinition{
		Aliases: []string{"my_prop"},
		Id:      "XXX",
		Name:    "my-prop",
		Schema:  ol.JSONString(schemaString),
	})
	propertyDefinitionInput := autopilot.Register[ol.PropertyDefinitionInput]("property_definition_input", ol.PropertyDefinitionInput{
		Name:   "my-prop",
		Schema: ol.NewJSON(schemaString),
	})
	testRequest := NewTestRequest(
		`"mutation PropertyDefinitionCreate($input:PropertyDefinitionInput!){propertyDefinitionCreate(input: $input){definition{aliases,id,name,schema},errors{message,path}}}"`,
		`{"input": {{ template "property_definition_input" }} }`,
		`{"data":{"propertyDefinitionCreate":{"definition":{{ template "expected_property_definition" }}, "errors":[] }}}`,
	)
	client := BestTestClient(t, "properties/definition_create", testRequest)

	// Act
	actualPropertyDefinition, err := client.CreatePropertyDefinition(propertyDefinitionInput)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedPropertyDefinition, *actualPropertyDefinition)
	autopilot.Equals(t, propertyDefinitionInput.Name, actualPropertyDefinition.Name)
	autopilot.Equals(t, propertyDefinitionInput.Schema, ol.JSON(actualPropertyDefinition.Schema.Map()))
}

func TestDeletePropertyDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation PropertyDefinitionDelete($input:IdentifierInput!){propertyDefinitionDelete(resource: $input){errors{message,path}}}"`,
		`{"input":{"alias":"my_prop"}}`,
		`{"data":{"propertyDefinitionDelete":{"errors":[]}}}`,
	)
	client := BestTestClient(t, "properties/definition_delete", testRequest)

	// Act
	err := client.DeletePropertyDefinition("my_prop")

	// Assert
	autopilot.Ok(t, err)
}

func TestGetPropertyDefinition(t *testing.T) {
	// Arrange
	expectedPropertyDefinition := autopilot.Register[ol.PropertyDefinition]("expected_property_definition",
		ol.PropertyDefinition{
			Aliases: []string{"my_prop"},
			Id:      "XXX",
			Name:    "my-prop",
			Schema:  ol.JSONString(schemaString),
		})
	testRequest := NewTestRequest(
		`"query PropertyDefinitionGet($input:IdentifierInput!){account{propertyDefinition(input: $input){aliases,id,name,schema}}}"`,
		`{"input":{"alias":"my_prop"}}`,
		`{"data":{"account":{"propertyDefinition": {{ template "expected_property_definition" }} }}}`,
	)
	client := BestTestClient(t, "properties/definition_get", testRequest)

	// Act
	property, err := client.GetPropertyDefinition("my_prop")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, expectedPropertyDefinition, *property)
	autopilot.Equals(t, "my-prop", property.Name)
}

func TestListPropertyDefinitions(t *testing.T) {
	// Arrange
	expectedPropDefsPageOne := autopilot.Register[[]ol.PropertyDefinition]("property_definitions", []ol.PropertyDefinition{
		{
			Aliases: []string{"prop1"},
			Id:      "XXX",
			Name:    "prop1",
			Schema:  ol.JSONString(schemaString),
		},
		{
			Aliases: []string{"prop2"},
			Id:      "XXX",
			Name:    "prop2",
			Schema:  ol.JSONString(schemaString),
		},
	})
	expectedPropDefPageTwo := autopilot.Register[ol.PropertyDefinition]("property_definition_3", ol.PropertyDefinition{
		Aliases: []string{"prop3"},
		Id:      "XXX",
		Name:    "prop3",
		Schema:  ol.JSONString(schemaString),
	})
	testRequestOne := NewTestRequest(
		`"query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,schema},{{ template "pagination_request" }}}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{"data":{"account":{"propertyDefinitions":{"nodes": {{ template "property_definitions" }},{{ template "pagination_initial_pageInfo_response" }}}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,schema},{{ template "pagination_request" }}}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{"data":{"account":{"propertyDefinitions":{"nodes":[{{ template "property_definition_3" }}],{{ template "pagination_second_pageInfo_response" }}}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "properties/definition_list", requests...)

	// Act
	properties, err := client.ListPropertyDefinitions(nil)
	result := properties.Nodes

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, expectedPropDefsPageOne[0], result[0])
	autopilot.Equals(t, expectedPropDefsPageOne[1], result[1])
	autopilot.Equals(t, expectedPropDefPageTwo, result[2])
}
