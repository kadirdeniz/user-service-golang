package entity

import "gorm.io/gorm"

// User represents the user entity in the system with associated details.
// swagger:model
type User struct {
	*gorm.Model

	// The first name of the user.
	Firstname string `mapstructure:"first_name"`

	// The last name of the user.
	Lastname string `mapstructure:"last_name"`

	// The email address of the user.
	// Format: email
	Email string `mapstructure:"email"`

	// The nickname for the user.
	Nickname string `mapstructure:"nickname"`

	// The password for the user (hashed).
	Password string `mapstructure:"password"`

	// Profile picture associated with the user.
	ProfilePicture UserPicture `mapstructure:"profile_picture"`

	// Interests associated with the user.
	Interests []Interest `gorm:"many2many:user_interests;" mapstructure:"interests"`

	// Posts created by the user.
	Posts []Post
}
