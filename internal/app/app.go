package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"starterA/internal/config"
	"starterA/internal/database"
	"starterA/internal/database/repo"
)

// Application is a dependency container for the application.
// It holds shared resources that can be used across different
// interfaces (HTTP, CLI, TUI, MCP, etc.)
type Application struct {
	AppCtx context.Context
	Logger *slog.Logger
	Config *config.Config
	DB     *repo.Queries
	DBConn *sql.DB
}

// New creates a new Application instance with the provided dependencies
func New(ctx context.Context, logger *slog.Logger, cfg *config.Config) (*Application, error) {

	dbConn, err := database.GetConnection(cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("error getting db connection: %w", err)
	}

	db := repo.New(dbConn)

	return &Application{
		AppCtx: ctx,
		Logger: logger,
		Config: cfg,
		DB:     db,
		DBConn: dbConn,
	}, nil
}
