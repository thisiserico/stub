package httpapi

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/thisiserico/stub/expectation"
	"goji.io"
	"goji.io/pat"
)

const port = ":8080"

// Client allows to run an http API.
type Client struct {
	mux   *goji.Mux
	store expectation.Store
}

// New constructs a new Client.
func New() *Client {
	return &Client{
		mux:   goji.NewMux(),
		store: expectation.NewStore(),
	}
}

// Serve runs an http server with a single handler to catch them all.
func (c *Client) Serve() {
	log.Println("registering stub routes")
	c.mux.HandleFunc(pat.Put("/expectation"), c.registerExpectation)
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

func flattenHeaders(headers map[string][]string) string {
	flattened := make([]string, 0, len(headers))
	for key, values := range headers {
		flattened = append(flattened, fmt.Sprintf("%s:%s", key, values))
	}

	return strings.Join(flattened, " ")
}
