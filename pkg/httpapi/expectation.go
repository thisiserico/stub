package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/thisiserico/stub/expectation"
)

type expectationRequest struct {
	Method   string          `json:"method"`
	Path     string          `json:"path"`
	Headers  headers         `json:"headers"`
	Response json.RawMessage `json:"response"`
}

type headers map[string][]string

func (c *Client) registerExpectation(w http.ResponseWriter, r *http.Request) {
	var req expectationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid expectation payload"))
		return
	}
	r.Body.Close()

	log.Printf(
		"registering new expectation for %s %s %s",
		strings.ToUpper(req.Method),
		req.Path,
		flattenHeaders(req.Headers),
	)

	endpoint, err := expectation.For(
		strings.ToUpper(req.Method),
		req.Path,
		req.Headers,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid expectation"))
		return
	}

	response, err := expectation.ReplyWith(req.Response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid expected response"))
		return
	}

	if err := c.store.Register(endpoint, response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed registering the given expectation"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
