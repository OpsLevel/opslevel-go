{{- define "property_assign_input" }}{"owner":{"id":"{{ template "id1_string" }}"},"definition":{"id":"{{ template "id2_string" }}"},"value":"true"}{{ end }}

{{- define "service_properties_page_1" }}
{
    "definition": {
        "id": "{{ template "id2_string" }}"
    },
    "locked": true,
    "owner": {
        "__typename": "Service",
        "id": "{{ template "id1_string" }}",
        "aliases": []
    },
    "validationErrors": [],
    "value": "true"
},
{
    "definition": {
        "id": "{{ template "id3_string" }}"
    },
    "locked": false,
    "owner": {
        "__typename": "Service",
        "id": "{{ template "id1_string" }}",
        "aliases": []
    },
    "validationErrors": [],
    "value": "false"
}
{{ end }}

{{- define "service_properties_page_2" }}
{
    "definition": {
        "id": "{{ template "id4_string" }}"
    },
    "locked": true,
    "owner": {
        "__typename": "Service",
        "id": "{{ template "id1_string" }}",
        "aliases": []
    },
    "validationErrors": [],
    "value": "\"Hello World!\""
}
{{ end }}
