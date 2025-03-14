# Go Web Application Template

A modern, scalable web application template built with Go. This template provides a solid foundation for developing web applications with a clean architecture, database integration, and responsive UI components.

## Features

- **Chi Router**: Fast and lightweight HTTP router for Go
- **Turso Database**: SQL database with libsql client integration
- **SQLC**: Type-safe SQL in Go with generated code
- **Gomponents**: Declarative HTML views in Go
- **Clean Configuration**: Environment-based configuration with cleanenv
- **HTTPS Support**: Automatic TLS certificate management with Let's Encrypt
- **CORS Support**: Configurable Cross-Origin Resource Sharing
- **Structured Logging**: Logging with slog
- **Graceful Shutdown**: Proper server shutdown handling

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Turso database account (for database functionality)

### Installation

You can use the `gonew` tool to create a new project based on this template:

```bash
# Install gonew if you haven't already
go install golang.org/x/tools/cmd/gonew@latest

# Create a new project based on this template
gonew github.com/mhpenta/app github.com/yourusername/your-project
```

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
│   ├── config/            # Configuration handling
│   ├── database/          # Database connection and queries
│   │   ├── queries/       # SQL query files
│   │   ├── repo/          # Generated code from sqlc
│   │   └── schema/        # Database schema
│   ├── routes/            # HTTP route definitions
│   └── ui/                # UI components with gomponents
├── config.toml            # Configuration file
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
└── sqlc.yaml              # SQLC configuration
```

## Database Management

This template uses [sqlc](https://sqlc.dev/) to generate type-safe Go code from SQL:

```bash
# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Generate database code
sqlc generate
```

## Development Workflow

1. Define your database schema in `internal/database/schema/`
2. Write SQL queries in `internal/database/queries/`
3. Generate Go code with `sqlc generate`
4. Create API routes in `internal/routes/`
5. Create UI components in `internal/ui/`
6. Configure your application in `config.toml`

## Technologies Used

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- [Turso Database](https://turso.tech/)
- [SQLC](https://sqlc.dev/)
- [Gomponents](https://github.com/maragudk/gomponents)
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- [autocert](https://golang.org/x/crypto/acme/autocert)
