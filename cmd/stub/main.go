package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thisiserico/stub/pkg/httpapi"
)

func main() {
	ctx := prepareShutdown()

	apiClient := httpapi.New()
	go apiClient.Serve()

	<-ctx.Done()
}

func prepareShutdown() context.Context {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Printf("stop signal captured: %v", <-stop)
		cancel()
	}()

	return ctx
}
