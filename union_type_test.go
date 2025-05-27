package opslevel_test

import (
	"encoding/json"
	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
	"testing"
)

func TestRelationshipResourceUnmarshalJSON(t *testing.T) {
	// Arrange
	testCases := []struct {
		name     string
		input    string
		expected ol.RelationshipResource
	}{
		{
			name:  "Domain resource",
			input: `{"__typename": "Domain", "id": "Z2lkOi8vb3BzbGV2ZWwvRG9tYWluLzE"}`,
			expected: ol.RelationshipResource{
				Type: ol.TypeIdentifier("Domain"),
				Domain: ol.DomainId{
					Id: "Z2lkOi8vb3BzbGV2ZWwvRG9tYWluLzE",
				},
			},
		},
		{
			name:  "Service resource",
			input: `{"__typename": "Service", "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8x"}`,
			expected: ol.RelationshipResource{
				Type: ol.TypeIdentifier("Service"),
				Service: ol.ServiceId{
					Id: ol.ID("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8x"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			var result ol.RelationshipResource
			err := json.Unmarshal([]byte(tc.input), &result)

			// Assert
			autopilot.Ok(t, err)
			autopilot.Equals(t, tc.expected.Type, result.Type)
			switch tc.expected.Type {
			case "Domain":
				autopilot.Equals(t, tc.expected.Domain, result.Domain)
			case "Service":
				autopilot.Equals(t, tc.expected.Service.Id, result.Service.Id)
			}
		})
	}
}
