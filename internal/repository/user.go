package repository

import (
	"context"
	"fmt"
	"user-service-golang/internal/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var userRepositoryInstance *userRepository

type UserRepositoryActions interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(userId uint) (entity.User, error)
	FindById(userId uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByNickname(email string) (entity.User, error)
	IsEmailExist(email string) bool
	IsNicknameExist(nickname string) bool
}

type userRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepositoryActions {
	if userRepositoryInstance == nil {
		userRepositoryInstance = &userRepository{
			db:    db,
			redis: redis,
		}
	}

	return userRepositoryInstance
}

func (u *userRepository) Create(user entity.User) (entity.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), user, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *userRepository) Update(user entity.User) (entity.User, error) {
	err := u.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), user, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *userRepository) Delete(userId uint) (entity.User, error) {
	user := entity.User{
		Model: &gorm.Model{
			ID: userId,
		},
	}
	err := u.db.Delete(&user).Error
	if err != nil {
		return user, err
	}

	err = u.redis.Del(context.Background(), fmt.Sprintf("users-%v", user.ID)).Err()
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *userRepository) FindById(userId uint) (entity.User, error) {
	var user entity.User

	err := u.redis.Get(context.Background(), fmt.Sprintf("users-%v", user.ID)).Scan(&user)
	if err != nil {
		fmt.Println(err)
	} else if user.ID != 0 {
		return user, err
	}

	err = u.db.First(&user, userId).Error
	if err != nil {
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), user, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := u.redis.Get(context.Background(), fmt.Sprintf("users-%v", user.ID)).Scan(&user)
	if err != nil {
		fmt.Println(err)
	} else if user.ID != 0 {
		return user, err
	}

	err = u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), user, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *userRepository) FindByNickname(nickname string) (entity.User, error) {
	var user entity.User

	err := u.redis.Get(context.Background(), fmt.Sprintf("users-%v", nickname)).Scan(&user)
	if err != nil {
		fmt.Println(err)
	} else if user.ID != 0 {
		return user, err
	}

	err = u.db.Where("nickname = ?", nickname).First(&user).Error
	if err != nil {
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), user, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *userRepository) IsEmailExist(email string) bool {
	var user entity.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return err == nil
}

func (u *userRepository) IsNicknameExist(nickname string) bool {
	var user entity.User
	err := u.db.Where("nickname = ?", nickname).First(&user).Error
	return err == nil
}
