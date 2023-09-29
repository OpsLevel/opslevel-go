{{- define "teamId_1" }}
{
"alias": "{{ template "alias1" }}",
"id": "{{ template "id1" }}"
}
{{ end }}
{{- define "teamId_2" }}
{
"alias": "{{ template "alias2" }}",
"id": "{{ template "id2" }}"
}
{{ end }}
{{- define "teamId_3" }}
{
"alias": "{{ template "alias3" }}",
"id": "{{ template "id3" }}"
}
{{ end }}



{{- define "team_1" }}
{
"alias": "{{ template "alias1" }}",
"id": "{{ template "id1" }}",
"aliases": [
  "{{ template "alias1" }}"
],
"contacts": [
  {{ template "contact_1" }}
],
"group": null,
"htmlUrl": "https://app.opslevel.com/teams/bots",
"manager": {{ template "user_1" }},
"members": {
  "nodes": [
    {{ template "user_2" }},
    {{ template "user_3" }}
  ],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 2
},
"name": "{{ template "name1" }}",
"responsibilities": null
}
{{ end }}
{{- define "team_2" }}
{
"alias": "{{ template "alias2" }}",
"id": "{{ template "id2" }}",
"aliases": [
  "{{ template "alias2" }}"
],
"contacts": [
  {{ template "contact_2" }}
],
"group": null,
"htmlUrl": "https://app.opslevel.com/teams/bots",
"manager": {{ template "user_1" }},
"members": {
  "nodes": [
    {{ template "user_2" }},
    {{ template "user_3" }}
  ],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 2
},
"name": "{{ template "name2" }}",
"responsibilities": null
}
{{ end }}
{{- define "team_3" }}
{
"alias": "{{ template "alias3" }}",
"id": "{{ template "id3" }}",
"aliases": [
  "{{ template "alias3" }}"
],
"contacts": [
  {{ template "contact_3" }}
],
"group": null,
"htmlUrl": "https://app.opslevel.com/teams/bots",
"manager": {{ template "user_1" }},
"members": {
  "nodes": [
    {{ template "user_2" }},
    {{ template "user_3" }}
  ],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 2
},
"name": "{{ template "name3" }}",
"responsibilities": null
}
{{ end }}
