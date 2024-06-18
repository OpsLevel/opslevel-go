package opslevel

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
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

func (client *Client) AssignTags(identifier string, tags map[string]string) ([]Tag, error) {
	input := TagAssignInput{
		Tags: []TagInput{},
	}
	for key, value := range tags {
		if err := ValidateTagKey(key); err != nil {
			return nil, err
		}
		input.Tags = append(input.Tags, TagInput{
			Key:   key,
			Value: value,
		})
	}
	if IsID(identifier) {
		input.Id = NewID(identifier)
	} else {
		input.Alias = &identifier
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
// Tags not in 'tagsWanted' will be deleted, new tags from 'tagsWanted' will be created. Reconciled tags are returned.
func (client *Client) ReconcileTags(resourceType TaggableResourceInterface, tagsWanted []string) ([]Tag, error) {
	var err error
	var tagConnection *TagConnection
	var assignedTags []Tag

	tagConnection, err = resourceType.GetTags(client, nil)
	if err != nil {
		return assignedTags, err
	}
	if tagConnection == nil {
		return assignedTags, fmt.Errorf("no tags found on %s with id '%s'", string(resourceType.ResourceType()), resourceType.ResourceId())
	}

	// delete tags found in resource but not listed in tagsWanted
	for _, tag := range tagConnection.Nodes {
		if !slices.Contains(tagsWanted, tag.Flatten()) {
			if err := client.DeleteTag(tag.Id); err != nil {
				return assignedTags, err
			}
		}
	}

	// format tags listed in Terraform config but not found in service
	tagInput := map[string]string{}
	for _, tag := range tagsWanted {
		parts := strings.Split(tag, ":")
		if len(parts) != 2 {
			return assignedTags, fmt.Errorf("[%s] invalid tag, should be in format 'key:value' (only a single colon between the key and value, no spaces or special characters)", tag)
		}
		key := parts[0]
		value := parts[1]
		tagInput[key] = value
	}
	// assign tags listed in Terraform config but not found in service
	assignedTags, err = client.AssignTags(string(resourceType.ResourceId()), tagInput)
	if err != nil {
		return assignedTags, err
	}

	return assignedTags, nil
}
