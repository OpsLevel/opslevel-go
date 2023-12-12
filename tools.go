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
	Category     ToolCategory `json:"category" validate:"required" default:"logs"`
	DisplayName  string       `json:"displayName" yaml:"displayName" default:"John Doe"`
	Url          string       `json:"url" default:"john.doe@example.com"`
	Environment  string       `json:"environment,omitempty" yaml:"environment,omitempty" default:"-"`
	ServiceId    ID           `json:"serviceId,omitempty" yaml:"serviceId,omitempty" default:"-"`
	ServiceAlias string       `json:"serviceAlias,omitempty" yaml:"serviceId,omitempty" default:"-"`
}

type ToolUpdateInput struct {
	Id          ID           `json:"id" default:"Z2lkOi8vMTIzNDU2Nzg5MTAK"`
	Category    ToolCategory `json:"category,omitempty" default:"-"`
	DisplayName string       `json:"displayName,omitempty" yaml:"displayName,omitempty" default:"John Doe"`
	Url         string       `json:"url,omitempty" yaml:"url,omitempty" default:"john.doe@example.com"`
	Environment string       `json:"environment,omitempty" yaml:"environment,omitempty" default:"-"`
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

// Deprecated: Use client.GetServiceWithAlias(alias).Tools instead
func (client *Client) GetToolsForServiceWithAlias(alias string) ([]Tool, error) {
	service, serviceErr := client.GetServiceWithAlias(alias)
	return service.Tools.Nodes, serviceErr
}

// Deprecated: Use GetToolsForService instead
func (client *Client) GetToolsForServiceWithId(id ID) ([]Tool, error) {
	return client.GetToolsForService(id, nil)
}

// Deprecated: Use client.GetService(id).Tools instead
func (client *Client) GetToolsForService(id ID, variables *PayloadVariables) ([]Tool, error) {
	service, serviceErr := client.GetService(id)
	return service.Tools.Nodes, serviceErr
}

// Deprecated: Use client.GetService(id).Tools.TotalCount instead
func (client *Client) GetToolCount(id ID) (int, error) {
	service, serviceErr := client.GetService(id)
	return service.Tools.TotalCount, serviceErr
}

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
