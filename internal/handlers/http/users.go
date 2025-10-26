package httphandlers

import (
	"encoding/json"
	"net/http"
	"starterA/internal/service"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// GetUsersHandler returns an HTTP handler for listing users
func (h *HTTPHandlers) GetUsersHandler() http.HandlerFunc {
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

		users, err := h.Service.GetUsers(r.Context(), limit, offset)
		if err != nil {
			h.serverError(w, err)
			return
		}

		h.respond(w, http.StatusOK, users)
	}
}

// CreateUserHandler returns an HTTP handler for creating a user
func (h *HTTPHandlers) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input service.CreateUserInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			h.badRequest(w, err)
			return
		}

		user, err := h.Service.CreateUser(r.Context(), &input)
		if err != nil {
			h.serverError(w, err)
			return
		}

		h.respond(w, http.StatusCreated, user)
	}
}

// GetUserHandler returns an HTTP handler for getting a single user
func (h *HTTPHandlers) GetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			h.badRequest(w, err)
			return
		}

		user, err := h.Service.GetUser(r.Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				h.notFound(w)
				return
			}
			h.serverError(w, err)
			return
		}

		h.respond(w, http.StatusOK, user)
	}
}

// UpdateUserHandler returns an HTTP handler for updating a user
func (h *HTTPHandlers) UpdateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			h.badRequest(w, err)
			return
		}

		var input service.UpdateUserInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			h.badRequest(w, err)
			return
		}

		user, err := h.Service.UpdateUser(r.Context(), id, &input)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				h.notFound(w)
				return
			}
			h.serverError(w, err)
			return
		}

		h.respond(w, http.StatusOK, user)
	}
}

// DeleteUserHandler returns an HTTP handler for deleting a user
func (h *HTTPHandlers) DeleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			h.badRequest(w, err)
			return
		}

		err = h.Service.DeleteUser(r.Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				h.notFound(w)
				return
			}
			h.serverError(w, err)
			return
		}

		h.respond(w, http.StatusNoContent, nil)
	}
}