package entity

import "gorm.io/gorm"

// Post represents a user's post in the system.
// swagger:model
type Post struct {
	*gorm.Model

	// The ID of the user who created the post.
	// Required: true
	UserId uint

	// The title of the post.
	// Required: true
	Title string

	// The URL or location of the picture associated with the post.
	Picture string
}
