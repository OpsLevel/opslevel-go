package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

// TODO: not sure if there is a better way to handle reusing a client
// Probably should be a feature of autopilot
func getTestRequestWithAlias() autopilot.TestRequest {
	return autopilot.NewTestRequest(
		`query TeamGet($alias:String!){account{team(alias: $alias){alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{"alias":"example"}`,
		`{ "data": {
    "account": {
      "team": {
        "alias": "example",
        {{ template "id1" }},
        "aliases": [
          "example"
        ],
        "contacts": [
          {
            "address": "#general",
            "displayName": "",
            {{ template "id2" }},
            "type": "slack"
          },
          {
            "address": "https://example.com",
            "displayName": "Homepage",
            {{ template "id3" }},
            "type": "web"
          }
        ],
        "htmlUrl": "https://app.opslevel.com/teams/example",
        "manager": {
          "email": "john@example.com",
          "htmlUrl": "https://app.opslevel.com/users/1098",
          {{ template "id2" }},
          "name": "John Example",
          "role": "admin"
        },
        "memberships": {
          "nodes": [
            {{ template "team_membership_1" }},
            {{ template "team_membership_2" }},
            {{ template "team_membership_3" }}
          ],
          {{ template "pagination_response_same_cursor" }}
        },
        "tags": {
          "nodes": [
            {
              {{ template "id2" }},
              "key": "k8s-app",
              "value": "kube-dns"
            },
            {
              {{ template "id3" }},
              "key": "imported",
              "value": "kubectl-opslevel"
            },
            {
              {{ template "id4" }},
              "key": "hello",
              "value": "world"
            }
          ],
          {{ template "pagination_response_different_cursor" }},
          "totalCount": 3
        },
        "name": "Example",
        "responsibilities": "Foo &amp; bar"
      }}}}`,
	)
}

func TestCreateTeam(t *testing.T) {
	// Arrange
	contacts := autopilot.Register[[]ol.ContactInput]("contact_input_slice",
		[]ol.ContactInput{
			ol.CreateContactSlackHandle("@mozzie", ol.NullString()),
			ol.CreateContactWeb("https://example.com", ol.RefOf("Homepage")),
		},
	)
	input := autopilot.Register[ol.TeamCreateInput]("team_create_input",
		ol.TeamCreateInput{
			Name:             "Example",
			Responsibilities: ol.RefOf("Foo & bar"),
			Contacts:         &contacts,
			ParentTeam:       ol.NewIdentifier("parent_team"),
		})
	testRequest := autopilot.NewTestRequest(
		`mutation TeamCreate($input:TeamCreateInput!){teamCreate(input: $input){team{alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
		`{"input": {{ template "team_create_input" }} }`,
		`{ "data": {
    "teamCreate": {
      "team": {
        "alias": "example",
        {{ template "id1" }},
        "aliases": [
          "example"
        ],
        "contacts": {{ template "contact_input_slice" }},
        "htmlUrl": "https://app.opslevel-staging.com/teams/example",
        "manager": {
          "email": "john@example.com",
          "htmlUrl": "https://app.opslevel-staging.com/users/410",
          {{ template "id2" }},
          "name": "John Example",
          "role": "admin"
        },
        "memberships": {
          "nodes": [ {{ template "team_membership_3" }} ],
          {{ template "pagination_response_same_cursor" }}
        },
        "name": "Example",
        "parentTeam": {
          {{ template "id2" }},
          "alias": "parent_team"
        },
        "responsibilities": "Foo &amp; bar"
      },
      "errors": []
    }}}`,
	)

	client := BestTestClient(t, "team/create", testRequest)
	// Act
	result, err := client.CreateTeam(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, "parent_team", result.ParentTeam.Alias)
}

func TestGetTeam(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query TeamGet($id:ID!){account{team(id: $id){alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "id1" }} }`,
		`{ "data": {
    "account": {
      "team": {
        "alias": "example",
        {{ template "id1" }},
        "aliases": [
          "example"
        ],
        "contacts": [
          {
            "address": "#general",
            "displayName": "",
            {{ template "id2" }},
            "type": "slack"
          },
          {
            "address": "https://example.com",
            "displayName": "Homepage",
            {{ template "id3" }},
            "type": "web"
          }
        ],
        "htmlUrl": "https://app.opslevel.com/teams/example",
        "manager": {
          "email": "john@example.com",
          "htmlUrl": "https://app.opslevel.com/users/1098",
          {{ template "id2" }},
          "name": "John Example",
          "role": "admin"
        },
        "memberships": {
          "nodes": [ {{ template "team_membership_2" }} ],
          {{ template "pagination_response_same_cursor" }}
        },
        "tags": {
          "nodes": [
            {
              {{ template "id2" }},
              "key": "k8s-app",
              "value": "kube-dns"
            },
            {
              {{ template "id3" }},
              "key": "imported",
              "value": "kubectl-opslevel"
            },
            {
              {{ template "id4" }},
              "key": "hello",
              "value": "world"
            }
          ],
          {{ template "pagination_response_different_cursor" }},
          "totalCount": 3
        },
        "name": "Example",
        "responsibilities": "Foo &amp; bar"
      }}}}`,
	)

	client := BestTestClient(t, "team/get", testRequest)
	// Act
	result, err := client.GetTeam(id1)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, id2, result.Memberships.Nodes[0].User.Id)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
}

func TestTeamMembers(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query TeamMembersList($after:String!$first:Int!$team:ID!){account{team(id: $team){memberships(after: $after, first: $first){nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "team": "{{ template "id4_string" }}" }`,
		`{ "data": { "account": { "team": { "memberships": { "nodes": [ {{ template "team_membership_1" }}, {{ template "team_membership_2"}} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query TeamMembersList($after:String!$first:Int!$team:ID!){account{team(id: $team){memberships(after: $after, first: $first){nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "team": "{{ template "id4_string" }}" }`,
		`{ "data": { "account": { "team": { "memberships": { "nodes": [ {{ template "team_membership_3"}} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "team/members", requests...)
	// Act
	team := ol.Team{
		TeamId: ol.TeamId{
			Id: id4,
		},
	}
	resp, err := team.GetMemberships(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "kyle@opslevel.com", result[0].User.Email)
}

func TestTeamTags(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query TeamTagsList($after:String!$first:Int!$team:ID!){account{team(id: $team){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "team": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "team": { "tags": { "nodes": [ { {{ template "id2" }}, "key": "prod", "value": "false" } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query TeamTagsList($after:String!$first:Int!$team:ID!){account{team(id: $team){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "team": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "team": { "tags": { "nodes": [ { {{ template "id3" }}, "key": "test", "value": "true" } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "team/tags", requests...)
	// Act
	team := ol.Team{
		TeamId: ol.TeamId{
			Id: id1,
		},
	}
	resp, err := team.GetTags(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, resp.TotalCount)
	autopilot.Equals(t, "prod", result[0].Key)
	autopilot.Equals(t, "false", result[0].Value)
	autopilot.Equals(t, "test", result[1].Key)
	autopilot.Equals(t, "true", result[1].Value)
}

func TestGetTeamWithAlias(t *testing.T) {
	// Arrange
	client := BestTestClient(t, "team/get_with_alias", getTestRequestWithAlias())
	// Act
	result, err := client.GetTeamWithAlias("example")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
}

func TestListTeams(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query TeamList($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": {
      "account": {
        "teams": {
          "nodes": [
            {
              "alias": "devops",
              "aliases": [
                "devops"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/devops",
              {{ template "id1" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "DevOps",
              "responsibilities": "Own Infra & Tools."
            },
            {
              "alias": "developers",
              "aliases": [
                "developers"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/developers",
              {{ template "id2" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "Developers",
              "responsibilities": null
            },
            {
              "alias": "marketing",
              "aliases": [
                "marketing"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/marketing",
              {{ template "id3" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "Marketing",
              "responsibilities": null
            }
          ],
        {{ template "pagination_initial_pageInfo_response" }},
        "totalCount": 3
        }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query TeamList($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": {
      "account": {
        "teams": {
          "nodes": [
            {
              "alias": "security",
              "aliases": [
                "security"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/security",
              {{ template "id4" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "Security",
              "responsibilities": null
            },
            {
              "alias": "vps",
              "aliases": [
                "vps"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/vps",
              {{ template "id4" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "VPs",
              "responsibilities": null
            }
          ],
{{ template "pagination_second_pageInfo_response" }},
			"totalCount": 2
        }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "team/list", requests...)
	// Act
	response, err := client.ListTeams(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, response.TotalCount)
	autopilot.Equals(t, "devops", result[0].Alias)
	autopilot.Equals(t, "developers", result[1].Alias)
	autopilot.Equals(t, "security", result[3].Alias)
	autopilot.Equals(t, "vps", result[4].Alias)
	autopilot.Equals(t, "Own Infra & Tools.", result[0].Responsibilities)
}

func TestListTeamsWithManager(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query TeamList($after:String!$email:String!$first:Int!){account{teams(managerEmail: $email, after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ "after": "", "first": 100, "email": "kyle@opslevel.com" }`,
		`{ "data": {
      "account": {
        "teams": {
          "nodes": [
            {
              "alias": "devops",
              "aliases": [
                "devops"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/devops",
              {{ template "id1" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "DevOps",
              "responsibilities": "Own Infra & Tools."
            },
            {
              "alias": "developers",
              "aliases": [
                "developers"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/developers",
              {{ template "id1" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "Developers",
              "responsibilities": null
            },
            {
              "alias": "marketing",
              "aliases": [
                "marketing"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/marketing",
              {{ template "id2" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "Marketing",
              "responsibilities": null
            }
          ],
        {{ template "pagination_initial_pageInfo_response" }},
        "totalCount": 3
        }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query TeamList($after:String!$email:String!$first:Int!){account{teams(managerEmail: $email, after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ "after": "OA", "first": 100, "email": "kyle@opslevel.com" }`,
		`{ "data": {
      "account": {
        "teams": {
          "nodes": [
            {
              "alias": "security",
              "aliases": [
                "security"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/security",
              {{ template "id3" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "Security",
              "responsibilities": null
            },
            {
              "alias": "vps",
              "aliases": [
                "vps"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/vps",
              {{ template "id4" }},
              "manager": {{ template "user_1" }},
              "memberships": {
                "nodes": [
                    {{ template "team_membership_1" }},
                    {{ template "team_membership_2" }},
                    {{ template "team_membership_3" }}
                ],
                {{ template "pagination_response_different_cursor" }}
              },
              "name": "VPs",
              "responsibilities": null
            }
          ],
        {{ template "pagination_second_pageInfo_response" }},
        "totalCount": 2
        }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "team/list_with_manager", requests...)
	// Act
	response, err := client.ListTeamsWithManager("kyle@opslevel.com", nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, response.TotalCount)
	autopilot.Equals(t, "devops", result[0].Alias)
	autopilot.Equals(t, "developers", result[1].Alias)
	autopilot.Equals(t, "security", result[3].Alias)
	autopilot.Equals(t, "vps", result[4].Alias)
	autopilot.Equals(t, "Own Infra & Tools.", result[0].Responsibilities)
}

func TestUpdateTeam(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.TeamUpdateInput]("team_update_input",
		ol.TeamUpdateInput{
			Id:               &id1,
			Responsibilities: ol.RefOf("Foo & bar"),
			ParentTeam:       ol.NewIdentifier("parent_team"),
		},
	)
	testRequest := autopilot.NewTestRequest(
		`mutation TeamUpdate($input:TeamUpdateInput!){teamUpdate(input: $input){team{alias,id,aliases,contacts{address,displayName,id,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{team{alias,id},role,user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},errors{message,path}}}`,
		`{"input": {{ template "team_update_input" }} }`,
		`{ "data": {
      "teamUpdate": {
        "team": {
          "alias": "example",
          {{ template "id1" }},
          "aliases": [
            "example"
          ],
          "contacts": [
            {
              "address": "#general",
              "displayName": "",
              {{ template "id2" }},
              "type": "slack"
            },
            {
              "address": "https://example.com",
              "displayName": "Homepage",
              {{ template "id4" }},
              "type": "web"
            }
          ],
          "htmlUrl": "https://app.opslevel.com/teams/example",
          "manager":               {
            "email": "ken@example.com",
            "htmlUrl": "https://app.opslevel.com/users/1099",
            {{ template "id1" }},
            "name": "Ken Example",
            "role": "user"
          },
          "memberships": {
            "nodes": [
              {{ template "team_membership_1" }},
              {{ template "team_membership_2" }},
              {{ template "team_membership_3" }}
            ],
            {{ template "pagination_response_same_cursor" }}
          },
          "parentTeam": {
            {{ template "id4" }},
            "alias": "parent_team"
          },
          "name": "Example",
          "responsibilities": "Foo &amp; bar"
        },
        "errors": []
       }}}`,
	)
	client := BestTestClient(t, "team/update", testRequest)
	// Act
	result, err := client.UpdateTeam(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "ken@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, "parent_team", result.ParentTeam.Alias)
}

func TestDeleteTeam(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TeamDelete($input:TeamDeleteInput!){teamDelete(input: $input){deletedTeamId,deletedTeamAlias,errors{message,path}}}`,
		`{"input": { {{ template "id3" }} } }`,
		`{"data": {"teamDelete": {"deletedTeamId": "{{ template "id3_string" }}", "deletedTeamAlias": "example", "errors": [] } }}`,
	)
	client := BestTestClient(t, "team/delete", testRequest)
	// Act
	err := client.DeleteTeam(id3)
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteTeamWithAlias(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TeamDelete($input:TeamDeleteInput!){teamDelete(input: $input){deletedTeamId,deletedTeamAlias,errors{message,path}}}`,
		`{"input": {"alias": "example" }}`,
		`{"data": {"teamDelete": {"deletedTeamId": "{{ template "id3_string" }}", "deletedTeamAlias": "example", "errors": [] }}}`,
	)
	client := BestTestClient(t, "team/delete_with_alias", testRequest)
	// Act
	err := client.DeleteTeamWithAlias("example")
	// Assert
	autopilot.Ok(t, err)
}

func TestTeamAddMemberhip(t *testing.T) {
	// Arrange
	testRequestWithTeamId := autopilot.NewTestRequest(
		`mutation TeamMembershipCreate($input:TeamMembershipCreateInput!){teamMembershipCreate(input: $input){memberships{team{alias,id},role,user{id,email}},errors{message,path}}}`,
		`{"input": {"teamId": "{{ template "id1_string" }}", "members": [ { {{ template "team_membership_user_input_1" }} } ] }}`,
		`{"data": {"teamMembershipCreate": {"memberships": [ {{ template "team_membership_1" }} ], "errors": [] }}}`,
	)

	clientWithTeamId := BestTestClient(t, "team/add_member", testRequestWithTeamId)
	clientWithAlias := BestTestClient(t, "team/get_with_alias_add_member", getTestRequestWithAlias())

	// Act
	team, _ := clientWithAlias.GetTeamWithAlias("example")
	newMembership := ol.TeamMembershipUserInput{
		Role: ol.RefOf("user"),
		User: &ol.UserIdentifierInput{Id: &id1, Email: ol.RefOf("kyle@opslevel.com")},
	}
	result, err := clientWithTeamId.AddMemberships(&team.TeamId, newMembership)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamRemoveMemberhip(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation TeamMembershipDelete($input:TeamMembershipDeleteInput!){teamMembershipDelete(input: $input){deletedMembers{id,email,htmlUrl,name,role},errors{message,path}}}`,
		`{"input": { "teamId": "{{ template "id1_string" }}", "members": [ { {{ template "team_membership_user_input_1" }} } ] }}`,
		`{ "data": {
    "teamMembershipDelete": {
      "deletedMembers": [
        {
          {{ template "user_id_email_1" }},
          "htmlUrl": "https://app.opslevel.com/users/3068",
          "name": "John",
          "role": "user"
        }
      ],
      "errors": []
    }}}`,
	)
	client1 := BestTestClient(t, "team/get_with_alias_rm_member", getTestRequestWithAlias())
	client2 := BestTestClient(t, "team/remove_member", testRequest)
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	membershipToDelete := ol.TeamMembershipUserInput{
		Role: ol.RefOf("user"),
		User: &ol.UserIdentifierInput{Id: &id1, Email: ol.RefOf("kyle@opslevel.com")},
	}

	result, err := client2.RemoveMemberships(&team.TeamId, membershipToDelete)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamAddContact(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ContactCreate($input:ContactCreateInput!){contactCreate(input: $input){contact{address,displayName,id,type},errors{message,path}}}`,
		`{"input": {"type":"slack", "address":"#general", "ownerId": "{{ template "id1_string" }}" }}`,
		`{"data": {"contactCreate": {"contact": {"address": "#general", "displayName": "Slack", {{ template "id2" }}, "type": "slack"}, "errors": [] } }}`,
	)
	client1 := BestTestClient(t, "team/get_with_alias_add_contact", getTestRequestWithAlias())
	client2 := BestTestClient(t, "team/alias_add_contact", testRequest)
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.AddContact(string(team.TeamId.Id), ol.CreateContactSlack("#general", nil))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "#general", result.Address)
}

func TestTeamUpdateContact(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.ContactInput]("contact_input_slack",
		ol.CreateContactSlack("#general", ol.RefOf("Main Channel")),
	)
	autopilot.Register[ol.ContactUpdateInput]("contact_update_input_slack",
		ol.ContactUpdateInput{
			Id:          id1,
			DisplayName: ol.RefOf(*input.DisplayName),
			Address:     ol.RefOf(input.Address),
			Type:        &input.Type,
		})
	testRequest := autopilot.NewTestRequest(
		`mutation ContactUpdate($input:ContactUpdateInput!){contactUpdate(input: $input){contact{address,displayName,id,type},errors{message,path}}}`,
		`{"input":  {{ template "contact_update_input_slack" }} }`,
		`{"data": {"contactUpdate":  {"contact": {"address": "#general", "displayName": "Main Channel", {{ template "id2" }}, "type": "slack" }, "errors": [] }}}`,
	)
	client := BestTestClient(t, "team/update_contact", testRequest)
	// Act
	result, err := client.UpdateContact(id1, input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Main Channel", result.DisplayName)
}

func TestTeamUpdateContactWithTypeNil(t *testing.T) {
	// Arrange
	input := autopilot.Register[ol.ContactInput]("contact_input",
		ol.ContactInput{Address: "#general", DisplayName: ol.RefOf("Main Channel")},
	)
	autopilot.Register[ol.ContactUpdateInput]("contact_update_input",
		ol.ContactUpdateInput{
			Id:          id2,
			DisplayName: ol.RefOf(*input.DisplayName),
			Address:     ol.RefOf(input.Address),
		})
	testRequest := autopilot.NewTestRequest(
		`mutation ContactUpdate($input:ContactUpdateInput!){contactUpdate(input: $input){contact{address,displayName,id,type},errors{message,path}}}`,
		`{ "input": {{ template "contact_update_input" }} }`,
		`{"data": {"contactUpdate": {"contact": {"address": "#general", "displayName": "Main Channel", {{ template "id2" }}, "type": "slack" }, "errors": [] }}}`,
	)
	client := BestTestClient(t, "team/update_contact_nil_type", testRequest)
	// Act
	result, err := client.UpdateContact(id2, input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Main Channel", result.DisplayName)
}

func TestTeamRemoveContact(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ContactDelete($input:ContactDeleteInput!){contactDelete(input: $input){deletedContactId,errors{message,path}}}`,
		`{"input": { {{ template "id3" }} }}`,
		`{"data": {"contactDelete": {"deletedContactId": "{{ template "id2_string" }}", "errors": [] } }}`,
	)
	client := BestTestClient(t, "team/remove_contact", testRequest)
	// Act
	err := client.RemoveContact(id3)
	// Assert
	autopilot.Ok(t, err)
}
