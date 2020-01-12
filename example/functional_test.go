package functional

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/lucsky/cuid"
)

const (
	correlationIDHeader = "x-correlation-id"
	knownPath           = "/target/path/in/my/service"
	knownResponse       = "known response"
)

var serviceAddress = "http://localhost:8080"

func TestUsingAGetRequest(t *testing.T) {
	t.Run("the service is not ready for", func(t *testing.T) {
		test := prepareTestHandler(t)

		test.whenHittingAGetEndpoint()
		test.thenANotImplementedErrorIsReturned()
	})

	t.Run("the service prepared for", func(t *testing.T) {
		test := prepareTestHandler(t)

		test.givenAMockedGetEndpoint()
		test.whenHittingAGetEndpoint()
		test.thenTheExpectedResponseIsReturned()
	})
}

type testHandler struct {
	*testing.T

	correlationID  string
	respStatusCode int
	respBody       interface{}
}

func prepareTestHandler(t *testing.T) *testHandler {
	return &testHandler{
		T:             t,
		correlationID: cuid.New(),
	}
}

type expectationPayload struct {
	Method     string              `json:"using_method"`
	Path       string              `json:"against_path"`
	ReqHeaders map[string][]string `json:"with_headers"`

	StatusCode int         `json:"returns_code"`
	Response   interface{} `json:"with_response"`
}

func (t *testHandler) givenAMockedGetEndpoint() {
	js, _ := json.Marshal(expectationPayload{
		Method:     http.MethodGet,
		Path:       knownPath,
		ReqHeaders: map[string][]string{correlationIDHeader: {t.correlationID}},

		StatusCode: http.StatusPartialContent,
		Response:   knownResponse,
	})

	uri := fmt.Sprintf("%s/expectation", serviceAddress)
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(js))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	t.respStatusCode = resp.StatusCode
}

func (t *testHandler) whenHittingAGetEndpoint() {
	uri := fmt.Sprintf("%s%s", serviceAddress, knownPath)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(correlationIDHeader, t.correlationID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var respBody interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&respBody); err == nil {
		t.respBody = respBody
	}

	t.respStatusCode = resp.StatusCode
}

func (t *testHandler) thenANotImplementedErrorIsReturned() {
	if t.respStatusCode != http.StatusNotImplemented {
		t.Fatalf("unexpected status code, want %d, got %d", http.StatusNotImplemented, t.respStatusCode)
	}
}

func (t *testHandler) thenTheExpectedResponseIsReturned() {
	if t.respStatusCode != http.StatusPartialContent {
		t.Fatalf("unexpected status code, want %d, got %d", http.StatusPartialContent, t.respStatusCode)
	}

	got := t.respBody.(string)
	if knownResponse != got {
		t.Fatalf("unexpected response, want %s, got %s", knownResponse, got)
	}
}
