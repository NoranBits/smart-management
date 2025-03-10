// ///////////////////////////////////////////////////////////////////////////
// src: ./internal/router/router.go										  //
// desc: Initializes and configures the HTTP routing for the application.//
// ////////////////////////////////////////////////////////////////////////
package router

import (
	"net/http"
	"time"

	handler "backend_server/internal/handler"
	repository "backend_server/internal/repository"
	service "backend_server/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// RouterDependencies groups dependencies required by the router.
type RouterDependencies struct {
	Repo repository.RepositoryInterface
}

// NewRouter initializes a new Chi router with routes and middleware.
func NewRouter(repo repository.RepositoryInterface) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Health check endpoint.
	r.Get("/health", handler.HealthCheck)

	// Mount user API routes.
	userService := service.NewUserService(repo)
	r.Mount("/api/users", handler.UserRouter(userService))

	return r
}
