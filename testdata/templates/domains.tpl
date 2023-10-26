{{- define "domain1_response" }}
{
    {{ template "id1" }},
    "aliases": [
      "platformdomain"
    ],
    "name": "PlatformDomain1",
    "description": "Our first Platform Domain!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/domains/platformdomain",
    "owner": {
      "groupAlias": "kyle_team",
      {{ template "id4" }}
    },
    "note": "{{ template "description" }}"
}
{{end}}
{{- define "domain2_response" }}
{
    {{ template "id2" }},
    "aliases": [
      "platformdomain2"
    ],
    "name": "PlatformDomain2",
    "description": "Our second domain!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/domains/platformdomain2",
    "owner": {
      "groupAlias": "kyle_team",
      {{ template "id4" }}
    },
    "note": "{{ template "description" }}"
}
{{end}}
{{- define "domain3_response" }}
{
    {{ template "id3" }},
    "aliases": [
      "platformdomain3"
    ],
    "name": "PlatformDomain3",
    "description": "Our third domain!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/domains/platformdomain3",
    "owner": {
      "teamAlias": "kyle_team",
      "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83AbC"
    },
    "note": "{{ template "description" }}"
  }
{{end}}
