package controller

import (
	"github.com/YogaRP/finansial/user-service/internal/pkg/response"
	"github.com/YogaRP/finansial/user-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
)

type UserControllerInterface interface {
	GetUserByEmail(c fiber.Ctx) error
	GetUserByID(c fiber.Ctx) error
}

type UserController struct {
	userService service.UserServiceInterface
}

// GetUserByEmail implements [UserControllerInterface].
func (u *UserController) GetUserByEmail(c fiber.Ctx) error {
	ctx := c.Context()
	email := c.Params("email")

	user, err := u.userService.GetUserByEmail(ctx, email)

	if err != nil {
		log.Errorf("[UserController] GetUserByEmail - 1: %v", err)
		return response.InternalError(c, err.Error())
	}

	return response.OK(c, "Success get user", user)
}

// GetUserByID implements [UserControllerInterface].
func (u *UserController) GetUserByID(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	parseId, _ := uuid.Parse(id)

	user, err := u.userService.GetUserByID(ctx, parseId)
	if err != nil {
		log.Errorf("[UserController] GetUserByID - 1: %v", err)
		return response.InternalError(c, err.Error())
	}

	return response.OK(c, "Success get user", user)
}

func NewUserController(userService service.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userService,
	}
}
