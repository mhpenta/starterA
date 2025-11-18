package httphandlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/mhpenta/starterA/internal/service"
)

// HTTPHandlers handles HTTP-specific request/response logic
type HTTPHandlers struct {
	Service *service.Service
	Logger  *slog.Logger
}

// New creates a new HTTPHandlers instance
func New(svc *service.Service, logger *slog.Logger) *HTTPHandlers {
	return &HTTPHandlers{
		Service: svc,
		Logger:  logger,
	}
}

// respond sends a JSON response with the given status code and data
func (h *HTTPHandlers) respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			h.Logger.Error("Failed to encode response", "error", err)
		}
	}
}

// respondError sends a JSON error response
func (h *HTTPHandlers) respondError(w http.ResponseWriter, status int, message string) {
	h.respond(w, status, map[string]string{"error": message})
}

// badRequest logs and returns a 400 Bad Request response
func (h *HTTPHandlers) badRequest(w http.ResponseWriter, err error) {
	h.Logger.Warn("Bad request", "error", err)
	h.respondError(w, http.StatusBadRequest, "Invalid request")
}

// notFound returns a 404 Not Found response
func (h *HTTPHandlers) notFound(w http.ResponseWriter) {
	h.respondError(w, http.StatusNotFound, "Resource not found")
}

// serverError logs and returns a 500 Internal Server Error response
func (h *HTTPHandlers) serverError(w http.ResponseWriter, err error) {
	h.Logger.Error("Server error", "error", err)
	h.respondError(w, http.StatusInternalServerError, "Internal server error")
}
