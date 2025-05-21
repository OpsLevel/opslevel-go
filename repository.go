package opslevel

import (
	"fmt"
	"slices"

	"github.com/relvacode/iso8601"
)

type RepositoryId struct {
	Id           ID
	DefaultAlias string
}

type Repository struct {
	ArchivedAt         iso8601.Time
	CreatedOn          iso8601.Time
	DefaultAlias       string
	DefaultBranch      string
	Description        string
	Forked             bool
	HtmlUrl            string
	Id                 ID
	Languages          []Language
	LastOwnerChangedAt iso8601.Time
	Locked             bool
	Name               string
	Organization       string
	Owner              TeamId
	Private            bool
	RepoKey            string
	Services           *RepositoryServiceConnection
	Tags               *TagConnection
	Tier               Tier
	Type               string
	Url                string
	Visible            bool
}

// RepositoryConnection The connection type for Repository
type RepositoryConnection struct {
	HiddenCount       int
	Nodes             []Repository
	OrganizationCount int
	OwnedCount        int
	PageInfo          PageInfo
	TotalCount        int `graphql:"-"`
	VisibleCount      int
}

type RepositoryServiceEdge struct {
	AtRoot              bool
	Node                ServiceId
	Paths               []RepositoryPath
	ServiceRepositories []ServiceRepository
}

// RepositoryServiceConnection The connection type for Service
type RepositoryServiceConnection struct {
	Edges      []RepositoryServiceEdge
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

type ServiceRepositoryEdge struct {
	Node                RepositoryId
	ServiceRepositories []ServiceRepository
}

// ServiceRepositoryConnection The connection type for Repository
type ServiceRepositoryConnection struct {
	Edges      []ServiceRepositoryEdge
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

func (repository *Repository) ResourceId() ID {
	return repository.Id
}

func (repository *Repository) ResourceType() TaggableResource {
	return TaggableResourceRepository
}

func (repository *Repository) GetService(service ID, directory string) *ServiceRepository {
	for _, edge := range repository.Services.Edges {
		for _, connection := range edge.ServiceRepositories {
			if connection.Service.Id == service && connection.BaseDirectory == directory {
				return &connection
			}
		}
	}
	return nil
}

func (repository *Repository) Hydrate(client *Client) error {
	if repository.Services == nil {
		repository.Services = &RepositoryServiceConnection{}
	}
	if repository.Services.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = repository.Services.PageInfo.End
		resp, err := repository.GetServices(client, variables)
		if err != nil {
			return err
		}
		repository.Services = resp
	}
	repository.Services.TotalCount = len(repository.Services.Edges)

	if repository.Tags == nil {
		repository.Tags = &TagConnection{}
	}
	if repository.Tags.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = repository.Tags.PageInfo.End
		_, err := repository.GetTags(client, variables)
		if err != nil {
			return err
		}
	}
	repository.Tags.TotalCount = len(repository.Tags.Nodes)
	return nil
}

func (repository *Repository) GetServices(client *Client, variables *PayloadVariables) (*RepositoryServiceConnection, error) {
	var q struct {
		Account struct {
			Repository struct {
				Services RepositoryServiceConnection `graphql:"services(after: $after, first: $first)"`
			} `graphql:"repository(id: $id)"`
		}
	}
	if repository.Id == "" {
		return nil, fmt.Errorf("unable to get Services, invalid repository id: '%s'", repository.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["id"] = repository.Id
	if err := client.Query(&q, *variables, WithName("RepositoryServicesList")); err != nil {
		return nil, err
	}
	if q.Account.Repository.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Repository.Services.PageInfo.End
		resp, err := repository.GetServices(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Repository.Services.Edges = append(q.Account.Repository.Services.Edges, resp.Edges...)
		q.Account.Repository.Services.PageInfo = resp.PageInfo
	}
	q.Account.Repository.Services.TotalCount = len(q.Account.Repository.Services.Edges)
	return &q.Account.Repository.Services, nil
}

func (repository *Repository) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			Repository struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"repository(id: $id)"`
		}
	}
	if repository.Id == "" {
		return nil, fmt.Errorf("unable to get Tags, invalid repository id: '%s'", repository.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["id"] = repository.Id
	if err := client.Query(&q, *variables, WithName("RepositoryTagsList")); err != nil {
		return nil, err
	}
	if repository.Tags == nil {
		repository.Tags = &TagConnection{}
	}
	// Add unique tags only
	for _, tagNode := range q.Account.Repository.Tags.Nodes {
		if !slices.Contains[[]Tag, Tag](repository.Tags.Nodes, tagNode) {
			repository.Tags.Nodes = append(repository.Tags.Nodes, tagNode)
		}
	}
	repository.Tags.PageInfo = q.Account.Repository.Tags.PageInfo
	if repository.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = repository.Tags.PageInfo.End
		_, err := repository.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
	}
	repository.Tags.TotalCount = len(repository.Tags.Nodes)
	return repository.Tags, nil
}

func (client *Client) ConnectServiceRepository(service *ServiceId, repository *Repository) (*ServiceRepository, error) {
	input := ServiceRepositoryCreateInput{
		Service:       *NewIdentifier(string(service.Id)),
		Repository:    *NewIdentifier(string(repository.Id)),
		BaseDirectory: RefOf("/"),
		DisplayName:   RefOf(fmt.Sprintf("%s/%s", repository.Organization, repository.Name)),
	}
	return client.CreateServiceRepository(input)
}

func (client *Client) CreateServiceRepository(input ServiceRepositoryCreateInput) (*ServiceRepository, error) {
	var m struct {
		Payload ServiceRepositoryCreatePayload `graphql:"serviceRepositoryCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ServiceRepositoryCreate"))
	return &m.Payload.ServiceRepository, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetRepositoryWithAlias(alias string) (*Repository, error) {
	var q struct {
		Account struct {
			Repository Repository `graphql:"repository(alias: $repo)"`
		}
	}
	v := PayloadVariables{
		"repo": alias,
	}
	if err := client.Query(&q, v, WithName("RepositoryGet")); err != nil {
		return nil, err
	}
	if err := q.Account.Repository.Hydrate(client); err != nil {
		return &q.Account.Repository, err
	}
	return &q.Account.Repository, nil
}

func (client *Client) GetRepository(id ID) (*Repository, error) {
	var q struct {
		Account struct {
			Repository Repository `graphql:"repository(id: $repo)"`
		}
	}
	v := PayloadVariables{
		"repo": id,
	}
	if err := client.Query(&q, v, WithName("RepositoryGet")); err != nil {
		return nil, err
	}
	if err := q.Account.Repository.Hydrate(client); err != nil {
		return &q.Account.Repository, err
	}
	return &q.Account.Repository, nil
}

func (client *Client) ListRepositories(variables *PayloadVariables) (*RepositoryConnection, error) {
	var q struct {
		Account struct {
			Repositories RepositoryConnection `graphql:"repositories(after: $after, first: $first, visible: $visible)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
		(*variables)["visible"] = true
	}
	if err := client.Query(&q, *variables, WithName("RepositoryList")); err != nil {
		return &q.Account.Repositories, err
	}
	if q.Account.Repositories.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Repositories.PageInfo.End
		resp, err := client.ListRepositories(variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			err := node.Hydrate(client)
			if err != nil {
				return nil, err
			}
			q.Account.Repositories.Nodes = append(q.Account.Repositories.Nodes, node)
		}
		q.Account.Repositories.PageInfo = resp.PageInfo
	}
	q.Account.Repositories.TotalCount = len(q.Account.Repositories.Nodes)
	return &q.Account.Repositories, nil
}

func (client *Client) ListRepositoriesWithTier(tier string, variables *PayloadVariables) (*RepositoryConnection, error) {
	var q struct {
		Account struct {
			Repositories RepositoryConnection `graphql:"repositories(tierAlias: $tier, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["tier"] = tier
	if err := client.Query(&q, *variables, WithName("RepositoryListWithTier")); err != nil {
		return &q.Account.Repositories, err
	}
	if q.Account.Repositories.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Repositories.PageInfo.End
		resp, err := client.ListRepositoriesWithTier(tier, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			err := node.Hydrate(client)
			if err != nil {
				return nil, err
			}
			q.Account.Repositories.Nodes = append(q.Account.Repositories.Nodes, node)
		}
		q.Account.Repositories.PageInfo = resp.PageInfo
	}
	q.Account.Repositories.TotalCount = len(q.Account.Repositories.Nodes)
	return &q.Account.Repositories, nil
}

func (client *Client) UpdateRepository(input RepositoryUpdateInput) (*Repository, error) {
	var m struct {
		Payload RepositoryUpdatePayload `graphql:"repositoryUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RepositoryUpdate"))
	return &m.Payload.Repository, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateServiceRepository(input ServiceRepositoryUpdateInput) (*ServiceRepository, error) {
	var m struct {
		Payload ServiceRepositoryUpdatePayload `graphql:"serviceRepositoryUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ServiceRepositoryUpdate"))
	return &m.Payload.ServiceRepository, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteServiceRepository(id ID) error {
	var m struct {
		Payload struct { // TODO: fix this
			Id     ID `graphql:"deletedId"`
			Errors []Error
		} `graphql:"serviceRepositoryDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("ServiceRepositoryDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
