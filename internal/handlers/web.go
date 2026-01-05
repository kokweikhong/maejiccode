package handlers

import "net/http"

// WebHandler handles web-specific requests
// Add any custom web handlers here in the future
type WebHandler struct{}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

func (h *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Custom web handling logic can go here
	http.NotFound(w, r)
}
