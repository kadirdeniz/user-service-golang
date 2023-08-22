package handler

import (
	"fmt"
	"strconv"
	"user-service-golang/internal/dto"
	"user-service-golang/internal/service"
	"user-service-golang/pkg"
	"user-service-golang/tools"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

var userHandlerInstance UserHandlerActions

type UserHandlerActions interface {
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindByEmail(c *fiber.Ctx) error
	FindByNickname(c *fiber.Ctx) error
}

type userHandler struct {
	userService service.UserServiceActions
	validator   *validator.Validate
}

func NewUserHandler(
	userService service.UserServiceActions,
) AuthHandlerActions {
	if userHandlerInstance == nil {
		userHandlerInstance = &userHandler{
			userService: userService,
			validator:   validator.New(),
		}

		user := tools.NewServer().Group(fmt.Sprintf("/api/%v/users", pkg.Configs.Application.Version))
		user.Get("/:id", userHandlerInstance.FindById)
		user.Get("/email/:email", userHandlerInstance.FindByEmail)
		user.Get("/nickname/:nickname", userHandlerInstance.FindByNickname)
		user.Put("/:id", userHandlerInstance.Update)
		user.Delete("/:id", userHandlerInstance.Delete)
	}

	return authHandlerInstance
}

func (u *userHandler) Update(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	var updateUserRequest dto.UpdateUserRequest
	err = c.BodyParser(&updateUserRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	err = u.validator.Struct(updateUserRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	user, err := u.userService.FindById(uint(userId))
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.ErrorResponse{
				Message: pkg.ErrUserNotFound.Error(),
			})
		}
	}

	user.Firstname = updateUserRequest.Firstname
	user.Lastname = updateUserRequest.Lastname
	user.Nickname = updateUserRequest.Nickname
	user.Email = updateUserRequest.Email

	err = u.userService.Update(user)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.ErrorResponse{
				Message: pkg.ErrUserNotFound.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (u *userHandler) Delete(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	err = u.userService.Delete(uint(userId))
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.ErrorResponse{
				Message: pkg.ErrUserNotFound.Error(),
			})
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (u *userHandler) FindById(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	user, err := u.userService.FindById(uint(userId))
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.ErrorResponse{
				Message: pkg.ErrUserNotFound.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	var userResponse dto.GetUserByIdResponse
	err = mapstructure.Decode(user, &userResponse)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func (u *userHandler) FindByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := u.userService.FindByEmail(email)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.ErrorResponse{
				Message: pkg.ErrUserNotFound.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	var userResponse dto.GetUserByIdResponse
	err = mapstructure.Decode(user, &userResponse)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(userResponse)
}

func (u *userHandler) FindByNickname(c *fiber.Ctx) error {
	nickname := c.Params("nickname")

	user, err := u.userService.FindByNickname(nickname)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.ErrorResponse{
				Message: pkg.ErrUserNotFound.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	var userResponse dto.GetUserByIdResponse
	err = mapstructure.Decode(user, &userResponse)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(userResponse)
}
