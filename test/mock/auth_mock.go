package mock

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"user-service-golang/internal/dto"
)

var MockRegisterRequest dto.RegisterRequest = dto.RegisterRequest{
	FirstName: "John",
	LastName:  "Doe",
	Email:     "johndoe@mail.com",
	Nickname:  "johndoe",
	Password:  "johndoe",
}

var MockLoginRequest dto.LoginRequest = dto.LoginRequest{
	Email:    "johndoe@mail.com",
	Password: "johndoe",
}

// var MockRegisterResponse dto.RegisterResponse = dto.RegisterResponse{
// 	Token: "string",
// }

func NewMockRegisterRequest(reqObj dto.RegisterRequest) *http.Request {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqObj)
	if err != nil {
		log.Fatal(err)
	}

	req := httptest.NewRequest("POST", "http://localhost:3000/auth/register", &buf)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func NewMockLoginRequest(reqObj dto.LoginRequest) *http.Request {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqObj)
	if err != nil {
		log.Fatal(err)
	}

	req := httptest.NewRequest("POST", "http://localhost:3000/auth/login", &buf)
	req.Header.Set("Content-Type", "application/json")

	return req
}
