package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"user-service-golang/internal/entity"
	"user-service-golang/pkg"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var userRepositoryInstance *userRepository
var userRedisRepositoryInstance *userRedisRepository
var userDBRepositoryInstance *userDBRepository

type UserRepositoryActions interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) error
	Delete(userId uint) error
	FindById(userId uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByNickname(email string) (entity.User, error)
	IsEmailExist(email string) bool
	IsNicknameExist(nickname string) bool
}

type userRepository struct {
	db    UserRepositoryActions
	redis UserRepositoryActions
}

func NewUserRepository(db UserRepositoryActions, redis UserRepositoryActions) UserRepositoryActions {
	if userRepositoryInstance == nil {
		userRepositoryInstance = &userRepository{
			db:    db,
			redis: redis,
		}
	}

	return userRepositoryInstance
}

func (u *userRepository) Create(user entity.User) (entity.User, error) {
	user, err := u.db.Create(user)
	if err != nil {
		return user, err
	}

	user, err = u.redis.Create(user)
	if err != nil {
		slog.Error(err.Error())
	}

	return user, err
}

func (u *userRepository) Update(user entity.User) error {
	err := u.redis.Update(user)
	if err != nil && err != pkg.ErrUserNotFound {
		slog.Error(err.Error())
	}

	err = u.db.Update(user)
	if err != nil {
		return err
	}

	return err
}

func (u *userRepository) Delete(userId uint) error {
	err := u.redis.Delete(userId)
	if err != nil && err != pkg.ErrUserNotFound {
		slog.Error(err.Error())
	}

	err = u.db.Delete(userId)
	if err != nil {
		return err
	}

	return err
}

func (u *userRepository) FindById(userId uint) (entity.User, error) {
	user, err := u.redis.FindById(userId)
	if err != nil && err != pkg.ErrUserNotFound {
		slog.Error(err.Error())
	}

	if !user.IsEmpty() {
		return user, err
	}

	user, err = u.db.FindById(userId)
	if err != nil {
		return user, err
	}

	user, err = u.redis.Create(user)
	if err != nil {
		slog.Error(err.Error())
	}

	return user, err
}

func (u *userRepository) FindByEmail(email string) (entity.User, error) {
	user, err := u.redis.FindByEmail(email)
	if err != nil && err != pkg.ErrUserNotFound {
		slog.Error(err.Error())
	}

	if !user.IsEmpty() {
		return user, err
	}

	user, err = u.db.FindByEmail(email)
	if err != nil {
		return user, err
	}

	user, err = u.redis.Create(user)
	if err != nil {
		slog.Error(err.Error())
	}

	return user, err
}

func (u *userRepository) FindByNickname(nickname string) (entity.User, error) {
	user, err := u.redis.FindByNickname(nickname)
	if err != nil && err != pkg.ErrUserNotFound {
		slog.Error(err.Error())
	}

	if !user.IsEmpty() {
		return user, err
	}

	user, err = u.db.FindByNickname(nickname)
	if err != nil {
		return user, err
	}

	user, err = u.redis.Create(user)
	if err != nil {
		slog.Error(err.Error())
	}

	return user, err
}

func (u *userRepository) IsEmailExist(email string) bool {
	if u.redis.IsEmailExist(email) {
		return true
	}

	return u.db.IsEmailExist(email)
}

func (u *userRepository) IsNicknameExist(nickname string) bool {
	if u.redis.IsNicknameExist(nickname) {
		return true
	}

	return u.db.IsNicknameExist(nickname)
}

type userRedisRepository struct {
	redis *redis.Client
}

func NewUserRedisRepository(redis *redis.Client) UserRepositoryActions {
	if userRedisRepositoryInstance == nil {
		userRedisRepositoryInstance = &userRedisRepository{
			redis: redis,
		}
	}

	return userRedisRepositoryInstance
}

func (u *userRedisRepository) Create(user entity.User) (entity.User, error) {
	userStringify, err := json.Marshal(user)
	if err != nil {
		slog.Error(err.Error())
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), userStringify, 0).Err()
	if err != nil {
		slog.Error(err.Error())
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-email-%v", user.Email), user.ID, 0).Err()
	if err != nil {
		slog.Error(err.Error())
		return user, err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-nickname-%v", user.Nickname), user.ID, 0).Err()
	if err != nil {
		slog.Error(err.Error())
		return user, err
	}

	return user, err
}

func (u *userRedisRepository) Update(user entity.User) error {
	userStringify, err := json.Marshal(user)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-%v", user.ID), userStringify, 0).Err()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-email-%v", user.Email), user.ID, 0).Err()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = u.redis.Set(context.Background(), fmt.Sprintf("users-nickname-%v", user.Nickname), user.ID, 0).Err()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return err
}

func (u *userRedisRepository) Delete(userId uint) error {
	userObj, err := u.FindById(userId)
	if err != nil {
		return err
	}

	err = u.redis.Del(context.Background(), fmt.Sprintf("users-%v", userId)).Err()
	if err != nil {
		return err
	}

	err = u.redis.Del(context.Background(), fmt.Sprintf("users-email-%v", userObj.Email)).Err()
	if err != nil {
		return err
	}

	err = u.redis.Del(context.Background(), fmt.Sprintf("users-nickname-%v", userObj.Nickname)).Err()
	if err != nil {
		return err
	}

	return err
}

func (u *userRedisRepository) FindById(userId uint) (entity.User, error) {
	var user entity.User

	userObj := u.redis.Get(context.Background(), fmt.Sprintf("users-%v", userId)).Val()
	if userObj == "" {
		return user, pkg.ErrUserNotFound
	}

	err := json.Unmarshal([]byte(userObj), &user)
	if err != nil {
		slog.Error(err.Error())
		return user, pkg.ErrUserNotFound
	}

	return user, err
}

func (u *userRedisRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	userId := u.redis.Get(context.Background(), fmt.Sprintf("users-email-%v", email)).Val()
	if userId == "" {
		return user, pkg.ErrUserNotFound
	}

	uintUserId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		slog.Error(err.Error())
		return user, pkg.ErrUserNotFound
	}

	user, err = u.FindById(uint(uintUserId))

	return user, err
}

func (u *userRedisRepository) FindByNickname(nickname string) (entity.User, error) {
	var user entity.User

	userId := u.redis.Get(context.Background(), fmt.Sprintf("users-nickname-%v", nickname)).Val()
	if userId == "" {
		return user, pkg.ErrUserNotFound
	}

	uintUserId, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		slog.Error(err.Error())
		return user, pkg.ErrUserNotFound
	}

	user, err = u.FindById(uint(uintUserId))

	return user, err
}

func (u *userRedisRepository) IsEmailExist(email string) bool {
	var count int64
	u.redis.Get(context.Background(), fmt.Sprintf("users-email-%v", email)).Scan(&count)
	return count > 0
}

func (u *userRedisRepository) IsNicknameExist(nickname string) bool {
	var count int64
	u.redis.Get(context.Background(), fmt.Sprintf("users-nickname-%v", nickname)).Scan(&count)
	return count > 0
}

type userDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(db *gorm.DB) UserRepositoryActions {
	if userDBRepositoryInstance == nil {
		userDBRepositoryInstance = &userDBRepository{
			db: db,
		}
	}

	return userDBRepositoryInstance
}

func (u *userDBRepository) Create(user entity.User) (entity.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (u *userDBRepository) Update(user entity.User) error {
	resp := u.db.Save(&user)
	if resp.RowsAffected == 0 {
		return pkg.ErrUserNotFound
	}
	err := resp.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return pkg.ErrUserNotFound
		}
		return err
	}

	return err
}

func (u *userDBRepository) Delete(userId uint) error {
	user := entity.User{
		Model: &gorm.Model{
			ID: userId,
		},
	}

	resp := u.db.Delete(&user)
	if resp.RowsAffected == 0 {
		return pkg.ErrUserNotFound
	}
	err := resp.Error
	if err != nil {
		return err
	}

	return err
}

func (u *userDBRepository) FindById(userId uint) (entity.User, error) {
	var user entity.User

	err := u.db.First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, pkg.ErrUserNotFound
		}

		return user, err
	}

	return user, err
}

func (u *userDBRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, pkg.ErrUserNotFound
		}

		return user, err
	}

	return user, err
}

func (u *userDBRepository) FindByNickname(nickname string) (entity.User, error) {
	var user entity.User

	err := u.db.Where("nickname = ?", nickname).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, pkg.ErrUserNotFound
		}

		return user, err
	}

	return user, err
}

func (u *userDBRepository) IsEmailExist(email string) bool {
	var count int64
	u.db.Model(&entity.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (u *userDBRepository) IsNicknameExist(nickname string) bool {
	var count int64
	u.db.Model(&entity.User{}).Where("nickname = ?", nickname).Count(&count)
	return count > 0
}
