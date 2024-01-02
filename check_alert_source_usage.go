package opslevel

type AlertSourceUsageCheckFragment struct {
	AlertSourceNamePredicate Predicate           `graphql:"alertSourceNamePredicate"`
	AlertSourceType          AlertSourceTypeEnum `graphql:"alertSourceType"`
}

// type CheckAlertSourceUsageCreateInput struct {
// 	CheckCreateInput

// 	AlertSourceType          AlertSourceTypeEnum `json:"alertSourceType,omitempty"`
// 	AlertSourceNamePredicate *PredicateInput     `json:"alertSourceNamePredicate,omitempty"`
// }

// type CheckAlertSourceUsageUpdateInput struct {
// 	CheckUpdateInput

// 	AlertSourceType          AlertSourceTypeEnum   `json:"alertSourceType,omitempty"`
// 	AlertSourceNamePredicate *PredicateUpdateInput `json:"alertSourceNamePredicate,omitempty"`
// }

func (client *Client) CreateCheckAlertSourceUsage(input CheckAlertSourceUsageCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkAlertSourceUsageCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckAlertSourceUsageCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckAlertSourceUsage(input CheckAlertSourceUsageUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkAlertSourceUsageUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckAlertSourceUsageUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
