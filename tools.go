package opslevel

type ToolConnection struct {
	Nodes      []Tool
	PageInfo   PageInfo
	TotalCount int
}

func (client *Client) CreateTool(input ToolCreateInput) (*Tool, error) {
	var m struct {
		Payload struct {
			Tool   Tool
			Errors []Error
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
			Errors []Error
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
			Errors []Error
		} `graphql:"toolDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": ToolDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("ToolDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
