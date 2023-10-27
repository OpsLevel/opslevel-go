{{- define "secrets_1" }}
{
    {{ template "id1" }},
    "alias": "{{ template "alias1" }}",
    "timestamps": {{ template "timestamps" }}
}
{{end}}
{{- define "secrets_2" }}
{
    {{ template "id2" }},
    "alias": "{{ template "alias2" }}",
    "owner": {{ template "teamId_2" }},
    "timestamps": {{ template "timestamps" }}
}
{{end}}
{{- define "secrets_3" }}
{
    {{ template "id3" }},
    "alias": "{{ template "alias3" }}",
    "owner": {{ template "teamId_3" }},
    "timestamps": {{ template "timestamps" }}
}
{{end}}
{{- define "secret_input_1" }}
{
  "value": "secret_value_1"
  "owner": {
    {{ template "id1" }}
  }
}
{{end}}
{{- define "secret_create_vars" }}
{
  "alias": "alias1",
  "input": {
    "value": "my-secret",
    "owner": {
      {{ template "id2" }}
    }
  }
}
{{end}}
{{- define "secret_delete_vars" }}
{
  "input": {
    {{ template "id1" }}
  }
}
{{end}}
{{- define "secret_get_vars" }}
{
  "input":{
    {{ template "id2" }}
  }
}

{{end}}
{{- define "secret_update_vars" }}
{
  "input": {
    "value": "secret_value_2",
    "owner": {
      {{ template "id2" }}
    }
  },
  "secret": {
    {{ template "id2" }}
  }
}
{{end}}
{{- define "secret_create_response" }}
{
  "data": {
    "secretsVaultsSecretCreate": {
      "secret": {{ template "secrets_2" }},
      "errors": []
    }
  }
}
{{end}}
{{- define "secret_delete_response" }}
{
  "data": {
		"secretsVaultsSecretDelete": {
			"errors": []
		}
  }
}
{{end}}
{{- define "secret_get_response" }}
{
  "data": {
    "account": {
      "secretsVaultsSecret": {{ template "secrets_2" }}
    }
  }
}
{{end}}
{{- define "secret_list_response_1" }}
{
  "data": {
    "account": {
      "secretsVaultsSecrets": {
        "nodes": [
            {{ template "secrets_1" }},
            {{ template "secrets_2" }}
        ],
        {{ template "pagination_initial_pageInfo_response" }}
      }
    }
  }
}
{{end}}
{{- define "secret_list_response_2" }}
{
  "data": {
    "account": {
      "secretsVaultsSecrets": {
        "nodes": [
          {{ template "secrets_3" }}
        ],
        {{ template "pagination_second_pageInfo_response" }}
        }
    }
  }
}
{{end}}
{{- define "secret_update_response" }}
{
  "data": {
    "secretsVaultsSecretUpdate": {
	    "secret": {{ template "secrets_2" }},
      "errors": []
    }
  }
}
{{end}}
