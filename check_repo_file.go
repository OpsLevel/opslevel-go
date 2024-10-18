package opslevel

type RepositoryFileCheckFragment struct {
	DirectorySearch       bool       `graphql:"directorySearch"`       // Whether the check looks for the existence of a directory instead of a file.
	FileContentsPredicate *Predicate `graphql:"fileContentsPredicate"` // Condition to match the file content.
	Filepaths             []string   `graphql:"filePaths"`             // Restrict the search to certain file paths.
	UseAbsoluteRoot       bool       `graphql:"useAbsoluteRoot"`       // Whether the checks looks at the absolute root of a repo or the relative root (the directory specified when attached a repo to a service).
}

func (client *Client) CreateCheckRepositoryFile(input CheckRepositoryFileCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryFileCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRepositoryFile(input CheckRepositoryFileUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryFileUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}
