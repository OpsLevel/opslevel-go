package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
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
	autopilot.Equals(t, &ol.Repository{}, result)
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
	client := ATestClient(t, "repository/list")
	// Act
	result, err := client.ListRepositories()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
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
