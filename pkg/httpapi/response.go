package httpapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) fetchResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf("fetching response for %s %s %s", r.Method, r.URL.RequestURI(), flattenHeaders(r))
}

func flattenHeaders(r *http.Request) string {
	flattened := make([]string, 0, len(r.Header))
	for key, values := range r.Header {
		flattened = append(flattened, fmt.Sprintf("%s:%s", key, values))
	}

	return strings.Join(flattened, " ")
}
