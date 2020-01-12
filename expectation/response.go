package expectation

import "encoding/json"

var nonExistingResponse = Response{}

// Response defines the actual response for any given endpoint.
type Response struct {
	statusCode int
	body       json.RawMessage
	headers    map[string][]string
}

// ReplyWith returns either a valid response or an error.
func ReplyWith(
	code int,
	body json.RawMessage,
	headers map[string][]string,
) (Response, error) {
	return Response{
		statusCode: code,
		body:       body,
		headers:    headers,
	}, nil
}

// MockedStatusCode returns the response status code.
func (r Response) MockedStatusCode() int {
	return r.statusCode
}

// MockedBody returns the response body payload.
func (r Response) MockedBody() json.RawMessage {
	return r.body
}

// MockedHeaders returns the response headers.
func (r Response) MockedHeaders() map[string][]string {
	return r.headers
}
