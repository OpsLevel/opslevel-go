package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2022"
	"github.com/rocktavious/autopilot"
	"testing"
)

func TestGetAlertSourceWithExternalIdentifier(t *testing.T) {
	//Arrange
	client := ATestClientAlt(t, "alert_source/get", "alert_source/get_with_external_identifier")
	// Act
	result, err := client.GetAlertSourceWithExternalIdentifier(ol.AlertSourceExternalIdentifier{
		Type: 			ol.AlertSourceTypeEnumDatadog,
		ExternalId: "12345678",
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "test", result.Description)
	autopilot.Equals(t, ol.AlertSourceTypeEnumDatadog, result.Type)
}

func TestGetAlertSource(t *testing.T) {
	//Arrange
	client := ATestClient(t, "alert_source/get")
	// Act
	result, err := client.GetAlertSource("Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "test", result.Description)
	autopilot.Equals(t, ol.AlertSourceTypeEnumDatadog, result.Type)
}