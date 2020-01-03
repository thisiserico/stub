package httpapi

import (
	"log"
	"net/http"
)

func (c *Client) registerExpectation(w http.ResponseWriter, r *http.Request) {
	log.Println("registering new expectation")

	w.WriteHeader(http.StatusCreated)
}
