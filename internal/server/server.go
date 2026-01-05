package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/kokweikhong/maejiccode/internal/handlers"
)

type Server struct {
	Port       string
	uiFS embed.FS
	mux        *http.ServeMux
}

// New creates a new server instance
func New(uiFS embed.FS) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := &Server{
		Port:       port,
		uiFS: uiFS,
		mux:        http.NewServeMux(),
	}

	s.setupRoutes()
	return s
}

// setupRoutes configures all application routes
func (s *Server) setupRoutes() {
	// Serve embedded frontend files
	buildFS, err := fs.Sub(s.uiFS, "build")
	if err != nil {
		panic(fmt.Sprintf("failed to load frontend files: %v", err))
	}

	// Future API routes will go here
	// Example: s.mux.HandleFunc("/api/v1/...", handlers.APIHandler)
	s.mux.Handle("/api/", handlers.NewAPIHandler())

	// Serve static files with SPA fallback
	s.mux.Handle("/", s.spaHandler(buildFS))
}

// spaHandler serves static files and falls back to index.html for client-side routing
func (s *Server) spaHandler(fsys fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(fsys))
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to open the requested file
		path := r.URL.Path
		if path == "" {
			path = "/"
		}
		
		_, err := fsys.Open(path[1:]) // Remove leading slash for fs.Open
		if err != nil {
			// File doesn't exist, serve index.html for SPA routing
			r.URL.Path = "/"
		}
		
		fileServer.ServeHTTP(w, r)
	})
}

// Start runs the HTTP server
func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.Port, s.mux)
}
