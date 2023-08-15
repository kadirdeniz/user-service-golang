package service

import (
	"user-service-golang/internal/entity"
	"user-service-golang/pkg"
)

var authServiceInstance *authService

type authService struct {
	userService UserServiceActions
	jwtService  JWTActions
	bcrypt      BcryptActions
}

type AuthServiceActions interface {
	Register(user entity.User) (string, error)
	Login(user entity.User) (string, error)
}

func NewAuthService(userService UserServiceActions) AuthServiceActions {
	if authServiceInstance == nil {
		authServiceInstance = &authService{
			userService: userService,
			jwtService:  NewJWT(),
			bcrypt:      NewBcrypt(),
		}
	}
	return authServiceInstance
}

func (a *authService) Register(user entity.User) (string, error) {
	userObj, err := a.userService.Create(user)
	if err != nil {
		return "", err
	}
	return a.jwtService.SetUserId(userObj.ID).CreateToken().GetToken(), nil
}

func (a *authService) Login(user entity.User) (string, error) {
	userObj, err := a.userService.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if userObj.ID == 0 {
		return "", pkg.ErrUserNotFound
	}

	if !a.bcrypt.IsPasswordCorrect(userObj.Password, user.Password) {
		return "", pkg.ErrPasswordIncorrect
	}
	return a.jwtService.SetUserId(userObj.ID).CreateToken().GetToken(), nil
}
