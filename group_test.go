package opslevel_test

import (
	"fmt"
	"testing"

	"github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

func TestCreateGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/create")
	// Act
	members := []opslevel.GroupMemberInput{
		opslevel.GroupMemberInput{Email: "edgar+test@opslevel.com"},
	}
	result, err := client.CreateGroup(opslevel.GroupCreateInput{
		Name:        "platform",
		Description: "Another test group",
		Members:     members,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "platform", result.Name)
	autopilot.Equals(t, "Another test group", result.Description)
	autopilot.Equals(t, "edgar+test@opslevel.com", result.Members.Nodes[0].Email)
}

func TestGetGroup(t *testing.T) {
	// Arrange
	client := ATestClient(t, "group/get")
	// Act
	result, err := client.GetGroup("Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI")
	// Assert
	fmt.Printf("%v", result)
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Alias)
	autopilot.Equals(t, nil, result.Parent.GroupId.Id)
}

func TestGetGroupWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "group/get", "group/get_with_alias")
	// Act
	result, err := client.GetGroupWithAlias("test_group_1")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "test_group_1", result.Alias)
	autopilot.Equals(t, nil, result.Parent.GroupId.Id)
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
