{{- define "service_1" }}
{
    "apiDocumentPath": "/src/swagger.json",
    "aliases": [
        "{{ template "alias1" }}"
    ],
    "description": "{{ template "description" }}",
    "framework": null,
    "htmlUrl": "https://app.opslevel.com/services/{{ template "alias1" }}",
    {{ template "id1" }},
    "language": null,
    "lifecycle": null,
    "locked": true,
    "name": "Foo",
    "owner": null,
    "preferredApiDocument": null,
    "preferredApiDocumentSource": "PULL",
    "product": null,
    "repos": {
        "edges": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    },
    "tags": {
        "nodes": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    },
    "tier": null,
    "tools": {
        "nodes": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    }
}
{{ end }}
{{- define "service_2" }}
{
    "apiDocumentPath": "/src/swagger.json",
    "aliases": [
        "{{ template "alias2" }}"
    ],
    "description": "{{ template "description" }}",
    "framework": null,
    "htmlUrl": "https://app.opslevel.com/services/{{ template "alias2" }}",
    {{ template "id1" }},
    "language": null,
    "lifecycle": null,
    "name": "Bar",
    "owner": null,
    "preferredApiDocument": null,
    "preferredApiDocumentSource": "PULL",
    "product": null,
    "defaultServiceRepository": {
        "baseDirectory": "backend",
        "displayName": "Backend",
        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZVJlcG9zaXRvcnkvMTg3",
        "repository": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvUmVwb3NpdG9yaWVzOjpCaXRidWNrZXQvMjYw",
            "defaultAlias": "bitbucket.org:raptors-store/Store Front"
        },
        "service": {
        "id": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS8zMQ",
        "aliases": [
          "Back End",
          "Backend Service",
          "a/b/c",
          "back end testing",
          "back_end",
          "fs-prod:deployment/bolt-http",
          "shopping_barts",
          "shopping_cart_service",
          "testing_1",
          "testing_11",
          "testing_12",
          "testing_123",
          "testing_1234",
          "testing_15",
          "testing_2",
          "testing_3",
          "testing_4",
          "testing_5",
          "testing_6",
          "testing_8"
        ]
        }
    },
    "repos": {
        "edges": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    },
    "tags": {
        "nodes": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    },
    "tier": null,
    "tools": {
        "nodes": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    }
}
{{ end }}
{{- define "service_with_note" }}
{
    "apiDocumentPath": "/src/swagger.json",
    "aliases": [
        "{{ template "alias1" }}"
    ],
    "description": "{{ template "description" }}",
    "framework": null,
    "htmlUrl": "https://app.opslevel.com/services/{{ template "alias1" }}",
    {{ template "id1" }},
    "language": null,
    "lifecycle": null,
    "locked": true,
    "name": "Foo",
    "note": "Foo note",
    "owner": null,
    "preferredApiDocument": null,
    "preferredApiDocumentSource": "PULL",
    "product": null,
    "repos": {
        "edges": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    },
    "tags": {
        "nodes": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    },
    "tier": null,
    "tools": {
        "nodes": [],
        "pageInfo": {
        "hasNextPage": false,
        "hasPreviousPage": false,
        "startCursor": null,
        "endCursor": null
        },
        "totalCount": 0
    }
}
{{ end }}
