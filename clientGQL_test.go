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

var (
	dataTemplater = NewTestDataTemplater()
	id1           = ol.ID(dataTemplater.ParseValue("id1_string"))
	id2           = ol.ID(dataTemplater.ParseValue("id2_string"))
	id3           = ol.ID(dataTemplater.ParseValue("id3_string"))
	id4           = ol.ID(dataTemplater.ParseValue("id4_string"))
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
	return strings.TrimSpace(response)
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

func NewTestRequest(request string, variables string, response string) TestRequest {
	testRequest := TestRequest{
		Request:   Templated(request),
		Variables: Templated(variables),
		Response:  Templated(response),
	}
	if !strings.HasPrefix(testRequest.Request, "\"") || !strings.HasSuffix(testRequest.Request, "\"") {
		panic(fmt.Errorf("testRequest Request should be wrapped in quotes: '%s'", testRequest.Request))
	}
	if !IsValidJson(testRequest.Variables) {
		panic(fmt.Errorf("testRequest Variables is not valid json: '%s'", testRequest.Variables))
	}
	jsonRequestWithVariables := fmt.Sprintf(
		`{"query": %s, "variables": %s}`,
		testRequest.Request,
		testRequest.Variables,
	)
	if !IsValidJson(jsonRequestWithVariables) {
		panic(fmt.Errorf("testRequest Request with Variables is not valid json: '%s'", jsonRequestWithVariables))
	}
	if !IsValidJson(testRequest.Response) {
		panic(fmt.Errorf("testRequest Response is not json: '%s'", testRequest.Response))
	}
	return testRequest
}

func NewTestDataTemplater(templateDirs ...string) *TestDataTemplater {
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

	output := TestDataTemplater{}
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"WrapWithCurlyBrackets": func(value string) string { return "{ " + value + " }" },
	})
	tmpl.Funcs(sprig.TxtFuncMap())
	tmpl, err := tmpl.ParseFiles(templateFiles...)
	if err != nil {
		panic(fmt.Errorf("error during template initialization: %s", err))
	}
	output.rootTemplate = tmpl
	return &output
}

type TestDataTemplater struct {
	rootTemplate *template.Template
}

func (t *TestDataTemplater) ParseValue(value string) string {
	return t.ParseTemplatedString(`{{ template "` + value + `" }}`)
}

func (t *TestDataTemplater) ParseTemplatedString(contents string) string {
	target, err := t.rootTemplate.Parse(contents)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %s", err))
	}
	data := bytes.NewBuffer([]byte{})
	if err = target.Execute(data, nil); err != nil {
		panic(err)
	}
	return strings.TrimSpace(data.String())
}

type TestRequest struct {
	Request   string
	Variables string
	Response  string
}

func IsValidJson(data string) bool {
	return json.Valid([]byte(data))
}

func RegisterPaginatedEndpoint(t *testing.T, endpoint string, requests ...TestRequest) string {
	url := fmt.Sprintf("/LOCAL_TESTING/%s", endpoint)
	requestCount := 0
	autopilot.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		GraphQLQueryTemplatedValidation(t, fmt.Sprintf(
			`{"query": %s, "variables": %s}`,
			requests[requestCount].Request,
			requests[requestCount].Variables),
		)(r)
		TemplatedResponse(requests[requestCount].Response)(w)
		requestCount += 1
	})
	return autopilot.Server.URL + url
}

func BestTestClient(t *testing.T, endpoint string, requests ...TestRequest) *ol.Client {
	url := RegisterPaginatedEndpoint(t, endpoint, requests...)
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(url))
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
