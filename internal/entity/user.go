package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Firstname      string
	Lastname       string
	Email          string
	Nickname       string
	Password       string
	ProfilePicture UserPicture
	Interests      []Interest `gorm:"many2many:user_interests;"`
	Posts          []Post
}
