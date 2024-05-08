package opslevel

import (
	"fmt"
	"slices"

	"github.com/relvacode/iso8601"
)

type Language struct {
	Name  string
	Usage float64
}

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

type RepositoryPath struct {
	Href string
	Path string
}

type ServiceRepository struct {
	BaseDirectory string
	DisplayName   string
	Id            ID
	Repository    RepositoryId
	Service       ServiceId
}

type RepositoryConnection struct {
	HiddenCount       int
	Nodes             []Repository
	OrganizationCount int
	OwnedCount        int
	PageInfo          PageInfo
	TotalCount        int
	VisibleCount      int
}

type RepositoryServiceEdge struct {
	AtRoot              bool
	Node                ServiceId
	Paths               []RepositoryPath
	ServiceRepositories []ServiceRepository
}

type RepositoryServiceConnection struct {
	Edges      []RepositoryServiceEdge
	PageInfo   PageInfo
	TotalCount int
}

type ServiceRepositoryEdge struct {
	Node                RepositoryId
	ServiceRepositories []ServiceRepository
}

type ServiceRepositoryConnection struct {
	Edges      []ServiceRepositoryEdge
	PageInfo   PageInfo
	TotalCount int
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
		_, err := repository.GetServices(client, variables)
		if err != nil {
			return err
		}
	}

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
	if repository.Services == nil {
		repository.Services = &RepositoryServiceConnection{}
	}
	repository.Services.Edges = append(repository.Services.Edges, q.Account.Repository.Services.Edges...)
	repository.Services.PageInfo = q.Account.Repository.Services.PageInfo
	repository.Services.TotalCount += q.Account.Repository.Services.TotalCount
	for repository.Services.PageInfo.HasNextPage {
		(*variables)["after"] = repository.Services.PageInfo.End
		_, err := repository.GetServices(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return repository.Services, nil
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
	repository.Tags.TotalCount += q.Account.Repository.Tags.TotalCount
	for repository.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = repository.Tags.PageInfo.End
		_, err := repository.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return repository.Tags, nil
}

func (repository *Repository) GetTeam(client *Client) (*Team, error) {
	if repository.Owner.Id == "" {
		return nil, nil
	}
	return client.GetTeam(repository.Owner.Id)
}

func (repository *Repository) GetTeamId() TeamId {
	return repository.Owner
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
		Payload struct {
			ServiceRepository ServiceRepository
			Errors            []OpsLevelErrors
		} `graphql:"serviceRepositoryCreate(input: $input)"`
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
	}
	if err := client.Query(&q, *variables, WithName("RepositoryList")); err != nil {
		return &q.Account.Repositories, err
	}
	for q.Account.Repositories.PageInfo.HasNextPage {
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
		q.Account.Repositories.TotalCount += resp.TotalCount
	}
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
	for q.Account.Repositories.PageInfo.HasNextPage {
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
		q.Account.Repositories.TotalCount += resp.TotalCount
	}
	return &q.Account.Repositories, nil
}

func (client *Client) UpdateRepository(input RepositoryUpdateInput) (*Repository, error) {
	var m struct {
		Payload struct {
			Repository Repository
			Errors     []OpsLevelErrors
		} `graphql:"repositoryUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RepositoryUpdate"))
	return &m.Payload.Repository, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateServiceRepository(input ServiceRepositoryUpdateInput) (*ServiceRepository, error) {
	var m struct {
		Payload struct {
			ServiceRepository ServiceRepository
			Errors            []OpsLevelErrors
		} `graphql:"serviceRepositoryUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ServiceRepositoryUpdate"))
	return &m.Payload.ServiceRepository, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteServiceRepository(id ID) error {
	var m struct {
		Payload struct {
			Id     ID `graphql:"deletedId"`
			Errors []OpsLevelErrors
		} `graphql:"serviceRepositoryDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("ServiceRepositoryDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
