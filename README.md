# Go Application Starter Template

A clean, transport-agnostic application template built with Go. Designed for flexibility - run as HTTP server, CLI, TUI, or MCP with the same core business logic.

## Features

- **Transport-Agnostic Architecture**: Core business logic separate from HTTP/CLI/TUI concerns
- **Chi Router**: Fast HTTP router
- **Turso Database**: SQL database with libsql client
- **SQLC**: Type-safe SQL code generation
- **Gomponents**: Type-safe HTML components
- **HTTPS Support**: Automatic TLS with Let's Encrypt
- **Graceful Shutdown**: Signal handling for clean exits

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Turso database account (for database functionality)

### Configuration

Create a `config.toml` file in your project root:

```toml
[Database]
TursoConnectionString = "libsql://your-database-url.turso.io?authToken=your-auth-token"

[Server]
Port = "8080"
EnableHTTPS = false
HTTPSPort = "443"
TimeOutInSeconds = 60
AllowedCorsURLs = ["http://localhost:3000"]
TaskTimeOutInSeconds = 3600
ServerDomain = "yourdomain.com"

[App]
Environment = "dev"  # Use "prod" for production
```

### Running the Application

```bash
# Build and run
go build -o app ./cmd/main.go
./app

# Or use go run
go run ./cmd/main.go

# With custom config path
go run ./cmd/main.go -c path/to/config.toml
```

## Project Structure

```
├── cmd/
│   └── main.go            # Application entry point
├── internal/
│   ├── app/               # Dependency container (infrastructure)
│   ├── service/           # Business logic (transport-agnostic)
│   ├── handlers/
│   │   └── http/          # HTTP transport layer
│   ├── routes/            # HTTP route definitions
│   ├── database/          # Database access
│   │   ├── queries/       # SQL query files
│   │   ├── repo/          # Generated code from sqlc
│   │   └── schema/        # Database schema
│   ├── config/            # Configuration
│   └── ui/                # UI components
├── config.toml            # Configuration file
└── sqlc.yaml              # SQLC configuration
```

## Architecture

Clean layered architecture with clear separation of concerns:

```
app (infrastructure) → service (business logic) → handlers (transport)
```

- **App** (`internal/app/`): Dependency container holding DB, Logger, Config, Context
- **Service** (`internal/service/`): Transport-agnostic business logic
- **Handlers** (`internal/handlers/`): Transport-specific adapters (HTTP, CLI, TUI, MCP)

This design lets you expose the same business logic through multiple interfaces without code duplication.

## Development Workflow

1. Define schema in `internal/database/schema/`
2. Write queries in `internal/database/queries/`
3. Generate code: `sqlc generate`
4. Add business logic in `internal/service/`
5. Add transport handlers in `internal/handlers/` (HTTP, CLI, etc.)
6. Register HTTP routes in `internal/routes/`

## Technologies Used

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- [Turso Database](https://turso.tech/)
- [SQLC](https://sqlc.dev/)
- [Gomponents](https://github.com/maragudk/gomponents)
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- [autocert](https://golang.org/x/crypto/acme/autocert)
