{{- define "alias1" }}example{{ end }}
{{- define "alias2" }}example_2{{ end }}
{{- define "name1" }}Example{{ end }}
{{- define "name2" }}Example 2{{ end }}
{{- define "id1" }}MTIzNDU2Nzg5MTIzNDU2Nzg5{{ end }}
{{- define "id2" }}OTg3NjU0MzIxOTg3NjU0MzIx{{ end }}
{{- define "eid1" }}NzQxODUyOTYzNzQxODUyOTYz{{ end }}
{{- define "description" }}An example description{{ end }}
{{- define "pagination_request" }}pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount{{ end }}
{{- define "no_pagination_response" }}"pageInfo": {
    "hasNextPage": false,
    "hasPreviousPage": false,
    "startCursor": "MQ",
    "endCursor": "NA"
}{{ end }}
{{- define "pagination_initial_query_variables" }}
"variables": {
    "after": "",
    "first": 100
}
{{ end }}
{{- define "pagination_second_query_variables" }}
"variables": {
	"after": "OA",
	"first": 100
}
{{ end }}
{{- define "pagination_initial_pageInfo_response" }}
"pageInfo": {
  "hasNextPage": true,
  "hasPreviousPage": false,
  "startCursor": "MQ",
  "endCursor": "OA"
}
{{ end }}
{{- define "pagination_second_pageInfo_response" }}
"pageInfo": {
  "hasNextPage": false,
  "hasPreviousPage": true,
  "startCursor": "OA",
  "endCursor": "EOf"
}
{{ end }}
{{- define "error1" }}{
    "message": "Example Error",
    "path": [
        "one",
        "two",
        "three"
    ]
}{{ end }}

