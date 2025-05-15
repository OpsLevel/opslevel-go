package opslevel

type ServiceDocument struct {
	Id         ID                    `graphql:"id" json:"id"`
	HtmlURL    string                `graphql:"htmlUrl" json:"htmlUrl,omitempty"`
	Source     ServiceDocumentSource `graphql:"source" json:"source"`
	Timestamps Timestamps            `graphql:"timestamps" json:"timestamps"`
}

type ServiceDocumentContent struct {
	ServiceDocument
	Content string `graphql:"content" json:"content,omitempty"`
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

func (client *Client) ListDocuments(variables *PayloadVariables) (*ServiceDocumentsConnection, error) {
	var q struct {
		Account struct {
			Documents ServiceDocumentsConnection `graphql:"documents(searchTerm: $searchTerm, after: $after, first: $first)"`
		}
	}

	variables = client.PopulatePaginationParams(variables)

	if (*variables)["searchTerm"] == nil {
		(*variables)["searchTerm"] = ""
	}

	if err := client.Query(&q, *variables, WithName("ListDocuments")); err != nil {
		return nil, err
	}
	q.Account.Documents.TotalCount = len(q.Account.Documents.Nodes)

	if q.Account.Documents.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Documents.PageInfo.End

		resp, err := client.ListDocuments(variables)
		if err != nil {
			return &q.Account.Documents, err
		}
		q.Account.Documents.Nodes = append(q.Account.Documents.Nodes, resp.Nodes...)
		q.Account.Documents.PageInfo = resp.PageInfo
		q.Account.Documents.TotalCount += resp.TotalCount
	}
	return &q.Account.Documents, nil
}

func (client *Client) GetDocument(id ID) (*ServiceDocumentContent, error) {
	var q struct {
		Account struct {
			Document ServiceDocumentContent `graphql:"document(id: $id)"`
		}
	}

	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("DocumentGet"))
	return &q.Account.Document, err
}
