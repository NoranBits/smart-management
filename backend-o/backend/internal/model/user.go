// /////////////////////////////////////
// src: ./internal/model/user.go	//
// desc: Represents a user account.//
// //////////////////////////////////
package model

import "gorm.io/gorm"

// User represents a user account.
type User struct {
	gorm.Model        // Provides ID, CreatedAt, UpdatedAt, DeletedAt.
	Name       string `gorm:"type:varchar(100);not null"`                             // User's full name.
	Email      string `gorm:"type:varchar(100);uniqueIndex:uni_users_email;not null"` // User's email address.
	Password   string `gorm:"type:varchar(255);not null"`                             // Hashed password.
	// additional fields
}
