// /////////////////////////////////////////////////////////
// src: ./internal/dto/user_dto.go 					     //
// desc: DTO package provides data transfer objects.	//
// //////////////////////////////////////////////////////
package dto

// UserDTO represents the public view of a user.
type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
