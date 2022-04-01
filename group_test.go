package opslevel_test

import (
	"fmt"
	"testing"

	"github.com/rocktavious/autopilot"
)

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
