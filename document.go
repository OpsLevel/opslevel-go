package opslevel

type ServiceDocument struct {
	Id         ID                    `graphql:"id" json:"id"`
	HtmlURL    string                `graphql:"htmlUrl" json:"htmUrl,omitempty"`
	Source     ServiceDocumentSource `graphql:"source" json:"source"`
	Timestamps Timestamps            `graphql:"timestamps" json:"timestamps"`
}

type ServiceDocumentContent struct {
	ServiceDocument
	Content string `graphql:"content" json:"content,omitempty"`
}

type ServiceDocumentsContentConnection struct {
	Nodes      []ServiceDocumentContent
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

func (client *Client) ServiceApiDocSettingsUpdate(service string, docPath string, docSource *ApiDocumentSourceEnum) (*Service, error) {
	var m struct {
		Payload ServiceUpdatePayload `graphql:"serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource)"`
	}
	v := PayloadVariables{
		"service":   *NewIdentifier(service),
		"docPath":   (*string)(nil),
		"docSource": docSource,
	}
	if docPath != "" {
		v["docPath"] = &docPath
	}
	err := client.Mutate(&m, v, WithName("ServiceApiDocSettingsUpdate"))
	return &m.Payload.Service, HandleErrors(err, m.Payload.Errors)
}

type DocumentsConnection struct {
	Nodes []ServiceDocument
}

func (client *Client) DocumentSearch(searchTerm string) (*DocumentsConnection, error) {
	var q struct {
		Account struct {
			Documents DocumentsConnection `graphql:"documents(searchTerm: $searchTerm)"`
		}
	}

	v := PayloadVariables{
		"searchTerm": &searchTerm,
	}
	err := client.Query(&q, v, WithName("DocumentSearch"))
	return &q.Account.Documents, err
}
