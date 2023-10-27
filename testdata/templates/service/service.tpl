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
