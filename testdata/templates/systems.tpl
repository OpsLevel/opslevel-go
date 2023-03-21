{{- define "system1_response" }}
{
    "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy",
    "aliases": [
      "platformsystem1"
    ],
    "name": "PlatformSystem1",
    "description": "Yolo!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem1",
    "owner": {
      "alias": "kyle_team",
      "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"
    }
}
{{end}}
{{- define "system2_response" }}
{
    "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMz",
    "aliases": [
      "platformsystem2"
    ],
    "name": "PlatformSystem2",
    "description": "Yolo2!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem2",
    "owner": {
      "alias": "kyle_team",
      "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"
    }
}
{{end}}
{{- define "system3_response" }}
{
    "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzM0",
    "aliases": [
      "platformsystem3"
    ],
    "name": "PlatformSystem3",
    "description": "Yolo3!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem3",
    "owner": {
      "alias": "kyle_team",
      "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"
    }
}
{{end}}
{{- define "system_create_response" }}
"systemCreate": {
  "system": {
    "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUz",
    "aliases": [
      "platformsystem3"
    ],
    "name": "PlatformSystem3",
    "description": "creating this for testing purposes",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem3",
    "owner": {
      "alias": "kyle_team",
      "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"
    },
    "parent": null
  },
  "errors": []
}
{{end}}
{{- define "system_getalias_response" }}
"account": {
  "system": {
    "id": "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzMy",
    "aliases": [
      "platformsystem1"
    ],
    "name": "PlatformSystem1",
    "description": "Yolo!",
    "htmlUrl": "https://app.opslevel-staging.com/catalog/systems/platformsystem1",
    "owner": {
      "alias": "kyle_team",
      "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"
    }
  }
}
{{end}}