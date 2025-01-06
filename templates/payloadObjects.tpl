{{- define "payload_objects" }}
// {{.Name | title}} {{.Desc | clean | endSentence}}
type {{.Name}} struct { {{ range .Fields }}
  {{- if eq (len .Arguments) 0 }}
    {{ title .Name }} {{ getFieldTypeForObject . | trimPrefix "*" }} `graphql:"{{ .Name }}"` // {{ .Desc | replace "\n" " " }}
  {{- end -}}{{- end }}
}
{{- end -}}
