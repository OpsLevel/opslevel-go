package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2026"
	"github.com/relvacode/iso8601"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_create_request" }}`,
		`{{ template "campaign_create_request_vars" }}`,
		`{{ template "campaign_create_response" }}`,
	)
	client := BestTestClient(t, "campaign/create", testRequest)

	brief := "A test campaign"
	// Act
	campaign, err := client.CreateCampaign(ol.CampaignCreateInput{
		Name:         "New Campaign",
		OwnerId:      id1,
		FilterId:     ol.RefOf(id2),
		ProjectBrief: &brief,
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "New Campaign", campaign.Name)
	autopilot.Equals(t, ol.CampaignStatusEnumDraft, campaign.Status)
	autopilot.Equals(t, id1, campaign.Owner.Id)
	autopilot.Equals(t, id2, campaign.Filter.Id)
	autopilot.Equals(t, "A test campaign", campaign.RawProjectBrief)
}

func TestGetCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_get_request" }}`,
		`{{ template "campaign_get_request_vars" }}`,
		`{{ template "campaign_get_response" }}`,
	)
	client := BestTestClient(t, "campaign/get", testRequest)

	// Act
	campaign, err := client.GetCampaign(id1)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, campaign.Id)
	autopilot.Equals(t, "Fetched Campaign", campaign.Name)
	autopilot.Equals(t, ol.CampaignStatusEnumScheduled, campaign.Status)
	autopilot.Equals(t, "2026-05-01 00:00:00 +0000 UTC", campaign.StartDate.String())
	autopilot.Equals(t, "2026-06-30 00:00:00 +0000 UTC", campaign.TargetDate.String())
}

func TestUpdateCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_update_request" }}`,
		`{{ template "campaign_update_request_vars" }}`,
		`{{ template "campaign_update_response" }}`,
	)
	client := BestTestClient(t, "campaign/update", testRequest)

	name := "Updated Campaign"
	// Act
	campaign, err := client.UpdateCampaign(ol.CampaignUpdateInput{
		Id:      id1,
		Name:    &name,
		OwnerId: ol.RefOf(id2),
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, campaign.Id)
	autopilot.Equals(t, "Updated Campaign", campaign.Name)
	autopilot.Equals(t, id2, campaign.Owner.Id)
}

func TestDeleteCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_delete_request" }}`,
		`{{ template "campaign_delete_request_vars" }}`,
		`{{ template "campaign_delete_response" }}`,
	)
	client := BestTestClient(t, "campaign/delete", testRequest)

	// Act
	err := client.DeleteCampaign(id1)

	// Assert
	autopilot.Ok(t, err)
}

func TestScheduleCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_schedule_request" }}`,
		`{{ template "campaign_schedule_request_vars" }}`,
		`{{ template "campaign_schedule_response" }}`,
	)
	client := BestTestClient(t, "campaign/schedule", testRequest)

	startDate, _ := iso8601.ParseString("2026-05-01T00:00:00Z")
	targetDate, _ := iso8601.ParseString("2026-06-30T00:00:00Z")

	// Act
	campaign, err := client.ScheduleCampaign(ol.CampaignScheduleUpdateInput{
		Id:         id1,
		StartDate:  iso8601.Time{Time: startDate},
		TargetDate: iso8601.Time{Time: targetDate},
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, campaign.Id)
	autopilot.Equals(t, ol.CampaignStatusEnumScheduled, campaign.Status)
	autopilot.Equals(t, "2026-05-01 00:00:00 +0000 UTC", campaign.StartDate.String())
	autopilot.Equals(t, "2026-06-30 00:00:00 +0000 UTC", campaign.TargetDate.String())
}

func TestUnscheduleCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_unschedule_request" }}`,
		`{{ template "campaign_unschedule_request_vars" }}`,
		`{{ template "campaign_unschedule_response" }}`,
	)
	client := BestTestClient(t, "campaign/unschedule", testRequest)

	// Act
	campaign, err := client.UnscheduleCampaign(id1)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, campaign.Id)
	autopilot.Equals(t, ol.CampaignStatusEnumDraft, campaign.Status)
	autopilot.Equals(t, true, campaign.StartDate.IsZero())
}

func TestCopyChecksToCampaign(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`{{ template "campaign_copy_checks_request" }}`,
		`{{ template "campaign_copy_checks_request_vars" }}`,
		`{{ template "campaign_copy_checks_response" }}`,
	)
	client := BestTestClient(t, "campaign/copy_checks", testRequest)

	// Act
	campaign, err := client.CopyChecksToCampaign(ol.ChecksCopyToCampaignInput{
		CampaignId: id1,
		CheckIds:   []ol.ID{id2, id3},
	})

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, campaign.Id)
	autopilot.Equals(t, 2, campaign.CheckStats.Total)
}

func TestListCampaigns(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query CampaignsList($after:String!$first:Int!$sortBy:CampaignSortEnum!$status:String!){account{campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}]){nodes{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "after": "", "first": 500, "sortBy": "start_date_DESC", "status": "in_progress" }`,
		`{ "data": { "account": { "campaigns": { "nodes": [ {{ template "campaign1_response" }}, {{ template "campaign2_response" }} ], {{ template "pagination_initial_pageInfo_response" }} }}}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query CampaignsList($after:String!$first:Int!$sortBy:CampaignSortEnum!$status:String!){account{campaigns(first: $first, after: $after, sortBy: $sortBy, filter: [{key: status, arg: $status}]){nodes{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}`,
		`{ "after": "OA", "first": 500, "sortBy": "start_date_DESC", "status": "in_progress" }`,
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
		`{ "after": "", "first": 500, "sortBy": "start_date_DESC", "status": "in_progress" }`,
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
