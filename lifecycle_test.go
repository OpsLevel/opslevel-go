package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2023"
)

func TestListLifecycles(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query LifecycleList{account{lifecycles{alias,description,id,index,name}}}`,
		`{}`,
		`{"data": { "account": {
		"lifecycles": [
			{{ template "lifecycle-pre-alpha" }},
			{{ template "lifecycle-alpha" }},
			{{ template "lifecycle-beta" }},
			{{ template "lifecycle-ga" }},
			{{ template "lifecycle-eol" }}
		]
    }}}`,
	)
	client := BestTestClient(t, "lifecycles", testRequest)
	// Act
	result, err := client.ListLifecycles()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "pre-alpha", result[0].Alias)
}
