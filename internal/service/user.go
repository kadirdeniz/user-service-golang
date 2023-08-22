package service

import (
	"user-service-golang/internal/entity"
	"user-service-golang/internal/repository"
	"user-service-golang/pkg"
)

type UserServiceActions interface {
	Create(entity.User) (entity.User, error)
	Update(entity.User) error
	Delete(userId uint) error
	FindById(userId uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByNickname(email string) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepositoryActions
	bcrypt         BcryptActions
}

func NewUserService(
	userRepository repository.UserRepositoryActions,
) UserServiceActions {
	return &userService{
		userRepository: userRepository,
		bcrypt:         NewBcrypt(),
	}
}

func (u *userService) Create(user entity.User) (entity.User, error) {
	if u.userRepository.IsEmailExist(user.Email) {
		return entity.User{}, pkg.ErrEmailAlreadyExist
	}

	if u.userRepository.IsNicknameExist(user.Nickname) {
		return entity.User{}, pkg.ErrNicknameAlreadyExist
	}

	user.Password = u.bcrypt.HashPassword(user.Password)

	return u.userRepository.Create(user)
}

func (u *userService) Update(user entity.User) error {

	if u.userRepository.IsEmailExist(user.Email) {
		return pkg.ErrEmailAlreadyExist
	}

	if u.userRepository.IsNicknameExist(user.Nickname) {
		return pkg.ErrNicknameAlreadyExist
	}

	return u.userRepository.Update(user)
}

func (u *userService) Delete(userId uint) error {
	return u.userRepository.Delete(userId)
}

func (u *userService) FindById(userId uint) (entity.User, error) {
	return u.userRepository.FindById(userId)
}

func (u *userService) FindByEmail(email string) (entity.User, error) {
	return u.userRepository.FindByEmail(email)
}

func (u *userService) FindByNickname(nickname string) (entity.User, error) {
	return u.userRepository.FindByNickname(nickname)
}
