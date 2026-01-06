package server

// TODO: Define Server struct with configuration and dependencies
// TODO: Implement HTTP handler for serving js files as es modules
// TODO: Implement on-demand transformation of ES modules
// TODO: Implement WebSocket support for HMR
// TODO: Handle static file serving for non-JS assets
import (
	"log"
	"net/http"
)

const (
	port    = ":3000"
	rootDir = "example"
)

// Start starts the HTTP server and serves html from the example directory.
func Start() error {
	fs := http.FileServer(http.Dir(rootDir))
	http.Handle("/", fs)

	log.Printf("Lumen dev server running at http://localhost%s", port)
	return http.ListenAndServe(port, nil)
}
