package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2023"
)

func TestListTiers(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query TierList{account{tiers{alias,description,id,index,name}}}`,
		`{}`,
		`{"data": { "account": { "tiers": [ {{ template "tier_1" }}, {{ template "tier_2" }}, {{ template "tier_3" }}, {{ template "tier_4" }} ] }}}`)
	client := BestTestClient(t, "tiers", testRequest)
	// Act
	result, err := client.ListTiers()
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 4, len(result))
	autopilot.Equals(t, "example", result[0].Alias)
}
