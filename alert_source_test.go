package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestGetAlertSourceWithExternalIdentifier(t *testing.T) {
	//Arrange
	request := `{"query":"query AlertSourceGet($externalIdentifier:AlertSourceExternalIdentifier!){account{alertSource(externalIdentifier: $externalIdentifier){name,description,id,type,externalId,integration{id,name,type},url}}}",
	"variables":{
		"externalIdentifier": {
			"type": "datadog",
			"externalId": "12345678"
		}
	}}`
	response := `{"data": {
	"account": {
		"alertSource": {
			"description": "test",
			"externalId": "12345678",
			"id": "Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE",
			"integration": {
				"name": "test-integration",
				"id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpQYWdlcmR1dHlJbnRlZ3JhdGlvbi8zMg",
				"type": "datadog"
			},
			"name": "Example",
			"type": "datadog",
			"url": "https://example.com"
		}
	}
	}}`
	client := ABetterTestClient(t, "alert_source/get_with_external_identifier", request, response)
	// Act
	result, err := client.GetAlertSourceWithExternalIdentifier(ol.AlertSourceExternalIdentifier{
		Type:       ol.AlertSourceTypeEnumDatadog,
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
	request := `{"query":"query AlertSourceGet($id:ID!){account{alertSource(id: $id){name,description,id,type,externalId,integration{id,name,type},url}}}",
	"variables":{
		"id": "Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE"
	}}`
	response := `{"data": {
	"account": {
		"alertSource": {
			"description": "test",
			"externalId": "12345678",
			"id": "Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE",
			"integration": {
				"name": "test-integration",
				"id": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpQYWdlcmR1dHlJbnRlZ3JhdGlvbi8zMg",
				"type": "datadog"
			},
			"name": "Example",
			"type": "datadog",
			"url": "https://example.com"
		}
	}
	}}`
	client := ABetterTestClient(t, "alert_source/get", request, response)
	// Act
	result, err := client.GetAlertSource("Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "test", result.Description)
	autopilot.Equals(t, ol.AlertSourceTypeEnumDatadog, result.Type)
}
