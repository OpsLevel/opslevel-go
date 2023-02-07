package opslevel

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hasura/go-graphql-client"
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
	retryClient.Logger = nil

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
	var url string
	if strings.Contains(settings.url, "/LOCAL_TESTING/") {
		url = settings.url
	} else {
		url = fmt.Sprintf("%s/graphql", settings.url)
	}

	return &Client{
		pageSize: graphql.Int(settings.pageSize),
		client:   graphql.NewClient(url, standardClient),
	}
}

func (client *Client) InitialPageVariables() PayloadVariables {
	return PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
	}
}

func (client *Client) InitialPageVariablesPointer() *PayloadVariables {
	v := PayloadVariables{
		"after": graphql.String(""),
		"first": client.pageSize,
	}
	return &v
}

func (client *Client) Query(q interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return client.QueryCTX(context.Background(), q, variables, options...)
}

func (client *Client) QueryCTX(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return client.client.Query(ctx, q, variables, options...)
}

func (client *Client) Mutate(m interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return client.MutateCTX(context.Background(), m, variables, options...)
}

func (client *Client) MutateCTX(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) error {
	return client.client.Mutate(ctx, m, variables, options...)
}

func (client *Client) Validate() error {
	var q struct {
		Account struct {
			Id ID
		}
	}
	err := client.Query(&q, nil)
	// TODO: we should probably use a custom OpsLevelClientError type - https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
	if err != nil {
		return fmt.Errorf("client validation error: %s", err.Error())
	}
	return nil
}

func WithName(name string) graphql.Option {
	return graphql.OperationName(name)
}
