package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestInviteUser(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation UserInvite($email:String!$input:UserInput!){userInvite(email: $email input: $input){user{id,email,htmlUrl,name,role},errors{message,path}}}`,
		`{"email": "kyle@opslevel.com", "input": { "name": "Kyle Rockman", "skipWelcomeEmail": false }}`,
		`{"data": { "userInvite": { "user": {{ template "user_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "user/invite", testRequest)
	// Act
	result, err := client.InviteUser("kyle@opslevel.com", ol.UserInput{
		Name:             ol.NewString("Kyle Rockman"),
		SkipWelcomeEmail: ol.Bool(false),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestGetUser(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query UserGet($input:UserIdentifierInput!){account{user(input: $input){id,email,htmlUrl,name,role}}}`,
		`{"input": { "email": "kyle@opslevel.com" }}`,
		`{"data": {"account": {"user": {{ template "user_1" }} }}}`,
	)

	client := BestTestClient(t, "user/get", testRequest)
	// Act
	result, err := client.GetUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestGetUserTeams(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query UserTeamsList($after:String!$first:Int!$user:ID!){account{user(id: $user){teams(after: $after, first: $first){nodes{alias,id},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "user": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "user": { "teams": { "nodes": [ { {{ template "teamId_1" }} }, { {{ template "teamId_2" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query UserTeamsList($after:String!$first:Int!$user:ID!){account{user(id: $user){teams(after: $after, first: $first){nodes{alias,id},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "user": "{{ template "id1_string" }}" }`,
		`{ "data": { "account": { "user": { "teams": { "nodes": [ { {{ template "teamId_3"}} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "user/teams", requests...)
	// Act
	user := ol.User{
		UserId: ol.UserId{
			Id: id1,
		},
	}
	resp, err := user.Teams(client, nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, resp.TotalCount)
	autopilot.Equals(t, "example", result[0].Alias)
	autopilot.Equals(t, "example_3", result[2].Alias)
}

func TestListUser(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query UserList($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount}}}`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "users": { "nodes": [ {{ template "user_1" }}, {{ template "user_2" }} ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query UserList($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},{{ template "pagination_request" }},totalCount}}}`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "users": { "nodes": [ {{ template "user_3" }} ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "user/list", requests...)
	// Act
	response, err := client.ListUsers(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Edgar Ochoa", result[1].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[1].Role)
	autopilot.Equals(t, "Matthew Brahms", result[2].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[2].Role)
}

func TestUpdateUser(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation UserUpdate($input:UserInput!$user:UserIdentifierInput!){userUpdate(user: $user input: $input){user{id,email,htmlUrl,name,role},errors{message,path}}}`,
		`{"input": {"role": "admin", "skipWelcomeEmail": false }, "user": {"email": "kyle@opslevel.com" }}`,
		`{"data": {"userUpdate": {"user": {{ template "user_1_update" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "user/update", testRequest)
	// Act
	result, err := client.UpdateUser("kyle@opslevel.com", ol.UserInput{
		Role:             ol.NewString(string(ol.UserRoleAdmin)),
		SkipWelcomeEmail: ol.Bool(false),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result.Role)
}

func TestDeleteUser(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation UserDelete($user:UserIdentifierInput!){userDelete(user: $user){errors{message,path}}}`,
		`{"user": {"email": "kyle@opslevel.com" }}`,
		`{"data": {"userDelete": {"errors": [] }}}`,
	)

	client := BestTestClient(t, "user/delete", testRequest)
	// Act
	err := client.DeleteUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteUserDoesNotExist(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation UserDelete($user:UserIdentifierInput!){userDelete(user: $user){errors{message,path}}}`,
		`{"user": {"email": "not-found@opslevel.com" }}`,
		`{"data": {"userDelete": {"errors": [{"message": "User with email 'not-found@opslevel.com' does not exist on this account", "path": ["user"] }] }}}`,
	)

	client := BestTestClient(t, "user/delete_not_found", testRequest)
	// Act
	err := client.DeleteUser("not-found@opslevel.com")
	// Assert
	autopilot.Equals(t, `OpsLevel API Errors:
	- 'user' User with email 'not-found@opslevel.com' does not exist on this account
`, err.Error())
}

func TestGetUserTags(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query UserTagsList($after:String!$first:Int!$user:ID!){account{user(id: $user){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "first_page_variables" }}, "user": "{{ template "id1_string" }}" }`,
		`{
        "data": {
          "account": {
            "user": {
              "tags": {
                "nodes": [
                  {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIw",
                    "key": "user-tag-1",
                    "value": "user-value-1"
                  },
                  {
                    "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIx",
                    "key": "user-tag-2",
                    "value": "user-value-2"
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
		`query UserTagsList($after:String!$first:Int!$user:ID!){account{user(id: $user){tags(after: $after, first: $first){nodes{id,key,value},{{ template "pagination_request" }},totalCount}}}}`,
		`{ {{ template "second_page_variables" }}, "user": "{{ template "id1_string"}}" }`,
		`{
        "data": {
          "account": {
            "user": {
              "tags": {
                "nodes": [
                {
                  "id": "Z2lkOi8vb3BzbGV2ZWwvVGFnLzEwODIz",
                  "key": "captain",
                  "value": "tuna"
                }
                ],
                {{ template "pagination_second_pageInfo_response" }},
                "totalCount": 1
              }
          }
          }}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "user/tags", requests...)
	// Act
	user := ol.User{
		UserId: ol.UserId{
			Id: id1,
		},
	}
	resp, err := user.GetTags(client, nil)
	autopilot.Ok(t, err)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, resp.TotalCount)
	autopilot.Equals(t, "user-tag-1", result[0].Key)
	autopilot.Equals(t, "user-value-1", result[0].Value)
	autopilot.Equals(t, "user-tag-2", result[1].Key)
	autopilot.Equals(t, "user-value-2", result[1].Value)
	autopilot.Equals(t, "db", result[2].Key)
	autopilot.Equals(t, "prod", result[2].Value)
	autopilot.Equals(t, "captain", result[3].Key)
	autopilot.Equals(t, "tuna", result[3].Value)
}
