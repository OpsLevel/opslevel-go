{{- define "service_1" }}
{
    "apiDocumentPath": "/src/swagger.json",
    "aliases": [
        "{{ template "alias1" }}"
    ],
    "description": "{{ template "description" }}",
    "framework": null,
    "htmlUrl": "https://app.opslevel.com/services/{{ template "alias1" }}",
    "id": "{{ template "id1" }}",
    "language": null,
    "lifecycle": null,
    "name": "Foo",
    "owner": null,
    "preferredApiDocument": null,
    "preferredApiDocumentSource": "PULL",
    "product": null,
    "tier": null
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
    "id": "{{ template "id1" }}",
    "language": null,
    "lifecycle": null,
    "name": "Bar",
    "owner": null,
    "preferredApiDocument": null,
    "preferredApiDocumentSource": "PULL",
    "product": null,
    "tier": null
}
{{ end }}