package opslevel

type ServicePropertyCheckFragment struct {
	Property  ServicePropertyTypeEnum `graphql:"serviceProperty"`
	Predicate *Predicate              `graphql:"propertyValuePredicate"`
}

type CheckServicePropertyCreateInput struct {
	CheckCreateInput

	Property  ServicePropertyTypeEnum `json:"serviceProperty"`
	Predicate *PredicateInput         `json:"propertyValuePredicate,omitempty"`
}

type CheckServicePropertyUpdateInput struct {
	CheckUpdateInput

	Property  ServicePropertyTypeEnum `json:"serviceProperty,omitempty"`
	Predicate *PredicateInput         `json:"propertyValuePredicate,omitempty"`
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
