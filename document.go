package opslevel

import "github.com/shurcooL/graphql"

type ServiceDocumentSource struct {
	Integration       `graphql:"... on ApiDocIntegration"`
	ServiceRepository `graphql:"... on ServiceRepository"`
}

type ServiceDocument struct {
	Id         graphql.ID            `graphql:"id" json:"id"`
	HtmlURL    string                `graphql:"htmlUrl" json:"htmUrl,omitempty"`
	Source     ServiceDocumentSource `graphql:"source" json:"source"`
	Timestamps Timestamps            `graphql:"timestamps" json:"timestamps"`
}

type ServiceDocumentContent struct {
	ServiceDocument
	Content string `graphql:"content" json:"content,omitempty"`
}

func (c *Client) ServiceApiDocSettingsUpdate(service string, docPath string, docSource *ApiDocumentSourceEnum) (*Service, error) {
	var m struct {
		Payload struct {
			Service Service
			Errors  []OpsLevelErrors
		} `graphql:"serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource)"`
	}
	v := PayloadVariables{
		"service":   *NewIdentifier(service),
		"docPath":   NewString(docPath),
		"docSource": docSource,
	}
	if err := c.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Service, FormatErrors(m.Payload.Errors)
}
