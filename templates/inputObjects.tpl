{{- define "inputs" }}
// {{ title .Name }} {{ clean .Desc | endSentence }}
type {{ .Name }} struct { {{ range .Values }}
  {{ title .Name.Name }} {{ getFieldType .Type }} `
      {{- jsonStructTag . }} {{ yamlStructTag . }}
      {{- exampleStructTag . }}` {{ fieldCommentDescription . }}
  {{- end }}
}
{{- end -}}
