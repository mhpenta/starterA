package app

import (
	"context"
	"log/slog"
	"starterA/internal/config"
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
}

// New creates a new Application instance with the provided dependencies
func New(ctx context.Context, logger *slog.Logger, cfg *config.Config, db *repo.Queries) *Application {
	return &Application{
		AppCtx: ctx,
		Logger: logger,
		Config: cfg,
		DB:     db,
	}
}
