package opslevel_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	queryPrefix string = `"query":`
	varPrefix   string = `"variables":`
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
	json.Unmarshal([]byte(Templated(request)), &exp)
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

func NewTestRequest(request string, variables string, response string) TestRequest {
	testRequest := TestRequest{templater: autopilot.NewFixtureTemplater()}
	if err := testRequest.ParseRequest(request); err != nil {
		panic(err)
	}
	if err := testRequest.ParseVariables(variables); err != nil {
		panic(err)
	}
	if err := testRequest.ParseResponse(response); err != nil {
		panic(err)
	}
	return testRequest
}

type TestRequest struct {
	templater *autopilot.FixtureTemplater
	Request   string
	Variables string
	Response  string
}

func (t *TestRequest) IsValidJson(data string) bool {
	return json.Valid([]byte(data))
}

func (t *TestRequest) ParseRequest(rawRequest string) error {
	queryValue, _ := strings.CutPrefix(rawRequest, queryPrefix)
	parsedRequest, err := t.templater.Use(fmt.Sprintf("{%s %s}", queryPrefix, queryValue))
	if err != nil {
		return err
	}

	if !t.IsValidJson(parsedRequest) {
		return fmt.Errorf("invalid json: %s", parsedRequest)
	}
	t.Request = strings.Trim(parsedRequest, "{}")
	return nil
}

func (t *TestRequest) ParseVariables(rawVariables string) error {
	variables, _ := strings.CutPrefix(rawVariables, varPrefix)
	parsedVariables, err := t.templater.Use(variables)
	if err != nil {
		return err
	}

	if !t.IsValidJson(fmt.Sprintf("{%s %s}", varPrefix, parsedVariables)) {
		return fmt.Errorf("invalid json: %s", parsedVariables)
	}
	t.Variables = fmt.Sprintf("%s %s", varPrefix, parsedVariables)
	return nil
}

func (t *TestRequest) ParseResponse(rawResponse string) error {
	parsedResponse, err := t.templater.Use(rawResponse)
	if err != nil {
		return err
	}
	if !t.IsValidJson(parsedResponse) {
		return fmt.Errorf("invalid json: %s", parsedResponse)
	}
	t.Response = parsedResponse
	return nil
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

func TmpBetterTestClient(t *testing.T, endpoint string, testRequest TestRequest) *ol.Client {
	oldStyleRequest := fmt.Sprintf(`{%s, %s}`, testRequest.Request, testRequest.Variables)
	return ABetterTestClient(t, endpoint, oldStyleRequest, testRequest.Response)
}

func TmpPaginatedTestClient(t *testing.T, endpoint string, requests ...TestRequest) *ol.Client {
	oldStyleRequests := []TestRequest{}
	for _, request := range requests {
		oldStyleRequest := TestRequest{
			Request:  fmt.Sprintf(`{%s, %s}`, request.Request, request.Variables),
			Response: request.Response,
		}
		oldStyleRequests = append(oldStyleRequests, oldStyleRequest)
	}
	return APaginatedTestClient(t, endpoint, oldStyleRequests...)
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
