package entity

import "gorm.io/gorm"

// UserPicture represents the different sizes of a user's picture.
// swagger:model
type UserPicture struct {
	*gorm.Model

	// The ID of the user to whom the picture belongs.
	// Required: true
	UserId uint `mapstructure:"user_id"`

	// The thumbnail version of the user's picture.
	Thumbnail string `mapstructure:"thumbnail"`

	// The original size of the user's picture.
	Original string `mapstructure:"original"`

	// The large version of the user's picture.
	Large string `mapstructure:"large"`
}
