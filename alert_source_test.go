package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateAlertSourceService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AlertSourceServiceCreate($input:AlertSourceServiceCreateInput!){alertSourceServiceCreate(input: $input){alertSourceService{alertSource{name,description,id,type,externalId,integration{id,name,type},url},id,service{id,aliases},status},errors{message,path}}}`,
		`{"input": { "alertSourceExternalIdentifier": { "externalId": "QWERTY", "type": "datadog" }, "service": { "alias": "example" }}}`,
		`{"data": { "alertSourceServiceCreate": { "alertSourceService": { "service": { "aliases": ["example"] }}}}}`,
	)

	client := BestTestClient(t, "alert_source/create", testRequest)
	// Act
	result, _ := client.CreateAlertSourceService(ol.AlertSourceServiceCreateInput{
		Service:                       *ol.NewIdentifier("example"),
		AlertSourceExternalIdentifier: ol.NewAlertSource(ol.AlertSourceTypeEnumDatadog, "QWERTY"),
	})
	// Assert
	autopilot.Equals(t, "example", result.Service.Aliases[0])
}

func TestGetAlertSourceWithExternalIdentifier(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query AlertSourceGet($externalIdentifier:AlertSourceExternalIdentifier!){account{alertSource(externalIdentifier: $externalIdentifier){name,description,id,type,externalId,integration{id,name,type},url}}}`,
		`{"externalIdentifier": { "type": "datadog", "externalId": "12345678" }}`,
		`{"data": {
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
        }}}`,
	)

	client := BestTestClient(t, "alert_source/get_with_external_identifier", testRequest)
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
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query AlertSourceGet($id:ID!){account{alertSource(id: $id){name,description,id,type,externalId,integration{id,name,type},url}}}`,
		`{"id": "Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE" }`,
		`{"data": {
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
        }}}`,
	)
	client := BestTestClient(t, "alert_source/get", testRequest)
	// Act
	result, err := client.GetAlertSource("Z2lkOi8vb3BzbGV2ZWwvQWxlcnRTb3VyY2VzOjpQYWdlcmR1dHkvNjE")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "test", result.Description)
	autopilot.Equals(t, ol.AlertSourceTypeEnumDatadog, result.Type)
}

func TestDeleteAlertSourceService(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation AlertSourceServiceDelete($input:AlertSourceDeleteInput!){alertSourceServiceDelete(input: $input){errors{message,path}}}`,
		`{"input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz" }}`,
		`{"data": { "alertSourceServiceDelete": { "errors": [] }}}`,
	)

	client := BestTestClient(t, "alert_source/delete", testRequest)
	// Act
	err := client.DeleteAlertSourceService("Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvODYz")
	// Assert
	autopilot.Equals(t, nil, err)
}
