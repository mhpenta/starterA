# Go Web Application Template

A simple but scalable web application template built with Go. This particular template is used for internal tools with maximum simplicity in mind.

It provides a solid foundation for developing web applications using a layered service architecture pattern (popularized by Alex Edwards), simple database integration, and hypermedia driven UI components.

## Features

- **Layered Architecture**: Clean separation between business logic (service layer) and HTTP handling (application layer)
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
gonew starterA github.com/yourusername/your-project
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
│   ├── application/       # Application layer (HTTP handlers)
│   │   ├── app.go         # Application struct and helpers
│   │   ├── users.go       # User-related HTTP handlers
│   │   └── home.go        # Home page handler
│   ├── service/           # Service layer (business logic)
│   │   └── service.go     # Service struct and business methods
│   ├── config/            # Configuration handling
│   ├── database/          # Database connection and queries
│   │   ├── queries/       # SQL query files
│   │   ├── repo/          # Generated code from sqlc
│   │   └── schema/        # Database schema
│   ├── routes/            # HTTP route definitions
│   └── ui/                # UI components with gomponents
├── ARCHITECTURE.md        # Detailed architecture documentation
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
4. Add business logic to the service layer in `internal/service/`
5. Create HTTP handlers in the application layer `internal/application/`
6. Register routes in `internal/routes/`
7. Create UI components in `internal/ui/`
8. Configure your application in `config.toml`

## Architecture

This template follows a layered service architecture pattern:

- **Service Layer** (`internal/service/`): Contains all business logic and dependencies
- **Application Layer** (`internal/application/`): Handles HTTP-specific concerns
- **Clean Separation**: Business logic is completely isolated from HTTP handling

For more details, see [ARCHITECTURE.md](ARCHITECTURE.md)

## Technologies Used

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- [Turso Database](https://turso.tech/)
- [SQLC](https://sqlc.dev/)
- [Gomponents](https://github.com/maragudk/gomponents)
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv)
- [autocert](https://golang.org/x/crypto/acme/autocert)
