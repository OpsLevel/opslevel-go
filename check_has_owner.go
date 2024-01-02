package opslevel

type ServiceOwnershipCheckFragment struct {
	RequireContactMethod *bool        `graphql:"requireContactMethod"`
	ContactMethod        *ContactType `graphql:"contactMethod"`
	TeamTagKey           string       `graphql:"tagKey"`
	TeamTagPredicate     *Predicate   `graphql:"tagPredicate"`
}

// type CheckServiceOwnershipCreateInput struct {
// 	CheckCreateInput

// 	RequireContactMethod *bool           `json:"requireContactMethod,omitempty"`
// 	ContactMethod        *ContactType    `json:"contactMethod,omitempty"`
// 	TeamTagKey           string          `json:"tagKey,omitempty"`
// 	TeamTagPredicate     *PredicateInput `json:"tagPredicate,omitempty"`
// }

// type CheckServiceOwnershipUpdateInput struct {
// 	CheckUpdateInput

// 	RequireContactMethod *bool                 `json:"requireContactMethod,omitempty"`
// 	ContactMethod        *ContactType          `json:"contactMethod,omitempty"`
// 	TeamTagKey           string                `json:"tagKey,omitempty"`
// 	TeamTagPredicate     *PredicateUpdateInput `json:"tagPredicate,omitempty"`
// }

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
