package opslevel_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	ol "github.com/opslevel/opslevel-go"
	"github.com/rocktavious/autopilot"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shurcooL/graphql"
)

func TestMain(m *testing.M) {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	log.Logger = log.Output(output)
	flag.Parse()
	teardown := autopilot.Setup()
	defer teardown()
	os.Exit(m.Run())
}

type GraphqlQuery struct {
	Query     string
	Variables map[string]interface{} `json:",omitempty"`
}

func Parse(r *http.Request) GraphqlQuery {
	output := GraphqlQuery{}
	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return output
	}
	json.Unmarshal(bytes, &output)
	return output
}

func LogRaw() autopilot.RequestValidation {
	return func(r *http.Request) {
		defer r.Body.Close()
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Err(err)
		}
		log.Info().Msg(string(bytes))
	}
}

func QueryValidation(t *testing.T, exp string) autopilot.RequestValidation {
	return func(r *http.Request) {
		q := Parse(r)
		autopilot.Equals(t, exp, q.Query)
	}
}

func FixtureQueryValidation(t *testing.T, fixture string) autopilot.RequestValidation {
	return func(r *http.Request) {
		q := Parse(r)
		exp := GraphqlQuery{}
		bytes := []byte(autopilot.Fixture(fixture))
		json.Unmarshal(bytes, &exp)
		autopilot.Equals(t, exp, q)
	}
}

func ATestClient(t *testing.T, endpoint string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetTestURL(autopilot.RegisterEndpoint(fmt.Sprintf("/%s", endpoint),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", endpoint)),
		FixtureQueryValidation(t, fmt.Sprintf("%s_request.json", endpoint)))))
}

func ATestClientAlt(t *testing.T, response string, request string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetTestURL(autopilot.RegisterEndpoint(fmt.Sprintf("/%s__%s", response, request),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", response)),
		FixtureQueryValidation(t, fmt.Sprintf("%s_request.json", request)))))
}

func ATestClientSkipRequest(t *testing.T, endpoint string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetTestURL(autopilot.RegisterEndpoint(fmt.Sprintf("/%s", endpoint),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", endpoint)),
		autopilot.SkipRequestValidation())))
}

func ATestClientLogRequest(t *testing.T, endpoint string) *ol.Client {
	return ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetTestURL(autopilot.RegisterEndpoint(fmt.Sprintf("/%s", endpoint),
		autopilot.FixtureResponse(fmt.Sprintf("%s_response.json", endpoint)),
		LogRaw())))
}

func TestClientQuery(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/account", autopilot.FixtureResponse("account_response.json"), QueryValidation(t, "{account{id}}"))
	client := ol.NewGQLClient(ol.SetAPIToken("x"), ol.SetTestURL(url))
	var q struct {
		Account struct {
			Id graphql.ID
		}
	}
	var v map[string]interface{}
	// Act
	err := client.Query(&q, v)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "1234", q.Account.Id)
}

/*
These tests don't work very well with our autopilot endpoint stuff because they need to make recursive calls
We need to figure out a better way to handle mapping different payloads to different requests

// This test should infinitly recurse on the Service pagination call
func TestClientQueryPagination(t *testing.T) {
	t.Parallel()
	// Arrange
	url := autopilot.RegisterEndpoint("/pagination", autopilot.FixtureResponse("pagination_response.json"), autopilot.SkipRequestValidation())
	client := NewGQLClient("X", SetURL(url))
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	// Act
	go func() {
		_, err := client.ListServices()
		autopilot.Ok(t, err)
		done <- true
	}()

	// Assert
	select {
	case <-timeout:
		// Test Was running infinitely in a pagination recursion - this is a success
	case <-done:
		t.Fatal("TestClientQueryPagination did not infinitely recurse on pagination of Services")
	}
}

// This test should infinitly recurse on the Service.Tags nested pagination call
func TestClientQueryNestedPagination(t *testing.T) {
	t.Parallel()
	// Arrange
	//url := autopilot.RegisterEndpoint("/query_nested_pagination", "query_nested_pagination.json")
	url := autopilot.RegisterEndpoint("/nested_pagination", autopilot.FixtureResponse("nested_pagination_response.json"), FixtureQueryValidation(t, "nested_pagination_request.json"))
	client := NewGQLClient("X", SetURL(url))
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	// Act
	go func() {
		_, err := client.ListServices()
		autopilot.Ok(t, err)
		done <- true
	}()

	// Assert
	select {
	case <-timeout:
		// Test Was running infinitely in a nested pagination recursion - this is a success
	case <-done:
		t.Fatal("TestClientQueryNestedPagination did not infinitely recurse on nested pagination of Service.Tags")
	}
}
*/
