package opslevel

type RepositoryGrepCheckFragment struct {
	DirectorySearch       bool      `graphql:"directorySearch"`       // Whether the check looks for the existence of a directory instead of a file.
	FileContentsPredicate Predicate `graphql:"fileContentsPredicate"` // Condition to match the file content.
	Filepaths             []string  `graphql:"filePaths"`             // Restrict the search to certain file paths.
}

func (client *Client) CreateCheckRepositoryGrep(input CheckRepositoryGrepCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryGrepCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryGrepCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRepositoryGrep(input CheckRepositoryGrepUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryGrepUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryGrepUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
