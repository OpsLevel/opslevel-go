package opslevel

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

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
	Query     string                 `json`
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

func TestClientQuery(t *testing.T) {
	// Arrange
	url := autopilot.RegisterEndpoint("/account", autopilot.FixtureResponse("account_response.json"), QueryValidation(t, "{account{id}}"))
	client := NewClient("X", SetURL(url))
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

// This test should infinitly recurse on the Service pagination call
func TestClientQueryPagination(t *testing.T) {
	t.Parallel()
	// Arrange
	url := autopilot.RegisterEndpoint("/pagination", autopilot.FixtureResponse("pagination_response.json"), autopilot.SkipRequestValidation())
	// url := RegisterEndpointWithRequestValidation("/query_pagination", "query_pagination.json", func(r *http.Request) { log.Info().Msg(ReadBody(r)) })
	client := NewClient("X", SetURL(url))
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

