package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"testing"

	"github.com/rocktavious/autopilot/v2022"
)

func TestGetRepositoryWithAliasNotFound(t *testing.T) {
	// Arrange
	request := `{
	"query": "query RepositoryGet($repo:String!){account{repository(alias: $repo){archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible}}}",
	"variables":{
		"repo": "github.com:rocktavious/autopilot"
   }
}`
	response := `{"data": {
	"account": {
		"repository": null
	}
}}`
	client := ABetterTestClient(t, "repository/get_not_found", request, response)
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, *ol.NewID(), result.Id)
}

func TestGetRepositoryWithAlias(t *testing.T) {
	// Arrange
	request := `{
	"query": "query RepositoryGet($repo:String!){account{repository(alias: $repo){archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible}}}",
	"variables":{
		"repo": "github.com:rocktavious/autopilot"
    }
}`
	response := `{"data": {
	"account": {
		"repository": {{ template "repository_1" }}
	}
}}`
	client := ABetterTestClient(t, "repository/get_with_alias", request, response)
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestGetRepository(t *testing.T) {
	// Arrange
	request := `{
	"query": "query RepositoryGet($repo:ID!){account{repository(id: $repo){archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible}}}",
	"variables":{
		"repo": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk"
    }
}`
	response := `{"data": {
	"account": {
		"repository": {{ template "repository_1" }}
	}
}}`
	client := ABetterTestClient(t, "repository/get", request, response)
	// Act
	result, err := client.GetRepository("Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestListRepositories(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query RepositoryList($after:String!$first:Int!){account{repositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}",
			"variables": {
				{{ template "first_page_variables" }}
			}
			}`,
			`{
					  "data": {
						"account": {
						  "repositories": {
							"nodes": [
							  {{ template "repository_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query RepositoryList($after:String!$first:Int!){account{repositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}",
			"variables": {
				{{ template "second_page_variables" }}
			}
			}`,
			`{
					  "data": {
						"account": {
						  "repositories": {
							"nodes": [
							  {{ template "repository_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "repositories/list", requests...)
	// Act
	resp, err := client.ListRepositories(nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "autopilot", result[0].Name)
	autopilot.Equals(t, "https://github.com/opslevel/cli", result[1].Url)
}

func TestListRepositoriesWithTier(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query RepositoryListWithTier($after:String!$first:Int!$tier:String!){account{repositories(tierAlias: $tier, after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"tier": "tier_1"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "repositories": {
							"nodes": [
							  {{ template "repository_1" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
		{`{"query": "query RepositoryListWithTier($after:String!$first:Int!$tier:String!){account{repositories(tierAlias: $tier, after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"tier": "tier_1"
			}
			}`,
			`{
					  "data": {
						"account": {
						  "repositories": {
							"nodes": [
							  {{ template "repository_2" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }
						}
					  }
					}`},
	}
	client := APaginatedTestClient(t, "repositories/list_with_tier", requests...)
	// Act
	resp, err := client.ListRepositoriesWithTier("tier_1", nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "autopilot", result[0].Name)
	autopilot.Equals(t, "https://github.com/opslevel/cli", result[1].Url)
}

func TestDeleteServiceRepository(t *testing.T) {
	// Arrange
	request := `{
	"query": "mutation ServiceRepositoryDelete($input:DeleteInput!){serviceRepositoryDelete(input: $input){deletedId,errors{message,path}}}",
	"variables":{
		"input": {
			"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3"
		}
    }
}`
	response := `{"data": {
	"serviceRepositoryDelete": {
		"deletedId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3",
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "repository/service_delete", request, response)
	// Act
	err := client.DeleteServiceRepository("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3")
	// Assert
	autopilot.Ok(t, err)
}

func TestGetServices(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query RepositoryServicesList($after:String!$first:Int!$id:ID!){account{repository(id: $id){services(after: $after, first: $first){edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"
			}
			}`,
			`{
				  "data": {
					"account": {
					  "repository": {
						"services": {
						  "edges": [
							{
							  "atRoot": true,
							  "node": {
								"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8xODc1",
								"aliases": [
								  "apple-running-app",
								  "catalog_service_4"
								]
							  },
							  "paths": [
								{
								  "href": "https://bitbucket.org/raptors-store/catalogue",
								  "path": ""
								}
							  ],
							  "serviceRepositories": [
								{
								  "baseDirectory": "",
								  "displayName": "raptors-store/Catalogue",
								  "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvMjQ1Mg",
								  "repository": {
									"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYx",
									"defaultAlias": "bitbucket.org:raptors-store/Catalogue"
								  },
								  "service": {
									"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8xODc1",
									"aliases": [
									  "apple-running-app",
									  "catalog_service_4"
									]
								  }
								}
							  ]
							},
							{
							  "atRoot": true,
							  "node": {
								"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS83NDc",
								"aliases": [
								  "Catalog_Shopping_test",
								  "Catalog_test_service",
								  "XYZ    DEF",
								  "XYZ DEF",
								  "catalog_service_2",
								  "no_service_alias",
								  "service_for_catalogue_repo",
								  "xyz_service_2",
								  "xyz_service_4"
								]
							  },
							  "paths": [
								{
								  "href": "https://bitbucket.org/raptors-store/catalogue",
								  "path": ""
								}
							  ],
							  "serviceRepositories": [
								{
								  "baseDirectory": "",
								  "displayName": "raptors-store/Catalogue",
								  "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvMTg1",
								  "repository": {
									"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYx",
									"defaultAlias": "bitbucket.org:raptors-store/Catalogue"
								  },
								  "service": {
									"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS83NDc",
									"aliases": [
									  "Catalog_Shopping_test",
									  "Catalog_test_service",
									  "XYZ    DEF",
									  "XYZ DEF",
									  "catalog_service_2",
									  "no_service_alias",
									  "service_for_catalogue_repo",
									  "xyz_service_2",
									  "xyz_service_4"
									]
								  }
								}
							  ]
							}
						  ],
						  {{ template "pagination_initial_pageInfo_response" }},
						  "totalCount": 2
						}
					  }
					}
				  }
				}`},
		{`{"query": "query RepositoryServicesList($after:String!$first:Int!$id:ID!){account{repository(id: $id){services(after: $after, first: $first){edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"
			}
			}`,
			`{
				  "data": {
					"account": {
					  "repository": {
						"services": {
						  "edges": [
							{
							  "atRoot": true,
							  "node": {
								"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8zMQ",
								"aliases": [
								  "Back End",
								  "Backend Service",
								  "a/b/c",
								  "back end testing",
								  "back_end",
								  "fs-prod:deployment/bolt-http",
								  "shopping_barts",
								  "shopping_cart_service",
								  "testing_1",
								  "testing_11",
								  "testing_12",
								  "testing_123",
								  "testing_1234",
								  "testing_15",
								  "testing_2",
								  "testing_3",
								  "testing_4",
								  "testing_5",
								  "testing_6",
								  "testing_8"
								]
							  },
							  "paths": [
								{
								  "href": "https://bitbucket.org/raptors-store/catalogue",
								  "path": ""
								}
							  ],
							  "serviceRepositories": [
								{
								  "baseDirectory": "",
								  "displayName": "raptors-store/Catalogue",
								  "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvMjYyMg",
								  "repository": {
									"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYx",
									"defaultAlias": "bitbucket.org:raptors-store/Catalogue"
								  },
								  "service": {
									"id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8zMQ",
									"aliases": [
									  "Back End",
									  "Backend Service",
									  "a/b/c",
									  "back end testing",
									  "back_end",
									  "fs-prod:deployment/bolt-http",
									  "shopping_barts",
									  "shopping_cart_service",
									  "testing_1",
									  "testing_11",
									  "testing_12",
									  "testing_123",
									  "testing_1234",
									  "testing_15",
									  "testing_2",
									  "testing_3",
									  "testing_4",
									  "testing_5",
									  "testing_6",
									  "testing_8"
									]
								  }
								}
							  ]
							}
						  ],
						  {{ template "pagination_second_pageInfo_response" }},
						  "totalCount": 1
						}
					  }
					}
				  }
				}`},
	}
	client := APaginatedTestClient(t, "repository/services", requests...)
	// Act
	repository := ol.Repository{
		Id: "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc",
	}
	resp, err := repository.GetServices(client, nil)
	result := resp.Edges
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, *ol.NewID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8xODc1"), result[0].Node.Id)
	autopilot.Equals(t, *ol.NewID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS83NDc"), result[1].Node.Id)
	autopilot.Equals(t, *ol.NewID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8zMQ"), result[2].Node.Id)
}

func TestGetTags(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query RepositoryTagsList($after:String!$first:Int!$id:ID!){account{repository(id: $id){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},
				"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"
			}
			}`,
			`{
				  "data": {
					"account": {
					  "repository": {
						"tags": {
						  "nodes": [
							{
							  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIw",
							  "key": "abc",
							  "value": "abc"
							},
							{
							  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIx",
							  "key": "db",
							  "value": "mongoqqqq"
							},
							{
							  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIy",
							  "key": "db",
							  "value": "prod"
							}
						  ],
						  {{ template "pagination_initial_pageInfo_response" }},
						  "totalCount": 3
						}
					  }
					}
				  }
				}`},
		{`{"query": "query RepositoryTagsList($after:String!$first:Int!$id:ID!){account{repository(id: $id){tags(after: $after, first: $first){nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},
				"id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc"
			}
			}`,
			`{
				  "data": {
					"account": {
					  "repository": {
						"tags": {
						  "nodes": [
							{
							  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIz",
							  "key": "env",
							  "value": "staging"
							}
						  ],
						  {{ template "pagination_second_pageInfo_response" }},
						  "totalCount": 1
						}
					  }
					}
				  }
				}`},
	}
	client := APaginatedTestClient(t, "repository/tags", requests...)
	// Act
	repository := ol.Repository{
		Id: "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc",
	}
	resp, err := repository.GetTags(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, resp.TotalCount)
	autopilot.Equals(t, "abc", result[0].Key)
	autopilot.Equals(t, "abc", result[0].Value)
	autopilot.Equals(t, "db", result[2].Key)
	autopilot.Equals(t, "prod", result[2].Value)
	autopilot.Equals(t, "env", result[3].Key)
	autopilot.Equals(t, "staging", result[3].Value)
}
