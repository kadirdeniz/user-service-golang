package entity

import "gorm.io/gorm"

// Interest represents a specific interest category.
// swagger:model
type Interest struct {
	*gorm.Model

	// The title of the interest.
	// Required: true
	Title string

	// The category of the interest.
	// Required: true
	Category string

	// List of users associated with this interest.
	Users []User `gorm:"many2many:user_interests;"`
}
