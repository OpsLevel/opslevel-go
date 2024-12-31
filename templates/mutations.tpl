{{ define "create_operation" -}}
{{- end }}

{{ define "update_operation" -}}
{{- end }}

{{ define "delete_operation" -}}
  {{- range .Fields }}
    {{- if hasSuffix "Delete" .Name }}
// {{ title $.Name }} {{ .Desc | replace "\n" " " | endSentence }}
func (client *Client) Delete{{ title .Name | trimSuffix "Delete" }}( {{- template "delete_input_args" . -}} ) error {
	var d struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"{{ .Name }}( {{- template "graphql_delete_input_args" . -}} )"`
	}
  v := PayloadVariables{
    {{ template "delete_payload_variables" . }}
  }
	err := client.Mutate(&d, v, WithName("{{ title .Name }}"))
	return HandleErrors(err, d.Payload.Errors)
}
    {{- end }}
  {{- end }}
{{- end }}

{{ define "delete_payload_variables" }}
  {{- range $index, $arg := .Arguments -}}
    {{- if hasPrefix "IdentifierInput" .Type.String -}}
      "{{.Name.Name}}": identifier, {{- nindent 4 "" -}}
    {{- else if contains (trimSuffix "!" .Type.String) (list "ContactDeleteInput" "DeleteInput" | join " " ) -}}
      "{{.Name.Name}}": {{ trimSuffix "!" .Type.String }}{ {{- nindent 4 "" -}}
        Id: id, {{- nindent 4 "" -}}
      {{- nindent 4 "}," -}}
    {{- else if contains (trimSuffix "!" .Type.String) (list "AliasDeleteInput" "String" | join " " ) -}}
      "{{.Name.Name}}": {{.Name.Name}}, {{- nindent 4 "" -}}
    {{- else -}}
      "{{.Name.Name}}": {{ trimSuffix "!" .Type.String }}{ {{- nindent 4 "" -}}
      {{- with deleteInputRename $arg -}}
        {{- if eq . "ID" -}}
          Id: id,
        {{- else -}}
          {{ $arg.Name.Name }}: {{ . }},
        {{ end -}}
      {{- end -}}
    {{- nindent 4 "}," -}}
    {{- end -}}
  {{- end -}}
{{- end }}

{{ define "delete_input_args" }}
  {{- range $index, $arg := .Arguments -}} {{- if gt $index 0 }}, {{ end -}}
    {{- with deleteInputRename $arg -}}
      {{- if eq . "ID" -}}
        id ID
      {{- else if eq . "String" -}}
        {{ $arg.Name.Name }} string
      {{- else if eq . "string" -}}
        identifier string
      {{- else -}}
        {{ $arg.Name.Name }} {{ . }}
      {{- end -}}
    {{- end -}}
  {{- end -}}
{{- end }}

{{ define "graphql_delete_input_args" }}
  {{- range $index, $arg := .Arguments -}} {{- if gt $index 0 }}, {{ end -}}
    {{ $arg.Name.Name }}: ${{ $arg.Name.Name }}
  {{- end -}}
{{- end }}
