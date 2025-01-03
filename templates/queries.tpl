{{- define "get_query" }}
  {{- if and (eq "OBJECT" .Kind) (not (contains "Connection" .String)) }}
    {{- range .Fields }}
      {{ if and (isObjectType .) (not (contains "Connection" .Type.String)) }}
      {{- if gt (len .Arguments) 0 }}
func ({{ lower $.Name }} *{{ $.Name }}) Get{{ title .Name }}(client *Client
          {{- range $index, $arg := .Arguments -}}
            , {{ $arg.Name.Name }} {{ trimSuffix "!" $arg.Type.String }}
          {{- end -}}
    ) (*{{ title .Name }}, error){
  var q struct {
    Account struct {
      {{ $.Name }} struct {
        {{ title .Name }} {{ trimSuffix "!" .Type.String }} `graphql:"{{ .Name }}(
            {{- range $index, $arg := .Arguments }}
              {{- if gt $index 0 }}, {{ end -}}
                {{ $arg.Name.Name }}: ${{ $arg.Name.Name }}
            {{- end }})"`
      } `graphql:"{{ lower $.Name }}(id: $id)"`
    }
  }
	if {{ lower $.Name }}.Id == "" {
		return nil, fmt.Errorf("unable to get {{ title .Name }}, invalid {{ lower $.Name }} id: '%s'", {{ lower $.Name }}.Id)
	}
  v := PayloadVariables{ {{ range .Arguments }}
    "{{.Name.Name}}": {{.Name.Name}}, {{ end }}
  }

	err := client.Query(&q, v, WithName("{{ $.Name }}{{ title .Name }}Get"))
  return &q.Account.{{ $.Name }}.{{ title .Name }}, HandleErrors(err, nil)
}
      {{ end }}
      {{ end }}
    {{ end }}
  {{- end }}
{{- end -}}

{{- define "list_query" }}
  {{- if eq "OBJECT" .Kind }}
    {{- range .Fields }}
      {{ if and (isObjectType .) (contains "Connection" .Type.String) }}
      {{- if gt (len .Arguments) 0 }}
func ({{ lower $.Name }} *{{ $.Name }}) List{{ title .Name | makePlural }}(client *Client, variables *PayloadVariables) (*{{ trimSuffix "!" .Type.String }}, error){
  var q struct {
    Account struct {
      {{ $.Name }} struct {
        {{ title .Name }} {{ trimSuffix "!" .Type.String }} `graphql:"{{ .Name }}(after: $after, first: $first)"`
      } `graphql:"{{ lower $.Name }}(id: $id)"`
    }
  }
	if {{ lower $.Name }}.Id == "" {
		return nil, fmt.Errorf("unable to get {{ title .Name }}, invalid {{ lower $.Name }} id: '%s'", {{ lower $.Name }}.Id)
	}
  if variables == nil {
    variables = client.InitialPageVariablesPointer()
  }

  if err := client.Query(&q, *variables, WithName("{{ title .Name }}List")); err != nil {
    return nil, err
  }
  for q.Account.{{ $.Name }}.{{ title .Name }}.PageInfo.HasNextPage {
    (*variables)["after"] = q.Account.{{ $.Name }}.{{ title .Name }}.PageInfo.End
    resp, err := {{ lower $.Name }}.List{{ title .Name | makePlural }}(client, variables)
    if err != nil {
      return nil, err
    }
    q.Account.{{ $.Name }}.{{ title .Name }}.Nodes = append(q.Account.{{ $.Name }}.{{ title .Name }}.Nodes, resp.Nodes...)
    q.Account.{{ $.Name }}.{{ title .Name }}.PageInfo = resp.PageInfo
    q.Account.{{ $.Name }}.{{ title .Name }}.TotalCount += resp.TotalCount
  }
  return &q.Account.{{ $.Name }}.{{ title .Name }}, nil
}
      {{ end }}
      {{ end }}
    {{ end }}
  {{- end }}
{{- end -}}

{{ define "account_get_query" -}}
  {{- range .Fields }}
    {{ if and (isObjectType .) (not (contains "Connection" .Type.String)) }}
func (client *Client) Get{{ title .Name }}(
          {{- range $index, $arg := .Arguments }}
            {{- if gt $index 0 }}, {{ end -}}
            {{ $arg.Name.Name }} {{ trimSuffix "!" $arg.Type.String }}
          {{- end -}}
    ) (*{{ trimSuffix "!" .Type.String }}, error){
  var q struct {
    Account struct {
      {{ title .Name }} {{ trimSuffix "!" .Type.String }} `graphql:"{{ .Name }}(
          {{- range $index, $arg := .Arguments }}
            {{- if gt $index 0 }}, {{ end -}}
            {{ $arg.Name.Name }}: ${{ $arg.Name.Name }}
          {{- end }})"`
    }
  }
  v := PayloadVariables{ {{ range .Arguments }}
    "{{.Name.Name}}": {{.Name.Name}}, {{ end }}
  }

	err := client.Query(&q, v, WithName("{{ title .Name  }}Get"))
  return &q.Account.{{ title .Name }}, HandleErrors(err, nil)
}
    {{ end }}
  {{- end }}
{{- end }}

{{ define "account_list_query" -}}
  {{- range .Fields }}
    {{- if contains "Connection" .Type.String }}
func (client *Client) List{{ title .Name | queryRename | makePlural }}(variables *PayloadVariables) (*{{ trimSuffix "!" .Type.String }}, error){
  var q struct {
    Account struct {
      {{ title .Name }} {{ trimSuffix "!" .Type.String }} `graphql:"{{ .Name }}(after: $after, first: $first)"`
    }
  }
  if variables == nil {
    variables = client.InitialPageVariablesPointer()
  }

  if err := client.Query(&q, *variables, WithName("{{ title .Name | queryRename }}List")); err != nil {
    return nil, err
  }
  for q.Account.{{ title .Name }}.PageInfo.HasNextPage {
    (*variables)["after"] = q.Account.{{ title .Name }}.PageInfo.End
    resp, err := client.List{{ title .Name | queryRename | makePlural }}(variables)
    if err != nil {
      return nil, err
    }
    q.Account.{{ title .Name }}.Nodes = append(q.Account.{{ title .Name }}.Nodes, resp.Nodes...)
    q.Account.{{ title .Name }}.PageInfo = resp.PageInfo
    q.Account.{{ title .Name }}.TotalCount += resp.TotalCount
  }
  return &q.Account.{{ title .Name }}, nil
}
    {{ end }}
  {{- end }}
{{- end }}
