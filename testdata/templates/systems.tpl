{{- define "system1_response" }}
{
    {{ template "id1" }},
    "aliases": [
      "platformsystem1"
    ],
    "name": "PlatformSystem1",
    "description": "Yolo!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem1",
    "owner": {
      {{ template "id4" }}
    },
    "parent": {{ template "domain1_response" }},
    "note": "{{ template "description" }}"
}
{{end}}
{{- define "system2_response" }}
{
    {{ template "id2" }},
    "aliases": [
      "platformsystem2"
    ],
    "name": "PlatformSystem2",
    "description": "Yolo2!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem2",
    "owner": {
      {{ template "id4" }}
    },
    "note": "{{ template "description" }}"
}
{{end}}
{{- define "system3_response" }}
{
    {{ template "id3" }},
    "aliases": [
      "platformsystem3"
    ],
    "name": "PlatformSystem3",
    "description": "Yolo3!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem3",
    "owner": {
      "teamAlias": "kyle_team",
      {{ template "id4" }}
    },
    "parent": {{ template "domain1_response" }},
    "note": "{{ template "description" }}"
}
{{end}}
