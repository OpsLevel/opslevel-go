{{- define "category_1" }}
{
  {{ template "id1" }},
  "name": "{{ template "name1" }}"
}
{{ end }}

{{- define "category_2" }}
{
  {{ template "id2" }},
  "name": "{{ template "name2" }}"
}
{{ end }}

{{- define "category_3" }}
{
  {{ template "id3" }},
  "name": "{{ template "name3" }}"
}
{{ end }}

{{- define "rubric_categories_response1" }}
  "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjAz",
  "name": "🔐 Security"
{{ end }}
{{- define "rubric_categories_response2" }}
  "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA0",
  "name": "🟢 Reliability"
{{ end }}
{{- define "rubric_categories_response3" }}
  "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
  "name": "🔍 Observability"
{{ end }}
