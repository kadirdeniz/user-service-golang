package service

import (
	"fmt"
	"strconv"
	"user-service-golang/pkg"

	"github.com/golang-jwt/jwt/v5"
)

var secret = pkg.Configs.JWT.Secret

type JWTActions interface {
	CreateToken() JWTActions
	ParseToken() JWTActions
	SetToken(token string) JWTActions
	GetToken() string
	GetUserId() uint
	SetUserId(userId uint) JWTActions
}

type jsonwebtoken struct {
	token  string
	claims jwt.RegisteredClaims
}

func NewJWT() JWTActions {
	return &jsonwebtoken{
		token:  "",
		claims: jwt.RegisteredClaims{},
	}
}

func (j *jsonwebtoken) SetToken(token string) JWTActions {
	j.token = token
	return j
}

func (j *jsonwebtoken) SetUserId(userId uint) JWTActions {
	j.claims.Issuer = strconv.FormatUint(uint64(userId), 10)
	return j
}

func (j *jsonwebtoken) GetUserId() uint {
	userId, err := strconv.ParseUint(j.claims.Subject, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	return uint(userId)
}

func (j *jsonwebtoken) GetToken() string {
	return j.token
}

func (j *jsonwebtoken) CreateToken() JWTActions {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims).SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	j.token = token

	return j
}

func (j *jsonwebtoken) ParseToken() JWTActions {
	token, err := jwt.ParseWithClaims(j.token, &j.claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		fmt.Println(claims.Issuer)
	}

	return j
}
