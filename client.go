package opslevel

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type ClientSettings struct {
	url            string
	userAgentExtra string
	timeout        time.Duration
	retries        int

	apiVisibility string // Only Used by GQL
	pageSize      int    // Only Used by GQL
}

type Option func(*ClientSettings)

func newClientSettings(options ...Option) *ClientSettings {
	settings := &ClientSettings{
		url:     "https://app.opslevel.com/",
		timeout: time.Second * 10,
		retries: 10,

		pageSize:      100,
		apiVisibility: "public",
	}
	for _, opt := range options {
		opt(settings)
	}
	return settings
}

func SetAPIToken(apiToken string) Option {
	return func(c *ClientSettings) {
		os.Setenv("OPSLEVEL_API_TOKEN", apiToken)
	}
}

func SetURL(url string) Option {
	return func(c *ClientSettings) {
		c.url = fmt.Sprintf("%s/graphql", strings.TrimRight(url, "/"))
	}
}

// SetTestUrl - Only use this when making test Clients
func SetTestURL(url string) Option {
	return func(c *ClientSettings) {
		c.url = url
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

func SetMaxRetries(amount int) Option {
	return func(c *ClientSettings) {
		c.retries = amount
	}
}

func SetAPIVisibility(visibility string) Option {
	return func(c *ClientSettings) {
		c.apiVisibility = visibility
	}
}

func SetPageSize(size int) Option {
	return func(c *ClientSettings) {
		c.pageSize = size
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
