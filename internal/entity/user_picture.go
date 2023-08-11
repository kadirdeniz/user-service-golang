package entity

import "gorm.io/gorm"

type UserPicture struct {
	*gorm.Model
	UserId    uint
	Thumbnail string
	Original  string
	Large     string
}
