package opslevel_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	ol "github.com/opslevel/opslevel-go/v2025"
	"github.com/rocktavious/autopilot/v2023"
)

const testRestClientResponse = `{ "result": "Hello World!" }`

func testRestClientResponseWriter() autopilot.ResponseWriter {
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, testRestClientResponse)
	}
}

func ATestRestClient(endpoint string) *resty.Client {
	return ol.NewRestClient(ol.SetURL(
		autopilot.RegisterEndpoint(
			fmt.Sprintf("/%s", endpoint),
			testRestClientResponseWriter(),
			autopilot.SkipRequestValidation(),
		),
	),
	)
}

func TestRestClientQuery(t *testing.T) {
	// Arrange
	client := ATestRestClient("rest/example")
	resp := &ol.RestResponse{}
	// Act
	_, err := client.R().SetResult(resp).Get("")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Hello World!", resp.Result)
}
