package entity

import "gorm.io/gorm"

type Interest struct {
	*gorm.Model
	Title    string
	Category string
}
