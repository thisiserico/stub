package expectation

var nonExistingResponse = Response{}

// Response defines the actual response for any given endpoint.
type Response struct{}

func ReplyWith(body interface{}) (Response, error) {
	return Response{}, nil
}
