package opslevel

type CustomEventCheckFragment struct {
	Integration      IntegrationId `graphql:"integration"`      // The integration this check uses.
	PassPending      bool          `graphql:"passPending"`      // True if this check should pass by default. Otherwise the default 'pending' state counts as a failure.
	ResultMessage    string        `graphql:"resultMessage"`    // The check result message template.
	ServiceSelector  string        `graphql:"serviceSelector"`  // A jq expression that will be ran against your payload to select the service.
	SuccessCondition string        `graphql:"successCondition"` // A jq expression that will be ran against your payload to evaluate the check result. A truthy value will result in the check passing.
}

func (client *Client) CreateCheckCustomEvent(input CheckCustomEventCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckCustomEventCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckCustomEvent(input CheckCustomEventUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckCustomEventUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
