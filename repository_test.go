package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
)

func TestGetRepositoryWithAliasNotFound(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "repository/get_not_found", "repository/get_with_alias")
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, &ol.Repository{}, result)
}

func TestGetRepositoryWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "repository/get", "repository/get_with_alias")
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestGetRepository(t *testing.T) {
	// Arrange
	client := ATestClient(t, "repository/get")
	// Act
	result, err := client.GetRepository("Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestListRepositories(t *testing.T) {
	// Arrange
	client := ATestClient(t, "repository/list")
	// Act
	result, err := client.ListRepositories()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestDeleteServiceRepository(t *testing.T) {
	// Arrange
	client := ATestClient(t, "repository/service/delete")
	// Act
	err := client.DeleteServiceRepository(ol.NewID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3"))
	// Assert
	autopilot.Ok(t, err)
}
