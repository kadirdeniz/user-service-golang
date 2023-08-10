package handler

import (
	"user-service-golang/internal/dto"
	"user-service-golang/internal/service"
	"user-service-golang/pkg"
	"user-service-golang/tools"

	"github.com/gofiber/fiber/v2"
)

var authHandlerInstance AuthHandlerActions

type AuthHandlerActions interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	userService service.UserServiceActions
	jwtService  service.JWTActions
	bcrypt      service.BcryptActions
}

func NewAuthHandler(
	userService service.UserServiceActions,
	jwtService service.JWTActions,
	bcrypt service.BcryptActions,
) AuthHandlerActions {
	if authHandlerInstance == nil {
		authHandlerInstance = &authHandler{
			userService: userService,
			jwtService:  jwtService,
			bcrypt:      bcrypt,
		}

		auth := tools.NewServer().Group("/auth")
		auth.Post("/register", authHandlerInstance.Register)
		auth.Post("/login", authHandlerInstance.Login)
	}

	return authHandlerInstance
}

func (a *authHandler) Register(c *fiber.Ctx) error {
	var dto dto.RegisterRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	user, err := a.userService.Create(dto.ToUser())
	if err != nil {
		return err
	}

	token := a.jwtService.SetUserId(user.ID).CreateToken().GetToken()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func (a *authHandler) Login(c *fiber.Ctx) error {
	var dto dto.LoginRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	user, err := a.userService.FindByEmail(dto.Email)
	if err != nil {
		return err
	}

	if !a.bcrypt.IsPasswordCorrect(user.Password, dto.Password) {
		return pkg.ErrPasswordIncorrect
	}

	token := a.jwtService.SetUserId(user.ID).CreateToken().GetToken()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
