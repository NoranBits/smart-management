package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func validateToken(tokenString string) bool {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET is not set")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method is HMAC.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Printf("Token parsing error: %v\n", err)
		return false
	}

	return token.Valid
}

// JWTAuth is a middleware that validates the JWT provided in the "Authorization" header.
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// DEBUG Check if authentication is disabled via environment variable.
		if os.Getenv("ENABLE_AUTH") == "false" {
			// UNSECURE Skip token validation and directly call the next handler.
			next.ServeHTTP(w, r)
			return
		}

		// Extract token from the Authorization header.
		// It should be sent as: "Bearer <token>"
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// OPTIONAL split the auth header if it contains 'Bearer ' prefix.
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		if !validateToken(tokenString) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If token is valid, pass the request to the next handler.
		next.ServeHTTP(w, r)
	})
}
