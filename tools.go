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
	DisplayName  string       `json:"displayName" yaml:"displayName"`
	Url          string       `json:"url"`
	Environment  string       `json:"environment,omitempty" yaml:"environment,omitempty"`
	ServiceId    ID           `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
	ServiceAlias string       `json:"serviceAlias,omitempty" yaml:"serviceId,omitempty"`
}

func (t *ToolCreateInput) Example() ToolCreateInput {
	return ToolCreateInput{
		Category:     ToolCategoryLogs,
		DisplayName:  ExampleName,
		Url:          ExampleUrl,
		Environment:  ExampleEnvironment,
		ServiceId:    *NewID(ExampleId),
		ServiceAlias: ExampleAlias,
	}
}

func (t *ToolCreateInput) ExampleJson() string {
	return GenJsonFrom[ToolCreateInput](t.Example())
}

func (t *ToolCreateInput) ExampleYaml() string {
	return GenYamlFrom[ToolCreateInput](t.Example())
}

type ToolUpdateInput struct {
	Id          ID           `json:"id"`
	Category    ToolCategory `json:"category,omitempty"`
	DisplayName string       `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Url         string       `json:"url,omitempty" yaml:"url,omitempty"`
	Environment string       `json:"environment,omitempty" yaml:"environment,omitempty"`
}

func (t *ToolUpdateInput) Example() ToolUpdateInput {
	return ToolUpdateInput{
		Id:          ExampleId,
		Category:    ToolCategoryLogs,
		DisplayName: ExampleName,
		Url:         ExampleUrl,
		Environment: ExampleEnvironment,
	}
}

func (t *ToolUpdateInput) ExampleJson() string {
	return GenJsonFrom[ToolUpdateInput](t.Example())
}

func (t *ToolUpdateInput) ExampleYaml() string {
	return GenYamlFrom[ToolUpdateInput](t.Example())
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
