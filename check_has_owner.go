package opslevel

type ServiceOwnershipCheckFragment struct {
	ContactMethod        *ContactType `graphql:"contactMethod"`        // The type of contact method that an owner should provide.
	RequireContactMethod *bool        `graphql:"requireContactMethod"` // Whether to require a contact method for a service owner or not.
	TeamTagKey           string       `graphql:"tagKey"`               // The tag key that should exist for a service owner.
	TeamTagPredicate     *Predicate   `graphql:"tagPredicate"`         // The condition that should be satisfied by the tag value.
}

func (client *Client) CreateCheckServiceOwnership(input CheckServiceOwnershipCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceOwnershipCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceOwnershipCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckServiceOwnership(input CheckServiceOwnershipUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceOwnershipUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceOwnershipUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
