package opslevel

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
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

func (t Tag) HasSameKeyValue(otherTag Tag) bool {
	return t.Key == otherTag.Key && t.Value == otherTag.Value
}

func (t Tag) Flatten() string {
	return fmt.Sprintf("%s:%s", t.Key, t.Value)
}

func (client *Client) GetTaggableResource(resourceType TaggableResource, identifier string) (TaggableResourceInterface, error) {
	var err error
	var taggableResource TaggableResourceInterface

	switch resourceType {
	case TaggableResourceService:
		if IsID(identifier) {
			taggableResource, err = client.GetService(identifier)
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
		err = fmt.Errorf("not a taggable resource type: %s", resourceType)
	}

	return taggableResource, err
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

func (client *Client) AssignTagsWithTagInputs(identifier string, tags []TagInput) ([]Tag, error) {
	input := TagAssignInput{
		Tags: tags,
	}
	if IsID(identifier) {
		input.Id = RefOf(ID(identifier))
	} else {
		input.Alias = RefOf(identifier)
	}
	return client.AssignTag(input)
}

func (client *Client) AssignTags(identifier string, tags map[string]string) ([]Tag, error) {
	var tagInputs []TagInput

	for key, value := range tags {
		if err := ValidateTagKey(key); err != nil {
			return nil, err
		}
		tagInputs = append(tagInputs, TagInput{Key: key, Value: value})
	}

	return client.AssignTagsWithTagInputs(identifier, tagInputs)
}

func (client *Client) AssignTag(input TagAssignInput) ([]Tag, error) {
	var m struct {
		Payload TagAssignPayload `graphql:"tagAssign(input: $input)"`
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
		Payload TagCreatePayload `graphql:"tagCreate(input: $input)"`
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
		Payload TagUpdatePayload `graphql:"tagUpdate(input: $input)"`
	}
	if input.Key != nil {
		if err := ValidateTagKey(*input.Key); err != nil {
			return nil, err
		}
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("TagUpdate"))
	return &m.Payload.Tag, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteTag(id ID) error {
	var m struct {
		Payload BasePayload `graphql:"tagDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": TagDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("TagDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

// ReconcileTags manages tags API operations for TaggableResourceInterface implementations
//
// Tags from `tagsDesired` are compared against current tags of TaggableResourceInterface and differences are either created or deleted.
func (client *Client) ReconcileTags(resourceType TaggableResourceInterface, tagsDesired []Tag) error {
	tagConnection, err := resourceType.GetTags(client, nil)
	if err != nil {
		return err
	}

	toCreate, toDelete := reconcileTags(tagConnection.Nodes, tagsDesired)
	for _, tag := range toCreate {
		taggableResourceType := resourceType.ResourceType()
		taggableResourceID := resourceType.ResourceId()
		_, err := client.CreateTag(TagCreateInput{
			Id:    &taggableResourceID,
			Type:  &taggableResourceType,
			Key:   tag.Key,
			Value: tag.Value,
		})
		if err != nil {
			return err
		}
	}
	for _, tag := range toDelete {
		err := client.DeleteTag(tag.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func reconcileTags(currentTags, desiredTags []Tag) ([]Tag, []Tag) {
	toCreate := make([]Tag, 0)
	toDelete := make([]Tag, 0)

	for _, tag := range currentTags {
		if slices.ContainsFunc(desiredTags, func(t Tag) bool { return tag.HasSameKeyValue(t) }) {
			continue
		}
		toDelete = append(toDelete, tag)
	}

	for _, tag := range desiredTags {
		if slices.ContainsFunc(currentTags, func(t Tag) bool { return tag.HasSameKeyValue(t) }) {
			continue
		}
		toCreate = append(toCreate, tag)
	}

	return toCreate, toDelete
}
