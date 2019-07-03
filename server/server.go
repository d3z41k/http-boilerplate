package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
)

// HelloWorld is a sample handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// NewRouter returns a new HTTP handler that implements the main server routes
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Set up our root handlers
	r.Get("/", HelloWorld)

	// Set up static file serving
	staticPath, _ := filepath.Abs("static")
	// fs := http.FileServer(unindexed.Dir(staticPath)) // safe static distrib
	fs := http.FileServer(http.Dir(staticPath))
	r.Handle("/*", fs)

	return r
}
