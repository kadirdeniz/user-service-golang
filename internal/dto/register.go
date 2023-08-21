package dto

import (
	"user-service-golang/internal/entity"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=16"`
	LastName  string `json:"last_name" validate:"required,min=3,max=16"`
	Email     string `json:"email" validate:"required,email"`
	Nickname  string `json:"nickname" validate:"required,min=3,max=16"`
	Password  string `json:"password" validate:"required,min=8,max=16"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (r *RegisterRequest) ToUser() entity.User {
	return entity.User{
		Firstname: r.FirstName,
		Lastname:  r.LastName,
		Email:     r.Email,
		Nickname:  r.Nickname,
		Password:  r.Password,
	}
}
