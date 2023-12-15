package opslevel

type HasDocumentationCheckFragment struct {
	DocumentType    HasDocumentationTypeEnum    `graphql:"documentType"`
	DocumentSubtype HasDocumentationSubtypeEnum `graphql:"documentSubtype"`
}

type CheckHasDocumentationCreateInput struct {
	CheckCreateInput

	DocumentType    HasDocumentationTypeEnum    `json:"documentType" yaml:"documentType" default:"api"`
	DocumentSubtype HasDocumentationSubtypeEnum `json:"documentSubtype" yaml:"documentSubtype" default:"openapi"`
}

type CheckHasDocumentationUpdateInput struct {
	CheckUpdateInput

	DocumentType    HasDocumentationTypeEnum    `json:"documentType"`
	DocumentSubtype HasDocumentationSubtypeEnum `json:"documentSubtype"`
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
