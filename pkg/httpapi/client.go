package httpapi

import (
	"log"
	"net"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

const port = ":8080"

// Client allows to run an http API.
type Client struct {
	mux *goji.Mux
}

// New constructs a new Client.
func New() *Client {
	return &Client{
		mux: goji.NewMux(),
	}
}

// Serve runs an http server with a single handler to catch them all.
func (c *Client) Serve() {
	log.Println("registering stub routes")
	c.mux.HandleFunc(pat.Post("/expectation"), c.registerExpectation)
	c.mux.HandleFunc(pat.New("/*"), c.fetchResponse)

	log.Printf("preparing network listener on port %s", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    port,
		Handler: c.mux,
	}

	log.Println("serving http requests")
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
