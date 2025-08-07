# Layered Service Architecture

This application follows a two-layer architecture pattern popularized by Alex Edwards, providing clean separation of concerns and improved testability.

## Architecture Overview

The application is organized into two main layers:

### 1. Service Layer (`internal/service/`)
- Contains all business logic
- Manages dependencies (database, logger, config)
- Provides a clean API for business operations
- Completely decoupled from HTTP concerns

**Key Components:**
- `Service` struct: Central dependency container
- Business methods: `CreateUser()`, `GetUsers()`, `UpdateUser()`, `DeleteUser()`
- Input/Output DTOs: Clean data transfer objects

### 2. Application Layer (`internal/application/`)
- Handles HTTP-specific concerns
- Request/response handling
- Input validation and error formatting
- Routes requests to appropriate service methods

**Key Components:**
- `Application` struct: HTTP handler container
- HTTP handlers: `GetUsersHandler()`, `CreateUserHandler()`, etc.
- Helper methods: `respond()`, `respondError()`, `badRequest()`, `serverError()`

## Benefits

1. **Separation of Concerns**: Business logic is completely separated from HTTP handling
2. **Testability**: Service layer can be tested without HTTP dependencies
3. **Maintainability**: Adding new dependencies only requires updating the Service struct
4. **Reusability**: Service layer can be used by different interfaces (HTTP, CLI, gRPC)
5. **Clean Dependencies**: Each layer has clear, minimal dependencies

## Directory Structure

```
internal/
├── application/        # Application layer (HTTP handlers)
│   ├── app.go         # Application struct and helpers
│   ├── users.go       # User-related HTTP handlers
│   └── home.go        # Home page handler
├── service/           # Service layer (business logic)
│   └── service.go     # Service struct and business methods
├── database/          # Database layer
│   ├── repo/          # Generated SQLC code
│   └── connect.go     # Database connection
├── routes/            # Route registration
│   └── routes.go      # HTTP route definitions
└── config/            # Configuration
    └── config.go      # Config loading and structs
```

## Request Flow

1. HTTP request arrives at router (`routes/routes.go`)
2. Router directs to appropriate application handler (`application/*.go`)
3. Application handler validates input and calls service method (`service/service.go`)
4. Service executes business logic and database operations (`database/repo/*.go`)
5. Service returns result to application layer
6. Application formats response and sends to client

## Example: Creating a User

```go
// 1. HTTP Handler (Application Layer)
func (app *Application) CreateUserHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var input service.CreateUserInput
        json.NewDecoder(r.Body).Decode(&input)
        
        user, err := app.Service.CreateUser(r.Context(), &input)
        if err != nil {
            app.serverError(w, err)
            return
        }
        
        app.respond(w, http.StatusCreated, user)
    }
}

// 2. Business Logic (Service Layer)
func (s *Service) CreateUser(ctx context.Context, input *CreateUserInput) (*repo.User, error) {
    s.Logger.Info("Creating user", "username", input.Username)
    
    user, err := s.DB.CreateUser(ctx, repo.CreateUserParams{
        Username: input.Username,
        Email:    input.Email,
    })
    
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return &user, nil
}
```

## Adding New Features

To add a new feature:

1. **Define business logic** in the service layer
2. **Create HTTP handlers** in the application layer
3. **Register routes** in `routes/routes.go`
4. **Add any new dependencies** to the Service struct

## Testing

The layered architecture makes testing straightforward:

- **Service Layer**: Test business logic with mocked dependencies
- **Application Layer**: Test HTTP handling with mocked service
- **Integration**: Test full flow with real database

## References

- [Alex Edwards - The Fat Service Pattern](https://www.alexedwards.net/blog/the-fat-service-pattern)
- [Organizing Database Access in Go](https://www.alexedwards.net/blog/organising-database-access)