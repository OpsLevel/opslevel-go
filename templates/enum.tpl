{{- define "enum" }}

// {{.Name | title}} {{.Desc | clean | endSentence}}
type {{.Name}} string

const (
{{- range .EnumValuesDefinition }}
	{{$.Name}}{{.EnumValue | enumIdentifier}} {{$.Name}} = {{.EnumValue | quote}} // {{ .Desc | clean | fullSentence }}
{{- end }}
)
// All {{$.Name}} as []string
var All{{$.Name}} = []string {
	{{range .EnumValuesDefinition}}string({{$.Name}}{{.EnumValue | enumIdentifier}}),
	{{end}}
}
{{- end -}}
