package opslevel

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	pageSize graphql.Int
	client   *graphql.Client
}

type customTransport struct {
	apiVisibility       string
	userAgent           string
	underlyingTransport http.RoundTripper
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("GraphQL-Visibility", t.apiVisibility)
	req.Header.Set("User-Agent", t.userAgent)
	return t.underlyingTransport.RoundTrip(req)
}

// Deprecated: Use NewGQLClient instead
func NewClient(apiToken string, options ...Option) *Client {
	options = append(options, SetAPIToken(apiToken))
	return NewGQLClient(options...)
}

func NewGQLClient(options ...Option) *Client {
	settings := newClientSettings(options...)

	httpToken := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("OPSLEVEL_API_TOKEN"), TokenType: "Bearer"},
	)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = settings.retries

	standardClient := retryClient.StandardClient()
	standardClient.Timeout = settings.timeout
	standardClient.Transport = &oauth2.Transport{
		Source: httpToken,
		Base: &customTransport{
			apiVisibility:       settings.apiVisibility,
			userAgent:           buildUserAgent(settings.userAgentExtra),
			underlyingTransport: standardClient.Transport,
		},
	}

	return &Client{
		pageSize: graphql.Int(settings.pageSize),
		client:   graphql.NewClient(settings.url, standardClient),
	}
}

func (client *Client) InitialPageVariables() PayloadVariables {
	return PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
	}
}

func (client *Client) Query(q interface{}, variables map[string]interface{}) error {
	return client.QueryCTX(context.Background(), q, variables)
}

func (client *Client) QueryCTX(ctx context.Context, q interface{}, variables map[string]interface{}) error {
	return client.client.Query(ctx, q, variables)
}

func (client *Client) Mutate(m interface{}, variables map[string]interface{}) error {
	return client.MutateCTX(context.Background(), m, variables)
}

func (client *Client) MutateCTX(ctx context.Context, m interface{}, variables map[string]interface{}) error {
	return client.client.Mutate(ctx, m, variables)
}

func (client *Client) Validate() error {
	var q struct {
		Account struct {
			Id graphql.ID
		}
	}
	err := client.Query(&q, nil)
	// TODO: we should probably use a custom OpsLevelClientError type - https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
	if err != nil {
		if strings.Contains(err.Error(), "401 Unauthorized") {
			return errors.New("Client Validation Error: Please provide a valid OpsLevel API token")
		}
		return fmt.Errorf("Client Validation Error: %s", err.Error())
	}
	return nil
}
