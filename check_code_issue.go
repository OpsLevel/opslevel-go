package opslevel

type CodeIssueCheckFragment struct {
	Constraint     CheckCodeIssueConstraintEnum `graphql:"constraint"`     // The type of constraint used in evaluation the code issues check.
	IssueName      string                       `graphql:"issueName"`      // The issue name used for code issue lookup.
	IssueType      string                       `graphql:"issueType"`      // The type of code issue to consider.
	MaxAllowed     string                       `graphql:"maxAllowed"`     // The threshold count of code issues beyond which the check starts failing.
	ResolutionTime CodeIssueResolutionTime      `graphql:"resolutionTime"` // The resolution time recommended by the reporting source of the code issue.
	Severity       []string                     `graphql:"severity"`       // The severity levels of the issue.
}

// CodeIssueResolutionTime represents how long a code issue has been detected.
type CodeIssueResolutionTime struct {
	Unit  CodeIssueResolutionTimeUnitEnum `graphql:"unit"`  // The name of duration of time.
	Value string                          `graphql:"value"` // The count value of the specified unit.
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
