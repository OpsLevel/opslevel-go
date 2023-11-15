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
	client := BestTestClient(t, "properties/create_definition", testRequest)

	// Act
	property, err := client.CreatePropertyDefinition(ol.PropertyDefinitionInput{
		Name:   "my-prop",
		Schema: "{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "my-prop", property.Name)
}
