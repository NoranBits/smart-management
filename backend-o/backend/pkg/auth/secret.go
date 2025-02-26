// /////////////////////////////////////////////////////////
// src: ./pkg/auth/secret.go							 //
// desc: Provides hashing and comparison functionality.	//
// //////////////////////////////////////////////////////
package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given plain-text password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with its possible plain-text equivalent.
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
