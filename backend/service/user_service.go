// /////////////////////////////////////////////////////////////////
// src: ./internal/service/user_service.go						 //
// desc: Provides business logic operations for user entities.	//
// //////////////////////////////////////////////////////////////
package service

import (
	DTO "backend_server/DTO"
	model "backend_server/internal/model"
	repository "backend_server/internal/repository"
	auth "backend_server/pkg/auth"
	"fmt"
	"os"
	"time"

	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

// UserService defines operations for users.
type UserService struct {
	Repo repository.RepositoryInterface
}

// NewUserService creates a new UserService with the provided repository.
func NewUserService(repo repository.RepositoryInterface) *UserService {
	return &UserService{Repo: repo}
}

// GetUserByID retrieves a user by their ID.
// Returns an error if the user is not found.
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// ListUsers retrieves all users.
func (s *UserService) ListUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}

// CreateUser handles the operation to create a new user.
func (s *UserService) CreateUser(u *model.User) error {
	// Hash the password before creating the user
	hashedPwd, err := auth.HashPassword(u.Password)
	if err != nil {
		return errors.New("error hashing password: %v")
	}
	u.Password = string(hashedPwd)
	return s.Repo.CreateUser(u)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) UpdateUser(u *model.User) error {
	// Hash the password before updating the user
	hashedPwd, err := auth.HashPassword(u.Password)
	if err != nil {
		return errors.New("error hashing password: %v")
	}
	u.Password = string(hashedPwd)
	return s.Repo.UpdateUser(u)
}

// LoginUser handles the operation to authenticate a user.
func (s *UserService) LoginUser(email, password string) error {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		fmt.Println("Error getting user by email:", err)
		return errors.New("invalid credentials") // More generic error message
	}

	fmt.Println("User found:", user) // Log the user

	matched, err := auth.CheckPassword(user.Password, password)
	if err != nil {
		fmt.Println("Error comparing passwords:", err) // Log the error
		return errors.New("invalid credentials")
	}
	if !matched {
		fmt.Println("Passwords do not match") // Log the mismatch
		return errors.New("invalid credentials")
	}

	// Authentication successful
	return nil
}

// GenerateJWT generates a JWT token for the user.
func (s *UserService) GenerateJWT(email string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

// ConvertUser maps an internal model.User to DTO.UserDTO.
func ConvertUser(u *model.User) DTO.UserDTO {
	return DTO.UserDTO{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Role:   u.Role,
		Active: u.Active,
	}
}
