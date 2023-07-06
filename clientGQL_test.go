package opslevel_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestMain(m *testing.M) {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	log.Logger = log.Output(output)
	flag.Parse()
	teardown := autopilot.Setup()
	defer teardown()
	os.Exit(m.Run())
}

func Templated(input string) string {
	response, err := autopilot.Templater.Use(input)
	if err != nil {
		panic(err)
	}
	return response
}

func TemplatedResponse(response string) autopilot.ResponseWriter {
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, Templated(response))
	}
}

func GraphQLQueryTemplate(request string) autopilot.GraphqlQuery {
	exp := autopilot.GraphqlQuery{
		Variables: nil,
	}
	if err := json.Unmarshal([]byte(Templated(request)), &exp); err != nil {
		panic(err)
	}
	return exp
}

func GraphQLQueryTemplatedValidation(t *testing.T, request string) autopilot.RequestValidation {
	return func(r *http.Request) {
		autopilot.Equals(t, autopilot.ToJson(GraphQLQueryTemplate(request)), autopilot.ToJson(autopilot.Parse(r)))
	}
}

func ABetterTestClient(t *testing.T, endpoint string, request string, response string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(autopilot.RegisterEndpoint(fmt.Sprintf("/LOCAL_TESTING/%s", endpoint),
		TemplatedResponse(response),
		GraphQLQueryTemplatedValidation(t, request))))
}

func ATestClient(t *testing.T, endpoint string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(autopilot.RegisterEndpoint(fmt.Sprintf("/LOCAL_TESTING/%s", endpoint),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", endpoint)),
		autopilot.GraphQLQueryFixtureValidation(t, fmt.Sprintf("%s_request.json", endpoint)))))
}

type TestRequest struct {
	Request  string
	Response string
}

func RegisterPaginatedEndpoint(t *testing.T, endpoint string, requests ...TestRequest) string {
	url := fmt.Sprintf("/LOCAL_TESTING/%s", endpoint)
	requestCount := 0
	autopilot.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		GraphQLQueryTemplatedValidation(t, requests[requestCount].Request)(r)
		TemplatedResponse(requests[requestCount].Response)(w)
		requestCount += 1
	})
	return autopilot.Server.URL + url
}

func APaginatedTestClient(t *testing.T, endpoint string, requests ...TestRequest) *ol.Client {
	url := RegisterPaginatedEndpoint(t, endpoint, requests...)
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(url))
}

func ATestClientAlt(t *testing.T, response string, request string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(autopilot.RegisterEndpoint(fmt.Sprintf("/LOCAL_TESTING/%s__%s", response, request),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", response)),
		autopilot.GraphQLQueryFixtureValidation(t, fmt.Sprintf("%s_request.json", request)))))
}

func ATestClientSkipRequest(t *testing.T, endpoint string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(autopilot.RegisterEndpoint(fmt.Sprintf("/LOCAL_TESTING/%s", endpoint),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", endpoint)),
		autopilot.SkipRequestValidation())))
}

func TestClientQuery(t *testing.T) {
	// Arrange
	headers := map[string]string{"x": "x"}
	request := `{
    "query": "{account{id}}",
	"variables":{}
}`
	response := `{"data": {
	"account": {
		"id": "1234"
	}
}}`
	url := autopilot.RegisterEndpoint("/LOCAL_TESTING/account",
		TemplatedResponse(response),
		GraphQLQueryTemplatedValidation(t, request))
	client := ol.NewGQLClient(
		ol.SetAPIToken("x"),
		ol.SetMaxRetries(0),
		ol.SetURL(url),
		ol.SetHeaders(headers),
		ol.SetUserAgentExtra("x"),
		ol.SetTimeout(0),
		ol.SetAPIVisibility("internal"),
		ol.SetPageSize(100))
	var q struct {
		Account struct {
			Id ol.ID
		}
	}
	var v map[string]interface{}
	// Act
	err := client.Query(&q, v)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1234", string(q.Account.Id))
}
