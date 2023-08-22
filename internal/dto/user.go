package dto

import (
	"user-service-golang/internal/entity"

	"gorm.io/gorm"
)

type GetUserByIdResponse struct {
	*gorm.Model
	Firstname string `mapstructure:"first_name"`
	Lastname  string `mapstructure:"last_name"`
	Email     string `mapstructure:"email"`
	Nickname  string `mapstructure:"nickname"`
}

type UpdateUserRequest struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
}

func (u UpdateUserRequest) ToUser() entity.User {
	return entity.User{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Nickname:  u.Nickname,
	}
}
