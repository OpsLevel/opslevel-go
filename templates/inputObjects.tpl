{{- define "inputs" }}
// {{ .Name }} {{ clean .Desc | endSentence }}
type {{ .Name }} struct { {{ range .Values }}
  {{ title .Name.Name }} {{ if and (eq $.Name "AliasCreateInput") (eq .Name.Name "ownerId") }}ID
   {{- else if and (eq $.Name "CheckPackageVersionUpdateInput") (eq .Name.Name "versionConstraintPredicate") }}*PredicateUpdateInput
   {{- else if and (eq $.Name "ManualCheckFrequencyUpdateInput") (eq .Name.Name "frequencyValue") }}*Nullable[int]
   {{- else if eq $.Name "IdentifierInput" }}
     {{- if eq .Type.String "String" }}*string
     {{- else }}*ID
     {{- end }}
   {{- else }}{{ getFieldTypeForInputObject .Type }}
                         {{- end }} `
      {{- jsonStructTag . }} {{ yamlStructTag . }}
      {{- exampleStructTag . }}` {{ fieldCommentDescription . }}
  {{- end }}
}
{{- end -}}
