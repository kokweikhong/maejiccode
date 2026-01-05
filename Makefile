.PHONY: help ui-install ui-dev ui-build ui-preview ui-check ui-check-watch ui-sync ui-clean ui-clean-all
.PHONY: go-build go-run go-test go-dev go-clean go-tidy go-fmt go-vet
.PHONY: build run dev clean all

# Default target
help:
	@echo "Available targets:"
	@echo ""
	@echo "UI targets:"
	@echo "  ui-install      - Install UI npm dependencies"
	@echo "  ui-dev          - Run UI development server"
	@echo "  ui-build        - Build UI for production"
	@echo "  ui-preview      - Preview UI production build"
	@echo "  ui-sync         - Sync SvelteKit files"
	@echo "  ui-check        - Run UI type checking"
	@echo "  ui-check-watch  - Run UI type checking in watch mode"
	@echo "  ui-clean        - Remove UI build artifacts"
	@echo "  ui-clean-all    - Remove UI build artifacts and node_modules"
	@echo ""
	@echo "Go targets:"
	@echo "  go-build        - Build Go application"
	@echo "  go-run          - Run Go application"
	@echo "  go-test         - Run Go tests"
	@echo "  go-dev          - Run Go with hot reload (air)"
	@echo "  go-clean        - Remove Go build artifacts"
	@echo "  go-tidy         - Tidy Go dependencies"
	@echo "  go-fmt          - Format Go code"
	@echo "  go-vet          - Run go vet"
	@echo ""
	@echo "Combined targets:"
	@echo "  build           - Build both UI and Go"
	@echo "  run             - Run Go application"
	@echo "  dev             - Run both UI dev server and Go with hot reload"
	@echo "  clean           - Clean both UI and Go artifacts"
	@echo "  all             - Install dependencies and build everything"

# Install UI dependencies
ui-install:
	cd ui && npm install

# Run UI development server
ui-dev:
	cd ui && npm run dev

# Build UI for production
ui-build:
	cd ui && npm run build

# Preview UI production build
ui-preview:
	cd ui && npm run preview

# Sync SvelteKit files
ui-sync:
	cd ui && npm run prepare

# Run UI type checking
ui-check:
	cd ui && npm run check

# Run UI type checking in watch mode
ui-check-watch:
	cd ui && npm run check:watch

# Clean UI build artifacts
ui-clean:
	cd ui && rm -rf build .svelte-kit

# Clean UI everything including node_modules
ui-clean-all: ui-clean
	cd ui && rm -rf node_modules

# Go targets

# Build Go application
go-build: ui-build
	go build -o bin/maejiccode.exe ./cmd/web

# Run Go application
go-run:
	go run ./cmd/web/main.go

# Run Go tests
go-test:
	go test -v ./...

# Run Go with hot reload (requires air: go install github.com/air-verse/air@latest)
go-dev:
	@where air >nul 2>nul && air || (echo Air not found. Install with: go install github.com/air-verse/air@latest && echo Running without hot reload... && go run ./cmd/web/main.go)

# Clean Go build artifacts
go-clean:
	rm -rf bin/

# Tidy Go dependencies
go-tidy:
	go mod tidy

# Format Go code
go-fmt:
	go fmt ./...

# Run go vet
go-vet:
	go vet ./...

# Combined targets

# Build everything
build: ui-build go-build

# Run the application
run: go-run

# Run development environment (UI dev server and Go with hot reload in parallel)
dev:
	@echo "Starting development environment..."
	@echo "Run 'make ui-dev' in one terminal and 'make go-dev' in another"
	@echo "Or use: start cmd /k \"make ui-dev\" && make go-dev"

# Clean everything
clean: ui-clean go-clean

# Install and build everything
all: ui-install ui-build go-tidy go-build
