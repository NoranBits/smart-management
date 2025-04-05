// ///////////////////////////////////////////////////////////////
// src: ./internal/repository/repository.go					   //
// desc: Provides database access methods for user management.//
// ////////////////////////////////////////////////////////////
package repository

import (
	model "backend_server/internal/model"
	auth "backend_server/pkg/auth"
	"fmt"

	"gorm.io/gorm"
)

// RepositoryInterface defines all data access methods your application needs.
type RepositoryInterface interface {
	// User-related methods
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

// repository is a concrete implementation of RepositoryInterface.
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new RepositoryInterface using the provided DB connection.
func NewRepository(db *gorm.DB) RepositoryInterface {
	return &repository{db: db}
}

// GetUserByID retrieves a single user by ID.
func (r *repository) GetUserByID(id uint) (*model.User, error) {
	var user *model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail retrieves a single user by email.
func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Unscoped().Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers returns a slice of all users.
func (r *repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser beszúr egy új felhasználót.
func (r *repository) CreateUser(user *model.User) error {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return fmt.Errorf("missing required fields")
	}
	hashedPwd, err := auth.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = hashedPwd

	fmt.Println("Creating user:", user) // Log the user data before insertion

	err = r.db.Create(user).Error
	if err != nil {
		fmt.Println("Error creating user:", err) // Log the error
		return err
	}

	fmt.Println("User created successfully:", user) // Log the user data after insertion

	return nil
}

// UpdateUser updates an existing user record.
func (r *repository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

// DeleteUser deletes a user record by ID.
func (r *repository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, "id = ?", id).Error
}
