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
	authService service.AuthServiceActions
}

func NewAuthHandler(
	authService service.AuthServiceActions,
) AuthHandlerActions {
	if authHandlerInstance == nil {
		authHandlerInstance = &authHandler{
			authService: authService,
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

	token, err := a.authService.Register(dto.ToUser())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func (a *authHandler) Login(c *fiber.Ctx) error {
	var dto dto.LoginRequest
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	token, err := a.authService.Login(dto.ToUser())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
