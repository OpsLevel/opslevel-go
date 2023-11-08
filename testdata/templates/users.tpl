{{- define "user_id_email_1" }}
{{ template "id1" }},
"email": "kyle@opslevel.com"
{{ end }}
{{- define "user_id_email_2" }}
{{ template "id2" }},
"email": "edgar@opslevel.com"
{{ end }}

{{- define "user_id_email_3" }}
{{ template "id3" }},
"email": "ken@opslevel.com"
{{ end }}

{{- define "user_1" }}
{
{{ template "user_id_email_1" }},
"name": "Kyle Rockman",
"role": "user"
}
{{ end }}
{{- define "user_1_update" }}
{
{{ template "user_id_email_1" }},
"name": "Kyle Rockman",
"role": "admin"
}
{{ end }}
{{- define "user_2" }}
{
{{ template "user_id_email_2" }},
"name": "Edgar Ochoa",
"role": "admin"
}
{{ end }}
{{- define "user_3" }}
{
"email": "matthew@opslevel.com",
"id": "3",
"name": "Matthew Brahms",
"role": "admin"
}
{{ end }}
{{- define "user_4" }}
{
"name": "Andrew Example",
"email": "example@opslevel.com"
}
{{ end }}
{{- define "user_5" }}
{
"name": "Ashley Example",
"email": "example@opslevel.com"
}
{{ end }}
{{- define "user_6" }}
{
"name": "Drew Example",
"email": "example@opslevel.com"
}
{{ end }}
