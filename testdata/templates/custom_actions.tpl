{{- define "custom_actions_request" }}{aliases,id,description,liquidTemplate,name,... on CustomActionsWebhookAction{headers,httpMethod,webhookUrl}}{{ end }}
{{- define "custom_actions_trigger_request" }}{action{aliases,id},aliases,description,filter{id,name},id,name,owner{alias,id},timestamps{createdAt,updatedAt}}{{ end }}
{{- define "custom_action1" }}{
    "aliases": [],
    "description": null,
    "id": "123456789",
    "liquidTemplate": "{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}",
    "name": "Deploy Rollback",
    "headers": {
        "Content-Type": "application/json"
    },
    "httpMethod": "POST",
    "webhookUrl": "https://gitlab.com/api/v4/projects/1/trigger/pipeline"
}{{ end }}
{{- define "custom_action2" }}{
    "aliases": [],
    "description": "Trigger a deploy freeze",
    "id": "987654321",
    "liquidTemplate": "{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"freeze\"}",
    "name": "Deploy Freeze",
    "headers": {
        "Accept": "application/vnd.github+json",
        "Content-Type": "application/json"
    },
    "httpMethod": "PUT",
    "webhookUrl": "https://gitlab.com/api/v4/projects/1/trigger/pipeline"
}{{ end }}
{{- define "custom_action_trigger1" }}{
    "action": {
      "aliases": [],
      "id": "987654321"
    },
    "aliases": [],
    "description": "Disables the Deploy Freeze",
    "filter": {
      "id": "987654321",
      "name": "Uses Ruby and Deploys Frozen"
    },
    "id": "123456789",
    "name": "Release",
    "owner": {
      "alias": "platform",
      "id": "123456789"
    },
    "timestamps": {
      "createdAt": "2022-12-15T01:34:00.289687Z",
      "updatedAt": "2022-12-15T15:01:08.832770Z"
    }
}{{ end }}
{{- define "custom_action_trigger2" }}{
    "action": {
      "aliases": [],
      "id": "123456789"
    },
    "aliases": [],
    "description": "Cuts a new Release",
    "filter": {
      "id": "123456789",
      "name": "Uses Ruby"
    },
    "id": "987654321",
    "name": "Release",
    "owner": {
      "alias": "platform",
      "id": "123456789"
    },
    "timestamps": {
      "createdAt": "2022-12-15T01:34:00.289687Z",
      "updatedAt": "2022-12-15T15:01:08.832770Z"
    }
}{{ end }}