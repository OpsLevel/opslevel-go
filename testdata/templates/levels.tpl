{{- define "level_1" }}
{
    "alias": "{{ template "alias1" }}",
    "description": "{{ template "description" }}",
    "id": "{{ template "id1" }}",
    "index": 1,
    "name": "{{ template "name1" }}"
}
{{ end }}

{{- define "level_2" }}
{
    "alias": "{{ template "alias2" }}",
    "description": "{{ template "description" }}",
    "id": "{{ template "id2" }}",
    "index": 2,
    "name": "{{ template "name2" }}"
}
{{ end }}

{{- define "level_3" }}
{
    "alias": "{{ template "alias3" }}",
    "description": "{{ template "description" }}",
    "id": "{{ template "id3" }}",
    "index": 3,
    "name": "{{ template "name3" }}"
}
{{ end }}