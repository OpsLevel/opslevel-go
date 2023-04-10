{{- define "serviceDependency" }}
{
    "id": "{{ template "id1" }}",
    "sourceService": {
        "id": "{{ template "id2" }}",
        "aliases": [
            "{{ template "alias2" }}"
        ]
    },
    "destinationService": {
        "id": "{{ template "id3" }}",
        "aliases": [
            "{{ template "alias3" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
{{- define "serviceDependencyEdge_1" }}
{
    "id": "{{ template "id1" }}",
    "locked": false,
    "node": {
        "id": "{{ template "id1" }}",
        "aliases": [
            "{{ template "alias1" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
{{- define "serviceDependencyEdge_2" }}
{
    "id": "{{ template "id2" }}",
    "locked": false,
    "node": {
        "id": "{{ template "id2" }}",
        "aliases": [
            "{{ template "alias2" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
{{- define "serviceDependencyEdge_3" }}
{
    "id": "{{ template "id3" }}",
    "locked": false,
    "node": {
        "id": "{{ template "id3" }}",
        "aliases": [
            "{{ template "alias3" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}