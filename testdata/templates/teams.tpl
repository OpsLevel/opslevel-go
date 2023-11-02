{{- define "teamId_1" }}
{
"alias": "{{ template "alias1" }}",
{{ template "id1" }}
}
{{ end }}
{{- define "teamId_2" }}
{
"alias": "{{ template "alias2" }}",
{{ template "id2" }}
}
{{ end }}
{{- define "teamId_3" }}
{
"alias": "{{ template "alias3" }}",
{{ template "id3" }}
}
{{ end }}



{{- define "team_1" }}
{
"alias": "{{ template "alias1" }}",
{{ template "id1" }},
"aliases": [
  "{{ template "alias1" }}"
],
"contacts": [
  {{ template "contact_1" }}
],
"htmlUrl": "https://app.opslevel.com/teams/bots",
"manager": {{ template "user_1" }},
"memberships": {
  "nodes": [
    {
      "team": {{ template "teamId_2" }},
      "user": {{ template "user_2" }},
      "role": "admin"
    },
    {
      "team": {{ template "teamId_3" }},
      "user": {{ template "user_3" }},
      "role": "admin"
    }
  ],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 2
},
"name": "{{ template "name1" }}",
"responsibilities": null,
"tags": {
  "nodes": [],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 0
}
}
{{ end }}
{{- define "team_2" }}
{
"alias": "{{ template "alias2" }}",
{{ template "id2" }},
"aliases": [
  "{{ template "alias2" }}"
],
"contacts": [
  {{ template "contact_2" }}
],
"htmlUrl": "https://app.opslevel.com/teams/bots",
"manager": {{ template "user_1" }},
"memberships": {
  "nodes": [
    {
      "team": {{ template "teamId_2" }},
      "user": {{ template "user_2" }},
      "role": "admin"
    },
    {
      "team": {{ template "teamId_3" }},
      "user": {{ template "user_3" }},
      "role": "admin"
    }
  ],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 2
},
"name": "{{ template "name2" }}",
"responsibilities": null,
"tags": {
  "nodes": [],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 0
}
}
{{ end }}
{{- define "team_3" }}
{
"alias": "{{ template "alias3" }}",
{{ template "id3" }},
"aliases": [
  "{{ template "alias3" }}"
],
"contacts": [
  {{ template "contact_3" }}
],
"htmlUrl": "https://app.opslevel.com/teams/bots",
"manager": {{ template "user_1" }},
"memberships": {
  "nodes": [
    {
      "team": {{ template "teamId_2" }},
      "user": {{ template "user_2" }},
      "role": "admin"
    },
    {
      "team": {{ template "teamId_3" }},
      "user": {{ template "user_3" }},
      "role": "admin"
    }
  ],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 2
},
"name": "{{ template "name3" }}",
"responsibilities": null,
"tags": {
  "nodes": [],
  "pageInfo": {{ template "next_page_false" }},
  "totalCount": 0
}
}
{{ end }}
