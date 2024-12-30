{{- define "inputs" }}
// {{ .Name }} {{ clean .Desc | endSentence }}
type {{ .Name }} struct { {{ range .Values }}
  {{ title .Name.Name }} {{ if and (eq $.Name "AliasCreateInput") (eq .Name.Name "ownerId") }}ID
   {{- else if and (eq $.Name "CheckPackageVersionUpdateInput") (eq .Name.Name "versionConstraintPredicate") }}*Nullable[PredicateUpdateInput]
   {{- else }}{{ getFieldTypeForInputObject .Type }}
                         {{- end }} `
      {{- jsonStructTag . }} {{ yamlStructTag . }}
      {{- exampleStructTag . }}` {{ fieldCommentDescription . }}
  {{- end }}
}
{{- end -}}
