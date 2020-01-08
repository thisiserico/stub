package expectation

import (
	"errors"
	"sync"
)

// Store provides utilities to register and fetch responses for endpoints.
type Store interface {
	Register(Endpoint, Response) error
	Fetch(Endpoint) (Response, error)
}

// NewStore initializes a new response in memory store.
func NewStore() Store {
	return &store{
		Mutex:        &sync.Mutex{},
		expectations: make([]expectation, 0),
	}
}

type store struct {
	*sync.Mutex

	expectations []expectation
}

func (s *store) Register(e Endpoint, r Response) error {
	s.Lock()
	defer s.Unlock()

	// TODO check whether it already exists, overwrite if so.
	s.expectations = append(s.expectations, responseForEndpoint(e, r))

	return nil
}

func (s *store) Fetch(e Endpoint) (Response, error) {
	for _, expectation := range s.expectations {
		if expectation.matches(e) {
			return expectation.response, nil
		}
	}

	return nonExistingResponse, errors.New("unknown endpoint")
}
