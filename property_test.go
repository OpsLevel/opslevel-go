package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

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
	autopilot.Equals(t, "my-prop", property.Name)
	autopilot.Equals(t, "{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}", property.Schema)
}
