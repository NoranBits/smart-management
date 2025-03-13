// /////////////////////////////////////////////////////////
// src: ./pkg/auth/secret.go							 //
// desc: Provides hashing and comparison functionality.	//
// //////////////////////////////////////////////////////
package auth

import (
	"errors"

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
