package opslevel

// ComponentTypeId Information about a particular component type
type ComponentTypeId Identifier

// ComponentType Information about a particular component type
type ComponentType struct {
	ComponentTypeId
	Description string            // The description of the component type (Optional)
	Href        string            // The relative path to link to the component type (Required)
	Icon        ComponentTypeIcon // The icon associated with the component type (Required)
	IsDefault   bool              // Whether or not the component type is the default (Required)
	Name        string            // The name of the component type (Required)
	Timestamps  Timestamps        // When the component type was created and updated (Required)
}

// ComponentTypeIcon The icon for a component type
type ComponentTypeIcon struct {
	Color string // The color, represented as a hexcode, for the icon (Optional)
	Name  string // The name of the icon in Phosphor icons for Vue, e.g. `PhBird`. See https://phosphoricons.com/ for a full list (Optional)
}

type ComponentTypeConnection struct {
	Nodes      []ComponentType `json:"nodes"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount" graphql:"-"`
}

func (client *Client) CreateComponentType(input ComponentTypeInput) (*ComponentType, error) {
	var m struct {
		Payload struct {
			ComponentType ComponentType
			Errors        []OpsLevelErrors
		} `graphql:"componentTypeCreate(input:$input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ComponentTypeCreate"))
	return &m.Payload.ComponentType, HandleErrors(err, m.Payload.Errors)
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

func (client *Client) UpdateComponentType(identifier string, input ComponentTypeInput) (*ComponentType, error) {
	var m struct {
		Payload struct {
			ComponentType ComponentType
			Errors        []OpsLevelErrors
		} `graphql:"componentTypeUpdate(componentType:$target,input:$input)"`
	}
	v := PayloadVariables{
		"target": *NewIdentifier(identifier),
		"input":  input,
	}
	err := client.Mutate(&m, v, WithName("ComponentTypeUpdate"))
	return &m.Payload.ComponentType, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteComponentType(identifier string) error {
	var d struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"componentTypeDelete(resource:$target)"`
	}
	v := PayloadVariables{
		"target": *NewIdentifier(identifier),
	}
	err := client.Mutate(&d, v, WithName("ComponentTypeDelete"))
	return HandleErrors(err, d.Payload.Errors)
}
