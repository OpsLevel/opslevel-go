package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
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
	requests := []TestRequest{
		{`{"query": "query GroupChildTeamsList($after:String!$first:Int!$group:ID!){account{group(id: $group){childTeams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},name,responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "first_page_variables" }},"group": "123456789"
			}
			}`,
			`{
				"data": {
					"account": {
						"group": {
							"childTeams": {
								"nodes": [
									{{ template "team_1" }},
									{{ template "team_2" }}
								],
								{{ template "pagination_initial_pageInfo_response" }},
								"totalCount": 2
							}
						  }}}}`},
		{`{"query": "query GroupChildTeamsList($after:String!$first:Int!$group:ID!){account{group(id: $group){childTeams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},name,responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}",
			"variables": {
				{{ template "second_page_variables" }},"group": "123456789"
			}
			}`,
			`{
				"data": {
					"account": {
						"group": {
							"childTeams": {
								"nodes": [
									{{ template "team_3" }}
								],
								{{ template "pagination_second_pageInfo_response" }},
								"totalCount": 1
							}
						  }}}}`},
	}

	client := APaginatedTestClient(t, "group/child_teams", requests...)
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
	client1 := getGroupWithAliasTestClient(t)
	client2 := ATestClient(t, "group/descendant_teams")
	// Act
	group, err := client1.GetGroupWithAlias("test_group_1")
	result, err := group.DescendantTeams(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "platform", result[0].Alias)
}

func TestDescendantRepositories(t *testing.T) {
	// Arrange
	client1 := getGroupWithAliasTestClient(t)
	client2 := ATestClient(t, "group/descendant_repositories")
	// Act
	group, _ := client1.GetGroupWithAlias("test_group_1")
	result, err := group.DescendantRepositories(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "github.com:OpsLevel/cli", result[0].DefaultAlias)
}

func TestDescendantServices(t *testing.T) {
	// Arrange
	client1 := getGroupWithAliasTestClient(t)
	client2 := ATestClient(t, "group/descendant_services")
	// Act
	group, _ := client1.GetGroupWithAlias("test_group_1")
	result, err := group.DescendantServices(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8yNjE5", string(result[0].Id))
	autopilot.Equals(t, "atlantis", result[0].Aliases[0])
}

func TestDescendantSubgroups(t *testing.T) {
	// Arrange
	client1 := getGroupWithAliasTestClient(t)
	client2 := ATestClient(t, "group/descendant_subgroups")
	// Act
	group, _ := client1.GetGroupWithAlias("test_group_1")
	result, err := group.DescendantSubgroups(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "platform", result[0].Alias)
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
	requests := []TestRequest{
		{
			`{
    "query": "query ($after:String!$first:Int!){account{groups(after: $after, first: $first){nodes{alias,id,description,htmlUrl,name,parent{alias,id}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
{{ template "pagination_initial_query_variables" }}
  }`,
			`{
  "data": {
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
      }
    }
  }
}`},
		{
			`{
    "query": "query ($after:String!$first:Int!){account{groups(after: $after, first: $first){nodes{alias,id,description,htmlUrl,name,parent{alias,id}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
{{ template "pagination_second_query_variables" }}
  }`,
			`{
  "data": {
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
      }
    }
  }
}`},
	}
	client := APaginatedTestClient(t, "group/list", requests...)
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
	client1 := getGroupWithAliasTestClient(t)
	client2 := ATestClient(t, "group/members")
	// Act
	group, _ := client1.GetGroupWithAlias("test_group_1")
	result, err := group.Members(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "edgar+test@opslevel.com", result[0].Email)
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
