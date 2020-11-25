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

type Client struct {
	url           string
	bearerToken   string
	graphqlClient *graphql.Client
}

func (c *Client) Do(ctx context.Context, query string, params map[string]interface{}, res interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Set("Authorization", c.bearerToken)
	for key, value := range params {
		req.Var(key, value)
	}
	return c.graphqlClient.Run(ctx, req, res)
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

type graphqlError struct {
	Path    []string
	Message string
}
