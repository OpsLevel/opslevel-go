package opslevel

type RelationshipDefinitionPayload struct {
	Definition RelationshipDefinitionType // The relationship that was defined.
	BasePayload
}

func (client *Client) CreateRelationshipDefinition(input RelationshipDefinitionInput) (*RelationshipDefinitionType, error) {
	var m struct {
		Payload RelationshipDefinitionPayload `graphql:"relationshipDefinitionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipDefinitionCreate"))
	return &m.Payload.Definition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateRelationship(input RelationshipDefinition) (*RelationshipType, error) {
	var m struct {
		Payload struct {
			Relationship RelationshipType
			Errors       []Error
		} `graphql:"relationshipCreate(relationshipDefinition: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipCreate"))
	return &m.Payload.Relationship, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetRelationshipDefinition(identifier string) (*RelationshipDefinitionType, error) {
	var q struct {
		Account struct {
			Resource RelationshipDefinitionType `graphql:"relationshipDefinition(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(identifier),
	}
	err := client.Query(&q, v, WithName("RelationshipDefinitionGet"))
	return &q.Account.Resource, HandleErrors(err, nil)
}

func (client *Client) ListRelationshipDefinitions(variables *PayloadVariables) (*RelationshipDefinitionConnection, error) {
	var q struct {
		Account struct {
			Resources RelationshipDefinitionConnection `graphql:"relationshipDefinitions(after: $after, first: $first, componentType: $componentType, resource: $resource)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	} else {
		if (*variables)["first"] == nil {
			(*variables)["first"] = 100
		}
		if (*variables)["after"] == nil {
			(*variables)["after"] = ""
		}
	}
	if (*variables)["componentType"] == nil {
		(*variables)["componentType"] = &IdentifierInput{}
	}
	if (*variables)["resource"] == nil {
		(*variables)["resource"] = NewID("")
	}
	if err := client.Query(&q, *variables, WithName("RelationshipDefinitionList")); err != nil {
		return nil, err
	}
	if q.Account.Resources.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Resources.PageInfo.End
		resp, err := client.ListRelationshipDefinitions(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Resources.Nodes = append(q.Account.Resources.Nodes, resp.Nodes...)
		q.Account.Resources.PageInfo = resp.PageInfo
	}
	q.Account.Resources.TotalCount = len(q.Account.Resources.Nodes)
	return &q.Account.Resources, nil
}

func (client *Client) UpdateRelationshipDefinition(identifier string, input RelationshipDefinitionInput) (*RelationshipDefinitionType, error) {
	var m struct {
		Payload RelationshipDefinitionPayload `graphql:"relationshipDefinitionUpdate(relationshipDefinition: $identifier, input: $input)"`
	}
	v := PayloadVariables{
		"identifier": *NewIdentifier(identifier),
		"input":      input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipDefinitionUpdate"))
	return &m.Payload.Definition, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteRelationshipDefinition(identifier string) (*ID, error) {
	var m struct {
		Payload struct {
			DeletedId ID      `graphql:"deletedId"`
			Errors    []Error `graphql:"errors"`
		} `graphql:"relationshipDefinitionDelete(resource: $input)"`
	}
	input := *NewIdentifier(identifier)
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("RelationshipDefinitionDelete"))
	return &m.Payload.DeletedId, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteRelationship(identifier string) (*ID, error) {
	var m struct {
		Payload struct {
			DeletedId ID      `graphql:"deletedId"`
			Errors    []Error `graphql:"errors"`
		} `graphql:"relationshipDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: ID(identifier)},
	}
	err := client.Mutate(&m, v, WithName("RelationshipDelete"))
	return &m.Payload.DeletedId, HandleErrors(err, m.Payload.Errors)
}
