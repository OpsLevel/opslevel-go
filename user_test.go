package opslevel_test

import (
	"errors"
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestInviteUser(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserInvite($email:String!$input:UserInput!){userInvite(email: $email input: $input){user{id,email,htmlUrl,name,role},errors{message,path}}}",
	"variables":{
		"email": "kyle@opslevel.com",
		"input": {
			"name": "Kyle Rockman"
		}
	}}`
	response := `{"data": {
	"userInvite": {
		"user": {{ template "user_1" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "user/invite", request, response)
	// Act
	result, err := client.InviteUser("kyle@opslevel.com", ol.UserInput{
		Name: "Kyle Rockman",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1", string(result.Id))
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestGetUser(t *testing.T) {
	// Arrange
	request := `{"query":"query UserGet($input:UserIdentifierInput!){account{user(input: $input){id,email,htmlUrl,name,role}}}",
	"variables":{
		"input": {
			"email": "kyle@opslevel.com"
		}
	}}`
	response := `{"data": {
	"account": {
		"user": {{ template "user_1" }}
	}
	}}`
	client := ABetterTestClient(t, "user/get", request, response)
	// Act
	result, err := client.GetUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1", string(result.Id))
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestListUser(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{
  "query": "query UserList($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
  {{ template "pagination_initial_query_variables" }}
}`,
			`{
  "data": {
    "account": {
      "users": {
        "nodes": [
          {{ template "user_1" }},
          {{ template "user_2" }}
        ],
        {{ template "pagination_initial_pageInfo_response" }},
		"totalCount": 2
      }
    }
  }
}`},
		{`{
  "query": "query UserList($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
  {{ template "pagination_second_query_variables" }}
}`,
			`{
  "data": {
    "account": {
      "users": {
        "nodes": [
			{{ template "user_3" }}
        ],
        {{ template "pagination_second_pageInfo_response" }},
		"totalCount": 1
      }
    }
  }
}`},
	}
	client := APaginatedTestClient(t, "user/list", requests...)
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
	request := `{"query":"mutation UserUpdate($input:UserInput!$user:UserIdentifierInput!){userUpdate(user: $user input: $input){user{id,email,htmlUrl,name,role},errors{message,path}}}",
	"variables":{
		"input": {
			"role": "admin"
		},
		"user": {
			"email": "kyle@opslevel.com"
		}
	}}`
	response := `{"data": {
	"userUpdate": {
		"user": {{ template "user_1_update" }},
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "user/update", request, response)
	// Act
	result, err := client.UpdateUser("kyle@opslevel.com", ol.UserInput{
		Role: ol.UserRoleAdmin,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1", string(result.Id))
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result.Role)
}

func TestDeleteUser(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserDelete($user:UserIdentifierInput!){userDelete(user: $user){errors{message,path}}}",
	"variables":{
		"user": {
			"email": "kyle@opslevel.com"
		}
	}}`
	response := `{"data": {
	"userDelete": {
		"errors": []
	}
	}}`
	client := ABetterTestClient(t, "user/delete", request, response)
	// Act
	err := client.DeleteUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteUserDoesNotExist(t *testing.T) {
	// Arrange
	request := `{"query":"mutation UserDelete($user:UserIdentifierInput!){userDelete(user: $user){errors{message,path}}}",
	"variables":{
		"user": {
			"email": "not-found@opslevel.com"
		}
	}}`
	response := `{"data": {
	"userDelete": {
		"errors": [{
			"message": "User with email 'not-found@opslevel.com' does not exist on this account",
			"path": ["user"]
		}]
	}
	}}`
	client := ABetterTestClient(t, "user/delete_not_found", request, response)
	// Act
	err := client.DeleteUser("not-found@opslevel.com")
	// Assert
	autopilot.Equals(t, errors.New("OpsLevel API Errors:\n\t* User with email 'not-found@opslevel.com' does not exist on this account"), err)
}
