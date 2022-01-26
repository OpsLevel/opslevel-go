package opslevel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

const defaultURL = "https://api.opslevel.com/graphql"

type ClientSettings struct {
	url            string
	apiVisibility  string
	userAgentExtra string
	timeout        time.Duration
	pageSize       int
	ctx            context.Context
}

type Client struct {
	url      string
	pageSize graphql.Int
	ctx      context.Context // Should this be here?
	client   *graphql.Client
}

type Option func(*ClientSettings)

func SetURL(url string) Option {
	return func(c *ClientSettings) {
		c.url = url
	}
}

func SetContext(ctx context.Context) Option {
	return func(c *ClientSettings) {
		c.ctx = ctx
	}
}

func SetPageSize(size int) Option {
	return func(c *ClientSettings) {
		c.pageSize = size
	}
}

func SetAPIVisibility(visibility string) Option {
	return func(c *ClientSettings) {
		c.apiVisibility = visibility
	}
}

func SetUserAgentExtra(extra string) Option {
	return func(c *ClientSettings) {
		c.userAgentExtra = extra
	}
}

func SetTimeout(amount time.Duration) Option {
	return func(c *ClientSettings) {
		c.timeout = amount
	}
}

type customTransport struct {
	apiVisibility string
	userAgent     string
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("GraphQL-Visibility", t.apiVisibility)
	req.Header.Set("User-Agent", t.userAgent)
	return http.DefaultTransport.RoundTrip(req)
}

func NewClient(apiToken string, options ...Option) *Client {
	httpToken := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken, TokenType: "Bearer"},
	)
	settings := &ClientSettings{
		url:           defaultURL,
		apiVisibility: "public",
		pageSize:      100,
		timeout:       time.Second * 10,
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
			Timeout: settings.timeout,
			Transport: &oauth2.Transport{
				Source: httpToken,
				Base: &customTransport{
					apiVisibility: settings.apiVisibility,
					userAgent:     buildUserAgent(settings.userAgentExtra),
				},
			},
		}),
	}
}

/*
Return a string suitable for use as a User-Agent header.
The string will be of the form:

<agent_name>/<agent_version> go/<go_ver> <plat_name>/<plat_ver> client/<code_extras> user/<user_extras>
*/
func buildUserAgent(extra string) string {
	base := fmt.Sprintf("opslevel-go/%s go/%s %s/%s", clientVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	if extra != "" {
		base = fmt.Sprintf("%s client/%s", base, extra)
	}
	if value, present := os.LookupEnv("OPSLEVEL_USER_AGENT_EXTRAS"); present {
		base = fmt.Sprintf("%s user/%s", base, value)
	}
	return base
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
