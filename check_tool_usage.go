package opslevel

type ToolUsageCheckFragment struct {
	EnvironmentPredicate *Predicate   `graphql:"environmentPredicate"` // The condition that the environment should satisfy to be evaluated.
	ToolCategory         ToolCategory `graphql:"toolCategory"`         // The category that the tool belongs to.
	ToolNamePredicate    *Predicate   `graphql:"toolNamePredicate"`    // The condition that the tool name should satisfy to be evaluated.
	ToolUrlPredicate     *Predicate   `graphql:"toolUrlPredicate"`     // The condition that the tool url should satisfy to be evaluated.
}

func (client *Client) CreateCheckToolUsage(input CheckToolUsageCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckToolUsageCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckToolUsage(input CheckToolUsageUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckToolUsageUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
