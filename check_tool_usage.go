package opslevel

type ToolUsageCheckFragment struct {
	ToolCategory         ToolCategory `graphql:"toolCategory"`
	ToolNamePredicate    *Predicate   `graphql:"toolNamePredicate"`
	ToolUrlPredicate     *Predicate   `graphql:"toolUrlPredicate"`
	EnvironmentPredicate *Predicate   `graphql:"environmentPredicate"`
}

// type CheckToolUsageCreateInput struct {
// 	CheckCreateInput

// 	ToolCategory         ToolCategory    `json:"toolCategory" yaml:"toolCategory" default:"backlog"`
// 	ToolNamePredicate    *PredicateInput `json:"toolNamePredicate,omitempty" yaml:"toolNamePredicate,omitempty"`
// 	ToolUrlPredicate     *PredicateInput `json:"toolUrlPredicate,omitempty" yaml:"toolUrlPredicate,omitempty"`
// 	EnvironmentPredicate *PredicateInput `json:"environmentPredicate,omitempty" yaml:"environmentPredicate,omitempty"`
// }

// type CheckToolUsageUpdateInput struct {
// 	CheckUpdateInput

// 	ToolCategory         ToolCategory    `json:"toolCategory,omitempty"`
// 	ToolNamePredicate    *PredicateInput `json:"toolNamePredicate,omitempty"`
// 	ToolUrlPredicate     *PredicateInput `json:"toolUrlPredicate,omitempty"`
// 	EnvironmentPredicate *PredicateInput `json:"environmentPredicate,omitempty"`
// }

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
