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
mutation CampaignCreate($input:CampaignCreateInput!){campaignCreate(input: $input){campaign{id,name,checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
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
query CampaignGet($id:ID!){account{campaign(id: $id){id,name,checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate}}}
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
mutation CampaignUpdate($input:CampaignUpdateInput!){campaignUpdate(input: $input){campaign{id,name,checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
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
mutation CampaignDelete($input:DeleteInput!){campaignDelete(input: $input){deletedId,errors{message,path}}}
{{ end }}

{{- define "campaign_delete_request_vars" }}
{"input":{"id":"{{ template "id1_string" }}"}}
{{ end }}

{{- define "campaign_delete_response" }}{
    "data":{"campaignDelete":{"deletedId":"{{ template "id1_string" }}","errors":[]}}
}{{ end }}

{{- define "campaign_schedule_request" }}
mutation CampaignScheduleUpdate($input:CampaignScheduleUpdateInput!){campaignScheduleUpdate(input: $input){campaign{id,name,checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
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
mutation CampaignUnschedule($input:CampaignUnscheduleInput!){campaignUnschedule(input: $input){campaign{id,name,checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},errors{message,path}}}
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

{{- define "campaign_list_checks_request" }}
query CampaignChecksList($after:String!$first:Int!$id:ID!){account{campaign(id: $id){checks(first: $first, after: $after){nodes{id,name,sourceCheck{id,name}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}}
{{ end }}

{{- define "campaign_list_checks_request_vars" }}
{"after":"","first":500,"id":"{{ template "id1_string" }}"}
{{ end }}

{{- define "campaign_list_checks_response" }}{
    "data":{"account":{"campaign":{"checks":{"nodes":[
        {"id":"{{ template "id2_string" }}","name":"Secret Rotation","sourceCheck":{"id":"{{ template "id3_string" }}","name":"Secret Rotation"}},
        {"id":"{{ template "id3_string" }}","name":"Dependency Scanning","sourceCheck":null}
    ],"pageInfo":{"hasNextPage":false,"hasPreviousPage":false,"startCursor":null,"endCursor":null}}}}}
}{{ end }}

{{- define "campaign_list_checks_empty_response" }}{
    "data":{"account":{"campaign":{"checks":{"nodes":[],"pageInfo":{"hasNextPage":false,"hasPreviousPage":false,"startCursor":null,"endCursor":null}}}}}
}{{ end }}

{{- define "campaign_copy_checks_request" }}
mutation ChecksCopyToCampaign($input:ChecksCopyToCampaignInput!){checksCopyToCampaign(input: $input){campaign{id,name,checkStats{total,totalSuccessful},endedDate,filter{id,name},htmlUrl,owner{alias,id},projectBrief,rawProjectBrief,reminder{channels,daysOfWeek,defaultSlackChannel,frequency,frequencyUnit,message,nextOccurrence,timeOfDay,timezone},serviceStats{total,totalSuccessful},startDate,status,targetDate},createdChecks{campaign{id,name},category{description,id,name},description,enableOn,enabled,filter{id,name,connective,htmlUrl,predicates{caseSensitive,key,keyData,type,value}},id,level{alias,checks{id,name},description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},sourceCheck{id,name},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CodeIssueCheck{constraint,issueName,issueType,maxAllowed,resolutionTime{unit,value},severity},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{frequencyTimeScale,frequencyValue,startingDate},updateRequiresComment},... on RepositoryFileCheck{directorySearch,fileContentsPredicate{type,value},filePaths,useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,fileContentsPredicate{type,value},filePaths},... on RepositorySearchCheck{fileContentsPredicate{type,value},fileExtensions},... on ServiceOwnershipCheck{contactMethod,requireContactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyDefinition{aliases,allowedInConfigFiles,id,name,description,displaySubtype,displayType,propertyDisplayStatus,lockedStatus,schema},propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{environmentPredicate{type,value},toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value}},... on HasDocumentationCheck{documentSubtype,documentType},... on PackageVersionCheck{missingPackageResult,packageConstraint,packageManager,packageName,packageNameIsRegex,versionConstraintPredicate{type,value}},... on RelationshipCheck{relationshipCountPredicate{type,value},relationshipDefinition{alias,componentType{id,aliases},description,id,managementRules{operator,sourceProperty,sourcePropertyBuiltin,targetCategory,targetProperty,targetPropertyBuiltin,targetType},metadata{allowedCategories,allowedTypes,maxItems,minItems},name}}},errors{message,path}}}
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
    },"createdChecks":[{ {{ template "common_check_response" }} }],"errors":[]}}
}{{ end }}
