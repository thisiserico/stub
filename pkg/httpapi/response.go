package httpapi

import (
	"log"
	"net/http"
	"strings"

	"github.com/thisiserico/stub/expectation"
)

func (c *Client) fetchResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"fetching response for %s %s %s",
		r.Method,
		r.URL.RequestURI(),
		flattenHeaders(r.Header),
	)

	endpoint, _ := expectation.For(
		strings.ToUpper(r.Method),
		r.URL.RequestURI(),
		r.Header,
	)

	response, err := c.store.Fetch(endpoint)
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	if statusCode := response.MockedStatusCode(); statusCode != 0 {
		w.WriteHeader(statusCode)
	}

	js, err := response.MockedBody().MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
