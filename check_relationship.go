package opslevel

type RelationshipCheckFragment struct {
	RelationshipCountPredicate *Predicate                 `graphql:"relationshipCountPredicate"` // The condition that should be satisfied by the number of RelatedTo relationships
	RelationshipDefinition     RelationshipDefinitionType `graphql:"relationshipDefinition"`     // The relationship definition that the check is based on.
}

func (client *Client) CreateCheckRelationship(input CheckRelationshipCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRelationshipCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRelationshipCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRelationship(input CheckRelationshipUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRelationshipUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRelationshipUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
