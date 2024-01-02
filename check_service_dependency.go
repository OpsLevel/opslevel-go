package opslevel

// type CheckServiceDependencyCreateInput struct {
// 	CheckCreateInput
// }

// type CheckServiceDependencyUpdateInput struct {
// 	CheckUpdateInput
// }

func (client *Client) CreateCheckServiceDependency(input CheckServiceDependencyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceDependencyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceDependencyCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckServiceDependency(input CheckServiceDependencyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceDependencyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceDependencyUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
