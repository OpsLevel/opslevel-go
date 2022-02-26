package opslevel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

type RestClient struct {
	baseURL    *url.URL
	httpClient *http.Client
}

type RestClientOption func(c *RestClient)

type RestResponse struct {
	Result string `json:"result"`
}

// WithBaseURL modifies the Client baseURL.
func WithBaseURL(baseURL string) RestClientOption {
	return func(c *RestClient) {
		bu, _ := url.Parse(baseURL)
		c.baseURL = bu
	}
}

// WithHTTPClient modifies the Client http.Client.
func WithHTTPClient(hc *http.Client) RestClientOption {
	return func(c *RestClient) {
		c.httpClient = hc
	}
}

// TODO: Can we replace this all with Resty

// NewClient returns a Client pointer
func NewRestClient(opts ...RestClientOption) *RestClient {
	baseURL, _ := url.Parse("https://app.opslevel.com")
	client := &RestClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
	for _, o := range opts {
		o(client)
	}
	return client
}

func (c *RestClient) Do(method string, url string, body interface{}) (*RestResponse, error) {
	var err error

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("%s\n%s", url, string(b))
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Debug().Msgf("Failed to send request to OpsLevel: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	log.Debug().Msgf("Received status code %d", resp.StatusCode)
	if resp.StatusCode != http.StatusAccepted {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, err
		}
		s := buf.String()
		return nil, fmt.Errorf("status %d; %s", resp.StatusCode, s)
	}

	output := RestResponse{}
	err = decoder.Decode(&output)
	if err != nil {
		log.Debug().Msgf("Failed to decode response from OpsLevel: %s", err.Error())
		return nil, err
	}
	return &output, nil
}

func (c *RestClient) Get(url string, body interface{}) (*RestResponse, error) {
	return c.Do("GET", url, body)
}

func (c *RestClient) Post(url string, body interface{}) (*RestResponse, error) {
	return c.Do("POST", url, body)
}
