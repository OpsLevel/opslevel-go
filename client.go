package opslevel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

const defaultURL = "https://api.opslevel.com/graphql"

type ClientSettings struct {
	url        string
	ctx        context.Context
	httpClient *http.Client
}

type Client struct {
	url    string
	ctx    context.Context // Should this be here?
	client *graphql.Client
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

func SetHttpClient(client *http.Client) option {
	return func(c *ClientSettings) {
		c.httpClient = client
	}
}

type customTransport struct{}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("GraphQL-Visibility", "internal")
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient(apiToken string, options ...option) *Client {
	httpToken := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken, TokenType: "Bearer"},
	)
	settings := &ClientSettings{
		url: defaultURL,
		ctx: context.Background(),
		httpClient: &http.Client{
			Transport: &oauth2.Transport{
				Source: httpToken,
				Base:   &customTransport{},
			},
		},
	}
	for _, opt := range options {
		opt(settings)
	}
	return &Client{
		url:    settings.url,
		ctx:    settings.ctx,
		client: graphql.NewClient(settings.url, settings.httpClient),
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
	if err != nil {
		return fmt.Errorf("Unable to Validate connection to OpsLevel API: %s", err.Error())
	}
	return nil
}
