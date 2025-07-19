# SyntaxRush Makefile
# Common development and build tasks

.PHONY: build install clean test dev help global-install completion

# Default target
help:
	@echo "🚀 SyntaxRush Development Commands"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@echo "Development:"
	@echo "  build         Build the binary locally"
	@echo "  dev           Build and run with Go sample"
	@echo "  test          Run tests (when available)"
	@echo "  clean         Clean build artifacts"
	@echo ""
	@echo "Installation:"
	@echo "  install       Interactive installation menu"
	@echo "  global-install Global installation to /usr/local/bin"
	@echo "  completion    Setup shell auto-completion"
	@echo ""
	@echo "Maintenance:"
	@echo "  fmt           Format all Go code"
	@echo "  lint          Run linting (requires golangci-lint)"
	@echo "  mod-tidy      Clean up go.mod and go.sum"

# Build the binary
build:
	@echo "🔨 Building SyntaxRush..."
	@go build -ldflags="-s -w" -o syntaxrush .
	@echo "✅ Build complete: ./syntaxrush"

# Development build and test
dev: build
	@echo "🚀 Starting SyntaxRush with Go sample..."
	@./syntaxrush practice go

# Interactive installation
install:
	@./install.sh

# Global installation
global-install:
	@./scripts/install-global.sh

# Setup shell completion
completion:
	@./scripts/setup-completion.sh

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -f syntaxrush syntaxrush.exe
	@rm -f syntaxrush_completion.*
	@echo "✅ Clean complete"

# Format Go code
fmt:
	@echo "📝 Formatting Go code..."
	@go fmt ./...
	@echo "✅ Code formatted"

# Run linting (requires golangci-lint)
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		echo "🔍 Running linter..."; \
		golangci-lint run; \
		echo "✅ Linting complete"; \
	else \
		echo "⚠️  golangci-lint not installed. Install with:"; \
		echo "   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.54.2"; \
	fi

# Tidy go modules
mod-tidy:
	@echo "📦 Tidying Go modules..."
	@go mod tidy
	@echo "✅ Modules tidied"

# Run tests (placeholder for future)
test:
	@echo "🧪 Running tests..."
	@go test ./... -v
	@echo "✅ Tests complete"

# Development cycle: clean, build, test
dev-cycle: clean fmt mod-tidy build
	@echo "🔄 Development cycle complete"

# Release preparation
release-prep: clean fmt mod-tidy test build
	@echo "📦 Release preparation complete"
	@echo "✅ Ready for release"

# Quick build for CI/CD
ci-build:
	@go mod download
	@go build -ldflags="-s -w" -o syntaxrush .

# Display project info
info:
	@echo "📊 SyntaxRush Project Information"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "Go Version: $$(go version)"
	@echo "Module: $$(head -1 go.mod)"
	@echo "Files: $$(find . -name '*.go' | wc -l) Go files"
	@echo "Size: $$(du -sh . | cut -f1) total"
	@if [ -f syntaxrush ]; then echo "Binary: $$(ls -lh syntaxrush | awk '{print $$5}')"; fi
