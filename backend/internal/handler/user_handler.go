// /////////////////////////////////////////////////////////////////////////////
// src: ./ internal/handler/user_handler.go									 //
// desc: Provides HTTP handlers for user management endpoints.				//
// //////////////////////////////////////////////////////////////////////////
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "backend_server/DTO"
	model "backend_server/internal/model"
	service "backend_server/service"

	chi "github.com/go-chi/chi/v5"
)

// UserRouter creates a new router for user endpoints.
func UserRouter(svc *service.UserService) http.Handler {
	r := chi.NewRouter()

	// Define the API routes.
	r.Get("/", listUsers(svc))
	r.Get("/{id}", getUser(svc))
	r.Post("/", createUser(svc))
	r.Delete("/{id}", deleteUser(svc))
	r.Put("/{id}", updateUser(svc))
	r.Get("/email/{email}", getUserByEmail(svc))

	return r
}

// listUsers returns an HTTP handler that writes a JSON-encoded list of users to the response.
func listUsers(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userList, err := svc.ListUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Convert model.User slice to []dto.UserDTO
		dtoList := make([]dto.UserDTO, 0, len(userList))
		for _, u := range userList {
			dtoList = append(dtoList, service.ConvertUser(&u))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dtoList)
	}
}

// getUser returns an HTTP handler that fetches a single user by their ID.
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
		// Convert the returned user to a DTO.
		dtoUser := service.ConvertUser(user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dtoUser)
	}
}

// createUser returns an HTTP handler that creates a new user from a JSON payload.
func createUser(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
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

		// Convert to DTO so that we don't expose the hashed password.
		dtoUser := service.ConvertUser(&user)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(dtoUser)
	}
}

// getUserByEmail returns an HTTP handler that fetches a single user by their Email.
func getUserByEmail(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the user Email from the URL parameters.
		emailParam := chi.URLParam(r, "email")

		user, err := svc.GetUserByEmail(emailParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		dtoUser := service.ConvertUser(user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dtoUser)
	}
}

// updateUser returns an HTTP handler that updates a user by their ID.
func deleteUser(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user id", http.StatusBadRequest)
			return
		}

		if err := svc.DeleteUser(uint(idInt)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent) // 204 No Content response status tells the client that the request has succeeded
	}
}

// updateUser returns an HTTP handler that updates a user by their ID.
func updateUser(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid user id", http.StatusBadRequest)
			return
		}

		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		user.ID = uint(idInt)
		if err := svc.UpdateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent) // 204 No Content response status tells the client that the request has succeeded
	}
}
