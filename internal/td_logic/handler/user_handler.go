package handler

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/controller"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUser(c *fiber.Ctx) error
	InsertUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error

	LoginUser(c *fiber.Ctx) error
}
type userHandler struct {
	ctx        context.Context
	controller controller.UserController
	lgr        *logger.Logger
}

func NewUserHandler(controller controller.UserController, lgr *logger.Logger) UserHandler {
	return &userHandler{
		controller: controller,
		ctx:        context.Background(),
	}

}
func (us *userHandler) GetUser(c *fiber.Ctx) error {
	var user *model.User
	tokenStr := c.Cookies("tokenAuth")
	user, err := us.controller.GetUser(c.Context(), tokenStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s", err)})
	}
	return c.Render("user", user)
}
func (us *userHandler) InsertUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if err := us.controller.InsertUser(c.Context(), user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully"})
}
func (us *userHandler) UpdateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	tokenStr := c.Cookies("tokenAuth")
	tokenStrNew, err := us.controller.UpdateUser(c.Context(), user, tokenStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot update user"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": fmt.Sprintf("%s", tokenStrNew)})
}
func (us *userHandler) DeleteUser(c *fiber.Ctx) error {
	tokenStr := c.Cookies("tokenAuth")
	if err := us.controller.DeleteUser(c.Context(), tokenStr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s", err)})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully"})
}

func (us *userHandler) LoginUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	jwtStr, err := us.controller.LoginUser(c.Context(), &user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": fmt.Sprintf("%s", jwtStr)})

}
