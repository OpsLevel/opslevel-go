package opslevel

type Component Service

type ComponentCreateInput ServiceCreateInput

type ComponentUpdateInput ServiceUpdateInput

type ComponentConnection ServiceConnection

type ComponentTypeConnection struct {
	Nodes      []ComponentType `json:"nodes"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount" graphql:"-"`
}

func (client *Client) CreateComponentType(input ComponentTypeInput) (*ComponentType, error) {
	var m struct {
		Payload ComponentTypePayload `graphql:"componentTypeCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ComponentTypeCreate"))
	return &m.Payload.ComponentType, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateComponent(input ComponentCreateInput) (*Component, error) {
	resource, err := client.CreateService(ServiceCreateInput(input))
	return any(resource).(*Component), err
}

func (client *Client) GetComponentType(identifier string) (*ComponentType, error) {
	var q struct {
		Account struct {
			ComponentType ComponentType `graphql:"componentType(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Query(&q, v, WithName("ComponentTypeGet"))
	return &q.Account.ComponentType, HandleErrors(err, nil)
}

func (client *Client) GetComponent(identifier string) (*Component, error) {
	resource, err := client.GetService(identifier)
	return any(resource).(*Component), err
}

func (client *Client) ListComponentTypes(variables *PayloadVariables) (*ComponentTypeConnection, error) {
	var q struct {
		Account struct {
			ComponentTypes ComponentTypeConnection `graphql:"componentTypes(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("ComponentTypeList")); err != nil {
		return nil, err
	}
	for q.Account.ComponentTypes.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.ComponentTypes.PageInfo.End
		resp, err := client.ListComponentTypes(variables)
		if err != nil {
			return nil, err
		}
		q.Account.ComponentTypes.Nodes = append(q.Account.ComponentTypes.Nodes, resp.Nodes...)
		q.Account.ComponentTypes.PageInfo = resp.PageInfo
	}
	q.Account.ComponentTypes.TotalCount = len(q.Account.ComponentTypes.Nodes)
	return &q.Account.ComponentTypes, nil
}

func (client *Client) ListComponents(variables *PayloadVariables) (*ComponentConnection, error) {
	resource, err := client.ListServices(variables)
	return any(resource).(*ComponentConnection), err
}

func (client *Client) UpdateComponentType(identifier string, input ComponentTypeInput) (*ComponentType, error) {
	var m struct {
		Payload ComponentTypePayload `graphql:"componentTypeUpdate(componentType:$target,input:$input)"`
	}
	v := PayloadVariables{
		"target": *NewIdentifier(identifier),
		"input":  input,
	}
	err := client.Mutate(&m, v, WithName("ComponentTypeUpdate"))
	return &m.Payload.ComponentType, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateComponent(input ComponentUpdateInput) (*Component, error) {
	resource, err := client.UpdateService(ServiceUpdateInput(input))
	return any(resource).(*Component), err
}

func (client *Client) DeleteComponentType(identifier string) error {
	var d struct {
		Payload BasePayload `graphql:"componentTypeDelete(resource:$target)"`
	}
	v := PayloadVariables{
		"target": *NewIdentifier(identifier),
	}
	err := client.Mutate(&d, v, WithName("ComponentTypeDelete"))
	return HandleErrors(err, d.Payload.Errors)
}

func (client *Client) DeleteComponent(identifier string) error {
	return client.DeleteService(identifier)
}
