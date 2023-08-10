package entity

import "gorm.io/gorm"

type Picture struct {
	*gorm.Model
	UserId    uint
	Thumbnail string
	Original  string
	Large     string
}
