package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2022"
)

func TestListLifecycles(t *testing.T) {
	// Arrange
	request := `{
    "query": "query LifecycleList{account{lifecycles{alias,description,id,index,name}}}",
	"variables":{}
}`
	response := `{"data": {
	"account": {
		"lifecycles": [
			{{ template "lifecycle-pre-alpha" }},
			{{ template "lifecycle-alpha" }},
			{{ template "lifecycle-beta" }},
			{{ template "lifecycle-ga" }},
			{{ template "lifecycle-eol" }}
		]
	}
}}`
	client := ABetterTestClient(t, "lifecycles", request, response)
	// Act
	result, err := client.ListLifecycles()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "pre-alpha", result[0].Alias)
}
