package opslevel

import (
	"testing"

	"github.com/rocktavious/autopilot"
)

// func TestCreateServiceRepository(t *testing.T) {
// 	// Arrange
// 	urlGetRepo := autopilot.RegisterEndpoint("/repository_with_alias_2", autopilot.FixtureResponse("repository_response.json"), FixtureQueryValidation(t, "repository_with_alias_request.json"))
// 	clientGetRepo := NewClient("X", SetURL(urlGetRepo))

// 	urlGetService := autopilot.RegisterEndpoint("/service_id_with_alias_2", autopilot.FixtureResponse("service_id_response.json"), FixtureQueryValidation(t, "service_id_with_alias_request.json"))
// 	clientGetService := NewClient("X", SetURL(urlGetService))

// 	url := autopilot.RegisterEndpoint("/connect_service_repository", autopilot.FixtureResponse("service_repository_response.json"), FixtureQueryValidation(t, "connect_service_repository_request.json"))
// 	client := NewClient("X", SetURL(url))
// 	// Act
// 	repo, repoErr := clientGetRepo.GetRepositoryWithAlias("github.com:rocktavious/autopilot")
// 	service, serviceErr := clientGetService.GetServiceIdWithAlias("coredns")
// 	result, err := client.ConnectServiceRepository(service, repo)
// 	// Assert
// 	autopilot.Ok(t, repoErr)
// 	autopilot.Ok(t, serviceErr)
// 	autopilot.Ok(t, err)
// 	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvNDIxNA", result.Id)
// }

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
