# MaejicCode

A Go web application with embedded SvelteKit frontend, bundled into a single executable.

## Features

- 🚀 Fast Go backend
- ⚡ Modern SvelteKit frontend
- 📦 Single executable bundle
- 🎨 Responsive design
- 🔥 Hot reload in development

## Pages

- **Home**: Welcome page with feature highlights
- **Services**: Overview of available services
- **About Us**: Company information and values
- **Contact Us**: Contact form and information

## Prerequisites

- Go 1.25.5 or higher
- Node.js 18+ and npm

## Development

### Setup

1. Clone the repository
2. Install frontend dependencies:
   ```bash
   cd frontend
   npm install
   cd ..
   ```

### Running in Development Mode

1. Start the SvelteKit dev server:

   ```bash
   cd frontend
   npm run dev
   ```

   The frontend will be available at http://localhost:5173

2. In another terminal, run the Go server (after building frontend at least once):
   ```bash
   go run ./cmd/web
   ```

### Building for Production

#### Windows

```cmd
build.bat
```

#### Linux/Mac

```bash
chmod +x build.sh
./build.sh
```

This will:

1. Install frontend dependencies
2. Build the SvelteKit app as static files
3. Compile the Go application with embedded frontend
4. Create a single executable (`maejiccode.exe` on Windows, `maejiccode` on Linux/Mac)

## Running the Application

After building, simply run the executable:

```bash
# Windows
maejiccode.exe

# Linux/Mac
./maejiccode
```

The application will start on http://localhost:8080

You can change the port by setting the `PORT` environment variable:

```bash
# Windows
set PORT=3000
maejiccode.exe

# Linux/Mac
PORT=3000 ./maejiccode
```

## Project Structure

```
.
├── cmd/
│   └── web/
│       └── main.go            # Application entry point
├── internal/
│   ├── server/
│   │   └── server.go          # HTTP server setup and routing
│   └── handlers/
│       ├── api.go             # API handlers (health check, future endpoints)
│       └── web.go             # Web-specific handlers
├── frontend/                  # SvelteKit frontend
│   ├── src/
│   │   ├── routes/           # Page components
│   │   │   ├── +page.svelte              # Home page
│   │   │   ├── +layout.svelte            # Layout wrapper
│   │   │   ├── services/+page.svelte     # Services page
│   │   │   ├── about/+page.svelte        # About page
│   │   │   └── contact/+page.svelte      # Contact page
│   │   ├── app.html          # HTML template
│   │   └── app.css           # Global styles
│   ├── static/               # Static assets (favicon, etc.)
│   ├── package.json          # Node dependencies
│   ├── svelte.config.js      # SvelteKit config
│   └── vite.config.js        # Vite config
├── go.mod                    # Go module definition
├── build.bat                 # Windows build script
├── build.sh                  # Linux/Mac build script
└── README.md                 # This file
```

### Directory Explanation

- **cmd/web/**: Main application entry point
- **internal/**: Private application code
  - **server/**: HTTP server configuration and routing
  - **handlers/**: HTTP request handlers (separated by concern)
- **frontend/**: SvelteKit frontend application (embedded at build time)

## API Endpoints

The application includes a placeholder API structure for future development:

### Available Endpoints

- `GET /api/health` - Health check endpoint
  ```json
  {
    "status": "ok",
    "message": "Server is running"
  }
  ```

### Adding New API Endpoints

To add new API endpoints, edit [internal/handlers/api.go](internal/handlers/api.go):

```go
func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/api/health":
        h.handleHealth(w, r)
    case "/api/your-endpoint":  // Add your new endpoint here
        h.handleYourEndpoint(w, r)
    default:
        http.Error(w, "API endpoint not found", http.StatusNotFound)
    }
}
```

## Technology Stack

- **Backend**: Go 1.25.5
- **Frontend**: SvelteKit 2.0
- **Build Tool**: Vite 5.0
- **Styling**: CSS with scoped styles

## License

MIT
