package opslevel

import (
	"github.com/go-resty/resty/v2"
)

type RestResponse struct {
	Result string `json:"result"`
}

func NewRestClient(options ...Option) *resty.Client {
	client := resty.New()
	settings := newClientSettings(options...)
	client.SetBaseURL(settings.url)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("User-Agent", buildUserAgent(settings.userAgentExtra))
	client.SetTimeout(settings.timeout)
	return client
}
