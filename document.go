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

type Document struct {
	Id ID `graphql:"id" json:"id"`
}

type DocumentsConnection struct {
	Nodes []Document
}

func (client *Client) DocumentSearch(searchTerm string) ([]Document, error) {
	var q struct {
		Account struct {
			Documents DocumentsConnection `graphql:"documents(searchTerm: $searchTerm)"`
		}
	}

	v := PayloadVariables{
		"searchTerm": &searchTerm,
	}
	err := client.Query(&q, v, WithName("DocumentSearch"))
	return q.Account.Documents.Nodes, err
}

// func (client *Client) ServiceDocumentSearch(searchTerm string) ([]Document, error) {
// 	var q struct {
// 		Account struct {
// 			Services struct {
// 				Documents DocumentsConnection `graphql:"documents(searchTerm: $searchTerm)"`
// 			}
// 		}
// 	}

// 	v := PayloadVariables{
// 		"searchTerm": &searchTerm,
// 	}
// 	err := client.Query(&q, v, WithName("ServiceDocumentSearch"))
// 	return q.Account.Services.Nodes, err
// }
