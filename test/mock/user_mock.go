package mock

import (
	"time"
	"user-service-golang/internal/entity"

	"gorm.io/gorm"
)

var MockUser entity.User = entity.User{
	Model: &gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	},
	Firstname: "John",
	Lastname:  "Doe",
	Email:     "johndoe@mail.com",
	Nickname:  "johndoe",
	Password:  "$2a$10$X3vy//lV9cMPGH/mN0JdweI3TXRkUpp4viyQ2RCoDOxGBt8Jvz6bS",
	ProfilePicture: entity.UserPicture{
		Model: &gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		UserId:    1,
		Thumbnail: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		Original:  "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		Large:     "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
	},
	Interests: []entity.Interest{
		{
			Model: &gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Title: "Basketball",
			Users: nil,
		},
		{
			Model: &gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Title: "Football",
			Users: nil,
		},
	},
	Posts: []entity.Post{
		{
			Model: &gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			UserId:  1,
			Title:   "Post 1",
			Picture: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		},
		{
			Model: &gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			UserId:  1,
			Title:   "Post 2",
			Picture: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		},
	},
}
