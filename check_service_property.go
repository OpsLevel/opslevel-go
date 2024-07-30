package opslevel

type ServicePropertyCheckFragment struct {
	Property           ServicePropertyTypeEnum `graphql:"serviceProperty"`        // The property of the service that the check will verify.
	PropertyDefinition *IdentifierInput        `graphql:"propertyDefinition"`     // The secondary key of the property that the check will verify (e.g. the specific custom property).
	Predicate          *Predicate              `graphql:"propertyValuePredicate"` // The condition that should be satisfied by the service property value.
}

func (client *Client) CreateCheckServiceProperty(input CheckServicePropertyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServicePropertyCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckServiceProperty(input CheckServicePropertyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServicePropertyUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
