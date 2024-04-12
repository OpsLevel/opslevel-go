{{- define "contact_1" }}
{
"address": "{{ template "email1" }}",
"displayName": "Email",
"displayType": "Primary",
"externalId": null,
{{ template "id1" }},
"isDefault": true,
"type": "email"
}
{{ end }}
{{- define "contact_2" }}
{
"address": "#engineering",
"displayName": "Slack Channel",
"displayType": "Secondary",
"externalId": "CN689A229",
{{ template "id2" }},
"isDefault": true,
"type": "slack"
}
{{ end }}
{{- define "contact_3" }}
{
"address": "#marketing",
"displayName": "Slack",
"displayType": null,
"externalId": null,
{{ template "id3" }},
"isDefault": false,
"type": "slack"
}
{{ end }}
