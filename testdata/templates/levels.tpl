{{- define "level_1" }}
{
    "alias": "{{ template "alias1" }}",
    "description": "{{ template "description" }}",
    {{ template "id1" }},
    "index": 1,
    "name": "{{ template "name1" }}"
}
{{ end }}

{{- define "level_2" }}
{
    "alias": "{{ template "alias2" }}",
    "description": "{{ template "description" }}",
    {{ template "id2" }},
    "index": 2,
    "name": "{{ template "name2" }}"
}
{{ end }}

{{- define "level_3" }}
{
    "alias": "{{ template "alias3" }}",
    "description": "{{ template "description" }}",
    {{ template "id3" }},
    "index": 3,
    "name": "{{ template "name3" }}"
}
{{ end }}

{{- define "level_bronze" }}
"level": {
  "alias": "bronze",
  "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
  "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
  "index": 1,
  "name": "Bronze"
}
{{ end }}
