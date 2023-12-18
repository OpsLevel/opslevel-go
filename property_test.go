package opslevel_test

import (
	"fmt"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

const (
	schemaString  = `{"$ref":"#/$defs/MyProp","$defs":{"MyProp":{"properties":{"name":{"type":"string","title":"the name","description":"The name of a friend","default":"alex","examples":["joe","lucy"]}},"additionalProperties":false,"type":"object","required":["name"]}}}`
	schemaString2 = `{"enum": ["red","green","blue"],"type": "string"}`
)

func TestCreatePropertyDefinition(t *testing.T) {
	// Arrange
	schema := ol.NewJSON(schemaString)
	expectedPropertyDefinition := autopilot.Register[ol.PropertyDefinition]("expected_property_definition", ol.PropertyDefinition{
		Aliases: []string{"my_prop"},
		Id:      "XXX",
		Name:    "my-prop",
		Schema:  schema,
	})
	propertyDefinitionInput := autopilot.Register[ol.PropertyDefinitionInput]("property_definition_input", ol.PropertyDefinitionInput{
		Name:   "my-prop",
		Schema: schema,
	})
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyDefinitionCreate($input:PropertyDefinitionInput!){propertyDefinitionCreate(input: $input){definition{aliases,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},errors{message,path}}}`,
		`{"input": {{ template "property_definition_input" }} }`,
		fmt.Sprintf(`{"data":{"propertyDefinitionCreate":{"definition": {"aliases":["my_prop"],"id":"XXX","name":"my-prop","schema": %s}, "errors":[] }}}`, schemaString),
	)
	client := BestTestClient(t, "properties/definition_create", testRequest)

	// Act
	actualPropertyDefinition, err := client.CreatePropertyDefinition(propertyDefinitionInput)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedPropertyDefinition, *actualPropertyDefinition)
	autopilot.Equals(t, expectedPropertyDefinition.Schema, actualPropertyDefinition.Schema)
}

func TestUpdatePropertyDefinition(t *testing.T) {
	// Arrange
	schema := ol.NewJSON(schemaString2)
	expectedPropertyDefinition := autopilot.Register[ol.PropertyDefinition]("expected_property_definition", ol.PropertyDefinition{
		Aliases:               []string{"my_prop"},
		Id:                    "XXX",
		Name:                  "my-prop",
		Description:           "this description was added",
		Schema:                schema,
		PropertyDisplayStatus: ol.PropertyDisplayStatusEnumHidden,
	})
	propertyDefinitionInput := autopilot.Register[ol.PropertyDefinitionInput]("property_definition_input", ol.PropertyDefinitionInput{
		Description:           "this description was added",
		Schema:                schema,
		PropertyDisplayStatus: ol.PropertyDisplayStatusEnumHidden,
	})
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyDefinitionUpdate($input:PropertyDefinitionInput!$propertyDefinition:IdentifierInput!){propertyDefinitionUpdate(propertyDefinition: $propertyDefinition, input: $input){definition{aliases,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},errors{message,path}}}`,
		`{"propertyDefinition":{"alias":"my_prop"}, "input": {{ template "property_definition_input" }} }`,
		fmt.Sprintf(`{"data":{"propertyDefinitionUpdate":{"definition": {"aliases":["my_prop"],"id":"XXX","name":"my-prop","description":"this description was added","propertyDisplayStatus":"hidden","schema": %s}, "errors":[] }}}`, schemaString2),
	)
	client := BestTestClient(t, "properties/definition_update", testRequest)

	// Act
	actualPropertyDefinition, err := client.UpdatePropertyDefinition("my_prop", propertyDefinitionInput)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedPropertyDefinition, *actualPropertyDefinition)
	autopilot.Equals(t, expectedPropertyDefinition.Schema, actualPropertyDefinition.Schema)
}

func TestDeletePropertyDefinition(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyDefinitionDelete($input:IdentifierInput!){propertyDefinitionDelete(resource: $input){errors{message,path}}}`,
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
	schema := ol.NewJSON(schemaString)
	expectedPropertyDefinition := autopilot.Register[ol.PropertyDefinition]("expected_property_definition",
		ol.PropertyDefinition{
			Aliases: []string{"my_prop"},
			Id:      "XXX",
			Name:    "my-prop",
			Schema:  schema,
		})
	testRequest := autopilot.NewTestRequest(
		`query PropertyDefinitionGet($input:IdentifierInput!){account{propertyDefinition(input: $input){aliases,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema}}}`,
		`{"input":{"alias":"my_prop"}}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinition": {"aliases":["my_prop"],"id":"XXX","name":"my-prop","schema": %s }}}}`, schemaString),
	)
	client := BestTestClient(t, "properties/definition_get", testRequest)

	// Act
	property, err := client.GetPropertyDefinition("my_prop")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, expectedPropertyDefinition, *property)
	autopilot.Equals(t, "my-prop", property.Name)
	autopilot.Equals(t, schema, property.Schema)
}

func TestListPropertyDefinitions(t *testing.T) {
	// Arrange
	schema := ol.NewJSON(schemaString)
	expectedPropDefsPageOne := autopilot.Register[[]ol.PropertyDefinition]("property_definitions", []ol.PropertyDefinition{
		{
			Aliases: []string{"prop1"},
			Id:      "XXX",
			Name:    "prop1",
			Schema:  ol.NewJSON(schemaString),
		},
		{
			Aliases: []string{"prop2"},
			Id:      "XXX",
			Name:    "prop2",
			Schema:  ol.NewJSON(schemaString),
		},
	})
	expectedPropDefPageTwo := autopilot.Register[ol.PropertyDefinition]("property_definition_3", ol.PropertyDefinition{
		Aliases: []string{"prop3"},
		Id:      "XXX",
		Name:    "prop3",
		Schema:  ol.NewJSON(schemaString),
	})
	testRequestOne := autopilot.NewTestRequest(
		`query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop1"],"id":"XXX","name":"prop1","schema": %s},{"aliases":["prop2"],"id":"XXX","name":"prop2","schema": %s}],{{ template "pagination_initial_pageInfo_response" }}}}}}`, schema.ToJSON(), schema.ToJSON()),
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,id,name,description,displaySubtype,displayType,propertyDisplayStatus,schema},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop3"],"id":"XXX","name":"prop3","schema": %s}],{{ template "pagination_second_pageInfo_response" }}}}}}`, schema.ToJSON()),
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}
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
	autopilot.Equals(t, expectedPropDefsPageOne[0].Schema, result[0].Schema)
	autopilot.Equals(t, expectedPropDefsPageOne[1].Schema, result[1].Schema)
	autopilot.Equals(t, expectedPropDefPageTwo.Schema, result[2].Schema)
}

func TestGetProperty(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,alias},owner{id,alias},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"XXX_PROPERTY_DEFINITION_ID_XXX"},"owner":{"id":"XXX_SERVICE_ID_XXX"},"validationErrors":[],"value":"true"}}}}`,
	)
	client := BestTestClient(t, "properties/property_get", testRequest)

	// Act
	property, err := client.GetProperty("monolith", "is_beta_feature")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "XXX_PROPERTY_DEFINITION_ID_XXX", string(*property.Definition.Id))
	autopilot.Equals(t, "XXX_SERVICE_ID_XXX", string(*property.Owner.Id))
	autopilot.Equals(t, 0, len(property.ValidationErrors))
	autopilot.Equals(t, "true", string(property.Value))
}

func TestAssignProperty(t *testing.T) {
	// Arrange
	input := ol.PropertyInput{
		Owner:      *ol.NewIdentifier("monolith"),
		Definition: *ol.NewIdentifier("is_beta_feature"),
		Value:      "true",
	}
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyAssign($input:PropertyInput!){propertyAssign(input: $input){property{definition{id,alias},owner{id,alias},validationErrors{message,path},value},errors{message,path}}}`,
		`{"input": {{ template "property_assign_input" }} }`,
		`{"data":{"propertyAssign":{"property":{"definition":{"alias":"is_beta_feature"},"owner":{"alias":"monolith"},"validationErrors":[],"value":"true"},"errors":[]}}}`,
	)
	client := BestTestClient(t, "properties/property_assign", testRequest)

	// Act
	property, err := client.PropertyAssign(input)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "is_beta_feature", string(*property.Definition.Alias))
	autopilot.Equals(t, "monolith", string(*property.Owner.Alias))
	autopilot.Equals(t, 0, len(property.ValidationErrors))
	autopilot.Equals(t, "true", string(property.Value))
}

func TestUnassignProperty(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyUnassign($definition:IdentifierInput!$owner:IdentifierInput!){propertyUnassign(owner: $owner, definition: $definition){errors{message,path}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"}}`,
		`{"data":{"propertyUnassign":{"errors":[]}}}`,
	)
	client := BestTestClient(t, "properties/property_unassign", testRequest)

	// Act
	err := client.PropertyUnassign("monolith", "is_beta_feature")

	// Assert
	autopilot.Ok(t, err)
}

func TestGetServiceProperties(t *testing.T) {
	// Arrange
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: id1,
		},
	}
	expectedPropsPageOne := autopilot.Register[[]ol.Property]("service_properties", []ol.Property{
		{
			Definition:       *ol.NewIdentifier("propdef1"),
			Owner:            *ol.NewIdentifier(string(service.Id)),
			ValidationErrors: []ol.OpsLevelErrors{},
			Value:            ol.JSONString("true"),
		},
		{
			Definition:       *ol.NewIdentifier("propdef2"),
			Owner:            *ol.NewIdentifier(string(service.Id)),
			ValidationErrors: []ol.OpsLevelErrors{},
			Value:            ol.JSONString("false"),
		},
	})
	expectedPropsPageTwo := autopilot.Register[[]ol.Property]("service_properties_3", []ol.Property{
		{
			Definition:       *ol.NewIdentifier("propdef3"),
			Owner:            *ol.NewIdentifier(string(service.Id)),
			ValidationErrors: []ol.OpsLevelErrors{},
			Value:            ol.JSONString("\"Hello World!\""),
		},
	})
	testRequestOne := autopilot.NewTestRequest(
		`query ServicePropertiesList($after:String!$first:Int!$service:ID!){account{service(id: $service){properties(after: $after, first: $first){edges{cursor,node{definition{id,alias},owner{id,alias},validationErrors{message,path},value}},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data":{"account":{"service":{"properties":{"edges":[{{ template "service_property_edge_1" }}],{{ template "pagination_initial_pageInfo_response" }}}}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServicePropertiesList($after:String!$first:Int!$service:ID!){account{service(id: $service){properties(after: $after, first: $first){edges{cursor,node{definition{id,alias},owner{id,alias},validationErrors{message,path},value}},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data":{"account":{"service":{"properties":{"edges":[{{ template "service_property_edge_3" }}],{{ template "pagination_second_pageInfo_response" }}}}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "service/get_properties", requests...)

	// Act
	properties, err := service.GetProperties(client, nil)
	result := properties.Edges

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, expectedPropsPageOne[0].Value, result[0].Node.Value)
	autopilot.Equals(t, expectedPropsPageOne[1].Value, result[1].Node.Value)
	autopilot.Equals(t, expectedPropsPageTwo[0].Value, result[2].Node.Value)
}
