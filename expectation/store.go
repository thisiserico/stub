package expectation

import (
	"errors"
	"sync"
)

const expectationNotFound = -1

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

	if index := s.findExpectation(e); index != expectationNotFound {
		s.expectations[index] = responseForEndpoint(e, r)
		return nil
	}

	s.expectations = append(s.expectations, responseForEndpoint(e, r))
	return nil
}

func (s *store) Fetch(e Endpoint) (Response, error) {
	if index := s.findExpectation(e); index != expectationNotFound {
		return s.expectations[index].response, nil
	}

	return nonExistingResponse, errors.New("unknown endpoint")
}

func (s *store) findExpectation(e Endpoint) int {
	for i, expectation := range s.expectations {
		if expectation.matches(e) {
			return i
		}
	}

	return expectationNotFound
}
