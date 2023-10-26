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
    "level": {
      "alias": "bronze",
      "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
      "index": 1,
      "name": "Bronze"
    },
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
    "filter": {
      "connective": null,
      "htmlUrl": "https://app.opslevel.com/filters/401",
      "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzQwMQ",
      "name": "Tier 1 Services",
      "predicates": [
        {
          "key": "tier_index",
          "keyData": null,
          "type": "equals",
          "value": "1"
        }
      ]
    },
    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpUb29sVXNhZ2UvMTMyNw",
    "level": {
      "alias": "bronze",
      "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
      "index": 1,
      "name": "Bronze"
    },
    "name": "Metrics Tool",
    "notes": null
{{ end }}
{{- define "owner_defined_check" }}
    "category": null,
	"description": "Verifies that the service has an owner defined.",
    "enabled": true,
    "filter": null,
    "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8xMzI9",
    "level": {
      "alias": "bronze",
      "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
      "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
      "index": 1,
      "name": "Bronze"
    },
    "name": "Owner Defined",
    "notes": null
{{ end }}
