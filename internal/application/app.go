package application

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"starterA/internal/config"
	"starterA/internal/service"
)

type Application struct {
	Service *service.Service
	Logger  *slog.Logger
	Config  *config.Config
}

func New(service *service.Service, logger *slog.Logger, cfg *config.Config) *Application {
	return &Application{
		Service: service,
		Logger:  logger,
		Config:  cfg,
	}
}

func (app *Application) respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			app.Logger.Error("Failed to encode response", "error", err)
		}
	}
}

func (app *Application) respondError(w http.ResponseWriter, status int, message string) {
	app.respond(w, status, map[string]string{"error": message})
}

func (app *Application) badRequest(w http.ResponseWriter, err error) {
	app.Logger.Warn("Bad request", "error", err)
	app.respondError(w, http.StatusBadRequest, "Invalid request")
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.respondError(w, http.StatusNotFound, "Resource not found")
}

func (app *Application) serverError(w http.ResponseWriter, err error) {
	app.Logger.Error("Server error", "error", err)
	app.respondError(w, http.StatusInternalServerError, "Internal server error")
}