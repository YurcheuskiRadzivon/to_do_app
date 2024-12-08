package handler

import (
	"context"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/controller"
	"github.com/gofiber/fiber/v2"
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
	lgr        *logger.Logger
}

func NewTaskHandler(controller controller.TaskController, lgr *logger.Logger) TaskHandler {
	return &taskHandler{
		controller: controller,
		ctx:        context.Background(),
		lgr:        lgr,
	}

}
func (th *taskHandler) GetTasks(c *fiber.Ctx) error {
	cookie := c.Cookies("tokenAuth")

	sortParam := c.Query("sort", "date")

	tasks, err := th.controller.GetTasks(c.Context(), cookie, sortParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
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
