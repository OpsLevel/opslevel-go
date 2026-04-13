{{- define "campaign1_response" }}
{
    {{ template "id1" }},
    "name": "Campaign 1",
    "htmlUrl": "https://app.opslevel.com/campaigns/campaign1",
    "status": "in_progress",
    "checkStats": {
        "total": 10,
        "totalSuccessful": 8
    },
    "serviceStats": {
        "total": 5,
        "totalSuccessful": 4
    },
    "owner": {
        {{ template "id1" }},
        "alias": "platform"
    },
    "startDate": "2024-01-01T00:00:00Z",
    "projectBrief": "First campaign for testing",
    "rawProjectBrief": "First campaign for testing",
    "filter": null,
    "reminder": {
              "channels": [
                "slack"
              ],
              "daysOfWeek": [
                "monday",
                "wednesday"
              ],
              "defaultSlackChannel": "#engineering",
              "frequency": 1,
              "frequencyUnit": "week",
              "message": "",
              "nextOccurrence": null,
              "timeOfDay": "10:31",
              "timezone": "America/Chicago"
            }
}
{{end}}
{{- define "campaign2_response" }}
{
    {{ template "id2" }},
    "name": "Campaign 2",
    "htmlUrl": "https://app.opslevel.com/campaigns/campaign2",
    "status": "in_progress",
    "checkStats": {
        "total": 5,
        "totalSuccessful": 5
    },
    "serviceStats": {
        "total": 8,
        "totalSuccessful": 7
    },
    "owner": null,
    "startDate": "2024-02-01T00:00:00Z",
    "endedDate": "2024-03-01T00:00:00Z",
    "projectBrief": "Second campaign for testing",
    "rawProjectBrief": "Second campaign for testing",
    "filter": {
              "id": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEwNDU",
              "name": "Uses Go"
            },
    "reminder": null
}
{{end}}
{{- define "campaign3_response" }}
{
    {{ template "id3" }},
    "name": "Campaign 3",
    "htmlUrl": "https://app.opslevel.com/campaigns/campaign3",
    "status": "in_progress",
    "checkStats": {
        "total": 1,
        "totalSuccessful": 0
    },
    "serviceStats": {
        "total": 1,
        "totalSuccessful": 0
    },
    "owner": {
        {{ template "id3" }},
        "alias": "frontend"
    },
    "projectBrief": "Third campaign for testing",
    "rawProjectBrief": "Third campaign for testing",
    "filter": null,
    "reminder": null
}
{{end}}

{{- define "campaign_create_request" }}
mutation CampaignCreate($input:CampaignCreateInput!){campaignCreate(input: $input){campaign{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
{{ end }}

{{- define "campaign_create_request_vars" }}
{"input":{"name":"New Campaign","ownerId":"{{ template "id1_string" }}","filterId":"{{ template "id2_string" }}","projectBrief":"A test campaign"}}
{{ end }}

{{- define "campaign_create_response" }}{
    "data":{"campaignCreate":{"campaign":{
        {{ template "id1" }},
        "name":"New Campaign",
        "htmlUrl":"https://app.opslevel.com/campaigns/new-campaign",
        "status":"draft",
        "checkStats":{"total":0,"totalSuccessful":0},
        "serviceStats":{"total":0,"totalSuccessful":0},
        "owner":{ {{ template "id1" }}, "alias":"platform" },
        "projectBrief":"A test campaign",
        "rawProjectBrief":"A test campaign",
        "filter":{ "id":"{{ template "id2_string" }}", "name":"Tier 1" },
        "reminder":null
    },"errors":[]}}
}{{ end }}

{{- define "campaign_get_request" }}
query CampaignGet($id:ID!){account{campaign(id: $id){checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate}}}
{{ end }}

{{- define "campaign_get_request_vars" }}
{"id":"{{ template "id1_string" }}"}
{{ end }}

{{- define "campaign_get_response" }}{
    "data":{"account":{"campaign":{
        {{ template "id1" }},
        "name":"Fetched Campaign",
        "htmlUrl":"https://app.opslevel.com/campaigns/fetched",
        "status":"scheduled",
        "checkStats":{"total":3,"totalSuccessful":1},
        "serviceStats":{"total":10,"totalSuccessful":5},
        "owner":{ {{ template "id1" }}, "alias":"platform" },
        "startDate":"2026-05-01T00:00:00Z",
        "targetDate":"2026-06-30T00:00:00Z",
        "projectBrief":"Fetched campaign brief",
        "rawProjectBrief":"Fetched campaign brief",
        "filter":null,
        "reminder":null
    }}}
}{{ end }}

{{- define "campaign_update_request" }}
mutation CampaignUpdate($input:CampaignUpdateInput!){campaignUpdate(input: $input){campaign{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
{{ end }}

{{- define "campaign_update_request_vars" }}
{"input":{"id":"{{ template "id1_string" }}","name":"Updated Campaign","ownerId":"{{ template "id2_string" }}"}}
{{ end }}

{{- define "campaign_update_response" }}{
    "data":{"campaignUpdate":{"campaign":{
        {{ template "id1" }},
        "name":"Updated Campaign",
        "htmlUrl":"https://app.opslevel.com/campaigns/updated",
        "status":"draft",
        "checkStats":{"total":0,"totalSuccessful":0},
        "serviceStats":{"total":0,"totalSuccessful":0},
        "owner":{ {{ template "id2" }}, "alias":"staff" },
        "projectBrief":"A test campaign",
        "rawProjectBrief":"A test campaign",
        "filter":null,
        "reminder":null
    },"errors":[]}}
}{{ end }}

{{- define "campaign_delete_request" }}
mutation CampaignDelete($input:DeleteInput!){campaignDelete(input: $input){deletedCampaignId,errors{message,path}}}
{{ end }}

{{- define "campaign_delete_request_vars" }}
{"input":{"id":"{{ template "id1_string" }}"}}
{{ end }}

{{- define "campaign_delete_response" }}{
    "data":{"campaignDelete":{"deletedCampaignId":"{{ template "id1_string" }}","errors":[]}}
}{{ end }}

{{- define "campaign_schedule_request" }}
mutation CampaignScheduleUpdate($input:CampaignScheduleUpdateInput!){campaignScheduleUpdate(input: $input){campaign{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
{{ end }}

{{- define "campaign_schedule_request_vars" }}
{"input":{"id":"{{ template "id1_string" }}","startDate":"2026-05-01T00:00:00Z","targetDate":"2026-06-30T00:00:00Z"}}
{{ end }}

{{- define "campaign_schedule_response" }}{
    "data":{"campaignScheduleUpdate":{"campaign":{
        {{ template "id1" }},
        "name":"New Campaign",
        "htmlUrl":"https://app.opslevel.com/campaigns/new-campaign",
        "status":"scheduled",
        "checkStats":{"total":0,"totalSuccessful":0},
        "serviceStats":{"total":0,"totalSuccessful":0},
        "owner":{ {{ template "id1" }}, "alias":"platform" },
        "startDate":"2026-05-01T00:00:00Z",
        "targetDate":"2026-06-30T00:00:00Z",
        "projectBrief":"A test campaign",
        "rawProjectBrief":"A test campaign",
        "filter":null,
        "reminder":null
    },"errors":[]}}
}{{ end }}

{{- define "campaign_unschedule_request" }}
mutation CampaignUnschedule($input:DeleteInput!){campaignUnschedule(input: $input){campaign{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
{{ end }}

{{- define "campaign_unschedule_request_vars" }}
{"input":{"id":"{{ template "id1_string" }}"}}
{{ end }}

{{- define "campaign_unschedule_response" }}{
    "data":{"campaignUnschedule":{"campaign":{
        {{ template "id1" }},
        "name":"New Campaign",
        "htmlUrl":"https://app.opslevel.com/campaigns/new-campaign",
        "status":"draft",
        "checkStats":{"total":0,"totalSuccessful":0},
        "serviceStats":{"total":0,"totalSuccessful":0},
        "owner":{ {{ template "id1" }}, "alias":"platform" },
        "projectBrief":"A test campaign",
        "rawProjectBrief":"A test campaign",
        "filter":null,
        "reminder":null
    },"errors":[]}}
}{{ end }}

{{- define "campaign_copy_checks_request" }}
mutation ChecksCopyToCampaign($input:ChecksCopyToCampaignInput!){checksCopyToCampaign(input: $input){campaign{checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,id,name,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
{{ end }}

{{- define "campaign_copy_checks_request_vars" }}
{"input":{"campaignId":"{{ template "id1_string" }}","checkIds":["{{ template "id2_string" }}","{{ template "id3_string" }}"]}}
{{ end }}

{{- define "campaign_copy_checks_response" }}{
    "data":{"checksCopyToCampaign":{"campaign":{
        {{ template "id1" }},
        "name":"New Campaign",
        "htmlUrl":"https://app.opslevel.com/campaigns/new-campaign",
        "status":"scheduled",
        "checkStats":{"total":2,"totalSuccessful":0},
        "serviceStats":{"total":10,"totalSuccessful":0},
        "owner":{ {{ template "id1" }}, "alias":"platform" },
        "startDate":"2026-05-01T00:00:00Z",
        "targetDate":"2026-06-30T00:00:00Z",
        "projectBrief":"A test campaign",
        "rawProjectBrief":"A test campaign",
        "filter":null,
        "reminder":null
    },"errors":[]}}
}{{ end }}
