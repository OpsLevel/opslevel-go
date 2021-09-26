package opslevel

import (
	"fmt"
	"regexp"

	"github.com/shurcooL/graphql"
)

type TagOwner string

const (
	TagOwnerService    TagOwner = "Service"
	TagOwnerRepository TagOwner = "Repository"
)

type Tag struct {
	Id    graphql.ID `json:"id"`
	Key   string     `json:"key"`
	Value string     `json:"value"`
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

var TagKeyRegex = regexp.MustCompile("\\A[a-z][0-9a-z_\\.\\/\\\\-]*\\z")
var TagKeyErrorMsg = "must start with a letter and only lowercase alphanumerics, underscores, hyphens, periods, and slashes are allowed."

func (t *TagInput) Validate() error {
	if !TagKeyRegex.MatchString(t.Key) {
		return fmt.Errorf("invalid tag key name '%s' - %s", t.Key, TagKeyErrorMsg)
	}
	return nil
}

type TagAssignInput struct {
	Id    graphql.ID       `json:"id,omitempty"`
	Alias string           `json:"alias,omitempty"`
	Type  TaggableResource `json:"type,omitempty"`
	Tags  []TagInput       `json:"tags"`
}

func (t *TagAssignInput) Validate() error {
	for _, tag := range t.Tags {
		if err := tag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TagCreateInput struct {
	Id    graphql.ID       `json:"id"`
	Alias string           `json:"alias,omitempty"`
	Type  TaggableResource `json:"type,omitempty"`
	Key   string           `json:"key"`
	Value string           `json:"value"`
}

func (t *TagCreateInput) Validate() error {
	if !TagKeyRegex.MatchString(t.Key) {
		return fmt.Errorf("invalid tag key name '%s' - %s", t.Key, TagKeyErrorMsg)
	}
	return nil
}

type TagUpdateInput struct {
	Id    graphql.ID `json:"id"`
	Key   string     `json:"key,omitempty"`
	Value string     `json:"value,omitempty"`
}

func (t *TagUpdateInput) Validate() error {
	if !TagKeyRegex.MatchString(t.Key) {
		return fmt.Errorf("invalid tag key name '%s' - %s", t.Key, TagKeyErrorMsg)
	}
	return nil
}

type TagDeleteInput struct {
	Id graphql.ID `json:"id"`
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

func (client *Client) AssignTagsForId(id graphql.ID, tags map[string]string) ([]Tag, error) {
	input := TagAssignInput{
		Id:   graphql.ID(id),
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

func (client *Client) AssignTagForId(id graphql.ID, key string, value string) ([]Tag, error) {
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
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return m.Payload.Tags, FormatErrors(m.Payload.Errors)
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

func (client *Client) CreateTagsForId(id graphql.ID, tags map[string]string) ([]Tag, error) {
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
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Tag, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (conn *TagConnection) Hydrate(service graphql.ID, client *Client) error {
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
func (client *Client) GetTagsForServiceWithId(service graphql.ID) ([]Tag, error) {
	return client.GetTagsForService(service)
}

func (client *Client) GetTagsForService(service graphql.ID) ([]Tag, error) {
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
	if err := client.Query(&q, v); err != nil {
		return q.Account.Service.Tags.Nodes, err
	}
	if err := q.Account.Service.Tags.Hydrate(service, client); err != nil {
		return q.Account.Service.Tags.Nodes, err
	}
	return q.Account.Service.Tags.Nodes, nil
}

func (client *Client) GetTagCount(service graphql.ID) (int, error) {
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
	if err := client.Query(&q, v); err != nil {
		return 0, err
	}
	return int(q.Account.Service.Tags.TotalCount), nil
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
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Tag, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteTag(id graphql.ID) error {
	var m struct {
		Payload struct {
			Id     graphql.ID `graphql:"deletedTagId"`
			Errors []OpsLevelErrors
		} `graphql:"tagDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TagDeleteInput{Id: id},
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
