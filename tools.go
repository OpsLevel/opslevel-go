package opslevel

import (
	"github.com/hasura/go-graphql-client"
)

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
	err := client.Mutate(&m, v)
	return &m.Payload.Tool, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (conn *ToolConnection) Hydrate(service ID, client *Client) error {
	var q struct {
		Account struct {
			Service struct {
				Tools ToolConnection `graphql:"tools(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": service,
		"first":   client.pageSize,
	}
	q.Account.Service.Tools.PageInfo = conn.PageInfo
	for q.Account.Service.Tools.PageInfo.HasNextPage {
		v["after"] = q.Account.Service.Tools.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Service.Tools.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (client *Client) GetToolsForServiceWithAlias(alias string) ([]Tool, error) {
	service, serviceErr := client.GetServiceIdWithAlias(alias)
	if serviceErr != nil {
		return nil, serviceErr
	}
	return client.GetToolsForService(service.Id)
}

// Deprecated: Use GetToolsForService instead
func (client *Client) GetToolsForServiceWithId(service ID) ([]Tool, error) {
	return client.GetToolsForService(service)
}

func (client *Client) GetToolsForService(service ID) ([]Tool, error) {
	var q struct {
		Account struct {
			Service struct {
				Tools ToolConnection `graphql:"tools(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": service,
		"after":   graphql.String(""),
		"first":   client.pageSize,
	}
	if err := client.Query(&q, v); err != nil {
		return q.Account.Service.Tools.Nodes, err
	}
	if err := q.Account.Service.Tools.Hydrate(service, client); err != nil {
		return q.Account.Service.Tools.Nodes, err
	}
	return q.Account.Service.Tools.Nodes, nil
}

func (client *Client) GetToolCount(service ID) (int, error) {
	var q struct {
		Account struct {
			Service struct {
				Tools struct {
					TotalCount int
				}
			} `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": service,
	}
	err := client.Query(&q, v)
	return int(q.Account.Service.Tools.TotalCount), HandleErrors(err, nil)
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
	err := client.Mutate(&m, v)
	return &m.Payload.Tool, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteTool(id ID) error {
	var m struct {
		Payload struct {
			Id     ID `graphql:"deletedToolId"`
			Errors []OpsLevelErrors
		} `graphql:"toolDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": ToolDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
