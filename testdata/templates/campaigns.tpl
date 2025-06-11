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
    "status": "ended",
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
    "status": "draft",
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
