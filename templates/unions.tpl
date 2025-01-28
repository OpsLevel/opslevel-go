{{- define "union" }}
// {{ title .Name }} {{.Desc | clean | endSentence}}
  {{- if eq "CheckOwner" .Name }}
    {{ template "check_owner" }}
  {{- else if eq "EntityOwner" .Name }}
    {{ template "entity_owner" }}
  {{- else if eq "ServiceDocumentSource" .Name }}
    {{ template "service_document_source" }}
  {{- else }}
type {{.Name}} struct {
 {{ range .TypeNames }}
    {{- if not (contains "Group" . ) }}
  {{.}} {{. -}}
      {{- if not (contains . (list "ApiDocIntegration" "InfrastructureResource" "ServiceRepository" | join " " )) -}}Id
      {{- end }} `graphql:"... on {{.}}"`
    {{- end }}
  {{- end }}
}
 {{- end }}
{{- end -}}

{{- define "check_owner" -}}
type CheckOwner struct {
	Team TeamId `graphql:"... on Team"`
	// User UserId `graphql:"... on User"` // TODO: will this be public?
}
{{ end }}

{{- define "entity_owner" -}}
type EntityOwner struct {
	OnTeam EntityOwnerTeam `graphql:"... on Team"`
}
{{ end }}

{{- define "service_document_source" -}}
type ServiceDocumentSource struct {
	IntegrationId     `graphql:"... on ApiDocIntegration"`
	ServiceRepository `graphql:"... on ServiceRepository"`
}
{{ end }}
