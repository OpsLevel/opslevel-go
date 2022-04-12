package opslevel_test

import (
	"testing"

	"github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

func TestCreateGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/create")
	// Act
	members := []opslevel.MemberInput{
		opslevel.MemberInput{Email: "edgar+test@opslevel.com"},
	}
	result, err := client.CreateGroup(opslevel.GroupInput{
		Name:        "platform",
		Description: "Another test group",
		Members:     members,
		Parent:      &opslevel.IdentifierInput{Alias: "test_group_1"},
		Teams: []opslevel.IdentifierInput{
			opslevel.IdentifierInput{Alias: "platform"},
		},
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
	err := client.DeleteGroupWithAlias("platform")
	// Assert
	autopilot.Ok(t, err)
}

func TestGetGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/get")
	// Act
	result, err := client.GetGroup("Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Alias)
	autopilot.Equals(t, nil, result.Parent.Id)
}

func TestGetGroupWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "group/get", "group/get_with_alias")
	// Act
	result, err := client.GetGroupWithAlias("test_group_1")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Alias)
	autopilot.Equals(t, nil, result.Parent.Id)
}

func TestListGroups(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/list")
	// Act
	result, err := client.ListGroups()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "test_group_2", result[0].Alias)
	autopilot.Equals(t, "test_group_1", result[1].Alias)
}

func TestUpdateGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/update")
	members := []opslevel.MemberInput{
		opslevel.MemberInput{Email: "edgar+test@opslevel.com"},
	}
	// Act
	result, err := client.UpdateGroup(opslevel.IdentifierInput{Id: "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI"}, opslevel.GroupInput{
		Description: "This is the first test group",
		Members:     members,
		Parent:      &opslevel.IdentifierInput{Alias: "test_group_2"},
		Teams: []opslevel.IdentifierInput{
			{Alias: "platform"},
		},
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Name)
	autopilot.Equals(t, "This is the first test group", result.Description)
	autopilot.Equals(t, "test_group_2", result.Parent.Alias)
}
