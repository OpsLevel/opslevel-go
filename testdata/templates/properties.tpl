{{- define "property_assign_input" }}{"owner":{"alias":"monolith"},"definition":{"alias":"is_beta_feature"},"value":"true"}{{ end }}

{{- define "service_property_edge_1" }}
{
    "cursor": "",
    "node": {
        "definition": {
            "alias": "propdef1"
        },
        "owner": {
            "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
        },
        "validationErrors": [],
        "value": "true"
    }
},
{
    "cursor": "",
    "node": {
        "definition": {
            "alias": "propdef2"
        },
        "owner": {
            "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
        },
        "validationErrors": [],
        "value": "false"
    }
}
{{ end }}

{{- define "service_property_edge_3" }}
{
    "cursor": "",
    "node": {
        "definition": {
            "alias": "propdef3"
        },
        "owner": {
            "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx"
        },
        "validationErrors": [],
        "value": "\"Hello World!\""
    }
}
{{ end }}
