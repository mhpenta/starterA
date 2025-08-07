package application

import (
	"encoding/json"
	"net/http"
	"starterA/internal/service"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (app *Application) GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Default pagination
		limit := int64(100)
		offset := int64(0)
		
		// Parse query parameters
		if l := r.URL.Query().Get("limit"); l != "" {
			if parsedLimit, err := strconv.ParseInt(l, 10, 64); err == nil && parsedLimit > 0 {
				limit = parsedLimit
			}
		}
		if o := r.URL.Query().Get("offset"); o != "" {
			if parsedOffset, err := strconv.ParseInt(o, 10, 64); err == nil && parsedOffset >= 0 {
				offset = parsedOffset
			}
		}
		
		users, err := app.Service.GetUsers(r.Context(), limit, offset)
		if err != nil {
			app.serverError(w, err)
			return
		}

		app.respond(w, http.StatusOK, users)
	}
}

func (app *Application) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input service.CreateUserInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			app.badRequest(w, err)
			return
		}

		user, err := app.Service.CreateUser(r.Context(), &input)
		if err != nil {
			app.serverError(w, err)
			return
		}

		app.respond(w, http.StatusCreated, user)
	}
}

func (app *Application) GetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			app.badRequest(w, err)
			return
		}

		user, err := app.Service.GetUser(r.Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				app.notFound(w)
				return
			}
			app.serverError(w, err)
			return
		}

		app.respond(w, http.StatusOK, user)
	}
}

func (app *Application) UpdateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			app.badRequest(w, err)
			return
		}

		var input service.UpdateUserInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			app.badRequest(w, err)
			return
		}

		user, err := app.Service.UpdateUser(r.Context(), id, &input)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				app.notFound(w)
				return
			}
			app.serverError(w, err)
			return
		}

		app.respond(w, http.StatusOK, user)
	}
}

func (app *Application) DeleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			app.badRequest(w, err)
			return
		}

		err = app.Service.DeleteUser(r.Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				app.notFound(w)
				return
			}
			app.serverError(w, err)
			return
		}

		app.respond(w, http.StatusNoContent, nil)
	}
}