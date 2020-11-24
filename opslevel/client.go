package opslevel

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

func NewClient(url, authToken string) Client {
	return Client{
		url:           url,
		bearerToken:   fmt.Sprintf("Bearer %s", authToken),
		graphqlClient: graphql.NewClient(url),
	}
}

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
	if err := c.makeReq(ctx, tagCreateMutation, params, &res); err != nil {
		return nil, err
	}
	// Check for application level errors
	if err := handleGraphqlErrs(res.TagCreate.Errors); err != nil {
		return nil, err
	}
	return res.TagCreate.Tag, nil
}

func (c *Client) GetTeam(ctx context.Context, alias string) (*Team, error) {
	params := map[string]interface{}{
		"teamAlias": alias,
	}
	var res teamResponse
	if err := c.makeReq(ctx, teamQuery, params, &res); err != nil {
		return nil, fmt.Errorf("could not find team: %w", err)
	}
	if res.Account.Team == nil {
		return nil, fmt.Errorf("no team was found by alias: %s", alias)
	}
	return res.Account.Team, nil
}

func (c *Client) makeReq(ctx context.Context, query string, params map[string]interface{}, res interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Set("Authorization", c.bearerToken)
	for key, value := range params {
		req.Var(key, value)
	}
	var resp map[string]interface{}
	return c.graphqlClient.Run(ctx, req, &resp)
}

func handleGraphqlErrs(errs []graphqlError) error {
	if len(errs) == 0 {
		return nil
	}
	var errMsg string
	for _, err := range errs {
		errMsg += fmt.Sprintf("%s path: %s", err.Message, err.Path)
	}
	return fmt.Errorf("could not create tag: %s", errMsg)
}
