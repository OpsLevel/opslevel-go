{{- define "liquid_template_freeze_string" -}}
{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"freeze\"}
{{- end }}

{{- define "liquid_template_rollback_string" -}}
{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}
{{- end }}

{{- define "liquid_template_freeze" -}}
"liquidTemplate": "{{- template "liquid_template_freeze_string" -}}"
{{ end }}

{{- define "liquid_template_rollback" }}
"liquidTemplate": "{{- template "liquid_template_rollback_string" -}}"
{{ end }}

{{- define "custom_action1_response" }}
    "aliases": [],
    "description": null,
    "id": "123456789",
    {{ template "liquid_template_rollback" }},
    "name": "Deploy Rollback",
    "headers": {
        "Content-Type": "application/json"
    },
    "httpMethod": "POST",
    "webhookUrl": "https://gitlab.com/api/v4/projects/1/trigger/pipeline"
{{ end }}
{{- define "custom_action2_response" }}
    "aliases": [],
    "description": "Trigger a deploy freeze",
    "id": "987654322",
    {{ template "liquid_template_freeze" }},
    "name": "Deploy Freeze",
    "headers": {
        "Accept": "application/vnd.github+json",
        "Content-Type": "application/json"
    },
    "httpMethod": "PUT",
    "webhookUrl": "https://gitlab.com/api/v4/projects/1/trigger/pipeline"
{{ end }}
{{- define "custom_action3_response" }}
    "aliases": [],
    "description": "Page the On-Call Engineer",
    "id": "987654323",
    {{ template "liquid_template_freeze" }},
    "name": "Page On-Call",
    "headers": {
        "Accept": "application/vnd.github+json",
        "Content-Type": "application/json"
    },
    "httpMethod": "PUT",
    "webhookUrl": "https://gitlab.com/api/v4/projects/1/trigger/pipeline"
{{ end }}
{{- define "custom_action_trigger1_response" }}
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
{{ end }}
{{- define "custom_action_trigger2_response" }}
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
{{ end }}
{{- define "custom_action_trigger3_response" }}
    "action": {
      "aliases": [],
      "id": "123456788"
    },
    "aliases": [],
    "description": "Rolls back to last good deploy",
    "filter": {
      "id": "123456788",
      "name": "Uses Go"
    },
    "id": "987654322",
    "name": "Rollback",
    "owner": {
      "alias": "platform",
      "id": "123456781"
    },
    "timestamps": {
      "createdAt": "2022-12-15T01:34:00.289687Z",
      "updatedAt": "2022-12-15T15:01:08.832770Z"
    }
{{ end }}
{{- define "custom_actions_request" }}{aliases,id,description,liquidTemplate,name,... on CustomActionsWebhookAction{headers,httpMethod,webhookUrl}}{{ end }}
{{- define "custom_actions_trigger_request" }}{action{aliases,id},aliases,description,filter{id,name},id,manualInputsDefinition,name,owner{alias,id},published,timestamps{createdAt,updatedAt},accessControl,responseTemplate,entityType}{{ end }}
{{- define "custom_action1" }}{
    "aliases": [],
    "description": null,
    "id": "123456789",
    {{ template "liquid_template_rollback" }},
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
    {{ template "liquid_template_freeze" }},
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
    "entityType": "SERVICE",
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
{{- define "extended_team_access_response_1" }}
{
  "data": {
    "account": {
      "customActionsTriggerDefinition": {
        "extendedTeamAccess": {
          "nodes": [ {{ template "team_1" }} ],
          {{ template "pagination_initial_pageInfo_response" }}
        }
      }
    }
  }
}
{{ end }}
{{- define "extended_team_access_response_2" }}
{
  "data": {
    "account": {
      "customActionsTriggerDefinition": {
        "extendedTeamAccess": {
          "nodes": [ {{ template "team_2" }} ],
          {{ template "pagination_second_pageInfo_response" }}
        }
      }
    }
  }
}
{{ end }}
{{- define "extended_team_access_get_vars_1" }}
{
  "input": { {{ template "id1" }} },
  {{ template "first_page_variables" }}
}
{{ end }}
{{- define "extended_team_access_get_vars_2" }}
{
  "input": { {{ template "id1" }} },
  {{ template "second_page_variables" }}
}
{{ end }}
