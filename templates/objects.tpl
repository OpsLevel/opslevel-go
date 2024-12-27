{{- define "objects" }}
// {{.Name | title}} {{.Desc | clean | endSentence}}
type {{.Name}} struct { {{ range .Fields }}
  {{- if eq (len .Arguments) 0 }}
    {{ title .Name }} {{ getFieldTypeForObject . }} `graphql:"{{ .Name }}" json:"{{ .Name }}"
      {{- if hasPrefix "JSON" .Type.String }} scalar:"true"{{ end }}` // {{ .Desc | replace "\n" " " }}
  {{- end -}}{{- end }}
}
{{- end -}}
