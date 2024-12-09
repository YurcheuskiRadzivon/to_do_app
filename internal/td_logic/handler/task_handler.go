package handler

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/controller"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
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

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	task, err := th.controller.GetTask(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Render("tasks_id", task)
}
func (th *taskHandler) InsertTask(c *fiber.Ctx) error {
	var task model.Task
	var t model.T

	if err := c.BodyParser(&t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	task.Title = t.Title
	task.Description = t.Description
	task.Status = false
	if t.Status == "true" {
		task.Status = true
	}
	tokenStr := c.Cookies("tokenAuth")
	if err := th.controller.InsertTask(c.Context(), task, tokenStr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s", err)})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully"})

}
func (th *taskHandler) UpdateTask(c *fiber.Ctx) error {
	th.lgr.InfoLogger.Println("Update task with ID: ", c.Params("id"))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var task model.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	tokenStr := c.Cookies("tokenAuth")
	err = th.controller.UpdateTask(c.Context(), task, tokenStr, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully"})
}
func (th *taskHandler) DeleteTask(c *fiber.Ctx) error {

	th.lgr.InfoLogger.Println("Update task with ID: ", c.Params("id"))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = th.controller.DeleteTask(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully"})
}
