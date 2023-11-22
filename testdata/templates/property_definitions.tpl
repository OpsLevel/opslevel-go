{{- define "json_schema_string" -}}
{\"$ref\":\"#/$defs/MyProp\",\"$defs\":{\"MyProp\":{\"properties\":{\"name\":{\"type\":\"string\",\"title\":\"the name\",\"description\":\"The name of a friend\",\"default\":\"alex\",\"examples\":[\"joe\",\"lucy\"]}},\"additionalProperties\":false,\"type\":\"object\",\"required\":[\"name\"]}}}
{{- end}}
