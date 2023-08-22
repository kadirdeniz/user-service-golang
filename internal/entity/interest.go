package entity

import "gorm.io/gorm"

// Interest represents a specific interest category.
// swagger:model
type Interest struct {
	*gorm.Model

	// The title of the interest.
	// Required: true
	Title string `mapstructure:"title"`

	// The category of the interest.
	// Required: true
	Category string `mapstructure:"category"`

	// List of users associated with this interest.
	Users []User `gorm:"many2many:user_interests;" mapstructure:"users"`
}
