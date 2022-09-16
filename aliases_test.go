package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2022"
)

func TestCreateAliases(t *testing.T) {
	// Arrange
	client := ATestClient(t, "aliases/create")
	// Act
	result, err := client.CreateAliases("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", []string{"MyAwesomeAlias"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "MyAwesomeAlias", result[1])
}

// TODO: Add CreateAliasesOwnerNotFound

func TestDeleteServiceAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "aliases/delete", "aliases/delete_service")
	// Act
	err := client.DeleteServiceAlias("MyAwesomeAlias")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteTeamAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "aliases/delete", "aliases/delete_team")
	// Act
	err := client.DeleteTeamAlias("MyAwesomeAlias")
	// Assert
	autopilot.Ok(t, err)
}

// TODO: Add DeleteAliasNotFound
