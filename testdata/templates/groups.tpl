{{- define "group_1" }}
{
    "alias": "test_group_1",
    "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI",
    "description": null,
    "htmlUrl": "https://app.opslevel-staging.com/groups/test_group_1",
    "name": "test_group_1",
    "parent": null
}
{{ end }}
{{- define "group_2" }}
{
    "alias": "test_group_2",
    "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTE",
    "description": "test1123",
    "htmlUrl": "https://app.opslevel-staging.com/groups/test_group_2",
    "name": "test_group_2",
    "parent": {
      "alias": "test_group_1",
      "id": "Z2lkOi8vb3BzbGV2ZWwvTmFtZXNwYWNlczo6R3JvdXAvMTI"
    }
}
{{ end }}