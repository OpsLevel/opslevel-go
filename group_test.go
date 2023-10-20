package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

// TODO: not sure if there is a better way to handle reusing a client
// Probably should be a feature of autopilot
var getGroupWithAliasClient *ol.Client

func getGroupWithAliasTestClient(t *testing.T) *ol.Client {
	if getGroupWithAliasClient == nil {
		getGroupWithAliasClient = ATestClientAlt(t, "group/get", "group/get_with_alias")
	}
	return getGroupWithAliasClient
}

func TestCreateGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/create")
	members := []ol.MemberInput{
		{Email: "edgar+test@opslevel.com"},
	}
	teams := []ol.IdentifierInput{
		{Alias: "platform"},
	}
	// Act

	result, err := client.CreateGroup(ol.GroupInput{
		Name:        "platform",
		Description: "Another test group",
		Members:     &members,
		Parent:      ol.NewIdentifier("test_group_1"),
		Teams:       &teams,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "platform", result.Name)
	autopilot.Equals(t, "Another test group", result.Description)
	autopilot.Equals(t, "test_group_1", result.Parent.Alias)
}

func TestDeleteGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/delete")
	// Act
	err := client.DeleteGroup("Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTc")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteGroupWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "group/delete", "group/delete_with_alias")
	// Act
	err := client.DeleteGroup("platform")
	// Assert
	autopilot.Ok(t, err)
}

func TestChildTeams(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query GroupChildTeamsList($after:String!$first:Int!$group:ID!){account{group(id: $group){childTeams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "group": "123456789" }`,
		`{ "data": { "account": { "group": { "childTeams": { "nodes": [ {{ template "team_1" }}, {{ template "team_2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query GroupChildTeamsList($after:String!$first:Int!$group:ID!){account{group(id: $group){childTeams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "group": "123456789" }`,
		`{ "data": { "account": { "group": { "childTeams": { "nodes": [ {{ template "team_3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "group/child_teams", requests...)
	group := ol.Group{
		GroupId: ol.GroupId{
			Id: "123456789",
		},
	}
	// Act
	resp, err := group.ChildTeams(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, "example_3", result[2].Alias)
}

func TestDescendantTeams(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query GroupDescendantTeamsList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantTeams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantTeams": { "nodes": [ {{ template "team_1" }}, {{ template "team_2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query GroupDescendantTeamsList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantTeams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantTeams": { "nodes": [ {{ template "team_3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "group/descendant_teams", requests...)
	// Act
	group := ol.Group{
		GroupId: ol.GroupId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
		},
	}
	resp, err := group.DescendantTeams(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, "example_3", result[2].Alias)
}

func TestDescendantRepositories(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query GroupDescendantRepositoriesList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantRepositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantRepositories": { "nodes": [ {{ template "repository_1"}} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query GroupDescendantRepositoriesList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantRepositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantRepositories": { "nodes": [ {{ template "repository_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "group/descendant_repositories", requests...)
	// Act
	group := ol.Group{
		GroupId: ol.GroupId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
		},
	}
	resp, err := group.DescendantRepositories(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "github.com:rocktavious/autopilot", result[0].DefaultAlias)
	autopilot.Equals(t, "github.com:OpsLevel/cli", result[1].DefaultAlias)
}

func TestDescendantServices(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query GroupDescendantServicesList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantServices": { "nodes": [ {{ template "service_1"}} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query GroupDescendantServicesList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantServices(after: $after, first: $first){nodes{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantServices": { "nodes": [ {{ template "service_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}
	client := BestTestClient(t, "group/descendant_services", requests...)

	// Act
	group := ol.Group{
		GroupId: ol.GroupId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
		},
	}
	resp, err := group.DescendantServices(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "example", result[0].Aliases[0])
	autopilot.Equals(t, "example_2", result[1].Aliases[0])
}

func TestDescendantSubgroups(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query GroupDescendantSubgroupsList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantSubgroups(after: $after, first: $first){nodes{alias,id,description,htmlUrl,name,parent{alias,id}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantSubgroups": { "nodes": [ {{ template "group_1"}} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query GroupDescendantSubgroupsList($after:String!$first:Int!$group:ID!){account{group(id: $group){descendantSubgroups(after: $after, first: $first){nodes{alias,id,description,htmlUrl,name,parent{alias,id}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "descendantSubgroups": { "nodes": [ {{ template "group_2" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "group/descendant_subgroups", requests...)
	// Act
	group := ol.Group{
		GroupId: ol.GroupId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
		},
	}
	resp, err := group.DescendantSubgroups(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "test_group_1", result[0].Alias)
	autopilot.Equals(t, "test_group_2", result[1].Alias)
}

func TestGetGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/get")
	// Act
	result, err := client.GetGroup("Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Alias)
	autopilot.Equals(t, "", string(result.Parent.Id))
}

func TestGetGroupWithAlias(t *testing.T) {
	// Arrange
	client := getGroupWithAliasTestClient(t)
	// Act
	result, err := client.GetGroupWithAlias("test_group_1")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Alias)
	autopilot.Equals(t, "", string(result.Parent.Id))
}

func TestListGroups(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query ($after:String!$first:Int!){account{groups(after: $after, first: $first){nodes{alias,id,description,htmlUrl,name,parent{alias,id}},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": {
    "account": {
      "groups": {
        "nodes": [
          {
            "alias": "test_group_2",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTE",
            "description": "test1123",
            "htmlUrl": "https://app.opslevel-staging.com/groups/test_group_2",
            "name": "test_group_2",
            "parent": {
              "alias": "test_group_1",
              "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI"
            }
          },
          {
            "alias": "test_group_1",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
            "description": null,
            "htmlUrl": "https://app.opslevel-staging.com/groups/test_group_1",
            "name": "test_group_1",
            "parent": null
          }
        ],
        {{ template "pagination_initial_pageInfo_response" }},
        "totalCount": 2
      }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query ($after:String!$first:Int!){account{groups(after: $after, first: $first){nodes{alias,id,description,htmlUrl,name,parent{alias,id}},{{ template "pagination_request" }},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{"data": {
    "account": {
      "groups": {
        "nodes": [
          {
            "alias": "test_group_3",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvNTE",
            "description": "third test group",
            "htmlUrl": "https://app.opslevel-staging.com/groups/test_group_3",
            "name": "test_group_3",
            "parent": {
              "alias": "test_group_1",
              "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI"
            }
          }
        ],
        {{ template "pagination_second_pageInfo_response" }},
        "totalCount": 1
      }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "group/list", requests...)
	// Act
	response, err := client.ListGroups(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "test_group_2", result[0].Alias)
	autopilot.Equals(t, "test_group_1", result[1].Alias)
	autopilot.Equals(t, "test_group_3", result[2].Alias)
}

func TestMembers(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query GroupMembersList($after:String!$first:Int!$group:ID!){account{group(id: $group){members(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "members": { "nodes": [ {{ template "user_1"}}, {{ template "user_2"}} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 } }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query GroupMembersList($after:String!$first:Int!$group:ID!){account{group(id: $group){members(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "group": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI" }`,
		`{ "data": { "account": { "group": { "members": { "nodes": [ {{ template "user_3"}} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "group/members", requests...)
	// Act
	group := ol.Group{
		GroupId: ol.GroupId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
		},
	}
	resp, err := group.Members(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "kyle@opslevel.com", result[0].Email)
	autopilot.Equals(t, "Matthew Brahms", result[2].Name)
}

func TestUpdateGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/update")
	members := []ol.MemberInput{
		{Email: "edgar+test@opslevel.com"},
	}
	teams := []ol.IdentifierInput{
		{Alias: "platform"},
	}
	// Act
	result, err := client.UpdateGroup("Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI", ol.GroupInput{
		Description: "This is the first test group",
		Members:     &members,
		Parent:      ol.NewIdentifier("test_group_2"),
		Teams:       &teams,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Name)
	autopilot.Equals(t, "This is the first test group", result.Description)
	autopilot.Equals(t, "test_group_2", result.Parent.Alias)
}
