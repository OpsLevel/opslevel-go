package opslevel

type RelationshipDefinitionPayload struct {
	Definition RelationshipDefinitionType // The relationship that was defined.
	BasePayload
}

func (client *Client) CreateRelationshipDefinition(input RelationshipDefinitionInput) (*RelationshipDefinitionType, error) {
	var m struct {
		Payload RelationshipDefinitionPayload `graphql:"relationshipDefinitionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipDefinitionCreate"))
	return &m.Payload.Definition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateRelationshipDefinition(identifier string, input RelationshipDefinitionInput) (*RelationshipDefinitionType, error) {
	var m struct {
		Payload RelationshipDefinitionPayload `graphql:"relationshipDefinitionUpdate(relationshipDefinition: $identifier, input: $input)"`
	}
	v := PayloadVariables{
		"identifier": NewIdentifier(identifier),
		"input":      input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipDefinitionUpdate"))
	return &m.Payload.Definition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteRelationshipDefinition(identifier string) (*ID, error) {
	var m struct {
		Payload struct {
			DeletedId ID      `graphql:"deletedId"`
			Errors    []Error `graphql:"errors"`
		} `graphql:"relationshipDefinitionDelete(resource: $input)"`
	}
	input := *NewIdentifier(identifier)
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipDefinitionDelete"))
	return &m.Payload.DeletedId, HandleErrors(err, m.Payload.Errors)
}
