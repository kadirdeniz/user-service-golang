package dto

import "user-service-golang/internal/entity"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

func (l *LoginRequest) ToUser() entity.User {
	return entity.User{
		Email:    l.Email,
		Password: l.Password,
	}
}
