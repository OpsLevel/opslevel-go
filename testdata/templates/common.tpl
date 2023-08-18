{{- define "alias1" }}example{{ end }}
{{- define "alias2" }}example_2{{ end }}
{{- define "alias3" }}example_3{{ end }}
{{- define "alias4" }}example_4{{ end }}
{{- define "name1" }}Example{{ end }}
{{- define "name2" }}Example 2{{ end }}
{{- define "name3" }}Example 3{{ end }}
{{- define "name4" }}Example 4{{ end }}
{{- define "email1" }}kyle@opslevel.com{{ end }}
{{- define "email2" }}edgar@opslevel.com{{ end }}
{{- define "email3" }}matthew@opslevel.com{{ end }}
{{- define "id1" }}Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx{{ end }}
{{- define "id2" }}Z2lkOi8vOTg3NjU0MzIxMTIzNDU2Nzg5{{ end }}
{{- define "id3" }}Z2lkOi8vMTkyODM3NDY1NTY0NzM4Mjkx{{ end }}
{{- define "id4" }}Z2lkOi8vMTkyODM3NDY1{{ end }}
{{- define "eid1" }}NzQxODUyOTYzNzQxODUyOTYz{{ end }}
{{- define "description" }}An example description{{ end }}
{{- define "pagination_request" }}pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount{{ end }}
{{- define "no_pagination_response" }}"pageInfo": {
    "hasNextPage": false,
    "hasPreviousPage": false,
    "startCursor": "MQ",
    "endCursor": "NA"
}{{ end }}
{{- define "first_page_variables" }}
"after": "",
"first": 100
{{ end }}
{{- define "second_page_variables" }}
"after": "OA",
"first": 100
{{ end }}
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
{{- define "next_page_false" }}
{
  "hasNextPage": false,
  "hasPreviousPage": true,
  "startCursor": "OA",
  "endCursor": "EOf"
}
{{ end }}
{{- define "next_page_true" }}
{
  "hasNextPage": true,
  "hasPreviousPage": false,
  "startCursor": "MQ",
  "endCursor": "OA"
}
{{ end }}
{{- define "pagination_initial_pageInfo_response" }}
"pageInfo": {{ template "next_page_true" }}
{{ end }}
{{- define "pagination_second_pageInfo_response" }}
"pageInfo": {{ template "next_page_false" }}
{{ end }}
{{- define "error1" }}{
    "message": "Example Error",
    "path": [
      "one",
      "two",
      "three"
    ]
}{{ end }}

