# Adding New API Endpoints

This guide explains how to add new API endpoints to the MaejicCode application.

## Project Structure Overview

```
.
├── cmd/web/main.go              # Entry point
├── embed.go                     # Embedded frontend files
├── internal/
│   ├── server/server.go         # Server setup and routing
│   └── handlers/
│       ├── api.go               # API handlers
│       └── web.go               # Web handlers
└── frontend/                    # SvelteKit app
```

## Step 1: Add Handler Function

Edit `internal/handlers/api.go` and add your new handler function:

```go
// Example: Add a new endpoint to get users
func (h *APIHandler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Your business logic here
    users := []map[string]string{
        {"id": "1", "name": "John Doe"},
        {"id": "2", "name": "Jane Smith"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
```

## Step 2: Register the Route

In the same file (`internal/handlers/api.go`), add your route to the `ServeHTTP` method:

```go
func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/api/health":
        h.handleHealth(w, r)
    case "/api/users":           // Add your new route here
        h.handleGetUsers(w, r)
    default:
        http.Error(w, "API endpoint not found", http.StatusNotFound)
    }
}
```

## Step 3: Test Your Endpoint

Rebuild and run the application:

```bash
# Windows
build.bat
maejiccode.exe

# Linux/Mac
./build.sh
./maejiccode
```

Test your endpoint:

```bash
curl http://localhost:8080/api/users
```

## Advanced: Adding Route Parameters

For endpoints with parameters (e.g., `/api/users/{id}`), you can parse them manually:

```go
func (h *APIHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
    // Extract ID from path
    // For /api/users/123, this would extract "123"
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 4 {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    userID := parts[3]

    // Your logic here
    user := map[string]string{
        "id":   userID,
        "name": "John Doe",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

Then register it:

```go
func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.URL.Path == "/api/health":
        h.handleHealth(w, r)
    case strings.HasPrefix(r.URL.Path, "/api/users/"):
        h.handleGetUser(w, r)
    case r.URL.Path == "/api/users":
        h.handleGetUsers(w, r)
    default:
        http.Error(w, "API endpoint not found", http.StatusNotFound)
    }
}
```

## Advanced: Using a Router Library

For complex routing, consider adding a third-party router like `chi` or `gorilla/mux`:

```bash
go get github.com/go-chi/chi/v5
```

Then update `internal/server/server.go`:

```go
import "github.com/go-chi/chi/v5"

func (s *Server) setupRoutes() {
    router := chi.NewRouter()

    // API routes
    router.Route("/api", func(r chi.Router) {
        r.Get("/health", handlers.HandleHealth)
        r.Get("/users", handlers.HandleGetUsers)
        r.Get("/users/{id}", handlers.HandleGetUser)
    })

    // Serve frontend (must be last)
    buildFS, _ := fs.Sub(s.frontendFS, "frontend/build")
    router.Handle("/*", http.FileServer(http.FS(buildFS)))

    s.mux = router
}
```

## Example: Complete API Endpoint

Here's a complete example of a POST endpoint that accepts JSON:

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserResponse struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Created string `json:"created"`
}

func (h *APIHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate input
    if req.Name == "" || req.Email == "" {
        http.Error(w, "Name and email are required", http.StatusBadRequest)
        return
    }

    // Create user (this would normally save to a database)
    resp := CreateUserResponse{
        ID:      "123",  // Generate a real ID
        Name:    req.Name,
        Email:   req.Email,
        Created: time.Now().Format(time.RFC3339),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(resp)
}
```

Register it:

```go
case "/api/users":
    if r.Method == http.MethodPost {
        h.handleCreateUser(w, r)
    } else if r.Method == http.MethodGet {
        h.handleGetUsers(w, r)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
```

Test it:

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

## Best Practices

1. **Validation**: Always validate input data
2. **Error Handling**: Return appropriate HTTP status codes
3. **Content-Type**: Set the correct Content-Type header
4. **Documentation**: Document your API endpoints
5. **Testing**: Write tests for your handlers
6. **Security**: Implement authentication/authorization as needed
7. **Logging**: Log important events and errors

## Current Available Endpoints

- `GET /api/health` - Health check endpoint

Add your new endpoints to this list as you create them!
