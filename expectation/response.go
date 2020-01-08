package expectation

import "encoding/json"

var nonExistingResponse = Response{}

// Response defines the actual response for any given endpoint.
type Response struct {
	body json.RawMessage
}

// ReplyWith returns either a valid response or an error.
func ReplyWith(body json.RawMessage) (Response, error) {
	return Response{
		body: body,
	}, nil
}

// MockedBody returns the response body payload.
func (r Response) MockedBody() json.RawMessage {
	return r.body
}
