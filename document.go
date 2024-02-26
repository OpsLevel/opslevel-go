package opslevel

// NOTE: test first, then replace this struct with commented out one below
type ServiceDocumentSource struct {
	IntegrationId     `graphql:"... on ApiDocIntegration"`
	ServiceRepository `graphql:"... on ServiceRepository"`
}

// type ServiceDocumentSource struct {
// 	IntegrationId     IntegrationId `graphql:"... on ApiDocIntegration"`
// 	ServiceRepository ServiceRepository `graphql:"... on ServiceRepository"`
// }

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
		"docPath":   NullString(),
		"docSource": docSource,
	}
	if docPath != "" {
		v["docPath"] = RefOf(docPath)
	}
	err := c.Mutate(&m, v, WithName("ServiceApiDocSettingsUpdate"))
	return &m.Payload.Service, HandleErrors(err, m.Payload.Errors)
}
