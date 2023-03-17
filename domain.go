package opslevel

type EntityOwner struct {
	GroupId `graphql:"... on Group"`
	TeamId  `graphql:"... on Team"`
}

type DomainId Identifier

type Domain struct {
	DomainId
	Name        string      `graphql:"name"`
	Description string      `graphql:"description"`
	HTMLUrl     string      `graphql:"htmlUrl"`
	Owner       EntityOwner `graphql:"owner"`
}

type DomainConnection struct {
	Nodes      []Domain `json:"nodes"`
	PageInfo   PageInfo `json:"pageInfo"`
	TotalCount int      `json:"totalCount"`
}

type DomainCreateInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       *ID    `json:"ownerId,omitempty"`
	Note        string `json:"note,omitempty"`
}

type DomainUpdateInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       *ID    `json:"ownerId,omitempty"`
	Note        string `json:"note,omitempty"`
}

func (s *DomainId) ChildSystems(client *Client, variables *PayloadVariables) (*SystemConnection, error) {
	return &SystemConnection{}, nil
}

func (s *DomainId) Tags(client *Client, variables *PayloadVariables) (*TagConnection, error) {
	return &TagConnection{}, nil
}

func (s *DomainId) AssignSystem(client *Client, systems ...string) error {
	return nil
}

func (c *Client) CreateDomain(input DomainCreateInput) (*Domain, error) {
	var m struct {
		Payload struct {
			Domain Domain
			Errors []OpsLevelErrors
		} `graphql:"domainCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := c.Mutate(&m, v, WithName("DomainCreate"))
	return &m.Payload.Domain, HandleErrors(err, m.Payload.Errors)
}

func (c *Client) GetDomain(identifier string) (*Domain, error) {
	var q struct {
		Account struct {
			Domain Domain `graphql:"domain(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": NewIdentifier(identifier),
	}
	err := c.Query(&q, v, WithName("DomainGet"))
	return &q.Account.Domain, HandleErrors(err, nil)
}

func (c *Client) ListDomains(variables *PayloadVariables) (*DomainConnection, error) {
	var q struct {
		Account struct {
			Domains DomainConnection `graphql:"domains(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = c.InitialPageVariablesPointer()
	}
	if err := c.Query(&q, *variables, WithName("DomainsList")); err != nil {
		return &DomainConnection{}, err
	}
	for q.Account.Domains.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Domains.PageInfo.End
		resp, err := c.ListDomains(variables)
		if err != nil {
			return &DomainConnection{}, err
		}
		q.Account.Domains.Nodes = append(q.Account.Domains.Nodes, resp.Nodes...)
		q.Account.Domains.PageInfo = resp.PageInfo
		q.Account.Domains.TotalCount += resp.TotalCount
	}
	return &q.Account.Domains, nil
}

func (c *Client) UpdateDomain(identifier string, input DomainUpdateInput) (*Domain, error) {
	return &Domain{}, nil
}

func (c *Client) DeleteDomain(identifier string) error {
	var d struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"domainDelete(resource: $input)"`
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := c.Mutate(&d, v, WithName("DomainDelete"))
	return HandleErrors(err, d.Payload.Errors)
}
