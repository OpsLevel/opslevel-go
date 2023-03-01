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

type ToolCreateInput struct {
	Category     ToolCategory `json:"category"`
	DisplayName  string       `json:"displayName"`
	Url          string       `json:"url"`
	Environment  string       `json:"environment,omitempty"`
	ServiceId    ID           `json:"serviceId,omitempty"`
	ServiceAlias string       `json:"serviceAlias,omitempty"`
}

type ToolUpdateInput struct {
	Id          ID           `json:"id"`
	Category    ToolCategory `json:"category,omitempty"`
	DisplayName string       `json:"displayName,omitempty"`
	Url         string       `json:"url,omitempty"`
	Environment string       `json:"environment,omitempty"`
}

type ToolDeleteInput struct {
	Id ID `json:"id"`
}

//#region Create

func (client *Client) CreateTool(input ToolCreateInput) (*Tool, error) {
	// TODO: validate - Category, DisplayName & Url are non nil - or throw err
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

//#endregion

//#region Retrieve

//#endregion

//#region Update

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

//#endregion

//#region Delete

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

//#endregion
