package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

func TestListCampaigns(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query CampaignsList($after:String!$first:Int!$sortBy:CampaignSortEnum!$status:String!){account{campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}]){nodes{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "after": "", "first": 100, "sortBy": "start_date_DESC", "status": "in_progress" }`,
		`{ "data": { "account": { "campaigns": { "nodes": [ {{ template "campaign1_response" }}, {{ template "campaign2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query CampaignsList($after:String!$first:Int!$sortBy:CampaignSortEnum!$status:String!){account{campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}]){nodes{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "after": "OA", "first": 100, "sortBy": "start_date_DESC", "status": "in_progress" }`,
		`{ "data": { "account": { "campaigns": { "nodes": [ {{ template "campaign3_response" }} ], {{ template "pagination_second_pageInfo_response" }} }}}}`,
	)
	requests := []autopilot.TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "campaign/list", requests...)
	// Act
	response, err := client.ListCampaigns(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Campaign 1", result[0].Name)
	autopilot.Equals(t, ol.CampaignStatusEnumInProgress, result[0].Status)
	autopilot.Equals(t, "Campaign 2", result[1].Name)
	autopilot.Equals(t, ol.CampaignStatusEnumInProgress, result[1].Status)
	autopilot.Equals(t, "Campaign 3", result[2].Name)
	autopilot.Equals(t, ol.CampaignStatusEnumInProgress, result[2].Status)
	autopilot.Equals(t, "#engineering", result[0].Reminder.DefaultSlackChannel)
	autopilot.Equals(t, "America/Chicago", result[0].Reminder.Timezone)
	autopilot.Equals(t, "Uses Go", result[1].Filter.Name)
	autopilot.Equals(t, "2024-01-01 00:00:00 +0000 UTC", result[0].StartDate.String())
}

// TestListCampaignsVariables_AsPayloadVariables verifies that ListCampaignsVariables produces the correct payload map.
func TestListCampaignsVariables_AsPayloadVariables(t *testing.T) {
	after := "cursor"
	first := 5
	sortBy := ol.CampaignSortEnumStartDateAsc
	status := ol.CampaignStatusEnumScheduled
	variables := (&ol.ListCampaignsVariables{
		After:  &after,
		First:  &first,
		SortBy: &sortBy,
		Status: &status,
	}).AsPayloadVariables()
	expected := ol.PayloadVariables{
		"after":  after,
		"first":  first,
		"sortBy": sortBy,
		"status": string(status),
	}
	autopilot.Equals(t, expected, *variables)
}

// TestListCampaignsWithCustomVariables verifies that custom ListCampaignsVariables values are sent in the GraphQL request.
func TestListCampaignsWithCustomVariables(t *testing.T) {
	after := "cursor"
	first := 5
	sortBy := ol.CampaignSortEnumStartDateAsc
	status := ol.CampaignStatusEnumDelayed
	testRequest := autopilot.NewTestRequest(
		`query CampaignsList($after:String!$first:Int!$sortBy:CampaignSortEnum!$status:String!){account{campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}]){nodes{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "after": "cursor", "first": 5, "sortBy": "start_date_ASC", "status": "delayed" }`,
		`{ "data": { "account": { "campaigns": { "nodes": [ {{ template "campaign1_response" }} ], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": null, "endCursor": null } }}}}`,
	)
	client := BestTestClient(t, "campaign/list_custom", testRequest)
	// Act
	response, err := client.ListCampaigns(&ol.ListCampaignsVariables{
		After:  &after,
		First:  &first,
		SortBy: &sortBy,
		Status: &status,
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, response.TotalCount)
	autopilot.Equals(t, "Campaign 1", response.Nodes[0].Name)
}

func TestListCampaignsEmpty(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`query CampaignsList($after:String!$first:Int!$sortBy:CampaignSortEnum!$status:String!){account{campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}]){nodes{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "after": "", "first": 100, "sortBy": "start_date_DESC", "status": "in_progress" }`,
		`{ "data": { "account": { "campaigns": { "nodes": [], "pageInfo": { "hasNextPage": false, "hasPreviousPage": false, "startCursor": null, "endCursor": null } }}}}`,
	)

	client := BestTestClient(t, "campaign/list_empty", testRequest)
	// Act
	response, err := client.ListCampaigns(nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 0, response.TotalCount)
	autopilot.Equals(t, 0, len(response.Nodes))
}
