package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestEntityOwnerGroupReturnsCorrectId(t *testing.T) {
	// Arrange
	owner := ol.EntityOwner{
		OnGroup: ol.EntityOwnerGroup{
			Id:    id1,
			Alias: "Example",
		},
	}
	// Act
	// Assert
	autopilot.Equals(t, id1, owner.Id())
	autopilot.Equals(t, id1, owner.OnGroup.AsGroup().Id)
	autopilot.Equals(t, "Example", owner.Alias())
	autopilot.Equals(t, "Example", owner.OnGroup.AsGroup().Alias)
}

func TestEntityOwnerTeamReturnsCorrectId(t *testing.T) {
	// Arrange
	owner := ol.EntityOwner{
		OnTeam: ol.EntityOwnerTeam{
			Id:    id1,
			Alias: "Example",
		},
	}
	// Act
	// Assert
	autopilot.Equals(t, id1, owner.Id())
	autopilot.Equals(t, id1, owner.OnTeam.AsTeam().Id)
	autopilot.Equals(t, "Example", owner.Alias())
	autopilot.Equals(t, "Example", owner.OnTeam.AsTeam().Alias)
}
