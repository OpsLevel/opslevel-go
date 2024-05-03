package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

func TestServiceTags(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceTagsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "service": { "tags": { "nodes": [ { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA5", "key": "prod", "value": "false" } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceTagsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "service": { "tags": { "nodes": [ { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA4", "key": "test", "value": "true" } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
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
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "prod", result[0].Key)
	autopilot.Equals(t, "false", result[0].Value)
	autopilot.Equals(t, "test", result[1].Key)
	autopilot.Equals(t, "true", result[1].Value)
}

func TestServiceTools(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceToolsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tools(after: $after, first: $first){nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}}}}`,
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
                                {{ template "pagination_initial_pageInfo_response" }},
                                "totalCount": 1
                            }
                          }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceToolsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tools(after: $after, first: $first){nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}}}}`,
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
                                {{ template "pagination_second_pageInfo_response" }},
                                "totalCount": 1
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
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "PagerDuty", result[0].DisplayName)
	autopilot.Equals(t, ol.ToolCategoryContinuousIntegration, result[1].Category)
}

func TestServiceRepositories(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repos(after: $after, first: $first){edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount}}}}`,
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
                                {{ template "pagination_initial_pageInfo_response" }},
                                "totalCount": 2
                            }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repos(after: $after, first: $first){edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount}}}}`,
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
                                {{ template "pagination_second_pageInfo_response" }},
                                "totalCount": 1
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
	result := resp.Edges
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw", string(result[0].Node.Id))
	autopilot.Equals(t, "bitbucket.org:raptors-store/Store Front", result[1].Node.DefaultAlias)
}

func TestCreateService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceCreate($input:ServiceCreateInput!){serviceCreate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
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

func TestCreateServiceWithParentSystem(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceCreate($input:ServiceCreateInput!){serviceCreate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
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
	autopilot.Equals(t, 1, len(result.Aliases))
}

func TestUpdateService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceUpdate($input:ServiceUpdateInput!){serviceUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
		`{"input":{"id": "123456789"}}`,
		`{"data": {"serviceUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/update", testRequest)

	// Act
	result, err := client.UpdateService(ol.ServiceUpdateInput{
		Id: ol.NewID("123456789"),
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Foo", result.Name)
}

func TestNewUpdateService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceUpdate($input:ServiceUpdateInputV2!){serviceUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
		`{"input":{"id": "123456789"}}`,
		`{"data": {"serviceUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/new_update", testRequest)

	// Act
	result, err := client.UpdateServiceV2(ol.ServiceUpdateInputV2{
		Id: ol.NewID("123456789"),
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Foo", result.Name)
}

func TestUpdateServiceWithSystem(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceUpdate($input:ServiceUpdateInput!){serviceUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
		`{"input":{"id": "123456789", "parent": {"alias": "FooSystem"}}}`,
		`{"data": {"serviceUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/update_with_system", testRequest)

	// Act
	result, err := client.UpdateService(ol.ServiceUpdateInput{
		Id:     ol.NewID("123456789"),
		Parent: ol.NewIdentifier("FooSystem"),
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Foo", result.Name)
}

func TestNewUpdateServiceWithFields(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceUpdate($input:ServiceUpdateInputV2!){serviceUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
		`{"input":{"id": "123456789", "language": null, "lifecycleAlias": "pre-alpha", "parent": {"alias": "FooSystem"}, "tierAlias": null}}`,
		`{"data": {"serviceUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/new_update_with_fields", testRequest)

	// Act
	result, err := client.UpdateServiceV2(ol.ServiceUpdateInputV2{
		Framework:      nil, // will do nothing - not included in request body
		Id:             ol.NewID("123456789"),
		Language:       ol.NewNullString(), // will unset the field - becomes `null` in request body
		LifecycleAlias: ol.NewOptionalString("pre-alpha"),
		Parent:         ol.NewIdentifier("FooSystem"),
		TierAlias:      ol.NewOptionalString(""), // will unset the field - becomes `null` in request body
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Foo", result.Name)
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
		`query ServiceGet($service:String!){account{service(alias: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}}}}`,
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
          {{ template "pagination_response_same_cursor" }},
          "totalCount": 1
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
          {{ template "pagination_response_different_cursor" }},
          "totalCount": 3
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
          {{ template "no_pagination_response" }},
          "totalCount": 4
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
		`query ServiceGet($service:ID!){account{service(id: $service){apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}}}}`,
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
          {{ template "pagination_response_same_cursor" }},
          "totalCount": 1
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
          {{ template "pagination_response_different_cursor" }},
          "totalCount": 3
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
          {{ template "no_pagination_response" }},
          "totalCount": 4
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
}

func TestGetServiceDocuments(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceDocumentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){documents(after: $after, first: $first){nodes{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ "service": "{{ template "id1_string" }}", {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "service": { "documents": { "nodes": [ {{ template "document_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceDocumentsList($after:String!$first:Int!$service:ID!){account{service(id: $service){documents(after: $after, first: $first){nodes{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ "service": "{{ template "id1_string" }}", {{ template "second_page_variables" }} }`,
		`{ "data": { "account": { "service": { "documents": { "nodes": [ {{ template "document_1" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
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

func TestListServices(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceList($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceList($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }} }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list", requests...)
	// Act
	resp, err := client.ListServices(nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithFilter(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithFilter($after:String!$filter:IdentifierInput$first:Int!){account{services(filterIdentifier: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "filter": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithFilter($after:String!$filter:IdentifierInput$first:Int!){account{services(filterIdentifier: $filter, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "filter": { {{ template "id1" }} } }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_filter", requests...)
	// Act
	response, err := client.ListServicesWithFilter(string(id1), nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithFramework(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithFramework($after:String!$first:Int!$framework:String!){account{services(framework: $framework, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "framework": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithFramework($after:String!$first:Int!$framework:String!){account{services(framework: $framework, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "framework": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_framework", requests...)
	// Act
	response, err := client.ListServicesWithFramework("postgres", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithLanguage(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithLanguage($after:String!$first:Int!$language:String!){account{services(language: $language, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "language": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithLanguage($after:String!$first:Int!$language:String!){account{services(language: $language, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "language": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_language", requests...)
	// Act
	response, err := client.ListServicesWithLanguage("postgres", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithOwner(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithOwner($after:String!$first:Int!$owner:String!){account{services(ownerAlias: $owner, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "owner": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithOwner($after:String!$first:Int!$owner:String!){account{services(ownerAlias: $owner, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "owner": "postgres" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_owner", requests...)
	// Act
	response, err := client.ListServicesWithOwner("postgres", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithTag(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithTag($after:String!$first:Int!$tag:TagArgs!){account{services(tag: $tag, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "tag": { "key": "app", "value": "worker" } }`,
		`{"data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithTag($after:String!$first:Int!$tag:TagArgs!){account{services(tag: $tag, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "tag": { "key": "app", "value": "worker" } }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_tag", requests...)
	// Act
	tagArgs, err := ol.NewTagArgs("app:worker")
	autopilot.Ok(t, err)
	response, err := client.ListServicesWithTag(tagArgs, nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithTier(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithTier($after:String!$first:Int!$tier:String!){account{services(tierAlias: $tier, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "tier": "tier_1" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithTier($after:String!$first:Int!$tier:String!){account{services(tierAlias: $tier, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "tier": "tier_1" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_tier", requests...)
	// Act
	response, err := client.ListServicesWithTier("tier_1", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithLifecycle(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithLifecycle($after:String!$first:Int!$lifecycle:String!){account{services(lifecycleAlias: $lifecycle, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "lifecycle": "alpha" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithLifecycle($after:String!$first:Int!$lifecycle:String!){account{services(lifecycleAlias: $lifecycle, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "lifecycle": "alpha" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_lifecycle", requests...)
	// Act
	response, err := client.ListServicesWithLifecycle("alpha", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithProduct(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query ServiceListWithProduct($after:String!$first:Int!$product:String!){account{services(product: $product, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "first_page_variables" }}, "product": "test" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query ServiceListWithProduct($after:String!$first:Int!$product:String!){account{services(product: $product, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},managedAliases,name,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ {{ template "second_page_variables" }}, "product": "test" }`,
		`{ "data": { "account": { "services": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "service/list_with_product", requests...)
	// Act
	response, err := client.ListServicesWithProduct("test", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
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
