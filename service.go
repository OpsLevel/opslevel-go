package opslevel

import (
	"fmt"
	"strings"

	"github.com/hasura/go-graphql-client"
)

type ServiceId struct {
	Id      ID       `json:"id"`
	Aliases []string `json:"aliases,omitempty"`
}

type Service struct {
	ApiDocumentPath string `json:"apiDocumentPath,omitempty"`
	Description     string `json:"description,omitempty"`
	Framework       string `json:"framework,omitempty"`
	HtmlURL         string `json:"htmlUrl"`
	ServiceId
	Language                   string                       `json:"language,omitempty"`
	Lifecycle                  Lifecycle                    `json:"lifecycle,omitempty"`
	Name                       string                       `json:"name,omitempty"`
	Owner                      TeamId                       `json:"owner,omitempty"`
	PreferredApiDocument       *ServiceDocument             `json:"preferredApiDocument,omitempty"`
	PreferredApiDocumentSource *ApiDocumentSourceEnum       `json:"preferredApiDocumentSource,omitempty"`
	Product                    string                       `json:"product,omitempty"`
	Repositories               *ServiceRepositoryConnection `json:"repos,omitempty" graphql:"repos"`
	Tags                       *TagConnection               `json:"tags,omitempty"`
	Tier                       Tier                         `json:"tier,omitempty"`
	Timestamps                 Timestamps                   `json:"timestamps"`
	Tools                      *ToolConnection              `json:"tools,omitempty"`
}

type ServiceConnection struct {
	Nodes      []Service
	PageInfo   PageInfo
	TotalCount int
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
	Id          ID     `json:"id,omitempty"`
	Alias       string `json:"alias,omitempty"`
	Name        string `json:"name,omitempty"`
	Product     string `json:"product,omitempty"`
	Description string `json:"description,omitempty"`
	Language    string `json:"language,omitempty"`
	Framework   string `json:"framework,omitempty"`
	Tier        string `json:"tierAlias,omitempty"`
	Owner       string `json:"ownerAlias,omitempty"`
	Lifecycle   string `json:"lifecycleAlias,omitempty"`
}

type ServiceDeleteInput struct {
	Id    ID     `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}

//#region ServiceHelpers

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

func (s *Service) HasTool(category ToolCategory, name string, environment string) bool {
	for _, tool := range s.Tools.Nodes {
		if tool.Category == category && tool.DisplayName == name && tool.Environment == environment {
			return true
		}
	}
	return false
}

func (s *Service) Hydrate(client *Client) error {
	if s.Tags.PageInfo.HasNextPage {
		variables := &PayloadVariables{}
		(*variables)["after"] = s.Tags.PageInfo.End
		_, err := s.GetTags(client, variables)
		if err != nil {
			return err
		}
	}

	if s.Tools.PageInfo.HasNextPage {
		variables := &PayloadVariables{}
		(*variables)["after"] = s.Tools.PageInfo.End
		resp, err := s.GetTools(client, variables)
		if err != nil {
			return err
		}
		s.Tools = resp
	}

	if s.Repositories.PageInfo.HasNextPage {
		variables := &PayloadVariables{}
		(*variables)["after"] = s.Repositories.PageInfo.End
		resp, err := s.GetRepositories(client, variables)
		if err != nil {
			return err
		}
		s.Repositories = resp
	}
	return nil
}

func (s *Service) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if s.Id == "" {
		return nil, fmt.Errorf("Unable to get Tags, invalid service id: '%s'", s.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = s.Id
	if err := client.Query(&q, *variables, WithName("ServiceTagsList")); err != nil {
		return nil, err
	}
	if s.Tags == nil {
		tags := TagConnection{}
		s.Tags = &tags
	}
	s.Tags.Nodes = append(s.Tags.Nodes, q.Account.Service.Tags.Nodes...)
	s.Tags.PageInfo = q.Account.Service.Tags.PageInfo
	s.Tags.TotalCount += q.Account.Service.Tags.TotalCount
	for s.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = s.Tags.PageInfo.End
		_, err := s.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return s.Tags, nil
}

func (s *Service) GetTools(client *Client, variables *PayloadVariables) (*ToolConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Tools ToolConnection `graphql:"tools(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if s.Id == "" {
		return nil, fmt.Errorf("Unable to get Tools, invalid service id: '%s'", s.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = s.Id
	if err := client.Query(&q, *variables, WithName("ServiceToolsList")); err != nil {
		return nil, err
	}
	if s.Tools == nil {
		tools := ToolConnection{}
		s.Tools = &tools
	}
	s.Tools.Nodes = append(s.Tools.Nodes, q.Account.Service.Tools.Nodes...)
	s.Tools.PageInfo = q.Account.Service.Tools.PageInfo
	s.Tools.TotalCount += q.Account.Service.Tools.TotalCount
	for s.Tools.PageInfo.HasNextPage {
		(*variables)["after"] = s.Tools.PageInfo.End
		_, err := s.GetTools(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return s.Tools, nil
}

func (s *Service) GetRepositories(client *Client, variables *PayloadVariables) (*ServiceRepositoryConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Repositories ServiceRepositoryConnection `graphql:"repos(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if s.Id == "" {
		return nil, fmt.Errorf("Unable to get Repositories, invalid service id: '%s'", s.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = s.Id
	if err := client.Query(&q, *variables, WithName("ServiceRepositoriesList")); err != nil {
		return nil, err
	}
	if s.Repositories == nil {
		repositories := ServiceRepositoryConnection{}
		s.Repositories = &repositories
	}
	s.Repositories.Nodes = append(s.Repositories.Nodes, q.Account.Service.Repositories.Nodes...)
	s.Repositories.PageInfo = q.Account.Service.Repositories.PageInfo
	s.Repositories.TotalCount += q.Account.Service.Repositories.TotalCount
	for s.Repositories.PageInfo.HasNextPage {
		(*variables)["after"] = s.Repositories.PageInfo.End
		_, err := s.GetRepositories(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return s.Repositories, nil
}

func (s *Service) Documents(client *Client) ([]ServiceDocument, error) {
	var q struct {
		Account struct {
			Service struct {
				Documents struct {
					Nodes    []ServiceDocument
					PageInfo PageInfo
				} `graphql:"documents(after: $after, first: $first)"`
			} `graphql:"service(id: $id)"`
		}
	}
	if s.Id == "" {
		return nil, fmt.Errorf("unable to get 'Documents', invalid service id: '%s'", s.Id)
	}
	v := PayloadVariables{
		"id":    s.Id,
		"first": client.pageSize,
		"after": graphql.String(""),
	}
	var output []ServiceDocument
	if err := client.Query(&q, v, WithName("ServiceDocumentsList")); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Service.Documents.Nodes...)
	for q.Account.Service.Documents.PageInfo.HasNextPage {
		v["after"] = q.Account.Service.Documents.PageInfo.End
		if err := client.Query(&q, v, WithName("ServiceDocumentsList")); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Service.Documents.Nodes...)
	}
	return output, nil
}

//#endregion

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
	if err := client.Mutate(&m, v, WithName("ServiceCreate")); err != nil {
		return nil, err
	}
	if err := m.Payload.Service.Hydrate(client); err != nil {
		return &m.Payload.Service, err
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
		"service": alias,
	}
	err := client.Query(&q, v, WithName("ServiceGet"))
	return &q.Account.Service, HandleErrors(err, nil)
}

func (client *Client) GetServiceWithAlias(alias string) (*Service, error) {
	var q struct {
		Account struct {
			Service Service `graphql:"service(alias: $service)"`
		}
	}
	v := PayloadVariables{
		"service": alias,
	}
	if err := client.Query(&q, v, WithName("ServiceGet")); err != nil {
		return nil, err
	}
	if err := q.Account.Service.Hydrate(client); err != nil {
		return &q.Account.Service, err
	}
	return &q.Account.Service, nil
}

// Deprecated: Use GetService instead
func (client *Client) GetServiceWithId(id ID) (*Service, error) {
	return client.GetService(id)
}

func (client *Client) GetService(id ID) (*Service, error) {
	var q struct {
		Account struct {
			Service Service `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": id,
	}
	if err := client.Query(&q, v, WithName("ServiceGet")); err != nil {
		return nil, err
	}
	if err := q.Account.Service.Hydrate(client); err != nil {
		return &q.Account.Service, err
	}
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
	err := client.Query(&q, nil, WithName("ServiceCountGet"))
	return int(q.Account.Services.TotalCount), HandleErrors(err, nil)
}

// TODO: maybe we can find a way to merge ServiceConnection.Query & Hydrate
func (conn *ServiceConnection) Query(client *Client, q interface{}, v PayloadVariables) ([]Service, error) {
	if err := client.Query(q, v, WithName("ServiceList")); err != nil {
		return conn.Nodes, err
	}
	if err := conn.Hydrate(client); err != nil {
		return conn.Nodes, err
	}
	return conn.Nodes, nil
}

func (conn *ServiceConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Services.PageInfo = conn.PageInfo
	for q.Account.Services.PageInfo.HasNextPage {
		v["after"] = q.Account.Services.PageInfo.End
		if err := client.Query(&q, v, WithName("ServiceList")); err != nil {
			return err
		}
		for _, item := range q.Account.Services.Nodes {
			if err := (&item).Hydrate(client); err != nil {
				return err
			}
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (client *Client) ListServices() ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	return q.Account.Services.Query(client, &q, v)
}

func (client *Client) ListServicesWithFramework(framework string) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(framework: $framework, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["framework"] = framework
	return q.Account.Services.Query(client, &q, v)
}

func (client *Client) ListServicesWithLanguage(language string) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(language: $language, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["language"] = language
	return q.Account.Services.Query(client, &q, v)
}

func (client *Client) ListServicesWithLifecycle(lifecycle string) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(lifecycleAlias: $lifecycle, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["lifecycle"] = lifecycle
	return q.Account.Services.Query(client, &q, v)
}

func (client *Client) ListServicesWithOwner(owner string) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(ownerAlias: $owner, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["owner"] = owner
	return q.Account.Services.Query(client, &q, v)
}

func (client *Client) ListServicesWithProduct(product string) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(product: $product, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["product"] = product
	return q.Account.Services.Query(client, &q, v)
}

type TagArgs struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func NewTagArgs(tag string) TagArgs {
	kv := strings.Split(tag, ":")
	switch len(kv) {
	case 1:
		return TagArgs{
			Key: kv[0],
		}
	case 2:
		return TagArgs{
			Key:   kv[0],
			Value: kv[1],
		}
	default: // TODO: is this the best we can do?
		return TagArgs{
			Key: tag,
		}
	}
}

func (client *Client) ListServicesWithTag(tag TagArgs) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(tag: $tag, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["tag"] = tag
	return q.Account.Services.Query(client, &q, v)
}

func (client *Client) ListServicesWithTier(tier string) ([]Service, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(tierAlias: $tier, after: $after, first: $first)"`
		}
	}
	v := client.InitialPageVariables()
	v["tier"] = graphql.String(tier)
	return q.Account.Services.Query(client, &q, v)
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
	if err := client.Mutate(&m, v, WithName("ServiceUpdate")); err != nil {
		return nil, err
	}
	if err := m.Payload.Service.Hydrate(client); err != nil {
		return &m.Payload.Service, err
	}
	return &m.Payload.Service, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

// TODO: we should have a method that takes and ID and that follows the convention of other delete functions
func (client *Client) DeleteService(input ServiceDeleteInput) error {
	var m struct {
		Payload struct {
			Id     ID               `graphql:"deletedServiceId"`
			Alias  graphql.String   `graphql:"deletedServiceAlias"`
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"serviceDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ServiceDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteServiceWithAlias(alias string) error {
	return client.DeleteService(ServiceDeleteInput{
		Alias: alias,
	})
}

//#endregion
