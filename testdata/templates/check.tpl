{{- define "check_base_vars" }}
        "name": "Hello World",
        "enabled": true,
        "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
        "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
        "notes": "Hello World Check"
{{ end }}
{{- define "common_check_response" }}
    "category": null,
    "description": "Verifies that the service has a repository integrated.",
    "enabled": true,
    "filter": null,
    {{ template "id1" }},
    "level": {{ template "level_1" }},
    "name": "Repository Integrated",
    "notes": null
{{ end }}
{{- define "manual_check_response" }}{
    {{ template "common_check_response" }},
    "updateFrequency": {
        "startingDate": "2021-07-26T20:22:44.427Z",
        "frequencyTimeScale": "week",
        "frequencyValue": 1
    },
}
{{ end }}
{{- define "metrics_tool_check" }}
    "category": null,
    "description": "The service has a metrics tool.",
    "enabled": true,
    "filter": { {{ template "filter_tier1service_response" }} },
    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpUb29sVXNhZ2UvMTMyNw",
    {{ template "level_bronze" }},
    "name": "Metrics Tool",
    "notes": null
{{ end }}
{{- define "owner_defined_check" }}
    "category": null,
    "description": "Verifies that the service has an owner defined.",
    "enabled": true,
    "filter": null,
    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8xMzI9",
    {{ template "level_bronze" }},
    "name": "Owner Defined",
    "notes": null
{{ end }}
