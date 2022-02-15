package opslevel_test

import (
	"testing"

	"github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"
)

// TODO: not sure if there is a better way to handle reusing a client
// Probably should be a feature of autopilot
var getWithAliasClient *opslevel.Client

func getWithAliasTestClient(t *testing.T) *opslevel.Client {
	if getWithAliasClient == nil {
		getWithAliasClient = ATestClientAlt(t, "team/get", "team/get_with_alias")
	}
	return getWithAliasClient
}

func TestCreateTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/create")
	// Act
	contacts := []opslevel.ContactInput{
		opslevel.CreateContactSlack("#general", ""),
		opslevel.CreateContactWeb("https://example.com", "Homepage"),
	}
	result, err := client.CreateTeam(opslevel.TeamCreateInput{
		Name:             "Example",
		ManagerEmail:     "john@example.com",
		Responsibilities: "Foo & bar",
		Contacts:         contacts,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
}

func TestGetTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/get")
	// Act
	result, err := client.GetTeam("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
}

func TestGetTeamWithAlias(t *testing.T) {
	// Arrange
	client := getWithAliasTestClient(t)
	// Act
	result, err := client.GetTeamWithAlias("example")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
}

func TestListTeams(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/list")
	// Act
	result, err := client.ListTeams()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "devops", result[0].Alias)
	autopilot.Equals(t, "developers", result[1].Alias)
	autopilot.Equals(t, "Own Infra & Tools.", result[0].Responsibilities)
}

func TestUpdateTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/update")
	// Act
	result, err := client.UpdateTeam(opslevel.TeamUpdateInput{
		Id:               opslevel.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ"),
		ManagerEmail:     "ken@example.com",
		Responsibilities: "Foo & bar",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "ken@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
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
	client1 := getWithAliasTestClient(t)
	client2 := ATestClient(t, "team/add_member")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.AddMember(&team.TeamId, "john@example.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamRemoveMember(t *testing.T) {
	// Arrange
	client1 := getWithAliasTestClient(t)
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
	client1 := getWithAliasTestClient(t)
	client2 := ATestClient(t, "team/add_contact")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.AddContact(&team.TeamId, opslevel.CreateContactSlack("#general", ""))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "#general", result.Address)
}

func TestTeamUpdateContact(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/update_contact")
	// Act
	result, err := client.UpdateContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4", opslevel.CreateContactSlack("#general", "Main Channel"))
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
