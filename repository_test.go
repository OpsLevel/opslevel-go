package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestGetRepositoryWithAliasNotFound(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/repository_with_alias_not_found", autopilot.FixtureResponse("repository_not_found_response.json"), FixtureQueryValidation(t, "repository_with_alias_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, &Repository{}, result)
}

func TestGetRepositoryWithAlias(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/repository_with_alias", autopilot.FixtureResponse("repository_response.json"), FixtureQueryValidation(t, "repository_with_alias_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "Developers", result.Owner.Name)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestGetRepository(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/repository", autopilot.FixtureResponse("repository_response.json"), FixtureQueryValidation(t, "repository_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.GetRepository("Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpHaXRodWIvMjY1MTk")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "main", result.DefaultBranch)
	autopilot.Equals(t, "autopilot", result.Name)
	autopilot.Equals(t, "359666903", result.RepoKey)
	autopilot.Equals(t, "Developers", result.Owner.Name)
	autopilot.Equals(t, "tier_2", result.Tier.Alias)
}

func TestListRepositories(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/list_repositories", autopilot.FixtureResponse("repository_list_response.json"), FixtureQueryValidation(t, "repository_list_request.json"))
	client := NewClient("X", SetURL(url))
	// Act
	result, err := client.ListRepositories()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}
