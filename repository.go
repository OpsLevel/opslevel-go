package opslevel

import (
	"fmt"

	"github.com/hasura/go-graphql-client"
	"github.com/relvacode/iso8601"
)

type Language struct {
	Name  string
	Usage float32
}

// Lightweight Repository struct used to make some API calls return less data
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
	Services           RepositoryServiceConnection
	Tags               RepositoryTagConnection
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
	TotalCount graphql.Int
}

type ServiceRepositoryEdge struct {
	Node                RepositoryId
	ServiceRepositories []ServiceRepository
}

type ServiceRepositoryConnection struct {
	Edges      []ServiceRepositoryEdge
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type RepositoryTagConnection struct {
	Nodes      []Tag
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type ServiceRepositoryCreateInput struct {
	Service       IdentifierInput `json:"service"`
	Repository    IdentifierInput `json:"repository"`
	BaseDirectory string          `json:"baseDirectory"`
	DisplayName   string          `json:"displayName,omitempty"`
}

type ServiceRepositoryUpdateInput struct {
	Id            ID     `json:"id"`
	BaseDirectory string `json:"baseDirectory,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
}

func (r *Repository) GetService(service ID, directory string) *ServiceRepository {
	for _, edge := range r.Services.Edges {
		for _, connection := range edge.ServiceRepositories {
			// TODO: this string cast is just for during the conversion so i don't have to change all of Service objects
			if string(connection.Service.Id) == string(service) && connection.BaseDirectory == directory {
				return &connection
			}
		}
	}
	return nil
}

func (r *Repository) Hydrate(client *Client) error {
	if err := r.Services.Hydrate(r.Id, client); err != nil {
		return err
	}
	if err := r.Tags.Hydrate(r.Id, client); err != nil {
		return err
	}
	return nil
}

//#region Create

func (client *Client) ConnectServiceRepository(service *ServiceId, repository *Repository) (*ServiceRepository, error) {
	input := ServiceRepositoryCreateInput{
		Service:       IdentifierInput{Id: ID(service.Id)}, // TODO: only temporary - can undo once we finally convert service objects
		Repository:    IdentifierInput{Id: repository.Id},
		BaseDirectory: "/",
		DisplayName:   fmt.Sprintf("%s/%s", repository.Organization, repository.Name),
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
	err := client.Mutate(&m, v)
	return &m.Payload.ServiceRepository, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetRepositoryWithAlias(alias string) (*Repository, error) {
	var q struct {
		Account struct {
			Repository Repository `graphql:"repository(alias: $repo)"`
		}
	}
	v := PayloadVariables{
		"repo": graphql.String(alias),
	}
	if err := client.Query(&q, v); err != nil {
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
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if err := q.Account.Repository.Hydrate(client); err != nil {
		return &q.Account.Repository, err
	}
	return &q.Account.Repository, nil
}

func (conn *RepositoryConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Repositories RepositoryConnection `graphql:"repositories(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Repositories.PageInfo = conn.PageInfo
	for q.Account.Repositories.PageInfo.HasNextPage {
		v["after"] = q.Account.Repositories.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Repositories.Nodes {
			if err := (&item).Hydrate(client); err != nil {
				return err
			}
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (conn *RepositoryServiceConnection) Hydrate(id ID, client *Client) error {
	var q struct {
		Account struct {
			Repository struct {
				Services RepositoryServiceConnection `graphql:"services(after: $after, first: $first)"`
			} `graphql:"repository(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id":    id,
		"first": client.pageSize,
	}
	q.Account.Repository.Services.PageInfo = conn.PageInfo
	for q.Account.Repository.Services.PageInfo.HasNextPage {
		v["after"] = q.Account.Repository.Services.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Repository.Services.Edges {
			conn.Edges = append(conn.Edges, item)
		}
	}
	return nil
}

func (conn *ServiceRepositoryConnection) Hydrate(id ID, client *Client) error {
	var q struct {
		Account struct {
			Service struct {
				Repos ServiceRepositoryConnection `graphql:"repos(after: $after, first: $first)"`
			} `graphql:"service(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id":    id,
		"first": client.pageSize,
	}
	q.Account.Service.Repos.PageInfo = conn.PageInfo
	for q.Account.Service.Repos.PageInfo.HasNextPage {
		v["after"] = q.Account.Service.Repos.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Service.Repos.Edges {
			conn.Edges = append(conn.Edges, item)
		}
	}
	return nil
}

func (conn *RepositoryTagConnection) Hydrate(id ID, client *Client) error {
	var q struct {
		Account struct {
			Repository struct {
				Tags RepositoryTagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"repository(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id":    id,
		"first": client.pageSize,
	}
	q.Account.Repository.Tags.PageInfo = conn.PageInfo
	for q.Account.Repository.Tags.PageInfo.HasNextPage {
		v["after"] = q.Account.Repository.Tags.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Repository.Tags.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (client *Client) ListRepositories() ([]Repository, error) {
	var q struct {
		Account struct {
			Repositories RepositoryConnection `graphql:"repositories(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
	}
	if err := client.Query(&q, v); err != nil {
		return q.Account.Repositories.Nodes, err
	}
	if err := q.Account.Repositories.Hydrate(client); err != nil {
		return q.Account.Repositories.Nodes, err
	}
	return q.Account.Repositories.Nodes, nil
}

func (client *Client) ListRepositoriesWithTier(tier string) ([]Repository, error) {
	var q struct {
		Account struct {
			Repositories RepositoryConnection `graphql:"repositories(tierAlias: $tier, after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
		"tier":  graphql.String(tier),
	}
	if err := client.Query(&q, v); err != nil {
		return q.Account.Repositories.Nodes, err
	}
	if err := q.Account.Repositories.Hydrate(client); err != nil {
		return q.Account.Repositories.Nodes, err
	}
	return q.Account.Repositories.Nodes, nil
}

//#endregion

//#region Update

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
	err := client.Mutate(&m, v)
	return &m.Payload.ServiceRepository, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

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
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
