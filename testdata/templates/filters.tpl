{{- define "filter_1" }}
{
    "connective": null,
    "htmlUrl": "https://app.opslevel.com/filters/1",
    {{ template "id1" }},
    "name": "{{ template "name1" }}",
    "predicates": []
}
{{ end }}

{{- define "filter_2" }}
{
    "connective": null,
    "htmlUrl": "https://app.opslevel.com/filters/2",
    {{ template "id2" }},
    "name": "{{ template "name2" }}",
    "predicates": [
        {
        "key": "tier_index",
        "keyData": null,
        "type": "equals",
        "value": "1"
        }
    ]
}
{{ end }}

{{- define "filter_3" }}
{
    "connective": null,
    "htmlUrl": "https://app.opslevel.com/filters/3",
    {{ template "id3" }},
    "name": "{{ template "name3" }}",
    "predicates": []
}
{{ end }}

{{- define "filter_kubernetes_response" }}
  "connective": null,
  "htmlUrl": "https://app.opslevel.com/filters/458",
  "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzQ1OA",
  "name": "Kubernetes",
  "predicates": []
{{ end }}

{{- define "filter_tier1service_response" }}
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
{{ end }}

{{- define "filter_complex_kubernetes_response" }}
  "connective": null,
  "htmlUrl": "https://app.opslevel.com/filters/452",
  "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tsaXN0LzQ1Mg",
  "name": "Kubernetes",
  "predicates": [
    {
      "key": "tier_index",
      "keyData": null,
      "type": "equals",
      "value": "1"
    }
  ]
{{ end }}

{{- define "create_filter_nested_input" }}
"name": "Self deployed or Rails",
"predicates": [
  {
    "key": "filter_id",
    "type": "matches",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg"
  },
  {
    "key": "filter_id",
    "type": "matches",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjQ"
  }
],
"connective": "or"
{{ end }}

{{- define "create_filter_nested_response" }}
"connective": "or",
"htmlUrl": "https://app.opslevel.com/filters/2346",
"id": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzIzNDY",
"name": "Self deployed or Rails",
"predicates": [
  {
    "key": "filter_id",
    "keyData": null,
    "type": "matches",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg"
  },
  {
    "key": "filter_id",
    "keyData": null,
    "type": "matches",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjQ"
  }
]
{{ end }}
{{- define "update_filter_nested_input" }}
"id": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzIzNDY",
"name": "Tier 1-2 not deployed by us",
"predicates": [
  {
    "key": "filter_id",
    "type": "does_not_match",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg"
  },
  {
    "key": "filter_id",
    "type": "matches",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjY"
  }
],
"connective": "and"
{{ end }}

{{- define "update_filter_nested_response" }}
"connective": "and",
"htmlUrl": "https://app.opslevel.com/filters/2346",
"id": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzIzNDY",
"name": "Tier 1-2 not deployed by us",
"predicates": [
  {
    "key": "filter_id",
    "keyData": null,
    "type": "does_not_match",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNTg"
  },
  {
    "key": "filter_id",
    "keyData": null,
    "type": "matches",
    "value": "Z2lkOi8vb3BzbGV2ZWwvRmlsdGVyLzEyNjY"
  }
]
{{ end }}
