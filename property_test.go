package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

// TODO: template variable for reused schema
// TODO: template variable for ID's
// TODO: template variable for prop1,prop2,prop3

func TestCreatePropertyDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation PropertyDefinitionCreate($input:PropertyDefinitionInput!){propertyDefinitionCreate(input: $input){definition{aliases,id,name,schema},errors{message,path}}}"`,
		`{"input":{"name":"my-prop","schema":"{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}"}}`,
		`{"data":{"propertyDefinitionCreate":{"definition":{"aliases":["my_prop"],"id":"XXX","name":"my-prop","schema":"{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}"},"errors":[]}}}`,
	)
	client := BestTestClient(t, "properties/definition_create", testRequest)

	// Act
	property, err := client.CreatePropertyDefinition(ol.PropertyDefinitionInput{
		Name:   "my-prop",
		Schema: "{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, "my-prop", property.Name)
	autopilot.Equals(t, "{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}", property.Schema)
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
	err := client.DeletePropertyDefinition(*ol.NewIdentifier("my_prop"))

	// Assert
	autopilot.Ok(t, err)
}

func TestGetPropertyDefinition(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query PropertyDefinitionGet($input:IdentifierInput!){account{propertyDefinition(input: $input){aliases,id,name,schema}}}"`,
		`{"input":{"alias":"my_prop"}}`,
		`{"data":{"account":{"propertyDefinition":{"aliases":["my_prop"],"id":"XXX","name":"my-prop","schema":"{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}"}}}}`,
	)
	client := BestTestClient(t, "properties/definition_get", testRequest)

	// Act
	property, err := client.GetPropertyDefinition(*ol.NewIdentifier("my_prop"))

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, "my-prop", property.Name)
	autopilot.Equals(t, "{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}", property.Schema)
}

func TestListPropertyDefinitions(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,schema},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop1"],"id":"XXX","name":"prop1","schema":"{\"key1\":\"val1\"}"},{"aliases":["prop2"],"id":"XXX","name":"prop2","schema":"{\"key2\":\"val2\"}"}],{{ template "pagination_initial_pageInfo_response" }},"totalCount":2}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,schema},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop3"],"id":"XXX","name":"prop3","schema":"{\"key3\":\"val3\"}"}],{{ template "pagination_second_pageInfo_response" }},"totalCount":1}}}}`,
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
