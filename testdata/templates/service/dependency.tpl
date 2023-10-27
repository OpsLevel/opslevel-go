{{- define "serviceDependency" }}
{
    {{ template "id1" }},
    "sourceService": {
        {{ template "id2" }},
        "aliases": [
            "{{ template "alias2" }}"
        ]
    },
    "destinationService": {
        {{ template "id3" }},
        "aliases": [
            "{{ template "alias3" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
{{- define "serviceDependencyEdge_1" }}
{
    {{ template "id1" }},
    "locked": false,
    "node": {
        {{ template "id1" }},
        "aliases": [
            "{{ template "alias1" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
{{- define "serviceDependencyEdge_2" }}
{
    {{ template "id2" }},
    "locked": false,
    "node": {
        {{ template "id2" }},
        "aliases": [
            "{{ template "alias2" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
{{- define "serviceDependencyEdge_3" }}
{
    {{ template "id3" }},
    "locked": false,
    "node": {
        {{ template "id3" }},
        "aliases": [
            "{{ template "alias3" }}"
        ]
    },
    "notes": "{{ template "description" }}"
}
{{ end }}
