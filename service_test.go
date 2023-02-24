package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

func TestServiceTags(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ServiceTagsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
					"account": {
						"service": {
							"tags": {
								"nodes": [
									{
									  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA5",
									  "key": "prod",
									  "value": "false"
									}
								],
								{{ template "pagination_initial_pageInfo_response" }},
								"totalCount": 1
							}
						  }}}}`},
		{`{"query": "query ServiceTagsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
					"account": {
						"service": {
							"tags": {
								"nodes": [
									{
									  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA4",
									  "key": "test",
									  "value": "true"
									}
								],
								{{ template "pagination_second_pageInfo_response" }},
								"totalCount": 1
							}
						  }}}}`},
	}
	client := APaginatedTestClient(t, "service/tags", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceToolsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tools(after: $after, first: $first){nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
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
						  }}}}`},
		{`{"query": "query ServiceToolsList($after:String!$first:Int!$service:ID!){account{service(id: $service){tools(after: $after, first: $first){nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
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
							}
						  }}}}`},
	}
	client := APaginatedTestClient(t, "service/tools", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repos(after: $after, first: $first){edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
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
							}
						  }}}}`},
		{`{"query": "query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repos(after: $after, first: $first){edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
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
							}
						  }}}}`},
	}
	client := APaginatedTestClient(t, "service/repositories", requests...)
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
	client := ATestClient(t, "service/create")
	// Act
	result, err := client.CreateService(ol.ServiceCreateInput{
		Name:        "Foo",
		Description: "Foo service",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result.Aliases))
}

func TestUpdateService(t *testing.T) {
	//Arrange
	request := `{"query":
		"mutation ServiceUpdate($input:ServiceUpdateInput!){serviceUpdate(input: $input){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},errors{message,path}}}",
		"variables":{"input":{"id": "123456789"}}
	}`
	response := `{"data": {"serviceUpdate": {
     "service": {{ template "service_1" }},
     "errors": []
 }}}`

	client := ABetterTestClient(t, "service/update", request, response)

	// Act
	_, err := client.UpdateService(ol.ServiceUpdateInput{
		Id: "123456789",
	})

	// Assert
	autopilot.Ok(t, err)
	//autopilot.Equals(t, "Deploy Rollback", action.Name)
}

func TestGetServiceIdWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/get_id", "service/get_id_with_alias")
	// Act
	result, err := client.GetServiceIdWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", string(result.Id))
}

func TestGetServiceWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/get", "service/get_with_alias")
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
}

func TestGetService(t *testing.T) {
	// Arrange
	client := ATestClient(t, "service/get")
	client2 := ATestClient(t, "service/get_documents")
	// Act
	result, err := client.GetService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx")
	docs, err := result.Documents(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, 1, len(docs))
	autopilot.Equals(t, "", docs[0].HtmlURL)
}

func TestListServices(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ServiceList($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }}
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceList($after:String!$first:Int!){account{services(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }}
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list", requests...)
	// Act
	resp, err := client.ListServices(nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithFramework(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ServiceListWithFramework($after:String!$first:Int!$framework:String!){account{services(framework: $framework, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"framework": "postgres"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithFramework($after:String!$first:Int!$framework:String!){account{services(framework: $framework, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"framework": "postgres"
				}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_framework", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceListWithLanguage($after:String!$first:Int!$language:String!){account{services(language: $language, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"language": "postgres"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithLanguage($after:String!$first:Int!$language:String!){account{services(language: $language, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"language": "postgres"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_language", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceListWithOwner($after:String!$first:Int!$owner:String!){account{services(ownerAlias: $owner, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"owner": "postgres"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithOwner($after:String!$first:Int!$owner:String!){account{services(ownerAlias: $owner, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"owner": "postgres"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_owner", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceListWithTag($after:String!$first:Int!$tag:TagArgs!){account{services(tag: $tag, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
			  {{ template "first_page_variables" }},
			  "tag": {
				"key": "app",
				"value": "worker"
			  }
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithTag($after:String!$first:Int!$tag:TagArgs!){account{services(tag: $tag, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
			  {{ template "second_page_variables" }},
			  "tag": {
				"key": "app",
				"value": "worker"
			  }
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_tag", requests...)
	// Act
	response, err := client.ListServicesWithTag(ol.NewTagArgs("app:worker"), nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, response.TotalCount)
	autopilot.Equals(t, "Foo", result[0].Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result[1].PreferredApiDocumentSource)
}

func TestListServicesWithTier(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query ServiceListWithTier($after:String!$first:Int!$tier:String!){account{services(tierAlias: $tier, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"tier": "tier_1"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithTier($after:String!$first:Int!$tier:String!){account{services(tierAlias: $tier, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"tier": "tier_1"
				}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_tier", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceListWithLifecycle($after:String!$first:Int!$lifecycle:String!){account{services(lifecycleAlias: $lifecycle, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"lifecycle": "alpha"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithLifecycle($after:String!$first:Int!$lifecycle:String!){account{services(lifecycleAlias: $lifecycle, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"lifecycle": "alpha"
				}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_lifecycle", requests...)
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
	requests := []TestRequest{
		{`{"query": "query ServiceListWithProduct($after:String!$first:Int!$product:String!){account{services(product: $product, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"product": "test"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query ServiceListWithProduct($after:String!$first:Int!$product:String!){account{services(product: $product, after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"product": "test"
				}
			}`,
			`{
					  "data": {
						"account": {
						  "services": {
							"nodes": [
							  {{ template "service_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "service/list_with_product", requests...)
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
	client := ATestClient(t, "service/delete")
	// Act
	err := client.DeleteService(ol.ServiceDeleteInput{Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3"})
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteServicesWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/delete", "service/delete_with_alias")
	// Act
	err := client.DeleteServiceWithAlias("db")
	// Assert
	autopilot.Ok(t, err)
}
