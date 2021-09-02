package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go"
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
	autopilot.Equals(t, "Developers", result.Owner.Name)
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
	autopilot.Equals(t, "Developers", result.Owner.Name)
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
