package dto

import "user-service-golang/internal/entity"

// LoginRequest represents the structure for an API login request.
// swagger:model
type LoginRequest struct {
	// The email address of the user.
	// Required: true
	// Format: email
	Email string `json:"email" validate:"required,email"`

	// The password for the user.
	// Required: true
	// Mininum: 8
	// Maximum: 16
	Password string `json:"password" validate:"required,min=8,max=16"`
}

// ToUser converts a LoginRequest to a User entity.
func (l *LoginRequest) ToUser() entity.User {
	return entity.User{
		Email:    l.Email,
		Password: l.Password,
	}
}
