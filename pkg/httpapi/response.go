package httpapi

import (
	"log"
	"net/http"
)

func (c *Client) fetchResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"fetching response for %s %s %s",
		r.Method,
		r.URL.RequestURI(),
		flattenHeaders(r.Header),
	)
}
