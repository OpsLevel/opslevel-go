package opslevel

import (
	"github.com/shurcooL/graphql"
)

type ServiceId struct {
	Id graphql.ID `json:"id"`
}

type Service struct {
	Aliases     []string  `json:"aliases"`
	Description string    `json:"description"`
	Framework   string    `json:"framework"`
	Language    string    `json:"language"`
	Lifecycle   Lifecycle `json:"lifecycle"`
	Name        string    `json:"name"`
	Owner       Team      `json:"owner"`
	Product     string    `json:"product"`
	Tier        Tier      `json:"tier"`
	Tags        TagConnection
	Tools       ToolConnection
	ServiceId
}

type ServiceConnection struct {
	Nodes      []Service
	PageInfo   PageInfo
	TotalCount graphql.Int
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

func (s *Service) HasAlias(alias string) bool {
	for _, a := range s.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}

func (s *Service) HasTag(key string, value string) bool {
	for _, tag := range s.Tags.Nodes {
		if tag.Key == key && tag.Value == value {
			return true
		}
	}
	return false
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

// This is a lightweight api call to lookup a service id by and alias - it does not return a full Service object
func (client *Client) GetServiceIdWithAlias(alias string) (*ServiceId, error) {
	var q struct {
		Account struct {
			Service ServiceId `graphql:"service(alias: $service)"`
		}
	}
	v := PayloadVariables{
		"service": graphql.String(alias),
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.Service, nil
}

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

// Deprecated: Use GetService instead
func (client *Client) GetServiceWithId(id graphql.ID) (*Service, error) {
	return client.GetService(id)
}

func (client *Client) GetService(id graphql.ID) (*Service, error) {
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
			Services ServiceConnection `graphql:"services(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"after": graphql.String(""),
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
