package entity

import "gorm.io/gorm"

type Interest struct {
	*gorm.Model
	Title    string
	Category string
	Users    []User `gorm:"many2many:user_interests;"`
}
