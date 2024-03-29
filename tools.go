package opslevel

type Tool struct {
	Category      ToolCategory
	CategoryAlias string `json:",omitempty"`
	DisplayName   string
	Environment   string `json:",omitempty"`
	Id            ID     `json:",omitempty"`
	Url           string
	Service       ServiceId
}

type ToolConnection struct {
	Nodes      []Tool
	PageInfo   PageInfo
	TotalCount int
}

func (client *Client) CreateTool(input ToolCreateInput) (*Tool, error) {
	var m struct {
		Payload struct {
			Tool   Tool
			Errors []OpsLevelErrors
		} `graphql:"toolCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ToolCreate"))
	return &m.Payload.Tool, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateTool(input ToolUpdateInput) (*Tool, error) {
	var m struct {
		Payload struct {
			Tool   Tool
			Errors []OpsLevelErrors
		} `graphql:"toolUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ToolUpdate"))
	return &m.Payload.Tool, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteTool(id ID) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"toolDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": ToolDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("ToolDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
