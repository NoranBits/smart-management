// /////////////////////////////////////////////////////////////////////////////
// src: ./internal/handler/health.go										//
// desc: Provides health check endpoints to indicate the server is running.//
// //////////////////////////////////////////////////////////////////////////

package handler

import (
	"net/http"
)

// HealthCheck returns a simple "OK" message to indicate the server is running.
// @Summary Health check
// @Description Checks if the server is running
// @Tags health
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
