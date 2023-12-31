package pkg

import "errors"

// swagger:model ErrorResponse
type ErrorResponse struct {
	// The error message.
	// example: email already exist
	// in: body
	Message string `json:"message"`
}

var ErrEmailAlreadyExist = errors.New("email already exist")
var ErrNicknameAlreadyExist = errors.New("nickname already exist")
var ErrPasswordIncorrect = errors.New("password incorrect")
var ErrUserNotFound = errors.New("user not found")
