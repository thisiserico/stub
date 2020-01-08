package expectation

// Endpoint encapsulates client requests.
type Endpoint struct{}

func For(method, path string, headers map[string][]string) (Endpoint, error) {
	return Endpoint{}, nil
}
