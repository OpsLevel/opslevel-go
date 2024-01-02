package opslevel

type TagDefinedCheckFragment struct {
	TagKey       string     `graphql:"tagKey"`
	TagPredicate *Predicate `graphql:"tagPredicate"`
}

// type CheckTagDefinedCreateInput struct {
// 	CheckCreateInput

// 	TagKey       string          `json:"tagKey" yaml:"tagKey" default:"app"`
// 	TagPredicate *PredicateInput `json:"tagPredicate,omitempty" yaml:"tagPredicate,omitempty" default:"{\"TagKey\":\"equals\", \"Value\":\"postgres\"}"`
// }

// type CheckTagDefinedUpdateInput struct {
// 	CheckUpdateInput

// 	TagKey       string          `json:"tagKey,omitempty"`
// 	TagPredicate *PredicateInput `json:"tagPredicate,omitempty"`
// }

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
