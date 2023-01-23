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