package opslevel

type TagDefinedCheckFragment struct {
	TagKey       string     `graphql:"tagKey"`       // The tag key where the tag predicate should be applied.
	TagPredicate *Predicate `graphql:"tagPredicate"` // The condition that should be satisfied by the tag value.
}

func (client *Client) CreateCheckTagDefined(input CheckTagDefinedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckTagDefinedCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckTagDefined(input CheckTagDefinedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckTagDefinedUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
