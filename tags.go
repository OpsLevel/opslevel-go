package opslevel

import (
	"github.com/hasura/go-graphql-client"
)

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

func (client *Client) AssignTagsForAlias(alias string, tags map[string]string) ([]Tag, error) {
	input := TagAssignInput{
		Alias: alias,
		Tags:  []TagInput{},
	}
	for key, value := range tags {
		input.Tags = append(input.Tags, TagInput{
			Key:   key,
			Value: value,
		})
	}
	return client.AssignTags(input)
}

func (client *Client) AssignTagForAlias(alias string, key string, value string) ([]Tag, error) {
	input := TagAssignInput{
		Alias: alias,
		Tags:  []TagInput{},
	}
	input.Tags = append(input.Tags, TagInput{
		Key:   key,
		Value: value,
	})
	return client.AssignTags(input)
}

func (client *Client) AssignTagsForId(id ID, tags map[string]string) ([]Tag, error) {
	input := TagAssignInput{
		Id:   id,
		Tags: []TagInput{},
	}
	for key, value := range tags {
		input.Tags = append(input.Tags, TagInput{
			Key:   key,
			Value: value,
		})
	}
	return client.AssignTags(input)
}

func (client *Client) AssignTagForId(id ID, key string, value string) ([]Tag, error) {
	input := TagAssignInput{
		Id:   id,
		Tags: []TagInput{},
	}
	input.Tags = append(input.Tags, TagInput{
		Key:   key,
		Value: value,
	})
	return client.AssignTags(input)
}

func (client *Client) AssignTags(input TagAssignInput) ([]Tag, error) {
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

func (client *Client) CreateTags(alias string, tags map[string]string) ([]Tag, error) {
	var output []Tag
	for key, value := range tags {
		input := TagCreateInput{
			Alias: alias,
			Key:   key,
			Value: value,
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

func (client *Client) CreateTagsForId(id ID, tags map[string]string) ([]Tag, error) {
	var output []Tag
	for key, value := range tags {
		input := TagCreateInput{
			Id:    id,
			Key:   key,
			Value: value,
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

func (client *Client) CreateTag(input TagCreateInput) (*Tag, error) {
	var m struct {
		Payload struct {
			Tag    Tag
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

func (conn *TagConnection) Hydrate(service ID, client *Client) error {
	var q struct {
		Account struct {
			Service struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": service,
		"first":   client.pageSize,
	}
	q.Account.Service.Tags.PageInfo = conn.PageInfo
	for q.Account.Service.Tags.PageInfo.HasNextPage {
		v["after"] = q.Account.Service.Tags.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Service.Tags.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
}

func (client *Client) GetTagsForServiceWithAlias(alias string) ([]Tag, error) {
	service, serviceErr := client.GetServiceIdWithAlias(alias)
	if serviceErr != nil {
		return nil, serviceErr
	}
	return client.GetTagsForService(service.Id)
}

// Deprecated: use GetTagsForService instead
func (client *Client) GetTagsForServiceWithId(service ID) ([]Tag, error) {
	return client.GetTagsForService(service)
}

func (client *Client) GetTagsForService(service ID) ([]Tag, error) {
	var q struct {
		Account struct {
			Service struct {
				Tags TagConnection `graphql:"tags(after: $after, first: $first)"`
			} `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": service,
		"after":   graphql.String(""),
		"first":   client.pageSize,
	}
	if err := client.Query(&q, v, WithName("ServiceTagsList")); err != nil {
		return q.Account.Service.Tags.Nodes, err
	}
	if err := q.Account.Service.Tags.Hydrate(service, client); err != nil {
		return q.Account.Service.Tags.Nodes, err
	}
	return q.Account.Service.Tags.Nodes, nil
}

func (client *Client) GetTagCount(service ID) (int, error) {
	var q struct {
		Account struct {
			Service struct {
				Tags struct {
					TotalCount int
				}
			} `graphql:"service(id: $service)"`
		}
	}
	v := PayloadVariables{
		"service": service,
	}
	err := client.Query(&q, v, WithName("ServiceCount"))
	return int(q.Account.Service.Tags.TotalCount), HandleErrors(err, nil)
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
			Id     ID `graphql:"deletedTagId"`
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
