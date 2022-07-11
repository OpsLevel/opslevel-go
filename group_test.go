package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
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
	err := client.DeleteGroupWithAlias("platform")
	// Assert
	autopilot.Ok(t, err)
}

func TestChildTeams(t *testing.T) {
	// Arrange
	client1 := getGroupWithAliasTestClient(t)
	client2 := ATestClient(t, "group/child_teams")
	// Act
	group, err := client1.GetGroupWithAlias("test_group_1")
	result, err := group.ChildTeams(client2)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "platform", result[0].Alias)
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
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8yNjE5", result[0].Id)
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
	autopilot.Equals(t, nil, result.Parent.Id)
}

func TestGetGroupWithAlias(t *testing.T) {
	// Arrange
	client := getGroupWithAliasTestClient(t)
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
