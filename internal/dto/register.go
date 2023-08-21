package dto

import (
	"user-service-golang/internal/entity"
)

// RegisterRequest represents the structure for an API registration request.
// swagger:model
type RegisterRequest struct {
	// The first name of the user.
	// Required: true
	// Mininum: 3
	// Maximum: 16
	FirstName string `json:"first_name" validate:"required,min=3,max=16"`

	// The last name of the user.
	// Required: true
	// Mininum: 3
	// Maximum: 16
	LastName string `json:"last_name" validate:"required,min=3,max=16"`

	// The email address of the user.
	// Required: true
	// Format: email
	Email string `json:"email" validate:"required,email"`

	// The nickname for the user.
	// Required: true
	// Mininum: 3
	// Maximum: 16
	Nickname string `json:"nickname" validate:"required,min=3,max=16"`

	// The password for the user.
	// Required: true
	// Mininum: 8
	// Maximum: 16
	Password string `json:"password" validate:"required,min=8,max=16"`
}

// TokenResponse represents the structure for an API token response.
// swagger:model
type TokenResponse struct {
	// The JWT token for authenticated users.
	// Required: true
	Token string `json:"token"`
}

// ToUser converts a RegisterRequest to a User entity.
func (r *RegisterRequest) ToUser() entity.User {
	return entity.User{
		Firstname: r.FirstName,
		Lastname:  r.LastName,
		Email:     r.Email,
		Nickname:  r.Nickname,
		Password:  r.Password,
	}
}
