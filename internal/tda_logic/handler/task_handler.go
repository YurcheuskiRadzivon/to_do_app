package handler

import (
	"context"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/controller"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
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
	image1Base64 := "iVBORw0KGgoAAAANSUhEUgAA...AAAAABJRU5ErkJggg=="
	image2Base64 := "iVBORw0KGgoAAAANSUhEUgAA...AAAAAElFTkSuQmCC"
	task := model.Task{
		Title:       "Sample Task",
		Description: "This is a sample task",
		Status:      true,
		AddedTime:   time.Now(),
		Images:      []string{image1Base64, image2Base64},
	}
	cookie := c.Cookies("tokenAuth")
	err := th.controller.InsertTask(c.Context(), task, cookie)
	if err != nil {
		log.Println(err)
		return err
	}
	tasks, err := th.controller.GetTasks(c.Context(), 9, cookie)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(tasks)
	return nil
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
