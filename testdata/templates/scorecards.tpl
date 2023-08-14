{{- define "scorecard_create_request" }}{
    "query": "mutation ScorecardCreate($input:ScorecardInput!){scorecardCreate(input: $input){scorecard{aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks},errors{message,path}}}",
    "variables": {"input":{"name":"new scorecard","owner":{"alias":"platform"}}}
}{{ end }}

{{- define "scorecard_create_response" }}{
    "data": {"scorecardCreate":{"scorecard":{"name":"new scorecard","owner":{"teamAlias":"platform"}},"errors":[]}}
}{{ end }}

{{- define "scorecard_update_request" }}{
    "query": "mutation ScorecardUpdate($input:ScorecardInput!$scorecard:IdentifierInput!){scorecardUpdate(scorecard: $scorecard, input: $input){scorecard{aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks},errors{message,path}}}",
    "variables": {"input":{"name":"updated scorecard","owner":{"id":"team-id"}},"scorecard":{"id":"scorecard-id"}}
}{{ end }}

{{- define "scorecard_update_response" }}{
    "data": {"scorecardUpdate":{"scorecard":{"name":"updated scorecard","owner":{"id":"team-id"}},"errors":[]}}
}{{ end }}

{{- define "scorecard_delete_request"}}{
    "query": "mutation ScorecardDelete($input:IdentifierInput!){scorecardDelete(input: $input){deletedScorecardId,errors{message,path}}}",
    "variables": {"input":{"alias":"scorecard-alias"}}
}{{ end }}

{{- define "scorecard_delete_response" }}{
    "data": {"scorecardDelete":{"deletedScorecardId":"scorecard-id","errors":[]}}
}{{ end }}

{{- define "scorecard_get_request" }}{
    "query": "query ScorecardGet($input:IdentifierInput!){account{scorecard(input: $input){aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks}}}",
    "variables": {"input":{"id":"scorecard-id"}}
}{{ end }}

{{- define "scorecard_get_response" }}{
    "data": {"account":{"scorecard":{"id":"scorecard-id", "name":"existing scorecard","owner":{"teamAlias":"platform"}}}}
}{{ end }}
