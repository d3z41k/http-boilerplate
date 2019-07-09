package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"projects/http-boilerplate/server"
)

func main() {
	// This is the domain the server should accept connections for.
	// domains := []string{"example.com", "www.example.com"}
	handler := server.NewRouter()
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	go srv.ListenAndServe()
	// go srv.Serve(autocert.NewListener(domains...))

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
