// /////////////////////////////////////
// src: ./internal/model/user.go	//
// desc: Represents a user account.//
// //////////////////////////////////
package model

import "gorm.io/gorm"

// User represents a user account.
type User struct {
	gorm.Model        // Provides ID, CreatedAt, UpdatedAt, DeletedAt.
	Name       string `gorm:"type:varchar(100);not null"`             // User's full name.
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null"` // User's email address.
	Password   string `gorm:"type:varchar(255);not null"`             // Hashed password.
	Role       string `gorm:"type:varchar(50);not null"`              // User's role.
	Active     bool   `gorm:"default:true;not null"`                  // Account status.
	// additional fields
}
