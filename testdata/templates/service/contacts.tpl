{{- define "contact_1" }}
{
"address": "{{ template "email1" }}",
"displayName": "Email",
{{ template "id1" }},
"type": "email"
}
{{ end }}
{{- define "contact_2" }}
{
"address": "{{ template "email2" }}",
"displayName": "Email",
{{ template "id2" }},
"type": "email"
}
{{ end }}
{{- define "contact_3" }}
{
"address": "{{ template "email3" }}",
"displayName": "Email",
{{ template "id3" }},
"type": "email"
}
{{ end }}
