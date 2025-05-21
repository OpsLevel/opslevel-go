package opslevel_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/hasura/go-graphql-client"
	ol "github.com/opslevel/opslevel-go/v2025"
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
	alias1        = dataTemplater.ParseValue("alias1")
	alias2        = dataTemplater.ParseValue("alias2")
	alias3        = dataTemplater.ParseValue("alias3")
	alias4        = dataTemplater.ParseValue("alias4")
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
		_, err := fmt.Fprint(w, Templated(response))
		if err != nil {
			panic(err)
		}
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

func NewTestDataTemplater() *TestDataTemplater {
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

func (testDataTemplater *TestDataTemplater) ParseValue(value string) string {
	return testDataTemplater.ParseTemplatedString(`{{ template "` + value + `" }}`)
}

func (testDataTemplater *TestDataTemplater) ParseTemplatedString(contents string) string {
	target, err := testDataTemplater.rootTemplate.Parse(contents)
	if err != nil {
		panic(fmt.Errorf("error parsing template: %s", err))
	}
	data := bytes.NewBuffer([]byte{})
	if err = target.Execute(data, nil); err != nil {
		panic(err)
	}
	return strings.TrimSpace(data.String())
}

func stripWhitespace(input string) string {
	// Remove newlines, tabs, and spaces
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "  ", "")
	return input
}

func BestTestClient(t *testing.T, endpoint string, requests ...autopilot.TestRequest) *ol.Client {
	for i, request := range requests {
		requests[i].Request.Query = stripWhitespace(request.Request.Query)
	}
	urlToRegister := fmt.Sprintf("/LOCAL_TESTING/%s", endpoint)
	registeredUrl := autopilot.RegisterPaginatedEndpoint(t, urlToRegister, requests...)
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetMaxRetries(0), ol.SetURL(registeredUrl))
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

func httpResponseByCode(statusCode int) autopilot.ResponseWriter {
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
	}
}

func TestMissingTeamIsAnOpsLevelApiError(t *testing.T) {
	testRequest := autopilot.NewTestRequest(
		`query TeamGet($id:ID!){account{team(id: $id){alias,id,aliases,managedAliases,contacts{address,displayName,displayType,externalId,id,isDefault,type},htmlUrl,manager{id,email,htmlUrl,name,provisionedBy,role},memberships{nodes{role,team{alias,id},user{id,email}},{{ template "pagination_request" }}},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }}}}}}`,
		`{ {{ template "id1" }} }`,
		`{"errors": [
				{
					"message": "Team with id '{{ template "id1_string" }}' does not exist on this account",
					"path": ["account", "team"],
					"locations": [{"line": 1, "column": 32}]
				}
     ]}`,
	)
	client := BestTestClient(t, "team/missing_team", testRequest)
	// Act
	_, err := client.GetTeam(id1)
	// Assert
	autopilot.Equals(t, true, ol.IsOpsLevelApiError(err))
}

func Test404ResponseNotAnOpsLevelApiError(t *testing.T) {
	// Arrange
	headers := map[string]string{"x": "x"}
	request := `{ "query": "{account{id}}", "variables":{} }`
	url := autopilot.RegisterEndpoint("/LOCAL_TESTING/test_404",
		httpResponseByCode(http.StatusNotFound),
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
	autopilot.Equals(t, false, ol.IsOpsLevelApiError(err))
}

func TestGenericHasuraErrorNotAnOpsLevelApiError(t *testing.T) {
	// Arrange
	headers := map[string]string{"x": "x"}
	request := `{ "query": "{account{id}}", "variables":{} }`
	url := autopilot.RegisterEndpoint("/LOCAL_TESTING/test_200",
		httpResponseByCode(http.StatusOK),
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
	_, isHasuraError := err.(graphql.Errors)
	autopilot.Equals(t, true, isHasuraError)
	autopilot.Equals(t, false, ol.IsOpsLevelApiError(err))
}

func TestGenericErrorIsNotAnOpsLevelApiError(t *testing.T) {
	autopilot.Equals(t, false, ol.IsOpsLevelApiError(errors.New("asdf")))
}

func TestNilErrorIsNotAnOpsLevelApiError(t *testing.T) {
	autopilot.Equals(t, false, ol.IsOpsLevelApiError(nil))
}
