package expectation

type expectation struct {
	endpoint Endpoint
	response Response
}

func responseForEndpoint(e Endpoint, r Response) expectation {
	return expectation{
		endpoint: e,
		response: r,
	}
}
