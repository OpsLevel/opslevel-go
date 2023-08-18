{{- define "scorecard_create_request" }}{
    "query": "mutation ScorecardCreate($input:ScorecardInput!){scorecardCreate(input: $input){scorecard{aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks},errors{message,path}}}",
    "variables": {"input":{"name":"new scorecard","ownerId":"Z2lkOi8vMTIzNDU2Nzg5Cg=="}}
}{{ end }}

{{- define "scorecard_create_response" }}{
    "data": {"scorecardCreate":{"scorecard":{"name":"new scorecard","owner":{"id": "Z2lkOi8vMTIzNDU2Nzg5Cg=="}},"errors":[]}}
}{{ end }}

{{- define "scorecard_update_request" }}{
    "query": "mutation ScorecardUpdate($input:ScorecardInput!$scorecard:IdentifierInput!){scorecardUpdate(scorecard: $scorecard, input: $input){scorecard{aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks},errors{message,path}}}",
    "variables": {"input":{"name":"updated scorecard","ownerId":"Z2lkOi8vMTIzNDU2Nzg5Cg=="},"scorecard":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK"}}
}{{ end }}

{{- define "scorecard_update_response" }}{
    "data": {"scorecardUpdate":{"scorecard":{"name":"updated scorecard","owner":{"id": "Z2lkOi8vMTIzNDU2Nzg5Cg=="}},"errors":[]}}
}{{ end }}

{{- define "scorecard_delete_request"}}{
    "query": "mutation ScorecardDelete($input:IdentifierInput!){scorecardDelete(input: $input){deletedScorecardId,errors{message,path}}}",
    "variables": {"input":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK"}}
}{{ end }}

{{- define "scorecard_delete_response" }}{
    "data": {"scorecardDelete":{"deletedScorecardId":"Z2lkOi8vMTIzNDU2Nzg5MTAK","errors":[]}}
}{{ end }}

{{- define "scorecard_get_request" }}{
    "query": "query ScorecardGet($input:IdentifierInput!){account{scorecard(input: $input){aliases,id,description,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value}},name,owner{... on Group{groupAlias:alias,id},... on Team{teamAlias:alias,id}},passingChecks,serviceCount,totalChecks}}}",
    "variables": {"input":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK"}}
}{{ end }}

{{- define "scorecard_get_response" }}{
    "data": {"account":{"scorecard":{"id":"Z2lkOi8vMTIzNDU2Nzg5MTAK", "name":"existing scorecard","owner":{"id":"Z2lkOi8vMTIzNDU2Nzg5Cg=="}}}}
}{{ end }}
