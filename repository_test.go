package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"

	"github.com/rocktavious/autopilot/v2023"
)

func TestConnectServiceRepository(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceRepositoryCreate($input:ServiceRepositoryCreateInput!){serviceRepositoryCreate(input: $input){serviceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},errors{message,path}}}`,
		`{ "input": { "service": { {{ template "id1" }} }, "repository": { {{ template "id1" }} }, "baseDirectory": "/", "displayName": "OpsLevel/opslevel" }}`,
		`{"data": {
    "serviceRepositoryCreate": {
        "serviceRepository": {
            "baseDirectory": "/",
            "displayName": "OpsLevel/opslevel",
            {{ template "id1" }},
            "repository": {
                {{ template "id2" }},
                "defaultAlias": "{{ template "alias1" }}"
            },
            "service": {
                {{ template "id3" }},
                "aliases": [
                  "{{ template "alias1" }}",
                  "{{ template "alias2" }}"
                ]
            }
        },
        "errors": []
    }}}`,
	)
	client := BestTestClient(t, "repository/connect", testRequest)
	service := ol.ServiceId{
		Id: id1,
	}
	repository := ol.Repository{
		Id:           id1,
		Name:         "opslevel",
		Organization: "OpsLevel",
	}
	// Act
	result, err := client.ConnectServiceRepository(&service, &repository)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "OpsLevel/opslevel", result.DisplayName)
}

func TestGetRepositoryWithAliasNotFound(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query RepositoryGet($repo:String!){account{repository(alias: $repo){archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible}}}`,
		`{ "repo": "github.com:rocktavious/autopilot" }`,
		`{"data": { "account": { "repository": null }}}`,
	)
	client := BestTestClient(t, "repository/get_not_found", testRequest)
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, *ol.NewID(), result.Id)
}

func TestGetRepositoryWithAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query RepositoryGet($repo:String!){account{repository(alias: $repo){archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible}}}`,
		`{"repo": "github.com:rocktavious/autopilot" }`,
		`{"data": { "account": { "repository": {{ template "repository_1" }} }}}`,
	)

	client := BestTestClient(t, "repository/get_with_alias", testRequest)
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
	testRequest := autopilot.NewTestRequest(
		`query RepositoryGet($repo:ID!){account{repository(id: $repo){archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible}}}`,
		`{"repo": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk" }`,
		`{"data": { "account": { "repository": {{ template "repository_1" }} }}}`,
	)

	client := BestTestClient(t, "repository/get", testRequest)
	// Act
	result, err := client.GetRepository("Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, true, result.Locked)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestListRepositories(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query RepositoryList($after:String!$first:Int!$visible:Boolean!){account{repositories(after: $after, first: $first, visible: $visible){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}`,
		`{ {{ template "first_page_variables" }}, "visible": true }`,
		`{ "data": { "account": { "repositories": { "nodes": [ {{ template "repository_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query RepositoryList($after:String!$first:Int!$visible:Boolean!){account{repositories(after: $after, first: $first, visible: $visible){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}`,
		`{ {{ template "second_page_variables" }}, "visible": true }`,
		`{ "data": { "account": { "repositories": { "nodes": [ {{ template "repository_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "repositories/list", requests...)
	// Act
	resp, err := client.ListRepositories(nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "autopilot", result[0].Name)
	autopilot.Equals(t, true, result[0].Locked)
	autopilot.Equals(t, false, result[1].Locked)
	autopilot.Equals(t, "https://github.com/opslevel/cli", result[1].Url)
}

func TestListRepositoriesWithTier(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query RepositoryListWithTier($after:String!$first:Int!$tier:String!){account{repositories(tierAlias: $tier, after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}`,
		`{ {{ template "first_page_variables" }}, "tier": "tier_1" }`,
		`{ "data": { "account": { "repositories": { "nodes": [ {{ template "repository_1" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query RepositoryListWithTier($after:String!$first:Int!$tier:String!){account{repositories(tierAlias: $tier, after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}`,
		`{ {{ template "second_page_variables" }}, "tier": "tier_1" }`,
		`{ "data": { "account": { "repositories": { "nodes": [ {{ template "repository_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "repositories/list_with_tier", requests...)
	// Act
	resp, err := client.ListRepositoriesWithTier("tier_1", nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "autopilot", result[0].Name)
	autopilot.Equals(t, "https://github.com/opslevel/cli", result[1].Url)
}

func TestUpdateRepository(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RepositoryUpdate($input:RepositoryUpdateInput!){repositoryUpdate(input: $input){repository{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},errors{message,path}}}`,
		`{ "input": { {{ template "id1" }}, "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": { "repositoryUpdate": { "repository": {{ template "repository_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "repositories/update", testRequest)
	// Act
	resp, err := client.UpdateRepository(ol.RepositoryUpdateInput{
		Id:      id1,
		OwnerId: ol.NewNullableFrom(ol.ID(string(id1))),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "developers", resp.Owner.Alias)
}

func TestRepositoryUpdateOwnerNotPresent(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RepositoryUpdate($input:RepositoryUpdateInput!){repositoryUpdate(input: $input){repository{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},errors{message,path}}}`,
		`{"input": { {{ template "id1" }} }}`,
		`{"data": { "repositoryUpdate": { "repository": {{ template "repository_2" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "repositories/update_owner_not_present", testRequest)
	// Act
	resp, err := client.UpdateRepository(ol.RepositoryUpdateInput{
		Id: *ol.NewID(string(id1)),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "platform", resp.Owner.Alias)
}

func TestRepositoryUpdateOwnerNull(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation RepositoryUpdate($input:RepositoryUpdateInput!){repositoryUpdate(input: $input){repository{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},errors{message,path}}}`,
		`{"input": { {{ template "id1" }}, "ownerId": null }}`,
		`{"data": { "repositoryUpdate": { "repository": {{ template "repository_3" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "repositories/update_owner_null", testRequest)
	// Act
	resp, err := client.UpdateRepository(ol.RepositoryUpdateInput{
		Id:      *ol.NewID(string(id1)),
		OwnerId: ol.NewNullOf[ol.ID](),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "", string(resp.Owner.Id))
}

func TestUpdateServiceRepository(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceRepositoryUpdate($input:ServiceRepositoryUpdateInput!){serviceRepositoryUpdate(input: $input){serviceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},errors{message,path}}}`,
		`{ "input": { {{ template "id1" }}, "displayName": "Foobar" }}`,
		`{"data": {
    "serviceRepositoryUpdate": {
        "serviceRepository": {
            "baseDirectory": "",
            "displayName": "Foobar",
            {{ template "id1" }},
            "repository": {
                {{ template "id2" }},
                "defaultAlias": "{{ template "alias1" }}"
            },
            "service": {
                {{ template "id3" }},
                "aliases": [
                  "{{ template "alias1" }}",
                  "{{ template "alias2" }}"
                ]
            }
        },
        "errors": []
    }}}`,
	)

	client := BestTestClient(t, "repository/service_update", testRequest)
	// Act
	resp, err := client.UpdateServiceRepository(ol.ServiceRepositoryUpdateInput{
		Id:          id1,
		DisplayName: ol.NewNullableFrom("Foobar"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Foobar", resp.DisplayName)
}

func TestDeleteServiceRepository(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceRepositoryDelete($input:DeleteInput!){serviceRepositoryDelete(input: $input){deletedId,errors{message,path}}}`,
		`{"input": { "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3" }}`,
		`{"data": { "serviceRepositoryDelete": { "deletedId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3", "errors": [] }}}`,
	)

	client := BestTestClient(t, "repository/service_delete", testRequest)
	// Act
	err := client.DeleteServiceRepository("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3")
	// Assert
	autopilot.Ok(t, err)
}

func TestGetServices(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query RepositoryServicesList($after:String!$first:Int!$id:ID!){account{repository(id: $id){services(after: $after, first: $first){edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc" }`,
		`{ "data": {
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
    }`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query RepositoryServicesList($after:String!$first:Int!$id:ID!){account{repository(id: $id){services(after: $after, first: $first){edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc" }`,
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
    }`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "repository/services", requests...)
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
	testRequestOne := autopilot.NewTestRequest(
		`query RepositoryTagsList($after:String!$first:Int!$id:ID!){account{repository(id: $id){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc" }`,
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
    }`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query RepositoryTagsList($after:String!$first:Int!$id:ID!){account{repository(id: $id){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRsYWIvMTA5ODc" }`,
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
    }`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "repository/tags", requests...)
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
