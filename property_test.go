package opslevel_test

import (
	"fmt"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

const (
	schemaString  = `{"$ref":"#/$defs/MyProp","$defs":{"MyProp":{"properties":{"name":{"type":"string","title":"the name","description":"The name of a friend","default":"alex","examples":["joe","lucy"]}},"additionalProperties":false,"type":"object","required":["name"]}}}`
	schemaString2 = `{"enum": ["red","green","blue"],"type": "string"}`
)

func TestCreatePropertyDefinition(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString)
	schemaAsJSON, schemaAsJSONErr := ol.NewJSONSchema(schemaString)
	autopilot.Ok(t, schemaErr)
	autopilot.Ok(t, schemaAsJSONErr)
	expectedPropertyDefinition := autopilot.Register("expected_property_definition", ol.PropertyDefinition{
		Aliases:              []string{"my_prop"},
		AllowedInConfigFiles: true,
		Id:                   "XXX",
		Name:                 "my-prop",
		Schema:               *schemaAsJSON,
	})
	propertyDefinitionInput := autopilot.Register("property_definition_input", ol.PropertyDefinitionInput{
		AllowedInConfigFiles: ol.RefOf(true),
		Name:                 ol.RefOf("my-prop"),
		Schema:               schema,
	})
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyDefinitionCreate($input:PropertyDefinitionInput!){propertyDefinitionCreate(input: $input){definition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,lockedStatus,schema},errors{message,path}}}`,
		`{"input": {{ template "property_definition_input" }} }`,
		fmt.Sprintf(`{"data":{"propertyDefinitionCreate":{"definition": {"aliases":["my_prop"],"allowedInConfigFiles":true,"id":"XXX","name":"my-prop","schema": %s}, "errors":[] }}}`, schemaString),
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
	schema, schemaErr := ol.NewJSONSchema(schemaString)
	schemaAsJSON, schemaAsJSONErr := ol.NewJSONSchema(schemaString2)
	autopilot.Ok(t, schemaErr)
	autopilot.Ok(t, schemaAsJSONErr)
	expectedPropertyDefinition := autopilot.Register("expected_property_definition", ol.PropertyDefinition{
		Aliases:               []string{"my_prop"},
		AllowedInConfigFiles:  false,
		Description:           "this description was added",
		Id:                    "XXX",
		Name:                  "my-prop",
		PropertyDisplayStatus: ol.PropertyDisplayStatusEnumHidden,
		Schema:                *schemaAsJSON,
	})
	propertyDefinitionInput := autopilot.Register("property_definition_input", ol.PropertyDefinitionInput{
		AllowedInConfigFiles:  ol.RefOf(false),
		Description:           ol.RefOf("this description was added"),
		PropertyDisplayStatus: &ol.PropertyDisplayStatusEnumHidden,
		Schema:                schema,
	})
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyDefinitionUpdate($input:PropertyDefinitionInput!$propertyDefinition:IdentifierInput!){propertyDefinitionUpdate(propertyDefinition: $propertyDefinition, input: $input){definition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,lockedStatus,schema},errors{message,path}}}`,
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
	schema, schemaErr := ol.NewJSONSchema(schemaString)
	autopilot.Ok(t, schemaErr)
	expectedPropertyDefinition := autopilot.Register("expected_property_definition",
		ol.PropertyDefinition{
			Aliases:              []string{"my_prop"},
			AllowedInConfigFiles: true,
			Id:                   "XXX",
			Name:                 "my-prop",
			Schema:               *schema,
		})
	testRequest := autopilot.NewTestRequest(
		`query PropertyDefinitionGet($input:IdentifierInput!){account{propertyDefinition(input: $input){aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,lockedStatus,schema}}}`,
		`{"input":{"alias":"my_prop"}}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinition": {"aliases":["my_prop"],"allowedInConfigFiles":true,"id":"XXX","name":"my-prop","schema": %s }}}}`, schemaString),
	)
	client := BestTestClient(t, "properties/definition_get", testRequest)

	// Act
	property, err := client.GetPropertyDefinition("my_prop")

	// Assert
	autopilot.Equals(t, "XXX", string(property.Id))
	autopilot.Equals(t, "my-prop", property.Name)
	autopilot.Equals(t, *schema, property.Schema)
	autopilot.Equals(t, expectedPropertyDefinition, *property)
	autopilot.Equals(t, true, property.AllowedInConfigFiles)
	autopilot.Ok(t, err)
}

func TestListPropertyDefinitions(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString)
	schemaPage1, schemaPage1Err := ol.NewJSONSchema(schemaString)
	schemaPage2, schemaPage2Err := ol.NewJSONSchema(schemaString)
	schemaPage3, schemaPage3Err := ol.NewJSONSchema(schemaString)
	autopilot.Ok(t, schemaErr)
	autopilot.Ok(t, schemaPage1Err)
	autopilot.Ok(t, schemaPage2Err)
	autopilot.Ok(t, schemaPage3Err)
	expectedPropDefsPageOne := autopilot.Register("property_definitions", []ol.PropertyDefinition{
		{
			AllowedInConfigFiles: true,
			Aliases:              []string{"prop1"},
			Id:                   "XXX",
			Name:                 "prop1",
			Schema:               *schemaPage1,
		},
		{
			AllowedInConfigFiles: false,
			Aliases:              []string{"prop2"},
			Id:                   "XXX",
			Name:                 "prop2",
			Schema:               *schemaPage2,
		},
	})
	expectedPropDefPageTwo := autopilot.Register("property_definition_3", ol.PropertyDefinition{
		AllowedInConfigFiles: true,
		Aliases:              []string{"prop3"},
		Id:                   "XXX",
		Name:                 "prop3",
		Schema:               *schemaPage3,
	})
	testRequestOne := autopilot.NewTestRequest(
		`query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,lockedStatus,schema},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop1"],"allowedInConfigFiles":true,"id":"XXX","name":"prop1","schema": %s},{"aliases":["prop2"],"allowedInConfigFiles":false,"id":"XXX","name":"prop2","schema": %s}],{{ template "pagination_initial_pageInfo_response" }}}}}}`, schema.AsString(), schema.AsString()),
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query PropertyDefinitionList($after:String!$first:Int!){account{propertyDefinitions(after: $after, first: $first){nodes{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,lockedStatus,schema},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"propertyDefinitions":{"nodes":[{"aliases":["prop3"],"allowedInConfigFiles":true,"id":"XXX","name":"prop3","schema": %s}],{{ template "pagination_second_pageInfo_response" }}}}}}`, schema.AsString()),
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
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,aliases},locked,owner{... on Service{id,aliases}},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":true,"owner":{"id":"{{ template "id1_string" }}"},"validationErrors":[],"value":"true"}}}}`,
	)
	client := BestTestClient(t, "properties/property_get", testRequest)

	// Act
	property, err := client.GetProperty("monolith", "is_beta_feature")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "true", string(*property.Value))
	autopilot.Equals(t, 0, len(property.ValidationErrors))
	autopilot.Equals(t, string(id1), string(property.Owner.Id()))
	autopilot.Equals(t, string(id2), string(property.Definition.Id))
	autopilot.Equals(t, true, property.Locked)
}

func TestGetPropertyHasErrors(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,aliases},locked,owner{... on Service{id,aliases}},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"dropdown"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":false,"owner":{"id":"{{ template "id1_string" }}"},"validationErrors":[{"message":"vmessage1","path":["vmp1","vmp2"]},{"message":"vmessage2","path":["vmp3"]}],"value":"\"orange\""}}}}`,
	)
	client := BestTestClient(t, "properties/property_get_has_errors", testRequest)

	// Act
	property, err := client.GetProperty("monolith", "dropdown")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(property.ValidationErrors))
	autopilot.Equals(t, false, property.Locked)
	autopilot.Equals(t, string(id1), string(property.Owner.Id()))
	autopilot.Equals(t, string(id2), string(property.Definition.Id))
	// validation error 1
	autopilot.Equals(t, "vmessage1", property.ValidationErrors[0].Message)
	autopilot.Equals(t, 2, len(property.ValidationErrors[0].Path))
	autopilot.Equals(t, "vmp1", property.ValidationErrors[0].Path[0])
	autopilot.Equals(t, "vmp2", property.ValidationErrors[0].Path[1])
	// validation error 2
	autopilot.Equals(t, "vmessage2", property.ValidationErrors[1].Message)
	autopilot.Equals(t, 1, len(property.ValidationErrors[1].Path))
	autopilot.Equals(t, "vmp3", property.ValidationErrors[1].Path[0])
	autopilot.Equals(t, "\"orange\"", string(*property.Value))
}

func TestGetPropertyHasNullValue(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,aliases},locked,owner{... on Service{id,aliases}},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":true,"owner":{"id":"{{ template "id1_string" }}"},"validationErrors":[],"value":null}}}}`,
	)
	client := BestTestClient(t, "properties/property_get_has_null_value", testRequest)

	// Act
	property, err := client.GetProperty("monolith", "is_beta_feature")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 0, len(property.ValidationErrors))
	autopilot.Equals(t, string(id1), string(property.Owner.Id()))
	autopilot.Equals(t, string(id2), string(property.Definition.Id))
	autopilot.Equals(t, true, property.Locked)
	autopilot.Equals(t, true, property.Value == nil)
}

func TestAssignProperty(t *testing.T) {
	// Arrange
	input := ol.PropertyInput{
		Owner:      *ol.NewIdentifier(string(id1)),
		Definition: *ol.NewIdentifier(string(id2)),
		Value:      "true",
	}
	testRequest := autopilot.NewTestRequest(
		`mutation PropertyAssign($input:PropertyInput!){propertyAssign(input: $input){property{definition{id,aliases},locked,owner{... on Service{id,aliases}},validationErrors{message,path},value},errors{message,path}}}`,
		`{"input": {{ template "property_assign_input" }} }`,
		`{"data":{"propertyAssign":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":true,"owner":{"id":"{{ template "id1_string" }}"},"validationErrors":[],"value":"true"},"errors":[]}}}`,
	)
	client := BestTestClient(t, "properties/property_assign", testRequest)

	// Act
	property, err := client.PropertyAssign(input)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "true", string(*property.Value))
	autopilot.Equals(t, 0, len(property.ValidationErrors))
	autopilot.Equals(t, string(id1), string(property.Owner.Id()))
	autopilot.Equals(t, string(id2), string(property.Definition.Id))
	autopilot.Equals(t, true, property.Locked)
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
	serviceId := ol.ServiceId{
		Id: id1,
	}
	service := ol.Service{
		ServiceId: serviceId,
	}
	owner := ol.EntityOwnerService{
		OnService: serviceId,
	}
	value1 := ol.JsonString("true")
	value2 := ol.JsonString("false")
	value3 := ol.JsonString("\"Hello World!\"")
	expectedPropsPageOne := autopilot.Register("service_properties", []ol.Property{
		{
			Locked: true,
			Definition: ol.PropertyDefinitionId{
				Id: id2,
			},
			Owner:            owner,
			ValidationErrors: []ol.Error{},
			Value:            &value1,
		},
		{
			Locked: false,
			Definition: ol.PropertyDefinitionId{
				Id: id3,
			},
			Owner:            owner,
			ValidationErrors: []ol.Error{},
			Value:            &value2,
		},
	})
	expectedPropsPageTwo := autopilot.Register("service_properties_3", []ol.Property{
		{
			Locked: true,
			Definition: ol.PropertyDefinitionId{
				Id: id4,
			},
			Owner:            owner,
			ValidationErrors: []ol.Error{},
			Value:            &value3,
		},
	})
	testRequestOne := autopilot.NewTestRequest(
		`query ServicePropertiesList($after:String!$first:Int!$service:ID!){account{service(id: $service){properties(after: $after, first: $first){nodes{definition{id,aliases},locked,owner{... on Service{id,aliases}},validationErrors{message,path},value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data":{"account":{"service":{"properties":{"nodes":[{{ template "service_properties_page_1" }}],{{ template "pagination_initial_pageInfo_response" }}}}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServicePropertiesList($after:String!$first:Int!$service:ID!){account{service(id: $service){properties(after: $after, first: $first){nodes{definition{id,aliases},locked,owner{... on Service{id,aliases}},validationErrors{message,path},value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data":{"account":{"service":{"properties":{"nodes":[{{ template "service_properties_page_2" }}],{{ template "pagination_second_pageInfo_response" }}}}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "service/get_properties", requests...)

	// Act
	properties, err := service.GetProperties(client, nil)
	result := properties.Nodes

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, expectedPropsPageOne[0], result[0])
	autopilot.Equals(t, expectedPropsPageOne[1], result[1])
	autopilot.Equals(t, expectedPropsPageTwo[0], result[2])
	autopilot.Equals(t, expectedPropsPageOne[0].Value, result[0].Value)
	autopilot.Equals(t, expectedPropsPageOne[1].Value, result[1].Value)
	autopilot.Equals(t, expectedPropsPageTwo[0].Value, result[2].Value)
}
