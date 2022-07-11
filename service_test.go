package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
)

func TestCreateService(t *testing.T) {
	// Arrange
	client := ATestClient(t, "service/create")
	// Act
	result, err := client.CreateService(ol.ServiceCreateInput{
		Name:        "Foo",
		Description: "Foo service",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result.Aliases))
}

func TestGetServiceIdWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/get_id", "service/get_id_with_alias")
	// Act
	result, err := client.GetServiceIdWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx", result.Id)
}

func TestGetServiceWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/get", "service/get_with_alias")
	// Act
	result, err := client.GetServiceWithAlias("coredns")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
	autopilot.Equals(t, 4, result.Tools.TotalCount)
}

func TestGetService(t *testing.T) {
	// Arrange
	client := ATestClient(t, "service/get")
	// Act
	result, err := client.GetService("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS81MzEx")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result.Aliases))
	autopilot.Equals(t, "alpha", result.Lifecycle.Alias)
	autopilot.Equals(t, "developers", result.Owner.Alias)
	autopilot.Equals(t, "tier_1", result.Tier.Alias)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
	autopilot.Equals(t, 4, result.Tools.TotalCount)
}

func TestListServices(t *testing.T) {
	// Arrange
	client := ATestClient(t, "service/list")
	// Act
	result, err := client.ListServices()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithFramework(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_framework")
	// Act
	result, err := client.ListServicesWithFramework("postgres")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithLanguage(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_language")
	// Act
	result, err := client.ListServicesWithLanguage("postgres")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithOwner(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_owner")
	// Act
	result, err := client.ListServicesWithOwner("postgres")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestListServicesWithTag(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/list", "service/list_with_tag")
	// Act
	result, err := client.ListServicesWithTag(ol.NewTagArgs("app:worker"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "generally_available", result[0].Lifecycle.Alias)
}

func TestDeleteService(t *testing.T) {
	// Arrange
	client := ATestClient(t, "service/delete")
	// Act
	err := client.DeleteService(ol.ServiceDeleteInput{Id: ol.NewID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82NzQ3")})
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteServicesWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "service/delete", "service/delete_with_alias")
	// Act
	err := client.DeleteServiceWithAlias("db")
	// Assert
	autopilot.Ok(t, err)
}
