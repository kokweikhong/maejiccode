package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// APIHandler handles API requests
type APIHandler struct{}

// ContactRequest represents the contact form submission
type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// ContactResponse represents the API response
type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NewAPIHandler creates a new API handler
func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

// ServeHTTP implements http.Handler interface
func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	switch r.URL.Path {
	case "/api/health":
		h.handleHealth(w, r)
	case "/api/contact":
		h.handleContact(w, r)
	default:
		// Return 404 for unimplemented API routes
		http.Error(w, "API endpoint not found", http.StatusNotFound)
	}
}

// handleHealth returns the health status of the application
func (h *APIHandler) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"message": "Server is running",
	})
}

// handleContact processes contact form submissions
func (h *APIHandler) handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Message: "Invalid request format",
		})
		return
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Message == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ContactResponse{
			Success: false,
			Message: "Name, email, and message are required",
		})
		return
	}

	// Log the contact submission (in a real app, you'd save to database or send email)
	fmt.Printf("[%s] Contact form submission:\n", time.Now().Format(time.RFC3339))
	fmt.Printf("  Name: %s\n", req.Name)
	fmt.Printf("  Email: %s\n", req.Email)
	fmt.Printf("  Subject: %s\n", req.Subject)
	fmt.Printf("  Message: %s\n", req.Message)
	fmt.Println("---")

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ContactResponse{
		Success: true,
		Message: "Thank you for your message! We'll get back to you soon.",
	})
}
