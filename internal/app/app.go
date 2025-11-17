package app

import (
	"context"
	"database/sql"
	"errors"
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

func (a *Application) Close() error {
	if a == nil {
		return nil
	}

	var allErrors []error

	if a.DBConn != nil {
		if err := a.DBConn.Close(); err != nil {
			allErrors = append(allErrors, fmt.Errorf("closing database: %w", err))
		}
		a.DBConn = nil
	}

	if len(allErrors) > 0 {
		if a.Logger == nil {
			a.Logger = slog.Default()
		}
		a.Logger.Warn("error closing application", "errors", len(allErrors))
		for i, err := range allErrors {
			a.Logger.Warn("Error", "id", i, "err", err.Error())
		}

		return errors.Join(allErrors...)
	}

	return nil
}
