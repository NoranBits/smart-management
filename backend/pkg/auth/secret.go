// /////////////////////////////////////////////////////////
// src: ./pkg/auth/secret.go							 //
// desc: Provides hashing and comparison functionality.	//
// //////////////////////////////////////////////////////
package auth

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given plain-text password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a bcrypt-hashed password with a plain-text candidate.
func CheckPassword(hashedPassword, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// Password is incorrect (hash doesn't match)
			return false, nil
		}
		// An actual error occurred while comparing
		return false, err
	}
	// The passwords match
	return true, nil
}

// ValidatePasswordStrength checks if the password meets the minimum strength requirements.
func ValidatePasswordStrength(password string) bool {
	// Minimum 8 characters, 1 uppercase, 1 lowercase, 1 number, 1 special character
	if len(password) < 8 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := true // TODO: Check for special characters if role is admin/manager
	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case regexp.MustCompile(`[!@#$%^&*()]`).MatchString(string(char)):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
