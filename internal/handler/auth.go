package handler

import (
	"user-service-golang/internal/dto"
	"user-service-golang/internal/service"
	"user-service-golang/pkg"
	"user-service-golang/tools"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

var authHandlerInstance AuthHandlerActions

type AuthHandlerActions interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthServiceActions
	validator   *validator.Validate
}

func NewAuthHandler(
	authService service.AuthServiceActions,
) AuthHandlerActions {
	if authHandlerInstance == nil {
		authHandlerInstance = &authHandler{
			authService: authService,
			validator:   validator.New(),
		}

		auth := tools.NewServer().Group("/auth")
		auth.Post("/register", authHandlerInstance.Register)
		auth.Post("/login", authHandlerInstance.Login)
	}

	return authHandlerInstance
}

func (a *authHandler) Register(c *fiber.Ctx) error {
	var dtoObj dto.RegisterRequest
	if err := c.BodyParser(&dtoObj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err := a.validator.Struct(dtoObj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	token, err := a.authService.Register(dtoObj.ToUser())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.TokenResponse{
		Token: token,
	})
}

func (a *authHandler) Login(c *fiber.Ctx) error {
	var dtoObj dto.LoginRequest
	if err := c.BodyParser(&dtoObj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	if err := a.validator.Struct(dtoObj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	token, err := a.authService.Login(dtoObj.ToUser())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.TokenResponse{
		Token: token,
	})
}
