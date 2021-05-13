package opslevel

import (
	"github.com/shurcooL/graphql"
)

type Service struct {
	Aliases     []graphql.String `json:"aliases"`
	Description graphql.String   `json:"description"`
	Framework   graphql.String   `json:"framework"`
	Id          graphql.ID       `json:"id"`
	Language    graphql.String   `json:"language"`
	Lifecycle   Lifecycle        `json:"lifecycle"`
	Name        graphql.String   `json:"name"`
	Owner       Team             `json:"owner"`
	Product     graphql.String   `json:"product"`
	Tier        Tier             `json:"tier"`
}

type ServiceCreateInput struct {
	Name        string `json:"name"`
	Product     string `json:"product,omitempty"`
	Description string `json:"description,omitempty"`
	Language    string `json:"language,omitempty"`
	Framework   string `json:"framework,omitempty"`
	Tier        string `json:"tierAlias,omitempty"`
	Owner       string `json:"ownerAlias,omitempty"`
	Lifecycle   string `json:"lifecycleAlias,omitempty"`
}

type ServiceUpdateInput struct {
	Id           graphql.ID `json:"id,omitempty"`
	Alias        string     `json:"alias,omitempty"`
	Name         string     `json:"name,omitempty"`
	Product      string     `json:"product,omitempty"`
	Descripition string     `json:"description,omitempty"`
	Language     string     `json:"language,omitempty"`
	Framework    string     `json:"framework,omitempty"`
	Tier         string     `json:"tierAlias,omitempty"`
	Owner        string     `json:"ownerAlias,omitempty"`
	Lifecycle    string     `json:"lifecycleAlias,omitempty"`
}

type ServiceDeleteInput struct {
	Id    graphql.ID `json:"id,omitempty"`
	Alias string     `json:"alias,omitempty"`
}

//#region Create

func (client *Client) CreateService(input ServiceCreateInput) (*Service, error) {
	var m struct {
		Payload struct {
			Service Service
			Errors  []OpsLevelErrors
		} `graphql:"serviceCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Service, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetServiceWithAlias(alias string) (*Service, error) {
	var q struct {
		Account struct {
			Service Service `graphql:"service(alias: $service)"`
		}
	}
	v := PayloadVariables{
		"service": graphql.String(alias),
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	// TODO: if q.Account.Service.Tags.PageInfo.HasNextPage - Do Further Paginate Query?!
	return &q.Account.Service, nil
}

func (client *Client) GetServiceWithId(id graphql.ID) (*Service, error) {
	var q struct {
		Account struct {
			Service Service `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	// TODO: if q.Account.Service.Tags.PageInfo.HasNextPage - Do Further Paginate Query?!
	return &q.Account.Service, nil
}

func (client *Client) GetServiceCount() (int, error) {
	var q struct {
		Account struct {
			Services struct {
				TotalCount graphql.Int
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return 0, err
	}
	return int(q.Account.Services.TotalCount), nil
}

func (client *Client) ListServices() ([]Service, error) {
	var output []Service
	var q struct {
		Account struct {
			Services struct {
				Nodes    []Service
				PageInfo PageInfo
			} `graphql:"services(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"after": "",
		"first": client.pageSize,
	}
	if err := client.Query(&q, v); err != nil {
		return output, err
	}
	for _, item := range q.Account.Services.Nodes {
		output = append(output, item)
	}
	for q.Account.Services.PageInfo.HasNextPage {
		v["after"] = q.Account.Services.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return output, err
		}
		for _, item := range q.Account.Services.Nodes {
			output = append(output, item)
		}
	}
	return output, nil
}

//#endregion

//#region Update

func (client *Client) UpdateService(input ServiceUpdateInput) (*Service, error) {
	var m struct {
		Payload struct {
			Service Service
			Errors  []OpsLevelErrors
		} `graphql:"serviceUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Service, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteService(input ServiceDeleteInput) error {
	var m struct {
		Payload struct {
			Id     graphql.ID       `graphql:"deletedServiceId"`
			Alias  graphql.String   `graphql:"deletedServiceAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"serviceDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
