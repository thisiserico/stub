package expectation

import (
	"fmt"
	"net/http"
	"strings"
)

// Endpoint encapsulates client requests.
type Endpoint struct {
	method  string
	path    string
	headers map[string][]string
}

// For returns either a valid endpoint or an error.
func For(method, path string, headers map[string][]string) (Endpoint, error) {
	if !isAllowedHTTPMethod(method) {
		return Endpoint{}, fmt.Errorf("forbidden method %s", method)
	}

	sanitizedHeaders := make(map[string][]string)
	for header, values := range headers {
		sanitizedValues := make([]string, 0, len(values))
		for _, value := range values {
			sanitizedValues = append(sanitizedValues, strings.ToLower(value))
		}

		sanitizedHeaders[strings.ToLower(header)] = sanitizedValues
	}

	return Endpoint{
		method:  method,
		path:    path,
		headers: sanitizedHeaders,
	}, nil
}

func isAllowedHTTPMethod(method string) bool {
	allowedMethods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	}

	for _, allowedMethod := range allowedMethods {
		if allowedMethod == method {
			return true
		}
	}

	return method == ""
}
