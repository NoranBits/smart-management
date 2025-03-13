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

	"errors"
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
	hashedPwd, err := auth.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	u.Password = string(hashedPwd)
	return s.Repo.CreateUser(u)
}

// Pseudocode for user login
func (s *UserService) LoginUser(email, password string) error {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	matched, err := auth.CheckPassword(user.Password, password)
	if err != nil {
		// Something went wrong comparing the hash
		return err
	}
	if !matched {
		// Password is incorrect
		return errors.New("invalid credentials")
	}

	// TODO: proceed with session / JWT token creation, etc.
	return nil
}

// ConvertUser maps an internal model.User to DTO.UserDTO.
func ConvertUser(u *model.User) DTO.UserDTO {
	return DTO.UserDTO{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
