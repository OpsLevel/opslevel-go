package opslevel

import "context"

func (c *Client) CreateTag(ctx context.Context, key, value, alias, resourceType string) (*Tag, error) {
	args := map[string]string{
		"alias": alias,
		"type":  resourceType,
		"key":   key,
		"value": value,
	}
	params := map[string]interface{}{
		"input": args,
	}
	var res createTagResponse
	if err := c.Do(ctx, tagCreateMutation, params, &res); err != nil {
		return nil, err
	}
	// Check for application level errors
	if err := handleGraphqlErrs(res.TagCreate.Errors); err != nil {
		return nil, err
	}
	return res.TagCreate.Tag, nil
}

type Tag struct {
	Id    string
	Owner string
	Key   string
	Value string
}

type createTagResponse struct {
	TagCreate struct {
		Tag    *Tag
		Errors []graphqlError
	}
}

const tagCreateMutation = `
mutation create($input: TagCreateInput!){
  tagCreate(input: $input){
    tag{
      key
      value
    }
    errors {
      path
      message
    }
  }
}
`
