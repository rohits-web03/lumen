package server

// TODO: Define Server struct with configuration and dependencies
// TODO: Implement HTTP handler for serving js files as es modules
// TODO: Implement on-demand transformation of ES modules
// TODO: Implement WebSocket support for HMR
// TODO: Handle static file serving for non-JS assets
import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Server struct {
	port    string
	rootDir string
}

func (s *Server) serveStatic(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir(s.rootDir))
	fs.ServeHTTP(w, r)
}

func (s *Server) serveJS(w http.ResponseWriter, r *http.Request) {
	// get the request js file from disk
	file_path := filepath.Join(s.rootDir, r.URL.Path)

	file_content, err := os.ReadFile(file_path)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	// set content type to application/javascript
	w.Header().Set("Content-Type", "application/javascript")
	// write the file content to the response
	w.Write(file_content)
}

// Start starts the HTTP server and serves html and js files from the example directory.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasSuffix(r.URL.Path, ".js"):
		s.serveJS(w, r)
	default:
		s.serveStatic(w, r)
	}
}

func Start(port string, rootDir string) error {
	srv := &Server{
		port:    port,
		rootDir: rootDir,
	}

	http.Handle("/", srv)

	log.Printf("Lumen dev server running at http://localhost%s", srv.port)
	return http.ListenAndServe(srv.port, nil)
}
