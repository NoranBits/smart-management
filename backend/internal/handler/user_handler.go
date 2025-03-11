// internal/handler/user_handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "backend_server/internal/model"
	service "backend_server/service"

	"github.com/go-chi/chi/v5"
)

// UserRouter creates a new router for user endpoints.
func UserRouter(svc *service.UserService) http.Handler {
	r := chi.NewRouter()

	// Define the API routes.
	r.Get("/", listUsers(svc))
	r.Post("/", createUser(svc))
	r.Get("/{id}", getUser(svc))

	return r
}

// listUsers returns an HTTP handler that writes a JSON-encoded list of users to the response.
// @Summary List users
// @Description Retrieves a list of all users.
// @Tags Users
// @Produce json
// @Success 200 {array} dto.UserDTO
// @Router / [get]
func listUsers(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := svc.ListUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// getUser returns an HTTP handler that fetches a single user by their ID.
// @Summary Get user by ID
// @Description Fetch a user based on provided ID.
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {string} string "Invalid user id"
// @Failure 404 {string} string "User not found"
// @Router /{id} [get]
func getUser(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the user ID from the URL parameters.
		idParam := chi.URLParam(r, "id")
		// Convert the string to an int using Atoi, then cast to uint.
		idInt, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user id", http.StatusBadRequest)
			return
		}

		id := uint(idInt)

		user, err := svc.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// createUser returns an HTTP handler that creates a new user from a JSON payload.
// @Summary Create user
// @Description Creates a new user.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User to create"
// @Success 201 {object} dto.UserDTO
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal server error"
// @Router / [post]
func createUser(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		// Decode the incoming JSON payload into the User model.
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Invoke the service to create the new user.
		if err := svc.CreateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
