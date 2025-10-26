package routes

import (
	"github.com/go-chi/chi/v5"
	httphandlers "starterA/internal/handlers/http"
)

// RegisterRoutes sets up all the routes for the application
func RegisterRoutes(r *chi.Mux, handlers *httphandlers.HTTPHandlers) {
	// No static files needed with Tailwind CSS via CDN

	// Register home route
	r.Get("/", handlers.HomeHandler())

	// Register API routes
	registerAPIRoutes(r, handlers)
}

// registerAPIRoutes sets up all API routes
func registerAPIRoutes(r *chi.Mux, handlers *httphandlers.HTTPHandlers) {
	r.Route("/api", func(r chi.Router) {
		// Users endpoints
		r.Route("/users", func(r chi.Router) {
			r.Get("/", handlers.GetUsersHandler())
			r.Post("/", handlers.CreateUserHandler())
			r.Get("/{id}", handlers.GetUserHandler())
			r.Put("/{id}", handlers.UpdateUserHandler())
			r.Delete("/{id}", handlers.DeleteUserHandler())
		})

		// Add more API routes as needed
	})
}
