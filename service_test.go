package opslevel_test

import (
	"fmt"
	"strings"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestServiceTags(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceTagsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "service": { "tags": { "nodes": [ { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA5", "key": "prod", "value": "false" } ], {{ template "pagination_initial_pageInfo_response" }} } }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceTagsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "service": { "tags": { "nodes": [ { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA4", "key": "test", "value": "true" } ], {{ template "pagination_second_pageInfo_response" }} } }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/tags", requests...)
	// Act
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
		},
	}
	resp, err := service.GetTags(client, nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert

	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "prod", result[0].Key)
	autopilot.Equals(t, "false", result[0].Value)
	autopilot.Equals(t, "test", result[1].Key)
	autopilot.Equals(t, "true", result[1].Value)
}

func TestServiceSystem(t *testing.T) {
	// Arrange
	request := autopilot.NewTestRequest(
		`query ServiceSystemGet($service:ID!){account{service(id: $service){parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}}}}}}`,
		`{ "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "service": { "parent": {{ template "system1_response" }} } } } }`,
	)
	client := BestTestClient(t, "service/system", request)
	// Act
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
		},
	}
	resp, err := service.GetSystem(client, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(resp.Id))
}

func TestServiceTools(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceToolsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tools(after: $after, first: $first){nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": {
                    "account": {
                        "service": {
                            "tools": {
                                "nodes": [
                                    {
                                      "category": "incidents",
                                      "categoryAlias": null,
                                      "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC84MDYz",
                                      "displayName": "PagerDuty",
                                      "environment": "Production",
                                      "service": {
                                        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
                                        "aliases": [
                                          "foo"
                                        ]
                                      },
                                      "url": "https://pagerduty.com"
                                    }
                                ],
                                {{ template "pagination_initial_pageInfo_response" }}
                            }
                          }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceToolsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tools(after: $after, first: $first){nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": {
                    "account": {
                        "service": {
                            "tools": {
                                "nodes": [
                                    {
                                      "category": "continuous_integration",
                                      "categoryAlias": null,
                                      "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC84MDY0",
                                      "displayName": "Gitlab CI",
                                      "environment": "Production",
                                      "service": {
                                        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
                                        "aliases": [
                                          "foo"
                                        ]
                                      },
                                      "url": "https://gitlab.com"
                                    }
                                ],
                                {{ template "pagination_second_pageInfo_response" }}
                            }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/tools", requests...)
	// Act
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
		},
	}
	resp, err := service.GetTools(client, nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert

	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "PagerDuty", result[0].DisplayName)
	autopilot.Equals(t, ol.ToolCategoryContinuousIntegration, result[1].Category)
}

func TestServiceRepositories(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repos(after: $after, first: $first){edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": {
                    "account": {
                        "service": {
                            "repos": {
                              "edges": [
                                {
                                  "node": {
                                    "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw",
                                    "defaultAlias": "bitbucket.org:raptors-store/Store Front"
                                  },
                                  "serviceRepositories": [
                                    {
                                      "baseDirectory": "shopping-cart",
                                      "displayName": "raptors-store/Store Front",
                                      "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvMjc2Nw",
                                      "repository": {
                                        "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw",
                                        "defaultAlias": "bitbucket.org:raptors-store/Store Front"
                                      },
                                      "service": {
                                        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8xOTQy",
                                        "aliases": [
                                          "dogfood",
                                          "opslevel-frontend",
                                          "opslevel_com",
                                          "service_alias",
                                          "shopping_cart",
                                          "shopping_cart 1",
                                          "shopping_cart_1235",
                                          "shopping_cart_2",
                                          "shopping_cart_service_2",
                                          "shopping_tart",
                                          "shopping_tarts"
                                        ]
                                      }
                                    }
                                  ]
                                }
                              ],
                                {{ template "pagination_initial_pageInfo_response" }}
                            }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repos(after: $after, first: $first){edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": {
                    "account": {
                        "service": {
                            "repos": {
                              "edges": [
                                {
                                  "node": {
                                    "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw",
                                    "defaultAlias": "bitbucket.org:raptors-store/Store Front"
                                  },
                                  "serviceRepositories": [
                                    {
                                      "baseDirectory": "shopping-cart",
                                      "displayName": "raptors-store/Store Front",
                                      "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvMjc2Nw",
                                      "repository": {
                                        "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw",
                                        "defaultAlias": "bitbucket.org:raptors-store/Store Front"
                                      },
                                      "service": {
                                        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8xOTQy",
                                        "aliases": [
                                          "dogfood",
                                          "opslevel-frontend",
                                          "opslevel_com",
                                          "service_alias",
                                          "shopping_cart",
                                          "shopping_cart 1",
                                          "shopping_cart_1235",
                                          "shopping_cart_2",
                                          "shopping_cart_service_2",
                                          "shopping_tart",
                                          "shopping_tarts"
                                        ]
                                      }
                                    }
                                  ]
                                }
                              ],
                                {{ template "pagination_second_pageInfo_response" }}
                            }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/repositories", requests...)
	// Act
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
		},
	}
	resp, err := service.GetRepositories(client, nil)
	autopilot.Ok(t, err)
	result := resp.Edges
	// Assert

	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw", string(result[0].Node.Id))
	autopilot.Equals(t, "bitbucket.org:raptors-store/Store Front", result[1].Node.DefaultAlias)
}

func TestCreateService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceCreate($input:ServiceCreateInput!){serviceCreate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
		`{ "input": { "name": "Foo", "description": "Foo service" } }`,
		`{ "data": { "serviceCreate": { "service": {{ template "service_1" }}, "errors": [] } }}`,
	)
	client := BestTestClient(t, "service/create", testRequest)
	// Act
	result, err := client.CreateService(ol.ServiceCreateInput{
		Name:        "Foo",
		Description: ol.RefOf("Foo service"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result.Aliases))
}

func TestCreateServiceWithNote(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`mutation ServiceCreate($input:ServiceCreateInput!){serviceCreate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
		`{ "input": { "name": "Foo", "description": "Foo service" } }`,
		`{ "data": { "serviceCreate": { "service": {{ template "service_1" }}, "errors": [] } }}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`mutation ServiceUpdateNote($input:ServiceNoteUpdateInput!){serviceNoteUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},errors{message,path}}}`,
		`{ "input": { "service": { {{ template "id1" }} }, "note": "Foo note" } }`,
		`{ "data": { "serviceNoteUpdate": { "service": {{ template "service_with_note" }}, "errors": [] } }}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/create_with_note", requests...)
	// Act
	service, servicErr := client.CreateService(ol.ServiceCreateInput{
		Name:        "Foo",
		Description: ol.RefOf("Foo service"),
	})
	autopilot.Ok(t, servicErr)
	note := "Foo note"
	service, noteErr := client.UpdateServiceNote(ol.ServiceNoteUpdateInput{
		Service: *ol.NewIdentifier(string(service.Id)),
		Note:    ol.RefOf(note),
	})
	// Assert
	autopilot.Ok(t, noteErr)
	autopilot.Equals(t, note, service.Note)
}

func TestCreateServiceWithParentSystem(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceCreate($input:ServiceCreateInput!){serviceCreate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
		`{ "input": { "name": "Foo", "description": "Foo service", "parent": {"alias": "FooSystem"} } }`,
		`{ "data": { "serviceCreate": { "service": {{ template "service_1" }}, "errors": [] } }}`,
	)
	client := BestTestClient(t, "service/create_with_system", testRequest)
	// Act
	result, err := client.CreateService(ol.ServiceCreateInput{
		Name:        "Foo",
		Description: ol.RefOf("Foo service"),
		Parent:      ol.NewIdentifier("FooSystem"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, true, result.Locked)
	autopilot.Equals(t, 1, len(result.Aliases))
}

func TestUpdateService(t *testing.T) {
	addVars := `{"input":{"description": "The quick brown fox", "framework": "django", "id": "123456789", "lifecycleAlias": "pre-alpha", "name": "Hello World", "parent": {"alias": "some_system"}, "tierAlias": "tier_4"}}`
	// delVars := `{"input":{"description": null, "framework": null, "id": "123456789", "lifecycleAlias": null, "parent": null, "tierAlias": null}}`
	delVarsV1DoesNotWorkExceptOnParent := `{"input":{"id": "123456789", "parent": null}}`
	zeroVars := `{"input":{"description": "", "framework": "", "id": "123456789"}}`
	type TestCase struct {
		Name  string
		Vars  string
		Input ol.ServiceUpdateInput
	}
	testCases := []TestCase{
		{
			Name: "add fields v1",
			Vars: addVars,
			Input: ol.ServiceUpdateInput{
				Parent:         ol.NewIdentifier("some_system"),
				Id:             ol.RefOf(ol.ID("123456789")),
				Name:           ol.RefOf("Hello World"),
				Description:    ol.RefOf("The quick brown fox"),
				Framework:      ol.RefOf("django"),
				TierAlias:      ol.RefOf("tier_4"),
				LifecycleAlias: ol.RefOf("pre-alpha"),
			},
		},
		// {
		// 	Name: "add fields v2",
		// 	Vars: addVars,
		// 	Input: ol.ServiceUpdateInputV2{
		// 		Parent:         ol.NewIdentifier("some_system"),
		// 		Id:             ol.NewID("123456789"),
		// 		Name:           ol.RefOf("Hello World"),
		// 		Description:    ol.RefOf("The quick brown fox"),
		// 		Framework:      ol.RefOf("django"),
		// 		TierAlias:      ol.RefOf("tier_4"),
		// 		LifecycleAlias: ol.RefOf("pre-alpha"),
		// 	},
		// },
		{
			Name: "unset fields v1 - does not work except on parent",
			Vars: delVarsV1DoesNotWorkExceptOnParent,
			Input: ol.ServiceUpdateInput{
				Parent:         &ol.IdentifierInput{},
				Id:             ol.RefOf(ol.ID("123456789")),
				Description:    nil,
				Framework:      nil,
				TierAlias:      nil,
				LifecycleAlias: nil,
			},
		},
		// {
		// 	Name: "unset fields v2 - works on all including parent",
		// 	Vars: delVars,
		// 	Input: ol.ServiceUpdateInputV2{
		// 		Parent:         ol.NewIdentifier(),
		// 		Id:             ol.NewID("123456789"),
		// 		Description:    ol.NewNull(),
		// 		Framework:      ol.NewNull(),
		// 		TierAlias:      ol.NewNull(),
		// 		LifecycleAlias: ol.NewNull(),
		// 	},
		// },
		{
			Name: "set fields to zero value v1",
			Vars: zeroVars,
			Input: ol.ServiceUpdateInput{
				Id:          ol.RefOf(ol.ID("123456789")),
				Description: ol.RefOf(""),
				Framework:   ol.RefOf(""),
			},
		},
		// {
		// 	Name: "set fields to zero value v2",
		// 	Vars: zeroVars,
		// 	Input: ol.ServiceUpdateInputV2{
		// 		Id:          ol.RefOf(ol.ID("123456789")),
		// 		Description: ol.NewNull(),
		// 		Framework:   ol.NewNull(),
		// 	},
		// },
	}

	for i, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			testRequest := autopilot.NewTestRequest(
				`mutation ServiceUpdate($input:ServiceUpdateInput!){serviceUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
				testCase.Vars,
				`{"data": {"serviceUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
			)

			client := BestTestClient(t, fmt.Sprintf("service/update_%d", i+1), testRequest)

			_, err := client.UpdateService(testCase.Input)
			if err != nil {
				t.Errorf("got unexpected error: '%+v'", err)
			}
		})
	}
}

func TestGetServiceIdWithAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceGet($service:String!){account{service(alias: $service){id,aliases}}}`,
		`{ "service": "coredns" }`,
		`{ "data": { "account": { "service": { "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx" } } } }`,
	)
	client := BestTestClient(t, "service/get_id_with_alias", testRequest)
	// Act
	result, err := client.GetServiceIdWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", string(result.Id))
}

func TestGetServiceWithAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceGet($service:String!){account{service(alias: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}}}}`,
		`{ "service": "coredns" }`,
		`{ "data": {
    "account": {
      "service": {
        "aliases": [
          "coredns",
          "coredns-kube-system"
        ],
        "apiDocumentPath": null,
        "description": null,
        "framework": null,
        "language": "go",
        "lifecycle": {
          "alias": "alpha",
          "description": "Service is supporting features used by others at the company, or a very small set of friendly customers.",
          "id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQyNw",
          "index": 2,
          "name": "Alpha"
        },
        "name": "coredns",
        "owner": {
          "alias": "developers",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NDk"
        },
        "preferredApiDocument": {
          "id": "Z2lkOi8vb3BzbGV2ZWwvRG9jdW1lbnRzOjpBcGkvOTU0MQ",
          "htmlUrl": null,
          "source": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkRvY3VtZW50czo6QXBpRG9jSW50ZWdyYXRpb24vMTgxNw",
            "name": "API Docs",
            "type": "apiDoc"
          },
          "timestamps": {
            "createdAt": "2022-07-22T17:01:53.080794Z",
            "updatedAt": "2022-07-22T17:01:53.101899Z"
          }
        },
        "preferredApiDocumentSource": "PUSH",
        "product": "MyProduct",
        "repos": {
          "edges": [
            {
              "node": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk",
                "defaultAlias": "github.com:rocktavious/autopilot"
              },
              "serviceRepositories": [
                {
                  "baseDirectory": "",
                  "displayName": "rocktavious/autopilot",
                  "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvNDIxNw",
                  "repository": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk",
                    "defaultAlias": "github.com:rocktavious/autopilot"
                  },
                  "service": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81NDI5"
                  }
                }
              ]
            }
          ],
          {{ template "pagination_response_same_cursor" }}
        },
        "tier": {
          "alias": "tier_1",
          "description": "Mission-critical service or repository. Failure could result in significant impact to revenue or reputation.",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVGllci8zNDE",
          "index": 1,
          "name": "Tier 1"
        },
        "tags": {
          "nodes": [
            {
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzExMDg4NA",
              "key": "k8s-app",
              "value": "kube-dns"
            },
            {
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzExMDg4NQ",
              "key": "imported",
              "value": "kubectl-opslevel"
            },
            {
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzExMDg4Ng",
              "key": "hello",
              "value": "world"
            }
          ],
          {{ template "pagination_response_different_cursor" }}
        },
        "timestamps": {
          "createdAt": "2022-07-22T16:59:20.361676Z",
          "updatedAt": "2022-07-22T16:59:38.940251Z"
        },
        "tools": {
          "nodes": [
            {
              "category": "code",
              "categoryAlias": null,
              "displayName": "GitHub",
              "environment": "prod",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNA",
              "url": "https://github.com/opslevel/coredns",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            },
            {
              "category": "code",
              "categoryAlias": null,
              "displayName": "github",
              "environment": null,
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNQ",
              "url": "https://github.com",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            },
            {
              "category": "logs",
              "categoryAlias": null,
              "displayName": "logz",
              "environment": null,
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNg",
              "url": "https://logz.com",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            },
            {
              "category": "other",
              "categoryAlias": null,
              "displayName": "datadog",
              "environment": null,
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNw",
              "url": "https://datadog.com",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            }
          ],
          {{ template "no_pagination_response" }}
        },
        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
      }
    }
  }
}`,
	)
	client := BestTestClient(t, "service/get_with_alias", testRequest)
	// Act
	result, err := client.GetServiceWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, "API Docs", result.PreferredApiDocument.Source.Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPush, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, true, result.HasAlias("coredns"))
	autopilot.Equals(t, false, result.HasAlias("opslevel-dns"))
	autopilot.Equals(t, true, result.HasTag("hello", "world"))
	autopilot.Equals(t, false, result.HasTag("provider", "opslevel"))
	autopilot.Equals(t, true, result.HasTool("code", "GitHub", "prod"))
	autopilot.Equals(t, false, result.HasTool("observability", "honeycomb", "certification"))
}

func TestGetService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceGet($service:ID!){account{service(id: $service){{ template "service_get" }}}}}`,
		`{ "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx" }`,
		`{ "data": {
    "account": {
      "service": {
        "aliases": [
          "coredns",
          "coredns-kube-system"
        ],
        "apiDocumentPath": null,
        "description": null,
        "framework": null,
        "language": "go",
        "lifecycle": {
          "alias": "alpha",
          "description": "Service is supporting features used by others at the company, or a very small set of friendly customers.",
          "id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQyNw",
          "index": 2,
          "name": "Alpha"
        },
        "name": "coredns",
        "owner": {
          "alias": "developers",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NDk"
        },
        "parent": {
          "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzExOTc",
          "aliases": ["just_updated_this_with_an_alias","taimoor_s_orange_system","update_2_lol"]
        },
        "preferredApiDocument": {
          "id": "Z2lkOi8vb3BzbGV2ZWwvRG9jdW1lbnRzOjpBcGkvOTU0MQ",
          "htmlUrl": null,
          "source": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkRvY3VtZW50czo6QXBpRG9jSW50ZWdyYXRpb24vMTgxNw",
            "name": "API Docs",
            "type": "apiDoc"
          },
          "timestamps": {
            "createdAt": "2022-07-22T17:01:53.080794Z",
            "updatedAt": "2022-07-22T17:01:53.101899Z"
          }
        },
        "preferredApiDocumentSource": "PUSH",
        "product": "MyProduct",
        "repos": {
          "edges": [
            {
              "node": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk",
                "defaultAlias": "github.com:rocktavious/autopilot"
              },
              "serviceRepositories": [
                {
                  "baseDirectory": "",
                  "displayName": "rocktavious/autopilot",
                  "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvNDIxNw",
                  "repository": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk",
                    "defaultAlias": "github.com:rocktavious/autopilot"
                  },
                  "service": {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81NDI5"
                  }
                }
              ]
            }
          ],
          {{ template "pagination_response_same_cursor" }}
        },
        "tier": {
          "alias": "tier_1",
          "description": "Mission-critical service or repository. Failure could result in significant impact to revenue or reputation.",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVGllci8zNDE",
          "index": 1,
          "name": "Tier 1"
        },
        "tags": {
          "nodes": [
            {
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzExMDg4NA",
              "key": "k8s-app",
              "value": "kube-dns"
            },
            {
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzExMDg4NQ",
              "key": "imported",
              "value": "kubectl-opslevel"
            },
            {
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzExMDg4Ng",
              "key": "hello",
              "value": "world"
            }
          ],
          {{ template "pagination_response_different_cursor" }}
        },
        "timestamps": {
          "createdAt": "2022-07-22T16:59:20.361676Z",
          "updatedAt": "2022-07-22T16:59:38.940251Z"
        },
        "tools": {
          "nodes": [
            {
              "category": "code",
              "categoryAlias": null,
              "displayName": "GitHub",
              "environment": "prod",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNA",
              "url": "https://github.com/opslevel/coredns",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            },
            {
              "category": "code",
              "categoryAlias": null,
              "displayName": "github",
              "environment": null,
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNQ",
              "url": "https://github.com",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            },
            {
              "category": "logs",
              "categoryAlias": null,
              "displayName": "logz",
              "environment": null,
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNg",
              "url": "https://logz.com",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            },
            {
              "category": "other",
              "categoryAlias": null,
              "displayName": "datadog",
              "environment": null,
              "id": "Z2lkOi8vb3BzbGV2ZWwvVG9vbC8yMjgxNw",
              "url": "https://datadog.com",
              "service": {
                "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
              }
            }
          ],
          {{ template "no_pagination_response" }}
        },
        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx"
      }
    }
  }
}`,
	)
	client := BestTestClient(t, "service/get", testRequest)
	// Act
	result, err := client.GetService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzExOTc", string(result.Parent.Id))
	autopilot.Equals(t, []string{"just_updated_this_with_an_alias", "taimoor_s_orange_system", "update_2_lol"}, result.Parent.Aliases)
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, (*ol.ServiceRepository)(nil), result.Repository)
}

func TestGetServiceDocuments(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceDocumentsList($after:String!$first:Int!$searchTerm:String!$service:ID!){account{service(id: $service){documents(searchTerm: $searchTerm, after: $after, first: $first){nodes{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},{{ template "pagination_request" }}}}}}`,
		`{ "service": "{{ template "id1_string" }}", {{ template "first_page_variables" }}, "searchTerm": "" }`,
		`{ "data": { "account": { "service": { "documents": { "nodes": [ {{ template "document_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceDocumentsList($after:String!$first:Int!$searchTerm:String!$service:ID!){account{service(id: $service){documents(searchTerm: $searchTerm, after: $after, first: $first){nodes{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},{{ template "pagination_request" }}}}}}`,
		`{ "service": "{{ template "id1_string" }}", {{ template "second_page_variables" }}, "searchTerm": "" }`,
		`{ "data": { "account": { "service": { "documents": { "nodes": [ {{ template "document_1" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/get_documents", requests...)
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: id1,
		},
	}
	// Act
	resp, err := service.GetDocuments(client, nil)
	// result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	// autopilot.Equals(t, "Foo", result[0].HtmlURL)
}

func TestGetServiceStats(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`query GetServiceStats($service:ID!){account{service(id: $service){serviceStats{rubric{categoryLevel{alias,checks{id,name},description,id,index,name},checkResults{byLevel{nodes{items{nodes{check{id,name},lastUpdated,message,service{id,aliases},serviceAlias,status},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},level{alias,checks{id,name},description,id,index,name}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},nextLevel{level{alias,checks{id,name},description,id,index,name}}},level{alias,checks{id,name},description,id,index,name}}}}}}`,
		`{ "service": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx" }`,
		`{ "data": {
    "account": {
      "service": {
        "serviceStats": {
          "rubric": {
            "categoryLevel": null,
            "checkResults": {
              "byLevel": {
                "nodes": [
                  {
                    "items": {
                      "nodes": [
                        {
                          "check": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpTZXJ2aWNlUHJvcGVydHkvNDc2OA",
                            "name": "Has Tier Defined"
                          },
                          "lastUpdated": "2025-06-05T19:52:54.177040Z",
                          "message": "The service has a tier.",
                          "service": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8x",
                            "aliases": []
                          },
                          "serviceAlias": "",
                          "status": "passed"
                        }
                      ],
                      "pageInfo": {
                        "hasNextPage": false,
                        "hasPreviousPage": false,
                        "startCursor": "MQ",
                        "endCursor": "Nw"
                      }
                    },
                    "level": {
                      "alias": "bronze",
                      "checks": [
                        {
                          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpTZXJ2aWNlUHJvcGVydHkvNDc2OA",
                          "name": "Has Tier Defined"
                        }
                      ],
                      "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
                      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI0",
                      "index": 1,
                      "name": "🥉 Bronze"
                    }
                  },
                  {
                    "items": {
                      "nodes": [
                        {
                          "check": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpBbGVydFNvdXJjZVVzYWdlLzczODA",
                            "name": "On-call configured for each service"
                          },
                          "lastUpdated": "2025-06-05T19:52:53.931285Z",
                          "message": "The component is using a On-call alert source.",
                          "service": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8x",
                            "aliases": []
                          },
                          "serviceAlias": "",
                          "status": "passed"
                        }
                      ],
                      "pageInfo": {
                        "hasNextPage": false,
                        "hasPreviousPage": false,
                        "startCursor": "MQ",
                        "endCursor": "Ng"
                      }
                    },
                    "level": {
                      "alias": "silver",
                      "checks": [
                        {
                          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpBbGVydFNvdXJjZVVzYWdlLzczODA",
                          "name": "On-call configured for each service"
                        }
                      ],
                      "description": "Services in this level satisfy important and critical checks. This is considered healthy but there is room for improvement.",
                      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI1",
                      "index": 2,
                      "name": "🥈 Silver"
                    }
                  },
                  {
                    "items": {
                      "nodes": [
                        {
                          "check": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpSZXBvRmlsZS82ODU1",
                            "name": "Has 'Readme.md'"
                          },
                          "lastUpdated": "2025-06-05T19:54:40.256998Z",
                          "message": "Repo files <a href='https://github.com/helpify/helpify/blob/main/README.md' target='_blank'>Readme.md, README.md, README.MD, readme.md or Readme.MD</a> exists.",
                          "service": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8x",
                            "aliases": []
                          },
                          "serviceAlias": "",
                          "status": "passed"
                        }
                      ],
                      "pageInfo": {
                        "hasNextPage": false,
                        "hasPreviousPage": false,
                        "startCursor": "MQ",
                        "endCursor": "Ng"
                      }
                    },
                    "level": {
                      "alias": "gold",
                      "checks": [
                        {
                          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpSZXBvRmlsZS82ODU1",
                          "name": "Has 'Readme.md'"
                        }
                      ],
                      "description": "Services in this level satisfy critical, important and useful checks. This is the level all services should aspire to be in.",
                      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI2",
                      "index": 3,
                      "name": "🥇 Gold"
                    }
                  },
                  {
                    "items": {
                      "nodes": [
                        {
                          "check": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzc1Njc",
                            "name": "No Version Bump PR/MR older than 14d"
                          },
                          "lastUpdated": "2025-06-05T12:16:53.897400Z",
                          "message": "### Check failed\n  Service **opslevel** has version bump MRs older than 14 days.",
                          "service": {
                            "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8x",
                            "aliases": [
                            ]
                          },
                          "serviceAlias": "",
                          "status": "failed"
                        }
                      ],
                      "pageInfo": {
                        "hasNextPage": false,
                        "hasPreviousPage": false,
                        "startCursor": "MQ",
                        "endCursor": "NQ"
                      }
                    },
                    "level": {
                      "alias": "platinum",
                      "checks": [
                        {
                          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzc1Njc",
                          "name": "No Version Bump PR/MR older than 14d"
                        }
                      ],
                      "description": "Services in this level satisfy above and beyond checks. This is the equivalent of getting an A+ and doing all your extra credit assignments.",
                      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMTAyOA",
                      "index": 4,
                      "name": "💎 Platinum"
                    }
                  }
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "NA"
                }
              },
              "nextLevel": {
                "level": {
                  "alias": "platinum",
                  "checks": [
                    {
                      "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzc1Njc",
                      "name": "No Version Bump PR/MR older than 14d"
                    }
                  ],
                  "description": "Services in this level satisfy above and beyond checks. This is the equivalent of getting an A+ and doing all your extra credit assignments.",
                  "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMTAyOA",
                  "index": 4,
                  "name": "💎 Platinum"
                }
              }
            },
            "level": {
              "alias": "gold",
              "checks": [
                {
                  "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpSZXBvRmlsZS82ODU1",
                  "name": "Has 'Readme.md'"
                }
              ],
              "description": "Services in this level satisfy critical, important and useful checks. This is the level all services should aspire to be in.",
              "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMjI2",
              "index": 3,
              "name": "🥇 Gold"
            }
          }
        }
      }
    }
  }
}
`,
	)
	client := BestTestClient(t, "service/get_service_stats", testRequest)
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: id1,
		},
	}
	// Act
	resp, err := service.GetServiceStats(client)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "🥇 Gold", resp.Rubric.Level.Name)
	autopilot.Equals(t, "Has 'Readme.md'", resp.Rubric.Level.Checks[0].Name)
	autopilot.Equals(t, ol.CheckStatus("passed"), resp.Rubric.CheckResults.ByLevel.Nodes[0].Items.Nodes[0].Status)
	autopilot.Equals(t, ol.CheckStatus("failed"), resp.Rubric.CheckResults.ByLevel.Nodes[3].Items.Nodes[0].Status)
}

func TestGetServiceStatsInvalidServiceId(t *testing.T) {
	client := BestTestClient(t, "service/get_service_stats_invalid_service")
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id: "",
		},
	}
	// Act
	_, err := service.GetServiceStats(client)
	// Assert
	autopilot.Equals(t, "unable to get 'ServiceStats', invalid service id: ''", err.Error())
}

func TestListServices(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceList($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceList($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list", requests...)
	// Act
	resp, err := client.ListServices(nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert

	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
	autopilot.Equals(t, "Backend", result[1].Repository.DisplayName)
}

func TestListServicesWithFilter(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithFilter($after:String!$filter:IdentifierInput$first:Int!){account{services(filterIdentifier: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ {{ template "first_page_variables" }}, "filter": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithFilter($after:String!$filter:IdentifierInput$first:Int!){account{services(filterIdentifier: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ {{ template "second_page_variables" }}, "filter": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_filter", requests...)
	// Act
	response, err := client.ListServicesWithFilter(string(id1), nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithFramework(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithFramework($after:String!$first:Int!$framework:String!){account{services(framework: $framework, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "framework": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithFramework($after:String!$first:Int!$framework:String!){account{services(framework: $framework, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "framework": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_framework", requests...)
	// Act
	response, err := client.ListServicesWithFramework("postgres", nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithLanguage(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithLanguage($after:String!$first:Int!$language:String!){account{services(language: $language, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "language": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithLanguage($after:String!$first:Int!$language:String!){account{services(language: $language, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "language": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_language", requests...)
	// Act
	response, err := client.ListServicesWithLanguage("postgres", nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithOwner(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithOwner($after:String!$first:Int!$owner:String!){account{services(ownerAlias: $owner, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "owner": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithOwner($after:String!$first:Int!$owner:String!){account{services(ownerAlias: $owner, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "owner": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_owner", requests...)
	// Act
	response, err := client.ListServicesWithOwner("postgres", nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithTag(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithTag($after:String!$first:Int!$tag:TagArgs!){account{services(tag: $tag, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "tag": { "key": "app", "value": "worker" } }`,
		`{"data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithTag($after:String!$first:Int!$tag:TagArgs!){account{services(tag: $tag, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "tag": { "key": "app", "value": "worker" } }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_tag", requests...)
	// Act
	tagArgs, err := ol.NewTagArgs("app:worker")
	autopilot.Ok(t, err)
	response, err := client.ListServicesWithTag(tagArgs, nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithTier(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithTier($after:String!$first:Int!$tier:String!){account{services(tierAlias: $tier, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "tier": "tier_1" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithTier($after:String!$first:Int!$tier:String!){account{services(tierAlias: $tier, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "tier": "tier_1" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_tier", requests...)
	// Act
	response, err := client.ListServicesWithTier("tier_1", nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithLifecycle(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithLifecycle($after:String!$first:Int!$lifecycle:String!){account{services(lifecycleAlias: $lifecycle, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "lifecycle": "alpha" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithLifecycle($after:String!$first:Int!$lifecycle:String!){account{services(lifecycleAlias: $lifecycle, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "lifecycle": "alpha" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_lifecycle", requests...)
	// Act
	response, err := client.ListServicesWithLifecycle("alpha", nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithProduct(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithProduct($after:String!$first:Int!$product:String!){account{services(product: $product, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "first_page_variables" }}, "product": "test" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithProduct($after:String!$first:Int!$product:String!){account{services(product: $product, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},{{ template "pagination_request" }}}}}`,
		`{ {{ template "second_page_variables" }}, "product": "test" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_product", requests...)
	// Act
	response, err := client.ListServicesWithProduct("test", nil)
	autopilot.Ok(t, err)
	result := response.Nodes
	// Assert

	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestDeleteService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceDelete($input:ServiceDeleteInput!){serviceDelete(input: $input){deletedServiceId,deletedServiceAlias,errors{message,path}}}`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3" } }`,
		`{ "data": { "serviceDelete": { "deletedServiceId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3", "deletedServiceAlias": "db", "errors": [] } } }`,
	)
	client := BestTestClient(t, "service/delete", testRequest)
	// Act
	err := client.DeleteService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteServicesWithAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceDelete($input:ServiceDeleteInput!){serviceDelete(input: $input){deletedServiceId,deletedServiceAlias,errors{message,path}}}`,
		`{ "input": { "alias": "db" } }`,
		`{ "data": { "serviceDelete": { "deletedServiceId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3", "deletedServiceAlias": "db", "errors": [] } } }`,
	)
	client := BestTestClient(t, "service/delete_with_alias", testRequest)
	// Act
	err := client.DeleteService("db")
	// Assert
	autopilot.Ok(t, err)
}

func TestServiceReconcileAliasesDeleteAll(t *testing.T) {
	// Arrange
	aliasesWanted := []string{}
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id:      id1,
			Aliases: []string{"one", "two"},
		},
		ManagedAliases: []string{"one", "two"},
	}

	// delete "one" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "one", "ownerType": "service" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "two" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerType": "service" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// get service
	testRequestThree := autopilot.NewTestRequest(
		`query ServiceGet($service:ID!){account{service(id: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}}}}`,
		`{ "service": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "service": { {{ template "id1" }}, "aliases": [], "managedAliases": [] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree}
	client := BestTestClient(t, "service/reconcile_aliases_delete_all", requests...)

	// Act
	err := service.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, len(service.Aliases), 0)
	autopilot.Equals(t, len(service.ManagedAliases), 0)
}

func TestServiceReconcileAliasesDeleteSome(t *testing.T) {
	// Arrange
	aliasesWanted := []string{"two"}
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id:      id1,
			Aliases: []string{"one", "two"},
		},
		ManagedAliases: []string{"one", "two"},
	}

	// delete "one" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "one", "ownerType": "service" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceGet($service:ID!){account{service(id: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}}}}`,
		`{ "service": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "service": { {{ template "id1" }}, "aliases": [ "two" ], "managedAliases": [ "two" ] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "service/reconcile_aliases_delete_some", requests...)

	// Act
	err := service.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, len(service.Aliases), 1)
	autopilot.Equals(t, len(service.ManagedAliases), 1)
}

func TestServiceReconcileAliases(t *testing.T) {
	// Arrange
	aliasesWanted := []string{"one", "two", "three"}
	service := ol.Service{
		ServiceId: ol.ServiceId{
			Id:      id1,
			Aliases: []string{"one", "alpha", "beta"},
		},
		ManagedAliases: []string{"one", "alpha", "beta"},
	}

	// delete "alpha" alias
	testRequestOne := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "alpha", "ownerType": "service" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// delete "beta" alias
	testRequestTwo := autopilot.NewTestRequest(
		`mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}`,
		`{"input":{ "alias": "beta", "ownerType": "service" }}`,
		`{"data": { "aliasDelete": {"errors": [] }}}`,
	)
	// create "two" alias
	testRequestThree := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input":{ "alias": "two", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "one", "two" ], "ownerId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)
	// create "three" alias
	testRequestFour := autopilot.NewTestRequest(
		`mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}`,
		`{"input":{ "alias": "three", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "aliasCreate": { "aliases": [ "one", "two", "three" ], "ownerId": "{{ template "id1_string" }}", "errors": [] }}}`,
	)
	// get service
	testRequestFive := autopilot.NewTestRequest(
		`query ServiceGet($service:ID!){account{service(id: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}}}}`,
		`{ "service": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "service": { {{ template "id1" }}, "aliases": ["one", "two", "three"], "managedAliases": ["one", "two", "three"] }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo, testRequestThree, testRequestFour, testRequestFive}
	client := BestTestClient(t, "service/reconcile_aliases", requests...)

	// Act
	err := service.ReconcileAliases(client, aliasesWanted)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, service.Aliases, aliasesWanted)
	autopilot.Equals(t, service.ManagedAliases, aliasesWanted)
}

func TestListServicesWithInputFilter(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "filter": [{ "key": "name", "arg": "Foo", "type": "equals", "caseSensitive": false }], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": "MQ", "endCursor": "MQ" } }}}}`,
	)
	client := BestTestClient(t, "service/list_with_input_filter", testRequest)
	// Act
	filter := ol.ServiceFilterInput{
		Arg:           "Foo",
		CaseSensitive: false,
		Key:           &ol.ServiceFilterEnumName,
		Type:          &ol.TypeEnumEquals,
	}
	response, err := client.ListServicesWithInputFilter(filter, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, response.TotalCount)
	if len(response.Nodes) > 0 {
		autopilot.Equals(t, "Foo", response.Nodes[0].Name)
		autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(response.Nodes[0].Id))
	}
}

func TestListServicesWithInputFilterAnd(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "filter": [{ "connective": "and", "caseSensitive": false, "predicates": [{ "key": "language", "arg": "Ruby", "type": "contains", "caseSensitive": false }, { "key": "tier_index", "arg": "1", "type": "equals", "caseSensitive": false }] }], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": "MQ", "endCursor": "MQ" } }}}}`,
	)
	client := BestTestClient(t, "service/list_with_input_filter_and", testRequest)

	// Act - Complex filter: language is Ruby AND tier index is 1
	predicates := []ol.ServiceFilterInput{
		{
			Arg:  "Ruby",
			Key:  &ol.ServiceFilterEnumLanguage,
			Type: &ol.TypeEnumContains,
		},
		{
			Arg:  "1",
			Key:  &ol.ServiceFilterEnumTierIndex,
			Type: &ol.TypeEnumEquals,
		},
	}

	complexFilter := ol.ServiceFilterInput{
		Connective: &ol.ConnectiveEnumAnd,
		Predicates: predicates,
	}

	response, err := client.ListServicesWithInputFilter(complexFilter, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, response.TotalCount)
}

func TestListServicesWithInputFilterOr(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "filter": [{ "connective": "or", "caseSensitive": false, "predicates": [{ "key": "language", "arg": "Ruby", "type": "contains", "caseSensitive": false }, { "key": "language", "arg": "Go", "type": "contains", "caseSensitive": false }] }], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }}, {{ template "service_2" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": "MQ", "endCursor": "Mg" } }}}}`,
	)
	client := BestTestClient(t, "service/list_with_input_filter_or", testRequest)

	// Act - OR filter: language is Ruby OR Go
	orPredicates := []ol.ServiceFilterInput{
		{
			Arg:  "Ruby",
			Key:  &ol.ServiceFilterEnumLanguage,
			Type: &ol.TypeEnumContains,
		},
		{
			Arg:  "Go",
			Key:  &ol.ServiceFilterEnumLanguage,
			Type: &ol.TypeEnumContains,
		},
	}

	orFilter := ol.ServiceFilterInput{
		Connective: &ol.ConnectiveEnumOr,
		Predicates: orPredicates,
	}

	response, err := client.ListServicesWithInputFilter(orFilter, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
}

func TestListServicesWithInputFilterNested(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "filter": [{ "connective": "and", "caseSensitive": false, "predicates": [{ "connective": "or", "caseSensitive": false, "predicates": [{ "key": "language", "arg": "Ruby", "type": "contains", "caseSensitive": false }, { "key": "language", "arg": "Go", "type": "contains", "caseSensitive": false }] }, { "key": "tier_index", "arg": "1", "type": "equals", "caseSensitive": false }] }], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": "MQ", "endCursor": "MQ" } }}}}`,
	)
	client := BestTestClient(t, "service/list_with_input_filter_nested", testRequest)

	// Act - Nested filter: (language is Ruby OR Go) AND tier index is 1
	orPredicates := []ol.ServiceFilterInput{
		{
			Arg:  "Ruby",
			Key:  &ol.ServiceFilterEnumLanguage,
			Type: &ol.TypeEnumContains,
		},
		{
			Arg:  "Go",
			Key:  &ol.ServiceFilterEnumLanguage,
			Type: &ol.TypeEnumContains,
		},
	}

	orFilter := ol.ServiceFilterInput{
		Connective: &ol.ConnectiveEnumOr,
		Predicates: orPredicates,
	}

	gigaPredicates := []ol.ServiceFilterInput{
		orFilter,
		{
			Arg:  "1",
			Key:  &ol.ServiceFilterEnumTierIndex,
			Type: &ol.TypeEnumEquals,
		},
	}

	nestedFilter := ol.ServiceFilterInput{
		Connective: &ol.ConnectiveEnumAnd,
		Predicates: gigaPredicates,
	}

	response, err := client.ListServicesWithInputFilter(nestedFilter, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, response.TotalCount)
}

func TestListServicesWithInputFilterMultiplePredicateTypes(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		// The filter is a single AND filter with multiple predicates, not a flat array
		`{ "filter": [{ "connective": "and", "caseSensitive": false, "predicates": [
				{ "key": "name", "arg": "api", "type": "contains", "caseSensitive": true },
				{ "key": "product", "arg": "test", "type": "does_not_contain", "caseSensitive": false },
				{ "key": "tier_index", "arg": "3", "type": "greater_than_or_equal_to", "caseSensitive": false }
			]}], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": "MQ", "endCursor": "MQ" } }}}}`,
	)
	client := BestTestClient(t, "service/list_with_input_filter_multiple_types", testRequest)

	// Act - Multiple filters with different predicate types (all in a single AND filter)
	predicates := []ol.ServiceFilterInput{
		{
			Arg:           "api",
			CaseSensitive: true,
			Key:           &ol.ServiceFilterEnumName,
			Type:          &ol.TypeEnumContains,
		},
		{
			Arg:  "test",
			Key:  &ol.ServiceFilterEnumProduct,
			Type: &ol.TypeEnumDoesNotContain,
		},
		{
			Arg:  "3",
			Key:  &ol.ServiceFilterEnumTierIndex,
			Type: &ol.TypeEnumGreaterThanOrEqualTo,
		},
	}

	filter := ol.ServiceFilterInput{
		Connective: &ol.ConnectiveEnumAnd,
		Predicates: predicates,
	}

	response, err := client.ListServicesWithInputFilter(filter, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, response.TotalCount)
}

func TestListServicesWithInputFilterCaseSensitivity(t *testing.T) {
	// Test case sensitive match
	testRequestMatch := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "filter": [{ "key": "name", "arg": "Foo", "type": "equals", "caseSensitive": true }], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": "MQ", "endCursor": "MQ" } }}}}`,
	)

	// Test case sensitive non-match (different case)
	testRequestNoMatch := autopilot.NewTestRequest(
		`query ServiceListWithInputFilter($after:String!$filter:[ServiceFilterInput!]!$first:Int!){account{services(filter: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "filter": [{ "key": "name", "arg": "foo", "type": "equals", "caseSensitive": true }], {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": null, "endCursor": null } }}}}`,
	)

	requests := []autopilot.TestRequest{testRequestMatch, testRequestNoMatch}
	client := BestTestClient(t, "service/list_with_input_filter_case_sensitivity", requests...)

	// Act & Assert - Test case sensitive match (should find service "Foo")
	filterMatch := ol.ServiceFilterInput{
		Arg:           "Foo",
		CaseSensitive: true,
		Key:           &ol.ServiceFilterEnumName,
		Type:          &ol.TypeEnumEquals,
	}
	response, err := client.ListServicesWithInputFilter(filterMatch, nil)
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, response.TotalCount)
	if len(response.Nodes) > 0 {
		autopilot.Equals(t, "Foo", response.Nodes[0].Name)
	}

	// Act & Assert - Test case sensitive non-match (should NOT find service with "foo" when name is "Foo")
	filterNoMatch := ol.ServiceFilterInput{
		Arg:           "foo", // lowercase - should not match "Foo" when case sensitive
		CaseSensitive: true,
		Key:           &ol.ServiceFilterEnumName,
		Type:          &ol.TypeEnumEquals,
	}
	response, err = client.ListServicesWithInputFilter(filterNoMatch, nil)
	autopilot.Ok(t, err)
	autopilot.Equals(t, 0, response.TotalCount)
}

func TestServiceGetWithHTMLEntities(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query ServiceGet($service:ID!){account{service(id: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},type{id,aliases}}}}`,
		`{ "service": "{{ template "id1_string" }}" }`,
		`{"data": {"account": {"service": {
			{{ template "id1" }},
			"aliases": ["test-service"],
			"name": "TestService",
			"description": "Service with &lt;em&gt;emphasis&lt;/em&gt; &amp; &quot;quotes&quot;",
			"htmlUrl": "https://app.opslevel.com/services/test-service",
			"note": "Additional notes: &lt;strong&gt;important&lt;/strong&gt; &amp; critical",
			"product": "Product with &amp; ampersand",
			"apiDocumentPath": null,
			"framework": null,
			"language": null,
			"lifecycle": null,
			"locked": false,
			"managedAliases": [],
			"maturityReport": null,
			"owner": null,
			"parent": null,
			"preferredApiDocument": null,
			"preferredApiDocumentSource": null,
			"repos": {"edges": [], "pageInfo": {"hasNextPage": false, "hasPreviousPage": false, "startCursor": "", "endCursor": ""}},
			"defaultServiceRepository": null,
			"tags": {"nodes": [], "pageInfo": {"hasNextPage": false, "hasPreviousPage": false, "startCursor": "", "endCursor": ""}},
			"tier": null,
			"timestamps": null,
			"tools": {"nodes": [], "pageInfo": {"hasNextPage": false, "hasPreviousPage": false, "startCursor": "", "endCursor": ""}},
			"type": null
		}}}}`,
	)

	client := BestTestClient(t, "service/html_entities", testRequest)
	// Act
	result, err := client.GetService(string(id1))
	// Assert
	autopilot.Ok(t, err)

	// Verify HTML entities are unescaped
	if strings.Contains(result.Description, "&lt;") || strings.Contains(result.Description, "&gt;") ||
		strings.Contains(result.Description, "&amp;") || strings.Contains(result.Description, "&quot;") {
		t.Errorf("Service Description still contains HTML entities: %s", result.Description)
	}
	if strings.Contains(result.Note, "&lt;") || strings.Contains(result.Note, "&gt;") ||
		strings.Contains(result.Note, "&amp;") {
		t.Errorf("Service Note still contains HTML entities: %s", result.Note)
	}
	if strings.Contains(result.Product, "&amp;") {
		t.Errorf("Service Product still contains HTML entities: %s", result.Product)
	}

	// Verify expected unescaped values
	expectedDescription := `Service with <em>emphasis</em> & "quotes"`
	expectedNote := "Additional notes: <strong>important</strong> & critical"
	expectedProduct := "Product with & ampersand"

	autopilot.Equals(t, expectedDescription, result.Description)
	autopilot.Equals(t, expectedNote, result.Note)
	autopilot.Equals(t, expectedProduct, result.Product)
}
