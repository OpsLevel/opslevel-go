package opslevel_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig/v3"
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
	testRequest := TestRequest{}
	testRequest.templater = NewFancyTemplater()
	testRequest.ParseRequest(request)
	testRequest.ParseVariables(variables)
	testRequest.ParseResponse(response)
	return testRequest
}

type FancyTemplater struct {
	rootTemplate *template.Template
}

func (t *FancyTemplater) ParseTemplatedString(contents string) (string, error) {
	// clone, err := t.rootTemplate.Clone()
	// if err != nil {
	// 	return "", fmt.Errorf("error cloning core template: %s", err)
	// }
	// target, err := clone.Parse(contents)
	target, err := t.rootTemplate.Parse(contents)
	if err != nil {
		return "", fmt.Errorf("error parsing template: %s", err)
	}
	data := bytes.NewBuffer([]byte{})
	err = target.Execute(data, nil)
	if err != nil {
		return "", err
	}
	return data.String(), nil
}

func (t *FancyTemplater) JsonIndent(contents string) (string, error) {
	output := bytes.NewBuffer([]byte{})
	output.WriteString(contents)
	contentsAsBytes := output.Bytes()
	output.Reset()

	err := json.Indent(output, contentsAsBytes, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error during json indenting: %s", err)
	}
	return output.String(), nil
}

func NewFancyTemplater(templateDirs ...string) *FancyTemplater {
	var templateFiles []string
	for _, dir := range []string{"./testdata/templates"} {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				templateFiles = append(templateFiles, path)
			}
			return nil
		})
		if err != nil {
			panic(fmt.Errorf("error during loading template files: %s", err))
		}
	}

	output := FancyTemplater{}
	tmpl, err := template.New("").Funcs(sprig.TxtFuncMap()).ParseFiles(templateFiles...)
	if err != nil {
		panic(fmt.Errorf("error during template initialization: %s", err))
	}
	output.rootTemplate = tmpl
	return &output
}

type TestRequest struct {
	templater *FancyTemplater
	Request   string
	Variables string
	Response  string
}

func (t *TestRequest) RequestWithVariables() string {
	jsonRequestWithVariables := fmt.Sprintf(`{%s, %s}`, t.Request, t.Variables)
	if t.IsValidJson(jsonRequestWithVariables) {
		return jsonRequestWithVariables
	}
	return ""
}

func (t *TestRequest) IsValidJson(data string) bool {
	return json.Valid([]byte(data))
}

func (t *TestRequest) ParseRequest(rawRequest string) {
	parsedRequest, _ := t.templater.ParseTemplatedString(rawRequest)
	// queryValue, _ := strings.CutPrefix(rawRequest, queryPrefix)
	jsonFormattedRequest := fmt.Sprintf("{%s}", parsedRequest)
	if !t.IsValidJson(jsonFormattedRequest) {
		panic(fmt.Errorf("test request could not be JSON formatted: %s", parsedRequest))
	}
	jsonFormattedRequest, err := t.templater.JsonIndent(fmt.Sprintf("{%s}", parsedRequest))
	if err != nil {
		panic(err)
	}
	t.Request = strings.TrimSpace(strings.Trim(jsonFormattedRequest, "{}"))
}

func (t *TestRequest) ParseVariables(rawVariables string) {
	parsedVariables, err := t.templater.ParseTemplatedString(rawVariables)
	if err != nil {
		panic(err)
	}

	jsonFormattedVariableObject := strings.TrimSpace(parsedVariables)
	jsonFormattedVariableObject, _ = strings.CutPrefix(jsonFormattedVariableObject, varPrefix)
	if !t.IsValidJson(jsonFormattedVariableObject) {
		panic(fmt.Errorf("test variables could not be JSON formatted: %s", parsedVariables))
	}
	t.Variables = fmt.Sprintf("%s %s", varPrefix, jsonFormattedVariableObject)
}

func (t *TestRequest) ParseResponse(rawResponse string) {
	parsedResponse, err := t.templater.ParseTemplatedString(rawResponse)
	if err != nil {
		panic(err)
	}
	if !t.IsValidJson(parsedResponse) {
		panic(fmt.Errorf("test response could not be JSON formatted: %s", parsedResponse))
	}
	t.Response = parsedResponse
}

func RegisterPaginatedEndpoint(t *testing.T, endpoint string, requests ...TestRequest) string {
	url := fmt.Sprintf("/LOCAL_TESTING/%s", endpoint)
	requestCount := 0
	autopilot.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		GraphQLQueryTemplatedValidation(t, requests[requestCount].RequestWithVariables())(r)
		TemplatedResponse(requests[requestCount].Response)(w)
		requestCount += 1
	})
	return autopilot.Server.URL + url
}

// APaginatedTestClient, TmpPaginatedTestClient, ATestClient(uses *.json, do later)
func BestTestClient(t *testing.T, endpoint string, requests ...TestRequest) *ol.Client {
	if len(requests) > 1 {
		return APaginatedTestClient(t, endpoint, requests...)
	}
	request := requests[0]
	urlOption := ol.SetURL(
		autopilot.RegisterEndpoint(
			fmt.Sprintf("/LOCAL_TESTING/%s", endpoint),
			TemplatedResponse(request.Response),
			GraphQLQueryTemplatedValidation(t, request.RequestWithVariables()),
		),
	)
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), urlOption)
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
