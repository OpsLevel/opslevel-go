package opslevel

type ServiceDocumentSource struct {
	IntegrationId     `graphql:"... on ApiDocIntegration"`
	ServiceRepository `graphql:"... on ServiceRepository"`
}

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
		Payload struct {
			Service Service
			Errors  []OpsLevelErrors
		} `graphql:"serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource)"`
	}
	v := PayloadVariables{
		"service":   *NewIdentifier(service),
		"docPath":   NullString(),
		"docSource": docSource,
	}
	if docPath != "" {
		v["docPath"] = RefOf(docPath)
	}
	err := client.Mutate(&m, v, WithName("ServiceApiDocSettingsUpdate"))
	return &m.Payload.Service, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) ServiceDocumentSearch(service string, searchTerm string, docType *HasDocumentationTypeEnum, hidden *bool) ([]ServiceDocumentContent, error) {
	var m struct {
		Payload struct {
			Service struct {
				Documents ServiceDocumentsContentConnection `graphql:"documents(searchTerm: $searchTerm, type: $type, hidden: $hidden)"`
			}
		} `graphql:"service(service: $service)"`
	}
	v := PayloadVariables{
		"service":    *NewIdentifier(service),
		"searchTerm": &searchTerm,
		"type":       docType,
		"hidden":     hidden,
	}
	err := client.Query(&m, v, WithName("ServiceDocumentSearch"))
	return m.Payload.Service.Documents.Nodes, err
}
