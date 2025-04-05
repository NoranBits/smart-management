package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CONTEXT Key for storing user information in the request context.
type contextKey string

const (
	userIDKey   contextKey = "userID"
	userRoleKey contextKey = "userRole"
)

func validateToken(tokenString string) (jwt.MapClaims, bool) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET is not set")
		return nil, false
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
		return nil, false
	}
	if !token.Valid {
		return nil, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, true
	}
	return nil, false
}

// JWTAuth is a middleware that validates the JWT provided in the "Authorization" header.
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// DEBUG Check if authentication is disabled via environment variable.
		//		if os.Getenv("ENABLE_AUTH") == "false" {
		//		next.ServeHTTP(w, r)
		//		return
		//	}

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

		claims, valid := validateToken(tokenString)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Extract user information from claims and add to context
		userID := claims["sub"].(string)
		userRole := claims["role"].(string)

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		ctx = context.WithValue(ctx, userRoleKey, userRole)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext retrieves the user ID from the context.
func GetUserIDFromContext(ctx context.Context) string {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok {
		fmt.Println("userID not found in context")
		return "userID not found"
	}
	return userID
}

// GetUserRoleFromContext retrieves the user role from the context.
func GetUserRoleFromContext(ctx context.Context) string {
	userRole, ok := ctx.Value(userRoleKey).(string)
	if !ok {
		fmt.Println("userRole not found in context")
		return "userRole not found"
	}
	return userRole
}
