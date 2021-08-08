package opslevel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

const defaultURL = "https://api.opslevel.com/graphql"

type ClientSettings struct {
	url           string
	apiVisibility string
	pageSize      int
	ctx           context.Context
}

type Client struct {
	url      string
	pageSize graphql.Int
	ctx      context.Context // Should this be here?
	client   *graphql.Client
}

type option func(*ClientSettings)

func SetURL(url string) option {
	return func(c *ClientSettings) {
		c.url = url
	}
}

func SetContext(ctx context.Context) option {
	return func(c *ClientSettings) {
		c.ctx = ctx
	}
}

func SetPageSize(size int) option {
	return func(c *ClientSettings) {
		c.pageSize = size
	}
}

func SetAPIVisibility(visibility string) option {
	return func(c *ClientSettings) {
		c.apiVisibility = visibility
	}
}

type customTransport struct {
	apiVisibility string
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("GraphQL-Visibility", t.apiVisibility)
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient(apiToken string, options ...option) *Client {
	httpToken := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken, TokenType: "Bearer"},
	)
	settings := &ClientSettings{
		url:           defaultURL,
		apiVisibility: "public",
		pageSize:      100,
		ctx:           context.Background(),
	}
	for _, opt := range options {
		opt(settings)
	}
	return &Client{
		url:      settings.url,
		pageSize: graphql.Int(settings.pageSize),
		ctx:      settings.ctx,
		client: graphql.NewClient(settings.url, &http.Client{
			Timeout: time.Second * 10,
			Transport: &oauth2.Transport{
				Source: httpToken,
				Base: &customTransport{
					apiVisibility: settings.apiVisibility,
				},
			},
		}),
	}
}

func (c *Client) InitialPageVariables() PayloadVariables {
	return PayloadVariables{
		"after": graphql.String(""),
		"first": c.pageSize,
	}
}

// Should we create a context for every query/mutate ?
func (c *Client) Query(q interface{}, variables map[string]interface{}) error {
	return c.client.Query(c.ctx, q, variables)
}

func (c *Client) Mutate(m interface{}, variables map[string]interface{}) error {
	return c.client.Mutate(c.ctx, m, variables)
}

func (c *Client) Validate() error {
	var q struct {
		Account struct {
			Id graphql.ID
		}
	}
	err := c.Query(&q, nil)
	// TODO: we should probably use a custom OpsLevelClientError type - https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
	if err != nil {
		if strings.Contains(err.Error(), "401 Unauthorized") {
			return errors.New("Client Validation Error: Please provide a valid OpsLevel API token")
		}
		return fmt.Errorf("Client Validation Error: %s", err.Error())
	}
	return nil
}
