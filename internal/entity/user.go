package entity

import "gorm.io/gorm"

// User represents the user entity in the system with associated details.
// swagger:model
type User struct {
	*gorm.Model

	// The first name of the user.
	Firstname string

	// The last name of the user.
	Lastname string

	// The email address of the user.
	// Format: email
	Email string

	// The nickname for the user.
	Nickname string

	// The password for the user (hashed).
	Password string

	// Profile picture associated with the user.
	ProfilePicture UserPicture

	// Interests associated with the user.
	Interests []Interest `gorm:"many2many:user_interests;"`

	// Posts created by the user.
	Posts []Post
}
