{{- define "tool_1" }}
{
    "category": "other",
    "categoryAlias": null,
    "displayName": "{{ template "name1" }}",
    "environment": null,
    "id": "{{ template "id1" }}",
    "service": {
        "id": "{{ template "id1" }}"
    },
    "url": "https://example.com"
}
{{ end }}

{{- define "tool_1_update" }}
{
    "category": "deployment",
    "categoryAlias": null,
    "displayName": "{{ template "name1" }}",
    "environment": "prod",
    "id": "{{ template "id1" }}",
    "service": {
        "id": "{{ template "id1" }}"
    },
    "url": "https://example.com"
}
{{ end }}