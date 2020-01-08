package httpapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (c *Client) fetchResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"fetching response for %s %s %s",
		r.Method,
		r.URL.RequestURI(),
		flattenHeaders(r.Header),
	)
}

func flattenHeaders(headers map[string][]string) string {
	flattened := make([]string, 0, len(headers))
	for key, values := range headers {
		flattened = append(flattened, fmt.Sprintf("%s:%s", key, values))
	}

	return strings.Join(flattened, " ")
}
