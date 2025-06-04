package opslevel

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type ServiceId struct {
	Id      ID       `json:"id"`
	Aliases []string `json:"aliases,omitempty"`
}

// TODO: Lifecycle, TeamId, Tier should probably be pointers.
type Service struct {
	ApiDocumentPath string `json:"apiDocumentPath,omitempty"`
	Description     string `json:"description,omitempty"`
	Framework       string `json:"framework,omitempty"`
	HtmlURL         string `json:"htmlUrl"`
	ServiceId
	Language                   string                         `json:"language,omitempty"`
	Lifecycle                  Lifecycle                      `json:"lifecycle,omitempty"`
	Locked                     bool                           `json:"locked" graphql:"locked"`
	ManagedAliases             []string                       `json:"managedAliases,omitempty"`
	MaturityReport             MaturityReportOnlyOverallLevel `json:"maturityReport" graphql:"maturityReport"`
	Name                       string                         `json:"name,omitempty"`
	Note                       string                         `json:"note,omitempty"`
	Owner                      TeamId                         `json:"owner,omitempty"`
	Parent                     *SystemId                      `json:"parent,omitempty" graphql:"parent"`
	PreferredApiDocument       *ServiceDocument               `json:"preferredApiDocument,omitempty"`
	PreferredApiDocumentSource *ApiDocumentSourceEnum         `json:"preferredApiDocumentSource,omitempty"`
	Product                    string                         `json:"product,omitempty"`
	Repositories               *ServiceRepositoryConnection   `json:"repos,omitempty" graphql:"repos"`
	Repository                 *ServiceRepository             `graphql:"defaultServiceRepository" json:"defaultServiceRepository"`
	Tags                       *TagConnection                 `json:"tags,omitempty"`
	Tier                       Tier                           `json:"tier,omitempty"`
	Timestamps                 Timestamps                     `json:"timestamps"`
	Tools                      *ToolConnection                `json:"tools,omitempty"`
	Type                       *ComponentTypeId               `json:"type,omitempty"`

	Dependencies *ServiceDependenciesConnection `graphql:"-"`
	Dependents   *ServiceDependentsConnection   `graphql:"-"`

	LastDeploy *Deploy                      `graphql:"-"`
	Properties *ServicePropertiesConnection `graphql:"-"`
}

// Returns unique identifiers created by OpsLevel, values in Aliases but not ManagedAliases
func (service *Service) UniqueIdentifiers() []string {
	uniqueIdentifiers := []string{}
	for _, alias := range service.Aliases {
		if !slices.Contains(service.ManagedAliases, alias) {
			uniqueIdentifiers = append(uniqueIdentifiers, alias)
		}
	}

	return uniqueIdentifiers
}

func (service *Service) ReconcileAliases(client *Client, aliasesWanted []string) error {
	aliasesToCreate, aliasesToDelete := extractAliases(service.Aliases, aliasesWanted)

	// reconcile wanted aliases with actual aliases
	deleteErr := client.DeleteAliases(AliasOwnerTypeEnumService, aliasesToDelete)
	_, createErr := client.CreateAliases(service.Id, aliasesToCreate)

	// update service to reflect API updates
	updatedService, getErr := client.GetServiceId(service.Id)
	if updatedService != nil {
		service.Aliases = updatedService.Aliases
		service.ManagedAliases = updatedService.ManagedAliases
	}

	return errors.Join(deleteErr, createErr, getErr)
}

func (service *Service) ResourceId() ID {
	return service.Id
}

func (service *Service) ResourceType() TaggableResource {
	return TaggableResourceService
}

func (service *Service) AliasableType() AliasOwnerTypeEnum {
	return AliasOwnerTypeEnumService
}

func (service *Service) HasAlias(alias string) bool {
	for _, a := range service.Aliases {
		if a == alias {
			return true
		}
	}
	return false
}

func (service *Service) HasTag(key string, value string) bool {
	for _, tag := range service.Tags.Nodes {
		if tag.Key == key && tag.Value == value {
			return true
		}
	}
	return false
}

func (service *Service) HasTool(category ToolCategory, name string, environment string) bool {
	for _, tool := range service.Tools.Nodes {
		if tool.Category == category && tool.DisplayName == name && tool.Environment == environment {
			return true
		}
	}
	return false
}

func (service *Service) Hydrate(client *Client) error {
	if service.Tags == nil {
		service.Tags = &TagConnection{}
	}
	if service.Tags.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = service.Tags.PageInfo.End
		resp, err := service.GetTags(client, variables)
		if err != nil {
			return err
		}
		service.Tags = resp
	}
	service.Tags.TotalCount = len(service.Tags.Nodes)

	if service.Tools == nil {
		service.Tools = &ToolConnection{}
	}
	if service.Tools.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = service.Tools.PageInfo.End
		resp, err := service.GetTools(client, variables)
		if err != nil {
			return err
		}
		service.Tools = resp
	}
	service.Tools.TotalCount = len(service.Tools.Nodes)

	if service.Repositories == nil {
		service.Repositories = &ServiceRepositoryConnection{}
	}
	if service.Repositories.PageInfo.HasNextPage {
		variables := client.InitialPageVariablesPointer()
		(*variables)["after"] = service.Repositories.PageInfo.End
		resp, err := service.GetRepositories(client, variables)
		if err != nil {
			return err
		}
		service.Repositories = resp
	}
	service.Repositories.TotalCount = len(service.Repositories.Edges)

	return nil
}

func (service *Service) GetSystem(client *Client, variables *PayloadVariables) (*System, error) {
	var q struct {
		Account struct {
			Service struct {
				System System `graphql:"system"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get system, invalid Service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceSystemGet")); err != nil {
		return nil, err
	}
	return &q.Account.Service.System, nil
}

func (service *Service) GetTags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get Tags, invalid service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceTagsList")); err != nil {
		return nil, err
	}
	if service.Tags == nil {
		service.Tags = &TagConnection{}
	}
	// Add unique tags only
	for _, resp := range q.Account.Service.Tags.Nodes {
		if !slices.Contains[[]Tag, Tag](service.Tags.Nodes, resp) {
			service.Tags.Nodes = append(service.Tags.Nodes, resp)
		}
	}
	service.Tags.PageInfo = q.Account.Service.Tags.PageInfo
	if service.Tags.PageInfo.HasNextPage {
		(*variables)["after"] = service.Tags.PageInfo.End
		_, err := service.GetTags(client, variables)
		if err != nil {
			return nil, err
		}
	}
	service.Tags.TotalCount = len(service.Tags.Nodes)
	return service.Tags, nil
}

func (service *Service) GetLastDeploy(client *Client, variables *PayloadVariables) (*Deploy, error) {
	var q struct {
		Account struct {
			Service struct {
				LastDeploy Deploy
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get LastDeploy, invalid service id: '%s'", service.Id)
	}

	if variables == nil {
		variables = &PayloadVariables{}
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceLastDeploy")); err != nil {
		return nil, err
	}
	return &q.Account.Service.LastDeploy, nil
}

func (service *Service) GetAliases() []string {
	return service.Aliases
}

func (service *Service) GetTools(client *Client, variables *PayloadVariables) (*ToolConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Tools ToolConnection `graphql:"tools(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get Tools, invalid service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceToolsList")); err != nil {
		return nil, err
	}
	if service.Tools == nil {
		tools := ToolConnection{}
		service.Tools = &tools
	}
	service.Tools.Nodes = append(service.Tools.Nodes, q.Account.Service.Tools.Nodes...)
	service.Tools.PageInfo = q.Account.Service.Tools.PageInfo
	if service.Tools.PageInfo.HasNextPage {
		(*variables)["after"] = service.Tools.PageInfo.End
		_, err := service.GetTools(client, variables)
		if err != nil {
			return nil, err
		}
	}
	service.Tools.TotalCount = len(service.Tools.Nodes)
	return service.Tools, nil
}

func (service *Service) GetRepositories(client *Client, variables *PayloadVariables) (*ServiceRepositoryConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Repositories ServiceRepositoryConnection `graphql:"repos(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get Repositories, invalid service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceRepositoriesList")); err != nil {
		return nil, err
	}
	if q.Account.Service.Repositories.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Service.Repositories.PageInfo.End
		resp, err := service.GetRepositories(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Service.Repositories.Edges = append(q.Account.Service.Repositories.Edges, resp.Edges...)
		q.Account.Service.Repositories.PageInfo = resp.PageInfo
	}
	q.Account.Service.Repositories.TotalCount = len(q.Account.Service.Repositories.Edges)
	return &q.Account.Service.Repositories, nil
}

func (service *Service) GetDocuments(client *Client, variables *PayloadVariables) (*ServiceDocumentConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Documents ServiceDocumentConnection `graphql:"documents(searchTerm: $searchTerm, after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get 'Documents', invalid service id: '%s'", service.Id)
	}

	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if (*variables)["searchTerm"] == nil {
		(*variables)["searchTerm"] = ""
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceDocumentsList")); err != nil {
		return nil, err
	}
	if q.Account.Service.Documents.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Service.Documents.PageInfo.End
		resp, err := service.GetDocuments(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Service.Documents.Nodes = append(q.Account.Service.Documents.Nodes, resp.Nodes...)
		q.Account.Service.Documents.PageInfo = resp.PageInfo
	}
	q.Account.Service.Documents.TotalCount = len(q.Account.Service.Documents.Nodes)
	return &q.Account.Service.Documents, nil
}

func (service *Service) GetServiceStats(client *Client) (*ServiceStats, error) {
	var q struct {
		Account struct {
			Service struct {
				ServiceStats ServiceStats
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get 'ServiceStats', invalid service id: '%s'", service.Id)
	}

	variables := &PayloadVariables{}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("GetServiceStats")); err != nil {
		return nil, err
	}
	return &q.Account.Service.ServiceStats, nil
}

func (client *Client) CreateService(input ServiceCreateInput) (*Service, error) {
	var m struct {
		Payload ServiceCreatePayload `graphql:"serviceCreate(input: $input)"`
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
	return &m.Payload.Service, HandleErrors(m.Payload.Errors)
}

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

func (client *Client) GetServiceId(id ID) (*Service, error) {
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

func (client *Client) GetService(identifier string) (*Service, error) {
	if IsID(identifier) {
		return client.GetServiceId(ID(identifier))
	} else {
		return client.GetServiceWithAlias(identifier)
	}
}

func (client *Client) GetServiceCount() (int, error) {
	var q struct {
		Account struct {
			Services struct {
				TotalCount int
			}
		}
	}
	err := client.Query(&q, nil, WithName("ServiceCountGet"))
	return q.Account.Services.TotalCount, HandleErrors(err, nil)
}

func (client *Client) ListServices(variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}

	if err := client.Query(&q, *variables, WithName("ServiceList")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServices(variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithFilter(filterIdentifier string, variables *PayloadVariables) (*ServiceConnection, error) {
	if !IsID(filterIdentifier) {
		return nil, fmt.Errorf("filterId must be an ID. Given: '%s'", filterIdentifier)
	}
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(filterIdentifier: $filter, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["filter"] = NewIdentifier(filterIdentifier)

	if err := client.Query(&q, *variables, WithName("ServiceListWithFilter")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithFilter(filterIdentifier, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithFramework(framework string, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(framework: $framework, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["framework"] = framework

	if err := client.Query(&q, *variables, WithName("ServiceListWithFramework")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithFramework(framework, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithLanguage(language string, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(language: $language, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["language"] = language

	if err := client.Query(&q, *variables, WithName("ServiceListWithLanguage")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithLanguage(language, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithLifecycle(lifecycle string, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(lifecycleAlias: $lifecycle, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["lifecycle"] = lifecycle

	if err := client.Query(&q, *variables, WithName("ServiceListWithLifecycle")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithLifecycle(lifecycle, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithOwner(owner string, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(ownerAlias: $owner, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["owner"] = owner

	if err := client.Query(&q, *variables, WithName("ServiceListWithOwner")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithOwner(owner, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithProduct(product string, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(product: $product, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["product"] = product

	if err := client.Query(&q, *variables, WithName("ServiceListWithProduct")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithProduct(product, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func NewTagArgs(tag string) (TagArgs, error) {
	kv := strings.Split(tag, ":")
	switch len(kv) {
	case 1:
		return TagArgs{
			Key: RefOf(kv[0]),
		}, nil
	case 2:
		return TagArgs{
			Key:   RefOf(kv[0]),
			Value: RefOf(kv[1]),
		}, nil
	default:
		return TagArgs{}, fmt.Errorf("cannot make a valid TagArg from: '%s' (not in format key:value)", tag)
	}
}

func (client *Client) ListServicesWithTag(tag TagArgs, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(tag: $tag, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["tag"] = tag

	if err := client.Query(&q, *variables, WithName("ServiceListWithTag")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithTag(tag, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) ListServicesWithTier(tier string, variables *PayloadVariables) (*ServiceConnection, error) {
	var q struct {
		Account struct {
			Services ServiceConnection `graphql:"services(tierAlias: $tier, after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["tier"] = tier

	if err := client.Query(&q, *variables, WithName("ServiceListWithTier")); err != nil {
		return nil, err
	}

	if q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesWithTier(tier, variables)
		if err != nil {
			return nil, err
		}
		for _, node := range resp.Nodes {
			if err := node.Hydrate(client); err != nil {
				return nil, err
			}
			q.Account.Services.Nodes = append(q.Account.Services.Nodes, node)
		}
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}

func (client *Client) UpdateService(input ServiceUpdateInput) (*Service, error) {
	var m struct {
		Payload ServiceUpdatePayload `graphql:"serviceUpdate(input: $input)"`
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
	return &m.Payload.Service, HandleErrors(m.Payload.Errors)
}

func (client *Client) UpdateServiceNote(input ServiceNoteUpdateInput) (*Service, error) {
	var m struct {
		Payload ServiceUpdatePayload `graphql:"serviceNoteUpdate(input: $input)"`
	}

	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v, WithName("ServiceUpdateNote")); err != nil {
		return nil, err
	}

	return &m.Payload.Service, HandleErrors(m.Payload.Errors)
}

func (client *Client) DeleteService(identifier string) error {
	input := ServiceDeleteInput{}
	if IsID(identifier) {
		input.Id = RefOf(ID(identifier))
	} else {
		input.Alias = RefOf(identifier)
	}

	var m struct {
		Payload struct { // TODO: fix this
			Id     ID      `graphql:"deletedServiceId"`
			Alias  string  `graphql:"deletedServiceAlias"`
			Errors []Error `graphql:"errors"`
		} `graphql:"serviceDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ServiceDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
