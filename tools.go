package opslevel

import (
	"github.com/shurcooL/graphql"
)

type ToolCategory string

const (
	ToolCategoryAdmin                 ToolCategory = "admin"
	ToolCategoryCode                  ToolCategory = "code"
	ToolCategoryContinuousIntegration ToolCategory = "continuous_integration"
	ToolCategoryDeployment            ToolCategory = "deployment"
	ToolCategoryErrors                ToolCategory = "errors"
	ToolCategoryFeatureFlag           ToolCategory = "feature_flag"
	ToolCategoryHealthChecks          ToolCategory = "health_checks"
	ToolCategoryIncidents             ToolCategory = "incidents"
	ToolCategoryLogs                  ToolCategory = "logs"
	ToolCategoryMetrics               ToolCategory = "metrics"
	ToolCategoryOrchestrator          ToolCategory = "orchestrator"
	ToolCategoryRunbooks              ToolCategory = "runbooks"
	ToolCategoryStatusPage            ToolCategory = "status_page"
	ToolCategoryWiki                  ToolCategory = "wiki"
	ToolCategoryOther                 ToolCategory = "other"
)

func GetToolCategoryTypes() []string {
	return []string{
		string(ToolCategoryAdmin),
		string(ToolCategoryCode),
		string(ToolCategoryContinuousIntegration),
		string(ToolCategoryDeployment),
		string(ToolCategoryErrors),
		string(ToolCategoryFeatureFlag),
		string(ToolCategoryHealthChecks),
		string(ToolCategoryIncidents),
		string(ToolCategoryLogs),
		string(ToolCategoryMetrics),
		string(ToolCategoryOrchestrator),
		string(ToolCategoryRunbooks),
		string(ToolCategoryStatusPage),
		string(ToolCategoryWiki),
		string(ToolCategoryOther),
	}
}

type Tool struct {
	Category      ToolCategory
	CategoryAlias string `json:",omitempty"`
	DisplayName   string
	Environment   string     `json:",omitempty"`
	Id            graphql.ID `json:",omitempty"`
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
	ServiceId    graphql.ID   `json:"serviceId,omitempty"`
	ServiceAlias string       `json:"serviceAlias,omitempty"`
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
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Tool, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (conn *ToolConnection) Hydrate(service graphql.ID, client *Client) error {
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
func (client *Client) GetToolsForServiceWithId(service graphql.ID) ([]Tool, error) {
	return client.GetToolsForService(service)
}

func (client *Client) GetToolsForService(service graphql.ID) ([]Tool, error) {
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

func (client *Client) GetToolCount(service graphql.ID) (int, error) {
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
	if err := client.Query(&q, v); err != nil {
		return 0, err
	}
	return int(q.Account.Service.Tools.TotalCount), nil
}

//#endregion
