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
		{`{"query": "query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
					"account": {
						"service": {
							"repositories": {
								"nodes": [
									{{ template "repository_1" }}
								],
								{{ template "pagination_initial_pageInfo_response" }},
								"totalCount": 1
							}
						  }}}}`},
		{`{"query": "query ServiceRepositoriesList($after:String!$first:Int!$service:ID!){account{service(id: $service){repositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"service": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4"
			}
			}`,
			`{
				"data": {
					"account": {
						"service": {
							"repositories": {
								"nodes": [
									{{ template "repository_2" }}
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
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "https://github.com/rocktavious/autopilot", result[0].Url)
	autopilot.Equals(t, "cli", result[1].Name)
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
	client := ATestClient(t, "service/list")
	// Act
	result, err := client.ListServices()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
	autopilot.Equals(t, "API Docs", result[0].PreferredApiDocument.Source.Name)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPush, *result[0].PreferredApiDocumentSource)
}

func TestListServicesWithFramework(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_framework")
	// Act
	result, err := client.ListServicesWithFramework("postgres")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithLanguage(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_language")
	// Act
	result, err := client.ListServicesWithLanguage("postgres")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithOwner(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_owner")
	// Act
	result, err := client.ListServicesWithOwner("postgres")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithTag(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_tag")
	// Act
	result, err := client.ListServicesWithTag(ol.NewTagArgs("app:worker"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
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
