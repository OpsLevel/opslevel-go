{{- define "scorecard_create_request" }}
mutation ScorecardCreate($input:ScorecardInput!){scorecardCreate(input: $input){scorecard{id,aliases,affectsOverallServiceLevels,description,filter{id,name,connective,htmlUrl,predicates{caseSensitive,key,keyData,type,value}},href,name,owner{... on Team{teamAlias:alias,id}},passingChecks,passingChecksPercentage,serviceCount,servicesReport{levelCounts{level{alias,checks{id,name},description,id,index,name},serviceCount}},slug,totalChecks},errors{message,path}}}
{{ end }}

{{- define "scorecard_create_request_vars" }}
{"input":{"affectsOverallServiceLevels":true,"description":"a new scorecard with an attached filter id","name":"new scorecard","ownerId":"Z2lkOi8vMTIzNDU2Nzg5Cg==","filterId":"Z2lkOi8vMTIzNDU2MTIzCg=="}}
{{ end }}

{{- define "scorecard_create_request_vars_affects_service_levels_false" }}
{"input":{"affectsOverallServiceLevels":false,"description":"a new scorecard with an attached filter id","name":"new scorecard","ownerId":"Z2lkOi8vMTIzNDU2Nzg5Cg==","filterId":"Z2lkOi8vMTIzNDU2MTIzCg=="}}
{{ end }}

{{- define "scorecard_create_response" }}{
    "data":{"scorecardCreate":{"scorecard":{"affectsOverallServiceLevels":true,"description":"a new scorecard with an attached filter id","filter":{"connective":null,"htmlUrl":"https://app.opslevel.com/filters/123456123","id":"Z2lkOi8vMTIzNDU2MTIzCg==","name":"some filter","predicates":[]},"name":"new scorecard","owner":{"id":"Z2lkOi8vMTIzNDU2Nzg5Cg=="}},"errors":[]}}
}{{ end }}

{{- define "scorecard_create_response_affects_service_levels_false" }}{
    "data":{"scorecardCreate":{"scorecard":{"affectsOverallServiceLevels":false,"description":"a new scorecard with an attached filter id","filter":{"connective":null,"htmlUrl":"https://app.opslevel.com/filters/123456123","id":"Z2lkOi8vMTIzNDU2MTIzCg==","name":"some filter","predicates":[]},"name":"new scorecard","owner":{"id":"Z2lkOi8vMTIzNDU2Nzg5Cg=="}},"errors":[]}}
}{{ end }}

{{- define "scorecard_update_request" }}
mutation ScorecardUpdate($input:ScorecardInput!$scorecard:IdentifierInput!){scorecardUpdate(scorecard: $scorecard, input: $input){scorecard{id,aliases,affectsOverallServiceLevels,description,filter{id,name,connective,htmlUrl,predicates{caseSensitive,key,keyData,type,value}},href,name,owner{... on Team{teamAlias:alias,id}},passingChecks,passingChecksPercentage,serviceCount,servicesReport{levelCounts{level{alias,checks{id,name},description,id,index,name},serviceCount}},slug,totalChecks},errors{message,path}}}
{{ end }}

{{- define "scorecard_update_request_vars" }}
{"input":{"description":"this scorecard was updated","filterId":"Z2lkOi8vMTIzNDU2NDU2Cg==","name":"updated scorecard","ownerId":"Z2lkOi8vMTIzNDU2Nzc3Cg=="},"scorecard":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK"}}
{{ end }}

{{- define "scorecard_update_response" }}{
    "data":{"scorecardUpdate":{"scorecard":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK","description":"this scorecard was updated","filter":{"connective":null,"htmlUrl":"https://app.opslevel.com/filters/123456456","id":"Z2lkOi8vMTIzNDU2NDU2Cg==","name":"some new filter"},"name":"updated scorecard","owner":{"id":"Z2lkOi8vMTIzNDU2Nzc3Cg=="}},"errors":[]}}
}{{ end }}

{{- define "scorecard_delete_request" }}
mutation ScorecardDelete($input:IdentifierInput!){scorecardDelete(input: $input){deletedScorecardId,errors{message,path}}}
{{ end }}

{{- define "scorecard_delete_request_vars" }}
{"input":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK"}}
{{ end }}

{{- define "scorecard_delete_response" }}{
    "data":{"scorecardDelete":{"deletedScorecardId":"Z2lkOi8vMTIzNDU2Nzg5MTAK","errors":[]}}
}{{ end }}

{{- define "scorecard_get_request" }}
query ScorecardGet($input:IdentifierInput!){account{scorecard(input: $input){id,aliases,affectsOverallServiceLevels,description,filter{id,name,connective,htmlUrl,predicates{caseSensitive,key,keyData,type,value}},href,name,owner{... on Team{teamAlias:alias,id}},passingChecks,passingChecksPercentage,serviceCount,servicesReport{levelCounts{level{alias,checks{id,name},description,id,index,name},serviceCount}},slug,totalChecks}}}
{{ end }}

{{- define "scorecard_get_request_vars" }}
{"input":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK"}}
{{ end }}

{{- define "scorecard_get_response" }}{
    "data":{"account":{"scorecard":{"aliases":["existing_scorecard"],"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK","description":"hello there!","filter":{"connective":null,"htmlUrl":"https://app.opslevel.com/filters/123456123","id":"Z2lkOi8vMTIzNDU2MTIzCg==","name":"some filter","predicates":[]},"name":"fetched scorecard","owner":{"id":"Z2lkOi8vMTIzNDU2Nzg5Cg=="},"passingChecks":10,"serviceCount":20,"totalChecks":30}}}
}{{ end }}

{{- define "scorecard_list_query" }}query ScorecardsList($after:String!$first:Int!){account{scorecards(after: $after, first: $first){nodes{id,aliases,affectsOverallServiceLevels,description,filter{id,name,connective,htmlUrl,predicates{caseSensitive,key,keyData,type,value}},href,name,owner{... on Team{teamAlias:alias,id}},passingChecks,passingChecksPercentage,serviceCount,servicesReport{levelCounts{level{alias,checks{id,name},description,id,index,name},serviceCount}},slug,totalChecks},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}{{ end }}

{{- define "scorecard_1_response" }}
    "id":"Z2lkOi8vMTExMTExMTEK",
    "aliases":["first_scorecard"],
    "name":"first scorecard",
    "description":"the first scorecard",
    "filter":{"connective":"and","htmlUrl":"https://app.opslevel.com/filters/123456123","id":"Z2lkOi8vMTIzNDU2MTIzCg==","name":"first filter","predicates":[]},
    "owner":{"id":"Z2lkOi8vMTIzNDU2Nzg5Cg=="},
    "passingChecks":92,
    "serviceCount":150,
    "totalChecks":100
{{ end }}

{{- define "scorecard_2_response" }}
    "id":"Z2lkOi8vMjIyMjIyMjIK",
    "aliases":["second_scorecard"],
    "name":"second scorecard",
    "description":"the second scorecard",
    "filter":{"connective":"or","htmlUrl":"https://app.opslevel.com/filters/123456456","id":"Z2lkOi8vMTIzNDU2NDU2Cg==","name":"second filter","predicates":[]},
    "owner":{"id":"Z2lkOi8vMTIzNDU2Nzg5Cg=="},
    "passingChecks":20,
    "serviceCount":22,
    "totalChecks":27
{{ end }}

{{- define "scorecard_3_response" }}
    "id":"Z2lkOi8vMzMzMzMzMzMK",
    "aliases":["third_scorecard"],
    "filter":{},
    "name":"third scorecard",
    "description":"the third scorecard",
    "owner":{"id":"Z2lkOi8vMTIzNDU2Nzc3Cg=="},
    "passingChecks":30,
    "serviceCount":33,
    "totalChecks":33
{{ end }}
