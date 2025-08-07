package routes

import (
	"github.com/go-chi/chi/v5"
	"starterA/internal/application"
)

// RegisterRoutes sets up all the routes for the application
func RegisterRoutes(r *chi.Mux, app *application.Application) {
	// No static files needed with Tailwind CSS via CDN

	// Register home route
	r.Get("/", app.HomeHandler())

	// Register API routes
	registerAPIRoutes(r, app)
}

// registerAPIRoutes sets up all API routes
func registerAPIRoutes(r *chi.Mux, app *application.Application) {
	r.Route("/api", func(r chi.Router) {
		// Users endpoints
		r.Route("/users", func(r chi.Router) {
			r.Get("/", app.GetUsersHandler())
			r.Post("/", app.CreateUserHandler())
			r.Get("/{id}", app.GetUserHandler())
			r.Put("/{id}", app.UpdateUserHandler())
			r.Delete("/{id}", app.DeleteUserHandler())
		})

		// Add more API routes as needed
	})
}
