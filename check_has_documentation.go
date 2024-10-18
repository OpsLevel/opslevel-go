package opslevel

type HasDocumentationCheckFragment struct {
	DocumentSubtype HasDocumentationSubtypeEnum `graphql:"documentSubtype"` // The subtype of the document.
	DocumentType    HasDocumentationTypeEnum    `graphql:"documentType"`    // The type of the document.
}

func (client *Client) CreateCheckHasDocumentation(input CheckHasDocumentationCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasDocumentationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckHasDocumentationCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckHasDocumentation(input CheckHasDocumentationUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasDocumentationUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckHasDocumentationUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
