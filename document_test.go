package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2022"
	"testing"

	"github.com/rocktavious/autopilot"
)

func TestServiceApiDocSettingsUpdate(t *testing.T) {
	// Arrange
	client := ATestClient(t, "service/api_doc_settings")
	// Act
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "/src/swagger.json", ol.ApiDocumentSourceEnumPull)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8zOTI4MQ", result.Id)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}
