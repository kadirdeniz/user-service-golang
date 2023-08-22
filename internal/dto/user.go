package dto

import "gorm.io/gorm"

type GetUserByIdResponse struct {
	*gorm.Model
	FirstName string `mapstructure:"first_name"`
	LastName  string `mapstructure:"last_name"`
	Email     string `mapstructure:"email"`
	Nickname  string `mapstructure:"nickname"`
}
