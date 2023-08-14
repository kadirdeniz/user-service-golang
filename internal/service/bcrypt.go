package service

import (
	hasher "golang.org/x/crypto/bcrypt"
)

var bcryptInstance BcryptActions

type BcryptActions interface {
	HashPassword(password string) string
	IsPasswordCorrect(hashedPassword, password string) bool
}

type bcrypt struct{}

func NewBcrypt() BcryptActions {
	if bcryptInstance == nil {
		bcryptInstance = &bcrypt{}
	}
	return bcryptInstance
}

func (b *bcrypt) HashPassword(password string) string {
	hashedPassword, err := hasher.GenerateFromPassword([]byte(password), hasher.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (b *bcrypt) IsPasswordCorrect(hashedPassword, password string) bool {
	err := hasher.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
