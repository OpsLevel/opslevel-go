package opslevel

type SystemId Identifier

type System struct {
	SystemId
	Name        string      `graphql:"name"`
	Description string      `graphql:"description"`
	HTMLUrl     string      `graphql:"htmlUrl"`
	Owner       EntityOwner `graphql:"owner"`
	Parent      Domain      `graphql:"parent"`
}

type SystemConnection struct {
	Nodes      []System `json:"nodes"`
	PageInfo   PageInfo `json:"pageInfo"`
	TotalCount int      `json:"totalCount"`
}

type SystemInput struct {
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Owner       *ID              `json:"ownerId,omitempty"`
	Parent      *IdentifierInput `json:"parent,omitempty"`
}

func (s *SystemId) ChildServices(client *Client, variables *PayloadVariables) (*ServiceConnection, error) {
	return &ServiceConnection{}, nil
}

func (s *SystemId) Tags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	return &TagConnection{}, nil
}

func (s *SystemId) AssignService(client *Client, services ...string) error {
	return nil
}

func (c *Client) CreateSystem(input SystemInput) (*System, error) {
	return &System{}, nil
}

func (c *Client) GetSystem(identifier string) (*System, error) {
	return &System{}, nil
}

func (c *Client) ListSystems(variables *PayloadVariables) (*SystemConnection, error) {
	var q struct {
		Account struct {
			Systems SystemConnection `graphql:"systems(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = c.InitialPageVariablesPointer()
	}
	if err := c.Query(&q, *variables, WithName("SystemsList")); err != nil {
		return &SystemConnection{}, err
	}
	for q.Account.Systems.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Systems.PageInfo.End
		resp, err := c.ListSystems(variables)
		if err != nil {
			return &SystemConnection{}, err
		}
		q.Account.Systems.Nodes = append(q.Account.Systems.Nodes, resp.Nodes...)
		q.Account.Systems.PageInfo = resp.PageInfo
		q.Account.Systems.TotalCount += resp.TotalCount
	}
	return &q.Account.Systems, nil
}

func (c *Client) UpdateSystem(identifier string, input SystemInput) (*System, error) {
	return &System{}, nil
}

func (c *Client) DeleteSystem(identifier string) error {
	return nil
}
