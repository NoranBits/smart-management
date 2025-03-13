// /////////////////////////////////////////////////////////////////////////////
// src: ./internal/handler/health.go										//
// desc: Provides health check endpoints to indicate the server is running.//
// //////////////////////////////////////////////////////////////////////////

package handler

import (
	"net/http"
)

// HealthCheck returns a simple "OK" message to indicate the server is running.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
