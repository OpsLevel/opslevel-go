package opslevel_test

import (
	"errors"
	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
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
	autopilot.Equals(t, "1", result.Id)
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
	autopilot.Equals(t, "1", result.Id)
	autopilot.Equals(t, "Kyle Rockman", result.Name)
	autopilot.Equals(t, ol.UserRoleUser, result.Role)
}

func TestListUser(t *testing.T) {
	// Arrange
	client := ATestClient(t, "user/list")
	// Act
	result, err := client.ListUsers()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "Edgar Ochoa", result[1].Name)
	autopilot.Equals(t, ol.UserRoleAdmin, result[1].Role)
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
	autopilot.Equals(t, "1", result.Id)
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
