package opslevel_test

import (
	"fmt"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2026"
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
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,aliases},locked,owner{__typename,... on Team{alias,id},... on Service{id,aliases}},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":true,"owner":{"__typename":"Service","id":"{{ template "id1_string" }}","aliases":[]},"validationErrors":[],"value":"true"}}}}`,
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
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,aliases},locked,owner{__typename,... on Team{alias,id},... on Service{id,aliases}},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"dropdown"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":false,"owner":{"__typename":"Service","id":"{{ template "id1_string" }}","aliases":[]},"validationErrors":[{"message":"vmessage1","path":["vmp1","vmp2"]},{"message":"vmessage2","path":["vmp3"]}],"value":"\"orange\""}}}}`,
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
		`query PropertyGet($definition:IdentifierInput!$owner:IdentifierInput!){account{property(owner: $owner, definition: $definition){definition{id,aliases},locked,owner{__typename,... on Team{alias,id},... on Service{id,aliases}},validationErrors{message,path},value}}}`,
		`{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"}}`,
		`{"data":{"account":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":true,"owner":{"__typename":"Service","id":"{{ template "id1_string" }}","aliases":[]},"validationErrors":[],"value":null}}}}`,
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
		`mutation PropertyAssign($input:PropertyInput!){propertyAssign(input: $input){property{definition{id,aliases},locked,owner{__typename,... on Team{alias,id},... on Service{id,aliases}},validationErrors{message,path},value},errors{message,path}}}`,
		`{"input": {{ template "property_assign_input" }} }`,
		`{"data":{"propertyAssign":{"property":{"definition":{"id":"{{ template "id2_string" }}"},"locked":true,"owner":{"__typename":"Service","id":"{{ template "id1_string" }}","aliases":[]},"validationErrors":[],"value":"true"},"errors":[]}}}`,
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

func TestCreateTeamPropertyDefinition(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString2)
	autopilot.Ok(t, schemaErr)
	expectedDefinition := autopilot.Register("expected_team_property_definition", ol.TeamPropertyDefinition{
		Alias:  "my_team_prop",
		Id:     id1,
		Name:   "my-team-prop",
		Schema: *schema,
	})
	input := autopilot.Register("team_property_definition_input", ol.TeamPropertyDefinitionInput{
		Alias:  "my_team_prop",
		Name:   "my-team-prop",
		Schema: *schema,
	})
	testRequest := autopilot.NewTestRequest(
		`mutation TeamPropertyDefinitionCreate($input:TeamPropertyDefinitionInput!){teamPropertyDefinitionCreate(input: $input){definition{alias,description,displaySubtype,displayType,id,lockedStatus,name,schema},errors{message,path}}}`,
		`{"input": {{ template "team_property_definition_input" }} }`,
		fmt.Sprintf(`{"data":{"teamPropertyDefinitionCreate":{"definition":{"alias":"my_team_prop","id":"%s","name":"my-team-prop","schema":%s},"errors":[]}}}`, id1, schemaString2),
	)
	client := BestTestClient(t, "properties/team_definition_create", testRequest)

	// Act
	result, err := client.CreateTeamPropertyDefinition(input)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedDefinition, *result)
	autopilot.Equals(t, expectedDefinition.Schema, result.Schema)
}

func TestUpdateTeamPropertyDefinition(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString2)
	autopilot.Ok(t, schemaErr)
	expectedDefinition := autopilot.Register("expected_team_property_definition", ol.TeamPropertyDefinition{
		Alias:       "my_team_prop",
		Description: "updated description",
		Id:          id1,
		Name:        "my-team-prop",
		Schema:      *schema,
	})
	input := autopilot.Register("team_property_definition_input", ol.TeamPropertyDefinitionInput{
		Alias:       "my_team_prop",
		Description: "updated description",
		Name:        "my-team-prop",
		Schema:      *schema,
	})
	testRequest := autopilot.NewTestRequest(
		`mutation TeamPropertyDefinitionUpdate($input:TeamPropertyDefinitionInput!$propertyDefinition:IdentifierInput!){teamPropertyDefinitionUpdate(propertyDefinition: $propertyDefinition, input: $input){definition{alias,description,displaySubtype,displayType,id,lockedStatus,name,schema},errors{message,path}}}`,
		`{"propertyDefinition":{"alias":"my_team_prop"}, "input": {{ template "team_property_definition_input" }} }`,
		fmt.Sprintf(`{"data":{"teamPropertyDefinitionUpdate":{"definition":{"alias":"my_team_prop","description":"updated description","id":"%s","name":"my-team-prop","schema":%s},"errors":[]}}}`, id1, schemaString2),
	)
	client := BestTestClient(t, "properties/team_definition_update", testRequest)

	// Act
	result, err := client.UpdateTeamPropertyDefinition("my_team_prop", input)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedDefinition, *result)
	autopilot.Equals(t, expectedDefinition.Schema, result.Schema)
}

func TestGetTeamPropertyDefinition(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString2)
	autopilot.Ok(t, schemaErr)
	expectedDefinition := autopilot.Register("expected_team_property_definition", ol.TeamPropertyDefinition{
		Alias:  "my_team_prop",
		Id:     id1,
		Name:   "my-team-prop",
		Schema: *schema,
	})
	testRequest := autopilot.NewTestRequest(
		`query TeamPropertyDefinitionGet($input:IdentifierInput!){account{teamPropertyDefinition(input: $input){alias,description,displaySubtype,displayType,id,lockedStatus,name,schema}}}`,
		`{"input":{"alias":"my_team_prop"}}`,
		fmt.Sprintf(`{"data":{"account":{"teamPropertyDefinition":{"alias":"my_team_prop","id":"%s","name":"my-team-prop","schema":%s}}}}`, id1, schemaString2),
	)
	client := BestTestClient(t, "properties/team_definition_get", testRequest)

	// Act
	result, err := client.GetTeamPropertyDefinition("my_team_prop")

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, expectedDefinition, *result)
	autopilot.Equals(t, expectedDefinition.Schema, result.Schema)
	autopilot.Equals(t, string(id1), string(result.Id))
}

func TestListTeamPropertyDefinitions(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString2)
	autopilot.Ok(t, schemaErr)
	expectedPage1 := autopilot.Register("team_property_definitions_page1", []ol.TeamPropertyDefinition{
		{Alias: "prop_a", Id: id1, Name: "prop-a", Schema: *schema},
		{Alias: "prop_b", Id: id2, Name: "prop-b", Schema: *schema},
	})
	expectedPage2 := autopilot.Register("team_property_definition_page2", ol.TeamPropertyDefinition{
		Alias: "prop_c", Id: id3, Name: "prop-c", Schema: *schema,
	})
	testRequestOne := autopilot.NewTestRequest(
		`query TeamPropertyDefinitionList($after:String!$first:Int!){account{teamPropertyDefinitions(after: $after, first: $first){nodes{alias,description,displaySubtype,displayType,id,lockedStatus,name,schema},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"teamPropertyDefinitions":{"nodes":[{"alias":"prop_a","id":"%s","name":"prop-a","schema":%s},{"alias":"prop_b","id":"%s","name":"prop-b","schema":%s}],{{ template "pagination_initial_pageInfo_response" }}}}}}`, id1, schemaString2, id2, schemaString2),
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query TeamPropertyDefinitionList($after:String!$first:Int!){account{teamPropertyDefinitions(after: $after, first: $first){nodes{alias,description,displaySubtype,displayType,id,lockedStatus,name,schema},{{ template "pagination_request" }}}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		fmt.Sprintf(`{"data":{"account":{"teamPropertyDefinitions":{"nodes":[{"alias":"prop_c","id":"%s","name":"prop-c","schema":%s}],{{ template "pagination_second_pageInfo_response" }}}}}}`, id3, schemaString2),
	)
	client := BestTestClient(t, "properties/team_definition_list", testRequestOne, testRequestTwo)

	// Act
	result, err := client.ListTeamPropertyDefinitions(nil)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result.Nodes))
	autopilot.Equals(t, expectedPage1[0], result.Nodes[0])
	autopilot.Equals(t, expectedPage1[1], result.Nodes[1])
	autopilot.Equals(t, expectedPage2, result.Nodes[2])
	autopilot.Equals(t, expectedPage1[0].Schema, result.Nodes[0].Schema)
	autopilot.Equals(t, 3, result.TotalCount)
}

func TestGetServiceProperties(t *testing.T) {
	// Arrange
	serviceId := ol.ServiceId{
		Id:      id1,
		Aliases: []string{},
	}
	service := ol.Service{
		ServiceId: serviceId,
	}
	owner := ol.PropertyOwner{
		Typename:  "Service",
		ServiceId: &serviceId,
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
		`query ServicePropertiesList($after:String!$first:Int!$service:ID!){account{service(id: $service){properties(after: $after, first: $first){nodes{definition{id,aliases},locked,owner{__typename,... on Team{alias,id},... on Service{id,aliases}},validationErrors{message,path},value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "{{ template "id1_string" }}" }`,
		`{"data":{"account":{"service":{"properties":{"nodes":[{{ template "service_properties_page_1" }}],{{ template "pagination_initial_pageInfo_response" }}}}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServicePropertiesList($after:String!$first:Int!$service:ID!){account{service(id: $service){properties(after: $after, first: $first){nodes{definition{id,aliases},locked,owner{__typename,... on Team{alias,id},... on Service{id,aliases}},validationErrors{message,path},value},{{ template "pagination_request" }}}}}}`,
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

func TestAssignTeamPropertyDefinitions(t *testing.T) {
	// Arrange
	schema, schemaErr := ol.NewJSONSchema(schemaString2)
	autopilot.Ok(t, schemaErr)
	input := autopilot.Register("team_property_definitions_assign_input", ol.TeamPropertyDefinitionsAssignInput{
		Properties: []ol.TeamPropertyDefinitionInput{
			{Alias: "prop_a", Name: "prop-a", Schema: *schema},
			{Alias: "prop_b", Name: "prop-b", Schema: *schema},
		},
	})
	expectedDefinitions := autopilot.Register("team_property_definitions_assigned", []ol.TeamPropertyDefinition{
		{Alias: "prop_a", Id: id1, Name: "prop-a", Schema: *schema},
		{Alias: "prop_b", Id: id2, Name: "prop-b", Schema: *schema},
	})
	testRequest := autopilot.NewTestRequest(
		`mutation TeamPropertyDefinitionsAssign($input:TeamPropertyDefinitionsAssignInput!){teamPropertyDefinitionsAssign(input: $input){properties{nodes{alias,description,displaySubtype,displayType,id,lockedStatus,name,schema},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},errors{message,path}}}`,
		`{"input": {{ template "team_property_definitions_assign_input" }} }`,
		fmt.Sprintf(`{"data":{"teamPropertyDefinitionsAssign":{"properties":{"nodes":[{"alias":"prop_a","id":"%s","name":"prop-a","schema":%s},{"alias":"prop_b","id":"%s","name":"prop-b","schema":%s}],"pageInfo":{"hasNextPage":false,"hasPreviousPage":false,"startCursor":"MQ","endCursor":"NA"}},"errors":[]}}}`, id1, schemaString2, id2, schemaString2),
	)
	client := BestTestClient(t, "properties/team_definitions_assign", testRequest)

	// Act
	result, err := client.AssignTeamPropertyDefinitions(input)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Nodes))
	autopilot.Equals(t, expectedDefinitions[0], result.Nodes[0])
	autopilot.Equals(t, expectedDefinitions[1], result.Nodes[1])
	autopilot.Equals(t, 2, result.TotalCount)
}
