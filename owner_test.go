package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestEntityOwnerGroupReturnsCorrectId(t *testing.T) {
	// Arrange
	owner := ol.EntityOwner{
		OnGroup: ol.GroupId{
			Id: "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		},
	}
	// Act
	// Assert
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), owner.Id())
}

func TestEntityOwnerTeamReturnsCorrectId(t *testing.T) {
	// Arrange
	owner := ol.EntityOwner{
		OnTeam: ol.TeamId{
			Id: "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		},
	}
	// Act
	// Assert
	autopilot.Equals(t, ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"), owner.Id())
}
