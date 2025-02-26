// //////////////////////////////////////////////////////////////////
// src: internal/middleware/logging.go					 		  //
// desc: Provides a middleware that logs incoming HTTP requests. //
// ///////////////////////////////////////////////////////////////
package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware" // For retrieving request ID.
	"github.com/rs/zerolog/log"
)

// Logger is a middleware that logs incoming HTTP requests with extra details.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Process the request.
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		// Retrieve the request ID set by chi's RequestID middleware.
		reqID := middleware.GetReqID(r.Context())
		// Use r.RemoteAddr for the client IP (or use a custom RealIP middleware).
		remoteIP := r.RemoteAddr

		// Log the request details in a structured format.
		log.Info().
			Str("req_id", reqID).
			Str("remote_ip", remoteIP).
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Dur("duration", duration).
			Msg("handled request")
	})
}
