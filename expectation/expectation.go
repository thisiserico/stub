package expectation

import "strings"

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

func (e expectation) matches(endpoint Endpoint) bool {
	if e.endpoint.method != endpoint.method {
		return false
	}

	if e.endpoint.path != "" && e.endpoint.path != endpoint.path {
		return false
	}

	for _, expectedHeader := range e.endpoint.headers {
		var headerMatches bool
		expectedHeaderStr := strings.Join(expectedHeader, ",")

		for _, givenHeader := range endpoint.headers {
			givenHeaderStr := strings.Join(givenHeader, ",")

			if expectedHeaderStr == givenHeaderStr {
				headerMatches = true
				break
			}
		}

		if !headerMatches {
			return false
		}
	}

	return true
}
