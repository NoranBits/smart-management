////////////////////////////////////////////////////////////////////////////////
// src: ./ internal/handler/auth_handler.go									 //
// desc: Provides HTTP handlers for authentication endpoints.				//
// //////////////////////////////////////////////////////////////////////////

package handler

import (
	"encoding/json"
	"net/http"

	"backend_server/internal/model"
	"backend_server/service"
)

// LoginRequest defines the structure for the login request payload.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest defines the structure for the register request payload.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// LoginHandler handles user login.
func LoginHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := svc.LoginUser(req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Generate JWT token
		token, err := svc.GenerateJWT(req.Email)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

// RegisterHandler handles user registration.
func RegisterHandler(svc *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		user := &model.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Role:     req.Role,
			Active:   true, // Set default value
		}

		err := svc.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User registered successfully"))
	}
}
