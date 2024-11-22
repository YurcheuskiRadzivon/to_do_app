package controller

import (
	"context"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/model"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/repository"
)

type TaskController interface {
	GetTasks(ctx context.Context, userId int) ([]model.Task, error)
	GetTask(ctx context.Context, id int) (*model.Task, error)
	InsertTask(ctx context.Context, Task model.Task) error
	UpdateTask(ctx context.Context, id int, Task model.Task) error
	DeleteTask(ctx context.Context, id int) error
}
type taskController struct {
	repo repository.UserRepository
}
