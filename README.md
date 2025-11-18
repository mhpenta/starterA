# Go Application Starter

Transport-agnostic Go template with HTTP server support. Core business logic stays separate from transport layer - easily add CLI, TUI, or MCP interfaces later.

## Stack

- **Chi** - HTTP router with middleware
- **Turso** - SQLite database with libsql, easily swapped out for sqlite
- **SQLC** - Type-safe SQL code generation
- **Gomponents** - Type-safe HTML templating
**- **Let's Encrypt** - Automatic HTTPS (optional)**

## Quick Start

```bash
# Run
go run ./cmd/main.go

# Build
go build -o app ./cmd/main.go

# Custom config
go run ./cmd/main.go -c path/to/config.toml
```

## Configuration

Edit `cmd/config.toml`:

```toml
[Database]
TursoConnectionString = "libsql://your-database.turso.io?authToken=your-token"

[Server]
Port = "8080"
EnableHTTPS = false
HTTPSPort = "443"
TimeOutInSeconds = 60
AllowedCorsURLs = ["http://localhost:3000"]
TaskTimeOutInSeconds = 3600
ServerDomain = "yourdomain.com"

[App]
Environment = "dev"
```

## Architecture

```
app (infrastructure) → service (business logic) → handlers (transport)
```

- `internal/app/` - Dependency container (DB, Logger, Config, Context)
- `internal/service/` - Transport-agnostic business logic
- `internal/handlers/` - Transport adapters (HTTP, CLI, TUI, etc.)
- `internal/database/` - Database access with SQLC-generated code

## Development

1. Define schema: `internal/database/schema/`
2. Write queries: `internal/database/queries/`
3. Generate code: `sqlc generate`
4. Add business logic: `internal/service/`
5. Add handlers: `internal/handlers/`
6. Register routes: `internal/routes/`

## License

MIT License - see [LICENSE](LICENSE) file for details.
