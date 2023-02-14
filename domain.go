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

type DomainInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       *ID    `json:"ownerId,omitempty"`
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

func (c *Client) CreateDomain(input DomainInput) (*Domain, error) {
	return &Domain{}, nil
}

func (c *Client) GetDomain(identifier string) (*Domain, error) {
	return &Domain{}, nil
}

func (c *Client) ListDomains(variables *PayloadVariables) (*DomainConnection, error) {
	return &DomainConnection{}, nil
}

func (c *Client) UpdateDomain(identifier string, input DomainInput) (*Domain, error) {
	return &Domain{}, nil
}

func (c *Client) DeleteDomain(identifier string) error {
	return nil
}
