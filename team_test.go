package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

// TODO: not sure if there is a better way to handle reusing a client
// Probably should be a feature of autopilot
var testRequestWithAlias = NewTestRequest(
	`"query TeamGet($alias:String!){account{team(alias: $alias){alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}"`,
	`{"alias":"example"}`,
	`{ "data": {
    "account": {
      "team": {
        "alias": "example",
        "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
        "aliases": [
          "example"
        ],
        "contacts": [
          {
            "address": "#general",
            "displayName": "",
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTgy",
            "type": "slack"
          },
          {
            "address": "https://example.com",
            "displayName": "Homepage",
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTgz",
            "type": "web"
          }
        ],
        "htmlUrl": "https://app.opslevel.com/teams/example",
        "manager": {
          "email": "john@example.com",
          "htmlUrl": "https://app.opslevel.com/users/1098",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci8xMDk3",
          "name": "John Example",
          "role": "admin"
        },
        "members": {
          "nodes": [
            {
              "email": "kyle@example.com",
              "htmlUrl": "https://app.opslevel.com/users/1097",
              "id": "Z2lkOiBvb38zbGV2ZWwvVXNlci8xMDk3",
              "name": "Kyle Example",
              "role": "admin"
            },
            {
              "email": "john@example.com",
              "htmlUrl": "https://app.opslevel.com/users/1098",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci8xMDk3",
              "name": "John Example",
              "role": "admin"
            },
            {
              "email": "ken@example.com",
              "htmlUrl": "https://app.opslevel.com/users/1099",
              "id": "Z2lkOi7vb3BzbBV2ZWwvVXNlci8xMDk3",
              "name": "Ken Example",
              "role": "user"
            }
          ],
          "pageInfo": {
            "hasNextPage": false,
            "hasPreviousPage": false,
            "startCursor": "MQ",
            "endCursor": "MQ"
          }
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
          "pageInfo": {
            "hasNextPage": false,
            "hasPreviousPage": false,
            "startCursor": "MQ",
            "endCursor": "Mw"
          },
          "totalCount": 3
        },
        "name": "Example",
        "responsibilities": "Foo &amp; bar"
      }}}}`,
)

func TestCreateTeam(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation TeamCreate($input:TeamCreateInput!){teamCreate(input: $input){team{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},errors{message,path}}}"`,
		`{"input": {"name": "Example", "managerEmail": "john@example.com", "parentTeam": {"alias": "parent_team"}, "responsibilities": "Foo & bar", "group": {"alias": "test_group"}, "contacts": [ {"type": "slack_handle", "address": "@mozzie"}, {"type": "slack", "displayName": "", "address": "#general"}, {"type": "web", "displayName": "Homepage", "address": "https://example.com"} ] }}`,
		`{ "data": {
    "teamCreate": {
      "team": {
        "alias": "example",
        "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzY",
        "aliases": [
          "example"
        ],
        "contacts": [
          {
            "address": "#general",
            "displayName": "",
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTk",
            "type": "slack"
          },
          {
            "address": "https://example.com",
            "displayName": "Homepage",
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNjA",
            "type": "web"
          }
        ],
        "group": {
          "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvNTI",
          "alias": "test_group"
        },
        "htmlUrl": "https://app.opslevel-staging.com/teams/example",
        "manager": {
          "email": "john@example.com",
          "htmlUrl": "https://app.opslevel-staging.com/users/410",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci80MTA",
          "name": "John Example",
          "role": "admin"
        },
        "members": {
          "nodes": [
            {
              "email": "john@example.com",
              "htmlUrl": "https://app.opslevel-staging.com/users/410",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci80MTA",
              "name": "John Example",
              "role": "admin"
            }
          ],
          "pageInfo": {
            "hasNextPage": false,
            "hasPreviousPage": false,
            "startCursor": "MQ",
            "endCursor": "MQ"
          }
        },
        "name": "Example",
        "parentTeam": {
          "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3",
          "alias": "parent_team"
        },
        "responsibilities": "Foo &amp; bar"
      },
      "errors": []
    }}}`,
	)

	client := BestTestClient(t, "team/create", testRequest)
	// client := ATestClient(t, "team/create")
	// Act
	contacts := []ol.ContactInput{
		ol.CreateContactSlackHandle("@mozzie", ol.NullString()),
		ol.CreateContactSlack("#general", ol.NewString("")),
		ol.CreateContactWeb("https://example.com", ol.NewString("Homepage")),
	}
	result, err := client.CreateTeam(ol.TeamCreateInput{
		Name:             "Example",
		ManagerEmail:     "john@example.com",
		Responsibilities: "Foo & bar",
		Contacts:         &contacts,
		Group:            ol.NewIdentifier("test_group"),
		ParentTeam:       ol.NewIdentifier("parent_team"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, "test_group", result.Group.Alias)
	autopilot.Equals(t, "parent_team", result.ParentTeam.Alias)
}

func TestGetTeam(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query TeamGet($id:ID!){account{team(id: $id){alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ "id":"Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ" }`,
		`{ "data": {
    "account": {
      "team": {
        "alias": "example",
        "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
        "aliases": [
          "example"
        ],
        "contacts": [
          {
            "address": "#general",
            "displayName": "",
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTgy",
            "type": "slack"
          },
          {
            "address": "https://example.com",
            "displayName": "Homepage",
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTgz",
            "type": "web"
          }
        ],
        "htmlUrl": "https://app.opslevel.com/teams/example",
        "manager": {
          "email": "john@example.com",
          "htmlUrl": "https://app.opslevel.com/users/1098",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci8xMDk3",
          "name": "John Example",
          "role": "admin"
        },
        "members": {
          "nodes": [
            {
              "email": "kyle@example.com",
              "htmlUrl": "https://app.opslevel.com/users/1097",
              "id": "Z2lkOiBvb38zbGV2ZWwvVXNlci8xMDk3",
              "name": "Kyle Example",
              "role": "admin"
            },
            {
              "email": "john@example.com",
              "htmlUrl": "https://app.opslevel.com/users/1098",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci8xMDk3",
              "name": "John Example",
              "role": "admin"
            },
            {
              "email": "ken@example.com",
              "htmlUrl": "https://app.opslevel.com/users/1099",
              "id": "Z2lkOi7vb3BzbBV2ZWwvVXNlci8xMDk3",
              "name": "Ken Example",
              "role": "user"
            }
          ],
          "pageInfo": {
            "hasNextPage": false,
            "hasPreviousPage": false,
            "startCursor": "MQ",
            "endCursor": "MQ"
          }
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
          "pageInfo": {
            "hasNextPage": false,
            "hasPreviousPage": false,
            "startCursor": "MQ",
            "endCursor": "Mw"
          },
          "totalCount": 3
        },
        "name": "Example",
        "responsibilities": "Foo &amp; bar"
      }}}}`,
	)

	client := BestTestClient(t, "team/get", testRequest)
	// Act
	result, err := client.GetTeam("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
}

func TestTeamMembers(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query TeamMembersList($after:String!$first:Int!$team:ID!){account{team(id: $team){members(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "team": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "team": { "members": { "nodes": [ {{ template "user_1"}}, {{ template "user_2"}} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query TeamMembersList($after:String!$first:Int!$team:ID!){account{team(id: $team){members(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "team": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "team": { "members": { "nodes": [ {{ template "user_3"}} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "team/members", requests...)
	// Act
	team := ol.Team{
		TeamId: ol.TeamId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
		},
	}
	resp, err := team.GetMembers(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "kyle@opslevel.com", result[0].Email)
	autopilot.Equals(t, "Matthew Brahms", result[2].Name)
}

func TestTeamTags(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query TeamTagsList($after:String!$first:Int!$team:ID!){account{team(id: $team){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "first_page_variables" }}, "team": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "team": { "tags": { "nodes": [ { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA5", "key": "prod", "value": "false" } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query TeamTagsList($after:String!$first:Int!$team:ID!){account{team(id: $team){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}"`,
		`{ {{ template "second_page_variables" }}, "team": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4" }`,
		`{ "data": { "account": { "team": { "tags": { "nodes": [ { "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODA4", "key": "test", "value": "true" } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 } }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "team/tags", requests...)
	// Act
	team := ol.Team{
		TeamId: ol.TeamId{
			Id: "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS85NjQ4",
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
	client := BestTestClient(t, "team/get_with_alias", testRequestWithAlias)
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
	testRequestOne := NewTestRequest(
		`"query TeamList($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}"`,
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NDk",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTA",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "Marketing",
              "responsibilities": null
            }
          ],
        {{ template "pagination_initial_pageInfo_response" }},
        "totalCount": 3
        }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query TeamList($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}"`,
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTE",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTI",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "VPs",
              "responsibilities": null
            }
          ],
{{ template "pagination_second_pageInfo_response" }},
			"totalCount": 2
        }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

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
	testRequestOne := NewTestRequest(
		`"query TeamList($after:String!$email:String!$first:Int!){account{teams(managerEmail: $email, after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}"`,
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NDk",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTA",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "Marketing",
              "responsibilities": null
            }
          ],
        {{ template "pagination_initial_pageInfo_response" }},
        "totalCount": 3
        }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query TeamList($after:String!$email:String!$first:Int!){account{teams(managerEmail: $email, after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}"`,
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTE",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
					{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
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
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTI",
              "manager": {{ template "user_1" }},
              "members": {
                "nodes": [
                  	{{ template "user_4" }},
					{{ template "user_5" }},
                    {{ template "user_6" }}
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "VPs",
              "responsibilities": null
            }
          ],
        {{ template "pagination_second_pageInfo_response" }},
        "totalCount": 2
        }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

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
	testRequest := NewTestRequest(
		`"mutation TeamUpdate($input:TeamUpdateInput!){teamUpdate(input: $input){team{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},errors{message,path}}}"`,
		`{"input": {"id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ", "managerEmail": "ken@example.com", "responsibilities": "Foo & bar", "parentTeam": {"alias": "parent_team"}, "group": {"alias": "test_group" }}}`,
		`{ "data": {
      "teamUpdate": {
        "team": {
          "alias": "example",
          "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
          "aliases": [
            "example"
          ],
          "contacts": [
            {
              "address": "#general",
              "displayName": "",
              "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTgy",
              "type": "slack"
            },
            {
              "address": "https://example.com",
              "displayName": "Homepage",
              "id": "Z2lkOi8vb3BzbGV2ZWwvQ29udGFjdC8xNTgz",
              "type": "web"
            }
          ],
          "group": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvNTI",
            "alias": "test_group"
          },
          "htmlUrl": "https://app.opslevel.com/teams/example",
          "manager":               {
            "email": "ken@example.com",
            "htmlUrl": "https://app.opslevel.com/users/1099",
            "id": "Z2lkOi7vb3BzbBV2ZWwvVXNlci8xMDk3",
            "name": "Ken Example",
            "role": "user"
          },
          "members": {
            "nodes": [
              {
                "email": "kyle@example.com",
                "htmlUrl": "https://app.opslevel.com/users/1097",
                "id": "Z2lkOiBvb38zbGV2ZWwvVXNlci8xMDk3",
                "name": "Kyle Example",
                "role": "admin"
              },
              {
                "email": "john@example.com",
                "htmlUrl": "https://app.opslevel.com/users/1098",
                "id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci8xMDk3",
                "name": "John Example",
                "role": "admin"
              },
              {
                "email": "ken@example.com",
                "htmlUrl": "https://app.opslevel.com/users/1099",
                "id": "Z2lkOi7vb3BzbBV2ZWwvVXNlci8xMDk3",
                "name": "Ken Example",
                "role": "user"
              }
            ],
            "pageInfo": {
              "hasNextPage": false,
              "hasPreviousPage": false,
              "startCursor": "MQ",
              "endCursor": "MQ"
            }
          },
          "parentTeam": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3",
            "alias": "parent_team"
          },
          "name": "Example",
          "responsibilities": "Foo &amp; bar"
        },
        "errors": []
       }}}`,
	)
	client := BestTestClient(t, "team/update", testRequest)
	// client := ATestClient(t, "team/update")
	// Act
	result, err := client.UpdateTeam(ol.TeamUpdateInput{
		Id:               "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
		ManagerEmail:     "ken@example.com",
		Responsibilities: "Foo & bar",
		Group:            ol.NewIdentifier("test_group"),
		ParentTeam:       ol.NewIdentifier("parent_team"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "ken@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, "test_group", result.Group.Alias)
	autopilot.Equals(t, "parent_team", result.ParentTeam.Alias)
}

func TestDeleteTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/delete")
	// Act
	err := client.DeleteTeam("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteTeamWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "team/delete", "team/delete_with_alias")
	// Act
	err := client.DeleteTeamWithAlias("example")
	// Assert
	autopilot.Ok(t, err)
}

func TestTeamAddMember(t *testing.T) {
	// Arrange
	testRequestWithTeamId := NewTestRequest(
		`"mutation TeamMembershipCreate($input:TeamMembershipCreateInput!){teamMembershipCreate(input: $input){memberships{team{alias,id},role,user{id,email}},errors{message,path}}}"`,
		`{"input": {"teamId": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ", "members": [ {"user": {"email": "john@example.com"}} ] }}`,
		`{"data": {"teamMembershipCreate": {"memberships": [ {"user": {"id": "Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4", "email": "john@example.com"}, "role": "admin"} ], "errors": [] }}}`,
	)

	clientWithTeamId := BestTestClient(t, "team/add_member", testRequestWithTeamId)
	clientWithAlias := BestTestClient(t, "team/get_with_alias_add_member", testRequestWithAlias)

	// Act
	team, _ := clientWithAlias.GetTeamWithAlias("example")
	result, err := clientWithTeamId.AddMember(&team.TeamId, "john@example.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamRemoveMember(t *testing.T) {
	// Arrange
	client1 := BestTestClient(t, "team/get_with_alias_rm_member", testRequestWithAlias)
	client2 := ATestClient(t, "team/remove_member")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.RemoveMember(&team.TeamId, "john@example.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamAddContact(t *testing.T) {
	// Arrange
	client1 := BestTestClient(t, "team/get_with_alias_add_contact", testRequestWithAlias)
	client2 := ATestClient(t, "team/add_contact")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.AddContact(string(team.TeamId.Id), ol.CreateContactSlack("#general", ol.NewString("")))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "#general", result.Address)
}

func TestTeamUpdateContact(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/update_contact")
	// Act
	result, err := client.UpdateContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4", ol.CreateContactSlack("#general", ol.NewString("Main Channel")))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Main Channel", result.DisplayName)
}

func TestTeamUpdateContactWithTypeNil(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "team/update_contact", "team/update_contact_nil_type")
	// Act
	result, err := client.UpdateContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4", ol.ContactInput{Address: "#general", DisplayName: ol.NewString("Main Channel")})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Main Channel", result.DisplayName)
}

func TestTeamRemoveContact(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/remove_contact")
	// Act
	err := client.RemoveContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4")
	// Assert
	autopilot.Ok(t, err)
}
