package opslevel

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
)

type TagOwner string

const (
	TagOwnerService    TagOwner = "Service"
	TagOwnerRepository TagOwner = "Repository"
)

type TaggableResourceInterface interface {
	GetTags(*Client, *PayloadVariables) (*TagConnection, error)
	ResourceId() ID
	ResourceType() TaggableResource
}

var (
	TagKeyRegex    = regexp.MustCompile(`\A[a-z][0-9a-z_\.\/\\-]*\z`)
	TagKeyErrorMsg = "tag key name '%s' must start with a letter and be only lowercase alphanumerics, underscores, hyphens, periods, and slashes"
)

type Tag struct {
	Id    ID     `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (t Tag) HasSameKeyValue(otherTag Tag) bool {
	return t.Key == otherTag.Key && t.Value == otherTag.Value
}

func (t Tag) Flatten() string {
	return fmt.Sprintf("%s:%s", t.Key, t.Value)
}

type TagConnection struct {
	Nodes      []Tag
	PageInfo   PageInfo
	TotalCount int
}

func (client *Client) GetTaggableResource(resourceType TaggableResource, identifier string) (TaggableResourceInterface, error) {
	var err error
	var taggableResource TaggableResourceInterface

	switch resourceType {
	case TaggableResourceService:
		if IsID(identifier) {
			taggableResource, err = client.GetService(ID(identifier))
		} else {
			taggableResource, err = client.GetServiceWithAlias(identifier)
		}
	case TaggableResourceRepository:
		if IsID(identifier) {
			taggableResource, err = client.GetRepository(ID(identifier))
		} else {
			taggableResource, err = client.GetRepositoryWithAlias(identifier)
		}
	case TaggableResourceTeam:
		if IsID(identifier) {
			taggableResource, err = client.GetTeam(ID(identifier))
		} else {
			taggableResource, err = client.GetTeamWithAlias(identifier)
		}
	case TaggableResourceDomain:
		taggableResource, err = client.GetDomain(identifier)
	case TaggableResourceInfrastructureresource:
		taggableResource, err = client.GetInfrastructure(identifier)
	case TaggableResourceSystem:
		taggableResource, err = client.GetSystem(identifier)
	case TaggableResourceUser:
		taggableResource, err = client.GetUser(identifier)
	default:
		return nil, fmt.Errorf("not a taggable resource type: %s" + string(resourceType))
	}

	if err != nil {
		return nil, err
	}
	return taggableResource, nil
}

func (tagConnection *TagConnection) GetTagById(tagId ID) (*Tag, error) {
	for _, tag := range tagConnection.Nodes {
		if tag.Id == tagId {
			return &tag, nil
		}
	}
	return nil, fmt.Errorf("tag with ID '%s' not found", tagId)
}

func ValidateTagKey(key string) error {
	if !TagKeyRegex.MatchString(key) {
		return fmt.Errorf(TagKeyErrorMsg, key)
	}
	return nil
}

func (client *Client) AssignTagsWithTags(identifier string, tags []Tag) ([]Tag, error) {
	var tagInput []TagInput
	for _, tag := range tags {
		tagInput = append(tagInput, TagInput{Key: tag.Key, Value: tag.Value})
	}
	input := TagAssignInput{
		Tags: tagInput,
	}
	if IsID(identifier) {
		input.Id = NewID(identifier)
	} else {
		input.Alias = &identifier
	}
	return client.AssignTag(input)
}

func (client *Client) AssignTags(identifier string, tags map[string]string) ([]Tag, error) {
	var tagsSlice []Tag

	for key, value := range tags {
		if err := ValidateTagKey(key); err != nil {
			return nil, err
		}
		tagsSlice = append(tagsSlice, Tag{
			Key:   key,
			Value: value,
		})
	}
	return client.AssignTagsWithTags(identifier, tagsSlice)
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

func (client *Client) CreateTags(identifier string, tags map[string]string) ([]Tag, error) {
	var output []Tag
	var allErrors error
	for key, value := range tags {
		if err := ValidateTagKey(key); err != nil {
			return nil, err
		}
		input := TagCreateInput{
			Key:   key,
			Value: value,
		}
		if IsID(identifier) {
			input.Id = NewID(identifier)
		} else {
			input.Alias = &identifier
		}
		newTag, err := client.CreateTag(input)
		if err != nil {
			allErrors = errors.Join(allErrors, err)
		} else {
			output = append(output, *newTag)
		}
	}
	return output, allErrors
}

func (client *Client) CreateTag(input TagCreateInput) (*Tag, error) {
	var m struct {
		Payload struct {
			Tag    Tag `json:"tag"`
			Errors []OpsLevelErrors
		} `graphql:"tagCreate(input: $input)"`
	}
	if err := ValidateTagKey(input.Key); err != nil {
		return nil, err
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TagCreate"))
	return &m.Payload.Tag, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateTag(input TagUpdateInput) (*Tag, error) {
	var m struct {
		Payload struct {
			Tag    Tag
			Errors []OpsLevelErrors
		} `graphql:"tagUpdate(input: $input)"`
	}
	if err := ValidateTagKey(*input.Key); err != nil {
		return nil, err
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TagUpdate"))
	return &m.Payload.Tag, HandleErrors(err, m.Payload.Errors)
}

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

// ReconcileTags manages tags API operations for TaggableResourceInterface implementations
//
// Tags not in 'tagsWanted' will be deleted, new tags from 'tagsWanted' will be created
func (client *Client) ReconcileTags(resourceType TaggableResourceInterface, tagsWanted []Tag) error {
	var allErrors, err error
	var tagConnection *TagConnection

	tagConnection, err = resourceType.GetTags(client, nil)
	if err != nil {
		return err
	}
	if tagConnection == nil {
		return fmt.Errorf("no tags found on %s with id '%s'", string(resourceType.ResourceType()), resourceType.ResourceId())
	}

	tagsToCreate, tagsToDelete := ExtractTags(tagConnection.Nodes, tagsWanted)
	// delete tags found in resource but not listed in tagsWanted
	for _, tag := range tagsToDelete {
		allErrors = errors.Join(allErrors, client.DeleteTag(tag.Id))
	}
	if allErrors != nil {
		return allErrors
	}

	if len(tagsToCreate) > 0 {
		_, err = client.AssignTagsWithTags(string(resourceType.ResourceId()), tagsToCreate)
		if err != nil {
			return err
		}
	}

	return nil
}

// Given actual tags and wanted tags, returns tagsToCreate and tagsToDelete lists
func ExtractTags(existingTags, tagsWanted []Tag) ([]Tag, []Tag) {
	var tagsToCreate, tagsToDelete []Tag

	// collect tagsToDelete - existing tags that are no longer wanted
	for _, existingTag := range existingTags {
		if !slices.ContainsFunc(tagsWanted, func(t Tag) bool { return existingTag.HasSameKeyValue(t) }) {
			tagsToDelete = append(tagsToDelete, existingTag)
		}
	}

	// collect tagsToCreate - wanted tags that do not yet exist
	for _, tagWanted := range tagsWanted {
		if !slices.ContainsFunc(existingTags, func(t Tag) bool { return tagWanted.HasSameKeyValue(t) }) {
			tagsToCreate = append(tagsToCreate, tagWanted)
		}
	}

	return tagsToCreate, tagsToDelete
}
