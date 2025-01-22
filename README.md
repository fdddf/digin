# Go Gin/Dig Modular Web Service

A production-ready web service template using Gin web framework and Uber's Dig for dependency injection, featuring automated module registration and scalable architecture.

## Features

- **Modular Architecture**: Domain-driven design with isolated modules (users, products, etc.)
- **Dependency Injection**: Powered by Uber's Dig for clean dependency management
- **Code Generation**: Auto-registration of modules via build-time code generation
- **Gin Framework**: High-performance HTTP routing and middleware support
- **Configuration Management**: Environment-aware configuration loading
- **Route Groups**: Versioned API endpoints with middleware support

## Getting Started

### Prerequisites

- Go 1.21+
- Go modules enabled
- `dig` and `gin` packages:
  ```bash
  go get go.uber.org/dig
  go get github.com/gin-gonic/gin
  ```

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/fdddf/digin.git myapp
   cd myapp
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Generate module registration code:
   ```bash
   go generate ./...
   ```

4. Build and run:
   ```bash
   go build -o myapp ./cmd/server/
   ./myapp
   ```

## Project Structure

```text
myapp/
├── cmd/
│   └── server/          # Main application entrypoint
├── internal/
│   ├── config/          # Configuration management
│   ├── core/            # DI container setup
│   ├── modules/         # Domain modules (users, products, etc.)
│   └── gen/             # Auto-generated code (do not edit manually)
├── scripts/             # Code generation utilities
├── go.mod
└── go.sum
```

## Code Generation

The project uses automatic module registration via code generation:

1. **Module Convention**:
   - Each domain package contains a `module.go` file
   - Implements `Register(*dig.Container)` function
   - Located in `internal/modules/<domain>/`

2. **Generate Code**:
   ```bash
   # Manual generation
   go run scripts/generate_modules.go
   
   # Using go generate
   go generate ./...
   ```

3. **Example Module** (`internal/modules/user/module.go`):
   ```go
   package user

   import "go.uber.org/dig"

   func Register(container *dig.Container) {
       container.Provide(NewService)
       container.Provide(NewHandler)
       container.Provide(NewUserRoutes, dig.Group("routes"))
   }
   ```

## API Examples

```bash
# Get user by ID
curl http://localhost:8080/api/v1/users/123

# Create product
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name":"New Product", "price":99.99}'
```

## Testing

Run the test suite:
```bash
go test -v ./...
```

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details

---

**Acknowledgments**
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Uber Dig](https://github.com/uber-go/dig)
```
