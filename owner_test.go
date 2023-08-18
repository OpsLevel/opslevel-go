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
			Id:    "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
			Alias: "Example",
		},
	}
	// Act
	// Assert
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), owner.Id())
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), owner.OnGroup.AsGroup().Id)
	autopilot.Equals(t, "Example", owner.Alias())
	autopilot.Equals(t, "Example", owner.OnGroup.AsGroup().Alias)
}

func TestEntityOwnerTeamReturnsCorrectId(t *testing.T) {
	// Arrange
	owner := ol.EntityOwner{
		OnTeam: ol.EntityOwnerTeam{
			Id:    "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
			Alias: "Example",
		},
	}
	// Act
	// Assert
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), owner.Id())
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), owner.OnTeam.AsTeam().Id)
	autopilot.Equals(t, "Example", owner.Alias())
	autopilot.Equals(t, "Example", owner.OnTeam.AsTeam().Alias)
}
