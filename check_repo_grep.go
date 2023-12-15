package opslevel

type RepositoryGrepCheckFragment struct {
	DirectorySearch       bool       `graphql:"directorySearch"`
	Filepaths             []string   `graphql:"filePaths"`
	FileContentsPredicate *Predicate `graphql:"fileContentsPredicate"`
}

type CheckRepositoryGrepCreateInput struct {
	CheckCreateInput

	DirectorySearch       bool            `json:"directorySearch" yaml:"directorySearch" default:"false"`
	Filepaths             []string        `json:"filePaths" yaml:"filePaths" default:"[\"**/hello.go\", \"src/**\"]"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty" yaml:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryGrepUpdateInput struct {
	CheckUpdateInput

	DirectorySearch       bool            `json:"directorySearch"`
	Filepaths             []string        `json:"filePaths,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
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
