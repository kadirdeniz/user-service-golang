package pkg

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var ErrEmailAlreadyExist = errors.New("email already exist")
var ErrNicknameAlreadyExist = errors.New("nickname already exist")
var ErrPasswordIncorrect = errors.New("password incorrect")
