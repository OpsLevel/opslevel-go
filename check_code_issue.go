package opslevel

type CodeIssueCheckFragment struct {
	Campaign string `graphql:"campaign"` // The campaign the check belongs to.
	Url      string `graphql:"url"`      // The url to the check.
}

func (client *Client) CreateCheckCodeIssue(input CheckCodeIssueCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCodeIssueCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckCodeIssueCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckCodeIssue(input CheckCodeIssueUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCodeIssueUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckCodeIssueUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
