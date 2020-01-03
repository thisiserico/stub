package httpapi

import (
	"log"
	"net"
	"net/http"
)

const port = ":8080"

// Client allows to run an http API.
type Client struct{}

// New constructs a new Client.
func New() *Client {
	return &Client{}
}

// Serve runs an http server with a single handler to catch them all.
func (c *Client) Serve() {
	log.Printf("preparing network listener on port %s", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(c.handler),
	}

	log.Println("serving http requests")
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
