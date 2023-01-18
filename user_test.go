package opslevel_test

import (
	"errors"
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestInviteUser(t *testing.T) {
	// Arrange
	client := ATestClient(t, "user/invite")
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
	client := ATestClient(t, "user/get")
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
	client := ATestClient(t, "user/list")
	// Act
	resp, err := client.ListUsers(nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "Edgar Ochoa", result[1].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[1].Role)
}

func TestListUserBetter(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			`{"query":
	"query ($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{email,htmlUrl,id,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
	"variables": {"after":"","first":100}
}`,
			`{"data": {"account": {"users": {
      "nodes": [
        {
          "email": "kyle@opslevel.com",
          "id": "1",
          "name": "Kyle Rockman",
          "role": "user"
        },
        {
          "email": "edgar@opslevel.com",
          "id": "2",
          "name": "Edgar Ochoa",
          "role": "admin"
        }
      ],
      "pageInfo": {
        "hasNextPage": true,
        "hasPreviousPage": false,
        "startCursor": "MQ",
        "endCursor": "NDc"
      }}}}}`,
		},
		{
			`"query": "query ($after:String!$first:Int!){account{users(after: $after, first: $first){nodes{email,htmlUrl,id,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
"variables": {
  "after":"NDc",
  "first":100
}}`,
			`{"data": {"account": {"users": {
      "nodes": [
        {
          "email": "matthew@opslevel.com",
          "id": "3",
          "name": "Matthew Brahms",
          "role": "admin"
        }
      ],
      "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": true,
        "startCursor": "NDc",
        "endCursor": "Mx"
      }}}}}`,
		},
	}
	client := APaginatedTestClient(t, "user/list_paginated", requests...)

	// Act
	resp, err := client.ListUsers(nil)
	result := resp.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, len(result))
	autopilot.Equals(t, "Edgar Ochoa", result[1].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[1].Role)
	autopilot.Equals(t, "Matthew Brahms", result[2].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[2].Role)
}

func TestUpdateUser(t *testing.T) {
	// Arrange
	client := ATestClient(t, "user/update")
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
	client := ATestClient(t, "user/delete")
	// Act
	err := client.DeleteUser("kyle@opslevel.com")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteUserDoesNotExist(t *testing.T) {
	// Arrange
	client := ATestClient(t, "user/delete_not_found")
	// Act
	err := client.DeleteUser("not-found@opslevel.com")
	// Assert
	autopilot.Equals(t, errors.New("OpsLevel API Errors:\n\t* User with email 'not-found@opslevel.com' does not exist on this account"), err)
}
