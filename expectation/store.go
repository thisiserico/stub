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
		Mutex:     &sync.Mutex{},
		endpoints: make(map[Endpoint]Response),
	}
}

type store struct {
	*sync.Mutex

	endpoints map[Endpoint]Response
}

func (s *store) Register(e Endpoint, r Response) error {
	s.Lock()
	defer s.Unlock()

	// TODO check whether it already exists, overwrite if so.

	s.endpoints[e] = r
	return nil
}

func (s *store) Fetch(e Endpoint) (Response, error) {
	return nonExistingResponse, errors.New("unknown endpoint")
}
