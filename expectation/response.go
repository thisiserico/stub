package expectation

var nonExistingResponse = Response{}

// Response defines the actual response for any given endpoint.
type Response struct {
	body []byte
}

// ReplyWith returns either a valid response or an error.
func ReplyWith(body []byte) (Response, error) {
	return Response{
		body: body,
	}, nil
}

// MockedBody returns the response body payload.
func (r Response) MockedBody() []byte {
	return r.body
}
