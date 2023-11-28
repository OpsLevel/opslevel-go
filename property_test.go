package opslevel_test

import (
	"fmt"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

const schemaString = `{"$ref":"#/$defs/MyProp","$defs":{"MyProp":{"properties":{"name":{"type":"string","title":"the name","description":"The name of a friend","default":"alex","examples":["joe","lucy"]}},"additionalProperties":false,"type":"object","required":["name"]}}}`

func TestCreatePropertyDefinition(t *testing.T) {
	// Arrange
	schema := ol.NewJSON(schemaString)
	propertyDefinition := autopilot.Register[ol.PropertyDefinitionInput]("property_definition_input", ol.PropertyDefinitionInput{
		Name:   "my-prop",
		Schema: schema,
	})
	testRequest := NewTestRequest(
		`"mutation PropertyDefinitionCreate($input:PropertyDefinitionInput!){propertyDefinitionCreate(input: $input){definition{aliases,id,name,schema},errors{message,path}}}"`,
		`{"input": {{ template "property_definition_input" }} }`,
		fmt.Sprintf(`{"data":{"propertyDefinitionCreate":{"definition":{"aliases":["my_prop"],"id":"XXX","name":"my-prop","schema": %s},"errors":[]}}}`, propertyDefinition.Schema.ToJSON()),
	)
	client := BestTestClient(t, "properties/definition_create", testRequest)

	// Act
	property, err := client.CreatePropertyDefinition(propertyDefinition)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, propertyDefinition.Name, property.Name)
	autopilot.Equals(t, propertyDefinition.Schema, property.Schema)
}

func TestDeletePropertyDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation PropertyDefinitionDelete($input:IdentifierInput!){propertyDefinitionDelete(resource: $input){deletedAlias,deletedId,errors{message,path}}}"`,
		`{"input":{"alias":"my_prop"}}`,
		`{"data":{"propertyDefinitionDelete":{"deletedAlias":"my_prop","deletedId":"XXX","errors":[]}}}`,
	)
	client := BestTestClient(t, "properties/definition_delete", testRequest)

	// Act
	err := client.DeletePropertyDefinition("my_prop")

	// Assert
	autopilot.Ok(t, err)
}

func TestGetPropertyDefinition(t *testing.T) {
	// Arrange
	schema := ol.NewJSON(schemaString)
	testRequest := NewTestRequest(
		`"query PropertyDefinitionGet($input:IdentifierInput!){account{propertyDefinition(input: $input){aliases,id,name,schema}}}"`,
		`{"input":{"alias":"my_prop"}}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinition":{"aliases":["my_prop"],"id":"XXX","name":"my-prop","schema": %s}}}}`, schema.ToJSON()),
	)
	client := BestTestClient(t, "properties/definition_get", testRequest)

	// Act
	property, err := client.GetPropertyDefinition("my_prop")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, "my-prop", property.Name)
	autopilot.Equals(t, schema, property.Schema)
}

func TestListPropertyDefinitions(t *testing.T) {
	// Arrange
	schema := ol.NewJSON(schemaString)
	testRequestOne := NewTestRequest(
		`"query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,schema},{{ template "pagination_request" }}}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop1"],"id":"XXX","name":"prop1","schema": %s},{"aliases":["prop2"],"id":"XXX","name":"prop2","schema": %s}],{{ template "pagination_initial_pageInfo_response" }}}}}}`, schema.ToJSON(), schema.ToJSON()),
	)
	testRequestTwo := NewTestRequest(
		`"query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,schema},{{ template "pagination_request" }}}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop3"],"id":"XXX","name":"prop3","schema": %s}],{{ template "pagination_second_pageInfo_response" }}}}}}`, schema.ToJSON()),
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "properties/definition_list", requests...)

	// Act
	properties, err := client.ListPropertyDefinitions(nil)
	result := properties.Nodes

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, "prop1", result[0].Name)
	autopilot.Equals(t, "prop2", result[1].Name)
	autopilot.Equals(t, "prop3", result[2].Name)
}
