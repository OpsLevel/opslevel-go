package opslevel

import "fmt"

type ServiceDependency struct {
	Id        ID        `graphql:"id"`
	Service   ServiceId `graphql:"sourceService"`
	DependsOn ServiceId `graphql:"destinationService"`
	Notes     string    `graphql:"notes"`
}

type ServiceDependenciesEdge struct {
	Id     ID         `graphql:"id"`
	Locked bool       `graphql:"locked"`
	Node   *ServiceId `graphql:"node"`
	Notes  string     `graphql:"notes"`
}

type ServiceDependentsEdge struct {
	Id     ID         `graphql:"id"`
	Locked bool       `graphql:"locked"`
	Node   *ServiceId `graphql:"node"`
	Notes  string     `graphql:"notes"`
}

type ServiceDependenciesConnection struct {
	Edges    []ServiceDependenciesEdge `graphql:"edges"`
	PageInfo PageInfo
}

type ServiceDependentsConnection struct {
	Edges    []ServiceDependentsEdge `graphql:"edges"`
	PageInfo PageInfo
}

func (client *Client) CreateServiceDependency(input ServiceDependencyCreateInput) (*ServiceDependency, error) {
	var m struct {
		Payload struct {
			ServiceDependency *ServiceDependency
			Errors            []OpsLevelErrors
		} `graphql:"serviceDependencyCreate(inputV2: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ServiceDependencyCreate"))
	return m.Payload.ServiceDependency, HandleErrors(err, m.Payload.Errors)
}

func (service *Service) GetDependencies(client *Client, variables *PayloadVariables) (*ServiceDependenciesConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Dependencies ServiceDependenciesConnection `graphql:"dependencies(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get Dependencies, invalid service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceDependenciesList")); err != nil {
		return nil, err
	}
	if service.Dependencies == nil {
		service.Dependencies = &ServiceDependenciesConnection{}
	}
	service.Dependencies.Edges = append(service.Dependencies.Edges, q.Account.Service.Dependencies.Edges...)
	service.Dependencies.PageInfo = q.Account.Service.Dependencies.PageInfo
	for service.Dependencies.PageInfo.HasNextPage {
		(*variables)["after"] = service.Dependencies.PageInfo.End
		_, err := service.GetDependencies(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return service.Dependencies, nil
}

func (service *Service) GetDependents(client *Client, variables *PayloadVariables) (*ServiceDependentsConnection, error) {
	var q struct {
		Account struct {
			Service struct {
				Dependents ServiceDependentsConnection `graphql:"dependents(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	if service.Id == "" {
		return nil, fmt.Errorf("unable to get Dependents, invalid service id: '%s'", service.Id)
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["service"] = service.Id
	if err := client.Query(&q, *variables, WithName("ServiceDependentsList")); err != nil {
		return nil, err
	}
	if service.Dependents == nil {
		service.Dependents = &ServiceDependentsConnection{}
	}
	service.Dependents.Edges = append(service.Dependents.Edges, q.Account.Service.Dependents.Edges...)
	service.Dependents.PageInfo = q.Account.Service.Dependents.PageInfo
	for service.Dependents.PageInfo.HasNextPage {
		(*variables)["after"] = service.Dependents.PageInfo.End
		_, err := service.GetDependents(client, variables)
		if err != nil {
			return nil, err
		}
	}
	return service.Dependents, nil
}

func (client *Client) DeleteServiceDependency(id ID) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"serviceDependencyDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("ServiceDependencyDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
