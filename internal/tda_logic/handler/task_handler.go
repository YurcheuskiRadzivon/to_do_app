package handler

import (
	"context"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/controller"
	"github.com/gofiber/fiber/v2"
	"log"
)

type TaskHandler interface {
	GetTasks(c *fiber.Ctx) error
	GetTask(c *fiber.Ctx) error
	InsertTask(c *fiber.Ctx) error
	UpdateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}
type taskHandler struct {
	ctx        context.Context
	controller controller.TaskController
}

func NewTaskHandler(controller controller.TaskController) TaskHandler {
	return &taskHandler{
		controller: controller,
		ctx:        context.Background(),
	}

}
func (th *taskHandler) GetTasks(c *fiber.Ctx) error {
	cookie := c.Cookies("tokenAuth")
	tasks, err := th.controller.GetTasks(c.Context(), cookie)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.Render("tasks", tasks)
}
func (th *taskHandler) GetTask(c *fiber.Ctx) error {
	return nil
}
func (th *taskHandler) InsertTask(c *fiber.Ctx) error {
	return nil
}
func (th *taskHandler) UpdateTask(c *fiber.Ctx) error {
	return nil
}
func (th *taskHandler) DeleteTask(c *fiber.Ctx) error {
	return nil
}
