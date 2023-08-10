package entity

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	UserId  uint
	Title   string
	Picture string
}
