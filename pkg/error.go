package pkg

import "errors"

var ErrEmailAlreadyExist = errors.New("email already exist")
var ErrNicknameAlreadyExist = errors.New("nickname already exist")
var ErrPasswordIncorrect = errors.New("password incorrect")
