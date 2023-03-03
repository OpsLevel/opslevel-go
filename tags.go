package opslevel

type TagOwner string

const (
	TagOwnerService    TagOwner = "Service"
	TagOwnerRepository TagOwner = "Repository"
)

type Tag struct {
	Id    ID     `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TagConnection struct {
	Nodes      []Tag
	PageInfo   PageInfo
	TotalCount int
}

type TagInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TagAssignInput struct {
	Id    ID               `json:"id,omitempty"`
	Alias string           `json:"alias,omitempty"`
	Type  TaggableResource `json:"type,omitempty"`
	Tags  []TagInput       `json:"tags"`
}

type TagCreateInput struct {
	Id    ID               `json:"id"`
	Alias string           `json:"alias,omitempty"`
	Type  TaggableResource `json:"type,omitempty"`
	Key   string           `json:"key"`
	Value string           `json:"value"`
}

type TagUpdateInput struct {
	Id    ID     `json:"id"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type TagDeleteInput struct {
	Id ID `json:"id"`
}

//#region Assign

// Deprecated: Use AssignTagsFor instead
func (client *Client) AssignTagsForAlias(alias string, tags map[string]string) ([]Tag, error) {
	return client.AssignTags(alias, tags)
}

// Deprecated: Use AssignTagFor instead
func (client *Client) AssignTagForAlias(alias string, key string, value string) ([]Tag, error) {
	return client.AssignTags(alias, map[string]string{key: value})
}

// Deprecated: Use AssignTagsFor instead
func (client *Client) AssignTagsForId(id ID, tags map[string]string) ([]Tag, error) {
	return client.AssignTags(string(id), tags)
}

// Deprecated: Use AssignTagFor instead
func (client *Client) AssignTagForId(id ID, key string, value string) ([]Tag, error) {
	return client.AssignTags(string(id), map[string]string{key: value})
}

func (client *Client) AssignTags(identifier string, tags map[string]string) ([]Tag, error) {
	input := TagAssignInput{
		Tags: []TagInput{},
	}
	for key, value := range tags {
		input.Tags = append(input.Tags, TagInput{
			Key:   key,
			Value: value,
		})
	}
	if IsID(identifier) {
		input.Id = ID(identifier)
	} else {
		input.Alias = identifier
	}
	return client.AssignTag(input)
}

func (client *Client) AssignTag(input TagAssignInput) ([]Tag, error) {
	var m struct {
		Payload struct {
			Tags   []Tag
			Errors []OpsLevelErrors
		} `graphql:"tagAssign(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TagAssign"))
	return m.Payload.Tags, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Create

func (client *Client) CreateTags(identifier string, tags map[string]string) ([]Tag, error) {
	var output []Tag
	for key, value := range tags {
		input := TagCreateInput{
			Key:   key,
			Value: value,
		}
		if IsID(identifier) {
			input.Id = ID(identifier)
		} else {
			input.Alias = identifier
		}
		newTag, err := client.CreateTag(input)
		if err != nil {
			// TODO: combind errors?
		} else {
			output = append(output, *newTag)
		}
	}
	return output, nil
}

// Deprecated: Use CreateTags instead
func (client *Client) CreateTagsForId(id ID, tags map[string]string) ([]Tag, error) {
	return client.CreateTags(string(id), tags)
}

func (client *Client) CreateTag(input TagCreateInput) (*Tag, error) {
	var m struct {
		Payload struct {
			Tag    Tag `json:"tag"`
			Errors []OpsLevelErrors
		} `graphql:"tagCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TagCreate"))
	return &m.Payload.Tag, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

// Deprecated: use client.GetServiceWithAlias(alias).Tags instead
func (client *Client) GetTagsForServiceWithAlias(alias string) ([]Tag, error) {
	service, err := client.GetServiceWithAlias(alias)
	return service.Tags.Nodes, err
}

// Deprecated: use client.GetService(id).Tags instead
func (client *Client) GetTagsForServiceWithId(id ID) ([]Tag, error) {
	service, err := client.GetService(id)
	return service.Tags.Nodes, err
}

// Deprecated: use client.GetService(id).Tags instead
func (client *Client) GetTagsForService(id ID) ([]Tag, error) {
	service, err := client.GetService(id)
	return service.Tags.Nodes, err
}

// Deprecated: use client.GetService(id).Tags.TotalCount instead
func (client *Client) GetTagCount(id ID) (int, error) {
	service, err := client.GetService(id)
	return service.Tags.TotalCount, err
}

//#endregion

//#region Update

func (client *Client) UpdateTag(input TagUpdateInput) (*Tag, error) {
	var m struct {
		Payload struct {
			Tag    Tag
			Errors []OpsLevelErrors
		} `graphql:"tagUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TagUpdate"))
	return &m.Payload.Tag, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteTag(id ID) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"tagDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TagDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("TagDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
