package dto

import "user-service-golang/internal/entity"

// RegisterRequest represents the structure for an API registration request.
// swagger:model RegisterRequest
type RegisterRequest struct {
	// FirstName is the first name of the user.
	// Required: true
	// Minimum: 3
	// Maximum: 16
	// example: John
	FirstName string `json:"first_name" validate:"required,min=3,max=16"`

	// LastName is the last name of the user.
	// Required: true
	// Minimum: 3
	// Maximum: 16
	// example: Doe
	LastName string `json:"last_name" validate:"required,min=3,max=16"`

	// Email is the email address of the user.
	// Required: true
	// Format: email
	// example: john.doe@example.com
	Email string `json:"email" validate:"required,email"`

	// Nickname is the nickname for the user.
	// Required: true
	// Minimum: 3
	// Maximum: 16
	// example: johndoe
	Nickname string `json:"nickname" validate:"required,min=3,max=16"`

	// Password is the password for the user.
	// Required: true
	// Minimum: 8
	// Maximum: 16
	// example: password1234
	Password string `json:"password" validate:"required,min=8,max=16"`
}

// TokenResponse represents the structure for an API token response.
// swagger:response TokenResponse
type TokenResponse struct {
	// The JWT token for authenticated users.
	// Required: true
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
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
