package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"starterA/internal/database/repo"
	"starterA/internal/handlers"
)

// RegisterRoutes sets up all the routes for the application
func RegisterRoutes(r *chi.Mux, db *repo.Queries) {
	// No static files needed with Tailwind CSS via CDN

	// Register home route
	r.Get("/", handlers.HomeHandler())

	// Register API routes
	registerAPIRoutes(r, db)
}

// registerAPIRoutes sets up all API routes
func registerAPIRoutes(r *chi.Mux, db *repo.Queries) {
	r.Route("/api", func(r chi.Router) {
		// Users endpoints
		r.Route("/users", func(r chi.Router) {
			r.Get("/", handleGetUsers(db))
			r.Post("/", handleCreateUser(db))
			r.Get("/{id}", handleGetUser(db))
			r.Put("/{id}", handleUpdateUser(db))
			r.Delete("/{id}", handleDeleteUser(db))
		})

		// Add more API routes as needed
	})
}

// User handlers
func handleGetUsers(db *repo.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation will go here
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func handleCreateUser(db *repo.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation will go here
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func handleGetUser(db *repo.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation will go here
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func handleUpdateUser(db *repo.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation will go here
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func handleDeleteUser(db *repo.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation will go here
		w.WriteHeader(http.StatusNotImplemented)
	}
}
