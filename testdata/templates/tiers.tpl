{{- define "tier_1" }}
{
  "alias": "{{ template "alias1" }}",
  "description": "Mission-critical service or repository. Failure could result in significant impact to revenue or reputation.",
  {{ template "id1" }},
  "index": 1,
  "name": "Tier 1"
}
{{ end }}
{{- define "tier_2" }}
{
  "alias": "{{ template "alias2" }}",
  "description": "Customer-facing service or repository. Failure results in degraded experience for customers, although without significant impact to revenue or reputation.",
  {{ template "id2" }},
  "index": 2,
  "name": "Tier 2"
}
{{ end }}
{{- define "tier_3" }}
{
  "alias": "{{ template "alias3" }}",
  "description": "Internal service or repository. Failure could result in productivity being compromised within the company.",
  {{ template "id3" }},
  "index": 3,
  "name": "Tier 3"
}
{{ end }}
{{- define "tier_4" }}
{
  "alias": "{{ template "alias4" }}",
  "description": "Other service or repository. Failure doesn't result in immediate or significant impact.",
  {{ template "id4" }},
  "index": 4,
  "name": "Tier 4"
}
{{ end }}
